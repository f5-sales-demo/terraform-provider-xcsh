---
page_title: "xcsh_network_policy_view reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_policy_view reference."
---

# xcsh_network_policy_view reference

<a id="canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133032001232203-1021201002213222-2010122110102212-1232003032211333-0003122320011123-0312221032020231-1001111202301012-3022001211200322"></a>

## Property reference — Property reference / 211300031113 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- Property reference

<a id="canonical-3302203022332001-2111001222001001-3312210330322200-2001013202232310-0222211322220233-3132313223321302-3211231233201121-1023200122112210"></a>

## Direct properties — Property reference / 211300031113 / 3

<a id="canonical-0313321231302321-3200023202001230-1132311000313213-3201333212230113-2301300002212202-2120002230200321-1022101203303103-3330103002201230"></a>

<a id="canonical-1311133030023000-3220131012112113-3103002100030010-1030100223002102-3032212323222110-0030011232313223-0023200232312033-0122101222032223"></a>

## annotations property — Property reference / 211300031113 / 4

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

<a id="canonical-2330031223021331-3231123023031230-3322212200003312-0133020211201013-3221030103131211-3002010202332011-2011300023311121-3213310023301202"></a>

<a id="canonical-2313100213210123-0212000303111221-0221202120110131-1301000232123212-3123101311112200-1323010030310132-2303012013103100-0312111210103333"></a>

## description property — Property reference / 211300031113 / 5

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

<a id="canonical-3202021133132213-3002011203223012-3100313333300301-3032212031330321-1213111020201320-0013033032030332-0300020331310322-0121031023023012"></a>

<a id="canonical-2201022220203331-2320210301132003-3332121110201100-2320203202101000-2323031023112213-3301211220303300-0113323303002100-1312211030030003"></a>

## disable property — Property reference / 211300031113 / 6

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

- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213): complete subsection reference.

