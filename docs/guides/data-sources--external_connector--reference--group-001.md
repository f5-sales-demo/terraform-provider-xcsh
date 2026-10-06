---
page_title: "xcsh_external_connector reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_external_connector reference."
---

# xcsh_external_connector reference

<a id="canonical-1233001313220202-0330031300031121-2122102203322002-0320020033102030-1000131023310231-2120311131221020-3130000033333030-3202120133013332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112)
- Property reference

<a id="canonical-1023120330310122-0230322220102300-1103030210302333-1202013212202321-3320033321222112-2233220333020210-3223313101212110-3312001221103200"></a>

### Direct properties for `xcsh_external_connector`

<a id="canonical-2221102201300331-2223113323201223-0333211131131233-0010221230302112-1231213313000103-2113111323112303-0200023102123322-2012301320102122"></a>

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

- [ce_site_reference](data-sources--external_connector--reference--group-001.md#canonical-0021313222132001-2203210302103211-2131132031033011-2312020221330132-1002132133021002-1002112203123220-0032201323123120-0313033230120130): complete subsection reference.

<a id="canonical-0032003113102320-2303213301112230-3201022333112213-2333221203130010-2101200123331201-0331003020022003-0023032301000300-3023122200323023"></a>

<a id="canonical-2231321113221310-3002002211100233-1130322022102103-0033130111223101-0330123232311212-3312033112010010-0220110232103123-2220003013303210"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the ExternalConnector.

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

- [gre](data-sources--external_connector--reference--group-001.md#canonical-2133011320013230-2201033111331033-2133120330113333-1201323012323010-0121303102130003-0012323311331211-2221202310220031-3312112211211021): complete subsection reference.

<a id="canonical-0312132001332101-1231101200201212-3300233030332221-0133132101201121-3110000303222101-1111110112020132-1330123032033232-1303113032030002"></a>

<a id="canonical-2132000102330301-3333120231133002-1333221112033220-1303200303101003-2301022100320030-2123210220112110-3003121001231201-3300300323132033"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-0111331000011310-0023220213313113-3101231221331210-3023110001002331-1103223322102302-3020032332011320-2011013120011211-2320213112003233): complete subsection reference.

<a id="canonical-1130220213020121-1302223213002202-2111120332002332-2032131311030031-1023013010301021-2332012032100102-0311330100012321-1001233301113201"></a>

<a id="canonical-3121010003002132-0223223231231023-2222210331221130-2201020131003013-3132223232212032-1212212121321101-2203123232022323-1330220122113132"></a>

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

<a id="canonical-0102311013202211-0220201221203121-2210021113332023-3331233202203211-0311313323021111-0012333111130202-0030221320120022-0003322023311211"></a>

<a id="canonical-2120233330221202-2301320322030231-1121230000333012-3133023223313101-0320033022103202-3001311233132303-3123303303000112-3101320322320321"></a>

#### `name` property

Type: `"string"`. Required.

Name of the ExternalConnector.

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

<a id="canonical-3001031103310112-1021311033210032-0200330322302312-0113031303132022-2312011201311320-3233303033021021-2010223233130111-3333123121133110"></a>

<a id="canonical-2001010000200100-3132031230231323-0002033300212002-1031013112300301-0202011200321003-1011122330211210-1012303231233121-1120311230021313"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the ExternalConnector exists.

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

<a id="canonical-3201230222103322-3301110331302222-0113300322113200-1031310313021100-1333300130002230-3033022013330102-2321002123112211-0000200010331023"></a>

### All schema paths for `xcsh_external_connector`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--external_connector--reference--group-001.md#canonical-2221102201300331-2223113323201223-0333211131131233-0010221230302112-1231213313000103-2113111323112303-0200023102123322-2012301320102122) |
| `ce_site_reference` | [ce_site_reference](data-sources--external_connector--reference--group-001.md#canonical-2101023032332302-3022320121201010-2111202222210303-1031230233303320-3121121131011303-1223200232223232-3202321000100010-2210130121221130) |
| `ce_site_reference.name` | [ce_site_reference.name](data-sources--external_connector--reference--group-001.md#canonical-3211323230020312-1020223001031123-2133132112331002-3220301322203232-3200231101301202-3202023030210103-3130231201233110-2022230100130211) |
| `ce_site_reference.namespace` | [ce_site_reference.namespace](data-sources--external_connector--reference--group-001.md#canonical-3122023133030012-2301100121323132-2203020220101023-1210110120123203-3113033301203013-1032230030323101-3131312021321323-0032133201021103) |
| `ce_site_reference.tenant` | [ce_site_reference.tenant](data-sources--external_connector--reference--group-001.md#canonical-3002102323313321-1120003102220011-2312333213103000-3132213001022023-3320313310112222-1003233333033233-1231332030023101-3121011102230012) |
| `description` | [description](data-sources--external_connector--reference--group-001.md#canonical-0032003113102320-2303213301112230-3201022333112213-2333221203130010-2101200123331201-0331003020022003-0023032301000300-3023122200323023) |
| `gre` | [gre](data-sources--external_connector--reference--group-001.md#canonical-3333120210131232-0001221121003202-0022010132310223-1202230221330131-0003301310200311-1120312331001031-0230111100322023-2030222131300202) |
| `gre.gre_parameters` | [gre.gre_parameters](data-sources--external_connector--reference--group-001.md#canonical-3212111311300101-0322020231230130-2210031100102300-1113211020200112-0000322121212033-2310332112311010-3203003302213030-1233121100112110) |
| `gre.gre_parameters.peer_ip_address` | [gre.gre_parameters.peer_ip_address](data-sources--external_connector--reference--group-001.md#canonical-1111222113122330-1101003203010103-3321103011213033-0300313230333332-2222221330322030-2202031131321112-1202021102300012-3213231121222312) |
| `gre.gre_parameters.peer_ip_address.addr` | [gre.gre_parameters.peer_ip_address.addr](data-sources--external_connector--reference--group-001.md#canonical-1300111302133011-1211213033313312-0200132031011133-0031332211210012-1200201321332021-2331312133112101-2312322010203310-2302020301122331) |
| `gre.gre_parameters.segment` | [gre.gre_parameters.segment](data-sources--external_connector--reference--group-001.md#canonical-3133231131212300-0120002111221020-1312310211233332-3113222132230223-3031022310233233-3030301011313311-1100133000312033-0321332210331213) |
| `gre.gre_parameters.segment.refs` | [gre.gre_parameters.segment.refs](data-sources--external_connector--reference--group-001.md#canonical-0333133010312121-1013112310112020-1030012002222312-0112112210021001-0013312012132210-0032032210321330-2013213230221020-2120332002310100) |
| `gre.gre_parameters.segment.refs.kind` | [gre.gre_parameters.segment.refs.kind](data-sources--external_connector--reference--group-001.md#canonical-1031013000211001-1300223102212210-0330330203231022-3023133200333032-0022133001200233-2103123312321011-3012310302330010-2210213202210113) |
| `gre.gre_parameters.segment.refs.name` | [gre.gre_parameters.segment.refs.name](data-sources--external_connector--reference--group-001.md#canonical-2201301031220312-2132130103130300-3330222223312212-0010320111033010-3022301000120113-1030330320211222-2112020313202230-2121220233200312) |
| `gre.gre_parameters.segment.refs.namespace` | [gre.gre_parameters.segment.refs.namespace](data-sources--external_connector--reference--group-001.md#canonical-3300031011212003-1021011223301011-2022212113112110-1010221022003211-0031333221321213-0320033321201300-2323002323021330-0320322310321302) |
| `gre.gre_parameters.segment.refs.tenant` | [gre.gre_parameters.segment.refs.tenant](data-sources--external_connector--reference--group-001.md#canonical-3232100010312312-1300213322030131-2122031222000211-1220122302230030-0331311001213321-0213111213210011-2021323121203321-0111313021331232) |
| `gre.gre_parameters.segment.refs.uid` | [gre.gre_parameters.segment.refs.uid](data-sources--external_connector--reference--group-001.md#canonical-1310003010230321-1100133210310001-3110312210213123-2212232331201332-0300112033311102-2222223220320230-2221320301012220-0223210203333332) |
| `gre.gre_parameters.site_local_inside_network` | [gre.gre_parameters.site_local_inside_network](data-sources--external_connector--reference--group-001.md#canonical-0320313222022121-0331332211322220-0323000233231100-3302020102101210-3000012230321200-3120201313232003-1213003213223022-3323003302111302) |
| `gre.gre_parameters.site_local_network` | [gre.gre_parameters.site_local_network](data-sources--external_connector--reference--group-001.md#canonical-0212212123332312-3313111313331332-0130231201121230-1021003022000100-0212210312320020-2130320333331130-0321100121001222-1011323333021312) |
| `gre.gre_parameters.tunnel_eps` | [gre.gre_parameters.tunnel_eps](data-sources--external_connector--reference--group-001.md#canonical-0220333121132331-2102312103321120-2031122133033322-2330100100001100-0200133231320033-0232012000133023-0221330331100022-3301033202230301) |
| `gre.gre_parameters.tunnel_eps.interface` | [gre.gre_parameters.tunnel_eps.interface](data-sources--external_connector--reference--group-001.md#canonical-2212113121012303-2302033120000030-1011010003231020-3010133010213000-1312312313303220-1122013223010102-0300123213122111-3232333122100333) |
| `gre.gre_parameters.tunnel_eps.local_tunnel_ip` | [gre.gre_parameters.tunnel_eps.local_tunnel_ip](data-sources--external_connector--reference--group-001.md#canonical-0130203123211202-1111021112200212-1301212303323103-0230222311201233-3010201021132231-2331300111101210-3023203013003133-2103210002000030) |
| `gre.gre_parameters.tunnel_eps.node` | [gre.gre_parameters.tunnel_eps.node](data-sources--external_connector--reference--group-001.md#canonical-2310330221312031-0112120120232213-3301321300202322-1023202123223320-1020221131120223-0332031020232012-2333233130323213-3100233301233312) |
| `gre.gre_parameters.tunnel_eps.remote_tunnel_ip` | [gre.gre_parameters.tunnel_eps.remote_tunnel_ip](data-sources--external_connector--reference--group-001.md#canonical-0002230133010213-3113020011123200-3031322332201013-1113213202013303-3213200003231230-2032013321032332-1003220300121132-1202003121030233) |
| `gre.gre_parameters.tunnel_mtu` | [gre.gre_parameters.tunnel_mtu](data-sources--external_connector--reference--group-001.md#canonical-0102030333312323-1322122031331113-2321113302003210-1102031030231121-3222303021033210-3123032231211011-3210200120221103-0322012200110103) |
| `id` | [ID](data-sources--external_connector--reference--group-001.md#canonical-0312132001332101-1231101200201212-3300233030332221-0133132101201121-3110000303222101-1111110112020132-1330123032033232-1303113032030002) |
| `ipsec` | [ipsec](data-sources--external_connector--reference--group-001.md#canonical-3100212313111230-3010003120203030-3331101322232031-2111313010230001-0212020031222102-2223232112201110-0230030303110001-0122003213201201) |
| `ipsec.ike_parameters` | [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-2031123322120301-1120103202011222-2022033312200010-1032233020032221-1132121310111133-1101022300101020-2231222203330231-1312003210003302) |
| `ipsec.ike_parameters.dpd_disabled` | [ipsec.ike_parameters.dpd_disabled](data-sources--external_connector--reference--group-001.md#canonical-0103211120020211-3210103301113002-1310103313130020-2222310300330302-2130201213223031-3312221233112110-3311313313102102-1132321330303110) |
| `ipsec.ike_parameters.dpd_keep_alive_timer` | [ipsec.ike_parameters.dpd_keep_alive_timer](data-sources--external_connector--reference--group-001.md#canonical-0332021223120310-2221301312312102-1212010220030320-1103022233200302-3010201233001003-3330312032213230-2202020123032203-3121030021312013) |
| `ipsec.ike_parameters.dpd_keep_alive_timer.timeout` | [ipsec.ike_parameters.dpd_keep_alive_timer.timeout](data-sources--external_connector--reference--group-001.md#canonical-0122230032122113-1132012012120313-3121011220022012-0101200022110033-1123033101323111-1332223303022113-3033320012122031-3310132302323030) |
| `ipsec.ike_parameters.ike_phase1_profile` | [ipsec.ike_parameters.ike_phase1_profile](data-sources--external_connector--reference--group-001.md#canonical-0132311030203100-3311013121230211-2030032001001102-0020030331103203-2221122002233321-0203221300012101-3203031003013122-1333231123100203) |
| `ipsec.ike_parameters.ike_phase1_profile.name` | [ipsec.ike_parameters.ike_phase1_profile.name](data-sources--external_connector--reference--group-001.md#canonical-0003323320303121-2113122031012013-1123103122303200-2113033331121131-1310233101102020-0330203232131123-3013300213232010-1123031031121332) |
| `ipsec.ike_parameters.ike_phase1_profile.namespace` | [ipsec.ike_parameters.ike_phase1_profile.namespace](data-sources--external_connector--reference--group-001.md#canonical-1012301033021123-3313102302103212-3332020001000230-3201030111332303-0220101332211033-3312011002330203-0013031220232231-3202230311211302) |
| `ipsec.ike_parameters.ike_phase1_profile.tenant` | [ipsec.ike_parameters.ike_phase1_profile.tenant](data-sources--external_connector--reference--group-001.md#canonical-3320221123113200-0201221301301200-2230322113021101-0103323032103303-0303322121030311-0323122222101312-0033132032231201-0100223120132202) |
| `ipsec.ike_parameters.ike_phase2_profile` | [ipsec.ike_parameters.ike_phase2_profile](data-sources--external_connector--reference--group-001.md#canonical-3200221131312313-1202231331232122-1322130203330123-2232122333003320-2023330201223132-2103022210130332-1330012020333130-3001012333000132) |
| `ipsec.ike_parameters.ike_phase2_profile.name` | [ipsec.ike_parameters.ike_phase2_profile.name](data-sources--external_connector--reference--group-001.md#canonical-2323010231311002-2132133001301123-3332230212103003-2122103111110222-0002211203300330-1022221112213131-3223211221101313-2123113220301312) |
| `ipsec.ike_parameters.ike_phase2_profile.namespace` | [ipsec.ike_parameters.ike_phase2_profile.namespace](data-sources--external_connector--reference--group-001.md#canonical-2110230010010133-2022103013310213-3220202101323231-0131013311022131-0122330311000323-0010301002201121-0230232220100233-1100210220113203) |
| `ipsec.ike_parameters.ike_phase2_profile.tenant` | [ipsec.ike_parameters.ike_phase2_profile.tenant](data-sources--external_connector--reference--group-001.md#canonical-1332123123310011-2121123310302312-1003003132211023-2132222222213322-1210313132302330-3100313021111103-3303123122012113-2203133330323213) |
| `ipsec.ike_parameters.initiator` | [ipsec.ike_parameters.initiator](data-sources--external_connector--reference--group-001.md#canonical-0123320113230012-1133311112320320-3021101100013203-3131033130021300-3221201130302330-2313130323230323-3001200102302000-3331123333301032) |
| `ipsec.ike_parameters.responder` | [ipsec.ike_parameters.responder](data-sources--external_connector--reference--group-001.md#canonical-1110031120302233-3222201312122030-3122322320313012-2222002200300000-1100122203102120-0221300032223130-2002103231003022-0001110032133032) |
| `ipsec.ike_parameters.rm_hostname` | [ipsec.ike_parameters.rm_hostname](data-sources--external_connector--reference--group-001.md#canonical-0113101211121200-2230330023303212-3020232230200203-2103301120223131-2120020313231310-2020011121300021-0000201003302223-3330231103202130) |
| `ipsec.ike_parameters.rm_ip_address` | [ipsec.ike_parameters.rm_ip_address](data-sources--external_connector--reference--group-001.md#canonical-2120132021023322-3130003113000002-2301330111013123-2213223110210321-1223320110100203-0002330303231303-2233032131113231-3022313123013032) |
| `ipsec.ike_parameters.rm_ip_address.dual_stack` | [ipsec.ike_parameters.rm_ip_address.dual_stack](data-sources--external_connector--reference--group-001.md#canonical-1112020333132021-3200012030133122-1002203032132212-2221230310113312-0332101220233301-2100301103110302-2033230231123231-1211131311011031) |
| `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4` | [ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4](data-sources--external_connector--reference--group-001.md#canonical-3110330203121322-1332111203202100-1313233113321000-1001222323030002-3110231212011120-3002200122332032-1233132211333301-2301210321123321) |
| `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4.addr` | [ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4.addr](data-sources--external_connector--reference--group-001.md#canonical-3321222331200023-0230030300322111-0313110220303330-1302223010121302-0023333111211303-0011121323101303-3313123003231133-2030331120322330) |
| `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6` | [ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6](data-sources--external_connector--reference--group-001.md#canonical-2131122120100333-1111223110112300-0012211120312120-3020202303201111-3122113102013101-2223203003223210-0133111023122122-3332011002032201) |
| `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6.addr` | [ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6.addr](data-sources--external_connector--reference--group-001.md#canonical-1300003030323233-2031201011323011-3312332221323222-1032231322122232-1002031122312102-0032122020200133-2312330113010103-2100100130002033) |
| `ipsec.ike_parameters.rm_ip_address.ipv4` | [ipsec.ike_parameters.rm_ip_address.ipv4](data-sources--external_connector--reference--group-001.md#canonical-2332012223020202-2110301010202231-2310020312132232-0201101013001131-2322102310022120-3233022121200300-2032113200321012-2112313021011200) |
| `ipsec.ike_parameters.rm_ip_address.ipv4.addr` | [ipsec.ike_parameters.rm_ip_address.ipv4.addr](data-sources--external_connector--reference--group-001.md#canonical-2233003201001100-2020112031203132-1113002330033002-3202232100221312-3113211202301002-3201322203101001-1201023122323000-3233311031112001) |
| `ipsec.ike_parameters.rm_ip_address.ipv6` | [ipsec.ike_parameters.rm_ip_address.ipv6](data-sources--external_connector--reference--group-001.md#canonical-0031313102302203-1310113101233230-3013012011320302-3002213130031320-1201022133032222-1103010023030023-0133202033111121-1322121222201031) |
| `ipsec.ike_parameters.rm_ip_address.ipv6.addr` | [ipsec.ike_parameters.rm_ip_address.ipv6.addr](data-sources--external_connector--reference--group-001.md#canonical-0133113232203130-3223000302120331-1233232021130323-0113322331213231-0212321112131020-2232003010023110-1232001322300123-1312230211110133) |
| `ipsec.ike_parameters.use_default_local_ike_id` | [ipsec.ike_parameters.use_default_local_ike_id](data-sources--external_connector--reference--group-001.md#canonical-3010212120213013-1333322112213100-0211111223133333-1112212133220012-0120330331033313-2302312002123103-1100101203201120-3121133201133002) |
| `ipsec.ike_parameters.use_default_remote_ike_id` | [ipsec.ike_parameters.use_default_remote_ike_id](data-sources--external_connector--reference--group-001.md#canonical-2000211022100003-3201300200133230-2010202021102031-0203232013222120-0211313103130223-1032100032122220-1020300123000020-1320002033111220) |
| `ipsec.ipsec_tunnel_parameters` | [ipsec.ipsec_tunnel_parameters](data-sources--external_connector--reference--group-001.md#canonical-0133303213132133-3320001310330001-3010313130212001-2031010120222302-3033033300020123-2013223103031222-0300000221133332-2232320002221333) |
| `ipsec.ipsec_tunnel_parameters.peer_ip_address` | [ipsec.ipsec_tunnel_parameters.peer_ip_address](data-sources--external_connector--reference--group-001.md#canonical-2002112313122312-3200200311230110-2222020300302202-2011001022020023-3001202221020220-0103311201101313-0303211130123313-2231200300120302) |
| `ipsec.ipsec_tunnel_parameters.peer_ip_address.addr` | [ipsec.ipsec_tunnel_parameters.peer_ip_address.addr](data-sources--external_connector--reference--group-001.md#canonical-0210223023323010-3331133033231011-2203303313013013-3320311230220003-1323010313033122-0011321103022011-3312320331112332-1222030230231023) |
| `ipsec.ipsec_tunnel_parameters.psk` | [ipsec.ipsec_tunnel_parameters.psk](data-sources--external_connector--reference--group-001.md#canonical-1031223000011003-0211012030032330-1211301003223210-3332100232102231-3030102012222230-3232333333200300-1320031110132301-3212100320200312) |
| `ipsec.ipsec_tunnel_parameters.segment` | [ipsec.ipsec_tunnel_parameters.segment](data-sources--external_connector--reference--group-001.md#canonical-1033232031211101-1103132101323103-1121310332320002-1010001333033213-3233201102302013-1300013012230220-3120101221101123-0201302303222301) |
| `ipsec.ipsec_tunnel_parameters.segment.refs` | [ipsec.ipsec_tunnel_parameters.segment.refs](data-sources--external_connector--reference--group-001.md#canonical-1213213012321022-2022213012113112-0100301001203030-1331203212011302-2121023333323220-1212021302100301-3120202332033331-2010130200233102) |
| `ipsec.ipsec_tunnel_parameters.segment.refs.kind` | [ipsec.ipsec_tunnel_parameters.segment.refs.kind](data-sources--external_connector--reference--group-001.md#canonical-2010333302031301-3131112310033221-2211012301132023-0112023130322030-2112310221202023-0323013110210310-3321312030203021-0322223203312031) |
| `ipsec.ipsec_tunnel_parameters.segment.refs.name` | [ipsec.ipsec_tunnel_parameters.segment.refs.name](data-sources--external_connector--reference--group-001.md#canonical-1211202321301212-1002311302230033-0231100303203101-3000023211230231-0123031330332030-0210301321310320-2012300320211310-2103122323021230) |
| `ipsec.ipsec_tunnel_parameters.segment.refs.namespace` | [ipsec.ipsec_tunnel_parameters.segment.refs.namespace](data-sources--external_connector--reference--group-001.md#canonical-0210200313230100-1323122013030303-3021330231010233-0112032200302200-2100331132333121-0202222000030002-3213332332330023-1320323311010331) |
| `ipsec.ipsec_tunnel_parameters.segment.refs.tenant` | [ipsec.ipsec_tunnel_parameters.segment.refs.tenant](data-sources--external_connector--reference--group-001.md#canonical-2100221233321300-2321003320330311-2133013211023313-1311132103231332-3202200232202332-3133000123101002-3132002022132310-0311101100100003) |
| `ipsec.ipsec_tunnel_parameters.segment.refs.uid` | [ipsec.ipsec_tunnel_parameters.segment.refs.uid](data-sources--external_connector--reference--group-001.md#canonical-0012131322130010-2303101312321322-3011030133321330-1123322120333003-1232120203020111-1220223101323213-1300132003100233-1220033213102113) |
| `ipsec.ipsec_tunnel_parameters.site_local_inside_network` | [ipsec.ipsec_tunnel_parameters.site_local_inside_network](data-sources--external_connector--reference--group-001.md#canonical-1223311202331013-3222103133101330-1013112112333321-2023112002113332-3200230033330233-1031122223000332-1032233222232003-1033331233023301) |
| `ipsec.ipsec_tunnel_parameters.site_local_network` | [ipsec.ipsec_tunnel_parameters.site_local_network](data-sources--external_connector--reference--group-001.md#canonical-1332101003020300-3011301121210203-0003003212322100-3020231320123213-0130030312121111-1231212330203221-0003122310011332-3030301213332000) |
| `ipsec.ipsec_tunnel_parameters.tunnel_eps` | [ipsec.ipsec_tunnel_parameters.tunnel_eps](data-sources--external_connector--reference--group-001.md#canonical-0030203223121131-0011020323332131-1233231022132110-2011132221032001-1012023311232233-0322230231122310-1302220000323023-0110313301120010) |
| `ipsec.ipsec_tunnel_parameters.tunnel_eps.interface` | [ipsec.ipsec_tunnel_parameters.tunnel_eps.interface](data-sources--external_connector--reference--group-001.md#canonical-3121133212220103-1020212030211133-0001302122321203-3010212000333002-3023023211100021-3033330110332323-3100310220000030-3022213123323002) |
| `ipsec.ipsec_tunnel_parameters.tunnel_eps.local_tunnel_ip` | [ipsec.ipsec_tunnel_parameters.tunnel_eps.local_tunnel_ip](data-sources--external_connector--reference--group-001.md#canonical-1203313231300112-0133333012310310-1321210333113101-3220230320213201-3231222002012330-2303033321133000-0330002232330022-3232132220213201) |
| `ipsec.ipsec_tunnel_parameters.tunnel_eps.node` | [ipsec.ipsec_tunnel_parameters.tunnel_eps.node](data-sources--external_connector--reference--group-001.md#canonical-3023321023210313-3330202203231130-1030030203102120-3032002031101020-2330300303312313-3300212211023301-0322112200331020-0313020303112303) |
| `ipsec.ipsec_tunnel_parameters.tunnel_eps.remote_tunnel_ip` | [ipsec.ipsec_tunnel_parameters.tunnel_eps.remote_tunnel_ip](data-sources--external_connector--reference--group-001.md#canonical-3020122231312022-0331013032300021-1333131001023130-3223311130300301-2000110131332320-2320111030110011-3232332320023323-1003202231012300) |
| `ipsec.ipsec_tunnel_parameters.tunnel_mtu` | [ipsec.ipsec_tunnel_parameters.tunnel_mtu](data-sources--external_connector--reference--group-001.md#canonical-0222000002300012-2233103300332032-3330001220021221-2213311023111333-2002311003010110-3123133212133031-0213033232302322-1211233302010312) |
| `labels` | [labels](data-sources--external_connector--reference--group-001.md#canonical-1130220213020121-1302223213002202-2111120332002332-2032131311030031-1023013010301021-2332012032100102-0311330100012321-1001233301113201) |
| `name` | [name](data-sources--external_connector--reference--group-001.md#canonical-0102311013202211-0220201221203121-2210021113332023-3331233202203211-0311313323021111-0012333111130202-0030221320120022-0003322023311211) |
| `namespace` | [namespace](data-sources--external_connector--reference--group-001.md#canonical-3001031103310112-1021311033210032-0200330322302312-0113031303132022-2312011201311320-3233303033021021-2010223233130111-3333123121133110) |

<a id="canonical-0021313222132001-2203210302103211-2131132031033011-2312020221330132-1002132133021002-1002112203123220-0032201323123120-0313033230120130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ce_site_reference` properties

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-1233001313220202-0330031300031121-2122102203322002-0320020033102030-1000131023310231-2120311131221020-3130000033333030-3202120133013332)
- ce_site_reference

<a id="canonical-2101023032332302-3022320121201010-2111202222210303-1031230233303320-3121121131011303-1223200232223232-3202321000100010-2210130121221130"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-1230313203201222-0303021200200033-1201022333132023-1330032111302132-3333210312113213-0231332301322302-3310111322203102-1203330313100232"></a>

### Direct properties for `ce_site_reference`

<a id="canonical-3211323230020312-1020223001031123-2133132112331002-3220301322203232-3200231101301202-3202023030210103-3130231201233110-2022230100130211"></a>

#### `ce_site_reference.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-3122023133030012-2301100121323132-2203020220101023-1210110120123203-3113033301203013-1032230030323101-3131312021321323-0032133201021103"></a>

<a id="canonical-0221122002010201-3300013132110022-2001021230000002-2112130303212231-0131010121233213-1333213002100121-0212233122131101-0311211202312320"></a>

#### `ce_site_reference.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-3002102323313321-1120003102220011-2312333213103000-3132213001022023-3320313310112222-1003233333033233-1231332030023101-3121011102230012"></a>

<a id="canonical-0332321200310011-0210023101102100-1012303031321210-3032301202032110-2031131223003122-3322231112203110-3032232000302333-3201130200201302"></a>

#### `ce_site_reference.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-2133011320013230-2201033111331033-2133120330113333-1201323012323010-0121303102130003-0012323311331211-2221202310220031-3312112211211021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gre` properties

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-1233001313220202-0330031300031121-2122102203322002-0320020033102030-1000131023310231-2120311131221020-3130000033333030-3202120133013332)
- gre

<a id="canonical-3333120210131232-0001221121003202-0022010132310223-1202230221330131-0003301310200311-1120312331001031-0230111100322023-2030222131300202"></a>

Type: `"single"`. Computed.

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

- [gre](data-sources--external_connector--reference--group-001.md#canonical-3333120210131232-0001221121003202-0022010132310223-1202230221330131-0003301310200311-1120312331001031-0230111100322023-2030222131300202)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-3100212313111230-3010003120203030-3331101322232031-2111313010230001-0212020031222102-2223232112201110-0230030303110001-0122003213201201)

Select alternatives according to the provider validators above.

<a id="canonical-3031302013213301-1323220323103313-0121120022210013-3322321032120013-1221101223332032-1011210301320300-1320211302303321-2303003030030221"></a>

### Direct properties for `gre`

- [gre_parameters](data-sources--external_connector--reference--group-001.md#canonical-3111101010031023-1103010101130011-1021023300000320-0213031013130000-0312022120020323-2100231323133133-3213003201030022-3233200333220233): complete subsection reference.

<a id="canonical-3111101010031023-1103010101130011-1021023300000320-0213031013130000-0312022120020323-2100231323133133-3213003201030022-3233200333220233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gre.gre_parameters` properties

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-1233001313220202-0330031300031121-2122102203322002-0320020033102030-1000131023310231-2120311131221020-3130000033333030-3202120133013332)
- [gre](data-sources--external_connector--reference--group-001.md#canonical-2133011320013230-2201033111331033-2133120330113333-1201323012323010-0121303102130003-0012323311331211-2221202310220031-3312112211211021)
- gre.gre_parameters

<a id="canonical-3212111311300101-0322020231230130-2210031100102300-1113211020200112-0000322121212033-2310332112311010-3203003302213030-1233121100112110"></a>

Type: `"single"`. Computed.

GRE configuration parameters required for GRE Connection type.

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

<a id="canonical-1332103130121213-2011100021010132-2301031320101013-3032130023313133-1223031022100331-0332102321023323-2103232320013232-1202200003220321"></a>

### Direct properties for `gre.gre_parameters`

- [peer_ip_address](data-sources--external_connector--reference--group-001.md#canonical-0232102201201332-1201210011113212-0123133032101020-1000012321300233-0113022130223121-3221103223321100-1102201211003021-2120121321222231): complete subsection reference.

- [segment](data-sources--external_connector--reference--group-001.md#canonical-3010322332310332-1103122330310323-3313213220203120-1332330202322012-1112112013111101-2003331112332000-2031133300111100-2013202322300102): complete subsection reference.

- [site_local_inside_network](data-sources--external_connector--reference--group-001.md#canonical-0011022200313223-0202222333303003-0100301121202333-3220213103022011-2320011331012113-1033222233301203-2022021132223003-3020231130301213): complete subsection reference.

- [site_local_network](data-sources--external_connector--reference--group-001.md#canonical-0020220311000002-2111003000212123-3213013103010103-0123032021221132-3033323202133013-0113212023232132-2000320032131033-3101021221002301): complete subsection reference.

- [tunnel_eps](data-sources--external_connector--reference--group-001.md#canonical-2322131213221130-0311133203112333-1001022101210013-1033122202301302-2301020122111301-0323330302132232-1321121233102010-3200332130033300): complete subsection reference.

<a id="canonical-0102030333312323-1322122031331113-2321113302003210-1102031030231121-3222303021033210-3123032231211011-3210200120221103-0322012200110103"></a>

<a id="canonical-2223310231003203-1323300003233012-2200231000301321-2013320202233113-0222110103232122-2022223331113130-2103112021212230-2311120321001000"></a>

#### `gre.gre_parameters.tunnel_mtu` property

Type: `"number"`. Computed.

Configure MTU for the GRE tunnel interface.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0232102201201332-1201210011113212-0123133032101020-1000012321300233-0113022130223121-3221103223321100-1102201211003021-2120121321222231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gre.gre_parameters.peer_ip_address` properties

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-1233001313220202-0330031300031121-2122102203322002-0320020033102030-1000131023310231-2120311131221020-3130000033333030-3202120133013332)
- [gre](data-sources--external_connector--reference--group-001.md#canonical-2133011320013230-2201033111331033-2133120330113333-1201323012323010-0121303102130003-0012323311331211-2221202310220031-3312112211211021)
- [gre.gre_parameters](data-sources--external_connector--reference--group-001.md#canonical-3111101010031023-1103010101130011-1021023300000320-0213031013130000-0312022120020323-2100231323133133-3213003201030022-3233200333220233)
- gre.gre_parameters.peer_ip_address

<a id="canonical-1111222113122330-1101003203010103-3321103011213033-0300313230333332-2222221330322030-2202031131321112-1202021102300012-3213231121222312"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0021120101231112-2101231120101112-0101013300330021-1223110330032002-1011013010331300-3211120012313110-3122231111121021-3232231130023112"></a>

### Direct properties for `gre.gre_parameters.peer_ip_address`

<a id="canonical-1300111302133011-1211213033313312-0200132031011133-0031332211210012-1200201321332021-2331312133112101-2312322010203310-2302020301122331"></a>

#### `gre.gre_parameters.peer_ip_address.addr` property

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

<a id="canonical-3010322332310332-1103122330310323-3313213220203120-1332330202322012-1112112013111101-2003331112332000-2031133300111100-2013202322300102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gre.gre_parameters.segment` properties

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-1233001313220202-0330031300031121-2122102203322002-0320020033102030-1000131023310231-2120311131221020-3130000033333030-3202120133013332)
- [gre](data-sources--external_connector--reference--group-001.md#canonical-2133011320013230-2201033111331033-2133120330113333-1201323012323010-0121303102130003-0012323311331211-2221202310220031-3312112211211021)
- [gre.gre_parameters](data-sources--external_connector--reference--group-001.md#canonical-3111101010031023-1103010101130011-1021023300000320-0213031013130000-0312022120020323-2100231323133133-3213003201030022-3233200333220233)
- gre.gre_parameters.segment

<a id="canonical-3133231131212300-0120002111221020-1312310211233332-3113222132230223-3031022310233233-3030301011313311-1100133000312033-0321332210331213"></a>

Type: `"single"`. Computed.

Segment Reference Type. Reference to Segment Object.

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

<a id="canonical-0303002121100000-0233111030010331-3210200113023001-2202100300303331-3221101002001133-3010111323113220-1030111222311001-1110211030230210"></a>

### Direct properties for `gre.gre_parameters.segment`

- [refs](data-sources--external_connector--reference--group-001.md#canonical-0303212331023212-0030122012122232-2003113222311011-3330112211131130-3112311101001033-1130102312031310-3112312110321323-0213300212112223): complete subsection reference.

<a id="canonical-0303212331023212-0030122012122232-2003113222311011-3330112211131130-3112311101001033-1130102312031310-3112312110321323-0213300212112223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gre.gre_parameters.segment.refs` properties

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-1233001313220202-0330031300031121-2122102203322002-0320020033102030-1000131023310231-2120311131221020-3130000033333030-3202120133013332)
- [gre](data-sources--external_connector--reference--group-001.md#canonical-2133011320013230-2201033111331033-2133120330113333-1201323012323010-0121303102130003-0012323311331211-2221202310220031-3312112211211021)
- [gre.gre_parameters](data-sources--external_connector--reference--group-001.md#canonical-3111101010031023-1103010101130011-1021023300000320-0213031013130000-0312022120020323-2100231323133133-3213003201030022-3233200333220233)
- [gre.gre_parameters.segment](data-sources--external_connector--reference--group-001.md#canonical-3010322332310332-1103122330310323-3313213220203120-1332330202322012-1112112013111101-2003331112332000-2031133300111100-2013202322300102)
- gre.gre_parameters.segment.refs

<a id="canonical-0333133010312121-1013112310112020-1030012002222312-0112112210021001-0013312012132210-0032032210321330-2013213230221020-2120332002310100"></a>

Type: `"list"`. Computed.

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-2032210000313233-0301333101021010-0232101310102232-0012033103200131-1322312233110331-0110212113231220-2030120210132203-3202321203023212"></a>

### Direct properties for `gre.gre_parameters.segment.refs`

<a id="canonical-1031013000211001-1300223102212210-0330330203231022-3023133200333032-0022133001200233-2103123312321011-3012310302330010-2210213202210113"></a>

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

<a id="canonical-2201301031220312-2132130103130300-3330222223312212-0010320111033010-3022301000120113-1030330320211222-2112020313202230-2121220233200312"></a>

<a id="canonical-1101330023320200-1013201323123322-1230320223200101-3112132203120032-3231333313320130-3010113300113310-1131032323022312-3322232210122121"></a>

#### `gre.gre_parameters.segment.refs.name` property

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

<a id="canonical-3300031011212003-1021011223301011-2022212113112110-1010221022003211-0031333221321213-0320033321201300-2323002323021330-0320322310321302"></a>

<a id="canonical-3220023122203301-0021323323123012-2222021103220123-0112033023030122-3123223212022001-1101211201113321-0310021131002311-2320123021322213"></a>

#### `gre.gre_parameters.segment.refs.namespace` property

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

<a id="canonical-3232100010312312-1300213322030131-2122031222000211-1220122302230030-0331311001213321-0213111213210011-2021323121203321-0111313021331232"></a>

<a id="canonical-3322113321020020-3132230022231012-1000310321220010-1132101311111121-0020101220203312-2012222231313320-0211221121103130-2030011001021320"></a>

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

<a id="canonical-1310003010230321-1100133210310001-3110312210213123-2212232331201332-0300112033311102-2222223220320230-2221320301012220-0223210203333332"></a>

<a id="canonical-3032123031031330-3323322222123031-3312001203320233-3230032111133030-0110123231131201-3321322320302333-0230123131133310-3200001111313202"></a>

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

<a id="canonical-0011022200313223-0202222333303003-0100301121202333-3220213103022011-2320011331012113-1033222233301203-2022021132223003-3020231130301213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gre.gre_parameters.site_local_inside_network` properties

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-1233001313220202-0330031300031121-2122102203322002-0320020033102030-1000131023310231-2120311131221020-3130000033333030-3202120133013332)
- [gre](data-sources--external_connector--reference--group-001.md#canonical-2133011320013230-2201033111331033-2133120330113333-1201323012323010-0121303102130003-0012323311331211-2221202310220031-3312112211211021)
- [gre.gre_parameters](data-sources--external_connector--reference--group-001.md#canonical-3111101010031023-1103010101130011-1021023300000320-0213031013130000-0312022120020323-2100231323133133-3213003201030022-3233200333220233)
- gre.gre_parameters.site_local_inside_network

<a id="canonical-0320313222022121-0331332211322220-0323000233231100-3302020102101210-3000012230321200-3120201313232003-1213003213223022-3323003302111302"></a>

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

<a id="canonical-0020220311000002-2111003000212123-3213013103010103-0123032021221132-3033323202133013-0113212023232132-2000320032131033-3101021221002301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gre.gre_parameters.site_local_network` properties

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-1233001313220202-0330031300031121-2122102203322002-0320020033102030-1000131023310231-2120311131221020-3130000033333030-3202120133013332)
- [gre](data-sources--external_connector--reference--group-001.md#canonical-2133011320013230-2201033111331033-2133120330113333-1201323012323010-0121303102130003-0012323311331211-2221202310220031-3312112211211021)
- [gre.gre_parameters](data-sources--external_connector--reference--group-001.md#canonical-3111101010031023-1103010101130011-1021023300000320-0213031013130000-0312022120020323-2100231323133133-3213003201030022-3233200333220233)
- gre.gre_parameters.site_local_network

<a id="canonical-0212212123332312-3313111313331332-0130231201121230-1021003022000100-0212210312320020-2130320333331130-0321100121001222-1011323333021312"></a>

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

<a id="canonical-2322131213221130-0311133203112333-1001022101210013-1033122202301302-2301020122111301-0323330302132232-1321121233102010-3200332130033300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gre.gre_parameters.tunnel_eps` properties

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-1233001313220202-0330031300031121-2122102203322002-0320020033102030-1000131023310231-2120311131221020-3130000033333030-3202120133013332)
- [gre](data-sources--external_connector--reference--group-001.md#canonical-2133011320013230-2201033111331033-2133120330113333-1201323012323010-0121303102130003-0012323311331211-2221202310220031-3312112211211021)
- [gre.gre_parameters](data-sources--external_connector--reference--group-001.md#canonical-3111101010031023-1103010101130011-1021023300000320-0213031013130000-0312022120020323-2100231323133133-3213003201030022-3233200333220233)
- gre.gre_parameters.tunnel_eps

<a id="canonical-0220333121132331-2102312103321120-2031122133033322-2330100100001100-0200133231320033-0232012000133023-0221330331100022-3301033202230301"></a>

Type: `"list"`. Computed.

Configure tunnel parameters, source, destination, IP addresses.

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

<a id="canonical-1203000033231012-2002133001010132-0123211222031002-1312030113320322-1032121230311130-1311022232130033-1130131022230120-2001023032332321"></a>

### Direct properties for `gre.gre_parameters.tunnel_eps`

<a id="canonical-2212113121012303-2302033120000030-1011010003231020-3010133010213000-1312312313303220-1122013223010102-0300123213122111-3232333122100333"></a>

#### `gre.gre_parameters.tunnel_eps.interface` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-0130203123211202-1111021112200212-1301212303323103-0230222311201233-3010201021132231-2331300111101210-3023203013003133-2103210002000030"></a>

<a id="canonical-1221231330020113-0112131202110133-3322300222003332-2002000212022321-0311133102213130-1003203000203313-1210000003300013-0000332012301103"></a>

#### `gre.gre_parameters.tunnel_eps.local_tunnel_ip` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-2310330221312031-0112120120232213-3301321300202322-1023202123223320-1020221131120223-0332031020232012-2333233130323213-3100233301233312"></a>

<a id="canonical-0301203222031021-1021301033130210-0031011111220233-1021003133113321-0132231220213133-1222031033312002-0030102302022111-2313212011331030"></a>

#### `gre.gre_parameters.tunnel_eps.node` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-0002230133010213-3113020011123200-3031322332201013-1113213202013303-3213200003231230-2032013321032332-1003220300121132-1202003121030233"></a>

<a id="canonical-0010002302001102-0011133021123020-1113133101032122-3220120022200203-2301330321202130-0202220001023220-3120311200031212-3223011101230123"></a>

#### `gre.gre_parameters.tunnel_eps.remote_tunnel_ip` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-0111331000011310-0023220213313113-3101231221331210-3023110001002331-1103223322102302-3020032332011320-2011013120011211-2320213112003233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec` properties

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-1233001313220202-0330031300031121-2122102203322002-0320020033102030-1000131023310231-2120311131221020-3130000033333030-3202120133013332)
- ipsec

<a id="canonical-3100212313111230-3010003120203030-3331101322232031-2111313010230001-0212020031222102-2223232112201110-0230030303110001-0122003213201201"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2110232121212312-1000213001031122-0310030132113013-1323131210132312-3003022013001311-2212302321200000-0333022113132210-2222210022323210"></a>

### Direct properties for `ipsec`

- [ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-1330002233213122-1233303203212331-0103313203231320-0223301331122010-3123211313221101-1303311322332213-2133122311330211-0022223103030223): complete subsection reference.

- [ipsec_tunnel_parameters](data-sources--external_connector--reference--group-001.md#canonical-1203031111130212-3100200313212112-2011133112022013-0223221300313313-3333012221300300-3130100330021102-2012202002001221-0023322303313133): complete subsection reference.

<a id="canonical-1330002233213122-1233303203212331-0103313203231320-0223301331122010-3123211313221101-1303311322332213-2133122311330211-0022223103030223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ike_parameters` properties

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-1233001313220202-0330031300031121-2122102203322002-0320020033102030-1000131023310231-2120311131221020-3130000033333030-3202120133013332)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-0111331000011310-0023220213313113-3101231221331210-3023110001002331-1103223322102302-3020032332011320-2011013120011211-2320213112003233)
- ipsec.ike_parameters

<a id="canonical-2031123322120301-1120103202011222-2022033312200010-1032233020032221-1132121310111133-1101022300101020-2231222203330231-1312003210003302"></a>

Type: `"single"`. Computed.

IKE configuration parameters required for IPsec Connection type.

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

<a id="canonical-3310213221203123-2321200112330012-3123103111321320-3223032003303110-3020110200031033-3133333233002302-1032303222120112-0123012121112200"></a>

### Direct properties for `ipsec.ike_parameters`

- [dpd_disabled](data-sources--external_connector--reference--group-001.md#canonical-1231013332303333-1210012203230001-0303022213201321-0320021320000110-2200032320100201-1212001201313213-2310100110000322-0011303300222202): complete subsection reference.

- [dpd_keep_alive_timer](data-sources--external_connector--reference--group-001.md#canonical-2122201322103313-2010210312210111-0332212213122101-1323230223100201-1302222022030111-3001121120121131-1221130012133231-2002100220100112): complete subsection reference.

- [ike_phase1_profile](data-sources--external_connector--reference--group-001.md#canonical-3013003212013132-2022031202211130-3311131000132122-3032330012323310-3333102103102010-3302202313103010-0310212002113303-2123310202021202): complete subsection reference.

- [ike_phase2_profile](data-sources--external_connector--reference--group-001.md#canonical-2213301220223023-1223131031200311-0000002122302211-0102200031102213-2320210012033121-1013213123320122-3223131021111131-3223113010012020): complete subsection reference.

- [initiator](data-sources--external_connector--reference--group-001.md#canonical-0021101033233033-0011212032103002-3122202222211131-0010031033132102-2102323003102310-2101310023123103-3311033312313112-1223320322232302): complete subsection reference.

- [responder](data-sources--external_connector--reference--group-001.md#canonical-1322131232202301-2131222230102131-1200312030121123-1322311101123301-1210100011210111-1001020132131003-3301332301031031-3002103100330210): complete subsection reference.

<a id="canonical-0113101211121200-2230330023303212-3020232230200203-2103301120223131-2120020313231310-2020011121300021-0000201003302223-3330231103202130"></a>

<a id="canonical-2030330231303110-2321132230113213-2132332212220322-0033112013230113-1321110130302333-1323100302131333-3203103330132122-0103222113132313"></a>

#### `ipsec.ike_parameters.rm_hostname` property

Type: `"string"`. Computed.

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

- [rm_ip_address](data-sources--external_connector--reference--group-001.md#canonical-1320330321100100-3203232011031232-2303223302030101-0013201220123111-0312110012130323-2031131023200100-3332021200222322-0032310120333312): complete subsection reference.

- [use_default_local_ike_id](data-sources--external_connector--reference--group-001.md#canonical-0213021013011111-0130320001031111-0001012023310211-1312031331123231-1322003230321021-3201210010223020-3303311203223033-1111020002232320): complete subsection reference.

- [use_default_remote_ike_id](data-sources--external_connector--reference--group-001.md#canonical-1120222221112210-2312210123222130-0303123013123301-0310321303103131-2210210123212131-1303010012102123-3130000131032202-3202103101300320): complete subsection reference.

<a id="canonical-1231013332303333-1210012203230001-0303022213201321-0320021320000110-2200032320100201-1212001201313213-2310100110000322-0011303300222202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ike_parameters.dpd_disabled` properties

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-1233001313220202-0330031300031121-2122102203322002-0320020033102030-1000131023310231-2120311131221020-3130000033333030-3202120133013332)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-0111331000011310-0023220213313113-3101231221331210-3023110001002331-1103223322102302-3020032332011320-2011013120011211-2320213112003233)
- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-1330002233213122-1233303203212331-0103313203231320-0223301331122010-3123211313221101-1303311322332213-2133122311330211-0022223103030223)
- ipsec.ike_parameters.dpd_disabled

<a id="canonical-0103211120020211-3210103301113002-1310103313130020-2222310300330302-2130201213223031-3312221233112110-3311313313102102-1132321330303110"></a>

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

<a id="canonical-2122201322103313-2010210312210111-0332212213122101-1323230223100201-1302222022030111-3001121120121131-1221130012133231-2002100220100112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ike_parameters.dpd_keep_alive_timer` properties

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-1233001313220202-0330031300031121-2122102203322002-0320020033102030-1000131023310231-2120311131221020-3130000033333030-3202120133013332)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-0111331000011310-0023220213313113-3101231221331210-3023110001002331-1103223322102302-3020032332011320-2011013120011211-2320213112003233)
- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-1330002233213122-1233303203212331-0103313203231320-0223301331122010-3123211313221101-1303311322332213-2133122311330211-0022223103030223)
- ipsec.ike_parameters.dpd_keep_alive_timer

<a id="canonical-0332021223120310-2221301312312102-1212010220030320-1103022233200302-3010201233001003-3330312032213230-2202020123032203-3121030021312013"></a>

Type: `"single"`. Computed.

Configuration parameter for dpd keep alive timer.

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

<a id="canonical-1232021021313021-3102130123232001-1030321330210330-0211310202201321-1320022010023232-1313322202223220-1002110222303210-2210000030020032"></a>

### Direct properties for `ipsec.ike_parameters.dpd_keep_alive_timer`

<a id="canonical-0122230032122113-1132012012120313-3121011220022012-0101200022110033-1123033101323111-1332223303022113-3033320012122031-3310132302323030"></a>

#### `ipsec.ike_parameters.dpd_keep_alive_timer.timeout` property

Type: `"number"`. Computed.

Keepalive Timer. Operation timeout duration

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
    "ves.io.schema.rules.uint32.lte": "5"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5"
  }
}
```

<a id="canonical-3013003212013132-2022031202211130-3311131000132122-3032330012323310-3333102103102010-3302202313103010-0310212002113303-2123310202021202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ike_parameters.ike_phase1_profile` properties

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-1233001313220202-0330031300031121-2122102203322002-0320020033102030-1000131023310231-2120311131221020-3130000033333030-3202120133013332)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-0111331000011310-0023220213313113-3101231221331210-3023110001002331-1103223322102302-3020032332011320-2011013120011211-2320213112003233)
- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-1330002233213122-1233303203212331-0103313203231320-0223301331122010-3123211313221101-1303311322332213-2133122311330211-0022223103030223)
- ipsec.ike_parameters.ike_phase1_profile

<a id="canonical-0132311030203100-3311013121230211-2030032001001102-0020030331103203-2221122002233321-0203221300012101-3203031003013122-1333231123100203"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-0122330020121032-1131132200113212-2112020221013131-3020101023100033-0120130021031032-3321132312030103-0220221200301213-2112023311031032"></a>

### Direct properties for `ipsec.ike_parameters.ike_phase1_profile`

<a id="canonical-0003323320303121-2113122031012013-1123103122303200-2113033331121131-1310233101102020-0330203232131123-3013300213232010-1123031031121332"></a>

#### `ipsec.ike_parameters.ike_phase1_profile.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-1012301033021123-3313102302103212-3332020001000230-3201030111332303-0220101332211033-3312011002330203-0013031220232231-3202230311211302"></a>

<a id="canonical-0002103211331313-0231300221123223-3211003022113020-3120033033301201-0122333311103231-0122211231210013-1021310312221020-1133303020103102"></a>

#### `ipsec.ike_parameters.ike_phase1_profile.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-3320221123113200-0201221301301200-2230322113021101-0103323032103303-0303322121030311-0323122222101312-0033132032231201-0100223120132202"></a>

<a id="canonical-0001231123133302-0132023031200130-2110233013103023-2020212023000031-3300000210132320-3120230213333032-0102213120100130-2003211310212221"></a>

#### `ipsec.ike_parameters.ike_phase1_profile.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-2213301220223023-1223131031200311-0000002122302211-0102200031102213-2320210012033121-1013213123320122-3223131021111131-3223113010012020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ike_parameters.ike_phase2_profile` properties

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-1233001313220202-0330031300031121-2122102203322002-0320020033102030-1000131023310231-2120311131221020-3130000033333030-3202120133013332)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-0111331000011310-0023220213313113-3101231221331210-3023110001002331-1103223322102302-3020032332011320-2011013120011211-2320213112003233)
- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-1330002233213122-1233303203212331-0103313203231320-0223301331122010-3123211313221101-1303311322332213-2133122311330211-0022223103030223)
- ipsec.ike_parameters.ike_phase2_profile

<a id="canonical-3200221131312313-1202231331232122-1322130203330123-2232122333003320-2023330201223132-2103022210130332-1330012020333130-3001012333000132"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-3001303101312003-0331211320201002-3021111110131222-0130012322121123-2330011303000013-1321201312313123-0011022121202303-0220010110233103"></a>

### Direct properties for `ipsec.ike_parameters.ike_phase2_profile`

<a id="canonical-2323010231311002-2132133001301123-3332230212103003-2122103111110222-0002211203300330-1022221112213131-3223211221101313-2123113220301312"></a>

#### `ipsec.ike_parameters.ike_phase2_profile.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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

<a id="canonical-2110230010010133-2022103013310213-3220202101323231-0131013311022131-0122330311000323-0010301002201121-0230232220100233-1100210220113203"></a>

<a id="canonical-3200033311103321-1301222021122032-3132121002022322-2210220111032211-1321020031012212-1102202013320032-3011223302010030-2202122320120212"></a>

#### `ipsec.ike_parameters.ike_phase2_profile.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-1332123123310011-2121123310302312-1003003132211023-2132222222213322-1210313132302330-3100313021111103-3303123122012113-2203133330323213"></a>

<a id="canonical-0222021312032311-3010010130010100-3300310021111331-2132230332203101-3330121020320121-0301120002013322-2033021333110002-0310331030312112"></a>

#### `ipsec.ike_parameters.ike_phase2_profile.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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

<a id="canonical-0021101033233033-0011212032103002-3122202222211131-0010031033132102-2102323003102310-2101310023123103-3311033312313112-1223320322232302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ike_parameters.initiator` properties

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-1233001313220202-0330031300031121-2122102203322002-0320020033102030-1000131023310231-2120311131221020-3130000033333030-3202120133013332)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-0111331000011310-0023220213313113-3101231221331210-3023110001002331-1103223322102302-3020032332011320-2011013120011211-2320213112003233)
- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-1330002233213122-1233303203212331-0103313203231320-0223301331122010-3123211313221101-1303311322332213-2133122311330211-0022223103030223)
- ipsec.ike_parameters.initiator

<a id="canonical-0123320113230012-1133311112320320-3021101100013203-3131033130021300-3221201130302330-2313130323230323-3001200102302000-3331123333301032"></a>

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

<a id="canonical-1322131232202301-2131222230102131-1200312030121123-1322311101123301-1210100011210111-1001020132131003-3301332301031031-3002103100330210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ike_parameters.responder` properties

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-1233001313220202-0330031300031121-2122102203322002-0320020033102030-1000131023310231-2120311131221020-3130000033333030-3202120133013332)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-0111331000011310-0023220213313113-3101231221331210-3023110001002331-1103223322102302-3020032332011320-2011013120011211-2320213112003233)
- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-1330002233213122-1233303203212331-0103313203231320-0223301331122010-3123211313221101-1303311322332213-2133122311330211-0022223103030223)
- ipsec.ike_parameters.responder

<a id="canonical-1110031120302233-3222201312122030-3122322320313012-2222002200300000-1100122203102120-0221300032223130-2002103231003022-0001110032133032"></a>

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

<a id="canonical-1320330321100100-3203232011031232-2303223302030101-0013201220123111-0312110012130323-2031131023200100-3332021200222322-0032310120333312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ike_parameters.rm_ip_address` properties

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-1233001313220202-0330031300031121-2122102203322002-0320020033102030-1000131023310231-2120311131221020-3130000033333030-3202120133013332)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-0111331000011310-0023220213313113-3101231221331210-3023110001002331-1103223322102302-3020032332011320-2011013120011211-2320213112003233)
- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-1330002233213122-1233303203212331-0103313203231320-0223301331122010-3123211313221101-1303311322332213-2133122311330211-0022223103030223)
- ipsec.ike_parameters.rm_ip_address

<a id="canonical-2120132021023322-3130003113000002-2301330111013123-2213223110210321-1223320110100203-0002330303231303-2233032131113231-3022313123013032"></a>

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

<a id="canonical-1220330220322101-1102233003310000-3103300132231013-3302030303022001-1123012233233112-1321332330232100-0233220212330101-3303032202303330"></a>

### Direct properties for `ipsec.ike_parameters.rm_ip_address`

- [dual_stack](data-sources--external_connector--reference--group-001.md#canonical-0312102131230303-2121100310021100-0103130003102031-1120022310132223-1200303013220302-3110333123101202-2120112122223112-2133231122133120): complete subsection reference.

- [IPv4](data-sources--external_connector--reference--group-001.md#canonical-3200202022202200-2220212130011323-3231233323231031-2031123102033032-1202030023031202-0020320001100303-2201202010113012-3330113011002020): complete subsection reference.

- [IPv6](data-sources--external_connector--reference--group-001.md#canonical-3013201120232021-0113231021111100-0311021103221130-2011000011003302-3022123233312120-0310103212121220-2212133011223213-3102031013020231): complete subsection reference.

<a id="canonical-0312102131230303-2121100310021100-0103130003102031-1120022310132223-1200303013220302-3110333123101202-2120112122223112-2133231122133120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ike_parameters.rm_ip_address.dual_stack` properties

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-1233001313220202-0330031300031121-2122102203322002-0320020033102030-1000131023310231-2120311131221020-3130000033333030-3202120133013332)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-0111331000011310-0023220213313113-3101231221331210-3023110001002331-1103223322102302-3020032332011320-2011013120011211-2320213112003233)
- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-1330002233213122-1233303203212331-0103313203231320-0223301331122010-3123211313221101-1303311322332213-2133122311330211-0022223103030223)
- [ipsec.ike_parameters.rm_ip_address](data-sources--external_connector--reference--group-001.md#canonical-1320330321100100-3203232011031232-2303223302030101-0013201220123111-0312110012130323-2031131023200100-3332021200222322-0032310120333312)
- ipsec.ike_parameters.rm_ip_address.dual_stack

<a id="canonical-1112020333132021-3200012030133122-1002203032132212-2221230310113312-0332101220233301-2100301103110302-2033230231123231-1211131311011031"></a>

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

<a id="canonical-2010102122132101-3000210033103131-0120211012300202-0031232131301211-0201310121321022-1332212003232112-0300232013120223-0021112120223221"></a>

### Direct properties for `ipsec.ike_parameters.rm_ip_address.dual_stack`

- [IPv4](data-sources--external_connector--reference--group-001.md#canonical-1003321223311321-0201032313311120-0120103321101313-1120112003113230-3300020332233010-0111011203123000-1003231311003130-0320000233120231): complete subsection reference.

- [IPv6](data-sources--external_connector--reference--group-001.md#canonical-1022133333323030-2020123213200013-3121303301121223-2232311210032031-1022312101120331-3200302323122303-2033220212201111-2232211032131101): complete subsection reference.

<a id="canonical-1003321223311321-0201032313311120-0120103321101313-1120112003113230-3300020332233010-0111011203123000-1003231311003130-0320000233120231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4` properties

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-1233001313220202-0330031300031121-2122102203322002-0320020033102030-1000131023310231-2120311131221020-3130000033333030-3202120133013332)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-0111331000011310-0023220213313113-3101231221331210-3023110001002331-1103223322102302-3020032332011320-2011013120011211-2320213112003233)
- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-1330002233213122-1233303203212331-0103313203231320-0223301331122010-3123211313221101-1303311322332213-2133122311330211-0022223103030223)
- [ipsec.ike_parameters.rm_ip_address](data-sources--external_connector--reference--group-001.md#canonical-1320330321100100-3203232011031232-2303223302030101-0013201220123111-0312110012130323-2031131023200100-3332021200222322-0032310120333312)
- [ipsec.ike_parameters.rm_ip_address.dual_stack](data-sources--external_connector--reference--group-001.md#canonical-0312102131230303-2121100310021100-0103130003102031-1120022310132223-1200303013220302-3110333123101202-2120112122223112-2133231122133120)
- ipsec.ike_parameters.rm_ip_address.dual_stack.IPv4

<a id="canonical-3110330203121322-1332111203202100-1313233113321000-1001222323030002-3110231212011120-3002200122332032-1233132211333301-2301210321123321"></a>

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

<a id="canonical-0031010301221231-0110311300231020-2011230021222302-2232110312110230-0310020223231203-1201202231231122-2030121100113202-1123213100013333"></a>

### Direct properties for `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4`

<a id="canonical-3321222331200023-0230030300322111-0313110220303330-1302223010121302-0023333111211303-0011121323101303-3313123003231133-2030331120322330"></a>

#### `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv4.addr` property

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

<a id="canonical-1022133333323030-2020123213200013-3121303301121223-2232311210032031-1022312101120331-3200302323122303-2033220212201111-2232211032131101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6` properties

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-1233001313220202-0330031300031121-2122102203322002-0320020033102030-1000131023310231-2120311131221020-3130000033333030-3202120133013332)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-0111331000011310-0023220213313113-3101231221331210-3023110001002331-1103223322102302-3020032332011320-2011013120011211-2320213112003233)
- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-1330002233213122-1233303203212331-0103313203231320-0223301331122010-3123211313221101-1303311322332213-2133122311330211-0022223103030223)
- [ipsec.ike_parameters.rm_ip_address](data-sources--external_connector--reference--group-001.md#canonical-1320330321100100-3203232011031232-2303223302030101-0013201220123111-0312110012130323-2031131023200100-3332021200222322-0032310120333312)
- [ipsec.ike_parameters.rm_ip_address.dual_stack](data-sources--external_connector--reference--group-001.md#canonical-0312102131230303-2121100310021100-0103130003102031-1120022310132223-1200303013220302-3110333123101202-2120112122223112-2133231122133120)
- ipsec.ike_parameters.rm_ip_address.dual_stack.IPv6

<a id="canonical-2131122120100333-1111223110112300-0012211120312120-3020202303201111-3122113102013101-2223203003223210-0133111023122122-3332011002032201"></a>

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

<a id="canonical-1030002322103230-0013102221321222-2233200121331131-3222222322222013-1020201302102001-0210312102132302-1213020201330123-2222210311010211"></a>

### Direct properties for `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6`

<a id="canonical-1300003030323233-2031201011323011-3312332221323222-1032231322122232-1002031122312102-0032122020200133-2312330113010103-2100100130002033"></a>

#### `ipsec.ike_parameters.rm_ip_address.dual_stack.ipv6.addr` property

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

<a id="canonical-3200202022202200-2220212130011323-3231233323231031-2031123102033032-1202030023031202-0020320001100303-2201202010113012-3330113011002020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ike_parameters.rm_ip_address.ipv4` properties

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-1233001313220202-0330031300031121-2122102203322002-0320020033102030-1000131023310231-2120311131221020-3130000033333030-3202120133013332)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-0111331000011310-0023220213313113-3101231221331210-3023110001002331-1103223322102302-3020032332011320-2011013120011211-2320213112003233)
- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-1330002233213122-1233303203212331-0103313203231320-0223301331122010-3123211313221101-1303311322332213-2133122311330211-0022223103030223)
- [ipsec.ike_parameters.rm_ip_address](data-sources--external_connector--reference--group-001.md#canonical-1320330321100100-3203232011031232-2303223302030101-0013201220123111-0312110012130323-2031131023200100-3332021200222322-0032310120333312)
- ipsec.ike_parameters.rm_ip_address.IPv4

<a id="canonical-2332012223020202-2110301010202231-2310020312132232-0201101013001131-2322102310022120-3233022121200300-2032113200321012-2112313021011200"></a>

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

<a id="canonical-0102131122021001-2213030222110323-0300030031120103-3213210111000111-1330332213213020-0131011100122313-0313310020030233-0100211231311001"></a>

### Direct properties for `ipsec.ike_parameters.rm_ip_address.ipv4`

<a id="canonical-2233003201001100-2020112031203132-1113002330033002-3202232100221312-3113211202301002-3201322203101001-1201023122323000-3233311031112001"></a>

#### `ipsec.ike_parameters.rm_ip_address.ipv4.addr` property

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

<a id="canonical-3013201120232021-0113231021111100-0311021103221130-2011000011003302-3022123233312120-0310103212121220-2212133011223213-3102031013020231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ike_parameters.rm_ip_address.ipv6` properties

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-1233001313220202-0330031300031121-2122102203322002-0320020033102030-1000131023310231-2120311131221020-3130000033333030-3202120133013332)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-0111331000011310-0023220213313113-3101231221331210-3023110001002331-1103223322102302-3020032332011320-2011013120011211-2320213112003233)
- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-1330002233213122-1233303203212331-0103313203231320-0223301331122010-3123211313221101-1303311322332213-2133122311330211-0022223103030223)
- [ipsec.ike_parameters.rm_ip_address](data-sources--external_connector--reference--group-001.md#canonical-1320330321100100-3203232011031232-2303223302030101-0013201220123111-0312110012130323-2031131023200100-3332021200222322-0032310120333312)
- ipsec.ike_parameters.rm_ip_address.IPv6

<a id="canonical-0031313102302203-1310113101233230-3013012011320302-3002213130031320-1201022133032222-1103010023030023-0133202033111121-1322121222201031"></a>

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

<a id="canonical-0103021020121131-2001031120010322-1201120121301330-1130123130030013-2303322213102122-1202023322331303-0111100231333130-3220102210311133"></a>

### Direct properties for `ipsec.ike_parameters.rm_ip_address.ipv6`

<a id="canonical-0133113232203130-3223000302120331-1233232021130323-0113322331213231-0212321112131020-2232003010023110-1232001322300123-1312230211110133"></a>

#### `ipsec.ike_parameters.rm_ip_address.ipv6.addr` property

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

<a id="canonical-0213021013011111-0130320001031111-0001012023310211-1312031331123231-1322003230321021-3201210010223020-3303311203223033-1111020002232320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ike_parameters.use_default_local_ike_id` properties

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-1233001313220202-0330031300031121-2122102203322002-0320020033102030-1000131023310231-2120311131221020-3130000033333030-3202120133013332)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-0111331000011310-0023220213313113-3101231221331210-3023110001002331-1103223322102302-3020032332011320-2011013120011211-2320213112003233)
- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-1330002233213122-1233303203212331-0103313203231320-0223301331122010-3123211313221101-1303311322332213-2133122311330211-0022223103030223)
- ipsec.ike_parameters.use_default_local_ike_id

<a id="canonical-3010212120213013-1333322112213100-0211111223133333-1112212133220012-0120330331033313-2302312002123103-1100101203201120-3121133201133002"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1120222221112210-2312210123222130-0303123013123301-0310321303103131-2210210123212131-1303010012102123-3130000131032202-3202103101300320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ike_parameters.use_default_remote_ike_id` properties

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-1233001313220202-0330031300031121-2122102203322002-0320020033102030-1000131023310231-2120311131221020-3130000033333030-3202120133013332)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-0111331000011310-0023220213313113-3101231221331210-3023110001002331-1103223322102302-3020032332011320-2011013120011211-2320213112003233)
- [ipsec.ike_parameters](data-sources--external_connector--reference--group-001.md#canonical-1330002233213122-1233303203212331-0103313203231320-0223301331122010-3123211313221101-1303311322332213-2133122311330211-0022223103030223)
- ipsec.ike_parameters.use_default_remote_ike_id

<a id="canonical-2000211022100003-3201300200133230-2010202021102031-0203232013222120-0211313103130223-1032100032122220-1020300123000020-1320002033111220"></a>

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

<a id="canonical-1203031111130212-3100200313212112-2011133112022013-0223221300313313-3333012221300300-3130100330021102-2012202002001221-0023322303313133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ipsec_tunnel_parameters` properties

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-1233001313220202-0330031300031121-2122102203322002-0320020033102030-1000131023310231-2120311131221020-3130000033333030-3202120133013332)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-0111331000011310-0023220213313113-3101231221331210-3023110001002331-1103223322102302-3020032332011320-2011013120011211-2320213112003233)
- ipsec.ipsec_tunnel_parameters

<a id="canonical-0133303213132133-3320001310330001-3010313130212001-2031010120222302-3033033300020123-2013223103031222-0300000221133332-2232320002221333"></a>

Type: `"single"`. Computed.

In this section, we will configure the tunnel parameters, source, destination, IP addresses, and
segment.

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

<a id="canonical-0201200322321002-3013130033200233-0111132233122001-0120013212123022-2222101112212132-0300122300102321-3013013121010202-3121101210021111"></a>

### Direct properties for `ipsec.ipsec_tunnel_parameters`

- [peer_ip_address](data-sources--external_connector--reference--group-001.md#canonical-2020200020320011-3022011320110023-0112111022113212-0122001013120221-0202002333212101-1013010210133122-3332203113022012-1121332120022011): complete subsection reference.

<a id="canonical-1031223000011003-0211012030032330-1211301003223210-3332100232102231-3030102012222230-3232333333200300-1320031110132301-3212100320200312"></a>

<a id="canonical-1312223121301300-1220132322222132-0311310300013102-1203311113113202-2203203312331131-3121322030000023-2230321110323032-1001010013212312"></a>

#### `ipsec.ipsec_tunnel_parameters.psk` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [segment](data-sources--external_connector--reference--group-001.md#canonical-1010001120002030-1111212020010231-3000003302003103-2301312322120220-1322311223231031-3233202013113121-0223211101232333-1222132123123001): complete subsection reference.

- [site_local_inside_network](data-sources--external_connector--reference--group-001.md#canonical-3030301300110310-2222333311203021-2100221102200323-0222021132233333-2312100210032333-3000023013302101-0020031000322302-2303303011322233): complete subsection reference.

- [site_local_network](data-sources--external_connector--reference--group-001.md#canonical-0010232121033221-0111222031223112-0333323321201200-2232032132313323-1013112003221332-1331322020230203-1030200312200232-0002320123312032): complete subsection reference.

- [tunnel_eps](data-sources--external_connector--reference--group-001.md#canonical-1203320103000013-1201323323000200-3123321232320022-3310203201110313-3330131333201213-3210222223331222-2023033222220310-3303012310100120): complete subsection reference.

<a id="canonical-0222000002300012-2233103300332032-3330001220021221-2213311023111333-2002311003010110-3123133212133031-0213033232302322-1211233302010312"></a>

<a id="canonical-3311232002011303-0003112002330101-2230333021032222-3012313002301322-3103313002201211-1333203333132202-3023310122203310-1230031301131010"></a>

#### `ipsec.ipsec_tunnel_parameters.tunnel_mtu` property

Type: `"number"`. Computed.

The tunnel MTU defines the maximum size of the packet that can be sent through the tunnel without
needing to be fragmented.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2020200020320011-3022011320110023-0112111022113212-0122001013120221-0202002333212101-1013010210133122-3332203113022012-1121332120022011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ipsec_tunnel_parameters.peer_ip_address` properties

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-1233001313220202-0330031300031121-2122102203322002-0320020033102030-1000131023310231-2120311131221020-3130000033333030-3202120133013332)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-0111331000011310-0023220213313113-3101231221331210-3023110001002331-1103223322102302-3020032332011320-2011013120011211-2320213112003233)
- [ipsec.ipsec_tunnel_parameters](data-sources--external_connector--reference--group-001.md#canonical-1203031111130212-3100200313212112-2011133112022013-0223221300313313-3333012221300300-3130100330021102-2012202002001221-0023322303313133)
- ipsec.ipsec_tunnel_parameters.peer_ip_address

<a id="canonical-2002112313122312-3200200311230110-2222020300302202-2011001022020023-3001202221020220-0103311201101313-0303211130123313-2231200300120302"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2220031002020011-2020113220000001-2332200312013111-2212331220220210-1012002321322303-0221313021212030-0002223222110110-0033223333201023"></a>

### Direct properties for `ipsec.ipsec_tunnel_parameters.peer_ip_address`

<a id="canonical-0210223023323010-3331133033231011-2203303313013013-3320311230220003-1323010313033122-0011321103022011-3312320331112332-1222030230231023"></a>

#### `ipsec.ipsec_tunnel_parameters.peer_ip_address.addr` property

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

<a id="canonical-1010001120002030-1111212020010231-3000003302003103-2301312322120220-1322311223231031-3233202013113121-0223211101232333-1222132123123001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ipsec_tunnel_parameters.segment` properties

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-1233001313220202-0330031300031121-2122102203322002-0320020033102030-1000131023310231-2120311131221020-3130000033333030-3202120133013332)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-0111331000011310-0023220213313113-3101231221331210-3023110001002331-1103223322102302-3020032332011320-2011013120011211-2320213112003233)
- [ipsec.ipsec_tunnel_parameters](data-sources--external_connector--reference--group-001.md#canonical-1203031111130212-3100200313212112-2011133112022013-0223221300313313-3333012221300300-3130100330021102-2012202002001221-0023322303313133)
- ipsec.ipsec_tunnel_parameters.segment

<a id="canonical-1033232031211101-1103132101323103-1121310332320002-1010001333033213-3233201102302013-1300013012230220-3120101221101123-0201302303222301"></a>

Type: `"single"`. Computed.

Segment Reference Type. Reference to Segment Object.

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

<a id="canonical-2013111201032231-2313013032330102-2302332000123312-2033303010221330-1110032220102131-1112311031232122-2012223102231122-1322232123320001"></a>

### Direct properties for `ipsec.ipsec_tunnel_parameters.segment`

- [refs](data-sources--external_connector--reference--group-001.md#canonical-2333113322103303-0022133132111331-2021102231333003-2330332230011130-0122313301202311-1032303202223111-2111231313212332-3233311020030010): complete subsection reference.

<a id="canonical-2333113322103303-0022133132111331-2021102231333003-2330332230011130-0122313301202311-1032303202223111-2111231313212332-3233311020030010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ipsec_tunnel_parameters.segment.refs` properties

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-1233001313220202-0330031300031121-2122102203322002-0320020033102030-1000131023310231-2120311131221020-3130000033333030-3202120133013332)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-0111331000011310-0023220213313113-3101231221331210-3023110001002331-1103223322102302-3020032332011320-2011013120011211-2320213112003233)
- [ipsec.ipsec_tunnel_parameters](data-sources--external_connector--reference--group-001.md#canonical-1203031111130212-3100200313212112-2011133112022013-0223221300313313-3333012221300300-3130100330021102-2012202002001221-0023322303313133)
- [ipsec.ipsec_tunnel_parameters.segment](data-sources--external_connector--reference--group-001.md#canonical-1010001120002030-1111212020010231-3000003302003103-2301312322120220-1322311223231031-3233202013113121-0223211101232333-1222132123123001)
- ipsec.ipsec_tunnel_parameters.segment.refs

<a id="canonical-1213213012321022-2022213012113112-0100301001203030-1331203212011302-2121023333323220-1212021302100301-3120202332033331-2010130200233102"></a>

Type: `"list"`. Computed.

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-2212302001331301-2321003321013211-0001211313011101-0002321001332010-0121112220232031-0203121203202230-1122210033231300-1011220021123023"></a>

### Direct properties for `ipsec.ipsec_tunnel_parameters.segment.refs`

<a id="canonical-2010333302031301-3131112310033221-2211012301132023-0112023130322030-2112310221202023-0323013110210310-3321312030203021-0322223203312031"></a>

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

<a id="canonical-1211202321301212-1002311302230033-0231100303203101-3000023211230231-0123031330332030-0210301321310320-2012300320211310-2103122323021230"></a>

<a id="canonical-3003012231333021-0121100102123300-3123111122231302-3012300011211120-2322000003311033-1202211310333213-3211133210132223-1113030330120032"></a>

#### `ipsec.ipsec_tunnel_parameters.segment.refs.name` property

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

<a id="canonical-0210200313230100-1323122013030303-3021330231010233-0112032200302200-2100331132333121-0202222000030002-3213332332330023-1320323311010331"></a>

<a id="canonical-2000111321311101-3023221021222210-3020332131123222-0002331033232011-1110010031032022-1130130311223323-0222213322300032-2311033332331213"></a>

#### `ipsec.ipsec_tunnel_parameters.segment.refs.namespace` property

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

<a id="canonical-2100221233321300-2321003320330311-2133013211023313-1311132103231332-3202200232202332-3133000123101002-3132002022132310-0311101100100003"></a>

<a id="canonical-3201102002213301-0113202100202210-0313131110202100-0231001123120000-1110212011001122-1032113110310333-1130020233032020-2333023310113021"></a>

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

<a id="canonical-0012131322130010-2303101312321322-3011030133321330-1123322120333003-1232120203020111-1220223101323213-1300132003100233-1220033213102113"></a>

<a id="canonical-1332132202320002-0023021001121333-2102320232200030-2023132033231110-2110311010220031-1232011003022221-2120212000313013-2020201023122212"></a>

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

<a id="canonical-3030301300110310-2222333311203021-2100221102200323-0222021132233333-2312100210032333-3000023013302101-0020031000322302-2303303011322233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ipsec_tunnel_parameters.site_local_inside_network` properties

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-1233001313220202-0330031300031121-2122102203322002-0320020033102030-1000131023310231-2120311131221020-3130000033333030-3202120133013332)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-0111331000011310-0023220213313113-3101231221331210-3023110001002331-1103223322102302-3020032332011320-2011013120011211-2320213112003233)
- [ipsec.ipsec_tunnel_parameters](data-sources--external_connector--reference--group-001.md#canonical-1203031111130212-3100200313212112-2011133112022013-0223221300313313-3333012221300300-3130100330021102-2012202002001221-0023322303313133)
- ipsec.ipsec_tunnel_parameters.site_local_inside_network

<a id="canonical-1223311202331013-3222103133101330-1013112112333321-2023112002113332-3200230033330233-1031122223000332-1032233222232003-1033331233023301"></a>

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

<a id="canonical-0010232121033221-0111222031223112-0333323321201200-2232032132313323-1013112003221332-1331322020230203-1030200312200232-0002320123312032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ipsec_tunnel_parameters.site_local_network` properties

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-1233001313220202-0330031300031121-2122102203322002-0320020033102030-1000131023310231-2120311131221020-3130000033333030-3202120133013332)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-0111331000011310-0023220213313113-3101231221331210-3023110001002331-1103223322102302-3020032332011320-2011013120011211-2320213112003233)
- [ipsec.ipsec_tunnel_parameters](data-sources--external_connector--reference--group-001.md#canonical-1203031111130212-3100200313212112-2011133112022013-0223221300313313-3333012221300300-3130100330021102-2012202002001221-0023322303313133)
- ipsec.ipsec_tunnel_parameters.site_local_network

<a id="canonical-1332101003020300-3011301121210203-0003003212322100-3020231320123213-0130030312121111-1231212330203221-0003122310011332-3030301213332000"></a>

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

<a id="canonical-1203320103000013-1201323323000200-3123321232320022-3310203201110313-3330131333201213-3210222223331222-2023033222220310-3303012310100120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipsec.ipsec_tunnel_parameters.tunnel_eps` properties

Breadcrumbs:

- [xcsh_external_connector](../data-sources/external_connector.md#canonical-0300201230203033-0013121300121302-3330301212212211-1322111000002113-3130220122033010-0213230233320030-2021033121223032-3220131000330112)
- [Property reference](data-sources--external_connector--reference--group-001.md#canonical-1233001313220202-0330031300031121-2122102203322002-0320020033102030-1000131023310231-2120311131221020-3130000033333030-3202120133013332)
- [ipsec](data-sources--external_connector--reference--group-001.md#canonical-0111331000011310-0023220213313113-3101231221331210-3023110001002331-1103223322102302-3020032332011320-2011013120011211-2320213112003233)
- [ipsec.ipsec_tunnel_parameters](data-sources--external_connector--reference--group-001.md#canonical-1203031111130212-3100200313212112-2011133112022013-0223221300313313-3333012221300300-3130100330021102-2012202002001221-0023322303313133)
- ipsec.ipsec_tunnel_parameters.tunnel_eps

<a id="canonical-0030203223121131-0011020323332131-1233231022132110-2011132221032001-1012023311232233-0322230231122310-1302220000323023-0110313301120010"></a>

Type: `"list"`. Computed.

Configure tunnel parameters, local and remote IP addresses.

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

<a id="canonical-3001222031132120-0203310213110320-2010302111210202-0312111021020013-3220321333001033-3122213013212120-1131130330212221-1222231203310220"></a>

### Direct properties for `ipsec.ipsec_tunnel_parameters.tunnel_eps`

<a id="canonical-3121133212220103-1020212030211133-0001302122321203-3010212000333002-3023023211100021-3033330110332323-3100310220000030-3022213123323002"></a>

#### `ipsec.ipsec_tunnel_parameters.tunnel_eps.interface` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-1203313231300112-0133333012310310-1321210333113101-3220230320213201-3231222002012330-2303033321133000-0330002232330022-3232132220213201"></a>

<a id="canonical-2021021321322112-3033200330120013-0322201211202312-1203300223003220-2113022013201121-0011303102231333-1331133131123013-2320011221301221"></a>

#### `ipsec.ipsec_tunnel_parameters.tunnel_eps.local_tunnel_ip` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-3023321023210313-3330202203231130-1030030203102120-3032002031101020-2330300303312313-3300212211023301-0322112200331020-0313020303112303"></a>

<a id="canonical-0130202110131311-0003131322133110-0312133302002120-2132001003023003-2323101100121020-0233330131000121-1013313231023211-1011123031223013"></a>

#### `ipsec.ipsec_tunnel_parameters.tunnel_eps.node` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-3020122231312022-0331013032300021-1333131001023130-3223311130300301-2000110131332320-2320111030110011-3232332320023323-1003202231012300"></a>

<a id="canonical-1100331030033332-2001222030121103-1230312021313030-3233103233113233-2222010201012100-3300033311333312-0300103130033223-2122321323102132"></a>

#### `ipsec.ipsec_tunnel_parameters.tunnel_eps.remote_tunnel_ip` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```
