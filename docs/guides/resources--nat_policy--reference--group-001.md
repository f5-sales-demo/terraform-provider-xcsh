---
page_title: "xcsh_nat_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nat_policy reference."
---

# xcsh_nat_policy reference

<a id="canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1100330220223133-3200200031132322-3032321200020330-3103132231013112-0100322223201302-1001023101202301-3310320130100201-1020002111313130"></a>

## Property reference — Property reference / 021313213121 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- Property reference

<a id="canonical-2220303000002003-0303203230123321-1221020213133221-2312123221201121-1132022321003233-1332200231321311-0233231001231000-0211203122030113"></a>

## Direct properties — Property reference / 021313213121 / 3

<a id="canonical-0332122232121112-3120320001021122-0022132310302112-0200103031322111-2013121321132211-3023112311101201-3122232330021021-3032333111003232"></a>

<a id="canonical-1113232232231210-0211213001121230-2213302001311213-3103213333330120-2213012301213213-1003211030012031-2011101332302120-3332233112320303"></a>

## annotations property — Property reference / 021313213121 / 4

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

<a id="canonical-1010220332122020-0321123122201311-2030212131022231-1203100301032031-3122111222230222-0020133200233210-0202222110131212-2321301003032221"></a>

<a id="canonical-1013210003310012-2123123330030120-0110302131302211-1102322333110021-3130203010213011-2033313012232122-1110220230110331-1002030303331331"></a>

## description property — Property reference / 021313213121 / 5

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

<a id="canonical-0131220133130202-3222231331100001-1100331323122211-0223301201332131-1203313132231101-0233232012320230-3021030010223300-3112203332020302"></a>

<a id="canonical-1221112111023200-2011211033031021-0232233232113101-0222022013210313-3130203313202210-2332113210123022-1302021232313231-0301120210302301"></a>

## disable property — Property reference / 021313213121 / 6

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

<a id="canonical-1320020300233233-3211123202312331-3132031102233022-2033120323023310-3233210120310303-2201310332300231-0323120031131032-1332130130312132"></a>

<a id="canonical-1232133202223012-3221123300300020-3311120000112021-0011012013103221-3322320103100010-0102330333320200-2221130023230122-1310311103311233"></a>

## ID property — Property reference / 021313213121 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1131231210012200-1132130213331303-2221111011223212-2213320111100111-0200002230200110-1232313120100210-3301110313002120-2203231002331200"></a>

<a id="canonical-0001000133022231-0030100012202001-2322211022313112-2310010232331320-2132102103110200-1303000003202022-2301113231023232-2300222222000310"></a>

## labels property — Property reference / 021313213121 / 8

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

<a id="canonical-1230212310130013-1000200231313100-2330023002323030-0131331202221130-3303323023332222-1000232311210112-3001313113122102-0310131033012110"></a>

<a id="canonical-3031130121303113-3032230210201302-1010013120120321-0111022031323210-0011303021310031-2030113010320122-2233213033232313-2330303220330330"></a>

## name property — Property reference / 021313213121 / 9

Type: `"string"`. Required.

Name of the NAT Policy. Must be unique within the namespace.

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

<a id="canonical-2200223211211102-2323121031031230-0111022031310222-1130031212113032-2133312210210022-0022132323331320-0111202321223301-3013322002130122"></a>

<a id="canonical-2223301031131200-0302031323110010-1133033132200223-0330001013303301-2221233121211212-3330121300323133-0032133211200122-3110122311112232"></a>

## namespace property — Property reference / 021313213121 / 10

Type: `"string"`. Required.

Namespace where the NAT Policy is created.

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

- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132): complete subsection reference.

