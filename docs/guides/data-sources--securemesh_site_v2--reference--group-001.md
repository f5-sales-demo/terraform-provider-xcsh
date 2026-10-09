---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md#canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001)
- Property reference

<a id="canonical-3203102203000211-2331101011103122-1023321310030232-1100002312332300-2313012300111222-1121232301232213-3203023301021022-1332120023231110"></a>

### Direct properties for `xcsh_securemesh_site_v2`

- [active_enhanced_firewall_policies](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1331221220031203-3301303103233220-1020321103300210-0212001010311133-3121031113032022-0002002212321231-3023122233121033-1133313030332020): complete subsection reference.

- [active_forward_proxy_policies](data-sources--securemesh_site_v2--reference--group-003.md#canonical-2312330000320133-0311220133203323-2330231032312003-2220020022330201-1232220311333223-2330303301022322-2210012112010331-1233200303031103): complete subsection reference.

- [admin_user_credentials](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1103103321213310-1213020121320230-0221023123031100-3220023132033011-1331231123232002-0102130230202022-2211100211020020-0313231031032131): complete subsection reference.

<a id="canonical-2310123220020203-3123200011023110-3002130330131333-3130302133310012-0311221021130120-2331232333321010-1003202210131131-2222101202011213"></a>

<a id="canonical-1233120033313233-0301123012300303-2102330211100012-3320012131232012-0101311302220031-2303201120132332-1223311103113103-1322002132131133"></a>

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

- [aws](data-sources--securemesh_site_v2--reference--group-003.md#canonical-1202311132020131-3033221033032312-1303201020312113-0010013030220000-0202012232223100-1012200232222203-3213112213020133-2021301103323331): complete subsection reference.

- [Azure](data-sources--securemesh_site_v2--reference--group-004.md#canonical-0020122133011221-1213033210121110-3133203312000032-1311221201230000-2333130202300010-2132002120101312-2033013031212212-3122103003303121): complete subsection reference.

- [baremetal](data-sources--securemesh_site_v2--reference--group-005.md#canonical-1123321011122102-0201133310320122-3113112022111023-0022010233222132-1221231211213321-0332120031323322-1001302031301121-3023321101013203): complete subsection reference.

- [block_all_services](data-sources--securemesh_site_v2--reference--group-006.md#canonical-0022133033201330-1120002103203123-2320032113020210-0120031322102012-2100011310330101-1110132030033300-1020033233132031-2201302023301101): complete subsection reference.

- [blocked_services](data-sources--securemesh_site_v2--reference--group-006.md#canonical-2331300013301233-3123202011131231-0002223230222322-0323132000330211-1330032112212213-0012313001102332-3200223201231313-0010032013213000): complete subsection reference.

- [custom_proxy](data-sources--securemesh_site_v2--reference--group-006.md#canonical-3133032200000010-2012232033033322-0220322222233113-0302010000011203-2301213330322212-2131333332221112-2010302123013211-1132111231230100): complete subsection reference.

- [custom_proxy_bypass](data-sources--securemesh_site_v2--reference--group-006.md#canonical-1031033310322022-0301132212030010-1323201221211321-3313320110313211-3210322323302033-2222331020312200-2202123022123110-3100133101310020): complete subsection reference.

- [dc_cluster_group_sli](data-sources--securemesh_site_v2--reference--group-006.md#canonical-0331333120221112-3321130121321302-3001003321012123-0010231022300022-1102031332232001-2102300220303010-3321230323103000-1132201303013000): complete subsection reference.

- [dc_cluster_group_slo](data-sources--securemesh_site_v2--reference--group-006.md#canonical-0130312023322303-0003210032232131-3333000311023213-3222232203111221-0130232232022112-1313223302102100-2120032002123000-0210133313022103): complete subsection reference.

<a id="canonical-1212333313222231-2203013012320001-2220111020302012-0021321132233232-2012330212002113-0013201030102102-1103200311322211-1113310213121020"></a>

<a id="canonical-0021203130322301-2101032223232230-2012332020303011-1112201310101030-1331310033011330-2123133312000200-2000020031020032-1012013222200303"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the SecuremeshSiteV2.

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

- [disable_advanced_delivery](data-sources--securemesh_site_v2--reference--group-006.md#canonical-1213023230210130-1032003120132101-2201122000211032-0322111123031323-3030102122310010-0013322302230131-1332131200210210-3120002013301322): complete subsection reference.

- [disable_ha](data-sources--securemesh_site_v2--reference--group-007.md#canonical-2013213202032232-1201030110001201-3020211321332212-2120302120002312-3223301330322102-1111020201033332-1202321012210330-1330223033112022): complete subsection reference.

- [disable_log_anonymization](data-sources--securemesh_site_v2--reference--group-007.md#canonical-2000123210022123-3130010030010020-3222133323002330-1330102221233312-2223302333231222-3030012330222023-1201101333332131-1003200010130101): complete subsection reference.

- [disable_management_network](data-sources--securemesh_site_v2--reference--group-007.md#canonical-3230123220103000-0033000120201100-0013210112221103-0321230123310232-2021103220331321-3202332010302123-1100131023010011-1121000100310033): complete subsection reference.

- [disable_url_categorization](data-sources--securemesh_site_v2--reference--group-007.md#canonical-1303210012021303-2101331133310130-1301011330021322-2131013133123011-2222000220311232-0321313100101021-0323131000000033-3113133310130133): complete subsection reference.

- [dns_ntp_config](data-sources--securemesh_site_v2--reference--group-007.md#canonical-3330210130131331-1111102213112232-3103010133222311-1323003033030132-2002312332303123-2122020221132202-3030100123303333-1322230230022212): complete subsection reference.

- [eks_k8s](data-sources--securemesh_site_v2--reference--group-007.md#canonical-2123212200201001-0222322202022031-0020100300002010-0113030330021212-1020223223032332-0101032101031133-3131323330222223-2211012301313223): complete subsection reference.

- [enable_advanced_delivery](data-sources--securemesh_site_v2--reference--group-008.md#canonical-3110231233201103-1211200233220331-1223220332233133-2230100132132213-1110323332211333-2100302112103010-2232030331320033-1331322222110122): complete subsection reference.

- [enable_ha](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0101002333321120-0012212210300003-3121103101312310-3331102111332301-2213213102231013-1030003131222231-1133022113031011-1323132311231201): complete subsection reference.

- [enable_log_anonymization](data-sources--securemesh_site_v2--reference--group-008.md#canonical-2323032112001202-2221032111003112-0322313203323200-3110222233312301-2300310111032030-2010230200012033-1231113132102102-0322030233330201): complete subsection reference.

- [enable_management_network](data-sources--securemesh_site_v2--reference--group-008.md#canonical-1100210112101123-3322221302032001-0231112220330223-1212133320020300-2330133111020123-3013120322122011-2022211022110312-0032022012331111): complete subsection reference.

- [enable_url_categorization](data-sources--securemesh_site_v2--reference--group-008.md#canonical-3202010212322121-2023003321100101-1021232022011021-3302013313030110-3132312302310113-0102231023022002-2330103221012233-0203020022030011): complete subsection reference.

- [equinix](data-sources--securemesh_site_v2--reference--group-008.md#canonical-0030310233033312-0313122321301031-3010231101203102-3130031113332300-0221201012022100-1202330212133002-2113000010310020-2110011122032003): complete subsection reference.

- [f5_proxy](data-sources--securemesh_site_v2--reference--group-009.md#canonical-2300101301002031-2100011030120022-0123330033310013-2200013131331122-0300010013002303-2111121113111031-2300322332300313-1313112003030122): complete subsection reference.

- [gcp](data-sources--securemesh_site_v2--reference--group-009.md#canonical-0302022103000033-1230031130021213-1223013302211010-0021322330110123-0001022330113211-1332223223221113-1222112020311202-3321232330203321): complete subsection reference.

<a id="canonical-1323000022032003-3312322310010203-2203211200220123-3112032101023232-2321012111031003-3230210012011123-3213220000210211-2002001030332230"></a>

<a id="canonical-3333100120101211-3021023001031201-0301333201201013-1230000122022300-3012212000210221-2111013312000311-0230232221302313-2023213110230001"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [kvm](data-sources--securemesh_site_v2--reference--group-010.md#canonical-1213013212011230-3201131232202211-0200003231333231-3210033231231002-3133001113032013-3132200332210333-3001130311322320-0231030112011032): complete subsection reference.

<a id="canonical-0210110002003023-0231121030202000-3100030303232110-3313202320111302-3321221013020021-0000113302030012-3112131210133100-3312200321023003"></a>

<a id="canonical-1221202213321212-3300132223232102-2102131101131132-2211023133300300-0233113311213233-2020201001001032-0220111233122132-0202232313332233"></a>

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

- [load_balancing](data-sources--securemesh_site_v2--reference--group-011.md#canonical-0121223111122001-3321121211131212-0130203002210033-1220320201001121-1121022313021321-2102030003320310-3112312203330120-0133131122212110): complete subsection reference.

- [local_vrf](data-sources--securemesh_site_v2--reference--group-011.md#canonical-3013302102012130-1200330033002031-0311100330013332-0130020210001233-3230110003202023-3003321101132311-0212113312202331-2013011213230211): complete subsection reference.

- [log_receiver_with_net](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3023110201200001-2313230223011122-3303112211100000-3132120202123300-1210303000011110-0232103302130300-0331113010221003-0103313111112332): complete subsection reference.

- [logs_streaming_disabled](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3312333002021230-2313121320113132-0202121101132333-1020313233323133-1212020320020130-0320131000002003-3213221321030320-2011010311010100): complete subsection reference.

<a id="canonical-3003313000332032-3001100103103022-1132313320112202-3213010210100010-1023222231220002-1022101302133210-1313032113010333-2133133320010222"></a>

<a id="canonical-0222233333312023-0230130212132020-2101110033131001-1103312331022222-1033222011112012-2013300211212221-1031331313102110-3011012111201320"></a>

#### `name` property

Type: `"string"`. Required.

Name of the SecuremeshSiteV2.

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

<a id="canonical-2312021203333220-0211133302110132-1002330223310110-1222303001032001-3223212111220101-0011320031003312-2331002112200001-1033200112000220"></a>

<a id="canonical-2102012100323311-1211231021033333-3013021033111022-3331021320300221-0032132301001102-1233213100311221-1123312332233302-2030112002102330"></a>

#### `namespace` property

Type: `"string"`. Optional, Computed.

Namespace where the SecuremeshSiteV2 exists.

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

- [no_forward_proxy](data-sources--securemesh_site_v2--reference--group-012.md#canonical-0233333211333233-2312020110021301-2123021233033012-3010102332313210-3231313033300120-1121330230203121-1210102010020212-3302330001121210): complete subsection reference.

- [no_network_policy](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2300302332233003-0012321023230001-0002203123133321-3301231303333030-2131311231231300-3321021131122301-1111003011023013-3302002202323103): complete subsection reference.

- [no_proxy_bypass](data-sources--securemesh_site_v2--reference--group-012.md#canonical-3320113330113312-0230001301120322-3002203321220201-2321010231000321-1003133122001312-0312021000200203-3112012033001130-1333100212331010): complete subsection reference.

- [no_s2s_connectivity_sli](data-sources--securemesh_site_v2--reference--group-012.md#canonical-1323111133021223-1232233001220010-1033101302021013-2102121030310221-2202010203020030-2300133202310213-2120230100101000-2011231333110010): complete subsection reference.

- [no_s2s_connectivity_slo](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2020332030132111-0211012220122133-0321312221332102-2200333121310232-3111223133212200-2331110022130102-3133103133003221-2231231211133133): complete subsection reference.

- [nutanix](data-sources--securemesh_site_v2--reference--group-012.md#canonical-2233222210303030-3023131102031002-1032103133000223-1033011101033002-0130331102212200-3120300200121131-0000122323133322-3333200220101112): complete subsection reference.

- [oci](data-sources--securemesh_site_v2--reference--group-013.md#canonical-2210330221012211-1132021000021330-3321213312321112-2000013102100230-3113000221221132-3220013203002203-2023320300230320-3121232120122311): complete subsection reference.

- [offline_survivability_mode](data-sources--securemesh_site_v2--reference--group-014.md#canonical-1021200122303311-3233121032110203-1312312313203131-3300220212031202-1330203323130121-3303330222102333-1033013213023320-1320312232330232): complete subsection reference.

- [openshift_virtualization](data-sources--securemesh_site_v2--reference--group-014.md#canonical-3000023023210002-3023010130232302-1031002021101023-2323012100311023-1033112120223000-1321031120001331-1321002322111103-2001113201120031): complete subsection reference.

- [openstack](data-sources--securemesh_site_v2--reference--group-015.md#canonical-2231200212123223-2011111021100223-3200311311202302-3313333002320030-3101033311333111-3232311330300000-3113101222113013-1131210321001203): complete subsection reference.

- [performance_enhancement_mode](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1210330222322011-3322102030101023-2323200312011122-3030103323232033-2231300322300102-1010130313313223-0330310332203011-1000012122032112): complete subsection reference.

- [private_adn](data-sources--securemesh_site_v2--reference--group-016.md#canonical-1202313021223030-0121131323321332-2312202322112231-3212210331300302-0303320322301001-1121333100211113-3321310112012112-3133330323023222): complete subsection reference.

- [re_select](data-sources--securemesh_site_v2--reference--group-016.md#canonical-2121113003333022-2122100030330221-0011311122211220-2010201213300022-0132320300020113-3222232131112223-3232100101011113-0213130312020332): complete subsection reference.

- [segment_vrf](data-sources--securemesh_site_v2--reference--group-016.md#canonical-0230130300320210-0011023031331123-3231201210213202-0320133333320113-3130021121223112-2102321332323131-1022213002332031-0333312130233302): complete subsection reference.

- [site_mesh_group_on_slo](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0130300332201231-2102323331103230-1113310112020311-2132031122200311-2020200210212122-1333131013011300-0222101112203210-1213212023001031): complete subsection reference.

<a id="canonical-3310110230011202-2321011230331001-2110002301303232-0030021133112210-1232200003130020-3120030010021003-2211200112023321-2022112221110231"></a>

<a id="canonical-0230100303011032-1213222020120200-1203002321222201-0102233102211200-3133103202000200-2123113311230033-1221213123101111-1000310330003001"></a>

#### `tunnel_dead_timeout` property

Type: `"number"`. Computed.

Time interval, in millisec, within which any IPsec / SSL connection from the site going down is
detected. When not set (== 0), a default value of 10000 msec will be used.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 180000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.uint32.lte": "180000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "180000"
  }
}
```

<a id="canonical-1132201212132022-1022323021030032-1100321023311233-2012312232113122-1120013001130111-3122113032031002-1303033111022230-2332301003031223"></a>

<a id="canonical-3001310200331130-0013111133022200-1012212121132001-0221322223003130-3220122002020011-1101030330223332-1323331101131332-3312321211030020"></a>

#### `tunnel_type` property

Type: `"string"`. Computed.

\[Enum:
SITE\_TO\_SITE\_TUNNEL\_IPSEC\_OR\_SSL|SITE\_TO\_SITE\_TUNNEL\_IPSEC|SITE\_TO\_SITE\_TUNNEL\_SSL\]
Tunnel encapsulation to be used between sites Tunnel can operate in both IPsec and SSL, with IPsec
being preferred over SSL. Tunnel is of type IPsec Tunnel is of type SSL. Possible values are
\`SITE\_TO\_SITE\_TUNNEL\_IPSEC\_OR\_SSL\`, \`SITE\_TO\_SITE\_TUNNEL\_IPSEC\`,
\`SITE\_TO\_SITE\_TUNNEL\_SSL\`. Defaults to \`SITE\_TO\_SITE\_TUNNEL\_IPSEC\_OR\_SSL\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_TO_SITE_TUNNEL_IPSEC_OR_SSL",
  "enum": [
    "SITE_TO_SITE_TUNNEL_IPSEC_OR_SSL",
    "SITE_TO_SITE_TUNNEL_IPSEC",
    "SITE_TO_SITE_TUNNEL_SSL"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [upgrade_settings](data-sources--securemesh_site_v2--reference--group-017.md#canonical-0030100133010212-1200030111021312-3123032332031021-0010103012023120-2202203012321113-0113233110320102-2113301111200001-0103012002023131): complete subsection reference.

- [vmware](data-sources--securemesh_site_v2--reference--group-017.md#canonical-2210000000111300-0230202103322211-3033112311032330-0323231101022110-0223023121132012-3311132320032113-3220232031322232-2322132210333232): complete subsection reference.
