---
page_title: "xcsh_cloud_link reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_link reference."
---

# xcsh_cloud_link reference

<a id="canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303103302111301-0313032023300121-0033321121122232-2010131331000111-3131103121301032-2121031033231331-1330122211200321-2311330233331001"></a>

## Property reference — Property reference / 011210300113 / 2

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)
- Property reference

<a id="canonical-1122002101122200-2321133303221012-3333310330302330-1030302330323102-0201132332201032-2101132000332201-1113013031330332-3210021001023200"></a>

## Direct properties — Property reference / 011210300113 / 3

<a id="canonical-3022203033020230-3120302010013332-2312300333323132-1210333231030013-1122221301310310-0322202223322221-3333300203210323-2110321223100012"></a>

<a id="canonical-3012212111212201-2033133232323110-0321303033132033-2023231311122320-1003000113100222-1123013011202102-1113022221123103-2000110313332332"></a>

## annotations property — Property reference / 011210300113 / 4

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

Upstream description:

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
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

<a id="canonical-0013302233130301-0332313031010012-3200033020031320-0323202223132200-2031111010030103-0132222210011132-0312133100232313-2211230210111132"></a>

## description property — Property reference / 011210300113 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2102300323112032-0312231120302230-3320300003022220-2331001332033013-0113030013010111-2331130223332310-3232302121221330-0020332021322000"></a>

## disable property — Property reference / 011210300113 / 6

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

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

<a id="canonical-0102122323122331-2110300300112002-1132033113110331-0310013013022312-0303110201211130-1131131112021011-2221023222301233-0201032232030310"></a>

## ID property — Property reference / 011210300113 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0330321222200303-1111032021200133-1201023123231220-1311320222303202-3112112133130233-0230202023102030-3022110103300012-1120330212030332"></a>

<a id="canonical-2113021330211013-0010123000322221-2231233031030121-2100123223030011-3013213022032323-2100312002202312-1330201222023211-0223232100003220"></a>

## labels property — Property reference / 011210300113 / 8

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

Upstream description:

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

<a id="canonical-3211123100211231-0330220321202201-3121223332101232-1123323013301320-3000231302310032-1023110000202301-3310121003031321-0312200021023222"></a>

## name property — Property reference / 011210300113 / 9

Type: `"string"`. Required.

Name of the Cloud Link. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2231122211210322-1223101301323301-1222332210022022-3102133033102220-0013233110010123-3303103111012202-2230212000012103-2102133231111120"></a>

## namespace property — Property reference / 011210300113 / 10

Type: `"string"`. Required.

Namespace where the Cloud Link is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3203211221300022-2331113131320032-3223223002231013-3021110000323323-2121000022121022-2300123302010303-1121333023123313-2101012211230103"></a>

## All schema paths — Property reference / 011210300113 / 11

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

<a id="canonical-3013233101130120-2321131231120033-2000222101321310-2120302231003332-2020023103233022-1300113202223102-2023311123133323-1110120330131013"></a>

## Next pages — Property reference / 011210300113 / 12

