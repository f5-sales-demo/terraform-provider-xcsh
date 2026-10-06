---
page_title: "xcsh_cloud_link reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_link reference."
---

# xcsh_cloud_link reference

<a id="canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)
- Property reference

<a id="canonical-3303103302111301-0313032023300121-0033321121122232-2010131331000111-3131103121301032-2121031033231331-1330122211200321-2311330233331001"></a>

### Direct properties for `xcsh_cloud_link`

<a id="canonical-3022203033020230-3120302010013332-2312300333323132-1210333231030013-1122221301310310-0322202223322221-3333300203210323-2110321223100012"></a>

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

- [aws](resources--cloud_link--reference--group-001.md#canonical-2122012130133020-2032323112221322-1211111102003323-1300012000223020-1120220220211310-0223321133201021-1300111122111112-2011112300122012): complete subsection reference.

<a id="canonical-3210200032230100-1213321020323220-0121321032210301-1133213312223301-2021313202120013-2123111113130221-2131220222333131-3032032320122203"></a>

<a id="canonical-1122002101122200-2321133303221012-3333310330302330-1030302330323102-0201132332201032-2101132000332201-1113013031330332-3210021001023200"></a>

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

<a id="canonical-1101322110110201-3212202113330331-0202223033122310-3212232120011123-2212212101303002-2011231312221220-3112203023111020-0012011132000112"></a>

<a id="canonical-3012212111212201-2033133232323110-0321303033132033-2023231311122320-1003000113100222-1123013011202102-1113022221123103-2000110313332332"></a>

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

- [disabled](resources--cloud_link--reference--group-001.md#canonical-1022131120022203-1330132311200122-3203320013133123-1312200301110003-2220313021320223-0321311322113312-1231023302122033-0222221310010130): complete subsection reference.

- [enabled](resources--cloud_link--reference--group-001.md#canonical-2320130011300110-3021312032122332-3011210131012030-1310310010301231-2131130022302300-3011131030311231-0221010200100112-3232312022020312): complete subsection reference.

- [gcp](resources--cloud_link--reference--group-001.md#canonical-2120230010303101-2202120333222110-3132020103101203-0002303133323002-1131210100111130-2330332130032310-1203120020022120-2302330200333121): complete subsection reference.

<a id="canonical-3332231312322102-0313303311113120-1030231023323210-0230112012203132-0102022322313202-1332302110320313-1100133001302313-3001221101211303"></a>

<a id="canonical-0013302233130301-0332313031010012-3200033020031320-0323202223132200-2031111010030103-0132222210011132-0312133100232313-2211230210111132"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0330321222200303-1111032021200133-1201023123231220-1311320222303202-3112112133130233-0230202023102030-3022110103300012-1120330212030332"></a>

<a id="canonical-2102300323112032-0312231120302230-3320300003022220-2331001332033013-0113030013010111-2331130223332310-3232302121221330-0020332021322000"></a>

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

<a id="canonical-3221000132113302-1122120202212013-1131200300001030-2333120111300101-2000030202031300-0010121020233231-0112303202300110-0323220033002320"></a>

<a id="canonical-0102122323122331-2110300300112002-1132033113110331-0310013013022312-0303110201211130-1131131112021011-2221023222301233-0201032232030310"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Cloud Link. Must be unique within the namespace.

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

<a id="canonical-0300330002321121-0033032101232002-0303112112120133-3321210120002330-1111312022311212-1312130302113022-2013032200323223-2011332131112332"></a>

<a id="canonical-2113021330211013-0010123000322221-2231233031030121-2100123223030011-3013213022032323-2100312002202312-1330201222023211-0223232100003220"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Cloud Link is created.

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

- [timeouts](resources--cloud_link--reference--group-001.md#canonical-1000132120000000-2030223002311110-0323103213230233-0103003302023032-3221033120201230-3112000110012223-3112303220021202-1023123321332000): complete subsection reference.

<a id="canonical-3211123100211231-0330220321202201-3121223332101232-1123323013301320-3000231302310032-1023110000202301-3310121003031321-0312200021023222"></a>

### All schema paths for `xcsh_cloud_link`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--cloud_link--reference--group-001.md#canonical-3022203033020230-3120302010013332-2312300333323132-1210333231030013-1122221301310310-0322202223322221-3333300203210323-2110321223100012) |
| `aws` | [aws](resources--cloud_link--reference--group-001.md#canonical-1002232311000300-0122311212200020-2301013330301220-3220323120032002-1120132211223201-2120131110022330-1301200133303220-0120110221333130) |
| `aws.aws_cred` | [aws.aws_cred](resources--cloud_link--reference--group-001.md#canonical-1011302120131021-2212331122112102-1101333323310203-1012030321311330-0103332011021323-2302012001101223-2131200322013222-1122122311021311) |
| `aws.aws_cred.name` | [aws.aws_cred.name](resources--cloud_link--reference--group-001.md#canonical-2222322303010112-2311321333120310-1230332333103032-0033232302300311-1311332320022210-1033210032013110-0211212311230030-0213212302003111) |
| `aws.aws_cred.namespace` | [aws.aws_cred.namespace](resources--cloud_link--reference--group-001.md#canonical-3231022221123001-0132110120113232-3112223230020110-1120022212301112-2122011202102220-3130000300021331-3310300301211331-2112031311131333) |
| `aws.aws_cred.tenant` | [aws.aws_cred.tenant](resources--cloud_link--reference--group-001.md#canonical-1303110300332012-3012011211212021-1322020003222300-0123031013003201-2213210110121212-0010330033021332-2010222110311021-1033001320223200) |
| `aws.byoc` | [aws.byoc](resources--cloud_link--reference--group-001.md#canonical-3303113222033001-2232322100312322-0120010032123313-0031221123033231-0213211133113211-3230013311123330-2300000000222331-1301001211212202) |
| `aws.byoc.connections` | [aws.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-2013221322121002-1020010230230202-3100112020003011-0332310302233011-0322210133112013-2010313111010230-2130223133302003-1231210011000233) |
| `aws.byoc.connections.auth_key` | [aws.byoc.connections.auth_key](resources--cloud_link--reference--group-001.md#canonical-2123131322000303-3203301103111133-0313220003202222-3002022112002020-1222122321033231-1321320223120222-2332001030121311-1131221231220223) |
| `aws.byoc.connections.auth_key.blindfold_secret_info` | [aws.byoc.connections.auth_key.blindfold_secret_info](resources--cloud_link--reference--group-001.md#canonical-0233122312311111-3313103001121222-3222101102212130-3210013233220201-3332330302130212-2032333311011012-2222331202220302-2110210210021121) |
| `aws.byoc.connections.auth_key.blindfold_secret_info.decryption_provider` | [aws.byoc.connections.auth_key.blindfold_secret_info.decryption_provider](resources--cloud_link--reference--group-001.md#canonical-1122003012110300-3020331002102003-3310100313311201-0110003010221222-2223312322313203-0330102010332111-0323303301313221-1131113131332011) |
| `aws.byoc.connections.auth_key.blindfold_secret_info.location` | [aws.byoc.connections.auth_key.blindfold_secret_info.location](resources--cloud_link--reference--group-001.md#canonical-0203230103131012-0202212300121203-3000201001223321-0030310112323132-0231321222131320-3312031132033112-3132313133113113-2231330233000111) |
| `aws.byoc.connections.auth_key.blindfold_secret_info.store_provider` | [aws.byoc.connections.auth_key.blindfold_secret_info.store_provider](resources--cloud_link--reference--group-001.md#canonical-1122120121102311-0112332013031023-1112232110012212-2300030020012301-2300232203002123-2322021303312323-3122020011312120-0022112031213003) |
| `aws.byoc.connections.auth_key.clear_secret_info` | [aws.byoc.connections.auth_key.clear_secret_info](resources--cloud_link--reference--group-001.md#canonical-1321031201323031-1302133231001000-3002012200310301-1302010322200112-3210220132210013-2310333013030302-0112022222220032-2311211103111102) |
| `aws.byoc.connections.auth_key.clear_secret_info.provider_ref` | [aws.byoc.connections.auth_key.clear_secret_info.provider_ref](resources--cloud_link--reference--group-001.md#canonical-3322132220031021-2023131301023331-2121220213112032-1220201120320031-1110300033231313-3313211222303213-1303210310100323-1101331322120233) |
| `aws.byoc.connections.auth_key.clear_secret_info.url` | [aws.byoc.connections.auth_key.clear_secret_info.url](resources--cloud_link--reference--group-001.md#canonical-0210101301112032-0120010313300323-1032211100320002-0031200322122120-0321033113330121-2000312113312110-0233230312310030-1103103003012130) |
| `aws.byoc.connections.bgp_asn` | [aws.byoc.connections.bgp_asn](resources--cloud_link--reference--group-001.md#canonical-0032213123000001-0311130020203113-0331121131233003-3132330120013220-2110311200022322-3033120110013120-3103231212321130-0223110001233012) |
| `aws.byoc.connections.connection_id` | [aws.byoc.connections.connection_id](resources--cloud_link--reference--group-001.md#canonical-1313312210100201-3220212100311312-2001110302233123-0312222230332123-0223231120033333-2133322022222322-3022020222021223-3332121102322032) |
| `aws.byoc.connections.ipv4` | [aws.byoc.connections.ipv4](resources--cloud_link--reference--group-001.md#canonical-2032322223222333-0300220221112323-1302311011123000-0130232231302322-3320230310031322-2112112001202112-0203203023001310-1132202230221232) |
| `aws.byoc.connections.ipv4.aws_router_peer_address` | [aws.byoc.connections.ipv4.aws_router_peer_address](resources--cloud_link--reference--group-001.md#canonical-0112201321201002-3031010030312012-0211120322033013-3312122110203211-2031223030010333-1203230121322130-0231322003330203-2123020010122123) |
| `aws.byoc.connections.ipv4.router_peer_address` | [aws.byoc.connections.ipv4.router_peer_address](resources--cloud_link--reference--group-001.md#canonical-0212310200233231-3322131301221321-0010130210000113-2121233002231033-1313033330212010-2012333001020200-1110010131023223-1200022002111033) |
| `aws.byoc.connections.metadata` | [aws.byoc.connections.metadata](resources--cloud_link--reference--group-001.md#canonical-3310123313221030-2303221020002000-2133022231032321-1033323200133123-1003200012221011-3212333110010011-2031211120033323-3123102220123211) |
| `aws.byoc.connections.metadata.description_spec` | [aws.byoc.connections.metadata.description_spec](resources--cloud_link--reference--group-001.md#canonical-1121110003223013-3002023201233333-2331310013210312-1230220023130322-1102013130003130-2300020210010023-1021022021000001-0022121202011301) |
| `aws.byoc.connections.metadata.name` | [aws.byoc.connections.metadata.name](resources--cloud_link--reference--group-001.md#canonical-1320131010301122-2032301230002231-2223231201013012-1033200320133012-3311303023212221-3323302202313233-0023323233131101-2220132203120012) |
| `aws.byoc.connections.region` | [aws.byoc.connections.region](resources--cloud_link--reference--group-001.md#canonical-1330010233000222-0332210031303121-2220012201022122-3303032022003331-0311210113122103-2000002131113201-1002012130200120-0332212011030223) |
| `aws.byoc.connections.system_generated_name` | [aws.byoc.connections.system_generated_name](resources--cloud_link--reference--group-001.md#canonical-2310203012133332-2130120010221012-1233303333003130-2122100122022300-0232012000031332-2122111122120002-3221103202210202-3111230231112220) |
| `aws.byoc.connections.tags` | [aws.byoc.connections.tags](resources--cloud_link--reference--group-001.md#canonical-0201332233212302-3321200113310233-0113312102202112-2213002323113310-0313230222132313-2133303122333101-0112133211112203-3302220030131031) |
| `aws.byoc.connections.user_assigned_name` | [aws.byoc.connections.user_assigned_name](resources--cloud_link--reference--group-001.md#canonical-2122201321312003-1012010130032321-3201313103121121-0030133302322221-1011303113302033-1231230130202221-3321223333132013-3221031223112102) |
| `aws.byoc.connections.virtual_interface_type` | [aws.byoc.connections.virtual_interface_type](resources--cloud_link--reference--group-001.md#canonical-1130103133231001-2212332301131223-0122013123200211-1102321323131102-3011031233301332-0212323200313110-3331111232313202-2102012101320302) |
| `aws.byoc.connections.vlan` | [aws.byoc.connections.vlan](resources--cloud_link--reference--group-001.md#canonical-1101022332223013-1312022022201010-2212013322001021-1131231330322312-3320000122113112-2230133320023223-2331333211312302-3231001100222112) |
| `aws.custom_asn` | [aws.custom_asn](resources--cloud_link--reference--group-001.md#canonical-3101223102032322-1001131013322303-1022112022021220-3202031033303013-0302020322321323-2002323201312031-1113021203112313-0230221232213300) |
| `description` | [description](resources--cloud_link--reference--group-001.md#canonical-3210200032230100-1213321020323220-0121321032210301-1133213312223301-2021313202120013-2123111113130221-2131220222333131-3032032320122203) |
| `disable` | [disable](resources--cloud_link--reference--group-001.md#canonical-1101322110110201-3212202113330331-0202223033122310-3212232120011123-2212212101303002-2011231312221220-3112203023111020-0012011132000112) |
| `disabled` | [disabled](resources--cloud_link--reference--group-001.md#canonical-1300023102330311-2213330331111333-0000312032231223-0000003330322130-1201313121133210-3202033222320332-2030012020301222-0302323213001331) |
| `enabled` | [enabled](resources--cloud_link--reference--group-001.md#canonical-1011102030130012-1323200331330210-2101200102102213-0130223201332012-1203103323331231-0231122103221012-0103320333013230-1102121301210011) |
| `enabled.cloudlink_network_name` | [enabled.cloudlink_network_name](resources--cloud_link--reference--group-001.md#canonical-3112130113130302-2010201302133302-1212222122130210-3130323132002220-3133121021010013-3323001121232311-0003233313211323-3020101103111320) |
| `gcp` | [gcp](resources--cloud_link--reference--group-001.md#canonical-0020202103220233-2301032030020003-2001201122010011-0003122303130120-0032302020103123-1313100213311333-0003312021333312-2323211022220201) |
| `gcp.byoc` | [gcp.byoc](resources--cloud_link--reference--group-001.md#canonical-0203111220311011-2010230223002321-0230322212111111-3020313113332321-2320000231101221-2233330320232321-2311032112201013-3012233331212001) |
| `gcp.byoc.connections` | [gcp.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-1331122223322010-1112223031030312-3130021221110223-3202111321213133-0233310101300023-2211301320013232-0102103002222113-1231021020322032) |
| `gcp.byoc.connections.interconnect_attachment_name` | [gcp.byoc.connections.interconnect_attachment_name](resources--cloud_link--reference--group-001.md#canonical-1320323133001232-0232230210211023-3103103320133203-2211221013330123-1133210233100021-0020222223321131-2031130212310101-3123333230012110) |
| `gcp.byoc.connections.metadata` | [gcp.byoc.connections.metadata](resources--cloud_link--reference--group-001.md#canonical-2332033132112212-1210003333112013-0010130232332212-1230210322301310-2321203120220200-0320110032111323-1010001300303222-3321220312103120) |
| `gcp.byoc.connections.metadata.description_spec` | [gcp.byoc.connections.metadata.description_spec](resources--cloud_link--reference--group-001.md#canonical-3122232012030023-0321311322003302-2321132133212020-1210011023000032-1033302111323312-1113130122321311-1122232231012300-1213310012122213) |
| `gcp.byoc.connections.metadata.name` | [gcp.byoc.connections.metadata.name](resources--cloud_link--reference--group-001.md#canonical-1003110021203231-2132122132010123-3023020130020120-2120102233230302-1010121301213133-2223211221303031-0021320201331303-2032312113302221) |
| `gcp.byoc.connections.project` | [gcp.byoc.connections.project](resources--cloud_link--reference--group-001.md#canonical-1000011331230332-2001113032123100-1200022030010211-3303230232021031-2011002121230230-1031111210331200-0033022112132112-0131203220231210) |
| `gcp.byoc.connections.region` | [gcp.byoc.connections.region](resources--cloud_link--reference--group-001.md#canonical-1033030330032132-2232300220301023-3232030313103030-3122031200221113-3332133330300131-2121131123121030-2021231001003101-1110301311133200) |
| `gcp.byoc.connections.same_as_credential` | [gcp.byoc.connections.same_as_credential](resources--cloud_link--reference--group-001.md#canonical-0122201002122132-2321300121112020-2222332020131213-2202000313011213-3133123200231021-0111103331023120-1323132320132333-3123000320000121) |
| `gcp.gcp_cred` | [gcp.gcp_cred](resources--cloud_link--reference--group-001.md#canonical-0002230030330100-1210320203120013-2321121010301233-3211000030332011-1010013132201013-0211233102111221-2231000313110113-2330333302110212) |
| `gcp.gcp_cred.name` | [gcp.gcp_cred.name](resources--cloud_link--reference--group-001.md#canonical-2300032232021230-0230313232031321-2112220013302112-0212102230332220-0112010111032030-2210203232210133-1321330320302212-0012231322322000) |
| `gcp.gcp_cred.namespace` | [gcp.gcp_cred.namespace](resources--cloud_link--reference--group-001.md#canonical-3023223113020103-3010313101233021-0202330301223000-3221321222322030-1123123220220202-3002101210321123-3000003201110032-1313332122120031) |
| `gcp.gcp_cred.tenant` | [gcp.gcp_cred.tenant](resources--cloud_link--reference--group-001.md#canonical-2112000120133012-1212130301301321-1012013211210211-2223302113220012-1303130000100031-0131310331101213-1200201330312030-2103112111000311) |
| `id` | [ID](resources--cloud_link--reference--group-001.md#canonical-3332231312322102-0313303311113120-1030231023323210-0230112012203132-0102022322313202-1332302110320313-1100133001302313-3001221101211303) |
| `labels` | [labels](resources--cloud_link--reference--group-001.md#canonical-0330321222200303-1111032021200133-1201023123231220-1311320222303202-3112112133130233-0230202023102030-3022110103300012-1120330212030332) |
| `name` | [name](resources--cloud_link--reference--group-001.md#canonical-3221000132113302-1122120202212013-1131200300001030-2333120111300101-2000030202031300-0010121020233231-0112303202300110-0323220033002320) |
| `namespace` | [namespace](resources--cloud_link--reference--group-001.md#canonical-0300330002321121-0033032101232002-0303112112120133-3321210120002330-1111312022311212-1312130302113022-2013032200323223-2011332131112332) |
| `timeouts` | [timeouts](resources--cloud_link--reference--group-001.md#canonical-1101221033113122-3131232332310222-2230302030330032-3222311202120012-3310213132101102-3301300112112203-3333333031111202-1223320310003233) |
| `timeouts.create` | [timeouts.create](resources--cloud_link--reference--group-001.md#canonical-1211000320332331-1311320121313023-1203033022003030-3130020012212220-2302100120231223-3120212203102103-0202131011200021-2323023212010210) |
| `timeouts.delete` | [timeouts.delete](resources--cloud_link--reference--group-001.md#canonical-1203330013323320-1102133132122222-2213031113323303-1031123321003323-0112011201302232-0233200020202130-2021021033020300-1201213212221031) |
| `timeouts.read` | [timeouts.read](resources--cloud_link--reference--group-001.md#canonical-0101133123203102-2011101112221032-2032313310102011-1301020023332122-3332303110131023-1201010022011101-0210031223102301-0210230101000022) |
| `timeouts.update` | [timeouts.update](resources--cloud_link--reference--group-001.md#canonical-1201130222112201-0123320202131133-1303210100030010-3131102000122200-1313122000021220-3312303312223123-3021131200013221-2103322023020133) |

<a id="canonical-2122012130133020-2032323112221322-1211111102003323-1300012000223020-1120220220211310-0223321133201021-1300111122111112-2011112300122012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws` properties

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311)
- aws

<a id="canonical-1002232311000300-0122311212200020-2301013330301220-3220323120032002-1120132211223201-2120131110022330-1301200133303220-0120110221333130"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: aws, gcp\] Amazon Web Services(AWS) CloudLink Provider. CloudLink for AWS Cloud Provider.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cloud_link_type": "[\"byoc\"]",
  "x-ves-oneof-field-direct_connect_gateway_asn_choice": "[\"custom_asn\"]"
}
```

OneOf alternatives in this subsection:

- [aws](resources--cloud_link--reference--group-001.md#canonical-1002232311000300-0122311212200020-2301013330301220-3220323120032002-1120132211223201-2120131110022330-1301200133303220-0120110221333130)
- [gcp](resources--cloud_link--reference--group-001.md#canonical-0020202103220233-2301032030020003-2001201122010011-0003122303130120-0032302020103123-1313100213311333-0003312021333312-2323211022220201)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
aws {
  # Configure direct properties listed below.
}
```

<a id="canonical-2302110311213021-0001322133032321-1112312111133020-3210113002221323-0222032312111101-2031331200013310-0230330100312023-3213210212000303"></a>

### Direct properties for `aws`

- [aws_cred](resources--cloud_link--reference--group-001.md#canonical-0323112032320101-1012021201210020-3000221330223010-2232021222213203-2311001330103203-2301223111222001-3031133231200210-1123020330323203): complete subsection reference.

- [byoc](resources--cloud_link--reference--group-001.md#canonical-2233321131011111-2001322330011202-3013000322213232-0323113233313302-1232101031313200-0230202230022321-1301101331100322-0213130120231202): complete subsection reference.

<a id="canonical-3101223102032322-1001131013322303-1022112022021220-3202031033303013-0302020322321323-2002323201312031-1113021203112313-0230221232213300"></a>

<a id="canonical-2301303310131333-3132011003323330-3003223123001300-3020332300200222-1330333320031010-2131232222220013-0311121302010302-2020012121110102"></a>

#### `aws.custom_asn` property

Type: `"number"`. Optional.

Exclusive with \[\] F5XC will use custom ASN to create a Direct Connect Gateway.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 64512, Maximum: 65534},
    validators.Int64Range{Minimum: 4200000000, Maximum: 4294967294},
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
    "maximum": 4294967294,
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
    "ves.io.schema.rules.uint32.ranges": "64512-65534, 4200000000-4294967294"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "64512-65534, 4200000000-4294967294"
  }
}
```

<a id="canonical-0323112032320101-1012021201210020-3000221330223010-2232021222213203-2311001330103203-2301223111222001-3031133231200210-1123020330323203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.aws_cred` properties

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311)
- [aws](resources--cloud_link--reference--group-001.md#canonical-2122012130133020-2032323112221322-1211111102003323-1300012000223020-1120220220211310-0223321133201021-1300111122111112-2011112300122012)
- aws.aws_cred

<a id="canonical-1011302120131021-2212331122112102-1101333323310203-1012030321311330-0103332011021323-2302012001101223-2131200322013222-1122122311021311"></a>

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
aws_cred {
  # Configure direct properties listed below.
}
```

<a id="canonical-3200203203100031-1030003130120120-2013303232111230-0222032210003011-0312322212303032-0321200223013022-0223121031221300-1301032102221103"></a>

### Direct properties for `aws.aws_cred`

<a id="canonical-2222322303010112-2311321333120310-1230332333103032-0033232302300311-1311332320022210-1033210032013110-0211212311230030-0213212302003111"></a>

#### `aws.aws_cred.name` property

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

<a id="canonical-3231022221123001-0132110120113232-3112223230020110-1120022212301112-2122011202102220-3130000300021331-3310300301211331-2112031311131333"></a>

<a id="canonical-1320123200233131-1112103000320112-2321230131231210-1222000133101301-3000111202110300-1320132310220122-2300213321031121-2300322030310211"></a>

#### `aws.aws_cred.namespace` property

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

<a id="canonical-1303110300332012-3012011211212021-1322020003222300-0123031013003201-2213210110121212-0010330033021332-2010222110311021-1033001320223200"></a>

<a id="canonical-2210021013210221-2010223131331103-0300210132213230-3130103300012312-3110131012131331-3210112010123023-2312130001030132-0233233003120332"></a>

#### `aws.aws_cred.tenant` property

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

<a id="canonical-2233321131011111-2001322330011202-3013000322213232-0323113233313302-1232101031313200-0230202230022321-1301101331100322-0213130120231202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.byoc` properties

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311)
- [aws](resources--cloud_link--reference--group-001.md#canonical-2122012130133020-2032323112221322-1211111102003323-1300012000223020-1120220220211310-0223321133201021-1300111122111112-2011112300122012)
- aws.byoc

<a id="canonical-3303113222033001-2232322100312322-0120010032123313-0031221123033231-0213211133113211-3230013311123330-2300000000222331-1301001211212202"></a>

Type: `"object"`. single nested block, Optional.

Bring Your Own Connections. List of Bring You Own Connection.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("connections")}
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
byoc {
  # Configure direct properties listed below.
}
```

<a id="canonical-2103133230030130-1113230002203312-0031031113030023-2100131023030210-1303300313313321-2000003203133011-0333333301323321-0030012000330111"></a>

### Direct properties for `aws.byoc`

- [connections](resources--cloud_link--reference--group-001.md#canonical-1023130221100320-0102323203033021-0313310312213332-2102300233220103-0233031333221203-1120322110031212-0121133300101023-3331230300212001): complete subsection reference.

<a id="canonical-1023130221100320-0102323203033021-0313310312213332-2102300233220103-0233031333221203-1120322110031212-0121133300101023-3331230300212001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.byoc.connections` properties

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311)
- [aws](resources--cloud_link--reference--group-001.md#canonical-2122012130133020-2032323112221322-1211111102003323-1300012000223020-1120220220211310-0223321133201021-1300111122111112-2011112300122012)
- [aws.byoc](resources--cloud_link--reference--group-001.md#canonical-2233321131011111-2001322330011202-3013000322213232-0323113233313302-1232101031313200-0230202230022321-1301101331100322-0213130120231202)
- aws.byoc.connections

<a id="canonical-2013221322121002-1020010230230202-3100112020003011-0332310302233011-0322210133112013-2010313111010230-2130223133302003-1231210011000233"></a>

Type: `"object"`. list nested block, Optional.

List of Bring You Own Connections. These AWS Direct Connect connections are not managed by F5XC but
will be used for connecting sites and REs.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("bgp_asn",
    "connection_id",
    "region",
    "vlan"),
  validators.ConflictingListObjectAttributes("system_generated_name",
    "user_assigned_name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 10,
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
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
connections {
  # Configure direct properties listed below.
}
```

<a id="canonical-1313302013330122-1323222111213320-0212030301301320-0000301232222222-0212001312130233-1032132230023132-3321231013101231-2011320100113001"></a>

### Direct properties for `aws.byoc.connections`

- [auth_key](resources--cloud_link--reference--group-001.md#canonical-3312321013123103-2313001210302022-2321001102333301-1011203300221103-0210100000232231-2012112022101010-2102132321212121-2011132021320031): complete subsection reference.

<a id="canonical-0032213123000001-0311130020203113-0331121131233003-3132330120013220-2110311200022322-3033120110013120-3103231212321130-0223110001233012"></a>

<a id="canonical-1122201233320231-2301012230020202-2103120103010001-1003212230330102-0020301100232123-2121022313021031-2222220100022130-3202331100131010"></a>

#### `aws.byoc.connections.bgp_asn` property

Type: `"number"`. Optional.

The Border Gateway Protocol (BGP) Autonomous System Number (ASN) of your on-premises router for the
new virtual interface to be configured on AWS.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 2147483647),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="canonical-1313312210100201-3220212100311312-2001110302233123-0312222230332123-0223231120033333-2133322022222322-3022020222021223-3332121102322032"></a>

<a id="canonical-1100023020100011-0121230211202110-3221210121133213-3203110221223023-0021132012132103-3220001013130231-3323210122131311-2311220200202213"></a>

#### `aws.byoc.connections.connection_id` property

Type: `"string"`. Optional.

ID of the existing AWS Direct Connect Connection.

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "pattern": "^(dxcon-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(dxcon-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(dxcon-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

- [IPv4](resources--cloud_link--reference--group-001.md#canonical-0302023103302130-0112101233231230-3203233032101133-1032133221323000-3100323202130322-3211310322230233-2233020131331030-3320031233333020): complete subsection reference.

- [metadata](resources--cloud_link--reference--group-001.md#canonical-3101033123112001-2031201223322300-1122112301102320-1220233032023011-0133301031211120-0120122122331011-1012000112000221-0200101110202333): complete subsection reference.

<a id="canonical-1330010233000222-0332210031303121-2220012201022122-3303032022003331-0311210113122103-2000002131113201-1002012130200120-0332212011030223"></a>

<a id="canonical-2032023001113132-0310202031330210-3122223011000130-3120003332331030-2223212133200112-1202230300121330-3113020232113133-1120321021003103"></a>

#### `aws.byoc.connections.region` property

Type: `"string"`. Optional.

\[Enum:
ap-northeast-1|ap-southeast-1|eu-central-1|eu-west-1|eu-west-3|sa-east-1|us-east-1|us-east-2|us-west-2|ca-central-1|af-south-1|ap-east-1|ap-south-1|ap-northeast-2|ap-southeast-2|eu-south-1|eu-north-1|eu-west-2|me-south-1|us-west-1|ap-southeast-3\]
Region. Region where the connection is setup. Possible values are \`ap-northeast-1\`,
\`ap-southeast-1\`, \`eu-central-1\`, \`eu-west-1\`, \`eu-west-3\`, \`sa-east-1\`, \`us-east-1\`,
\`us-east-2\`, \`us-west-2\`, \`ca-central-1\`, \`af-south-1\`, \`ap-east-1\`, \`ap-south-1\`,
\`ap-northeast-2\`, \`ap-southeast-2\`, \`eu-south-1\`, \`eu-north-1\`, \`eu-west-2\`,
\`me-south-1\`, \`us-west-1\`, \`ap-southeast-3\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["af-south-1","ap-east-1","ap-northeast-1","ap-northeast-2","ap-south-1","ap-southeast-1","ap-southeast-2","ap-southeast-3","ca-central-1","eu-central-1","eu-north-1","eu-south-1","eu-west-1","eu-west-2","eu-west-3","me-south-1","sa-east-1","us-east-1","us-east-2","us-west-1","us-west-2"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("ap-northeast-1",
    "ap-southeast-1",
    "eu-central-1",
    "eu-west-1",
    "eu-west-3",
    "sa-east-1",
    "us-east-1",
    "us-east-2",
    "us-west-2",
    "ca-central-1",
    "af-south-1",
    "ap-east-1",
    "ap-south-1",
    "ap-northeast-2",
    "ap-southeast-2",
    "eu-south-1",
    "eu-north-1",
    "eu-west-2",
    "me-south-1",
    "us-west-1",
    "ap-southeast-3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ap-northeast-1",
    "ap-southeast-1",
    "eu-central-1",
    "eu-west-1",
    "eu-west-3",
    "sa-east-1",
    "us-east-1",
    "us-east-2",
    "us-west-2",
    "ca-central-1",
    "af-south-1",
    "ap-east-1",
    "ap-south-1",
    "ap-northeast-2",
    "ap-southeast-2",
    "eu-south-1",
    "eu-north-1",
    "eu-west-2",
    "me-south-1",
    "us-west-1",
    "ap-southeast-3"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"ap-northeast-1\\\",\\\"ap-southeast-1\\\",\\\"eu-central-1\\\",\\\"eu-west-1\\\",\\\"eu-west-3\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-2\\\",\\\"ca-central-1\\\",\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-south-1\\\",\\\"ap-northeast-2\\\",\\\"ap-southeast-2\\\",\\\"eu-south-1\\\",\\\"eu-north-1\\\",\\\"eu-west-2\\\",\\\"me-south-1\\\",\\\"us-west-1\\\",\\\"ap-southeast-3\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ap-northeast-1\\\",\\\"ap-southeast-1\\\",\\\"eu-central-1\\\",\\\"eu-west-1\\\",\\\"eu-west-3\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-2\\\",\\\"ca-central-1\\\",\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-south-1\\\",\\\"ap-northeast-2\\\",\\\"ap-southeast-2\\\",\\\"eu-south-1\\\",\\\"eu-north-1\\\",\\\"eu-west-2\\\",\\\"me-south-1\\\",\\\"us-west-1\\\",\\\"ap-southeast-3\\\"]"
  }
}
```

- [system_generated_name](resources--cloud_link--reference--group-001.md#canonical-3112123022131020-2031211021331013-1230000030031201-3112332010322101-3123310213321302-0230223330213222-1220113133003333-2233230301222321): complete subsection reference.

<a id="canonical-0201332233212302-3321200113310233-0113312102202112-2213002323113310-0313230222132313-2133303122333101-0112133211112203-3302220030131031"></a>

<a id="canonical-0103322233303303-1230222201022210-0221120232311313-0001323033202012-1113310330033121-3322202200330302-0030212330132120-0132131031230200"></a>

#### `aws.byoc.connections.tags` property

Type: `["map", "string"]`. Optional.

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console. Specified tags will be added to Virtual
interface along with any F5XC specific tags.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":40},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"40\",\"ves.io.schema.rules.map.values.string.max_len\":\"256\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":256,\"minLength\":1,\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 40
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "40",
      "ves.io.schema.rules.map.values.string.max_len": "256",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 256,
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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "256",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "256",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-2122201321312003-1012010130032321-3201313103121121-0030133302322221-1011303113302033-1231230130202221-3321223333132013-3221031223112102"></a>

<a id="canonical-2121132100002020-2201011132033221-3131202000002000-1330233332002220-2331233113201201-1201100233301232-1132133300232331-0103110131132333"></a>

#### `aws.byoc.connections.user_assigned_name` property

Type: `"string"`. Optional.

Exclusive with \[system\_generated\_name\] User is managing the AWS resource name.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1130103133231001-2212332301131223-0122013123200211-1102321323131102-3011031233301332-0212323200313110-3331111232313202-2102012101320302"></a>

<a id="canonical-1213132021313130-0322222321313130-1133203222323212-2131002110321101-1322010100312312-1113330210322101-0103200112032312-1023330003230311"></a>

#### `aws.byoc.connections.virtual_interface_type` property

Type: `"string"`. Optional.

\[Enum: PRIVATE\] Defines the type of virtual interface that needs to be configured on AWS -
PRIVATE: Private A private virtual interface should be used to access an Amazon VPC using private IP
addresses. - TRANSIT: Transit A transit virtual interface is a VLAN that transports traffic from a
Direct Connect.. The only possible value is \`PRIVATE\`. Defaults to \`PRIVATE\`.

Additional upstream details:

Defines the type of virtual interface that needs to be configured on AWS

&#8203;- PRIVATE: Private

A private virtual interface should be used to access an Amazon VPC using private IP addresses.
&#8203;- TRANSIT: Transit

A transit virtual interface is a VLAN that transports traffic from a Direct Connect gateway to one
or more transit gateways.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["PRIVATE"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("PRIVATE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "PRIVATE",
  "enum": [
    "PRIVATE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1101022332223013-1312022022201010-2212013322001021-1131231330322312-3320000122113112-2230133320023223-2331333211312302-3231001100222112"></a>

<a id="canonical-2133110230023000-3002302010112220-1302102331321301-2300101123231130-2023323033201330-3322113133112330-3031230133003012-2222102220131220"></a>

#### `aws.byoc.connections.vlan` property

Type: `"number"`. Optional.

Virtual Local Area Network number for the new virtual interface to be configured on the AWS. This
tag is required for any traffic traversing the AWS Direct Connect connection.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 4094),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 4094,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4094"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4094"
  }
}
```

<a id="canonical-3312321013123103-2313001210302022-2321001102333301-1011203300221103-0210100000232231-2012112022101010-2102132321212121-2011132021320031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.byoc.connections.auth_key` properties

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311)
- [aws](resources--cloud_link--reference--group-001.md#canonical-2122012130133020-2032323112221322-1211111102003323-1300012000223020-1120220220211310-0223321133201021-1300111122111112-2011112300122012)
- [aws.byoc](resources--cloud_link--reference--group-001.md#canonical-2233321131011111-2001322330011202-3013000322213232-0323113233313302-1232101031313200-0230202230022321-1301101331100322-0213130120231202)
- [aws.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-1023130221100320-0102323203033021-0313310312213332-2102300233220103-0233031333221203-1120322110031212-0121133300101023-3331230300212001)
- aws.byoc.connections.auth_key

<a id="canonical-2123131322000303-3203301103111133-0313220003202222-3002022112002020-1222122321033231-1321320223120222-2332001030121311-1131221231220223"></a>

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
auth_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-1212013113023013-2303010031130021-3333101213231222-3203120331320131-0021302020221032-3001223001030103-3221303203313303-0020102021003002"></a>

### Direct properties for `aws.byoc.connections.auth_key`

- [blindfold_secret_info](resources--cloud_link--reference--group-001.md#canonical-1101213132201311-1122100123323211-0200100023322011-0112331132012202-1102111113212231-0302313200302333-0100012231031311-3321000130131010): complete subsection reference.

- [clear_secret_info](resources--cloud_link--reference--group-001.md#canonical-2321010012302131-3013230130310311-2131211021031330-1001312012130310-1131323103010111-0001312110132022-1032232100032102-0101013133031230): complete subsection reference.

<a id="canonical-1101213132201311-1122100123323211-0200100023322011-0112331132012202-1102111113212231-0302313200302333-0100012231031311-3321000130131010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.byoc.connections.auth_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311)
- [aws](resources--cloud_link--reference--group-001.md#canonical-2122012130133020-2032323112221322-1211111102003323-1300012000223020-1120220220211310-0223321133201021-1300111122111112-2011112300122012)
- [aws.byoc](resources--cloud_link--reference--group-001.md#canonical-2233321131011111-2001322330011202-3013000322213232-0323113233313302-1232101031313200-0230202230022321-1301101331100322-0213130120231202)
- [aws.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-1023130221100320-0102323203033021-0313310312213332-2102300233220103-0233031333221203-1120322110031212-0121133300101023-3331230300212001)
- [aws.byoc.connections.auth_key](resources--cloud_link--reference--group-001.md#canonical-3312321013123103-2313001210302022-2321001102333301-1011203300221103-0210100000232231-2012112022101010-2102132321212121-2011132021320031)
- aws.byoc.connections.auth_key.blindfold_secret_info

<a id="canonical-0233122312311111-3313103001121222-3222101102212130-3210013233220201-3332330302130212-2032333311011012-2222331202220302-2110210210021121"></a>

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

<a id="canonical-3200303200130231-2033320001202133-3001223020032031-2023323001313323-0011221012113100-2301100303032310-3310013220023123-0033222111112120"></a>

### Direct properties for `aws.byoc.connections.auth_key.blindfold_secret_info`

<a id="canonical-1122003012110300-3020331002102003-3310100313311201-0110003010221222-2223312322313203-0330102010332111-0323303301313221-1131113131332011"></a>

#### `aws.byoc.connections.auth_key.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-0203230103131012-0202212300121203-3000201001223321-0030310112323132-0231321222131320-3312031132033112-3132313133113113-2231330233000111"></a>

<a id="canonical-3023301112222011-0330311002222101-3220211102333011-1121121131133222-2000123023232103-2222312031330102-3101203023201221-0223220030002302"></a>

#### `aws.byoc.connections.auth_key.blindfold_secret_info.location` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1122120121102311-0112332013031023-1112232110012212-2300030020012301-2300232203002123-2322021303312323-3122020011312120-0022112031213003"></a>

<a id="canonical-3321230230021023-1021321311103020-0123023221031312-0230101032122233-2300111230310133-0230212201013320-3221112300221231-3033032232210131"></a>

#### `aws.byoc.connections.auth_key.blindfold_secret_info.store_provider` property

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

<a id="canonical-2321010012302131-3013230130310311-2131211021031330-1001312012130310-1131323103010111-0001312110132022-1032232100032102-0101013133031230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.byoc.connections.auth_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311)
- [aws](resources--cloud_link--reference--group-001.md#canonical-2122012130133020-2032323112221322-1211111102003323-1300012000223020-1120220220211310-0223321133201021-1300111122111112-2011112300122012)
- [aws.byoc](resources--cloud_link--reference--group-001.md#canonical-2233321131011111-2001322330011202-3013000322213232-0323113233313302-1232101031313200-0230202230022321-1301101331100322-0213130120231202)
- [aws.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-1023130221100320-0102323203033021-0313310312213332-2102300233220103-0233031333221203-1120322110031212-0121133300101023-3331230300212001)
- [aws.byoc.connections.auth_key](resources--cloud_link--reference--group-001.md#canonical-3312321013123103-2313001210302022-2321001102333301-1011203300221103-0210100000232231-2012112022101010-2102132321212121-2011132021320031)
- aws.byoc.connections.auth_key.clear_secret_info

<a id="canonical-1321031201323031-1302133231001000-3002012200310301-1302010322200112-3210220132210013-2310333013030302-0112022222220032-2311211103111102"></a>

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

<a id="canonical-3311100112320120-1032113200303223-2320311120012203-0031333213322301-3011102311330221-1123311023101102-2300023101230210-3032110020100222"></a>

### Direct properties for `aws.byoc.connections.auth_key.clear_secret_info`

<a id="canonical-3322132220031021-2023131301023331-2121220213112032-1220201120320031-1110300033231313-3313211222303213-1303210310100323-1101331322120233"></a>

#### `aws.byoc.connections.auth_key.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0210101301112032-0120010313300323-1032211100320002-0031200322122120-0321033113330121-2000312113312110-0233230312310030-1103103003012130"></a>

<a id="canonical-0201312103301211-1201000333222000-0202013303000130-1112302330100133-2001033101232330-3221320112211232-2231201301322123-1021131002220300"></a>

#### `aws.byoc.connections.auth_key.clear_secret_info.url` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0302023103302130-0112101233231230-3203233032101133-1032133221323000-3100323202130322-3211310322230233-2233020131331030-3320031233333020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.byoc.connections.ipv4` properties

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311)
- [aws](resources--cloud_link--reference--group-001.md#canonical-2122012130133020-2032323112221322-1211111102003323-1300012000223020-1120220220211310-0223321133201021-1300111122111112-2011112300122012)
- [aws.byoc](resources--cloud_link--reference--group-001.md#canonical-2233321131011111-2001322330011202-3013000322213232-0323113233313302-1232101031313200-0230202230022321-1301101331100322-0213130120231202)
- [aws.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-1023130221100320-0102323203033021-0313310312213332-2102300233220103-0233031333221203-1120322110031212-0121133300101023-3331230300212001)
- aws.byoc.connections.IPv4

<a id="canonical-2032322223222333-0300220221112323-1302311011123000-0130232231302322-3320230310031322-2112112001202112-0203203023001310-1132202230221232"></a>

Type: `"object"`. single nested block, Optional.

Configure BGP IPv4 peering for endpoints.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("aws_router_peer_address",
    "router_peer_address")}
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
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-3312302200012002-0221012000220213-2202012230100023-3132122103333320-0023200220333121-2302103320231203-3310223211103032-2212110213030332"></a>

### Direct properties for `aws.byoc.connections.ipv4`

<a id="canonical-0112201321201002-3031010030312012-0211120322033013-3312122110203211-2031223030010333-1203230121322130-0231322003330203-2123020010122123"></a>

#### `aws.byoc.connections.ipv4.aws_router_peer_address` property

Type: `"string"`. Optional.

The BGP peer IP configured on the AWS endpoint.

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
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "32",
    "ves.io.schema.rules.string.min_ip_prefix_length": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "32",
    "ves.io.schema.rules.string.min_ip_prefix_length": "1"
  }
}
```

<a id="canonical-0212310200233231-3322131301221321-0010130210000113-2121233002231033-1313033330212010-2012333001020200-1110010131023223-1200022002111033"></a>

<a id="canonical-1322021322321010-0232313130321033-0323201102330110-2100121213012211-0011012133231011-1130130012321133-0100013110320221-2230123103032303"></a>

#### `aws.byoc.connections.ipv4.router_peer_address` property

Type: `"string"`. Optional.

The BGP peer IP configured on your (customer) endpoint.

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
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "32",
    "ves.io.schema.rules.string.min_ip_prefix_length": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "32",
    "ves.io.schema.rules.string.min_ip_prefix_length": "1"
  }
}
```

<a id="canonical-3101033123112001-2031201223322300-1122112301102320-1220233032023011-0133301031211120-0120122122331011-1012000112000221-0200101110202333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.byoc.connections.metadata` properties

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311)
- [aws](resources--cloud_link--reference--group-001.md#canonical-2122012130133020-2032323112221322-1211111102003323-1300012000223020-1120220220211310-0223321133201021-1300111122111112-2011112300122012)
- [aws.byoc](resources--cloud_link--reference--group-001.md#canonical-2233321131011111-2001322330011202-3013000322213232-0323113233313302-1232101031313200-0230202230022321-1301101331100322-0213130120231202)
- [aws.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-1023130221100320-0102323203033021-0313310312213332-2102300233220103-0233031333221203-1120322110031212-0121133300101023-3331230300212001)
- aws.byoc.connections.metadata

<a id="canonical-3310123313221030-2303221020002000-2133022231032321-1033323200133123-1003200012221011-3212333110010011-2031211120033323-3123102220123211"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-2132331113231302-2333330021331331-1301133301123211-3120020233031011-2232222202203131-0112311322300330-1310311031112303-3330032102213031"></a>

### Direct properties for `aws.byoc.connections.metadata`

<a id="canonical-1121110003223013-3002023201233333-2331310013210312-1230220023130322-1102013130003130-2300020210010023-1021022021000001-0022121202011301"></a>

#### `aws.byoc.connections.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-1320131010301122-2032301230002231-2223231201013012-1033200320133012-3311303023212221-3323302202313233-0023323233131101-2220132203120012"></a>

<a id="canonical-3301221132112022-2100132100320101-3232000203022121-3112211020201030-3123310221223331-3121031333302320-3202310010002323-0113312121220123"></a>

#### `aws.byoc.connections.metadata.name` property

Type: `"string"`. Optional.

This is the name of the message. The value of name has to follow DNS-1035 format.

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
  "minLength": 1,
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-3112123022131020-2031211021331013-1230000030031201-3112332010322101-3123310213321302-0230223330213222-1220113133003333-2233230301222321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aws.byoc.connections.system_generated_name` properties

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311)
- [aws](resources--cloud_link--reference--group-001.md#canonical-2122012130133020-2032323112221322-1211111102003323-1300012000223020-1120220220211310-0223321133201021-1300111122111112-2011112300122012)
- [aws.byoc](resources--cloud_link--reference--group-001.md#canonical-2233321131011111-2001322330011202-3013000322213232-0323113233313302-1232101031313200-0230202230022321-1301101331100322-0213130120231202)
- [aws.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-1023130221100320-0102323203033021-0313310312213332-2102300233220103-0233031333221203-1120322110031212-0121133300101023-3331230300212001)
- aws.byoc.connections.system_generated_name

<a id="canonical-2310203012133332-2130120010221012-1233303333003130-2122100122022300-0232012000031332-2122111122120002-3221103202210202-3111230231112220"></a>

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
system_generated_name = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1022131120022203-1330132311200122-3203320013133123-1312200301110003-2220313021320223-0321311322113312-1231023302122033-0222221310010130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disabled` properties

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311)
- disabled

<a id="canonical-1300023102330311-2213330331111333-0000312032231223-0000003330322130-1201313121133210-3202033222320332-2030012020301222-0302323213001331"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disabled, enabled\] Enable this option

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

- [disabled](resources--cloud_link--reference--group-001.md#canonical-1300023102330311-2213330331111333-0000312032231223-0000003330322130-1201313121133210-3202033222320332-2030012020301222-0302323213001331)
- [enabled](resources--cloud_link--reference--group-001.md#canonical-1011102030130012-1323200331330210-2101200102102213-0130223201332012-1203103323331231-0231122103221012-0103320333013230-1102121301210011)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2320130011300110-3021312032122332-3011210131012030-1310310010301231-2131130022302300-3011131030311231-0221010200100112-3232312022020312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enabled` properties

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311)
- enabled

<a id="canonical-1011102030130012-1323200331330210-2101200102102213-0130223201332012-1203103323331231-0231122103221012-0103320333013230-1102121301210011"></a>

Type: `"object"`. single nested block, Optional.

CloudLink ADN Network Config.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("cloudlink_network_name")}
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
enabled {
  # Configure direct properties listed below.
}
```

<a id="canonical-2002220202233210-1220210300212020-0002313130013313-1222211233013112-3310032311213313-0211310001032223-3321020301202013-1200312021301300"></a>

### Direct properties for `enabled`

<a id="canonical-3112130113130302-2010201302133302-1212222122130210-3130323132002220-3133121021010013-3323001121232311-0003233313211323-3020101103111320"></a>

#### `enabled.cloudlink_network_name` property

Type: `"string"`. Optional.

Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN
network. To provision a Private ADN network, please contact F5 Distributed Cloud support.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2120230010303101-2202120333222110-3132020103101203-0002303133323002-1131210100111130-2330332130032310-1203120020022120-2302330200333121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp` properties

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311)
- gcp

<a id="canonical-0020202103220233-2301032030020003-2001201122010011-0003122303130120-0032302020103123-1313100213311333-0003312021333312-2323211022220201"></a>

Type: `"object"`. single nested block, Optional.

Google Cloud Platform (GCP) CloudLink Provider. CloudLink for GCP Cloud Provider.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cloud_link_type": "[\"byoc\"]"
}
```

Terraform syntax:

```terraform
gcp {
  # Configure direct properties listed below.
}
```

<a id="canonical-2132010202331210-2020011023222303-1001001220112331-1111103220001301-1213031020302310-2203031300101320-1000122011303021-3132103210231322"></a>

### Direct properties for `gcp`

- [byoc](resources--cloud_link--reference--group-001.md#canonical-3203032012032130-2203123221123112-2333003311102003-3330133201333003-3210212310211012-3233001103310111-0132322112122332-1332312020101333): complete subsection reference.

- [gcp_cred](resources--cloud_link--reference--group-001.md#canonical-3123231132230302-2000202222031322-1313322313312231-0011203110123103-3123132132132123-1100101001001230-1032310011123330-0300211020002302): complete subsection reference.

<a id="canonical-3203032012032130-2203123221123112-2333003311102003-3330133201333003-3210212310211012-3233001103310111-0132322112122332-1332312020101333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.byoc` properties

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311)
- [gcp](resources--cloud_link--reference--group-001.md#canonical-2120230010303101-2202120333222110-3132020103101203-0002303133323002-1131210100111130-2330332130032310-1203120020022120-2302330200333121)
- gcp.byoc

<a id="canonical-0203111220311011-2010230223002321-0230322212111111-3020313113332321-2320000231101221-2233330320232321-2311032112201013-3012233331212001"></a>

Type: `"object"`. single nested block, Optional.

GCP Bring Your Own Connections. List of GCP Bring You Own Connections.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("connections")}
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
byoc {
  # Configure direct properties listed below.
}
```

<a id="canonical-0130130220023113-0103312302000223-2022012032130203-1221120322302322-3213132001333231-3032010113321113-1032321022013012-0033033330030013"></a>

### Direct properties for `gcp.byoc`

- [connections](resources--cloud_link--reference--group-001.md#canonical-3300300311011020-1232112233202120-3121231302233032-0323330120012132-1203212002211002-0212331312203112-2320223102010211-0101312032120201): complete subsection reference.

<a id="canonical-3300300311011020-1232112233202120-3121231302233032-0323330120012132-1203212002211002-0212331312203112-2320223102010211-0101312032120201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.byoc.connections` properties

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311)
- [gcp](resources--cloud_link--reference--group-001.md#canonical-2120230010303101-2202120333222110-3132020103101203-0002303133323002-1131210100111130-2330332130032310-1203120020022120-2302330200333121)
- [gcp.byoc](resources--cloud_link--reference--group-001.md#canonical-3203032012032130-2203123221123112-2333003311102003-3330133201333003-3210212310211012-3233001103310111-0132322112122332-1332312020101333)
- gcp.byoc.connections

<a id="canonical-1331122223322010-1112223031030312-3130021221110223-3202111321213133-0233310101300023-2211301320013232-0102103002222113-1231021020322032"></a>

Type: `"object"`. list nested block, Optional.

Each 'Bring Your Own Connection' represents a virtual connection that the customer has provisioned
in the Cloud (example: AWS Direct Connect). F5XC will orchestrate networking resources in the cloud
to facilitate seamless private connectivity.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("interconnect_attachment_name",
    "region"),
  validators.ConflictingListObjectAttributes("project",
    "same_as_credential")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 10,
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
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
connections {
  # Configure direct properties listed below.
}
```

<a id="canonical-0220113130002220-2202211322331111-2122113101122110-0100003201122123-1202110031030023-2230130320310200-1132131103301321-2300100201101123"></a>

### Direct properties for `gcp.byoc.connections`

<a id="canonical-1320323133001232-0232230210211023-3103103320133203-2211221013330123-1133210233100021-0020222223321131-2031130212310101-3123333230012110"></a>

#### `gcp.byoc.connections.interconnect_attachment_name` property

Type: `"string"`. Optional.

Name of already-existing GCP Cloud Interconnect Attachment.

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
  "maxLength": 63,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 63,
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
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [metadata](resources--cloud_link--reference--group-001.md#canonical-1231033111003012-3023121311330213-1113012103223011-3332101032200000-3321331130231323-0023222021213213-1202002212230131-3011212301103210): complete subsection reference.

<a id="canonical-1000011331230332-2001113032123100-1200022030010211-3303230232021031-2011002121230230-1031111210331200-0033022112132112-0131203220231210"></a>

<a id="canonical-2031213212000131-3010201301203213-3021330333230000-2302303220132111-1313021230133303-3113030312100102-3223323020200223-1033022011232220"></a>

#### `gcp.byoc.connections.project` property

Type: `"string"`. Optional.

Exclusive with \[same\_as\_credential\] Specify a GCP Project for the interconnect attachment.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 30),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 30,
  "minLength": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 30,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "30",
    "ves.io.schema.rules.string.min_len": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "30",
    "ves.io.schema.rules.string.min_len": "4"
  }
}
```

<a id="canonical-1033030330032132-2232300220301023-3232030313103030-3122031200221113-3332133330300131-2121131123121030-2021231001003101-1110301311133200"></a>

<a id="canonical-0333133032132111-2311002130303222-2000113130230033-2100231110132113-2003321030120230-0022313311300310-0212303231013200-3103300332230131"></a>

#### `gcp.byoc.connections.region` property

Type: `"string"`. Optional.

\[Enum:
asia-east1|asia-east2|asia-northeast1|asia-northeast2|asia-northeast3|asia-southeast1|asia-southeast2|europe-central2|europe-north1|europe-west1|europe-west2|europe-west3|europe-west4|europe-west6|europe-west8|europe-west9|europe-west10|europe-west12|europe-southwest1|me-west1|me-central1|me-central2|northamerica-northeast1|northamerica-northeast2|us-central1|us-east1|us-east4|us-east5|us-south1|us-west1|us-west2|us-west3|us-west4|southamerica-east1|southamerica-west1|australia-southeast1|australia-southeast2|asia-south1|asia-south2\]
GCP Region in which the GCP Cloud Interconnect attachment is configured. Possible values are
\`asia-east1\`, \`asia-east2\`, \`asia-northeast1\`, \`asia-northeast2\`, \`asia-northeast3\`,
\`asia-southeast1\`, \`asia-southeast2\`, \`europe-central2\`, \`europe-north1\`, \`europe-west1\`,
\`europe-west2\`, \`europe-west3\`, \`europe-west4\`, \`europe-west6\`, \`europe-west8\`,
\`europe-west9\`, \`europe-west10\`, \`europe-west12\`, \`europe-southwest1\`, \`me-west1\`,
\`me-central1\`, \`me-central2\`, \`northamerica-northeast1\`, \`northamerica-northeast2\`,
\`us-central1\`, \`us-east1\`, \`us-east4\`, \`us-east5\`, \`us-south1\`, \`us-west1\`,
\`us-west2\`, \`us-west3\`, \`us-west4\`, \`southamerica-east1\`, \`southamerica-west1\`,
\`australia-southeast1\`, \`australia-southeast2\`, \`asia-south1\`, \`asia-south2\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["asia-east1","asia-east2","asia-northeast1","asia-northeast2","asia-northeast3","asia-south1","asia-south2","asia-southeast1","asia-southeast2","australia-southeast1","australia-southeast2","europe-central2","europe-north1","europe-southwest1","europe-west1","europe-west10","europe-west12","europe-west2","europe-west3","europe-west4","europe-west6","europe-west8","europe-west9","me-central1","me-central2","me-west1","northamerica-northeast1","northamerica-northeast2","southamerica-east1","southamerica-west1","us-central1","us-east1","us-east4","us-east5","us-south1","us-west1","us-west2","us-west3","us-west4"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("asia-east1",
    "asia-east2",
    "asia-northeast1",
    "asia-northeast2",
    "asia-northeast3",
    "asia-southeast1",
    "asia-southeast2",
    "europe-central2",
    "europe-north1",
    "europe-west1",
    "europe-west2",
    "europe-west3",
    "europe-west4",
    "europe-west6",
    "europe-west8",
    "europe-west9",
    "europe-west10",
    "europe-west12",
    "europe-southwest1",
    "me-west1",
    "me-central1",
    "me-central2",
    "northamerica-northeast1",
    "northamerica-northeast2",
    "us-central1",
    "us-east1",
    "us-east4",
    "us-east5",
    "us-south1",
    "us-west1",
    "us-west2",
    "us-west3",
    "us-west4",
    "southamerica-east1",
    "southamerica-west1",
    "australia-southeast1",
    "australia-southeast2",
    "asia-south1",
    "asia-south2"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "asia-east1",
    "asia-east2",
    "asia-northeast1",
    "asia-northeast2",
    "asia-northeast3",
    "asia-southeast1",
    "asia-southeast2",
    "europe-central2",
    "europe-north1",
    "europe-west1",
    "europe-west2",
    "europe-west3",
    "europe-west4",
    "europe-west6",
    "europe-west8",
    "europe-west9",
    "europe-west10",
    "europe-west12",
    "europe-southwest1",
    "me-west1",
    "me-central1",
    "me-central2",
    "northamerica-northeast1",
    "northamerica-northeast2",
    "us-central1",
    "us-east1",
    "us-east4",
    "us-east5",
    "us-south1",
    "us-west1",
    "us-west2",
    "us-west3",
    "us-west4",
    "southamerica-east1",
    "southamerica-west1",
    "australia-southeast1",
    "australia-southeast2",
    "asia-south1",
    "asia-south2"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"asia-east1\\\",\\\"asia-east2\\\",\\\"asia-northeast1\\\",\\\"asia-northeast2\\\",\\\"asia-northeast3\\\",\\\"asia-southeast1\\\",\\\"asia-southeast2\\\",\\\"europe-central2\\\",\\\"europe-north1\\\",\\\"europe-west1\\\",\\\"europe-west2\\\",\\\"europe-west3\\\",\\\"europe-west4\\\",\\\"europe-west6\\\",\\\"europe-west8\\\",\\\"europe-west9\\\",\\\"europe-west10\\\",\\\"europe-west12\\\",\\\"europe-southwest1\\\",\\\"me-west1\\\",\\\"me-central1\\\",\\\"me-central2\\\",\\\"northamerica-northeast1\\\",\\\"northamerica-northeast2\\\",\\\"us-central1\\\",\\\"us-east1\\\",\\\"us-east4\\\",\\\"us-east5\\\",\\\"us-south1\\\",\\\"us-west1\\\",\\\"us-west2\\\",\\\"us-west3\\\",\\\"us-west4\\\",\\\"southamerica-east1\\\",\\\"southamerica-west1\\\",\\\"australia-southeast1\\\",\\\"australia-southeast2\\\",\\\"asia-south1\\\",\\\"asia-south2\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"asia-east1\\\",\\\"asia-east2\\\",\\\"asia-northeast1\\\",\\\"asia-northeast2\\\",\\\"asia-northeast3\\\",\\\"asia-southeast1\\\",\\\"asia-southeast2\\\",\\\"europe-central2\\\",\\\"europe-north1\\\",\\\"europe-west1\\\",\\\"europe-west2\\\",\\\"europe-west3\\\",\\\"europe-west4\\\",\\\"europe-west6\\\",\\\"europe-west8\\\",\\\"europe-west9\\\",\\\"europe-west10\\\",\\\"europe-west12\\\",\\\"europe-southwest1\\\",\\\"me-west1\\\",\\\"me-central1\\\",\\\"me-central2\\\",\\\"northamerica-northeast1\\\",\\\"northamerica-northeast2\\\",\\\"us-central1\\\",\\\"us-east1\\\",\\\"us-east4\\\",\\\"us-east5\\\",\\\"us-south1\\\",\\\"us-west1\\\",\\\"us-west2\\\",\\\"us-west3\\\",\\\"us-west4\\\",\\\"southamerica-east1\\\",\\\"southamerica-west1\\\",\\\"australia-southeast1\\\",\\\"australia-southeast2\\\",\\\"asia-south1\\\",\\\"asia-south2\\\"]"
  }
}
```

- [same_as_credential](resources--cloud_link--reference--group-001.md#canonical-0202331032230310-3101023230033102-3111320023333023-3220321311231102-3311231202011211-2122010332233212-1033013133032021-1121332032212113): complete subsection reference.

<a id="canonical-1231033111003012-3023121311330213-1113012103223011-3332101032200000-3321331130231323-0023222021213213-1202002212230131-3011212301103210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.byoc.connections.metadata` properties

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311)
- [gcp](resources--cloud_link--reference--group-001.md#canonical-2120230010303101-2202120333222110-3132020103101203-0002303133323002-1131210100111130-2330332130032310-1203120020022120-2302330200333121)
- [gcp.byoc](resources--cloud_link--reference--group-001.md#canonical-3203032012032130-2203123221123112-2333003311102003-3330133201333003-3210212310211012-3233001103310111-0132322112122332-1332312020101333)
- [gcp.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-3300300311011020-1232112233202120-3121231302233032-0323330120012132-1203212002211002-0212331312203112-2320223102010211-0101312032120201)
- gcp.byoc.connections.metadata

<a id="canonical-2332033132112212-1210003333112013-0010130232332212-1230210322301310-2321203120220200-0320110032111323-1010001300303222-3321220312103120"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-1333113302102320-0012331032031302-1020021222311202-1012323311203201-3200300212300331-1103022312031112-0330212023033012-0233001231101101"></a>

### Direct properties for `gcp.byoc.connections.metadata`

<a id="canonical-3122232012030023-0321311322003302-2321132133212020-1210011023000032-1033302111323312-1113130122321311-1122232231012300-1213310012122213"></a>

#### `gcp.byoc.connections.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-1003110021203231-2132122132010123-3023020130020120-2120102233230302-1010121301213133-2223211221303031-0021320201331303-2032312113302221"></a>

<a id="canonical-0303211001331030-1130333002301300-1222100011032333-0110001302033320-2130211020112220-2030231020011031-3312013100213120-3331212203102213"></a>

#### `gcp.byoc.connections.metadata.name` property

Type: `"string"`. Optional.

This is the name of the message. The value of name has to follow DNS-1035 format.

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
  "minLength": 1,
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-0202331032230310-3101023230033102-3111320023333023-3220321311231102-3311231202011211-2122010332233212-1033013133032021-1121332032212113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.byoc.connections.same_as_credential` properties

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311)
- [gcp](resources--cloud_link--reference--group-001.md#canonical-2120230010303101-2202120333222110-3132020103101203-0002303133323002-1131210100111130-2330332130032310-1203120020022120-2302330200333121)
- [gcp.byoc](resources--cloud_link--reference--group-001.md#canonical-3203032012032130-2203123221123112-2333003311102003-3330133201333003-3210212310211012-3233001103310111-0132322112122332-1332312020101333)
- [gcp.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-3300300311011020-1232112233202120-3121231302233032-0323330120012132-1203212002211002-0212331312203112-2320223102010211-0101312032120201)
- gcp.byoc.connections.same_as_credential

<a id="canonical-0122201002122132-2321300121112020-2222332020131213-2202000313011213-3133123200231021-0111103331023120-1323132320132333-3123000320000121"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for same as credential.

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
same_as_credential = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3123231132230302-2000202222031322-1313322313312231-0011203110123103-3123132132132123-1100101001001230-1032310011123330-0300211020002302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp.gcp_cred` properties

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311)
- [gcp](resources--cloud_link--reference--group-001.md#canonical-2120230010303101-2202120333222110-3132020103101203-0002303133323002-1131210100111130-2330332130032310-1203120020022120-2302330200333121)
- gcp.gcp_cred

<a id="canonical-0002230030330100-1210320203120013-2321121010301233-3211000030332011-1010013132201013-0211233102111221-2231000313110113-2330333302110212"></a>

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
gcp_cred {
  # Configure direct properties listed below.
}
```

<a id="canonical-2022202313223212-3101323211202200-2220133123111220-2103121000222211-2223023333310301-1120022232333132-0211200121013311-3311201203202131"></a>

### Direct properties for `gcp.gcp_cred`

<a id="canonical-2300032232021230-0230313232031321-2112220013302112-0212102230332220-0112010111032030-2210203232210133-1321330320302212-0012231322322000"></a>

#### `gcp.gcp_cred.name` property

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

<a id="canonical-3023223113020103-3010313101233021-0202330301223000-3221321222322030-1123123220220202-3002101210321123-3000003201110032-1313332122120031"></a>

<a id="canonical-0203010301121013-2131202323301112-1303300223303212-0032021220233322-3323203211023001-3221000300130301-0233011303100203-0132211131131121"></a>

#### `gcp.gcp_cred.namespace` property

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

<a id="canonical-2112000120133012-1212130301301321-1012013211210211-2223302113220012-1303130000100031-0131310331101213-1200201330312030-2103112111000311"></a>

<a id="canonical-1010302010201122-0023021322031013-1133333110022122-2022303302011130-0012331310021001-3300031123003211-2332331120031121-0032110121222002"></a>

#### `gcp.gcp_cred.tenant` property

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

<a id="canonical-1000132120000000-2030223002311110-0323103213230233-0103003302023032-3221033120201230-3112000110012223-3112303220021202-1023123321332000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311)
- timeouts

<a id="canonical-1101221033113122-3131232332310222-2230302030330032-3222311202120012-3310213132101102-3301300112112203-3333333031111202-1223320310003233"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3223200122233300-2210122133333112-3210002033112311-0100103122112332-3222022000311221-3202000223031233-3030232102301231-3312213012133213"></a>

### Direct properties for `timeouts`

<a id="canonical-1211000320332331-1311320121313023-1203033022003030-3130020012212220-2302100120231223-3120212203102103-0202131011200021-2323023212010210"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1203330013323320-1102133132122222-2213031113323303-1031123321003323-0112011201302232-0233200020202130-2021021033020300-1201213212221031"></a>

<a id="canonical-0020020301233202-1333222010003220-1022300323331301-0322101101300223-0233321210031211-2300031123020201-1312010120102330-3203012210220111"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0101133123203102-2011101112221032-2032313310102011-1301020023332122-3332303110131023-1201010022011101-0210031223102301-0210230101000022"></a>

<a id="canonical-3200000103233310-3321112002001110-3013122102032110-1221031221133003-1300312030003113-1320302023122333-2213130220231233-0032322000023211"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1201130222112201-0123320202131133-1303210100030010-3131102000122200-1313122000021220-3312303312223123-3021131200013221-2103322023020133"></a>

<a id="canonical-2032112102122302-2213212320211122-0331201223013131-0121123330302131-2132022123030303-2030113222213322-2202313021321022-3223233133013302"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
