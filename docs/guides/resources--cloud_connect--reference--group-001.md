---
page_title: "xcsh_cloud_connect reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_connect reference."
---

# xcsh_cloud_connect reference

<a id="canonical-3000033123020112-3330311232230011-1111030120033023-0320023332201330-1233113231331330-1101323002122132-1323101202220113-3023111000220221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0332321332111210-1212011313332202-0122322103023210-0211201032230032-3100301231321121-2022132120323232-0223331100123003-3101310001212302"></a>

## Property reference — Property reference / 322220320130 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)
- Property reference

<a id="canonical-1222122211301203-2220322321230103-1133311112011210-2002210201020203-0121121100200130-0011323130012013-2201011120031131-3202301111213023"></a>

## Direct properties — Property reference / 322220320130 / 3

<a id="canonical-1330212111212302-2032202111013213-2010011010010202-3330113010320013-1011302023001131-3123112212031312-1311120100311220-0011222102132303"></a>

<a id="canonical-3133302222013321-3211031002303022-0020333011021222-0001010222101331-2131333103321003-0113332312331210-0210201310001011-1003011112220232"></a>

## annotations property — Property reference / 322220320130 / 4

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

- [aws_provider](resources--cloud_connect--reference--group-001.md#canonical-3103101220221300-3222333301231021-3123032111132110-0202212222332110-2002030232032101-0022221112102322-2300121023130210-2030003033012003): complete subsection reference.

- [azure_vnet_site](resources--cloud_connect--reference--group-001.md#canonical-3021103322130320-1200003010013121-0010020300322100-0312322031223102-3220323002313103-2002300323233103-0021222100212022-3300231133311003): complete subsection reference.

<a id="canonical-0211012101003021-3011211013220312-0121210003133301-3303303312230123-2000333001302213-0123200001320101-2013311310033023-3101130301311112"></a>

<a id="canonical-3331003133222030-1130012203101110-1021312332311013-2211111021212330-0111030000321010-3121023331210220-0201021101213202-3203211123331121"></a>

## description property — Property reference / 322220320130 / 5

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

<a id="canonical-2002103011121231-2131213021333230-0021320331133021-0321332321332001-0130021130000311-1232122001101313-1001102331222310-3331300032223223"></a>

<a id="canonical-1103320233100012-3022000333021330-3331131313200013-2222230131323210-1210310231223301-2230331300323331-1220012303233130-0333321202131201"></a>

## disable property — Property reference / 322220320130 / 6

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

<a id="canonical-1101130020323322-3103030231132332-0111120103313033-2113222022121220-0120212131302123-3221020312303213-0310213010233010-1110133122300232"></a>

<a id="canonical-0002032121031213-2121110021013322-1102133221032311-1012210331203103-0113012303123001-3313312201031210-3032120313030033-1103103300202032"></a>

## ID property — Property reference / 322220320130 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2010022323230302-3012022201130123-2010011111233301-0031312231333112-0122300101233002-3321232123131120-0022010110022133-0322300002012010"></a>

<a id="canonical-3101102130303002-0211222111311202-3232222321020101-3113132113122131-1012221222000102-1123003020030330-2311000312111310-1220311033112022"></a>

## labels property — Property reference / 322220320130 / 8

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

<a id="canonical-0021113210322033-1333302231023203-0333121022130213-2103130313030323-3330133113321020-3133121020123032-3031011131110323-2000322230003012"></a>

<a id="canonical-0212022131310131-0112012222003322-2001321322000113-3030211302033223-2100331013111011-0221010023331101-2103302310013013-0230320201311030"></a>

## name property — Property reference / 322220320130 / 9

Type: `"string"`. Required.

Name of the Cloud Connect. Must be unique within the namespace.

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

<a id="canonical-2330030302200202-0320230001000303-0020000230123011-2330313332122332-2332213003110011-1211212013221330-3020221003021333-0002233211202133"></a>

<a id="canonical-1213002113220022-1200303231330321-2121031213121211-3011313030131330-3011102023310023-3221120321011022-3312013113102312-3222131213000102"></a>

## namespace property — Property reference / 322220320130 / 10

Type: `"string"`. Required.

Namespace where the Cloud Connect is created.

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

- [segment](resources--cloud_connect--reference--group-001.md#canonical-0212201120030202-1202123000003111-2102023221221232-3200021132222310-2120132330002232-1032231132002213-2133103223032230-0321330113002232): complete subsection reference.

- [timeouts](resources--cloud_connect--reference--group-001.md#canonical-0102000200203213-0133023212110222-3021012023202203-0022100231003022-1321013002322210-0303322031111330-1031321201022011-1032333113131332): complete subsection reference.

<a id="canonical-3312213310000131-1133201203201201-1320300310313103-1102100030322000-0103331121003110-2003013200322222-1211212122003321-0032102311022222"></a>

## All schema paths — Property reference / 322220320130 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--cloud_connect--reference--group-001.md#canonical-1330212111212302-2032202111013213-2010011010010202-3330113010320013-1011302023001131-3123112212031312-1311120100311220-0011222102132303) |
| `aws_provider` | [aws_provider](resources--cloud_connect--reference--group-001.md#canonical-2230012111131020-2120033011322111-2223301331332033-3312023303133300-2113310120102102-0032103002200102-3013120330102121-1132120111201131) |
| `aws_provider.aws_tgw_site` | [aws_provider.aws_tgw_site](resources--cloud_connect--reference--group-001.md#canonical-2302230323110333-1302012120202322-1223132033322313-3322200332311120-3333302211331321-2120113221033313-3312231112032012-1022310213030132) |
| `aws_provider.aws_tgw_site.cred` | [aws_provider.aws_tgw_site.cred](resources--cloud_connect--reference--group-001.md#canonical-2021121222113232-3212130230302103-1322223022102223-2201130030030111-1313323313212011-1011111301003232-1121112002202120-3231010223132102) |
| `aws_provider.aws_tgw_site.cred.name` | [aws_provider.aws_tgw_site.cred.name](resources--cloud_connect--reference--group-001.md#canonical-2013312221311102-0131220032201121-2223201311001320-2011302100032001-0321233130000313-0101130322010102-2122331113100121-1203322011222010) |
| `aws_provider.aws_tgw_site.cred.namespace` | [aws_provider.aws_tgw_site.cred.namespace](resources--cloud_connect--reference--group-001.md#canonical-1130123201101301-1320130002120020-2011223230110233-3320003310132123-1010223130233120-2333231303022223-0201000321213013-3221110210220021) |
| `aws_provider.aws_tgw_site.cred.tenant` | [aws_provider.aws_tgw_site.cred.tenant](resources--cloud_connect--reference--group-001.md#canonical-0213302300210012-2100210031311300-2121302101101203-3100233120012221-3100222122312021-1023031133011132-1032232102220003-3032322001311303) |
| `aws_provider.aws_tgw_site.site` | [aws_provider.aws_tgw_site.site](resources--cloud_connect--reference--group-001.md#canonical-1001102123012203-2331122220331011-2220111302330333-2203311102111323-0330300203321101-2311303133300123-2213100103323122-2223211223113311) |
| `aws_provider.aws_tgw_site.site.name` | [aws_provider.aws_tgw_site.site.name](resources--cloud_connect--reference--group-001.md#canonical-2011120113101211-1000010023100132-0103313101303301-2323123223230003-1313201211231113-0022110000133002-0311003231322313-0211020131131230) |
| `aws_provider.aws_tgw_site.site.namespace` | [aws_provider.aws_tgw_site.site.namespace](resources--cloud_connect--reference--group-001.md#canonical-1002132220302330-2302012320223130-2213221330213131-2211212211321030-3223313312101012-3330130002103332-3321032203003010-3000001001312110) |
| `aws_provider.aws_tgw_site.site.tenant` | [aws_provider.aws_tgw_site.site.tenant](resources--cloud_connect--reference--group-001.md#canonical-1201113101221201-0112313203101200-2102201213311003-2300020011133201-2203132231312211-1320020000121111-0032211232011312-1002032320100121) |
| `aws_provider.aws_tgw_site.vpc_attachments` | [aws_provider.aws_tgw_site.vpc_attachments](resources--cloud_connect--reference--group-001.md#canonical-1202113311232111-3312102312222122-3321132113200102-2022001222201003-0012032112132302-1100301312302202-0331010201120031-1302303222023110) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](resources--cloud_connect--reference--group-001.md#canonical-0102231233020322-1000300323213233-2001322010211001-2103221020212110-2122002301010013-1030320230230203-3021212330203223-2120201311102013) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing](resources--cloud_connect--reference--group-001.md#canonical-2020120123120311-2200220210302120-0232300200102213-2133103121312200-1313000023303111-2211320320112311-1332131323211102-1001010120312313) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables](resources--cloud_connect--reference--group-001.md#canonical-2102012233301121-2332210013322102-2311011300021002-1003121320120203-0222023311330330-0000011130211112-3310321031122103-1231233230220033) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables.route_table_id` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables.route_table_id](resources--cloud_connect--reference--group-001.md#canonical-3212310020301311-1122322200233333-2023132222011313-2012321133131020-0200321021322222-1223133032220101-2312311301320111-3201320121333201) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables.static_routes` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables.static_routes](resources--cloud_connect--reference--group-001.md#canonical-3311103230133103-2212123211132010-2011230030030113-2322002011110213-1323202303320310-0203320331302313-3203321132122033-1301312133203202) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route](resources--cloud_connect--reference--group-001.md#canonical-2213001110130003-1313220312011111-1200330303211113-3230033023003200-2122322123310331-1302220111011033-0101002121201103-3302002120210001) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_tables` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_tables](resources--cloud_connect--reference--group-001.md#canonical-2002021300202131-2112021313220322-3322003130031333-0103211000103333-2023112230010032-3132021003031033-3022130322302132-0023101233202110) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables](resources--cloud_connect--reference--group-001.md#canonical-3032220302331311-3303311233201011-2230131321302222-0120211113220323-3102300221131030-1030021113222211-0002300222130131-0103120330210220) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables.route_table_id` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables.route_table_id](resources--cloud_connect--reference--group-001.md#canonical-3011003220122222-0312332121020130-3131330133220310-2103322003103301-2122212001030233-2020032200313032-3230210213132213-1022231010023000) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels](resources--cloud_connect--reference--group-001.md#canonical-2112111113022322-0311101112001213-1313111013031302-2202122330323232-3232023130203230-2223110203103312-3021322210223000-0211110013213233) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.manual_routing` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.manual_routing](resources--cloud_connect--reference--group-001.md#canonical-0221212230300231-1222201233032320-2330000100103113-3311211031110223-0110101132112021-2011023312030113-2321031022220133-0030330302333201) |
| `aws_provider.aws_tgw_site.vpc_attachments.vpc_list.vpc_id` | [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.vpc_id](resources--cloud_connect--reference--group-001.md#canonical-0132322013012131-0001012303101323-3210333103102023-1323102301322010-3002320300213021-3012310102300303-2133100120233313-2212231103332012) |
| `azure_vnet_site` | [azure_vnet_site](resources--cloud_connect--reference--group-001.md#canonical-3320232211002023-2222321132133113-1030211310323103-0001310332330003-3301010333202123-3021122230233322-1321323003011310-3220330000221212) |
| `azure_vnet_site.site` | [azure_vnet_site.site](resources--cloud_connect--reference--group-001.md#canonical-0222323330102223-3300031002033201-3222331010003211-1011202111110320-0002310001031112-0331232322203322-3003030302322130-0032020032303000) |
| `azure_vnet_site.site.name` | [azure_vnet_site.site.name](resources--cloud_connect--reference--group-001.md#canonical-2333332223122210-1233223031210130-3133210220320012-2001001100033321-2033021123201320-3203333300201023-3333130201113132-3122221112332030) |
| `azure_vnet_site.site.namespace` | [azure_vnet_site.site.namespace](resources--cloud_connect--reference--group-001.md#canonical-0202313331012320-0012132010101002-3212012030220100-2322122320032300-1211032311002010-0120002013230331-1330023012103203-1213013303320031) |
| `azure_vnet_site.site.tenant` | [azure_vnet_site.site.tenant](resources--cloud_connect--reference--group-001.md#canonical-3113013123002231-2231203011120332-1033302122122230-0133030301022301-3121101103003130-0210232013332003-1032303211133131-0113113012120122) |
| `azure_vnet_site.vnet_attachments` | [azure_vnet_site.vnet_attachments](resources--cloud_connect--reference--group-001.md#canonical-3003210202311111-1003301310211211-3321021232232313-3202221201331320-1123301133303203-2030223311022303-1230111123313213-1321103330110032) |
| `azure_vnet_site.vnet_attachments.vnet_list` | [azure_vnet_site.vnet_attachments.vnet_list](resources--cloud_connect--reference--group-001.md#canonical-0303021332031221-2323220033203023-0332320121031300-0132013331103203-2120320022301001-3233120021131122-3032311032333121-1312003230233303) |
| `azure_vnet_site.vnet_attachments.vnet_list.custom_routing` | [azure_vnet_site.vnet_attachments.vnet_list.custom_routing](resources--cloud_connect--reference--group-001.md#canonical-2211030201023203-2000132221231211-3301200320032323-0112032110022211-0111111130001103-3122330222120220-1103102001312232-0030221300210331) |
| `azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables` | [azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables](resources--cloud_connect--reference--group-001.md#canonical-2121301200130110-2000111222231002-1111131003312123-1021003331030230-0021313223331302-1023202301021222-1202231200233021-0013113113033310) |
| `azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables.route_table_id` | [azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables.route_table_id](resources--cloud_connect--reference--group-001.md#canonical-3211033203120010-2122301223120122-1120311332211213-2303330210011102-2102221300021110-2112011320320021-1021012322311010-1033322100022131) |
| `azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables.static_routes` | [azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables.static_routes](resources--cloud_connect--reference--group-001.md#canonical-3123020300030303-2200120211111312-2113023322013223-2132001010210002-0223300110321232-3313111011023320-1103133033103113-2203320100320113) |
| `azure_vnet_site.vnet_attachments.vnet_list.default_route` | [azure_vnet_site.vnet_attachments.vnet_list.default_route](resources--cloud_connect--reference--group-001.md#canonical-3120203333233030-1032230201302211-2311333313330321-2132113232120001-3321213122023021-1220201212000303-2213120322011202-0001320102203303) |
| `azure_vnet_site.vnet_attachments.vnet_list.default_route.all_route_tables` | [azure_vnet_site.vnet_attachments.vnet_list.default_route.all_route_tables](resources--cloud_connect--reference--group-001.md#canonical-3311201022122013-1332103001231331-2223133032113103-2313122101113000-0213111300120113-3000030132123110-2110321231012333-2100331231202333) |
| `azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables` | [azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables](resources--cloud_connect--reference--group-001.md#canonical-0130213311112211-2222111120002031-0132310120000223-3221210312131033-2213123202303213-3120133202222122-3323230310122231-1033313132331302) |
| `azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables.route_table_id` | [azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables.route_table_id](resources--cloud_connect--reference--group-001.md#canonical-0330021123033013-3230222122303330-2122013312010323-2102110023213330-1131221133133210-2102320012231202-1102231110013003-0332121121301121) |
| `azure_vnet_site.vnet_attachments.vnet_list.labels` | [azure_vnet_site.vnet_attachments.vnet_list.labels](resources--cloud_connect--reference--group-001.md#canonical-0222002120233301-0223113320230300-0121013133001330-1132203010210000-3133322112020100-0322003131200121-3313113331130221-0232331321121300) |
| `azure_vnet_site.vnet_attachments.vnet_list.manual_routing` | [azure_vnet_site.vnet_attachments.vnet_list.manual_routing](resources--cloud_connect--reference--group-001.md#canonical-2200130320303311-3232302300231322-3231011333031313-0001021031232211-2100122013303200-3131313012021120-0221222232120103-0311323301100002) |
| `azure_vnet_site.vnet_attachments.vnet_list.subscription_id` | [azure_vnet_site.vnet_attachments.vnet_list.subscription_id](resources--cloud_connect--reference--group-001.md#canonical-2332303101311210-2112113202111210-3001220123212023-1323320212122101-1030211121122311-1023223331213300-0132032233002130-0221122022321323) |
| `azure_vnet_site.vnet_attachments.vnet_list.vnet_id` | [azure_vnet_site.vnet_attachments.vnet_list.vnet_id](resources--cloud_connect--reference--group-001.md#canonical-3123112001323223-1213033320320233-1320300321322200-2221210012133303-0333233002003101-3002031313210203-3111331312233211-0303300233013121) |
| `description` | [description](resources--cloud_connect--reference--group-001.md#canonical-0211012101003021-3011211013220312-0121210003133301-3303303312230123-2000333001302213-0123200001320101-2013311310033023-3101130301311112) |
| `disable` | [disable](resources--cloud_connect--reference--group-001.md#canonical-2002103011121231-2131213021333230-0021320331133021-0321332321332001-0130021130000311-1232122001101313-1001102331222310-3331300032223223) |
| `id` | [id](resources--cloud_connect--reference--group-001.md#canonical-1101130020323322-3103030231132332-0111120103313033-2113222022121220-0120212131302123-3221020312303213-0310213010233010-1110133122300232) |
| `labels` | [labels](resources--cloud_connect--reference--group-001.md#canonical-2010022323230302-3012022201130123-2010011111233301-0031312231333112-0122300101233002-3321232123131120-0022010110022133-0322300002012010) |
| `name` | [name](resources--cloud_connect--reference--group-001.md#canonical-0021113210322033-1333302231023203-0333121022130213-2103130313030323-3330133113321020-3133121020123032-3031011131110323-2000322230003012) |
| `namespace` | [namespace](resources--cloud_connect--reference--group-001.md#canonical-2330030302200202-0320230001000303-0020000230123011-2330313332122332-2332213003110011-1211212013221330-3020221003021333-0002233211202133) |
| `segment` | [segment](resources--cloud_connect--reference--group-001.md#canonical-1221133032112201-2320030323022320-1012301202303203-1031211200300032-0102123323230122-1231200112101123-2323021213021002-2320111120032123) |
| `segment.name` | [segment.name](resources--cloud_connect--reference--group-001.md#canonical-0000100301313103-1321202332030333-1232321133233130-2112101213321112-1100212103133211-3233121003301013-0312222112200032-1110222302320320) |
| `segment.namespace` | [segment.namespace](resources--cloud_connect--reference--group-001.md#canonical-2311201100330300-0312120312220313-2121002301033020-1012330221133201-3332212220010031-0203201111013310-3022000021313222-2121013110013120) |
| `segment.tenant` | [segment.tenant](resources--cloud_connect--reference--group-001.md#canonical-3112010120000011-0311123022202000-2201213331132313-2310131023201012-0200121332223332-0121110331330310-2211102100002323-2311012210201322) |
| `timeouts` | [timeouts](resources--cloud_connect--reference--group-001.md#canonical-2101123122131111-0123323323022131-2010031211110200-3230221033101310-3030222020021103-2121332120313013-1331202001222003-3111211303211020) |
| `timeouts.create` | [timeouts.create](resources--cloud_connect--reference--group-001.md#canonical-3302322231002222-1000001230320012-2102230021320232-3111033003212112-1223123133301210-0202112013211332-2000021100212333-0113210221100301) |
| `timeouts.delete` | [timeouts.delete](resources--cloud_connect--reference--group-001.md#canonical-3312330100023333-3111223132200021-0112323220312113-1303210320130003-1300333120300032-3323211033032201-0323122333330330-3220132122101122) |
| `timeouts.read` | [timeouts.read](resources--cloud_connect--reference--group-001.md#canonical-3021202320130010-0131121003003122-1012032203031030-2120032221300112-1210322223301222-2110003000233030-1101132312012231-3001322231222323) |
| `timeouts.update` | [timeouts.update](resources--cloud_connect--reference--group-001.md#canonical-2210233302213033-3011121223011033-1230112030202133-3231000131102231-0123130002203033-3313003022330002-3323123333301332-3103133120120211) |

<a id="canonical-0032323220023230-0000101322133130-2222123101330313-3023101232132123-0311012023231233-2303323130331100-1001231022320020-1131301011111131"></a>

## Next pages — Property reference / 322220320130 / 12

- [aws_provider](resources--cloud_connect--reference--group-001.md#canonical-3103101220221300-3222333301231021-3123032111132110-0202212222332110-2002030232032101-0022221112102322-2300121023130210-2030003033012003)
- [azure_vnet_site](resources--cloud_connect--reference--group-001.md#canonical-3021103322130320-1200003010013121-0010020300322100-0312322031223102-3220323002313103-2002300323233103-0021222100212022-3300231133311003)
- [segment](resources--cloud_connect--reference--group-001.md#canonical-0212201120030202-1202123000003111-2102023221221232-3200021132222310-2120132330002232-1032231132002213-2133103223032230-0321330113002232)
- [timeouts](resources--cloud_connect--reference--group-001.md#canonical-0102000200203213-0133023212110222-3021012023202203-0022100231003022-1321013002322210-0303322031111330-1031321201022011-1032333113131332)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)

<a id="canonical-3103101220221300-3222333301231021-3123032111132110-0202212222332110-2002030232032101-0022221112102322-2300121023130210-2030003033012003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320001113032003-3012232202321201-2231032311310200-1032110110323033-1033013122020330-2020210231000111-3133133120022130-0213320122002211"></a>

## aws_provider — aws_provider / 131020011130 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-3000033123020112-3330311232230011-1111030120033023-0320023332201330-1233113231331330-1101323002122132-1323101202220113-3023111000220221)
- aws_provider

<a id="canonical-2230012111131020-2120033011322111-2223301331332033-3312023303133300-2113310120102102-0032103002200102-3013120330102121-1132120111201131"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: aws\_provider, Azure\_vnet\_site\] Configuration parameter for aws provider.

Upstream description:

Cloud Connect with AWS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-site_type": "[\"aws_tgw_site\"]"
}
```

OneOf alternatives in this subsection:

- [aws_provider](resources--cloud_connect--reference--group-001.md#canonical-2230012111131020-2120033011322111-2223301331332033-3312023303133300-2113310120102102-0032103002200102-3013120330102121-1132120111201131)
- [azure_vnet_site](resources--cloud_connect--reference--group-001.md#canonical-3320232211002023-2222321132133113-1030211310323103-0001310332330003-3301010333202123-3021122230233322-1321323003011310-3220330000221212)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
aws_provider {
  # Configure direct properties listed below.
}
```

<a id="canonical-1101313303312012-0032101223001031-0002022110112323-1003232223232100-3131021111120033-0313100002300032-1213230000122221-3023223302030221"></a>

## Direct properties — aws_provider / 131020011130 / 3

- [aws_tgw_site](resources--cloud_connect--reference--group-001.md#canonical-1100320102231322-1033313001221030-2113103123301132-3232000101000312-1011301322100112-3011313110020121-3031331221113311-2313211203222203): complete subsection reference.

<a id="canonical-2113102322033333-0021210232320212-1200031001131211-2022113122230330-0212212222330200-3330333013013201-3333002030301330-2100002030102213"></a>

## Next pages — aws_provider / 131020011130 / 4

- [aws_provider.aws_tgw_site](resources--cloud_connect--reference--group-001.md#canonical-1100320102231322-1033313001221030-2113103123301132-3232000101000312-1011301322100112-3011313110020121-3031331221113311-2313211203222203)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-3000033123020112-3330311232230011-1111030120033023-0320023332201330-1233113231331330-1101323002122132-1323101202220113-3023111000220221)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)

<a id="canonical-1100320102231322-1033313001221030-2113103123301132-3232000101000312-1011301322100112-3011313110020121-3031331221113311-2313211203222203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122221303213301-2133021233000121-2111123330232200-0223221011221231-0012323220123223-3003222123310223-1211221113012220-2230321232031123"></a>

## aws_provider.aws_tgw_site — aws_tgw_site / 011332031021 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-3000033123020112-3330311232230011-1111030120033023-0320023332201330-1233113231331330-1101323002122132-1323101202220113-3023111000220221)
- [aws_provider](resources--cloud_connect--reference--group-001.md#canonical-3103101220221300-3222333301231021-3123032111132110-0202212222332110-2002030232032101-0022221112102322-2300121023130210-2030003033012003)
- aws_provider.aws_tgw_site

<a id="canonical-2302230323110333-1302012120202322-1223132033322313-3322200332311120-3333302211331321-2120113221033313-3312231112032012-1022310213030132"></a>

Type: `"object"`. single nested block, Optional.

AWS TGW Site Type. Cloud Connect AWS TGW Site Type.

Upstream description:

Cloud Connect AWS TGW Site Type.

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
aws_tgw_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-2203001202313122-0230002111213320-0031130320010030-0033003310101213-1010233311302332-1321330202000112-0201022133012313-3031331033303123"></a>

## Direct properties — aws_tgw_site / 011332031021 / 3

- [cred](resources--cloud_connect--reference--group-001.md#canonical-3312332321213110-3213330110212003-3310122110111331-1231321322333232-3233202120123330-1133000031113111-1113123132210321-3223120103022131): complete subsection reference.

- [site](resources--cloud_connect--reference--group-001.md#canonical-0233330102101201-0312302011112010-3311212032012230-3123312202232022-3321133002331133-1301013302132103-3120000132221021-2231203211032000): complete subsection reference.

- [vpc_attachments](resources--cloud_connect--reference--group-001.md#canonical-2012322120311013-2100022121212203-0111001020222200-1332313201300222-3001000002033020-0011230321132211-1201033112130331-1100132320030110): complete subsection reference.

<a id="canonical-1111322301002303-2211111130111011-1103322102112001-0332301102220020-1331033000211232-0330210301122000-1223001132011002-3323220301323202"></a>

## Next pages — aws_tgw_site / 011332031021 / 4

- [aws_provider.aws_tgw_site.cred](resources--cloud_connect--reference--group-001.md#canonical-3312332321213110-3213330110212003-3310122110111331-1231321322333232-3233202120123330-1133000031113111-1113123132210321-3223120103022131)
- [aws_provider.aws_tgw_site.site](resources--cloud_connect--reference--group-001.md#canonical-0233330102101201-0312302011112010-3311212032012230-3123312202232022-3321133002331133-1301013302132103-3120000132221021-2231203211032000)
- [aws_provider.aws_tgw_site.vpc_attachments](resources--cloud_connect--reference--group-001.md#canonical-2012322120311013-2100022121212203-0111001020222200-1332313201300222-3001000002033020-0011230321132211-1201033112130331-1100132320030110)
- [aws_provider](resources--cloud_connect--reference--group-001.md#canonical-3103101220221300-3222333301231021-3123032111132110-0202212222332110-2002030232032101-0022221112102322-2300121023130210-2030003033012003)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)

<a id="canonical-3312332321213110-3213330110212003-3310122110111331-1231321322333232-3233202120123330-1133000031113111-1113123132210321-3223120103022131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102213212302331-0230010103223010-1221232031323133-3211100302031312-1333221211122322-0021110211212323-2310102131132311-3221112230231123"></a>

## aws_provider.aws_tgw_site.cred — cred / 113211012030 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-3000033123020112-3330311232230011-1111030120033023-0320023332201330-1233113231331330-1101323002122132-1323101202220113-3023111000220221)
- [aws_provider](resources--cloud_connect--reference--group-001.md#canonical-3103101220221300-3222333301231021-3123032111132110-0202212222332110-2002030232032101-0022221112102322-2300121023130210-2030003033012003)
- [aws_provider.aws_tgw_site](resources--cloud_connect--reference--group-001.md#canonical-1100320102231322-1033313001221030-2113103123301132-3232000101000312-1011301322100112-3011313110020121-3031331221113311-2313211203222203)
- aws_provider.aws_tgw_site.cred

<a id="canonical-2021121222113232-3212130230302103-1322223022102223-2201130030030111-1313323313212011-1011111301003232-1121112002202120-3231010223132102"></a>

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
cred {
  # Configure direct properties listed below.
}
```

<a id="canonical-1121112011002122-3203132013302202-1231020321230310-2210013322203110-2332210332123202-2001200200111130-2203130321332101-2013310201011321"></a>

## Direct properties — cred / 113211012030 / 3

<a id="canonical-2013312221311102-0131220032201121-2223201311001320-2011302100032001-0321233130000313-0101130322010102-2122331113100121-1203322011222010"></a>

<a id="canonical-3333002021101323-1010012012220221-3332012332121121-3030303331011112-0003213131030033-2130301203122030-1303020010101303-1211232011213200"></a>

## name property — cred / 113211012030 / 4

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

<a id="canonical-1130123201101301-1320130002120020-2011223230110233-3320003310132123-1010223130233120-2333231303022223-0201000321213013-3221110210220021"></a>

<a id="canonical-3033121101101013-0233012212012100-1330100323032012-2212232212122022-0102321013101311-1130013021112220-3323000230233020-1212232323121031"></a>

## namespace property — cred / 113211012030 / 5

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

<a id="canonical-0213302300210012-2100210031311300-2121302101101203-3100233120012221-3100222122312021-1023031133011132-1032232102220003-3032322001311303"></a>

<a id="canonical-1101213232231320-2121223221002031-1312220203130202-2102210210113232-3302202113201313-0030111032302312-0001301232112031-0031233112013213"></a>

## tenant property — cred / 113211012030 / 6

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

<a id="canonical-3021132023112221-3223310312113221-3022312233120012-1310220322321223-0220212212033202-3011303020320222-1200032312113110-1022223213330133"></a>

## Next pages — cred / 113211012030 / 7

- [aws_provider.aws_tgw_site](resources--cloud_connect--reference--group-001.md#canonical-1100320102231322-1033313001221030-2113103123301132-3232000101000312-1011301322100112-3011313110020121-3031331221113311-2313211203222203)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)

<a id="canonical-0233330102101201-0312302011112010-3311212032012230-3123312202232022-3321133002331133-1301013302132103-3120000132221021-2231203211032000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233023331022133-0111213322320123-1001323123133330-1321030000121300-0012322230001313-1001010312102030-3323100323120130-0111122111031021"></a>

## aws_provider.aws_tgw_site.site — site / 103103220221 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-3000033123020112-3330311232230011-1111030120033023-0320023332201330-1233113231331330-1101323002122132-1323101202220113-3023111000220221)
- [aws_provider](resources--cloud_connect--reference--group-001.md#canonical-3103101220221300-3222333301231021-3123032111132110-0202212222332110-2002030232032101-0022221112102322-2300121023130210-2030003033012003)
- [aws_provider.aws_tgw_site](resources--cloud_connect--reference--group-001.md#canonical-1100320102231322-1033313001221030-2113103123301132-3232000101000312-1011301322100112-3011313110020121-3031331221113311-2313211203222203)
- aws_provider.aws_tgw_site.site

<a id="canonical-1001102123012203-2331122220331011-2220111302330333-2203311102111323-0330300203321101-2311303133300123-2213100103323122-2223211223113311"></a>

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-2133223021020032-1033203111300111-1323110330131233-1011122320323003-3222113101230112-2022221132203122-1301122001312001-3210131310122131"></a>

## Direct properties — site / 103103220221 / 3

<a id="canonical-2011120113101211-1000010023100132-0103313101303301-2323123223230003-1313201211231113-0022110000133002-0311003231322313-0211020131131230"></a>

<a id="canonical-1010023031000100-2231132301320131-1200103130301310-1112130231111002-3011210210210102-1022230030103013-3110323021020021-3321203232223123"></a>

## name property — site / 103103220221 / 4

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

<a id="canonical-1002132220302330-2302012320223130-2213221330213131-2211212211321030-3223313312101012-3330130002103332-3321032203003010-3000001001312110"></a>

<a id="canonical-3313123100233103-1213313003020231-1302223202113213-0020200031231310-3231021213320330-2223112122220122-0002001003221111-0113210102223233"></a>

## namespace property — site / 103103220221 / 5

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

<a id="canonical-1201113101221201-0112313203101200-2102201213311003-2300020011133201-2203132231312211-1320020000121111-0032211232011312-1002032320100121"></a>

<a id="canonical-0022332033301231-0021103203010333-0032230100103320-1311121201221122-0103030103010230-1011012332000301-0311011032120203-1312111331100113"></a>

## tenant property — site / 103103220221 / 6

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

<a id="canonical-1000010023331031-1103013111010110-0113331101121221-3122132222300002-2322210210133121-1032232023102322-3131330221032320-3303133102212331"></a>

## Next pages — site / 103103220221 / 7

- [aws_provider.aws_tgw_site](resources--cloud_connect--reference--group-001.md#canonical-1100320102231322-1033313001221030-2113103123301132-3232000101000312-1011301322100112-3011313110020121-3031331221113311-2313211203222203)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)

<a id="canonical-2012322120311013-2100022121212203-0111001020222200-1332313201300222-3001000002033020-0011230321132211-1201033112130331-1100132320030110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313133023130033-3130232220313132-3021232313322103-3030331321310103-1012213202320032-3202023202003022-1111011132300020-0112231131211010"></a>

## aws_provider.aws_tgw_site.vpc_attachments — vpc_attachments / 233101101121 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-3000033123020112-3330311232230011-1111030120033023-0320023332201330-1233113231331330-1101323002122132-1323101202220113-3023111000220221)
- [aws_provider](resources--cloud_connect--reference--group-001.md#canonical-3103101220221300-3222333301231021-3123032111132110-0202212222332110-2002030232032101-0022221112102322-2300121023130210-2030003033012003)
- [aws_provider.aws_tgw_site](resources--cloud_connect--reference--group-001.md#canonical-1100320102231322-1033313001221030-2113103123301132-3232000101000312-1011301322100112-3011313110020121-3031331221113311-2313211203222203)
- aws_provider.aws_tgw_site.vpc_attachments

<a id="canonical-1202113311232111-3312102312222122-3321132113200102-2022001222201003-0012032112132302-1100301312302202-0331010201120031-1302303222023110"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for vpc attachments.

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
vpc_attachments {
  # Configure direct properties listed below.
}
```

<a id="canonical-1101330012111322-2321313230231321-0010002032220300-0122331031221220-0032023030122110-3223113322122210-0003010312000133-3031133110121101"></a>

## Direct properties — vpc_attachments / 233101101121 / 3

- [vpc_list](resources--cloud_connect--reference--group-001.md#canonical-3200300212101213-0322103321201312-3031332200301200-0130131130213120-0121121301223133-1102023111022131-1132102113110112-0030232031110101): complete subsection reference.

<a id="canonical-0113002002112333-2000220001223003-0120013221130300-2032331321230320-3230122031122200-3031330312012222-0102322202033021-3310010010113223"></a>

## Next pages — vpc_attachments / 233101101121 / 4

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](resources--cloud_connect--reference--group-001.md#canonical-3200300212101213-0322103321201312-3031332200301200-0130131130213120-0121121301223133-1102023111022131-1132102113110112-0030232031110101)
- [aws_provider.aws_tgw_site](resources--cloud_connect--reference--group-001.md#canonical-1100320102231322-1033313001221030-2113103123301132-3232000101000312-1011301322100112-3011313110020121-3031331221113311-2313211203222203)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)

<a id="canonical-3200300212101213-0322103321201312-3031332200301200-0130131130213120-0121121301223133-1102023111022131-1132102113110112-0030232031110101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300013111123020-2210200131003032-2301212223121223-0111330010303122-0330332000302321-2322330213203120-1223133113112302-1303310030111202"></a>

## aws_provider.aws_tgw_site.vpc_attachments.vpc_list — vpc_list / 031233031330 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-3000033123020112-3330311232230011-1111030120033023-0320023332201330-1233113231331330-1101323002122132-1323101202220113-3023111000220221)
- [aws_provider](resources--cloud_connect--reference--group-001.md#canonical-3103101220221300-3222333301231021-3123032111132110-0202212222332110-2002030232032101-0022221112102322-2300121023130210-2030003033012003)
- [aws_provider.aws_tgw_site](resources--cloud_connect--reference--group-001.md#canonical-1100320102231322-1033313001221030-2113103123301132-3232000101000312-1011301322100112-3011313110020121-3031331221113311-2313211203222203)
- [aws_provider.aws_tgw_site.vpc_attachments](resources--cloud_connect--reference--group-001.md#canonical-2012322120311013-2100022121212203-0111001020222200-1332313201300222-3001000002033020-0011230321132211-1201033112130331-1100132320030110)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list

<a id="canonical-0102231233020322-1000300323213233-2001322010211001-2103221020212110-2122002301010013-1030320230230203-3021212330203223-2120201311102013"></a>

Type: `"object"`. list nested block, Optional.

VPC List. Collection of items or values

Upstream description:

Collection of items or values

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("vpc_id"),
  validators.ConflictingListObjectAttributes("custom_routing",
    "default_route"),
  validators.ConflictingListObjectAttributes("custom_routing",
    "manual_routing"),
  validators.ConflictingListObjectAttributes("default_route",
    "manual_routing")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

Terraform syntax:

```terraform
vpc_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0312020011310202-1331122122112130-2212103001131000-1113020130230333-2300103230331103-2122212302231303-3202313310312103-2132102031112321"></a>

## Direct properties — vpc_list / 031233031330 / 3

- [custom_routing](resources--cloud_connect--reference--group-001.md#canonical-3113202333110312-1102031311301233-1330222132123100-3010313232010100-2211120033002201-0101123022332001-1002320221233133-3031201133012021): complete subsection reference.

- [default_route](resources--cloud_connect--reference--group-001.md#canonical-3233310033321120-2110120300102121-1222212213121011-3211203323031022-0323101010102011-3032033101120022-1321001101212331-1323210021310313): complete subsection reference.

- [labels](resources--cloud_connect--reference--group-001.md#canonical-0310333211010221-3003201212210333-1202231323203330-3001221120111001-0120030123123332-1120302002132233-2123230323133300-3122312331330310): complete subsection reference.

- [manual_routing](resources--cloud_connect--reference--group-001.md#canonical-1020113333302101-2030202001320221-3010320333121102-2301330000232212-0323202310103020-2220221103200131-1112302001210021-2320020130203231): complete subsection reference.

<a id="canonical-0132322013012131-0001012303101323-3210333103102023-1323102301322010-3002320300213021-3012310102300303-2133100120233313-2212231103332012"></a>

<a id="canonical-1332300202322113-1111010101220210-2012320300001232-1313201232201130-0112122022101331-3001230233132223-0311222130000233-2230312200103012"></a>

## vpc_id property — vpc_list / 031233031330 / 4

Type: `"string"`. Optional.

Enter the VPC ID of the VPC to be attached.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
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
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

<a id="canonical-3013332132113133-3110112101301021-1003323020011300-2133033011333200-0220210212122100-3133002011212323-1113310120202220-1210320112030333"></a>

## Next pages — vpc_list / 031233031330 / 5

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing](resources--cloud_connect--reference--group-001.md#canonical-3113202333110312-1102031311301233-1330222132123100-3010313232010100-2211120033002201-0101123022332001-1002320221233133-3031201133012021)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route](resources--cloud_connect--reference--group-001.md#canonical-3233310033321120-2110120300102121-1222212213121011-3211203323031022-0323101010102011-3032033101120022-1321001101212331-1323210021310313)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels](resources--cloud_connect--reference--group-001.md#canonical-0310333211010221-3003201212210333-1202231323203330-3001221120111001-0120030123123332-1120302002132233-2123230323133300-3122312331330310)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.manual_routing](resources--cloud_connect--reference--group-001.md#canonical-1020113333302101-2030202001320221-3010320333121102-2301330000232212-0323202310103020-2220221103200131-1112302001210021-2320020130203231)
- [aws_provider.aws_tgw_site.vpc_attachments](resources--cloud_connect--reference--group-001.md#canonical-2012322120311013-2100022121212203-0111001020222200-1332313201300222-3001000002033020-0011230321132211-1201033112130331-1100132320030110)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)

<a id="canonical-3113202333110312-1102031311301233-1330222132123100-3010313232010100-2211120033002201-0101123022332001-1002320221233133-3031201133012021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131303000122120-1122300332103302-0331331312133120-0010010320100303-0210121132001220-3332103220130120-1000332100012231-3301321001012221"></a>

## aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing — custom_routing / 103032313330 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-3000033123020112-3330311232230011-1111030120033023-0320023332201330-1233113231331330-1101323002122132-1323101202220113-3023111000220221)
- [aws_provider](resources--cloud_connect--reference--group-001.md#canonical-3103101220221300-3222333301231021-3123032111132110-0202212222332110-2002030232032101-0022221112102322-2300121023130210-2030003033012003)
- [aws_provider.aws_tgw_site](resources--cloud_connect--reference--group-001.md#canonical-1100320102231322-1033313001221030-2113103123301132-3232000101000312-1011301322100112-3011313110020121-3031331221113311-2313211203222203)
- [aws_provider.aws_tgw_site.vpc_attachments](resources--cloud_connect--reference--group-001.md#canonical-2012322120311013-2100022121212203-0111001020222200-1332313201300222-3001000002033020-0011230321132211-1201033112130331-1100132320030110)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](resources--cloud_connect--reference--group-001.md#canonical-3200300212101213-0322103321201312-3031332200301200-0130131130213120-0121121301223133-1102023111022131-1132102113110112-0030232031110101)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing

<a id="canonical-2020120123120311-2200220210302120-0232300200102213-2133103121312200-1313000023303111-2211320320112311-1332131323211102-1001010120312313"></a>

Type: `"object"`. single nested block, Optional.

AWS Route Table List. AWS Route Table List.

Upstream description:

AWS Route Table List.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("route_tables")}
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
custom_routing {
  # Configure direct properties listed below.
}
```

<a id="canonical-3202030003310022-2113222212120230-2301013021121213-3333323111131313-1122233213023200-1232020201123211-2033311213211210-0031323123122322"></a>

## Direct properties — custom_routing / 103032313330 / 3

- [route_tables](resources--cloud_connect--reference--group-001.md#canonical-2133311032131320-3013333312010132-1213001102130021-2033211121203110-2230021233331030-1230112231022231-0121022113130102-3031030331211000): complete subsection reference.

<a id="canonical-2031213332211211-0011232132011330-2130210130302023-0313123100233121-0230012213131331-2111202232013031-0132110130122121-1330332121201231"></a>

## Next pages — custom_routing / 103032313330 / 4

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables](resources--cloud_connect--reference--group-001.md#canonical-2133311032131320-3013333312010132-1213001102130021-2033211121203110-2230021233331030-1230112231022231-0121022113130102-3031030331211000)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](resources--cloud_connect--reference--group-001.md#canonical-3200300212101213-0322103321201312-3031332200301200-0130131130213120-0121121301223133-1102023111022131-1132102113110112-0030232031110101)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)

<a id="canonical-2133311032131320-3013333312010132-1213001102130021-2033211121203110-2230021233331030-1230112231022231-0121022113130102-3031030331211000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3111122333010131-0300223031312130-1301301233020023-2113113130111331-3220220022022201-3011311312311302-2020133322011231-0023131032311021"></a>

## aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables — route_tables / 130311213201 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-3000033123020112-3330311232230011-1111030120033023-0320023332201330-1233113231331330-1101323002122132-1323101202220113-3023111000220221)
- [aws_provider](resources--cloud_connect--reference--group-001.md#canonical-3103101220221300-3222333301231021-3123032111132110-0202212222332110-2002030232032101-0022221112102322-2300121023130210-2030003033012003)
- [aws_provider.aws_tgw_site](resources--cloud_connect--reference--group-001.md#canonical-1100320102231322-1033313001221030-2113103123301132-3232000101000312-1011301322100112-3011313110020121-3031331221113311-2313211203222203)
- [aws_provider.aws_tgw_site.vpc_attachments](resources--cloud_connect--reference--group-001.md#canonical-2012322120311013-2100022121212203-0111001020222200-1332313201300222-3001000002033020-0011230321132211-1201033112130331-1100132320030110)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](resources--cloud_connect--reference--group-001.md#canonical-3200300212101213-0322103321201312-3031332200301200-0130131130213120-0121121301223133-1102023111022131-1132102113110112-0030232031110101)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing](resources--cloud_connect--reference--group-001.md#canonical-3113202333110312-1102031311301233-1330222132123100-3010313232010100-2211120033002201-0101123022332001-1002320221233133-3031201133012021)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing.route_tables

<a id="canonical-2102012233301121-2332210013322102-2311011300021002-1003121320120203-0222023311330330-0000011130211112-3310321031122103-1231233230220033"></a>

Type: `"object"`. list nested block, Optional.

List of route tables. Route Tables.

Upstream description:

Route Tables.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("static_routes")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 200,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "200",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "200",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
route_tables {
  # Configure direct properties listed below.
}
```

<a id="canonical-0000210333020020-3323331220030320-0112103110210013-0301323110130311-3211321230111122-2320320131022001-3330010010222320-3003133220001313"></a>

## Direct properties — route_tables / 130311213201 / 3

<a id="canonical-3212310020301311-1122322200233333-2023132222011313-2012321133131020-0200321021322222-1223133032220101-2312311301320111-3201320121333201"></a>

<a id="canonical-0320210313230003-3120102312030122-2313313303212110-0231033202211302-2233131331130031-0330300230010130-1303301332030030-1012023203113122"></a>

## route_table_id property — route_tables / 130311213201 / 4

Type: `"string"`. Optional.

Route table ID. Route table ID.

Upstream description:

Route table ID.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^(rtb-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(rtb-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(rtb-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

<a id="canonical-3311103230133103-2212123211132010-2011230030030113-2322002011110213-1323202303320310-0203320331302313-3203321132122033-1301312133203202"></a>

<a id="canonical-0111321101133333-3131022000320100-0130302333131331-3230202013113112-1301311331031021-0130130322300022-2120220131211230-2220111100223001"></a>

## static_routes property — route_tables / 130311213201 / 5

Type: `["list", "string"]`. Optional.

Static Routes. List of Static Routes.

Upstream description:

List of Static Routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 50),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 50,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 50,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0131133211233231-1322122302023032-2100213000030111-0030322200132110-1113113103230221-1300221203102211-2302103323100322-0123123101201101"></a>

## Next pages — route_tables / 130311213201 / 6

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.custom_routing](resources--cloud_connect--reference--group-001.md#canonical-3113202333110312-1102031311301233-1330222132123100-3010313232010100-2211120033002201-0101123022332001-1002320221233133-3031201133012021)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)

<a id="canonical-3233310033321120-2110120300102121-1222212213121011-3211203323031022-0323101010102011-3032033101120022-1321001101212331-1323210021310313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233233223332011-3110332303101221-0100021203200132-0122312023310212-2022100131032210-2320331030212302-1301130120101033-1102231113033320"></a>

## aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route — default_route / 023022000033 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-3000033123020112-3330311232230011-1111030120033023-0320023332201330-1233113231331330-1101323002122132-1323101202220113-3023111000220221)
- [aws_provider](resources--cloud_connect--reference--group-001.md#canonical-3103101220221300-3222333301231021-3123032111132110-0202212222332110-2002030232032101-0022221112102322-2300121023130210-2030003033012003)
- [aws_provider.aws_tgw_site](resources--cloud_connect--reference--group-001.md#canonical-1100320102231322-1033313001221030-2113103123301132-3232000101000312-1011301322100112-3011313110020121-3031331221113311-2313211203222203)
- [aws_provider.aws_tgw_site.vpc_attachments](resources--cloud_connect--reference--group-001.md#canonical-2012322120311013-2100022121212203-0111001020222200-1332313201300222-3001000002033020-0011230321132211-1201033112130331-1100132320030110)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](resources--cloud_connect--reference--group-001.md#canonical-3200300212101213-0322103321201312-3031332200301200-0130131130213120-0121121301223133-1102023111022131-1132102113110112-0030232031110101)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route

<a id="canonical-2213001110130003-1313220312011111-1200330303211113-3230033023003200-2122322123310331-1302220111011033-0101002121201103-3302002120210001"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for default route.

Upstream description:

Select Override Default Route Choice.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("all_route_tables",
    "selective_route_tables")}
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
  "x-ves-oneof-field-default_route_choice": "[\"all_route_tables\",\"selective_route_tables\"]"
}
```

Terraform syntax:

```terraform
default_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-3200011330203202-2022132302201210-3123320312012022-0001303101010031-0013320110211100-3213221213123201-1332202331103233-0222001110202210"></a>

## Direct properties — default_route / 023022000033 / 3

- [all_route_tables](resources--cloud_connect--reference--group-001.md#canonical-2010210330133012-1113331121220211-2301211111021022-0220023012332111-2233033303102131-3330031330102111-0330312233333010-3133001110201101): complete subsection reference.

- [selective_route_tables](resources--cloud_connect--reference--group-001.md#canonical-2221223103020031-3233303133133123-1311301323021132-0222313210010320-0121213023132100-2030020101221211-3222121030210213-1112030230300030): complete subsection reference.

<a id="canonical-1313202131201333-2001030221232211-0232301311201011-2230001121323020-1133011032133222-2033122301321021-2021333031323220-3220230033001130"></a>

## Next pages — default_route / 023022000033 / 4

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_tables](resources--cloud_connect--reference--group-001.md#canonical-2010210330133012-1113331121220211-2301211111021022-0220023012332111-2233033303102131-3330031330102111-0330312233333010-3133001110201101)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables](resources--cloud_connect--reference--group-001.md#canonical-2221223103020031-3233303133133123-1311301323021132-0222313210010320-0121213023132100-2030020101221211-3222121030210213-1112030230300030)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](resources--cloud_connect--reference--group-001.md#canonical-3200300212101213-0322103321201312-3031332200301200-0130131130213120-0121121301223133-1102023111022131-1132102113110112-0030232031110101)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)

<a id="canonical-2010210330133012-1113331121220211-2301211111021022-0220023012332111-2233033303102131-3330031330102111-0330312233333010-3133001110201101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0132200220221131-0202102222010102-1212131320032313-0103232333223233-1001222230202032-2021013030330133-3303133230020010-0230033302322212"></a>

## aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_tables — all_route_tables / 012103120120 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-3000033123020112-3330311232230011-1111030120033023-0320023332201330-1233113231331330-1101323002122132-1323101202220113-3023111000220221)
- [aws_provider](resources--cloud_connect--reference--group-001.md#canonical-3103101220221300-3222333301231021-3123032111132110-0202212222332110-2002030232032101-0022221112102322-2300121023130210-2030003033012003)
- [aws_provider.aws_tgw_site](resources--cloud_connect--reference--group-001.md#canonical-1100320102231322-1033313001221030-2113103123301132-3232000101000312-1011301322100112-3011313110020121-3031331221113311-2313211203222203)
- [aws_provider.aws_tgw_site.vpc_attachments](resources--cloud_connect--reference--group-001.md#canonical-2012322120311013-2100022121212203-0111001020222200-1332313201300222-3001000002033020-0011230321132211-1201033112130331-1100132320030110)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](resources--cloud_connect--reference--group-001.md#canonical-3200300212101213-0322103321201312-3031332200301200-0130131130213120-0121121301223133-1102023111022131-1132102113110112-0030232031110101)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route](resources--cloud_connect--reference--group-001.md#canonical-3233310033321120-2110120300102121-1222212213121011-3211203323031022-0323101010102011-3032033101120022-1321001101212331-1323210021310313)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.all_route_tables

<a id="canonical-2002021300202131-2112021313220322-3322003130031333-0103211000103333-2023112230010032-3132021003031033-3022130322302132-0023101233202110"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all route tables.

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
all_route_tables = {}
```

<a id="canonical-2333302220321121-3213300112113113-2303012103102300-3323233231300321-3303331201311213-1313211213201133-0332022100112133-1323201030211130"></a>

## Direct properties — all_route_tables / 012103120120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0000011020233230-0022012311223030-2332000130112012-2132100133131322-3002012301232320-3310221330110300-1123132321203212-3020000300300023"></a>

## Next pages — all_route_tables / 012103120120 / 4

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route](resources--cloud_connect--reference--group-001.md#canonical-3233310033321120-2110120300102121-1222212213121011-3211203323031022-0323101010102011-3032033101120022-1321001101212331-1323210021310313)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)

<a id="canonical-2221223103020031-3233303133133123-1311301323021132-0222313210010320-0121213023132100-2030020101221211-3222121030210213-1112030230300030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032202203120000-0003320310033102-1131120030310132-0203113210210022-1210111112202100-3021110331333303-1222213311123103-1023002311021112"></a>

## aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables — selective_route_tables / 123103101102 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-3000033123020112-3330311232230011-1111030120033023-0320023332201330-1233113231331330-1101323002122132-1323101202220113-3023111000220221)
- [aws_provider](resources--cloud_connect--reference--group-001.md#canonical-3103101220221300-3222333301231021-3123032111132110-0202212222332110-2002030232032101-0022221112102322-2300121023130210-2030003033012003)
- [aws_provider.aws_tgw_site](resources--cloud_connect--reference--group-001.md#canonical-1100320102231322-1033313001221030-2113103123301132-3232000101000312-1011301322100112-3011313110020121-3031331221113311-2313211203222203)
- [aws_provider.aws_tgw_site.vpc_attachments](resources--cloud_connect--reference--group-001.md#canonical-2012322120311013-2100022121212203-0111001020222200-1332313201300222-3001000002033020-0011230321132211-1201033112130331-1100132320030110)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](resources--cloud_connect--reference--group-001.md#canonical-3200300212101213-0322103321201312-3031332200301200-0130131130213120-0121121301223133-1102023111022131-1132102113110112-0030232031110101)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route](resources--cloud_connect--reference--group-001.md#canonical-3233310033321120-2110120300102121-1222212213121011-3211203323031022-0323101010102011-3032033101120022-1321001101212331-1323210021310313)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route.selective_route_tables

<a id="canonical-3032220302331311-3303311233201011-2230131321302222-0120211113220323-3102300221131030-1030021113222211-0002300222130131-0103120330210220"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for selective route tables.

Upstream description:

AWS Route Table.

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
selective_route_tables {
  # Configure direct properties listed below.
}
```

<a id="canonical-3113111100001322-2133311232202103-3233133132302003-1311121303303020-1122011131023001-0303313102200010-0311122303211110-2101133331110230"></a>

## Direct properties — selective_route_tables / 123103101102 / 3

<a id="canonical-3011003220122222-0312332121020130-3131330133220310-2103322003103301-2122212001030233-2020032200313032-3230210213132213-1022231010023000"></a>

<a id="canonical-3002100002200232-1132210301100332-3211111232102033-1223313203232031-3210113302033133-1302213021201130-1310301002220201-3302320202101131"></a>

## route_table_id property — selective_route_tables / 123103101102 / 4

Type: `["list", "string"]`. Optional.

Route table ID. Route table ID.

Upstream description:

Route table ID.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.pattern": "^(rtb-)([a-z0-9]{8}|[a-z0-9]{17})$",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.pattern": "^(rtb-)([a-z0-9]{8}|[a-z0-9]{17})$",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2320012220112132-2002320322300023-2032033012313132-0002320120202102-1323121320010313-3223212210213001-1133203303032103-2121132233103312"></a>

## Next pages — selective_route_tables / 123103101102 / 5

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list.default_route](resources--cloud_connect--reference--group-001.md#canonical-3233310033321120-2110120300102121-1222212213121011-3211203323031022-0323101010102011-3032033101120022-1321001101212331-1323210021310313)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)

<a id="canonical-0310333211010221-3003201212210333-1202231323203330-3001221120111001-0120030123123332-1120302002132233-2123230323133300-3122312331330310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223030302112133-2123031300221120-1313103133323322-0311320020112020-0100022022303000-2203310000330332-3130303313012011-0013003013221331"></a>

## aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels — labels / 100113011020 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-3000033123020112-3330311232230011-1111030120033023-0320023332201330-1233113231331330-1101323002122132-1323101202220113-3023111000220221)
- [aws_provider](resources--cloud_connect--reference--group-001.md#canonical-3103101220221300-3222333301231021-3123032111132110-0202212222332110-2002030232032101-0022221112102322-2300121023130210-2030003033012003)
- [aws_provider.aws_tgw_site](resources--cloud_connect--reference--group-001.md#canonical-1100320102231322-1033313001221030-2113103123301132-3232000101000312-1011301322100112-3011313110020121-3031331221113311-2313211203222203)
- [aws_provider.aws_tgw_site.vpc_attachments](resources--cloud_connect--reference--group-001.md#canonical-2012322120311013-2100022121212203-0111001020222200-1332313201300222-3001000002033020-0011230321132211-1201033112130331-1100132320030110)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](resources--cloud_connect--reference--group-001.md#canonical-3200300212101213-0322103321201312-3031332200301200-0130131130213120-0121121301223133-1102023111022131-1132102113110112-0030232031110101)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.labels

<a id="canonical-2112111113022322-0311101112001213-1313111013031302-2202122330323232-3232023130203230-2223110203103312-3021322210223000-0211110013213233"></a>

Type: `"object"`. single nested block, Optional.

Add labels for the VPC attachment. These labels can then be used in policies such as enhanced
firewall.

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
labels {}
```

<a id="canonical-3100032013010103-0111030321212333-2221003032033210-2020333331312300-3223320302301021-3002001003023311-3203103021030113-0000113003033330"></a>

## Direct properties — labels / 100113011020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0321301032230123-3010120002031133-0212303022032010-3031323212031303-2001322022330332-2331103133210032-3333330011322021-2130132223102023"></a>

## Next pages — labels / 100113011020 / 4

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](resources--cloud_connect--reference--group-001.md#canonical-3200300212101213-0322103321201312-3031332200301200-0130131130213120-0121121301223133-1102023111022131-1132102113110112-0030232031110101)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)

<a id="canonical-1020113333302101-2030202001320221-3010320333121102-2301330000232212-0323202310103020-2220221103200131-1112302001210021-2320020130203231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302201103131010-0130332301121030-0310001222133002-3133311101231322-3110220032301333-3233231100210320-3212210131302303-3122220200321221"></a>

## aws_provider.aws_tgw_site.vpc_attachments.vpc_list.manual_routing — manual_routing / 002113013111 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-3000033123020112-3330311232230011-1111030120033023-0320023332201330-1233113231331330-1101323002122132-1323101202220113-3023111000220221)
- [aws_provider](resources--cloud_connect--reference--group-001.md#canonical-3103101220221300-3222333301231021-3123032111132110-0202212222332110-2002030232032101-0022221112102322-2300121023130210-2030003033012003)
- [aws_provider.aws_tgw_site](resources--cloud_connect--reference--group-001.md#canonical-1100320102231322-1033313001221030-2113103123301132-3232000101000312-1011301322100112-3011313110020121-3031331221113311-2313211203222203)
- [aws_provider.aws_tgw_site.vpc_attachments](resources--cloud_connect--reference--group-001.md#canonical-2012322120311013-2100022121212203-0111001020222200-1332313201300222-3001000002033020-0011230321132211-1201033112130331-1100132320030110)
- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](resources--cloud_connect--reference--group-001.md#canonical-3200300212101213-0322103321201312-3031332200301200-0130131130213120-0121121301223133-1102023111022131-1132102113110112-0030232031110101)
- aws_provider.aws_tgw_site.vpc_attachments.vpc_list.manual_routing

<a id="canonical-0221212230300231-1222201233032320-2330000100103113-3311211031110223-0110101132112021-2011023312030113-2321031022220133-0030330302333201"></a>

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
manual_routing = {}
```

<a id="canonical-1123231233202203-3033100103210211-0111100302212000-0231003031133331-2011321223033113-3211012202121023-2123031130122331-1312310003131302"></a>

## Direct properties — manual_routing / 002113013111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1023303013333231-1200311233020113-2223102021320230-3231002220032111-0021322013030123-3002212001321301-3033121103010200-2133323320211120"></a>

## Next pages — manual_routing / 002113013111 / 4

- [aws_provider.aws_tgw_site.vpc_attachments.vpc_list](resources--cloud_connect--reference--group-001.md#canonical-3200300212101213-0322103321201312-3031332200301200-0130131130213120-0121121301223133-1102023111022131-1132102113110112-0030232031110101)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)

<a id="canonical-3021103322130320-1200003010013121-0010020300322100-0312322031223102-3220323002313103-2002300323233103-0021222100212022-3300231133311003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011303331331322-2131220333331102-1132313301002000-0100321022010113-3323303030113023-1221102022223131-0331201101002111-3030211310033200"></a>

## azure_vnet_site — azure_vnet_site / 011320311212 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-3000033123020112-3330311232230011-1111030120033023-0320023332201330-1233113231331330-1101323002122132-1323101202220113-3023111000220221)
- azure_vnet_site

<a id="canonical-3320232211002023-2222321132133113-1030211310323103-0001310332330003-3301010333202123-3021122230233322-1321323003011310-3220330000221212"></a>

Type: `"object"`. single nested block, Optional.

Azure VNet Site Type. Cloud Connect Azure VNet Site Type.

Upstream description:

Cloud Connect Azure VNet Site Type.

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
azure_vnet_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-2000201213231220-1321332132003231-3133133211311312-3211101111232312-0113030003331032-1123121320333132-0113301311001023-2120302120220313"></a>

## Direct properties — azure_vnet_site / 011320311212 / 3

- [site](resources--cloud_connect--reference--group-001.md#canonical-1210100313321313-2113000230032213-2231210032203023-0132313332201110-1332101111332330-2131122210112033-1230210322321031-0211002210020212): complete subsection reference.

- [vnet_attachments](resources--cloud_connect--reference--group-001.md#canonical-3022033033322013-0223203001230023-1301122301311130-3323100132012112-1213320010021300-3231311210113021-1200002123210200-2332333310000100): complete subsection reference.

<a id="canonical-3300303332202112-2123313031220023-1310121023003210-3133131121212202-0010323032331231-2001020022113223-1030211201110101-3000001331031313"></a>

## Next pages — azure_vnet_site / 011320311212 / 4

- [azure_vnet_site.site](resources--cloud_connect--reference--group-001.md#canonical-1210100313321313-2113000230032213-2231210032203023-0132313332201110-1332101111332330-2131122210112033-1230210322321031-0211002210020212)
- [azure_vnet_site.vnet_attachments](resources--cloud_connect--reference--group-001.md#canonical-3022033033322013-0223203001230023-1301122301311130-3323100132012112-1213320010021300-3231311210113021-1200002123210200-2332333310000100)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-3000033123020112-3330311232230011-1111030120033023-0320023332201330-1233113231331330-1101323002122132-1323101202220113-3023111000220221)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)

<a id="canonical-1210100313321313-2113000230032213-2231210032203023-0132313332201110-1332101111332330-2131122210112033-1230210322321031-0211002210020212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222120220120210-2332131213232011-3010011132001102-2113332303312100-0200233113301020-0013012003320033-3220202010223311-1202020123033212"></a>

## azure_vnet_site.site — site / 313001020231 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-3000033123020112-3330311232230011-1111030120033023-0320023332201330-1233113231331330-1101323002122132-1323101202220113-3023111000220221)
- [azure_vnet_site](resources--cloud_connect--reference--group-001.md#canonical-3021103322130320-1200003010013121-0010020300322100-0312322031223102-3220323002313103-2002300323233103-0021222100212022-3300231133311003)
- azure_vnet_site.site

<a id="canonical-0222323330102223-3300031002033201-3222331010003211-1011202111110320-0002310001031112-0331232322203322-3003030302322130-0032020032303000"></a>

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-3012102332222331-2331030020230230-3012113113122312-0113233311200001-0230303102313323-2232322320103022-3223210221010020-2120202332202130"></a>

## Direct properties — site / 313001020231 / 3

<a id="canonical-2333332223122210-1233223031210130-3133210220320012-2001001100033321-2033021123201320-3203333300201023-3333130201113132-3122221112332030"></a>

<a id="canonical-0003301102133323-0322130032222022-3331133311332120-1022330031232022-2123012203133020-3320130131021121-2121311003232100-0012011202113010"></a>

## name property — site / 313001020231 / 4

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

<a id="canonical-0202313331012320-0012132010101002-3212012030220100-2322122320032300-1211032311002010-0120002013230331-1330023012103203-1213013303320031"></a>

<a id="canonical-3303223223022312-2330312321030212-0112112010010033-3211200022011030-1201123033133001-1331200102121110-2032330301210010-0010201002223221"></a>

## namespace property — site / 313001020231 / 5

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

<a id="canonical-3113013123002231-2231203011120332-1033302122122230-0133030301022301-3121101103003130-0210232013332003-1032303211133131-0113113012120122"></a>

<a id="canonical-0010323122311312-0223023122122033-0012011110203213-1300332210110233-1100101122322020-3033121312110223-3100121302030033-3100321111130001"></a>

## tenant property — site / 313001020231 / 6

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

<a id="canonical-0322320103301102-1212230313333132-3233312310222332-3000023312021212-1020321120333103-1203010320000132-0101331213222313-3123031101110100"></a>

## Next pages — site / 313001020231 / 7

- [azure_vnet_site](resources--cloud_connect--reference--group-001.md#canonical-3021103322130320-1200003010013121-0010020300322100-0312322031223102-3220323002313103-2002300323233103-0021222100212022-3300231133311003)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)

<a id="canonical-3022033033322013-0223203001230023-1301122301311130-3323100132012112-1213320010021300-3231311210113021-1200002123210200-2332333310000100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122233311122132-0110020122100001-1300010132002010-2302033332333103-3202301333122101-1232101233110300-1223120012103122-3202211122120201"></a>

## azure_vnet_site.vnet_attachments — vnet_attachments / 212222133211 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-3000033123020112-3330311232230011-1111030120033023-0320023332201330-1233113231331330-1101323002122132-1323101202220113-3023111000220221)
- [azure_vnet_site](resources--cloud_connect--reference--group-001.md#canonical-3021103322130320-1200003010013121-0010020300322100-0312322031223102-3220323002313103-2002300323233103-0021222100212022-3300231133311003)
- azure_vnet_site.vnet_attachments

<a id="canonical-3003210202311111-1003301310211211-3321021232232313-3202221201331320-1123301133303203-2030223311022303-1230111123313213-1321103330110032"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for vnet attachments.

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
vnet_attachments {
  # Configure direct properties listed below.
}
```

<a id="canonical-0122220320031221-0130310100000210-3220130021332210-2120321312231300-2030331032000031-3102123021301232-2313310200211133-1130222301100312"></a>

## Direct properties — vnet_attachments / 212222133211 / 3

- [vnet_list](resources--cloud_connect--reference--group-001.md#canonical-2103202202212310-0301333133333230-1313222101031310-2110032331113122-1302101123011231-2220302331003321-1211030020313312-2313300020200222): complete subsection reference.

<a id="canonical-0032102023333000-0231111202311111-3100213120133303-2330333313332102-3123030112032110-3033103232121103-3310123302011332-2032121230233122"></a>

## Next pages — vnet_attachments / 212222133211 / 4

- [azure_vnet_site.vnet_attachments.vnet_list](resources--cloud_connect--reference--group-001.md#canonical-2103202202212310-0301333133333230-1313222101031310-2110032331113122-1302101123011231-2220302331003321-1211030020313312-2313300020200222)
- [azure_vnet_site](resources--cloud_connect--reference--group-001.md#canonical-3021103322130320-1200003010013121-0010020300322100-0312322031223102-3220323002313103-2002300323233103-0021222100212022-3300231133311003)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)

<a id="canonical-2103202202212310-0301333133333230-1313222101031310-2110032331113122-1302101123011231-2220302331003321-1211030020313312-2313300020200222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111101030013323-1230002231203333-2330333223100121-0000233121300030-1122010131302031-3200211023330221-2322111023013223-2300030102111130"></a>

## azure_vnet_site.vnet_attachments.vnet_list — vnet_list / 232311003321 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-3000033123020112-3330311232230011-1111030120033023-0320023332201330-1233113231331330-1101323002122132-1323101202220113-3023111000220221)
- [azure_vnet_site](resources--cloud_connect--reference--group-001.md#canonical-3021103322130320-1200003010013121-0010020300322100-0312322031223102-3220323002313103-2002300323233103-0021222100212022-3300231133311003)
- [azure_vnet_site.vnet_attachments](resources--cloud_connect--reference--group-001.md#canonical-3022033033322013-0223203001230023-1301122301311130-3323100132012112-1213320010021300-3231311210113021-1200002123210200-2332333310000100)
- azure_vnet_site.vnet_attachments.vnet_list

<a id="canonical-0303021332031221-2323220033203023-0332320121031300-0132013331103203-2120320022301001-3233120021131122-3032311032333121-1312003230233303"></a>

Type: `"object"`. list nested block, Optional.

VNet List. Collection of items or values

Upstream description:

Collection of items or values

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("subscription_id",
    "vnet_id"),
  validators.ConflictingListObjectAttributes("custom_routing",
    "default_route"),
  validators.ConflictingListObjectAttributes("custom_routing",
    "manual_routing"),
  validators.ConflictingListObjectAttributes("default_route",
    "manual_routing")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

Terraform syntax:

```terraform
vnet_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1021003112133120-2023123030103121-0022301333233313-3120111103011320-2122023012103212-0130101123211223-3230023223111232-1223132112321122"></a>

## Direct properties — vnet_list / 232311003321 / 3

- [custom_routing](resources--cloud_connect--reference--group-001.md#canonical-2200110000212110-0211223122101311-0132320111230123-3202303313123111-3032112321013112-3003302001202121-0223200133101303-3220101131112012): complete subsection reference.

- [default_route](resources--cloud_connect--reference--group-001.md#canonical-3220201320311203-2020313030312302-1330213222131120-3312022320300031-0211232111213303-1132323111103322-2121000101222302-0211103211201013): complete subsection reference.

- [labels](resources--cloud_connect--reference--group-001.md#canonical-2202230122213313-2312032221021231-1230313333300022-2000331210312310-1002233121022033-2213301331013233-3003120110013113-2112131221230020): complete subsection reference.

- [manual_routing](resources--cloud_connect--reference--group-001.md#canonical-3210032132201123-2302223311020303-2330130200133121-0133200323321032-0002012102122332-0110321202210120-1220322230013030-3113020302121222): complete subsection reference.

<a id="canonical-2332303101311210-2112113202111210-3001220123212023-1323320212122101-1030211121122311-1023223331213300-0132032233002130-0221122022321323"></a>

<a id="canonical-2230301121111113-3330130110022213-3312103310113220-3101330011321330-1122222010311020-3313132321320331-3102123112233023-1100010120231130"></a>

## subscription_id property — vnet_list / 232311003321 / 4

Type: `"string"`. Optional.

Enter the Subscription ID of the VNet to be attached.

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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-3123112001323223-1213033320320233-1320300321322200-2221210012133303-0333233002003101-3002031313210203-3111331312233211-0303300233013121"></a>

<a id="canonical-3310023030101301-3103212210233320-0231322302213013-1102121030131123-3022112303122030-2210132121322202-0322230300133203-2330123123123020"></a>

## vnet_id property — vnet_list / 232311003321 / 5

Type: `"string"`. Optional.

Enter the VNet ID of the VNet to be attached in format
/&lt;resource-group-name&gt;/&lt;VNet-name&gt;.

Upstream description:

Enter the VNet ID of the VNet to be attached in format
/&lt;resource-group-name&gt;/&lt;VNet-name&gt;

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2210211223100022-0023310132132030-0331321023122113-2321200010103200-1123331113232032-2201322003103333-2003132033031210-1301031002110222"></a>

## Next pages — vnet_list / 232311003321 / 6

- [azure_vnet_site.vnet_attachments.vnet_list.custom_routing](resources--cloud_connect--reference--group-001.md#canonical-2200110000212110-0211223122101311-0132320111230123-3202303313123111-3032112321013112-3003302001202121-0223200133101303-3220101131112012)
- [azure_vnet_site.vnet_attachments.vnet_list.default_route](resources--cloud_connect--reference--group-001.md#canonical-3220201320311203-2020313030312302-1330213222131120-3312022320300031-0211232111213303-1132323111103322-2121000101222302-0211103211201013)
- [azure_vnet_site.vnet_attachments.vnet_list.labels](resources--cloud_connect--reference--group-001.md#canonical-2202230122213313-2312032221021231-1230313333300022-2000331210312310-1002233121022033-2213301331013233-3003120110013113-2112131221230020)
- [azure_vnet_site.vnet_attachments.vnet_list.manual_routing](resources--cloud_connect--reference--group-001.md#canonical-3210032132201123-2302223311020303-2330130200133121-0133200323321032-0002012102122332-0110321202210120-1220322230013030-3113020302121222)
- [azure_vnet_site.vnet_attachments](resources--cloud_connect--reference--group-001.md#canonical-3022033033322013-0223203001230023-1301122301311130-3323100132012112-1213320010021300-3231311210113021-1200002123210200-2332333310000100)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)

<a id="canonical-2200110000212110-0211223122101311-0132320111230123-3202303313123111-3032112321013112-3003302001202121-0223200133101303-3220101131112012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0113033212331001-0112113032021121-2021310120320200-0331201313130030-1223111102113022-0202100220230111-2100210020131233-1123220210230202"></a>

## azure_vnet_site.vnet_attachments.vnet_list.custom_routing — custom_routing / 211012200313 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-3000033123020112-3330311232230011-1111030120033023-0320023332201330-1233113231331330-1101323002122132-1323101202220113-3023111000220221)
- [azure_vnet_site](resources--cloud_connect--reference--group-001.md#canonical-3021103322130320-1200003010013121-0010020300322100-0312322031223102-3220323002313103-2002300323233103-0021222100212022-3300231133311003)
- [azure_vnet_site.vnet_attachments](resources--cloud_connect--reference--group-001.md#canonical-3022033033322013-0223203001230023-1301122301311130-3323100132012112-1213320010021300-3231311210113021-1200002123210200-2332333310000100)
- [azure_vnet_site.vnet_attachments.vnet_list](resources--cloud_connect--reference--group-001.md#canonical-2103202202212310-0301333133333230-1313222101031310-2110032331113122-1302101123011231-2220302331003321-1211030020313312-2313300020200222)
- azure_vnet_site.vnet_attachments.vnet_list.custom_routing

<a id="canonical-2211030201023203-2000132221231211-3301200320032323-0112032110022211-0111111130001103-3122330222120220-1103102001312232-0030221300210331"></a>

Type: `"object"`. single nested block, Optional.

List Azure Route Table with Static Route.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("route_tables")}
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
custom_routing {
  # Configure direct properties listed below.
}
```

<a id="canonical-1223212101321322-3301100003330332-0301331222131023-3230002010001322-0330013333230231-2102121311111231-0302100203322221-0103212122300013"></a>

## Direct properties — custom_routing / 211012200313 / 3

- [route_tables](resources--cloud_connect--reference--group-001.md#canonical-1333020222202213-0312323112021211-3232130011213302-1223001300200100-2221133021022132-0200212002203133-2221033221323120-3011313330102330): complete subsection reference.

<a id="canonical-0323302123301120-2231300010122011-3002022123322203-3222200013300132-3100333121022031-2310200321030213-3322331000000122-1013113323103123"></a>

## Next pages — custom_routing / 211012200313 / 4

- [azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables](resources--cloud_connect--reference--group-001.md#canonical-1333020222202213-0312323112021211-3232130011213302-1223001300200100-2221133021022132-0200212002203133-2221033221323120-3011313330102330)
- [azure_vnet_site.vnet_attachments.vnet_list](resources--cloud_connect--reference--group-001.md#canonical-2103202202212310-0301333133333230-1313222101031310-2110032331113122-1302101123011231-2220302331003321-1211030020313312-2313300020200222)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)

<a id="canonical-1333020222202213-0312323112021211-3232130011213302-1223001300200100-2221133021022132-0200212002203133-2221033221323120-3011313330102330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310230220302333-2111200212002002-2112221303230300-0210303201302301-3113322210232013-1330322220330212-1322011201010310-2122211302222010"></a>

## azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables — route_tables / 321102111200 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-3000033123020112-3330311232230011-1111030120033023-0320023332201330-1233113231331330-1101323002122132-1323101202220113-3023111000220221)
- [azure_vnet_site](resources--cloud_connect--reference--group-001.md#canonical-3021103322130320-1200003010013121-0010020300322100-0312322031223102-3220323002313103-2002300323233103-0021222100212022-3300231133311003)
- [azure_vnet_site.vnet_attachments](resources--cloud_connect--reference--group-001.md#canonical-3022033033322013-0223203001230023-1301122301311130-3323100132012112-1213320010021300-3231311210113021-1200002123210200-2332333310000100)
- [azure_vnet_site.vnet_attachments.vnet_list](resources--cloud_connect--reference--group-001.md#canonical-2103202202212310-0301333133333230-1313222101031310-2110032331113122-1302101123011231-2220302331003321-1211030020313312-2313300020200222)
- [azure_vnet_site.vnet_attachments.vnet_list.custom_routing](resources--cloud_connect--reference--group-001.md#canonical-2200110000212110-0211223122101311-0132320111230123-3202303313123111-3032112321013112-3003302001202121-0223200133101303-3220101131112012)
- azure_vnet_site.vnet_attachments.vnet_list.custom_routing.route_tables

<a id="canonical-2121301200130110-2000111222231002-1111131003312123-1021003331030230-0021313223331302-1023202301021222-1202231200233021-0013113113033310"></a>

Type: `"object"`. list nested block, Optional.

List of route tables with static routes. Route Tables with static routes.

Upstream description:

Route Tables with static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("static_routes")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 200,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "200",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "200",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
route_tables {
  # Configure direct properties listed below.
}
```

<a id="canonical-0210121132213031-0320120212313213-2030200013320132-3012331113310311-0222001202231202-0101212320332111-1202232102010112-0231303123201332"></a>

## Direct properties — route_tables / 321102111200 / 3

<a id="canonical-3211033203120010-2122301223120122-1120311332211213-2303330210011102-2102221300021110-2112011320320021-1021012322311010-1033322100022131"></a>

<a id="canonical-2321320333213321-1101130103132201-3203222210123332-0020301303333231-1223000100231133-0230030123110123-2203023121321333-3130111331332311"></a>

## route_table_id property — route_tables / 321102111200 / 4

Type: `"string"`. Optional.

Route table ID in the format /&lt;resource-group-name&gt;/&lt;route-table-name&gt;.

Upstream description:

Route table ID in the format /&lt;resource-group-name&gt;/&lt;route-table-name&gt;

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^\\\\/[-\\\\w\\\\._\\\\(\\\\)]+\\\\/[a-zA-Z0-9][a-zA-Z0-9-._]+[a-zA-Z0-9_]$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^\\\\/[-\\\\w\\\\._\\\\(\\\\)]+\\\\/[a-zA-Z0-9][a-zA-Z0-9-._]+[a-zA-Z0-9_]$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.pattern": "^\\\\/[-\\\\w\\\\._\\\\(\\\\)]+\\\\/[a-zA-Z0-9][a-zA-Z0-9-._]+[a-zA-Z0-9_]$"
  }
}
```

<a id="canonical-3123020300030303-2200120211111312-2113023322013223-2132001010210002-0223300110321232-3313111011023320-1103133033103113-2203320100320113"></a>

<a id="canonical-0232013232200102-2102233223110212-3012322021303333-0330100023103122-0211012231011020-2133202312112002-3122231000333313-1122000101123310"></a>

## static_routes property — route_tables / 321102111200 / 5

Type: `["list", "string"]`. Optional.

Static Routes. List of Static Routes.

Upstream description:

List of Static Routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 50),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 50,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 50,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3113200002212121-3122232210112312-2303211133133111-2122310003313030-0231003023032320-1001210031112020-3210023011200020-3320003331110313"></a>

## Next pages — route_tables / 321102111200 / 6

- [azure_vnet_site.vnet_attachments.vnet_list.custom_routing](resources--cloud_connect--reference--group-001.md#canonical-2200110000212110-0211223122101311-0132320111230123-3202303313123111-3032112321013112-3003302001202121-0223200133101303-3220101131112012)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)

<a id="canonical-3220201320311203-2020313030312302-1330213222131120-3312022320300031-0211232111213303-1132323111103322-2121000101222302-0211103211201013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211220230313120-2000302132232032-2312221000202011-0313123010203203-3332230320330121-1032022232230010-2020121303323131-1201010132031210"></a>

## azure_vnet_site.vnet_attachments.vnet_list.default_route — default_route / 313013310210 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-3000033123020112-3330311232230011-1111030120033023-0320023332201330-1233113231331330-1101323002122132-1323101202220113-3023111000220221)
- [azure_vnet_site](resources--cloud_connect--reference--group-001.md#canonical-3021103322130320-1200003010013121-0010020300322100-0312322031223102-3220323002313103-2002300323233103-0021222100212022-3300231133311003)
- [azure_vnet_site.vnet_attachments](resources--cloud_connect--reference--group-001.md#canonical-3022033033322013-0223203001230023-1301122301311130-3323100132012112-1213320010021300-3231311210113021-1200002123210200-2332333310000100)
- [azure_vnet_site.vnet_attachments.vnet_list](resources--cloud_connect--reference--group-001.md#canonical-2103202202212310-0301333133333230-1313222101031310-2110032331113122-1302101123011231-2220302331003321-1211030020313312-2313300020200222)
- azure_vnet_site.vnet_attachments.vnet_list.default_route

<a id="canonical-3120203333233030-1032230201302211-2311333313330321-2132113232120001-3321213122023021-1220201212000303-2213120322011202-0001320102203303"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for default route.

Upstream description:

Select Override Default Route Choice.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("all_route_tables",
    "selective_route_tables")}
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
  "x-ves-oneof-field-default_route_choice": "[\"all_route_tables\",\"selective_route_tables\"]"
}
```

Terraform syntax:

```terraform
default_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-3133113313212202-1232111322320210-2213231130110222-2223111013333320-0220032000000130-2101020310120121-2100010210002112-2331102201013301"></a>

## Direct properties — default_route / 313013310210 / 3

- [all_route_tables](resources--cloud_connect--reference--group-001.md#canonical-2321310022221102-0211100223023023-2101002122232231-0223133232223133-0122330111013112-0122120022030302-3232322322110001-3131022211321301): complete subsection reference.

- [selective_route_tables](resources--cloud_connect--reference--group-001.md#canonical-1330323010211330-0223132112002123-1311312313131103-1003013021230300-0031320131121332-2330301012303023-1313113122010021-1231303321321130): complete subsection reference.

<a id="canonical-0230120101231210-0111212132213313-2132301332221230-0310011302310013-0001113233022011-3021031133130103-0303332202202331-3200322013122010"></a>

## Next pages — default_route / 313013310210 / 4

- [azure_vnet_site.vnet_attachments.vnet_list.default_route.all_route_tables](resources--cloud_connect--reference--group-001.md#canonical-2321310022221102-0211100223023023-2101002122232231-0223133232223133-0122330111013112-0122120022030302-3232322322110001-3131022211321301)
- [azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables](resources--cloud_connect--reference--group-001.md#canonical-1330323010211330-0223132112002123-1311312313131103-1003013021230300-0031320131121332-2330301012303023-1313113122010021-1231303321321130)
- [azure_vnet_site.vnet_attachments.vnet_list](resources--cloud_connect--reference--group-001.md#canonical-2103202202212310-0301333133333230-1313222101031310-2110032331113122-1302101123011231-2220302331003321-1211030020313312-2313300020200222)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)

<a id="canonical-2321310022221102-0211100223023023-2101002122232231-0223133232223133-0122330111013112-0122120022030302-3232322322110001-3131022211321301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030313303211023-0001320030230132-2110130010233103-1102020223113312-3212102110131033-0303100223113012-3212032113230301-2020122321002312"></a>

## azure_vnet_site.vnet_attachments.vnet_list.default_route.all_route_tables — all_route_tables / 211201202023 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-3000033123020112-3330311232230011-1111030120033023-0320023332201330-1233113231331330-1101323002122132-1323101202220113-3023111000220221)
- [azure_vnet_site](resources--cloud_connect--reference--group-001.md#canonical-3021103322130320-1200003010013121-0010020300322100-0312322031223102-3220323002313103-2002300323233103-0021222100212022-3300231133311003)
- [azure_vnet_site.vnet_attachments](resources--cloud_connect--reference--group-001.md#canonical-3022033033322013-0223203001230023-1301122301311130-3323100132012112-1213320010021300-3231311210113021-1200002123210200-2332333310000100)
- [azure_vnet_site.vnet_attachments.vnet_list](resources--cloud_connect--reference--group-001.md#canonical-2103202202212310-0301333133333230-1313222101031310-2110032331113122-1302101123011231-2220302331003321-1211030020313312-2313300020200222)
- [azure_vnet_site.vnet_attachments.vnet_list.default_route](resources--cloud_connect--reference--group-001.md#canonical-3220201320311203-2020313030312302-1330213222131120-3312022320300031-0211232111213303-1132323111103322-2121000101222302-0211103211201013)
- azure_vnet_site.vnet_attachments.vnet_list.default_route.all_route_tables

<a id="canonical-3311201022122013-1332103001231331-2223133032113103-2313122101113000-0213111300120113-3000030132123110-2110321231012333-2100331231202333"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all route tables.

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
all_route_tables = {}
```

<a id="canonical-3330233203201133-3133333011213202-0012310333012201-0333003010312313-3002221313121103-2321031023101002-2302122001131133-2330200230333103"></a>

## Direct properties — all_route_tables / 211201202023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2132221313101332-2130033311121312-1323110222023321-0320223120320331-3222021023033302-1331101022212123-1032332323100303-0220122212323030"></a>

## Next pages — all_route_tables / 211201202023 / 4

- [azure_vnet_site.vnet_attachments.vnet_list.default_route](resources--cloud_connect--reference--group-001.md#canonical-3220201320311203-2020313030312302-1330213222131120-3312022320300031-0211232111213303-1132323111103322-2121000101222302-0211103211201013)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)

<a id="canonical-1330323010211330-0223132112002123-1311312313131103-1003013021230300-0031320131121332-2330301012303023-1313113122010021-1231303321321130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0233313133320330-2001323312120321-0210133331020200-3302303111210210-2013320321302331-2013033300303001-3032002331320302-1213020032320013"></a>

## azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables — selective_route_tables / 223113322102 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-3000033123020112-3330311232230011-1111030120033023-0320023332201330-1233113231331330-1101323002122132-1323101202220113-3023111000220221)
- [azure_vnet_site](resources--cloud_connect--reference--group-001.md#canonical-3021103322130320-1200003010013121-0010020300322100-0312322031223102-3220323002313103-2002300323233103-0021222100212022-3300231133311003)
- [azure_vnet_site.vnet_attachments](resources--cloud_connect--reference--group-001.md#canonical-3022033033322013-0223203001230023-1301122301311130-3323100132012112-1213320010021300-3231311210113021-1200002123210200-2332333310000100)
- [azure_vnet_site.vnet_attachments.vnet_list](resources--cloud_connect--reference--group-001.md#canonical-2103202202212310-0301333133333230-1313222101031310-2110032331113122-1302101123011231-2220302331003321-1211030020313312-2313300020200222)
- [azure_vnet_site.vnet_attachments.vnet_list.default_route](resources--cloud_connect--reference--group-001.md#canonical-3220201320311203-2020313030312302-1330213222131120-3312022320300031-0211232111213303-1132323111103322-2121000101222302-0211103211201013)
- azure_vnet_site.vnet_attachments.vnet_list.default_route.selective_route_tables

<a id="canonical-0130213311112211-2222111120002031-0132310120000223-3221210312131033-2213123202303213-3120133202222122-3323230310122231-1033313132331302"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for selective route tables.

Upstream description:

Azure Route Table.

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
selective_route_tables {
  # Configure direct properties listed below.
}
```

<a id="canonical-1023103211210213-1023233011033121-2332321201313103-2232030320130210-2130202100233031-1330130332122320-3222312323030322-2120132201310100"></a>

## Direct properties — selective_route_tables / 223113322102 / 3

<a id="canonical-0330021123033013-3230222122303330-2122013312010323-2102110023213330-1131221133133210-2102320012231202-1102231110013003-0332121121301121"></a>

<a id="canonical-2232222010222123-3311121221101213-3302230220322330-2123321130321200-3120012013232122-1003213203321033-2112201113220001-2302311010210313"></a>

## route_table_id property — selective_route_tables / 223113322102 / 4

Type: `["list", "string"]`. Optional.

Route table ID in the format /&lt;resource-group-name&gt;/&lt;route-table-name&gt;.

Upstream description:

Route table ID in the format /&lt;resource-group-name&gt;/&lt;route-table-name&gt;

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3320123113330230-3013003031022331-3022201221023221-3310333320202030-2221222022311310-3100201210103221-3202331310221211-3312130023110021"></a>

## Next pages — selective_route_tables / 223113322102 / 5

- [azure_vnet_site.vnet_attachments.vnet_list.default_route](resources--cloud_connect--reference--group-001.md#canonical-3220201320311203-2020313030312302-1330213222131120-3312022320300031-0211232111213303-1132323111103322-2121000101222302-0211103211201013)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)

<a id="canonical-2202230122213313-2312032221021231-1230313333300022-2000331210312310-1002233121022033-2213301331013233-3003120110013113-2112131221230020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3030030122123111-0112030330332032-1000310222022122-3323313311312002-2300001020310202-3003333220111102-1102121002003022-2312331233300200"></a>

## azure_vnet_site.vnet_attachments.vnet_list.labels — labels / 331220201222 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-3000033123020112-3330311232230011-1111030120033023-0320023332201330-1233113231331330-1101323002122132-1323101202220113-3023111000220221)
- [azure_vnet_site](resources--cloud_connect--reference--group-001.md#canonical-3021103322130320-1200003010013121-0010020300322100-0312322031223102-3220323002313103-2002300323233103-0021222100212022-3300231133311003)
- [azure_vnet_site.vnet_attachments](resources--cloud_connect--reference--group-001.md#canonical-3022033033322013-0223203001230023-1301122301311130-3323100132012112-1213320010021300-3231311210113021-1200002123210200-2332333310000100)
- [azure_vnet_site.vnet_attachments.vnet_list](resources--cloud_connect--reference--group-001.md#canonical-2103202202212310-0301333133333230-1313222101031310-2110032331113122-1302101123011231-2220302331003321-1211030020313312-2313300020200222)
- azure_vnet_site.vnet_attachments.vnet_list.labels

<a id="canonical-0222002120233301-0223113320230300-0121013133001330-1132203010210000-3133322112020100-0322003131200121-3313113331130221-0232331321121300"></a>

Type: `"object"`. single nested block, Optional.

Add labels for the VNet attachments. These labels can then be used in policies such as enhanced
firewall policies.

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
labels {}
```

<a id="canonical-2202202023001233-2013003303332100-0333210211222031-2220001212010232-2221121010111132-3311113112121313-3322021032033001-2302203303031113"></a>

## Direct properties — labels / 331220201222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3011301120003021-1002011020110223-2110010213310103-0330333010010202-2230233023110120-1011010023100112-3120102022201110-3330333010232301"></a>

## Next pages — labels / 331220201222 / 4

- [azure_vnet_site.vnet_attachments.vnet_list](resources--cloud_connect--reference--group-001.md#canonical-2103202202212310-0301333133333230-1313222101031310-2110032331113122-1302101123011231-2220302331003321-1211030020313312-2313300020200222)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)

<a id="canonical-3210032132201123-2302223311020303-2330130200133121-0133200323321032-0002012102122332-0110321202210120-1220322230013030-3113020302121222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220313302022120-2033311222300003-2022212303110012-3120233011111122-3011022113203003-0313023001102013-0333213233132311-2223132212010231"></a>

## azure_vnet_site.vnet_attachments.vnet_list.manual_routing — manual_routing / 103031102302 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-3000033123020112-3330311232230011-1111030120033023-0320023332201330-1233113231331330-1101323002122132-1323101202220113-3023111000220221)
- [azure_vnet_site](resources--cloud_connect--reference--group-001.md#canonical-3021103322130320-1200003010013121-0010020300322100-0312322031223102-3220323002313103-2002300323233103-0021222100212022-3300231133311003)
- [azure_vnet_site.vnet_attachments](resources--cloud_connect--reference--group-001.md#canonical-3022033033322013-0223203001230023-1301122301311130-3323100132012112-1213320010021300-3231311210113021-1200002123210200-2332333310000100)
- [azure_vnet_site.vnet_attachments.vnet_list](resources--cloud_connect--reference--group-001.md#canonical-2103202202212310-0301333133333230-1313222101031310-2110032331113122-1302101123011231-2220302331003321-1211030020313312-2313300020200222)
- azure_vnet_site.vnet_attachments.vnet_list.manual_routing

<a id="canonical-2200130320303311-3232302300231322-3231011333031313-0001021031232211-2100122013303200-3131313012021120-0221222232120103-0311323301100002"></a>

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
manual_routing = {}
```

<a id="canonical-2120220030123211-3031310101120003-3221101302232301-3120300111233121-0101302202231210-0013021301101100-1021130001232303-3102332330330300"></a>

## Direct properties — manual_routing / 103031102302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3220312331212111-2010213213012021-3311330110001012-3120020003302233-1103223000020301-3010131322102013-0311320302102202-1132101310310121"></a>

## Next pages — manual_routing / 103031102302 / 4

- [azure_vnet_site.vnet_attachments.vnet_list](resources--cloud_connect--reference--group-001.md#canonical-2103202202212310-0301333133333230-1313222101031310-2110032331113122-1302101123011231-2220302331003321-1211030020313312-2313300020200222)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)

<a id="canonical-0212201120030202-1202123000003111-2102023221221232-3200021132222310-2120132330002232-1032231132002213-2133103223032230-0321330113002232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212003120011203-0312233013112111-3032103211233102-3132113202201213-3313020031003200-3301233211332011-1022122330203023-0213300001110212"></a>

## segment — segment / 013200020020 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-3000033123020112-3330311232230011-1111030120033023-0320023332201330-1233113231331330-1101323002122132-1323101202220113-3023111000220221)
- segment

<a id="canonical-1221133032112201-2320030323022320-1012301202303203-1031211200300032-0102123323230122-1231200112101123-2323021213021002-2320111120032123"></a>

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
segment {
  # Configure direct properties listed below.
}
```

<a id="canonical-0110330121313131-2210223132230221-2013221310010021-2332120002110130-0111202130210031-1013222112203133-2122033211033021-2021023021131033"></a>

## Direct properties — segment / 013200020020 / 3

<a id="canonical-0000100301313103-1321202332030333-1232321133233130-2112101213321112-1100212103133211-3233121003301013-0312222112200032-1110222302320320"></a>

<a id="canonical-1320202003211033-0200230113011213-3321200010303320-0012312123203123-3333230001112133-0120210332332130-0102003012030212-1223233012321211"></a>

## name property — segment / 013200020020 / 4

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

<a id="canonical-2311201100330300-0312120312220313-2121002301033020-1012330221133201-3332212220010031-0203201111013310-3022000021313222-2121013110013120"></a>

<a id="canonical-1322133011233131-0103323130001300-3210120220301122-3332130200020012-3321330012101032-2213011113303233-1103020122020300-2013131202021022"></a>

## namespace property — segment / 013200020020 / 5

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

<a id="canonical-3112010120000011-0311123022202000-2201213331132313-2310131023201012-0200121332223332-0121110331330310-2211102100002323-2311012210201322"></a>

<a id="canonical-1033102213230002-3213303130101203-0312232033220013-2112332032221020-3330031002102210-2122111232002113-0011222010200213-2123001311230233"></a>

## tenant property — segment / 013200020020 / 6

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

<a id="canonical-0100031022120130-1312203032213300-1011202012202103-0223311203311233-2110112000132011-3331103230120101-1110223210300021-0211020310200313"></a>

## Next pages — segment / 013200020020 / 7

- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-3000033123020112-3330311232230011-1111030120033023-0320023332201330-1233113231331330-1101323002122132-1323101202220113-3023111000220221)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)

<a id="canonical-0102000200203213-0133023212110222-3021012023202203-0022100231003022-1321013002322210-0303322031111330-1031321201022011-1032333113131332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120303300320202-1120222132023332-0130332113323332-3031222013320311-2132021133201033-2102312213003331-3210311001011023-3110001220231002"></a>

## timeouts — timeouts / 121013300131 / 2

Breadcrumbs:

- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)
- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-3000033123020112-3330311232230011-1111030120033023-0320023332201330-1233113231331330-1101323002122132-1323101202220113-3023111000220221)
- timeouts

<a id="canonical-2101123122131111-0123323323022131-2010031211110200-3230221033101310-3030222020021103-2121332120313013-1331202001222003-3111211303211020"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3310323230320023-2022231011023100-1011113301201232-1201301032102100-1220322233223322-0011023023123333-2210332320001122-1323112011230010"></a>

## Direct properties — timeouts / 121013300131 / 3

<a id="canonical-3302322231002222-1000001230320012-2102230021320232-3111033003212112-1223123133301210-0202112013211332-2000021100212333-0113210221100301"></a>

<a id="canonical-3121230212001323-1301202201003121-1130123102200011-2230013222011132-2010020213310121-1332132013320320-0121303220102302-2202031022020233"></a>

## create property — timeouts / 121013300131 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3312330100023333-3111223132200021-0112323220312113-1303210320130003-1300333120300032-3323211033032201-0323122333330330-3220132122101122"></a>

<a id="canonical-2311302120232220-0332022011121213-0200202213020312-3311311222300102-0321211303133100-0032231332003331-0110112031121223-3332122032000213"></a>

## delete property — timeouts / 121013300131 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-3021202320130010-0131121003003122-1012032203031030-2120032221300112-1210322223301222-2110003000233030-1101132312012231-3001322231222323"></a>

<a id="canonical-1103211321113332-1000120303013032-1301333203300211-2331213013130233-2223210111011020-0330020123233333-0231001103322032-3022223133332200"></a>

## read property — timeouts / 121013300131 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2210233302213033-3011121223011033-1230112030202133-3231000131102231-0123130002203033-3313003022330002-3323123333301332-3103133120120211"></a>

<a id="canonical-1022121201202202-1322003130001130-2021223211013033-1013130013032030-2031130300222321-1211131312031122-1032212111113321-0331202010202121"></a>

## update property — timeouts / 121013300131 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1100103100203100-3333303301321010-3300130122310310-3010223310112003-2013221033212121-0023001002020223-0201030111210313-1030001212231100"></a>

## Next pages — timeouts / 121013300131 / 8

- [Property reference](resources--cloud_connect--reference--group-001.md#canonical-3000033123020112-3330311232230011-1111030120033023-0320023332201330-1233113231331330-1101323002122132-1323101202220113-3023111000220221)
- [xcsh_cloud_connect](../resources/cloud_connect.md#canonical-3211233122320300-2011210020323101-0200131013121302-1233133220013202-0122133132120331-2023300223331013-1010032320020312-0333212122300033)
