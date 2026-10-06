---
page_title: "xcsh_network_policy_view reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_policy_view reference."
---

# xcsh_network_policy_view reference

<a id="canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- Property reference

<a id="canonical-3133032001232203-1021201002213222-2010122110102212-1232003032211333-0003122320011123-0312221032020231-1001111202301012-3022001211200322"></a>

### Direct properties for `xcsh_network_policy_view`

<a id="canonical-0313321231302321-3200023202001230-1132311000313213-3201333212230113-2301300002212202-2120002230200321-1022101203303103-3330103002201230"></a>

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

<a id="canonical-2330031223021331-3231123023031230-3322212200003312-0133020211201013-3221030103131211-3002010202332011-2011300023311121-3213310023301202"></a>

<a id="canonical-3302203022332001-2111001222001001-3312210330322200-2001013202232310-0222211322220233-3132313223321302-3211231233201121-1023200122112210"></a>

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

<a id="canonical-3202021133132213-3002011203223012-3100313333300301-3032212031330321-1213111020201320-0013033032030332-0300020331310322-0121031023023012"></a>

<a id="canonical-1311133030023000-3220131012112113-3103002100030010-1030100223002102-3032212323222110-0030011232313223-0023200232312033-0122101222032223"></a>

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

- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213): complete subsection reference.

- [endpoint](resources--network_policy_view--reference--group-001.md#canonical-3300312010233011-1332123212033122-3330303200202222-0112211301022200-0121210323131100-1100322331002021-3232201003132022-3201310313012123): complete subsection reference.

<a id="canonical-3031030011210200-3320332030221122-3121323301101120-3203200310102113-0111031230213323-2300003123032233-1032221322020032-2022312323323212"></a>

<a id="canonical-2313100213210123-0212000303111221-0221202120110131-1301000232123212-3123101311112200-1323010030310132-2303012013103100-0312111210103333"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013): complete subsection reference.

<a id="canonical-0103002231123333-2123121123233122-2011132213311221-3301223222020313-3333032233032133-3211023033313302-0133020003213322-2201000000103202"></a>

<a id="canonical-2201022220203331-2320210301132003-3332121110201100-2320203202101000-2323031023112213-3301211220303300-0113323303002100-1312211030030003"></a>

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

<a id="canonical-0003110202032023-3210031220230011-3203301222113212-3123322003132000-3132130000111322-0032303113010020-1113212021123011-0011130211311002"></a>

<a id="canonical-1223121110300211-3123012120120132-1123320213203003-3201021030331110-0331311123112132-1131301101101031-3213313310303333-0233323122313221"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Network Policy View. Must be unique within the namespace.

Additional upstream details:

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

<a id="canonical-2103030022131000-1203230322101231-1332312330002121-2301231121000022-1320111323312323-0331332302231322-3300112322021031-3330221330033023"></a>

<a id="canonical-0000000223332021-2221112203102210-3102130210130230-1033012212333033-3022310333323103-0131130233201232-0333011121030222-2330333020302010"></a>

#### `namespace` property

Type: `"string"`. Optional, Computed.

Namespace for the Network Policy View. The F5 XC API restricts this resource to the system
namespace; it defaults to that value and may be omitted.

Additional upstream details:

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