- [site](resources--nat_policy--reference--group-001.md#canonical-3001223222011220-0003203113332212-3332203303130003-1310223301012333-2223122022201100-2321110020002221-0120100011321320-2020123030023230): complete subsection reference.

- [timeouts](resources--nat_policy--reference--group-001.md#canonical-0332310313320102-1233203001320300-0120211031101101-2101310332223201-1311001110220122-3212202313031202-2103103310200032-1323023301223232): complete subsection reference.

<a id="canonical-0323201211032001-1200231011223310-0333233002200033-3201032012102230-1013101023132032-2131000303100020-1212200201313302-1020102021301321"></a>

## All schema paths — Property reference / 021313213121 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--nat_policy--reference--group-001.md#canonical-0332122232121112-3120320001021122-0022132310302112-0200103031322111-2013121321132211-3023112311101201-3122232330021021-3032333111003232) |
| `description` | [description](resources--nat_policy--reference--group-001.md#canonical-1010220332122020-0321123122201311-2030212131022231-1203100301032031-3122111222230222-0020133200233210-0202222110131212-2321301003032221) |
| `disable` | [disable](resources--nat_policy--reference--group-001.md#canonical-0131220133130202-3222231331100001-1100331323122211-0223301201332131-1203313132231101-0233232012320230-3021030010223300-3112203332020302) |
| `id` | [id](resources--nat_policy--reference--group-001.md#canonical-1320020300233233-3211123202312331-3132031102233022-2033120323023310-3233210120310303-2201310332300231-0323120031131032-1332130130312132) |
| `labels` | [labels](resources--nat_policy--reference--group-001.md#canonical-1131231210012200-1132130213331303-2221111011223212-2213320111100111-0200002230200110-1232313120100210-3301110313002120-2203231002331200) |
| `name` | [name](resources--nat_policy--reference--group-001.md#canonical-1230212310130013-1000200231313100-2330023002323030-0131331202221130-3303323023332222-1000232311210112-3001313113122102-0310131033012110) |
| `namespace` | [namespace](resources--nat_policy--reference--group-001.md#canonical-2200223211211102-2323121031031230-0111022031310222-1130031212113032-2133312210210022-0022132323331320-0111202321223301-3013322002130122) |
| `rules` | [rules](resources--nat_policy--reference--group-001.md#canonical-1213230032120130-1310201032022230-2210021020301121-2022301310033333-0312210220102203-0002001213333330-0032031233202320-3311221233122321) |
| `rules.action` | [rules.action](resources--nat_policy--reference--group-001.md#canonical-3132032130003123-1023102133012110-3300213111331201-1022301232311113-1131330123030300-2311000022133200-3023200003111300-2202023003310132) |
| `rules.action.dynamic` | [rules.action.dynamic](resources--nat_policy--reference--group-001.md#canonical-2022332100311212-1103130032012132-1022200020013013-3111222320303112-1202002300320222-1121211231120102-2201003221113102-1321231211112233) |
| `rules.action.dynamic.elastic_ips` | [rules.action.dynamic.elastic_ips](resources--nat_policy--reference--group-001.md#canonical-1130011120302103-3120010010110020-2013030332331321-3103333032111003-1023333310312123-3130022330100002-3132221123211201-0113223100230020) |
| `rules.action.dynamic.elastic_ips.refs` | [rules.action.dynamic.elastic_ips.refs](resources--nat_policy--reference--group-001.md#canonical-3230002133311231-2010102100120113-1113203231030110-1000220310213203-1312312323302111-2322303111211303-2222321233303012-0211011001102220) |
| `rules.action.dynamic.elastic_ips.refs.kind` | [rules.action.dynamic.elastic_ips.refs.kind](resources--nat_policy--reference--group-001.md#canonical-1113212230003330-0031300331133012-3200033020012322-2330123011203102-1121220200102020-2132103212231022-3021013323231102-1211130323313032) |
| `rules.action.dynamic.elastic_ips.refs.name` | [rules.action.dynamic.elastic_ips.refs.name](resources--nat_policy--reference--group-001.md#canonical-1213022111010220-1301111222110013-2203220211303222-0121212331203100-3020112210330330-2222103002310121-0230113100020023-2321002232022210) |
| `rules.action.dynamic.elastic_ips.refs.namespace` | [rules.action.dynamic.elastic_ips.refs.namespace](resources--nat_policy--reference--group-001.md#canonical-1223231033310010-0101033100210003-1031311333032123-3233030032213303-3033230320330231-0322233213333232-1313113131202003-3100202010010330) |
| `rules.action.dynamic.elastic_ips.refs.tenant` | [rules.action.dynamic.elastic_ips.refs.tenant](resources--nat_policy--reference--group-001.md#canonical-1021313213300012-3003022030012320-3231111020023113-0013332322111102-0221213030313230-0011300332103313-0320110330132121-3333032310131220) |
| `rules.action.dynamic.elastic_ips.refs.uid` | [rules.action.dynamic.elastic_ips.refs.uid](resources--nat_policy--reference--group-001.md#canonical-0322010231113103-0223301132323002-3020223311312302-3210323030201020-0303202333010310-2302032211020123-2030031311203021-2203213223023120) |
| `rules.action.dynamic.pools` | [rules.action.dynamic.pools](resources--nat_policy--reference--group-001.md#canonical-2333122032123311-2132123202201123-0002023121210131-2011022133331210-0213001020203103-0233132120101020-1000122210322003-1213012112103332) |
| `rules.action.dynamic.pools.prefixes` | [rules.action.dynamic.pools.prefixes](resources--nat_policy--reference--group-001.md#canonical-3131030022302230-2321000123112123-0002113101103322-2103131111001120-1130201203233302-2221232123033023-0031131321022211-2213110303201200) |
| `rules.action.virtual_cidr` | [rules.action.virtual_cidr](resources--nat_policy--reference--group-001.md#canonical-1203200121312221-0103022130321003-0132020131202300-0113311020012122-3230130211100231-2303030333302100-2023002312200303-2021223113202212) |
| `rules.cloud_connect` | [rules.cloud_connect](resources--nat_policy--reference--group-001.md#canonical-1111312031133000-0013212212212233-3310210223200331-0313021033103133-1121003013233000-3300330003220212-2313222100113121-3030310310330302) |
| `rules.cloud_connect.refs` | [rules.cloud_connect.refs](resources--nat_policy--reference--group-001.md#canonical-1311001123333020-0003323000131120-2023112330112021-0023023323332100-1220122112322101-1232131110302212-1102210202213233-0102000022110010) |
| `rules.cloud_connect.refs.kind` | [rules.cloud_connect.refs.kind](resources--nat_policy--reference--group-001.md#canonical-2310102002130212-0032001312002113-1222232213203222-2312302200000113-1221231333123301-0330003303331132-3123000111103001-1220303011302312) |
| `rules.cloud_connect.refs.name` | [rules.cloud_connect.refs.name](resources--nat_policy--reference--group-001.md#canonical-0002001231223021-0311200313203211-3033330320232303-1000132203003103-3211023120123330-2200302330332302-3013202023102211-0012002000331211) |
| `rules.cloud_connect.refs.namespace` | [rules.cloud_connect.refs.namespace](resources--nat_policy--reference--group-001.md#canonical-0021132233010122-0230110310230313-0220002303313232-1132031312311203-1130112010202133-3333330212333320-0320000000231010-2322201013321201) |
| `rules.cloud_connect.refs.tenant` | [rules.cloud_connect.refs.tenant](resources--nat_policy--reference--group-001.md#canonical-2101113310233323-1233020220231013-0310133023202222-2212003223032103-0001031321012323-2220110302122003-2111301200231310-3331223113130300) |
| `rules.cloud_connect.refs.uid` | [rules.cloud_connect.refs.uid](resources--nat_policy--reference--group-001.md#canonical-2110320100020102-3001000210330313-2302002112330200-1030122121200203-0110201032331233-2311122133232201-1311103031111301-2010301103300010) |
| `rules.criteria` | [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-1331311020201203-1012203013201322-2210101102101012-3133221110100000-3013310133201020-3030000201200230-0222331331001031-0020120020310023) |
| `rules.criteria.any` | [rules.criteria.any](resources--nat_policy--reference--group-001.md#canonical-0133132230022201-1030232020331001-3211112113202332-0002113111111302-1022033212300333-0203031212322311-3313321112230031-2031023113312311) |
| `rules.criteria.destination_cidr` | [rules.criteria.destination_cidr](resources--nat_policy--reference--group-001.md#canonical-3023321213011103-1213233303032210-2233300212311212-3320113031112023-2033333213311012-2003101212100201-3111203021232020-0032111201121202) |
| `rules.criteria.icmp` | [rules.criteria.icmp](resources--nat_policy--reference--group-001.md#canonical-3221222003032231-1132232102223023-2000112020303312-1033021301313112-0013331010212211-3133120103101232-3302220011012303-2302312230101322) |
| `rules.criteria.site_local_inside_network` | [rules.criteria.site_local_inside_network](resources--nat_policy--reference--group-001.md#canonical-0102112121222011-0231101310300032-3213100022112212-3101331233111122-2201102212112202-0133120220201231-2131021002023221-2023211013030202) |
| `rules.criteria.site_local_network` | [rules.criteria.site_local_network](resources--nat_policy--reference--group-001.md#canonical-3203013202301233-0012223313333001-2103302131222332-3321232030003221-0021233102331120-1331101223312330-1103022303121133-3020333311200333) |
| `rules.criteria.source_cidr` | [rules.criteria.source_cidr](resources--nat_policy--reference--group-001.md#canonical-2231203023123131-1330030301021101-0031120201310102-2321322212011330-1122300221231311-0102013232323112-1221000200303000-2001223222231213) |
| `rules.criteria.tcp` | [rules.criteria.tcp](resources--nat_policy--reference--group-001.md#canonical-0210221123330011-1131202312010100-3130302232123233-3233200210032020-2230300221202331-1201320221003212-1310130022312231-3032312002133220) |
| `rules.criteria.tcp.destination_port` | [rules.criteria.tcp.destination_port](resources--nat_policy--reference--group-001.md#canonical-1033332331020031-0202200021331310-2012033200300103-1303023223113200-1331232010021331-2222123330021122-1311113330012031-2323311110213003) |
| `rules.criteria.tcp.destination_port.no_port_match` | [rules.criteria.tcp.destination_port.no_port_match](resources--nat_policy--reference--group-001.md#canonical-3111012322132130-2222201223111313-3000232132032130-0123302322332310-0002002231021212-0121232200231023-3022223231323332-3211323220100311) |
| `rules.criteria.tcp.destination_port.port` | [rules.criteria.tcp.destination_port.port](resources--nat_policy--reference--group-001.md#canonical-3102213323303103-3323102302201301-1222122300131211-0003201122113001-1222030010211321-3002200103231023-0022011301012311-2101313202012331) |
| `rules.criteria.tcp.destination_port.port_ranges` | [rules.criteria.tcp.destination_port.port_ranges](resources--nat_policy--reference--group-001.md#canonical-0221111332133021-2201103330202120-1000112130111101-3303212130333302-0222023233013310-1022130310211001-1220233102320011-2100221032313313) |
| `rules.criteria.tcp.source_port` | [rules.criteria.tcp.source_port](resources--nat_policy--reference--group-001.md#canonical-3102320021110001-1201320132120300-1120310103221333-1231330312020013-3122022222320220-0330003320203103-1003233333011223-3001322013132011) |
| `rules.criteria.tcp.source_port.no_port_match` | [rules.criteria.tcp.source_port.no_port_match](resources--nat_policy--reference--group-001.md#canonical-3322230212210221-3330012023122231-2300303232201022-2023021000112232-3132120132003331-2100131111000333-1022303130110101-3303233232010230) |
| `rules.criteria.tcp.source_port.port` | [rules.criteria.tcp.source_port.port](resources--nat_policy--reference--group-001.md#canonical-0133301301033003-1023022100221302-0013322012103223-0113111211031123-2213120133011032-1120033203322020-2233322322320033-3010101320233231) |
| `rules.criteria.tcp.source_port.port_ranges` | [rules.criteria.tcp.source_port.port_ranges](resources--nat_policy--reference--group-001.md#canonical-0032101111133113-1321021031323302-2010132000023113-2020101113020020-1002222223022003-2003212200313320-3203322113122223-3233002331011222) |
| `rules.criteria.udp` | [rules.criteria.udp](resources--nat_policy--reference--group-001.md#canonical-0223102112002322-3313133203331103-2212323000211302-0333312230312003-0110021322221312-3131022233022032-0331203011231233-0210132001313232) |
| `rules.criteria.udp.destination_port` | [rules.criteria.udp.destination_port](resources--nat_policy--reference--group-001.md#canonical-2213211112020101-0112332022023200-0322303202310001-2233103113201101-0230303232112221-1302020330200302-3032031232203023-2223321213032321) |
| `rules.criteria.udp.destination_port.no_port_match` | [rules.criteria.udp.destination_port.no_port_match](resources--nat_policy--reference--group-001.md#canonical-1033332131110232-2101313200011232-0023301223220131-2300101210203113-0322312232211013-0231003102323020-1223210322302121-0100013301202002) |
| `rules.criteria.udp.destination_port.port` | [rules.criteria.udp.destination_port.port](resources--nat_policy--reference--group-001.md#canonical-0332001133021312-3331102313121022-3312030330113303-0001303331203230-2001132022101300-0231100101123200-1322213201032022-1131230000112111) |
| `rules.criteria.udp.destination_port.port_ranges` | [rules.criteria.udp.destination_port.port_ranges](resources--nat_policy--reference--group-001.md#canonical-0111201230022022-0120011011222131-2232111111030223-0233003000201122-1202300131310011-1003112102100203-2311100312222110-3221032011233011) |
| `rules.criteria.udp.source_port` | [rules.criteria.udp.source_port](resources--nat_policy--reference--group-001.md#canonical-2310211313233212-3112103031332033-2111232310212331-1122113310203223-2231232121120030-0320023321200210-3133332011133133-3032323233322001) |
| `rules.criteria.udp.source_port.no_port_match` | [rules.criteria.udp.source_port.no_port_match](resources--nat_policy--reference--group-001.md#canonical-3210010222231233-1200201011123322-1221321210213120-0210300101101300-3010300031011202-0113321310010002-2302111320013233-2321201011211111) |
| `rules.criteria.udp.source_port.port` | [rules.criteria.udp.source_port.port](resources--nat_policy--reference--group-001.md#canonical-3231201210302201-2002123223123203-2013101332111132-2100121303103200-3122322202021022-2311323312021003-0302120113311131-2132030231332103) |
| `rules.criteria.udp.source_port.port_ranges` | [rules.criteria.udp.source_port.port_ranges](resources--nat_policy--reference--group-001.md#canonical-1231203121121100-0232313332300023-0133123120122203-3112033212023210-2330220333001021-0031302131203323-2133333030023003-0031020003030311) |
| `rules.disable_spec` | [rules.disable_spec](resources--nat_policy--reference--group-001.md#canonical-1230201131132230-0312113000231211-0321211220110203-0100033112202320-1331003310133122-1020132131022322-3001232110133020-3012213022202102) |
| `rules.enable` | [rules.enable](resources--nat_policy--reference--group-001.md#canonical-2111322330111130-1222222221011221-0003131112130332-3013311202332230-2022213322300203-2102101223003130-1000103202223021-3012301120302301) |
| `rules.name` | [rules.name](resources--nat_policy--reference--group-001.md#canonical-2012021302001111-1101030211031331-2322213001200132-3213300212322330-3231020300331200-0132101030012113-2123322132330211-0110320301221133) |
| `rules.node_interface` | [rules.node_interface](resources--nat_policy--reference--group-001.md#canonical-0320333020103122-0220031302222200-1020233221300333-2130230333210031-1020012322200323-1001011003332000-2211233121021101-2102131210002123) |
| `rules.node_interface.list` | [rules.node_interface.list](resources--nat_policy--reference--group-001.md#canonical-0000001202123023-1310210330222313-0121030022302300-1222001230113003-2303331012003003-3323213300222302-1312103320012133-1201111213123010) |
| `rules.node_interface.list.interface` | [rules.node_interface.list.interface](resources--nat_policy--reference--group-001.md#canonical-2312213120023023-2021323303020323-0303202330123110-3023001213103110-3130202321102212-2133201012122011-3313023030311301-3332003320212312) |
| `rules.node_interface.list.interface.kind` | [rules.node_interface.list.interface.kind](resources--nat_policy--reference--group-001.md#canonical-2222101030310100-0212331213312030-2033320312022233-2013033230131112-1002001010322200-0323311030020110-2032033012112000-0313010330301112) |
| `rules.node_interface.list.interface.name` | [rules.node_interface.list.interface.name](resources--nat_policy--reference--group-001.md#canonical-0333013321212002-2101330123002032-1133303211220230-2132233010033023-1133023011031311-2302321330012231-3312302110020130-2030100000233221) |
| `rules.node_interface.list.interface.namespace` | [rules.node_interface.list.interface.namespace](resources--nat_policy--reference--group-001.md#canonical-3120132331302202-0103012223330103-2313110013111210-1131213302211220-1102202303023310-2230230002331202-3013213000023011-3200032032102033) |
| `rules.node_interface.list.interface.tenant` | [rules.node_interface.list.interface.tenant](resources--nat_policy--reference--group-001.md#canonical-2331013322323300-2222203123100003-0220033111220221-2022223020203231-1300313033132101-2300210220012021-3230102222300000-2230302322220331) |
| `rules.node_interface.list.interface.uid` | [rules.node_interface.list.interface.uid](resources--nat_policy--reference--group-001.md#canonical-0312221311000331-0200023020230313-1103211102210121-3303132033032323-2320113100121223-2321110031023112-0302221123121123-2102120312002222) |
| `rules.node_interface.list.node` | [rules.node_interface.list.node](resources--nat_policy--reference--group-001.md#canonical-2322312222011333-3303132232310100-2000012231230003-2322012002210320-2200131030321012-2003202031311323-0111132332003132-3300312233110021) |
| `rules.segment` | [rules.segment](resources--nat_policy--reference--group-001.md#canonical-1002003323223203-1013303333103203-0031010010033300-2212110101202121-1300320122201213-3231233212331122-1212101101011202-0220322331222011) |
| `rules.segment.refs` | [rules.segment.refs](resources--nat_policy--reference--group-001.md#canonical-0332133133132122-0101200311001100-3022303322210031-2023223203302111-1223231110231233-1122103220022103-1112213110323303-3002232330120213) |
| `rules.segment.refs.kind` | [rules.segment.refs.kind](resources--nat_policy--reference--group-001.md#canonical-1303333120301303-1230231030131121-1323200123002233-1120323300233211-3321203333232330-1012221222320010-2100103202032023-2003010223202121) |
| `rules.segment.refs.name` | [rules.segment.refs.name](resources--nat_policy--reference--group-001.md#canonical-1202300030210222-3222310113200332-1031030123210223-0100010111033233-0130003310220133-1322233313100013-1211001222201130-1210221222301301) |
| `rules.segment.refs.namespace` | [rules.segment.refs.namespace](resources--nat_policy--reference--group-001.md#canonical-3312203201010010-0133302102133231-3012310312011233-3001111123030112-0111303101302333-2223310103101303-2130212031111103-0231320023211331) |
| `rules.segment.refs.tenant` | [rules.segment.refs.tenant](resources--nat_policy--reference--group-001.md#canonical-3011132330120302-2102022123210132-3030202013331002-0202120120300120-0331313100202031-2113223020003331-2233122103023022-3003303113201233) |
| `rules.segment.refs.uid` | [rules.segment.refs.uid](resources--nat_policy--reference--group-001.md#canonical-2002001331321210-1122130110101100-3113012002331321-3302112101133133-0210303323122200-0330303320212133-2323212003333032-1010301030333113) |
| `rules.virtual_network` | [rules.virtual_network](resources--nat_policy--reference--group-001.md#canonical-1002012311021032-1021333321012032-2230202231321302-0022021101022210-3232101102203121-0320102230102002-1112310123131332-1033233021312231) |
| `rules.virtual_network.refs` | [rules.virtual_network.refs](resources--nat_policy--reference--group-001.md#canonical-2222011321033210-1120130020123120-3002112101023023-0130313030222231-1312202103120313-0300002100120101-1332132132330120-2032022112011330) |
| `rules.virtual_network.refs.kind` | [rules.virtual_network.refs.kind](resources--nat_policy--reference--group-001.md#canonical-1113132123121211-0303111122032020-1131320332210001-1303103003210211-2011332210303312-1131021320201012-0012132321001002-1233232232110112) |
| `rules.virtual_network.refs.name` | [rules.virtual_network.refs.name](resources--nat_policy--reference--group-001.md#canonical-1200312121220021-3201311310302112-2321232000130330-1233101203002220-3313233313023000-2131111221101323-1212333230003323-3312102000110322) |
| `rules.virtual_network.refs.namespace` | [rules.virtual_network.refs.namespace](resources--nat_policy--reference--group-001.md#canonical-1331302123001313-2303231032113013-1122311111311323-2021312322020330-2233233220130203-2221030231012300-2221320133320102-0122300011113033) |
| `rules.virtual_network.refs.tenant` | [rules.virtual_network.refs.tenant](resources--nat_policy--reference--group-001.md#canonical-2311013331330322-2231300312021311-3111210002331320-0312130121300121-3121110213001321-1331032122033030-0300333330223232-1302000202031230) |
| `rules.virtual_network.refs.uid` | [rules.virtual_network.refs.uid](resources--nat_policy--reference--group-001.md#canonical-0002211022010032-1131320322310020-3121012302202322-2332101220330210-0202213230110030-0210031020123233-3323010002321110-3032323123010312) |
| `site` | [site](resources--nat_policy--reference--group-001.md#canonical-2030103300120200-0101122100203031-3033132321322000-3002330010012232-2100220201321330-2233122102122111-2032023000031113-2222130223311331) |
| `site.refs` | [site.refs](resources--nat_policy--reference--group-001.md#canonical-2220012132012310-2211132013030100-1111123232003233-0211301020200002-1023002332331212-1122311130321300-1030020213330233-3132213231110303) |
| `site.refs.kind` | [site.refs.kind](resources--nat_policy--reference--group-001.md#canonical-0311112013112312-3310001003100122-0001311320213111-1023133301200203-1300221330233123-1020023300202322-1303013001310120-0001101011110022) |
| `site.refs.name` | [site.refs.name](resources--nat_policy--reference--group-001.md#canonical-1320033122131000-1021012033133222-3031313330001310-0210311211202220-3220231230001230-1013003300131111-0312210212213300-1203110101233022) |
| `site.refs.namespace` | [site.refs.namespace](resources--nat_policy--reference--group-001.md#canonical-1113031312331003-3333331030333330-2322101131312222-0023300231132111-3312121000212110-0013120221202333-0131201033322001-0103200013031212) |
| `site.refs.tenant` | [site.refs.tenant](resources--nat_policy--reference--group-001.md#canonical-2220311222103222-2321312313102321-1313310223012110-0101212021133213-1231002112033112-0011322330021323-0301322032021330-2301202301330123) |
| `site.refs.uid` | [site.refs.uid](resources--nat_policy--reference--group-001.md#canonical-3102123331211033-2312333333213302-0002303203202010-1122100202001232-3030203013103311-0120032230111011-2011230100032000-0120131002301132) |
| `timeouts` | [timeouts](resources--nat_policy--reference--group-001.md#canonical-1123000232032010-1111231111122300-2030311322101302-2321002131302202-2332021102231101-3221031131021123-0022222103212310-0102113200200102) |
| `timeouts.create` | [timeouts.create](resources--nat_policy--reference--group-001.md#canonical-1011322020203030-1113033222120201-3232102302110222-1002103103033312-1231230320020332-2002030000002002-1230230122333103-1101312320012122) |
| `timeouts.delete` | [timeouts.delete](resources--nat_policy--reference--group-001.md#canonical-1310323321110022-1330102030300120-2223201133211122-1301322020122320-2033222302133133-2233331220303233-0230122100111232-3330010110012132) |
| `timeouts.read` | [timeouts.read](resources--nat_policy--reference--group-001.md#canonical-0110023222221033-0002200102110023-3030031311231003-1203100202221020-3212302031021013-2310003322212122-1311330311302002-0133113213112023) |
| `timeouts.update` | [timeouts.update](resources--nat_policy--reference--group-001.md#canonical-0212231123312211-1221133020303323-3112212231232233-0322320320030013-2332112203022230-1110023101210220-2021110323000211-3213310020303212) |

<a id="canonical-1333230113122223-3230012132013031-2003120123133233-3130203112110222-2101301031212110-3202132121102033-3222221122112210-3121212333233311"></a>

## Next pages — Property reference / 021313213121 / 12

- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [site](resources--nat_policy--reference--group-001.md#canonical-3001223222011220-0003203113332212-3332203303130003-1310223301012333-2223122022201100-2321110020002221-0120100011321320-2020123030023230)
- [timeouts](resources--nat_policy--reference--group-001.md#canonical-0332310313320102-1233203001320300-0120211031101101-2101310332223201-1311001110220122-3212202313031202-2103103310200032-1323023301223232)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)

<a id="canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112322003023233-0002210111230221-2210110331032120-3031001232123312-3211302123210333-2132300011121333-1233033001232003-1033131330112121"></a>

## rules — rules / 213202020200 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- rules

<a id="canonical-1213230032120130-1310201032022230-2210021020301121-2022301310033333-0312210220102203-0002001213333330-0032031233202320-3311221233122321"></a>

Type: `"object"`. list nested block, Optional.

List of rules to apply under the NAT Policy. Rule that matches first would be applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("cloud_connect",
    "node_interface"),
  validators.ConflictingListObjectAttributes("cloud_connect",
    "segment"),
  validators.ConflictingListObjectAttributes("cloud_connect",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("disable_spec",
    "enable"),
  validators.ConflictingListObjectAttributes("node_interface",
    "segment"),
  validators.ConflictingListObjectAttributes("node_interface",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("segment",
    "virtual_network")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-1110330220330002-0002102023000011-2113101333300330-1112123020013100-2200012213113032-3330001212202301-3120303302202032-1330212032230321"></a>

## Direct properties — rules / 213202020200 / 3

- [action](resources--nat_policy--reference--group-001.md#canonical-1122221021030123-2131233221312132-0230001013023100-0130321320322011-2220032330303113-0023002023013103-3130110202211032-1022000313331212): complete subsection reference.

- [cloud_connect](resources--nat_policy--reference--group-001.md#canonical-1003100233302122-0331210123231312-3101222012030102-0011000230001133-0230002101033103-3330233311000231-2133020120103333-2030021301103211): complete subsection reference.

- [criteria](resources--nat_policy--reference--group-001.md#canonical-0211202132100101-0101202032220222-0011011112111300-3220020012220110-0202012111300120-3331002203310331-0223221133023011-2213211332320022): complete subsection reference.

- [disable_spec](resources--nat_policy--reference--group-001.md#canonical-2333111032320031-2222102332001330-1220210111201123-1011122013230003-3231123330013221-3103132010233220-0010221330011032-0200130112321100): complete subsection reference.

- [enable](resources--nat_policy--reference--group-001.md#canonical-3301200022202232-1013113223132333-1111013013320103-3201130310232122-1233023330321002-0020013000032003-0121010123113111-2203023302031312): complete subsection reference.

<a id="canonical-2012021302001111-1101030211031331-2322213001200132-3213300212322330-3231020300331200-0132101030012113-2123322132330211-0110320301221133"></a>

<a id="canonical-0002101300032200-2210110130110313-3002330201003200-0013303323120300-2020123122132120-1212312103303011-1223120030132032-0123301022312001"></a>

## name property — rules / 213202020200 / 4

Type: `"string"`. Optional.

Name. Name of the Rule.

Upstream description:

Name of the Rule.

Provider validators and defaults (from schema source):

```go
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

- [node_interface](resources--nat_policy--reference--group-001.md#canonical-0003122032201330-0222033131003300-3123021011023102-3103311200332223-1132130323200320-1223032020223203-0130321232310300-0121313320300220): complete subsection reference.

- [segment](resources--nat_policy--reference--group-001.md#canonical-2000331302003211-2112122330312222-0231223203303123-2331332301321123-1023103133311320-1120102103133013-3031202021101232-2321033122322230): complete subsection reference.

- [virtual_network](resources--nat_policy--reference--group-001.md#canonical-1203111110012112-3101113022030330-0033312021330011-0221130000132103-2001003103030300-2130101120332322-3210203132100203-1033032202202201): complete subsection reference.

<a id="canonical-3211222023023301-3120120321102003-1030232122201231-2203112002102330-3232102232202312-1023302100121331-1203033313233312-0231232303130211"></a>

## Next pages — rules / 213202020200 / 5

- [rules.action](resources--nat_policy--reference--group-001.md#canonical-1122221021030123-2131233221312132-0230001013023100-0130321320322011-2220032330303113-0023002023013103-3130110202211032-1022000313331212)
- [rules.cloud_connect](resources--nat_policy--reference--group-001.md#canonical-1003100233302122-0331210123231312-3101222012030102-0011000230001133-0230002101033103-3330233311000231-2133020120103333-2030021301103211)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-0211202132100101-0101202032220222-0011011112111300-3220020012220110-0202012111300120-3331002203310331-0223221133023011-2213211332320022)
- [rules.disable_spec](resources--nat_policy--reference--group-001.md#canonical-2333111032320031-2222102332001330-1220210111201123-1011122013230003-3231123330013221-3103132010233220-0010221330011032-0200130112321100)
- [rules.enable](resources--nat_policy--reference--group-001.md#canonical-3301200022202232-1013113223132333-1111013013320103-3201130310232122-1233023330321002-0020013000032003-0121010123113111-2203023302031312)
- [rules.node_interface](resources--nat_policy--reference--group-001.md#canonical-0003122032201330-0222033131003300-3123021011023102-3103311200332223-1132130323200320-1223032020223203-0130321232310300-0121313320300220)
- [rules.segment](resources--nat_policy--reference--group-001.md#canonical-2000331302003211-2112122330312222-0231223203303123-2331332301321123-1023103133311320-1120102103133013-3031202021101232-2321033122322230)
- [rules.virtual_network](resources--nat_policy--reference--group-001.md#canonical-1203111110012112-3101113022030330-0033312021330011-0221130000132103-2001003103030300-2130101120332322-3210203132100203-1033032202202201)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)

<a id="canonical-1122221021030123-2131233221312132-0230001013023100-0130321320322011-2220032330303113-0023002023013103-3130110202211032-1022000313331212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0201100313020023-3201003211010102-0111123330023103-1210102020111132-3110102233113300-0122212120232223-2100233020233110-1202232123212312"></a>

## rules.action — action / 333032021203 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- rules.action

<a id="canonical-3132032130003123-1023102133012110-3300213111331201-1022301232311113-1131330123030300-2311000022133200-3023200003111300-2202023003310132"></a>

Type: `"object"`. single nested block, Optional.

Action to apply on the packet if the NAT rule is applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dynamic",
    "virtual_cidr")}
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
  "x-ves-oneof-field-source_nat_choice": "[\"dynamic\",\"virtual_cidr\"]"
}
```

Terraform syntax:

```terraform
action {
  # Configure direct properties listed below.
}
```

<a id="canonical-2200223020031102-2003023010210112-2000033312232310-1103312310021313-1133020310111233-1103321120301201-2200023121012120-1202033323101032"></a>

## Direct properties — action / 333032021203 / 3

- [dynamic](resources--nat_policy--reference--group-001.md#canonical-0221213332001013-1211231020320202-2110123331210011-3311233300323212-0323332201123333-3023103321300300-1013003221001113-1100103221120012): complete subsection reference.

<a id="canonical-1203200121312221-0103022130321003-0132020131202300-0113311020012122-3230130211100231-2303030333302100-2023002312200303-2021223113202212"></a>

<a id="canonical-3200000311200211-0333131111011303-2031311120112232-0120300233010332-2203322101100013-3231301010011121-1302101223230322-3130311100310101"></a>

## virtual_cidr property — action / 333032021203 / 4

Type: `"string"`. Optional.

Exclusive with \[dynamic\] Virtual Subnet NAT is static NAT that does a one-to-one translation
between the real source IP CIDR in the policy and the virtual CIDR in a bidirectional fashion. The
range of the real CIDR and virtual CIDRs should be the same (e.g. If the real CIDR has the CIDR..

Upstream description:

Exclusive with \[dynamic\] Virtual Subnet NAT is static NAT that does a one-to-one translation
between the real source IP CIDR in the policy and the virtual CIDR in a bidirectional fashion. The
range of the real CIDR and virtual CIDRs should be the same (e.g. If the real CIDR has the CIDR
192.0.2.0/24, the virtual CIDR has 100.100.100.0/24.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.CIDRValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

<a id="canonical-0323220230100132-1221221110121320-3032112031201110-2312110010310133-2003201003322002-2131132120022313-1323031310233113-3100031002132311"></a>

## Next pages — action / 333032021203 / 5

- [rules.action.dynamic](resources--nat_policy--reference--group-001.md#canonical-0221213332001013-1211231020320202-2110123331210011-3311233300323212-0323332201123333-3023103321300300-1013003221001113-1100103221120012)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)

<a id="canonical-0221213332001013-1211231020320202-2110123331210011-3311233300323212-0323332201123333-3023103321300300-1013003221001113-1100103221120012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223010222221211-0300002313231300-0100122212213201-3013133023311223-0220310030331320-3132311222021210-1220321113131232-3130000033301303"></a>

## rules.action.dynamic — dynamic / 120000202002 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [rules.action](resources--nat_policy--reference--group-001.md#canonical-1122221021030123-2131233221312132-0230001013023100-0130321320322011-2220032330303113-0023002023013103-3130110202211032-1022000313331212)
- rules.action.dynamic

<a id="canonical-2022332100311212-1103130032012132-1022200020013013-3111222320303112-1202002300320222-1121211231120102-2201003221113102-1321231211112233"></a>

Type: `"object"`. single nested block, Optional.

Dynamic Pool. Dynamic Pool Configuration.

Upstream description:

Dynamic Pool Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("elastic_ips",
    "pools")}
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
  "x-ves-oneof-field-pool_choice": "[\"elastic_ips\",\"pools\"]"
}
```

Terraform syntax:

```terraform
dynamic {
  # Configure direct properties listed below.
}
```

<a id="canonical-2211130103031213-2223022213311023-0122000231010212-2132122023013131-3203202203021132-0300033211102222-1223201120120013-2012033022110010"></a>

## Direct properties — dynamic / 120000202002 / 3

- [elastic_ips](resources--nat_policy--reference--group-001.md#canonical-1320332012230123-1200023330031312-3010211321230103-0312220130203323-1021023110233100-1332012000310330-0322320302003223-2032133322302003): complete subsection reference.

- [pools](resources--nat_policy--reference--group-001.md#canonical-3101330020232121-1300120130322033-1120200032101230-0031011222111032-3100321032313012-3300112132101200-0333021110213112-2300123021320102): complete subsection reference.

<a id="canonical-0112132120003310-2333010012223023-1201231303300213-1111311222221331-3002332323121021-3122212310303112-0000331020030232-2201231033310201"></a>

## Next pages — dynamic / 120000202002 / 4

- [rules.action.dynamic.elastic_ips](resources--nat_policy--reference--group-001.md#canonical-1320332012230123-1200023330031312-3010211321230103-0312220130203323-1021023110233100-1332012000310330-0322320302003223-2032133322302003)
- [rules.action.dynamic.pools](resources--nat_policy--reference--group-001.md#canonical-3101330020232121-1300120130322033-1120200032101230-0031011222111032-3100321032313012-3300112132101200-0333021110213112-2300123021320102)
- [rules.action](resources--nat_policy--reference--group-001.md#canonical-1122221021030123-2131233221312132-0230001013023100-0130321320322011-2220032330303113-0023002023013103-3130110202211032-1022000313331212)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)

<a id="canonical-1320332012230123-1200023330031312-3010211321230103-0312220130203323-1021023110233100-1332012000310330-0322320302003223-2032133322302003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0311221232311232-1211320233102200-0221213130013211-0013102121123320-3313330213313322-1113123330303113-0133030110110301-0022031321222101"></a>

## rules.action.dynamic.elastic_ips — elastic_ips / 131013211302 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [rules.action](resources--nat_policy--reference--group-001.md#canonical-1122221021030123-2131233221312132-0230001013023100-0130321320322011-2220032330303113-0023002023013103-3130110202211032-1022000313331212)
- [rules.action.dynamic](resources--nat_policy--reference--group-001.md#canonical-0221213332001013-1211231020320202-2110123331210011-3311233300323212-0323332201123333-3023103321300300-1013003221001113-1100103221120012)
- rules.action.dynamic.elastic_ips

<a id="canonical-1130011120302103-3120010010110020-2013030332331321-3103333032111003-1023333310312123-3130022330100002-3132221123211201-0113223100230020"></a>

Type: `"object"`. single nested block, Optional.

List of references to Cloud Elastic IP Object.

Provider validators and defaults (from schema source):

```go
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
elastic_ips {
  # Configure direct properties listed below.
}
```

<a id="canonical-2203012133003200-1100130320202111-3120210011302010-0120132222131103-1012030320121031-2222313020021102-0122201100233030-2120132213332033"></a>

## Direct properties — elastic_ips / 131013211302 / 3

- [refs](resources--nat_policy--reference--group-001.md#canonical-3021303223113021-3220020210100310-0021033122120022-2201011001330101-1320221000010220-0033132112200311-3031202133011201-1033121211131023): complete subsection reference.

<a id="canonical-1031332122032320-0212300001331013-0203023313031303-3013031032332230-1011202330310013-2131230311301110-0311223210223322-2333103212332003"></a>

## Next pages — elastic_ips / 131013211302 / 4

- [rules.action.dynamic.elastic_ips.refs](resources--nat_policy--reference--group-001.md#canonical-3021303223113021-3220020210100310-0021033122120022-2201011001330101-1320221000010220-0033132112200311-3031202133011201-1033121211131023)
- [rules.action.dynamic](resources--nat_policy--reference--group-001.md#canonical-0221213332001013-1211231020320202-2110123331210011-3311233300323212-0323332201123333-3023103321300300-1013003221001113-1100103221120012)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)

<a id="canonical-3021303223113021-3220020210100310-0021033122120022-2201011001330101-1320221000010220-0033132112200311-3031202133011201-1033121211131023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121331233303003-0313310010131030-1030121221321220-2310130222211220-1202021012231201-3320032011333323-1121221301222111-2201202100231100"></a>

## rules.action.dynamic.elastic_ips.refs — refs / 120232130023 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [rules.action](resources--nat_policy--reference--group-001.md#canonical-1122221021030123-2131233221312132-0230001013023100-0130321320322011-2220032330303113-0023002023013103-3130110202211032-1022000313331212)
- [rules.action.dynamic](resources--nat_policy--reference--group-001.md#canonical-0221213332001013-1211231020320202-2110123331210011-3311233300323212-0323332201123333-3023103321300300-1013003221001113-1100103221120012)
- [rules.action.dynamic.elastic_ips](resources--nat_policy--reference--group-001.md#canonical-1320332012230123-1200023330031312-3010211321230103-0312220130203323-1021023110233100-1332012000310330-0322320302003223-2032133322302003)
- rules.action.dynamic.elastic_ips.refs

<a id="canonical-3230002133311231-2010102100120113-1113203231030110-1000220310213203-1312312323302111-2322303111211303-2222321233303012-0211011001102220"></a>

Type: `"object"`. list nested block, Optional.

Reference to one or more cloud elastic IP objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
refs {
  # Configure direct properties listed below.
}
```

<a id="canonical-0210213302111223-1010211022303300-1232220100013203-0232003230211023-2000222333132210-0013112003302233-2121222102032023-3200133122311311"></a>

## Direct properties — refs / 120232130023 / 3

<a id="canonical-1113212230003330-0031300331133012-3200033020012322-2330123011203102-1121220200102020-2132103212231022-3021013323231102-1211130323313032"></a>

<a id="canonical-1001232232012223-0332002210013210-0210022123021222-2121202333320200-1202223323213000-3203222313103110-1100323313331210-2331020013012200"></a>

## kind property — refs / 120232130023 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-1213022111010220-1301111222110013-2203220211303222-0121212331203100-3020112210330330-2222103002310121-0230113100020023-2321002232022210"></a>

<a id="canonical-1123313313311231-3133321323313120-0330021200231022-1222000322113310-0023220313111032-2020233111003022-2203033313333320-0011222332233333"></a>

## name property — refs / 120232130023 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-1223231033310010-0101033100210003-1031311333032123-3233030032213303-3033230320330231-0322233213333232-1313113131202003-3100202010010330"></a>

<a id="canonical-2203311331131103-0010220320213111-1311003020300101-1003231012020121-1312022201231123-1021312211123013-1121331310212230-3003022222321310"></a>

## namespace property — refs / 120232130023 / 6

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

<a id="canonical-1021313213300012-3003022030012320-3231111020023113-0013332322111102-0221213030313230-0011300332103313-0320110330132121-3333032310131220"></a>

<a id="canonical-3112113232030231-3021001330321312-2131022113323301-2231021332321022-3013222112322223-3113300211100220-1310213103213030-2001230001133102"></a>

## tenant property — refs / 120232130023 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-0322010231113103-0223301132323002-3020223311312302-3210323030201020-0303202333010310-2302032211020123-2030031311203021-2203213223023120"></a>

<a id="canonical-2320002033000120-3032120301230313-1002111332132312-1131333332122111-3213333322122213-1033111211311222-3112013312211221-0133130010210103"></a>

## uid property — refs / 120232130023 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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

<a id="canonical-1233330203322110-1211221321112300-0201022103311223-3320332200311300-3212312221323303-1233222102210331-2332000121300303-3301100120132133"></a>

## Next pages — refs / 120232130023 / 9

- [rules.action.dynamic.elastic_ips](resources--nat_policy--reference--group-001.md#canonical-1320332012230123-1200023330031312-3010211321230103-0312220130203323-1021023110233100-1332012000310330-0322320302003223-2032133322302003)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)

<a id="canonical-3101330020232121-1300120130322033-1120200032101230-0031011222111032-3100321032313012-3300112132101200-0333021110213112-2300123021320102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223202332231000-1302321230010110-0100300031230002-1213103032000333-0231030033233302-0003220330303003-3201013111023222-0031231021002100"></a>

## rules.action.dynamic.pools — pools / 130233320303 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [rules.action](resources--nat_policy--reference--group-001.md#canonical-1122221021030123-2131233221312132-0230001013023100-0130321320322011-2220032330303113-0023002023013103-3130110202211032-1022000313331212)
- [rules.action.dynamic](resources--nat_policy--reference--group-001.md#canonical-0221213332001013-1211231020320202-2110123331210011-3311233300323212-0323332201123333-3023103321300300-1013003221001113-1100103221120012)
- rules.action.dynamic.pools

<a id="canonical-2333122032123311-2132123202201123-0002023121210131-2011022133331210-0213001020203103-0233132120101020-1000122210322003-1213012112103332"></a>

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
pools {
  # Configure direct properties listed below.
}
```

<a id="canonical-3211132103002221-0030003020020223-0221130333312223-0200222010013332-3320000300321100-3200112100023312-3132112031011110-0130311313322002"></a>

## Direct properties — pools / 130233320303 / 3

<a id="canonical-3131030022302230-2321000123112123-0002113101103322-2103131111001120-1130201203233302-2221232123033023-0031131321022211-2213110303201200"></a>

<a id="canonical-2203322330301201-3211210320213303-1211321113012010-1122012122001312-0331002322223100-3012131133202322-0333331131100121-0201201133330131"></a>

## prefixes property — pools / 130233320303 / 4

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0113023103001120-1322023100332031-0111320101031322-2000230331131332-1310230202013103-2001230032131121-1032003102132221-1311230200133321"></a>

## Next pages — pools / 130233320303 / 5

- [rules.action.dynamic](resources--nat_policy--reference--group-001.md#canonical-0221213332001013-1211231020320202-2110123331210011-3311233300323212-0323332201123333-3023103321300300-1013003221001113-1100103221120012)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)

<a id="canonical-1003100233302122-0331210123231312-3101222012030102-0011000230001133-0230002101033103-3330233311000231-2133020120103333-2030021301103211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302221310333212-1312212000201032-3322100220123012-3020000021322230-3002001102321330-3330021303330033-1022001302030102-1021120012020301"></a>

## rules.cloud_connect — cloud_connect / 000120131033 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- rules.cloud_connect

<a id="canonical-1111312031133000-0013212212212233-3310210223200331-0313021033103133-1121003013233000-3300330003220212-2313222100113121-3030310310330302"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for cloud connect.

Upstream description:

Reference to Cloud connect Object.

Provider validators and defaults (from schema source):

```go
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
cloud_connect {
  # Configure direct properties listed below.
}
```

<a id="canonical-2023202010022311-0012320330321121-2200320302132223-2230130131022203-0023131011032203-0233131332131303-2333323101330300-2003321233330201"></a>

## Direct properties — cloud_connect / 000120131033 / 3

- [refs](resources--nat_policy--reference--group-001.md#canonical-0301332012010322-0232321001330023-0201001213213300-2021111320100201-0303303133232022-2201221302102232-2123332010200222-2130300311213023): complete subsection reference.

<a id="canonical-0203013223222200-3023020203023032-1133202102110031-1313000020333000-2000030332333313-1232101230023122-3311022223001213-2323003213231312"></a>

## Next pages — cloud_connect / 000120131033 / 4

- [rules.cloud_connect.refs](resources--nat_policy--reference--group-001.md#canonical-0301332012010322-0232321001330023-0201001213213300-2021111320100201-0303303133232022-2201221302102232-2123332010200222-2130300311213023)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)

<a id="canonical-0301332012010322-0232321001330023-0201001213213300-2021111320100201-0303303133232022-2201221302102232-2123332010200222-2130300311213023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202310201321113-2200323232220031-3113321320032011-0220312010103322-3311311332300011-1320203023221333-0031012111023213-0311311020022303"></a>

## rules.cloud_connect.refs — refs / 131312100300 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [rules.cloud_connect](resources--nat_policy--reference--group-001.md#canonical-1003100233302122-0331210123231312-3101222012030102-0011000230001133-0230002101033103-3330233311000231-2133020120103333-2030021301103211)
- rules.cloud_connect.refs

<a id="canonical-1311001123333020-0003323000131120-2023112330112021-0023023323332100-1220122112322101-1232131110302212-1102210202213233-0102000022110010"></a>

Type: `"object"`. list nested block, Optional.

Cloud Connect. Reference to Cloud Connect Object.

Upstream description:

Reference to Cloud Connect Object.

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

<a id="canonical-2213032213023311-1221033002321000-1222030122322310-3323301100010131-0222121010112223-2002113123113200-1223101130013313-2201322031003312"></a>

## Direct properties — refs / 131312100300 / 3

<a id="canonical-2310102002130212-0032001312002113-1222232213203222-2312302200000113-1221231333123301-0330003303331132-3123000111103001-1220303011302312"></a>

<a id="canonical-1200210111033011-3123022221021031-3122221032310310-2211011130223120-1133000332101132-1310302102300333-3032301000130312-3210113302312020"></a>

## kind property — refs / 131312100300 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-0002001231223021-0311200313203211-3033330320232303-1000132203003103-3211023120123330-2200302330332302-3013202023102211-0012002000331211"></a>

<a id="canonical-1322323333223211-2201210231011001-2201311013102223-2310312133023330-0131201200301330-0021023303320211-3222000023110320-3232322101313133"></a>

## name property — refs / 131312100300 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-0021132233010122-0230110310230313-0220002303313232-1132031312311203-1130112010202133-3333330212333320-0320000000231010-2322201013321201"></a>

<a id="canonical-1221202131002222-2231210222020321-2132222030003333-2301120021102020-3000103311332200-3230120002031100-0311302320122333-0121012012020201"></a>

## namespace property — refs / 131312100300 / 6

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

<a id="canonical-2101113310233323-1233020220231013-0310133023202222-2212003223032103-0001031321012323-2220110302122003-2111301200231310-3331223113130300"></a>

<a id="canonical-3131230111210110-1010210222133322-1020000032333310-1333211101322222-2100200010301103-0110330210030113-2132333322011021-1200213100022301"></a>

## tenant property — refs / 131312100300 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-2110320100020102-3001000210330313-2302002112330200-1030122121200203-0110201032331233-2311122133232201-1311103031111301-2010301103300010"></a>

<a id="canonical-3003022313232132-3301211300103032-0011101211020313-1301022020221233-2012313121323233-1202231210020233-2000003011303312-3312022011130210"></a>

## uid property — refs / 131312100300 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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

<a id="canonical-2032131212011031-2132300333333022-1311130021001122-2302332013111122-1120100313001221-1031203312000130-1320020121221220-0323201321311312"></a>

## Next pages — refs / 131312100300 / 9

- [rules.cloud_connect](resources--nat_policy--reference--group-001.md#canonical-1003100233302122-0331210123231312-3101222012030102-0011000230001133-0230002101033103-3330233311000231-2133020120103333-2030021301103211)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)

<a id="canonical-0211202132100101-0101202032220222-0011011112111300-3220020012220110-0202012111300120-3331002203310331-0223221133023011-2213211332320022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301120203111013-3123033030333320-0113332321323300-2331001223120021-2210000202301021-3300112113001201-1023330222021211-3233210023313120"></a>

## rules.criteria — criteria / 000123022331 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- rules.criteria

<a id="canonical-1331311020201203-1012203013201322-2210101102101012-3133221110100000-3013310133201020-3030000201200230-0222331331001031-0020120020310023"></a>

Type: `"object"`. single nested block, Optional.

Match criteria of the packet to apply the NAT Rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("any",
    "icmp"),
  validators.ConflictingObjectAttributes("any",
    "tcp"),
  validators.ConflictingObjectAttributes("any",
    "udp"),
  validators.ConflictingObjectAttributes("icmp",
    "tcp"),
  validators.ConflictingObjectAttributes("icmp",
    "udp"),
  validators.ConflictingObjectAttributes("site_local_inside_network",
    "site_local_network"),
  validators.ConflictingObjectAttributes("tcp",
    "udp")}
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
  "x-ves-oneof-field-network_choice": "[\"site_local_inside_network\",\"site_local_network\"]",
  "x-ves-oneof-field-protocol_choice": "[\"any\",\"icmp\",\"tcp\",\"udp\"]"
}
```

Terraform syntax:

```terraform
criteria {
  # Configure direct properties listed below.
}
```

<a id="canonical-0103033221003320-1132322122132200-0132222233332302-0300233321302031-2221021003132321-1321230323031303-3031013001333010-0110122103223233"></a>

## Direct properties — criteria / 000123022331 / 3

- [any](resources--nat_policy--reference--group-001.md#canonical-3322130111322231-2211210112122132-1321020032320211-3030100213033221-1010202021312230-2202001320230330-0221301011133123-3300330103301332): complete subsection reference.

<a id="canonical-3023321213011103-1213233303032210-2233300212311212-3320113031112023-2033333213311012-2003101212100201-3111203021232020-0032111201121202"></a>

<a id="canonical-1320203300210303-0032031202302210-3131001131232000-1203301133223122-0220100202320301-1202101110330002-2101200212130003-1012222012220211"></a>

## destination_cidr property — criteria / 000123022331 / 4

Type: `["list", "string"]`. Optional.

Destination IP. Destination IP of the packet to match.

Upstream description:

Destination IP of the packet to match.

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
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

- [icmp](resources--nat_policy--reference--group-001.md#canonical-0311131033003010-2323213031003000-3111030000001003-1120111033230110-3331100233220022-3213331013230301-3111231201013101-3231122322001100): complete subsection reference.

- [site_local_inside_network](resources--nat_policy--reference--group-001.md#canonical-0111331033301131-2121203023323213-0031131103130323-2033322133233230-0111333322203113-2030022002031300-1111002013323031-0201223132031321): complete subsection reference.

- [site_local_network](resources--nat_policy--reference--group-001.md#canonical-3001002101100121-1103201310220210-0313020302130112-1003032003320230-3230010021213223-3301223012231123-2312011201231033-3102101023203103): complete subsection reference.

<a id="canonical-2231203023123131-1330030301021101-0031120201310102-2321322212011330-1122300221231311-0102013232323112-1221000200303000-2001223222231213"></a>

<a id="canonical-3020301211213231-2200111221130232-1112103213033300-2123312123011231-0322103113212121-2112110132111331-0011223132222231-0130210331223221"></a>

## source_cidr property — criteria / 000123022331 / 5

Type: `["list", "string"]`. Optional.

Source IP. Source IP of the packet to match.

Upstream description:

Source IP of the packet to match.

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
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

- [tcp](resources--nat_policy--reference--group-001.md#canonical-2301233133233231-1233311230111113-1300112232202320-0210231103202222-3030232230030201-0000232313210120-1011330101003002-1102333311310033): complete subsection reference.

- [udp](resources--nat_policy--reference--group-001.md#canonical-2211331033232333-2100001210303031-2131003002002112-0301201011323211-0132122121123001-3133302122301131-3113232033232301-1020210102220213): complete subsection reference.

<a id="canonical-0323311203123110-0332100231331033-3000110221232101-1232132222022131-2020313323022223-0031021113020103-1302102322101033-1021213002031021"></a>

## Next pages — criteria / 000123022331 / 6

- [rules.criteria.any](resources--nat_policy--reference--group-001.md#canonical-3322130111322231-2211210112122132-1321020032320211-3030100213033221-1010202021312230-2202001320230330-0221301011133123-3300330103301332)
- [rules.criteria.icmp](resources--nat_policy--reference--group-001.md#canonical-0311131033003010-2323213031003000-3111030000001003-1120111033230110-3331100233220022-3213331013230301-3111231201013101-3231122322001100)
- [rules.criteria.site_local_inside_network](resources--nat_policy--reference--group-001.md#canonical-0111331033301131-2121203023323213-0031131103130323-2033322133233230-0111333322203113-2030022002031300-1111002013323031-0201223132031321)
- [rules.criteria.site_local_network](resources--nat_policy--reference--group-001.md#canonical-3001002101100121-1103201310220210-0313020302130112-1003032003320230-3230010021213223-3301223012231123-2312011201231033-3102101023203103)
- [rules.criteria.tcp](resources--nat_policy--reference--group-001.md#canonical-2301233133233231-1233311230111113-1300112232202320-0210231103202222-3030232230030201-0000232313210120-1011330101003002-1102333311310033)
- [rules.criteria.udp](resources--nat_policy--reference--group-001.md#canonical-2211331033232333-2100001210303031-2131003002002112-0301201011323211-0132122121123001-3133302122301131-3113232033232301-1020210102220213)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)

<a id="canonical-3322130111322231-2211210112122132-1321020032320211-3030100213033221-1010202021312230-2202001320230330-0221301011133123-3300330103301332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101333122321113-2310030200000210-3223010130113122-3231320002122033-3202122112323212-0312032021023330-1133003223311331-2132132033012301"></a>

## rules.criteria.any — any / 121112232332 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-0211202132100101-0101202032220222-0011011112111300-3220020012220110-0202012111300120-3331002203310331-0223221133023011-2213211332320022)
- rules.criteria.any

<a id="canonical-0133132230022201-1030232020331001-3211112113202332-0002113111111302-1022033212300333-0203031212322311-3313321112230031-2031023113312311"></a>

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
any = {}
```

<a id="canonical-0132203330333000-1223221310133130-2301213002033330-0013031020030001-2123333330310010-1231122202021231-1001211220233112-3001001110032012"></a>

## Direct properties — any / 121112232332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2211211013223020-3330011020110321-1331102320312312-1202323210012133-2300101312323112-2323210301001220-3311111320313033-1130322212100203"></a>

## Next pages — any / 121112232332 / 4

- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-0211202132100101-0101202032220222-0011011112111300-3220020012220110-0202012111300120-3331002203310331-0223221133023011-2213211332320022)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)

<a id="canonical-0311131033003010-2323213031003000-3111030000001003-1120111033230110-3331100233220022-3213331013230301-3111231201013101-3231122322001100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0201110200331233-1133220123011031-0323211212011000-2302013222121302-3122230113102013-3220033021103332-0303212033321111-1221022321310212"></a>

## rules.criteria.icmp — icmp / 113111112032 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-0211202132100101-0101202032220222-0011011112111300-3220020012220110-0202012111300120-3331002203310331-0223221133023011-2213211332320022)
- rules.criteria.icmp

<a id="canonical-3221222003032231-1132232102223023-2000112020303312-1033021301313112-0013331010212211-3133120103101232-3302220011012303-2302312230101322"></a>

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
icmp = {}
```

<a id="canonical-0311231201302122-0020331301001112-3331233331033301-0101123013132212-1330203223003030-3032330002122130-1233121313312021-1303021001032023"></a>

## Direct properties — icmp / 113111112032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0201203120311301-3223232022030032-0213003101132312-1212220032321300-0303223311232000-1112311201011103-2121110112220003-1302320302022230"></a>

## Next pages — icmp / 113111112032 / 4

- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-0211202132100101-0101202032220222-0011011112111300-3220020012220110-0202012111300120-3331002203310331-0223221133023011-2213211332320022)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)

<a id="canonical-0111331033301131-2121203023323213-0031131103130323-2033322133233230-0111333322203113-2030022002031300-1111002013323031-0201223132031321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333232232103212-1231303301031313-2313310311103101-1130123013203010-1200023123302322-1131113120313001-0310300020233130-2030102301003220"></a>

## rules.criteria.site_local_inside_network — site_local_inside_network / 100010003133 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-0211202132100101-0101202032220222-0011011112111300-3220020012220110-0202012111300120-3331002203310331-0223221133023011-2213211332320022)
- rules.criteria.site_local_inside_network

<a id="canonical-0102112121222011-0231101310300032-3213100022112212-3101331233111122-2201102212112202-0133120220201231-2131021002023221-2023211013030202"></a>

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
site_local_inside_network = {}
```

<a id="canonical-3211203112012323-2301230123011302-2231121003001013-1121223303313010-0022013001012211-3121031123322322-0030003110310113-2203111210010011"></a>

## Direct properties — site_local_inside_network / 100010003133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3212003012213202-0223020133031033-3101300312103223-0032222112221101-1023213100330213-1320210000300231-3110313220103332-0212113133213213"></a>

## Next pages — site_local_inside_network / 100010003133 / 4

- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-0211202132100101-0101202032220222-0011011112111300-3220020012220110-0202012111300120-3331002203310331-0223221133023011-2213211332320022)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)

<a id="canonical-3001002101100121-1103201310220210-0313020302130112-1003032003320230-3230010021213223-3301223012231123-2312011201231033-3102101023203103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2323022103223230-0210302321210003-1120111231231310-0003003102023333-3000331321121103-0002231221030001-2310303130103133-2212210232323012"></a>

## rules.criteria.site_local_network — site_local_network / 213020011022 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-0211202132100101-0101202032220222-0011011112111300-3220020012220110-0202012111300120-3331002203310331-0223221133023011-2213211332320022)
- rules.criteria.site_local_network

<a id="canonical-3203013202301233-0012223313333001-2103302131222332-3321232030003221-0021233102331120-1331101223312330-1103022303121133-3020333311200333"></a>

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
site_local_network = {}
```

<a id="canonical-0202213312101222-3213202003322011-0211102101213302-2222000031312030-0201031103203012-1333113312311010-3131112100323101-1220030033100312"></a>

## Direct properties — site_local_network / 213020011022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3210223232032223-2121122000212111-3301330101032023-1222312113311110-2031322123321030-0030313300021123-1130132101031030-3023130320100202"></a>

## Next pages — site_local_network / 213020011022 / 4

- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-0211202132100101-0101202032220222-0011011112111300-3220020012220110-0202012111300120-3331002203310331-0223221133023011-2213211332320022)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)

<a id="canonical-2301233133233231-1233311230111113-1300112232202320-0210231103202222-3030232230030201-0000232313210120-1011330101003002-1102333311310033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000213213332331-2313200232111213-3213031321122220-2100300210111021-1313022102121321-0132311012100313-0331203111300322-0222021033003131"></a>

## rules.criteria.tcp — tcp / 000031211303 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-0211202132100101-0101202032220222-0011011112111300-3220020012220110-0202012111300120-3331002203310331-0223221133023011-2213211332320022)
- rules.criteria.tcp

<a id="canonical-0210221123330011-1131202312010100-3130302232123233-3233200210032020-2230300221202331-1201320221003212-1310130022312231-3032312002133220"></a>

Type: `"object"`. single nested block, Optional.

Action to apply on the packet if the NAT rule is applied.

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
tcp {
  # Configure direct properties listed below.
}
```

<a id="canonical-3122002230333121-3023001011020031-1123121212200023-1231303231200100-1012222111002200-1101303203320313-1332022332130100-2213010203230211"></a>

## Direct properties — tcp / 000031211303 / 3

- [destination_port](resources--nat_policy--reference--group-001.md#canonical-1122021131311023-1130133130321113-0303003123323203-2002323012121021-1221233233333311-2033011120113223-1303311330222230-1000300030200112): complete subsection reference.

- [source_port](resources--nat_policy--reference--group-001.md#canonical-0012333302323200-0202010231202231-0101230233201311-0133120031031103-2102021330123213-1121101113130303-0012332310201231-3332303022132030): complete subsection reference.

<a id="canonical-1311220201031330-0112211002033133-0212330130202313-3202100101302211-3323121321202221-2320301330303102-2221232312022323-0031122330112202"></a>

## Next pages — tcp / 000031211303 / 4

- [rules.criteria.tcp.destination_port](resources--nat_policy--reference--group-001.md#canonical-1122021131311023-1130133130321113-0303003123323203-2002323012121021-1221233233333311-2033011120113223-1303311330222230-1000300030200112)
- [rules.criteria.tcp.source_port](resources--nat_policy--reference--group-001.md#canonical-0012333302323200-0202010231202231-0101230233201311-0133120031031103-2102021330123213-1121101113130303-0012332310201231-3332303022132030)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-0211202132100101-0101202032220222-0011011112111300-3220020012220110-0202012111300120-3331002203310331-0223221133023011-2213211332320022)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)

<a id="canonical-1122021131311023-1130133130321113-0303003123323203-2002323012121021-1221233233333311-2033011120113223-1303311330222230-1000300030200112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123111013301320-3121033021032031-3102331113223133-3331123003303131-2300112310312330-2000323321012131-0223313311203111-0002231211333130"></a>

## rules.criteria.tcp.destination_port — destination_port / 231121223020 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-0211202132100101-0101202032220222-0011011112111300-3220020012220110-0202012111300120-3331002203310331-0223221133023011-2213211332320022)
- [rules.criteria.tcp](resources--nat_policy--reference--group-001.md#canonical-2301233133233231-1233311230111113-1300112232202320-0210231103202222-3030232230030201-0000232313210120-1011330101003002-1102333311310033)
- rules.criteria.tcp.destination_port

<a id="canonical-1033332331020031-0202200021331310-2012033200300103-1303023223113200-1331232010021331-2222123330021122-1311113330012031-2323311110213003"></a>

Type: `"object"`. single nested block, Optional.

Port match of the request can be a range or a specific port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_port_match",
    "port"),
  validators.ConflictingObjectAttributes("no_port_match",
    "port_ranges"),
  validators.ConflictingObjectAttributes("port",
    "port_ranges")}
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
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

Terraform syntax:

```terraform
destination_port {
  # Configure direct properties listed below.
}
```

<a id="canonical-1321322320121120-3032301310230123-1301010030013210-2131331122002320-3231100323023302-2200231201213002-1132333302002002-2030331333123333"></a>

## Direct properties — destination_port / 231121223020 / 3

- [no_port_match](resources--nat_policy--reference--group-001.md#canonical-0023103200122233-2030011202213110-0331310103333012-0121301122133133-1313110212003031-2003012021111102-0311333212121033-2200022221020302): complete subsection reference.

<a id="canonical-3102213323303103-3323102302201301-1222122300131211-0003201122113001-1222030010211321-3002200103231023-0022011301012311-2101313202012331"></a>

<a id="canonical-0301210302032221-3102300211303333-2132221020333033-1113313112113312-0032332322331303-0100221021200020-1113210102111100-0032101010232310"></a>

## port property — destination_port / 231121223020 / 4

Type: `"number"`. Optional.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Upstream description:

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0221111332133021-2201103330202120-1000112130111101-3303212130333302-0222023233013310-1022130310211001-1220233102320011-2100221032313313"></a>

<a id="canonical-1023222001332132-1130333202021222-0000100223232113-3001103102022322-1021023010333220-1132112300213322-2322311033131110-1313200200221333"></a>

## port_ranges property — destination_port / 231121223020 / 5

Type: `"string"`. Optional.

Exclusive with \[no\_port\_match port\] Port range to match.

Upstream description:

Exclusive with \[no\_port\_match port\] Port range to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 32,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  }
}
```

<a id="canonical-3331331301103221-1330001021102230-2321332222111203-0311011031123012-0030122212303232-2023213300021210-1122123213133001-2201320323010333"></a>

## Next pages — destination_port / 231121223020 / 6

- [rules.criteria.tcp.destination_port.no_port_match](resources--nat_policy--reference--group-001.md#canonical-0023103200122233-2030011202213110-0331310103333012-0121301122133133-1313110212003031-2003012021111102-0311333212121033-2200022221020302)
- [rules.criteria.tcp](resources--nat_policy--reference--group-001.md#canonical-2301233133233231-1233311230111113-1300112232202320-0210231103202222-3030232230030201-0000232313210120-1011330101003002-1102333311310033)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)

<a id="canonical-0023103200122233-2030011202213110-0331310103333012-0121301122133133-1313110212003031-2003012021111102-0311333212121033-2200022221020302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3322003321301213-2211202012322002-0002103232102232-2300213303201313-0302313011320021-2032311131203212-1200202333211123-0022002302202230"></a>

## rules.criteria.tcp.destination_port.no_port_match — no_port_match / 021121021323 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-0211202132100101-0101202032220222-0011011112111300-3220020012220110-0202012111300120-3331002203310331-0223221133023011-2213211332320022)
- [rules.criteria.tcp](resources--nat_policy--reference--group-001.md#canonical-2301233133233231-1233311230111113-1300112232202320-0210231103202222-3030232230030201-0000232313210120-1011330101003002-1102333311310033)
- [rules.criteria.tcp.destination_port](resources--nat_policy--reference--group-001.md#canonical-1122021131311023-1130133130321113-0303003123323203-2002323012121021-1221233233333311-2033011120113223-1303311330222230-1000300030200112)
- rules.criteria.tcp.destination_port.no_port_match

<a id="canonical-3111012322132130-2222201223111313-3000232132032130-0123302322332310-0002002231021212-0121232200231023-3022223231323332-3211323220100311"></a>

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
no_port_match = {}
```

<a id="canonical-0101000333211220-1331000322323323-1031212233003331-3200200012212111-0030032211021210-0000231000013020-1333133112311020-2131202002103221"></a>

## Direct properties — no_port_match / 021121021323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1020231032131332-3300333030103112-1012201213011231-2312322330012101-2121330222302333-3002332221301333-2032020100010032-1303003100033020"></a>

## Next pages — no_port_match / 021121021323 / 4

- [rules.criteria.tcp.destination_port](resources--nat_policy--reference--group-001.md#canonical-1122021131311023-1130133130321113-0303003123323203-2002323012121021-1221233233333311-2033011120113223-1303311330222230-1000300030200112)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)

<a id="canonical-0012333302323200-0202010231202231-0101230233201311-0133120031031103-2102021330123213-1121101113130303-0012332310201231-3332303022132030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122230133320223-3021120120310033-3020003013010333-1113002023310222-2030120231101030-0320233200213302-2320032323000130-0310132310213130"></a>

## rules.criteria.tcp.source_port — source_port / 323220200333 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-0211202132100101-0101202032220222-0011011112111300-3220020012220110-0202012111300120-3331002203310331-0223221133023011-2213211332320022)
- [rules.criteria.tcp](resources--nat_policy--reference--group-001.md#canonical-2301233133233231-1233311230111113-1300112232202320-0210231103202222-3030232230030201-0000232313210120-1011330101003002-1102333311310033)
- rules.criteria.tcp.source_port

<a id="canonical-3102320021110001-1201320132120300-1120310103221333-1231330312020013-3122022222320220-0330003320203103-1003233333011223-3001322013132011"></a>

Type: `"object"`. single nested block, Optional.

Port match of the request can be a range or a specific port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_port_match",
    "port"),
  validators.ConflictingObjectAttributes("no_port_match",
    "port_ranges"),
  validators.ConflictingObjectAttributes("port",
    "port_ranges")}
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
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

Terraform syntax:

```terraform
source_port {
  # Configure direct properties listed below.
}
```

<a id="canonical-0030112001012212-2202313130320222-0133110211312121-2120123211302320-0002020232101211-1030121221213312-0320231102021101-3113132003233032"></a>

## Direct properties — source_port / 323220200333 / 3

- [no_port_match](resources--nat_policy--reference--group-001.md#canonical-0131000301210222-0102320313012332-0232123211210021-3111322001321320-2323113111133310-2201111330222033-2322303110110020-0211031332011311): complete subsection reference.

<a id="canonical-0133301301033003-1023022100221302-0013322012103223-0113111211031123-2213120133011032-1120033203322020-2233322322320033-3010101320233231"></a>

<a id="canonical-1021222303213332-2213102002020013-3013321231203111-0012103310022322-3233030122333012-2110303111032021-0321100003313022-3112313203300213"></a>

## port property — source_port / 323220200333 / 4

Type: `"number"`. Optional.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Upstream description:

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0032101111133113-1321021031323302-2010132000023113-2020101113020020-1002222223022003-2003212200313320-3203322113122223-3233002331011222"></a>

<a id="canonical-1223230013322300-2122012121123202-0000000122022220-1101122011223321-1203132121132120-0120211002213201-1012100320312123-3321212212202013"></a>

## port_ranges property — source_port / 323220200333 / 5

Type: `"string"`. Optional.

Exclusive with \[no\_port\_match port\] Port range to match.

Upstream description:

Exclusive with \[no\_port\_match port\] Port range to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 32,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  }
}
```

<a id="canonical-3230233120003120-3021200011023121-0302000032100033-3211120311010330-1200012121021130-3021321203010132-1311031133202110-1120121012102331"></a>

## Next pages — source_port / 323220200333 / 6

- [rules.criteria.tcp.source_port.no_port_match](resources--nat_policy--reference--group-001.md#canonical-0131000301210222-0102320313012332-0232123211210021-3111322001321320-2323113111133310-2201111330222033-2322303110110020-0211031332011311)
- [rules.criteria.tcp](resources--nat_policy--reference--group-001.md#canonical-2301233133233231-1233311230111113-1300112232202320-0210231103202222-3030232230030201-0000232313210120-1011330101003002-1102333311310033)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)

<a id="canonical-0131000301210222-0102320313012332-0232123211210021-3111322001321320-2323113111133310-2201111330222033-2322303110110020-0211031332011311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200202301131111-0121223111222332-2010132121222222-3000220001320213-0233102223331321-3132012103123033-0300322033230202-1201230110303001"></a>

## rules.criteria.tcp.source_port.no_port_match — no_port_match / 011332311130 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-0211202132100101-0101202032220222-0011011112111300-3220020012220110-0202012111300120-3331002203310331-0223221133023011-2213211332320022)
- [rules.criteria.tcp](resources--nat_policy--reference--group-001.md#canonical-2301233133233231-1233311230111113-1300112232202320-0210231103202222-3030232230030201-0000232313210120-1011330101003002-1102333311310033)
- [rules.criteria.tcp.source_port](resources--nat_policy--reference--group-001.md#canonical-0012333302323200-0202010231202231-0101230233201311-0133120031031103-2102021330123213-1121101113130303-0012332310201231-3332303022132030)
- rules.criteria.tcp.source_port.no_port_match

<a id="canonical-3322230212210221-3330012023122231-2300303232201022-2023021000112232-3132120132003331-2100131111000333-1022303130110101-3303233232010230"></a>

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
no_port_match = {}
```

<a id="canonical-3220220320332101-0210112003233021-1302111213200220-3231221030203322-3210013022130302-3031330330320333-2201323131311001-0021030033221302"></a>

## Direct properties — no_port_match / 011332311130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0111132033303233-0232113313212031-1222032102111111-2222113320001131-2123303120021222-3302202121313233-2100202201231333-1222023310101231"></a>

## Next pages — no_port_match / 011332311130 / 4

- [rules.criteria.tcp.source_port](resources--nat_policy--reference--group-001.md#canonical-0012333302323200-0202010231202231-0101230233201311-0133120031031103-2102021330123213-1121101113130303-0012332310201231-3332303022132030)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)

<a id="canonical-2211331033232333-2100001210303031-2131003002002112-0301201011323211-0132122121123001-3133302122301131-3113232033232301-1020210102220213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1013012230101021-2001102001133232-3103200103130033-0302022333303300-0100333000301033-1223130023000322-2301311123011021-1311021323230132"></a>

## rules.criteria.udp — udp / 203212003200 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-0211202132100101-0101202032220222-0011011112111300-3220020012220110-0202012111300120-3331002203310331-0223221133023011-2213211332320022)
- rules.criteria.udp

<a id="canonical-0223102112002322-3313133203331103-2212323000211302-0333312230312003-0110021322221312-3131022233022032-0331203011231233-0210132001313232"></a>

Type: `"object"`. single nested block, Optional.

Action to apply on the packet if the NAT rule is applied.

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
udp {
  # Configure direct properties listed below.
}
```

<a id="canonical-3220301321120102-3200000030011220-2302023122101012-0012012121333201-3332300102113232-3301130110022330-2200320332303021-3323003123013110"></a>

## Direct properties — udp / 203212003200 / 3

- [destination_port](resources--nat_policy--reference--group-001.md#canonical-3130130033112021-3211102103000131-0212331101101001-3212003010202320-0030001002100110-2200011123222321-0202123221211202-0132131211322021): complete subsection reference.

- [source_port](resources--nat_policy--reference--group-001.md#canonical-0122012203231130-0202333101303112-0020300323310323-1023103331210023-0122120331233113-2312223111231220-2202230111212321-3201321202121231): complete subsection reference.

<a id="canonical-1332032300221233-0022133310131013-1101323232032213-1331200320213111-2103323232332111-0300133122210121-2113220233312130-1032112121323123"></a>

## Next pages — udp / 203212003200 / 4

- [rules.criteria.udp.destination_port](resources--nat_policy--reference--group-001.md#canonical-3130130033112021-3211102103000131-0212331101101001-3212003010202320-0030001002100110-2200011123222321-0202123221211202-0132131211322021)
- [rules.criteria.udp.source_port](resources--nat_policy--reference--group-001.md#canonical-0122012203231130-0202333101303112-0020300323310323-1023103331210023-0122120331233113-2312223111231220-2202230111212321-3201321202121231)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-0211202132100101-0101202032220222-0011011112111300-3220020012220110-0202012111300120-3331002203310331-0223221133023011-2213211332320022)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)

<a id="canonical-3130130033112021-3211102103000131-0212331101101001-3212003010202320-0030001002100110-2200011123222321-0202123221211202-0132131211322021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113203132210101-0011012121211202-2302030223112203-1002000203001133-3121223012211303-2012103221130223-3123203331320003-3233133102012013"></a>

## rules.criteria.udp.destination_port — destination_port / 133323030100 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-0211202132100101-0101202032220222-0011011112111300-3220020012220110-0202012111300120-3331002203310331-0223221133023011-2213211332320022)
- [rules.criteria.udp](resources--nat_policy--reference--group-001.md#canonical-2211331033232333-2100001210303031-2131003002002112-0301201011323211-0132122121123001-3133302122301131-3113232033232301-1020210102220213)
- rules.criteria.udp.destination_port

<a id="canonical-2213211112020101-0112332022023200-0322303202310001-2233103113201101-0230303232112221-1302020330200302-3032031232203023-2223321213032321"></a>

Type: `"object"`. single nested block, Optional.

Port match of the request can be a range or a specific port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_port_match",
    "port"),
  validators.ConflictingObjectAttributes("no_port_match",
    "port_ranges"),
  validators.ConflictingObjectAttributes("port",
    "port_ranges")}
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
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

Terraform syntax:

```terraform
destination_port {
  # Configure direct properties listed below.
}
```

<a id="canonical-3130220232223213-0201133113310201-2232202211110221-2230031123022212-2012300112202110-0323123002223232-1000201133303220-0030122332300111"></a>

## Direct properties — destination_port / 133323030100 / 3

- [no_port_match](resources--nat_policy--reference--group-001.md#canonical-1023202301200301-2012021332333110-2111222121000301-1212113003123122-1313200211112022-3212022321033131-3320130213113231-2333021020301213): complete subsection reference.

<a id="canonical-0332001133021312-3331102313121022-3312030330113303-0001303331203230-2001132022101300-0231100101123200-1322213201032022-1131230000112111"></a>

<a id="canonical-3013110132123333-0133021002320101-0031013001333000-0230101201123203-2030132223013023-2021102233132002-1300001321100012-2323012121233203"></a>

## port property — destination_port / 133323030100 / 4

Type: `"number"`. Optional.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Upstream description:

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0111201230022022-0120011011222131-2232111111030223-0233003000201122-1202300131310011-1003112102100203-2311100312222110-3221032011233011"></a>

<a id="canonical-1300121222111131-1223030123031020-1220101102211323-1333220111012023-3301210120031330-2033210230320301-2000023021233122-0301200003111211"></a>

## port_ranges property — destination_port / 133323030100 / 5

Type: `"string"`. Optional.

Exclusive with \[no\_port\_match port\] Port range to match.

Upstream description:

Exclusive with \[no\_port\_match port\] Port range to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 32,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  }
}
```

<a id="canonical-1121030302000212-0101322223011201-2100010322233233-0121223202001110-2100231033133310-2121211102201133-2021033010333032-3331331313301233"></a>

## Next pages — destination_port / 133323030100 / 6

- [rules.criteria.udp.destination_port.no_port_match](resources--nat_policy--reference--group-001.md#canonical-1023202301200301-2012021332333110-2111222121000301-1212113003123122-1313200211112022-3212022321033131-3320130213113231-2333021020301213)
- [rules.criteria.udp](resources--nat_policy--reference--group-001.md#canonical-2211331033232333-2100001210303031-2131003002002112-0301201011323211-0132122121123001-3133302122301131-3113232033232301-1020210102220213)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)

<a id="canonical-1023202301200301-2012021332333110-2111222121000301-1212113003123122-1313200211112022-3212022321033131-3320130213113231-2333021020301213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000310101302223-2332302102302312-2021213102222110-2300323023101021-2303321332330302-1011123123221200-3320003033330221-0203313011223321"></a>

## rules.criteria.udp.destination_port.no_port_match — no_port_match / 100113121203 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-0211202132100101-0101202032220222-0011011112111300-3220020012220110-0202012111300120-3331002203310331-0223221133023011-2213211332320022)
- [rules.criteria.udp](resources--nat_policy--reference--group-001.md#canonical-2211331033232333-2100001210303031-2131003002002112-0301201011323211-0132122121123001-3133302122301131-3113232033232301-1020210102220213)
- [rules.criteria.udp.destination_port](resources--nat_policy--reference--group-001.md#canonical-3130130033112021-3211102103000131-0212331101101001-3212003010202320-0030001002100110-2200011123222321-0202123221211202-0132131211322021)
- rules.criteria.udp.destination_port.no_port_match

<a id="canonical-1033332131110232-2101313200011232-0023301223220131-2300101210203113-0322312232211013-0231003102323020-1223210322302121-0100013301202002"></a>

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
no_port_match = {}
```

<a id="canonical-0013330232113211-1222032131023012-1020131220311113-3323111320303332-0023020331311112-2030312102213020-3223311200231131-0213132031312231"></a>

## Direct properties — no_port_match / 100113121203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1301112003010201-2033332230100130-2010120332032012-0033100223120131-1310012122123100-0012200202012000-0312200230233121-3113313122021212"></a>

## Next pages — no_port_match / 100113121203 / 4

- [rules.criteria.udp.destination_port](resources--nat_policy--reference--group-001.md#canonical-3130130033112021-3211102103000131-0212331101101001-3212003010202320-0030001002100110-2200011123222321-0202123221211202-0132131211322021)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)

<a id="canonical-0122012203231130-0202333101303112-0020300323310323-1023103331210023-0122120331233113-2312223111231220-2202230111212321-3201321202121231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123113301021130-2220321321013221-3222031100232333-1021002233301312-0121101113310002-1221232301113032-0132330132010032-2111203122102302"></a>

## rules.criteria.udp.source_port — source_port / 132021202300 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-0211202132100101-0101202032220222-0011011112111300-3220020012220110-0202012111300120-3331002203310331-0223221133023011-2213211332320022)
- [rules.criteria.udp](resources--nat_policy--reference--group-001.md#canonical-2211331033232333-2100001210303031-2131003002002112-0301201011323211-0132122121123001-3133302122301131-3113232033232301-1020210102220213)
- rules.criteria.udp.source_port

<a id="canonical-2310211313233212-3112103031332033-2111232310212331-1122113310203223-2231232121120030-0320023321200210-3133332011133133-3032323233322001"></a>

Type: `"object"`. single nested block, Optional.

Port match of the request can be a range or a specific port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_port_match",
    "port"),
  validators.ConflictingObjectAttributes("no_port_match",
    "port_ranges"),
  validators.ConflictingObjectAttributes("port",
    "port_ranges")}
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
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

Terraform syntax:

```terraform
source_port {
  # Configure direct properties listed below.
}
```

<a id="canonical-2302002131322011-0122320001021120-1013020201123010-0121023003032200-2013302032212033-2231033230302030-3220113010030212-1210120311003130"></a>

## Direct properties — source_port / 132021202300 / 3

- [no_port_match](resources--nat_policy--reference--group-001.md#canonical-3000122033003101-0033322313222201-2101220213110032-3230123310221113-3313100332223012-3100130302003103-1310202330033011-1311013113120332): complete subsection reference.

<a id="canonical-3231201210302201-2002123223123203-2013101332111132-2100121303103200-3122322202021022-2311323312021003-0302120113311131-2132030231332103"></a>

<a id="canonical-1102230030223011-1001201110023330-3301102231003231-2032332321323213-3110123210013102-1203202022222110-3023030203000332-3232013300011133"></a>

## port property — source_port / 132021202300 / 4

Type: `"number"`. Optional.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Upstream description:

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1231203121121100-0232313332300023-0133123120122203-3112033212023210-2330220333001021-0031302131203323-2133333030023003-0031020003030311"></a>

<a id="canonical-3300200300210010-3112201100220023-0213102021113022-2002010222223200-1001130202311120-0333112001001130-1333330222101101-1233231200200113"></a>

## port_ranges property — source_port / 132021202300 / 5

Type: `"string"`. Optional.

Exclusive with \[no\_port\_match port\] Port range to match.

Upstream description:

Exclusive with \[no\_port\_match port\] Port range to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 32,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  }
}
```

<a id="canonical-2103220213133003-2003302300123023-0031221103032131-0320132103112210-3103011212313030-1031212200201023-3331322022131321-1213313100322223"></a>

## Next pages — source_port / 132021202300 / 6

- [rules.criteria.udp.source_port.no_port_match](resources--nat_policy--reference--group-001.md#canonical-3000122033003101-0033322313222201-2101220213110032-3230123310221113-3313100332223012-3100130302003103-1310202330033011-1311013113120332)
- [rules.criteria.udp](resources--nat_policy--reference--group-001.md#canonical-2211331033232333-2100001210303031-2131003002002112-0301201011323211-0132122121123001-3133302122301131-3113232033232301-1020210102220213)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)

<a id="canonical-3000122033003101-0033322313222201-2101220213110032-3230123310221113-3313100332223012-3100130302003103-1310202330033011-1311013113120332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0111312333310311-3000212322321001-0102132103212231-3302121030233120-3310300013221233-2330321013211023-0031231223013023-1021031220212133"></a>

## rules.criteria.udp.source_port.no_port_match — no_port_match / 133230100132 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-0211202132100101-0101202032220222-0011011112111300-3220020012220110-0202012111300120-3331002203310331-0223221133023011-2213211332320022)
- [rules.criteria.udp](resources--nat_policy--reference--group-001.md#canonical-2211331033232333-2100001210303031-2131003002002112-0301201011323211-0132122121123001-3133302122301131-3113232033232301-1020210102220213)
- [rules.criteria.udp.source_port](resources--nat_policy--reference--group-001.md#canonical-0122012203231130-0202333101303112-0020300323310323-1023103331210023-0122120331233113-2312223111231220-2202230111212321-3201321202121231)
- rules.criteria.udp.source_port.no_port_match

<a id="canonical-3210010222231233-1200201011123322-1221321210213120-0210300101101300-3010300031011202-0113321310010002-2302111320013233-2321201011211111"></a>

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
no_port_match = {}
```

<a id="canonical-0200231232120113-0323332011012121-2122322221300010-1202023310323331-3121103321002333-2032120220111001-1010023110213202-3312033112222131"></a>

## Direct properties — no_port_match / 133230100132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0333010132002313-2233200203310203-3300220112030211-1022320312001100-0110010310103130-1103311122022211-1103133201333130-0120123113012002"></a>

## Next pages — no_port_match / 133230100132 / 4

- [rules.criteria.udp.source_port](resources--nat_policy--reference--group-001.md#canonical-0122012203231130-0202333101303112-0020300323310323-1023103331210023-0122120331233113-2312223111231220-2202230111212321-3201321202121231)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)

<a id="canonical-2333111032320031-2222102332001330-1220210111201123-1011122013230003-3231123330013221-3103132010233220-0010221330011032-0200130112321100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313101100013322-1310112223110230-1021012033223031-3333031311331100-0333202011032102-1011331001221101-3211221220310302-0032202032130031"></a>

## rules.disable_spec — disable_spec / 011211233002 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- rules.disable_spec

<a id="canonical-1230201131132230-0312113000231211-0321211220110203-0100033112202320-1331003310133122-1020132131022322-3001232110133020-3012213022202102"></a>

Type: `["object", {}]`. Optional.

Enable this option

Terraform syntax:

```terraform
disable_spec = {}
```

<a id="canonical-2231112321330011-2023200220311103-2323311033000121-2023010321310023-3200323321000132-0202230130212310-3020332320100000-2122013303333221"></a>

## Direct properties — disable_spec / 011211233002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0222132230221101-0302323011320131-0201332002112023-1312223200203132-2331232121120222-2332311302213022-2111331320212220-3303231132313003"></a>

## Next pages — disable_spec / 011211233002 / 4

- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)

<a id="canonical-3301200022202232-1013113223132333-1111013013320103-3201130310232122-1233023330321002-0020013000032003-0121010123113111-2203023302031312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0323013100000302-3200211232123303-2133300100202132-1301101032301103-0302200120322222-0131300012101312-2311312021013300-3202112013012222"></a>

## rules.enable — enable / 233312203032 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- rules.enable

<a id="canonical-2111322330111130-1222222221011221-0003131112130332-3013311202332230-2022213322300203-2102101223003130-1000103202223021-3012301120302301"></a>

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
enable = {}
```

<a id="canonical-2110021030122022-0322222022310221-1103121101020012-0101222212302111-0112303021011303-3011202202131330-2120112210320110-3102010100230001"></a>

## Direct properties — enable / 233312203032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0322120210021002-0223001131031313-1100233312232111-3022230223233010-1133330003231021-3103233031003321-2032212111131202-1311021102032013"></a>

## Next pages — enable / 233312203032 / 4

- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)

<a id="canonical-0003122032201330-0222033131003300-3123021011023102-3103311200332223-1132130323200320-1223032020223203-0130321232310300-0121313320300220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023202232310112-1021111212030230-3120313022030300-0130112313212302-3311303021002312-1301211310202333-3332322230200300-2331201221223122"></a>

## rules.node_interface — node_interface / 323121102311 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- rules.node_interface

<a id="canonical-0320333020103122-0220031302222200-1020233221300333-2130230333210031-1020012322200323-1001011003332000-2211233121021101-2102131210002123"></a>

Type: `"object"`. single nested block, Optional.

On multinode site, this type holds the information about per node interfaces.

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
node_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-0201012221132203-0110022221022030-2301013120331100-2003020122033213-3102032210201201-0030321301130221-1310313022311213-0331222223211111"></a>

## Direct properties — node_interface / 323121102311 / 3

- [list](resources--nat_policy--reference--group-001.md#canonical-2122022232302031-2333032101032021-0231201133031013-2321111030231130-2101303131110130-2003002331021120-0231312302331333-3110002222230032): complete subsection reference.

<a id="canonical-3221330212323233-3122320302200202-1010002231101202-0132221033330102-1023301023023223-0100330100201322-0131200320222311-1202310333302210"></a>

## Next pages — node_interface / 323121102311 / 4

- [rules.node_interface.list](resources--nat_policy--reference--group-001.md#canonical-2122022232302031-2333032101032021-0231201133031013-2321111030231130-2101303131110130-2003002331021120-0231312302331333-3110002222230032)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)

<a id="canonical-2122022232302031-2333032101032021-0231201133031013-2321111030231130-2101303131110130-2003002331021120-0231312302331333-3110002222230032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333322133320031-2202323011010333-3231033022202301-2313110302031022-3323010023112231-1123123200131333-0203220200031332-1110120220133113"></a>

## rules.node_interface.list — list / 232113000332 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [rules.node_interface](resources--nat_policy--reference--group-001.md#canonical-0003122032201330-0222033131003300-3123021011023102-3103311200332223-1132130323200320-1223032020223203-0130321232310300-0121313320300220)
- rules.node_interface.list

<a id="canonical-0000001202123023-1310210330222313-0121030022302300-1222001230113003-2303331012003003-3323213300222302-1312103320012133-1201111213123010"></a>

Type: `"object"`. list nested block, Optional.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

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
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

Terraform syntax:

```terraform
list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0320023312203312-0030021220113312-1323122232131132-1322223031113333-2023322123232223-3102022200212310-0101133222213323-0332230232312232"></a>

## Direct properties — list / 232113000332 / 3

- [interface](resources--nat_policy--reference--group-001.md#canonical-0212201101030223-3030213222002202-1030030320322233-3210310211222203-3132131210320231-3000222200032223-0113132133301103-2011011231112332): complete subsection reference.

<a id="canonical-2322312222011333-3303132232310100-2000012231230003-2322012002210320-2200131030321012-2003202031311323-0111132332003132-3300312233110021"></a>

<a id="canonical-0010230112302030-2100020301102101-2321033220032101-2031223333311011-1312320120310211-2020030210212033-2332311100302312-1220312133211322"></a>

## node property — list / 232113000332 / 4

Type: `"string"`. Optional.

Node. Node name on this site.

Upstream description:

Node name on this site.

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

<a id="canonical-1130313331131333-1303310002201322-2033100233020131-0011312231323300-1113222210113333-0200010312113201-2100130323131103-1213023332120223"></a>

## Next pages — list / 232113000332 / 5

- [rules.node_interface.list.interface](resources--nat_policy--reference--group-001.md#canonical-0212201101030223-3030213222002202-1030030320322233-3210310211222203-3132131210320231-3000222200032223-0113132133301103-2011011231112332)
- [rules.node_interface](resources--nat_policy--reference--group-001.md#canonical-0003122032201330-0222033131003300-3123021011023102-3103311200332223-1132130323200320-1223032020223203-0130321232310300-0121313320300220)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)

<a id="canonical-0212201101030223-3030213222002202-1030030320322233-3210310211222203-3132131210320231-3000222200032223-0113132133301103-2011011231112332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0313011002003233-2331222012300322-2113333200112110-2221222132203030-2033300130122110-3312011000220332-2101313313203211-1020030131023020"></a>

## rules.node_interface.list.interface — interface / 020303031202 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [rules.node_interface](resources--nat_policy--reference--group-001.md#canonical-0003122032201330-0222033131003300-3123021011023102-3103311200332223-1132130323200320-1223032020223203-0130321232310300-0121313320300220)
- [rules.node_interface.list](resources--nat_policy--reference--group-001.md#canonical-2122022232302031-2333032101032021-0231201133031013-2321111030231130-2101303131110130-2003002331021120-0231312302331333-3110002222230032)
- rules.node_interface.list.interface

<a id="canonical-2312213120023023-2021323303020323-0303202330123110-3023001213103110-3130202321102212-2133201012122011-3313023030311301-3332003320212312"></a>

Type: `"object"`. list nested block, Optional.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-1121310312120313-0220202001122030-2133302130123203-0000230013022302-1223103121302333-1313103111013113-0011212123010330-2121031201113033"></a>

## Direct properties — interface / 020303031202 / 3

<a id="canonical-2222101030310100-0212331213312030-2033320312022233-2013033230131112-1002001010322200-0323311030020110-2032033012112000-0313010330301112"></a>

<a id="canonical-1110131020120010-2230003110311100-3030213310023223-0020312122300313-3322030103121120-3333030322133323-0010320120332320-0303331002331003"></a>

## kind property — interface / 020303031202 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-0333013321212002-2101330123002032-1133303211220230-2132233010033023-1133023011031311-2302321330012231-3312302110020130-2030100000233221"></a>

<a id="canonical-2100021310202310-1023301122331130-2302130200203101-0330130030011003-0013331201201121-0302211010322313-2121333200221223-2002300210010031"></a>

## name property — interface / 020303031202 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-3120132331302202-0103012223330103-2313110013111210-1131213302211220-1102202303023310-2230230002331202-3013213000023011-3200032032102033"></a>

<a id="canonical-3311213121210030-2222302303110112-3112113033303233-3033310023330012-3020232031201121-3121011131113311-0233113202313320-1000130030303132"></a>

## namespace property — interface / 020303031202 / 6

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

<a id="canonical-2331013322323300-2222203123100003-0220033111220221-2022223020203231-1300313033132101-2300210220012021-3230102222300000-2230302322220331"></a>

<a id="canonical-3011222130230302-0012312321001011-1300330021112023-2221300032333132-2220211300010312-0330213332330130-2231323212111213-3002032230220113"></a>

## tenant property — interface / 020303031202 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-0312221311000331-0200023020230313-1103211102210121-3303132033032323-2320113100121223-2321110031023112-0302221123121123-2102120312002222"></a>

<a id="canonical-1021212133012001-0200310323233232-2330130223113031-0223123123132200-2303201013012331-0112102222133300-2221123311133120-2021012132030110"></a>

## uid property — interface / 020303031202 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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

<a id="canonical-0300033301302331-2010001220231003-0233003032301002-0330010321120203-0211311302030233-1212010113021113-2122223311122200-0121310023212230"></a>

## Next pages — interface / 020303031202 / 9

- [rules.node_interface.list](resources--nat_policy--reference--group-001.md#canonical-2122022232302031-2333032101032021-0231201133031013-2321111030231130-2101303131110130-2003002331021120-0231312302331333-3110002222230032)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)

<a id="canonical-2000331302003211-2112122330312222-0231223203303123-2331332301321123-1023103133311320-1120102103133013-3031202021101232-2321033122322230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003001221301112-0022102222013102-0303020120002113-1030110021112021-2003210003222112-2001221000311133-3032312033313221-1120333001003323"></a>

## rules.segment — segment / 223332022230 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- rules.segment

<a id="canonical-1002003323223203-1013303333103203-0031010010033300-2212110101202121-1300320122201213-3231233212331122-1212101101011202-0220322331222011"></a>

Type: `"object"`. single nested block, Optional.

Segment Reference Type. Reference to Segment Object.

Upstream description:

Reference to Segment Object.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0301212010133130-0013230323203220-0310123311201331-3002220212303302-1233201020003320-1123022330310222-3221211212010102-0031321330003223"></a>

## Direct properties — segment / 223332022230 / 3

- [refs](resources--nat_policy--reference--group-001.md#canonical-1103103320230130-2322002110112311-3333012233010000-2120130121230220-1210210322000211-0322321211012203-3223233031210023-0210021200001032): complete subsection reference.

<a id="canonical-0203121200031023-1111320313203112-1120100221132020-3021010020033011-3121011033020023-2310310202200000-3230320012122121-3010212222010003"></a>

## Next pages — segment / 223332022230 / 4

- [rules.segment.refs](resources--nat_policy--reference--group-001.md#canonical-1103103320230130-2322002110112311-3333012233010000-2120130121230220-1210210322000211-0322321211012203-3223233031210023-0210021200001032)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)

<a id="canonical-1103103320230130-2322002110112311-3333012233010000-2120130121230220-1210210322000211-0322321211012203-3223233031210023-0210021200001032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202230212003331-2222002112113003-2100320301111320-3012312033332003-3101130212223200-2300303222210001-2103013033010221-2112321003232013"></a>

## rules.segment.refs — refs / 133122003123 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [rules.segment](resources--nat_policy--reference--group-001.md#canonical-2000331302003211-2112122330312222-0231223203303123-2331332301321123-1023103133311320-1120102103133013-3031202021101232-2321033122322230)
- rules.segment.refs

<a id="canonical-0332133133132122-0101200311001100-3022303322210031-2023223203302111-1223231110231233-1122103220022103-1112213110323303-3002232330120213"></a>

Type: `"object"`. list nested block, Optional.

Segment. Reference to Segment Object.

Upstream description:

Reference to Segment Object.

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

<a id="canonical-3111021122021100-1031323221023010-3113120320131202-2211301132131333-2103330220130320-1201301111220233-3310330323300002-3100010001113123"></a>

## Direct properties — refs / 133122003123 / 3

<a id="canonical-1303333120301303-1230231030131121-1323200123002233-1120323300233211-3321203333232330-1012221222320010-2100103202032023-2003010223202121"></a>

<a id="canonical-0133213103033101-2301312310123103-2212123232220212-2332300332012320-1120310103001031-3230331020030103-1200303200331121-0132221201332300"></a>

## kind property — refs / 133122003123 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-1202300030210222-3222310113200332-1031030123210223-0100010111033233-0130003310220133-1322233313100013-1211001222201130-1210221222301301"></a>

<a id="canonical-3111112310012110-1213333001112211-2112212122323013-3022132203221123-1222232033120012-1002031013300331-1031013323203222-0212023010100201"></a>

## name property — refs / 133122003123 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-3312203201010010-0133302102133231-3012310312011233-3001111123030112-0111303101302333-2223310103101303-2130212031111103-0231320023211331"></a>

<a id="canonical-1220233003112011-3021020321320030-3010302222022121-2012113110121133-2110303011132220-2100302301222310-0000103221232021-1310032321221010"></a>

## namespace property — refs / 133122003123 / 6

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

<a id="canonical-3011132330120302-2102022123210132-3030202013331002-0202120120300120-0331313100202031-2113223020003331-2233122103023022-3003303113201233"></a>

<a id="canonical-1020112022000223-2312223211032303-0303000131213033-1132203133310132-1123210133220313-0030331201331320-2122032313113210-2330102033001031"></a>

## tenant property — refs / 133122003123 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-2002001331321210-1122130110101100-3113012002331321-3302112101133133-0210303323122200-0330303320212133-2323212003333032-1010301030333113"></a>

<a id="canonical-3231033313003033-3200321013232300-3010220000212031-1331211203320311-2113103131300220-2203310001133313-2121002013332301-2232111131330110"></a>

## uid property — refs / 133122003123 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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

<a id="canonical-3121312233230233-3112123323122201-1011220302110232-2202303110323233-1330021003001331-3211003300113321-1033332312333303-1112003130300202"></a>

## Next pages — refs / 133122003123 / 9

- [rules.segment](resources--nat_policy--reference--group-001.md#canonical-2000331302003211-2112122330312222-0231223203303123-2331332301321123-1023103133311320-1120102103133013-3031202021101232-2321033122322230)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)

<a id="canonical-1203111110012112-3101113022030330-0033312021330011-0221130000132103-2001003103030300-2130101120332322-3210203132100203-1033032202202201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320330133333001-2112212002323331-2310010003203231-2011212000322023-1003320231012221-1121033003033211-1202000302130323-3023201001222130"></a>

## rules.virtual_network — virtual_network / 112120003202 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- rules.virtual_network

<a id="canonical-1002012311021032-1021333321012032-2230202231321302-0022021101022210-3232101102203121-0320102230102002-1112310123131332-1033233021312231"></a>

Type: `"object"`. single nested block, Optional.

Carries the reference to virtual network.

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

<a id="canonical-2003301303122213-1102110113000323-2111200311031303-2112010023103321-3122233103001302-0011200130300100-2002220111133302-3331212131120120"></a>

## Direct properties — virtual_network / 112120003202 / 3

- [refs](resources--nat_policy--reference--group-001.md#canonical-0230300110110022-3020301330020222-1230303132313022-2321212030320313-3102031232000203-3023210223230311-0222030322031213-3213133112201100): complete subsection reference.

<a id="canonical-2102102130331011-0110211211320311-2103300122132101-1032021332300020-3233323200322212-1303313111302103-3310002120332001-3133203322031010"></a>

## Next pages — virtual_network / 112120003202 / 4

- [rules.virtual_network.refs](resources--nat_policy--reference--group-001.md#canonical-0230300110110022-3020301330020222-1230303132313022-2321212030320313-3102031232000203-3023210223230311-0222030322031213-3213133112201100)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)

<a id="canonical-0230300110110022-3020301330020222-1230303132313022-2321212030320313-3102031232000203-3023210223230311-0222030322031213-3213133112201100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202031221101110-0201211301310300-3021100302210010-3231233102233103-0223301323212322-0210230112312112-0130310101110302-0000002333222010"></a>

## rules.virtual_network.refs — refs / 232332211201 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [rules.virtual_network](resources--nat_policy--reference--group-001.md#canonical-1203111110012112-3101113022030330-0033312021330011-0221130000132103-2001003103030300-2130101120332322-3210203132100203-1033032202202201)
- rules.virtual_network.refs

<a id="canonical-2222011321033210-1120130020123120-3002112101023023-0130313030222231-1312202103120313-0300002100120101-1332132132330120-2032022112011330"></a>

Type: `"object"`. list nested block, Optional.

Virtual Network Reference. Reference to virtual network.

Upstream description:

Reference to virtual network.

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
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

<a id="canonical-0223123100001312-1121033311112120-1031300230120311-3303113201313101-0211233131122322-3202033132322102-1020011021003113-2003212231022230"></a>

## Direct properties — refs / 232332211201 / 3

<a id="canonical-1113132123121211-0303111122032020-1131320332210001-1303103003210211-2011332210303312-1131021320201012-0012132321001002-1233232232110112"></a>

<a id="canonical-0322332012113321-2322121032113032-3102133201003011-3010331132130311-2301303111213131-3130211110001232-2331313003020111-3320212313232320"></a>

## kind property — refs / 232332211201 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-1200312121220021-3201311310302112-2321232000130330-1233101203002220-3313233313023000-2131111221101323-1212333230003323-3312102000110322"></a>

<a id="canonical-0211133313112120-3120102113301223-0130122013120111-3023002022031013-1313212031301301-1321012110013333-2123121030330332-0303333301221212"></a>

## name property — refs / 232332211201 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-1331302123001313-2303231032113013-1122311111311323-2021312322020330-2233233220130203-2221030231012300-2221320133320102-0122300011113033"></a>

<a id="canonical-3121102210213113-0213210221201011-2313102223323213-0010300221002213-0203123323013310-1023020101232320-1013300031230111-2330331031101220"></a>

## namespace property — refs / 232332211201 / 6

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

<a id="canonical-2311013331330322-2231300312021311-3111210002331320-0312130121300121-3121110213001321-1331032122033030-0300333330223232-1302000202031230"></a>

<a id="canonical-1000210110131112-0113331202133012-3221001023103022-2121102002103301-0233002020003301-3102032202001032-3101201232221111-1312203032213333"></a>

## tenant property — refs / 232332211201 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-0002211022010032-1131320322310020-3121012302202322-2332101220330210-0202213230110030-0210031020123233-3323010002321110-3032323123010312"></a>

<a id="canonical-1033220312232130-3002131232320033-0203211122010103-0212013200322231-1302202021220323-2003233133030123-2012031000002222-0303321023020013"></a>

## uid property — refs / 232332211201 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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

<a id="canonical-2213232030220123-0011101023132231-1003223220310023-0211331301331112-3321313230232310-2322233033131122-0232031100212232-2301002113101023"></a>

## Next pages — refs / 232332211201 / 9

- [rules.virtual_network](resources--nat_policy--reference--group-001.md#canonical-1203111110012112-3101113022030330-0033312021330011-0221130000132103-2001003103030300-2130101120332322-3210203132100203-1033032202202201)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)

<a id="canonical-3001223222011220-0003203113332212-3332203303130003-1310223301012333-2223122022201100-2321110020002221-0120100011321320-2020123030023230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130002132103222-2100213221013023-2231133233012003-2013001311133223-0300311030302111-1112011332233203-1031220230122132-0333311020311221"></a>

## site — site / 122332302111 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- site

<a id="canonical-2030103300120200-0101122100203031-3033132321322000-3002330010012232-2100220201321330-2233122102122111-2032023000031113-2222130223311331"></a>

Type: `"object"`. single nested block, Optional.

Site Reference Type. Reference to Site Object.

Upstream description:

Reference to Site Object.

Provider validators and defaults (from schema source):

```go
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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-3222230022333320-0122102233320322-0033113003212211-2021311312222131-3020333203003331-1311223330320332-1031100100021301-1032221112111320"></a>

## Direct properties — site / 122332302111 / 3

- [refs](resources--nat_policy--reference--group-001.md#canonical-3213022203101121-2022201233311000-0211320032011333-1131101222312203-1033201102013032-3312313120300222-2010221220230233-3023100210320103): complete subsection reference.

<a id="canonical-1310031032210322-0220212230200330-2032220301131133-2211000310010113-3230111220100003-2313031212012322-3011210220010030-3300231103131122"></a>

## Next pages — site / 122332302111 / 4

- [site.refs](resources--nat_policy--reference--group-001.md#canonical-3213022203101121-2022201233311000-0211320032011333-1131101222312203-1033201102013032-3312313120300222-2010221220230233-3023100210320103)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)

<a id="canonical-3213022203101121-2022201233311000-0211320032011333-1131101222312203-1033201102013032-3312313120300222-2010221220230233-3023100210320103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101230122221121-2303101122300203-2132030031031311-0121030200101201-3130310103313000-0101120333212122-1110010121002311-2001123010203033"></a>

## site.refs — refs / 020311201033 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [site](resources--nat_policy--reference--group-001.md#canonical-3001223222011220-0003203113332212-3332203303130003-1310223301012333-2223122022201100-2321110020002221-0120100011321320-2020123030023230)
- site.refs

<a id="canonical-2220012132012310-2211132013030100-1111123232003233-0211301020200002-1023002332331212-1122311130321300-1030020213330233-3132213231110303"></a>

Type: `"object"`. list nested block, Optional.

Site. Reference to Site Object.

Upstream description:

Reference to Site Object.

Receipt-pinned upstream constraints:

```json
{
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

Terraform syntax:

```terraform
refs {
  # Configure direct properties listed below.
}
```

<a id="canonical-2021222232003223-1131103002323031-1132333031010031-3110313331021232-0103002220112011-1022200103103230-1211132232121300-0123121110312011"></a>

## Direct properties — refs / 020311201033 / 3

<a id="canonical-0311112013112312-3310001003100122-0001311320213111-1023133301200203-1300221330233123-1020023300202322-1303013001310120-0001101011110022"></a>

<a id="canonical-3221031103032200-0230123132102210-1330120333301302-0121123002002320-0313132320300113-2330113201120132-0310231321032112-2331000301103131"></a>

## kind property — refs / 020311201033 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. "route")

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

<a id="canonical-1320033122131000-1021012033133222-3031313330001310-0210311211202220-3220231230001230-1013003300131111-0312210212213300-1203110101233022"></a>

<a id="canonical-0221131121231323-0321231102310003-2020200100312123-1003211002320031-0212210231123221-1001002302000100-3201200012020202-0011031312231211"></a>

## name property — refs / 020311201033 / 5

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-1113031312331003-3333331030333330-2322101131312222-0023300231132111-3312121000212110-0013120221202333-0131201033322001-0103200013031212"></a>

<a id="canonical-1301230020102203-1131103232313220-0210101311203033-1002212010203012-1301113031103333-0130113233330132-1321021312101002-1232221110210331"></a>

## namespace property — refs / 020311201033 / 6

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

<a id="canonical-2220311222103222-2321312313102321-1313310223012110-0101212021133213-1231002112033112-0011322330021323-0301322032021330-2301202301330123"></a>

<a id="canonical-3321303320203120-0101312131312221-2322002322201210-0102032313131220-3131331233031302-0030201010321121-3011320223221320-2211130010310200"></a>

## tenant property — refs / 020311201033 / 7

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-3102123331211033-2312333333213302-0002303203202010-1122100202001232-3030203013103311-0120032230111011-2011230100032000-0120131002301132"></a>

<a id="canonical-1012331003231003-0123210030200323-1300220102113110-0313030132101233-2131312313330231-2222010230010300-3120113122001130-0322111302130232"></a>

## uid property — refs / 020311201033 / 8

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Upstream description:

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

<a id="canonical-0031312333001311-0313011013302212-1312101013013103-3023000303322303-2202010323323331-0002312301033000-1302201202102022-3020120232023220"></a>

## Next pages — refs / 020311201033 / 9

- [site](resources--nat_policy--reference--group-001.md#canonical-3001223222011220-0003203113332212-3332203303130003-1310223301012333-2223122022201100-2321110020002221-0120100011321320-2020123030023230)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)

<a id="canonical-0332310313320102-1233203001320300-0120211031101101-2101310332223201-1311001110220122-3212202313031202-2103103310200032-1323023301223232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332312303101022-2020012200020312-0020323300120112-1312230033030303-2320112013030312-0223132322022313-2011102200221022-0003101213320112"></a>

## timeouts — timeouts / 301201333121 / 2

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- timeouts

<a id="canonical-1123000232032010-1111231111122300-2030311322101302-2321002131302202-2332021102231101-3221031131021123-0022222103212310-0102113200200102"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3213020203110111-2012120111220330-0223133332131000-1300131333121200-0122002131113202-1031021311301123-1101221113112020-3121113011100123"></a>

## Direct properties — timeouts / 301201333121 / 3

<a id="canonical-1011322020203030-1113033222120201-3232102302110222-1002103103033312-1231230320020332-2002030000002002-1230230122333103-1101312320012122"></a>

<a id="canonical-2223203033211320-2002010011022021-3202300010000011-3010010013301332-2212202300332012-2222310120333010-0301013020322130-1000013320000130"></a>

## create property — timeouts / 301201333121 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1310323321110022-1330102030300120-2223201133211122-1301322020122320-2033222302133133-2233331220303233-0230122100111232-3330010110012132"></a>

<a id="canonical-2320101202102201-2120132033121321-0322013001302332-0000223022311212-2123203000320300-1123111111113113-3000313033202203-2223311022231201"></a>

## delete property — timeouts / 301201333121 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0110023222221033-0002200102110023-3030031311231003-1203100202221020-3212302031021013-2310003322212122-1311330311302002-0133113213112023"></a>

<a id="canonical-2312233301212322-1123332132331021-2123232112001203-1222101113323311-1011323331200221-2321203301033320-2032330323233110-2111223032100100"></a>

## read property — timeouts / 301201333121 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0212231123312211-1221133020303323-3112212231232233-0322320320030013-2332112203022230-1110023101210220-2021110323000211-3213310020303212"></a>

<a id="canonical-2033202312233111-0213230022013111-2333132011133320-2310302321111131-1133000333000331-3130102120212020-2113221012113233-3220100032101233"></a>

## update property — timeouts / 301201333121 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0300123011303011-0220212230231031-0122303100222133-2022001322011312-1130210130233220-2213211323221210-1313122201013013-3130333123130022"></a>

## Next pages — timeouts / 301201333121 / 8

- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
