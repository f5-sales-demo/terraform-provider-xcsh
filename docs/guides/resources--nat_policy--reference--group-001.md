---
page_title: "xcsh_nat_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nat_policy reference."
---

# xcsh_nat_policy reference

<a id="canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- Property reference

<a id="canonical-1100330220223133-3200200031132322-3032321200020330-3103132231013112-0100322223201302-1001023101202301-3310320130100201-1020002111313130"></a>

### Direct properties for `xcsh_nat_policy`

<a id="canonical-0332122232121112-3120320001021122-0022132310302112-0200103031322111-2013121321132211-3023112311101201-3122232330021021-3032333111003232"></a>

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

<a id="canonical-1010220332122020-0321123122201311-2030212131022231-1203100301032031-3122111222230222-0020133200233210-0202222110131212-2321301003032221"></a>

<a id="canonical-2220303000002003-0303203230123321-1221020213133221-2312123221201121-1132022321003233-1332200231321311-0233231001231000-0211203122030113"></a>

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1113232232231210-0211213001121230-2213302001311213-3103213333330120-2213012301213213-1003211030012031-2011101332302120-3332233112320303"></a>

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

<a id="canonical-1320020300233233-3211123202312331-3132031102233022-2033120323023310-3233210120310303-2201310332300231-0323120031131032-1332130130312132"></a>

<a id="canonical-1013210003310012-2123123330030120-0110302131302211-1102322333110021-3130203010213011-2033313012232122-1110220230110331-1002030303331331"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1131231210012200-1132130213331303-2221111011223212-2213320111100111-0200002230200110-1232313120100210-3301110313002120-2203231002331200"></a>

<a id="canonical-1221112111023200-2011211033031021-0232233232113101-0222022013210313-3130203313202210-2332113210123022-1302021232313231-0301120210302301"></a>

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

<a id="canonical-1230212310130013-1000200231313100-2330023002323030-0131331202221130-3303323023332222-1000232311210112-3001313113122102-0310131033012110"></a>

<a id="canonical-1232133202223012-3221123300300020-3311120000112021-0011012013103221-3322320103100010-0102330333320200-2221130023230122-1310311103311233"></a>

#### `name` property

Type: `"string"`. Required.

Name of the NAT Policy. Must be unique within the namespace.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0001000133022231-0030100012202001-2322211022313112-2310010232331320-2132102103110200-1303000003202022-2301113231023232-2300222222000310"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the NAT Policy is created.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3031130121303113-3032230210201302-1010013120120321-0111022031323210-0011303021310031-2030113010320122-2233213033232313-2330303220330330"></a>

