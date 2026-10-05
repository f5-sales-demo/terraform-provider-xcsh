---
page_title: "xcsh_origin_pool reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_origin_pool reference."
---

# xcsh_origin_pool reference

<a id="canonical-0303320322232021-3230331021221230-3233300121123203-0323131322210333-1211213200013022-3121333212320111-1131212232300133-2132123230303020"></a>

## labels property — origin_servers / 330013001011 / 4

Type: `["map", "string"]`. Optional.

Add Labels for this origin server, these labels can be used to form subset.

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

- [private_ip](resources--origin_pool--reference--group-002.md#canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003): complete subsection reference.

- [private_name](resources--origin_pool--reference--group-002.md#canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323): complete subsection reference.

- [public_ip](resources--origin_pool--reference--group-002.md#canonical-2012202102220023-2022233311031000-2131130202320231-1200001300020311-2032322312323113-2110322301123032-3132132030003331-0200031213133012): complete subsection reference.

- [public_name](resources--origin_pool--reference--group-002.md#canonical-0101123032132221-2211032101001201-1000311221111203-2133111213011322-0333020102301313-3100023301322230-0212222321323322-2203130310321323): complete subsection reference.

- [vn_private_ip](resources--origin_pool--reference--group-002.md#canonical-2130122232222011-2300321030110011-1031031023222200-1001031200312013-2120002123011312-1310133101323000-3001103213310123-2023322321313313): complete subsection reference.

- [vn_private_name](resources--origin_pool--reference--group-002.md#canonical-1013210030120200-0300013133323110-0211222301220003-2111231123223232-2130003221211000-0010213100230320-3322112211323121-2102120111321112): complete subsection reference.

<a id="canonical-2303013110110210-2023133023021110-0230000132133000-0312323123123312-2032010231003321-1013102213003120-2221231021012132-3132322312010012"></a>

## Next pages — origin_servers / 330013001011 / 5

- [origin_servers.cbip_service](resources--origin_pool--reference--group-002.md#canonical-1010331102312323-2231313310123113-1310231000100331-3110100312012302-3331111003312110-2102012030220303-1220320003131210-3032300320110001)
- [origin_servers.consul_service](resources--origin_pool--reference--group-002.md#canonical-3213211103303131-0231121111132123-2203320211303212-3333220121203231-1310101322112130-2012321333303031-3200113301012210-0120330123110331)
- [origin_servers.custom_endpoint_object](resources--origin_pool--reference--group-002.md#canonical-1013001331311011-2021321010021121-0102302220012123-3123220033111332-2222001121311212-2212111320022231-0232110210110023-2030010232123321)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-2000133232313121-0012133003021110-2020130000113133-3222221313310322-1112103023033022-3023021201311203-3003323121302310-2122200312303300)
- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003)
- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323)
- [origin_servers.public_ip](resources--origin_pool--reference--group-002.md#canonical-2012202102220023-2022233311031000-2131130202320231-1200001300020311-2032322312323113-2110322301123032-3132132030003331-0200031213133012)
- [origin_servers.public_name](resources--origin_pool--reference--group-002.md#canonical-0101123032132221-2211032101001201-1000311221111203-2133111213011322-0333020102301313-3100023301322230-0212222321323322-2203130310321323)
- [origin_servers.vn_private_ip](resources--origin_pool--reference--group-002.md#canonical-2130122232222011-2300321030110011-1031031023222200-1001031200312013-2120002123011312-1310133101323000-3001103213310123-2023322321313313)
- [origin_servers.vn_private_name](resources--origin_pool--reference--group-002.md#canonical-1013210030120200-0300013133323110-0211222301220003-2111231123223232-2130003221211000-0010213100230320-3322112211323121-2102120111321112)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-1010331102312323-2231313310123113-1310231000100331-3110100312012302-3331111003312110-2102012030220303-1220320003131210-3032300320110001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320233331223122-2002210302311120-1222021231312022-2113121031021300-0023303301122311-3333100003200221-0113231100002101-2020032110233201"></a>

## origin_servers.cbip_service — cbip_service / 321122032222 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- origin_servers.cbip_service

<a id="canonical-2232011323023031-2123202320003122-3001223113102021-0300230212012002-2302023131212310-0111230202310022-0330201131312223-3300232200032111"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with Classic BIG-IP Service (Virtual Server).

Upstream description:

Specify origin server with Classic BIG-IP Service (Virtual Server)

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("service_name")}
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
cbip_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-1010030012032220-0113213123000210-2312300001320030-0311023223222230-0013333202211113-0311232000012202-3222303112122123-1000313001122200"></a>

## Direct properties — cbip_service / 321122032222 / 3

<a id="canonical-2122320211101032-2021033213233330-0032001101221103-3123301221332330-3111312222001230-2133120330200002-1231313132021133-0133331103123121"></a>

<a id="canonical-2322223200232313-1122032301113320-3323121331032132-2032220100111111-3021312121221102-2132112012112333-1113320323212111-0221333023023332"></a>

## service_name property — cbip_service / 321122032222 / 4

Type: `"string"`. Optional.

Name of the discovered Classic BIG-IP virtual server to be used as origin.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-0013123111113300-0100131311331302-2333202102123111-1102000121303223-0133212313130333-1001332213120311-3312031110012310-2300210200013120"></a>

## Next pages — cbip_service / 321122032222 / 5

- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-3213211103303131-0231121111132123-2203320211303212-3333220121203231-1310101322112130-2012321333303031-3200113301012210-0120330123110331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3213030312010132-1211023113320302-3030003220330123-0010020001210111-1313333031312002-1200331302203101-0302030220202230-1313131030102310"></a>

## origin_servers.consul_service — consul_service / 232020213232 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- origin_servers.consul_service

<a id="canonical-2333113120233112-1200100221101203-0113323303001231-0100302303021323-2001233003012230-0230313310333100-2311302022213231-3223011030233031"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with HashiCorp Consul service name and site information.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("service_name"),
  validators.ConflictingObjectAttributes("inside_network",
    "outside_network")}
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
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\"]"
}
```

Terraform syntax:

```terraform
consul_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-0303031221320233-2321130232300200-0001003313121310-2220223120312103-0101203112101330-1223310003321120-3333013233031013-2202212332202002"></a>

## Direct properties — consul_service / 232020213232 / 3

- [inside_network](resources--origin_pool--reference--group-002.md#canonical-1022100230111330-3022300312122303-2132133132302213-2122332200123032-2213020330200203-3213021210010112-0121223021212031-0122222223110311): complete subsection reference.

- [outside_network](resources--origin_pool--reference--group-002.md#canonical-0120120323322100-1121111000013303-3122320303023003-3132333221212132-2133011220232110-1130011012010320-2302311321001210-1203333312313133): complete subsection reference.

<a id="canonical-3331121103330031-1122123103111100-1022301220121030-2033300032111112-0000032312020012-0202330133300332-1201103120300330-2322010231020003"></a>

<a id="canonical-1322100123033121-2221010010300121-1000202302013330-1210231322231302-3000133101000203-3103122021210113-3003332330322021-2023312233232202"></a>

## service_name property — consul_service / 232020213232 / 4

Type: `"string"`. Optional.

Consul service name of this origin server will be listed, including cluster-ID. The format is
servicename:cluster-ID.

Upstream description:

Consul service name of this origin server will be listed, including cluster-ID. The format is
servicename:cluster-ID.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [site_locator](resources--origin_pool--reference--group-002.md#canonical-1312101201231112-0002322310011331-1110023112333331-3121231133013002-1301120200102100-0321023313111333-3320211030210200-0131101300310123): complete subsection reference.

- [snat_pool](resources--origin_pool--reference--group-002.md#canonical-0122132301110200-1020030232012211-1022022322333032-3321101001220302-2112322220211303-1213200030003121-0103212313232122-0313012103320203): complete subsection reference.

<a id="canonical-3123303232030312-2233320030022200-1001222313303121-0210222103103310-0213111320032012-2100003212332012-2112002323130303-2200330221201200"></a>

## Next pages — consul_service / 232020213232 / 5

- [origin_servers.consul_service.inside_network](resources--origin_pool--reference--group-002.md#canonical-1022100230111330-3022300312122303-2132133132302213-2122332200123032-2213020330200203-3213021210010112-0121223021212031-0122222223110311)
- [origin_servers.consul_service.outside_network](resources--origin_pool--reference--group-002.md#canonical-0120120323322100-1121111000013303-3122320303023003-3132333221212132-2133011220232110-1130011012010320-2302311321001210-1203333312313133)
- [origin_servers.consul_service.site_locator](resources--origin_pool--reference--group-002.md#canonical-1312101201231112-0002322310011331-1110023112333331-3121231133013002-1301120200102100-0321023313111333-3320211030210200-0131101300310123)
- [origin_servers.consul_service.snat_pool](resources--origin_pool--reference--group-002.md#canonical-0122132301110200-1020030232012211-1022022322333032-3321101001220302-2112322220211303-1213200030003121-0103212313232122-0313012103320203)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-1022100230111330-3022300312122303-2132133132302213-2122332200123032-2213020330200203-3213021210010112-0121223021212031-0122222223110311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1230312302033031-0331122010223033-2023110200111202-0303101220201122-3030003210110112-1022110102323102-0021200110233210-3013332123210112"></a>

## origin_servers.consul_service.inside_network — inside_network / 202122132202 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.consul_service](resources--origin_pool--reference--group-002.md#canonical-3213211103303131-0231121111132123-2203320211303212-3333220121203231-1310101322112130-2012321333303031-3200113301012210-0120330123110331)
- origin_servers.consul_service.inside_network

<a id="canonical-1110003230323333-3131233003322202-3030200013122101-0010032013131331-3233122313003310-1131122003320020-0020331133123233-3320333111001322"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inside network.

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
inside_network = {}
```

<a id="canonical-1323221100231113-3333003330032231-0333111232023002-3021313032120201-1211033011112031-1012301321122331-0302200100222301-1100302210122222"></a>

## Direct properties — inside_network / 202122132202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3221321213202311-2112122300301123-0231023322033300-1032201213222330-1100302230102022-0002223111213103-3333302103220230-2232332111011320"></a>

## Next pages — inside_network / 202122132202 / 4

- [origin_servers.consul_service](resources--origin_pool--reference--group-002.md#canonical-3213211103303131-0231121111132123-2203320211303212-3333220121203231-1310101322112130-2012321333303031-3200113301012210-0120330123110331)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-0120120323322100-1121111000013303-3122320303023003-3132333221212132-2133011220232110-1130011012010320-2302311321001210-1203333312313133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001320311002002-2000113302002002-1011000033102223-0222232212232313-0231112033103100-2223130011130123-1010200112132330-0213331103012100"></a>

## origin_servers.consul_service.outside_network — outside_network / 232030310231 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.consul_service](resources--origin_pool--reference--group-002.md#canonical-3213211103303131-0231121111132123-2203320211303212-3333220121203231-1310101322112130-2012321333303031-3200113301012210-0120330123110331)
- origin_servers.consul_service.outside_network

<a id="canonical-2232001021110332-2302033301112320-1120323103323230-0013222133123302-2233200031101031-3002003300003130-3333002210133323-0331331331012220"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for outside network.

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
outside_network = {}
```

<a id="canonical-0221302303101021-3202221111220203-0202021101023033-1221132320011102-2230321201131301-2121230311331232-1222323301230322-2102110100231213"></a>

## Direct properties — outside_network / 232030310231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3130202001202001-2331112132310200-2232022212200113-2010223322033233-1001130203322302-3323210032000002-3303222332011311-1312333010322300"></a>

## Next pages — outside_network / 232030310231 / 4

- [origin_servers.consul_service](resources--origin_pool--reference--group-002.md#canonical-3213211103303131-0231121111132123-2203320211303212-3333220121203231-1310101322112130-2012321333303031-3200113301012210-0120330123110331)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-1312101201231112-0002322310011331-1110023112333331-3121231133013002-1301120200102100-0321023313111333-3320211030210200-0131101300310123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1000211012213220-0230223031303313-2110232000110012-2221021323211001-3312312211203331-0122233302201022-1110100301021323-0321310201330330"></a>

## origin_servers.consul_service.site_locator — site_locator / 112210310132 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.consul_service](resources--origin_pool--reference--group-002.md#canonical-3213211103303131-0231121111132123-2203320211303212-3333220121203231-1310101322112130-2012321333303031-3200113301012210-0120330123110331)
- origin_servers.consul_service.site_locator

<a id="canonical-2310223132002331-1020201010020121-2312303330130331-0103132010130222-1303311030100202-3122103303022111-2231301021212331-0213310312013010"></a>

Type: `"object"`. single nested block, Optional.

Message defines a reference to a site or virtual site object.

Upstream description:

This message defines a reference to a site or virtual site object.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
    "virtual_site")}
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
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

Terraform syntax:

```terraform
site_locator {
  # Configure direct properties listed below.
}
```

<a id="canonical-3113112323222112-3111323202122002-3100321003010320-0322301111313311-0202213320213213-3312103301021231-0213100331230212-0203332100012033"></a>

## Direct properties — site_locator / 112210310132 / 3

- [site](resources--origin_pool--reference--group-002.md#canonical-3001002303122033-1023221130211231-3322312210110112-0220222332023231-0032302101213030-0013122331030312-3310120323321201-3231132211322033): complete subsection reference.

- [virtual_site](resources--origin_pool--reference--group-002.md#canonical-0300103323003203-3201321001103133-0110200012122022-2311312330000332-2001323230310032-2332333011100303-0212110200230323-2331012020113330): complete subsection reference.

<a id="canonical-0000103231010233-1221033020330110-2323113123233200-2312202000103232-1003310121212100-2122233303022112-0133302203123212-3001021110011220"></a>

## Next pages — site_locator / 112210310132 / 4

- [origin_servers.consul_service.site_locator.site](resources--origin_pool--reference--group-002.md#canonical-3001002303122033-1023221130211231-3322312210110112-0220222332023231-0032302101213030-0013122331030312-3310120323321201-3231132211322033)
- [origin_servers.consul_service.site_locator.virtual_site](resources--origin_pool--reference--group-002.md#canonical-0300103323003203-3201321001103133-0110200012122022-2311312330000332-2001323230310032-2332333011100303-0212110200230323-2331012020113330)
- [origin_servers.consul_service](resources--origin_pool--reference--group-002.md#canonical-3213211103303131-0231121111132123-2203320211303212-3333220121203231-1310101322112130-2012321333303031-3200113301012210-0120330123110331)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-3001002303122033-1023221130211231-3322312210110112-0220222332023231-0032302101213030-0013122331030312-3310120323321201-3231132211322033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212210210221202-2321212232230203-3213233213022302-2003231220213023-3200201312000211-1122210002002112-3323133113021100-2102000201303030"></a>

## origin_servers.consul_service.site_locator.site — site / 210310200213 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.consul_service](resources--origin_pool--reference--group-002.md#canonical-3213211103303131-0231121111132123-2203320211303212-3333220121203231-1310101322112130-2012321333303031-3200113301012210-0120330123110331)
- [origin_servers.consul_service.site_locator](resources--origin_pool--reference--group-002.md#canonical-1312101201231112-0002322310011331-1110023112333331-3121231133013002-1301120200102100-0321023313111333-3320211030210200-0131101300310123)
- origin_servers.consul_service.site_locator.site

<a id="canonical-0102200331031133-2200301213330313-2022230102112212-3200231331300203-2122202133333020-2021003201001112-3211021131320011-1133231100030030"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-0030323001312200-2011303222233203-3221111031312212-1203121113121021-3310000332303002-1101222220102230-1331002130121013-1113013232000202"></a>

## Direct properties — site / 210310200213 / 3

<a id="canonical-0230133131031313-0020111122332201-0020233003303311-0033013122321333-2111022222011123-3020232022131123-3203320220031031-0103320321101012"></a>

<a id="canonical-2213323012123102-2311300232200333-3213302113211030-0123320102220223-2213232323102312-2232332210012233-3231300100131202-2101121121332332"></a>

## name property — site / 210310200213 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-1320320210002112-3300303032132200-3130312230203123-2221223121311322-2121212001010102-0021130300220122-2223001210123221-1332303013312222"></a>

<a id="canonical-3222230221330001-0010012221311133-0310330013333201-3312201013220130-1222031230323210-0100001222233033-3022233302232321-1011023322101203"></a>

## namespace property — site / 210310200213 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-1222232010212320-2203033013203330-0310302322212333-2102021022303033-3032030230213112-2312322023023231-2303010121230031-2011203301333002"></a>

<a id="canonical-2110330123210002-1201020003120110-1230000121010221-0313220321321022-3203113033211122-0320230233110010-2010212320113200-1303123123212230"></a>

## tenant property — site / 210310200213 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-3313102323201301-2123032132302333-2331021322301220-3003311201222322-2010310212112030-3301103112011013-0212120100003103-3031000103223132"></a>

## Next pages — site / 210310200213 / 7

- [origin_servers.consul_service.site_locator](resources--origin_pool--reference--group-002.md#canonical-1312101201231112-0002322310011331-1110023112333331-3121231133013002-1301120200102100-0321023313111333-3320211030210200-0131101300310123)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-0300103323003203-3201321001103133-0110200012122022-2311312330000332-2001323230310032-2332333011100303-0212110200230323-2331012020113330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030011022031211-2300101120011312-1330231113222302-2220013002023122-2121323213322101-0330111332110201-0312203330120102-0333013013030122"></a>

## origin_servers.consul_service.site_locator.virtual_site — virtual_site / 201031323232 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.consul_service](resources--origin_pool--reference--group-002.md#canonical-3213211103303131-0231121111132123-2203320211303212-3333220121203231-1310101322112130-2012321333303031-3200113301012210-0120330123110331)
- [origin_servers.consul_service.site_locator](resources--origin_pool--reference--group-002.md#canonical-1312101201231112-0002322310011331-1110023112333331-3121231133013002-1301120200102100-0321023313111333-3320211030210200-0131101300310123)
- origin_servers.consul_service.site_locator.virtual_site

<a id="canonical-1021233000001300-3231302203020232-0310123302131333-0313320330311122-2011300001133002-2021222311202230-0031000233303033-2203112002201112"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-0223233331100001-0012321000002302-3010003121231212-3333030220012030-3111220302210110-2130320233331112-3100132310212031-3130020232330230"></a>

## Direct properties — virtual_site / 201031323232 / 3

<a id="canonical-0103103310212302-1122012323220112-2213131202322203-0032231032310320-3101011003101031-2020100200332310-1220333200011201-0221000123020300"></a>

<a id="canonical-3013313133210122-0003003001321302-2311210201211131-1122123322002303-0130321201211311-2113022311203321-0000123201233032-3333232210002012"></a>

## name property — virtual_site / 201031323232 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-3201013012023030-0011210031222311-0223113032203201-3012301322020320-3022111333111132-1131031132223230-2113320032123311-1012200033021133"></a>

<a id="canonical-2112313302313200-2231030000223120-0022320313323310-2122110030010111-0213310123032132-2231322102312213-2133013011302100-0301021320203032"></a>

## namespace property — virtual_site / 201031323232 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-0330031120312222-2102221012032032-1332201113313100-1101011301112313-0210002301311302-1011101120020310-1211313302013311-1220312201301132"></a>

<a id="canonical-1211230310002302-2220102221111103-2021313333111111-2310022232121331-1210100300020331-2321310030221102-3213031302023010-2111201110123133"></a>

## tenant property — virtual_site / 201031323232 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-3013233112222111-2302312230223203-1100321331112210-0330330223222312-1132200112030021-3022323313032321-3320332133000303-2202301320111210"></a>

## Next pages — virtual_site / 201031323232 / 7

- [origin_servers.consul_service.site_locator](resources--origin_pool--reference--group-002.md#canonical-1312101201231112-0002322310011331-1110023112333331-3121231133013002-1301120200102100-0321023313111333-3320211030210200-0131101300310123)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-0122132301110200-1020030232012211-1022022322333032-3321101001220302-2112322220211303-1213200030003121-0103212313232122-0313012103320203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323130223232232-2010322203200221-1211203112113000-2301102200221111-3131232030111213-3010103013332101-3200322032302202-2322203312010223"></a>

## origin_servers.consul_service.snat_pool — snat_pool / 230023001302 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.consul_service](resources--origin_pool--reference--group-002.md#canonical-3213211103303131-0231121111132123-2203320211303212-3333220121203231-1310101322112130-2012321333303031-3200113301012210-0120330123110331)
- origin_servers.consul_service.snat_pool

<a id="canonical-3310112312023012-2030333223002320-3223031313200013-3132013031131332-0011203013002010-0320303321120323-2023303001223110-0120100321122332"></a>

Type: `"object"`. single nested block, Optional.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_snat_pool",
    "snat_pool")}
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
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

Terraform syntax:

```terraform
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-3230313301101001-0031130111113321-2120310303303132-2202012200321201-0102333202333113-1033321122312210-1110231100112011-3022321133232011"></a>

## Direct properties — snat_pool / 230023001302 / 3

- [no_snat_pool](resources--origin_pool--reference--group-002.md#canonical-1311301202032301-0010131323121321-3023220101220133-1303221300230203-0331201203223102-0020103130221031-2032003303012010-0220233221133320): complete subsection reference.

- [snat_pool](resources--origin_pool--reference--group-002.md#canonical-0110020331220233-1201321101022232-0300121022312111-2222310201003030-3302110201020011-3202132230010200-1010330121130302-2023031233211302): complete subsection reference.

<a id="canonical-1130333202313020-3131112002320012-0022112211132231-0300332300201321-2122003303302100-0113113102221231-2320220232322030-2031002203313103"></a>

## Next pages — snat_pool / 230023001302 / 4

- [origin_servers.consul_service.snat_pool.no_snat_pool](resources--origin_pool--reference--group-002.md#canonical-1311301202032301-0010131323121321-3023220101220133-1303221300230203-0331201203223102-0020103130221031-2032003303012010-0220233221133320)
- [origin_servers.consul_service.snat_pool.snat_pool](resources--origin_pool--reference--group-002.md#canonical-0110020331220233-1201321101022232-0300121022312111-2222310201003030-3302110201020011-3202132230010200-1010330121130302-2023031233211302)
- [origin_servers.consul_service](resources--origin_pool--reference--group-002.md#canonical-3213211103303131-0231121111132123-2203320211303212-3333220121203231-1310101322112130-2012321333303031-3200113301012210-0120330123110331)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-1311301202032301-0010131323121321-3023220101220133-1303221300230203-0331201203223102-0020103130221031-2032003303012010-0220233221133320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203333020220201-2321112320233000-2322032332312033-2131232020033303-0031102032213021-0023323030303221-1313302031100002-3013300211021110"></a>

## origin_servers.consul_service.snat_pool.no_snat_pool — no_snat_pool / 232310113112 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.consul_service](resources--origin_pool--reference--group-002.md#canonical-3213211103303131-0231121111132123-2203320211303212-3333220121203231-1310101322112130-2012321333303031-3200113301012210-0120330123110331)
- [origin_servers.consul_service.snat_pool](resources--origin_pool--reference--group-002.md#canonical-0122132301110200-1020030232012211-1022022322333032-3321101001220302-2112322220211303-1213200030003121-0103212313232122-0313012103320203)
- origin_servers.consul_service.snat_pool.no_snat_pool

<a id="canonical-1211011012230320-1213331000002222-3332122112303332-0320110033300211-2331230010122230-1222003100012031-3311012222110131-1113130022031221"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no snat pool.

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
no_snat_pool = {}
```

<a id="canonical-1232233003231110-0332311100022310-1203321312200213-0031011032230003-3200313031202301-1221323011200100-2212322230330121-0200102203321133"></a>

## Direct properties — no_snat_pool / 232310113112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3130031001202233-1121221120131130-2002321030333332-3213320102233210-3112330312120331-1021111322022032-1330232333113010-2010323130313023"></a>

## Next pages — no_snat_pool / 232310113112 / 4

- [origin_servers.consul_service.snat_pool](resources--origin_pool--reference--group-002.md#canonical-0122132301110200-1020030232012211-1022022322333032-3321101001220302-2112322220211303-1213200030003121-0103212313232122-0313012103320203)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-0110020331220233-1201321101022232-0300121022312111-2222310201003030-3302110201020011-3202132230010200-1010330121130302-2023031233211302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113110032321113-0101030223210031-1021103010003331-0022213223331001-1210100321031323-3311110112331013-0133120123011231-3113331023112211"></a>

## origin_servers.consul_service.snat_pool.snat_pool — snat_pool / 133312213330 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.consul_service](resources--origin_pool--reference--group-002.md#canonical-3213211103303131-0231121111132123-2203320211303212-3333220121203231-1310101322112130-2012321333303031-3200113301012210-0120330123110331)
- [origin_servers.consul_service.snat_pool](resources--origin_pool--reference--group-002.md#canonical-0122132301110200-1020030232012211-1022022322333032-3321101001220302-2112322220211303-1213200030003121-0103212313232122-0313012103320203)
- origin_servers.consul_service.snat_pool.snat_pool

<a id="canonical-0211201212313032-1220211223222131-2221020302131322-3033301111320201-3122311310023333-1202131220131030-0321032330022113-0020231023331000"></a>

Type: `"object"`. single nested block, Optional.

List of IPv4 prefixes that represent an endpoint.

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
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-2130231033222233-0121201300331201-3120213000211200-3021002032033232-1332001022023101-3220021023010131-1020320233320230-2121300122233120"></a>

## Direct properties — snat_pool / 133312213330 / 3

<a id="canonical-2020120101220132-1133213231212013-1230130013023000-0032220131233001-2202322011030131-0301310311112132-2012320021032132-3100332302323020"></a>

<a id="canonical-1121312001232221-0020002223020332-1310321002301002-3103110002002303-3220002113300022-0232333121331122-1103202103311023-3220003130001023"></a>

## prefixes property — snat_pool / 133312213330 / 4

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3123102130102331-3310001020332331-2003201230201230-2201232123230202-0100102312111111-3311003001030312-0033323203202113-1222031201032130"></a>

## Next pages — snat_pool / 133312213330 / 5

- [origin_servers.consul_service.snat_pool](resources--origin_pool--reference--group-002.md#canonical-0122132301110200-1020030232012211-1022022322333032-3321101001220302-2112322220211303-1213200030003121-0103212313232122-0313012103320203)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-1013001331311011-2021321010021121-0102302220012123-3123220033111332-2222001121311212-2212111320022231-0232110210110023-2030010232123321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011002211000202-3202211103322012-1003231033030230-1013311031112323-3032121011130313-2012010122031000-0021033122330202-0212233330301102"></a>

## origin_servers.custom_endpoint_object — custom_endpoint_object / 133331310003 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- origin_servers.custom_endpoint_object

<a id="canonical-1302022212330303-2203322202131021-3122332203210233-0223300301103330-1033111103332132-0013311211103211-3121130131121103-0320231132011301"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with a reference to endpoint object.

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
custom_endpoint_object {
  # Configure direct properties listed below.
}
```

<a id="canonical-3212310102300033-3033121233032221-0233030012321031-2312023300011033-3122303020302311-0201321003222031-0333223211302221-2322323130131332"></a>

## Direct properties — custom_endpoint_object / 133331310003 / 3

- [endpoint](resources--origin_pool--reference--group-002.md#canonical-0323112201323111-2303311221000332-3212210303320022-1112320113010011-1130220013100022-0103333330220020-2003313100112011-1132200331103122): complete subsection reference.

<a id="canonical-2313003331312011-3030330332301100-3120103113213333-1033200313111313-3021301201112103-3130321223232112-1212331003132110-0310202222121201"></a>

## Next pages — custom_endpoint_object / 133331310003 / 4

- [origin_servers.custom_endpoint_object.endpoint](resources--origin_pool--reference--group-002.md#canonical-0323112201323111-2303311221000332-3212210303320022-1112320113010011-1130220013100022-0103333330220020-2003313100112011-1132200331103122)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-0323112201323111-2303311221000332-3212210303320022-1112320113010011-1130220013100022-0103333330220020-2003313100112011-1132200331103122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1313113220132233-1221030312312023-3112003231103112-3000002120210123-0011322110021122-3221121321230121-2232012200211302-2021113310133121"></a>

## origin_servers.custom_endpoint_object.endpoint — endpoint / 113313221233 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.custom_endpoint_object](resources--origin_pool--reference--group-002.md#canonical-1013001331311011-2021321010021121-0102302220012123-3123220033111332-2222001121311212-2212111320022231-0232110210110023-2030010232123321)
- origin_servers.custom_endpoint_object.endpoint

<a id="canonical-2220012331112233-1013030220230023-2202120000200202-3202200203332230-0100201030310220-1230133131232303-3022312310213330-1223130103311121"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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
endpoint {
  # Configure direct properties listed below.
}
```

<a id="canonical-2102220312003311-2122131101132213-2022321321333021-1311221303200331-0322332310012300-3303212223123110-0311133203121332-1230323233001321"></a>

## Direct properties — endpoint / 113313221233 / 3

<a id="canonical-0003320112113300-0233311102103212-0032210321031003-1311112023212300-2021230322230020-2311330032120001-0231210233323020-0121012232203121"></a>

<a id="canonical-0233030013210003-3133211301300303-1100120220012223-3323032232013330-1200003121113202-0303113033130013-1321301013201223-2123303032122233"></a>

## name property — endpoint / 113313221233 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-1213101011112013-1233103301132302-2003300300212122-1133220220002211-1333110113112221-1233221232000313-0300111311300232-0312003120102112"></a>

<a id="canonical-2212030110100322-0133130302330232-0131311230322132-3213332203031002-2301222020133023-3122020301200110-3032120211031010-2312222033100031"></a>

## namespace property — endpoint / 113313221233 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-3330231212331013-2101100320123121-2312213221032033-1100233211312111-0320330002222031-3200032130101023-3130301011122012-2210101201330221"></a>

<a id="canonical-2103030210120300-0323030213233231-3200233232233032-0022012030302000-2331330002322311-3233120112020312-3303122032020123-0313210230221100"></a>

## tenant property — endpoint / 113313221233 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-2130233213232032-2011000112330301-2201131122320013-3033230103110032-0032130301121333-0303013203320031-2133213110313013-3003333321310113"></a>

## Next pages — endpoint / 113313221233 / 7

- [origin_servers.custom_endpoint_object](resources--origin_pool--reference--group-002.md#canonical-1013001331311011-2021321010021121-0102302220012123-3123220033111332-2222001121311212-2212111320022231-0232110210110023-2030010232123321)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-2000133232313121-0012133003021110-2020130000113133-3222221313310322-1112103023033022-3023021201311203-3003323121302310-2122200312303300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030313230012320-1211023112131133-3112312221130221-1002000101301331-3133232230210213-0232030121011033-0000003021312000-2322222332120333"></a>

## origin_servers.k8s_service — k8s_service / 102002112203 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- origin_servers.k8s_service

<a id="canonical-1030300312310013-0232032320031233-0321120322100000-1203211033132131-0011131303210313-3312332103010313-1112021202002011-0333111102120233"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with K8s service name and site information.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("inside_network",
    "outside_network"),
  validators.ConflictingObjectAttributes("inside_network",
    "vk8s_networks"),
  validators.ConflictingObjectAttributes("outside_network",
    "vk8s_networks")}
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
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"vk8s_networks\"]",
  "x-ves-oneof-field-service_info": "[\"service_name\"]"
}
```

Terraform syntax:

```terraform
k8s_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-2332222121113123-2000223112123222-0323130001312103-0020223122123033-1033322201212102-0213010133001333-1323211333102223-1130121323310102"></a>

## Direct properties — k8s_service / 102002112203 / 3

- [inside_network](resources--origin_pool--reference--group-002.md#canonical-0001033333001312-1110000223233102-1301013231111010-2202132121010201-0001211300033022-1122123130321202-1323202300302333-1310032301322011): complete subsection reference.

- [outside_network](resources--origin_pool--reference--group-002.md#canonical-2032033302122022-1103111020001001-2303301333213001-3102322130322113-3113231020002223-3033131021221332-2021231222100202-3011311100002022): complete subsection reference.

<a id="canonical-0311121300123202-2221210100230022-1233300103220023-2013022122201102-0101123301122032-0120131013101312-0112020211120133-1230322023012303"></a>

<a id="canonical-0212033311321021-3100203232300001-0132110132120230-2220330201130323-1033103220021233-1323330013010110-2003013222013333-1000103211011110"></a>

## protocol property — k8s_service / 102002112203 / 4

Type: `"string"`. Optional.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_UDP\] Type of protocol - PROTOCOL\_TCP: TCP - PROTOCOL\_UDP: UDP.
Possible values are \`PROTOCOL\_TCP\`, \`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Upstream description:

Type of protocol

&#8203;- PROTOCOL\_TCP: TCP

&#8203;- PROTOCOL\_UDP: UDP.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["PROTOCOL_TCP","PROTOCOL_UDP"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("PROTOCOL_TCP",
    "PROTOCOL_UDP"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "PROTOCOL_TCP",
  "enum": [
    "PROTOCOL_TCP",
    "PROTOCOL_UDP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3121002331113031-3230020310313222-2210212222223132-3211103300021220-0001213102332021-0302210231311112-0303111122301210-2130131333030202"></a>

<a id="canonical-2001233222022313-1223301113012201-2211021101012302-2113212203032303-3112112122320210-0030321332020332-1211131022211233-1000030313230220"></a>

## service_name property — k8s_service / 102002112203 / 5

Type: `"string"`. Optional.

Exclusive with \[\] K8s service name of the origin server will be listed, including the namespace
and cluster-ID. For vK8s services, you need to enter a string with the format
servicename.namespace:example-namespace'frontend', namespace is 'speedtest' and cluster-ID is
'prod', then you will enter..

Upstream description:

Exclusive with \[\] K8s service name of the origin server will be listed, including the namespace
and cluster-ID. For vK8s services, you need to enter a string with the format
servicename.namespace:example-namespace"frontend", namespace is "speedtest" and cluster-ID is
"prod", then you will enter "frontend.speedtest:prod". Both namespace and cluster-ID are optional.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ves_service_namespace_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ves_service_namespace_name": "true"
  }
}
```

- [site_locator](resources--origin_pool--reference--group-002.md#canonical-3332302001200033-1220112003110233-0022130323331023-3131130131013222-1033121010101012-2023030012220332-1102301113320022-2113123012211320): complete subsection reference.

- [snat_pool](resources--origin_pool--reference--group-002.md#canonical-2201230313322101-1312003232333023-0103301123101322-2112130013322233-1332230303023002-1110333220300031-2210012332300023-3202121221332102): complete subsection reference.

- [vk8s_networks](resources--origin_pool--reference--group-002.md#canonical-2010112100222333-2130310003111231-1311320001103301-0331012012231200-3121203332212212-1103022100102313-2310201311120330-2120010221201230): complete subsection reference.

<a id="canonical-1030010113100231-3033003123233321-2131010122022332-1113211113023300-0330303303121313-2101032312201221-1021021012000210-3222013102221233"></a>

## Next pages — k8s_service / 102002112203 / 6

- [origin_servers.k8s_service.inside_network](resources--origin_pool--reference--group-002.md#canonical-0001033333001312-1110000223233102-1301013231111010-2202132121010201-0001211300033022-1122123130321202-1323202300302333-1310032301322011)
- [origin_servers.k8s_service.outside_network](resources--origin_pool--reference--group-002.md#canonical-2032033302122022-1103111020001001-2303301333213001-3102322130322113-3113231020002223-3033131021221332-2021231222100202-3011311100002022)
- [origin_servers.k8s_service.site_locator](resources--origin_pool--reference--group-002.md#canonical-3332302001200033-1220112003110233-0022130323331023-3131130131013222-1033121010101012-2023030012220332-1102301113320022-2113123012211320)
- [origin_servers.k8s_service.snat_pool](resources--origin_pool--reference--group-002.md#canonical-2201230313322101-1312003232333023-0103301123101322-2112130013322233-1332230303023002-1110333220300031-2210012332300023-3202121221332102)
- [origin_servers.k8s_service.vk8s_networks](resources--origin_pool--reference--group-002.md#canonical-2010112100222333-2130310003111231-1311320001103301-0331012012231200-3121203332212212-1103022100102313-2310201311120330-2120010221201230)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-0001033333001312-1110000223233102-1301013231111010-2202132121010201-0001211300033022-1122123130321202-1323202300302333-1310032301322011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310112112223332-2202101311130131-0021333131213122-1320223200112013-3310221110202131-2111100311122300-3232112032310300-3033103222233013"></a>

## origin_servers.k8s_service.inside_network — inside_network / 131021232211 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-2000133232313121-0012133003021110-2020130000113133-3222221313310322-1112103023033022-3023021201311203-3003323121302310-2122200312303300)
- origin_servers.k8s_service.inside_network

<a id="canonical-1023300302332022-1002212333203103-0200330012333210-0302332010031111-1320013012311210-0100021132200333-1022300201013233-2313233023220331"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inside network.

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
inside_network = {}
```

<a id="canonical-3120123030020321-3033201101212331-3120213120203301-1333303321323020-0222232121130000-2002033103130130-3233121030030102-2212331111220113"></a>

## Direct properties — inside_network / 131021232211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0331322331032003-2322012330202320-0202133212300210-3220233121233000-2202003221331300-1123330223213132-1323132131122201-3002313210021310"></a>

## Next pages — inside_network / 131021232211 / 4

- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-2000133232313121-0012133003021110-2020130000113133-3222221313310322-1112103023033022-3023021201311203-3003323121302310-2122200312303300)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-2032033302122022-1103111020001001-2303301333213001-3102322130322113-3113231020002223-3033131021221332-2021231222100202-3011311100002022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0230132113323012-2212320031212020-3301131220202022-2233230221103031-1130103022303302-3321021302030022-3001113212132113-2330002330322220"></a>

## origin_servers.k8s_service.outside_network — outside_network / 133030021003 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-2000133232313121-0012133003021110-2020130000113133-3222221313310322-1112103023033022-3023021201311203-3003323121302310-2122200312303300)
- origin_servers.k8s_service.outside_network

<a id="canonical-0200213003113132-2023222231331310-2230113323301202-2211222031132011-0033332200031130-2013120323322321-1330102032233232-3101213331030010"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for outside network.

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
outside_network = {}
```

<a id="canonical-1313002023002331-3003130111103121-0133303130133122-0301203300313201-2023033021332113-0331020203230011-0032031203203210-1213233023013003"></a>

## Direct properties — outside_network / 133030021003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1313101321212032-1001123213213203-1031212133120323-3331112303333210-0311202112001321-0000112000113221-3300023102133030-0233031001302330"></a>

## Next pages — outside_network / 133030021003 / 4

- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-2000133232313121-0012133003021110-2020130000113133-3222221313310322-1112103023033022-3023021201311203-3003323121302310-2122200312303300)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-3332302001200033-1220112003110233-0022130323331023-3131130131013222-1033121010101012-2023030012220332-1102301113320022-2113123012211320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3012002323010230-1113330103013311-1023222321223120-3132013013101210-3300030200123321-1310012002202012-0330220131322123-3202030322120222"></a>

## origin_servers.k8s_service.site_locator — site_locator / 320321103322 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-2000133232313121-0012133003021110-2020130000113133-3222221313310322-1112103023033022-3023021201311203-3003323121302310-2122200312303300)
- origin_servers.k8s_service.site_locator

<a id="canonical-0210110101223301-1332211132103003-2122101201000300-3232202220003030-0323132200322021-1031030230301100-3013303132331112-1301002231102332"></a>

Type: `"object"`. single nested block, Optional.

Message defines a reference to a site or virtual site object.

Upstream description:

This message defines a reference to a site or virtual site object.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
    "virtual_site")}
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
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

Terraform syntax:

```terraform
site_locator {
  # Configure direct properties listed below.
}
```

<a id="canonical-1303320233003003-2220201210330121-0032322030133320-3313123311322112-3111012001100121-1102223222332203-2313121313030211-3223322131030100"></a>

## Direct properties — site_locator / 320321103322 / 3

- [site](resources--origin_pool--reference--group-002.md#canonical-3100133211320001-3312100030232320-3031301010230301-3330030203200230-0033031210223320-2010200033333010-3223011122110301-3320211111203031): complete subsection reference.

- [virtual_site](resources--origin_pool--reference--group-002.md#canonical-0022302012110203-2130330132230330-3101130123013031-2332313300310132-0313113133213110-3100010223101033-2233301002233221-1303300221201021): complete subsection reference.

<a id="canonical-2202323310331211-3022310331232020-2020222311121231-0232213002120103-1200121231210222-2301313002220333-2020333300332002-3213012132100320"></a>

## Next pages — site_locator / 320321103322 / 4

- [origin_servers.k8s_service.site_locator.site](resources--origin_pool--reference--group-002.md#canonical-3100133211320001-3312100030232320-3031301010230301-3330030203200230-0033031210223320-2010200033333010-3223011122110301-3320211111203031)
- [origin_servers.k8s_service.site_locator.virtual_site](resources--origin_pool--reference--group-002.md#canonical-0022302012110203-2130330132230330-3101130123013031-2332313300310132-0313113133213110-3100010223101033-2233301002233221-1303300221201021)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-2000133232313121-0012133003021110-2020130000113133-3222221313310322-1112103023033022-3023021201311203-3003323121302310-2122200312303300)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-3100133211320001-3312100030232320-3031301010230301-3330030203200230-0033031210223320-2010200033333010-3223011122110301-3320211111203031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122132202132123-3212130020033233-0121113011310113-1330111311000332-2300203322111001-2201310310202101-2032212023212212-0221011120221211"></a>

## origin_servers.k8s_service.site_locator.site — site / 320012033122 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-2000133232313121-0012133003021110-2020130000113133-3222221313310322-1112103023033022-3023021201311203-3003323121302310-2122200312303300)
- [origin_servers.k8s_service.site_locator](resources--origin_pool--reference--group-002.md#canonical-3332302001200033-1220112003110233-0022130323331023-3131130131013222-1033121010101012-2023030012220332-1102301113320022-2113123012211320)
- origin_servers.k8s_service.site_locator.site

<a id="canonical-3132212002001320-1011333321223332-1320123123221213-3111323223031202-3012133121203111-0202303002033130-3131332300331001-2333322330123022"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-1031231100223010-3212110013333133-1112132300222031-0120333331310013-0202310313223323-2202313303203222-1300322110222311-2212211112031102"></a>

## Direct properties — site / 320012033122 / 3

<a id="canonical-1231302310302222-1101203001232232-1130302320321033-0000320130103230-0001002231012230-3311130121303031-3032110323331223-1113013012233012"></a>

<a id="canonical-0022022201110031-2031023123203111-0031131233022312-3011211221302322-2000023112122003-2321301030302012-2201202130320033-3301213022200303"></a>

## name property — site / 320012033122 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-2000013301321012-2001230112000330-0201112100203003-1113213330132032-1300121001130102-0122300001311330-0002333212332130-0011100301221220"></a>

<a id="canonical-2310313223021321-3000222323023121-1303330001210133-0031113330303211-2101103021320122-1333322213033221-3311011112220011-3212320100101220"></a>

## namespace property — site / 320012033122 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-2210010001330323-0311122232312312-2111301223332213-2200102000133300-0032322003322331-3211200212121033-2333002310233103-2320202311320212"></a>

<a id="canonical-0333030302320001-3230311300113220-1323312310321112-0003202311200013-3110231202220031-3132102111101013-1203103003130233-1311000322330021"></a>

## tenant property — site / 320012033122 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-0121312121302303-0110001232332300-2230112010223003-0213011201101101-3120322131012011-1300332003000021-2031201022200330-0101222013003300"></a>

## Next pages — site / 320012033122 / 7

- [origin_servers.k8s_service.site_locator](resources--origin_pool--reference--group-002.md#canonical-3332302001200033-1220112003110233-0022130323331023-3131130131013222-1033121010101012-2023030012220332-1102301113320022-2113123012211320)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-0022302012110203-2130330132230330-3101130123013031-2332313300310132-0313113133213110-3100010223101033-2233301002233221-1303300221201021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121303313030031-1113033223011330-2320120211223113-2220233312212013-3102100111201330-0023312113132323-1023210020233233-1321112221100130"></a>

## origin_servers.k8s_service.site_locator.virtual_site — virtual_site / 232331011121 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-2000133232313121-0012133003021110-2020130000113133-3222221313310322-1112103023033022-3023021201311203-3003323121302310-2122200312303300)
- [origin_servers.k8s_service.site_locator](resources--origin_pool--reference--group-002.md#canonical-3332302001200033-1220112003110233-0022130323331023-3131130131013222-1033121010101012-2023030012220332-1102301113320022-2113123012211320)
- origin_servers.k8s_service.site_locator.virtual_site

<a id="canonical-2202122132021231-3321020200323200-1310202030313310-1330230003113200-2023120102022020-3101323010320321-0320133022010001-3312132023330131"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-0021011021222132-3033301003213320-3032020013223203-3302230230003023-1203220311101010-2130312120233212-1023122122323110-0031222113010101"></a>

## Direct properties — virtual_site / 232331011121 / 3

<a id="canonical-1111333110031032-3213113231122321-3132230132020123-1123302333233320-2201113113032032-3301013313120111-0211131021101103-3011311232300232"></a>

<a id="canonical-3301033033133201-0112002310011101-0031211101201223-0022321200021130-2032022202030000-0303101002113032-3300233321203331-0103131323210303"></a>

## name property — virtual_site / 232331011121 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-0312030111303223-0203213101223030-1001010012130330-3010221020030021-0031302012311123-2301331321021112-1202202012301222-0111103021310012"></a>

<a id="canonical-0302031301113321-0121311233001032-0000323222012203-2132301001320222-1133001311101313-3302222211220232-1110001331103130-3120131303303323"></a>

## namespace property — virtual_site / 232331011121 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-2113321233311132-1131200212200302-2320031233131202-1111221313012210-0131110123231321-3201303030032220-0202212002031331-1203222023210300"></a>

<a id="canonical-2121011211301112-1333322300303221-2133110231000132-1100220012121310-2330001021002031-0131103021111333-1301101111121330-1203010332332120"></a>

## tenant property — virtual_site / 232331011121 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-1031303032303033-2121103201223231-3230233101321202-3022132310010300-3120303023312300-2200112331132030-3201003120301312-0020120120122311"></a>

## Next pages — virtual_site / 232331011121 / 7

- [origin_servers.k8s_service.site_locator](resources--origin_pool--reference--group-002.md#canonical-3332302001200033-1220112003110233-0022130323331023-3131130131013222-1033121010101012-2023030012220332-1102301113320022-2113123012211320)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-2201230313322101-1312003232333023-0103301123101322-2112130013322233-1332230303023002-1110333220300031-2210012332300023-3202121221332102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121111310130013-3111221031010102-1232112020230330-0133123000130321-2310002201202211-1211001202202030-0303330123000311-2220033130302002"></a>

## origin_servers.k8s_service.snat_pool — snat_pool / 031231231302 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-2000133232313121-0012133003021110-2020130000113133-3222221313310322-1112103023033022-3023021201311203-3003323121302310-2122200312303300)
- origin_servers.k8s_service.snat_pool

<a id="canonical-1313332303200303-0131030312220311-3230121111020130-1203211302210213-3301211032330300-0012033312120030-3301312322223121-1032312222131033"></a>

Type: `"object"`. single nested block, Optional.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_snat_pool",
    "snat_pool")}
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
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

Terraform syntax:

```terraform
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-3113302121010220-3332023233323111-3330102111213022-3123012203313302-1102021233311301-0123322100321031-0033332021023310-2301130002332122"></a>

## Direct properties — snat_pool / 031231231302 / 3

- [no_snat_pool](resources--origin_pool--reference--group-002.md#canonical-3120030123300002-3302220202331111-2130210203032133-1210003121120212-1331220113313303-3023331330303323-2302302121321021-1131001030133123): complete subsection reference.

- [snat_pool](resources--origin_pool--reference--group-002.md#canonical-0203100021110200-1222213130002321-1303223321201100-1222003111122102-3031102313132102-2212033122301233-3100012230012301-3221313203112130): complete subsection reference.

<a id="canonical-2132303202220321-1332113101133333-2230222022030122-1333202003201123-0331131132032203-2003203233221333-3133020210102323-0330002210233113"></a>

## Next pages — snat_pool / 031231231302 / 4

- [origin_servers.k8s_service.snat_pool.no_snat_pool](resources--origin_pool--reference--group-002.md#canonical-3120030123300002-3302220202331111-2130210203032133-1210003121120212-1331220113313303-3023331330303323-2302302121321021-1131001030133123)
- [origin_servers.k8s_service.snat_pool.snat_pool](resources--origin_pool--reference--group-002.md#canonical-0203100021110200-1222213130002321-1303223321201100-1222003111122102-3031102313132102-2212033122301233-3100012230012301-3221313203112130)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-2000133232313121-0012133003021110-2020130000113133-3222221313310322-1112103023033022-3023021201311203-3003323121302310-2122200312303300)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-3120030123300002-3302220202331111-2130210203032133-1210003121120212-1331220113313303-3023331330303323-2302302121321021-1131001030133123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3212312333101203-0113003133012323-2101313131032220-2211130311112300-2221022102000020-1213202022003020-3101110221330322-3030200023220333"></a>

## origin_servers.k8s_service.snat_pool.no_snat_pool — no_snat_pool / 223203022331 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-2000133232313121-0012133003021110-2020130000113133-3222221313310322-1112103023033022-3023021201311203-3003323121302310-2122200312303300)
- [origin_servers.k8s_service.snat_pool](resources--origin_pool--reference--group-002.md#canonical-2201230313322101-1312003232333023-0103301123101322-2112130013322233-1332230303023002-1110333220300031-2210012332300023-3202121221332102)
- origin_servers.k8s_service.snat_pool.no_snat_pool

<a id="canonical-3210010303322003-3021021200000033-3323212120110303-1332000331023212-1133233102021230-1322212231333012-2220131303130310-3120202011301203"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no snat pool.

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
no_snat_pool = {}
```

<a id="canonical-3100130321220023-2330210221323000-3201111012031211-2033000110001200-0322211000322220-2100101003101230-1323002320212312-3011112330003331"></a>

## Direct properties — no_snat_pool / 223203022331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0003123111112101-0322102222001230-3302310103030101-0210210011032212-0331301222011301-1232020103220212-3032332102333311-1021322000203131"></a>

## Next pages — no_snat_pool / 223203022331 / 4

- [origin_servers.k8s_service.snat_pool](resources--origin_pool--reference--group-002.md#canonical-2201230313322101-1312003232333023-0103301123101322-2112130013322233-1332230303023002-1110333220300031-2210012332300023-3202121221332102)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-0203100021110200-1222213130002321-1303223321201100-1222003111122102-3031102313132102-2212033122301233-3100012230012301-3221313203112130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3023301020131321-2031332013013221-1111331211230221-2212321110001203-1310123212220100-3301203131220312-0333113020013013-1130011331100132"></a>

## origin_servers.k8s_service.snat_pool.snat_pool — snat_pool / 011033011001 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-2000133232313121-0012133003021110-2020130000113133-3222221313310322-1112103023033022-3023021201311203-3003323121302310-2122200312303300)
- [origin_servers.k8s_service.snat_pool](resources--origin_pool--reference--group-002.md#canonical-2201230313322101-1312003232333023-0103301123101322-2112130013322233-1332230303023002-1110333220300031-2210012332300023-3202121221332102)
- origin_servers.k8s_service.snat_pool.snat_pool

<a id="canonical-2301002222010202-0121131311200100-3203331322232323-3313020223203210-2101213201012213-2113030120331002-0013111231101320-3202312001002031"></a>

Type: `"object"`. single nested block, Optional.

List of IPv4 prefixes that represent an endpoint.

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
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-1322300213000211-2332131001332111-1330323211013012-3231033121313302-3303102133230313-3110203323112032-1031102000113332-0333001332200012"></a>

## Direct properties — snat_pool / 011033011001 / 3

<a id="canonical-0232203132122110-0010012323011322-1033300302032200-2222120301013310-1310102210222222-0033302213333223-0023221332213222-1123212130320123"></a>

<a id="canonical-0300223203122012-3012100223121221-3331203311110021-0112113321021030-3032111031112113-2010333212231232-0032012321122113-2201333301003333"></a>

## prefixes property — snat_pool / 011033011001 / 4

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3310032003233303-3202010300130101-3232133120333023-2212320233132312-2122203122231200-0101013121222201-0200030221202320-3322201130213113"></a>

## Next pages — snat_pool / 011033011001 / 5

- [origin_servers.k8s_service.snat_pool](resources--origin_pool--reference--group-002.md#canonical-2201230313322101-1312003232333023-0103301123101322-2112130013322233-1332230303023002-1110333220300031-2210012332300023-3202121221332102)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-2010112100222333-2130310003111231-1311320001103301-0331012012231200-3121203332212212-1103022100102313-2310201311120330-2120010221201230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300122100302220-1200001000321020-1102333103103002-3302211003123021-3230121032232321-0230010003231000-1100333231230133-2230211112201310"></a>

## origin_servers.k8s_service.vk8s_networks — vk8s_networks / 003323131011 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-2000133232313121-0012133003021110-2020130000113133-3222221313310322-1112103023033022-3023021201311203-3003323121302310-2122200312303300)
- origin_servers.k8s_service.vk8s_networks

<a id="canonical-3232303302233102-2031023011300221-3110021223111212-0022120200222200-0311213000033220-2201201100032011-0100022002301201-2221320203122020"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for vk8s networks.

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
vk8s_networks = {}
```

<a id="canonical-1031111313302130-3110203330022311-0223132000021311-3310322311113100-1322111200231301-2211332123330031-0322020312302100-2222300211221121"></a>

## Direct properties — vk8s_networks / 003323131011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2013300220301021-0023221012003100-2020230231001133-2003202131002210-3312011000012321-2330113301013132-1013030110010201-1201031113003021"></a>

## Next pages — vk8s_networks / 003323131011 / 4

- [origin_servers.k8s_service](resources--origin_pool--reference--group-002.md#canonical-2000133232313121-0012133003021110-2020130000113133-3222221313310322-1112103023033022-3023021201311203-3003323121302310-2122200312303300)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1113211030313311-0110022301202300-1223320100321111-3311020100113322-1112333213133101-1233311310011130-1222030112123021-1001202012322201"></a>

## origin_servers.private_ip — private_ip / 022030111210 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- origin_servers.private_ip

<a id="canonical-0132122003323203-1113202003201132-3321320022031212-2230321230311222-2210101323121301-3303110010030232-3303110003233000-0303302222121003"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with private or public IP address and site information.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("inside_network",
    "outside_network"),
  validators.ConflictingObjectAttributes("inside_network",
    "segment"),
  validators.ConflictingObjectAttributes("outside_network",
    "segment")}
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
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"segment\"]",
  "x-ves-oneof-field-private_ip_choice": "[\"ip\"]"
}
```

Terraform syntax:

```terraform
private_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-2212033312111111-2020120120121321-1022023233122332-1113113003102210-0311201113130113-1322332231212120-0212023123333331-1133111123233020"></a>

## Direct properties — private_ip / 022030111210 / 3

- [inside_network](resources--origin_pool--reference--group-002.md#canonical-2333033312003333-2020012201210130-2333321300102023-2223301201133333-1131320133122101-3322011331100320-1203013102230121-3220223120201200): complete subsection reference.

<a id="canonical-3131321310122213-2013013222233311-2013100113320312-1111123221322122-3013030033122032-2032133200310211-2223132120100033-2301131030132321"></a>

<a id="canonical-0211123103301220-3231120111030302-3211002213012313-3121300112021232-2132230331212133-1203301321100133-2313130320211021-1333002010310202"></a>

## ip property — private_ip / 022030111210 / 4

Type: `"string"`. Optional.

IP. Exclusive with \[\] Private IPv4 address.

Upstream description:

Exclusive with \[\] Private IPv4 address.

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [outside_network](resources--origin_pool--reference--group-002.md#canonical-2332312302210220-3001332333131122-1100220121131321-2233230013312312-2320320220221301-1130102233012220-2120000130122303-2120113222003130): complete subsection reference.

- [segment](resources--origin_pool--reference--group-002.md#canonical-3330110221310112-1110312123133100-2123130300111320-2011312003131331-2323030230133212-2300303013130222-3200300203112331-0000030201112330): complete subsection reference.

- [site_locator](resources--origin_pool--reference--group-002.md#canonical-1000032120021203-1313202131313321-2332131231320103-1101202113310332-1231010321113221-0311211102202130-1302223110320022-1103003330102200): complete subsection reference.

- [snat_pool](resources--origin_pool--reference--group-002.md#canonical-2203322130130331-0013113233122110-1033001232011010-2112001333020221-3201120301003311-0110310200231311-3321002311100001-1022132122001012): complete subsection reference.

<a id="canonical-0213210123311112-3122311201030013-1302311323211312-2321212030102031-3231223201032123-0032132030133212-2130001300312111-1101110222220111"></a>

## Next pages — private_ip / 022030111210 / 5

- [origin_servers.private_ip.inside_network](resources--origin_pool--reference--group-002.md#canonical-2333033312003333-2020012201210130-2333321300102023-2223301201133333-1131320133122101-3322011331100320-1203013102230121-3220223120201200)
- [origin_servers.private_ip.outside_network](resources--origin_pool--reference--group-002.md#canonical-2332312302210220-3001332333131122-1100220121131321-2233230013312312-2320320220221301-1130102233012220-2120000130122303-2120113222003130)
- [origin_servers.private_ip.segment](resources--origin_pool--reference--group-002.md#canonical-3330110221310112-1110312123133100-2123130300111320-2011312003131331-2323030230133212-2300303013130222-3200300203112331-0000030201112330)
- [origin_servers.private_ip.site_locator](resources--origin_pool--reference--group-002.md#canonical-1000032120021203-1313202131313321-2332131231320103-1101202113310332-1231010321113221-0311211102202130-1302223110320022-1103003330102200)
- [origin_servers.private_ip.snat_pool](resources--origin_pool--reference--group-002.md#canonical-2203322130130331-0013113233122110-1033001232011010-2112001333020221-3201120301003311-0110310200231311-3321002311100001-1022132122001012)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-2333033312003333-2020012201210130-2333321300102023-2223301201133333-1131320133122101-3322011331100320-1203013102230121-3220223120201200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2033230320221112-3231123001231000-2222230120212113-3330030012131213-0323223020210020-3102232101121321-1233201032130133-0131133032013200"></a>

## origin_servers.private_ip.inside_network — inside_network / 130011103322 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003)
- origin_servers.private_ip.inside_network

<a id="canonical-1010312132210302-2122330213320320-3002121000121223-2321111221333320-3233220102120023-3111220020033121-0022020003220201-2003211121111023"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inside network.

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
inside_network = {}
```

<a id="canonical-3311232031311002-3103330013312012-1323202300032300-3111233001033311-3202022033210113-0332303103010133-2221101023331210-0133121012222320"></a>

## Direct properties — inside_network / 130011103322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3322122022001321-2221300202211301-0223223100002310-1300122221113333-2212300012021310-3132201231132131-1232303311331020-3222200111021032"></a>

## Next pages — inside_network / 130011103322 / 4

- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-2332312302210220-3001332333131122-1100220121131321-2233230013312312-2320320220221301-1130102233012220-2120000130122303-2120113222003130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1211301132132103-0122323333110312-3020330321003001-2203211031031133-3023212300303333-3132123203200331-0221000111112023-0003011130320210"></a>

## origin_servers.private_ip.outside_network — outside_network / 301322121210 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003)
- origin_servers.private_ip.outside_network

<a id="canonical-3020123121020013-1233223013301111-3001332131013133-2303021113203001-3120321223301031-2112101011011222-0221013312230300-3030013213102321"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for outside network.

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
outside_network = {}
```

<a id="canonical-0112232123130213-0201302332223312-3321220222102320-1132203031220132-3332230011123211-2103111113000110-3132113211013130-0300230233131000"></a>

## Direct properties — outside_network / 301322121210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0023000002311312-2120212332133312-3013201001003023-1311033022210300-0032200223132020-3131101100300220-2102202312211332-2212022001310110"></a>

## Next pages — outside_network / 301322121210 / 4

- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-3330110221310112-1110312123133100-2123130300111320-2011312003131331-2323030230133212-2300303013130222-3200300203112331-0000030201112330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101010032323132-1332120322210102-2231300020301131-3330003023310313-2120323131312300-2301021222011130-2030012232022330-1130110101231302"></a>

## origin_servers.private_ip.segment — segment / 120300330203 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003)
- origin_servers.private_ip.segment

<a id="canonical-0122201332003012-3133322332232022-1210211032033221-1113102110020313-3132031203113300-1100300320102320-1110112012323103-2100131212133310"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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
segment {
  # Configure direct properties listed below.
}
```

<a id="canonical-1121210302211112-2102331332022031-1213001222330111-2220210300030101-2302202110201300-2231023332321331-3000101101022232-0123032010313203"></a>

## Direct properties — segment / 120300330203 / 3

<a id="canonical-1223002113120231-2231221222212323-3200032101100123-0100301000333301-1310301122013232-0130201212230011-1121333303122332-1211330012331332"></a>

<a id="canonical-2320011113321101-3300131001001003-0112231333023222-1300023323013010-1100031331100210-0210231002210023-3103023222331030-3230220333120011"></a>

## name property — segment / 120300330203 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-0002023023112303-3000002202301301-2122112011220210-2300210311301112-1310303122333301-3023310112222330-3112020232133100-3013231300223301"></a>

<a id="canonical-0133000310322310-0121113133310003-2231210220130213-3110122011101000-1223020113213120-2230212010321103-1312322122023320-0302202302302121"></a>

## namespace property — segment / 120300330203 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-2211231312012130-3330102313111132-0010213002332311-3333322201120203-3210311203321223-2202121223213200-0210312203332122-3122031211132310"></a>

<a id="canonical-3233313311020021-1022033021102211-1300232302031320-1001000031313311-2333313023033221-0322232200302133-3223101233201330-0032302000001222"></a>

## tenant property — segment / 120300330203 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-3310033001032030-2123333222013132-1133033303203323-0223222301030020-3021113000333111-0322211311023033-2311021011120033-0120002113223133"></a>

## Next pages — segment / 120300330203 / 7

- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-1000032120021203-1313202131313321-2332131231320103-1101202113310332-1231010321113221-0311211102202130-1302223110320022-1103003330102200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033120212203030-0023120232002213-2303013201112122-3001121202333312-3201111012312210-2332310320311202-0110132023310310-3312322210202212"></a>

## origin_servers.private_ip.site_locator — site_locator / 011331110033 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003)
- origin_servers.private_ip.site_locator

<a id="canonical-2300231310100223-1133011131133232-1022201302332233-0201232012331201-2310030012112201-1312110221032323-1012213021333220-1123101310031303"></a>

Type: `"object"`. single nested block, Optional.

Message defines a reference to a site or virtual site object.

Upstream description:

This message defines a reference to a site or virtual site object.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
    "virtual_site")}
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
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

Terraform syntax:

```terraform
site_locator {
  # Configure direct properties listed below.
}
```

<a id="canonical-0211021111021221-3011301121233022-2312201031022332-0233213022310231-3211023331213101-3123000021120303-0030120310200231-2321200203101101"></a>

## Direct properties — site_locator / 011331110033 / 3

- [site](resources--origin_pool--reference--group-002.md#canonical-0033003220302032-1233233113310201-1221132212023222-0023132332022003-3233200113230112-0210031031003020-1120331220022312-1302013132002100): complete subsection reference.

- [virtual_site](resources--origin_pool--reference--group-002.md#canonical-2333130022331103-1112130211201220-0203022112321012-1323001330210201-2223232033131220-1223103200121120-2323132201200131-1232331312223321): complete subsection reference.

<a id="canonical-3302220200212031-1332002100012020-0121002311320003-1122303210232311-1103312120122013-2010310033232103-2031033210201131-0230311310211213"></a>

## Next pages — site_locator / 011331110033 / 4

- [origin_servers.private_ip.site_locator.site](resources--origin_pool--reference--group-002.md#canonical-0033003220302032-1233233113310201-1221132212023222-0023132332022003-3233200113230112-0210031031003020-1120331220022312-1302013132002100)
- [origin_servers.private_ip.site_locator.virtual_site](resources--origin_pool--reference--group-002.md#canonical-2333130022331103-1112130211201220-0203022112321012-1323001330210201-2223232033131220-1223103200121120-2323132201200131-1232331312223321)
- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-0033003220302032-1233233113310201-1221132212023222-0023132332022003-3233200113230112-0210031031003020-1120331220022312-1302013132002100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032230231000332-0110101002103330-3232031222020030-3213301313013121-3203311110102220-3130123303111100-1231201211210032-3012020000010113"></a>

## origin_servers.private_ip.site_locator.site — site / 121301301021 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003)
- [origin_servers.private_ip.site_locator](resources--origin_pool--reference--group-002.md#canonical-1000032120021203-1313202131313321-2332131231320103-1101202113310332-1231010321113221-0311211102202130-1302223110320022-1103003330102200)
- origin_servers.private_ip.site_locator.site

<a id="canonical-2331020232213023-3031110031113100-3210011130022031-1310130031320000-1301021211133300-0320112000033032-2013320111122100-1133023031220313"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-0310330121031203-0021000020111322-3321313103101321-2232002000031133-2023031121002111-3102112200012321-0111010130230213-0311230122021210"></a>

## Direct properties — site / 121301301021 / 3

<a id="canonical-0213220102003313-2321021101003033-0112323000222111-3312212011301311-1330003232101003-1220312113223211-1331121211323313-3121012013301312"></a>

<a id="canonical-3013333010012010-1201221213132121-3111100121312122-0001203202102100-0301312010022102-1333302302000202-0332302310033331-3232112102230310"></a>

## name property — site / 121301301021 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-0310332310121302-0012001110320332-2133332133310001-3302331031221212-1122023203031010-3121012220302020-0011011002012001-3331221202100202"></a>

<a id="canonical-3333033323313211-0231130010313131-1332100112302020-3202210311313131-1203002013333120-2111131031313200-2201031020003333-1202320332121322"></a>

## namespace property — site / 121301301021 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-3333130012100110-3133010102330332-2230012130101130-2032111033113120-2101210230323002-3121102032110030-0031220330012002-1022302023022111"></a>

<a id="canonical-2120210232103020-0122111231011201-0132201022302031-3112012010013201-1021321332120332-1323231103312230-3100001221130213-3231222201103132"></a>

## tenant property — site / 121301301021 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-2223023003232220-0302023001320123-1323123212231323-0230102013002130-3232202000102101-3201222232211211-0133223233333102-2211212312210020"></a>

## Next pages — site / 121301301021 / 7

- [origin_servers.private_ip.site_locator](resources--origin_pool--reference--group-002.md#canonical-1000032120021203-1313202131313321-2332131231320103-1101202113310332-1231010321113221-0311211102202130-1302223110320022-1103003330102200)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-2333130022331103-1112130211201220-0203022112321012-1323001330210201-2223232033131220-1223103200121120-2323132201200131-1232331312223321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112323023131331-2102100322102021-0313301033002032-2232333001132202-3121313231100302-0111310323122131-3002302322221300-0331312021323230"></a>

## origin_servers.private_ip.site_locator.virtual_site — virtual_site / 112200112113 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003)
- [origin_servers.private_ip.site_locator](resources--origin_pool--reference--group-002.md#canonical-1000032120021203-1313202131313321-2332131231320103-1101202113310332-1231010321113221-0311211102202130-1302223110320022-1103003330102200)
- origin_servers.private_ip.site_locator.virtual_site

<a id="canonical-2222230231021133-1310332222203132-2213203023213220-0203330233023131-0122101031012301-0203303032321110-2230021233220330-2121021123221222"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-0011202022201311-0112210121132020-1320100000311033-1112222013000101-0103203202213112-1100300212013021-0232302310302223-0330022313111301"></a>

## Direct properties — virtual_site / 112200112113 / 3

<a id="canonical-1013210020122120-3113310033130023-2203233010131300-1232120113110230-1012120320130103-2210000321001023-2232103311023030-1331321022012011"></a>

<a id="canonical-3132110211120020-3323320110100210-1212200221121321-0123011121211200-0002322013033021-2202021123233010-0020201203202211-2303001020031211"></a>

## name property — virtual_site / 112200112113 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-2200220020123203-3033213310312002-0113221220033300-1103023110203201-3032000302331022-2001131023203233-0200213310311000-0133101003022121"></a>

<a id="canonical-1320111221131322-2322023030333120-0200102120222000-0010321230121012-0212233223110112-2213121001332011-0202032203331233-0101130021100230"></a>

## namespace property — virtual_site / 112200112113 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-0323020000231022-2321300331100223-1221131022231231-3322133121100023-3202320023222230-0110102231120112-0202310122123211-0220310023123201"></a>

<a id="canonical-2223001220010210-3202231113220103-1311122233333301-3332221312133012-3332110202322231-0212122231011311-0031321321101323-3011312032313033"></a>

## tenant property — virtual_site / 112200112113 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-2130221321312210-2002311201210130-1221022200030013-3110011212330223-2210200230320232-1102100101113311-2323030013211030-2233100000233212"></a>

## Next pages — virtual_site / 112200112113 / 7

- [origin_servers.private_ip.site_locator](resources--origin_pool--reference--group-002.md#canonical-1000032120021203-1313202131313321-2332131231320103-1101202113310332-1231010321113221-0311211102202130-1302223110320022-1103003330102200)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-2203322130130331-0013113233122110-1033001232011010-2112001333020221-3201120301003311-0110310200231311-3321002311100001-1022132122001012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3120001111101032-2131212220033110-2321013103311030-0332332220122320-0011310022101011-0333222323023003-2323002003103230-2003330233101130"></a>

## origin_servers.private_ip.snat_pool — snat_pool / 033023022213 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003)
- origin_servers.private_ip.snat_pool

<a id="canonical-1230022110130033-0122300323220323-2233303320012120-2123001031323003-0331000322013032-3031212033032211-0210003203200313-3110000013232222"></a>

Type: `"object"`. single nested block, Optional.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_snat_pool",
    "snat_pool")}
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
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

Terraform syntax:

```terraform
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-3111031022233122-3010111011010231-2003103022321021-0211113123011233-0130112223320032-2123303100110023-2312331132330003-1221230211310312"></a>

## Direct properties — snat_pool / 033023022213 / 3

- [no_snat_pool](resources--origin_pool--reference--group-002.md#canonical-2223123321310311-0022101202202202-2313330102010022-2220011122020023-1030013111210103-1122131100131311-1000123233211222-1311201133113313): complete subsection reference.

- [snat_pool](resources--origin_pool--reference--group-002.md#canonical-0223213203101123-1313311203312233-1030000202323301-1000222111110203-0212020020331110-0120321020133203-1022120113221100-3331221220332032): complete subsection reference.

<a id="canonical-3313202100012012-1222312121232213-1100210230120232-1220332023232013-3002231232332202-1320201310200323-0323311132223131-2101222203022222"></a>

## Next pages — snat_pool / 033023022213 / 4

- [origin_servers.private_ip.snat_pool.no_snat_pool](resources--origin_pool--reference--group-002.md#canonical-2223123321310311-0022101202202202-2313330102010022-2220011122020023-1030013111210103-1122131100131311-1000123233211222-1311201133113313)
- [origin_servers.private_ip.snat_pool.snat_pool](resources--origin_pool--reference--group-002.md#canonical-0223213203101123-1313311203312233-1030000202323301-1000222111110203-0212020020331110-0120321020133203-1022120113221100-3331221220332032)
- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-2223123321310311-0022101202202202-2313330102010022-2220011122020023-1030013111210103-1122131100131311-1000123233211222-1311201133113313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120301311122103-2211221130322330-1333323101313132-2101121112322001-3022323031112032-0120331022013323-0001023120232323-1323013002211010"></a>

## origin_servers.private_ip.snat_pool.no_snat_pool — no_snat_pool / 121213112330 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003)
- [origin_servers.private_ip.snat_pool](resources--origin_pool--reference--group-002.md#canonical-2203322130130331-0013113233122110-1033001232011010-2112001333020221-3201120301003311-0110310200231311-3321002311100001-1022132122001012)
- origin_servers.private_ip.snat_pool.no_snat_pool

<a id="canonical-1213000033011202-1103310310032130-1003133221100132-2020230011311220-1103111231331110-0213111131321230-2330010102223102-1132132222213302"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no snat pool.

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
no_snat_pool = {}
```

<a id="canonical-0320001002302030-3000332023020000-0201233232022311-3113010322132032-2110101012031223-0213221130112031-1302331300232200-2110202231023331"></a>

## Direct properties — no_snat_pool / 121213112330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1230132111231313-2322003020032313-3100321220331222-3022201112130221-1110111230300022-3102300021101213-3013103321121013-0131113313210001"></a>

## Next pages — no_snat_pool / 121213112330 / 4

- [origin_servers.private_ip.snat_pool](resources--origin_pool--reference--group-002.md#canonical-2203322130130331-0013113233122110-1033001232011010-2112001333020221-3201120301003311-0110310200231311-3321002311100001-1022132122001012)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-0223213203101123-1313311203312233-1030000202323301-1000222111110203-0212020020331110-0120321020133203-1022120113221100-3331221220332032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311331302231020-3003331322301210-3021122011010110-1312211003213010-2111133100223030-2033003202010302-3013121223313322-3013330023031033"></a>

## origin_servers.private_ip.snat_pool.snat_pool — snat_pool / 300133203002 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_ip](resources--origin_pool--reference--group-002.md#canonical-0230123122031010-0231002312323012-1310211030230020-2100002130223132-3313222333000000-0212233130001122-3202122200300023-1011131120231003)
- [origin_servers.private_ip.snat_pool](resources--origin_pool--reference--group-002.md#canonical-2203322130130331-0013113233122110-1033001232011010-2112001333020221-3201120301003311-0110310200231311-3321002311100001-1022132122001012)
- origin_servers.private_ip.snat_pool.snat_pool

<a id="canonical-0033010020000203-0223123013022223-0200013323230300-3310100102021313-3100202201211321-3000122310222333-3233331313312012-0111022110330203"></a>

Type: `"object"`. single nested block, Optional.

List of IPv4 prefixes that represent an endpoint.

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
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-1023112331203211-1320212021102311-3010302203332202-2212300330332201-1100010202312300-0111032311331300-2300100003130213-2001313301321221"></a>

## Direct properties — snat_pool / 300133203002 / 3

<a id="canonical-2310332121303123-3333230133231031-0211120300333113-1202323232220302-0032213011303110-1313113201221230-2313031222100132-2103101033132003"></a>

<a id="canonical-3203132000212102-0011210211030313-2320220300230133-3323320021003123-1030302130120231-2211323133332101-0213110330030100-2002130322020333"></a>

## prefixes property — snat_pool / 300133203002 / 4

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2011002323312112-0003130200233133-2102330232221233-2132302012310033-0323303300213130-2222030012322023-3312330223031102-0301033220200313"></a>

## Next pages — snat_pool / 300133203002 / 5

- [origin_servers.private_ip.snat_pool](resources--origin_pool--reference--group-002.md#canonical-2203322130130331-0013113233122110-1033001232011010-2112001333020221-3201120301003311-0110310200231311-3321002311100001-1022132122001012)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111331002200033-2001132132101202-1331331202320200-3033231230031331-0003330310233110-0302320001210220-1033113332233201-0311323132210013"></a>

## origin_servers.private_name — private_name / 202303231110 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- origin_servers.private_name

<a id="canonical-3022300021030132-2220222012003231-1032303111232231-2212212013120222-2323130212232113-0112320023321113-2313230210023201-3013310312203012"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with private or public DNS name and site information.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("dns_name"),
  validators.ConflictingObjectAttributes("inside_network",
    "outside_network"),
  validators.ConflictingObjectAttributes("inside_network",
    "segment"),
  validators.ConflictingObjectAttributes("outside_network",
    "segment")}
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
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"segment\"]"
}
```

Terraform syntax:

```terraform
private_name {
  # Configure direct properties listed below.
}
```

<a id="canonical-0032311010112203-2211032022002131-0123203110023121-1122101210231021-0021230110323230-2033021302111032-2110211230123011-0213121132331010"></a>

## Direct properties — private_name / 202303231110 / 3

<a id="canonical-0030320020312221-2112213011032231-1010223120203032-0030320313030110-2330203033200010-2000333313312010-1000133300301332-0301233013232223"></a>

<a id="canonical-2200312122331302-3120031310132220-2101133110203123-0331230120201230-2333031001311303-3332123302313120-0000310333023210-1332132320221323"></a>

## dns_name property — private_name / 202303231110 / 4

Type: `"string"`. Optional.

DNS Name. DNS Name

Upstream description:

DNS Name

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.95,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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

- [inside_network](resources--origin_pool--reference--group-002.md#canonical-2203123102332202-2332013300310111-3022033222313002-2213133113121032-1122013013203332-2023020212333100-3023233133301122-0311232310323331): complete subsection reference.

- [outside_network](resources--origin_pool--reference--group-002.md#canonical-3120201313033031-2111011302310031-3230023222102013-3032031102130021-1212010010233033-0030201131221331-0112102022203121-1231311201021120): complete subsection reference.

<a id="canonical-1001321013320301-3231130133303213-1012133100211202-0202310003112322-0320013300303023-2210013102000202-3300202320012331-0013231020013230"></a>

<a id="canonical-0130112221330303-0010201111020123-1031000111300202-3011210312011230-0312023131001331-2010232323211102-2232230321323323-2012212323220020"></a>

## refresh_interval property — private_name / 202303231110 / 5

Type: `"number"`. Optional.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Upstream description:

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 10, Maximum: 604800},
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
    "maximum": 604800,
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
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

- [segment](resources--origin_pool--reference--group-002.md#canonical-1200322213030001-3020120330301111-1323013131020133-1113202333233202-1000222302203230-1223201032001001-3323222331313231-1020213223310111): complete subsection reference.

- [site_locator](resources--origin_pool--reference--group-002.md#canonical-0212303223003130-1332322112033213-1131000022003312-2202112033132223-0122030223312100-3230301312001113-2030313001123011-1232322003133222): complete subsection reference.

- [snat_pool](resources--origin_pool--reference--group-002.md#canonical-0321112302111110-2332032022123300-3223231331023321-2023012232200220-1111231331321130-2333013103212300-2310131332030223-2112023131320203): complete subsection reference.

<a id="canonical-3331010103323103-0300031032213220-1100300333310112-0110022011110232-2031032221202111-3111003201002121-2002320102113031-1100132212313001"></a>

## Next pages — private_name / 202303231110 / 6

- [origin_servers.private_name.inside_network](resources--origin_pool--reference--group-002.md#canonical-2203123102332202-2332013300310111-3022033222313002-2213133113121032-1122013013203332-2023020212333100-3023233133301122-0311232310323331)
- [origin_servers.private_name.outside_network](resources--origin_pool--reference--group-002.md#canonical-3120201313033031-2111011302310031-3230023222102013-3032031102130021-1212010010233033-0030201131221331-0112102022203121-1231311201021120)
- [origin_servers.private_name.segment](resources--origin_pool--reference--group-002.md#canonical-1200322213030001-3020120330301111-1323013131020133-1113202333233202-1000222302203230-1223201032001001-3323222331313231-1020213223310111)
- [origin_servers.private_name.site_locator](resources--origin_pool--reference--group-002.md#canonical-0212303223003130-1332322112033213-1131000022003312-2202112033132223-0122030223312100-3230301312001113-2030313001123011-1232322003133222)
- [origin_servers.private_name.snat_pool](resources--origin_pool--reference--group-002.md#canonical-0321112302111110-2332032022123300-3223231331023321-2023012232200220-1111231331321130-2333013103212300-2310131332030223-2112023131320203)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-2203123102332202-2332013300310111-3022033222313002-2213133113121032-1122013013203332-2023020212333100-3023233133301122-0311232310323331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3230102101232311-2022011002202033-3103011231032301-1020321031102130-2033000033333013-0231012233103233-3310203201021000-1311330023103211"></a>

## origin_servers.private_name.inside_network — inside_network / 313231310010 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323)
- origin_servers.private_name.inside_network

<a id="canonical-2322312020100021-0201033330302321-0022320311110130-2203302111310130-2112111212210332-0030012002320200-2020220333231022-0100020021213131"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inside network.

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
inside_network = {}
```

<a id="canonical-3013222012113331-2113130321033203-3122012303101223-0203202331233210-2300023121100323-1101332323313210-1123230330320100-3011101223211202"></a>

## Direct properties — inside_network / 313231310010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2012220212203122-3323200321100230-0302102212222132-1012121102030311-1313320111023123-1202113332121333-2101112203221201-0223100113120230"></a>

## Next pages — inside_network / 313231310010 / 4

- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-3120201313033031-2111011302310031-3230023222102013-3032031102130021-1212010010233033-0030201131221331-0112102022203121-1231311201021120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1121333122302003-1012023110212031-2030200110033120-3203013121000102-1312231120133030-1222123133221012-2213303122302130-0311220100333312"></a>

## origin_servers.private_name.outside_network — outside_network / 122120330111 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323)
- origin_servers.private_name.outside_network

<a id="canonical-0013231003130331-3120123023103021-2320131211031233-3223331122101021-2103123112122130-2020233001233000-0013323130033302-2011301100101231"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for outside network.

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
outside_network = {}
```

<a id="canonical-3323030212222313-1101200123133212-2013303220011130-0123011111333230-2300201223331011-0331221310012321-2013022112232110-2201222002100211"></a>

## Direct properties — outside_network / 122120330111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0010132020131303-0311110102332333-1112210203111021-1213012030110010-3210002331013132-3212321023100132-2223223212132320-3021111030130302"></a>

## Next pages — outside_network / 122120330111 / 4

- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-1200322213030001-3020120330301111-1323013131020133-1113202333233202-1000222302203230-1223201032001001-3323222331313231-1020213223310111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201010011012121-1012130020321211-1230100301130220-3313032311113221-3013112033111220-3220030010021122-1013330020102123-3212202121112010"></a>

## origin_servers.private_name.segment — segment / 213110101330 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323)
- origin_servers.private_name.segment

<a id="canonical-1133101023013231-2111002010302323-2132111310121133-1101003132011232-1310123201023212-3333020132201122-2000001333102003-0200320101132013"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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
segment {
  # Configure direct properties listed below.
}
```

<a id="canonical-3011031011021330-2012122131230112-2132333303032132-3320032030331330-0002011321111231-3312231023101113-0020202231221222-0110223213211112"></a>

## Direct properties — segment / 213110101330 / 3

<a id="canonical-3130300300300122-2021121312321213-2203003013113330-1312322123211312-2132330301132333-0003220131030032-1010311023201132-0303113001021320"></a>

<a id="canonical-0301033310011303-3313222012210101-2001331301122303-1131202330020220-1103001331113203-1120211330211321-0000321333011330-1233231202220031"></a>

## name property — segment / 213110101330 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-2213202322130032-1302102012000111-0130011113033110-0332003132300102-3013320111020321-0123202013202103-2200010222222002-1313303132130320"></a>

<a id="canonical-2003331302013111-0221311022000113-2101133232111133-2120100200331221-1303213031003333-1000101231211303-1223021331022103-2322103233112302"></a>

## namespace property — segment / 213110101330 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-0020132303013213-0000202322312013-0130223103003201-2320022211103132-1011010001103111-3210213313122021-0323000330112332-0230310211022122"></a>

<a id="canonical-0200111322301202-0222130230033110-2231312303203023-0112031100111031-0120110100112232-3033122332010323-2013311202132223-1310111030322033"></a>

## tenant property — segment / 213110101330 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-1132312210301312-2223302113111131-3100202011230321-0333003200201302-2210133310101022-0300132002320301-1322320310122033-1113000121333121"></a>

## Next pages — segment / 213110101330 / 7

- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-0212303223003130-1332322112033213-1131000022003312-2202112033132223-0122030223312100-3230301312001113-2030313001123011-1232322003133222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322231131300002-1110001131112130-0203301221221232-3010103202112330-0131033102200100-2332101322333231-3321102330312301-0330230021303001"></a>

## origin_servers.private_name.site_locator — site_locator / 222201023301 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323)
- origin_servers.private_name.site_locator

<a id="canonical-1020001220021113-3201030222113010-3233123302013110-2033130012033221-0330122021331100-1113302312232110-2312031322302123-3113203103030302"></a>

Type: `"object"`. single nested block, Optional.

Message defines a reference to a site or virtual site object.

Upstream description:

This message defines a reference to a site or virtual site object.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
    "virtual_site")}
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
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

Terraform syntax:

```terraform
site_locator {
  # Configure direct properties listed below.
}
```

<a id="canonical-1233112102213320-2030203122231000-0233213201001222-1003232301133102-0321100132213231-1311233100111332-2310030023220322-3110103033110231"></a>

## Direct properties — site_locator / 222201023301 / 3

- [site](resources--origin_pool--reference--group-002.md#canonical-3202010311011020-1202031221320113-3122030101221230-2200102322211100-2123330202232233-3001123100320133-3112032000130220-0233002310102130): complete subsection reference.

- [virtual_site](resources--origin_pool--reference--group-002.md#canonical-0310120101132210-0322321121303212-2123011032321233-0312231230321133-2231302322012302-2001330223222303-0223221321130131-3023110131010102): complete subsection reference.

<a id="canonical-3011012223221000-1333022303033302-2301233321211230-2023031230331320-3003200022021201-3322111231023002-3330000103000020-3010010231220302"></a>

## Next pages — site_locator / 222201023301 / 4

- [origin_servers.private_name.site_locator.site](resources--origin_pool--reference--group-002.md#canonical-3202010311011020-1202031221320113-3122030101221230-2200102322211100-2123330202232233-3001123100320133-3112032000130220-0233002310102130)
- [origin_servers.private_name.site_locator.virtual_site](resources--origin_pool--reference--group-002.md#canonical-0310120101132210-0322321121303212-2123011032321233-0312231230321133-2231302322012302-2001330223222303-0223221321130131-3023110131010102)
- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-3202010311011020-1202031221320113-3122030101221230-2200102322211100-2123330202232233-3001123100320133-3112032000130220-0233002310102130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003330323212123-2312011010210021-2200100231122312-2130002123113110-1013110200213022-1303010121210103-3101003203320000-3013212012103132"></a>

## origin_servers.private_name.site_locator.site — site / 012212103103 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323)
- [origin_servers.private_name.site_locator](resources--origin_pool--reference--group-002.md#canonical-0212303223003130-1332322112033213-1131000022003312-2202112033132223-0122030223312100-3230301312001113-2030313001123011-1232322003133222)
- origin_servers.private_name.site_locator.site

<a id="canonical-2310110312030231-0323110012101103-1231313132110123-3322003102003131-1320220002300020-0210320122211021-3233202103231300-0312122000232233"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-1122331112312113-2332233321220222-0031330211120013-1112103201313213-1202333120030333-3332011232200231-3031121031121230-3131121232311210"></a>

## Direct properties — site / 012212103103 / 3

<a id="canonical-1233332330010203-2103101120011230-3302022311211002-3131021321120101-1030011110033031-3201111031230123-3322103220002102-0023103113220323"></a>

<a id="canonical-1131323032213200-2211331220220101-2030321331132230-1323023210112331-1121322111120230-2323000302323102-3320011121031111-2213220303110013"></a>

## name property — site / 012212103103 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-2030110230122011-0121312231332013-0103303202223333-0202112032330303-0111300300121002-0020021300130230-3302131213012013-1302221311122312"></a>

<a id="canonical-0331331211311231-1002330101302001-0023023211103222-0231212203002200-3211200112022212-1103233211132013-3203112203210111-3212303210311302"></a>

## namespace property — site / 012212103103 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-3122311300323320-0211222202200211-3102330002320302-3333310303332213-1212202113332222-3222331232300233-2232323000323223-1033321013103023"></a>

<a id="canonical-1022311031012232-3121201011013333-3333112031213203-2123233302300101-0211213113230201-1032021321113120-3013212230220100-2300113121313023"></a>

## tenant property — site / 012212103103 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-2322303113132011-3333002000212231-0330112000321302-0232000232100212-1213110231233303-2223120311001023-0032223333312021-0301312312033130"></a>

## Next pages — site / 012212103103 / 7

- [origin_servers.private_name.site_locator](resources--origin_pool--reference--group-002.md#canonical-0212303223003130-1332322112033213-1131000022003312-2202112033132223-0122030223312100-3230301312001113-2030313001123011-1232322003133222)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-0310120101132210-0322321121303212-2123011032321233-0312231230321133-2231302322012302-2001330223222303-0223221321130131-3023110131010102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112031122031333-0302033201200002-3210233133231200-3012011222303332-0331210010122033-1002321233200030-0310331123201233-0120321303300213"></a>

## origin_servers.private_name.site_locator.virtual_site — virtual_site / 322320033333 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323)
- [origin_servers.private_name.site_locator](resources--origin_pool--reference--group-002.md#canonical-0212303223003130-1332322112033213-1131000022003312-2202112033132223-0122030223312100-3230301312001113-2030313001123011-1232322003133222)
- origin_servers.private_name.site_locator.virtual_site

<a id="canonical-0310202213011000-3111221033301021-0000311230003001-3133110320333332-2332303200200020-1133002033223110-1113111032333320-0303223221312122"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-3013012201102123-1033102123030303-1020323132133130-1100013131103302-3131130232010301-3000123221010203-3023233211113221-0122210201130012"></a>

## Direct properties — virtual_site / 322320033333 / 3

<a id="canonical-3230020311221133-3320203223320221-3003211320333331-0100110201013203-0013132002113213-1033002220010200-0032321303320130-3003202000212202"></a>

<a id="canonical-0310111030121122-1023013000311201-1221021332003310-0132102323220013-1312022323323322-3201311033012302-3102231330021130-1002321213220010"></a>

## name property — virtual_site / 322320033333 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-1211202113232100-2103013212012101-3223031333022230-2020222132122330-1013333210103312-2133112122232022-0231130221203310-1301032203320222"></a>

<a id="canonical-2333333301023323-0322123332133012-2311032210230102-0322123130330020-3223012301203222-0011131003100233-3311323022221133-2023011131001021"></a>

## namespace property — virtual_site / 322320033333 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-3333120331021200-0122122122112221-1033232311130311-2001001031121012-1323101121200221-3031223233321322-1020232202213333-0032031311223013"></a>

<a id="canonical-2220120032020231-2020003131101130-0100133102133320-3222111202333331-1311320210011221-2022012003023110-3022300220113302-0111021313011201"></a>

## tenant property — virtual_site / 322320033333 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-2123202212021222-0231100212100133-2233210221100103-1212022030212021-3202320211032133-3103223232232100-2020021313222332-2012330000230113"></a>

## Next pages — virtual_site / 322320033333 / 7

- [origin_servers.private_name.site_locator](resources--origin_pool--reference--group-002.md#canonical-0212303223003130-1332322112033213-1131000022003312-2202112033132223-0122030223312100-3230301312001113-2030313001123011-1232322003133222)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-0321112302111110-2332032022123300-3223231331023321-2023012232200220-1111231331321130-2333013103212300-2310131332030223-2112023131320203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0102113211320023-2000222013122112-2332023220301331-3023103222231232-3131211313122102-0200033213301101-1203021113122130-2323333331223002"></a>

## origin_servers.private_name.snat_pool — snat_pool / 233113133133 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323)
- origin_servers.private_name.snat_pool

<a id="canonical-1012322211101010-0203023132110100-0321333121033313-1121311033211030-0012132301213012-1133130003230120-0020130231013323-3113120111010133"></a>

Type: `"object"`. single nested block, Optional.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_snat_pool",
    "snat_pool")}
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
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

Terraform syntax:

```terraform
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-3011333022132102-2333331213100030-1033033020010100-1320113332233013-1022221112110020-1311203131201113-2213212103302223-3303013031223030"></a>

## Direct properties — snat_pool / 233113133133 / 3

- [no_snat_pool](resources--origin_pool--reference--group-002.md#canonical-3130023221023210-3020131101000132-2331203133232103-1230023310301102-2223220110200130-0220200112120302-3310132133311233-3200112023312112): complete subsection reference.

- [snat_pool](resources--origin_pool--reference--group-002.md#canonical-2221131331021133-1131233031222100-2123110022013001-1003330121330001-2001002310023110-1111101220230032-3332220113200220-3133303312002021): complete subsection reference.

<a id="canonical-0331033313330013-3010000320331302-0031212033103003-1200011203133112-0331200003203003-3122323132230232-0100211020032221-3030031333202313"></a>

## Next pages — snat_pool / 233113133133 / 4

- [origin_servers.private_name.snat_pool.no_snat_pool](resources--origin_pool--reference--group-002.md#canonical-3130023221023210-3020131101000132-2331203133232103-1230023310301102-2223220110200130-0220200112120302-3310132133311233-3200112023312112)
- [origin_servers.private_name.snat_pool.snat_pool](resources--origin_pool--reference--group-002.md#canonical-2221131331021133-1131233031222100-2123110022013001-1003330121330001-2001002310023110-1111101220230032-3332220113200220-3133303312002021)
- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-3130023221023210-3020131101000132-2331203133232103-1230023310301102-2223220110200130-0220200112120302-3310132133311233-3200112023312112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221000301222333-2232032211013320-2101122012300300-1323133322310330-3201331101220202-0002021000131233-1300231230101102-1312230023313322"></a>

## origin_servers.private_name.snat_pool.no_snat_pool — no_snat_pool / 330221332322 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323)
- [origin_servers.private_name.snat_pool](resources--origin_pool--reference--group-002.md#canonical-0321112302111110-2332032022123300-3223231331023321-2023012232200220-1111231331321130-2333013103212300-2310131332030223-2112023131320203)
- origin_servers.private_name.snat_pool.no_snat_pool

<a id="canonical-0113203312310213-2110203230230023-2212211321322013-2102102302112231-3013121131100003-2120303003130323-0233330323032322-0223131100133121"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no snat pool.

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
no_snat_pool = {}
```

<a id="canonical-1213122323331200-0313121201320232-0200122101312031-0312232132031222-1313310011011310-2013323032102021-0113222322333212-2010110223103321"></a>

## Direct properties — no_snat_pool / 330221332322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0210032221032211-3023012112312330-0220223011331331-2102333112202003-2110202003201003-0332033333112302-3000013330211102-1000001223222031"></a>

## Next pages — no_snat_pool / 330221332322 / 4

- [origin_servers.private_name.snat_pool](resources--origin_pool--reference--group-002.md#canonical-0321112302111110-2332032022123300-3223231331023321-2023012232200220-1111231331321130-2333013103212300-2310131332030223-2112023131320203)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-2221131331021133-1131233031222100-2123110022013001-1003330121330001-2001002310023110-1111101220230032-3332220113200220-3133303312002021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013002003210123-0322131311021100-1302011131302222-2033301121130200-1001233010331301-3230111132010130-1131202232033131-0323121322302222"></a>

## origin_servers.private_name.snat_pool.snat_pool — snat_pool / 132301103002 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.private_name](resources--origin_pool--reference--group-002.md#canonical-1133022131320232-3123302201233302-1212311001131130-2012010123333133-0102031133330232-1000330202203203-2203003330013202-1120202100000323)
- [origin_servers.private_name.snat_pool](resources--origin_pool--reference--group-002.md#canonical-0321112302111110-2332032022123300-3223231331023321-2023012232200220-1111231331321130-2333013103212300-2310131332030223-2112023131320203)
- origin_servers.private_name.snat_pool.snat_pool

<a id="canonical-1012022110102030-0322310000331011-1121232011213022-0223112103323312-2103220000323222-3113313200023112-3120001311231113-0031212113330001"></a>

Type: `"object"`. single nested block, Optional.

List of IPv4 prefixes that represent an endpoint.

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
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-1123322200112113-0210011133232003-3321323011321223-3202023002101101-2113022303102233-2313201231221230-1213000223103331-0020021323030320"></a>

## Direct properties — snat_pool / 132301103002 / 3

<a id="canonical-0022211003233101-1123312210101003-0022330320112012-0103323033222030-1222232031131222-1020220031211132-0033330032101220-2332200002003220"></a>

<a id="canonical-0303332122110232-0213322321121203-2012231203223202-3210022111003220-1321202102032221-0302020331130102-3023000301011033-1130311111220330"></a>

## prefixes property — snat_pool / 132301103002 / 4

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3031211200113223-3131301201233101-2013112012111131-3303120000333201-1312221031021332-3133113012313003-1222131033001312-1100213011230320"></a>

## Next pages — snat_pool / 132301103002 / 5

- [origin_servers.private_name.snat_pool](resources--origin_pool--reference--group-002.md#canonical-0321112302111110-2332032022123300-3223231331023321-2023012232200220-1111231331321130-2333013103212300-2310131332030223-2112023131320203)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-2012202102220023-2022233311031000-2131130202320231-1200001300020311-2032322312323113-2110322301123032-3132132030003331-0200031213133012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0100100002113111-2332111002331120-1003002202021103-2112321021213200-0132202020010323-0021200210201032-1301113330230321-0032223233022302"></a>

## origin_servers.public_ip — public_ip / 311021221133 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- origin_servers.public_ip

<a id="canonical-2330102103323312-2311122332333212-2111322303321021-1202231123132011-2230321120110033-1323002102021311-3322132202133231-3130033013112100"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with public IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-public_ip_choice": "[\"ip\"]"
}
```

Terraform syntax:

```terraform
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-1031222203021320-1322122220322020-2210011323203200-0202100320230113-2032031020203130-1113223100030030-1300320012001313-3312213330020122"></a>

## Direct properties — public_ip / 311021221133 / 3

<a id="canonical-0301113230322103-0122131333211003-2200221332132002-0002323200320120-0322023121113332-2030330012133033-0131102333011030-1221302101232132"></a>

<a id="canonical-0122133203120101-0010111020123213-2103130333313031-1101033311133133-2301210211332220-2102220122102011-3301112300021233-0211012120000033"></a>

## ip property — public_ip / 311021221133 / 4

Type: `"string"`. Optional.

Public IPv4. Exclusive with \[\] Public IPv4 address.

Upstream description:

Exclusive with \[\] Public IPv4 address.

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-0201031110000013-3133133100001320-0102011200233011-2011003202202003-1013131201121321-1233210120233130-1021121301323120-1131011210112311"></a>

## Next pages — public_ip / 311021221133 / 5

- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-0101123032132221-2211032101001201-1000311221111203-2133111213011322-0333020102301313-3100023301322230-0212222321323322-2203130310321323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131321220010200-1302000320021330-1102220110202212-1030222100120002-2023031300111111-1211322301122203-1122212331312221-0111133023101132"></a>

## origin_servers.public_name — public_name / 120210311223 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- origin_servers.public_name

<a id="canonical-0222302133210320-2001221203003010-3022300302110110-3000333100030103-2313301012110103-2212210120131013-1012130222122313-3202312023112330"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with public DNS name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("dns_name")}
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
public_name {
  # Configure direct properties listed below.
}
```

<a id="canonical-0303200232132000-2120313211022310-0203102030330113-2131303031213111-1111221120000203-1022202100212201-3231313020200012-1202232113321200"></a>

## Direct properties — public_name / 120210311223 / 3

<a id="canonical-2111122303201022-2233210133031300-0320303100012233-1103103233312021-2032231230310321-1120300132031101-2002033321232112-1233310112133232"></a>

<a id="canonical-2123112321203203-2000033002023122-1233033330333001-1213133032032020-0222202322313023-1220030111202111-1233223013330223-1120320012222130"></a>

## dns_name property — public_name / 120210311223 / 4

Type: `"string"`. Optional.

DNS Name. DNS Name

Upstream description:

DNS Name

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3021102323312312-2101333023321220-1102102211203121-3133300301203323-3213311103001301-2121213022101223-1023111301321210-1011010022020222"></a>

<a id="canonical-1003001031111131-1000032311202231-1323102212001312-0230113211012303-3000202312320111-2123232102030003-1303213133331320-3312101020013001"></a>

## refresh_interval property — public_name / 120210311223 / 5

Type: `"number"`. Optional.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Upstream description:

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 10, Maximum: 604800},
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
    "maximum": 604800,
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
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

<a id="canonical-1122111131233022-3213331122332000-0212001231210132-0211121230210101-1131321201300211-0320013031103012-2311110013323011-3321031211203133"></a>

## Next pages — public_name / 120210311223 / 6

- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-2130122232222011-2300321030110011-1031031023222200-1001031200312013-2120002123011312-1310133101323000-3001103213310123-2023322321313313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3221012132202023-1311312101133310-3110223313302012-0121313002310333-1101331020220201-1231201332110332-2001020022021033-3221223100313211"></a>

## origin_servers.vn_private_ip — vn_private_ip / 312020100301 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- origin_servers.vn_private_ip

<a id="canonical-1212301103220103-0222133311031102-2310030020003310-3011032130010130-0333113121203022-1121200301130110-1121333133231213-2122303322012222"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with IP on Virtual Network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-virtual_network_ip_choice": "[\"ip\"]"
}
```

Terraform syntax:

```terraform
vn_private_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-2011211221000133-3213133112032133-2012322003031320-3311003312313131-1202201333201003-2102233030302200-3211112012232113-2013210221113111"></a>

## Direct properties — vn_private_ip / 312020100301 / 3

<a id="canonical-0303031013321122-3312002010330312-3300200032023111-1022212011213113-2232202312122132-0131302301220001-3101202011123200-0312301013023023"></a>

<a id="canonical-3310302122202333-0203312210121010-1303233003023230-1003211200313202-1313033320230201-0200133333001200-3132022200211232-3313120311302130"></a>

## ip property — vn_private_ip / 312020100301 / 4

Type: `"string"`. Optional.

IPv4. Exclusive with \[\] IPv4 address.

Upstream description:

Exclusive with \[\] IPv4 address.

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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [virtual_network](resources--origin_pool--reference--group-002.md#canonical-3122100030131322-0113010211030321-3203333130200133-1111310203002032-0020302020200012-0010322030112332-3310313213332211-3231002332012330): complete subsection reference.

<a id="canonical-3312002313311300-1222013110311222-2301100111300230-1202000300311121-1001032013222320-3233300333131200-3301211322011231-3002113230220003"></a>

## Next pages — vn_private_ip / 312020100301 / 5

- [origin_servers.vn_private_ip.virtual_network](resources--origin_pool--reference--group-002.md#canonical-3122100030131322-0113010211030321-3203333130200133-1111310203002032-0020302020200012-0010322030112332-3310313213332211-3231002332012330)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-3122100030131322-0113010211030321-3203333130200133-1111310203002032-0020302020200012-0010322030112332-3310313213332211-3231002332012330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3022211213221120-3202311013310330-3122310013202011-0132002112233211-3132302330032021-2010301123033221-3121033322232201-2023012331211132"></a>

## origin_servers.vn_private_ip.virtual_network — virtual_network / 122033023230 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.vn_private_ip](resources--origin_pool--reference--group-002.md#canonical-2130122232222011-2300321030110011-1031031023222200-1001031200312013-2120002123011312-1310133101323000-3001103213310123-2023322321313313)
- origin_servers.vn_private_ip.virtual_network

<a id="canonical-0003212300322202-3222000200122020-3203010013001101-1311323200003200-2102323331103133-0230313302203300-3302211300312303-3202332311301130"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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
virtual_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-2211231211330221-2032031113021110-1213213100021003-1033233311120332-1320210323301002-2232301111113131-1332221203312030-1100003133312110"></a>

## Direct properties — virtual_network / 122033023230 / 3

<a id="canonical-0000202302112102-1112012101000233-3302013312100000-0023221022201010-0201113330200201-2003201103322210-1330030120230101-2002210032322331"></a>

<a id="canonical-0310322312201233-2000332330303222-2312030031212312-1011333103200111-0132332212221311-0211213000102210-0221100201203231-1022220113220320"></a>

## name property — virtual_network / 122033023230 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-0223223101122021-2213133211011111-1110303311033212-2202212310201012-1232122321032311-0022200133002320-1203111330100322-3332031233320120"></a>

<a id="canonical-0232100012000003-0200323230122012-3133333020112233-2130322211010003-1232211102211003-0233011303013323-1031122331000322-2320102313300023"></a>

## namespace property — virtual_network / 122033023230 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-0111322301320300-0322201111323203-1212030301001200-0211310021223012-3123201121032322-2030010332303230-2231323310021331-0120323133322133"></a>

<a id="canonical-0332322222112221-0122312211222102-3320133212200103-1203223031013313-2312211210323021-1012332123222003-3031230003330301-0001233122111031"></a>

## tenant property — virtual_network / 122033023230 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-1110330203120132-0320300221011133-0100222332310021-3103013010210211-2013211213031303-3230332312301102-2302321311003132-2233003010221010"></a>

## Next pages — virtual_network / 122033023230 / 7

- [origin_servers.vn_private_ip](resources--origin_pool--reference--group-002.md#canonical-2130122232222011-2300321030110011-1031031023222200-1001031200312013-2120002123011312-1310133101323000-3001103213310123-2023322321313313)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-1013210030120200-0300013133323110-0211222301220003-2111231123223232-2130003221211000-0010213100230320-3322112211323121-2102120111321112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003101232323020-3001020300300133-1023230033123020-3212311012112300-2020230032322001-0120100110103023-2320320020331110-1213301032320100"></a>

## origin_servers.vn_private_name — vn_private_name / 232021033321 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- origin_servers.vn_private_name

<a id="canonical-1133331232233003-1320111130012311-1123222301130121-1020122210302320-1322110011220312-3110310013210300-1130202300310203-1332113133000131"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with DNS name on Virtual Network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("dns_name")}
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
vn_private_name {
  # Configure direct properties listed below.
}
```

<a id="canonical-0032130030013310-2203313201333220-0021203203211100-2033323010223311-2102213222102211-1312021331212233-1312123303320230-2020010313300020"></a>

## Direct properties — vn_private_name / 232021033321 / 3

<a id="canonical-3002030022032132-3311301011121132-2213132122130322-0202002203100100-0201311301120102-1020000023121010-2020131020121000-2302221200332000"></a>

<a id="canonical-1330330213323301-3111231123220210-0303022202230033-2312220011033230-3001303230332301-0100001021101100-3222000000111333-3133131303110033"></a>

## dns_name property — vn_private_name / 232021033321 / 4

Type: `"string"`. Optional.

DNS Name. DNS Name

Upstream description:

DNS Name

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
  stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.95,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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

- [private_network](resources--origin_pool--reference--group-002.md#canonical-3313202330121310-0332121101021311-0020212111023330-2010022013100332-0131010011102330-1300031011320320-3223132122001110-3220213212120000): complete subsection reference.

<a id="canonical-2203331123130100-2121300230221303-0332003302113001-2313331001020133-3321131310233322-1311020221323120-0132212220133322-2130133123232323"></a>

## Next pages — vn_private_name / 232021033321 / 5

- [origin_servers.vn_private_name.private_network](resources--origin_pool--reference--group-002.md#canonical-3313202330121310-0332121101021311-0020212111023330-2010022013100332-0131010011102330-1300031011320320-3223132122001110-3220213212120000)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-3313202330121310-0332121101021311-0020212111023330-2010022013100332-0131010011102330-1300031011320320-3223132122001110-3220213212120000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210130011123132-1221333223113212-3203103010121211-0030031233023212-0322030303332112-2021002113021323-2101221221003002-3030013211200300"></a>

## origin_servers.vn_private_name.private_network — private_network / 230013120002 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [origin_servers](resources--origin_pool--reference--group-001.md#canonical-3213210211103302-0113102102311332-3301020313122313-2300311021121330-3110303331323010-3002302023032312-3331303322001030-2133000010210120)
- [origin_servers.vn_private_name](resources--origin_pool--reference--group-002.md#canonical-1013210030120200-0300013133323110-0211222301220003-2111231123223232-2130003221211000-0010213100230320-3322112211323121-2102120111321112)
- origin_servers.vn_private_name.private_network

<a id="canonical-1133303321023212-2001300031121022-1011301211231302-0001021121033021-2323331011202103-0022221131021002-1223211000123212-0002221220103230"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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
private_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-0012301331302022-1303130000001202-3200103020002222-0320223322203303-0320201320021013-2222221200112120-2200011131132120-2020211101323323"></a>

## Direct properties — private_network / 230013120002 / 3

<a id="canonical-0211012010133233-1230101313033231-2203330313202001-0230213001201123-0023222333323120-1021301221110101-3130101312333010-0022323203212030"></a>

<a id="canonical-0303321311102333-1230110211213001-3123232212213210-2333121021122233-1223223203203332-0330001220230130-3311303133033201-0013012223113021"></a>

## name property — private_network / 230013120002 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-3001031121122201-2021001100001030-1223130321201021-2103201022330133-0031030202301102-2003332030311030-1013301002031230-0333123312320300"></a>

<a id="canonical-2132220312310103-2112102022022323-3113332023200133-3031022331201100-1112130301313302-0113123211312222-1212032020032202-1233321332230113"></a>

## namespace property — private_network / 230013120002 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-3203203021110231-0111012212301333-0231202321023031-0200023321203132-2122323200121130-1202232210232212-2112323113223002-3222230212232320"></a>

<a id="canonical-2033233200323323-0012131213103231-0312113101102002-1013012011203021-0232030230331321-0310320303001023-2023231223321013-1001103302233332"></a>

## tenant property — private_network / 230013120002 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-3102021032102012-2002103101033321-3100223230302200-0120211132122120-3213332132320011-1200201122312033-3230200031232013-1221021231232122"></a>

## Next pages — private_network / 230013120002 / 7

- [origin_servers.vn_private_name](resources--origin_pool--reference--group-002.md#canonical-1013210030120200-0300013133323110-0211222301220003-2111231123223232-2130003221211000-0010213100230320-3322112211323121-2102120111321112)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-3110220132232202-3223311123032232-3333333030113010-2133200310011011-3010332223210221-3113000210333321-1221001103212033-0210001201112033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203111031321212-2020012310022033-2121002303011221-0223302322101203-2311230332222101-2330301323103330-3330103231322201-1032220322233102"></a>

## same_as_endpoint_port — same_as_endpoint_port / 311001001321 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)
- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- same_as_endpoint_port

<a id="canonical-2003120102323102-0320322222233123-0110310222321212-1303023300232011-2231113032010222-0322332011312201-0332033220131211-2002332200121120"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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
same_as_endpoint_port = {}
```

<a id="canonical-0102100303313220-3031301320221100-1130013221121233-2011320012112311-1221101303232110-0200300033110301-1223122122010212-1200222331001033"></a>

## Direct properties — same_as_endpoint_port / 311001001321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1131102013003323-0113131023213211-2112203312100120-0131001332310130-1100023020310033-3331101200121001-0213133331101222-1212110230120202"></a>

## Next pages — same_as_endpoint_port / 311001001321 / 4

- [Property reference](resources--origin_pool--reference--group-001.md#canonical-3212012302103133-0133102110013031-2131221222021021-2000212221110230-2323033002013033-2112321312312211-0213033122012333-3133112201301013)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-1300222123211010-0221103312033303-3103220201011310-0312202100212232-2132331032032202-3102123303123231-2302302300131121-2101021333132222)

<a id="canonical-1322321330320022-2021311033121221-0210313231232031-3331330210233122-2333300130103032-3230223313322102-2023330330022001-1110122123303032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