- [timeouts](resources--network_policy_view--reference--group-001.md#canonical-2031001202231122-1230130201233212-0211330111232122-2233011300203033-2021113233233022-0313233323303300-1003201003202013-1113233202101200): complete subsection reference.

<a id="canonical-2200121030103221-1330223032100321-1033020310120223-1032123101323230-2131220120033121-0130132121032000-0233211301112012-0031212231313132"></a>

### All schema paths for `xcsh_network_policy_view`

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
| `id` | [ID](resources--network_policy_view--reference--group-001.md#canonical-3031030011210200-3320332030221122-3121323301101120-3203200310102113-0111031230213323-2300003123032233-1032221322020032-2022312323323212) |
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

<a id="canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `egress_rules` properties

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

<a id="canonical-3301321213200010-1012213200303230-0333223131122023-1301321002133032-2121111113103320-1311322332221202-0100132331023300-0122020203230020"></a>

### Direct properties for `egress_rules`

<a id="canonical-3201233030203230-1220003012300002-0110123303230111-0031111333210013-0000022131101313-1033033112231331-0112101332123232-1232202200020333"></a>

#### `egress_rules.action` property

Type: `"string"`. Optional.

\[Enum: DENY|ALLOW\] Network policy rule action configures the action to be taken on rule match
Apply deny action on rule match Apply allow action on rule match. Possible values are \`DENY\`,
\`ALLOW\`. Defaults to \`DENY\`.

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

<a id="canonical-2232011000132112-1002200220113102-1200001112021013-3033203312332210-0312310211112230-0102311133210131-0133220111313023-0302121232223013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `egress_rules.adv_action` properties

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

<a id="canonical-2111330002130201-0333132322113211-1110210022120311-0313221233103300-1221133200310203-1003221222000030-0323123121023121-0221220301311121"></a>

### Direct properties for `egress_rules.adv_action`

<a id="canonical-1303221222003202-1123012001331220-2013123321000320-2031332321211203-1102303210233011-2133000011100022-0003030001112013-3120012220033223"></a>

#### `egress_rules.adv_action.action` property

Type: `"string"`. Optional.

\[Enum: NOLOG|LOG\] Choice to choose logging or no logging This works together with option selected
via NetworkPolicyRuleAction or any other action specified x-. Possible values are \`NOLOG\`,
\`LOG\`. Defaults to \`NOLOG\`.

Additional upstream details:

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

<a id="canonical-2211323333213302-0122330231321012-2003012331133031-3032330112331303-2112033000331313-1102033003111333-2001201112112000-3123211233120312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `egress_rules.all_tcp_traffic` properties

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- egress_rules.all_tcp_traffic

<a id="canonical-3031100210213220-2013321323021230-0021302000211230-1311023131100322-3021322121233122-0200311113111100-2211300012013101-3001111131130011"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all tcp traffic.

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
all_tcp_traffic = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3110333102010022-1122120020201210-2320131323333303-3301002321111131-1032300213212010-1331110221312130-3002021012311323-3333011011001303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `egress_rules.all_traffic` properties

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- egress_rules.all_traffic

<a id="canonical-3231203022201023-3113103023330133-2300120123121011-1013322020221211-0010123201103320-1301212003030120-2123311231323221-3203032313323013"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all traffic.

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
all_traffic = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2223203031131111-3221222202113213-3311003231011301-1102302020200030-0130331011000232-3321132103233013-3220111133032330-1020012233232333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `egress_rules.all_udp_traffic` properties

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- egress_rules.all_udp_traffic

<a id="canonical-1120310323133333-0113210332200033-3123300300110021-0310201201230111-3313220133230303-0313301230011122-0213111033130223-0223011200300101"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all udp traffic.

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
all_udp_traffic = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0320102012023323-2123323033100130-0203201101010210-1123301121232300-1100101122003210-0333330331113220-1310021000112312-0330131001200022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `egress_rules.any` properties

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- egress_rules.any

<a id="canonical-2202200132303132-2313310331132101-0210030000201020-2001333102123103-0220033201100332-1331112121312012-2100120312310313-1213313032001010"></a>

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

<a id="canonical-2231120103111332-1031333103211212-2112303213120211-3301032110010030-3003202221103300-1231111330101233-3101302210133303-3313002010201213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `egress_rules.applications` properties

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- egress_rules.applications

<a id="canonical-2110323221132233-3303033230302132-1202011312202032-2032113002320313-0022013322020300-0323100221200121-1321023103001311-0200011131331222"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for applications.

Additional upstream details:

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

<a id="canonical-3320020023320302-3332333230132103-2330220003101111-2120011220302100-2102003002100230-3233232022222312-1230320322020331-1200213032100010"></a>

### Direct properties for `egress_rules.applications`

<a id="canonical-3200323030211122-3202130200233131-1220221133021300-1100213021332220-0020011031323012-3211023221030322-1123130133303200-0110003200030100"></a>

#### `egress_rules.applications.applications` property

Type: `["list", "string"]`. Optional.

\[Enum: APPLICATION\_HTTP|APPLICATION\_HTTPS|APPLICATION\_SNMP|APPLICATION\_DNS\] Application
Protocols. Application protocols like HTTP, SNMP. Possible values are \`APPLICATION\_HTTP\`,
\`APPLICATION\_HTTPS\`, \`APPLICATION\_SNMP\`, \`APPLICATION\_DNS\`. Defaults to
\`APPLICATION\_HTTP\`.

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

<a id="canonical-2210133131300203-1311111030103333-2101221313020211-3202313313223233-1033033110213023-3203321220022331-3010313101300122-2110003202320001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `egress_rules.inside_endpoints` properties

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- egress_rules.inside_endpoints

<a id="canonical-2211210300200003-2231330322232223-0232033133310320-1000113311333210-1322202332213200-3100332002210202-3121201310210313-1233200113333113"></a>

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
inside_endpoints = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2133121303032023-1302021303121323-1032331322022130-1010311332130110-1330131300233300-3113210122132120-3030030223300200-3102232030321122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `egress_rules.ip_prefix_set` properties

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- egress_rules.ip_prefix_set

<a id="canonical-2322031013313101-3120111320023003-2131110303233110-0120200303310302-2000312033311322-3123022331111000-2220223132203230-0011312311323101"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-0112212101332220-0130022110032012-2101232220223201-0223231300201003-3100110101302213-2200202300323012-0120231132200103-2111222022013121"></a>

### Direct properties for `egress_rules.ip_prefix_set`

- [ref](resources--network_policy_view--reference--group-001.md#canonical-0103211003302222-0210110311002123-0203233232131011-0021222130303020-1102203100312102-0223121212102230-0101231113332022-0212203310213032): complete subsection reference.

<a id="canonical-0103211003302222-0210110311002123-0203233232131011-0021222130303020-1102203100312102-0223121212102230-0101231113332022-0212203310213032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `egress_rules.ip_prefix_set.ref` properties

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- [egress_rules.ip_prefix_set](resources--network_policy_view--reference--group-001.md#canonical-2133121303032023-1302021303121323-1032331322022130-1010311332130110-1330131300233300-3113210122132120-3030030223300200-3102232030321122)
- egress_rules.ip_prefix_set.ref

<a id="canonical-1203321132132033-1010213320330023-0310120311320032-1122333300230322-0011233100132331-0000010210202302-0210302202201331-3203231012022113"></a>

Type: `"object"`. list nested block, Optional.

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

<a id="canonical-1311101301233212-2003201330023132-3201121223233201-3013101212133130-0230210303332201-1313201031220001-1210330303021321-0130013201322310"></a>

### Direct properties for `egress_rules.ip_prefix_set.ref`

<a id="canonical-3222333133200233-1023221020023120-0122301231322013-1120232231223132-2103210333000120-1121321011021313-2110011222111102-3303012331321012"></a>

#### `egress_rules.ip_prefix_set.ref.kind` property

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

<a id="canonical-0102122322100112-0112310303301030-1113202011323232-1202033222111131-3311000101011231-3323300001222100-2123230333211113-2331023133301200"></a>

<a id="canonical-3132313110300122-1021102212333130-1232233123033132-0312102020031103-0011011100322102-1101333321112130-2121112010313012-1311300011100333"></a>

#### `egress_rules.ip_prefix_set.ref.name` property

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

<a id="canonical-2133130322100232-2013123303223331-3223203203322321-3010011130023211-2220300100111331-3213200222220130-0002120210331211-3103100203011310"></a>

<a id="canonical-3230312201230000-2330120021210232-0010312103033130-1023103222313230-3212123330230131-3022332023001313-2032300023213322-0003201201113320"></a>

#### `egress_rules.ip_prefix_set.ref.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-3121300221231201-0332201213111312-0300020002121111-1030001120021202-0020223233001310-3330011210211121-0010331232131302-2130012132113302"></a>

<a id="canonical-2013001112121210-0233133131112211-2023030203302000-1322112101030133-0331003313303000-1011312031203220-2033220300312020-0323020210100001"></a>

#### `egress_rules.ip_prefix_set.ref.tenant` property

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

<a id="canonical-2210212230332300-3221332132110122-1230103121233302-2033333322213333-3203312002200120-1310000012130311-1033110212131222-0303221122132123"></a>

<a id="canonical-3201111230323220-1203100311311022-3032301012132233-0001002220330013-1222223103213232-3201311321100003-2101230210130332-3010212102203312"></a>

#### `egress_rules.ip_prefix_set.ref.uid` property

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

<a id="canonical-2220100223102200-3232110033110311-1312000013102023-2221303023102312-3133200332001321-1003130330230321-3302113101122102-0110322211001221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `egress_rules.label_matcher` properties

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- egress_rules.label_matcher

<a id="canonical-2031213323131011-1032221212111101-1123311011110130-0111131103223013-2233032023220213-1212311030001221-0210011222001331-2022202201201000"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-0131013021322323-2010101033111013-3003210101313231-3012003010000021-1100231013001122-1220130013111201-3123022320312103-3222130333201132"></a>

### Direct properties for `egress_rules.label_matcher`

<a id="canonical-1121122213322233-1131223231222033-2001003111221133-0301023131320300-0100330001120213-0112211033310301-2302202023313122-1332333201321231"></a>

#### `egress_rules.label_matcher.keys` property

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

<a id="canonical-0321003102212200-3312220331213011-3333131120110220-3311000313001112-1011212112233002-3023331111310321-2223330303322130-1210002202203001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `egress_rules.label_selector` properties

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- egress_rules.label_selector

<a id="canonical-3313320331122110-3100300102212213-0113211033102223-0032333032003133-3022330110320012-0323122232233132-3221303223303321-3012020000111031"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-2102112022111330-3122233213012203-2333123222110023-0111333023101223-2211201230132203-2121123132333021-1110023231310303-0223313212011202"></a>

### Direct properties for `egress_rules.label_selector`

<a id="canonical-3023100220120103-0101031311121111-1021131201130323-3232020030023330-1313203131210033-0202111112322201-1122013132232300-1122112001013110"></a>

#### `egress_rules.label_selector.expressions` property

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

<a id="canonical-1222103211120113-3103300030011023-0233333130213311-0333200023321321-1020133113303311-0111120132230302-3021320210022032-1133002121212122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `egress_rules.metadata` properties

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- egress_rules.metadata

<a id="canonical-1133333210033320-0213313223312022-3100133120200200-1311201103300223-1101100213310103-2122323211130330-3133231133321113-0121021232120311"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-2310231001323312-2321023020202322-2323120030003033-1132033332102330-2031211332101313-0231213320001330-2321231121211333-0132213233201032"></a>

### Direct properties for `egress_rules.metadata`

<a id="canonical-2100023301031322-1333213033331333-0223102233110301-2033133133301022-1323010121111300-3203220032123133-0100200101201012-1110210210010333"></a>

#### `egress_rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-1002312131132132-1021023131131020-0020111222203230-1300320321012231-1202110032221112-1222311023331002-1231020233311021-0133333200101001"></a>

<a id="canonical-3111111230023002-1333131333330213-3210223121223123-1310123231010133-3202333330221322-3200213133201311-2011221022332211-0322221202233313"></a>

#### `egress_rules.metadata.name` property

Type: `"string"`. Optional.

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

<a id="canonical-0133031212300131-1013020101300120-2113132020123301-3030301301013330-2133333333113301-1022310211302320-0233030013332220-0300132103230210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `egress_rules.outside_endpoints` properties

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- egress_rules.outside_endpoints

<a id="canonical-0131203302322221-0111312200120022-2201111101210301-0222013003213331-1312211020310222-0110320233311022-1211200032203000-3111210313000130"></a>

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
outside_endpoints = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2300333020003231-3212220233212321-2210311311032002-3203301103132011-3311321223320233-2131333121121011-1110223131100003-1200323333111232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `egress_rules.prefix_list` properties

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

<a id="canonical-2330020303030131-2300322111022320-1201031000311112-3131103031210300-3103031001011131-1101212020332123-0230322111210320-3112031330313331"></a>

### Direct properties for `egress_rules.prefix_list`

<a id="canonical-0122301232101003-3012013312302011-1123220133033301-2310003303022222-2102321201103302-3010013221222310-3021331121120020-1113232331330010"></a>

#### `egress_rules.prefix_list.prefixes` property

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

<a id="canonical-1333012223013200-2012222233331323-2101313222223312-1023133002032311-2230321032202131-0013122123103331-1331211303023100-0110312013100232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `egress_rules.protocol_port_range` properties

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [egress_rules](resources--network_policy_view--reference--group-001.md#canonical-1122013032230322-1312023003102312-3022203301130200-0113000221212100-3222322001302103-2112310210200331-1032000103302110-1300103102230213)
- egress_rules.protocol_port_range

<a id="canonical-2210312321003321-3310132233021212-1200001201130322-3302213011023133-0012222223102321-3321200131331101-2312313323123332-3033210233302312"></a>

Type: `"object"`. single nested block, Optional.

Protocol and Port. Protocol and Port ranges.

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

<a id="canonical-3020312201031321-3302302132211133-3302113133111121-1122032203021202-2032120012130110-3103213201110211-3023231321022223-1113202203202033"></a>

### Direct properties for `egress_rules.protocol_port_range`

<a id="canonical-2220322102131333-1132210022213210-2131321211222133-0213113322313332-2233101200022102-1301321123020010-1101010002120312-3011320322333133"></a>

#### `egress_rules.protocol_port_range.port_ranges` property

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

<a id="canonical-3120201020030331-0233122300000311-1021123030302001-0001232332231222-3221121221123023-1312230211030201-0212012123002022-1023021302033213"></a>

#### `egress_rules.protocol_port_range.protocol` property

Type: `"string"`. Optional.

\[Enum: ALL|TCP|UDP|ICMP\] Protocol in IP packet to be used as match criteria Values are TCP, UDP,
and icmp. Possible values are \`ALL\`, \`TCP\`, \`UDP\`, \`ICMP\`.

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
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  }
}
```

<a id="canonical-3300312010233011-1332123212033122-3330303200202222-0112211301022200-0121210323131100-1100322331002021-3232201003132022-3201310313012123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint` properties

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

<a id="canonical-2131103113220332-3332102200212010-3321001021210001-2220132113330311-2333322122103233-1132301102033221-3320331021110120-1113201103203331"></a>

### Direct properties for `endpoint`

- [any](resources--network_policy_view--reference--group-001.md#canonical-2223331013222303-2002212323031111-3001001220000333-1033202233330120-3011200101221222-2003332210021030-0211132200330120-1333110022331301): complete subsection reference.

- [inside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-0212031231300103-0131103033232231-3123111111003223-3110320333310210-0323303213000233-1320011322331022-0330333130232000-1010202322012210): complete subsection reference.

- [label_selector](resources--network_policy_view--reference--group-001.md#canonical-1023102030123222-3122020332012122-3233221013200010-3202230320033011-3001233012330223-2010303122220213-3020123311300220-0211300303320012): complete subsection reference.

- [outside_endpoints](resources--network_policy_view--reference--group-001.md#canonical-2112000022120120-0200101111220011-3013033011113300-0332020332303221-3030013232221120-1130322330031322-3210003023003333-0001030220223321): complete subsection reference.

- [prefix_list](resources--network_policy_view--reference--group-001.md#canonical-2313120223033033-1330123331220210-3310200222331312-1121223022332023-0220330312000011-0113002220013133-1102323120111223-2123330320332221): complete subsection reference.

<a id="canonical-2223331013222303-2002212323031111-3001001220000333-1033202233330120-3011200101221222-2003332210021030-0211132200330120-1333110022331301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint.any` properties

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [endpoint](resources--network_policy_view--reference--group-001.md#canonical-3300312010233011-1332123212033122-3330303200202222-0112211301022200-0121210323131100-1100322331002021-3232201003132022-3201310313012123)
- endpoint.any

<a id="canonical-0011021021131320-2110310312312030-0012022303103212-3010031113021132-0321011203123301-1303001103320331-2322010200200310-2000200032113012"></a>

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

<a id="canonical-0212031231300103-0131103033232231-3123111111003223-3110320333310210-0323303213000233-1320011322331022-0330333130232000-1010202322012210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint.inside_endpoints` properties

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [endpoint](resources--network_policy_view--reference--group-001.md#canonical-3300312010233011-1332123212033122-3330303200202222-0112211301022200-0121210323131100-1100322331002021-3232201003132022-3201310313012123)
- endpoint.inside_endpoints

<a id="canonical-2320011312210032-2032011121120002-3030310223011231-2021013031101200-3013121011100102-1032022333101203-3332121033223300-0132122002111211"></a>

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
inside_endpoints = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1023102030123222-3122020332012122-3233221013200010-3202230320033011-3001233012330223-2010303122220213-3020123311300220-0211300303320012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint.label_selector` properties

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [endpoint](resources--network_policy_view--reference--group-001.md#canonical-3300312010233011-1332123212033122-3330303200202222-0112211301022200-0121210323131100-1100322331002021-3232201003132022-3201310313012123)
- endpoint.label_selector

<a id="canonical-1331230011301201-1322321021332110-0010312332112002-3132010231010203-3031022030320200-3120211201231321-3321010233132232-2123331210032100"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-3031020322231021-3222320120203200-2002223122020120-1103121000202013-3002120222021211-1210112310013031-1313000321011301-3200030021211113"></a>

### Direct properties for `endpoint.label_selector`

<a id="canonical-0103121001200002-0103101321120101-1333002210200030-0322011330220113-0130331130211203-0002122030022333-0023023332321013-0100213103130122"></a>

#### `endpoint.label_selector.expressions` property

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

<a id="canonical-2112000022120120-0200101111220011-3013033011113300-0332020332303221-3030013232221120-1130322330031322-3210003023003333-0001030220223321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint.outside_endpoints` properties

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [endpoint](resources--network_policy_view--reference--group-001.md#canonical-3300312010233011-1332123212033122-3330303200202222-0112211301022200-0121210323131100-1100322331002021-3232201003132022-3201310313012123)
- endpoint.outside_endpoints

<a id="canonical-3013333322033313-3223112002202113-3030112120202002-1033130202232203-0312032232301322-1122130031001300-1323013321112220-3322303322330322"></a>

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
outside_endpoints = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2313120223033033-1330123331220210-3310200222331312-1121223022332023-0220330312000011-0113002220013133-1102323120111223-2123330320332221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `endpoint.prefix_list` properties

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

<a id="canonical-3120101213322210-0001332311211310-2111222123232303-3032230113331132-1103203313121322-0310333313301303-2310231111232110-2230013231110031"></a>

### Direct properties for `endpoint.prefix_list`

<a id="canonical-3320320313332022-0301131330102203-2111332030110312-3313200012100023-3220222312331013-3300133311203032-1012033120111102-0203001123021211"></a>

#### `endpoint.prefix_list.prefixes` property

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

<a id="canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ingress_rules` properties

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

<a id="canonical-0211313311303112-3120230311102132-2223121101011021-2333233321221233-3323202013313011-2133323323221212-1301131023000023-2030013302132230"></a>

### Direct properties for `ingress_rules`

<a id="canonical-2300233102010202-3110002031230200-3333103011122311-3103001123313120-1023320131211032-2321201333131023-1012133312201303-2032013022210303"></a>

#### `ingress_rules.action` property

Type: `"string"`. Optional.

\[Enum: DENY|ALLOW\] Network policy rule action configures the action to be taken on rule match
Apply deny action on rule match Apply allow action on rule match. Possible values are \`DENY\`,
\`ALLOW\`. Defaults to \`DENY\`.

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

<a id="canonical-2221201022221322-1203331112301322-2230223112212012-2231233033133130-3312312301113323-3012311113300322-1200022330023011-3320103102003012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ingress_rules.adv_action` properties

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

<a id="canonical-3233032132330320-2311002202320001-3321010122023331-1011202221201033-0111133030103121-2210120113220322-3212231032233013-0122023023233213"></a>

### Direct properties for `ingress_rules.adv_action`

<a id="canonical-2102321231013201-0221032213121132-0321011223000021-2110011103100330-3101102313312030-3103323202020001-1112130200311111-2010231003231133"></a>

#### `ingress_rules.adv_action.action` property

Type: `"string"`. Optional.

\[Enum: NOLOG|LOG\] Choice to choose logging or no logging This works together with option selected
via NetworkPolicyRuleAction or any other action specified x-. Possible values are \`NOLOG\`,
\`LOG\`. Defaults to \`NOLOG\`.

Additional upstream details:

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

<a id="canonical-2333113102103122-0102000132323310-2222010213002101-2121201130032232-0332323023013333-0133300302302131-2202202311223232-2223023011220222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ingress_rules.all_tcp_traffic` properties

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- ingress_rules.all_tcp_traffic

<a id="canonical-1122011302333321-1103312301323033-0303033110321212-3101133300130321-0320030230032203-3232101231001312-1321221221032210-3221120112320132"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all tcp traffic.

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
all_tcp_traffic = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2001021201212330-1101122120112032-3110202311132200-1122201233322230-2000100212302123-3200131313211011-2130211011101200-2132222021133300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ingress_rules.all_traffic` properties

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- ingress_rules.all_traffic

<a id="canonical-0112301033103003-1203213101212021-2201113210102311-2101331100210013-0020103330203000-2101010020213001-0100132011101132-0133333023023331"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all traffic.

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
all_traffic = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3320201223301323-3313023012130031-0210122101120000-2122030331121000-3312101320211130-2331210101310001-2233300313031300-2003330323300033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ingress_rules.all_udp_traffic` properties

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- ingress_rules.all_udp_traffic

<a id="canonical-0000302302211202-1130200201012112-2122230201303011-2321300103031002-2322000331220012-1303331211301111-0123302203013103-1322233220220311"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all udp traffic.

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
all_udp_traffic = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0110023221122000-1132101133202201-3023120323102212-3020232013300203-2000221202212311-1311223103102230-2121331101302003-1333233203223211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ingress_rules.any` properties

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- ingress_rules.any

<a id="canonical-3212202002113000-2302020200203320-2203010302213003-3203201223333333-0001220201021122-2000210023000012-3002311211123302-1023301123320210"></a>

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

<a id="canonical-3321100113013322-3031310032113122-3120010100003132-0322002231012023-2131102022312012-2222013121123012-0222332321101321-2131322210202131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ingress_rules.applications` properties

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- ingress_rules.applications

<a id="canonical-2011223130023002-0322120103300111-1003233332221030-1103031121032110-3320333110221313-1313301310232010-1011131030232322-0000031212132320"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for applications.

Additional upstream details:

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

<a id="canonical-1301232133212232-3122012011032230-1010110232133103-0331131212021022-2230332222032103-1021000020123233-2103111203232312-2212220122302020"></a>

### Direct properties for `ingress_rules.applications`

<a id="canonical-0122002101023202-3000320300022130-1323000131023322-1222323311232211-3332330221121131-0011202131210010-1333200303110013-1333210301001300"></a>

#### `ingress_rules.applications.applications` property

Type: `["list", "string"]`. Optional.

\[Enum: APPLICATION\_HTTP|APPLICATION\_HTTPS|APPLICATION\_SNMP|APPLICATION\_DNS\] Application
Protocols. Application protocols like HTTP, SNMP. Possible values are \`APPLICATION\_HTTP\`,
\`APPLICATION\_HTTPS\`, \`APPLICATION\_SNMP\`, \`APPLICATION\_DNS\`. Defaults to
\`APPLICATION\_HTTP\`.

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

<a id="canonical-0301123001002332-1132220200300010-0221033130112133-0320233000120200-3312220030103300-1223010201133123-0132313123232223-0112200101320102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ingress_rules.inside_endpoints` properties

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- ingress_rules.inside_endpoints

<a id="canonical-0033221223013132-2130312031221222-0312321200201002-2213300212132112-0111000033303120-1021202133311111-1012200112220210-2233012221221203"></a>

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
inside_endpoints = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1211210213223123-3302321313332133-1032323310130311-3010323200133321-1000301113213203-3301130030303133-0012303232301301-2102212232011203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ingress_rules.ip_prefix_set` properties

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- ingress_rules.ip_prefix_set

<a id="canonical-0333103123302002-3310201122332301-3212313310011122-3321213031330231-2003033030030102-2013202330331031-3202232012113212-0201303331321332"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-3210003221000232-2321101022101122-3012030231012231-2332103023103111-2002221122012123-2113030020001321-3032121130130202-0313211113331112"></a>

### Direct properties for `ingress_rules.ip_prefix_set`

- [ref](resources--network_policy_view--reference--group-001.md#canonical-3310222211223322-0002310303021002-2230020133013303-3020232303311001-0203210201132301-0213121022121021-1000302201332021-1333330001323100): complete subsection reference.

<a id="canonical-3310222211223322-0002310303021002-2230020133013303-3020232303311001-0203210201132301-0213121022121021-1000302201332021-1333330001323100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ingress_rules.ip_prefix_set.ref` properties

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- [ingress_rules.ip_prefix_set](resources--network_policy_view--reference--group-001.md#canonical-1211210213223123-3302321313332133-1032323310130311-3010323200133321-1000301113213203-3301130030303133-0012303232301301-2102212232011203)
- ingress_rules.ip_prefix_set.ref

<a id="canonical-3200103102320121-3032230123303220-1102223011212212-1210232200320300-2332131031031122-3300321111210201-2221331120303022-3231111332202302"></a>

Type: `"object"`. list nested block, Optional.

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

<a id="canonical-3020311322030211-3112132011032102-2100223303130013-1032131303020202-2223302303112333-3033112213213113-1102203003302331-2233231320103113"></a>

### Direct properties for `ingress_rules.ip_prefix_set.ref`

<a id="canonical-1312220231033031-1202333032231310-0200302232223001-0021010111230021-2033331220321032-2332110111202120-2100013123103033-1212330222320332"></a>

#### `ingress_rules.ip_prefix_set.ref.kind` property

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

<a id="canonical-3112333010301100-3200333012012233-0220032033212221-2212330310210233-0203103131021030-0131010200321333-2210211103122223-1011102300123100"></a>

<a id="canonical-1223100102122300-2123233300302110-1021213002013222-1031013111300332-3331223203233131-1322102331123011-3312100200001223-0032232310023120"></a>

#### `ingress_rules.ip_prefix_set.ref.name` property

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

<a id="canonical-1332102000322111-2130010322300301-3122210123332230-2312013130303212-0233210212133300-0111221312311211-2233313333221221-1000033011031311"></a>

<a id="canonical-1200313303000333-1003321012130030-0212120030231200-1311001301121010-3032230020222301-0310230301200223-1202012322002332-3100130032113313"></a>

#### `ingress_rules.ip_prefix_set.ref.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-2311222031120300-0001010213021223-2201223331123111-1111131233220301-3100102103330231-3120231130133321-1313322110233213-3013021111102111"></a>

<a id="canonical-0230102231022130-0322123313000013-3120031201322320-2001020202221010-3301221201210102-3032133312233310-1213203100300311-0133103232322131"></a>

#### `ingress_rules.ip_prefix_set.ref.tenant` property

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

<a id="canonical-3020301013222100-3011210030332121-0330330323000133-2113222100131311-3100320332013131-3201323111300321-3200021111202202-1121321233133232"></a>

<a id="canonical-1222101231132310-0100131031020301-2103302031130202-3122230033113101-2102200221320233-3202313003000321-2330000020223332-2301333103203331"></a>

#### `ingress_rules.ip_prefix_set.ref.uid` property

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

<a id="canonical-3121200322232301-2312301122303121-0013302311032330-3321102013023003-1032200112031110-1120200012022103-0200020130212312-0212010133320303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ingress_rules.label_matcher` properties

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- ingress_rules.label_matcher

<a id="canonical-3100223201201121-0211122332122102-1023221321133302-2102001012332012-0123302221313103-0230331312323111-0331222331301201-2323211023322233"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-3112113323233021-2110100121200112-3121132020201313-0213130110313113-0122300023013110-3223302122201321-2111013023112311-3033132103131202"></a>

### Direct properties for `ingress_rules.label_matcher`

<a id="canonical-3221313323201032-1200131111020203-1222110133233331-3103223113232323-3122223020032231-3130110012020110-3231231201101120-0301210211113122"></a>

#### `ingress_rules.label_matcher.keys` property

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

<a id="canonical-1223210302033033-1103020331300023-0320212220032002-2111212122331023-2210300232233201-0333321013230302-1203123220231011-0320032220121002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ingress_rules.label_selector` properties

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- ingress_rules.label_selector

<a id="canonical-3310220230013032-1031212232310313-2032213123112313-0210313010332320-3303022132210121-2103301131301122-1132310221023231-1001112133221231"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-1033321331311230-2001333223220110-1311112213102132-3020112112032012-3203300231213110-2110022110221301-3223311002302132-0123223030120012"></a>

### Direct properties for `ingress_rules.label_selector`

<a id="canonical-1331000121133311-2330112300223111-1213230231003321-1212233311030310-0131102331331233-0111002000102322-3213130023122102-3033300223110031"></a>

#### `ingress_rules.label_selector.expressions` property

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

<a id="canonical-3332022032030020-0212311133021021-1200223300231102-1111312312203303-2131211232300013-0010133103122333-0003312313302213-1111322322211300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ingress_rules.metadata` properties

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- ingress_rules.metadata

<a id="canonical-3132201022011010-2122212031333301-0022130123031221-2332202020220310-0022300311313312-2111311212100332-1220202031302201-0303033023302023"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-1213223020203322-3212211200320230-0003313123202002-3321332232032211-2202113322331302-0002212130231312-1020033032102222-0310203200330303"></a>

### Direct properties for `ingress_rules.metadata`

<a id="canonical-1030100303231003-1102031230101102-3202103220133011-3001100213232302-1333131221313032-2110022331311021-0010230133212231-1011111120323031"></a>

#### `ingress_rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-0112230121221113-1113303201103130-1002312212333222-0322110100031312-1211013210120000-3321101223000000-1003301222323220-0310203320211130"></a>

<a id="canonical-0030012330011331-0130210003002031-0232111133001232-1003021030320230-0032302321200012-2233130030331031-0112020100213120-1331210331333310"></a>

#### `ingress_rules.metadata.name` property

Type: `"string"`. Optional.

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

<a id="canonical-0221002002303113-2012333311200103-2301331132330201-0300112022023002-0230212323331232-1233312110022210-1132103103312221-2021113223313113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ingress_rules.outside_endpoints` properties

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- ingress_rules.outside_endpoints

<a id="canonical-1202222230310011-2010012031011101-2232320232332031-0113213233023313-2310200232303121-2032310211030321-0102112203212321-3211120232221101"></a>

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
outside_endpoints = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3331230311130320-0310212101222212-3000331010213321-1312123232010030-3021022221101330-1232323112202130-2300300013310113-0233211332110102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ingress_rules.prefix_list` properties

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

<a id="canonical-0231203013101320-1333320300311303-1220231123032002-2133123133331032-2011011111100011-2211023312003321-0011200320012202-2130113233331102"></a>

### Direct properties for `ingress_rules.prefix_list`

<a id="canonical-2111323233101123-2003222212331022-3233132322113112-1001202010201213-1133320232021100-0301231211203020-2211311112133031-0133202103331030"></a>

#### `ingress_rules.prefix_list.prefixes` property

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

<a id="canonical-3032320123112320-2022212300213331-0311300303203133-1011331303231311-1232001110003000-0213223122212023-3313122032332132-1210313132220133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ingress_rules.protocol_port_range` properties

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md#canonical-2003230030032310-3021013203131210-2022130301302300-2023210312333232-2131222013010200-0110110323220020-3112202330121123-3113322100310103)
- [Property reference](resources--network_policy_view--reference--group-001.md#canonical-0310331320313010-1032201302011331-3021123122231102-1033332111222021-2222212232111023-1020003123033331-3001213313122020-2213000133033101)
- [ingress_rules](resources--network_policy_view--reference--group-001.md#canonical-0331022301123000-1032302322101100-3013101303031112-2200031300003121-1132031121312113-1223333333232021-0013331232010332-0310310031000013)
- ingress_rules.protocol_port_range

<a id="canonical-0021001012103200-3002223020013310-3313022221231321-3301312123110300-3011332031231120-3201100100120121-0001300133323202-2201332021301013"></a>

Type: `"object"`. single nested block, Optional.

Protocol and Port. Protocol and Port ranges.

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

<a id="canonical-3313232110113122-3100021212013120-2113310331211232-1233310220211202-1033113330101303-0032113103011230-1102131123201133-0312312312332123"></a>

### Direct properties for `ingress_rules.protocol_port_range`

<a id="canonical-1222210113200213-2313203221320222-3122301202220021-2021010322021332-1032322103322320-3202202203232020-3203320332321303-1330020313133311"></a>

#### `ingress_rules.protocol_port_range.port_ranges` property

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

<a id="canonical-0220311213122021-1003322203232123-3303323012312330-2323033320333202-2212132300110300-0020200031031233-0300222231010300-3112112331333232"></a>

#### `ingress_rules.protocol_port_range.protocol` property

Type: `"string"`. Optional.

\[Enum: ALL|TCP|UDP|ICMP\] Protocol in IP packet to be used as match criteria Values are TCP, UDP,
and icmp. Possible values are \`ALL\`, \`TCP\`, \`UDP\`, \`ICMP\`.

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
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"ALL\\\",\\\"TCP\\\",\\\"UDP\\\",\\\"ICMP\\\"]"
  }
}
```

<a id="canonical-2031001202231122-1230130201233212-0211330111232122-2233011300203033-2021113233233022-0313233323303300-1003201003202013-1113233202101200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

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

<a id="canonical-3232311211113322-0200322213311003-2130313011232020-3030331300133121-3120121220113323-1012103032223103-3031011233032030-0120111310112311"></a>

### Direct properties for `timeouts`

<a id="canonical-1220133213331221-1002311301311112-3321221202320001-3211122213010333-1021233013102202-0022301012023302-0233213323001301-1310222123202111"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0030200120130312-1111000011331020-3331113201101300-1213233300211232-3323203002111231-0201333210310113-0213323132031013-3001012021231023"></a>

<a id="canonical-2211203201013202-0022320013003202-0320331320011122-3201113002203202-1212312001020002-2231000220211233-2010002201220030-2121233311122311"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1331021301300312-1303030110322213-3030022322321001-3300111100033112-3333332310001212-3123033220000033-1120233231330031-2011300123121311"></a>

<a id="canonical-3210233220121002-0201203021320133-0001331122120210-3132313331223113-0002100303123231-1101320211120321-3000011221331213-0210322203100302"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1002210332013320-2310212323001213-0300100003132100-2013211231220210-2031233300211232-0320020012003130-0032033223220013-3330303012300011"></a>

<a id="canonical-0030331132330332-0320122002020003-0330030003112102-0203032100313223-2020113002211101-0201232101332020-2300201222020022-1300323210201010"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
