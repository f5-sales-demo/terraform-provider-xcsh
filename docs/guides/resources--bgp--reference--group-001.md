---
page_title: "xcsh_bgp reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bgp reference."
---

# xcsh_bgp reference

<a id="canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- Property reference

<a id="canonical-3222101200301320-0210220000301101-3232032102211011-1022001220210323-1103322320313023-3102030201233320-1220223022200231-1033122200111020"></a>

### Direct properties for `xcsh_bgp`

<a id="canonical-3222311011332211-3022322101110330-1030112203131321-0032132330133203-0132003002301211-0001021301311212-3333223121133131-1301301313220032"></a>

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

- [bgp_parameters](resources--bgp--reference--group-001.md#canonical-3103003320212223-3323032330310220-0112102013012011-3303212132102001-2201310010330000-1120110231210311-2311000111033323-0133120330101211): complete subsection reference.

<a id="canonical-2112310113113222-3232210130222032-2201220231022102-3121211212303311-1330033102113302-2010202202130011-0112313011113003-2112011101002333"></a>

<a id="canonical-0111320003003302-0303031033131323-2132001332012310-2112120310103113-3120012200210200-0001022103022000-1103322031032101-2003001023213203"></a>

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

<a id="canonical-1011330323020301-1213000102331313-0231311131203010-1230222233012011-1311133111002213-0321120031103103-1110220033231332-0221131101001212"></a>

<a id="canonical-0312032121302302-3301001121301203-1213113003303110-3122110131023213-2123110121013110-1100101131110213-3212313301002001-3321233300132232"></a>

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

<a id="canonical-1122102222113110-2201020032002302-3321310313120312-2200010232213130-3223011233321201-3220022203112313-1331011102333133-2213032210211313"></a>

<a id="canonical-1123203322103222-1330201203020203-2201021133013001-2231101102322223-2033021321123022-2001210301330323-0321311223331310-0120201302122230"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3223123231330222-0022030031123013-0303012210112030-3310103313313302-0102001032011013-0011213233322323-3010310203121012-2011130122321010"></a>

<a id="canonical-3331123102200201-1101220122323003-1303101323000233-1312212010011201-3211321211232232-0021131323031203-0031322300323322-2230200010311310"></a>

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

<a id="canonical-2323131130202302-2012123233210020-1000100221033233-2001223133021202-0301232333123202-1220000103131303-3133220001100110-1032130112313213"></a>

<a id="canonical-3202313332033200-1112310032331103-0132020222111312-2113203121122322-3120031122002301-3221313333110303-0313030310223321-3103302130131310"></a>

#### `name` property

Type: `"string"`. Required.

Name of the BGP. Must be unique within the namespace.

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

<a id="canonical-3020322032120002-1032113103002032-0021133030301201-0102132023032300-3223031000100003-1032322212203202-1110211120032013-1313003311003321"></a>

<a id="canonical-1300030101333211-0210330320223200-0110003120012213-2302133132120133-2101100130123230-3203121123020030-3100131331203233-0103232011002023"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the BGP is created.

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

- [peers](resources--bgp--reference--group-001.md#canonical-3132002313303222-3210021222211033-2013001002330232-0123121112132233-0321012202000010-0112130200010312-2201100311111101-1320000132020103): complete subsection reference.

- [timeouts](resources--bgp--reference--group-002.md#canonical-3300212323303233-3123033100000031-0113223222020021-3122201130331211-0030003011113133-0300212300011301-1323133010230310-2232202103011033): complete subsection reference.

- [where](resources--bgp--reference--group-002.md#canonical-3310113311010321-2211233202102322-3001232112032331-0322032021203310-0031203331133113-2121220003210231-0303000101131332-1233013013233231): complete subsection reference.

<a id="canonical-0102103220133131-0203201132023111-0011033331211231-2233310031020313-3110003122000311-0031120010123003-2122311322010223-0032113212021301"></a>

### All schema paths for `xcsh_bgp`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--bgp--reference--group-001.md#canonical-3222311011332211-3022322101110330-1030112203131321-0032132330133203-0132003002301211-0001021301311212-3333223121133131-1301301313220032) |
| `bgp_parameters` | [bgp_parameters](resources--bgp--reference--group-001.md#canonical-0312322210330131-3232021030003310-2001103001223211-3110312220033113-1222332202113310-2100033021033230-0311030032001133-3311301130030123) |
| `bgp_parameters.asn` | [bgp_parameters.asn](resources--bgp--reference--group-001.md#canonical-2210011010202021-1112233002011032-1303133231110233-0013021320121213-0223023202010230-1123002003112120-2331202131231130-2222102331332133) |
| `bgp_parameters.from_site` | [bgp_parameters.from_site](resources--bgp--reference--group-001.md#canonical-2113202301310011-2311002103002231-3323001321333201-2301102302033103-2213002031211231-0222000313311120-0111003022132023-2210301231003333) |
| `bgp_parameters.ip_address` | [bgp_parameters.ip_address](resources--bgp--reference--group-001.md#canonical-3322221000221021-1111122120131301-3320013312220323-0231133002232121-0132112313311221-0131112220032300-3002211122122220-0102320211321111) |
| `bgp_parameters.local_address` | [bgp_parameters.local_address](resources--bgp--reference--group-001.md#canonical-0302311312133023-1320022223303223-0012111201033211-0300312232012212-1322321031202032-1123210332312013-0333130031213103-1030322322200032) |
| `description` | [description](resources--bgp--reference--group-001.md#canonical-2112310113113222-3232210130222032-2201220231022102-3121211212303311-1330033102113302-2010202202130011-0112313011113003-2112011101002333) |
| `disable` | [disable](resources--bgp--reference--group-001.md#canonical-1011330323020301-1213000102331313-0231311131203010-1230222233012011-1311133111002213-0321120031103103-1110220033231332-0221131101001212) |
| `id` | [ID](resources--bgp--reference--group-001.md#canonical-1122102222113110-2201020032002302-3321310313120312-2200010232213130-3223011233321201-3220022203112313-1331011102333133-2213032210211313) |
| `labels` | [labels](resources--bgp--reference--group-001.md#canonical-3223123231330222-0022030031123013-0303012210112030-3310103313313302-0102001032011013-0011213233322323-3010310203121012-2011130122321010) |
| `name` | [name](resources--bgp--reference--group-001.md#canonical-2323131130202302-2012123233210020-1000100221033233-2001223133021202-0301232333123202-1220000103131303-3133220001100110-1032130112313213) |
| `namespace` | [namespace](resources--bgp--reference--group-001.md#canonical-3020322032120002-1032113103002032-0021133030301201-0102132023032300-3223031000100003-1032322212203202-1110211120032013-1313003311003321) |
| `peers` | [peers](resources--bgp--reference--group-001.md#canonical-2331002010131000-1222130010211330-2332203203301122-0311312231230211-3210123131213021-2203101011320030-0133100031132323-0011022011223210) |
| `peers.bfd_disabled` | [peers.bfd_disabled](resources--bgp--reference--group-001.md#canonical-1320021021213230-0121013103231032-3033012132101202-2112231330103033-0021310211233212-3210000031130311-0120002301130103-0332033102333331) |
| `peers.bfd_enabled` | [peers.bfd_enabled](resources--bgp--reference--group-001.md#canonical-0022030032232021-3230222003321222-2002123211130030-1103322302013120-3000013330311213-2322123312101021-1333003022313121-2013322121212333) |
| `peers.bfd_enabled.multiplier` | [peers.bfd_enabled.multiplier](resources--bgp--reference--group-001.md#canonical-0121212102201310-0200133211230321-2023003003003212-1100103022233133-1122122121301203-0111121120222202-2122130110232002-0203030011222012) |
| `peers.bfd_enabled.receive_interval_milliseconds` | [peers.bfd_enabled.receive_interval_milliseconds](resources--bgp--reference--group-001.md#canonical-0211302112202210-2122120231123031-0312221013232220-0201013011313202-0022132023332000-3223001131012111-3001120111020132-0320022310020212) |
| `peers.bfd_enabled.transmit_interval_milliseconds` | [peers.bfd_enabled.transmit_interval_milliseconds](resources--bgp--reference--group-001.md#canonical-0010211011000011-1002310030223330-3331122300210010-0221023210223212-1312012123133102-0131332012010002-0323303132333212-0020122111112212) |
| `peers.disable_spec` | [peers.disable_spec](resources--bgp--reference--group-001.md#canonical-1321331310123103-1122023202021200-1100013311120021-0120332022331121-3033020023331123-2103311132221023-2122132011011200-0210130230320301) |
| `peers.ebgp_multihop_disabled` | [peers.ebgp_multihop_disabled](resources--bgp--reference--group-001.md#canonical-0331333310011323-1302320302330230-3313100321111320-1230102121201023-0300213323323301-2122022303003022-1012312301133212-1102311103310211) |
| `peers.ebgp_multihop_enabled` | [peers.ebgp_multihop_enabled](resources--bgp--reference--group-001.md#canonical-1310331200012212-1131011331210331-0030132003311033-3001203111203210-3300223203130132-0123030010010132-1020333120320230-3322201312100210) |
| `peers.external` | [peers.external](resources--bgp--reference--group-001.md#canonical-2000211200300200-3333300230331212-2112331202300310-1133112322301320-3130111020320021-1002232011133032-1010302030313013-3302132222003210) |
| `peers.external.address` | [peers.external.address](resources--bgp--reference--group-001.md#canonical-2132323300330013-3002121001320131-0300233300100011-1233212201222132-3330032033121301-3023300311011211-1301003331000110-1031110030112211) |
| `peers.external.address_ipv6` | [peers.external.address_ipv6](resources--bgp--reference--group-001.md#canonical-2302302231300223-1201123330212031-1103020203002032-1230030032013322-2102102301100102-2002221001302031-3011201132122330-2333003020132223) |
| `peers.external.asn` | [peers.external.asn](resources--bgp--reference--group-001.md#canonical-1213230011013003-2212331323131112-0233011230222300-3103013031013010-3103131011112202-0211302100323233-3030310131000101-1000121100311102) |
| `peers.external.default_gateway` | [peers.external.default_gateway](resources--bgp--reference--group-001.md#canonical-1113031133210210-0123331232112132-1302223112210321-3302013122032132-2231203100002312-0033210020321232-0221232301002113-2022000031331230) |
| `peers.external.default_gateway_v6` | [peers.external.default_gateway_v6](resources--bgp--reference--group-001.md#canonical-0013232310220330-0132311233232121-3330232023211331-3111312201102330-3131103202200232-2313123023233310-1220111310303210-1302110202222020) |
| `peers.external.disable_spec` | [peers.external.disable_spec](resources--bgp--reference--group-001.md#canonical-0120123011321323-0101222322033033-1131313130100232-1010100030022212-0122133230131210-2211123330302031-0001211010121232-0101111313112110) |
| `peers.external.disable_v6` | [peers.external.disable_v6](resources--bgp--reference--group-001.md#canonical-1302000010212312-0021120012002013-3201230003031012-1021103202020302-3100001102311320-3020220113302100-1323213011033320-2322231101032133) |
| `peers.external.external_connector` | [peers.external.external_connector](resources--bgp--reference--group-001.md#canonical-0321330001221023-1302303330111333-1322021211300331-3031310123213303-1133002000210012-0033321300323131-0302320232210311-0001211022010200) |
| `peers.external.family_inet` | [peers.external.family_inet](resources--bgp--reference--group-001.md#canonical-1312313031110203-2100113022323312-1333112102111220-2132312110103010-3030203222222010-0010013201023102-3133211121110102-2302332311103303) |
| `peers.external.family_inet.disable_spec` | [peers.external.family_inet.disable_spec](resources--bgp--reference--group-001.md#canonical-1312122231020110-1022002103230303-2010313212321131-2210122303110321-0233012110201331-1322322122010200-0312011122333230-0311000100212302) |
| `peers.external.family_inet.enable` | [peers.external.family_inet.enable](resources--bgp--reference--group-001.md#canonical-3033123202300311-1023032022332321-3320113212223320-2110210100013012-3331201130332003-3121222313003032-0131300331230000-2023300121011103) |
| `peers.external.family_inet.enable.aggregation` | [peers.external.family_inet.enable.aggregation](resources--bgp--reference--group-001.md#canonical-2221232100031210-3231201200300331-0022303303101131-2210032112033201-3002313010023212-1211203123320221-2311103212003032-0122311202031100) |
| `peers.external.family_inet.enable.aggregation.ip_prefix` | [peers.external.family_inet.enable.aggregation.ip_prefix](resources--bgp--reference--group-001.md#canonical-0203022330221313-1000333012130200-1303232123121100-1011002321001031-1022113321033223-3322313330200302-0003000331323333-2212303201300331) |
| `peers.external.family_inet.enable.aggregation.options` | [peers.external.family_inet.enable.aggregation.options](resources--bgp--reference--group-001.md#canonical-2322030231220312-3322333223221211-2030222101221100-1101320102130032-3121310200032313-2223201110311013-3030212111233111-3201221032113320) |
| `peers.external.family_inet.enable.aggregation.options.summary_only` | [peers.external.family_inet.enable.aggregation.options.summary_only](resources--bgp--reference--group-001.md#canonical-2323111200023033-2110321033101002-2000320232010310-2000032100333300-1212200301132133-1311310212022100-1332113323022120-2323011313220003) |
| `peers.external.from_site` | [peers.external.from_site](resources--bgp--reference--group-001.md#canonical-3332010223212103-0030123011103022-0110022131033110-3213201312113132-3123232310332122-3023232331021102-2020010333012100-2122120031003233) |
| `peers.external.from_site_v6` | [peers.external.from_site_v6](resources--bgp--reference--group-001.md#canonical-3301000212131021-0020301130303122-3333002302233022-0223001012232123-1023112200020000-0020300201113002-2001020201010301-2021230011300113) |
| `peers.external.interface` | [peers.external.interface](resources--bgp--reference--group-001.md#canonical-2222321312303320-1001112032003002-3203131103020213-3130222312330210-3000220201012220-1103232130101301-3130312121110220-0212100312010022) |
| `peers.external.interface.name` | [peers.external.interface.name](resources--bgp--reference--group-001.md#canonical-3010231102122333-1012030103021130-0212313232322312-1330232221233101-3112221211330030-2033011001002133-2321103011022003-2123100032130211) |
| `peers.external.interface.namespace` | [peers.external.interface.namespace](resources--bgp--reference--group-001.md#canonical-3033232232020333-0212212032023120-0333311313113002-2033131102213332-3231030102321210-2322221012301221-1303323313000302-0013300111112301) |
| `peers.external.interface.tenant` | [peers.external.interface.tenant](resources--bgp--reference--group-001.md#canonical-0012100303331212-0003022311002012-0000300000302100-1003330001212120-0201003300301321-2223001313331322-3230203002120030-0300331230121101) |
| `peers.external.interface_list` | [peers.external.interface_list](resources--bgp--reference--group-002.md#canonical-3010023100000333-0132232020110110-1313203002321022-1002210013020010-2301332303231233-2023100121232222-2332110202121120-3101221133121211) |
| `peers.external.interface_list.interfaces` | [peers.external.interface_list.interfaces](resources--bgp--reference--group-002.md#canonical-1332320012301200-3311220113333022-3230210010301011-0113313313212133-1313012332203001-2022031331100230-1201112110330332-2032022133320331) |
| `peers.external.interface_list.interfaces.name` | [peers.external.interface_list.interfaces.name](resources--bgp--reference--group-002.md#canonical-2001231001102201-3302101131131132-1202112121023233-3332302003000213-1030320013031031-1100212033110111-3201322201301123-0313110121312331) |
| `peers.external.interface_list.interfaces.namespace` | [peers.external.interface_list.interfaces.namespace](resources--bgp--reference--group-002.md#canonical-0131023001131032-1122133011320202-1020330113022131-0032223011311132-0001311302300220-3330001200232230-2332233300232103-3102000102323300) |
| `peers.external.interface_list.interfaces.tenant` | [peers.external.interface_list.interfaces.tenant](resources--bgp--reference--group-002.md#canonical-2311221131010231-3320301333331102-3232000020100122-3213101221132010-1000200311031023-2102002321310200-0012330203021221-1010132323311112) |
| `peers.external.md5_auth_key` | [peers.external.md5_auth_key](resources--bgp--reference--group-001.md#canonical-0022210203333322-0013131030000000-0232320203202012-1121122322103210-1131033101321103-3001233013300122-3323300003333030-0113111311312120) |
| `peers.external.no_authentication` | [peers.external.no_authentication](resources--bgp--reference--group-002.md#canonical-3331131031013000-0123322221321102-3001103211302301-1312030020112001-2130223211000220-2011002333003121-2100123022301231-1132312010002121) |
| `peers.external.port` | [peers.external.port](resources--bgp--reference--group-001.md#canonical-3312112130333030-2003101100100301-3211220300331030-2301112012103001-2131232002222201-1022303203233331-3313232320212011-3220021230023220) |
| `peers.external.subnet_begin_offset` | [peers.external.subnet_begin_offset](resources--bgp--reference--group-001.md#canonical-1131023001230033-2131320001131012-2222110133012021-1331021331111102-3230113122001020-2210022013322003-2313311030331002-1311033220102133) |
| `peers.external.subnet_begin_offset_v6` | [peers.external.subnet_begin_offset_v6](resources--bgp--reference--group-001.md#canonical-3012101020021013-3120021310310203-1012330213231222-2123122022323100-2210210022221221-1120013202131331-0102332132113100-2321131033013102) |
| `peers.external.subnet_end_offset` | [peers.external.subnet_end_offset](resources--bgp--reference--group-001.md#canonical-0020033300211101-2010103202023233-0103003201023010-0330031330313311-0301132100131101-1002021101033300-2130022333212020-1221021020201122) |
| `peers.external.subnet_end_offset_v6` | [peers.external.subnet_end_offset_v6](resources--bgp--reference--group-001.md#canonical-3032022233302322-2000123013313122-1320321331221101-1023303023321120-0001310113031331-0333133202133333-3221311001231001-0223323313103230) |
| `peers.label` | [peers.label](resources--bgp--reference--group-001.md#canonical-1120332223200211-1002002011120110-1110120212101301-0220311002123000-1100010231023013-3201122123302033-1220313323132102-1203311131132103) |
| `peers.metadata` | [peers.metadata](resources--bgp--reference--group-002.md#canonical-3123021121030231-1321121201110330-2231213223012321-2231100202001102-1230331233033310-0120131322131110-2120302011121132-1313110021113300) |
| `peers.metadata.description_spec` | [peers.metadata.description_spec](resources--bgp--reference--group-002.md#canonical-2321232310222311-2033202210113321-3211230220001103-1122211013312211-3220200100023320-2013111001002320-3301312011201130-2302103300101330) |
| `peers.metadata.name` | [peers.metadata.name](resources--bgp--reference--group-002.md#canonical-2201302213321233-1331331200311312-2212022110032301-2233330133133312-0121122030032230-0022112210120212-0100002020030330-2020020222033323) |
| `peers.passive_mode_disabled` | [peers.passive_mode_disabled](resources--bgp--reference--group-002.md#canonical-2212020233023212-2222232202312122-2020000113303321-2021001111300031-3313301122332212-1020122230002331-2032310231112310-0212231103210122) |
| `peers.passive_mode_enabled` | [peers.passive_mode_enabled](resources--bgp--reference--group-002.md#canonical-1211231311112210-2302200201102202-3222212121032320-0033302331200303-2310002333132111-2333330210220302-1220233203300200-0331131300012001) |
| `peers.routing_policies` | [peers.routing_policies](resources--bgp--reference--group-002.md#canonical-1222011310120302-1020320101232012-1012212323223221-3030013303323030-1100221312103133-3013110032213120-1222103032012133-2301000320111113) |
| `peers.routing_policies.route_policy` | [peers.routing_policies.route_policy](resources--bgp--reference--group-002.md#canonical-3022131123031203-1002112023220131-3200102201102031-1232113311200030-3211302303313123-2210130111233323-3303102011112021-0031313200022303) |
| `peers.routing_policies.route_policy.all_nodes` | [peers.routing_policies.route_policy.all_nodes](resources--bgp--reference--group-002.md#canonical-2102120230022023-0322332220121120-0303223021202313-3320210222313113-1233032133012203-1123133131002201-1032131113301020-1203010003212100) |
| `peers.routing_policies.route_policy.inbound` | [peers.routing_policies.route_policy.inbound](resources--bgp--reference--group-002.md#canonical-3232300212311103-0231023320112211-0201112100310020-2333233111102223-3012310210213101-1310221031332321-3100320211213013-1321223320020023) |
| `peers.routing_policies.route_policy.node_name` | [peers.routing_policies.route_policy.node_name](resources--bgp--reference--group-002.md#canonical-2100303231002121-2031220231011032-3202001321030203-1232023203231321-3332100302333032-3101231023133131-1030132221110023-0322132230213121) |
| `peers.routing_policies.route_policy.node_name.node` | [peers.routing_policies.route_policy.node_name.node](resources--bgp--reference--group-002.md#canonical-2012331002333101-1133322030310302-3110002120212201-1233332023121132-2001200233030321-3221301011212101-0211030013313000-0312130222232111) |
| `peers.routing_policies.route_policy.object_refs` | [peers.routing_policies.route_policy.object_refs](resources--bgp--reference--group-002.md#canonical-0231022312211000-1133122020231123-3012212311023001-3020023201200312-1301103120120133-2202222323223231-3311330120312211-1211101030131030) |
| `peers.routing_policies.route_policy.object_refs.kind` | [peers.routing_policies.route_policy.object_refs.kind](resources--bgp--reference--group-002.md#canonical-3303100000322320-0303320303313123-0213023320110113-1201332110022123-2023302122021133-1101301223022233-2112233031020201-3323333023222233) |
| `peers.routing_policies.route_policy.object_refs.name` | [peers.routing_policies.route_policy.object_refs.name](resources--bgp--reference--group-002.md#canonical-3213230102102032-2202232011323010-0031220011332033-0220030231033212-3233322133022122-0002301302303210-1220023330323330-2321230232021301) |
| `peers.routing_policies.route_policy.object_refs.namespace` | [peers.routing_policies.route_policy.object_refs.namespace](resources--bgp--reference--group-002.md#canonical-0202323012200310-1101123320331130-0011133130002310-3220120300202002-0313101032033320-1312122133113210-0021030030312023-0013100320223020) |
| `peers.routing_policies.route_policy.object_refs.tenant` | [peers.routing_policies.route_policy.object_refs.tenant](resources--bgp--reference--group-002.md#canonical-2200300120323213-0223001322322320-1011322031112311-2002021312000021-0023230311232222-2302032132111303-3033312120022331-2011222331013200) |
| `peers.routing_policies.route_policy.object_refs.uid` | [peers.routing_policies.route_policy.object_refs.uid](resources--bgp--reference--group-002.md#canonical-3112023030132101-0120002322112123-0133222113213330-1231120313200122-3200202020300312-1331112302230032-2313130001032101-2320011311301022) |
| `peers.routing_policies.route_policy.outbound` | [peers.routing_policies.route_policy.outbound](resources--bgp--reference--group-002.md#canonical-1132311230203222-1203220010222010-2021121231222300-0202223200102012-3202330200123302-0233023022112021-2102110211111202-0001212201102310) |
| `timeouts` | [timeouts](resources--bgp--reference--group-002.md#canonical-3021132332201311-2300111330103202-2031210111010222-2323023300232102-1121013002001221-2302111301321311-1300210212212221-1133120311021213) |
| `timeouts.create` | [timeouts.create](resources--bgp--reference--group-002.md#canonical-0001123213110033-1101132023002101-0301032133310021-3110122311333033-2002121333210203-2031310220331120-1002212000023023-0113020122302013) |
| `timeouts.delete` | [timeouts.delete](resources--bgp--reference--group-002.md#canonical-1303300133212212-3113003210211101-0213130023022033-2130101330100030-1131132212033020-3330223222213231-3031021332213321-0123130130333021) |
| `timeouts.read` | [timeouts.read](resources--bgp--reference--group-002.md#canonical-0213202220010231-3132220332213101-0032133111301032-3023113333022022-1311132202320330-2300023113101013-3322233133200220-0210211330323113) |
| `timeouts.update` | [timeouts.update](resources--bgp--reference--group-002.md#canonical-1211222311012113-0322013113212200-3032010131111123-1130310123120112-2332301123030032-2003033213012030-1031032222213220-3031300023320130) |
| `where` | [where](resources--bgp--reference--group-002.md#canonical-3220103110002133-1333320020233033-3313320203230220-1020021110111113-1203113112011300-1131212123120133-0030303222123001-0021021002032023) |
| `where.site` | [where.site](resources--bgp--reference--group-002.md#canonical-3230002111332233-2023012332221333-0300112011323011-3021321320110121-3301033130223130-2320000113002011-0312111033010033-0230123331200302) |
| `where.site.disable_internet_vip` | [where.site.disable_internet_vip](resources--bgp--reference--group-002.md#canonical-1223003311301221-3303101102323112-0130122122213132-3213033313020013-0332110231202011-0311010303133210-1233303323101111-2213302212020112) |
| `where.site.enable_internet_vip` | [where.site.enable_internet_vip](resources--bgp--reference--group-002.md#canonical-2130001020322233-1313231130113100-2213110233332300-0132130310331332-2030132012303100-3111212022301103-1210023031320011-1221101302100320) |
| `where.site.network_type` | [where.site.network_type](resources--bgp--reference--group-002.md#canonical-1131203123322230-2323011230100021-1101000301011201-3202201332003031-1002312203020221-0330322033223131-3013133331210020-2322112211102113) |
| `where.site.ref` | [where.site.ref](resources--bgp--reference--group-002.md#canonical-1022131212131330-1012011213112133-1133032212312132-1321322130010022-2133012301012301-1003202133302321-1230212333013323-1213113201102213) |
| `where.site.ref.kind` | [where.site.ref.kind](resources--bgp--reference--group-002.md#canonical-1221313221202101-3133032112000212-0303210001112330-2023023231010330-1120002203333000-1322013211120023-2320223311030212-1230110231102333) |
| `where.site.ref.name` | [where.site.ref.name](resources--bgp--reference--group-002.md#canonical-3000033133131221-1033311203321011-2231120033010230-1002123011200331-2002130211031233-3200033123022322-2102032200020030-3312023010130101) |
| `where.site.ref.namespace` | [where.site.ref.namespace](resources--bgp--reference--group-002.md#canonical-2131232113200123-1302022101201111-2210303202010221-2131101120132310-3220223320020301-2302301112132333-1111131232201011-2032312203111202) |
| `where.site.ref.tenant` | [where.site.ref.tenant](resources--bgp--reference--group-002.md#canonical-0121001320211210-3211223111302030-2210313003210231-3102231201020320-3013323321333302-0233013000300232-1311110103321330-1221301131112300) |
| `where.site.ref.uid` | [where.site.ref.uid](resources--bgp--reference--group-002.md#canonical-3313011120031311-3202203133010111-2231011312210220-0021201021230130-3323302331233231-2203003223322033-0202233310321100-0310130100132322) |
| `where.virtual_site` | [where.virtual_site](resources--bgp--reference--group-002.md#canonical-1133111233031221-3102311311110012-2201202300020232-0322320321311213-3110230103020321-0230221330332120-2111011013030311-3121133232001101) |
| `where.virtual_site.disable_internet_vip` | [where.virtual_site.disable_internet_vip](resources--bgp--reference--group-002.md#canonical-0300131333311032-0222313321100222-3221230201010202-1003022200310020-0022010312020023-0331303101133331-3123213210333210-3201130230330330) |
| `where.virtual_site.enable_internet_vip` | [where.virtual_site.enable_internet_vip](resources--bgp--reference--group-002.md#canonical-3003323211221211-0232122113322013-2220301222332123-1121311313220012-1033303201311333-1330130101023011-3122132312022100-2220231213322111) |
| `where.virtual_site.network_type` | [where.virtual_site.network_type](resources--bgp--reference--group-002.md#canonical-0112202232222320-0110213112121033-3210233131322023-1110011200012203-3011200100021022-0010021013122001-0211110200113232-1313312213312210) |
| `where.virtual_site.ref` | [where.virtual_site.ref](resources--bgp--reference--group-002.md#canonical-0102011002200311-0011100111101313-2313132101001233-0331122211313201-2132120201311322-3102033101132231-2132102022023013-1223102030322302) |
| `where.virtual_site.ref.kind` | [where.virtual_site.ref.kind](resources--bgp--reference--group-002.md#canonical-0032033333233030-2131302231320222-2003130333112330-1203210132201301-2011133212311313-1023111001111032-1020012120220001-2022103030000230) |
| `where.virtual_site.ref.name` | [where.virtual_site.ref.name](resources--bgp--reference--group-002.md#canonical-2110132321123303-1002131313200122-0200200333011121-1330001022021312-2110221012111332-3001101223012032-3210030012111203-0323031033120030) |
| `where.virtual_site.ref.namespace` | [where.virtual_site.ref.namespace](resources--bgp--reference--group-002.md#canonical-3032301222230330-1203133332231311-2332313130213303-3210102002210120-1313013221312100-0021001203323312-3322221230021222-2111102101313010) |
| `where.virtual_site.ref.tenant` | [where.virtual_site.ref.tenant](resources--bgp--reference--group-002.md#canonical-0133120002032030-2133313313320012-3312233100321231-0102201113022123-2222223322221030-2203100021113101-2300222123323111-2031003200200200) |
| `where.virtual_site.ref.uid` | [where.virtual_site.ref.uid](resources--bgp--reference--group-002.md#canonical-3101100121323210-3021302301200202-1122032321131321-3222132031031123-1131102123030131-2033013230032131-1021110103302222-3103322230113331) |

<a id="canonical-3103003320212223-3323032330310220-0112102013012011-3303212132102001-2201310010330000-1120110231210311-2311000111033323-0133120330101211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bgp_parameters` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- bgp_parameters

<a id="canonical-0312322210330131-3232021030003310-2001103001223211-3110312220033113-1222332202113310-2100033021033230-0311030032001133-3311301130030123"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for bgp parameters.

Additional upstream details:

BGP parameters for the local site.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("asn"),
  validators.ConflictingObjectAttributes("from_site",
    "ip_address"),
  validators.ConflictingObjectAttributes("from_site",
    "local_address"),
  validators.ConflictingObjectAttributes("ip_address",
    "local_address")}
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
  "x-ves-oneof-field-router_id_choice": "[\"from_site\",\"ip_address\",\"local_address\"]"
}
```

Terraform syntax:

```terraform
bgp_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-0332233130231133-1233013012221213-3011110311001003-3003022313201103-2333132302321322-3012130023012010-2310321323020331-3200303102210322"></a>

### Direct properties for `bgp_parameters`

<a id="canonical-2210011010202021-1112233002011032-1303133231110233-0013021320121213-0223023202010230-1123002003112120-2331202131231130-2222102331332133"></a>

#### `bgp_parameters.asn` property

Type: `"number"`. Optional.

ASN. Autonomous System Number.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [from_site](resources--bgp--reference--group-001.md#canonical-1203231031133321-1213030210133022-1021311000321330-3111332112320132-1123101222032032-2112120103330130-2103030333023230-2203023110223211): complete subsection reference.

<a id="canonical-3322221000221021-1111122120131301-3320013312220323-0231133002232121-0132112313311221-0131112220032300-3002211122122220-0102320211321111"></a>

<a id="canonical-0102003311102311-3302223233100230-1301323301320001-1302003033301222-1123131101231023-2311203330330230-0330130023002303-1013133203320203"></a>

#### `bgp_parameters.ip_address` property

Type: `"string"`. Optional.

Exclusive with \[from\_site local\_address\] Use the configured IPv4 Address as Router ID.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
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
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
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

- [local_address](resources--bgp--reference--group-001.md#canonical-0322212333332301-3100201110222010-3210102120330212-2321123232001100-2011033130030100-3032303131220330-2123012131323112-3321230120321320): complete subsection reference.

<a id="canonical-1203231031133321-1213030210133022-1021311000321330-3111332112320132-1123101222032032-2112120103330130-2103030333023230-2203023110223211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bgp_parameters.from_site` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [bgp_parameters](resources--bgp--reference--group-001.md#canonical-3103003320212223-3323032330310220-0112102013012011-3303212132102001-2201310010330000-1120110231210311-2311000111033323-0133120330101211)
- bgp_parameters.from_site

<a id="canonical-2113202301310011-2311002103002231-3323001321333201-2301102302033103-2213002031211231-0222000313311120-0111003022132023-2210301231003333"></a>

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
from_site = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0322212333332301-3100201110222010-3210102120330212-2321123232001100-2011033130030100-3032303131220330-2123012131323112-3321230120321320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bgp_parameters.local_address` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [bgp_parameters](resources--bgp--reference--group-001.md#canonical-3103003320212223-3323032330310220-0112102013012011-3303212132102001-2201310010330000-1120110231210311-2311000111033323-0133120330101211)
- bgp_parameters.local_address

<a id="canonical-0302311312133023-1320022223303223-0012111201033211-0300312232012212-1322321031202032-1123210332312013-0333130031213103-1030322322200032"></a>

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
local_address = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3132002313303222-3210021222211033-2013001002330232-0123121112132233-0321012202000010-0112130200010312-2201100311111101-1320000132020103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- peers

<a id="canonical-2331002010131000-1222130010211330-2332203203301122-0311312231230211-3210123131213021-2203101011320030-0133100031132323-0011022011223210"></a>

Type: `"object"`. list nested block, Optional.

Peers. List of peers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("bfd_disabled",
    "bfd_enabled"),
  validators.ConflictingListObjectAttributes("disable_spec",
    "routing_policies"),
  validators.ConflictingListObjectAttributes("ebgp_multihop_disabled",
    "ebgp_multihop_enabled"),
  validators.ConflictingListObjectAttributes("passive_mode_disabled",
    "passive_mode_enabled")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
peers {
  # Configure direct properties listed below.
}
```

<a id="canonical-1021313222330300-0232101103123111-0022010213011003-3321223012112201-3200023312330032-3122130012332013-1111110111222322-1120212130133030"></a>

### Direct properties for `peers`

- [bfd_disabled](resources--bgp--reference--group-001.md#canonical-2013121310213103-2313110101031330-3022330022010111-3213002001220000-3213020112221333-0023131022303332-2320301131032221-1321302031210310): complete subsection reference.

- [bfd_enabled](resources--bgp--reference--group-001.md#canonical-1111132102000210-3112203200302003-3322231311220121-1202313230202011-3110310023313123-2021200210323032-0220020320031302-1232321101020121): complete subsection reference.

- [disable_spec](resources--bgp--reference--group-001.md#canonical-3231001033223101-3130333131100122-2020103032003023-2211220231101111-1012333102111000-1032100000010322-1131132332132210-3110131031023223): complete subsection reference.

- [ebgp_multihop_disabled](resources--bgp--reference--group-001.md#canonical-0113332302222133-0010030322122122-1123031002202320-2021122212101331-2200032002111131-3023210221322221-1002000223112121-2202021112303120): complete subsection reference.

- [ebgp_multihop_enabled](resources--bgp--reference--group-001.md#canonical-0301303123130031-3001033002320311-0301331000013201-0102300210132311-2110203011200130-3322230203303011-0201210022213302-1032032102021122): complete subsection reference.

- [external](resources--bgp--reference--group-001.md#canonical-0001200002001233-2222102330023121-3230223320030311-0130211020130023-3230133120112331-2233132120332220-1210032131313111-3130022130010331): complete subsection reference.

<a id="canonical-1120332223200211-1002002011120110-1110120212101301-0220311002123000-1100010231023013-3201122123302033-1220313323132102-1203311131132103"></a>

<a id="canonical-3310210231031200-0321122112011302-3030013102212311-3313220210230333-0001100233132222-0233123310311203-2001233002211023-0000222032013301"></a>

#### `peers.label` property

Type: `"string"`. Optional.

Label. Specify whether this peer should be.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "labeling",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [metadata](resources--bgp--reference--group-002.md#canonical-1102332321320322-3302110300310021-2011013232000101-0212323123110010-0032202032311310-3223212030010113-2101332010233031-3022021202302020): complete subsection reference.

- [passive_mode_disabled](resources--bgp--reference--group-002.md#canonical-3001222122101221-1132332210123021-3233102220031330-0130111121100131-0022302320313231-1211122232300220-3033321021113002-1312113030132001): complete subsection reference.

- [passive_mode_enabled](resources--bgp--reference--group-002.md#canonical-0232120302301113-3232111320111310-2033121311313201-3030333321022232-1011013210322100-2321202101133221-3333110022231200-3330323000303031): complete subsection reference.

- [routing_policies](resources--bgp--reference--group-002.md#canonical-0222130023021003-1202011230202002-3321120110203013-1131302030313202-3030311220230000-0033202301032000-1313212311121113-3320021330020313): complete subsection reference.

<a id="canonical-2013121310213103-2313110101031330-3022330022010111-3213002001220000-3213020112221333-0023131022303332-2320301131032221-1321302031210310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.bfd_disabled` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [peers](resources--bgp--reference--group-001.md#canonical-3132002313303222-3210021222211033-2013001002330232-0123121112132233-0321012202000010-0112130200010312-2201100311111101-1320000132020103)
- peers.bfd_disabled

<a id="canonical-1320021021213230-0121013103231032-3033012132101202-2112231330103033-0021310211233212-3210000031130311-0120002301130103-0332033102333331"></a>

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
bfd_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1111132102000210-3112203200302003-3322231311220121-1202313230202011-3110310023313123-2021200210323032-0220020320031302-1232321101020121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.bfd_enabled` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [peers](resources--bgp--reference--group-001.md#canonical-3132002313303222-3210021222211033-2013001002330232-0123121112132233-0321012202000010-0112130200010312-2201100311111101-1320000132020103)
- peers.bfd_enabled

<a id="canonical-0022030032232021-3230222003321222-2002123211130030-1103322302013120-3000013330311213-2322123312101021-1333003022313121-2013322121212333"></a>

Type: `"object"`. single nested block, Optional.

BFD. BFD parameters.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("multiplier",
    "receive_interval_milliseconds",
    "transmit_interval_milliseconds")}
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
bfd_enabled {
  # Configure direct properties listed below.
}
```

<a id="canonical-1322020120212132-0332222122312301-3002102003323113-0112222332202110-2302003001113112-1322220022220120-2011232111303120-0000013130102103"></a>

### Direct properties for `peers.bfd_enabled`

<a id="canonical-0121212102201310-0200133211230321-2023003003003212-1100103022233133-1122122121301203-0111121120222202-2122130110232002-0203030011222012"></a>

#### `peers.bfd_enabled.multiplier` property

Type: `"number"`. Optional.

Specify Number of missed packets to bring session down'.

Additional upstream details:

Specify Number of missed packets to bring session down"

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(2, 255),
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="canonical-0211302112202210-2122120231123031-0312221013232220-0201013011313202-0022132023332000-3223001131012111-3001120111020132-0320022310020212"></a>

<a id="canonical-1030332101220111-2021031132111030-3002320121010210-1112311203013010-0103003103301212-1202212321231003-1320122211312222-1123223230303310"></a>

#### `peers.bfd_enabled.receive_interval_milliseconds` property

Type: `"number"`. Optional.

BFD receive interval timer, in milliseconds.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(300, 60000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 300
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-0010211011000011-1002310030223330-3331122300210010-0221023210223212-1312012123133102-0131332012010002-0323303132333212-0020122111112212"></a>

<a id="canonical-0200100213022001-3130003012111211-2003331312323210-3201003023001201-2103013001202302-0101213221302330-2022100102111312-1012203111332223"></a>

#### `peers.bfd_enabled.transmit_interval_milliseconds` property

Type: `"number"`. Optional.

BFD transmit interval timer, in milliseconds.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(300, 60000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 300
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "300",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-3231001033223101-3130333131100122-2020103032003023-2211220231101111-1012333102111000-1032100000010322-1131132332132210-3110131031023223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.disable_spec` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [peers](resources--bgp--reference--group-001.md#canonical-3132002313303222-3210021222211033-2013001002330232-0123121112132233-0321012202000010-0112130200010312-2201100311111101-1320000132020103)
- peers.disable_spec

<a id="canonical-1321331310123103-1122023202021200-1100013311120021-0120332022331121-3033020023331123-2103311132221023-2122132011011200-0210130230320301"></a>

Type: `["object", {}]`. Optional.

Enable this option

Terraform syntax:

```terraform
disable_spec = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0113332302222133-0010030322122122-1123031002202320-2021122212101331-2200032002111131-3023210221322221-1002000223112121-2202021112303120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.ebgp_multihop_disabled` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [peers](resources--bgp--reference--group-001.md#canonical-3132002313303222-3210021222211033-2013001002330232-0123121112132233-0321012202000010-0112130200010312-2201100311111101-1320000132020103)
- peers.ebgp_multihop_disabled

<a id="canonical-0331333310011323-1302320302330230-3313100321111320-1230102121201023-0300213323323301-2122022303003022-1012312301133212-1102311103310211"></a>

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
ebgp_multihop_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0301303123130031-3001033002320311-0301331000013201-0102300210132311-2110203011200130-3322230203303011-0201210022213302-1032032102021122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.ebgp_multihop_enabled` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [peers](resources--bgp--reference--group-001.md#canonical-3132002313303222-3210021222211033-2013001002330232-0123121112132233-0321012202000010-0112130200010312-2201100311111101-1320000132020103)
- peers.ebgp_multihop_enabled

<a id="canonical-1310331200012212-1131011331210331-0030132003311033-3001203111203210-3300223203130132-0123030010010132-1020333120320230-3322201312100210"></a>

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
ebgp_multihop_enabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0001200002001233-2222102330023121-3230223320030311-0130211020130023-3230133120112331-2233132120332220-1210032131313111-3130022130010331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [peers](resources--bgp--reference--group-001.md#canonical-3132002313303222-3210021222211033-2013001002330232-0123121112132233-0321012202000010-0112130200010312-2201100311111101-1320000132020103)
- peers.external

<a id="canonical-2000211200300200-3333300230331212-2112331202300310-1133112322301320-3130111020320021-1002232011133032-1010302030313013-3302132222003210"></a>

Type: `"object"`. single nested block, Optional.

External BGP Peer. External BGP Peer parameters.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("asn",
    "port"),
  validators.ConflictingObjectAttributes("address",
    "default_gateway"),
  validators.ConflictingObjectAttributes("address",
    "disable_spec"),
  validators.ConflictingObjectAttributes("address",
    "external_connector"),
  validators.ConflictingObjectAttributes("address",
    "from_site"),
  validators.ConflictingObjectAttributes("address",
    "subnet_begin_offset"),
  validators.ConflictingObjectAttributes("address",
    "subnet_end_offset"),
  validators.ConflictingObjectAttributes("address_ipv6",
    "default_gateway_v6"),
  validators.ConflictingObjectAttributes("address_ipv6",
    "disable_v6"),
  validators.ConflictingObjectAttributes("address_ipv6",
    "from_site_v6"),
  validators.ConflictingObjectAttributes("address_ipv6",
    "subnet_begin_offset_v6"),
  validators.ConflictingObjectAttributes("address_ipv6",
    "subnet_end_offset_v6"),
  validators.ConflictingObjectAttributes("default_gateway",
    "disable_spec"),
  validators.ConflictingObjectAttributes("default_gateway",
    "external_connector"),
  validators.ConflictingObjectAttributes("default_gateway",
    "from_site"),
  validators.ConflictingObjectAttributes("default_gateway",
    "subnet_begin_offset"),
  validators.ConflictingObjectAttributes("default_gateway",
    "subnet_end_offset"),
  validators.ConflictingObjectAttributes("default_gateway_v6",
    "disable_v6"),
  validators.ConflictingObjectAttributes("default_gateway_v6",
    "from_site_v6"),
  validators.ConflictingObjectAttributes("default_gateway_v6",
    "subnet_begin_offset_v6"),
  validators.ConflictingObjectAttributes("default_gateway_v6",
    "subnet_end_offset_v6"),
  validators.ConflictingObjectAttributes("disable_spec",
    "external_connector"),
  validators.ConflictingObjectAttributes("disable_spec",
    "from_site"),
  validators.ConflictingObjectAttributes("disable_spec",
    "subnet_begin_offset"),
  validators.ConflictingObjectAttributes("disable_spec",
    "subnet_end_offset"),
  validators.ConflictingObjectAttributes("disable_v6",
    "from_site_v6"),
  validators.ConflictingObjectAttributes("disable_v6",
    "subnet_begin_offset_v6"),
  validators.ConflictingObjectAttributes("disable_v6",
    "subnet_end_offset_v6"),
  validators.ConflictingObjectAttributes("external_connector",
    "from_site"),
  validators.ConflictingObjectAttributes("external_connector",
    "subnet_begin_offset"),
  validators.ConflictingObjectAttributes("external_connector",
    "subnet_end_offset"),
  validators.ConflictingObjectAttributes("from_site",
    "subnet_begin_offset"),
  validators.ConflictingObjectAttributes("from_site",
    "subnet_end_offset"),
  validators.ConflictingObjectAttributes("from_site_v6",
    "subnet_begin_offset_v6"),
  validators.ConflictingObjectAttributes("from_site_v6",
    "subnet_end_offset_v6"),
  validators.ConflictingObjectAttributes("interface",
    "interface_list"),
  validators.ConflictingObjectAttributes("md5_auth_key",
    "no_authentication"),
  validators.ConflictingObjectAttributes("subnet_begin_offset",
    "subnet_end_offset"),
  validators.ConflictingObjectAttributes("subnet_begin_offset_v6",
    "subnet_end_offset_v6")}
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
  "x-ves-oneof-field-address_choice": "[\"address\",\"default_gateway\",\"disable\",\"external_connector\",\"from_site\",\"subnet_begin_offset\",\"subnet_end_offset\"]",
  "x-ves-oneof-field-address_choice_v6": "[\"address_ipv6\",\"default_gateway_v6\",\"disable_v6\",\"from_site_v6\",\"subnet_begin_offset_v6\",\"subnet_end_offset_v6\"]",
  "x-ves-oneof-field-auth_choice": "[\"md5_auth_key\",\"no_authentication\"]",
  "x-ves-oneof-field-interface_choice": "[\"interface\",\"interface_list\"]"
}
```

Terraform syntax:

```terraform
external {
  # Configure direct properties listed below.
}
```

<a id="canonical-0012010010230311-0030132300213132-1131111131231230-1002211300030133-0332121021303202-0123232030232032-3003131033031112-3112131331102311"></a>

### Direct properties for `peers.external`

<a id="canonical-2132323300330013-3002121001320131-0300233300100011-1233212201222132-3330032033121301-3023300311011211-1301003331000110-1031110030112211"></a>

#### `peers.external.address` property

Type: `"string"`. Optional.

Exclusive with \[default\_gateway disable external\_connector from\_site subnet\_begin\_offset
subnet\_end\_offset\] Specify IPv4 peer address.

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-2302302231300223-1201123330212031-1103020203002032-1230030032013322-2102102301100102-2002221001302031-3011201132122330-2333003020132223"></a>

<a id="canonical-1223000322330110-0313100123321002-1231121300022131-1321300022010130-0311223131111322-0103301220032230-1311023303131012-3303112221112022"></a>

#### `peers.external.address_ipv6` property

Type: `"string"`. Optional.

Exclusive with \[default\_gateway\_v6 disable\_v6 from\_site\_v6 subnet\_begin\_offset\_v6
subnet\_end\_offset\_v6\] Specify peer IPv6 address.

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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-1213230011013003-2212331323131112-0233011230222300-3103013031013010-3103131011112202-0211302100323233-3030310131000101-1000121100311102"></a>

<a id="canonical-0212031302320223-1211322013310102-0331223213121203-1220003100333311-0120012211030312-3202333303123101-2211002132033322-2201330321112033"></a>

#### `peers.external.asn` property

Type: `"number"`. Optional.

ASN. Autonomous System Number for BGP peer.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [default_gateway](resources--bgp--reference--group-001.md#canonical-2212101233113330-0031311303322101-0312120233233010-1211321000332010-0100221300213003-0213010231021133-0310023031002311-0010311201321203): complete subsection reference.

- [default_gateway_v6](resources--bgp--reference--group-001.md#canonical-2300330311232023-1011301211201113-0220120311311001-2111332232300200-2232230210023311-0203113330330022-1233333102302313-0030013313220131): complete subsection reference.

- [disable_spec](resources--bgp--reference--group-001.md#canonical-1011022011002203-2220112011132201-3213121331220030-3201302120310001-2331220223231323-0222120113121012-2131320233111232-2330333123012112): complete subsection reference.

- [disable_v6](resources--bgp--reference--group-001.md#canonical-0013311310322112-3001000331330021-0110102110210113-3301310012111201-0211033322123012-0102203210313302-3321000213302030-3201223321111001): complete subsection reference.

- [external_connector](resources--bgp--reference--group-001.md#canonical-3033122311300320-1233031202203321-3000023113132311-3200130312000311-3121100022122022-2103101010201030-0320001012211132-0110023021321220): complete subsection reference.

- [family_inet](resources--bgp--reference--group-001.md#canonical-0113212100231110-0013013110022232-3211013200330312-2210023313122002-3331020102102121-3233203020233100-2133131030203003-1011022003130203): complete subsection reference.

- [from_site](resources--bgp--reference--group-001.md#canonical-0132302311122112-1301303020201320-1220203233112002-2013031213311202-0011202223333303-1001200120131221-2131113223230003-0200221211302222): complete subsection reference.

- [from_site_v6](resources--bgp--reference--group-001.md#canonical-0131210123203122-1210010133133000-2321323123132230-0322210212322222-1110003222312110-0010112302302231-1122222133212210-3100031020223332): complete subsection reference.

- [interface](resources--bgp--reference--group-001.md#canonical-3223001231322213-1010111113010202-0212323310131132-2201030020233130-0033301131001132-3323121211013230-1333032000321221-2121121010113321): complete subsection reference.

- [interface_list](resources--bgp--reference--group-002.md#canonical-0221000322321030-1113211202122102-0133131202033231-1331110121001123-0111220333300112-3122311132013122-0321033031322012-3310302200131001): complete subsection reference.

<a id="canonical-0022210203333322-0013131030000000-0232320203202012-1121122322103210-1131033101321103-3001233013300122-3323300003333030-0113111311312120"></a>

<a id="canonical-2111320022330301-3013121021130030-1310211011112120-1011121222110103-3130220123130313-0300212130330110-1330223020232211-1112201211302022"></a>

#### `peers.external.md5_auth_key` property

Type: `"string"`. Optional.

Exclusive with \[no\_authentication\] MD5 key for protecting BGP Sessions (RFC 2385).

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [no_authentication](resources--bgp--reference--group-002.md#canonical-0330223011231313-0131302323100133-2232102110312200-0331033133303032-1301100120202101-2320101223111322-3100000322300223-0032212000310102): complete subsection reference.

<a id="canonical-3312112130333030-2003101100100301-3211220300331030-2301112012103001-2131232002222201-1022303203233331-3313232320212011-3220021230023220"></a>

<a id="canonical-3211300310112331-1303223130000220-3030003103110320-2130030121001102-1000111220222301-2201331310110111-1002333110300233-3210301211110122"></a>

#### `peers.external.port` property

Type: `"number"`. Optional.

Peer Port. Peer TCP port number.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1131023001230033-2131320001131012-2222110133012021-1331021331111102-3230113122001020-2210022013322003-2313311030331002-1311033220102133"></a>

<a id="canonical-3103330210301302-3301022001223310-0201310001300222-0002122222310203-1202223231022333-1011210132031131-2213120303300310-3003002120102012"></a>

#### `peers.external.subnet_begin_offset` property

Type: `"number"`. Optional.

Exclusive with \[address default\_gateway disable external\_connector from\_site
subnet\_end\_offset\] Calculate peer address using offset from the beginning of the subnet.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 32),
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-3012101020021013-3120021310310203-1012330213231222-2123122022323100-2210210022221221-1120013202131331-0102332132113100-2321131033013102"></a>

<a id="canonical-2113102012300121-3000320212120220-3223123213120112-2310211330220133-2122321011212130-3023030021212132-0232112203013021-2002112122320102"></a>

#### `peers.external.subnet_begin_offset_v6` property

Type: `"number"`. Optional.

Exclusive with \[address\_ipv6 default\_gateway\_v6 disable\_v6 from\_site\_v6
subnet\_end\_offset\_v6\] Calculate peer address using offset from the beginning of the subnet.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 32),
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-0020033300211101-2010103202023233-0103003201023010-0330031330313311-0301132100131101-1002021101033300-2130022333212020-1221021020201122"></a>

<a id="canonical-1120012320210213-2321232202003212-3103323013320303-3200110212112222-2300332030320201-2131202033132023-1000003223132013-1121301212022231"></a>

#### `peers.external.subnet_end_offset` property

Type: `"number"`. Optional.

Exclusive with \[address default\_gateway disable external\_connector from\_site
subnet\_begin\_offset\] Calculate peer address using offset from the end of the subnet.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 32),
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-3032022233302322-2000123013313122-1320321331221101-1023303023321120-0001310113031331-0333133202133333-3221311001231001-0223323313103230"></a>

<a id="canonical-1132302013111003-1030133123123320-2033220222122103-1130331011212000-1301110220000001-3131121302122102-0102031210100030-1131121322202101"></a>

#### `peers.external.subnet_end_offset_v6` property

Type: `"number"`. Optional.

Exclusive with \[address\_ipv6 default\_gateway\_v6 disable\_v6 from\_site\_v6
subnet\_begin\_offset\_v6\] Calculate peer address using offset from the end of the subnet.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 32),
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-2212101233113330-0031311303322101-0312120233233010-1211321000332010-0100221300213003-0213010231021133-0310023031002311-0010311201321203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external.default_gateway` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [peers](resources--bgp--reference--group-001.md#canonical-3132002313303222-3210021222211033-2013001002330232-0123121112132233-0321012202000010-0112130200010312-2201100311111101-1320000132020103)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0001200002001233-2222102330023121-3230223320030311-0130211020130023-3230133120112331-2233132120332220-1210032131313111-3130022130010331)
- peers.external.default_gateway

<a id="canonical-1113031133210210-0123331232112132-1302223112210321-3302013122032132-2231203100002312-0033210020321232-0221232301002113-2022000031331230"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default gateway.

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
default_gateway = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2300330311232023-1011301211201113-0220120311311001-2111332232300200-2232230210023311-0203113330330022-1233333102302313-0030013313220131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external.default_gateway_v6` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [peers](resources--bgp--reference--group-001.md#canonical-3132002313303222-3210021222211033-2013001002330232-0123121112132233-0321012202000010-0112130200010312-2201100311111101-1320000132020103)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0001200002001233-2222102330023121-3230223320030311-0130211020130023-3230133120112331-2233132120332220-1210032131313111-3130022130010331)
- peers.external.default_gateway_v6

<a id="canonical-0013232310220330-0132311233232121-3330232023211331-3111312201102330-3131103202200232-2313123023233310-1220111310303210-1302110202222020"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default gateway v6.

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
default_gateway_v6 = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1011022011002203-2220112011132201-3213121331220030-3201302120310001-2331220223231323-0222120113121012-2131320233111232-2330333123012112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external.disable_spec` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [peers](resources--bgp--reference--group-001.md#canonical-3132002313303222-3210021222211033-2013001002330232-0123121112132233-0321012202000010-0112130200010312-2201100311111101-1320000132020103)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0001200002001233-2222102330023121-3230223320030311-0130211020130023-3230133120112331-2233132120332220-1210032131313111-3130022130010331)
- peers.external.disable_spec

<a id="canonical-0120123011321323-0101222322033033-1131313130100232-1010100030022212-0122133230131210-2211123330302031-0001211010121232-0101111313112110"></a>

Type: `["object", {}]`. Optional.

Enable this option

Terraform syntax:

```terraform
disable_spec = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0013311310322112-3001000331330021-0110102110210113-3301310012111201-0211033322123012-0102203210313302-3321000213302030-3201223321111001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external.disable_v6` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [peers](resources--bgp--reference--group-001.md#canonical-3132002313303222-3210021222211033-2013001002330232-0123121112132233-0321012202000010-0112130200010312-2201100311111101-1320000132020103)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0001200002001233-2222102330023121-3230223320030311-0130211020130023-3230133120112331-2233132120332220-1210032131313111-3130022130010331)
- peers.external.disable_v6

<a id="canonical-1302000010212312-0021120012002013-3201230003031012-1021103202020302-3100001102311320-3020220113302100-1323213011033320-2322231101032133"></a>

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
disable_v6 = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3033122311300320-1233031202203321-3000023113132311-3200130312000311-3121100022122022-2103101010201030-0320001012211132-0110023021321220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external.external_connector` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [peers](resources--bgp--reference--group-001.md#canonical-3132002313303222-3210021222211033-2013001002330232-0123121112132233-0321012202000010-0112130200010312-2201100311111101-1320000132020103)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0001200002001233-2222102330023121-3230223320030311-0130211020130023-3230133120112331-2233132120332220-1210032131313111-3130022130010331)
- peers.external.external_connector

<a id="canonical-0321330001221023-1302303330111333-1322021211300331-3031310123213303-1133002000210012-0033321300323131-0302320232210311-0001211022010200"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for external connector.

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
external_connector = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0113212100231110-0013013110022232-3211013200330312-2210023313122002-3331020102102121-3233203020233100-2133131030203003-1011022003130203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external.family_inet` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [peers](resources--bgp--reference--group-001.md#canonical-3132002313303222-3210021222211033-2013001002330232-0123121112132233-0321012202000010-0112130200010312-2201100311111101-1320000132020103)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0001200002001233-2222102330023121-3230223320030311-0130211020130023-3230133120112331-2233132120332220-1210032131313111-3130022130010331)
- peers.external.family_inet

<a id="canonical-1312313031110203-2100113022323312-1333112102111220-2132312110103010-3030203222222010-0010013201023102-3133211121110102-2302332311103303"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for family inet.

Additional upstream details:

Parameters for inet family.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_spec",
    "enable")}
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
  "x-ves-oneof-field-enable_choice": "[\"disable\",\"enable\"]"
}
```

Terraform syntax:

```terraform
family_inet {
  # Configure direct properties listed below.
}
```

<a id="canonical-0130312333130023-2123002203112303-2022311210003231-1110232301020213-1320130323321120-3000131310031322-0230103223212031-0123223110331003"></a>

### Direct properties for `peers.external.family_inet`

- [disable_spec](resources--bgp--reference--group-001.md#canonical-3201320022231221-0320233111300311-0030232223110330-3032230033301103-0131331322211133-1222122310122320-2111021100313020-2122130022013022): complete subsection reference.

- [enable](resources--bgp--reference--group-001.md#canonical-0021013112222321-0231001111030131-2100032231101313-2311200201001031-0223133202020303-2331013121230332-1330210233302220-1300301211203301): complete subsection reference.

<a id="canonical-3201320022231221-0320233111300311-0030232223110330-3032230033301103-0131331322211133-1222122310122320-2111021100313020-2122130022013022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external.family_inet.disable_spec` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [peers](resources--bgp--reference--group-001.md#canonical-3132002313303222-3210021222211033-2013001002330232-0123121112132233-0321012202000010-0112130200010312-2201100311111101-1320000132020103)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0001200002001233-2222102330023121-3230223320030311-0130211020130023-3230133120112331-2233132120332220-1210032131313111-3130022130010331)
- [peers.external.family_inet](resources--bgp--reference--group-001.md#canonical-0113212100231110-0013013110022232-3211013200330312-2210023313122002-3331020102102121-3233203020233100-2133131030203003-1011022003130203)
- peers.external.family_inet.disable_spec

<a id="canonical-1312122231020110-1022002103230303-2010313212321131-2210122303110321-0233012110201331-1322322122010200-0312011122333230-0311000100212302"></a>

Type: `["object", {}]`. Optional.

Enable this option

Terraform syntax:

```terraform
disable_spec = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0021013112222321-0231001111030131-2100032231101313-2311200201001031-0223133202020303-2331013121230332-1330210233302220-1300301211203301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external.family_inet.enable` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [peers](resources--bgp--reference--group-001.md#canonical-3132002313303222-3210021222211033-2013001002330232-0123121112132233-0321012202000010-0112130200010312-2201100311111101-1320000132020103)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0001200002001233-2222102330023121-3230223320030311-0130211020130023-3230133120112331-2233132120332220-1210032131313111-3130022130010331)
- [peers.external.family_inet](resources--bgp--reference--group-001.md#canonical-0113212100231110-0013013110022232-3211013200330312-2210023313122002-3331020102102121-3233203020233100-2133131030203003-1011022003130203)
- peers.external.family_inet.enable

<a id="canonical-3033123202300311-1023032022332321-3320113212223320-2110210100013012-3331201130332003-3121222313003032-0131300331230000-2023300121011103"></a>

Type: `"object"`. single nested block, Optional.

Unicast IPv4. IPv4 Unicast.

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
enable {
  # Configure direct properties listed below.
}
```

<a id="canonical-1223011320011011-1222033333113231-2323221333203230-1311113311002211-2033200203111312-3021310313203203-2323133233023130-0033312323020331"></a>

### Direct properties for `peers.external.family_inet.enable`

- [aggregation](resources--bgp--reference--group-001.md#canonical-3121211121012033-0303111220033220-1121211302201112-0312011131300230-0033333300331030-2021131003100323-3030310320023320-2001121333332313): complete subsection reference.

<a id="canonical-3121211121012033-0303111220033220-1121211302201112-0312011131300230-0033333300331030-2021131003100323-3030310320023320-2001121333332313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external.family_inet.enable.aggregation` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [peers](resources--bgp--reference--group-001.md#canonical-3132002313303222-3210021222211033-2013001002330232-0123121112132233-0321012202000010-0112130200010312-2201100311111101-1320000132020103)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0001200002001233-2222102330023121-3230223320030311-0130211020130023-3230133120112331-2233132120332220-1210032131313111-3130022130010331)
- [peers.external.family_inet](resources--bgp--reference--group-001.md#canonical-0113212100231110-0013013110022232-3211013200330312-2210023313122002-3331020102102121-3233203020233100-2133131030203003-1011022003130203)
- [peers.external.family_inet.enable](resources--bgp--reference--group-001.md#canonical-0021013112222321-0231001111030131-2100032231101313-2311200201001031-0223133202020303-2331013121230332-1330210233302220-1300301211203301)
- peers.external.family_inet.enable.aggregation

<a id="canonical-2221232100031210-3231201200300331-0022303303101131-2210032112033201-3002313010023212-1211203123320221-2311103212003032-0122311202031100"></a>

Type: `"object"`. list nested block, Optional.

BGP aggregation prefixes are shared among all peers, aggregation configured under any peer will take
effect on all peers. Aggregation in BGP occurs only when more specific routes exist in the routing
table and applies to outbound advertisements.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
aggregation {
  # Configure direct properties listed below.
}
```

<a id="canonical-2122020322213112-3310202211200201-2123000303100032-1312000021231031-1110112332021121-1200123111100032-3110121002033103-1222130220200103"></a>

### Direct properties for `peers.external.family_inet.enable.aggregation`

<a id="canonical-0203022330221313-1000333012130200-1303232123121100-1011002321001031-1022113321033223-3322313330200302-0003000331323333-2212303201300331"></a>

#### `peers.external.family_inet.enable.aggregation.ip_prefix` property

Type: `"string"`. Optional.

IP Prefix. Specify IPv4 subnet for aggregation.

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
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

- [options](resources--bgp--reference--group-001.md#canonical-3220221002330213-2220220311222320-2232032232101021-1011011200301021-2210323313020322-3322000210321201-3301033001112220-1330001221223233): complete subsection reference.

<a id="canonical-3220221002330213-2220220311222320-2232032232101021-1011011200301021-2210323313020322-3322000210321201-3301033001112220-1330001221223233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external.family_inet.enable.aggregation.options` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [peers](resources--bgp--reference--group-001.md#canonical-3132002313303222-3210021222211033-2013001002330232-0123121112132233-0321012202000010-0112130200010312-2201100311111101-1320000132020103)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0001200002001233-2222102330023121-3230223320030311-0130211020130023-3230133120112331-2233132120332220-1210032131313111-3130022130010331)
- [peers.external.family_inet](resources--bgp--reference--group-001.md#canonical-0113212100231110-0013013110022232-3211013200330312-2210023313122002-3331020102102121-3233203020233100-2133131030203003-1011022003130203)
- [peers.external.family_inet.enable](resources--bgp--reference--group-001.md#canonical-0021013112222321-0231001111030131-2100032231101313-2311200201001031-0223133202020303-2331013121230332-1330210233302220-1300301211203301)
- [peers.external.family_inet.enable.aggregation](resources--bgp--reference--group-001.md#canonical-3121211121012033-0303111220033220-1121211302201112-0312011131300230-0033333300331030-2021131003100323-3030310320023320-2001121333332313)
- peers.external.family_inet.enable.aggregation.options

<a id="canonical-2322030231220312-3322333223221211-2030222101221100-1101320102130032-3121310200032313-2223201110311013-3030212111233111-3201221032113320"></a>

Type: `"object"`. list nested block, Optional.

Aggregation OPTIONS. Configuration parameter for options

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
options {
  # Configure direct properties listed below.
}
```

<a id="canonical-2321103301023022-3311233132020303-0310332130320111-1121131213201122-1112123103320001-1012121111030123-2032003103331132-3230023303000200"></a>

### Direct properties for `peers.external.family_inet.enable.aggregation.options`

- [summary_only](resources--bgp--reference--group-001.md#canonical-2131301300010020-3333123330223310-2121203323032233-1233320303222232-2003113320200011-2201320112301012-2002121303102130-1131232211003200): complete subsection reference.

<a id="canonical-2131301300010020-3333123330223310-2121203323032233-1233320303222232-2003113320200011-2201320112301012-2002121303102130-1131232211003200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external.family_inet.enable.aggregation.options.summary_only` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [peers](resources--bgp--reference--group-001.md#canonical-3132002313303222-3210021222211033-2013001002330232-0123121112132233-0321012202000010-0112130200010312-2201100311111101-1320000132020103)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0001200002001233-2222102330023121-3230223320030311-0130211020130023-3230133120112331-2233132120332220-1210032131313111-3130022130010331)
- [peers.external.family_inet](resources--bgp--reference--group-001.md#canonical-0113212100231110-0013013110022232-3211013200330312-2210023313122002-3331020102102121-3233203020233100-2133131030203003-1011022003130203)
- [peers.external.family_inet.enable](resources--bgp--reference--group-001.md#canonical-0021013112222321-0231001111030131-2100032231101313-2311200201001031-0223133202020303-2331013121230332-1330210233302220-1300301211203301)
- [peers.external.family_inet.enable.aggregation](resources--bgp--reference--group-001.md#canonical-3121211121012033-0303111220033220-1121211302201112-0312011131300230-0033333300331030-2021131003100323-3030310320023320-2001121333332313)
- [peers.external.family_inet.enable.aggregation.options](resources--bgp--reference--group-001.md#canonical-3220221002330213-2220220311222320-2232032232101021-1011011200301021-2210323313020322-3322000210321201-3301033001112220-1330001221223233)
- peers.external.family_inet.enable.aggregation.options.summary_only

<a id="canonical-2323111200023033-2110321033101002-2000320232010310-2000032100333300-1212200301132133-1311310212022100-1332113323022120-2323011313220003"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for summary only.

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
summary_only {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0132302311122112-1301303020201320-1220203233112002-2013031213311202-0011202223333303-1001200120131221-2131113223230003-0200221211302222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external.from_site` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [peers](resources--bgp--reference--group-001.md#canonical-3132002313303222-3210021222211033-2013001002330232-0123121112132233-0321012202000010-0112130200010312-2201100311111101-1320000132020103)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0001200002001233-2222102330023121-3230223320030311-0130211020130023-3230133120112331-2233132120332220-1210032131313111-3130022130010331)
- peers.external.from_site

<a id="canonical-3332010223212103-0030123011103022-0110022131033110-3213201312113132-3123232310332122-3023232331021102-2020010333012100-2122120031003233"></a>

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
from_site = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0131210123203122-1210010133133000-2321323123132230-0322210212322222-1110003222312110-0010112302302231-1122222133212210-3100031020223332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external.from_site_v6` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [peers](resources--bgp--reference--group-001.md#canonical-3132002313303222-3210021222211033-2013001002330232-0123121112132233-0321012202000010-0112130200010312-2201100311111101-1320000132020103)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0001200002001233-2222102330023121-3230223320030311-0130211020130023-3230133120112331-2233132120332220-1210032131313111-3130022130010331)
- peers.external.from_site_v6

<a id="canonical-3301000212131021-0020301130303122-3333002302233022-0223001012232123-1023112200020000-0020300201113002-2001020201010301-2021230011300113"></a>

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
from_site_v6 = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3223001231322213-1010111113010202-0212323310131132-2201030020233130-0033301131001132-3323121211013230-1333032000321221-2121121010113321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `peers.external.interface` properties

Breadcrumbs:

- [xcsh_bgp](../resources/bgp.md#canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200)
- [Property reference](resources--bgp--reference--group-001.md#canonical-0222022111032003-2132031113023131-2122301210220323-1120011021332102-1232011202301021-3013311310110100-1303130223213123-1323030111303030)
- [peers](resources--bgp--reference--group-001.md#canonical-3132002313303222-3210021222211033-2013001002330232-0123121112132233-0321012202000010-0112130200010312-2201100311111101-1320000132020103)
- [peers.external](resources--bgp--reference--group-001.md#canonical-0001200002001233-2222102330023121-3230223320030311-0130211020130023-3230133120112331-2233132120332220-1210032131313111-3130022130010331)
- peers.external.interface

<a id="canonical-2222321312303320-1001112032003002-3203131103020213-3130222312330210-3000220201012220-1103232130101301-3130312121110220-0212100312010022"></a>

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
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-3121030302123302-3300332332212001-2201230201201301-0112200201100022-1303132030030101-2310320120032133-3120131000213100-3132130133331201"></a>

### Direct properties for `peers.external.interface`

<a id="canonical-3010231102122333-1012030103021130-0212313232322312-1330232221233101-3112221211330030-2033011001002133-2321103011022003-2123100032130211"></a>

#### `peers.external.interface.name` property

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

<a id="canonical-3033232232020333-0212212032023120-0333311313113002-2033131102213332-3231030102321210-2322221012301221-1303323313000302-0013300111112301"></a>

<a id="canonical-2021021201330011-3000320003023103-3330312032132010-3033130313231032-3013013100320133-0310330313320100-1302200123023303-3000022202330321"></a>

#### `peers.external.interface.namespace` property

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

<a id="canonical-0012100303331212-0003022311002012-0000300000302100-1003330001212120-0201003300301321-2223001313331322-3230203002120030-0300331230121101"></a>

<a id="canonical-0311302202110332-1233030303033103-3031322322231213-2211000303031002-2301221102001003-2013123012032102-2100212133020103-1213101101130001"></a>

#### `peers.external.interface.tenant` property

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
