---
page_title: "xcsh_external_connector reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_external_connector reference."
---

# xcsh_external_connector reference

<a id="canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- Property reference

<a id="canonical-3330123223200332-0030230123310032-0123002133301002-0002201331220021-1133333021302011-2111212331332220-0130010210312111-2112103120111121"></a>

### Direct properties for `xcsh_external_connector`

<a id="canonical-2332011111031230-2112213001122323-1110232300100302-3022330022203110-2022113130032122-1232100232210133-3123101300032001-2123333001010321"></a>

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

- [ce_site_reference](resources--external_connector--reference--group-001.md#canonical-2311120011233120-0122300223313033-3100003032220323-0112113103013023-2003110312112000-3233013001302220-1322131223200322-1113210121122322): complete subsection reference.

<a id="canonical-3000102311300333-3311333030120120-0303022010333221-2101132312211212-2313003110030320-2200310203121111-2021022102202133-1302033022032000"></a>

<a id="canonical-0300303133301211-3110133330121312-1223232321011031-1230303121233330-3211302010020102-0300211122333202-2230022131002312-1012003212011121"></a>

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

<a id="canonical-1021322132023030-1103103302103001-0003313220233330-1011033030310220-3123101322223231-0032133021211202-3231321120300221-2032002133232223"></a>

<a id="canonical-3103203031222020-1300003210321000-3112201011022220-1233100132102013-2132201230111030-3231112002221300-1320100321023312-3210201301322203"></a>

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

- [gre](resources--external_connector--reference--group-001.md#canonical-1033000200113321-0321330103312322-3230321011131113-2322301213310303-2112331120223022-0212022112222302-3112221022312333-0111121133033131): complete subsection reference.

<a id="canonical-1111200123221223-3210101100331121-3300121222102101-3213221312311313-1110022303300101-1031233121012130-2230223233001031-0001220021221201"></a>

<a id="canonical-0202311303130011-0203033131323103-0210230033321310-0202302021210211-3313221101120313-0233331010203002-1132011002313111-0022110122003012"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ipsec](resources--external_connector--reference--group-001.md#canonical-2222210331120321-3310203330131102-1332111323102002-1300213031003231-2312210013333323-3221022213020223-3312333001311121-0222322020131133): complete subsection reference.

<a id="canonical-1232330010000111-0323122223021011-3310131102013121-2301113220201222-3322123322313111-0322120100212303-0111202003203103-3212302033133101"></a>

<a id="canonical-3103233331120131-0320303133102123-1001232310133310-3302003000210000-3121022101332212-3112213323112103-1120112003000212-3322330203033331"></a>

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

<a id="canonical-2210113103323213-1002011320212000-3212110010311232-1213032232121003-3231301333333211-1231232311333133-0023021200022022-2312031002203122"></a>

<a id="canonical-3023130012311032-3121310300231211-0010331323301033-1103211203313112-3030323030001033-2213231021211313-0121311120233013-3301212113320220"></a>

#### `name` property

Type: `"string"`. Required.

Name of the External Connector. Must be unique within the namespace.

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

<a id="canonical-1033313020131232-3301101133221323-0011012233001100-3330213100123031-0011321312001220-1023213103100332-3021321132231333-1110111201131200"></a>

<a id="canonical-1012223133113201-0222022321022302-0003023323223311-1001210211332211-3312312101000120-2330023321230232-0121303013332110-1201231000332321"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the External Connector is created.

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

- [timeouts](resources--external_connector--reference--group-001.md#canonical-2023312120203110-1221003013330002-1112111012032321-3130202022101320-2101222231102210-3302313300130100-1023233130323102-1003201233323103): complete subsection reference.

<a id="canonical-0133333321110222-2311001303302133-0101020011111001-2021011230131201-3002212303222131-3112002012220011-0202212203003000-1223220010030300"></a>

### All schema paths for `xcsh_external_connector`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--external_connector--reference--group-001.md#canonical-2332011111031230-2112213001122323-1110232300100302-3022330022203110-2022113130032122-1232100232210133-3123101300032001-2123333001010321) |
| `ce_site_reference` | [ce_site_reference](resources--external_connector--reference--group-001.md#canonical-2023200113302113-0133031010221113-2010011122320103-2311320112100210-3213202322310300-1031003133331132-0300221002312310-3331131003312021) |
| `ce_site_reference.name` | [ce_site_reference.name](resources--external_connector--reference--group-001.md#canonical-2233331332333321-1310210311230021-1023311131010031-3202333020311312-3222201201220202-0201223300032121-3122102232230112-2033223002020201) |
| `ce_site_reference.namespace` | [ce_site_reference.namespace](resources--external_connector--reference--group-001.md#canonical-2201023211311322-2012020210032201-0322303321132120-0322330330020220-0130013112033213-3012231312233113-0011000332231321-1123232303312121) |
| `ce_site_reference.tenant` | [ce_site_reference.tenant](resources--external_connector--reference--group-001.md#canonical-3233130000030003-1322110032022301-3022311011112303-3132223320122022-0021222000032201-3103023321101103-2202301102332221-2011103131303220) |
| `description` | [description](resources--external_connector--reference--group-001.md#canonical-3000102311300333-3311333030120120-0303022010333221-2101132312211212-2313003110030320-2200310203121111-2021022102202133-1302033022032000) |
| `disable` | [disable](resources--external_connector--reference--group-001.md#canonical-1021322132023030-1103103302103001-0003313220233330-1011033030310220-3123101322223231-0032133021211202-3231321120300221-2032002133232223) |
| `gre` | [gre](resources--external_connector--reference--group-001.md#canonical-3230112310011220-1330123331311020-2021203202103110-2010120202330022-1000020230310321-0213333200210211-0331233322212202-1123010003003002) |
| `gre.gre_parameters` | [gre.gre_parameters](resources--external_connector--reference--group-001.md#canonical-1210321232330033-0330021132131133-2033210213322213-1231330013332121-3023113110032313-3200330002123320-3231131110002120-0210100222032031) |
| `gre.gre_parameters.peer_ip_address` | [gre.gre_parameters.peer_ip_address](resources--external_connector--reference--group-001.md#canonical-3210311303130320-1033222023002330-3311222201102223-1211102330003012-1113121023202022-0201100303302113-0002332233121230-1321322333323232) |
| `gre.gre_parameters.peer_ip_address.addr` | [gre.gre_parameters.peer_ip_address.addr](resources--external_connector--reference--group-001.md#canonical-0301013121313211-3210233013320021-2200020110210023-3203033120111113-2001202111103123-2033332010231021-0113021231213123-3203122230322001) |
| `gre.gre_parameters.segment` | [gre.gre_parameters.segment](resources--external_connector--reference--group-001.md#canonical-0332023221112211-0323212223021011-1313101120122120-0101330020321202-2321002321333111-0221202330302212-3121220332213210-1030301222230212) |
| `gre.gre_parameters.segment.refs` | [gre.gre_parameters.segment.refs](resources--external_connector--reference--group-001.md#canonical-0010000300323033-0022032013221130-2322222221232311-1122221133112211-2232203310322130-0022202131221101-1213013322200010-3101110031220303) |
| `gre.gre_parameters.segment.refs.kind` | [gre.gre_parameters.segment.refs.kind](resources--external_connector--reference--group-001.md#canonical-0103030303021000-2101310322333112-2001213232000230-1301230032233020-2233213322102112-2221322312013130-1301112310301331-1211231120202222) |
| `gre.gre_parameters.segment.refs.name` | [gre.gre_parameters.segment.refs.name](resources--external_connector--reference--group-001.md#canonical-3010321133130223-3021232131012221-2122331313200321-2230130221232222-1021200121302212-2231302110032220-0130021100002313-0221000310313131) |
| `gre.gre_parameters.segment.refs.namespace` | [gre.gre_parameters.segment.refs.namespace](resources--external_connector--reference--group-001.md#canonical-0021033022203200-3031022110113223-2130032130123112-2112133313300032-1302203200313122-2131221121001221-2212231022103110-3133011212020131) |
| `gre.gre_parameters.segment.refs.tenant` | [gre.gre_parameters.segment.refs.tenant](resources--external_connector--reference--group-001.md#canonical-0230210033030203-1032233320203100-1300313323120312-3133301130331121-0100102211021123-3013012123020230-2100312031303200-3112130110230201) |
| `gre.gre_parameters.segment.refs.uid` | [gre.gre_parameters.segment.refs.uid](resources--external_connector--reference--group-001.md#canonical-0212230033000330-1032001321300001-2030010303332330-3323011231302303-1322011120132303-2322001130023322-3212222033123003-3301201022300000) |
| `gre.gre_parameters.site_local_inside_network` | [gre.gre_parameters.site_local_inside_network](resources--external_connector--reference--group-001.md#canonical-1103230332203123-0220303032111032-0132031212302223-3222300011102222-1030103323233303-0120131232320132-3010002021201310-2133102133202332) |
| `gre.gre_parameters.site_local_network` | [gre.gre_parameters.site_local_network](resources--external_connector--reference--group-001.md#canonical-1332122232202101-0233101121131033-2201202133123132-3300100012200221-0003312120200130-3303102021100000-0200200112102320-1100310232310101) |
| `gre.gre_parameters.tunnel_eps` | [gre.gre_parameters.tunnel_eps](resources--external_connector--reference--group-001.md#canonical-1101102200220002-3003001312330223-3230320201310223-0210223032220330-2230030003321033-0222010222110112-2330113303110212-0233333203202110) |
| `gre.gre_parameters.tunnel_eps.interface` | [gre.gre_parameters.tunnel_eps.interface](resources--external_connector--reference--group-001.md#canonical-0110323212230223-3131000323232030-1200333302101220-3123210101332222-3022232223033130-1110130030222212-3010120331200032-0333022330131200) |
| `gre.gre_parameters.tunnel_eps.local_tunnel_ip` | [gre.gre_parameters.tunnel_eps.local_tunnel_ip](resources--external_connector--reference--group-001.md#canonical-3232202312310122-0032013211100111-2003113212012132-0110310333331210-2102100320033220-2112101202120113-1230320303120021-2001123132012010) |
| `gre.gre_parameters.tunnel_eps.node` | [gre.gre_parameters.tunnel_eps.node](resources--external_connector--reference--group-001.md#canonical-3321202131232111-1220301321113311-1233103003120123-3313013023103213-0133011323310322-0000323330003103-2301211202130031-2302232103130002) |
| `gre.gre_parameters.tunnel_eps.remote_tunnel_ip` | [gre.gre_parameters.tunnel_eps.remote_tunnel_ip](resources--external_connector--reference--group-001.md#canonical-0220233113010313-0331002212223130-1013232232002103-2213010133012031-3131122331032100-0312220310033113-3013100201032121-3331303022312101) |
| `gre.gre_parameters.tunnel_mtu` | [gre.gre_parameters.tunnel_mtu](resources--external_connector--reference--group-001.md#canonical-0333103301230323-2231112012002110-3000330123332211-2033320313103200-0220221011312301-2210223203222002-1230300211121210-0231310000033300) |
| `id` | [ID](resources--external_connector--reference--group-001.md#canonical-1111200123221223-3210101100331121-3300121222102101-3213221312311313-1110022303300101-1031233121012130-2230223233001031-0001220021221201) |
| `ipsec` | [ipsec](resources--external_connector--reference--group-001.md#canonical-3133300020202102-0113220121231003-3300111103230201-1030001130313133-3330021030100002-0231111303222310-1311121032213320-3013230330021310) |
| `ipsec.ike_parameters` | [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-0330120232320022-3220123130123003-1313222120000100-1332132211221112-0300032323111010-1010000212110312-2230023033132301-0021031322133201) |
| `ipsec.ike_parameters.dpd_disabled` | [ipsec.ike_parameters.dpd_disabled](resources--external_connector--reference--group-001.md#canonical-0000322123033030-0113323230301010-0001212300113300-0001313320201231-3233231313301130-2300000210320313-3312211220333310-1303110210003133) |
| `ipsec.ike_parameters.dpd_keep_alive_timer` | [ipsec.ike_parameters.dpd_keep_alive_timer](resources--external_connector--reference--group-001.md#canonical-1100302121330002-1303023203330132-0312112131003212-1133231102030011-2001332100232003-2100123230133030-0313331330111313-0100121130012133) |
| `ipsec.ike_parameters.dpd_keep_alive_timer.timeout` | [ipsec.ike_parameters.dpd_keep_alive_timer.timeout](resources--external_connector--reference--group-001.md#canonical-3110222211130003-0202233200323000-2032230303030213-2213022112033122-1032102110132212-3001220200333001-2101332333123133-0130211130103132) |
| `ipsec.ike_parameters.ike_phase1_profile` | [ipsec.ike_parameters.ike_phase1_profile](resources--external_connector--reference--group-001.md#canonical-3303330003011130-2202110133200021-3100201211211313-0201113213012203-0221302313012032-2220230332333331-0022003010212133-1212221131332110) |
| `ipsec.ike_parameters.ike_phase1_profile.name` | [ipsec.ike_parameters.ike_phase1_profile.name](resources--external_connector--reference--group-001.md#canonical-1201220200032212-0221113032020022-3313030222133332-0300120230130322-1220031233010320-0013322032302231-2313231332011331-3210223213001310) |
| `ipsec.ike_parameters.ike_phase1_profile.namespace` | [ipsec.ike_parameters.ike_phase1_profile.namespace](resources--external_connector--reference--group-001.md#canonical-1300221203003220-0230321020330132-0101100321232331-0301030013032121-0321222111133331-1001033231200000-3021100013201201-1010030211332310) |
| `ipsec.ike_parameters.ike_phase1_profile.tenant` | [ipsec.ike_parameters.ike_phase1_profile.tenant](resources--external_connector--reference--group-001.md#canonical-0123110221231022-1303313320002212-2013130102123330-0313223330020333-0112203203201312-1132103031133101-1022200310332033-1303100200311003) |
| `ipsec.ike_parameters.ike_phase2_profile` | [ipsec.ike_parameters.ike_phase2_profile](resources--external_connector--reference--group-001.md#canonical-0321132300310133-2111132121003130-1111231210330313-2030131013021213-2233003332333002-2230221000211032-3002202312213230-1200000121130130) |
| `ipsec.ike_parameters.ike_phase2_profile.name` | [ipsec.ike_parameters.ike_phase2_profile.name](resources--external_connector--reference--group-001.md#canonical-3323212131322032-1020010321310133-3313122100220332-2101212033003332-3133032112100302-2212121330013100-1320001130330320-3330312211111201) |
| `ipsec.ike_parameters.ike_phase2_profile.namespace` | [ipsec.ike_parameters.ike_phase2_profile.namespace](resources--external_connector--reference--group-001.md#canonical-0331233113033213-2322000002100130-3111012022321000-0121003301021113-2212103221301220-2030212232022101-0231221001110000-0122110032131100) |
| `ipsec.ike_parameters.ike_phase2_profile.tenant` | [ipsec.ike_parameters.ike_phase2_profile.tenant](resources--external_connector--reference--group-001.md#canonical-1101003003020023-1223222030301021-1123003221130323-0233122300121320-1001320311213220-1223221223030101-2232220100021123-1213002111322311) |
| `ipsec.ike_parameters.initiator` | [ipsec.ike_parameters.initiator](resources--external_connector--reference--group-001.md#canonical-1113101311333010-1200221020201110-1200201230313103-0112320202133013-1322303202033133-0030230220103103-3302220231313112-3122321210313301) |
| `ipsec.ike_parameters.responder` | [ipsec.ike_parameters.responder](resources--external_connector--reference--group-001.md#canonical-3313300310132211-1202231013233333-0113311323020313-1022023032200113-0032210030022320-2033220113200311-1132001102331121-2112221310023023) |
| `ipsec.ike_parameters.rm_hostname` | [ipsec.ike_parameters.rm_hostname](resources--external_connector--reference--group-001.md#canonical-2112202111231303-1200210111301003-1003122021212201-2333120113310230-3130310122301233-0332311332020001-1322201030330030-0213011332031230) |
| `ipsec.ike_parameters.rm_ip_address` | [ipsec.ike_parameters.rm_ip_address](resources--external_connector--reference--group-001.md#canonical-0230313113122031-1320313221013222-3122120212302303-2010201030030012-0321213320220022-2003323031212131-0212110223003301-0310103001320220) |
| `ipsec.ike_parameters.rm_ip_address.dual_stack` | [ipsec.ike_parameters.rm_ip_address.dual_stack](resources--external_connector--reference--group-001.md#canonical-3031310213111222-3230210113100312-0001333211101220-1313132133231321-1211313202231231-0303031022322120-3100013032220331-2110300231023033) |
| `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4` | [ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4](resources--external_connector--reference--group-001.md#canonical-1303331322122110-2120230300203133-2031003123211200-2130100211120300-2201123321030121-3332333020022000-1000102331133223-1032311030110310) |
| `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4.addr` | [ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4.addr](resources--external_connector--reference--group-001.md#canonical-0121331031211032-3332010003301210-1032333113121321-1233121203303200-3322312201022210-0001332102032012-3121223022312221-2032013231101221) |
| `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6` | [ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6](resources--external_connector--reference--group-001.md#canonical-3133123131002330-2303232132022223-1221020003012131-2113112200103010-0113211030111010-1222122311033133-3133322002132120-1320310131130322) |
| `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6.addr` | [ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6.addr](resources--external_connector--reference--group-001.md#canonical-2333022221332003-3313322321011301-3213330220101133-3033201030003232-1302103131230003-1022131320310302-0133302120221232-0320301312112310) |
| `ipsec.ike_parameters.rm_ip_address.ipv4` | [ipsec.ike_parameters.rm_ip_address.ipv4](resources--external_connector--reference--group-001.md#canonical-2201011100230130-1231303231023313-1321003333332110-3323132121100222-1003031333011312-3313213323030112-2202232311313111-0021100021303133) |
| `ipsec.ike_parameters.rm_ip_address.ipv4.addr` | [ipsec.ike_parameters.rm_ip_address.ipv4.addr](resources--external_connector--reference--group-001.md#canonical-3222212023201101-0232322110211022-3000123021321123-2020101001000331-1320121013233212-1120100023230233-0200321230213131-1022233232012112) |
| `ipsec.ike_parameters.rm_ip_address.ipv6` | [ipsec.ike_parameters.rm_ip_address.ipv6](resources--external_connector--reference--group-001.md#canonical-2101333132333030-0311131312013313-1303303022313123-3003322320132031-3011230123013101-0203232120303202-3032120003022333-0201032203312022) |
| `ipsec.ike_parameters.rm_ip_address.ipv6.addr` | [ipsec.ike_parameters.rm_ip_address.ipv6.addr](resources--external_connector--reference--group-001.md#canonical-1333033130231231-2231212230111120-0120102122323211-2320322103120003-3111231332221113-2103212000113122-1033223112313210-2332101001310203) |
| `ipsec.ike_parameters.use_default_local_ike_id` | [ipsec.ike_parameters.use_default_local_ike_id](resources--external_connector--reference--group-001.md#canonical-3002301301123320-0213210132211202-1313310322310322-2112331103301101-1001330331233213-0212133100032301-1030322020221220-0020113302111212) |
| `ipsec.ike_parameters.use_default_remote_ike_id` | [ipsec.ike_parameters.use_default_remote_ike_id](resources--external_connector--reference--group-001.md#canonical-1013200302312030-0302213022310213-1132120312001120-2212022300003010-0332020102310233-2022122210230032-1321331032321011-1133213010231301) |
| `ipsec.ipsec_tunnel_parameters` | [ipsec.ipsec_tunnel_parameters](resources--external_connector--reference--group-001.md#canonical-2000211302111030-1002232222001013-2322020300211012-1113000032010330-1212331330001100-3330133121121202-1323120203220023-0130021223020223) |
| `ipsec.ipsec_tunnel_parameters.peer_ip_address` | [ipsec.ipsec_tunnel_parameters.peer_ip_address](resources--external_connector--reference--group-001.md#canonical-1222202100122001-1132303203012323-2132202203003233-1002012203001302-3333010203032123-0221331113312103-2110112231212130-1022120103220010) |
| `ipsec.ipsec_tunnel_parameters.peer_ip_address.addr` | [ipsec.ipsec_tunnel_parameters.peer_ip_address.addr](resources--external_connector--reference--group-001.md#canonical-3132311233002322-2032111302322130-1133310233103323-3022131333121130-2203303100333130-3320310121133221-2131333111233220-3103310103003223) |
| `ipsec.ipsec_tunnel_parameters.psk` | [ipsec.ipsec_tunnel_parameters.psk](resources--external_connector--reference--group-001.md#canonical-2130331102103110-1323031001132332-2021230011131113-0122230123113023-3300331221121200-1301301213230231-0211221322202331-0121232303132003) |
| `ipsec.ipsec_tunnel_parameters.segment` | [ipsec.ipsec_tunnel_parameters.segment](resources--external_connector--reference--group-001.md#canonical-3220203300230031-3111233233231131-0223010232032001-0121232120211000-2312030023112231-0002023121120213-1321022100130331-0320123203032103) |
| `ipsec.ipsec_tunnel_parameters.segment.refs` | [ipsec.ipsec_tunnel_parameters.segment.refs](resources--external_connector--reference--group-001.md#canonical-1111112212330121-0310001020031230-2331312023101001-2322110130302102-0212130303002133-1333300301113001-1133130130100000-2331321332110112) |
| `ipsec.ipsec_tunnel_parameters.segment.refs.kind` | [ipsec.ipsec_tunnel_parameters.segment.refs.kind](resources--external_connector--reference--group-001.md#canonical-1002102112031002-3311023333132113-3111013103222300-2230033301213023-0322122002322122-3102100130131123-0231332110102032-2322201020333000) |
| `ipsec.ipsec_tunnel_parameters.segment.refs.name` | [ipsec.ipsec_tunnel_parameters.segment.refs.name](resources--external_connector--reference--group-001.md#canonical-3210202131322220-2023002213321012-2000003301210202-2233010311112321-1000100111102213-0212102310130112-2303203032320130-1011032311232213) |
| `ipsec.ipsec_tunnel_parameters.segment.refs.namespace` | [ipsec.ipsec_tunnel_parameters.segment.refs.namespace](resources--external_connector--reference--group-001.md#canonical-3003102303302012-1112312013102011-1010132223013202-2233300032112320-1020100112312020-3230213321130331-0310100200201210-3300031121111320) |
| `ipsec.ipsec_tunnel_parameters.segment.refs.tenant` | [ipsec.ipsec_tunnel_parameters.segment.refs.tenant](resources--external_connector--reference--group-001.md#canonical-1320110112123002-1302303013331001-2133313210231001-3011322130312013-0132022023322223-0013223331331202-2002210211132011-2312220213132022) |
| `ipsec.ipsec_tunnel_parameters.segment.refs.uid` | [ipsec.ipsec_tunnel_parameters.segment.refs.uid](resources--external_connector--reference--group-001.md#canonical-1313003130002031-1201322331012123-3203202211302120-3221212300210032-0313022122213012-0102130212230102-2232113022213030-0313303120030112) |
| `ipsec.ipsec_tunnel_parameters.site_local_inside_network` | [ipsec.ipsec_tunnel_parameters.site_local_inside_network](resources--external_connector--reference--group-001.md#canonical-2032233211000222-0300201322231100-2233210323233132-3312010033333132-0211233030321102-3230100120311130-2022320330302003-1311120013021211) |
| `ipsec.ipsec_tunnel_parameters.site_local_network` | [ipsec.ipsec_tunnel_parameters.site_local_network](resources--external_connector--reference--group-001.md#canonical-1001132321202233-2221021003321323-0221222100122102-0033122023112010-2110211330321332-2303112021221200-1313020331123220-3230203120323002) |
| `ipsec.ipsec_tunnel_parameters.tunnel_eps` | [ipsec.ipsec_tunnel_parameters.tunnel_eps](resources--external_connector--reference--group-001.md#canonical-2102102030212130-0320210230212102-0133012021233111-2102201211311201-2030223202331330-1332322103311310-0100311202022101-1211110310101113) |
| `ipsec.ipsec_tunnel_parameters.tunnel_eps.interface` | [ipsec.ipsec_tunnel_parameters.tunnel_eps.interface](resources--external_connector--reference--group-001.md#canonical-1121323222313320-0003203130001233-3120313313220101-2113110033021133-2223002110033323-2232132113313010-1011333001321113-1012301112320132) |
| `ipsec.ipsec_tunnel_parameters.tunnel_eps.local_tunnel_ip` | [ipsec.ipsec_tunnel_parameters.tunnel_eps.local_tunnel_ip](resources--external_connector--reference--group-001.md#canonical-1210311120321111-0011102103000131-1030333333113133-2032322021020233-3213313032233101-1002323001030313-0322010300202311-0210130232300130) |
| `ipsec.ipsec_tunnel_parameters.tunnel_eps.node` | [ipsec.ipsec_tunnel_parameters.tunnel_eps.node](resources--external_connector--reference--group-001.md#canonical-3100333331302313-3110213313220212-1103310001323223-3100203111022012-0323121000311303-0133223101122201-2302121221210031-3132322301321003) |
| `ipsec.ipsec_tunnel_parameters.tunnel_eps.remote_tunnel_ip` | [ipsec.ipsec_tunnel_parameters.tunnel_eps.remote_tunnel_ip](resources--external_connector--reference--group-001.md#canonical-1031233323111312-3303132112110133-0112133103111333-2121223001010122-0123021123111231-3303331023210023-2332212002020303-1131313002322131) |
| `ipsec.ipsec_tunnel_parameters.tunnel_mtu` | [ipsec.ipsec_tunnel_parameters.tunnel_mtu](resources--external_connector--reference--group-001.md#canonical-2202330331123200-3030312012022023-2230002222221013-1122032321232010-1303223202021331-3211321112033000-3212302203131100-3021303311022013) |
| `labels` | [labels](resources--external_connector--reference--group-001.md#canonical-1232330010000111-0323122223021011-3310131102013121-2301113220201222-3322123322313111-0322120100212303-0111202003203103-3212302033133101) |
| `name` | [name](resources--external_connector--reference--group-001.md#canonical-2210113103323213-1002011320212000-3212110010311232-1213032232121003-3231301333333211-1231232311333133-0023021200022022-2312031002203122) |
| `namespace` | [namespace](resources--external_connector--reference--group-001.md#canonical-1033313020131232-3301101133221323-0011012233001100-3330213100123031-0011321312001220-1023213103100332-3021321132231333-1110111201131200) |
| `timeouts` | [timeouts](resources--external_connector--reference--group-001.md#canonical-2022231212123301-1222323121000031-3211121303202020-0102200131010310-1233332332313002-3202232231123333-1111122222031021-1121232220222221) |
| `timeouts.create` | [timeouts.create](resources--external_connector--reference--group-001.md#canonical-3301313321202123-2213120303020031-2122320122322000-3330202021130332-1011122322033300-1323003020031323-1131023321203231-0300201223020320) |
| `timeouts.delete` | [timeouts.delete](resources--external_connector--reference--group-001.md#canonical-2223320110321020-0013221113011303-3230030101312310-3123203002300302-0202322230010001-2113010122202320-3001001030101321-0331103010131202) |
| `timeouts.read` | [timeouts.read](resources--external_connector--reference--group-001.md#canonical-2012121233322212-2330321030222001-2020222111320022-1120222032010132-3220213201203333-2302123212012212-3101002331311132-1301033230032200) |
| `timeouts.update` | [timeouts.update](resources--external_connector--reference--group-001.md#canonical-2010010233330333-0101000131233121-1300122032311002-0331222313011301-2002201131220030-1002120320230222-3310330223221033-1200113033302112) |

<a id="canonical-2311120011233120-0122300223313033-3100003032220323-0112113103013023-2003110312112000-3233013001302220-1322131223200322-1113210121122322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ce_site_reference` properties

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232)
- ce_site_reference

<a id="canonical-2023200113302113-0133031010221113-2010011122320103-2311320112100210-3213202322310300-1031003133331132-0300221002312310-3331131003312021"></a>

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
ce_site_reference {
  # Configure direct properties listed below.
}
```

<a id="canonical-3131303101001032-2012222310313130-2110110331020120-3010011223312321-1223331022010131-3002021203300130-1101230211320021-1332230220002102"></a>

### Direct properties for `ce_site_reference`

<a id="canonical-2233331332333321-1310210311230021-1023311131010031-3202333020311312-3222201201220202-0201223300032121-3122102232230112-2033223002020201"></a>

#### `ce_site_reference.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2201023211311322-2012020210032201-0322303321132120-0322330330020220-0130013112033213-3012231312233113-0011000332231321-1123232303312121"></a>

<a id="canonical-0211331232002320-3003300211101120-0313321310320203-3231131012302331-1132231023031021-2321111132002031-2200200002003020-0003121301213330"></a>

#### `ce_site_reference.namespace` property

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3233130000030003-1322110032022301-3022311011112303-3132223320122022-0021222000032201-3103023321101103-2202301102332221-2011103131303220"></a>

<a id="canonical-1012032132000313-1323032320200210-1002333121321231-0222321020232233-0311321310111010-0112120100000300-2132122020233101-2002201002201221"></a>

#### `ce_site_reference.tenant` property

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1033000200113321-0321330103312322-3230321011131113-2322301213310303-2112331120223022-0212022112222302-3112221022312333-0111121133033131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gre` properties

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232)
- gre

<a id="canonical-3230112310011220-1330123331311020-2021203202103110-2010120202330022-1000020230310321-0213333200210211-0331233322212202-1123010003003002"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: gre, ipsec\] GRE. External Connector with GRE tunnel.

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

- [gre](resources--external_connector--reference--group-001.md#canonical-3230112310011220-1330123331311020-2021203202103110-2010120202330022-1000020230310321-0213333200210211-0331233322212202-1123010003003002)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-3133300020202102-0113220121231003-3300111103230201-1030001130313133-3330021030100002-0231111303222310-1311121032213320-3013230330021310)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
gre {
  # Configure direct properties listed below.
}
```

<a id="canonical-3223131310012320-0031223221031330-2221033212233121-0110210002210311-2203121103203021-2132003333331322-1310031302212232-2201232312131123"></a>

### Direct properties for `gre`

- [gre_parameters](resources--external_connector--reference--group-001.md#canonical-3101301101332121-2110312002000013-2231001111210031-0111032211133210-1203222202331320-2002302122213201-1100123102223120-3313131020301200): complete subsection reference.

<a id="canonical-3101301101332121-2110312002000013-2231001111210031-0111032211133210-1203222202331320-2002302122213201-1100123102223120-3313131020301200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gre.gre_parameters` properties

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232)
- [gre](resources--external_connector--reference--group-001.md#canonical-1033000200113321-0321330103312322-3230321011131113-2322301213310303-2112331120223022-0212022112222302-3112221022312333-0111121133033131)
- gre.gre_parameters

<a id="canonical-1210321232330033-0330021132131133-2033210213322213-1231330013332121-3023113110032313-3200330002123320-3231131110002120-0210100222032031"></a>

Type: `"object"`. single nested block, Optional.

GRE configuration parameters required for GRE Connection type.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("tunnel_eps",
    "tunnel_mtu"),
  validators.ConflictingObjectAttributes("segment",
    "site_local_inside_network"),
  validators.ConflictingObjectAttributes("segment",
    "site_local_network"),
  validators.ConflictingObjectAttributes("site_local_inside_network",
    "site_local_network")}
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
  "x-ves-oneof-field-tunnel_source_vn": "[\"segment\",\"site_local_inside_network\",\"site_local_network\"]"
}
```

Terraform syntax:

```terraform
gre_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-2232323302230103-3203301210222231-2031003012213103-3331133323132313-3022010332313221-0123103002101111-2212332221101021-2033311033231130"></a>

### Direct properties for `gre.gre_parameters`

- [peer_ip_address](resources--external_connector--reference--group-001.md#canonical-2311222021113110-1311130120233133-2310033001200202-2211000203311010-3121133233201013-0121113123033132-2301313102013302-0321312330121122): complete subsection reference.

- [segment](resources--external_connector--reference--group-001.md#canonical-3033011103100101-0331031313031322-1321223022000220-2121023221103112-1031123102133113-0310331122002210-2232000331012220-1200323030030230): complete subsection reference.

- [site_local_inside_network](resources--external_connector--reference--group-001.md#canonical-0323303202120112-0013320211312000-0133101303331021-1301221032313300-0333021113300110-2331121000200131-3332323022031321-2101203110313011): complete subsection reference.

- [site_local_network](resources--external_connector--reference--group-001.md#canonical-3122301323222230-3110131220331223-3200121221312232-2110020023332300-2222012103130123-1001102032222202-2013331101003112-2200220032212012): complete subsection reference.

- [tunnel_eps](resources--external_connector--reference--group-001.md#canonical-3010112031211331-0301021120032210-0212123101022030-3302111100331202-2220301013332022-1320010300030001-0003013201312303-1013110321001212): complete subsection reference.

<a id="canonical-0333103301230323-2231112012002110-3000330123332211-2033320313103200-0220221011312301-2210223203222002-1230300211121210-0231310000033300"></a>

<a id="canonical-2133100121100230-1101221200303233-0020021220302331-1301300200112202-3011322213221131-1221300103212030-2333123223020133-1030110222112120"></a>

#### `gre.gre_parameters.tunnel_mtu` property

Type: `"number"`. Optional.

Configure MTU for the GRE tunnel interface.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(512, 1370),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1370,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 512
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "512",
    "ves.io.schema.rules.uint32.lte": "1370"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "512",
    "ves.io.schema.rules.uint32.lte": "1370"
  }
}
```

<a id="canonical-2311222021113110-1311130120233133-2310033001200202-2211000203311010-3121133233201013-0121113123033132-2301313102013302-0321312330121122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gre.gre_parameters.peer_ip_address` properties

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232)
- [gre](resources--external_connector--reference--group-001.md#canonical-1033000200113321-0321330103312322-3230321011131113-2322301213310303-2112331120223022-0212022112222302-3112221022312333-0111121133033131)
- [gre.gre_parameters](resources--external_connector--reference--group-001.md#canonical-3101301101332121-2110312002000013-2231001111210031-0111032211133210-1203222202331320-2002302122213201-1100123102223120-3313131020301200)
- gre.gre_parameters.peer_ip_address

<a id="canonical-3210311303130320-1033222023002330-3311222201102223-1211102330003012-1113121023202022-0201100303302113-0002332233121230-1321322333323232"></a>

Type: `"object"`. single nested block, Optional.

IPv4 Address. IPv4 Address in dot-decimal notation.

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
peer_ip_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-2031003332113133-1323211131312311-1213220110233132-3311303033133331-2122213212332333-3232221311230122-3120111033110303-1223222033022233"></a>

### Direct properties for `gre.gre_parameters.peer_ip_address`

<a id="canonical-0301013121313211-3210233013320021-2200020110210023-3203033120111113-2001202111103123-2033332010231021-0113021231213123-3203122230322001"></a>

#### `gre.gre_parameters.peer_ip_address.addr` property

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

<a id="canonical-3033011103100101-0331031313031322-1321223022000220-2121023221103112-1031123102133113-0310331122002210-2232000331012220-1200323030030230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gre.gre_parameters.segment` properties

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232)
- [gre](resources--external_connector--reference--group-001.md#canonical-1033000200113321-0321330103312322-3230321011131113-2322301213310303-2112331120223022-0212022112222302-3112221022312333-0111121133033131)
- [gre.gre_parameters](resources--external_connector--reference--group-001.md#canonical-3101301101332121-2110312002000013-2231001111210031-0111032211133210-1203222202331320-2002302122213201-1100123102223120-3313131020301200)
- gre.gre_parameters.segment

<a id="canonical-0332023221112211-0323212223021011-1313101120122120-0101330020321202-2321002321333111-0221202330302212-3121220332213210-1030301222230212"></a>

Type: `"object"`. single nested block, Optional.

Segment Reference Type. Reference to Segment Object.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("refs")}
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
segment {
  # Configure direct properties listed below.
}
```

<a id="canonical-2023233031201220-2223210301323001-2323231103133102-3310221003321033-2013232101123113-0323113213022013-0012320211001122-0233130033210003"></a>

### Direct properties for `gre.gre_parameters.segment`

- [refs](resources--external_connector--reference--group-001.md#canonical-0130021011120302-0123332323033300-0101311013330013-2321022120030010-3233022010333221-0200101213313323-0203212000000121-3203303200100211): complete subsection reference.

<a id="canonical-0130021011120302-0123332323033300-0101311013330013-2321022120030010-3233022010333221-0200101213313323-0203212000000121-3203303200100211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gre.gre_parameters.segment.refs` properties

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232)
- [gre](resources--external_connector--reference--group-001.md#canonical-1033000200113321-0321330103312322-3230321011131113-2322301213310303-2112331120223022-0212022112222302-3112221022312333-0111121133033131)
- [gre.gre_parameters](resources--external_connector--reference--group-001.md#canonical-3101301101332121-2110312002000013-2231001111210031-0111032211133210-1203222202331320-2002302122213201-1100123102223120-3313131020301200)
- [gre.gre_parameters.segment](resources--external_connector--reference--group-001.md#canonical-3033011103100101-0331031313031322-1321223022000220-2121023221103112-1031123102133113-0310331122002210-2232000331012220-1200323030030230)
- gre.gre_parameters.segment.refs

<a id="canonical-0010000300323033-0022032013221130-2322222221232311-1122221133112211-2232203310322130-0022202131221101-1213013322200010-3101110031220303"></a>

Type: `"object"`. list nested block, Optional.

Segment. Reference to Segment Object.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
refs {
  # Configure direct properties listed below.
}
```

<a id="canonical-1112200131012031-1232202132323023-3301032320212312-3020230212323222-1031111332303023-3001210030002320-3033312323323303-2031320222313102"></a>

### Direct properties for `gre.gre_parameters.segment.refs`

<a id="canonical-0103030303021000-2101310322333112-2001213232000230-1301230032233020-2233213322102112-2221322312013130-1301112310301331-1211231120202222"></a>

#### `gre.gre_parameters.segment.refs.kind` property

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

<a id="canonical-3010321133130223-3021232131012221-2122331313200321-2230130221232222-1021200121302212-2231302110032220-0130021100002313-0221000310313131"></a>

<a id="canonical-1220012311021331-1110313133201000-0203300023120112-2333131003102132-3201310101201012-0333311231203220-0202123020101120-2002021033103103"></a>

#### `gre.gre_parameters.segment.refs.name` property

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

<a id="canonical-0021033022203200-3031022110113223-2130032130123112-2112133313300032-1302203200313122-2131221121001221-2212231022103110-3133011212020131"></a>

<a id="canonical-1101312102123221-0221313213011323-0102123233321211-2010302232132122-3213113133230321-2112322111223103-0001103123210130-2010131232311221"></a>

#### `gre.gre_parameters.segment.refs.namespace` property

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

<a id="canonical-0230210033030203-1032233320203100-1300313323120312-3133301130331121-0100102211021123-3013012123020230-2100312031303200-3112130110230201"></a>

<a id="canonical-1122302112002003-1321312223110213-3223012030131202-3103013320022001-2311323100230203-0011030003022011-3231111330312200-1001032122323020"></a>

#### `gre.gre_parameters.segment.refs.tenant` property

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

<a id="canonical-0212230033000330-1032001321300001-2030010303332330-3323011231302303-1322011120132303-2322001130023322-3212222033123003-3301201022300000"></a>

<a id="canonical-1320333301003321-0010023332321303-0032113011021133-1321312021212013-2233130213130201-1213320212232303-0130223000312303-3200011031111131"></a>

#### `gre.gre_parameters.segment.refs.uid` property

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

<a id="canonical-0323303202120112-0013320211312000-0133101303331021-1301221032313300-0333021113300110-2331121000200131-3332323022031321-2101203110313011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gre.gre_parameters.site_local_inside_network` properties

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232)
- [gre](resources--external_connector--reference--group-001.md#canonical-1033000200113321-0321330103312322-3230321011131113-2322301213310303-2112331120223022-0212022112222302-3112221022312333-0111121133033131)
- [gre.gre_parameters](resources--external_connector--reference--group-001.md#canonical-3101301101332121-2110312002000013-2231001111210031-0111032211133210-1203222202331320-2002302122213201-1100123102223120-3313131020301200)
- gre.gre_parameters.site_local_inside_network

<a id="canonical-1103230332203123-0220303032111032-0132031212302223-3222300011102222-1030103323233303-0120131232320132-3010002021201310-2133102133202332"></a>

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
site_local_inside_network = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3122301323222230-3110131220331223-3200121221312232-2110020023332300-2222012103130123-1001102032222202-2013331101003112-2200220032212012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gre.gre_parameters.site_local_network` properties

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232)
- [gre](resources--external_connector--reference--group-001.md#canonical-1033000200113321-0321330103312322-3230321011131113-2322301213310303-2112331120223022-0212022112222302-3112221022312333-0111121133033131)
- [gre.gre_parameters](resources--external_connector--reference--group-001.md#canonical-3101301101332121-2110312002000013-2231001111210031-0111032211133210-1203222202331320-2002302122213201-1100123102223120-3313131020301200)
- gre.gre_parameters.site_local_network

<a id="canonical-1332122232202101-0233101121131033-2201202133123132-3300100012200221-0003312120200130-3303102021100000-0200200112102320-1100310232310101"></a>

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
site_local_network = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3010112031211331-0301021120032210-0212123101022030-3302111100331202-2220301013332022-1320010300030001-0003013201312303-1013110321001212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gre.gre_parameters.tunnel_eps` properties

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232)
- [gre](resources--external_connector--reference--group-001.md#canonical-1033000200113321-0321330103312322-3230321011131113-2322301213310303-2112331120223022-0212022112222302-3112221022312333-0111121133033131)
- [gre.gre_parameters](resources--external_connector--reference--group-001.md#canonical-3101301101332121-2110312002000013-2231001111210031-0111032211133210-1203222202331320-2002302122213201-1100123102223120-3313131020301200)
- gre.gre_parameters.tunnel_eps

<a id="canonical-1101102200220002-3003001312330223-3230320201310223-0210223032220330-2230030003321033-0222010222110112-2330113303110212-0233333203202110"></a>

Type: `"object"`. list nested block, Optional.

Configure tunnel parameters, source, destination, IP addresses.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("interface",
    "local_tunnel_ip",
    "node",
    "remote_tunnel_ip")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
tunnel_eps {
  # Configure direct properties listed below.
}
```

<a id="canonical-2212003320010302-2231333001101103-0332212213031013-2221332202023211-3212210313011111-3020021002033011-3131200003203020-2012020333031213"></a>

### Direct properties for `gre.gre_parameters.tunnel_eps`

<a id="canonical-0110323212230223-3131000323232030-1200333302101220-3123210101332222-3022232223033130-1110130030222212-3010120331200032-0333022330131200"></a>

#### `gre.gre_parameters.tunnel_eps.interface` property

Type: `"string"`. Optional.

For the chosen node, specify the interface that will be the tunnel source.

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

<a id="canonical-3232202312310122-0032013211100111-2003113212012132-0110310333331210-2102100320033220-2112101202120113-1230320303120021-2001123132012010"></a>

<a id="canonical-0100221211120101-1133303310022233-0133200230020020-1313130320233102-0301033220330310-3213032002110131-0112222111220020-0323300020222001"></a>

#### `gre.gre_parameters.tunnel_eps.local_tunnel_ip` property

Type: `"string"`. Optional.

For a particular tunnel on a node, specify the local tunnel IP Address i.e. The IP address of the
tunnel on the CE node itself and a subnet prefix length.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-3321202131232111-1220301321113311-1233103003120123-3313013023103213-0133011323310322-0000323330003103-2301211202130031-2302232103130002"></a>

<a id="canonical-1332312330112121-0333001111311221-0310132311032210-2213031130321310-1130101031312233-0032103232300231-1200312100022123-1003202120002320"></a>

#### `gre.gre_parameters.tunnel_eps.node` property

Type: `"string"`. Optional.

A CE site is composed of multiple nodes. Choose a node that will be part of this external
connection.

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

<a id="canonical-0220233113010313-0331002212223130-1013232232002103-2213010133012031-3131122331032100-0312220310033113-3013100201032121-3331303022312101"></a>

<a id="canonical-3002211200301303-3300331322023233-2222103322300020-1020211312233001-0023223220002133-3121333001220201-0223222031103211-0331230213102013"></a>

#### `gre.gre_parameters.tunnel_eps.remote_tunnel_ip` property

Type: `"string"`. Optional.

For a particular tunnel on a node, specify the remote tunnel IP Address i.e. The IP address of the
tunnel on the remote gateway and a subnet prefix length.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-2222210331120321-3310203330131102-1332111323102002-1300213031003231-2312210013333323-3221022213020223-3312333001311121-0222322020131133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec` properties

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232)
- ipsec

<a id="canonical-3133300020202102-0113220121231003-3300111103230201-1030001130313133-3330021030100002-0231111303222310-1311121032213320-3013230330021310"></a>

Type: `"object"`. single nested block, Optional.

IPsec. External Connector with IPsec tunnel.

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

<a id="canonical-0231122200330332-3000120013100320-2033223330201110-1003203020110003-0230010023222301-3201133331330221-0332122110130121-3030111030220112"></a>

### Direct properties for `ipsec`

- [ike_parameters](resources--external_connector--reference--group-001.md#canonical-2110112022302230-2103222202112010-1012130322030213-3100331333231022-3132330220301303-2012000231322123-0021133111013121-3222203023211013): complete subsection reference.

- [ipsec_tunnel_parameters](resources--external_connector--reference--group-001.md#canonical-1111312210033110-1323001002020013-0301303102011303-3323022223203100-3130333022311330-3033000102231322-1330001332110221-0323222300300120): complete subsection reference.

<a id="canonical-2110112022302230-2103222202112010-1012130322030213-3100331333231022-3132330220301303-2012000231322123-0021133111013121-3222203023211013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ike_parameters` properties

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-2222210331120321-3310203330131102-1332111323102002-1300213031003231-2312210013333323-3221022213020223-3312333001311121-0222322020131133)
- ipsec.ike_parameters

<a id="canonical-0330120232320022-3220123130123003-1313222120000100-1332132211221112-0300032323111010-1010000212110312-2230023033132301-0021031322133201"></a>

Type: `"object"`. single nested block, Optional.

IKE configuration parameters required for IPsec Connection type.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("dpd_disabled",
    "dpd_keep_alive_timer"),
  validators.ConflictingObjectAttributes("initiator",
    "responder"),
  validators.ConflictingObjectAttributes("rm_hostname",
    "rm_ip_address"),
  validators.ConflictingObjectAttributes("rm_hostname",
    "use_default_remote_ike_id"),
  validators.ConflictingObjectAttributes("rm_ip_address",
    "use_default_remote_ike_id")}
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
  "x-ves-oneof-field-dpd_choice": "[\"dpd_disabled\",\"dpd_keep_alive_timer\"]",
  "x-ves-oneof-field-local_ike_id": "[\"use_default_local_ike_id\"]",
  "x-ves-oneof-field-mode_choice": "[\"initiator\",\"responder\"]",
  "x-ves-oneof-field-remote_ike_id": "[\"rm_hostname\",\"rm_ip_address\",\"use_default_remote_ike_id\"]"
}
```

Terraform syntax:

```terraform
ike_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-2033023130120113-3111112110111000-0311002202203013-1022010220211021-1213102013103100-3202013221010223-2213132032212300-3010102230112321"></a>

### Direct properties for `ipsec.ike_parameters`

- [dpd_disabled](resources--external_connector--reference--group-001.md#canonical-1111321100210220-0111023123121121-2212313103303031-3021233311333001-0021003022332002-3313103311302312-3133311311333210-2232012003033323): complete subsection reference.

- [dpd_keep_alive_timer](resources--external_connector--reference--group-001.md#canonical-0201200323101322-3010031113332333-0210103123123123-1012200001010033-3021311222223022-1213032313113101-3130323023010031-0220302232322010): complete subsection reference.

- [ike_phase1_profile](resources--external_connector--reference--group-001.md#canonical-0211232120122301-1331132102322012-2030111102213223-2011102332130111-1111331000133003-1303123212120310-2201100121323223-1103323232220013): complete subsection reference.

- [ike_phase2_profile](resources--external_connector--reference--group-001.md#canonical-3111311323322130-0103333120230033-0312100202300122-2130222013013001-1120011132233320-1030022111132100-0112003132302101-0232131200220223): complete subsection reference.

- [initiator](resources--external_connector--reference--group-001.md#canonical-0302121201233210-2202133200011000-2021020131303223-3103331000102231-1033120300201012-3220330201011102-1112221032012003-2222313010101103): complete subsection reference.

- [responder](resources--external_connector--reference--group-001.md#canonical-0032203013123300-0333020112200010-0110312323131012-3221130002003033-2311223031332310-1331121120012120-1020133212213233-0103121320131033): complete subsection reference.

<a id="canonical-2112202111231303-1200210111301003-1003122021212201-2333120113310230-3130310122301233-0332311332020001-1322201030330030-0213011332031230"></a>

<a id="canonical-3310303123023131-0000331313033222-2012121030302323-0222202200031310-2331103002013002-3131321221233310-1233320300220003-2103223110110330"></a>

#### `ipsec.ike_parameters.rm_hostname` property

Type: `"string"`. Optional.

Exclusive with \[rm\_ip\_address use\_default\_remote\_ike\_id\] Configure an hostname Remote IKE
ID.

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

- [rm_ip_address](resources--external_connector--reference--group-001.md#canonical-1021320210230233-0300130321222101-0332113201003312-2123232100220012-3201003310002210-3322122233233202-3111300100310100-0211113122201313): complete subsection reference.

- [use_default_local_ike_id](resources--external_connector--reference--group-001.md#canonical-2203332212111231-1233232210033231-3302200313011103-1320320031132222-3020011013212211-2011021031300333-1202320013123100-1231103113130112): complete subsection reference.

- [use_default_remote_ike_id](resources--external_connector--reference--group-001.md#canonical-2332221223221312-2330120323211321-3330111123031100-3130000121101013-2201300311311122-2332203102103332-2032230023131002-0213122010103222): complete subsection reference.

<a id="canonical-1111321100210220-0111023123121121-2212313103303031-3021233311333001-0021003022332002-3313103311302312-3133311311333210-2232012003033323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ike_parameters.dpd_disabled` properties

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-2222210331120321-3310203330131102-1332111323102002-1300213031003231-2312210013333323-3221022213020223-3312333001311121-0222322020131133)
- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-2110112022302230-2103222202112010-1012130322030213-3100331333231022-3132330220301303-2012000231322123-0021133111013121-3222203023211013)
- ipsec.ike_parameters.dpd_disabled

<a id="canonical-0000322123033030-0113323230301010-0001212300113300-0001313320201231-3233231313301130-2300000210320313-3312211220333310-1303110210003133"></a>

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
dpd_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0201200323101322-3010031113332333-0210103123123123-1012200001010033-3021311222223022-1213032313113101-3130323023010031-0220302232322010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ike_parameters.dpd_keep_alive_timer` properties

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-2222210331120321-3310203330131102-1332111323102002-1300213031003231-2312210013333323-3221022213020223-3312333001311121-0222322020131133)
- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-2110112022302230-2103222202112010-1012130322030213-3100331333231022-3132330220301303-2012000231322123-0021133111013121-3222203023211013)
- ipsec.ike_parameters.dpd_keep_alive_timer

<a id="canonical-1100302121330002-1303023203330132-0312112131003212-1133231102030011-2001332100232003-2100123230133030-0313331330111313-0100121130012133"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for dpd keep alive timer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("timeout")}
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
dpd_keep_alive_timer {
  # Configure direct properties listed below.
}
```

<a id="canonical-1210212221300223-3232130311030210-0130200010323202-0223220300310312-0010000032213100-3220001131122030-0333220031320233-0221223111203022"></a>

### Direct properties for `ipsec.ike_parameters.dpd_keep_alive_timer`

<a id="canonical-3110222211130003-0202233200323000-2032230303030213-2213022112033122-1032102110132212-3001220200333001-2101332333123133-0130211130103132"></a>

#### `ipsec.ike_parameters.dpd_keep_alive_timer.timeout` property

Type: `"number"`. Optional.

Keepalive Timer. Operation timeout duration

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 5),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.uint32.lte": "5"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  }
}
```

<a id="canonical-0211232120122301-1331132102322012-2030111102213223-2011102332130111-1111331000133003-1303123212120310-2201100121323223-1103323232220013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ike_parameters.ike_phase1_profile` properties

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-2222210331120321-3310203330131102-1332111323102002-1300213031003231-2312210013333323-3221022213020223-3312333001311121-0222322020131133)
- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-2110112022302230-2103222202112010-1012130322030213-3100331333231022-3132330220301303-2012000231322123-0021133111013121-3222203023211013)
- ipsec.ike_parameters.ike_phase1_profile

<a id="canonical-3303330003011130-2202110133200021-3100201211211313-0201113213012203-0221302313012032-2220230332333331-0022003010212133-1212221131332110"></a>

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
ike_phase1_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-1013223013031110-3310312011122003-2020112020111220-0302102303213220-3232012032210211-0203211031213032-3222330330323021-3200330333303133"></a>

### Direct properties for `ipsec.ike_parameters.ike_phase1_profile`

<a id="canonical-1201220200032212-0221113032020022-3313030222133332-0300120230130322-1220031233010320-0013322032302231-2313231332011331-3210223213001310"></a>

#### `ipsec.ike_parameters.ike_phase1_profile.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1300221203003220-0230321020330132-0101100321232331-0301030013032121-0321222111133331-1001033231200000-3021100013201201-1010030211332310"></a>

<a id="canonical-0031331220321002-0012103003200013-0103330300222100-3312231130023212-0230311321311331-2103001031301002-0220112300212030-3300033013131211"></a>

#### `ipsec.ike_parameters.ike_phase1_profile.namespace` property

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0123110221231022-1303313320002212-2013130102123330-0313223330020333-0112203203201312-1132103031133101-1022200310332033-1303100200311003"></a>

<a id="canonical-1020231201200331-0013203011131200-2220222020101201-0232002301002131-3003230012103211-0030221001011333-1201313223231030-3131203302023023"></a>

#### `ipsec.ike_parameters.ike_phase1_profile.tenant` property

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3111311323322130-0103333120230033-0312100202300122-2130222013013001-1120011132233320-1030022111132100-0112003132302101-0232131200220223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ike_parameters.ike_phase2_profile` properties

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-2222210331120321-3310203330131102-1332111323102002-1300213031003231-2312210013333323-3221022213020223-3312333001311121-0222322020131133)
- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-2110112022302230-2103222202112010-1012130322030213-3100331333231022-3132330220301303-2012000231322123-0021133111013121-3222203023211013)
- ipsec.ike_parameters.ike_phase2_profile

<a id="canonical-0321132300310133-2111132121003130-1111231210330313-2030131013021213-2233003332333002-2230221000211032-3002202312213230-1200000121130130"></a>

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
ike_phase2_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-1133300202210110-0312023101102111-1230310133222010-1203033231220021-0002203002232201-1321311000212212-1000033020030022-0133320301202232"></a>

### Direct properties for `ipsec.ike_parameters.ike_phase2_profile`

<a id="canonical-3323212131322032-1020010321310133-3313122100220332-2101212033003332-3133032112100302-2212121330013100-1320001130330320-3330312211111201"></a>

#### `ipsec.ike_parameters.ike_phase2_profile.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0331233113033213-2322000002100130-3111012022321000-0121003301021113-2212103221301220-2030212232022101-0231221001110000-0122110032131100"></a>

<a id="canonical-1332333121130302-1330323323112023-0031121030212110-0131332031123323-1313233311020111-3210021103010000-3021010233300230-2211311300220213"></a>

#### `ipsec.ike_parameters.ike_phase2_profile.namespace` property

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1101003003020023-1223222030301021-1123003221130323-0233122300121320-1001320311213220-1223221223030101-2232220100021123-1213002111322311"></a>

<a id="canonical-2231233122032121-3323233102110132-1233330122221102-2102003113313221-3001232320222203-0022100030103021-2131113231121231-0000111310013301"></a>

#### `ipsec.ike_parameters.ike_phase2_profile.tenant` property

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0302121201233210-2202133200011000-2021020131303223-3103331000102231-1033120300201012-3220330201011102-1112221032012003-2222313010101103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ike_parameters.initiator` properties

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-2222210331120321-3310203330131102-1332111323102002-1300213031003231-2312210013333323-3221022213020223-3312333001311121-0222322020131133)
- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-2110112022302230-2103222202112010-1012130322030213-3100331333231022-3132330220301303-2012000231322123-0021133111013121-3222203023211013)
- ipsec.ike_parameters.initiator

<a id="canonical-1113101311333010-1200221020201110-1200201230313103-0112320202133013-1322303202033133-0030230220103103-3302220231313112-3122321210313301"></a>

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
initiator = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0032203013123300-0333020112200010-0110312323131012-3221130002003033-2311223031332310-1331121120012120-1020133212213233-0103121320131033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ike_parameters.responder` properties

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-2222210331120321-3310203330131102-1332111323102002-1300213031003231-2312210013333323-3221022213020223-3312333001311121-0222322020131133)
- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-2110112022302230-2103222202112010-1012130322030213-3100331333231022-3132330220301303-2012000231322123-0021133111013121-3222203023211013)
- ipsec.ike_parameters.responder

<a id="canonical-3313300310132211-1202231013233333-0113311323020313-1022023032200113-0032210030022320-2033220113200311-1132001102331121-2112221310023023"></a>

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
responder = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1021320210230233-0300130321222101-0332113201003312-2123232100220012-3201003310002210-3322122233233202-3111300100310100-0211113122201313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ike_parameters.rm_ip_address` properties

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-2222210331120321-3310203330131102-1332111323102002-1300213031003231-2312210013333323-3221022213020223-3312333001311121-0222322020131133)
- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-2110112022302230-2103222202112010-1012130322030213-3100331333231022-3132330220301303-2012000231322123-0021133111013121-3222203023211013)
- ipsec.ike_parameters.rm_ip_address

<a id="canonical-0230313113122031-1320313221013222-3122120212302303-2010201030030012-0321213320220022-2003323031212131-0212110223003301-0310103001320220"></a>

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
rm_ip_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-1311111201110220-3310100123123000-2312133100210212-0103331300300101-3132121130311000-3131210113231302-0112223103220302-1230212110220203"></a>

### Direct properties for `ipsec.ike_parameters.rm_ip_address`

- [dual_stack](resources--external_connector--reference--group-001.md#canonical-3300211300210031-3300320111200102-0211330332013320-3223110023133020-3121232220000031-2220132331310003-0013312222201003-1212303112111323): complete subsection reference.

- [IPv4](resources--external_connector--reference--group-001.md#canonical-2100211023110201-1212230201221210-2212333202032200-2001010322201123-1211313211001102-0322322133323112-2103313031331330-1003121303203221): complete subsection reference.

- [IPv6](resources--external_connector--reference--group-001.md#canonical-0113131123101301-2232011211100010-1303113030120332-1113202110330200-3010002111121003-2033222212100031-0300133103323112-2101203111100230): complete subsection reference.

<a id="canonical-3300211300210031-3300320111200102-0211330332013320-3223110023133020-3121232220000031-2220132331310003-0013312222201003-1212303112111323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ike_parameters.rm_ip_address.dual_stack` properties

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-2222210331120321-3310203330131102-1332111323102002-1300213031003231-2312210013333323-3221022213020223-3312333001311121-0222322020131133)
- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-2110112022302230-2103222202112010-1012130322030213-3100331333231022-3132330220301303-2012000231322123-0021133111013121-3222203023211013)
- [ipsec.ike_parameters.rm_ip_address](resources--external_connector--reference--group-001.md#canonical-1021320210230233-0300130321222101-0332113201003312-2123232100220012-3201003310002210-3322122233233202-3111300100310100-0211113122201313)
- ipsec.ike_parameters.rm_ip_address.dual_stack

<a id="canonical-3031310213111222-3230210113100312-0001333211101220-1313132133231321-1211313202231231-0303031022322120-3100013032220331-2110300231023033"></a>

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

<a id="canonical-1223301000113001-3020220221321331-1201022113320101-0210330031133011-0012221312001322-1221210111021213-2130332223230000-3233130011211310"></a>

### Direct properties for `ipsec.ike_parameters.rm_ip_address.dual_stack`

- [IPv4](resources--external_connector--reference--group-001.md#canonical-1010330312021233-2123202333231112-2022201013102013-2103331131222212-0120102311022020-3313010202301333-0022010220230110-3330013313013110): complete subsection reference.

- [IPv6](resources--external_connector--reference--group-001.md#canonical-0233101221313000-3011121112231320-3313002013030010-0120222210200121-3212210103110010-2302022331302132-0330131020123302-3230133310000203): complete subsection reference.

<a id="canonical-1010330312021233-2123202333231112-2022201013102013-2103331131222212-0120102311022020-3313010202301333-0022010220230110-3330013313013110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4` properties

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-2222210331120321-3310203330131102-1332111323102002-1300213031003231-2312210013333323-3221022213020223-3312333001311121-0222322020131133)
- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-2110112022302230-2103222202112010-1012130322030213-3100331333231022-3132330220301303-2012000231322123-0021133111013121-3222203023211013)
- [ipsec.ike_parameters.rm_ip_address](resources--external_connector--reference--group-001.md#canonical-1021320210230233-0300130321222101-0332113201003312-2123232100220012-3201003310002210-3322122233233202-3111300100310100-0211113122201313)
- [ipsec.ike_parameters.rm_ip_address.dual_stack](resources--external_connector--reference--group-001.md#canonical-3300211300210031-3300320111200102-0211330332013320-3223110023133020-3121232220000031-2220132331310003-0013312222201003-1212303112111323)
- ipsec.ike_parameters.rm_ip_address.dual_stack.IPv4

<a id="canonical-1303331322122110-2120230300203133-2031003123211200-2130100211120300-2201123321030121-3332333020022000-1000102331133223-1032311030110310"></a>

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

<a id="canonical-3230232231002313-1002032032200310-3323203213013013-0333022131230013-3021203200333330-0200111110302213-1302000331032130-0222030223011033"></a>

### Direct properties for `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4`

<a id="canonical-0121331031211032-3332010003301210-1032333113121321-1233121203303200-3322312201022210-0001332102032012-3121223022312221-2032013231101221"></a>

#### `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4.addr` property

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

<a id="canonical-0233101221313000-3011121112231320-3313002013030010-0120222210200121-3212210103110010-2302022331302132-0330131020123302-3230133310000203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6` properties

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-2222210331120321-3310203330131102-1332111323102002-1300213031003231-2312210013333323-3221022213020223-3312333001311121-0222322020131133)
- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-2110112022302230-2103222202112010-1012130322030213-3100331333231022-3132330220301303-2012000231322123-0021133111013121-3222203023211013)
- [ipsec.ike_parameters.rm_ip_address](resources--external_connector--reference--group-001.md#canonical-1021320210230233-0300130321222101-0332113201003312-2123232100220012-3201003310002210-3322122233233202-3111300100310100-0211113122201313)
- [ipsec.ike_parameters.rm_ip_address.dual_stack](resources--external_connector--reference--group-001.md#canonical-3300211300210031-3300320111200102-0211330332013320-3223110023133020-3121232220000031-2220132331310003-0013312222201003-1212303112111323)
- ipsec.ike_parameters.rm_ip_address.dual_stack.IPv6

<a id="canonical-3133123131002330-2303232132022223-1221020003012131-2113112200103010-0113211030111010-1222122311033133-3133322002132120-1320310131130322"></a>

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

<a id="canonical-2002031020320102-0211022303333231-0303200013210310-2001121023301321-3112021012122102-3023113110012131-2330012320013121-2320203130210311"></a>

### Direct properties for `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6`

<a id="canonical-2333022221332003-3313322321011301-3213330220101133-3033201030003232-1302103131230003-1022131320310302-0133302120221232-0320301312112310"></a>

#### `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6.addr` property

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

<a id="canonical-2100211023110201-1212230201221210-2212333202032200-2001010322201123-1211313211001102-0322322133323112-2103313031331330-1003121303203221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ike_parameters.rm_ip_address.ipv4` properties

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-2222210331120321-3310203330131102-1332111323102002-1300213031003231-2312210013333323-3221022213020223-3312333001311121-0222322020131133)
- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-2110112022302230-2103222202112010-1012130322030213-3100331333231022-3132330220301303-2012000231322123-0021133111013121-3222203023211013)
- [ipsec.ike_parameters.rm_ip_address](resources--external_connector--reference--group-001.md#canonical-1021320210230233-0300130321222101-0332113201003312-2123232100220012-3201003310002210-3322122233233202-3111300100310100-0211113122201313)
- ipsec.ike_parameters.rm_ip_address.IPv4

<a id="canonical-2201011100230130-1231303231023313-1321003333332110-3323132121100222-1003031333011312-3313213323030112-2202232311313111-0021100021303133"></a>

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

<a id="canonical-2333303332120031-2221111103010303-1300323232331331-0132232102133000-2313333310012323-0200002113112002-2211223311130020-0333302211333323"></a>

### Direct properties for `ipsec.ike_parameters.rm_ip_address.ipv4`

<a id="canonical-3222212023201101-0232322110211022-3000123021321123-2020101001000331-1320121013233212-1120100023230233-0200321230213131-1022233232012112"></a>

#### `ipsec.ike_parameters.rm_ip_address.ipv4.addr` property

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

<a id="canonical-0113131123101301-2232011211100010-1303113030120332-1113202110330200-3010002111121003-2033222212100031-0300133103323112-2101203111100230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ike_parameters.rm_ip_address.ipv6` properties

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-2222210331120321-3310203330131102-1332111323102002-1300213031003231-2312210013333323-3221022213020223-3312333001311121-0222322020131133)
- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-2110112022302230-2103222202112010-1012130322030213-3100331333231022-3132330220301303-2012000231322123-0021133111013121-3222203023211013)
- [ipsec.ike_parameters.rm_ip_address](resources--external_connector--reference--group-001.md#canonical-1021320210230233-0300130321222101-0332113201003312-2123232100220012-3201003310002210-3322122233233202-3111300100310100-0211113122201313)
- ipsec.ike_parameters.rm_ip_address.IPv6

<a id="canonical-2101333132333030-0311131312013313-1303303022313123-3003322320132031-3011230123013101-0203232120303202-3032120003022333-0201032203312022"></a>

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

<a id="canonical-1303330331213222-1111310101113231-0330001311333212-3320330030321030-1323310303231123-3233121301020302-0000322232030210-1200203130330231"></a>

### Direct properties for `ipsec.ike_parameters.rm_ip_address.ipv6`

<a id="canonical-1333033130231231-2231212230111120-0120102122323211-2320322103120003-3111231332221113-2103212000113122-1033223112313210-2332101001310203"></a>

#### `ipsec.ike_parameters.rm_ip_address.ipv6.addr` property

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

<a id="canonical-2203332212111231-1233232210033231-3302200313011103-1320320031132222-3020011013212211-2011021031300333-1202320013123100-1231103113130112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ike_parameters.use_default_local_ike_id` properties

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-2222210331120321-3310203330131102-1332111323102002-1300213031003231-2312210013333323-3221022213020223-3312333001311121-0222322020131133)
- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-2110112022302230-2103222202112010-1012130322030213-3100331333231022-3132330220301303-2012000231322123-0021133111013121-3222203023211013)
- ipsec.ike_parameters.use_default_local_ike_id

<a id="canonical-3002301301123320-0213210132211202-1313310322310322-2112331103301101-1001330331233213-0212133100032301-1030322020221220-0020113302111212"></a>

Type: `"object"`. single nested block, Optional.

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
use_default_local_ike_id {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2332221223221312-2330120323211321-3330111123031100-3130000121101013-2201300311311122-2332203102103332-2032230023131002-0213122010103222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ike_parameters.use_default_remote_ike_id` properties

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-2222210331120321-3310203330131102-1332111323102002-1300213031003231-2312210013333323-3221022213020223-3312333001311121-0222322020131133)
- [ipsec.ike_parameters](resources--external_connector--reference--group-001.md#canonical-2110112022302230-2103222202112010-1012130322030213-3100331333231022-3132330220301303-2012000231322123-0021133111013121-3222203023211013)
- ipsec.ike_parameters.use_default_remote_ike_id

<a id="canonical-1013200302312030-0302213022310213-1132120312001120-2212022300003010-0332020102310233-2022122210230032-1321331032321011-1133213010231301"></a>

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
use_default_remote_ike_id = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1111312210033110-1323001002020013-0301303102011303-3323022223203100-3130333022311330-3033000102231322-1330001332110221-0323222300300120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ipsec_tunnel_parameters` properties

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-2222210331120321-3310203330131102-1332111323102002-1300213031003231-2312210013333323-3221022213020223-3312333001311121-0222322020131133)
- ipsec.ipsec_tunnel_parameters

<a id="canonical-2000211302111030-1002232222001013-2322020300211012-1113000032010330-1212331330001100-3330133121121202-1323120203220023-0130021223020223"></a>

Type: `"object"`. single nested block, Optional.

In this section, we will configure the tunnel parameters, source, destination, IP addresses, and
segment.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("psk",
    "tunnel_eps",
    "tunnel_mtu"),
  validators.ConflictingObjectAttributes("segment",
    "site_local_inside_network"),
  validators.ConflictingObjectAttributes("segment",
    "site_local_network"),
  validators.ConflictingObjectAttributes("site_local_inside_network",
    "site_local_network")}
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
  "x-ves-oneof-field-tunnel_source_vn": "[\"segment\",\"site_local_inside_network\",\"site_local_network\"]"
}
```

Terraform syntax:

```terraform
ipsec_tunnel_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-0331013210123202-1212100323132013-2031222221332323-0330303100302021-3330201000032203-3333312121200022-1022031122203322-3233222003011222"></a>

### Direct properties for `ipsec.ipsec_tunnel_parameters`

- [peer_ip_address](resources--external_connector--reference--group-001.md#canonical-3021301113200132-1210020301222122-0132011012332233-0303100330222213-3111333010321220-0302310301201301-0332113022313333-0202300202033212): complete subsection reference.

<a id="canonical-2130331102103110-1323031001132332-2021230011131113-0122230123113023-3300331221121200-1301301213230231-0211221322202331-0121232303132003"></a>

<a id="canonical-2010003303021223-2333030323323020-2000212231301030-2123332322000033-2203220222300322-3030022011201312-1122012213013100-1303212101322311"></a>

#### `ipsec.ipsec_tunnel_parameters.psk` property

Type: `"string"`. Optional.

The IKE pre-shared key (PSK) is required to ensure the IKE peers can authenticate one another within
IKE phase 1 negotiation.

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

- [segment](resources--external_connector--reference--group-001.md#canonical-1322200131300021-2200232001312301-2123232000201310-3131302201230021-0000003021103031-1222031021130120-1133230031113323-3321303122021132): complete subsection reference.

- [site_local_inside_network](resources--external_connector--reference--group-001.md#canonical-2332112022211302-1302311111102000-0323130303200002-0313033021210121-3323011232312210-3320302013132310-3023330010323130-1012320130200130): complete subsection reference.

- [site_local_network](resources--external_connector--reference--group-001.md#canonical-0212133002302120-2211331113002301-3013233312320033-3023101102111312-1020110212313333-3300212010102331-0320232210302132-0133023333323012): complete subsection reference.

- [tunnel_eps](resources--external_connector--reference--group-001.md#canonical-1131202200302232-3323130320112131-0032233332221201-0333102102123321-2102313203000200-0221133230230021-2231211011303003-1233332202212302): complete subsection reference.

<a id="canonical-2202330331123200-3030312012022023-2230002222221013-1122032321232010-1303223202021331-3211321112033000-3212302203131100-3021303311022013"></a>

<a id="canonical-3210032102221020-1232022123011320-1222232230202112-1010001211003233-3310032130020020-3001311220111100-1000322323202311-0032221013023311"></a>

#### `ipsec.ipsec_tunnel_parameters.tunnel_mtu` property

Type: `"number"`. Optional.

The tunnel MTU defines the maximum size of the packet that can be sent through the tunnel without
needing to be fragmented.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(512, 1370),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1370,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 512
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "512",
    "ves.io.schema.rules.uint32.lte": "1370"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "512",
    "ves.io.schema.rules.uint32.lte": "1370"
  }
}
```

<a id="canonical-3021301113200132-1210020301222122-0132011012332233-0303100330222213-3111333010321220-0302310301201301-0332113022313333-0202300202033212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ipsec_tunnel_parameters.peer_ip_address` properties

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-2222210331120321-3310203330131102-1332111323102002-1300213031003231-2312210013333323-3221022213020223-3312333001311121-0222322020131133)
- [ipsec.ipsec_tunnel_parameters](resources--external_connector--reference--group-001.md#canonical-1111312210033110-1323001002020013-0301303102011303-3323022223203100-3130333022311330-3033000102231322-1330001332110221-0323222300300120)
- ipsec.ipsec_tunnel_parameters.peer_ip_address

<a id="canonical-1222202100122001-1132303203012323-2132202203003233-1002012203001302-3333010203032123-0221331113312103-2110112231212130-1022120103220010"></a>

Type: `"object"`. single nested block, Optional.

IPv4 Address. IPv4 Address in dot-decimal notation.

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
peer_ip_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-1102101331203320-2112113120220120-2132333011111211-3002013010020130-1300131331101021-3231321222302313-0012213012032322-3221331030031233"></a>

### Direct properties for `ipsec.ipsec_tunnel_parameters.peer_ip_address`

<a id="canonical-3132311233002322-2032111302322130-1133310233103323-3022131333121130-2203303100333130-3320310121133221-2131333111233220-3103310103003223"></a>

#### `ipsec.ipsec_tunnel_parameters.peer_ip_address.addr` property

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

<a id="canonical-1322200131300021-2200232001312301-2123232000201310-3131302201230021-0000003021103031-1222031021130120-1133230031113323-3321303122021132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ipsec_tunnel_parameters.segment` properties

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-2222210331120321-3310203330131102-1332111323102002-1300213031003231-2312210013333323-3221022213020223-3312333001311121-0222322020131133)
- [ipsec.ipsec_tunnel_parameters](resources--external_connector--reference--group-001.md#canonical-1111312210033110-1323001002020013-0301303102011303-3323022223203100-3130333022311330-3033000102231322-1330001332110221-0323222300300120)
- ipsec.ipsec_tunnel_parameters.segment

<a id="canonical-3220203300230031-3111233233231131-0223010232032001-0121232120211000-2312030023112231-0002023121120213-1321022100130331-0320123203032103"></a>

Type: `"object"`. single nested block, Optional.

Segment Reference Type. Reference to Segment Object.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("refs")}
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
segment {
  # Configure direct properties listed below.
}
```

<a id="canonical-0021330312013313-2220130032000332-0323101200003320-2110021132222300-0320311311203013-0120010202230113-0303222321331132-0220110331203020"></a>

### Direct properties for `ipsec.ipsec_tunnel_parameters.segment`

- [refs](resources--external_connector--reference--group-001.md#canonical-1003032313213122-0230220010230120-2331111030321011-2303003103110013-3300012130130110-3131021102332200-1130312332303122-2122002111022321): complete subsection reference.

<a id="canonical-1003032313213122-0230220010230120-2331111030321011-2303003103110013-3300012130130110-3131021102332200-1130312332303122-2122002111022321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ipsec_tunnel_parameters.segment.refs` properties

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-2222210331120321-3310203330131102-1332111323102002-1300213031003231-2312210013333323-3221022213020223-3312333001311121-0222322020131133)
- [ipsec.ipsec_tunnel_parameters](resources--external_connector--reference--group-001.md#canonical-1111312210033110-1323001002020013-0301303102011303-3323022223203100-3130333022311330-3033000102231322-1330001332110221-0323222300300120)
- [ipsec.ipsec_tunnel_parameters.segment](resources--external_connector--reference--group-001.md#canonical-1322200131300021-2200232001312301-2123232000201310-3131302201230021-0000003021103031-1222031021130120-1133230031113323-3321303122021132)
- ipsec.ipsec_tunnel_parameters.segment.refs

<a id="canonical-1111112212330121-0310001020031230-2331312023101001-2322110130302102-0212130303002133-1333300301113001-1133130130100000-2331321332110112"></a>

Type: `"object"`. list nested block, Optional.

Segment. Reference to Segment Object.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
refs {
  # Configure direct properties listed below.
}
```

<a id="canonical-3131130130103030-2300330201321122-2332021300312130-0110003113312122-3302113323121202-2112313301012201-2333301322121210-2313100331200212"></a>

### Direct properties for `ipsec.ipsec_tunnel_parameters.segment.refs`

<a id="canonical-1002102112031002-3311023333132113-3111013103222300-2230033301213023-0322122002322122-3102100130131123-0231332110102032-2322201020333000"></a>

#### `ipsec.ipsec_tunnel_parameters.segment.refs.kind` property

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

<a id="canonical-3210202131322220-2023002213321012-2000003301210202-2233010311112321-1000100111102213-0212102310130112-2303203032320130-1011032311232213"></a>

<a id="canonical-1003120303130033-3202012023003023-1232013233223211-2313010302332300-2033331033231021-0010310012303320-1331332120111212-0303033322221032"></a>

#### `ipsec.ipsec_tunnel_parameters.segment.refs.name` property

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

<a id="canonical-3003102303302012-1112312013102011-1010132223013202-2233300032112320-1020100112312020-3230213321130331-0310100200201210-3300031121111320"></a>

<a id="canonical-1130012010322213-2310303122122200-3132002000012220-1012302123133202-0023102313012110-0210021332122000-1310131033200311-2101331210031020"></a>

#### `ipsec.ipsec_tunnel_parameters.segment.refs.namespace` property

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

<a id="canonical-1320110112123002-1302303013331001-2133313210231001-3011322130312013-0132022023322223-0013223331331202-2002210211132011-2312220213132022"></a>

<a id="canonical-2213130012230201-2232332031223220-3312120311313202-1220101010020213-1333220010110320-3102023103310033-3021232103332001-1333211313220001"></a>

#### `ipsec.ipsec_tunnel_parameters.segment.refs.tenant` property

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

<a id="canonical-1313003130002031-1201322331012123-3203202211302120-3221212300210032-0313022122213012-0102130212230102-2232113022213030-0313303120030112"></a>

<a id="canonical-2202312101211101-0333322332202301-3123023231201102-2210022202010333-1133103223033201-3200321221011222-3030233133310202-1320021103001313"></a>

#### `ipsec.ipsec_tunnel_parameters.segment.refs.uid` property

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

<a id="canonical-2332112022211302-1302311111102000-0323130303200002-0313033021210121-3323011232312210-3320302013132310-3023330010323130-1012320130200130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ipsec_tunnel_parameters.site_local_inside_network` properties

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-2222210331120321-3310203330131102-1332111323102002-1300213031003231-2312210013333323-3221022213020223-3312333001311121-0222322020131133)
- [ipsec.ipsec_tunnel_parameters](resources--external_connector--reference--group-001.md#canonical-1111312210033110-1323001002020013-0301303102011303-3323022223203100-3130333022311330-3033000102231322-1330001332110221-0323222300300120)
- ipsec.ipsec_tunnel_parameters.site_local_inside_network

<a id="canonical-2032233211000222-0300201322231100-2233210323233132-3312010033333132-0211233030321102-3230100120311130-2022320330302003-1311120013021211"></a>

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
site_local_inside_network = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0212133002302120-2211331113002301-3013233312320033-3023101102111312-1020110212313333-3300212010102331-0320232210302132-0133023333323012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ipsec_tunnel_parameters.site_local_network` properties

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-2222210331120321-3310203330131102-1332111323102002-1300213031003231-2312210013333323-3221022213020223-3312333001311121-0222322020131133)
- [ipsec.ipsec_tunnel_parameters](resources--external_connector--reference--group-001.md#canonical-1111312210033110-1323001002020013-0301303102011303-3323022223203100-3130333022311330-3033000102231322-1330001332110221-0323222300300120)
- ipsec.ipsec_tunnel_parameters.site_local_network

<a id="canonical-1001132321202233-2221021003321323-0221222100122102-0033122023112010-2110211330321332-2303112021221200-1313020331123220-3230203120323002"></a>

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
site_local_network = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1131202200302232-3323130320112131-0032233332221201-0333102102123321-2102313203000200-0221133230230021-2231211011303003-1233332202212302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ipsec_tunnel_parameters.tunnel_eps` properties

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232)
- [ipsec](resources--external_connector--reference--group-001.md#canonical-2222210331120321-3310203330131102-1332111323102002-1300213031003231-2312210013333323-3221022213020223-3312333001311121-0222322020131133)
- [ipsec.ipsec_tunnel_parameters](resources--external_connector--reference--group-001.md#canonical-1111312210033110-1323001002020013-0301303102011303-3323022223203100-3130333022311330-3033000102231322-1330001332110221-0323222300300120)
- ipsec.ipsec_tunnel_parameters.tunnel_eps

<a id="canonical-2102102030212130-0320210230212102-0133012021233111-2102201211311201-2030223202331330-1332322103311310-0100311202022101-1211110310101113"></a>

Type: `"object"`. list nested block, Optional.

Configure tunnel parameters, local and remote IP addresses.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("interface",
    "local_tunnel_ip",
    "node",
    "remote_tunnel_ip")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
tunnel_eps {
  # Configure direct properties listed below.
}
```

<a id="canonical-3010212001130310-0220303002222123-1000123012212230-2313130301123303-0131003223321302-3130033123101331-2030213223223013-0002133123120232"></a>

### Direct properties for `ipsec.ipsec_tunnel_parameters.tunnel_eps`

<a id="canonical-1121323222313320-0003203130001233-3120313313220101-2113110033021133-2223002110033323-2232132113313010-1011333001321113-1012301112320132"></a>

#### `ipsec.ipsec_tunnel_parameters.tunnel_eps.interface` property

Type: `"string"`. Optional.

For the chosen node, specify the interface that will be the tunnel source.

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

<a id="canonical-1210311120321111-0011102103000131-1030333333113133-2032322021020233-3213313032233101-1002323001030313-0322010300202311-0210130232300130"></a>

<a id="canonical-3331300332231222-1102030233303023-2111111230112003-2210223330220203-3323230313022323-3021021023011333-1303203012103200-3321012222030202"></a>

#### `ipsec.ipsec_tunnel_parameters.tunnel_eps.local_tunnel_ip` property

Type: `"string"`. Optional.

For a particular tunnel on a node, specify the local tunnel IP Address i.e. The IP address of the
tunnel on the CE node itself and a subnet prefix length.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-3100333331302313-3110213313220212-1103310001323223-3100203111022012-0323121000311303-0133223101122201-2302121221210031-3132322301321003"></a>

<a id="canonical-2033021010111310-2232233332013210-2220313211022203-1313000021233323-2312210121001033-0223300120011233-2220302332001231-3231013111023302"></a>

#### `ipsec.ipsec_tunnel_parameters.tunnel_eps.node` property

Type: `"string"`. Optional.

A CE site is composed of multiple nodes. Choose a node that will be part of this external
connection.

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

<a id="canonical-1031233323111312-3303132112110133-0112133103111333-2121223001010122-0123021123111231-3303331023210023-2332212002020303-1131313002322131"></a>

<a id="canonical-1021022313231122-0000112103211120-2033301011032321-2321021131002302-2203232130011332-2331102011312302-2122020123133131-2001232303331031"></a>

#### `ipsec.ipsec_tunnel_parameters.tunnel_eps.remote_tunnel_ip` property

Type: `"string"`. Optional.

For a particular tunnel on a node, specify the remote tunnel IP Address i.e. The IP address of the
tunnel on the remote gateway and a subnet prefix length.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-2023312120203110-1221003013330002-1112111012032321-3130202022101320-2101222231102210-3302313300130100-1023233130323102-1003201233323103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_external_connector](../resources/external_connector.md#canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222)
- [Property reference](resources--external_connector--reference--group-001.md#canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232)
- timeouts

<a id="canonical-2022231212123301-1222323121000031-3211121303202020-0102200131010310-1233332332313002-3202232231123333-1111122222031021-1121232220222221"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3032233131000101-1230221303330002-2300110210210031-0230113013113031-2100303000002302-1202113313321311-2203323033332110-2332201001333201"></a>

### Direct properties for `timeouts`

<a id="canonical-3301313321202123-2213120303020031-2122320122322000-3330202021130332-1011122322033300-1323003020031323-1131023321203231-0300201223020320"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2223320110321020-0013221113011303-3230030101312310-3123203002300302-0202322230010001-2113010122202320-3001001030101321-0331103010131202"></a>

<a id="canonical-0330323311323203-0010232000102002-2310223211010303-1333223320002111-3022321102001301-3230103030101331-2121220210001102-2031222023312033"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2012121233322212-2330321030222001-2020222111320022-1120222032010132-3220213201203333-2302123212012212-3101002331311132-1301033230032200"></a>

<a id="canonical-1123232003231022-3203331323223313-3210110320101303-2233211321230301-2000110320011301-1232032221012032-2132133333020002-1201101310221112"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2010010233330333-0101000131233121-1300122032311002-0331222313011301-2002201131220030-1002120320230222-3310330223221033-1200113033302112"></a>

<a id="canonical-2003221133223110-0233221300022012-3311210130003321-0132012333012221-1213000213303303-2110001320133212-0022130333203303-1310032320312321"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