- [endpoint](resources--network_policy_view--reference--group-001.md#canonical-3300312010233011-1332123212033122-3330303200202222-0112211301022200-0121210323131100-1100322331002021-3232201003132022-3201310313012123): complete subsection reference.

<a id="canonical-3031030011210200-3320332030221122-3121323301101120-3203200310102113-0111031230213323-2300003123032233-1032221322020032-2022312323323212"></a>

<a id="canonical-1223121110300211-3123012120120132-1123320213203003-3201021030331110-0331311123112132-1131301101101031-3213313310303333-0233323122313221"></a>

## ID property — Property reference / 211300031113 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013): complete subsection reference.

<a id="canonical-0103002231123333-2123121123233122-2011132213311221-3301223222020313-3333032233032133-3211023033313302-0133020003213322-2201000000103202"></a>

<a id="canonical-0000000223332021-2221112203102210-3102130210130230-1033012212333033-3022310333323103-0131130233201232-0333011121030222-2330333020302010"></a>

## labels property — Property reference / 211300031113 / 8

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

<a id="canonical-0003110202032023-3210031220230011-3203301222113212-3123322003132000-3132130000111322-0032303113010020-1113212021123011-0011130211311002"></a>

<a id="canonical-2200121030103221-1330223032100321-1033020310120223-1032123101323230-2131220120033121-0130132121032000-0233211301112012-0031212231313132"></a>

## name property — Property reference / 211300031113 / 9

Type: `"string"`. Required.

Name of the Network Policy View. Must be unique within the namespace.

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

<a id="canonical-2103030022131000-1203230322101231-1332312330002121-2301231121000022-1320111323312323-0331332302231322-3300112322021031-3330221330033023"></a>

<a id="canonical-0132230203312321-2213111102112232-1322320310133213-1102100102000312-2331211203001232-2211033303012003-2121300020011102-2211312232213032"></a>

## namespace property — Property reference / 211300031113 / 10

Type: `"string"`. Optional, Computed.

Namespace for the Network Policy View. The F5 XC API restricts this resource to the system
namespace; it defaults to that value and may be omitted.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Default: stringdefault.StaticString("system")
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

- [timeouts](resources--network_policy_view--reference--group-001.md#canonical-2031001202231122-1230130201233212-0211330111232122-2233011300203033-2021113233233022-0313233323303300-1003201003202013-1113233202101200): complete subsection reference.

<a id="canonical-3212223133222010-0113012203010300-3011212132132231-3230331132232113-1011211320130103-1132230201211013-2010030010331322-0021210303201123"></a>

## All schema paths — Property reference / 211300031113 / 11

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--network_policy_view--reference--group-001.md#canonical-0313321231302321-3200023202001230-1132311000313213-3201333212230113-2301300002212202-2120002230200321-1022101203303103-3330103002201230) |
| `description` | [description](resources--network_policy_view--reference--group-001.md#canonical-2330031223021331-3231123023031230-3322212200003312-0133020211201013-3221030103131211-3002010202332011-2011300023311121-3213310023301202) |
| `disable` | [disable](resources--network_policy_view--reference--group-001.md#canonical-3202021133132213-3002011203223012-3100313333300301-3032212031330321-1213111020201320-0013033032030332-0300020331310322-0121031023023012) |
| `egress_rules` | [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-0121010023023121-2223003311233332-0200323131321011-1121230011003320-2233023200331012-3202200101211332-0303012310332122-0322230232111322) |
| `egress_rules.action` | [egress_rules.action](resources--network_policy_view--reference--group-001.md#canonical-3201233030203230-1220003012300002-0110123303230111-0031111333210013-0000022131101313-1033033112231331-0112101332123232-1232202200020333) |
| `egress_rules.adv_action` | [egress_rules.adv_action](resources--network_policy_view--reference--group-001.md#canonical-3123001103322030-1223111033310200-0202003210310120-0103133003333011-0123103102032131-3203232303012311-2033132310222030-1203210222213223) |
| `egress_rules.adv_action.action` | [egress_rules.adv_action.action](resources--network_policy_view--reference--group-001.md#canonical-1303221222003202-1123012001331220-2013123321000320-2031332321211203-1102303210233011-2133000011100022-0003030001112013-3120012220033223) |
| `egress_rules.all_tcp_traffic` | [egress_rules.all_tcp_traffic](resources--network_policy_view--reference--group-001.md#canonical-3031100210213220-2013321323021230-0021302000211230-1311023131100322-3021322121233122-0200311113111100-2211300012013101-3001111131130011) |
| `egress_rules.all_traffic` | [egress_rules.all_traffic](resources--network_policy_view--reference--group-001.md#canonical-3231203022201023-3113103023330133-2300120123121011-1013322020221211-0010123201103320-1301212003030120-2123311231323221-3203032313323013) |
| `egress_rules.all_udp_traffic` | [egress_rules.all_udp_traffic](resources--network_policy_view--reference--group-001.md#canonical-1120310323133333-0113210332200033-3123300300110021-0310201201230111-3313220133230303-0313301230011122-0213111033130223-0223011200300101) |
| `egress_rules.any` | [egress_rules.any](resources--network_policy_view--reference--group-001.md#canonical-2202200132303132-2313310331132101-0210030000201020-2001333102123103-0220033201100332-1331112121312012-2100120312310313-1213313032001010) |
| `egress_rules.applications` | [egress_rules.applications](resources--network_policy_view--reference--group-001.md#canonical-2110323221132233-3303033230302132-1202011312202032-2032113002320313-0022013322020300-0323100221200121-1321023103001311-0200011131331222) |
| `egress_rules.applications.applications` | [egress_rules.applications.applications](resources--network_policy_view--reference--group-001.md#canonical-3200323030211122-3202130200233131-1220221133021300-1100213021332220-0020011031323012-3211023221030322-1123130133303200-0110003200030100) |
| `egress_rules.inside_endpoints` | [egress_rules.inside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-2211210300200003-2231330322232223-0232033133310320-1000113311333210-1322202332213200-3100332002210202-3121201310210313-1233200113333113) |
| `egress_rules.ip_prefix_set` | [egress_rules.ip_prefix_set](resources--network_policy_view--reference--group-001.md#canonical-2322031013313101-3120111320023003-2131110303233110-0120200303310302-2000312033311322-3123022331111000-2220223132203230-0011312311323101) |
| `egress_rules.ip_prefix_set.ref` | [egress_rules.ip_prefix_set.ref](resources--network_policy_view--reference--group-001.md#canonical-1203321132132033-1010213320330023-0310120311320032-1122333300230322-0011233100132331-0000010210202302-0210302202201331-3203231012022113) |
| `egress_rules.ip_prefix_set.ref.kind` | [egress_rules.ip_prefix_set.ref.kind](resources--network_policy_view--reference--group-001.md#canonical-3222333133200233-1023221020023120-0122301231322013-1120232231223132-2103210333000120-1121321011021313-2110011222111102-3303012331321012) |
| `egress_rules.ip_prefix_set.ref.name` | [egress_rules.ip_prefix_set.ref.name](resources--network_policy_view--reference--group-001.md#canonical-0102122322100112-0112310303301030-1113202011323232-1202033222111131-3311000101011231-3323300001222100-2123230333211113-2331023133301200) |
| `egress_rules.ip_prefix_set.ref.namespace` | [egress_rules.ip_prefix_set.ref.namespace](resources--network_policy_view--reference--group-001.md#canonical-2133130322100232-2013123303223331-3223203203322321-3010011130023211-2220300100111331-3213200222220130-0002120210331211-3103100203011310) |
| `egress_rules.ip_prefix_set.ref.tenant` | [egress_rules.ip_prefix_set.ref.tenant](resources--network_policy_view--reference--group-001.md#canonical-3121300221231201-0332201213111312-0300020002121111-1030001120021202-0020223233001310-3330011210211121-0010331232131302-2130012132113302) |
| `egress_rules.ip_prefix_set.ref.uid` | [egress_rules.ip_prefix_set.ref.uid](resources--network_policy_view--reference--group-001.md#canonical-2210212230332300-3221332132110122-1230103121233302-2033333322213333-3203312002200120-1310000012130311-1033110212131222-0303221122132123) |
| `egress_rules.label_matcher` | [egress_rules.label_matcher](resources--network_policy_view--reference--group-001.md#canonical-2031213323131011-1032221212111101-1123311011110130-0111131103223013-2233032023220213-1212311030001221-0210011222001331-2022202201201000) |
| `egress_rules.label_matcher.keys` | [egress_rules.label_matcher.keys](resources--network_policy_view--reference--group-001.md#canonical-1121122213322233-1131223231222033-2001003111221133-0301023131320300-0100330001120213-0112211033310301-2302202023313122-1332333201321231) |
| `egress_rules.label_selector` | [egress_rules.label_selector](resources--network_policy_view--reference--group-001.md#canonical-3313320331122110-3100300102212213-0113211033102223-0032333032003133-3022330110320012-0323122232233132-3221303223303321-3012020000111031) |
| `egress_rules.label_selector.expressions` | [egress_rules.label_selector.expressions](resources--network_policy_view--reference--group-001.md#canonical-3023100220120103-0101031311121111-1021131201130323-3232020030023330-1313203131210033-0202111112322201-1122013132232300-1122112001013110) |
| `egress_rules.metadata` | [egress_rules.metadata](resources--network_policy_view--reference--group-001.md#canonical-1133333210033320-0213313223312022-3100133120200200-1311201103300223-1101100213310103-2122323211130330-3133231133321113-0121021232120311) |
| `egress_rules.metadata.description_spec` | [egress_rules.metadata.description_spec](resources--network_policy_view--reference--group-001.md#canonical-2100023301031322-1333213033331333-0223102233110301-2033133133301022-1323010121111300-3203220032123133-0100200101201012-1110210210010333) |
| `egress_rules.metadata.name` | [egress_rules.metadata.name](resources--network_policy_view--reference--group-001.md#canonical-1002312131132132-1021023131131020-0020111222203230-1300320321012231-1202110032221112-1222311023331002-1231020233311021-0133333200101001) |
| `egress_rules.outside_endpoints` | [egress_rules.outside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-0131203302322221-0111312200120022-2201111101210301-0222013003213331-1312211020310222-0110320233311022-1211200032203000-3111210313000130) |
| `egress_rules.prefix_list` | [egress_rules.prefix_list](resources--network_policy_view--reference--group-001.md#canonical-3112330012132110-0001113010032331-1332010110002330-0000232320130230-3031321111200323-2333121123331210-0122023321322133-0212001212230331) |
| `egress_rules.prefix_list.prefixes` | [egress_rules.prefix_list.prefixes](resources--network_policy_view--reference--group-001.md#canonical-0122301232101003-3012013312302011-1123220133033301-2310003303022222-2102321201103302-3010013221222310-3021331121120020-1113232331330010) |
| `egress_rules.protocol_port_range` | [egress_rules.protocol_port_range](resources--network_policy_view--reference--group-001.md#canonical-2210312321003321-3310132233021212-1200001201130322-3302213011023133-0012222223102321-3321200131331101-2312313323123332-3033210233302312) |
| `egress_rules.protocol_port_range.port_ranges` | [egress_rules.protocol_port_range.port_ranges](resources--network_policy_view--reference--group-001.md#canonical-2220322102131333-1132210022213210-2131321211222133-0213113322313332-2233101200022102-1301321123020010-1101010002120312-3011320322333133) |
| `egress_rules.protocol_port_range.protocol` | [egress_rules.protocol_port_range.protocol](resources--network_policy_view--reference--group-001.md#canonical-1333033030000331-1013212331322222-3230213021030030-2132230012211102-3120222110120111-2000322003131212-3033003322212123-1202311331332212) |
| `endpoint` | [endpoint](resources--network_policy_view--reference--group-001.md#canonical-0001103100311123-1022002020001032-0200112303300210-0101102101220010-3022100133110122-0120021312002311-2213030202333331-3013010130000100) |
| `endpoint.any` | [endpoint.any](resources--network_policy_view--reference--group-001.md#canonical-0011021021131320-2110310312312030-0012022303103212-3010031113021132-0321011203123301-1303001103320331-2322010200200310-2000200032113012) |
| `endpoint.inside_endpoints` | [endpoint.inside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-2320011312210032-2032011121120002-3030310223011231-2021013031101200-3013121011100102-1032022333101203-3332121033223300-0132122002111211) |
| `endpoint.label_selector` | [endpoint.label_selector](resources--network_policy_view--reference--group-001.md#canonical-1331230011301201-1322321021332110-0010312332112002-3132010231010203-3031022030320200-3120211201231321-3321010233132232-2123331210032100) |
| `endpoint.label_selector.expressions` | [endpoint.label_selector.expressions](resources--network_policy_view--reference--group-001.md#canonical-0103121001200002-0103101321120101-1333002210200030-0322011330220113-0130331130211203-0002122030022333-0023023332321013-0100213103130122) |
| `endpoint.outside_endpoints` | [endpoint.outside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-3013333322033313-3223112002202113-3030112120202002-1033130202232203-0312032232301322-1122130031001300-1323013321112220-3322303322330322) |
| `endpoint.prefix_list` | [endpoint.prefix_list](resources--network_policy_view--reference--group-001.md#canonical-2332101330132211-2022223210230102-3201223321021013-0302011313231232-2131302031200023-0320110311003300-2130213332233032-1202330210320130) |
| `endpoint.prefix_list.prefixes` | [endpoint.prefix_list.prefixes](resources--network_policy_view--reference--group-001.md#canonical-3320320313332022-0301131330102203-2111332030110312-3313200012100023-3220222312331013-3300133311203032-1012033120111102-0203001123021211) |
| `id` | [id](resources--network_policy_view--reference--group-001.md#canonical-3031030011210200-3320332030221122-3121323301101120-3203200310102113-0111031230213323-2300003123032233-1032221322020032-2022312323323212) |
| `ingress_rules` | [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-3011103302300222-0333123202233311-3321223300032122-2112332210101103-3310113131131002-2003013323003131-2332200132322031-0313101210313120) |
| `ingress_rules.action` | [ingress_rules.action](resources--network_policy_view--reference--group-001.md#canonical-2300233102010202-3110002031230200-3333103011122311-3103001123313120-1023320131211032-2321201333131023-1012133312201303-2032013022210303) |
| `ingress_rules.adv_action` | [ingress_rules.adv_action](resources--network_policy_view--reference--group-001.md#canonical-2210230222201200-0030311332331131-1110021021200112-2003110210033333-0133220201102020-1200323030113100-1112112333000002-2113220102320300) |
| `ingress_rules.adv_action.action` | [ingress_rules.adv_action.action](resources--network_policy_view--reference--group-001.md#canonical-2102321231013201-0221032213121132-0321011223000021-2110011103100330-3101102313312030-3103323202020001-1112130200311111-2010231003231133) |
| `ingress_rules.all_tcp_traffic` | [ingress_rules.all_tcp_traffic](resources--network_policy_view--reference--group-001.md#canonical-1122011302333321-1103312301323033-0303033110321212-3101133300130321-0320030230032203-3232101231001312-1321221221032210-3221120112320132) |
| `ingress_rules.all_traffic` | [ingress_rules.all_traffic](resources--network_policy_view--reference--group-001.md#canonical-0112301033103003-1203213101212021-2201113210102311-2101331100210013-0020103330203000-2101010020213001-0100132011101132-0133333023023331) |
| `ingress_rules.all_udp_traffic` | [ingress_rules.all_udp_traffic](resources--network_policy_view--reference--group-001.md#canonical-0000302302211202-1130200201012112-2122230201303011-2321300103031002-2322000331220012-1303331211301111-0123302203013103-1322233220220311) |
| `ingress_rules.any` | [ingress_rules.any](resources--network_policy_view--reference--group-001.md#canonical-3212202002113000-2302020200203320-2203010302213003-3203201223333333-0001220201021122-2000210023000012-3002311211123302-1023301123320210) |
| `ingress_rules.applications` | [ingress_rules.applications](resources--network_policy_view--reference--group-001.md#canonical-2011223130023002-0322120103300111-1003233332221030-1103031121032110-3320333110221313-1313301310232010-1011131030232322-0000031212132320) |
| `ingress_rules.applications.applications` | [ingress_rules.applications.applications](resources--network_policy_view--reference--group-001.md#canonical-0122002101023202-3000320300022130-1323000131023322-1222323311232211-3332330221121131-0011202131210010-1333200303110013-1333210301001300) |
| `ingress_rules.inside_endpoints` | [ingress_rules.inside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-0033221223013132-2130312031221222-0312321200201002-2213300212132112-0111000033303120-1021202133311111-1012200112220210-2233012221221203) |
| `ingress_rules.ip_prefix_set` | [ingress_rules.ip_prefix_set](resources--network_policy_view--reference--group-001.md#canonical-0333103123302002-3310201122332301-3212313310011122-3321213031330231-2003033030030102-2013202330331031-3202232012113212-0201303331321332) |
| `ingress_rules.ip_prefix_set.ref` | [ingress_rules.ip_prefix_set.ref](resources--network_policy_view--reference--group-001.md#canonical-3200103102320121-3032230123303220-1102223011212212-1210232200320300-2332131031031122-3300321111210201-2221331120303022-3231111332202302) |
| `ingress_rules.ip_prefix_set.ref.kind` | [ingress_rules.ip_prefix_set.ref.kind](resources--network_policy_view--reference--group-001.md#canonical-1312220231033031-1202333032231310-0200302232223001-0021010111230021-2033331220321032-2332110111202120-2100013123103033-1212330222320332) |
| `ingress_rules.ip_prefix_set.ref.name` | [ingress_rules.ip_prefix_set.ref.name](resources--network_policy_view--reference--group-001.md#canonical-3112333010301100-3200333012012233-0220032033212221-2212330310210233-0203103131021030-0131010200321333-2210211103122223-1011102300123100) |
| `ingress_rules.ip_prefix_set.ref.namespace` | [ingress_rules.ip_prefix_set.ref.namespace](resources--network_policy_view--reference--group-001.md#canonical-1332102000322111-2130010322300301-3122210123332230-2312013130303212-0233210212133300-0111221312311211-2233313333221221-1000033011031311) |
| `ingress_rules.ip_prefix_set.ref.tenant` | [ingress_rules.ip_prefix_set.ref.tenant](resources--network_policy_view--reference--group-001.md#canonical-2311222031120300-0001010213021223-2201223331123111-1111131233220301-3100102103330231-3120231130133321-1313322110233213-3013021111102111) |
| `ingress_rules.ip_prefix_set.ref.uid` | [ingress_rules.ip_prefix_set.ref.uid](resources--network_policy_view--reference--group-001.md#canonical-3020301013222100-3011210030332121-0330330323000133-2113222100131311-3100320332013131-3201323111300321-3200021111202202-1121321233133232) |
| `ingress_rules.label_matcher` | [ingress_rules.label_matcher](resources--network_policy_view--reference--group-001.md#canonical-3100223201201121-0211122332122102-1023221321133302-2102001012332012-0123302221313103-0230331312323111-0331222331301201-2323211023322233) |
| `ingress_rules.label_matcher.keys` | [ingress_rules.label_matcher.keys](resources--network_policy_view--reference--group-001.md#canonical-3221313323201032-1200131111020203-1222110133233331-3103223113232323-3122223020032231-3130110012020110-3231231201101120-0301210211113122) |
| `ingress_rules.label_selector` | [ingress_rules.label_selector](resources--network_policy_view--reference--group-001.md#canonical-3310220230013032-1031212232310313-2032213123112313-0210313010332320-3303022132210121-2103301131301122-1132310221023231-1001112133221231) |
| `ingress_rules.label_selector.expressions` | [ingress_rules.label_selector.expressions](resources--network_policy_view--reference--group-001.md#canonical-1331000121133311-2330112300223111-1213230231003321-1212233311030310-0131102331331233-0111002000102322-3213130023122102-3033300223110031) |
| `ingress_rules.metadata` | [ingress_rules.metadata](resources--network_policy_view--reference--group-001.md#canonical-3132201022011010-2122212031333301-0022130123031221-2332202020220310-0022300311313312-2111311212100332-1220202031302201-0303033023302023) |
| `ingress_rules.metadata.description_spec` | [ingress_rules.metadata.description_spec](resources--network_policy_view--reference--group-001.md#canonical-1030100303231003-1102031230101102-3202103220133011-3001100213232302-1333131221313032-2110022331311021-0010230133212231-1011111120323031) |
| `ingress_rules.metadata.name` | [ingress_rules.metadata.name](resources--network_policy_view--reference--group-001.md#canonical-0112230121221113-1113303201103130-1002312212333222-0322110100031312-1211013210120000-3321101223000000-1003301222323220-0310203320211130) |
| `ingress_rules.outside_endpoints` | [ingress_rules.outside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-1202222230310011-2010012031011101-2232320232332031-0113213233023313-2310200232303121-2032310211030321-0102112203212321-3211120232221101) |
| `ingress_rules.prefix_list` | [ingress_rules.prefix_list](resources--network_policy_view--reference--group-001.md#canonical-3023321000233203-3021102310232203-3200000103121110-1000023310120111-0010100023122123-3201210003103102-1002322112232211-0232212100112000) |
| `ingress_rules.prefix_list.prefixes` | [ingress_rules.prefix_list.prefixes](resources--network_policy_view--reference--group-001.md#canonical-2111323233101123-2003222212331022-3233132322113112-1001202010201213-1133320232021100-0301231211203020-2211311112133031-0133202103331030) |
| `ingress_rules.protocol_port_range` | [ingress_rules.protocol_port_range](resources--network_policy_view--reference--group-001.md#canonical-0021001012103200-3002223020013310-3313022221231321-3301312123110300-3011332031231120-3201100100120121-0001300133323202-2201332021301013) |
| `ingress_rules.protocol_port_range.port_ranges` | [ingress_rules.protocol_port_range.port_ranges](resources--network_policy_view--reference--group-001.md#canonical-1222210113200213-2313203221320222-3122301202220021-2021010322021332-1032322103322320-3202202203232020-3203320332321303-1330020313133311) |
| `ingress_rules.protocol_port_range.protocol` | [ingress_rules.protocol_port_range.protocol](resources--network_policy_view--reference--group-001.md#canonical-3233230201333110-2021012133100121-0222003130133223-1232113022022120-3202021131013201-1132013032321223-2101003120321010-2103031112300022) |
| `labels` | [labels](resources--network_policy_view--reference--group-001.md#canonical-0103002231123333-2123121123233122-2011132213311221-3301223222020313-3333032233032133-3211023033313302-0133020003213322-2201000000103202) |
| `name` | [name](resources--network_policy_view--reference--group-001.md#canonical-0003110202032023-3210031220230011-3203301222113212-3123322003132000-3132130000111322-0032303113010020-1113212021123011-0011130211311002) |
| `namespace` | [namespace](resources--network_policy_view--reference--group-001.md#canonical-2103030022131000-1203230322101231-1332312330002121-2301231121000022-1320111323312323-0331332302231322-3300112322021031-3330221330033023) |
| `timeouts` | [timeouts](resources--network_policy_view--reference--group-001.md#canonical-0122023323013000-2111320203103200-0203101203131222-3232020320230320-1021213223031021-0221210300032321-1002323231303101-3200330211202300) |
| `timeouts.create` | [timeouts.create](resources--network_policy_view--reference--group-001.md#canonical-1220133213331221-1002311301311112-3321221202320001-3211122213010333-1021233013102202-0022301012023302-0233213323001301-1310222123202111) |
| `timeouts.delete` | [timeouts.delete](resources--network_policy_view--reference--group-001.md#canonical-0030200120130312-1111000011331020-3331113201101300-1213233300211232-3323203002111231-0201333210310113-0213323132031013-3001012021231023) |
| `timeouts.read` | [timeouts.read](resources--network_policy_view--reference--group-001.md#canonical-1331021301300312-1303030110322213-3030022322321001-3300111100033112-3333332310001212-3123033220000033-1120233231330031-2011300123121311) |
| `timeouts.update` | [timeouts.update](resources--network_policy_view--reference--group-001.md#canonical-1002210332013320-2310212323001213-0300100003132100-2013211231220210-2031233300211232-0320020012003130-0032033223220013-3330303012300011) |

<a id="canonical-0320202113010232-3300003121222013-1101102011322003-0022221200030122-1020010013011030-2232122200112322-2113210212313221-1021121120330201"></a>

## Next pages — Property reference / 211300031113 / 12

- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- [endpoint](resources--network_policy_view--reference--group-001.md#canonical-3300312010233011-1332123212033122-3330303200202222-0112211301022200-0121210323131100-1100322331002021-3232201003132022-3201310313012123)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- [timeouts](resources--network_policy_view--reference--group-001.md#canonical-2031001202231122-1230130201233212-0211330111232122-2233011300203033-2021113233233022-0313233323303300-1003201003202013-1113233202101200)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301321213200010-1012213200303230-0333223131122023-1301321002133032-2121111113103320-1311322332221202-0100132331023300-0122020203230020"></a>

## egress_rules — egress_rules / 102303203030 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- egress_rules

<a id="canonical-0121010023023121-2223003311233332-0200323131321011-1121230011003320-2233023200331012-3202200101211332-0303012310332122-0322230232111322"></a>

Type: `"object"`. list nested block, Optional.

Ordered list of rules applied to connections from policy endpoints.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "all_traffic"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "all_udp_traffic"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "applications"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("all_traffic",
    "all_udp_traffic"),
  validators.ConflictingListObjectAttributes("all_traffic",
    "applications"),
  validators.ConflictingListObjectAttributes("all_traffic",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("all_udp_traffic",
    "applications"),
  validators.ConflictingListObjectAttributes("all_udp_traffic",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("any",
    "inside_endpoints"),
  validators.ConflictingListObjectAttributes("any",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("any",
    "label_selector"),
  validators.ConflictingListObjectAttributes("any",
    "outside_endpoints"),
  validators.ConflictingListObjectAttributes("any",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("applications",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("inside_endpoints",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("inside_endpoints",
    "label_selector"),
  validators.ConflictingListObjectAttributes("inside_endpoints",
    "outside_endpoints"),
  validators.ConflictingListObjectAttributes("inside_endpoints",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "label_selector"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "outside_endpoints"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("label_selector",
    "outside_endpoints"),
  validators.ConflictingListObjectAttributes("label_selector",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("outside_endpoints",
    "prefix_list")}
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
egress_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-0322212032312023-3332303313322210-0032223202131223-2121213000001133-2013210120320020-2122320322110231-2300321323010312-3013012122030022"></a>

## Direct properties — egress_rules / 102303203030 / 3

<a id="canonical-3201233030203230-1220003012300002-0110123303230111-0031111333210013-0000022131101313-1033033112231331-0112101332123232-1232202200020333"></a>

<a id="canonical-2001221011020313-2121021303331003-2112232123321322-1330002101113012-1133333320210323-0333323100121003-3033330130132123-0020032330231131"></a>

## action property — egress_rules / 102303203030 / 4

Type: `"string"`. Optional.

\[Enum: DENY|ALLOW\] Network policy rule action configures the action to be taken on rule match
Apply deny action on rule match Apply allow action on rule match. Possible values are \`DENY\`,
\`ALLOW\`. Defaults to \`DENY\`.

Upstream description:

Network policy rule action configures the action to be taken on rule match

Apply deny action on rule match Apply allow action on rule match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DENY",
    "ALLOW"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "DENY",
  "enum": [
    "DENY",
    "ALLOW"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [adv_action](resources--network_policy_view--reference--group-001.md#canonical-2232011000132112-1002200220113102-1200001112021013-3033203312332210-0312310211112230-0102311133210131-0133220111313023-0302121232223013): complete subsection reference.

- [all_tcp_traffic](resources--network_policy_view--reference--group-001.md#canonical-2211323333213302-0122330231321012-2003012331133031-3032330112331303-2112033000331313-1102033003111333-2001201112112000-3123211233120312): complete subsection reference.

- [all_traffic](resources--network_policy_view--reference--group-001.md#canonical-3110333102010022-1122120020201210-2320131323333303-3301002321111131-1032300213212010-1331110221312130-3002021012311323-3333011011001303): complete subsection reference.

- [all_udp_traffic](resources--network_policy_view--reference--group-001.md#canonical-2223203031131111-3221222202113213-3311003231011301-1102302020200030-0130331011000232-3321132103233013-3220111133032330-1020012233232333): complete subsection reference.

- [any](resources--network_policy_view--reference--group-001.md#canonical-0320102012023323-2123323033100130-0203201101010210-1123301121232300-1100101122003210-0333330331113220-1310021000112312-0330131001200022): complete subsection reference.

- [applications](resources--network_policy_view--reference--group-001.md#canonical-2231120103111332-1031333103211212-2112303213120211-3301032110010030-3003202221103300-1231111330101233-3101302210133303-3313002010201213): complete subsection reference.

- [inside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-2210133131300203-1311111030103333-2101221313020211-3202313313223233-1033033110213023-3203321220022331-3010313101300122-2110003202320001): complete subsection reference.

- [ip_prefix_set](resources--network_policy_view--reference--group-001.md#canonical-2133121303032023-1302021303121323-1032331322022130-1010311332130110-1330131300233300-3113210122132120-3030030223300200-3102232030321122): complete subsection reference.

- [label_matcher](resources--network_policy_view--reference--group-001.md#canonical-2220100223102200-3232110033110311-1312000013102023-2221303023102312-3133200332001321-1003130330230321-3302113101122102-0110322211001221): complete subsection reference.

- [label_selector](resources--network_policy_view--reference--group-001.md#canonical-0321003102212200-3312220331213011-3333131120110220-3311000313001112-1011212112233002-3023331111310321-2223330303322130-1210002202203001): complete subsection reference.

- [metadata](resources--network_policy_view--reference--group-001.md#canonical-1222103211120113-3103300030011023-0233333130213311-0333200023321321-1020133113303311-0111120132230302-3021320210022032-1133002121212122): complete subsection reference.

- [outside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-0133031212300131-1013020101300120-2113132020123301-3030301301013330-2133333333113301-1022310211302320-0233030013332220-0300132103230210): complete subsection reference.

- [prefix_list](resources--network_policy_view--reference--group-001.md#canonical-2300333020003231-3212220233212321-2210311311032002-3203301103132011-3311321223320233-2131333121121011-1110223131100003-1200323333111232): complete subsection reference.

- [protocol_port_range](resources--network_policy_view--reference--group-001.md#canonical-1333012223013200-2012222233331323-2101313222223312-1023133002032311-2230321032202131-0013122123103331-1331211303023100-0110312013100232): complete subsection reference.

<a id="canonical-2323110021212230-1233231321020302-2013333322200313-3200101003222331-2000111301210213-3322121332122330-0010000112302012-0201220010323223"></a>

## Next pages — egress_rules / 102303203030 / 5

- [egress_rules.adv_action](resources--network_policy_view--reference--group-001.md#canonical-2232011000132112-1002200220113102-1200001112021013-3033203312332210-0312310211112230-0102311133210131-0133220111313023-0302121232223013)
- [egress_rules.all_tcp_traffic](resources--network_policy_view--reference--group-001.md#canonical-2211323333213302-0122330231321012-2003012331133031-3032330112331303-2112033000331313-1102033003111333-2001201112112000-3123211233120312)
- [egress_rules.all_traffic](resources--network_policy_view--reference--group-001.md#canonical-3110333102010022-1122120020201210-2320131323333303-3301002321111131-1032300213212010-1331110221312130-3002021012311323-3333011011001303)
- [egress_rules.all_udp_traffic](resources--network_policy_view--reference--group-001.md#canonical-2223203031131111-3221222202113213-3311003231011301-1102302020200030-0130331011000232-3321132103233013-3220111133032330-1020012233232333)
- [egress_rules.any](resources--network_policy_view--reference--group-001.md#canonical-0320102012023323-2123323033100130-0203201101010210-1123301121232300-1100101122003210-0333330331113220-1310021000112312-0330131001200022)
- [egress_rules.applications](resources--network_policy_view--reference--group-001.md#canonical-2231120103111332-1031333103211212-2112303213120211-3301032110010030-3003202221103300-1231111330101233-3101302210133303-3313002010201213)
- [egress_rules.inside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-2210133131300203-1311111030103333-2101221313020211-3202313313223233-1033033110213023-3203321220022331-3010313101300122-2110003202320001)
- [egress_rules.ip_prefix_set](resources--network_policy_view--reference--group-001.md#canonical-2133121303032023-1302021303121323-1032331322022130-1010311332130110-1330131300233300-3113210122132120-3030030223300200-3102232030321122)
- [egress_rules.label_matcher](resources--network_policy_view--reference--group-001.md#canonical-2220100223102200-3232110033110311-1312000013102023-2221303023102312-3133200332001321-1003130330230321-3302113101122102-0110322211001221)
- [egress_rules.label_selector](resources--network_policy_view--reference--group-001.md#canonical-0321003102212200-3312220331213011-3333131120110220-3311000313001112-1011212112233002-3023331111310321-2223330303322130-1210002202203001)
- [egress_rules.metadata](resources--network_policy_view--reference--group-001.md#canonical-1222103211120113-3103300030011023-0233333130213311-0333200023321321-1020133113303311-0111120132230302-3021320210022032-1133002121212122)
- [egress_rules.outside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-0133031212300131-1013020101300120-2113132020123301-3030301301013330-2133333333113301-1022310211302320-0233030013332220-0300132103230210)
- [egress_rules.prefix_list](resources--network_policy_view--reference--group-001.md#canonical-2300333020003231-3212220233212321-2210311311032002-3203301103132011-3311321223320233-2131333121121011-1110223131100003-1200323333111232)
- [egress_rules.protocol_port_range](resources--network_policy_view--reference--group-001.md#canonical-1333012223013200-2012222233331323-2101313222223312-1023133002032311-2230321032202131-0013122123103331-1331211303023100-0110312013100232)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-2232011000132112-1002200220113102-1200001112021013-3033203312332210-0312310211112230-0102311133210131-0133220111313023-0302121232223013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111330002130201-0333132322113211-1110210022120311-0313221233103300-1221133200310203-1003221222000030-0323123121023121-0221220301311121"></a>

## egress_rules.adv_action — adv_action / 323120311102 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- egress_rules.adv_action

<a id="canonical-3123001103322030-1223111033310200-0202003210310120-0103133003333011-0123103102032131-3203232303012311-2033132310222030-1203210222213223"></a>

Type: `"object"`. single nested block, Optional.

Network Policy Rule Advanced Action provides additional OPTIONS along with RuleAction and
PBRRuleAction.

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
adv_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-3212331031302132-2232230323301222-1332301130022233-1202332200001123-3121120313202120-3001112000120210-0222312020031333-0111211032231001"></a>

## Direct properties — adv_action / 323120311102 / 3

<a id="canonical-1303221222003202-1123012001331220-2013123321000320-2031332321211203-1102303210233011-2133000011100022-0003030001112013-3120012220033223"></a>

<a id="canonical-3232111112132021-1020310300023230-0331030232302032-3303330133131321-0030312321232232-3202121310300111-1232313103312210-2311010323132102"></a>

## action property — adv_action / 323120311102 / 4

Type: `"string"`. Optional.

\[Enum: NOLOG|LOG\] Choice to choose logging or no logging This works together with option selected
via NetworkPolicyRuleAction or any other action specified x-. Possible values are \`NOLOG\`,
\`LOG\`. Defaults to \`NOLOG\`.

Upstream description:

Choice to choose logging or no logging This works together with option selected via
NetworkPolicyRuleAction or any other action specified x-example: (No Selection in
NetworkPolicyRuleAction + AdvancedAction as LOG) = LOG Only, (ALLOW/DENY in NetworkPolicyRuleAction
&#8203;+ AdvancedAction as LOG) = Log and Allow/Deny, (ALLOW/DENY in NetworkPolicyRuleAction + NOLOG in
AdvancedAction) = Allow/Deny with no log

Don't sample the traffic hitting the rule Sample the traffic hitting the rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NOLOG",
    "LOG"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NOLOG",
  "enum": [
    "NOLOG",
    "LOG"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0112011011331010-2231210123330133-3202012123232210-0002201120232201-3203030231211131-0111001100233111-2220200302003101-2323133021302112"></a>

## Next pages — adv_action / 323120311102 / 5

- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-2211323333213302-0122330231321012-2003012331133031-3032330112331303-2112033000331313-1102033003111333-2001201112112000-3123211233120312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2021100132101221-1233000312000213-2233300301232302-2220130313123101-1230333222030101-3220003023331320-1202030202203120-3200301320131113"></a>

## egress_rules.all_tcp_traffic — all_tcp_traffic / 021233322301 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- egress_rules.all_tcp_traffic

<a id="canonical-3031100210213220-2013321323021230-0021302000211230-1311023131100322-3021322121233122-0200311113111100-2211300012013101-3001111131130011"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all tcp traffic.

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
all_tcp_traffic = {}
```

<a id="canonical-2001002022102333-2120322323202110-3102312031011032-0300322210112022-0231203330301101-1002112103210131-1101002132200222-2021102311112032"></a>

## Direct properties — all_tcp_traffic / 021233322301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3033002300330211-1333230221020022-2101012003013320-3031132321021131-1103330222223021-2210232203303130-2110101033131232-3202011302121302"></a>

## Next pages — all_tcp_traffic / 021233322301 / 4

- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-3110333102010022-1122120020201210-2320131323333303-3301002321111131-1032300213212010-1331110221312130-3002021012311323-3333011011001303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102111030133231-1011311021102110-2221113321023100-3030003022311122-0233232220232012-1100203003131221-1101312011131232-1110210013220202"></a>

## egress_rules.all_traffic — all_traffic / 211123320302 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- egress_rules.all_traffic

<a id="canonical-3231203022201023-3113103023330133-2300120123121011-1013322020221211-0010123201103320-1301212003030120-2123311231323221-3203032313323013"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all traffic.

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
all_traffic = {}
```

<a id="canonical-3011213330213020-2122133200312103-1200101210021112-1022332001323132-1202202203012301-0110322332332010-1121232203221220-1220022103202011"></a>

## Direct properties — all_traffic / 211123320302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3310013132232331-3103132023020331-1323310013000012-2102121231103312-1310101003020322-1102010003313112-2131120300202003-3333002210213023"></a>

## Next pages — all_traffic / 211123320302 / 4

- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-2223203031131111-3221222202113213-3311003231011301-1102302020200030-0130331011000232-3321132103233013-3220111133032330-1020012233232333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0022320210011100-3101032013312102-0313202231223022-3032323023013221-0332103230333100-2301220103000322-3010121332210103-3223202323111212"></a>

## egress_rules.all_udp_traffic — all_udp_traffic / 011111201310 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- egress_rules.all_udp_traffic

<a id="canonical-1120310323133333-0113210332200033-3123300300110021-0310201201230111-3313220133230303-0313301230011122-0213111033130223-0223011200300101"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all udp traffic.

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
all_udp_traffic = {}
```

<a id="canonical-2200310332121032-0313212230212110-3323233112033213-3233220223131202-2332113031211202-3002132100100001-1101123231012201-1223211011233201"></a>

## Direct properties — all_udp_traffic / 011111201310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1120121303100311-3331013101202222-1011033202030010-0100131211223203-3021123312202300-2302122121202222-3102232310022002-1102112232201033"></a>

## Next pages — all_udp_traffic / 011111201310 / 4

- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-0320102012023323-2123323033100130-0203201101010210-1123301121232300-1100101122003210-0333330331113220-1310021000112312-0330131001200022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020222123211021-0031003330332332-3222303221333130-1312113211021201-0101010211300131-2332223003031032-2010122103000221-2333223100111003"></a>

## egress_rules.any — any / 202123223122 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- egress_rules.any

<a id="canonical-2202200132303132-2313310331132101-0210030000201020-2001333102123103-0220033201100332-1331112121312012-2100120312310313-1213313032001010"></a>

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

<a id="canonical-0130320101132021-0230322302311123-3101321120002322-3333203123222000-3122233313022130-0333221112112223-0113021010110022-0011000221130032"></a>

## Direct properties — any / 202123223122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2301001022320113-2103313311110222-1323213133222133-3103020132010221-0311130300330223-0221000122112031-3222312302103313-0002223010222010"></a>

## Next pages — any / 202123223122 / 4

- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-2231120103111332-1031333103211212-2112303213120211-3301032110010030-3003202221103300-1231111330101233-3101302210133303-3313002010201213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320020023320302-3332333230132103-2330220003101111-2120011220302100-2102003002100230-3233232022222312-1230320322020331-1200213032100010"></a>

## egress_rules.applications — applications / 101112010112 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- egress_rules.applications

<a id="canonical-2110323221132233-3303033230302132-1202011312202032-2032113002320313-0022013322020300-0323100221200121-1321023103001311-0200011131331222"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for applications.

Upstream description:

Application protocols like HTTP, SNMP.

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
applications {
  # Configure direct properties listed below.
}
```

<a id="canonical-3031332230112303-3013100012130203-3310101002311023-3233130311220301-0303231132331012-1013020321230323-2330031023331313-2222200032323020"></a>

## Direct properties — applications / 101112010112 / 3

<a id="canonical-3200323030211122-3202130200233131-1220221133021300-1100213021332220-0020011031323012-3211023221030322-1123130133303200-0110003200030100"></a>

<a id="canonical-3033123222001130-1022023221111021-1332120012332011-3113100203311102-0013133002310013-2123101321210100-2311303030232112-3111100232021212"></a>

## applications property — applications / 101112010112 / 4

Type: `["list", "string"]`. Optional.

\[Enum: APPLICATION\_HTTP|APPLICATION\_HTTPS|APPLICATION\_SNMP|APPLICATION\_DNS\] Application
Protocols. Application protocols like HTTP, SNMP. Possible values are \`APPLICATION\_HTTP\`,
\`APPLICATION\_HTTPS\`, \`APPLICATION\_SNMP\`, \`APPLICATION\_DNS\`. Defaults to
\`APPLICATION\_HTTP\`.

Upstream description:

Application protocols like HTTP, SNMP.

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

<a id="canonical-2132310100030113-1031210333030312-3331110330201020-1230332232021010-3112210032112310-2010030130333321-3103210300202013-1101323221031210"></a>

## Next pages — applications / 101112010112 / 5

- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-2210133131300203-1311111030103333-2101221313020211-3202313313223233-1033033110213023-3203321220022331-3010313101300122-2110003202320001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322112130030000-3112323122213011-0132021332332230-0311310000031312-3213030222003330-0202212201112303-1133200022211030-3010111110321132"></a>

## egress_rules.inside_endpoints — inside_endpoints / 222321320113 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- egress_rules.inside_endpoints

<a id="canonical-2211210300200003-2231330322232223-0232033133310320-1000113311333210-1322202332213200-3100332002210202-3121201310210313-1233200113333113"></a>

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
inside_endpoints = {}
```

<a id="canonical-2323123221331121-3013322110203220-3002310131220010-0231300100323001-0322102223011303-1102101302233230-1132110001101010-1123030002232232"></a>

## Direct properties — inside_endpoints / 222321320113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2130100021333002-0022020203201321-1202110210131300-2112101103020000-2211200320223310-1013231222310031-0200133030213200-1231302330330211"></a>

## Next pages — inside_endpoints / 222321320113 / 4

- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-2133121303032023-1302021303121323-1032331322022130-1010311332130110-1330131300233300-3113210122132120-3030030223300200-3102232030321122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112212101332220-0130022110032012-2101232220223201-0223231300201003-3100110101302213-2200202300323012-0120231132200103-2111222022013121"></a>

## egress_rules.ip_prefix_set — ip_prefix_set / 111203021031 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- egress_rules.ip_prefix_set

<a id="canonical-2322031013313101-3120111320023003-2131110303233110-0120200303310302-2000312033311322-3123022331111000-2220223132203230-0011312311323101"></a>

Type: `"object"`. single nested block, Optional.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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
ip_prefix_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-1100310122301030-3001000233213310-1230023032020312-2300021323223111-0122020130031132-3331310213320113-1012330002202231-2223331322332333"></a>

## Direct properties — ip_prefix_set / 111203021031 / 3

- [ref](resources--network_policy_view--reference--group-001.md#canonical-0103211003302222-0210110311002123-0203233232131011-0021222130303020-1102203100312102-0223121212102230-0101231113332022-0212203310213032): complete subsection reference.

<a id="canonical-3320333210202102-1331221230312303-1310130113130102-1333130223231101-3113012310212030-2313231021131033-2133200302200333-2230303130123301"></a>

## Next pages — ip_prefix_set / 111203021031 / 4

- [egress_rules.ip_prefix_set.ref](resources--network_policy_view--reference--group-001.md#canonical-0103211003302222-0210110311002123-0203233232131011-0021222130303020-1102203100312102-0223121212102230-0101231113332022-0212203310213032)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-0103211003302222-0210110311002123-0203233232131011-0021222130303020-1102203100312102-0223121212102230-0101231113332022-0212203310213032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311101301233212-2003201330023132-3201121223233201-3013101212133130-0230210303332201-1313201031220001-1210330303021321-0130013201322310"></a>

## egress_rules.ip_prefix_set.ref — ref / 113021322010 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- [egress_rules.ip_prefix_set](resources--network_policy_view--reference--group-001.md#canonical-2133121303032023-1302021303121323-1032331322022130-1010311332130110-1330131300233300-3113210122132120-3030030223300200-3102232030321122)
- egress_rules.ip_prefix_set.ref

<a id="canonical-1203321132132033-1010213320330023-0310120311320032-1122333300230322-0011233100132331-0000010210202302-0210302202201331-3203231012022113"></a>

Type: `"object"`. list nested block, Optional.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-3132313110300122-1021102212333130-1232233123033132-0312102020031103-0011011100322102-1101333321112130-2121112010313012-1311300011100333"></a>

## Direct properties — ref / 113021322010 / 3

<a id="canonical-3222333133200233-1023221020023120-0122301231322013-1120232231223132-2103210333000120-1121321011021313-2110011222111102-3303012331321012"></a>

<a id="canonical-3230312201230000-2330120021210232-0010312103033130-1023103222313230-3212123330230131-3022332023001313-2032300023213322-0003201201113320"></a>

## kind property — ref / 113021322010 / 4

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

<a id="canonical-0102122322100112-0112310303301030-1113202011323232-1202033222111131-3311000101011231-3323300001222100-2123230333211113-2331023133301200"></a>

<a id="canonical-2013001112121210-0233133131112211-2023030203302000-1322112101030133-0331003313303000-1011312031203220-2033220300312020-0323020210100001"></a>

## name property — ref / 113021322010 / 5

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

<a id="canonical-2133130322100232-2013123303223331-3223203203322321-3010011130023211-2220300100111331-3213200222220130-0002120210331211-3103100203011310"></a>

<a id="canonical-3201111230323220-1203100311311022-3032301012132233-0001002220330013-1222223103213232-3201311321100003-2101230210130332-3010212102203312"></a>

## namespace property — ref / 113021322010 / 6

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

<a id="canonical-3121300221231201-0332201213111312-0300020002121111-1030001120021202-0020223233001310-3330011210211121-0010331232131302-2130012132113302"></a>

<a id="canonical-1320013320021122-3122131313230230-1120101101300020-3032222303103032-3122322032300301-1022203101230320-0131322002031032-0113013003113112"></a>

## tenant property — ref / 113021322010 / 7

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

<a id="canonical-2210212230332300-3221332132110122-1230103121233302-2033333322213333-3203312002200120-1310000012130311-1033110212131222-0303221122132123"></a>

<a id="canonical-2301323100113211-3203212200132232-2202303133232012-1031203303213102-2232302121300231-2321323102032230-3111022333122012-1113111230212222"></a>

## uid property — ref / 113021322010 / 8

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

<a id="canonical-0313332330302011-0213003011021013-1111301001100200-2013121001211220-2102213033230233-1213032321032321-3321311102010013-2213003103031011"></a>

## Next pages — ref / 113021322010 / 9

- [egress_rules.ip_prefix_set](resources--network_policy_view--reference--group-001.md#canonical-2133121303032023-1302021303121323-1032331322022130-1010311332130110-1330131300233300-3113210122132120-3030030223300200-3102232030321122)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-2220100223102200-3232110033110311-1312000013102023-2221303023102312-3133200332001321-1003130330230321-3302113101122102-0110322211001221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131013021322323-2010101033111013-3003210101313231-3012003010000021-1100231013001122-1220130013111201-3123022320312103-3222130333201132"></a>

## egress_rules.label_matcher — label_matcher / 132211302010 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- egress_rules.label_matcher

<a id="canonical-2031213323131011-1032221212111101-1123311011110130-0111131103223013-2233032023220213-1212311030001221-0210011222001331-2022202201201000"></a>

Type: `"object"`. single nested block, Optional.

Label matcher specifies a list of label keys whose values need to match for source/client and
destination/server. Note that the actual label values are not specified and do not matter. This
allows an ability to scope grouping by the label key name.

Upstream description:

A label matcher specifies a list of label keys whose values need to match for source/client and
destination/server. Note that the actual label values are not specified and do not matter. This
allows an ability to scope grouping by the label key name.

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
label_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-3002300030332011-1101000330333212-2130333013130113-0202313110021001-1301133013331023-2103231003130102-0121131313221232-1000302021032201"></a>

## Direct properties — label_matcher / 132211302010 / 3

<a id="canonical-1121122213322233-1131223231222033-2001003111221133-0301023131320300-0100330001120213-0112211033310301-2302202023313122-1332333201321231"></a>

<a id="canonical-1031002113210113-0033320221310012-2112030030213231-2113220020130010-3301033111201201-0023012321320220-1110213130302120-2221303311111223"></a>

## keys property — label_matcher / 132211302010 / 4

Type: `["list", "string"]`. Optional.

The list of label key names that have to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2300030013220102-3111013033200102-0101232113010103-3311130211100110-3123221312220123-1212001201012322-1221320103301320-2030231322200120"></a>

## Next pages — label_matcher / 132211302010 / 5

- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-0321003102212200-3312220331213011-3333131120110220-3311000313001112-1011212112233002-3023331111310321-2223330303322130-1210002202203001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102112022111330-3122233213012203-2333123222110023-0111333023101223-2211201230132203-2121123132333021-1110023231310303-0223313212011202"></a>

## egress_rules.label_selector — label_selector / 331320101313 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- egress_rules.label_selector

<a id="canonical-3313320331122110-3100300102212213-0113211033102223-0032333032003133-3022330110320012-0323122232233132-3221303223303321-3012020000111031"></a>

Type: `"object"`. single nested block, Optional.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
label_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-0100022021202231-1221231301031323-1111211320120112-0202000323201201-0221132022120110-3033011131320110-1011220013001231-2100033302123221"></a>

## Direct properties — label_selector / 331320101313 / 3

<a id="canonical-3023100220120103-0101031311121111-1021131201130323-3232020030023330-1313203131210033-0202111112322201-1122013132232300-1122112001013110"></a>

<a id="canonical-3021032022122203-2210103122030300-0330023101313230-3200203110333012-0330201022330101-3201100333202022-3210130233000220-3111323233022122"></a>

## expressions property — label_selector / 331320101313 / 4

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-3330023122113312-1013021023131120-2211010221031101-0013212011130300-2300103013201001-3023020003300333-1111011232212213-1220322023232301"></a>

## Next pages — label_selector / 331320101313 / 5

- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-1222103211120113-3103300030011023-0233333130213311-0333200023321321-1020133113303311-0111120132230302-3021320210022032-1133002121212122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310231001323312-2321023020202322-2323120030003033-1132033332102330-2031211332101313-0231213320001330-2321231121211333-0132213233201032"></a>

## egress_rules.metadata — metadata / 102303232222 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- egress_rules.metadata

<a id="canonical-1133333210033320-0213313223312022-3100133120200200-1311201103300223-1101100213310103-2122323211130330-3133231133321113-0121021232120311"></a>

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

<a id="canonical-3111111230023002-1333131333330213-3210223121223123-1310123231010133-3202333330221322-3200213133201311-2011221022332211-0322221202233313"></a>

## Direct properties — metadata / 102303232222 / 3

<a id="canonical-2100023301031322-1333213033331333-0223102233110301-2033133133301022-1323010121111300-3203220032123133-0100200101201012-1110210210010333"></a>

<a id="canonical-3233221221020011-0301300031102111-1202001211000333-2323333223101111-1232230202012313-2030322100212011-1101121201121112-1102113030301321"></a>

## description_spec property — metadata / 102303232222 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-1002312131132132-1021023131131020-0020111222203230-1300320321012231-1202110032221112-1222311023331002-1231020233311021-0133333200101001"></a>

<a id="canonical-0200333013310010-0131112320331213-1332032002301001-2133331300003310-2320222300301102-0012230321310021-2000301003100321-2030003233022232"></a>

## name property — metadata / 102303232222 / 5

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

<a id="canonical-2202000230113132-2222102012020202-3100232232031110-2130202102130322-3223002111301220-0310020002210033-1323202122220233-0210011212232121"></a>

## Next pages — metadata / 102303232222 / 6

- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-0133031212300131-1013020101300120-2113132020123301-3030301301013330-2133333333113301-1022310211302320-0233030013332220-0300132103230210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3211103030022302-3100122002010002-2222200330203322-1113030231301232-3200220003323203-2032221021303230-1102000113101221-0311000003330003"></a>

## egress_rules.outside_endpoints — outside_endpoints / 000310332103 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- egress_rules.outside_endpoints

<a id="canonical-0131203302322221-0111312200120022-2201111101210301-0222013003213331-1312211020310222-0110320233311022-1211200032203000-3111210313000130"></a>

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
outside_endpoints = {}
```

<a id="canonical-2100023301012012-1200200231112310-1312011023101201-0312011330022030-1013112120101002-0203100112321130-1233020322032123-3230213133221032"></a>

## Direct properties — outside_endpoints / 000310332103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2210311021123201-3131113130123103-2120020202011130-1331232112102020-3033310112132131-0230231132330012-2213123021110313-1111201213303313"></a>

## Next pages — outside_endpoints / 000310332103 / 4

- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-2300333020003231-3212220233212321-2210311311032002-3203301103132011-3311321223320233-2131333121121011-1110223131100003-1200323333111232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2330020303030131-2300322111022320-1201031000311112-3131103031210300-3103031001011131-1101212020332123-0230322111210320-3112031330313331"></a>

## egress_rules.prefix_list — prefix_list / 002303012013 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- egress_rules.prefix_list

<a id="canonical-3112330012132110-0001113010032331-1332010110002330-0000232320130230-3031321111200323-2333121123331210-0122023321322133-0212001212230331"></a>

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
prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0210101210200022-0311102131302011-3110330111303221-0121001211211203-1323001132000232-1020311013331020-1212120102102022-2223000332113033"></a>

## Direct properties — prefix_list / 002303012013 / 3

<a id="canonical-0122301232101003-3012013312302011-1123220133033301-2310003303022222-2102321201103302-3010013221222310-3021331121120020-1113232331330010"></a>

<a id="canonical-3031320102302132-3103121230101020-3330110223102233-2020310231312033-3202031123210300-0311112002000302-0233131203321232-0113012131223302"></a>

## prefixes property — prefix_list / 002303012013 / 4

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

<a id="canonical-2130213002032201-1113300122221110-2102332322230010-2321101200131032-0021320130211022-3220310100313100-1023311233320230-1220101110221111"></a>

## Next pages — prefix_list / 002303012013 / 5

- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-1333012223013200-2012222233331323-2101313222223312-1023133002032311-2230321032202131-0013122123103331-1331211303023100-0110312013100232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3020312201031321-3302302132211133-3302113133111121-1122032203021202-2032120012130110-3103213201110211-3023231321022223-1113202203202033"></a>

## egress_rules.protocol_port_range — protocol_port_range / 222202122123 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- egress_rules.protocol_port_range

<a id="canonical-2210312321003321-3310132233021212-1200001201130322-3302213011023133-0012222223102321-3321200131331101-2312313323123332-3033210233302312"></a>

Type: `"object"`. single nested block, Optional.

Protocol and Port. Protocol and Port ranges.

Upstream description:

Protocol and Port ranges.

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
protocol_port_range {
  # Configure direct properties listed below.
}
```

<a id="canonical-3120201020030331-0233122300000311-1021123030302001-0001232332231222-3221121221123023-1312230211030201-0212012123002022-1023021302033213"></a>

## Direct properties — protocol_port_range / 222202122123 / 3

<a id="canonical-2220322102131333-1132210022213210-2131321211222133-0213113322313332-2233101200022102-1301321123020010-1101010002120312-3011320322333133"></a>

<a id="canonical-1002313001301111-2312011020021013-2121010313130022-2332333101002200-0030321200030202-1021310233003021-2200133113310012-3113203121112031"></a>

## port_ranges property — protocol_port_range / 222202122123 / 4

Type: `["list", "string"]`. Optional.

List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-1333033030000331-1013212331322222-3230213021030030-2132230012211102-3120222110120111-2000322003131212-3033003322212123-1202311331332212"></a>

<a id="canonical-0201102303212023-3320220233111031-0300312330111033-0133320223132230-0101220320303003-3011321123323212-0112302303313021-0103301020213332"></a>

## protocol property — protocol_port_range / 222202122123 / 5

Type: `"string"`. Optional.

\[Enum: ALL|TCP|UDP|ICMP\] Protocol in IP packet to be used as match criteria Values are TCP, UDP,
and icmp. Possible values are \`ALL\`, \`TCP\`, \`UDP\`, \`ICMP\`.

Upstream description:

Protocol in IP packet to be used as match criteria Values are TCP, UDP, and icmp.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ALL",
    "TCP",
    "UDP",
    "ICMP"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ALL",
    "TCP",
    "UDP",
    "ICMP"
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  }
}
```

<a id="canonical-0321311020200332-3221001203323112-1012130212330221-2332102030201121-1110223133311202-0103321131102000-1121302223031320-2321220303111332"></a>

## Next pages — protocol_port_range / 222202122123 / 6

- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-3300312010233011-1332123212033122-3330303200202222-0112211301022200-0121210323131100-1100322331002021-3232201003132022-3201310313012123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131103113220332-3332102200212010-3321001021210001-2220132113330311-2333322122103233-1132301102033221-3320331021110120-1113201103203331"></a>

## endpoint — endpoint / 333221132011 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- endpoint

<a id="canonical-0001103100311123-1022002020001032-0200112303300210-0101102101220010-3022100133110122-0120021312002311-2213030202333331-3013010130000100"></a>

Type: `"object"`. single nested block, Optional.

Shape of the endpoint choices for a view.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("any",
    "inside_endpoints"),
  validators.ConflictingObjectAttributes("any",
    "label_selector"),
  validators.ConflictingObjectAttributes("any",
    "outside_endpoints"),
  validators.ConflictingObjectAttributes("any",
    "prefix_list"),
  validators.ConflictingObjectAttributes("inside_endpoints",
    "label_selector"),
  validators.ConflictingObjectAttributes("inside_endpoints",
    "outside_endpoints"),
  validators.ConflictingObjectAttributes("inside_endpoints",
    "prefix_list"),
  validators.ConflictingObjectAttributes("label_selector",
    "outside_endpoints"),
  validators.ConflictingObjectAttributes("label_selector",
    "prefix_list"),
  validators.ConflictingObjectAttributes("outside_endpoints",
    "prefix_list")}
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
  "x-ves-oneof-field-endpoint_choice": "[\"any\",\"inside_endpoints\",\"label_selector\",\"outside_endpoints\",\"prefix_list\"]"
}
```

Terraform syntax:

```terraform
endpoint {
  # Configure direct properties listed below.
}
```

<a id="canonical-0201123123120030-1202031112100021-2323302033200003-2213332120013001-0012122330231100-1133130033231023-2022122302113321-3300123020231221"></a>

## Direct properties — endpoint / 333221132011 / 3

- [any](resources--network_policy_view--reference--group-001.md#canonical-2223331013222303-2002212323031111-3001001220000333-1033202233330120-3011200101221222-2003332210021030-0211132200330120-1333110022331301): complete subsection reference.

- [inside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-0212031231300103-0131103033232231-3123111111003223-3110320333310210-0323303213000233-1320011322331022-0330333130232000-1010202322012210): complete subsection reference.

- [label_selector](resources--network_policy_view--reference--group-001.md#canonical-1023102030123222-3122020332012122-3233221013200010-3202230320033011-3001233012330223-2010303122220213-3020123311300220-0211300303320012): complete subsection reference.

- [outside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-2112000022120120-0200101111220011-3013033011113300-0332020332303221-3030013232221120-1130322330031322-3210003023003333-0001030220223321): complete subsection reference.

- [prefix_list](resources--network_policy_view--reference--group-001.md#canonical-2313120223033033-1330123331220210-3310200222331312-1121223022332023-0220330312000011-0113002220013133-1102323120111223-2123330320332221): complete subsection reference.

<a id="canonical-0133121003131331-0123301200223013-3220130233220203-2003121122222131-2002332011230011-1230320303112013-2213210231031221-3002023010200003"></a>

## Next pages — endpoint / 333221132011 / 4

- [endpoint.any](resources--network_policy_view--reference--group-001.md#canonical-2223331013222303-2002212323031111-3001001220000333-1033202233330120-3011200101221222-2003332210021030-0211132200330120-1333110022331301)
- [endpoint.inside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-0212031231300103-0131103033232231-3123111111003223-3110320333310210-0323303213000233-1320011322331022-0330333130232000-1010202322012210)
- [endpoint.label_selector](resources--network_policy_view--reference--group-001.md#canonical-1023102030123222-3122020332012122-3233221013200010-3202230320033011-3001233012330223-2010303122220213-3020123311300220-0211300303320012)
- [endpoint.outside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-2112000022120120-0200101111220011-3013033011113300-0332020332303221-3030013232221120-1130322330031322-3210003023003333-0001030220223321)
- [endpoint.prefix_list](resources--network_policy_view--reference--group-001.md#canonical-2313120223033033-1330123331220210-3310200222331312-1121223022332023-0220330312000011-0113002220013133-1102323120111223-2123330320332221)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-2223331013222303-2002212323031111-3001001220000333-1033202233330120-3011200101221222-2003332210021030-0211132200330120-1333110022331301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2330302121333330-0113200031211013-3101010220012311-0121220112333211-1032321132112032-0312112310313010-1012021011123112-1133002103213021"></a>

## endpoint.any — any / 013100123213 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [endpoint](resources--network_policy_view--reference--group-001.md#canonical-3300312010233011-1332123212033122-3330303200202222-0112211301022200-0121210323131100-1100322331002021-3232201003132022-3201310313012123)
- endpoint.any

<a id="canonical-0011021021131320-2110310312312030-0012022303103212-3010031113021132-0321011203123301-1303001103320331-2322010200200310-2000200032113012"></a>

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

<a id="canonical-0233013123331113-1331331121332211-3001013002302312-3022102303131202-1011203333002301-1310122121121211-0220311010101120-1111021222030002"></a>

## Direct properties — any / 013100123213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2103323312100220-1312021303123331-3202100311321113-1321000002032330-2301200332302013-3202323310032113-3311030102102121-0303022033133313"></a>

## Next pages — any / 013100123213 / 4

- [endpoint](resources--network_policy_view--reference--group-001.md#canonical-3300312010233011-1332123212033122-3330303200202222-0112211301022200-0121210323131100-1100322331002021-3232201003132022-3201310313012123)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-0212031231300103-0131103033232231-3123111111003223-3110320333310210-0323303213000233-1320011322331022-0330333130232000-1010202322012210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0311002001111132-2100012022130330-2011312123033020-1332123011030203-0120330031103132-3303122231333320-2313213222111122-2320221202300231"></a>

## endpoint.inside_endpoints — inside_endpoints / 130330101213 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [endpoint](resources--network_policy_view--reference--group-001.md#canonical-3300312010233011-1332123212033122-3330303200202222-0112211301022200-0121210323131100-1100322331002021-3232201003132022-3201310313012123)
- endpoint.inside_endpoints

<a id="canonical-2320011312210032-2032011121120002-3030310223011231-2021013031101200-3013121011100102-1032022333101203-3332121033223300-0132122002111211"></a>

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
inside_endpoints = {}
```

<a id="canonical-0011332223001112-0211023231333303-1230202312200332-0001212123311012-2133021133012023-2032110330013213-1200123003021033-3031312303110332"></a>

## Direct properties — inside_endpoints / 130330101213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1232102103022103-1111032323002230-2012303133200121-3112103102313321-0000013033032321-3312130312031132-1120313313013022-3202100131100130"></a>

## Next pages — inside_endpoints / 130330101213 / 4

- [endpoint](resources--network_policy_view--reference--group-001.md#canonical-3300312010233011-1332123212033122-3330303200202222-0112211301022200-0121210323131100-1100322331002021-3232201003132022-3201310313012123)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-1023102030123222-3122020332012122-3233221013200010-3202230320033011-3001233012330223-2010303122220213-3020123311300220-0211300303320012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031020322231021-3222320120203200-2002223122020120-1103121000202013-3002120222021211-1210112310013031-1313000321011301-3200030021211113"></a>

## endpoint.label_selector — label_selector / 331112223100 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [endpoint](resources--network_policy_view--reference--group-001.md#canonical-3300312010233011-1332123212033122-3330303200202222-0112211301022200-0121210323131100-1100322331002021-3232201003132022-3201310313012123)
- endpoint.label_selector

<a id="canonical-1331230011301201-1322321021332110-0010312332112002-3132010231010203-3031022030320200-3120211201231321-3321010233132232-2123331210032100"></a>

Type: `"object"`. single nested block, Optional.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
label_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-2311231230210023-1112113111232303-0100020203031330-2022020120001020-1333203101131103-2130002312032223-1223113121311333-3112130332032231"></a>

## Direct properties — label_selector / 331112223100 / 3

<a id="canonical-0103121001200002-0103101321120101-1333002210200030-0322011330220113-0130331130211203-0002122030022333-0023023332321013-0100213103130122"></a>

<a id="canonical-3002320010233221-2001332002113223-1123303330203232-3220211303200123-3130230313032103-0311322323201013-0022030110210222-3201202122312102"></a>

## expressions property — label_selector / 331112223100 / 4

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-3120330332320032-2212230133030113-2321233103122223-0222113033132103-1131130111013303-3332020021210132-2102231220031113-0312121023020120"></a>

## Next pages — label_selector / 331112223100 / 5

- [endpoint](resources--network_policy_view--reference--group-001.md#canonical-3300312010233011-1332123212033122-3330303200202222-0112211301022200-0121210323131100-1100322331002021-3232201003132022-3201310313012123)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-2112000022120120-0200101111220011-3013033011113300-0332020332303221-3030013232221120-1130322330031322-3210003023003333-0001030220223321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1000012132323000-0001231222301310-3220232010222102-2003323320231311-1100023221323230-3220122003101310-3310031223233232-0000010303231031"></a>

## endpoint.outside_endpoints — outside_endpoints / 100330323111 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [endpoint](resources--network_policy_view--reference--group-001.md#canonical-3300312010233011-1332123212033122-3330303200202222-0112211301022200-0121210323131100-1100322331002021-3232201003132022-3201310313012123)
- endpoint.outside_endpoints

<a id="canonical-3013333322033313-3223112002202113-3030112120202002-1033130202232203-0312032232301322-1122130031001300-1323013321112220-3322303322330322"></a>

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
outside_endpoints = {}
```

<a id="canonical-0022000002002100-1031320200232321-1202130203012113-0131302130010032-3122313032312213-1032322321123213-2031232230330331-1323023110331332"></a>

## Direct properties — outside_endpoints / 100330323111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2031020322132033-3310303130210221-1221133123233112-0301331222312221-2202003310021113-0221222113102101-0111001300220132-2311011122321302"></a>

## Next pages — outside_endpoints / 100330323111 / 4

- [endpoint](resources--network_policy_view--reference--group-001.md#canonical-3300312010233011-1332123212033122-3330303200202222-0112211301022200-0121210323131100-1100322331002021-3232201003132022-3201310313012123)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-2313120223033033-1330123331220210-3310200222331312-1121223022332023-0220330312000011-0113002220013133-1102323120111223-2123330320332221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3120101213322210-0001332311211310-2111222123232303-3032230113331132-1103203313121322-0310333313301303-2310231111232110-2230013231110031"></a>

## endpoint.prefix_list — prefix_list / 211103122130 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [endpoint](resources--network_policy_view--reference--group-001.md#canonical-3300312010233011-1332123212033122-3330303200202222-0112211301022200-0121210323131100-1100322331002021-3232201003132022-3201310313012123)
- endpoint.prefix_list

<a id="canonical-2332101330132211-2022223210230102-3201223321021013-0302011313231232-2131302031200023-0320110311003300-2130213332233032-1202330210320130"></a>

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
prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0310120122021330-2213133022322102-0303033022113210-3302101221322311-2321323020212302-1323233011112011-1123330113233232-0020222002312123"></a>

## Direct properties — prefix_list / 211103122130 / 3

<a id="canonical-3320320313332022-0301131330102203-2111332030110312-3313200012100023-3220222312331013-3300133311203032-1012033120111102-0203001123021211"></a>

<a id="canonical-3001333131212021-3100112211030010-3013221131101123-2211233232232233-1130130123232301-3232203311123201-1032211323022131-3011010300120023"></a>

## prefixes property — prefix_list / 211103122130 / 4

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

<a id="canonical-0101302203003112-2023120111220102-3222212222330022-0213301231230321-3031332013022312-2312220333002111-2023220022122230-1002312131302200"></a>

## Next pages — prefix_list / 211103122130 / 5

- [endpoint](resources--network_policy_view--reference--group-001.md#canonical-3300312010233011-1332123212033122-3330303200202222-0112211301022200-0121210323131100-1100322331002021-3232201003132022-3201310313012123)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211313311303112-3120230311102132-2223121101011021-2333233321221233-3323202013313011-2133323323221212-1301131023000023-2030013302132230"></a>

## ingress_rules — ingress_rules / 112013030220 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- ingress_rules

<a id="canonical-3011103302300222-0333123202233311-3321223300032122-2112332210101103-3310113131131002-2003013323003131-2332200132322031-0313101210313120"></a>

Type: `"object"`. list nested block, Optional.

Ordered list of rules applied to connections to policy endpoints.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "all_traffic"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "all_udp_traffic"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "applications"),
  validators.ConflictingListObjectAttributes("all_tcp_traffic",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("all_traffic",
    "all_udp_traffic"),
  validators.ConflictingListObjectAttributes("all_traffic",
    "applications"),
  validators.ConflictingListObjectAttributes("all_traffic",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("all_udp_traffic",
    "applications"),
  validators.ConflictingListObjectAttributes("all_udp_traffic",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("any",
    "inside_endpoints"),
  validators.ConflictingListObjectAttributes("any",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("any",
    "label_selector"),
  validators.ConflictingListObjectAttributes("any",
    "outside_endpoints"),
  validators.ConflictingListObjectAttributes("any",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("applications",
    "protocol_port_range"),
  validators.ConflictingListObjectAttributes("inside_endpoints",
    "ip_prefix_set"),
  validators.ConflictingListObjectAttributes("inside_endpoints",
    "label_selector"),
  validators.ConflictingListObjectAttributes("inside_endpoints",
    "outside_endpoints"),
  validators.ConflictingListObjectAttributes("inside_endpoints",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "label_selector"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "outside_endpoints"),
  validators.ConflictingListObjectAttributes("ip_prefix_set",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("label_selector",
    "outside_endpoints"),
  validators.ConflictingListObjectAttributes("label_selector",
    "prefix_list"),
  validators.ConflictingListObjectAttributes("outside_endpoints",
    "prefix_list")}
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
ingress_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-2100302332123022-3101022033301312-0201232323201130-3112320031113111-0233230002010222-0330202132111322-1210321110222000-2000302210032113"></a>

## Direct properties — ingress_rules / 112013030220 / 3

<a id="canonical-2300233102010202-3110002031230200-3333103011122311-3103001123313120-1023320131211032-2321201333131023-1012133312201303-2032013022210303"></a>

<a id="canonical-2301001232220102-0120002312330212-1020102000301032-1221022300111010-1303231301031102-0132221021112321-0230221220130121-1010221203113333"></a>

## action property — ingress_rules / 112013030220 / 4

Type: `"string"`. Optional.

\[Enum: DENY|ALLOW\] Network policy rule action configures the action to be taken on rule match
Apply deny action on rule match Apply allow action on rule match. Possible values are \`DENY\`,
\`ALLOW\`. Defaults to \`DENY\`.

Upstream description:

Network policy rule action configures the action to be taken on rule match

Apply deny action on rule match Apply allow action on rule match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DENY",
    "ALLOW"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "DENY",
  "enum": [
    "DENY",
    "ALLOW"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [adv_action](resources--network_policy_view--reference--group-001.md#canonical-2221201022221322-1203331112301322-2230223112212012-2231233033133130-3312312301113323-3012311113300322-1200022330023011-3320103102003012): complete subsection reference.

- [all_tcp_traffic](resources--network_policy_view--reference--group-001.md#canonical-2333113102103122-0102000132323310-2222010213002101-2121201130032232-0332323023013333-0133300302302131-2202202311223232-2223023011220222): complete subsection reference.

- [all_traffic](resources--network_policy_view--reference--group-001.md#canonical-2001021201212330-1101122120112032-3110202311132200-1122201233322230-2000100212302123-3200131313211011-2130211011101200-2132222021133300): complete subsection reference.

- [all_udp_traffic](resources--network_policy_view--reference--group-001.md#canonical-3320201223301323-3313023012130031-0210122101120000-2122030331121000-3312101320211130-2331210101310001-2233300313031300-2003330323300033): complete subsection reference.

- [any](resources--network_policy_view--reference--group-001.md#canonical-0110023221122000-1132101133202201-3023120323102212-3020232013300203-2000221202212311-1311223103102230-2121331101302003-1333233203223211): complete subsection reference.

- [applications](resources--network_policy_view--reference--group-001.md#canonical-3321100113013322-3031310032113122-3120010100003132-0322002231012023-2131102022312012-2222013121123012-0222332321101321-2131322210202131): complete subsection reference.

- [inside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-0301123001002332-1132220200300010-0221033130112133-0320233000120200-3312220030103300-1223010201133123-0132313123232223-0112200101320102): complete subsection reference.

- [ip_prefix_set](resources--network_policy_view--reference--group-001.md#canonical-1211210213223123-3302321313332133-1032323310130311-3010323200133321-1000301113213203-3301130030303133-0012303232301301-2102212232011203): complete subsection reference.

- [label_matcher](resources--network_policy_view--reference--group-001.md#canonical-3121200322232301-2312301122303121-0013302311032330-3321102013023003-1032200112031110-1120200012022103-0200020130212312-0212010133320303): complete subsection reference.

- [label_selector](resources--network_policy_view--reference--group-001.md#canonical-1223210302033033-1103020331300023-0320212220032002-2111212122331023-2210300232233201-0333321013230302-1203123220231011-0320032220121002): complete subsection reference.

- [metadata](resources--network_policy_view--reference--group-001.md#canonical-3332022032030020-0212311133021021-1200223300231102-1111312312203303-2131211232300013-0010133103122333-0003312313302213-1111322322211300): complete subsection reference.

- [outside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-0221002002303113-2012333311200103-2301331132330201-0300112022023002-0230212323331232-1233312110022210-1132103103312221-2021113223313113): complete subsection reference.

- [prefix_list](resources--network_policy_view--reference--group-001.md#canonical-3331230311130320-0310212101222212-3000331010213321-1312123232010030-3021022221101330-1232323112202130-2300300013310113-0233211332110102): complete subsection reference.

- [protocol_port_range](resources--network_policy_view--reference--group-001.md#canonical-3032320123112320-2022212300213331-0311300303203133-1011331303231311-1232001110003000-0213223122212023-3313122032332132-1210313132220133): complete subsection reference.

<a id="canonical-3203321111123222-0031310300211313-1312233233331033-3133200310310230-0301320012232323-0133020202000032-3111323033102221-1101130211003320"></a>

## Next pages — ingress_rules / 112013030220 / 5

- [ingress_rules.adv_action](resources--network_policy_view--reference--group-001.md#canonical-2221201022221322-1203331112301322-2230223112212012-2231233033133130-3312312301113323-3012311113300322-1200022330023011-3320103102003012)
- [ingress_rules.all_tcp_traffic](resources--network_policy_view--reference--group-001.md#canonical-2333113102103122-0102000132323310-2222010213002101-2121201130032232-0332323023013333-0133300302302131-2202202311223232-2223023011220222)
- [ingress_rules.all_traffic](resources--network_policy_view--reference--group-001.md#canonical-2001021201212330-1101122120112032-3110202311132200-1122201233322230-2000100212302123-3200131313211011-2130211011101200-2132222021133300)
- [ingress_rules.all_udp_traffic](resources--network_policy_view--reference--group-001.md#canonical-3320201223301323-3313023012130031-0210122101120000-2122030331121000-3312101320211130-2331210101310001-2233300313031300-2003330323300033)
- [ingress_rules.any](resources--network_policy_view--reference--group-001.md#canonical-0110023221122000-1132101133202201-3023120323102212-3020232013300203-2000221202212311-1311223103102230-2121331101302003-1333233203223211)
- [ingress_rules.applications](resources--network_policy_view--reference--group-001.md#canonical-3321100113013322-3031310032113122-3120010100003132-0322002231012023-2131102022312012-2222013121123012-0222332321101321-2131322210202131)
- [ingress_rules.inside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-0301123001002332-1132220200300010-0221033130112133-0320233000120200-3312220030103300-1223010201133123-0132313123232223-0112200101320102)
- [ingress_rules.ip_prefix_set](resources--network_policy_view--reference--group-001.md#canonical-1211210213223123-3302321313332133-1032323310130311-3010323200133321-1000301113213203-3301130030303133-0012303232301301-2102212232011203)
- [ingress_rules.label_matcher](resources--network_policy_view--reference--group-001.md#canonical-3121200322232301-2312301122303121-0013302311032330-3321102013023003-1032200112031110-1120200012022103-0200020130212312-0212010133320303)
- [ingress_rules.label_selector](resources--network_policy_view--reference--group-001.md#canonical-1223210302033033-1103020331300023-0320212220032002-2111212122331023-2210300232233201-0333321013230302-1203123220231011-0320032220121002)
- [ingress_rules.metadata](resources--network_policy_view--reference--group-001.md#canonical-3332022032030020-0212311133021021-1200223300231102-1111312312203303-2131211232300013-0010133103122333-0003312313302213-1111322322211300)
- [ingress_rules.outside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-0221002002303113-2012333311200103-2301331132330201-0300112022023002-0230212323331232-1233312110022210-1132103103312221-2021113223313113)
- [ingress_rules.prefix_list](resources--network_policy_view--reference--group-001.md#canonical-3331230311130320-0310212101222212-3000331010213321-1312123232010030-3021022221101330-1232323112202130-2300300013310113-0233211332110102)
- [ingress_rules.protocol_port_range](resources--network_policy_view--reference--group-001.md#canonical-3032320123112320-2022212300213331-0311300303203133-1011331303231311-1232001110003000-0213223122212023-3313122032332132-1210313132220133)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-2221201022221322-1203331112301322-2230223112212012-2231233033133130-3312312301113323-3012311113300322-1200022330023011-3320103102003012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233032132330320-2311002202320001-3321010122023331-1011202221201033-0111133030103121-2210120113220322-3212231032233013-0122023023233213"></a>

## ingress_rules.adv_action — adv_action / 200122230332 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- ingress_rules.adv_action

<a id="canonical-2210230222201200-0030311332331131-1110021021200112-2003110210033333-0133220201102020-1200323030113100-1112112333000002-2113220102320300"></a>

Type: `"object"`. single nested block, Optional.

Network Policy Rule Advanced Action provides additional OPTIONS along with RuleAction and
PBRRuleAction.

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
adv_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-3230132120230232-1220212031301003-1100010330030221-0101031322023111-2011220030033110-1000231233222200-1302303331313121-0121201302020020"></a>

## Direct properties — adv_action / 200122230332 / 3

<a id="canonical-2102321231013201-0221032213121132-0321011223000021-2110011103100330-3101102313312030-3103323202020001-1112130200311111-2010231003231133"></a>

<a id="canonical-1313310012202332-3330202230320002-0023030020200131-2021131200103022-1223322320033222-0312303012033021-2230310222103232-2012102231021003"></a>

## action property — adv_action / 200122230332 / 4

Type: `"string"`. Optional.

\[Enum: NOLOG|LOG\] Choice to choose logging or no logging This works together with option selected
via NetworkPolicyRuleAction or any other action specified x-. Possible values are \`NOLOG\`,
\`LOG\`. Defaults to \`NOLOG\`.

Upstream description:

Choice to choose logging or no logging This works together with option selected via
NetworkPolicyRuleAction or any other action specified x-example: (No Selection in
NetworkPolicyRuleAction + AdvancedAction as LOG) = LOG Only, (ALLOW/DENY in NetworkPolicyRuleAction
&#8203;+ AdvancedAction as LOG) = Log and Allow/Deny, (ALLOW/DENY in NetworkPolicyRuleAction + NOLOG in
AdvancedAction) = Allow/Deny with no log

Don't sample the traffic hitting the rule Sample the traffic hitting the rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NOLOG",
    "LOG"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NOLOG",
  "enum": [
    "NOLOG",
    "LOG"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2110211221011032-3032332133331122-3321110000203203-0303221333330001-3101010032130002-2313210332031233-1210222333012102-3310203201132001"></a>

## Next pages — adv_action / 200122230332 / 5

- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-2333113102103122-0102000132323310-2222010213002101-2121201130032232-0332323023013333-0133300302302131-2202202311223232-2223023011220222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233312020021012-2032210002233211-0033320321233113-3202010130230003-2023210122121333-2121120222203012-1220302211302310-3200223303021103"></a>

## ingress_rules.all_tcp_traffic — all_tcp_traffic / 030221303222 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- ingress_rules.all_tcp_traffic

<a id="canonical-1122011302333321-1103312301323033-0303033110321212-3101133300130321-0320030230032203-3232101231001312-1321221221032210-3221120112320132"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all tcp traffic.

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
all_tcp_traffic = {}
```

<a id="canonical-0202222331021232-0301021022120312-3001003332323333-3100313102322101-2221203020031133-2003201010003312-1121203333223221-1131222012021001"></a>

## Direct properties — all_tcp_traffic / 030221303222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3300011010022012-0000001213323313-3221120021111331-0303302223123323-2301002221110203-3303033110212122-2320121023333212-0111232122330300"></a>

## Next pages — all_tcp_traffic / 030221303222 / 4

- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-2001021201212330-1101122120112032-3110202311132200-1122201233322230-2000100212302123-3200131313211011-2130211011101200-2132222021133300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101033323313332-0300133021032211-1331220111311332-3122222112331321-0021230113221212-3021012333222321-1032332113300211-1130003023121332"></a>

## ingress_rules.all_traffic — all_traffic / 021330000312 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- ingress_rules.all_traffic

<a id="canonical-0112301033103003-1203213101212021-2201113210102311-2101331100210013-0020103330203000-2101010020213001-0100132011101132-0133333023023331"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all traffic.

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
all_traffic = {}
```

<a id="canonical-2223100132100210-0333101222100300-3100020322313021-0130132200201101-1323133330131100-3322301000200210-1002232001133221-3032210113033213"></a>

## Direct properties — all_traffic / 021330000312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1023211131321301-2002212310030110-1033301220202103-1323331113223230-0011331131212120-2030333031222002-3211101202101123-1232200132032232"></a>

## Next pages — all_traffic / 021330000312 / 4

- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-3320201223301323-3313023012130031-0210122101120000-2122030331121000-3312101320211130-2331210101310001-2233300313031300-2003330323300033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0113321031331211-3003111331133200-0233121013100232-1221122233100010-0232023300031230-2233200232030300-2123132200313131-1030103311231031"></a>

## ingress_rules.all_udp_traffic — all_udp_traffic / 031222130211 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- ingress_rules.all_udp_traffic

<a id="canonical-0000302302211202-1130200201012112-2122230201303011-2321300103031002-2322000331220012-1303331211301111-0123302203013103-1322233220220311"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all udp traffic.

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
all_udp_traffic = {}
```

<a id="canonical-1021233200303230-1323010022111011-3223122311120320-2211301200112113-1201003101130121-1110232000001003-2132130212330312-1313322112022322"></a>

## Direct properties — all_udp_traffic / 031222130211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1300000103102011-2320210020002212-3101020103230330-1333222131320322-2200302020233311-2233232211311132-1222103103230210-2031220231232230"></a>

## Next pages — all_udp_traffic / 031222130211 / 4

- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-0110023221122000-1132101133202201-3023120323102212-3020232013300203-2000221202212311-1311223103102230-2121331101302003-1333233203223211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012211132020202-1232333221301211-2311032120133102-2013200120303230-1022331322112300-1012111310230100-1033322033321231-2210201233210333"></a>

## ingress_rules.any — any / 023303003222 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- ingress_rules.any

<a id="canonical-3212202002113000-2302020200203320-2203010302213003-3203201223333333-0001220201021122-2000210023000012-3002311211123302-1023301123320210"></a>

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

<a id="canonical-0130323333200302-3003212022320302-1123323122132122-0313121013022022-3213000023030212-3330120132100203-2332313230220123-3013132301230013"></a>

## Direct properties — any / 023303003222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3121133220132120-2311201032300020-1232211213032000-3233330211223300-3311213123101010-2120302320232110-3003213202102233-3132321133331203"></a>

## Next pages — any / 023303003222 / 4

- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-3321100113013322-3031310032113122-3120010100003132-0322002231012023-2131102022312012-2222013121123012-0222332321101321-2131322210202131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301232133212232-3122012011032230-1010110232133103-0331131212021022-2230332222032103-1021000020123233-2103111203232312-2212220122302020"></a>

## ingress_rules.applications — applications / 021332122313 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- ingress_rules.applications

<a id="canonical-2011223130023002-0322120103300111-1003233332221030-1103031121032110-3320333110221313-1313301310232010-1011131030232322-0000031212132320"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for applications.

Upstream description:

Application protocols like HTTP, SNMP.

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
applications {
  # Configure direct properties listed below.
}
```

<a id="canonical-0033203033333313-1113001010302212-2032020230112221-0321213230303020-2002101133032222-3002031010303322-3313221110012031-2102300103200312"></a>

## Direct properties — applications / 021332122313 / 3

<a id="canonical-0122002101023202-3000320300022130-1323000131023322-1222323311232211-3332330221121131-0011202131210010-1333200303110013-1333210301001300"></a>

<a id="canonical-0131203330110331-0121311032222323-2100231021132101-1300300022322010-1201331110201002-3113100221112222-3311332003312332-2111101001310131"></a>

## applications property — applications / 021332122313 / 4

Type: `["list", "string"]`. Optional.

\[Enum: APPLICATION\_HTTP|APPLICATION\_HTTPS|APPLICATION\_SNMP|APPLICATION\_DNS\] Application
Protocols. Application protocols like HTTP, SNMP. Possible values are \`APPLICATION\_HTTP\`,
\`APPLICATION\_HTTPS\`, \`APPLICATION\_SNMP\`, \`APPLICATION\_DNS\`. Defaults to
\`APPLICATION\_HTTP\`.

Upstream description:

Application protocols like HTTP, SNMP.

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

<a id="canonical-2202131322032103-3013111231030321-3201033202013322-2112203312221122-2002133022321033-1010030332222132-3310022321213110-1220322322123203"></a>

## Next pages — applications / 021332122313 / 5

- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-0301123001002332-1132220200300010-0221033130112133-0320233000120200-3312220030103300-1223010201133123-0132313123232223-0112200101320102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1100132210203322-3122111033333110-2023202120202133-0111230302311132-1330202320101013-0302223230030323-2210310003202032-2303323111333203"></a>

## ingress_rules.inside_endpoints — inside_endpoints / 323002121100 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- ingress_rules.inside_endpoints

<a id="canonical-0033221223013132-2130312031221222-0312321200201002-2213300212132112-0111000033303120-1021202133311111-1012200112220210-2233012221221203"></a>

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
inside_endpoints = {}
```

<a id="canonical-3202202003113131-0303231003222133-3310222120103223-2203333132000332-3330100220021012-0313230320203303-3220323020211201-3301232001311212"></a>

## Direct properties — inside_endpoints / 323002121100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2032212103330130-0232122102001110-1130123033220311-1122122121203302-0121110131332123-2113322213133313-0123120003303202-0201101133223223"></a>

## Next pages — inside_endpoints / 323002121100 / 4

- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-1211210213223123-3302321313332133-1032323310130311-3010323200133321-1000301113213203-3301130030303133-0012303232301301-2102212232011203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210003221000232-2321101022101122-3012030231012231-2332103023103111-2002221122012123-2113030020001321-3032121130130202-0313211113331112"></a>

## ingress_rules.ip_prefix_set — ip_prefix_set / 132322032212 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- ingress_rules.ip_prefix_set

<a id="canonical-0333103123302002-3310201122332301-3212313310011122-3321213031330231-2003033030030102-2013202330331031-3202232012113212-0201303331321332"></a>

Type: `"object"`. single nested block, Optional.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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
ip_prefix_set {
  # Configure direct properties listed below.
}
```

<a id="canonical-3030101021321011-1332012011131031-1320101112031332-1312302203230022-1013032313033303-0001100231032111-2020203213203231-2130133223312000"></a>

## Direct properties — ip_prefix_set / 132322032212 / 3

- [ref](resources--network_policy_view--reference--group-001.md#canonical-3310222211223322-0002310303021002-2230020133013303-3020232303311001-0203210201132301-0213121022121021-1000302201332021-1333330001323100): complete subsection reference.

<a id="canonical-3100003331311030-0031101203220211-1120110222330303-3122322030201330-2111311011032033-0001030001313122-0333203203131002-1302331331000330"></a>

## Next pages — ip_prefix_set / 132322032212 / 4

- [ingress_rules.ip_prefix_set.ref](resources--network_policy_view--reference--group-001.md#canonical-3310222211223322-0002310303021002-2230020133013303-3020232303311001-0203210201132301-0213121022121021-1000302201332021-1333330001323100)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-3310222211223322-0002310303021002-2230020133013303-3020232303311001-0203210201132301-0213121022121021-1000302201332021-1333330001323100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3020311322030211-3112132011032102-2100223303130013-1032131303020202-2223302303112333-3033112213213113-1102203003302331-2233231320103113"></a>

## ingress_rules.ip_prefix_set.ref — ref / 011323010103 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- [ingress_rules.ip_prefix_set](resources--network_policy_view--reference--group-001.md#canonical-1211210213223123-3302321313332133-1032323310130311-3010323200133321-1000301113213203-3301130030303133-0012303232301301-2102212232011203)
- ingress_rules.ip_prefix_set.ref

<a id="canonical-3200103102320121-3032230123303220-1102223011212212-1210232200320300-2332131031031122-3300321111210201-2221331120303022-3231111332202302"></a>

Type: `"object"`. list nested block, Optional.

List of references to ip\_prefix\_set objects.

Upstream description:

A list of references to ip\_prefix\_set objects.

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
ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-1223100102122300-2123233300302110-1021213002013222-1031013111300332-3331223203233131-1322102331123011-3312100200001223-0032232310023120"></a>

## Direct properties — ref / 011323010103 / 3

<a id="canonical-1312220231033031-1202333032231310-0200302232223001-0021010111230021-2033331220321032-2332110111202120-2100013123103033-1212330222320332"></a>

<a id="canonical-1200313303000333-1003321012130030-0212120030231200-1311001301121010-3032230020222301-0310230301200223-1202012322002332-3100130032113313"></a>

## kind property — ref / 011323010103 / 4

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

<a id="canonical-3112333010301100-3200333012012233-0220032033212221-2212330310210233-0203103131021030-0131010200321333-2210211103122223-1011102300123100"></a>

<a id="canonical-0230102231022130-0322123313000013-3120031201322320-2001020202221010-3301221201210102-3032133312233310-1213203100300311-0133103232322131"></a>

## name property — ref / 011323010103 / 5

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

<a id="canonical-1332102000322111-2130010322300301-3122210123332230-2312013130303212-0233210212133300-0111221312311211-2233313333221221-1000033011031311"></a>

<a id="canonical-1222101231132310-0100131031020301-2103302031130202-3122230033113101-2102200221320233-3202313003000321-2330000020223332-2301333103203331"></a>

## namespace property — ref / 011323010103 / 6

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

<a id="canonical-2311222031120300-0001010213021223-2201223331123111-1111131233220301-3100102103330231-3120231130133321-1313322110233213-3013021111102111"></a>

<a id="canonical-3200112031233011-2312011003310102-3000033003211312-3301302203110110-1021012121123203-2331212031312130-2300313322101221-0103221111020111"></a>

## tenant property — ref / 011323010103 / 7

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

<a id="canonical-3020301013222100-3011210030332121-0330330323000133-2113222100131311-3100320332013131-3201323111300321-3200021111202202-1121321233133232"></a>

<a id="canonical-3310301030312303-0330022220322332-0202223130311301-3010133233203112-1202123231003013-2023320001002030-0313003012130231-0300323130111213"></a>

## uid property — ref / 011323010103 / 8

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

<a id="canonical-1023312233300103-1033203320232210-0233302212232321-3100222100311021-3102323200232322-0032202013332033-3301322232303211-2311203100211020"></a>

## Next pages — ref / 011323010103 / 9

- [ingress_rules.ip_prefix_set](resources--network_policy_view--reference--group-001.md#canonical-1211210213223123-3302321313332133-1032323310130311-3010323200133321-1000301113213203-3301130030303133-0012303232301301-2102212232011203)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-3121200322232301-2312301122303121-0013302311032330-3321102013023003-1032200112031110-1120200012022103-0200020130212312-0212010133320303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112113323233021-2110100121200112-3121132020201313-0213130110313113-0122300023013110-3223302122201321-2111013023112311-3033132103131202"></a>

## ingress_rules.label_matcher — label_matcher / 321220212123 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- ingress_rules.label_matcher

<a id="canonical-3100223201201121-0211122332122102-1023221321133302-2102001012332012-0123302221313103-0230331312323111-0331222331301201-2323211023322233"></a>

Type: `"object"`. single nested block, Optional.

Label matcher specifies a list of label keys whose values need to match for source/client and
destination/server. Note that the actual label values are not specified and do not matter. This
allows an ability to scope grouping by the label key name.

Upstream description:

A label matcher specifies a list of label keys whose values need to match for source/client and
destination/server. Note that the actual label values are not specified and do not matter. This
allows an ability to scope grouping by the label key name.

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
label_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-2032023301211010-0232220021332030-0030223033330231-3232123023020013-1012222012000000-3003210323101120-1331202020000230-1112323333211113"></a>

## Direct properties — label_matcher / 321220212123 / 3

<a id="canonical-3221313323201032-1200131111020203-1222110133233331-3103223113232323-3122223020032231-3130110012020110-3231231201101120-0301210211113122"></a>

<a id="canonical-1012230230001320-2223302112321010-0201331130203102-2000233001202202-2131011122100031-3001111123213023-3123312220323103-3300112110332210"></a>

## keys property — label_matcher / 321220212123 / 4

Type: `["list", "string"]`. Optional.

The list of label key names that have to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2221210311303130-3111021302122132-0033132102122210-3202200132021330-1021111101123222-2210001003232330-0301232132103232-0031120222121230"></a>

## Next pages — label_matcher / 321220212123 / 5

- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-1223210302033033-1103020331300023-0320212220032002-2111212122331023-2210300232233201-0333321013230302-1203123220231011-0320032220121002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1033321331311230-2001333223220110-1311112213102132-3020112112032012-3203300231213110-2110022110221301-3223311002302132-0123223030120012"></a>

## ingress_rules.label_selector — label_selector / 301201332223 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- ingress_rules.label_selector

<a id="canonical-3310220230013032-1031212232310313-2032213123112313-0210313010332320-3303022132210121-2103301131301122-1132310221023231-1001112133221231"></a>

Type: `"object"`. single nested block, Optional.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
label_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-0020121111032012-2300201112321122-3112030021010322-2313010122232021-2120312030331223-2130210223212213-1321322033103212-2311221221323011"></a>

## Direct properties — label_selector / 301201332223 / 3

<a id="canonical-1331000121133311-2330112300223111-1213230231003321-1212233311030310-0131102331331233-0111002000102322-3213130023122102-3033300223110031"></a>

<a id="canonical-0301300300300332-2033031033013202-3323211012323001-1311230033102312-3302103001333133-0101311301132201-1310023130210022-2200232011223213"></a>

## expressions property — label_selector / 301201332223 / 4

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-1101233003331002-0123001132213020-3023030000331200-1223121012233302-1031001231303010-1223231323233313-0323303011032321-3321220231322032"></a>

## Next pages — label_selector / 301201332223 / 5

- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-3332022032030020-0212311133021021-1200223300231102-1111312312203303-2131211232300013-0010133103122333-0003312313302213-1111322322211300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213223020203322-3212211200320230-0003313123202002-3321332232032211-2202113322331302-0002212130231312-1020033032102222-0310203200330303"></a>

## ingress_rules.metadata — metadata / 023233303233 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- ingress_rules.metadata

<a id="canonical-3132201022011010-2122212031333301-0022130123031221-2332202020220310-0022300311313312-2111311212100332-1220202031302201-0303033023302023"></a>

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

<a id="canonical-0030012330011331-0130210003002031-0232111133001232-1003021030320230-0032302321200012-2233130030331031-0112020100213120-1331210331333310"></a>

## Direct properties — metadata / 023233303233 / 3

<a id="canonical-1030100303231003-1102031230101102-3202103220133011-3001100213232302-1333131221313032-2110022331311021-0010230133212231-1011111120323031"></a>

<a id="canonical-3022120222013102-2330032313122231-0321301310303303-2133101231020232-3213020020101012-2020221311031012-2131201332302033-0223000031010030"></a>

## description_spec property — metadata / 023233303233 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-0112230121221113-1113303201103130-1002312212333222-0322110100031312-1211013210120000-3321101223000000-1003301222323220-0310203320211130"></a>

<a id="canonical-3311211203033002-0202320012121231-0022220021013102-3010332323220322-3111110330130212-2001123012121033-1232203301203011-3112122213202213"></a>

## name property — metadata / 023233303233 / 5

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

<a id="canonical-1120113312010002-1010210032022122-2001200103311103-3012033301221011-2230201130121101-1223110023311102-0132003322113221-1013233202033111"></a>

## Next pages — metadata / 023233303233 / 6

- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-0221002002303113-2012333311200103-2301331132330201-0300112022023002-0230212323331232-1233312110022210-1132103103312221-2021113223313113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202310232001200-2330012301111122-0211310210100110-3132130011333011-0221013211320201-2011333222122113-1312122333201311-2011212120003133"></a>

## ingress_rules.outside_endpoints — outside_endpoints / 030121023000 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- ingress_rules.outside_endpoints

<a id="canonical-1202222230310011-2010012031011101-2232320232332031-0113213233023313-2310200232303121-2032310211030321-0102112203212321-3211120232221101"></a>

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
outside_endpoints = {}
```

<a id="canonical-1002302331030211-0202331000321322-0132230030202132-1300323002032020-0012202101331123-0203123321331223-0220112221003100-2222031210013323"></a>

## Direct properties — outside_endpoints / 030121023000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1202221120030203-1002100101300010-3332313232320222-3113323211031301-2021002300002123-1330203003120033-3321202310333030-3021123000120031"></a>

## Next pages — outside_endpoints / 030121023000 / 4

- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-3331230311130320-0310212101222212-3000331010213321-1312123232010030-3021022221101330-1232323112202130-2300300013310113-0233211332110102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0231203013101320-1333320300311303-1220231123032002-2133123133331032-2011011111100011-2211023312003321-0011200320012202-2130113233331102"></a>

## ingress_rules.prefix_list — prefix_list / 010312331001 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- ingress_rules.prefix_list

<a id="canonical-3023321000233203-3021102310232203-3200000103121110-1000023310120111-0010100023122123-3201210003103102-1002322112232211-0232212100112000"></a>

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
prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0131231212210203-0000223210300303-1010001133303122-2010011032231102-2122232213121323-3320200022101310-0000113020022211-1322233300030233"></a>

## Direct properties — prefix_list / 010312331001 / 3

<a id="canonical-2111323233101123-2003222212331022-3233132322113112-1001202010201213-1133320232021100-0301231211203020-2211311112133031-0133202103331030"></a>

<a id="canonical-2313320303203220-1111303010013302-0113331010022020-1230022320323203-1322232130112222-2330303220223003-0112310212000213-2230030002133010"></a>

## prefixes property — prefix_list / 010312331001 / 4

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

<a id="canonical-0311210020230201-0301001201311122-3022133033232313-0021112031233121-3320113012023101-3111032202233010-1331223200233033-0201312020303013"></a>

## Next pages — prefix_list / 010312331001 / 5

- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-3032320123112320-2022212300213331-0311300303203133-1011331303231311-1232001110003000-0213223122212023-3313122032332132-1210313132220133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313232110113122-3100021212013120-2113310331211232-1233310220211202-1033113330101303-0032113103011230-1102131123201133-0312312312332123"></a>

## ingress_rules.protocol_port_range — protocol_port_range / 212321100022 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- ingress_rules.protocol_port_range

<a id="canonical-0021001012103200-3002223020013310-3313022221231321-3301312123110300-3011332031231120-3201100100120121-0001300133323202-2201332021301013"></a>

Type: `"object"`. single nested block, Optional.

Protocol and Port. Protocol and Port ranges.

Upstream description:

Protocol and Port ranges.

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
protocol_port_range {
  # Configure direct properties listed below.
}
```

<a id="canonical-0220311213122021-1003322203232123-3303323012312330-2323033320333202-2212132300110300-0020200031031233-0300222231010300-3112112331333232"></a>

## Direct properties — protocol_port_range / 212321100022 / 3

<a id="canonical-1222210113200213-2313203221320222-3122301202220021-2021010322021332-1032322103322320-3202202203232020-3203320332321303-1330020313133311"></a>

<a id="canonical-3130301302320222-0301131331322023-3201321201310130-2233303003312221-2313232313000202-1032100212302321-0322120032022013-3332031030033130"></a>

## port_ranges property — protocol_port_range / 212321100022 / 4

Type: `["list", "string"]`. Optional.

List of port ranges. Each range is a single port or a pair of start and end ports e.g. 8080-8192.

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.port_range": "true",
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

<a id="canonical-3233230201333110-2021012133100121-0222003130133223-1232113022022120-3202021131013201-1132013032321223-2101003120321010-2103031112300022"></a>

<a id="canonical-1033112321223103-1201100001202103-3211110031112212-2333131233321211-3003003320210333-1102202121322032-3201203313331130-3011112200101032"></a>

## protocol property — protocol_port_range / 212321100022 / 5

Type: `"string"`. Optional.

\[Enum: ALL|TCP|UDP|ICMP\] Protocol in IP packet to be used as match criteria Values are TCP, UDP,
and icmp. Possible values are \`ALL\`, \`TCP\`, \`UDP\`, \`ICMP\`.

Upstream description:

Protocol in IP packet to be used as match criteria Values are TCP, UDP, and icmp.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ALL",
    "TCP",
    "UDP",
    "ICMP"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ALL",
    "TCP",
    "UDP",
    "ICMP"
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  }
}
```

<a id="canonical-3130001230301332-1101012022330322-2010101001231331-3322131232100132-3323113222110121-2012002230310303-0232122132303202-2202212233330323"></a>

## Next pages — protocol_port_range / 212321100022 / 6

- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)

<a id="canonical-2031001202231122-1230130201233212-0211330111232122-2233011300203033-2021113233233022-0313233323303300-1003201003202013-1113233202101200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3232311211113322-0200322213311003-2130313011232020-3030331300133121-3120121220113323-1012103032223103-3031011233032030-0120111310112311"></a>

## timeouts — timeouts / 130213323122 / 2

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- timeouts

<a id="canonical-0122023323013000-2111320203103200-0203101203131222-3232020320230320-1021213223031021-0221210300032321-1002323231303101-3200330211202300"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2211203201013202-0022320013003202-0320331320011122-3201113002203202-1212312001020002-2231000220211233-2010002201220030-2121233311122311"></a>

## Direct properties — timeouts / 130213323122 / 3

<a id="canonical-1220133213331221-1002311301311112-3321221202320001-3211122213010333-1021233013102202-0022301012023302-0233213323001301-1310222123202111"></a>

<a id="canonical-3210233220121002-0201203021320133-0001331122120210-3132313331223113-0002100303123231-1101320211120321-3000011221331213-0210322203100302"></a>

## create property — timeouts / 130213323122 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0030200120130312-1111000011331020-3331113201101300-1213233300211232-3323203002111231-0201333210310113-0213323132031013-3001012021231023"></a>

<a id="canonical-0030331132330332-0320122002020003-0330030003112102-0203032100313223-2020113002211101-0201232101332020-2300201222020022-1300323210201010"></a>

## delete property — timeouts / 130213323122 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1331021301300312-1303030110322213-3030022322321001-3300111100033112-3333332310001212-3123033220000033-1120233231330031-2011300123121311"></a>

<a id="canonical-0212223333032032-1321203011021200-2000230211000102-3133311223211131-0033013202321332-2223310120332111-3011231302333132-0333020233111020"></a>

## read property — timeouts / 130213323122 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1002210332013320-2310212323001213-0300100003132100-2013211231220210-2031233300211232-0320020012003130-0032033223220013-3330303012300011"></a>

<a id="canonical-1232210233213112-0211233122011212-3201302333310200-0300112103000221-1212333130300112-1302013300232320-0332130210331002-1021232212332010"></a>

## update property — timeouts / 130213323122 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2001122330012111-1031101313031000-3232231023000103-3022212233222203-0220023233122311-0332013320213321-1323233033303300-2312000313231313"></a>

## Next pages — timeouts / 130213323122 / 8

- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