### All schema paths for `xcsh_nat_policy`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--nat_policy--reference--group-001.md#canonical-0332122232121112-3120320001021122-0022132310302112-0200103031322111-2013121321132211-3023112311101201-3122232330021021-3032333111003232) |
| `description` | [description](resources--nat_policy--reference--group-001.md#canonical-1010220332122020-0321123122201311-2030212131022231-1203100301032031-3122111222230222-0020133200233210-0202222110131212-2321301003032221) |
| `disable` | [disable](resources--nat_policy--reference--group-001.md#canonical-0131220133130202-3222231331100001-1100331323122211-0223301201332131-1203313132231101-0233232012320230-3021030010223300-3112203332020302) |
| `id` | [ID](resources--nat_policy--reference--group-001.md#canonical-1320020300233233-3211123202312331-3132031102233022-2033120323023310-3233210120310303-2201310332300231-0323120031131032-1332130130312132) |
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

<a id="canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules` properties

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- rules

<a id="canonical-1213230032120130-1310201032022230-2210021020301121-2022301310033333-0312210220102203-0002001213333330-0032031233202320-3311221233122321"></a>

Type: `"object"`. list nested block, Optional.

List of rules to apply under the NAT Policy. Rule that matches first would be applied.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1112322003023233-0002210111230221-2210110331032120-3031001232123312-3211302123210333-2132300011121333-1233033001232003-1033131330112121"></a>

### Direct properties for `rules`

- [action](resources--nat_policy--reference--group-001.md#canonical-1122221021030123-2131233221312132-0230001013023100-0130321320322011-2220032330303113-0023002023013103-3130110202211032-1022000313331212): complete subsection reference.

- [cloud_connect](resources--nat_policy--reference--group-001.md#canonical-1003100233302122-0331210123231312-3101222012030102-0011000230001133-0230002101033103-3330233311000231-2133020120103333-2030021301103211): complete subsection reference.

- [criteria](resources--nat_policy--reference--group-001.md#canonical-0211202132100101-0101202032220222-0011011112111300-3220020012220110-0202012111300120-3331002203310331-0223221133023011-2213211332320022): complete subsection reference.

- [disable_spec](resources--nat_policy--reference--group-001.md#canonical-2333111032320031-2222102332001330-1220210111201123-1011122013230003-3231123330013221-3103132010233220-0010221330011032-0200130112321100): complete subsection reference.

- [enable](resources--nat_policy--reference--group-001.md#canonical-3301200022202232-1013113223132333-1111013013320103-3201130310232122-1233023330321002-0020013000032003-0121010123113111-2203023302031312): complete subsection reference.

<a id="canonical-2012021302001111-1101030211031331-2322213001200132-3213300212322330-3231020300331200-0132101030012113-2123322132330211-0110320301221133"></a>

<a id="canonical-1110330220330002-0002102023000011-2113101333300330-1112123020013100-2200012213113032-3330001212202301-3120303302202032-1330212032230321"></a>

#### `rules.name` property

Type: `"string"`. Optional.

Name. Name of the Rule.

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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1122221021030123-2131233221312132-0230001013023100-0130321320322011-2220032330303113-0023002023013103-3130110202211032-1022000313331212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.action` properties

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
EnumExtractionComplete: false
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

<a id="canonical-0201100313020023-3201003211010102-0111123330023103-1210102020111132-3110102233113300-0122212120232223-2100233020233110-1202232123212312"></a>

### Direct properties for `rules.action`

- [dynamic](resources--nat_policy--reference--group-001.md#canonical-0221213332001013-1211231020320202-2110123331210011-3311233300323212-0323332201123333-3023103321300300-1013003221001113-1100103221120012): complete subsection reference.

<a id="canonical-1203200121312221-0103022130321003-0132020131202300-0113311020012122-3230130211100231-2303030333302100-2023002312200303-2021223113202212"></a>

<a id="canonical-2200223020031102-2003023010210112-2000033312232310-1103312310021313-1133020310111233-1103321120301201-2200023121012120-1202033323101032"></a>

#### `rules.action.virtual_cidr` property

Type: `"string"`. Optional.

Exclusive with \[dynamic\] Virtual Subnet NAT is static NAT that does a one-to-one translation
between the real source IP CIDR in the policy and the virtual CIDR in a bidirectional fashion. The
range of the real CIDR and virtual CIDRs should be the same (e.g. If the real CIDR has the CIDR
192.0.2.0/24, the virtual CIDR has 100.100.100.0/24.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0221213332001013-1211231020320202-2110123331210011-3311233300323212-0323332201123333-3023103321300300-1013003221001113-1100103221120012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.action.dynamic` properties

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [rules.action](resources--nat_policy--reference--group-001.md#canonical-1122221021030123-2131233221312132-0230001013023100-0130321320322011-2220032330303113-0023002023013103-3130110202211032-1022000313331212)
- rules.action.dynamic

<a id="canonical-2022332100311212-1103130032012132-1022200020013013-3111222320303112-1202002300320222-1121211231120102-2201003221113102-1321231211112233"></a>

Type: `"object"`. single nested block, Optional.

Dynamic Pool. Dynamic Pool Configuration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2223010222221211-0300002313231300-0100122212213201-3013133023311223-0220310030331320-3132311222021210-1220321113131232-3130000033301303"></a>

### Direct properties for `rules.action.dynamic`

- [elastic_ips](resources--nat_policy--reference--group-001.md#canonical-1320332012230123-1200023330031312-3010211321230103-0312220130203323-1021023110233100-1332012000310330-0322320302003223-2032133322302003): complete subsection reference.

- [pools](resources--nat_policy--reference--group-001.md#canonical-3101330020232121-1300120130322033-1120200032101230-0031011222111032-3100321032313012-3300112132101200-0333021110213112-2300123021320102): complete subsection reference.

<a id="canonical-1320332012230123-1200023330031312-3010211321230103-0312220130203323-1021023110233100-1332012000310330-0322320302003223-2032133322302003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.action.dynamic.elastic_ips` properties

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
elastic_ips {
  # Configure direct properties listed below.
}
```

<a id="canonical-0311221232311232-1211320233102200-0221213130013211-0013102121123320-3313330213313322-1113123330303113-0133030110110301-0022031321222101"></a>

### Direct properties for `rules.action.dynamic.elastic_ips`

- [refs](resources--nat_policy--reference--group-001.md#canonical-3021303223113021-3220020210100310-0021033122120022-2201011001330101-1320221000010220-0033132112200311-3031202133011201-1033121211131023): complete subsection reference.

<a id="canonical-3021303223113021-3220020210100310-0021033122120022-2201011001330101-1320221000010220-0033132112200311-3031202133011201-1033121211131023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.action.dynamic.elastic_ips.refs` properties

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3121331233303003-0313310010131030-1030121221321220-2310130222211220-1202021012231201-3320032011333323-1121221301222111-2201202100231100"></a>

### Direct properties for `rules.action.dynamic.elastic_ips.refs`

<a id="canonical-1113212230003330-0031300331133012-3200033020012322-2330123011203102-1121220200102020-2132103212231022-3021013323231102-1211130323313032"></a>

#### `rules.action.dynamic.elastic_ips.refs.kind` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0210213302111223-1010211022303300-1232220100013203-0232003230211023-2000222333132210-0013112003302233-2121222102032023-3200133122311311"></a>

#### `rules.action.dynamic.elastic_ips.refs.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1001232232012223-0332002210013210-0210022123021222-2121202333320200-1202223323213000-3203222313103110-1100323313331210-2331020013012200"></a>

#### `rules.action.dynamic.elastic_ips.refs.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1123313313311231-3133321323313120-0330021200231022-1222000322113310-0023220313111032-2020233111003022-2203033313333320-0011222332233333"></a>

#### `rules.action.dynamic.elastic_ips.refs.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2203311331131103-0010220320213111-1311003020300101-1003231012020121-1312022201231123-1021312211123013-1121331310212230-3003022222321310"></a>

#### `rules.action.dynamic.elastic_ips.refs.uid` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3101330020232121-1300120130322033-1120200032101230-0031011222111032-3100321032313012-3300112132101200-0333021110213112-2300123021320102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.action.dynamic.pools` properties

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

<a id="canonical-1223202332231000-1302321230010110-0100300031230002-1213103032000333-0231030033233302-0003220330303003-3201013111023222-0031231021002100"></a>

### Direct properties for `rules.action.dynamic.pools`

<a id="canonical-3131030022302230-2321000123112123-0002113101103322-2103131111001120-1130201203233302-2221232123033023-0031131321022211-2213110303201200"></a>

#### `rules.action.dynamic.pools.prefixes` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1003100233302122-0331210123231312-3101222012030102-0011000230001133-0230002101033103-3330233311000231-2133020120103333-2030021301103211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.cloud_connect` properties

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- rules.cloud_connect

<a id="canonical-1111312031133000-0013212212212233-3310210223200331-0313021033103133-1121003013233000-3300330003220212-2313222100113121-3030310310330302"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for cloud connect.

Additional upstream details:

Reference to Cloud connect Object.

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
cloud_connect {
  # Configure direct properties listed below.
}
```

<a id="canonical-0302221310333212-1312212000201032-3322100220123012-3020000021322230-3002001102321330-3330021303330033-1022001302030102-1021120012020301"></a>

### Direct properties for `rules.cloud_connect`

- [refs](resources--nat_policy--reference--group-001.md#canonical-0301332012010322-0232321001330023-0201001213213300-2021111320100201-0303303133232022-2201221302102232-2123332010200222-2130300311213023): complete subsection reference.

<a id="canonical-0301332012010322-0232321001330023-0201001213213300-2021111320100201-0303303133232022-2201221302102232-2123332010200222-2130300311213023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.cloud_connect.refs` properties

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [rules.cloud_connect](resources--nat_policy--reference--group-001.md#canonical-1003100233302122-0331210123231312-3101222012030102-0011000230001133-0230002101033103-3330233311000231-2133020120103333-2030021301103211)
- rules.cloud_connect.refs

<a id="canonical-1311001123333020-0003323000131120-2023112330112021-0023023323332100-1220122112322101-1232131110302212-1102210202213233-0102000022110010"></a>

Type: `"object"`. list nested block, Optional.

Cloud Connect. Reference to Cloud Connect Object.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3202310201321113-2200323232220031-3113321320032011-0220312010103322-3311311332300011-1320203023221333-0031012111023213-0311311020022303"></a>

### Direct properties for `rules.cloud_connect.refs`

<a id="canonical-2310102002130212-0032001312002113-1222232213203222-2312302200000113-1221231333123301-0330003303331132-3123000111103001-1220303011302312"></a>

#### `rules.cloud_connect.refs.kind` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2213032213023311-1221033002321000-1222030122322310-3323301100010131-0222121010112223-2002113123113200-1223101130013313-2201322031003312"></a>

#### `rules.cloud_connect.refs.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1200210111033011-3123022221021031-3122221032310310-2211011130223120-1133000332101132-1310302102300333-3032301000130312-3210113302312020"></a>

#### `rules.cloud_connect.refs.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1322323333223211-2201210231011001-2201311013102223-2310312133023330-0131201200301330-0021023303320211-3222000023110320-3232322101313133"></a>

#### `rules.cloud_connect.refs.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1221202131002222-2231210222020321-2132222030003333-2301120021102020-3000103311332200-3230120002031100-0311302320122333-0121012012020201"></a>

#### `rules.cloud_connect.refs.uid` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0211202132100101-0101202032220222-0011011112111300-3220020012220110-0202012111300120-3331002203310331-0223221133023011-2213211332320022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.criteria` properties

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
EnumExtractionComplete: false
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

<a id="canonical-0301120203111013-3123033030333320-0113332321323300-2331001223120021-2210000202301021-3300112113001201-1023330222021211-3233210023313120"></a>

### Direct properties for `rules.criteria`

- [any](resources--nat_policy--reference--group-001.md#canonical-3322130111322231-2211210112122132-1321020032320211-3030100213033221-1010202021312230-2202001320230330-0221301011133123-3300330103301332): complete subsection reference.

<a id="canonical-3023321213011103-1213233303032210-2233300212311212-3320113031112023-2033333213311012-2003101212100201-3111203021232020-0032111201121202"></a>

<a id="canonical-0103033221003320-1132322122132200-0132222233332302-0300233321302031-2221021003132321-1321230323031303-3031013001333010-0110122103223233"></a>

#### `rules.criteria.destination_cidr` property

Type: `["list", "string"]`. Optional.

Destination IP. Destination IP of the packet to match.

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

<a id="canonical-1320203300210303-0032031202302210-3131001131232000-1203301133223122-0220100202320301-1202101110330002-2101200212130003-1012222012220211"></a>

#### `rules.criteria.source_cidr` property

Type: `["list", "string"]`. Optional.

Source IP. Source IP of the packet to match.

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

<a id="canonical-3322130111322231-2211210112122132-1321020032320211-3030100213033221-1010202021312230-2202001320230330-0221301011133123-3300330103301332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.criteria.any` properties

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-0211202132100101-0101202032220222-0011011112111300-3220020012220110-0202012111300120-3331002203310331-0223221133023011-2213211332320022)
- rules.criteria.any

<a id="canonical-0133132230022201-1030232020331001-3211112113202332-0002113111111302-1022033212300333-0203031212322311-3313321112230031-2031023113312311"></a>

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
any = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0311131033003010-2323213031003000-3111030000001003-1120111033230110-3331100233220022-3213331013230301-3111231201013101-3231122322001100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.criteria.icmp` properties

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-0211202132100101-0101202032220222-0011011112111300-3220020012220110-0202012111300120-3331002203310331-0223221133023011-2213211332320022)
- rules.criteria.icmp

<a id="canonical-3221222003032231-1132232102223023-2000112020303312-1033021301313112-0013331010212211-3133120103101232-3302220011012303-2302312230101322"></a>

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
icmp = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0111331033301131-2121203023323213-0031131103130323-2033322133233230-0111333322203113-2030022002031300-1111002013323031-0201223132031321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.criteria.site_local_inside_network` properties

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-0211202132100101-0101202032220222-0011011112111300-3220020012220110-0202012111300120-3331002203310331-0223221133023011-2213211332320022)
- rules.criteria.site_local_inside_network

<a id="canonical-0102112121222011-0231101310300032-3213100022112212-3101331233111122-2201102212112202-0133120220201231-2131021002023221-2023211013030202"></a>

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

<a id="canonical-3001002101100121-1103201310220210-0313020302130112-1003032003320230-3230010021213223-3301223012231123-2312011201231033-3102101023203103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.criteria.site_local_network` properties

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [rules.criteria](resources--nat_policy--reference--group-001.md#canonical-0211202132100101-0101202032220222-0011011112111300-3220020012220110-0202012111300120-3331002203310331-0223221133023011-2213211332320022)
- rules.criteria.site_local_network

<a id="canonical-3203013202301233-0012223313333001-2103302131222332-3321232030003221-0021233102331120-1331101223312330-1103022303121133-3020333311200333"></a>

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

<a id="canonical-2301233133233231-1233311230111113-1300112232202320-0210231103202222-3030232230030201-0000232313210120-1011330101003002-1102333311310033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.criteria.tcp` properties

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

<a id="canonical-3000213213332331-2313200232111213-3213031321122220-2100300210111021-1313022102121321-0132311012100313-0331203111300322-0222021033003131"></a>

### Direct properties for `rules.criteria.tcp`

- [destination_port](resources--nat_policy--reference--group-001.md#canonical-1122021131311023-1130133130321113-0303003123323203-2002323012121021-1221233233333311-2033011120113223-1303311330222230-1000300030200112): complete subsection reference.

- [source_port](resources--nat_policy--reference--group-001.md#canonical-0012333302323200-0202010231202231-0101230233201311-0133120031031103-2102021330123213-1121101113130303-0012332310201231-3332303022132030): complete subsection reference.

<a id="canonical-1122021131311023-1130133130321113-0303003123323203-2002323012121021-1221233233333311-2033011120113223-1303311330222230-1000300030200112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.criteria.tcp.destination_port` properties

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
EnumExtractionComplete: false
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

<a id="canonical-1123111013301320-3121033021032031-3102331113223133-3331123003303131-2300112310312330-2000323321012131-0223313311203111-0002231211333130"></a>

### Direct properties for `rules.criteria.tcp.destination_port`

- [no_port_match](resources--nat_policy--reference--group-001.md#canonical-0023103200122233-2030011202213110-0331310103333012-0121301122133133-1313110212003031-2003012021111102-0311333212121033-2200022221020302): complete subsection reference.

<a id="canonical-3102213323303103-3323102302201301-1222122300131211-0003201122113001-1222030010211321-3002200103231023-0022011301012311-2101313202012331"></a>

<a id="canonical-1321322320121120-3032301310230123-1301010030013210-2131331122002320-3231100323023302-2200231201213002-1132333302002002-2030331333123333"></a>

#### `rules.criteria.tcp.destination_port.port` property

Type: `"number"`. Optional.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0301210302032221-3102300211303333-2132221020333033-1113313112113312-0032332322331303-0100221021200020-1113210102111100-0032101010232310"></a>

#### `rules.criteria.tcp.destination_port.port_ranges` property

Type: `"string"`. Optional.

Exclusive with \[no\_port\_match port\] Port range to match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0023103200122233-2030011202213110-0331310103333012-0121301122133133-1313110212003031-2003012021111102-0311333212121033-2200022221020302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.criteria.tcp.destination_port.no_port_match` properties

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
no_port_match = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0012333302323200-0202010231202231-0101230233201311-0133120031031103-2102021330123213-1121101113130303-0012332310201231-3332303022132030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.criteria.tcp.source_port` properties

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
EnumExtractionComplete: false
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

<a id="canonical-3122230133320223-3021120120310033-3020003013010333-1113002023310222-2030120231101030-0320233200213302-2320032323000130-0310132310213130"></a>

### Direct properties for `rules.criteria.tcp.source_port`

- [no_port_match](resources--nat_policy--reference--group-001.md#canonical-0131000301210222-0102320313012332-0232123211210021-3111322001321320-2323113111133310-2201111330222033-2322303110110020-0211031332011311): complete subsection reference.

<a id="canonical-0133301301033003-1023022100221302-0013322012103223-0113111211031123-2213120133011032-1120033203322020-2233322322320033-3010101320233231"></a>

<a id="canonical-0030112001012212-2202313130320222-0133110211312121-2120123211302320-0002020232101211-1030121221213312-0320231102021101-3113132003233032"></a>

#### `rules.criteria.tcp.source_port.port` property

Type: `"number"`. Optional.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1021222303213332-2213102002020013-3013321231203111-0012103310022322-3233030122333012-2110303111032021-0321100003313022-3112313203300213"></a>

#### `rules.criteria.tcp.source_port.port_ranges` property

Type: `"string"`. Optional.

Exclusive with \[no\_port\_match port\] Port range to match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0131000301210222-0102320313012332-0232123211210021-3111322001321320-2323113111133310-2201111330222033-2322303110110020-0211031332011311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.criteria.tcp.source_port.no_port_match` properties

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
no_port_match = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2211331033232333-2100001210303031-2131003002002112-0301201011323211-0132122121123001-3133302122301131-3113232033232301-1020210102220213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.criteria.udp` properties

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

<a id="canonical-1013012230101021-2001102001133232-3103200103130033-0302022333303300-0100333000301033-1223130023000322-2301311123011021-1311021323230132"></a>

### Direct properties for `rules.criteria.udp`

- [destination_port](resources--nat_policy--reference--group-001.md#canonical-3130130033112021-3211102103000131-0212331101101001-3212003010202320-0030001002100110-2200011123222321-0202123221211202-0132131211322021): complete subsection reference.

- [source_port](resources--nat_policy--reference--group-001.md#canonical-0122012203231130-0202333101303112-0020300323310323-1023103331210023-0122120331233113-2312223111231220-2202230111212321-3201321202121231): complete subsection reference.

<a id="canonical-3130130033112021-3211102103000131-0212331101101001-3212003010202320-0030001002100110-2200011123222321-0202123221211202-0132131211322021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.criteria.udp.destination_port` properties

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
EnumExtractionComplete: false
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

<a id="canonical-3113203132210101-0011012121211202-2302030223112203-1002000203001133-3121223012211303-2012103221130223-3123203331320003-3233133102012013"></a>

### Direct properties for `rules.criteria.udp.destination_port`

- [no_port_match](resources--nat_policy--reference--group-001.md#canonical-1023202301200301-2012021332333110-2111222121000301-1212113003123122-1313200211112022-3212022321033131-3320130213113231-2333021020301213): complete subsection reference.

<a id="canonical-0332001133021312-3331102313121022-3312030330113303-0001303331203230-2001132022101300-0231100101123200-1322213201032022-1131230000112111"></a>

<a id="canonical-3130220232223213-0201133113310201-2232202211110221-2230031123022212-2012300112202110-0323123002223232-1000201133303220-0030122332300111"></a>

#### `rules.criteria.udp.destination_port.port` property

Type: `"number"`. Optional.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3013110132123333-0133021002320101-0031013001333000-0230101201123203-2030132223013023-2021102233132002-1300001321100012-2323012121233203"></a>

#### `rules.criteria.udp.destination_port.port_ranges` property

Type: `"string"`. Optional.

Exclusive with \[no\_port\_match port\] Port range to match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1023202301200301-2012021332333110-2111222121000301-1212113003123122-1313200211112022-3212022321033131-3320130213113231-2333021020301213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.criteria.udp.destination_port.no_port_match` properties

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
no_port_match = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0122012203231130-0202333101303112-0020300323310323-1023103331210023-0122120331233113-2312223111231220-2202230111212321-3201321202121231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.criteria.udp.source_port` properties

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
EnumExtractionComplete: false
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

<a id="canonical-1123113301021130-2220321321013221-3222031100232333-1021002233301312-0121101113310002-1221232301113032-0132330132010032-2111203122102302"></a>

### Direct properties for `rules.criteria.udp.source_port`

- [no_port_match](resources--nat_policy--reference--group-001.md#canonical-3000122033003101-0033322313222201-2101220213110032-3230123310221113-3313100332223012-3100130302003103-1310202330033011-1311013113120332): complete subsection reference.

<a id="canonical-3231201210302201-2002123223123203-2013101332111132-2100121303103200-3122322202021022-2311323312021003-0302120113311131-2132030231332103"></a>

<a id="canonical-2302002131322011-0122320001021120-1013020201123010-0121023003032200-2013302032212033-2231033230302030-3220113010030212-1210120311003130"></a>

#### `rules.criteria.udp.source_port.port` property

Type: `"number"`. Optional.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1102230030223011-1001201110023330-3301102231003231-2032332321323213-3110123210013102-1203202022222110-3023030203000332-3232013300011133"></a>

#### `rules.criteria.udp.source_port.port_ranges` property

Type: `"string"`. Optional.

Exclusive with \[no\_port\_match port\] Port range to match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3000122033003101-0033322313222201-2101220213110032-3230123310221113-3313100332223012-3100130302003103-1310202330033011-1311013113120332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.criteria.udp.source_port.no_port_match` properties

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
no_port_match = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2333111032320031-2222102332001330-1220210111201123-1011122013230003-3231123330013221-3103132010233220-0010221330011032-0200130112321100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.disable_spec` properties

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3301200022202232-1013113223132333-1111013013320103-3201130310232122-1233023330321002-0020013000032003-0121010123113111-2203023302031312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.enable` properties

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- rules.enable

<a id="canonical-2111322330111130-1222222221011221-0003131112130332-3013311202332230-2022213322300203-2102101223003130-1000103202223021-3012301120302301"></a>

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
enable = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0003122032201330-0222033131003300-3123021011023102-3103311200332223-1132130323200320-1223032020223203-0130321232310300-0121313320300220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.node_interface` properties

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

<a id="canonical-2023202232310112-1021111212030230-3120313022030300-0130112313212302-3311303021002312-1301211310202333-3332322230200300-2331201221223122"></a>

### Direct properties for `rules.node_interface`

- [list](resources--nat_policy--reference--group-001.md#canonical-2122022232302031-2333032101032021-0231201133031013-2321111030231130-2101303131110130-2003002331021120-0231312302331333-3110002222230032): complete subsection reference.

<a id="canonical-2122022232302031-2333032101032021-0231201133031013-2321111030231130-2101303131110130-2003002331021120-0231312302331333-3110002222230032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.node_interface.list` properties

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3333322133320031-2202323011010333-3231033022202301-2313110302031022-3323010023112231-1123123200131333-0203220200031332-1110120220133113"></a>

### Direct properties for `rules.node_interface.list`

- [interface](resources--nat_policy--reference--group-001.md#canonical-0212201101030223-3030213222002202-1030030320322233-3210310211222203-3132131210320231-3000222200032223-0113132133301103-2011011231112332): complete subsection reference.

<a id="canonical-2322312222011333-3303132232310100-2000012231230003-2322012002210320-2200131030321012-2003202031311323-0111132332003132-3300312233110021"></a>

<a id="canonical-0320023312203312-0030021220113312-1323122232131132-1322223031113333-2023322123232223-3102022200212310-0101133222213323-0332230232312232"></a>

#### `rules.node_interface.list.node` property

Type: `"string"`. Optional.

Node. Node name on this site.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0212201101030223-3030213222002202-1030030320322233-3210310211222203-3132131210320231-3000222200032223-0113132133301103-2011011231112332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.node_interface.list.interface` properties

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0313011002003233-2331222012300322-2113333200112110-2221222132203030-2033300130122110-3312011000220332-2101313313203211-1020030131023020"></a>

### Direct properties for `rules.node_interface.list.interface`

<a id="canonical-2222101030310100-0212331213312030-2033320312022233-2013033230131112-1002001010322200-0323311030020110-2032033012112000-0313010330301112"></a>

#### `rules.node_interface.list.interface.kind` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1121310312120313-0220202001122030-2133302130123203-0000230013022302-1223103121302333-1313103111013113-0011212123010330-2121031201113033"></a>

#### `rules.node_interface.list.interface.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1110131020120010-2230003110311100-3030213310023223-0020312122300313-3322030103121120-3333030322133323-0010320120332320-0303331002331003"></a>

#### `rules.node_interface.list.interface.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2100021310202310-1023301122331130-2302130200203101-0330130030011003-0013331201201121-0302211010322313-2121333200221223-2002300210010031"></a>

#### `rules.node_interface.list.interface.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3311213121210030-2222302303110112-3112113033303233-3033310023330012-3020232031201121-3121011131113311-0233113202313320-1000130030303132"></a>

#### `rules.node_interface.list.interface.uid` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2000331302003211-2112122330312222-0231223203303123-2331332301321123-1023103133311320-1120102103133013-3031202021101232-2321033122322230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.segment` properties

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- rules.segment

<a id="canonical-1002003323223203-1013303333103203-0031010010033300-2212110101202121-1300320122201213-3231233212331122-1212101101011202-0220322331222011"></a>

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

<a id="canonical-2003001221301112-0022102222013102-0303020120002113-1030110021112021-2003210003222112-2001221000311133-3032312033313221-1120333001003323"></a>

### Direct properties for `rules.segment`

- [refs](resources--nat_policy--reference--group-001.md#canonical-1103103320230130-2322002110112311-3333012233010000-2120130121230220-1210210322000211-0322321211012203-3223233031210023-0210021200001032): complete subsection reference.

<a id="canonical-1103103320230130-2322002110112311-3333012233010000-2120130121230220-1210210322000211-0322321211012203-3223233031210023-0210021200001032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.segment.refs` properties

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [rules.segment](resources--nat_policy--reference--group-001.md#canonical-2000331302003211-2112122330312222-0231223203303123-2331332301321123-1023103133311320-1120102103133013-3031202021101232-2321033122322230)
- rules.segment.refs

<a id="canonical-0332133133132122-0101200311001100-3022303322210031-2023223203302111-1223231110231233-1122103220022103-1112213110323303-3002232330120213"></a>

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3202230212003331-2222002112113003-2100320301111320-3012312033332003-3101130212223200-2300303222210001-2103013033010221-2112321003232013"></a>

### Direct properties for `rules.segment.refs`

<a id="canonical-1303333120301303-1230231030131121-1323200123002233-1120323300233211-3321203333232330-1012221222320010-2100103202032023-2003010223202121"></a>

#### `rules.segment.refs.kind` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3111021122021100-1031323221023010-3113120320131202-2211301132131333-2103330220130320-1201301111220233-3310330323300002-3100010001113123"></a>

#### `rules.segment.refs.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0133213103033101-2301312310123103-2212123232220212-2332300332012320-1120310103001031-3230331020030103-1200303200331121-0132221201332300"></a>

#### `rules.segment.refs.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3111112310012110-1213333001112211-2112212122323013-3022132203221123-1222232033120012-1002031013300331-1031013323203222-0212023010100201"></a>

#### `rules.segment.refs.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1220233003112011-3021020321320030-3010302222022121-2012113110121133-2110303011132220-2100302301222310-0000103221232021-1310032321221010"></a>

#### `rules.segment.refs.uid` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1203111110012112-3101113022030330-0033312021330011-0221130000132103-2001003103030300-2130101120332322-3210203132100203-1033032202202201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.virtual_network` properties

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

<a id="canonical-0320330133333001-2112212002323331-2310010003203231-2011212000322023-1003320231012221-1121033003033211-1202000302130323-3023201001222130"></a>

### Direct properties for `rules.virtual_network`

- [refs](resources--nat_policy--reference--group-001.md#canonical-0230300110110022-3020301330020222-1230303132313022-2321212030320313-3102031232000203-3023210223230311-0222030322031213-3213133112201100): complete subsection reference.

<a id="canonical-0230300110110022-3020301330020222-1230303132313022-2321212030320313-3102031232000203-3023210223230311-0222030322031213-3213133112201100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `rules.virtual_network.refs` properties

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [rules](resources--nat_policy--reference--group-001.md#canonical-2012000321231320-0111310003001103-1103213311023201-2231113112111332-1323003130112121-2020001112021003-1133301220231213-3111322002030132)
- [rules.virtual_network](resources--nat_policy--reference--group-001.md#canonical-1203111110012112-3101113022030330-0033312021330011-0221130000132103-2001003103030300-2130101120332322-3210203132100203-1033032202202201)
- rules.virtual_network.refs

<a id="canonical-2222011321033210-1120130020123120-3002112101023023-0130313030222231-1312202103120313-0300002100120101-1332132132330120-2032022112011330"></a>

Type: `"object"`. list nested block, Optional.

Virtual Network Reference. Reference to virtual network.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3202031221101110-0201211301310300-3021100302210010-3231233102233103-0223301323212322-0210230112312112-0130310101110302-0000002333222010"></a>

### Direct properties for `rules.virtual_network.refs`

<a id="canonical-1113132123121211-0303111122032020-1131320332210001-1303103003210211-2011332210303312-1131021320201012-0012132321001002-1233232232110112"></a>

#### `rules.virtual_network.refs.kind` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0223123100001312-1121033311112120-1031300230120311-3303113201313101-0211233131122322-3202033132322102-1020011021003113-2003212231022230"></a>

#### `rules.virtual_network.refs.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0322332012113321-2322121032113032-3102133201003011-3010331132130311-2301303111213131-3130211110001232-2331313003020111-3320212313232320"></a>

#### `rules.virtual_network.refs.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0211133313112120-3120102113301223-0130122013120111-3023002022031013-1313212031301301-1321012110013333-2123121030330332-0303333301221212"></a>

#### `rules.virtual_network.refs.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3121102210213113-0213210221201011-2313102223323213-0010300221002213-0203123323013310-1023020101232320-1013300031230111-2330331031101220"></a>

#### `rules.virtual_network.refs.uid` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3001223222011220-0003203113332212-3332203303130003-1310223301012333-2223122022201100-2321110020002221-0120100011321320-2020123030023230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site` properties

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- site

<a id="canonical-2030103300120200-0101122100203031-3033132321322000-3002330010012232-2100220201321330-2233122102122111-2032023000031113-2222130223311331"></a>

Type: `"object"`. single nested block, Optional.

Site Reference Type. Reference to Site Object.

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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-2130002132103222-2100213221013023-2231133233012003-2013001311133223-0300311030302111-1112011332233203-1031220230122132-0333311020311221"></a>

### Direct properties for `site`

- [refs](resources--nat_policy--reference--group-001.md#canonical-3213022203101121-2022201233311000-0211320032011333-1131101222312203-1033201102013032-3312313120300222-2010221220230233-3023100210320103): complete subsection reference.

<a id="canonical-3213022203101121-2022201233311000-0211320032011333-1131101222312203-1033201102013032-3312313120300222-2010221220230233-3023100210320103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site.refs` properties

Breadcrumbs:

- [xcsh_nat_policy](../resources/nat_policy.md#canonical-1331123111210011-3210022110003110-2000223220101101-3221312231313300-1121211331030213-2203120232111121-1210202302022302-1310120003320310)
- [Property reference](resources--nat_policy--reference--group-001.md#canonical-3020222133133201-3022303223012310-1120231003312300-0101031322031013-1020101232211220-3311111200100321-3011121133111320-0102131300033201)
- [site](resources--nat_policy--reference--group-001.md#canonical-3001223222011220-0003203113332212-3332203303130003-1310223301012333-2223122022201100-2321110020002221-0120100011321320-2020123030023230)
- site.refs

<a id="canonical-2220012132012310-2211132013030100-1111123232003233-0211301020200002-1023002332331212-1122311130321300-1030020213330233-3132213231110303"></a>

Type: `"object"`. list nested block, Optional.

Site. Reference to Site Object.

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

<a id="canonical-1101230122221121-2303101122300203-2132030031031311-0121030200101201-3130310103313000-0101120333212122-1110010121002311-2001123010203033"></a>

### Direct properties for `site.refs`

<a id="canonical-0311112013112312-3310001003100122-0001311320213111-1023133301200203-1300221330233123-1020023300202322-1303013001310120-0001101011110022"></a>

#### `site.refs.kind` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2021222232003223-1131103002323031-1132333031010031-3110313331021232-0103002220112011-1022200103103230-1211132232121300-0123121110312011"></a>

#### `site.refs.name` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3221031103032200-0230123132102210-1330120333301302-0121123002002320-0313132320300113-2330113201120132-0310231321032112-2331000301103131"></a>

#### `site.refs.namespace` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0221131121231323-0321231102310003-2020200100312123-1003211002320031-0212210231123221-1001002302000100-3201200012020202-0011031312231211"></a>

#### `site.refs.tenant` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1301230020102203-1131103232313220-0210101311203033-1002212010203012-1301113031103333-0130113233330132-1321021312101002-1232221110210331"></a>

#### `site.refs.uid` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0332310313320102-1233203001320300-0120211031101101-2101310332223201-1311001110220122-3212202313031202-2103103310200032-1323023301223232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

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

<a id="canonical-3332312303101022-2020012200020312-0020323300120112-1312230033030303-2320112013030312-0223132322022313-2011102200221022-0003101213320112"></a>

### Direct properties for `timeouts`

<a id="canonical-1011322020203030-1113033222120201-3232102302110222-1002103103033312-1231230320020332-2002030000002002-1230230122333103-1101312320012122"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1310323321110022-1330102030300120-2223201133211122-1301322020122320-2033222302133133-2233331220303233-0230122100111232-3330010110012132"></a>

<a id="canonical-3213020203110111-2012120111220330-0223133332131000-1300131333121200-0122002131113202-1031021311301123-1101221113112020-3121113011100123"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-0110023222221033-0002200102110023-3030031311231003-1203100202221020-3212302031021013-2310003322212122-1311330311302002-0133113213112023"></a>

<a id="canonical-2223203033211320-2002010011022021-3202300010000011-3010010013301332-2212202300332012-2222310120333010-0301013020322130-1000013320000130"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0212231123312211-1221133020303323-3112212231232233-0322320320030013-2332112203022230-1110023101210220-2021110323000211-3213310020303212"></a>

<a id="canonical-2320101202102201-2120132033121321-0322013001302332-0000223022311212-2123203000320300-1123111111113113-3000313033202203-2223311022231201"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
