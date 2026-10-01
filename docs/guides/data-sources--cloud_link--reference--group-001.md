---
page_title: "xcsh_cloud_link reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_link reference."
---

# xcsh_cloud_link reference

<a id="canonical-1213002220021231-0321302131202132-0123330031233010-0020203103220000-1121330212132301-3103210130020013-1201230020112100-0121202133202012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1013322233213102-2110301320120113-2003122313230031-2012010311113113-2101122303202302-2320303000031122-1223312302100311-1301030122231231"></a>

## Property reference — Property reference / 011031131011 / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)
- Property reference

<a id="canonical-2223301310222131-0323200133033123-2211330101321202-1122023300031302-1200221313101123-1212331310003220-0213013123321133-2201300212210031"></a>

## Direct properties — Property reference / 011031131011 / 3

<a id="canonical-3300331211131232-0002022021230033-0232032232121302-2021330211332012-0102112312231210-3102112122202031-3020220021023131-2212100233321111"></a>

<a id="canonical-0223312012201101-3121013213123000-2231320223133110-1031130231121102-2111213330202013-0230213000223211-1331321113222111-0311213310132113"></a>

## annotations property — Property reference / 011031131011 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

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

- [aws](data-sources--cloud_link--reference--group-001.md#canonical-3332012130312211-3232211102111223-2232333321133131-0322312033131212-1013100303010313-0303323220310300-0010310203000000-3200110102030211): complete subsection reference.

<a id="canonical-3321221033013013-0131313003100223-3210000322103120-3130302322220230-3033121120113223-0210231312203320-0022322000102003-3203310022020111"></a>

<a id="canonical-1000200213023133-2010310133301031-0133133023300330-0201200101010113-0211013133133012-3311231312212213-2333311103021120-3231030100020100"></a>

## description property — Property reference / 011031131011 / 5

Type: `"string"`. Computed.

Description of the CloudLink.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [disabled](data-sources--cloud_link--reference--group-001.md#canonical-1121222321001003-0301130222132322-3133021003331210-1023121231212320-2033201232121230-1210032112221212-2122331210213200-2213211333131300): complete subsection reference.

- [enabled](data-sources--cloud_link--reference--group-001.md#canonical-2133211113302331-1111103313201201-2300313320232302-1001210232102002-3322223320212110-1202033223333021-1002113012122212-3121000132101120): complete subsection reference.

- [gcp](data-sources--cloud_link--reference--group-001.md#canonical-1132023021123233-0321010311212023-0313303231201123-3010111320000110-0312123230211210-2100233101120312-2330020333300003-0210301202220230): complete subsection reference.

<a id="canonical-1000332000300323-0301310331203210-2133112011311233-3212120112130321-1002010022001120-2012100102212200-0110302010221023-2302113223130033"></a>

<a id="canonical-3110202133000023-1221132203021212-3302120122230322-0011230000321022-0122220000332031-3110303123220221-2123021310012301-0220321301120112"></a>

## ID property — Property reference / 011031131011 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2103202012133211-0020002033331102-0010102120001133-0312232331333112-3022302132121030-2201221132230200-1212320322331303-1220012011332112"></a>

<a id="canonical-3010213130303012-2123313211021301-1033120200111200-1100100213122201-0112101120101232-2223230022013131-1312312130033003-3122001130130222"></a>

## labels property — Property reference / 011031131011 / 7

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

<a id="canonical-3033021032121013-1220110122123003-2230220002302022-0200220320120011-2103221012333210-2112230001221100-2103001120130100-0000231032221002"></a>

<a id="canonical-0100101333012122-0102101031132103-0110030322002112-2233031001322321-3112101033201113-1331122212113130-2330002221331211-3213323213330313"></a>

## name property — Property reference / 011031131011 / 8

Type: `"string"`. Required.

Name of the CloudLink.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2233330012321133-0003121312000321-3331031023110202-0330313103303231-0212011033320003-1011221031213210-0131330202003230-0213120113320220"></a>

<a id="canonical-1120221101110112-1101102230102110-3322322003113301-2110023012000110-1220230103200300-0333101221223123-0203022133300033-2121113231210202"></a>

## namespace property — Property reference / 011031131011 / 9

Type: `"string"`. Required.

Namespace where the CloudLink exists.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3333132031110222-2111000131001022-0200303022200003-0103302121011002-2333211312320301-3000300232223110-0122332222133210-1312132330323202"></a>

## All schema paths — Property reference / 011031131011 / 10

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--cloud_link--reference--group-001.md#canonical-3300331211131232-0002022021230033-0232032232121302-2021330211332012-0102112312231210-3102112122202031-3020220021023131-2212100233321111) |
| `aws` | [aws](data-sources--cloud_link--reference--group-001.md#canonical-3031303133033101-3311223010011111-0101013013133002-0030311302330310-1101110023200203-1003300302012003-1331111223002131-2310120111100002) |
| `aws.aws_cred` | [aws.aws_cred](data-sources--cloud_link--reference--group-001.md#canonical-1020233032022131-0231000332113311-0130203031002202-3010212131200230-2212231120220113-3033110332132122-1131210031232101-0133131030111120) |
| `aws.aws_cred.name` | [aws.aws_cred.name](data-sources--cloud_link--reference--group-001.md#canonical-1122213133000221-3131232203113201-0123212132231201-3220022122022300-1131022130213220-0012033003101222-1232000332132123-3010331200323133) |
| `aws.aws_cred.namespace` | [aws.aws_cred.namespace](data-sources--cloud_link--reference--group-001.md#canonical-1303001212020302-2223302313112223-3230221010332333-1020201103112303-3312302031330010-3012012021231110-0101201333001000-1210031200223000) |
| `aws.aws_cred.tenant` | [aws.aws_cred.tenant](data-sources--cloud_link--reference--group-001.md#canonical-1110223110013011-3212111021003103-1030300032231023-3330021233212221-2110330201020311-0122203301330302-0311111220132210-3023103231333132) |
| `aws.byoc` | [aws.byoc](data-sources--cloud_link--reference--group-001.md#canonical-3020012132000210-3313320300023300-1331020023022232-0202211310120001-3323313031102310-0130002112033102-3021003203303130-1030311301113021) |
| `aws.byoc.connections` | [aws.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-0231300131332300-3211011011233313-2313030323020013-1320032021313303-2232123031233302-1030200333103223-3221320212302230-1303310110321102) |
| `aws.byoc.connections.auth_key` | [aws.byoc.connections.auth_key](data-sources--cloud_link--reference--group-001.md#canonical-3103001103331010-3322031102213231-3132001012133213-1121210030100120-2032020312012333-0102032103023011-0011033200021203-3202031313100021) |
| `aws.byoc.connections.auth_key.blindfold_secret_info` | [aws.byoc.connections.auth_key.blindfold_secret_info](data-sources--cloud_link--reference--group-001.md#canonical-1003020011012211-1110323013031331-2020320313300130-2031330032133203-1332312013130031-3320012213232212-1312320123331303-1001223100222033) |
| `aws.byoc.connections.auth_key.blindfold_secret_info.decryption_provider` | [aws.byoc.connections.auth_key.blindfold_secret_info.decryption_provider](data-sources--cloud_link--reference--group-001.md#canonical-3313213201003312-1310200110231331-2230303103222022-1320230012100233-0221002213210331-2023333330230002-3120101111121113-2200203203320302) |
| `aws.byoc.connections.auth_key.blindfold_secret_info.location` | [aws.byoc.connections.auth_key.blindfold_secret_info.location](data-sources--cloud_link--reference--group-001.md#canonical-2030033022020303-3223222012222032-3031013323322302-0021130111220112-1032113102030033-3133011333222232-1130023002213230-1313313002023303) |
| `aws.byoc.connections.auth_key.blindfold_secret_info.store_provider` | [aws.byoc.connections.auth_key.blindfold_secret_info.store_provider](data-sources--cloud_link--reference--group-001.md#canonical-0011110031131131-0311032322300200-3130301303313100-1301012121212011-2032322321321210-3030023130002321-2103130212321121-1111211131130123) |
| `aws.byoc.connections.auth_key.clear_secret_info` | [aws.byoc.connections.auth_key.clear_secret_info](data-sources--cloud_link--reference--group-001.md#canonical-3200303121210030-0113021203021002-2023202031220231-2220320001211211-2113311020303123-3200021002113100-3000211103313032-3100232020330003) |
| `aws.byoc.connections.auth_key.clear_secret_info.provider_ref` | [aws.byoc.connections.auth_key.clear_secret_info.provider_ref](data-sources--cloud_link--reference--group-001.md#canonical-2113033331003203-2023202032311131-0123123232101133-2101000221123212-2030010122202332-2230232120333031-1311130332312313-0222000122220111) |
| `aws.byoc.connections.auth_key.clear_secret_info.url` | [aws.byoc.connections.auth_key.clear_secret_info.url](data-sources--cloud_link--reference--group-001.md#canonical-1100222210031212-1312303012111120-3121012101131303-1233002003322301-3130220020202233-2032132203332212-3230100231101100-1011101020231022) |
| `aws.byoc.connections.bgp_asn` | [aws.byoc.connections.bgp_asn](data-sources--cloud_link--reference--group-001.md#canonical-3133222302222102-0010202201002211-1333303310312003-3331111303012311-0233000023232001-1130330121223333-3130312223321130-3303322003222311) |
| `aws.byoc.connections.connection_id` | [aws.byoc.connections.connection_id](data-sources--cloud_link--reference--group-001.md#canonical-1202311033110303-0333301333130330-2133213002112333-2030002323300322-1000101333200203-3221123221011311-0011000301211301-1311323000303330) |
| `aws.byoc.connections.ipv4` | [aws.byoc.connections.ipv4](data-sources--cloud_link--reference--group-001.md#canonical-3330001321022100-2123331012113313-2020113223231022-1312332201131122-0031312101221211-0213021223122032-2203323220032103-1013010201111010) |
| `aws.byoc.connections.ipv4.aws_router_peer_address` | [aws.byoc.connections.ipv4.aws_router_peer_address](data-sources--cloud_link--reference--group-001.md#canonical-0101031100023123-1313021103323101-3300220013220211-2333022032132100-3220103120103320-1000322233121300-3301123130212203-1220012210313032) |
| `aws.byoc.connections.ipv4.router_peer_address` | [aws.byoc.connections.ipv4.router_peer_address](data-sources--cloud_link--reference--group-001.md#canonical-2001020331131130-1030232211103231-0021222032221231-3211312023313012-0320220200333113-3212110103010120-2301233320322230-3021331233220131) |
| `aws.byoc.connections.metadata` | [aws.byoc.connections.metadata](data-sources--cloud_link--reference--group-001.md#canonical-1013112002012100-3110213113020303-2323200233203302-1311103222122320-3212303030132023-0000230202031300-2111111001321313-0333221213333333) |
| `aws.byoc.connections.metadata.description_spec` | [aws.byoc.connections.metadata.description_spec](data-sources--cloud_link--reference--group-001.md#canonical-1323321302020332-2110020300311002-1032320113320002-1213121120310323-1220321013200220-1203200030322233-1113222310312022-3103330213000030) |
| `aws.byoc.connections.metadata.name` | [aws.byoc.connections.metadata.name](data-sources--cloud_link--reference--group-001.md#canonical-0301003120300003-0001010030202223-0211013212123100-1222032333333301-2003331332111000-1110013010312310-3310333222232302-0321020000012101) |
| `aws.byoc.connections.region` | [aws.byoc.connections.region](data-sources--cloud_link--reference--group-001.md#canonical-3223013100332133-0121013102332312-0221222331210213-1331003003013332-0222031031112130-3300031120110211-1031203310210133-0001222002101011) |
| `aws.byoc.connections.system_generated_name` | [aws.byoc.connections.system_generated_name](data-sources--cloud_link--reference--group-001.md#canonical-1123131121020211-0011101130001121-3323323333013332-3210322012131133-1022201311221323-3010303132000331-3003303321032203-1210020123323322) |
| `aws.byoc.connections.tags` | [aws.byoc.connections.tags](data-sources--cloud_link--reference--group-001.md#canonical-2010323113101333-3302131130231103-3303300212213122-0331110203132103-0031332121321213-1033020313212301-1323133103331203-3210200002202331) |
| `aws.byoc.connections.user_assigned_name` | [aws.byoc.connections.user_assigned_name](data-sources--cloud_link--reference--group-001.md#canonical-2132103133130221-3102302011120120-3001103223001130-1010210330130003-2103011211013330-2111000023213333-2210131020001000-1201322011002031) |
| `aws.byoc.connections.virtual_interface_type` | [aws.byoc.connections.virtual_interface_type](data-sources--cloud_link--reference--group-001.md#canonical-2032033012313031-1013000320322201-3322033311003330-3210112122213021-3001331322121201-3201313113021212-1211313032322331-3022133211212001) |
| `aws.byoc.connections.vlan` | [aws.byoc.connections.vlan](data-sources--cloud_link--reference--group-001.md#canonical-1231102231223223-2310323203323332-0221232131100110-3210323210331311-1223122012231331-2011302333312321-2333311101221113-1232113012220003) |
| `aws.custom_asn` | [aws.custom_asn](data-sources--cloud_link--reference--group-001.md#canonical-2123300013221100-0203331202310130-3203302030220300-3202212213300323-1101032122032110-2212200201113330-3313200022233220-3031002132221132) |
| `description` | [description](data-sources--cloud_link--reference--group-001.md#canonical-3321221033013013-0131313003100223-3210000322103120-3130302322220230-3033121120113223-0210231312203320-0022322000102003-3203310022020111) |
| `disabled` | [disabled](data-sources--cloud_link--reference--group-001.md#canonical-0023213231102121-1110303231330113-2132223121123132-3223022331310132-2202121011001302-3230100010233023-3132220320232013-2331032032213330) |
| `enabled` | [enabled](data-sources--cloud_link--reference--group-001.md#canonical-2322221213320111-1300302120212320-3010203021000130-1203112101301001-0333123122220310-1023311000210321-0122123133010200-1001131223300232) |
| `enabled.cloudlink_network_name` | [enabled.cloudlink_network_name](data-sources--cloud_link--reference--group-001.md#canonical-1120310010202113-3200033033010032-0200113333220231-1020000210102020-0123000233231233-2013232220101100-3333303312203302-3101330331301001) |
| `gcp` | [gcp](data-sources--cloud_link--reference--group-001.md#canonical-3030201223313101-3333102310033121-0303300311223113-2330011201130221-3310300212322123-1133132330001231-2313323111313003-3023000132120223) |
| `gcp.byoc` | [gcp.byoc](data-sources--cloud_link--reference--group-001.md#canonical-1200222333111311-3210200113212201-3320112232321003-1320023122121011-1032030312330222-0101012303003321-2130322310003321-3031200323300031) |
| `gcp.byoc.connections` | [gcp.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-0032210302131311-2210232330312230-2213220002030103-2301022211132102-0122112230113203-2100333023122020-2201301312213123-3322133011011013) |
| `gcp.byoc.connections.interconnect_attachment_name` | [gcp.byoc.connections.interconnect_attachment_name](data-sources--cloud_link--reference--group-001.md#canonical-2323212021122031-3330323120320220-0201323012223101-1110223230232300-0320210203102032-3312221333021132-1033103133222200-1320212130033020) |
| `gcp.byoc.connections.metadata` | [gcp.byoc.connections.metadata](data-sources--cloud_link--reference--group-001.md#canonical-2223011221323302-3101230110031131-3020230122301023-1230010213023332-1332010310010323-2202010301011100-0333112013333111-0330001223122303) |
| `gcp.byoc.connections.metadata.description_spec` | [gcp.byoc.connections.metadata.description_spec](data-sources--cloud_link--reference--group-001.md#canonical-1001023331231332-3012031213123233-3010233232113202-0023301131202201-2022313322021302-1131202100122201-1130312230021120-1311030332300332) |
| `gcp.byoc.connections.metadata.name` | [gcp.byoc.connections.metadata.name](data-sources--cloud_link--reference--group-001.md#canonical-2032123001020111-0210010111110132-2300011101111112-1321202320102011-1132222122022030-2002133303211220-1120210200032100-2230212322323323) |
| `gcp.byoc.connections.project` | [gcp.byoc.connections.project](data-sources--cloud_link--reference--group-001.md#canonical-1333002302032010-1002003320003222-1103133310003001-1223022020010222-0013100313012131-1030023202021231-2113321201030033-1232131121111221) |
| `gcp.byoc.connections.region` | [gcp.byoc.connections.region](data-sources--cloud_link--reference--group-001.md#canonical-3313113100112223-2303203101223222-0313322202221103-2300322230221322-2300230202313122-3030020202011113-0021303330303232-3212130303030102) |
| `gcp.byoc.connections.same_as_credential` | [gcp.byoc.connections.same_as_credential](data-sources--cloud_link--reference--group-001.md#canonical-0123303130232102-2123010330232121-3301131110101233-1032000021122103-2003320003213111-0310113200203233-3110023202230210-0320231111022033) |
| `gcp.gcp_cred` | [gcp.gcp_cred](data-sources--cloud_link--reference--group-001.md#canonical-1120210203310010-3231322002013110-0002323131031222-3330022301023031-0013332201330201-2332032201333012-2111020311102321-3210000230110123) |
| `gcp.gcp_cred.name` | [gcp.gcp_cred.name](data-sources--cloud_link--reference--group-001.md#canonical-3223020220320031-0321122330301212-3201321101123133-1011320021122032-3310223301032120-2032311222322332-3303000222011203-0030333211210122) |
| `gcp.gcp_cred.namespace` | [gcp.gcp_cred.namespace](data-sources--cloud_link--reference--group-001.md#canonical-3323032302021223-3311331011133010-0313100320130313-1322313030223213-2231100302213322-3202312333203212-3231331221201301-0103121030111302) |
| `gcp.gcp_cred.tenant` | [gcp.gcp_cred.tenant](data-sources--cloud_link--reference--group-001.md#canonical-1102020100103001-1212032301331132-1220130312020313-0111032130001313-1020033033332330-2202022121322223-3022120302331323-0210210232311313) |
| `id` | [ID](data-sources--cloud_link--reference--group-001.md#canonical-1000332000300323-0301310331203210-2133112011311233-3212120112130321-1002010022001120-2012100102212200-0110302010221023-2302113223130033) |
| `labels` | [labels](data-sources--cloud_link--reference--group-001.md#canonical-2103202012133211-0020002033331102-0010102120001133-0312232331333112-3022302132121030-2201221132230200-1212320322331303-1220012011332112) |
| `name` | [name](data-sources--cloud_link--reference--group-001.md#canonical-3033021032121013-1220110122123003-2230220002302022-0200220320120011-2103221012333210-2112230001221100-2103001120130100-0000231032221002) |
| `namespace` | [namespace](data-sources--cloud_link--reference--group-001.md#canonical-2233330012321133-0003121312000321-3331031023110202-0330313103303231-0212011033320003-1011221031213210-0131330202003230-0213120113320220) |

<a id="canonical-0213200000031332-3310102010102011-2031321203011211-3102213123001221-1210033210002121-3303233212313130-0013233301022323-0113130313203311"></a>

## Next pages — Property reference / 011031131011 / 11

- [aws](data-sources--cloud_link--reference--group-001.md#canonical-3332012130312211-3232211102111223-2232333321133131-0322312033131212-1013100303010313-0303323220310300-0010310203000000-3200110102030211)
- [disabled](data-sources--cloud_link--reference--group-001.md#canonical-1121222321001003-0301130222132322-3133021003331210-1023121231212320-2033201232121230-1210032112221212-2122331210213200-2213211333131300)
- [enabled](data-sources--cloud_link--reference--group-001.md#canonical-2133211113302331-1111103313201201-2300313320232302-1001210232102002-3322223320212110-1202033223333021-1002113012122212-3121000132101120)
- [gcp](data-sources--cloud_link--reference--group-001.md#canonical-1132023021123233-0321010311212023-0313303231201123-3010111320000110-0312123230211210-2100233101120312-2330020333300003-0210301202220230)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)

<a id="canonical-3332012130312211-3232211102111223-2232333321133131-0322312033131212-1013100303010313-0303323220310300-0010310203000000-3200110102030211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1332320120103120-0111130232020312-2220021103020000-0222030313310000-1331202111302033-3123101013013101-0102333100121323-2302111022321121"></a>

## aws — aws / 122202233002 / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-1213002220021231-0321302131202132-0123330031233010-0020203103220000-1121330212132301-3103210130020013-1201230020112100-0121202133202012)
- aws

<a id="canonical-3031303133033101-3311223010011111-0101013013133002-0030311302330310-1101110023200203-1003300302012003-1331111223002131-2310120111100002"></a>

Type: `"single"`. Computed.

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

- [aws](data-sources--cloud_link--reference--group-001.md#canonical-3031303133033101-3311223010011111-0101013013133002-0030311302330310-1101110023200203-1003300302012003-1331111223002131-2310120111100002)
- [gcp](data-sources--cloud_link--reference--group-001.md#canonical-3030201223313101-3333102310033121-0303300311223113-2330011201130221-3310300212322123-1133132330001231-2313323111313003-3023000132120223)

Select alternatives according to the provider validators above.

<a id="canonical-1232301202012122-1320310213031123-0033032032100323-3022231203123022-1301011232011101-3223210110123013-0101330320010000-0203023013002310"></a>

## Direct properties — aws / 122202233002 / 3

- [aws_cred](data-sources--cloud_link--reference--group-001.md#canonical-0222002203323203-1300301323132032-3022013300121312-3110212233213333-2203302103211303-0103331130022100-2332033100121301-3133310221121021): complete subsection reference.

- [byoc](data-sources--cloud_link--reference--group-001.md#canonical-1033213001233100-3203112202003030-0310131022131311-2330001032132113-0301131303010200-1301121103323332-0232201231230233-2133300030122323): complete subsection reference.

<a id="canonical-2123300013221100-0203331202310130-3203302030220300-3202212213300323-1101032122032110-2212200201113330-3313200022233220-3031002132221132"></a>

<a id="canonical-3013303032300202-1220321233023102-3023231102130311-2012302021310102-1330201230011313-3210323322333220-1002022031133123-3030232302203302"></a>

## custom_asn property — aws / 122202233002 / 4

Type: `"number"`. Computed.

Exclusive with \[\] F5XC will use custom ASN to create a Direct Connect Gateway.

Upstream description:

Exclusive with \[\] F5XC will use custom ASN to create a Direct Connect Gateway.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1131212032323130-1111330232033122-1103010222233232-3222033232232231-3302332020213301-0203333031001100-3023000231132003-1110121333012122"></a>

## Next pages — aws / 122202233002 / 5

- [aws.aws_cred](data-sources--cloud_link--reference--group-001.md#canonical-0222002203323203-1300301323132032-3022013300121312-3110212233213333-2203302103211303-0103331130022100-2332033100121301-3133310221121021)
- [aws.byoc](data-sources--cloud_link--reference--group-001.md#canonical-1033213001233100-3203112202003030-0310131022131311-2330001032132113-0301131303010200-1301121103323332-0232201231230233-2133300030122323)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-1213002220021231-0321302131202132-0123330031233010-0020203103220000-1121330212132301-3103210130020013-1201230020112100-0121202133202012)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)

<a id="canonical-0222002203323203-1300301323132032-3022013300121312-3110212233213333-2203302103211303-0103331130022100-2332033100121301-3133310221121021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203232001201231-0122222030010010-2301312221010031-3111223112122211-1302331220231331-1033313113111331-1333030101300323-2203320012011320"></a>

## aws.aws_cred — aws_cred / 300311122110 / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-1213002220021231-0321302131202132-0123330031233010-0020203103220000-1121330212132301-3103210130020013-1201230020112100-0121202133202012)
- [aws](data-sources--cloud_link--reference--group-001.md#canonical-3332012130312211-3232211102111223-2232333321133131-0322312033131212-1013100303010313-0303323220310300-0010310203000000-3200110102030211)
- aws.aws_cred

<a id="canonical-1020233032022131-0231000332113311-0130203031002202-3010212131200230-2212231120220113-3033110332132122-1131210031232101-0133131030111120"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-0010233030211102-2310220200123220-0303231023103331-3321010021212032-3300011113331213-3202033020002322-0320021311122020-2111212302130110"></a>

## Direct properties — aws_cred / 300311122110 / 3

<a id="canonical-1122213133000221-3131232203113201-0123212132231201-3220022122022300-1131022130213220-0012033003101222-1232000332132123-3010331200323133"></a>

<a id="canonical-1133122032202030-0023103131112000-1113303122221312-0000330132123330-0330221311332311-1302112232220012-0111100320221011-1332320100123030"></a>

## name property — aws_cred / 300311122110 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1303001212020302-2223302313112223-3230221010332333-1020201103112303-3312302031330010-3012012021231110-0101201333001000-1210031200223000"></a>

<a id="canonical-2100231322112302-2111213000032203-0231032233001233-1222100302302011-3132100023122230-1131333001103110-1201203121113211-0030303230132013"></a>

## namespace property — aws_cred / 300311122110 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1110223110013011-3212111021003103-1030300032231023-3330021233212221-2110330201020311-0122203301330302-0311111220132210-3023103231333132"></a>

<a id="canonical-0200313110133001-0332310222203323-0023310211011221-2231223022113210-3010223222032313-1000312020232301-0320030233311312-3321003021312103"></a>

## tenant property — aws_cred / 300311122110 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3320133002303010-0003311313123112-1131111311200310-2201320201313013-1321201232123123-3233333210323211-0220201001310031-2023333222221011"></a>

## Next pages — aws_cred / 300311122110 / 7

- [aws](data-sources--cloud_link--reference--group-001.md#canonical-3332012130312211-3232211102111223-2232333321133131-0322312033131212-1013100303010313-0303323220310300-0010310203000000-3200110102030211)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)

<a id="canonical-1033213001233100-3203112202003030-0310131022131311-2330001032132113-0301131303010200-1301121103323332-0232201231230233-2133300030122323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1020211111220220-3020322221030322-2321111122020213-3311123001311012-0231332031022010-2223032231300100-2312133333310202-2130200313012022"></a>

## aws.byoc — byoc / 133001300101 / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-1213002220021231-0321302131202132-0123330031233010-0020203103220000-1121330212132301-3103210130020013-1201230020112100-0121202133202012)
- [aws](data-sources--cloud_link--reference--group-001.md#canonical-3332012130312211-3232211102111223-2232333321133131-0322312033131212-1013100303010313-0303323220310300-0010310203000000-3200110102030211)
- aws.byoc

<a id="canonical-3020012132000210-3313320300023300-1331020023022232-0202211310120001-3323313031102310-0130002112033102-3021003203303130-1030311301113021"></a>

Type: `"single"`. Computed.

Bring Your Own Connections. List of Bring You Own Connection.

Upstream description:

List of Bring You Own Connection.

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

<a id="canonical-1201221221112233-3102221220120222-0123330103003202-3201130232022200-1103223122022011-2100031321001021-1332022112021232-1302121201310132"></a>

## Direct properties — byoc / 133001300101 / 3

- [connections](data-sources--cloud_link--reference--group-001.md#canonical-1032332332221332-0102020012232333-0310303130112211-3011110333000321-3020013110000301-1101131123102203-3123131221012002-1012333233132121): complete subsection reference.

<a id="canonical-3330321111120232-3213320331123003-1011123230032201-1221220230311303-3033030101312121-1223231022312100-3102323103021203-3203211212201023"></a>

## Next pages — byoc / 133001300101 / 4

- [aws.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-1032332332221332-0102020012232333-0310303130112211-3011110333000321-3020013110000301-1101131123102203-3123131221012002-1012333233132121)
- [aws](data-sources--cloud_link--reference--group-001.md#canonical-3332012130312211-3232211102111223-2232333321133131-0322312033131212-1013100303010313-0303323220310300-0010310203000000-3200110102030211)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)

<a id="canonical-1032332332221332-0102020012232333-0310303130112211-3011110333000321-3020013110000301-1101131123102203-3123131221012002-1012333233132121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301010130012130-3131200331001133-1001123032123203-2130012032212233-0122122211212023-3302121133233011-1002313320311321-3230103103002000"></a>

## aws.byoc.connections — connections / 202013230210 / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-1213002220021231-0321302131202132-0123330031233010-0020203103220000-1121330212132301-3103210130020013-1201230020112100-0121202133202012)
- [aws](data-sources--cloud_link--reference--group-001.md#canonical-3332012130312211-3232211102111223-2232333321133131-0322312033131212-1013100303010313-0303323220310300-0010310203000000-3200110102030211)
- [aws.byoc](data-sources--cloud_link--reference--group-001.md#canonical-1033213001233100-3203112202003030-0310131022131311-2330001032132113-0301131303010200-1301121103323332-0232201231230233-2133300030122323)
- aws.byoc.connections

<a id="canonical-0231300131332300-3211011011233313-2313030323020013-1320032021313303-2232123031233302-1030200333103223-3221320212302230-1303310110321102"></a>

Type: `"list"`. Computed.

List of Bring You Own Connections. These AWS Direct Connect connections are not managed by F5XC but
will be used for connecting sites and REs.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2313011220222223-2213033100211311-2000223120333232-3300213313000011-1300023112112312-0020232210302201-0301021212310123-2020132313200123"></a>

## Direct properties — connections / 202013230210 / 3

- [auth_key](data-sources--cloud_link--reference--group-001.md#canonical-3013332200201212-2221123020103131-0033013301231101-2031233110103331-1322031011032202-0131121303312003-3332112101310312-1022031211032101): complete subsection reference.

<a id="canonical-3133222302222102-0010202201002211-1333303310312003-3331111303012311-0233000023232001-1130330121223333-3130312223321130-3303322003222311"></a>

<a id="canonical-1113110322000011-0220122031013110-3100023303013232-1020003201301133-1330321030031322-1201010113222021-1210321220231102-2131331233331000"></a>

## bgp_asn property — connections / 202013230210 / 4

Type: `"number"`. Computed.

The Border Gateway Protocol (BGP) Autonomous System Number (ASN) of your on-premises router for the
new virtual interface to be configured on AWS.

Upstream description:

The Border Gateway Protocol (BGP) Autonomous System Number (ASN) of your on-premises router for the
new virtual interface to be configured on AWS.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1202311033110303-0333301333130330-2133213002112333-2030002323300322-1000101333200203-3221123221011311-0011000301211301-1311323000303330"></a>

<a id="canonical-3200223010033202-3331303232222111-0102100201210221-2000333112213313-0021312212303103-0203132110222302-1012211122231333-0033102313303122"></a>

## connection_id property — connections / 202013230210 / 5

Type: `"string"`. Computed.

ID of the existing AWS Direct Connect Connection.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [ipv4](data-sources--cloud_link--reference--group-001.md#canonical-1223231323121200-0331331200303102-3033122032332332-2330012220000111-2313000001020112-1202033210133001-3132303113020302-1300321221311220): complete subsection reference.

- [metadata](data-sources--cloud_link--reference--group-001.md#canonical-1223130301101102-1303122312123122-1122011333113020-3223300122331102-2233110331213020-2000220202323332-3112023302000123-3022010332333320): complete subsection reference.

<a id="canonical-3223013100332133-0121013102332312-0221222331210213-1331003003013332-0222031031112130-3300031120110211-1031203310210133-0001222002101011"></a>

<a id="canonical-3002120233212301-0310300102310331-0211012202021213-1133031021213213-2130031233120033-3310323223232131-1110112220102321-1323211122120112"></a>

## region property — connections / 202013230210 / 6

Type: `"string"`. Computed.

\[Enum:
ap-northeast-1|ap-southeast-1|eu-central-1|eu-west-1|eu-west-3|sa-east-1|us-east-1|us-east-2|us-west-2|ca-central-1|af-south-1|ap-east-1|ap-south-1|ap-northeast-2|ap-southeast-2|eu-south-1|eu-north-1|eu-west-2|me-south-1|us-west-1|ap-southeast-3\]
Region. Region where the connection is setup. Possible values are \`ap-northeast-1\`,
\`ap-southeast-1\`, \`eu-central-1\`, \`eu-west-1\`, \`eu-west-3\`, \`sa-east-1\`, \`us-east-1\`,
\`us-east-2\`, \`us-west-2\`, \`ca-central-1\`, \`af-south-1\`, \`ap-east-1\`, \`ap-south-1\`,
\`ap-northeast-2\`, \`ap-southeast-2\`, \`eu-south-1\`, \`eu-north-1\`, \`eu-west-2\`,
\`me-south-1\`, \`us-west-1\`, \`ap-southeast-3\`.

Upstream description:

Region where the connection is setup.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [system_generated_name](data-sources--cloud_link--reference--group-001.md#canonical-2212100312030011-2222203211011323-2122230012032120-2113321120213000-2013301113031103-2200210132321211-1132201103012200-0010010121302130): complete subsection reference.

<a id="canonical-2010323113101333-3302131130231103-3303300212213122-0331110203132103-0031332121321213-1033020313212301-1323133103331203-3210200002202331"></a>

<a id="canonical-0031303001210222-0000301202103313-3121211231021033-0101121110320020-1202033212113210-3131030001223033-2123310122111220-1203332303111221"></a>

## tags property — connections / 202013230210 / 7

Type: `["map", "string"]`. Computed.

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

<a id="canonical-2132103133130221-3102302011120120-3001103223001130-1010210330130003-2103011211013330-2111000023213333-2210131020001000-1201322011002031"></a>

<a id="canonical-3322001032101301-0113000201003022-0220221322030031-3222103232133122-1313102132233102-0200231313010212-2122002112212023-0322300113332232"></a>

## user_assigned_name property — connections / 202013230210 / 8

Type: `"string"`. Computed.

Exclusive with \[system\_generated\_name\] User is managing the AWS resource name.

Upstream description:

Exclusive with \[system\_generated\_name\] User is managing the AWS resource name.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2032033012313031-1013000320322201-3322033311003330-3210112122213021-3001331322121201-3201313113021212-1211313032322331-3022133211212001"></a>

<a id="canonical-0012202021212100-1101123222210310-1231100221301023-1002220322232112-1210003030001102-2311120230122123-0010221302120200-0123002310122031"></a>

## virtual_interface_type property — connections / 202013230210 / 9

Type: `"string"`. Computed.

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

<a id="canonical-1231102231223223-2310323203323332-0221232131100110-3210323210331311-1223122012231331-2011302333312321-2333311101221113-1232113012220003"></a>

<a id="canonical-3031332212002320-1332313300313301-1101320222233312-1110311002030122-0113133111201322-0132103123220213-0223011013000233-2022303130313011"></a>

## vlan property — connections / 202013230210 / 10

Type: `"number"`. Computed.

Virtual Local Area Network number for the new virtual interface to be configured on the AWS. This
tag is required for any traffic traversing the AWS Direct Connect connection.

Upstream description:

Virtual Local Area Network number for the new virtual interface to be configured on the AWS. This
tag is required for any traffic traversing the AWS Direct Connect connection.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2322130020330203-3330100132303101-1003131332000033-2102133331023123-2333023101121301-3301121223032100-3210300100300233-0121002032230233"></a>

## Next pages — connections / 202013230210 / 11

- [aws.byoc.connections.auth_key](data-sources--cloud_link--reference--group-001.md#canonical-3013332200201212-2221123020103131-0033013301231101-2031233110103331-1322031011032202-0131121303312003-3332112101310312-1022031211032101)
- [aws.byoc.connections.ipv4](data-sources--cloud_link--reference--group-001.md#canonical-1223231323121200-0331331200303102-3033122032332332-2330012220000111-2313000001020112-1202033210133001-3132303113020302-1300321221311220)
- [aws.byoc.connections.metadata](data-sources--cloud_link--reference--group-001.md#canonical-1223130301101102-1303122312123122-1122011333113020-3223300122331102-2233110331213020-2000220202323332-3112023302000123-3022010332333320)
- [aws.byoc.connections.system_generated_name](data-sources--cloud_link--reference--group-001.md#canonical-2212100312030011-2222203211011323-2122230012032120-2113321120213000-2013301113031103-2200210132321211-1132201103012200-0010010121302130)
- [aws.byoc](data-sources--cloud_link--reference--group-001.md#canonical-1033213001233100-3203112202003030-0310131022131311-2330001032132113-0301131303010200-1301121103323332-0232201231230233-2133300030122323)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)

<a id="canonical-3013332200201212-2221123020103131-0033013301231101-2031233110103331-1322031011032202-0131121303312003-3332112101310312-1022031211032101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123102033012213-2122213302102303-1132321300102123-2011302021233222-2113200303133321-2113003230232223-3211330230003322-0200330323221332"></a>

## aws.byoc.connections.auth_key — auth_key / 030211223221 / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-1213002220021231-0321302131202132-0123330031233010-0020203103220000-1121330212132301-3103210130020013-1201230020112100-0121202133202012)
- [aws](data-sources--cloud_link--reference--group-001.md#canonical-3332012130312211-3232211102111223-2232333321133131-0322312033131212-1013100303010313-0303323220310300-0010310203000000-3200110102030211)
- [aws.byoc](data-sources--cloud_link--reference--group-001.md#canonical-1033213001233100-3203112202003030-0310131022131311-2330001032132113-0301131303010200-1301121103323332-0232201231230233-2133300030122323)
- [aws.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-1032332332221332-0102020012232333-0310303130112211-3011110333000321-3020013110000301-1101131123102203-3123131221012002-1012333233132121)
- aws.byoc.connections.auth_key

<a id="canonical-3103001103331010-3322031102213231-3132001012133213-1121210030100120-2032020312012333-0102032103023011-0011033200021203-3202031313100021"></a>

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

<a id="canonical-3110223222022100-1213311332222300-2002222122022323-0132313312320123-1220323111132123-1323303012030303-1320202130301202-1301123220333211"></a>

## Direct properties — auth_key / 030211223221 / 3

- [blindfold_secret_info](data-sources--cloud_link--reference--group-001.md#canonical-0020032102322010-2220212013303122-2302311232111113-3112020221233121-3001232001301131-1102233110122211-1212003233130032-1223303011303031): complete subsection reference.

- [clear_secret_info](data-sources--cloud_link--reference--group-001.md#canonical-1102312021033213-1231133330231313-2101233232212102-2032233231133123-2031221033132033-0202303102102011-2101223111332030-1203021210212301): complete subsection reference.

<a id="canonical-3323302133022023-0110131132100313-0103121200012133-3113312323200032-2021101102130101-2223001310212001-3012100221003122-2333203132222130"></a>

## Next pages — auth_key / 030211223221 / 4

- [aws.byoc.connections.auth_key.blindfold_secret_info](data-sources--cloud_link--reference--group-001.md#canonical-0020032102322010-2220212013303122-2302311232111113-3112020221233121-3001232001301131-1102233110122211-1212003233130032-1223303011303031)
- [aws.byoc.connections.auth_key.clear_secret_info](data-sources--cloud_link--reference--group-001.md#canonical-1102312021033213-1231133330231313-2101233232212102-2032233231133123-2031221033132033-0202303102102011-2101223111332030-1203021210212301)
- [aws.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-1032332332221332-0102020012232333-0310303130112211-3011110333000321-3020013110000301-1101131123102203-3123131221012002-1012333233132121)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)

<a id="canonical-0020032102322010-2220212013303122-2302311232111113-3112020221233121-3001232001301131-1102233110122211-1212003233130032-1223303011303031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0101123233120032-2012003301013333-0202020023102112-0232120232013310-0123002211111123-0111113031321302-0211103001013321-0001013223322030"></a>

## aws.byoc.connections.auth_key.blindfold_secret_info — blindfold_secret_info / 302232211021 / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-1213002220021231-0321302131202132-0123330031233010-0020203103220000-1121330212132301-3103210130020013-1201230020112100-0121202133202012)
- [aws](data-sources--cloud_link--reference--group-001.md#canonical-3332012130312211-3232211102111223-2232333321133131-0322312033131212-1013100303010313-0303323220310300-0010310203000000-3200110102030211)
- [aws.byoc](data-sources--cloud_link--reference--group-001.md#canonical-1033213001233100-3203112202003030-0310131022131311-2330001032132113-0301131303010200-1301121103323332-0232201231230233-2133300030122323)
- [aws.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-1032332332221332-0102020012232333-0310303130112211-3011110333000321-3020013110000301-1101131123102203-3123131221012002-1012333233132121)
- [aws.byoc.connections.auth_key](data-sources--cloud_link--reference--group-001.md#canonical-3013332200201212-2221123020103131-0033013301231101-2031233110103331-1322031011032202-0131121303312003-3332112101310312-1022031211032101)
- aws.byoc.connections.auth_key.blindfold_secret_info

<a id="canonical-1003020011012211-1110323013031331-2020320313300130-2031330032133203-1332312013130031-3320012213232212-1312320123331303-1001223100222033"></a>

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

<a id="canonical-2322123321031113-0210003102111103-0331333113132303-2130232111211201-1232321313202121-2231302311120022-3333221110331021-3220121301203321"></a>

## Direct properties — blindfold_secret_info / 302232211021 / 3

<a id="canonical-3313213201003312-1310200110231331-2230303103222022-1320230012100233-0221002213210331-2023333330230002-3120101111121113-2200203203320302"></a>

<a id="canonical-1331101330333311-3310333332200022-3013030331132203-2020103331201303-0333131002033331-2310313212032211-1202211303031331-3133300313123110"></a>

## decryption_provider property — blindfold_secret_info / 302232211021 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2030033022020303-3223222012222032-3031013323322302-0021130111220112-1032113102030033-3133011333222232-1130023002213230-1313313002023303"></a>

<a id="canonical-2212023222131211-2011121202223132-3100101221122121-0022111003131220-2132103232301220-1021011110302100-3131132102300011-0112123203001300"></a>

## location property — blindfold_secret_info / 302232211021 / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0011110031131131-0311032322300200-3130301303313100-1301012121212011-2032322321321210-3030023130002321-2103130212321121-1111211131130123"></a>

<a id="canonical-2011323022201020-0233032033222110-0201030020120312-0020310121002311-3100322122002033-2023111330212133-0213202333211233-2213331232013112"></a>

## store_provider property — blindfold_secret_info / 302232211021 / 6

Type: `"string"`. Computed.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3002001331201321-0112010200020233-1221213121100020-0223131003032000-0013013202322120-1232320230312130-0233221301121301-0223101300123132"></a>

## Next pages — blindfold_secret_info / 302232211021 / 7

- [aws.byoc.connections.auth_key](data-sources--cloud_link--reference--group-001.md#canonical-3013332200201212-2221123020103131-0033013301231101-2031233110103331-1322031011032202-0131121303312003-3332112101310312-1022031211032101)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)

<a id="canonical-1102312021033213-1231133330231313-2101233232212102-2032233231133123-2031221033132033-0202303102102011-2101223111332030-1203021210212301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203213210013130-0320332203231112-2320011120333113-1202031201232003-2013111303123103-1013303000101023-0232313103301202-3003111130332212"></a>

## aws.byoc.connections.auth_key.clear_secret_info — clear_secret_info / 301132102301 / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-1213002220021231-0321302131202132-0123330031233010-0020203103220000-1121330212132301-3103210130020013-1201230020112100-0121202133202012)
- [aws](data-sources--cloud_link--reference--group-001.md#canonical-3332012130312211-3232211102111223-2232333321133131-0322312033131212-1013100303010313-0303323220310300-0010310203000000-3200110102030211)
- [aws.byoc](data-sources--cloud_link--reference--group-001.md#canonical-1033213001233100-3203112202003030-0310131022131311-2330001032132113-0301131303010200-1301121103323332-0232201231230233-2133300030122323)
- [aws.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-1032332332221332-0102020012232333-0310303130112211-3011110333000321-3020013110000301-1101131123102203-3123131221012002-1012333233132121)
- [aws.byoc.connections.auth_key](data-sources--cloud_link--reference--group-001.md#canonical-3013332200201212-2221123020103131-0033013301231101-2031233110103331-1322031011032202-0131121303312003-3332112101310312-1022031211032101)
- aws.byoc.connections.auth_key.clear_secret_info

<a id="canonical-3200303121210030-0113021203021002-2023202031220231-2220320001211211-2113311020303123-3200021002113100-3000211103313032-3100232020330003"></a>

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

<a id="canonical-3200213231213111-3212301022212231-1232112010301023-3132123000201200-3330233022330032-3233111210123301-2012112320300331-1022331303021233"></a>

## Direct properties — clear_secret_info / 301132102301 / 3

<a id="canonical-2113033331003203-2023202032311131-0123123232101133-2101000221123212-2030010122202332-2230232120333031-1311130332312313-0222000122220111"></a>

<a id="canonical-1313213222221301-2121223031111301-3031113033211311-1331222320220133-2111201201323032-2221232120332320-1233212210110330-2212030313333011"></a>

## provider_ref property — clear_secret_info / 301132102301 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1100222210031212-1312303012111120-3121012101131303-1233002003322301-3130220020202233-2032132203332212-3230100231101100-1011101020231022"></a>

<a id="canonical-0332120130130121-0123023222001012-3332123300220202-3110221213332210-0011330003203300-0011033110101013-2323021013033020-3212020100331201"></a>

## URL property — clear_secret_info / 301132102301 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0211312301231302-3032212322320033-1130130013333203-2013131332200200-3232210110131001-2233110302212232-3011212002000222-0232112100201013"></a>

## Next pages — clear_secret_info / 301132102301 / 6

- [aws.byoc.connections.auth_key](data-sources--cloud_link--reference--group-001.md#canonical-3013332200201212-2221123020103131-0033013301231101-2031233110103331-1322031011032202-0131121303312003-3332112101310312-1022031211032101)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)

<a id="canonical-1223231323121200-0331331200303102-3033122032332332-2330012220000111-2313000001020112-1202033210133001-3132303113020302-1300321221311220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311021200300303-1111030130322211-3322312013213231-0101133230011331-2333131210301220-1100221203330303-0111210101113323-0032001002310330"></a>

## aws.byoc.connections.IPv4 — IPv4 / 002222202202 / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-1213002220021231-0321302131202132-0123330031233010-0020203103220000-1121330212132301-3103210130020013-1201230020112100-0121202133202012)
- [aws](data-sources--cloud_link--reference--group-001.md#canonical-3332012130312211-3232211102111223-2232333321133131-0322312033131212-1013100303010313-0303323220310300-0010310203000000-3200110102030211)
- [aws.byoc](data-sources--cloud_link--reference--group-001.md#canonical-1033213001233100-3203112202003030-0310131022131311-2330001032132113-0301131303010200-1301121103323332-0232201231230233-2133300030122323)
- [aws.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-1032332332221332-0102020012232333-0310303130112211-3011110333000321-3020013110000301-1101131123102203-3123131221012002-1012333233132121)
- aws.byoc.connections.IPv4

<a id="canonical-3330001321022100-2123331012113313-2020113223231022-1312332201131122-0031312101221211-0213021223122032-2203323220032103-1013010201111010"></a>

Type: `"single"`. Computed.

Configure BGP IPv4 peering for endpoints.

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

<a id="canonical-1022220221321101-1132130211322013-0013131033102212-1201100121223313-0332011112201001-0332100002232003-0132230102311331-0212203203232022"></a>

## Direct properties — IPv4 / 002222202202 / 3

<a id="canonical-0101031100023123-1313021103323101-3300220013220211-2333022032132100-3220103120103320-1000322233121300-3301123130212203-1220012210313032"></a>

<a id="canonical-1231111213311203-0210001212002132-0121222010023200-3100201213102123-0332323120223000-2233303011303020-3302331303313231-0201010203013301"></a>

## aws_router_peer_address property — IPv4 / 002222202202 / 4

Type: `"string"`. Computed.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2001020331131130-1030232211103231-0021222032221231-3211312023313012-0320220200333113-3212110103010120-2301233320322230-3021331233220131"></a>

<a id="canonical-0002211323011032-0301203103233030-0112231213302323-0301000002033111-3211003302130303-0031133331012333-1010211333033132-0101120121113330"></a>

## router_peer_address property — IPv4 / 002222202202 / 5

Type: `"string"`. Computed.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2021231223311231-3323121101013300-2011332332201200-2120013223031322-2313232200320213-0221222211311103-0113332221223203-0022011113132033"></a>

## Next pages — IPv4 / 002222202202 / 6

- [aws.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-1032332332221332-0102020012232333-0310303130112211-3011110333000321-3020013110000301-1101131123102203-3123131221012002-1012333233132121)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)

<a id="canonical-1223130301101102-1303122312123122-1122011333113020-3223300122331102-2233110331213020-2000220202323332-3112023302000123-3022010332333320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2322022102020012-3131312201133101-2230233311100231-1233131312220311-0201203010111210-0222330131220011-0210233321202101-0011233103110013"></a>

## aws.byoc.connections.metadata — metadata / 003020223300 / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-1213002220021231-0321302131202132-0123330031233010-0020203103220000-1121330212132301-3103210130020013-1201230020112100-0121202133202012)
- [aws](data-sources--cloud_link--reference--group-001.md#canonical-3332012130312211-3232211102111223-2232333321133131-0322312033131212-1013100303010313-0303323220310300-0010310203000000-3200110102030211)
- [aws.byoc](data-sources--cloud_link--reference--group-001.md#canonical-1033213001233100-3203112202003030-0310131022131311-2330001032132113-0301131303010200-1301121103323332-0232201231230233-2133300030122323)
- [aws.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-1032332332221332-0102020012232333-0310303130112211-3011110333000321-3020013110000301-1101131123102203-3123131221012002-1012333233132121)
- aws.byoc.connections.metadata

<a id="canonical-1013112002012100-3110213113020303-2323200233203302-1311103222122320-3212303030132023-0000230202031300-2111111001321313-0333221213333333"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-2122032320322231-3033010310313300-0233113000203323-2223122030311332-2031133220120022-2310010231330322-3222211032202002-2011123002022112"></a>

## Direct properties — metadata / 003020223300 / 3

<a id="canonical-1323321302020332-2110020300311002-1032320113320002-1213121120310323-1220321013200220-1203200030322233-1113222310312022-3103330213000030"></a>

<a id="canonical-1302010202332002-0010220131220120-1103011001002130-3101030020032302-2010300001223311-1012010210332023-0013320302331333-2123232201201012"></a>

## description_spec property — metadata / 003020223300 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-0301003120300003-0001010030202223-0211013212123100-1222032333333301-2003331332111000-1110013010312310-3310333222232302-0321020000012101"></a>

<a id="canonical-2200003011102013-1313200122200212-2313132321313311-3001030233132311-1201222111333211-3220200021202022-1033230233213021-1202301220302030"></a>

## name property — metadata / 003020223300 / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2111011132331203-1113132233121322-3312020322213212-0013112223302032-3022311102223300-0021031320133012-3130000111122233-2130002320332001"></a>

## Next pages — metadata / 003020223300 / 6

- [aws.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-1032332332221332-0102020012232333-0310303130112211-3011110333000321-3020013110000301-1101131123102203-3123131221012002-1012333233132121)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)

<a id="canonical-2212100312030011-2222203211011323-2122230012032120-2113321120213000-2013301113031103-2200210132321211-1132201103012200-0010010121302130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212301300221011-0132032000302011-1203322031023131-1211101020333300-0311201002233332-1201003330110300-2321010002023032-3023313010322210"></a>

## aws.byoc.connections.system_generated_name — system_generated_name / 101311203131 / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-1213002220021231-0321302131202132-0123330031233010-0020203103220000-1121330212132301-3103210130020013-1201230020112100-0121202133202012)
- [aws](data-sources--cloud_link--reference--group-001.md#canonical-3332012130312211-3232211102111223-2232333321133131-0322312033131212-1013100303010313-0303323220310300-0010310203000000-3200110102030211)
- [aws.byoc](data-sources--cloud_link--reference--group-001.md#canonical-1033213001233100-3203112202003030-0310131022131311-2330001032132113-0301131303010200-1301121103323332-0232201231230233-2133300030122323)
- [aws.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-1032332332221332-0102020012232333-0310303130112211-3011110333000321-3020013110000301-1101131123102203-3123131221012002-1012333233132121)
- aws.byoc.connections.system_generated_name

<a id="canonical-1123131121020211-0011101130001121-3323323333013332-3210322012131133-1022201311221323-3010303132000331-3003303321032203-1210020123323322"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1020120012131213-0131132110200321-0022230231302123-2231333120223332-2220203031122330-0231033131133302-1023102101100333-1130132011133202"></a>

## Direct properties — system_generated_name / 101311203131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3210103221033100-3102131201013013-2332320222211110-2112020221011300-1211133123001131-1030320211212012-2021302202002000-2100313030223020"></a>

## Next pages — system_generated_name / 101311203131 / 4

- [aws.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-1032332332221332-0102020012232333-0310303130112211-3011110333000321-3020013110000301-1101131123102203-3123131221012002-1012333233132121)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)

<a id="canonical-1121222321001003-0301130222132322-3133021003331210-1023121231212320-2033201232121230-1210032112221212-2122331210213200-2213211333131300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323022023120212-0121101322220001-2102213200322302-1011020102330222-1103122331302113-0013002223321022-0101001011301332-2332221323302111"></a>

## disabled — disabled / 123033222033 / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-1213002220021231-0321302131202132-0123330031233010-0020203103220000-1121330212132301-3103210130020013-1201230020112100-0121202133202012)
- disabled

<a id="canonical-0023213231102121-1110303231330113-2132223121123132-3223022331310132-2202121011001302-3230100010233023-3132220320232013-2331032032213330"></a>

Type: `["object", {}]`. Computed.

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

- [disabled](data-sources--cloud_link--reference--group-001.md#canonical-0023213231102121-1110303231330113-2132223121123132-3223022331310132-2202121011001302-3230100010233023-3132220320232013-2331032032213330)
- [enabled](data-sources--cloud_link--reference--group-001.md#canonical-2322221213320111-1300302120212320-3010203021000130-1203112101301001-0333123122220310-1023311000210321-0122123133010200-1001131223300232)

Select alternatives according to the provider validators above.

<a id="canonical-2112033000221220-0123002102322001-3120120013303020-3302133301212221-3013122223122120-3032013112101230-3212200202312211-1131303101110211"></a>

## Direct properties — disabled / 123033222033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2233232132223321-1111202102322100-3221203322231223-2301300020201310-2221200003301301-0132222133130322-3100232300223332-1301323032220202"></a>

## Next pages — disabled / 123033222033 / 4

- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-1213002220021231-0321302131202132-0123330031233010-0020203103220000-1121330212132301-3103210130020013-1201230020112100-0121202133202012)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)

<a id="canonical-2133211113302331-1111103313201201-2300313320232302-1001210232102002-3322223320212110-1202033223333021-1002113012122212-3121000132101120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320233021121011-0301110103302111-2222303120201010-2023331121231303-0133113300230222-0132311013221020-0312123122023002-1231113332230313"></a>

## enabled — enabled / 130132332003 / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-1213002220021231-0321302131202132-0123330031233010-0020203103220000-1121330212132301-3103210130020013-1201230020112100-0121202133202012)
- enabled

<a id="canonical-2322221213320111-1300302120212320-3010203021000130-1203112101301001-0333123122220310-1023311000210321-0122123133010200-1001131223300232"></a>

Type: `"single"`. Computed.

CloudLink ADN Network Config.

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

<a id="canonical-3021130102131223-0330220202000202-0312111202003133-1100131210033321-2320132031231330-2231332123020333-1233002321033200-2033120031311332"></a>

## Direct properties — enabled / 130132332003 / 3

<a id="canonical-1120310010202113-3200033033010032-0200113333220231-1020000210102020-0123000233231233-2013232220101100-3333303312203302-3101330331301001"></a>

<a id="canonical-0332222313210000-0233313300012333-2323001233212223-3220100300232320-3131213223122121-3021320031231003-0023031220333010-1032210202323232"></a>

## cloudlink_network_name property — enabled / 130132332003 / 4

Type: `"string"`. Computed.

Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN
network. To provision a Private ADN network, please contact F5 Distributed Cloud support.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0331121301311301-1321201200103330-3303331322210310-3021200103022323-1021012332331322-2322231123102200-2331112122103200-0222303021220100"></a>

## Next pages — enabled / 130132332003 / 5

- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-1213002220021231-0321302131202132-0123330031233010-0020203103220000-1121330212132301-3103210130020013-1201230020112100-0121202133202012)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)

<a id="canonical-1132023021123233-0321010311212023-0313303231201123-3010111320000110-0312123230211210-2100233101120312-2330020333300003-0210301202220230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121030213020323-1020211110212202-3100331220303312-1321000212201131-2111230032021003-2331113123110321-1122301233212212-3321323032333323"></a>

## gcp — gcp / 321022310012 / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-1213002220021231-0321302131202132-0123330031233010-0020203103220000-1121330212132301-3103210130020013-1201230020112100-0121202133202012)
- gcp

<a id="canonical-3030201223313101-3333102310033121-0303300311223113-2330011201130221-3310300212322123-1133132330001231-2313323111313003-3023000132120223"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0010221031231010-0210203232003332-2003313010300323-2110201002021212-3002323112130023-2233013023332211-1031101302212223-0021103331203201"></a>

## Direct properties — gcp / 321022310012 / 3

- [byoc](data-sources--cloud_link--reference--group-001.md#canonical-1002211013312320-2302013121111313-3102213302121321-0120030320121333-0233200223230111-1211103211032121-0021300212220200-2323030011301021): complete subsection reference.

- [gcp_cred](data-sources--cloud_link--reference--group-001.md#canonical-2320030130032030-0102123132202311-3132323022011300-3013011010330131-2302322013333213-2120300333221032-1100311003120012-1331231330022012): complete subsection reference.

<a id="canonical-3312010131131320-3233312022211321-1002230100202311-2031220020112022-3202001203322220-3322000323030021-3203131331122101-1112130302333331"></a>

## Next pages — gcp / 321022310012 / 4

- [gcp.byoc](data-sources--cloud_link--reference--group-001.md#canonical-1002211013312320-2302013121111313-3102213302121321-0120030320121333-0233200223230111-1211103211032121-0021300212220200-2323030011301021)
- [gcp.gcp_cred](data-sources--cloud_link--reference--group-001.md#canonical-2320030130032030-0102123132202311-3132323022011300-3013011010330131-2302322013333213-2120300333221032-1100311003120012-1331231330022012)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-1213002220021231-0321302131202132-0123330031233010-0020203103220000-1121330212132301-3103210130020013-1201230020112100-0121202133202012)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)

<a id="canonical-1002211013312320-2302013121111313-3102213302121321-0120030320121333-0233200223230111-1211103211032121-0021300212220200-2323030011301021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011312330222320-0102320100213331-2120303030131333-1302332010233013-1313312001322121-0213230202123020-3111013320321202-0331321330203130"></a>

## gcp.byoc — byoc / 012030223133 / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-1213002220021231-0321302131202132-0123330031233010-0020203103220000-1121330212132301-3103210130020013-1201230020112100-0121202133202012)
- [gcp](data-sources--cloud_link--reference--group-001.md#canonical-1132023021123233-0321010311212023-0313303231201123-3010111320000110-0312123230211210-2100233101120312-2330020333300003-0210301202220230)
- gcp.byoc

<a id="canonical-1200222333111311-3210200113212201-3320112232321003-1320023122121011-1032030312330222-0101012303003321-2130322310003321-3031200323300031"></a>

Type: `"single"`. Computed.

GCP Bring Your Own Connections. List of GCP Bring You Own Connections.

Upstream description:

List of GCP Bring You Own Connections.

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

<a id="canonical-1221022210321002-0033233013210321-1013132231133210-1122110203302033-2133111012110313-0001101031123231-0230202301011203-0203301223123301"></a>

## Direct properties — byoc / 012030223133 / 3

- [connections](data-sources--cloud_link--reference--group-001.md#canonical-3211332303233212-3111110333000133-2201033013212221-0333222331221212-0131011203301013-1222023212130213-0101330021003310-2020323330011333): complete subsection reference.

<a id="canonical-2232203131030130-3303123222320030-2212333111311020-1112200211020322-3333122021021102-1330201113111221-0122131222100132-0313310002022230"></a>

## Next pages — byoc / 012030223133 / 4

- [gcp.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-3211332303233212-3111110333000133-2201033013212221-0333222331221212-0131011203301013-1222023212130213-0101330021003310-2020323330011333)
- [gcp](data-sources--cloud_link--reference--group-001.md#canonical-1132023021123233-0321010311212023-0313303231201123-3010111320000110-0312123230211210-2100233101120312-2330020333300003-0210301202220230)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)

<a id="canonical-3211332303233212-3111110333000133-2201033013212221-0333222331221212-0131011203301013-1222023212130213-0101330021003310-2020323330011333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2112223303001302-1313231203020120-1310222033020102-0023133101220132-2123211102230203-2023033010311310-3320121122022332-2313231131300013"></a>

## gcp.byoc.connections — connections / 121123021330 / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-1213002220021231-0321302131202132-0123330031233010-0020203103220000-1121330212132301-3103210130020013-1201230020112100-0121202133202012)
- [gcp](data-sources--cloud_link--reference--group-001.md#canonical-1132023021123233-0321010311212023-0313303231201123-3010111320000110-0312123230211210-2100233101120312-2330020333300003-0210301202220230)
- [gcp.byoc](data-sources--cloud_link--reference--group-001.md#canonical-1002211013312320-2302013121111313-3102213302121321-0120030320121333-0233200223230111-1211103211032121-0021300212220200-2323030011301021)
- gcp.byoc.connections

<a id="canonical-0032210302131311-2210232330312230-2213220002030103-2301022211132102-0122112230113203-2100333023122020-2201301312213123-3322133011011013"></a>

Type: `"list"`. Computed.

Each 'Bring Your Own Connection' represents a virtual connection that the customer has provisioned
in the Cloud (.

Upstream description:

Each 'Bring Your Own Connection' represents a virtual connection that the customer has provisioned
in the Cloud (example: AWS Direct Connect). F5XC will orchestrate networking resources in the cloud
to facilitate seamless private connectivity.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2110003201233013-3332132211230101-2012202201100031-0230100012020031-1211310122331000-2222131221222321-2303331102011121-2003113223033011"></a>

## Direct properties — connections / 121123021330 / 3

<a id="canonical-2323212021122031-3330323120320220-0201323012223101-1110223230232300-0320210203102032-3312221333021132-1033103133222200-1320212130033020"></a>

<a id="canonical-0110120222211102-3210301001330323-2020001002202332-2003230230120122-3232303231131310-0013313031312033-1021310013310203-1033222210012320"></a>

## interconnect_attachment_name property — connections / 121123021330 / 4

Type: `"string"`. Computed.

Name of already-existing GCP Cloud Interconnect Attachment.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [metadata](data-sources--cloud_link--reference--group-001.md#canonical-0203032011003003-3112100022023310-0210203211333311-0011230300100111-1211232013233100-0000313102310212-2231031203202222-1113232110110211): complete subsection reference.

<a id="canonical-1333002302032010-1002003320003222-1103133310003001-1223022020010222-0013100313012131-1030023202021231-2113321201030033-1232131121111221"></a>

<a id="canonical-3131201301112330-0110013201330030-1322230111203233-2311013220023230-2202333220033333-2102132010120221-3032331202313021-2000131232100020"></a>

## project property — connections / 121123021330 / 5

Type: `"string"`. Computed.

Exclusive with \[same\_as\_credential\] Specify a GCP Project for the interconnect attachment.

Upstream description:

Exclusive with \[same\_as\_credential\] Specify a GCP Project for the interconnect attachment.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3313113100112223-2303203101223222-0313322202221103-2300322230221322-2300230202313122-3030020202011113-0021303330303232-3212130303030102"></a>

<a id="canonical-2232021100011112-0031002101220330-3211220210000320-2112222123201101-2021213123132301-2310001200103121-3211002212222212-0300030330320120"></a>

## region property — connections / 121123021330 / 6

Type: `"string"`. Computed.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [same_as_credential](data-sources--cloud_link--reference--group-001.md#canonical-0211022031021233-2323031302221212-3231231233333323-2321232311230331-3222120312310302-1102113133023123-1120030110330131-2003131030300110): complete subsection reference.

<a id="canonical-2120210303332213-3133032101233020-0220312000123213-3020303332332100-1120211020123033-0022013032102110-2201203130202320-2302012133110100"></a>

## Next pages — connections / 121123021330 / 7

- [gcp.byoc.connections.metadata](data-sources--cloud_link--reference--group-001.md#canonical-0203032011003003-3112100022023310-0210203211333311-0011230300100111-1211232013233100-0000313102310212-2231031203202222-1113232110110211)
- [gcp.byoc.connections.same_as_credential](data-sources--cloud_link--reference--group-001.md#canonical-0211022031021233-2323031302221212-3231231233333323-2321232311230331-3222120312310302-1102113133023123-1120030110330131-2003131030300110)
- [gcp.byoc](data-sources--cloud_link--reference--group-001.md#canonical-1002211013312320-2302013121111313-3102213302121321-0120030320121333-0233200223230111-1211103211032121-0021300212220200-2323030011301021)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)

<a id="canonical-0203032011003003-3112100022023310-0210203211333311-0011230300100111-1211232013233100-0000313102310212-2231031203202222-1113232110110211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201313213112010-1000311012110211-2132313011032230-1312013022012310-3112100122010202-1123331101232103-1301232102013221-1221031112130102"></a>

## gcp.byoc.connections.metadata — metadata / 131202020303 / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-1213002220021231-0321302131202132-0123330031233010-0020203103220000-1121330212132301-3103210130020013-1201230020112100-0121202133202012)
- [gcp](data-sources--cloud_link--reference--group-001.md#canonical-1132023021123233-0321010311212023-0313303231201123-3010111320000110-0312123230211210-2100233101120312-2330020333300003-0210301202220230)
- [gcp.byoc](data-sources--cloud_link--reference--group-001.md#canonical-1002211013312320-2302013121111313-3102213302121321-0120030320121333-0233200223230111-1211103211032121-0021300212220200-2323030011301021)
- [gcp.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-3211332303233212-3111110333000133-2201033013212221-0333222331221212-0131011203301013-1222023212130213-0101330021003310-2020323330011333)
- gcp.byoc.connections.metadata

<a id="canonical-2223011221323302-3101230110031131-3020230122301023-1230010213023332-1332010310010323-2202010301011100-0333112013333111-0330001223122303"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-1220030313200110-0013300303202310-0211121330320313-0210200010033210-1121201130322212-2120201221002332-2232303323221200-3331033102121030"></a>

## Direct properties — metadata / 131202020303 / 3

<a id="canonical-1001023331231332-3012031213123233-3010233232113202-0023301131202201-2022313322021302-1131202100122201-1130312230021120-1311030332300332"></a>

<a id="canonical-3003211211001331-2120012331223220-0003130102122023-3300020230322310-0200333113302131-1301031211303301-3212311020133231-2210110313233133"></a>

## description_spec property — metadata / 131202020303 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-2032123001020111-0210010111110132-2300011101111112-1321202320102011-1132222122022030-2002133303211220-1120210200032100-2230212322323323"></a>

<a id="canonical-3212020221003020-2101211020203003-1231020210111011-3310122323013210-1023001110112320-0011300300123113-1030311002323302-2112221121112030"></a>

## name property — metadata / 131202020303 / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1232132302320111-1020321031132110-2231312310222333-1333223310121112-1232200122111113-2110010210120003-2032003310012032-3002022000030013"></a>

## Next pages — metadata / 131202020303 / 6

- [gcp.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-3211332303233212-3111110333000133-2201033013212221-0333222331221212-0131011203301013-1222023212130213-0101330021003310-2020323330011333)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)

<a id="canonical-0211022031021233-2323031302221212-3231231233333323-2321232311230331-3222120312310302-1102113133023123-1120030110330131-2003131030300110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310111110323230-3301322130211002-2211202321023322-1220302302203122-1102333303122313-2003212330232132-3132130233221321-1313130133131022"></a>

## gcp.byoc.connections.same_as_credential — same_as_credential / 133030123033 / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-1213002220021231-0321302131202132-0123330031233010-0020203103220000-1121330212132301-3103210130020013-1201230020112100-0121202133202012)
- [gcp](data-sources--cloud_link--reference--group-001.md#canonical-1132023021123233-0321010311212023-0313303231201123-3010111320000110-0312123230211210-2100233101120312-2330020333300003-0210301202220230)
- [gcp.byoc](data-sources--cloud_link--reference--group-001.md#canonical-1002211013312320-2302013121111313-3102213302121321-0120030320121333-0233200223230111-1211103211032121-0021300212220200-2323030011301021)
- [gcp.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-3211332303233212-3111110333000133-2201033013212221-0333222331221212-0131011203301013-1222023212130213-0101330021003310-2020323330011333)
- gcp.byoc.connections.same_as_credential

<a id="canonical-0123303130232102-2123010330232121-3301131110101233-1032000021122103-2003320003213111-0310113200203233-3110023202230210-0320231111022033"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0312200210302212-2022101201032131-0213132321223210-0303131220021323-2133001120331231-3222132112330100-1023011010211032-1101101003221321"></a>

## Direct properties — same_as_credential / 133030123033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1110303203220203-3320302033200213-3001113011032202-2202131201002311-1133300012120103-2303300300201333-2310332111113312-3001001113221033"></a>

## Next pages — same_as_credential / 133030123033 / 4

- [gcp.byoc.connections](data-sources--cloud_link--reference--group-001.md#canonical-3211332303233212-3111110333000133-2201033013212221-0333222331221212-0131011203301013-1222023212130213-0101330021003310-2020323330011333)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)

<a id="canonical-2320030130032030-0102123132202311-3132323022011300-3013011010330131-2302322013333213-2120300333221032-1100311003120012-1331231330022012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103100011031232-2021313000103131-0230110233100100-1101301113003332-3000212032130210-3033003123201101-3020333333111122-1202001232020011"></a>

## gcp.gcp_cred — gcp_cred / 010233211321 / 2

Breadcrumbs:

- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)
- [Property reference](data-sources--cloud_link--reference--group-001.md#canonical-1213002220021231-0321302131202132-0123330031233010-0020203103220000-1121330212132301-3103210130020013-1201230020112100-0121202133202012)
- [gcp](data-sources--cloud_link--reference--group-001.md#canonical-1132023021123233-0321010311212023-0313303231201123-3010111320000110-0312123230211210-2100233101120312-2330020333300003-0210301202220230)
- gcp.gcp_cred

<a id="canonical-1120210203310010-3231322002013110-0002323131031222-3330022301023031-0013332201330201-2332032201333012-2111020311102321-3210000230110123"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-0113032230023202-1133203123332033-0232323332103122-3312320230001133-0013221210203333-2313003132132321-2321000030220230-2113321021213230"></a>

## Direct properties — gcp_cred / 010233211321 / 3

<a id="canonical-3223020220320031-0321122330301212-3201321101123133-1011320021122032-3310223301032120-2032311222322332-3303000222011203-0030333211210122"></a>

<a id="canonical-2301300012303200-2113023133221210-2333002201313003-1103330210013001-1313123032320012-0121030322103012-2031302112312332-2000123000113230"></a>

## name property — gcp_cred / 010233211321 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3323032302021223-3311331011133010-0313100320130313-1322313030223213-2231100302213322-3202312333203212-3231331221201301-0103121030111302"></a>

<a id="canonical-1021033030000013-3011113003100111-1122331131200012-3313033211020200-3311203310131330-2322230220210323-1300110121333100-2000231111120230"></a>

## namespace property — gcp_cred / 010233211321 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1102020100103001-1212032301331132-1220130312020313-0111032130001313-1020033033332330-2202022121322223-3022120302331323-0210210232311313"></a>

<a id="canonical-3021231112032300-2023103002131231-1012120032221003-3323031030112212-2232320001020000-0211211321230131-1232110033031122-2031021223110330"></a>

## tenant property — gcp_cred / 010233211321 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1231110103301130-1121223223210311-1101313112203103-1322032323222232-2233122332010302-3323223031312132-0120020032033111-3130033212230320"></a>

## Next pages — gcp_cred / 010233211321 / 7

- [gcp](data-sources--cloud_link--reference--group-001.md#canonical-1132023021123233-0321010311212023-0313303231201123-3010111320000110-0312123230211210-2100233101120312-2330020333300003-0210301202220230)
- [xcsh_cloud_link](../data-sources/cloud_link.md#canonical-0332330102021113-0233121121011202-1311220033203332-0030133101313001-2301002303112120-1100301103023223-3212103000330301-3301101211000103)