- [aws](resources--cloud_link--reference--group-001.md#canonical-2122012130133020-2032323112221322-1211111102003323-1300012000223020-1120220220211310-0223321133201021-1300111122111112-2011112300122012)
- [disabled](resources--cloud_link--reference--group-001.md#canonical-1022131120022203-1330132311200122-3203320013133123-1312200301110003-2220313021320223-0321311322113312-1231023302122033-0222221310010130)
- [enabled](resources--cloud_link--reference--group-001.md#canonical-2320130011300110-3021312032122332-3011210131012030-1310310010301231-2131130022302300-3011131030311231-0221010200100112-3232312022020312)
- [gcp](resources--cloud_link--reference--group-001.md#canonical-2120230010303101-2202120333222110-3132020103101203-0002303133323002-1131210100111130-2330332130032310-1203120020022120-2302330200333121)
- [timeouts](resources--cloud_link--reference--group-001.md#canonical-1000132120000000-2030223002311110-0323103213230233-0103003302023032-3221033120201230-3112000110012223-3112303220021202-1023123321332000)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)

<a id="canonical-2122012130133020-2032323112221322-1211111102003323-1300012000223020-1120220220211310-0223321133201021-1300111122111112-2011112300122012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302110311213021-0001322133032321-1112312111133020-3210113002221323-0222032312111101-2031331200013310-0230330100312023-3213210212000303"></a>

## aws — aws / 130010332130 / 2

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311)
- aws

<a id="canonical-1002232311000300-0122311212200020-2301013330301220-3220323120032002-1120132211223201-2120131110022330-1301200133303220-0120110221333130"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: aws, gcp\] Amazon Web Services(AWS) CloudLink Provider. CloudLink for AWS Cloud Provider.

Upstream description:

CloudLink for AWS Cloud Provider.

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

<a id="canonical-2301303310131333-3132011003323330-3003223123001300-3020332300200222-1330333320031010-2131232222220013-0311121302010302-2020012121110102"></a>

## Direct properties — aws / 130010332130 / 3

- [aws_cred](resources--cloud_link--reference--group-001.md#canonical-0323112032320101-1012021201210020-3000221330223010-2232021222213203-2311001330103203-2301223111222001-3031133231200210-1123020330323203): complete subsection reference.

- [byoc](resources--cloud_link--reference--group-001.md#canonical-2233321131011111-2001322330011202-3013000322213232-0323113233313302-1232101031313200-0230202230022321-1301101331100322-0213130120231202): complete subsection reference.

<a id="canonical-3101223102032322-1001131013322303-1022112022021220-3202031033303013-0302020322321323-2002323201312031-1113021203112313-0230221232213300"></a>

<a id="canonical-1131201131312310-1110130330113222-3313330322122131-2012202101232131-1100322322322330-2133332211220120-0323033031302120-1132212221102022"></a>

## custom_asn property — aws / 130010332130 / 4

Type: `"number"`. Optional.

Exclusive with \[\] F5XC will use custom ASN to create a Direct Connect Gateway.

Upstream description:

Exclusive with \[\] F5XC will use custom ASN to create a Direct Connect Gateway.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3322112313001031-0333323222321121-2330303332301322-1131330223310120-2033003131110101-2212100011011100-1310003020110123-2122101221203203"></a>

## Next pages — aws / 130010332130 / 5

- [aws.aws_cred](resources--cloud_link--reference--group-001.md#canonical-0323112032320101-1012021201210020-3000221330223010-2232021222213203-2311001330103203-2301223111222001-3031133231200210-1123020330323203)
- [aws.byoc](resources--cloud_link--reference--group-001.md#canonical-2233321131011111-2001322330011202-3013000322213232-0323113233313302-1232101031313200-0230202230022321-1301101331100322-0213130120231202)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)

<a id="canonical-0323112032320101-1012021201210020-3000221330223010-2232021222213203-2311001330103203-2301223111222001-3031133231200210-1123020330323203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200203203100031-1030003130120120-2013303232111230-0222032210003011-0312322212303032-0321200223013022-0223121031221300-1301032102221103"></a>

## aws.aws_cred — aws_cred / 213133121012 / 2

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311)
- [aws](resources--cloud_link--reference--group-001.md#canonical-2122012130133020-2032323112221322-1211111102003323-1300012000223020-1120220220211310-0223321133201021-1300111122111112-2011112300122012)
- aws.aws_cred

<a id="canonical-1011302120131021-2212331122112102-1101333323310203-1012030321311330-0103332011021323-2302012001101223-2131200322013222-1122122311021311"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1320123200233131-1112103000320112-2321230131231210-1222000133101301-3000111202110300-1320132310220122-2300213321031121-2300322030310211"></a>

## Direct properties — aws_cred / 213133121012 / 3

<a id="canonical-2222322303010112-2311321333120310-1230332333103032-0033232302300311-1311332320022210-1033210032013110-0211212311230030-0213212302003111"></a>

<a id="canonical-2210021013210221-2010223131331103-0300210132213230-3130103300012312-3110131012131331-3210112010123023-2312130001030132-0233233003120332"></a>

## name property — aws_cred / 213133121012 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0113300212332022-1313200233030131-2220201303113322-3131320211222323-0020122211133300-3130230223020223-2202233300031331-3113121200201101"></a>

## namespace property — aws_cred / 213133121012 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1301120300101230-0232103203032003-1301231330102302-0120031213230323-0110322023301031-2130323300033110-3321221021302233-2003300312213322"></a>

## tenant property — aws_cred / 213133121012 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2321131130233303-3232333222103021-2033010010300303-2200330222202020-0222033103131131-0020100230020201-1021133113233002-3302101030213333"></a>

## Next pages — aws_cred / 213133121012 / 7

- [aws](resources--cloud_link--reference--group-001.md#canonical-2122012130133020-2032323112221322-1211111102003323-1300012000223020-1120220220211310-0223321133201021-1300111122111112-2011112300122012)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)

<a id="canonical-2233321131011111-2001322330011202-3013000322213232-0323113233313302-1232101031313200-0230202230022321-1301101331100322-0213130120231202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103133230030130-1113230002203312-0031031113030023-2100131023030210-1303300313313321-2000003203133011-0333333301323321-0030012000330111"></a>

## aws.byoc — byoc / 102010100023 / 2

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311)
- [aws](resources--cloud_link--reference--group-001.md#canonical-2122012130133020-2032323112221322-1211111102003323-1300012000223020-1120220220211310-0223321133201021-1300111122111112-2011112300122012)
- aws.byoc

<a id="canonical-3303113222033001-2232322100312322-0120010032123313-0031221123033231-0213211133113211-3230013311123330-2300000000222331-1301001211212202"></a>

Type: `"object"`. single nested block, Optional.

Bring Your Own Connections. List of Bring You Own Connection.

Upstream description:

List of Bring You Own Connection.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2213300300300003-1233003331112031-0223330013210103-1310123020101330-3000211310201213-1013132013001203-3103330030300220-3002111003333103"></a>

## Direct properties — byoc / 102010100023 / 3

- [connections](resources--cloud_link--reference--group-001.md#canonical-1023130221100320-0102323203033021-0313310312213332-2102300233220103-0233031333221203-1120322110031212-0121133300101023-3331230300212001): complete subsection reference.

<a id="canonical-1011130012100123-3230213010002302-1101311132321032-3212232001011131-2010330222333332-2121302200201233-3003103313313012-3313000321330002"></a>

## Next pages — byoc / 102010100023 / 4

- [aws.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-1023130221100320-0102323203033021-0313310312213332-2102300233220103-0233031333221203-1120322110031212-0121133300101023-3331230300212001)
- [aws](resources--cloud_link--reference--group-001.md#canonical-2122012130133020-2032323112221322-1211111102003323-1300012000223020-1120220220211310-0223321133201021-1300111122111112-2011112300122012)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)

<a id="canonical-1023130221100320-0102323203033021-0313310312213332-2102300233220103-0233031333221203-1120322110031212-0121133300101023-3331230300212001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1313302013330122-1323222111213320-0212030301301320-0000301232222222-0212001312130233-1032132230023132-3321231013101231-2011320100113001"></a>

## aws.byoc.connections — connections / 322101300120 / 2

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1122201233320231-2301012230020202-2103120103010001-1003212230330102-0020301100232123-2121022313021031-2222220100022130-3202331100131010"></a>

## Direct properties — connections / 322101300120 / 3

- [auth_key](resources--cloud_link--reference--group-001.md#canonical-3312321013123103-2313001210302022-2321001102333301-1011203300221103-0210100000232231-2012112022101010-2102132321212121-2011132021320031): complete subsection reference.

<a id="canonical-0032213123000001-0311130020203113-0331121131233003-3132330120013220-2110311200022322-3033120110013120-3103231212321130-0223110001233012"></a>

<a id="canonical-1100023020100011-0121230211202110-3221210121133213-3203110221223023-0021132012132103-3220001013130231-3323210122131311-2311220200202213"></a>

## bgp_asn property — connections / 322101300120 / 4

Type: `"number"`. Optional.

The Border Gateway Protocol (BGP) Autonomous System Number (ASN) of your on-premises router for the
new virtual interface to be configured on AWS.

Upstream description:

The Border Gateway Protocol (BGP) Autonomous System Number (ASN) of your on-premises router for the
new virtual interface to be configured on AWS.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2032023001113132-0310202031330210-3122223011000130-3120003332331030-2223212133200112-1202230300121330-3113020232113133-1120321021003103"></a>

## connection_id property — connections / 322101300120 / 5

Type: `"string"`. Optional.

ID of the existing AWS Direct Connect Connection.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0103322233303303-1230222201022210-0221120232311313-0001323033202012-1113310330033121-3322202200330302-0030212330132120-0132131031230200"></a>

## region property — connections / 322101300120 / 6

Type: `"string"`. Optional.

\[Enum:
ap-northeast-1|ap-southeast-1|eu-central-1|eu-west-1|eu-west-3|sa-east-1|us-east-1|us-east-2|us-west-2|ca-central-1|af-south-1|ap-east-1|ap-south-1|ap-northeast-2|ap-southeast-2|eu-south-1|eu-north-1|eu-west-2|me-south-1|us-west-1|ap-southeast-3\]
Region. Region where the connection is setup. Possible values are \`ap-northeast-1\`,
\`ap-southeast-1\`, \`eu-central-1\`, \`eu-west-1\`, \`eu-west-3\`, \`sa-east-1\`, \`us-east-1\`,
\`us-east-2\`, \`us-west-2\`, \`ca-central-1\`, \`af-south-1\`, \`ap-east-1\`, \`ap-south-1\`,
\`ap-northeast-2\`, \`ap-southeast-2\`, \`eu-south-1\`, \`eu-north-1\`, \`eu-west-2\`,
\`me-south-1\`, \`us-west-1\`, \`ap-southeast-3\`.

Upstream description:

Region where the connection is setup.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2121132100002020-2201011132033221-3131202000002000-1330233332002220-2331233113201201-1201100233301232-1132133300232331-0103110131132333"></a>

## tags property — connections / 322101300120 / 7

Type: `["map", "string"]`. Optional.

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console. Specified tags will be added to Virtual
interface along with any F5XC specific tags.

Upstream description:

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console. Specified tags will be added to Virtual
interface along with any F5XC specific tags.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "object",
    "maxProperties": 40,
    "metadata": {
      "confidence": 0.75,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1213132021313130-0322222321313130-1133203222323212-2131002110321101-1322010100312312-1113330210322101-0103200112032312-1023330003230311"></a>

## user_assigned_name property — connections / 322101300120 / 8

Type: `"string"`. Optional.

Exclusive with \[system\_generated\_name\] User is managing the AWS resource name.

Upstream description:

Exclusive with \[system\_generated\_name\] User is managing the AWS resource name.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2133110230023000-3002302010112220-1302102331321301-2300101123231130-2023323033201330-3322113133112330-3031230133003012-2222102220131220"></a>

## virtual_interface_type property — connections / 322101300120 / 9

Type: `"string"`. Optional.

\[Enum: PRIVATE\] Defines the type of virtual interface that needs to be configured on AWS -
PRIVATE: Private A private virtual interface should be used to access an Amazon VPC using private IP
addresses. - TRANSIT: Transit A transit virtual interface is a VLAN that transports traffic from a
Direct Connect.. The only possible value is \`PRIVATE\`. Defaults to \`PRIVATE\`.

Upstream description:

Defines the type of virtual interface that needs to be configured on AWS

&#8203;- PRIVATE: Private

A private virtual interface should be used to access an Amazon VPC using private IP addresses.
&#8203;- TRANSIT: Transit

A transit virtual interface is a VLAN that transports traffic from a Direct Connect gateway to one
or more transit gateways.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3302322122323102-0121121113200000-2100123133120132-2111211200112111-0100321011010330-0320131203311000-1223001303300232-0003011110032103"></a>

## vlan property — connections / 322101300120 / 10

Type: `"number"`. Optional.

Virtual Local Area Network number for the new virtual interface to be configured on the AWS. This
tag is required for any traffic traversing the AWS Direct Connect connection.

Upstream description:

Virtual Local Area Network number for the new virtual interface to be configured on the AWS. This
tag is required for any traffic traversing the AWS Direct Connect connection.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3201300322200032-0022222113111030-2131000021221111-2200333231010022-2001201211230221-3010122132333220-3220303223303010-2102112012332102"></a>

## Next pages — connections / 322101300120 / 11

- [aws.byoc.connections.auth_key](resources--cloud_link--reference--group-001.md#canonical-3312321013123103-2313001210302022-2321001102333301-1011203300221103-0210100000232231-2012112022101010-2102132321212121-2011132021320031)
- [aws.byoc.connections.ipv4](resources--cloud_link--reference--group-001.md#canonical-0302023103302130-0112101233231230-3203233032101133-1032133221323000-3100323202130322-3211310322230233-2233020131331030-3320031233333020)
- [aws.byoc.connections.metadata](resources--cloud_link--reference--group-001.md#canonical-3101033123112001-2031201223322300-1122112301102320-1220233032023011-0133301031211120-0120122122331011-1012000112000221-0200101110202333)
- [aws.byoc.connections.system_generated_name](resources--cloud_link--reference--group-001.md#canonical-3112123022131020-2031211021331013-1230000030031201-3112332010322101-3123310213321302-0230223330213222-1220113133003333-2233230301222321)
- [aws.byoc](resources--cloud_link--reference--group-001.md#canonical-2233321131011111-2001322330011202-3013000322213232-0323113233313302-1232101031313200-0230202230022321-1301101331100322-0213130120231202)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)

<a id="canonical-3312321013123103-2313001210302022-2321001102333301-1011203300221103-0210100000232231-2012112022101010-2102132321212121-2011132021320031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212013113023013-2303010031130021-3333101213231222-3203120331320131-0021302020221032-3001223001030103-3221303203313303-0020102021003002"></a>

## aws.byoc.connections.auth_key — auth_key / 331101002020 / 2

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

<a id="canonical-1032311301323122-1333133202000020-1202112213003131-0303311331132113-3012311022131203-2310102331303120-3012012210221122-0132003303323211"></a>

## Direct properties — auth_key / 331101002020 / 3

- [blindfold_secret_info](resources--cloud_link--reference--group-001.md#canonical-1101213132201311-1122100123323211-0200100023322011-0112331132012202-1102111113212231-0302313200302333-0100012231031311-3321000130131010): complete subsection reference.

- [clear_secret_info](resources--cloud_link--reference--group-001.md#canonical-2321010012302131-3013230130310311-2131211021031330-1001312012130310-1131323103010111-0001312110132022-1032232100032102-0101013133031230): complete subsection reference.

<a id="canonical-1133101212302033-1111300112033300-2003103112222000-2232210231022221-3230122320113002-0213232031133220-2031002030101213-2032101302202001"></a>

## Next pages — auth_key / 331101002020 / 4

- [aws.byoc.connections.auth_key.blindfold_secret_info](resources--cloud_link--reference--group-001.md#canonical-1101213132201311-1122100123323211-0200100023322011-0112331132012202-1102111113212231-0302313200302333-0100012231031311-3321000130131010)
- [aws.byoc.connections.auth_key.clear_secret_info](resources--cloud_link--reference--group-001.md#canonical-2321010012302131-3013230130310311-2131211021031330-1001312012130310-1131323103010111-0001312110132022-1032232100032102-0101013133031230)
- [aws.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-1023130221100320-0102323203033021-0313310312213332-2102300233220103-0233031333221203-1120322110031212-0121133300101023-3331230300212001)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)

<a id="canonical-1101213132201311-1122100123323211-0200100023322011-0112331132012202-1102111113212231-0302313200302333-0100012231031311-3321000130131010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200303200130231-2033320001202133-3001223020032031-2023323001313323-0011221012113100-2301100303032310-3310013220023123-0033222111112120"></a>

## aws.byoc.connections.auth_key.blindfold_secret_info — blindfold_secret_info / 022022032000 / 2

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

<a id="canonical-3023301112222011-0330311002222101-3220211102333011-1121121131133222-2000123023232103-2222312031330102-3101203023201221-0223220030002302"></a>

## Direct properties — blindfold_secret_info / 022022032000 / 3

<a id="canonical-1122003012110300-3020331002102003-3310100313311201-0110003010221222-2223312322313203-0330102010332111-0323303301313221-1131113131332011"></a>

<a id="canonical-3321230230021023-1021321311103020-0123023221031312-0230101032122233-2300111230310133-0230212201013320-3221112300221231-3033032232210131"></a>

## decryption_provider property — blindfold_secret_info / 022022032000 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1002001103002122-2330201201301010-0130110311332113-3020230332023012-0100333130231023-0032200323003021-1030023012220130-2033002230212230"></a>

## location property — blindfold_secret_info / 022022032000 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2223302033033011-2201210220001111-0010230101132011-3232102020020201-2201123111211011-3203332003303301-1112233032201232-2321302133012222"></a>

## store_provider property — blindfold_secret_info / 022022032000 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1102120213310122-1131111120120033-1113302021331121-1202312021031120-2332211210011110-2121332220010020-0321203302131021-3132011112100001"></a>

## Next pages — blindfold_secret_info / 022022032000 / 7

- [aws.byoc.connections.auth_key](resources--cloud_link--reference--group-001.md#canonical-3312321013123103-2313001210302022-2321001102333301-1011203300221103-0210100000232231-2012112022101010-2102132321212121-2011132021320031)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)

<a id="canonical-2321010012302131-3013230130310311-2131211021031330-1001312012130310-1131323103010111-0001312110132022-1032232100032102-0101013133031230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311100112320120-1032113200303223-2320311120012203-0031333213322301-3011102311330221-1123311023101102-2300023101230210-3032110020100222"></a>

## aws.byoc.connections.auth_key.clear_secret_info — clear_secret_info / 101121023021 / 2

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

<a id="canonical-0201312103301211-1201000333222000-0202013303000130-1112302330100133-2001033101232330-3221320112211232-2231201301322123-1021131002220300"></a>

## Direct properties — clear_secret_info / 101121023021 / 3

<a id="canonical-3322132220031021-2023131301023331-2121220213112032-1220201120320031-1110300033231313-3313211222303213-1303210310100323-1101331322120233"></a>

<a id="canonical-3202310332312020-1032021011003133-0112331123133030-0120211121112011-0003320023110311-3013023102313332-2133310021113323-1020030231122222"></a>

## provider_ref property — clear_secret_info / 101121023021 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0210101301112032-0120010313300323-1032211100320002-0031200322122120-0321033113330121-2000312113312110-0233230312310030-1103103003012130"></a>

<a id="canonical-3001313120312223-1233212100200223-2020113012123031-2001312200313333-2120222011001232-1031020102102303-3221321220000331-0213330030232001"></a>

## URL property — clear_secret_info / 101121023021 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1323301100303003-3120322332213023-2300011033200230-1100012013101232-2003131221132222-1332320102303113-2010010230221202-2010312101000122"></a>

## Next pages — clear_secret_info / 101121023021 / 6

- [aws.byoc.connections.auth_key](resources--cloud_link--reference--group-001.md#canonical-3312321013123103-2313001210302022-2321001102333301-1011203300221103-0210100000232231-2012112022101010-2102132321212121-2011132021320031)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)

<a id="canonical-0302023103302130-0112101233231230-3203233032101133-1032133221323000-3100323202130322-3211310322230233-2233020131331030-3320031233333020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312302200012002-0221012000220213-2202012230100023-3132122103333320-0023200220333121-2302103320231203-3310223211103032-2212110213030332"></a>

## aws.byoc.connections.IPv4 — IPv4 / 330212120022 / 2

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

<a id="canonical-1322021322321010-0232313130321033-0323201102330110-2100121213012211-0011012133231011-1130130012321133-0100013110320221-2230123103032303"></a>

## Direct properties — IPv4 / 330212120022 / 3

<a id="canonical-0112201321201002-3031010030312012-0211120322033013-3312122110203211-2031223030010333-1203230121322130-0231322003330203-2123020010122123"></a>

<a id="canonical-1231022231011002-2223331012123331-2230002313233022-3221010202110100-3230031110101200-0031012223000332-1232122020222021-1021313121220210"></a>

## aws_router_peer_address property — IPv4 / 330212120022 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3333333110331001-2110112102122033-3230022333303220-2320132131232203-0113120223301330-3213100112021302-3213120321303201-3330030201113313"></a>

## router_peer_address property — IPv4 / 330212120022 / 5

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0000200030100333-2321222101020303-2002130010022013-1101100100321233-2101031230101330-0210001303033313-2110132012223032-0023312302003100"></a>

## Next pages — IPv4 / 330212120022 / 6

- [aws.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-1023130221100320-0102323203033021-0313310312213332-2102300233220103-0233031333221203-1120322110031212-0121133300101023-3331230300212001)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)

<a id="canonical-3101033123112001-2031201223322300-1122112301102320-1220233032023011-0133301031211120-0120122122331011-1012000112000221-0200101110202333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132331113231302-2333330021331331-1301133301123211-3120020233031011-2232222202203131-0112311322300330-1310311031112303-3330032102213031"></a>

## aws.byoc.connections.metadata — metadata / 001021310232 / 2

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
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3301221132112022-2100132100320101-3232000203022121-3112211020201030-3123310221223331-3121031333302320-3202310010002323-0113312121220123"></a>

## Direct properties — metadata / 001021310232 / 3

<a id="canonical-1121110003223013-3002023201233333-2331310013210312-1230220023130322-1102013130003130-2300020210010023-1021022021000001-0022121202011301"></a>

<a id="canonical-3020110322332010-0213300022202001-0012131301122313-0133223210012320-1003120231010111-0002022133232233-2322023010102331-2201023031133102"></a>

## description_spec property — metadata / 001021310232 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-1320131010301122-2032301230002231-2223231201013012-1033200320133012-3311303023212221-3323302202313233-0023323233131101-2220132203120012"></a>

<a id="canonical-0310231122112132-1100330133002312-2110223232300110-0023203111211023-3230023220122020-1020313211311300-1310332333302100-1020323300133320"></a>

## name property — metadata / 001021310232 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0202320313211323-3211130230021112-1021231211333021-0323232231210213-0212130303332322-2323100220201033-0312330300131020-2333330120320213"></a>

## Next pages — metadata / 001021310232 / 6

- [aws.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-1023130221100320-0102323203033021-0313310312213332-2102300233220103-0233031333221203-1120322110031212-0121133300101023-3331230300212001)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)

<a id="canonical-3112123022131020-2031211021331013-1230000030031201-3112332010322101-3123310213321302-0230223330213222-1220113133003333-2233230301222321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320201013211020-2301032103120001-3020312302233103-0011133332111212-1213011102101102-0232030110222113-2122102313331220-0221203022332203"></a>

## aws.byoc.connections.system_generated_name — system_generated_name / 020121101233 / 2

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

Upstream description:

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

<a id="canonical-1010030200031120-3000130110102103-2211210013102222-2032132110111121-2020023120100033-3000030201312330-0220110201232202-0323313100021022"></a>

## Direct properties — system_generated_name / 020121101233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2212312210301000-1211231303310332-1322232300003300-0220303013031221-1033100323220301-3221213111300130-3033201313001013-3000002111021223"></a>

## Next pages — system_generated_name / 020121101233 / 4

- [aws.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-1023130221100320-0102323203033021-0313310312213332-2102300233220103-0233031333221203-1120322110031212-0121133300101023-3331230300212001)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)

<a id="canonical-1022131120022203-1330132311200122-3203320013133123-1312200301110003-2220313021320223-0321311322113312-1231023302122033-0222221310010130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101323332102112-3322200000312221-1303322120213010-3131232013000000-2113322323110001-1333001223132122-2002313120300011-2221200122123202"></a>

## disabled — disabled / 330112120102 / 2

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311)
- disabled

<a id="canonical-1300023102330311-2213330331111333-0000312032231223-0000003330322130-1201313121133210-3202033222320332-2030012020301222-0302323213001331"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disabled, enabled\] Enable this option

Upstream description:

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

<a id="canonical-1310000311132222-2112320103002100-2022132120303332-1101033100220102-3303302321020030-3300233313221321-1133003113100021-1113322331222232"></a>

## Direct properties — disabled / 330112120102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3021112231010013-1312231103022010-0312131130230333-1002220000333102-0123230012331221-0102223103103213-2121111211223333-2012111220333220"></a>

## Next pages — disabled / 330112120102 / 4

- [Property reference](resources--cloud_link--reference--group-001.md#canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)

<a id="canonical-2320130011300110-3021312032122332-3011210131012030-1310310010301231-2131130022302300-3011131030311231-0221010200100112-3232312022020312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2002220202233210-1220210300212020-0002313130013313-1222211233013112-3310032311213313-0211310001032223-3321020301202013-1200312021301300"></a>

## enabled — enabled / 012300331323 / 2

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311)
- enabled

<a id="canonical-1011102030130012-1323200331330210-2101200102102213-0130223201332012-1203103323331231-0231122103221012-0103320333013230-1102121301210011"></a>

Type: `"object"`. single nested block, Optional.

CloudLink ADN Network Config.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0010022131103212-1100203321300310-2101233022121031-1023033320013113-2020120110233102-2200020003120313-0301302103121212-2203330313111120"></a>

## Direct properties — enabled / 012300331323 / 3

<a id="canonical-3112130113130302-2010201302133302-1212222122130210-3130323132002220-3133121021010013-3323001121232311-0003233313211323-3020101103111320"></a>

<a id="canonical-1011213321330130-2320300133103230-2003020311032222-0121322321210133-0030322012312030-3233032302011130-1001002323120103-1233010230222111"></a>

## cloudlink_network_name property — enabled / 012300331323 / 4

Type: `"string"`. Optional.

Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN
network. To provision a Private ADN network, please contact F5 Distributed Cloud support.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1032323332103111-1131213211320031-1110110322232121-0000210211130033-0020202311030313-2010201110003022-0301031211010221-2133311120103013"></a>

## Next pages — enabled / 012300331323 / 5

- [Property reference](resources--cloud_link--reference--group-001.md#canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)

<a id="canonical-2120230010303101-2202120333222110-3132020103101203-0002303133323002-1131210100111130-2330332130032310-1203120020022120-2302330200333121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132010202331210-2020011023222303-1001001220112331-1111103220001301-1213031020302310-2203031300101320-1000122011303021-3132103210231322"></a>

## gcp — gcp / 102122012220 / 2

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311)
- gcp

<a id="canonical-0020202103220233-2301032030020003-2001201122010011-0003122303130120-0032302020103123-1313100213311333-0003312021333312-2323211022220201"></a>

Type: `"object"`. single nested block, Optional.

Google Cloud Platform (GCP) CloudLink Provider. CloudLink for GCP Cloud Provider.

Upstream description:

CloudLink for GCP Cloud Provider.

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

<a id="canonical-0110333121031011-2322103102221330-2222213223110012-0200210020303123-1033131103011033-1203323230232210-3212330332311313-0232010200111032"></a>

## Direct properties — gcp / 102122012220 / 3

- [byoc](resources--cloud_link--reference--group-001.md#canonical-3203032012032130-2203123221123112-2333003311102003-3330133201333003-3210212310211012-3233001103310111-0132322112122332-1332312020101333): complete subsection reference.

- [gcp_cred](resources--cloud_link--reference--group-001.md#canonical-3123231132230302-2000202222031322-1313322313312231-0011203110123103-3123132132132123-1100101001001230-1032310011123330-0300211020002302): complete subsection reference.

<a id="canonical-0023301123030110-0010001313131111-1202001302121022-2100310311120212-0223233023110012-3003300331310330-1233021203312330-2302300022331031"></a>

## Next pages — gcp / 102122012220 / 4

- [gcp.byoc](resources--cloud_link--reference--group-001.md#canonical-3203032012032130-2203123221123112-2333003311102003-3330133201333003-3210212310211012-3233001103310111-0132322112122332-1332312020101333)
- [gcp.gcp_cred](resources--cloud_link--reference--group-001.md#canonical-3123231132230302-2000202222031322-1313322313312231-0011203110123103-3123132132132123-1100101001001230-1032310011123330-0300211020002302)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)

<a id="canonical-3203032012032130-2203123221123112-2333003311102003-3330133201333003-3210212310211012-3233001103310111-0132322112122332-1332312020101333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130130220023113-0103312302000223-2022012032130203-1221120322302322-3213132001333231-3032010113321113-1032321022013012-0033033330030013"></a>

## gcp.byoc — byoc / 010000303221 / 2

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311)
- [gcp](resources--cloud_link--reference--group-001.md#canonical-2120230010303101-2202120333222110-3132020103101203-0002303133323002-1131210100111130-2330332130032310-1203120020022120-2302330200333121)
- gcp.byoc

<a id="canonical-0203111220311011-2010230223002321-0230322212111111-3020313113332321-2320000231101221-2233330320232321-2311032112201013-3012233331212001"></a>

Type: `"object"`. single nested block, Optional.

GCP Bring Your Own Connections. List of GCP Bring You Own Connections.

Upstream description:

List of GCP Bring You Own Connections.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3012311310112002-2323321100032310-3231330023132003-0312231311013223-2002211311220030-0310310020302223-1100103013313232-1300030013202300"></a>

## Direct properties — byoc / 010000303221 / 3

- [connections](resources--cloud_link--reference--group-001.md#canonical-3300300311011020-1232112233202120-3121231302233032-0323330120012132-1203212002211002-0212331312203112-2320223102010211-0101312032120201): complete subsection reference.

<a id="canonical-0313003302001022-3011122032100020-3121203320202330-0032301113013112-2103323011231003-2102113102310023-3222120303011000-3003310101203312"></a>

## Next pages — byoc / 010000303221 / 4

- [gcp.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-3300300311011020-1232112233202120-3121231302233032-0323330120012132-1203212002211002-0212331312203112-2320223102010211-0101312032120201)
- [gcp](resources--cloud_link--reference--group-001.md#canonical-2120230010303101-2202120333222110-3132020103101203-0002303133323002-1131210100111130-2330332130032310-1203120020022120-2302330200333121)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)

<a id="canonical-3300300311011020-1232112233202120-3121231302233032-0323330120012132-1203212002211002-0212331312203112-2320223102010211-0101312032120201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0220113130002220-2202211322331111-2122113101122110-0100003201122123-1202110031030023-2230130320310200-1132131103301321-2300100201101123"></a>

## gcp.byoc.connections — connections / 120001332120 / 2

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311)
- [gcp](resources--cloud_link--reference--group-001.md#canonical-2120230010303101-2202120333222110-3132020103101203-0002303133323002-1131210100111130-2330332130032310-1203120020022120-2302330200333121)
- [gcp.byoc](resources--cloud_link--reference--group-001.md#canonical-3203032012032130-2203123221123112-2333003311102003-3330133201333003-3210212310211012-3233001103310111-0132322112122332-1332312020101333)
- gcp.byoc.connections

<a id="canonical-1331122223322010-1112223031030312-3130021221110223-3202111321213133-0233310101300023-2211301320013232-0102103002222113-1231021020322032"></a>

Type: `"object"`. list nested block, Optional.

Each 'Bring Your Own Connection' represents a virtual connection that the customer has provisioned
in the Cloud (.

Upstream description:

Each 'Bring Your Own Connection' represents a virtual connection that the customer has provisioned
in the Cloud (example: AWS Direct Connect). F5XC will orchestrate networking resources in the cloud
to facilitate seamless private connectivity.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2031213212000131-3010201301203213-3021330333230000-2302303220132111-1313021230133303-3113030312100102-3223323020200223-1033022011232220"></a>

## Direct properties — connections / 120001332120 / 3

<a id="canonical-1320323133001232-0232230210211023-3103103320133203-2211221013330123-1133210233100021-0020222223321131-2031130212310101-3123333230012110"></a>

<a id="canonical-0333133032132111-2311002130303222-2000113130230033-2100231110132113-2003321030120230-0022313311300310-0212303231013200-3103300332230131"></a>

## interconnect_attachment_name property — connections / 120001332120 / 4

Type: `"string"`. Optional.

Name of already-existing GCP Cloud Interconnect Attachment.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2033231312113303-0012033113330122-1001012002313000-3311301311232200-1223323200113230-2310122003100203-1310002121132231-1320223013323310"></a>

## project property — connections / 120001332120 / 5

Type: `"string"`. Optional.

Exclusive with \[same\_as\_credential\] Specify a GCP Project for the interconnect attachment.

Upstream description:

Exclusive with \[same\_as\_credential\] Specify a GCP Project for the interconnect attachment.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2101002111302230-0033230212211110-2010313010000121-0130021032210321-2033320033222103-3100222101311032-1113213320111313-0313313013100323"></a>

## region property — connections / 120001332120 / 6

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

Upstream description:

GCP Region in which the GCP Cloud Interconnect attachment is configured.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0131201302322303-3112323101113321-1320012013311003-0320233103213230-0323212122330222-3302030003101331-2333232121221032-1003211102021001"></a>

## Next pages — connections / 120001332120 / 7

- [gcp.byoc.connections.metadata](resources--cloud_link--reference--group-001.md#canonical-1231033111003012-3023121311330213-1113012103223011-3332101032200000-3321331130231323-0023222021213213-1202002212230131-3011212301103210)
- [gcp.byoc.connections.same_as_credential](resources--cloud_link--reference--group-001.md#canonical-0202331032230310-3101023230033102-3111320023333023-3220321311231102-3311231202011211-2122010332233212-1033013133032021-1121332032212113)
- [gcp.byoc](resources--cloud_link--reference--group-001.md#canonical-3203032012032130-2203123221123112-2333003311102003-3330133201333003-3210212310211012-3233001103310111-0132322112122332-1332312020101333)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)

<a id="canonical-1231033111003012-3023121311330213-1113012103223011-3332101032200000-3321331130231323-0023222021213213-1202002212230131-3011212301103210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1333113302102320-0012331032031302-1020021222311202-1012323311203201-3200300212300331-1103022312031112-0330212023033012-0233001231101101"></a>

## gcp.byoc.connections.metadata — metadata / 300132221020 / 2

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
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0303211001331030-1130333002301300-1222100011032333-0110001302033320-2130211020112220-2030231020011031-3312013100213120-3331212203102213"></a>

## Direct properties — metadata / 300132221020 / 3

<a id="canonical-3122232012030023-0321311322003302-2321132133212020-1210011023000032-1033302111323312-1113130122321311-1122232231012300-1213310012122213"></a>

<a id="canonical-3120202101120102-1230221022113120-2321112011230021-2232201120032223-2121001301010002-1213000020121212-3223013320322133-3312011131132231"></a>

## description_spec property — metadata / 300132221020 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-1003110021203231-2132122132010123-3023020130020120-2120102233230302-1010121301213133-2223211221303031-0021320201331303-2032312113302221"></a>

<a id="canonical-1131303020101132-0030313330200123-2223213331123220-0010133203202120-3031303021310130-1011011321101203-0310333201010211-1013033220332022"></a>

## name property — metadata / 300132221020 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0121101313232230-1331022211001022-0103012222113131-0223303101101213-2012121013201010-1300321200111210-1000113032321023-0203133211230213"></a>

## Next pages — metadata / 300132221020 / 6

- [gcp.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-3300300311011020-1232112233202120-3121231302233032-0323330120012132-1203212002211002-0212331312203112-2320223102010211-0101312032120201)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)

<a id="canonical-0202331032230310-3101023230033102-3111320023333023-3220321311231102-3311231202011211-2122010332233212-1033013133032021-1121332032212113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312232313333022-0131100211021300-2323232323003331-1220313310010021-3221130223033310-2303223301202001-2021130000132120-0001211002303202"></a>

## gcp.byoc.connections.same_as_credential — same_as_credential / 221313130020 / 2

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

Upstream description:

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

<a id="canonical-1102310132003300-1301202311233132-2030032233213032-3122011111231323-3203133020022301-1110311100022213-0321112332302331-2101220212201332"></a>

## Direct properties — same_as_credential / 221313130020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0031011321233021-0311221210031002-2213121312002003-2131103232031320-1213230021322003-3222031003103111-0000313312121110-2021123111220320"></a>

## Next pages — same_as_credential / 221313130020 / 4

- [gcp.byoc.connections](resources--cloud_link--reference--group-001.md#canonical-3300300311011020-1232112233202120-3121231302233032-0323330120012132-1203212002211002-0212331312203112-2320223102010211-0101312032120201)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)

<a id="canonical-3123231132230302-2000202222031322-1313322313312231-0011203110123103-3123132132132123-1100101001001230-1032310011123330-0300211020002302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2022202313223212-3101323211202200-2220133123111220-2103121000222211-2223023333310301-1120022232333132-0211200121013311-3311201203202131"></a>

## gcp.gcp_cred — gcp_cred / 123320332211 / 2

Breadcrumbs:

- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)
- [Property reference](resources--cloud_link--reference--group-001.md#canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311)
- [gcp](resources--cloud_link--reference--group-001.md#canonical-2120230010303101-2202120333222110-3132020103101203-0002303133323002-1131210100111130-2330332130032310-1203120020022120-2302330200333121)
- gcp.gcp_cred

<a id="canonical-0002230030330100-1210320203120013-2321121010301233-3211000030332011-1010013132201013-0211233102111221-2231000313110113-2330333302110212"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0203010301121013-2131202323301112-1303300223303212-0032021220233322-3323203211023001-3221000300130301-0233011303100203-0132211131131121"></a>

## Direct properties — gcp_cred / 123320332211 / 3

<a id="canonical-2300032232021230-0230313232031321-2112220013302112-0212102230332220-0112010111032030-2210203232210133-1321330320302212-0012231322322000"></a>

<a id="canonical-1010302010201122-0023021322031013-1133333110022122-2022303302011130-0012331310021001-3300031123003211-2332331120031121-0032110121222002"></a>

## name property — gcp_cred / 123320332211 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0323310132010022-2203033331122022-1113102332212221-2012132002211123-2123002012332121-2322323331331221-3200011321100122-1300033111202302"></a>

## namespace property — gcp_cred / 123320332211 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3212222300212323-3131322320010303-1200002330113203-1310321132330112-0233211023132321-3100000330232022-0033202010001011-1313332230111100"></a>

## tenant property — gcp_cred / 123320332211 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2003100000021010-3102003023132201-0023200332120211-0203102303120202-1012221311331330-1300022010203311-2320222123130002-2113211003213332"></a>

## Next pages — gcp_cred / 123320332211 / 7

- [gcp](resources--cloud_link--reference--group-001.md#canonical-2120230010303101-2202120333222110-3132020103101203-0002303133323002-1131210100111130-2330332130032310-1203120020022120-2302330200333121)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)

<a id="canonical-1000132120000000-2030223002311110-0323103213230233-0103003302023032-3221033120201230-3112000110012223-3112303220021202-1023123321332000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3223200122233300-2210122133333112-3210002033112311-0100103122112332-3222022000311221-3202000223031233-3030232102301231-3312213012133213"></a>

## timeouts — timeouts / 320301332313 / 2

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

<a id="canonical-0020020301233202-1333222010003220-1022300323331301-0322101101300223-0233321210031211-2300031123020201-1312010120102330-3203012210220111"></a>

## Direct properties — timeouts / 320301332313 / 3

<a id="canonical-1211000320332331-1311320121313023-1203033022003030-3130020012212220-2302100120231223-3120212203102103-0202131011200021-2323023212010210"></a>

<a id="canonical-3200000103233310-3321112002001110-3013122102032110-1221031221133003-1300312030003113-1320302023122333-2213130220231233-0032322000023211"></a>

## create property — timeouts / 320301332313 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1203330013323320-1102133132122222-2213031113323303-1031123321003323-0112011201302232-0233200020202130-2021021033020300-1201213212221031"></a>

<a id="canonical-2032112102122302-2213212320211122-0331201223013131-0121123330302131-2132022123030303-2030113222213322-2202313021321022-3223233133013302"></a>

## delete property — timeouts / 320301332313 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0101133123203102-2011101112221032-2032313310102011-1301020023332122-3332303110131023-1201010022011101-0210031223102301-0210230101000022"></a>

<a id="canonical-2031001011000330-2012031300220023-3000020021001103-1311210302210322-1233003202110220-2320222313230210-1002122000021331-3011033001001320"></a>

## read property — timeouts / 320301332313 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1201130222112201-0123320202131133-1303210100030010-3131102000122200-1313122000021220-3312303312223123-3021131200013221-2103322023020133"></a>

<a id="canonical-1232001210022030-3320212103230110-2231130020303113-3232323331120020-2300321331002233-2311021003220000-3232330000302232-3312320303100121"></a>

## update property — timeouts / 320301332313 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3211000003212310-1003232033020210-3111120011111033-1310320110002303-0222202311211103-0123211102122013-2023210303111312-2221211320133221"></a>

## Next pages — timeouts / 320301332313 / 8

- [Property reference](resources--cloud_link--reference--group-001.md#canonical-3321331113023223-1031311003111130-3033013110311022-2013130121323310-2022303312211132-1101113213222000-2113332022221332-0022211233220311)
- [xcsh_cloud_link](../resources/cloud_link.md#canonical-3202233132133131-2111003233330213-0332022100000001-3133321302110212-0022033031313130-0332303221332001-0302301320123002-0221013333221233)
