---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- Property reference

<a id="canonical-0023213003101031-3122200131201122-1100000233333332-2030003000023122-0212232231123011-0222222233133021-3022132112313232-1230122301332002"></a>

### Direct properties for `xcsh_securemesh_site_v2`

- [active_enhanced_firewall_policies](resources--securemesh_site_v2--reference--group-003.md#canonical-3310233122203010-0332320002300320-2331123231131321-0313313323121201-1111200133001032-0002303100331000-3011120302021001-1133033311121312): complete subsection reference.

- [active_forward_proxy_policies](resources--securemesh_site_v2--reference--group-003.md#canonical-3010320022010021-3103003201202111-2112211003112101-3112013110332012-1323333333221100-2220023130312020-2323032311203013-0213121322032231): complete subsection reference.

- [admin_user_credentials](resources--securemesh_site_v2--reference--group-003.md#canonical-2323112033210020-2113012223203221-0322233010022032-2122320020101010-2102301101033021-3002303312010131-2012011322012232-0231103113201210): complete subsection reference.

<a id="canonical-0033012201232313-3033022001201220-1130012302300012-0332320320001031-3012201320313110-3112220320212302-0221102222120311-2222010210203132"></a>

<a id="canonical-2000031201301001-3222302213101323-3102132003002132-3003310333000021-1302311232331332-1201011131011322-2103010110312323-1223301033111232"></a>

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

- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223): complete subsection reference.

- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320): complete subsection reference.

- [baremetal](resources--securemesh_site_v2--reference--group-005.md#canonical-2201113133130333-1301221030003121-3232200130003031-2003232222313211-3311100232302121-2001323120021322-0012000213003031-1030112013211303): complete subsection reference.

- [block_all_services](resources--securemesh_site_v2--reference--group-005.md#canonical-0213320123112011-2202110211322101-0013303221220031-1312221223201333-2013012110330312-2222232103123133-2323003000320332-2232323130313122): complete subsection reference.

- [blocked_services](resources--securemesh_site_v2--reference--group-005.md#canonical-1023020113032331-1130112021101312-2020122123031022-1200330020130000-0332012013303030-0103033110120230-3200102311032133-1212201133110122): complete subsection reference.

- [custom_proxy](resources--securemesh_site_v2--reference--group-006.md#canonical-1002312300302030-1233211121221322-1133130202001302-3111103230122000-0211003002110003-1012033010002203-3321112312311233-3033212203133131): complete subsection reference.

- [custom_proxy_bypass](resources--securemesh_site_v2--reference--group-006.md#canonical-1201120122021230-3013113003303010-0312222210310120-1032233221132010-0003011010201111-3212223322123210-1233330103221220-3111220102212230): complete subsection reference.

- [dc_cluster_group_sli](resources--securemesh_site_v2--reference--group-006.md#canonical-1020321133311233-2110133311213121-2123333100012012-3200301200123213-1020103130321222-1021002013333211-1002020122222102-3103333110002323): complete subsection reference.

- [dc_cluster_group_slo](resources--securemesh_site_v2--reference--group-006.md#canonical-3100323033003130-1302331203022103-1233101210123010-3220222322211011-0332110232310011-1003002331302211-2010313222320033-2132103002200131): complete subsection reference.

<a id="canonical-3100302131222112-3200010221002120-1100131023233332-0011203101212320-3022221332130120-3111021332233301-3300301012033322-1033320210003311"></a>

<a id="canonical-3020123103031101-0122021013121331-0101101102032313-2202322112100021-1101032332301023-0312123113003211-2013221232123320-3121302233021023"></a>

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

<a id="canonical-1001313101320223-1031002010022333-1110032003000110-0023322202123213-0312011012123332-1112223320120211-2300320013020220-2020221120113232"></a>

<a id="canonical-2122021130011223-1232030302311021-0131010132203101-0331122131021310-0021021030112122-3013310123212323-0233310023303133-2100001333001310"></a>

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

- [disable_advanced_delivery](resources--securemesh_site_v2--reference--group-006.md#canonical-1112203010312210-0132130311102202-0022222202032000-2131322332010323-3210102110223000-3200001233331101-1223311301211212-3200100111113210): complete subsection reference.

- [disable_ha](resources--securemesh_site_v2--reference--group-006.md#canonical-3231002302323332-2013010223311332-0021120301210112-2232331120122110-0301210303200310-2002220333321202-2022103131122023-3101003100301002): complete subsection reference.

- [disable_log_anonymization](resources--securemesh_site_v2--reference--group-006.md#canonical-3203112302132303-1313202311322332-0222023201233202-3221021303032331-1111112120320030-1100333003113200-0133320331321111-1203121122310001): complete subsection reference.

- [disable_management_network](resources--securemesh_site_v2--reference--group-006.md#canonical-3131220221232201-2023022102131221-2033023020002031-0101312232133333-3212110223303200-0212322023231221-2122222012123200-2002031123021212): complete subsection reference.

- [disable_url_categorization](resources--securemesh_site_v2--reference--group-006.md#canonical-0302003210031020-0002223310123233-0230133013212201-1212020012311333-0112301130131323-0032203021212031-1033011301012302-3123112120202231): complete subsection reference.

- [dns_ntp_config](resources--securemesh_site_v2--reference--group-006.md#canonical-0212020030133010-3223023110132101-2003232322020333-0130023111013232-0023300023001223-0131110322121120-0330310333311013-0300201302020000): complete subsection reference.

- [eks_k8s](resources--securemesh_site_v2--reference--group-006.md#canonical-0022130130300012-0230333000310233-1322033213223312-3321101002313103-1201131030031302-2021203021021121-0133011100321201-1030300311013100): complete subsection reference.

- [enable_advanced_delivery](resources--securemesh_site_v2--reference--group-008.md#canonical-1120312013323111-2001313211213303-3013322321101221-2302001312022101-1002300000201112-0101023103222121-0113230102321331-1332233101200023): complete subsection reference.

- [enable_ha](resources--securemesh_site_v2--reference--group-008.md#canonical-1220133303231202-3120323110021310-0221103231002213-1302123301122110-0113201221110233-2033320311323331-0000121000323311-3123322022023132): complete subsection reference.

- [enable_log_anonymization](resources--securemesh_site_v2--reference--group-008.md#canonical-3313333110020022-1323331321323133-1131231332001000-0330200311301223-2112002130210103-2212202210323103-0001322331303020-1112333301131100): complete subsection reference.

- [enable_management_network](resources--securemesh_site_v2--reference--group-008.md#canonical-3101121100030333-3231330312100021-1120003212311331-1332211112201100-3020232221303003-0132200233113110-1011133120133331-0113033310032120): complete subsection reference.

- [enable_url_categorization](resources--securemesh_site_v2--reference--group-008.md#canonical-2110022213030001-1113110301320221-2101200333030220-3333200320222031-3001312302021121-3133213020013103-1112002120223030-0213300210232131): complete subsection reference.

- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022): complete subsection reference.

- [f5_proxy](resources--securemesh_site_v2--reference--group-009.md#canonical-3021301201330310-0230130120101310-2231211023332220-1302200202031013-3111011231130033-1300202330311330-3213031120313332-2203211111320211): complete subsection reference.

- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012): complete subsection reference.

<a id="canonical-2120221000003031-3300201203322021-0011233110212322-1012323112031122-2122313110322312-3302312322212103-1203313002223200-2103013322020212"></a>

<a id="canonical-0100013202003331-2013323311113211-3133332010022112-2111103220101112-0300123321212131-2333332122002101-1011010111020122-0202300120100002"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [kvm](resources--securemesh_site_v2--reference--group-009.md#canonical-2330221111302030-3002303131302300-2010230031111120-1101330230301332-1330202312303122-0010020230132231-3211003200310221-1112230211023123): complete subsection reference.

<a id="canonical-3222032310231120-0320100321302303-3032212220221200-1300102300003321-2322023031110312-0203230200333202-3131300032112022-2010310322123103"></a>

<a id="canonical-0311133221112232-3213333330000302-3133010011312311-3220001100311022-0233031002322200-1213013322023121-3021231031121223-0112010022312302"></a>

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

- [load_balancing](resources--securemesh_site_v2--reference--group-010.md#canonical-2013023210221220-2012002230302233-0332210211312003-2210303202232112-2101201000333121-3101333310310123-2112213032220120-1121200031201313): complete subsection reference.

- [local_vrf](resources--securemesh_site_v2--reference--group-010.md#canonical-1103030301332311-0012312112031332-3113233131020203-3312111201312111-2011011123321011-3032202132331002-3233300301021211-1000123313112021): complete subsection reference.

- [log_receiver_with_net](resources--securemesh_site_v2--reference--group-011.md#canonical-2121032223220303-3321210233231023-1001332020100333-3000331113000212-2133111021333102-2331032012311032-2322012211021133-3030101010311030): complete subsection reference.

- [logs_streaming_disabled](resources--securemesh_site_v2--reference--group-011.md#canonical-2210233313020231-0310131332331201-3031332230130222-0133101033233310-0223230331130213-2103112112102321-1230002312101232-1211013133212221): complete subsection reference.

<a id="canonical-3222331110121033-3311231110131120-1121332002032212-0111113301122310-0303221000030333-2113211131333211-3303000223311103-1202010023213303"></a>

<a id="canonical-1021212200030202-2030023221113102-2132021330213032-2122222003131103-1311102033313023-2312122230203033-2000101100110012-2230321212300033"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Securemesh Site V2. Must be unique within the namespace. Must be at most 63 characters
(DNS-1035).

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  validators.NameValidator(),
  stringvalidator.LengthAtMost(63),
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

<a id="canonical-1121211130331033-0202102233223100-1200121023032300-0301320010130022-3200110200330100-3120130230101121-0120230233110100-1222323322333111"></a>

<a id="canonical-0122133030231112-2310300003111233-0013000020212122-0212300211220032-1121233113211203-3112002022122013-0010001020120303-2013023032232110"></a>

#### `namespace` property

Type: `"string"`. Optional, Computed.

Namespace for the Securemesh Site V2. The F5 XC API restricts this resource to the system namespace;
it defaults to that value and may be omitted.

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

- [no_forward_proxy](resources--securemesh_site_v2--reference--group-011.md#canonical-3231001003211033-1113201231132322-0013121220032323-2031203003230221-1133223030313302-3131212323231233-1301021322010201-1030310211331131): complete subsection reference.

- [no_network_policy](resources--securemesh_site_v2--reference--group-011.md#canonical-1111331220330301-2030301032232003-0321010121010110-1033012330011003-3321221320333313-2222101321233313-3201213303322110-3011030103300011): complete subsection reference.

- [no_proxy_bypass](resources--securemesh_site_v2--reference--group-011.md#canonical-3123331333122110-3122333032323313-1322031023033330-3022013201322102-1030322011123202-1032113110203010-1103213002302323-3011223220033322): complete subsection reference.

- [no_s2s_connectivity_sli](resources--securemesh_site_v2--reference--group-011.md#canonical-3322012222010332-2230100331221211-0132203200323130-2003020232110130-0012210112022001-1220312313020321-1303222003211323-2020120133231203): complete subsection reference.

- [no_s2s_connectivity_slo](resources--securemesh_site_v2--reference--group-011.md#canonical-1032112221231302-2123032103000023-1212321030120112-1301100133021333-3122002203103112-2022222012000331-0333232203322311-3201133120130202): complete subsection reference.

- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321): complete subsection reference.

- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311): complete subsection reference.

- [offline_survivability_mode](resources--securemesh_site_v2--reference--group-013.md#canonical-2001312130230101-3132012012311122-0230112101013213-1333302011302001-3223331311230002-0320013221303203-3122311112320113-1320320233110001): complete subsection reference.

- [openshift_virtualization](resources--securemesh_site_v2--reference--group-014.md#canonical-0121302321233330-0012112200322203-1301113300000200-3112222001032003-2331112123023303-3200031230001123-2210131121011110-1303130031122213): complete subsection reference.

- [openstack](resources--securemesh_site_v2--reference--group-015.md#canonical-3103330221113213-1101002321121001-1323321202302031-0033322000331021-3233221231330113-2212103122303202-3302221330110020-0111211323131322): complete subsection reference.

- [performance_enhancement_mode](resources--securemesh_site_v2--reference--group-016.md#canonical-2130013131033231-3232223210311301-1121020200130212-1011012213303102-1120012133331003-0202331032111120-2302012123333131-0300312001101100): complete subsection reference.

- [private_adn](resources--securemesh_site_v2--reference--group-016.md#canonical-1132333020000033-3121313020113130-1113011323131231-1303130201313202-3331202201110330-3232030020101012-0011123211010231-1301321320123330): complete subsection reference.

- [re_select](resources--securemesh_site_v2--reference--group-016.md#canonical-3302000210133302-2333120223333302-1211202010321103-0020233023121312-3323122111210033-0120212210023120-0333102200010203-3320002122220132): complete subsection reference.

- [segment_vrf](resources--securemesh_site_v2--reference--group-016.md#canonical-2013033031131331-2213313223020200-2112133031330202-2303320232020100-1213211121321303-2130010121132312-0000300320202222-1331023001013310): complete subsection reference.

- [site_mesh_group_on_slo](resources--securemesh_site_v2--reference--group-017.md#canonical-1132320030201122-0222122012131332-0210322212020301-1032300301203333-3202102311232202-3000202023203231-1030101202313102-0300121202300311): complete subsection reference.

- [software_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-0022231303311111-0320032323222122-1103000112000213-1212111100333213-2111302213021020-3313130203211320-3301133033221310-2322002021330133): complete subsection reference.

- [timeouts](resources--securemesh_site_v2--reference--group-017.md#canonical-2131133220012220-0222312200312320-2332131230210133-2323001331232302-0312012222021021-0213223112001210-2110313123231232-3312321032110030): complete subsection reference.

<a id="canonical-3133122210320312-0132331031011200-0312000330010012-1233002100001302-0221201332103112-0231302311020222-0122300213201231-2210113013013120"></a>

<a id="canonical-2222320132122210-3113331313233003-3310021120333202-1230213022000211-1113312200130031-3012123212123101-2032220201120211-3022103030131002"></a>

#### `tunnel_dead_timeout` property

Type: `"number"`. Optional, Computed.

Time interval, in millisec, within which any IPsec / SSL connection from the site going down is
detected. When not set (== 0), a default value of 10000 msec will be used.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 180000),
}
```

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

<a id="canonical-0021201232320123-0313020001313211-0113113100301020-0231113202121110-1220220301322111-1202013213220111-1322121002210033-1201233000120133"></a>

<a id="canonical-1212211011021010-0200001332302201-0302031321233022-1300230302231221-1313033200230023-1313103031002300-0222011200232003-3102020222111310"></a>

#### `tunnel_type` property

Type: `"string"`. Optional, Computed.

\[Enum:
SITE\_TO\_SITE\_TUNNEL\_IPSEC\_OR\_SSL|SITE\_TO\_SITE\_TUNNEL\_IPSEC|SITE\_TO\_SITE\_TUNNEL\_SSL\]
Tunnel encapsulation to be used between sites Tunnel can operate in both IPsec and SSL, with IPsec
being preferred over SSL. Tunnel is of type IPsec Tunnel is of type SSL. Possible values are
\`SITE\_TO\_SITE\_TUNNEL\_IPSEC\_OR\_SSL\`, \`SITE\_TO\_SITE\_TUNNEL\_IPSEC\`,
\`SITE\_TO\_SITE\_TUNNEL\_SSL\`. Defaults to \`SITE\_TO\_SITE\_TUNNEL\_IPSEC\_OR\_SSL\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["SITE_TO_SITE_TUNNEL_IPSEC","SITE_TO_SITE_TUNNEL_IPSEC_OR_SSL","SITE_TO_SITE_TUNNEL_SSL"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("SITE_TO_SITE_TUNNEL_IPSEC_OR_SSL",
    "SITE_TO_SITE_TUNNEL_IPSEC",
    "SITE_TO_SITE_TUNNEL_SSL"),
}
```

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

- [upgrade_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-2122031130011121-2120222320030031-2032103312220112-3121002222230322-2332133023022103-2321103112123110-2013202113333333-2203321000331223): complete subsection reference.

- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330): complete subsection reference.
