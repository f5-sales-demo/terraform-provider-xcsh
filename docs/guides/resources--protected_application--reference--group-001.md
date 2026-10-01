---
page_title: "xcsh_protected_application reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_application reference."
---

# xcsh_protected_application reference

<a id="canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032012213210212-2333132122310213-2130033231030113-1200133323310101-1201130101130310-1223012100311330-2131001311311221-1333110001201231"></a>

## Property reference — Property reference / 321122212010 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- Property reference

<a id="canonical-2000332030033311-1032103303203203-1332103331121132-3302001310311030-1222220131211212-0131102020230111-1220221313311102-1101013113011013"></a>

## Direct properties — Property reference / 321122212010 / 3

- [adobe_commerce_connector](resources--protected_application--reference--group-001.md#canonical-3332022203110112-1213232021331323-0313203313032202-0011122310122123-2101032101331003-1001330230023123-1032122322110301-1202320102202323): complete subsection reference.

<a id="canonical-2221331122010331-1322231121032213-2032100102021021-3222030202230302-3231112313222000-1312003200113103-2332210333102320-0332100110123011"></a>

<a id="canonical-3213323232313303-0203103230130130-2213013001012303-3000122212221220-3332121112233312-2330111130110220-3323032020131030-1320303223302233"></a>

## annotations property — Property reference / 321122212010 / 4

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

- [big_ip_iapp](resources--protected_application--reference--group-001.md#canonical-0303110002303222-0203110110101122-1011300032311021-0233330201031112-2323013212132131-0031102331100032-2101323231221123-3110022201001020): complete subsection reference.

- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111): complete subsection reference.

- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320): complete subsection reference.

- [custom_connector](resources--protected_application--reference--group-004.md#canonical-0001110101211002-1320031003201022-1000020112012123-0310131001312003-2012011132333013-2320221010020212-3001121100313033-3301333013201232): complete subsection reference.

<a id="canonical-0310030210023000-0303300212022110-0133320002000023-3200110122200232-2310000122021122-0111300133120203-3111023303022023-2330331301302013"></a>

<a id="canonical-0200111321002221-0002020303303103-0112320012121103-3130332120120112-0011132023000211-1032312202023102-3303122330322303-0220110233300320"></a>

## description property — Property reference / 321122212010 / 5

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

<a id="canonical-3331202201311302-3100213223133201-2112211213330220-3311032210030223-0102201102011310-3013113121112300-0323103323300231-0120301311203102"></a>

<a id="canonical-0330223303101010-1000231313121130-0102322100031302-2103200321312101-1110233220012230-1113230232023031-2300233301120000-3320032230313223"></a>

## disable property — Property reference / 321122212010 / 6

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

- [f5_big_ip](resources--protected_application--reference--group-004.md#canonical-0021222233012000-1203110331110030-2201103112320103-0232023023031221-1032102302131331-1113203023033213-3301223331110310-1100003322003121): complete subsection reference.

<a id="canonical-1013101201003121-3122123120202310-3201003322223023-1302232110012110-2123110302320103-0322130212230222-2033111012121102-1311100222120130"></a>

<a id="canonical-0130013103123233-0122300203011301-0231203321323110-2102100222331303-1320131221111321-2301312011301113-3131101022122213-1300301133030202"></a>

## ID property — Property reference / 321122212010 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0300320112320301-3332213301331101-3123220011303031-3230013032330002-3113013021201323-2221012101211323-0013133131010311-3323012221313230"></a>

<a id="canonical-1322321110233101-1033223112310310-0001333203200203-2032130223122300-1320033311010100-1133032122030313-0300302001230110-3303323121333122"></a>

## labels property — Property reference / 321122212010 / 8

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

<a id="canonical-1130101130233101-1130013132003001-2301213010202330-0011101222321321-3310112022012321-0331102331231221-2310121003221102-3322120210310231"></a>

<a id="canonical-3030331301012103-2200301322203333-0122201322112311-3102331122300102-3003310133303203-3122212233212030-2302023312101022-2200312313121302"></a>

## name property — Property reference / 321122212010 / 9

Type: `"string"`. Required.

Name of the Protected Application. Must be unique within the namespace.

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

<a id="canonical-1023330200000131-1120010020231112-1122322221012012-0230320033313213-3023301103133100-2321033301221310-3111022301201012-1301320110213121"></a>

<a id="canonical-2231333133222121-2122322022320023-1033230222113321-3332210030300013-3010120130303131-1212023010302021-3222130031302330-3323102022001021"></a>

## namespace property — Property reference / 321122212010 / 10

Type: `"string"`. Required.

Namespace where the Protected Application is created.

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

<a id="canonical-2101013102123222-3000113033310022-2331013332112101-1300310332110221-1231022001313320-0303210302300313-3220130121113002-1012302033022232"></a>

<a id="canonical-2111133013121100-2313112003233320-1200120313103303-1001023003023303-0120031212221122-3300021230033013-1311213323203100-3100020032313110"></a>

## region property — Property reference / 321122212010 / 11

Type: `"string"`. Optional, Computed.

\[Enum: US|EU|ASIA|CA\] Defines a selection for Bot Defense region - US: US United States of America
&#8203;- EU: EU European Union - ASIA: ASIA Asia - CA: CA Canada. Possible values are \`US\`, \`EU\`,
\`ASIA\`, \`CA\`. Defaults to \`US\`.

Upstream description:

Defines a selection for Bot Defense region

&#8203;- US: US

United States of America &#8203;- EU: EU

European Union &#8203;- ASIA: ASIA

Asia &#8203;- CA: CA

Canada.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("US",
    "EU",
    "ASIA",
    "CA"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "US",
  "enum": [
    "US",
    "EU",
    "ASIA",
    "CA"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [salesforce_commerce_connector](resources--protected_application--reference--group-004.md#canonical-2332033220132333-1032222222002332-0311220332202110-2323303102320312-2200130201223220-1133120130131231-2221031302220301-1011010202321210): complete subsection reference.

- [timeouts](resources--protected_application--reference--group-004.md#canonical-1321113002012002-0332322213123310-0023211222301233-1132002120201303-2223010212022303-1302330323232220-1331130002131110-3012220100132010): complete subsection reference.

<a id="canonical-2212012331332011-3320332223033032-3122201300001333-2330210020231100-1310310100120003-0312232002011200-3320123302220110-1013003030031133"></a>

## All schema paths — Property reference / 321122212010 / 12

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `adobe_commerce_connector` | [adobe_commerce_connector](resources--protected_application--reference--group-001.md#canonical-1010323121332303-2221100033303323-2230301120101230-1322010201212203-1313233311112130-2203323221131212-3223230222001003-1022221333123312) |
| `annotations` | [annotations](resources--protected_application--reference--group-001.md#canonical-2221331122010331-1322231121032213-2032100102021021-3222030202230302-3231112313222000-1312003200113103-2332210333102320-0332100110123011) |
| `big_ip_iapp` | [big_ip_iapp](resources--protected_application--reference--group-001.md#canonical-1333211210302130-2202023111203221-3121230010312031-2133011210203103-1000301000230212-1303023121132113-3210201231031331-0031110003302003) |
| `cloudflare` | [Cloudflare](resources--protected_application--reference--group-001.md#canonical-0033200130032012-3112302102100200-3202331202111311-1301221120333011-1320010020113111-0022021003211313-0112320013122321-1213230012002011) |
| `cloudflare.continue_mitigation_action_hdr` | [cloudflare.continue_mitigation_action_hdr](resources--protected_application--reference--group-001.md#canonical-0232122110333012-2330231310300323-0031100101133002-3111023233022003-3130323222000002-2231321223222322-0301000031301112-2002211120110030) |
| `cloudflare.disable_js_insert` | [cloudflare.disable_js_insert](resources--protected_application--reference--group-001.md#canonical-0130200133201310-1311301011120032-0213210313230222-0130312322023003-0321013213211130-0202013310210211-2212101210110210-2101231331220303) |
| `cloudflare.disable_mobile_sdk` | [cloudflare.disable_mobile_sdk](resources--protected_application--reference--group-001.md#canonical-3020303130321110-0230231210302023-1013113131131011-3213100211131133-3021231322120222-3013310123101211-2320303333103221-0211321102020012) |
| `cloudflare.js_insertion_rules` | [cloudflare.js_insertion_rules](resources--protected_application--reference--group-001.md#canonical-0222323213013232-0020031110233020-0222022223002111-3022230120223312-3001203211213231-2031202003000333-3023112232303013-2200232000201032) |
| `cloudflare.js_insertion_rules.exclude_list` | [cloudflare.js_insertion_rules.exclude_list](resources--protected_application--reference--group-001.md#canonical-0330213300030033-1110030201333123-1310023130220130-0011213301322023-2220120123123033-0230103302232031-3032211322202112-1313102320333300) |
| `cloudflare.js_insertion_rules.exclude_list.any_domain` | [cloudflare.js_insertion_rules.exclude_list.any_domain](resources--protected_application--reference--group-001.md#canonical-2031330111111121-1021113020310323-2303123310132330-3010221300300221-0001203012323122-3131102110222131-3211232113222022-1030020120221201) |
| `cloudflare.js_insertion_rules.exclude_list.domain` | [cloudflare.js_insertion_rules.exclude_list.domain](resources--protected_application--reference--group-001.md#canonical-3033003002001321-3032023022031233-2131001001023210-2111323201122213-0112100312323322-2300210213013120-3112011331010112-2120130231230001) |
| `cloudflare.js_insertion_rules.exclude_list.domain.exact_value` | [cloudflare.js_insertion_rules.exclude_list.domain.exact_value](resources--protected_application--reference--group-001.md#canonical-1330123232333122-3031111330313232-1133303123010213-0210301001113332-1011323031110311-3222310321212001-2302131031132321-2322232301021001) |
| `cloudflare.js_insertion_rules.exclude_list.domain.regex_value` | [cloudflare.js_insertion_rules.exclude_list.domain.regex_value](resources--protected_application--reference--group-001.md#canonical-0232003313133323-2231033303211131-0131302333221303-1100301231020030-3321311120211010-3200331132300012-2010231000013233-1301220313023031) |
| `cloudflare.js_insertion_rules.exclude_list.domain.suffix_value` | [cloudflare.js_insertion_rules.exclude_list.domain.suffix_value](resources--protected_application--reference--group-001.md#canonical-2312101303000212-1021131030030012-3233232023332003-3320201102130313-0212132200013233-3012212013122021-3033121032130033-1213300203100212) |
| `cloudflare.js_insertion_rules.exclude_list.metadata` | [cloudflare.js_insertion_rules.exclude_list.metadata](resources--protected_application--reference--group-001.md#canonical-2321032222102220-0303103210112032-2002030320130020-1131123202111002-2123231301223231-1132113103222322-3231023202203200-2113011103300200) |
| `cloudflare.js_insertion_rules.exclude_list.metadata.description_spec` | [cloudflare.js_insertion_rules.exclude_list.metadata.description_spec](resources--protected_application--reference--group-001.md#canonical-1321001102203001-0121122003121332-2322121303031012-3110211312322202-3002101323220221-1201203000301210-2321110332120120-0020331211013123) |
| `cloudflare.js_insertion_rules.exclude_list.metadata.name` | [cloudflare.js_insertion_rules.exclude_list.metadata.name](resources--protected_application--reference--group-001.md#canonical-3232132001112220-0131320302111200-2301021320211022-3300310203213210-0301222323322102-0230023000030330-0212220100002120-0320122330010230) |
| `cloudflare.js_insertion_rules.exclude_list.path` | [cloudflare.js_insertion_rules.exclude_list.path](resources--protected_application--reference--group-001.md#canonical-2132122231021221-0233010130120121-3232212003123312-1013333213103031-2212302101011010-2210111333100323-1012213330230022-3202002300113302) |
| `cloudflare.js_insertion_rules.exclude_list.path.path` | [cloudflare.js_insertion_rules.exclude_list.path.path](resources--protected_application--reference--group-001.md#canonical-2100023312320323-2112322122201022-1330021323032112-0310002001203021-0011033022220210-2300232000001000-0000332331112130-3321102221313321) |
| `cloudflare.js_insertion_rules.exclude_list.path.prefix` | [cloudflare.js_insertion_rules.exclude_list.path.prefix](resources--protected_application--reference--group-001.md#canonical-2000111232200000-1103133303121302-0033320110130323-0322033310032111-2323013032012232-3232001333123232-2312033001033302-1102103111110103) |
| `cloudflare.js_insertion_rules.exclude_list.path.regex` | [cloudflare.js_insertion_rules.exclude_list.path.regex](resources--protected_application--reference--group-001.md#canonical-1020211323210030-2031303011231032-1133223020022030-2213220313221310-1220100312123022-3200131023202101-0301333231023110-3031012100333210) |
| `cloudflare.js_insertion_rules.javascript_location` | [cloudflare.js_insertion_rules.javascript_location](resources--protected_application--reference--group-001.md#canonical-3002103322210303-1323313112022320-0120323103231022-3333012032331203-1222333231010100-1030013331210203-1310102113110211-3210002120300130) |
| `cloudflare.js_insertion_rules.js_download_path` | [cloudflare.js_insertion_rules.js_download_path](resources--protected_application--reference--group-001.md#canonical-1132331121321113-0301111132202321-0110302131322333-3020033000113103-2132212032031331-0232313203213000-0230132112010230-0010203132331222) |
| `cloudflare.js_insertion_rules.rules` | [cloudflare.js_insertion_rules.rules](resources--protected_application--reference--group-001.md#canonical-0213231112230230-2001130113200021-2321313322200002-3003332300112130-3011323003230121-1312023203011201-2122321332032321-3330322333200231) |
| `cloudflare.js_insertion_rules.rules.any_domain` | [cloudflare.js_insertion_rules.rules.any_domain](resources--protected_application--reference--group-001.md#canonical-3000113222132100-3121130210002030-2012101223131123-2003331220012313-0203200322312020-0323302322233030-0300031223202313-1231302231203112) |
| `cloudflare.js_insertion_rules.rules.domain` | [cloudflare.js_insertion_rules.rules.domain](resources--protected_application--reference--group-001.md#canonical-1312020132101210-1200331102023112-0012222131010033-3320120112103010-2103122232003330-1212020330220102-2123233233211101-3232113223220001) |
| `cloudflare.js_insertion_rules.rules.domain.exact_value` | [cloudflare.js_insertion_rules.rules.domain.exact_value](resources--protected_application--reference--group-001.md#canonical-1311113103220100-1232310110012013-2321323121011310-0021303130023012-3020031212133333-0030101030101031-0121012001223021-1312331212113301) |
| `cloudflare.js_insertion_rules.rules.domain.regex_value` | [cloudflare.js_insertion_rules.rules.domain.regex_value](resources--protected_application--reference--group-001.md#canonical-3203332331033212-0000320103223322-2013232301210132-2201300103333023-0211112321120301-0231221223132111-1210322200322231-0203302312121330) |
| `cloudflare.js_insertion_rules.rules.domain.suffix_value` | [cloudflare.js_insertion_rules.rules.domain.suffix_value](resources--protected_application--reference--group-001.md#canonical-2131122113001331-1100220120212201-0003022230021123-2311302001322111-2300212023323110-2300213332331113-3100000202331302-2010200232311331) |
| `cloudflare.js_insertion_rules.rules.exact_path` | [cloudflare.js_insertion_rules.rules.exact_path](resources--protected_application--reference--group-001.md#canonical-0333320100031011-1233001330102232-2303333220333002-1123221210333003-2120133123102020-2130133031210100-3231120311212223-3310021302103233) |
| `cloudflare.js_insertion_rules.rules.glob` | [cloudflare.js_insertion_rules.rules.glob](resources--protected_application--reference--group-001.md#canonical-2300123200201111-2303330233112130-1033013032012223-1230023133311212-2110003010200302-3232123212311013-2003113012021023-0123110011231310) |
| `cloudflare.js_insertion_rules.rules.metadata` | [cloudflare.js_insertion_rules.rules.metadata](resources--protected_application--reference--group-001.md#canonical-2100202210303030-0321111231302333-0103011101332011-3213311233213333-0112013023002032-2121232322021130-2132022010331212-1221221002322301) |
| `cloudflare.js_insertion_rules.rules.metadata.description_spec` | [cloudflare.js_insertion_rules.rules.metadata.description_spec](resources--protected_application--reference--group-001.md#canonical-2302033303312113-1002300101221003-0001003133110012-1103232220200231-2132201033011022-1112212301011111-1210010330000103-2113232320021311) |
| `cloudflare.js_insertion_rules.rules.metadata.name` | [cloudflare.js_insertion_rules.rules.metadata.name](resources--protected_application--reference--group-001.md#canonical-3200212322210020-0300312312203002-1231232010313200-3231113201102023-0111303013220032-2210203010220110-0211020322313302-0111122030310110) |
| `cloudflare.js_insertion_rules.rules.prefix` | [cloudflare.js_insertion_rules.rules.prefix](resources--protected_application--reference--group-001.md#canonical-1031323110301022-0033213333030012-3132132112133132-3111311220121001-0012223301312201-2000002013131110-3211330133332110-3102020022301231) |
| `cloudflare.loglevel` | [cloudflare.loglevel](resources--protected_application--reference--group-001.md#canonical-2012323212312333-1030111102102303-0111331221132201-2121232321313303-2311123320003123-0302230311020310-0231102111033033-0211310020222232) |
| `cloudflare.manual_js_insert` | [cloudflare.manual_js_insert](resources--protected_application--reference--group-001.md#canonical-2210200121122101-3311010033010210-2200302230103311-0033110231131020-0120321133221223-3121010023303103-2332031313022312-0001033200121120) |
| `cloudflare.manual_js_insert.js_download_path` | [cloudflare.manual_js_insert.js_download_path](resources--protected_application--reference--group-001.md#canonical-3100232132100013-1120300013012302-0223123312311223-1033303220011020-3223322223333333-3122312211313333-0223310212013131-1203203231001110) |
| `cloudflare.mobile_sdk_config` | [cloudflare.mobile_sdk_config](resources--protected_application--reference--group-001.md#canonical-0213022302033113-1300012000332200-0231123131211121-1320130133331310-2312110301200210-0112300020222023-3230112310333212-0200000322102223) |
| `cloudflare.mobile_sdk_config.mobile_identifier` | [cloudflare.mobile_sdk_config.mobile_identifier](resources--protected_application--reference--group-001.md#canonical-1021312012012320-1202220212211212-0211020132233221-2201101130021112-1312323302233001-0231200323000220-1031001132101231-0303203312331133) |
| `cloudflare.mobile_sdk_config.mobile_identifier.headers` | [cloudflare.mobile_sdk_config.mobile_identifier.headers](resources--protected_application--reference--group-001.md#canonical-1201231030100212-0233131131010302-0201231221233300-1311223223132012-0212101030030202-1010230221303313-1002012310123232-3023223311223131) |
| `cloudflare.mobile_sdk_config.mobile_identifier.headers.exact` | [cloudflare.mobile_sdk_config.mobile_identifier.headers.exact](resources--protected_application--reference--group-001.md#canonical-3010133000331023-2333120330031212-0220201232301132-0033023232320100-2101303113331231-0012222121322201-1333011102213210-1121101132123321) |
| `cloudflare.mobile_sdk_config.mobile_identifier.headers.name` | [cloudflare.mobile_sdk_config.mobile_identifier.headers.name](resources--protected_application--reference--group-001.md#canonical-3321100210233001-1003301231333331-0312230022111013-1031203230302321-3300113312313111-0011003300022303-0102223122130212-0023001233011313) |
| `cloudflare.mobile_sdk_config.mobile_identifier.headers.regex` | [cloudflare.mobile_sdk_config.mobile_identifier.headers.regex](resources--protected_application--reference--group-001.md#canonical-2113222220113023-2310021120233033-2023020310320113-3121032233102123-0133020001002303-0300300002020201-2202010232200113-2113300102320212) |
| `cloudflare.protected_endpoints` | [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3212033021020010-1201311130333123-0323230221231130-1210101103312231-2101031010300201-2333230031223101-1030211303013203-0120330110233201) |
| `cloudflare.protected_endpoints.any_domain` | [cloudflare.protected_endpoints.any_domain](resources--protected_application--reference--group-002.md#canonical-0113102131232031-1312323012002322-1132303201201221-3112033330031330-2322123200120211-3010223033010213-2202201103221202-0300322302303120) |
| `cloudflare.protected_endpoints.domain` | [cloudflare.protected_endpoints.domain](resources--protected_application--reference--group-002.md#canonical-1333302022232233-2120300310110333-2220012103022002-2010100130123120-1323012012121011-0002302021300122-0223033233300002-2210302213130322) |
| `cloudflare.protected_endpoints.domain.exact_value` | [cloudflare.protected_endpoints.domain.exact_value](resources--protected_application--reference--group-002.md#canonical-3101013123133023-0012221223200320-0201132031322323-2123211232320231-3110303232233201-2132030232231303-2300033200030011-0000203012112223) |
| `cloudflare.protected_endpoints.domain.regex_value` | [cloudflare.protected_endpoints.domain.regex_value](resources--protected_application--reference--group-002.md#canonical-3211203212221110-1033223001113201-1312122322021113-2103102030012110-1103130323232123-3211311230102313-1321113302001013-3101133021221013) |
| `cloudflare.protected_endpoints.domain.suffix_value` | [cloudflare.protected_endpoints.domain.suffix_value](resources--protected_application--reference--group-002.md#canonical-0330230120020223-0220212123012221-2333301210133030-3112330000330020-1101312101300311-2303201222213032-2230033312101023-1023123130231212) |
| `cloudflare.protected_endpoints.http_methods` | [cloudflare.protected_endpoints.http_methods](resources--protected_application--reference--group-001.md#canonical-2030333122003212-2320213012110221-0000211203030033-2010333203201321-1221013033311311-1113331000113300-3110200320331133-1032223033121312) |
| `cloudflare.protected_endpoints.metadata` | [cloudflare.protected_endpoints.metadata](resources--protected_application--reference--group-002.md#canonical-0012232122120332-0212123200120100-3203013212301131-1331123202113111-0303102111122322-1220121123011023-3033001020132013-0202020101220212) |
| `cloudflare.protected_endpoints.metadata.description_spec` | [cloudflare.protected_endpoints.metadata.description_spec](resources--protected_application--reference--group-002.md#canonical-3221122320311101-3102002030133112-3203201021321111-2022133332000202-3330213320012312-2120022311120320-2230320003003331-1110033312203132) |
| `cloudflare.protected_endpoints.metadata.name` | [cloudflare.protected_endpoints.metadata.name](resources--protected_application--reference--group-002.md#canonical-2030311322033323-0021123213031010-1021020331002310-3001030102130330-1001003230032023-2021311000322220-0203121303211201-2200220221113023) |
| `cloudflare.protected_endpoints.mobile_client` | [cloudflare.protected_endpoints.mobile_client](resources--protected_application--reference--group-002.md#canonical-2133232322011301-2100323220222002-3110231121020313-1033120023012213-1122032001200220-0122213120003103-3223303323201123-3332133221231032) |
| `cloudflare.protected_endpoints.mobile_client.block` | [cloudflare.protected_endpoints.mobile_client.block](resources--protected_application--reference--group-002.md#canonical-3313112003013223-3123220113312120-0222223320111203-0311211123120220-3032123231230233-1010213222120213-3101112232233223-0200202012131201) |
| `cloudflare.protected_endpoints.mobile_client.block.body` | [cloudflare.protected_endpoints.mobile_client.block.body](resources--protected_application--reference--group-002.md#canonical-3132003201111223-3100312001230332-3001310012210021-3233101110033021-0220323100011111-1001333131121032-3131223320211101-1133202132303032) |
| `cloudflare.protected_endpoints.mobile_client.block.content_type` | [cloudflare.protected_endpoints.mobile_client.block.content_type](resources--protected_application--reference--group-002.md#canonical-0333303101323332-3230313221321033-2102011022011123-3211330320230323-3110310202322030-2231100112030131-0002120031203110-1133112303333213) |
| `cloudflare.protected_endpoints.mobile_client.block.status` | [cloudflare.protected_endpoints.mobile_client.block.status](resources--protected_application--reference--group-002.md#canonical-1213312213002123-0001030201320001-0032122233220021-2230020030020001-2022112201321110-2012220331332131-2303101223021100-1003213111013230) |
| `cloudflare.protected_endpoints.mobile_client.continue` | [cloudflare.protected_endpoints.mobile_client.continue](resources--protected_application--reference--group-002.md#canonical-1112321121120010-3300310011221011-2002002331301012-3101101332033303-3030301112310132-2030003032113333-0021100020032323-2320013312333332) |
| `cloudflare.protected_endpoints.mobile_client.continue.add_header` | [cloudflare.protected_endpoints.mobile_client.continue.add_header](resources--protected_application--reference--group-002.md#canonical-3111123130103230-3300331220322230-2001231120311310-3013132112223312-3311103111123023-3012230110321233-1332000000012110-2230013333212331) |
| `cloudflare.protected_endpoints.mobile_client.continue.no_header` | [cloudflare.protected_endpoints.mobile_client.continue.no_header](resources--protected_application--reference--group-002.md#canonical-1330223321223301-3333020230312200-1110313331002033-1113033303130100-2220030013133031-2000010323313013-2223311110103110-3233230000210011) |
| `cloudflare.protected_endpoints.path` | [cloudflare.protected_endpoints.path](resources--protected_application--reference--group-002.md#canonical-0333011131231110-0012122202023320-2001223310020313-3303022331330010-2202301330113223-3210000332102332-0211010130002132-1231332222231331) |
| `cloudflare.protected_endpoints.path.caseinsensitive` | [cloudflare.protected_endpoints.path.caseinsensitive](resources--protected_application--reference--group-002.md#canonical-0333032133120321-2231201221112312-1001122231233123-2212022222233311-3321030130330130-0011100311301320-3311203021310221-0000111100130200) |
| `cloudflare.protected_endpoints.path.path` | [cloudflare.protected_endpoints.path.path](resources--protected_application--reference--group-002.md#canonical-1031231002332013-0103102322020032-2233331220213111-3323103000020321-1330320110130011-2012003031221313-2102220222323013-2121132200303001) |
| `cloudflare.protected_endpoints.query` | [cloudflare.protected_endpoints.query](resources--protected_application--reference--group-001.md#canonical-3033211113312003-0001210311213230-2200020200322332-3222012331003121-3000201230112203-1222030102113213-1132313012322123-0111001313321323) |
| `cloudflare.protected_endpoints.web_client` | [cloudflare.protected_endpoints.web_client](resources--protected_application--reference--group-002.md#canonical-3121310223103111-0301010320012013-3120110030132220-3202231230330102-1023232201201213-0320310200133333-1122201101321010-0333232030022222) |
| `cloudflare.protected_endpoints.web_client.block` | [cloudflare.protected_endpoints.web_client.block](resources--protected_application--reference--group-002.md#canonical-3113230031312132-1013333030110310-1103201101313101-2010020112300212-3221033123032102-0033012120323122-2132012332211031-1132012021211023) |
| `cloudflare.protected_endpoints.web_client.block.body` | [cloudflare.protected_endpoints.web_client.block.body](resources--protected_application--reference--group-002.md#canonical-1000303100301120-1313211001111013-3203111203313013-2010223223103301-0300221221100200-3013101311313000-1030032320321303-3200212200122001) |
| `cloudflare.protected_endpoints.web_client.block.content_type` | [cloudflare.protected_endpoints.web_client.block.content_type](resources--protected_application--reference--group-002.md#canonical-1033011230312201-2203003033100320-3223111330102330-1201133210210331-1033202101120213-3213323121201311-2100221113300310-1201122212320333) |
| `cloudflare.protected_endpoints.web_client.block.status` | [cloudflare.protected_endpoints.web_client.block.status](resources--protected_application--reference--group-002.md#canonical-2211320110302003-3110131210202212-3111200323130010-1201030132030103-0220012122300213-3022212210331323-1303223030333102-3323030012231101) |
| `cloudflare.protected_endpoints.web_client.continue` | [cloudflare.protected_endpoints.web_client.continue](resources--protected_application--reference--group-002.md#canonical-0103301000301232-2102202133001231-3101010102310102-0132321123013312-2233113022313221-2322021020021312-2320120323301222-3010203201332102) |
| `cloudflare.protected_endpoints.web_client.continue.add_header` | [cloudflare.protected_endpoints.web_client.continue.add_header](resources--protected_application--reference--group-002.md#canonical-2321303303011121-3102333113320202-0102303311323302-1220232333322011-2222131000233023-3323030232210031-2230332230102130-0000021200033300) |
| `cloudflare.protected_endpoints.web_client.continue.no_header` | [cloudflare.protected_endpoints.web_client.continue.no_header](resources--protected_application--reference--group-002.md#canonical-0311232032311101-3201331022132222-0311230333032322-0300332202203301-2201021130112012-3312200211101303-3123032023133000-1231202233203030) |
| `cloudflare.protected_endpoints.web_client.redirect` | [cloudflare.protected_endpoints.web_client.redirect](resources--protected_application--reference--group-002.md#canonical-2330331123021232-2310211101130032-3220122333323223-1212121020101113-3120103311003003-2322300032133210-2011013003012000-2331103030003211) |
| `cloudflare.protected_endpoints.web_client.redirect.location` | [cloudflare.protected_endpoints.web_client.redirect.location](resources--protected_application--reference--group-002.md#canonical-3303322023003220-0220321223113110-2132311020120323-3102233031222011-2022222132023003-1130120000030212-2200011212301301-1302323303310203) |
| `cloudflare.protected_endpoints.web_client.redirect.status` | [cloudflare.protected_endpoints.web_client.redirect.status](resources--protected_application--reference--group-002.md#canonical-3200300102212301-2232031211301231-2210311321000113-1133101333331113-3223322220001323-0332001323333222-3313001313000320-2232123133132102) |
| `cloudflare.protected_endpoints.web_mobile_client` | [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-2123131020201201-3112310030020222-1000120231311011-2121330012333300-0322020211303031-3333221011223102-3030032332030222-0120232031011120) |
| `cloudflare.protected_endpoints.web_mobile_client.block_mobile` | [cloudflare.protected_endpoints.web_mobile_client.block_mobile](resources--protected_application--reference--group-002.md#canonical-0133011313331013-1123020310131212-1220023213112000-2210111021201323-1011301000302311-0112002121211201-1331110331131223-3232113110202010) |
| `cloudflare.protected_endpoints.web_mobile_client.block_mobile.body` | [cloudflare.protected_endpoints.web_mobile_client.block_mobile.body](resources--protected_application--reference--group-002.md#canonical-0303122013133012-1311200103102331-1033010130233000-1033133110301332-3002110223131000-3001310232330030-0203323300323320-1330213020332333) |
| `cloudflare.protected_endpoints.web_mobile_client.block_mobile.content_type` | [cloudflare.protected_endpoints.web_mobile_client.block_mobile.content_type](resources--protected_application--reference--group-002.md#canonical-0003011101233020-1102323323011113-0211200332231122-3333330221323123-1101333130213100-2103323322113101-0103021003113131-2031213203313112) |
| `cloudflare.protected_endpoints.web_mobile_client.block_mobile.status` | [cloudflare.protected_endpoints.web_mobile_client.block_mobile.status](resources--protected_application--reference--group-002.md#canonical-1111102302303232-0302333100310210-0102323001312231-3031110122222102-0312230133002222-0103321311201202-2022122321031121-3321220222120131) |
| `cloudflare.protected_endpoints.web_mobile_client.block_web` | [cloudflare.protected_endpoints.web_mobile_client.block_web](resources--protected_application--reference--group-002.md#canonical-2121130001221211-0013111003301112-2133210303201331-1133310022221301-1000321321310303-2311123121010010-2210200222132000-3030230201230212) |
| `cloudflare.protected_endpoints.web_mobile_client.block_web.body` | [cloudflare.protected_endpoints.web_mobile_client.block_web.body](resources--protected_application--reference--group-002.md#canonical-2302110221123123-3332211302310103-0121133000032001-2021002300103130-2002032303031301-1011021112301100-1111303211110012-2303302031001000) |
| `cloudflare.protected_endpoints.web_mobile_client.block_web.content_type` | [cloudflare.protected_endpoints.web_mobile_client.block_web.content_type](resources--protected_application--reference--group-002.md#canonical-0230230222313103-3122022003233301-0100003013301013-1013112130311330-2123020110012301-1332123333202032-2133230201212221-2013231302223121) |
| `cloudflare.protected_endpoints.web_mobile_client.block_web.status` | [cloudflare.protected_endpoints.web_mobile_client.block_web.status](resources--protected_application--reference--group-002.md#canonical-3123331322020200-0030320132330200-3200132200233223-0130333321201231-0323233332223322-3013131003120010-3001232120301002-2303220332333203) |
| `cloudflare.protected_endpoints.web_mobile_client.continue_mobile` | [cloudflare.protected_endpoints.web_mobile_client.continue_mobile](resources--protected_application--reference--group-002.md#canonical-0020312020023110-0332102220321200-3230300020103030-1132122301212201-0233213121010131-3203022302000111-0123111032103312-2102120101012013) |
| `cloudflare.protected_endpoints.web_mobile_client.continue_mobile.add_header` | [cloudflare.protected_endpoints.web_mobile_client.continue_mobile.add_header](resources--protected_application--reference--group-002.md#canonical-1202022210201011-0123213303012133-0021013221331133-1130233300313303-3012203120122012-1022313201231031-3133230211232202-2211212333113303) |
| `cloudflare.protected_endpoints.web_mobile_client.continue_mobile.no_header` | [cloudflare.protected_endpoints.web_mobile_client.continue_mobile.no_header](resources--protected_application--reference--group-002.md#canonical-3303302132232022-1331033131311302-2121311121201222-2030211302222123-3302123132132333-0131003203330201-3020330222201202-0302132133330303) |
| `cloudflare.protected_endpoints.web_mobile_client.continue_web` | [cloudflare.protected_endpoints.web_mobile_client.continue_web](resources--protected_application--reference--group-002.md#canonical-1231113030022102-2000200322232122-2321312123020001-0113033202330212-0330111012330220-0303213222231103-1010020132330221-2102320221012112) |
| `cloudflare.protected_endpoints.web_mobile_client.continue_web.add_header` | [cloudflare.protected_endpoints.web_mobile_client.continue_web.add_header](resources--protected_application--reference--group-002.md#canonical-3112123101221322-0223303102103030-3230202330103013-1233213330203102-1010233232312332-0001122023102003-3021210221121023-3001311310310102) |
| `cloudflare.protected_endpoints.web_mobile_client.continue_web.no_header` | [cloudflare.protected_endpoints.web_mobile_client.continue_web.no_header](resources--protected_application--reference--group-002.md#canonical-0211301332003020-1133323021030330-0010121000200202-2310012110103130-0000210211103002-0331032303232133-3220200013231022-0002333302010130) |
| `cloudflare.protected_endpoints.web_mobile_client.redirect_web` | [cloudflare.protected_endpoints.web_mobile_client.redirect_web](resources--protected_application--reference--group-002.md#canonical-3321121013300203-2221300202100130-2001121130003103-2223110132233113-0030113330210212-3033221230011011-3021233113031113-3001101201201123) |
| `cloudflare.protected_endpoints.web_mobile_client.redirect_web.location` | [cloudflare.protected_endpoints.web_mobile_client.redirect_web.location](resources--protected_application--reference--group-002.md#canonical-1033102313130121-1212111200320312-2103010022303313-0320211001010201-2231122222210312-1130002201302220-1022010303201320-1022331001203312) |
| `cloudflare.protected_endpoints.web_mobile_client.redirect_web.status` | [cloudflare.protected_endpoints.web_mobile_client.redirect_web.status](resources--protected_application--reference--group-002.md#canonical-1231212322331113-2002101310332100-3002103011131103-0233313001232011-0113123001313021-2301311123033110-3200212033223130-2311023300212001) |
| `cloudflare.timeout` | [cloudflare.timeout](resources--protected_application--reference--group-001.md#canonical-2113320001022300-3113131103021220-0133100322322202-3002011010103123-1322301023213313-2023200222010120-2231103320102112-2311131001112103) |
| `cloudflare.trusted_clients` | [cloudflare.trusted_clients](resources--protected_application--reference--group-002.md#canonical-1131333132113202-2301303320003333-2310101333111312-1111122101112132-0233321030333333-2010121213002123-3233022020022211-3220210000021010) |
| `cloudflare.trusted_clients.http_header` | [cloudflare.trusted_clients.http_header](resources--protected_application--reference--group-002.md#canonical-2213020301213230-2223211032033033-0202320213002123-3210032202222300-3202103221203122-3213131331002000-1012201022100103-3200023101103113) |
| `cloudflare.trusted_clients.http_header.headers` | [cloudflare.trusted_clients.http_header.headers](resources--protected_application--reference--group-002.md#canonical-3330231022313213-3223033010203031-0313200320102002-0022023110303003-0302212332312310-2113001220001210-3233023111233123-0002131211310310) |
| `cloudflare.trusted_clients.http_header.headers.exact` | [cloudflare.trusted_clients.http_header.headers.exact](resources--protected_application--reference--group-002.md#canonical-1333300212333232-3301012321211102-1103332103020112-3213011313013220-1223202100120100-1331223101203002-1302302210112010-0022002331123111) |
| `cloudflare.trusted_clients.http_header.headers.name` | [cloudflare.trusted_clients.http_header.headers.name](resources--protected_application--reference--group-002.md#canonical-2321222303030310-2210030221113123-2032101101310213-1003332013310120-0020312120133031-2303222032133101-3331133000332130-0223201100032231) |
| `cloudflare.trusted_clients.http_header.headers.regex` | [cloudflare.trusted_clients.http_header.headers.regex](resources--protected_application--reference--group-002.md#canonical-2230121101211233-0032331121321033-0300113110313301-0121212211202000-1322132012300033-3320013310323303-0022113313233133-3111203311012013) |
| `cloudflare.trusted_clients.ip_prefix` | [cloudflare.trusted_clients.ip_prefix](resources--protected_application--reference--group-002.md#canonical-2201222032231023-1100111113322000-0310001022010001-0102110101131113-1102200003213313-0102301312330023-1222313303121203-2221110231322320) |
| `cloudflare.trusted_clients.metadata` | [cloudflare.trusted_clients.metadata](resources--protected_application--reference--group-002.md#canonical-2201322213223112-0221310303320031-0310103213131102-3310301311330211-3030033213020102-3313111313313131-1033300012313331-2023323332223203) |
| `cloudflare.trusted_clients.metadata.description_spec` | [cloudflare.trusted_clients.metadata.description_spec](resources--protected_application--reference--group-002.md#canonical-0111211300022002-3211121013120131-2301311331223031-2012223102013002-1301321232223013-0030102321103232-3100220330330301-2123320202310321) |
| `cloudflare.trusted_clients.metadata.name` | [cloudflare.trusted_clients.metadata.name](resources--protected_application--reference--group-002.md#canonical-1323133113332330-1010301320001032-3201230023321303-3132322322330310-3332130233322011-1203321023322012-2203311333222201-0010112022030203) |
| `cloudfront` | [cloudfront](resources--protected_application--reference--group-002.md#canonical-0323011203311321-0002123031021111-2112223320220213-1231321321221222-3112000200110112-2331021213200101-2130200001232020-2310023233330330) |
| `cloudfront.aws_configuration_id_selector` | [cloudfront.aws_configuration_id_selector](resources--protected_application--reference--group-002.md#canonical-3203022033232033-3300020132202320-1102030032203330-2003000231210111-3312100132300310-3123020102121330-3011301311010131-0202031002313021) |
| `cloudfront.aws_configuration_id_selector.ids` | [cloudfront.aws_configuration_id_selector.ids](resources--protected_application--reference--group-002.md#canonical-3303221233233132-0320333202103230-3223102230210232-1233010332032202-2100132030001101-1333323132221330-2013021033103023-1332331330212200) |
| `cloudfront.aws_configuration_tag_selector` | [cloudfront.aws_configuration_tag_selector](resources--protected_application--reference--group-002.md#canonical-3111200223233223-1200320210211330-0230003313032323-1021100122121321-3103023312000001-3333002130312022-1220101211330211-3303323330333200) |
| `cloudfront.aws_configuration_tag_selector.tags` | [cloudfront.aws_configuration_tag_selector.tags](resources--protected_application--reference--group-002.md#canonical-3020023302330110-2111212012031200-1033132330322123-0311203023233002-0302330302031012-1002012103021021-2220302301233330-2330202313221023) |
| `cloudfront.continue_mitigation_action_hdr` | [cloudfront.continue_mitigation_action_hdr](resources--protected_application--reference--group-002.md#canonical-3023102100102300-2020110030232123-2220003032101000-0131310021331123-3232113012330302-1231031333101223-0311223313303212-2202020002221333) |
| `cloudfront.data_sample` | [cloudfront.data_sample](resources--protected_application--reference--group-002.md#canonical-2122303003123230-3100332312320120-0100202001332333-3031321110321312-3123012101232213-0112120220231323-2003032103203231-0300310223013010) |
| `cloudfront.disable_aws_configuration` | [cloudfront.disable_aws_configuration](resources--protected_application--reference--group-002.md#canonical-3010332231112002-0212022212200200-0102313303312133-3321122002011231-0313303023203122-3313003310233133-3302302230013120-0123113001011332) |
| `cloudfront.disable_js_insert` | [cloudfront.disable_js_insert](resources--protected_application--reference--group-002.md#canonical-1230102223233022-0213201032200202-2121230130310300-0201101222312303-3200023102330103-3110320331321110-2320112112030201-0211233121022103) |
| `cloudfront.disable_mobile_sdk` | [cloudfront.disable_mobile_sdk](resources--protected_application--reference--group-002.md#canonical-0023020223203211-3222322130010222-1303112203011322-0002030320131221-3231003133223110-1013210213202031-3320203303220120-2110212223330220) |
| `cloudfront.js_insertion_rules` | [cloudfront.js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-3210233031332321-1322302011001003-1221210302212332-2320120111323113-0031000121030112-2102222233000002-3200320002220101-0130303100011223) |
| `cloudfront.js_insertion_rules.exclude_list` | [cloudfront.js_insertion_rules.exclude_list](resources--protected_application--reference--group-002.md#canonical-1101033233220033-1113023000031000-0030302122212213-2022303303330232-2112200022103331-3100233310302212-1303323211233213-0031011213222123) |
| `cloudfront.js_insertion_rules.exclude_list.any_domain` | [cloudfront.js_insertion_rules.exclude_list.any_domain](resources--protected_application--reference--group-002.md#canonical-3302113222033311-2330222021330030-1232232230111230-3132310211021313-0303002021030313-2302010301330002-3302331002201222-1100003210233000) |
| `cloudfront.js_insertion_rules.exclude_list.domain` | [cloudfront.js_insertion_rules.exclude_list.domain](resources--protected_application--reference--group-002.md#canonical-0110023101111330-2100003101201210-3220210301313112-3122012123320131-2203001331203003-0000122002021020-1212112020000021-0032011100222031) |
| `cloudfront.js_insertion_rules.exclude_list.domain.exact_value` | [cloudfront.js_insertion_rules.exclude_list.domain.exact_value](resources--protected_application--reference--group-002.md#canonical-3333212031332210-1233001202012210-1000012210011312-0211203132313301-1030032121202301-0202332231300320-0210222211000213-2130312303020320) |
| `cloudfront.js_insertion_rules.exclude_list.domain.regex_value` | [cloudfront.js_insertion_rules.exclude_list.domain.regex_value](resources--protected_application--reference--group-002.md#canonical-1212312013033113-3312110132332300-0231320133000112-0012322322210203-3020120112230002-3112222233213232-2233232100203321-0331231103333122) |
| `cloudfront.js_insertion_rules.exclude_list.domain.suffix_value` | [cloudfront.js_insertion_rules.exclude_list.domain.suffix_value](resources--protected_application--reference--group-002.md#canonical-3010312020320321-3021213012322230-3031023033233131-2321330012130212-0011120322101223-1232122311121032-1312213211013311-1033232103200021) |
| `cloudfront.js_insertion_rules.exclude_list.metadata` | [cloudfront.js_insertion_rules.exclude_list.metadata](resources--protected_application--reference--group-002.md#canonical-0331131003120202-3321122131300201-2120330321122301-3103100301232002-0000223312033102-3121300220220003-1112331321203033-1100312310231133) |
| `cloudfront.js_insertion_rules.exclude_list.metadata.description_spec` | [cloudfront.js_insertion_rules.exclude_list.metadata.description_spec](resources--protected_application--reference--group-002.md#canonical-2100000131003312-2022111030200000-3111123023030023-2000323110003021-3132300120130122-3032321112010011-0230111020001131-1333320233222210) |
| `cloudfront.js_insertion_rules.exclude_list.metadata.name` | [cloudfront.js_insertion_rules.exclude_list.metadata.name](resources--protected_application--reference--group-002.md#canonical-2231111332221102-1022111101313032-1023111230022112-0233102111121013-3330311302003012-0212321203211303-3132222012130122-2203321011313212) |
| `cloudfront.js_insertion_rules.exclude_list.path` | [cloudfront.js_insertion_rules.exclude_list.path](resources--protected_application--reference--group-002.md#canonical-1212333221130231-2003301110201313-2232333231213322-1023001213302123-1301011201211301-2220213111211321-3222130333003200-3221103010032323) |
| `cloudfront.js_insertion_rules.exclude_list.path.path` | [cloudfront.js_insertion_rules.exclude_list.path.path](resources--protected_application--reference--group-002.md#canonical-1332230000021302-3310310103030120-1023000022321332-3310201023011011-2102301103111110-0033112122000113-3113220300033031-0322111222120020) |
| `cloudfront.js_insertion_rules.exclude_list.path.prefix` | [cloudfront.js_insertion_rules.exclude_list.path.prefix](resources--protected_application--reference--group-002.md#canonical-0312021110333212-0133023012130221-3313130013330000-0223010212101331-3111002101320321-2230203003231321-3322011120020212-0203202302032220) |
| `cloudfront.js_insertion_rules.exclude_list.path.regex` | [cloudfront.js_insertion_rules.exclude_list.path.regex](resources--protected_application--reference--group-002.md#canonical-1330210002002220-3200033311123313-1313330130330212-2020313310103020-1011330330212113-3120023031110102-1102112211311022-0233010132000313) |
| `cloudfront.js_insertion_rules.javascript_location` | [cloudfront.js_insertion_rules.javascript_location](resources--protected_application--reference--group-002.md#canonical-2202113333201211-0202033333110130-3122112132112320-0102033330112232-0120131210132022-3323330202111303-0032112220023001-0230103022020231) |
| `cloudfront.js_insertion_rules.javascript_mode` | [cloudfront.js_insertion_rules.javascript_mode](resources--protected_application--reference--group-002.md#canonical-1130113100230033-3222100303210203-2033033033211020-2231210332003110-1021200220210211-3022213130222300-0011333102002313-3021131033130331) |
| `cloudfront.js_insertion_rules.js_download_path` | [cloudfront.js_insertion_rules.js_download_path](resources--protected_application--reference--group-002.md#canonical-1220121003202302-1221023020103302-3333133331011023-2113030202202010-3120020131211223-3110113030200310-3210233230210213-1321322232010110) |
| `cloudfront.js_insertion_rules.rules` | [cloudfront.js_insertion_rules.rules](resources--protected_application--reference--group-002.md#canonical-2232301211223321-3221213223300310-2233200033002102-0010210130210031-1102112312112332-1322001133222310-3302100211330323-2001211333132331) |
| `cloudfront.js_insertion_rules.rules.any_domain` | [cloudfront.js_insertion_rules.rules.any_domain](resources--protected_application--reference--group-002.md#canonical-2132311030110110-1123210001112230-1120313320130332-3013011220000113-3030012003101131-1121000303311021-2113331331203333-3322333030110332) |
| `cloudfront.js_insertion_rules.rules.domain` | [cloudfront.js_insertion_rules.rules.domain](resources--protected_application--reference--group-002.md#canonical-0013100022002003-0302333331023221-2001211133322011-3033012111230113-0122131231331000-2302133312331300-1210000301330103-0112000111012303) |
| `cloudfront.js_insertion_rules.rules.domain.exact_value` | [cloudfront.js_insertion_rules.rules.domain.exact_value](resources--protected_application--reference--group-002.md#canonical-0313103222213221-2320320302100212-1203013322231112-1203012021312131-1020203233213233-3002120033100003-2130303013103211-1310320113202121) |
| `cloudfront.js_insertion_rules.rules.domain.regex_value` | [cloudfront.js_insertion_rules.rules.domain.regex_value](resources--protected_application--reference--group-002.md#canonical-1200233221202121-0322022332012233-1123223020103322-3223110021030232-1113113123120113-2112030201033010-3310211110332131-3133120321230221) |
| `cloudfront.js_insertion_rules.rules.domain.suffix_value` | [cloudfront.js_insertion_rules.rules.domain.suffix_value](resources--protected_application--reference--group-002.md#canonical-3210320010130001-0220111310231323-3000133310023222-0133103323012132-1122323030322200-1030012132302010-1020122231303030-2323210032031102) |
| `cloudfront.js_insertion_rules.rules.exact_path` | [cloudfront.js_insertion_rules.rules.exact_path](resources--protected_application--reference--group-002.md#canonical-0323311013030222-3112110323330323-0233023332333221-2331300003222313-3013231300321133-2211110231220200-3103011023320312-1000212103101020) |
| `cloudfront.js_insertion_rules.rules.glob` | [cloudfront.js_insertion_rules.rules.glob](resources--protected_application--reference--group-002.md#canonical-2023020032320232-1121333112011120-0031322303211220-1001303331022202-1133300331031122-2330202221130231-2123101233211112-1012213100011300) |
| `cloudfront.js_insertion_rules.rules.metadata` | [cloudfront.js_insertion_rules.rules.metadata](resources--protected_application--reference--group-002.md#canonical-1031010132131112-1022320021232131-3131213322032030-1122331330333012-2122113332003112-2220333131220202-0022122112001111-3303133012331302) |
| `cloudfront.js_insertion_rules.rules.metadata.description_spec` | [cloudfront.js_insertion_rules.rules.metadata.description_spec](resources--protected_application--reference--group-002.md#canonical-1012322111320232-0210033131222102-3113200102233330-3221112010312322-1123233302313332-1200001101020331-2231022303312001-2101130001202201) |
| `cloudfront.js_insertion_rules.rules.metadata.name` | [cloudfront.js_insertion_rules.rules.metadata.name](resources--protected_application--reference--group-002.md#canonical-1101222332210323-2303302130210230-0113003103232131-0330223020030202-3200221223103311-3303020310000332-3212333020211031-0221311202310323) |
| `cloudfront.js_insertion_rules.rules.prefix` | [cloudfront.js_insertion_rules.rules.prefix](resources--protected_application--reference--group-002.md#canonical-3223031203233032-2003012201202212-2232133100013233-3202123130232202-2330313013202000-1013211002131002-2302122321032232-0311121222322110) |
| `cloudfront.loglevel` | [cloudfront.loglevel](resources--protected_application--reference--group-002.md#canonical-0200213222021013-2012212303233100-0013222033230112-2123103232303031-3332032132212033-2213310230003002-0123211310101320-2332102230131121) |
| `cloudfront.manual_js_insert` | [cloudfront.manual_js_insert](resources--protected_application--reference--group-003.md#canonical-3111311000322030-0022100000021020-0232123211332032-3221310303123120-2102331112132202-3131330031332023-2332203032012203-2102001302122230) |
| `cloudfront.manual_js_insert.javascript_mode` | [cloudfront.manual_js_insert.javascript_mode](resources--protected_application--reference--group-003.md#canonical-0030001220222311-1221333222033303-3133111222220222-1111301122022302-3023123312000231-2301133132020010-0211211111333111-0123003022312231) |
| `cloudfront.manual_js_insert.js_download_path` | [cloudfront.manual_js_insert.js_download_path](resources--protected_application--reference--group-003.md#canonical-0011003301222102-1310233120312210-0121013002313230-0023220030200033-1022123020103032-0023121000313121-3021222010133130-3302112032133333) |
| `cloudfront.mobile_sdk_config` | [cloudfront.mobile_sdk_config](resources--protected_application--reference--group-003.md#canonical-0003102230233322-1123130302302312-1310011103113111-1100130222313321-3032223220113130-1200231202311210-1023311301023022-2330013100022023) |
| `cloudfront.mobile_sdk_config.mobile_identifier` | [cloudfront.mobile_sdk_config.mobile_identifier](resources--protected_application--reference--group-003.md#canonical-1023231021012323-3021220201203210-3212312302100032-1202130310111132-0203300131111000-3300021023130102-0313033101021011-0211110232200321) |
| `cloudfront.mobile_sdk_config.mobile_identifier.headers` | [cloudfront.mobile_sdk_config.mobile_identifier.headers](resources--protected_application--reference--group-003.md#canonical-0010301331001010-3302032031132133-1213010000220323-2312312212233301-2330302330313100-0110133110010001-0223222323333030-3211032012112333) |
| `cloudfront.mobile_sdk_config.mobile_identifier.headers.exact` | [cloudfront.mobile_sdk_config.mobile_identifier.headers.exact](resources--protected_application--reference--group-003.md#canonical-1110312221023320-1233012033220103-2021320222013312-3131222112300002-3122013313301022-2120321223222332-3132223023200200-0100113132031102) |
| `cloudfront.mobile_sdk_config.mobile_identifier.headers.name` | [cloudfront.mobile_sdk_config.mobile_identifier.headers.name](resources--protected_application--reference--group-003.md#canonical-1133001011303302-1320233333013220-3130223023133000-2333030031131101-3032233310123311-1232122311001202-0300030223323311-2013021221111312) |
| `cloudfront.mobile_sdk_config.mobile_identifier.headers.regex` | [cloudfront.mobile_sdk_config.mobile_identifier.headers.regex](resources--protected_application--reference--group-003.md#canonical-0202020213311113-0123123032120011-0131032202020323-3310113301300312-1311102312222020-2022031101231223-2230122112101212-3201012030230230) |
| `cloudfront.protected_endpoints` | [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-0130322112222030-3010310233213230-0222002100231013-0112310133121111-0103222031032213-0333230030103221-3213332030100332-0030303213021213) |
| `cloudfront.protected_endpoints.any_domain` | [cloudfront.protected_endpoints.any_domain](resources--protected_application--reference--group-003.md#canonical-0102331320201233-3232020032100022-0321033121230120-0323031320233133-3122103330110110-2233111033112222-2113302030111310-3010101103210113) |
| `cloudfront.protected_endpoints.domain` | [cloudfront.protected_endpoints.domain](resources--protected_application--reference--group-003.md#canonical-1313110312021023-0123030231003101-1313310231010102-0032233321311031-1320300301233022-2311030020313023-1333303301302310-3010010332011201) |
| `cloudfront.protected_endpoints.domain.exact_value` | [cloudfront.protected_endpoints.domain.exact_value](resources--protected_application--reference--group-003.md#canonical-3120122233021112-1131311312011113-3201013020212100-2000310000110132-3132033010000031-0032302122122331-0103333323020323-3001212013012103) |
| `cloudfront.protected_endpoints.domain.regex_value` | [cloudfront.protected_endpoints.domain.regex_value](resources--protected_application--reference--group-003.md#canonical-1203301210030200-0102012033303010-2112213203030121-3110302223131003-2202203011110300-3223031322302200-3331231333202330-0022302132223023) |
| `cloudfront.protected_endpoints.domain.suffix_value` | [cloudfront.protected_endpoints.domain.suffix_value](resources--protected_application--reference--group-003.md#canonical-3203233103200221-0323013331312231-0130021311300333-3301203002322332-3312022001231332-3012000232330302-2211103120313301-1220121300323122) |
| `cloudfront.protected_endpoints.flow_label` | [cloudfront.protected_endpoints.flow_label](resources--protected_application--reference--group-003.md#canonical-2003220302020001-2223030300311220-0103320100333031-0001113022333230-1000133323103202-2303321023221102-2323311330302202-2333120220201021) |
| `cloudfront.protected_endpoints.flow_label.account_management` | [cloudfront.protected_endpoints.flow_label.account_management](resources--protected_application--reference--group-003.md#canonical-2222223332333010-2101003220111022-3331301103100110-0100323032022002-1211131202302023-1011322320032110-0300232020112203-3201231331233113) |
| `cloudfront.protected_endpoints.flow_label.account_management.create` | [cloudfront.protected_endpoints.flow_label.account_management.create](resources--protected_application--reference--group-003.md#canonical-1123023232301222-1213312321112321-0321330131330212-2312032020211233-0212123021011113-0021312031032021-1213010300202102-0021033312221312) |
| `cloudfront.protected_endpoints.flow_label.account_management.password_reset` | [cloudfront.protected_endpoints.flow_label.account_management.password_reset](resources--protected_application--reference--group-003.md#canonical-3113321021003202-0010000333020131-1213112333220030-0131213131112103-2012323220231000-1010231010220220-3223202131113111-0200300100201122) |
| `cloudfront.protected_endpoints.flow_label.authentication` | [cloudfront.protected_endpoints.flow_label.authentication](resources--protected_application--reference--group-003.md#canonical-2120011210000210-1312030131323033-1323203212003010-0111030211221013-0200303123202213-2011010032323302-0010223122220331-3020031320323210) |
| `cloudfront.protected_endpoints.flow_label.authentication.login` | [cloudfront.protected_endpoints.flow_label.authentication.login](resources--protected_application--reference--group-003.md#canonical-2100220030133031-2222020120132102-2303203130021202-0211202303130230-0023131020330202-2331013021222200-2313113322231331-1321200303120202) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.disable_transaction_result` | [cloudfront.protected_endpoints.flow_label.authentication.login.disable_transaction_result](resources--protected_application--reference--group-003.md#canonical-0023303301013131-1120311233013031-2201231101300033-0033123210213203-2103022323110122-3232002223103223-2003331010003112-1233303012210200) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result](resources--protected_application--reference--group-003.md#canonical-1312321011003321-0331102232310132-0213001002210213-3222001021211133-0101312320033223-2113212211103033-3202012121333121-0331211222013331) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions](resources--protected_application--reference--group-003.md#canonical-2322020233033113-1232313311222022-2301130220121332-1320023013012212-1321010130320103-2232322123211213-3332021202003010-0321130300032302) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.name` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.name](resources--protected_application--reference--group-003.md#canonical-2003111202101211-3323130303000020-0321001231332101-3032103303220312-3110310213011012-1120131233021313-0100021323100300-1223222000200033) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.regex_values` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.regex_values](resources--protected_application--reference--group-003.md#canonical-1011020021211122-2023323010222303-1112320303200130-2103202313010222-3033133331200032-0100130013110331-2132020020122130-2233100320012200) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.status` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.failure_conditions.status](resources--protected_application--reference--group-003.md#canonical-1133313212032100-3113020223310312-1200221030112031-3213303310310221-2322000000330033-1131223013203131-0102313001231300-3101122121011202) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions](resources--protected_application--reference--group-003.md#canonical-2303233312312033-3332223123122300-0321130001300301-3222023120031103-0130030330110120-3030332333100003-3101023013330233-0313012231320010) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.name` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.name](resources--protected_application--reference--group-003.md#canonical-3122223223030113-1301132113013330-3032301010310201-3310020213232120-0203231232223121-1003131130002011-2102000313112302-1212320321122332) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.regex_values` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.regex_values](resources--protected_application--reference--group-003.md#canonical-2222222000010333-1012102201231123-1210022332033022-0130113033213011-3321100120132333-2313120231310032-0102320332330131-3111323222121133) |
| `cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.status` | [cloudfront.protected_endpoints.flow_label.authentication.login.transaction_result.success_conditions.status](resources--protected_application--reference--group-003.md#canonical-0323030010303113-2122111223300123-3301213302121013-3002123213120323-3102302033332022-3312212001013332-2003201022102113-0200221310323031) |
| `cloudfront.protected_endpoints.flow_label.authentication.login_mfa` | [cloudfront.protected_endpoints.flow_label.authentication.login_mfa](resources--protected_application--reference--group-003.md#canonical-3003300202212103-1231230001310010-1300203213002131-1023312020311022-3101031300203030-1301333320232322-1333210212302311-3011231321020000) |
| `cloudfront.protected_endpoints.flow_label.authentication.login_partner` | [cloudfront.protected_endpoints.flow_label.authentication.login_partner](resources--protected_application--reference--group-003.md#canonical-1132303132132333-3012301232210213-3232320112121110-1202320312223012-3230300032233001-2021211011230333-1130123211131332-1322122122333033) |
| `cloudfront.protected_endpoints.flow_label.authentication.logout` | [cloudfront.protected_endpoints.flow_label.authentication.logout](resources--protected_application--reference--group-003.md#canonical-3210322333030011-0220230330311302-0311003201123012-2231031002312301-1022313021002022-0301133021100133-2032030300032320-1212200213323031) |
| `cloudfront.protected_endpoints.flow_label.authentication.token_refresh` | [cloudfront.protected_endpoints.flow_label.authentication.token_refresh](resources--protected_application--reference--group-003.md#canonical-0133020132230213-1032230031202221-0233133023130312-2112323011122320-1221133031211230-0121113031103022-2320110121300230-3332212320130031) |
| `cloudfront.protected_endpoints.flow_label.financial_services` | [cloudfront.protected_endpoints.flow_label.financial_services](resources--protected_application--reference--group-003.md#canonical-3311023122210020-0012103130131013-2112303010103103-0223101301311113-0230213211331112-3322023232310120-0131230302031121-2033011023231133) |
| `cloudfront.protected_endpoints.flow_label.financial_services.apply` | [cloudfront.protected_endpoints.flow_label.financial_services.apply](resources--protected_application--reference--group-003.md#canonical-1110203300103111-3200322301300013-3101120200310203-3103111322303212-3310333133301011-0333302103303311-2032120232031001-3030211000303332) |
| `cloudfront.protected_endpoints.flow_label.financial_services.money_transfer` | [cloudfront.protected_endpoints.flow_label.financial_services.money_transfer](resources--protected_application--reference--group-003.md#canonical-0012330332312330-2122100113301031-2222200313101032-2200323331210200-2032303022212221-1211033301033222-1233212210301313-1102223030121200) |
| `cloudfront.protected_endpoints.flow_label.flight` | [cloudfront.protected_endpoints.flow_label.flight](resources--protected_application--reference--group-003.md#canonical-3301321131023313-0302330113013030-3310232110030222-3030313232312020-1320032232323023-1332203032111130-3030023102213320-1202033201333023) |
| `cloudfront.protected_endpoints.flow_label.flight.checkin` | [cloudfront.protected_endpoints.flow_label.flight.checkin](resources--protected_application--reference--group-003.md#canonical-2121032003132110-1132231101023120-1112231211031222-1111032101213012-1113331010201312-2123103221033222-2313130313220023-0303132332000302) |
| `cloudfront.protected_endpoints.flow_label.profile_management` | [cloudfront.protected_endpoints.flow_label.profile_management](resources--protected_application--reference--group-003.md#canonical-3201201231303312-0232023323010330-2030123010331000-3111113012323111-0230302312122221-1112332120211321-2032232201311203-3011220030110030) |
| `cloudfront.protected_endpoints.flow_label.profile_management.create` | [cloudfront.protected_endpoints.flow_label.profile_management.create](resources--protected_application--reference--group-003.md#canonical-0002003313032132-1031130003310220-1121310123223313-3120202003033212-3121300202100133-2201200333311232-1012210312200303-3131212133121220) |
| `cloudfront.protected_endpoints.flow_label.profile_management.update` | [cloudfront.protected_endpoints.flow_label.profile_management.update](resources--protected_application--reference--group-003.md#canonical-1222122222100100-0122031012302131-1110020310222101-2232022232023022-3330233030011322-1033220003003331-0132310101222101-0020211102210130) |
| `cloudfront.protected_endpoints.flow_label.profile_management.view` | [cloudfront.protected_endpoints.flow_label.profile_management.view](resources--protected_application--reference--group-003.md#canonical-1112203213212331-2131201122320122-0212301200031202-0101321021221031-2330010313321102-0302110220131030-1031320220332223-1311121322213133) |
| `cloudfront.protected_endpoints.flow_label.search` | [cloudfront.protected_endpoints.flow_label.search](resources--protected_application--reference--group-003.md#canonical-1033122000013003-1200113000130310-0310300330202030-3203011102111231-3201320111230302-2010023012131133-1100320212203323-2122322022311023) |
| `cloudfront.protected_endpoints.flow_label.search.flight_search` | [cloudfront.protected_endpoints.flow_label.search.flight_search](resources--protected_application--reference--group-003.md#canonical-2032120321100120-0223320231030200-0103022021031132-0322211323312123-3110010111030123-2033133101223323-1323213101222033-2021222203330102) |
| `cloudfront.protected_endpoints.flow_label.search.product_search` | [cloudfront.protected_endpoints.flow_label.search.product_search](resources--protected_application--reference--group-003.md#canonical-3121202110033003-2010031021211002-1303322113120130-2120233010133213-2021020301301201-1013022002212021-3111213030133102-1233120032300121) |
| `cloudfront.protected_endpoints.flow_label.search.reservation_search` | [cloudfront.protected_endpoints.flow_label.search.reservation_search](resources--protected_application--reference--group-003.md#canonical-1320033130302113-0311213122120002-1300101111033202-1103102012013333-0123201300233201-3313112021331231-1200120001210012-2202020113212020) |
| `cloudfront.protected_endpoints.flow_label.search.room_search` | [cloudfront.protected_endpoints.flow_label.search.room_search](resources--protected_application--reference--group-003.md#canonical-2322003113001213-3333313133323000-0300212122032113-1330232313311313-1331030131231222-1021200111102121-1310221130333013-0113222013330020) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards](resources--protected_application--reference--group-003.md#canonical-1332302023322230-2010133223021213-1222003233112000-0203312001221213-3121011231323300-3231331201211230-0331321121300010-3232120003221202) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card](resources--protected_application--reference--group-003.md#canonical-2221032332113013-2121320230101312-2332002130322213-0310233211233210-0210301233012033-0032202120112201-1010123223222231-0112301020221233) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_validation` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_validation](resources--protected_application--reference--group-003.md#canonical-0220132213210103-3023321222211031-0121322331311112-3111132203010113-2313021122132323-3310131020121203-1131212301323331-3003221333313023) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart](resources--protected_application--reference--group-003.md#canonical-1012031230320210-2230212003000030-0113100220102223-0130121203310203-3100112320033323-2303023221112223-1021332020213103-1212111030022331) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_checkout` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_checkout](resources--protected_application--reference--group-003.md#canonical-2202132013210321-0102003121332010-2311211103012313-2313121101223222-2322223131031333-0031120201300233-2310312201201032-3031203113320011) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_choose_seat` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_choose_seat](resources--protected_application--reference--group-003.md#canonical-3103311323333301-0102113113121023-3200330012333230-1000200113110233-1123320001133123-0323113323333012-3220001320113212-1310103311332033) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission](resources--protected_application--reference--group-003.md#canonical-0101100103001210-3023212213331231-2201322332002003-2110030333321010-1112303001102320-0220200330211302-1201033113131320-3300231032030112) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_make_payment` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_make_payment](resources--protected_application--reference--group-003.md#canonical-0023010221111200-1330212031130113-1223223000213223-2301302033222120-3030103103230103-1221022133110233-1333221102130002-2011020102221123) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_order` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_order](resources--protected_application--reference--group-003.md#canonical-0202312022130002-3322132233123022-3010013302333211-2131202101222120-3123333311311130-1122021230303322-2001200122113202-3021011031212120) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry](resources--protected_application--reference--group-003.md#canonical-1033333122321133-0102111123210322-0220002122232303-1321301303223032-1111132110012323-1313303120031032-0312212001133112-1102301301011000) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation](resources--protected_application--reference--group-003.md#canonical-3003231110112122-0001210302023011-3233230030321310-1310123230121101-0310230231133223-0212123021300103-3022301313213123-2111210200311202) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card](resources--protected_application--reference--group-004.md#canonical-3133102122122210-0101313002123223-1013022301001211-2233012133323113-1311211121121213-1022120130130122-2003123321230203-3330301020333322) |
| `cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_update_quantity` | [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_update_quantity](resources--protected_application--reference--group-004.md#canonical-2233130000231102-3323221132302223-0003330000113201-1001122113312201-3333223112123223-2233201213331113-3232330112113110-3222101200220113) |
| `cloudfront.protected_endpoints.http_methods` | [cloudfront.protected_endpoints.http_methods](resources--protected_application--reference--group-003.md#canonical-3312011312103121-1203001110201303-1013212002130311-0102102303021330-1213003211000202-1230230002103232-2122213311313311-0130022312320232) |
| `cloudfront.protected_endpoints.metadata` | [cloudfront.protected_endpoints.metadata](resources--protected_application--reference--group-004.md#canonical-3012231232200303-0311223303321120-0023101320233201-2022011120300223-3323132223203330-1220312200031213-3100101211123210-1233013033021000) |
| `cloudfront.protected_endpoints.metadata.description_spec` | [cloudfront.protected_endpoints.metadata.description_spec](resources--protected_application--reference--group-004.md#canonical-3212133120323133-0211120013303123-3301213120223201-2023103022230211-1120323031323201-1112001123112102-1202301022102012-2103232000203021) |
| `cloudfront.protected_endpoints.metadata.name` | [cloudfront.protected_endpoints.metadata.name](resources--protected_application--reference--group-004.md#canonical-3312121203133002-3002213213212200-0111000301010120-0210033122301310-2333031202233001-1313012212322102-0320200001133003-2013100103232213) |
| `cloudfront.protected_endpoints.mobile_client` | [cloudfront.protected_endpoints.mobile_client](resources--protected_application--reference--group-004.md#canonical-1120230012210022-1101332010003322-2203012120223202-2231130332311211-2321221013211123-2100001103131210-3112221200203301-2212120130131132) |
| `cloudfront.protected_endpoints.mobile_client.block` | [cloudfront.protected_endpoints.mobile_client.block](resources--protected_application--reference--group-004.md#canonical-3130000003000121-3103002213020202-0111211031032023-1203030103320110-2110013201111302-3213122330111102-1213021201101333-2001123211301302) |
| `cloudfront.protected_endpoints.mobile_client.block.body` | [cloudfront.protected_endpoints.mobile_client.block.body](resources--protected_application--reference--group-004.md#canonical-3001330131233100-0132113201033120-1211201223212110-3103122210021302-0102023130310213-1020110103321101-1202110211101112-1021321100110000) |
| `cloudfront.protected_endpoints.mobile_client.block.content_type` | [cloudfront.protected_endpoints.mobile_client.block.content_type](resources--protected_application--reference--group-004.md#canonical-2201113100121332-3030322003310103-0121012133313133-0230111010223233-2031231101310123-3012012133120112-1120231132300121-0122110020223022) |
| `cloudfront.protected_endpoints.mobile_client.block.status` | [cloudfront.protected_endpoints.mobile_client.block.status](resources--protected_application--reference--group-004.md#canonical-2021003113132033-0202210320223302-0221223030201331-3223123313110213-0110120030113212-3233310323331332-0333302030230110-0021110200202210) |
| `cloudfront.protected_endpoints.mobile_client.continue` | [cloudfront.protected_endpoints.mobile_client.continue](resources--protected_application--reference--group-004.md#canonical-0122102112022013-0130323310212113-0102202231330032-1032221122112022-0112013303022132-1320110030330320-1123220121010022-3021322203303202) |
| `cloudfront.protected_endpoints.mobile_client.continue.add_header` | [cloudfront.protected_endpoints.mobile_client.continue.add_header](resources--protected_application--reference--group-004.md#canonical-2120111002010203-3113000010301133-3131122330320103-1220122211211133-1101111332232321-3101123322223123-0311132102030111-0312123013000313) |
| `cloudfront.protected_endpoints.mobile_client.continue.no_header` | [cloudfront.protected_endpoints.mobile_client.continue.no_header](resources--protected_application--reference--group-004.md#canonical-0202022312130123-1312311001100122-0032331210232111-1230123120031122-2133032131000030-3023231300003020-2220302332233223-1302213032120333) |
| `cloudfront.protected_endpoints.path` | [cloudfront.protected_endpoints.path](resources--protected_application--reference--group-003.md#canonical-0212211313212221-2231220220232323-2301013100022303-1232220021231120-3120213233130301-1310120231311131-1320311121012003-3203211131233131) |
| `cloudfront.protected_endpoints.query` | [cloudfront.protected_endpoints.query](resources--protected_application--reference--group-003.md#canonical-1131322110103100-2310231121012331-3202011010121210-1020311122113120-1232031322202123-0120000323301212-0121032121111003-2031001022330102) |
| `cloudfront.protected_endpoints.undefined_flow_label` | [cloudfront.protected_endpoints.undefined_flow_label](resources--protected_application--reference--group-004.md#canonical-3310011313321022-2123212110023030-1322032200322203-2222033202113002-2213312212302010-2323000331211201-1211012031202221-2113212133333203) |
| `cloudfront.protected_endpoints.web_client` | [cloudfront.protected_endpoints.web_client](resources--protected_application--reference--group-004.md#canonical-0031032023002233-3302232310310013-0222010011220311-1212023113200230-1302033000111202-1002010102220002-1100222122002101-3303120002001020) |
| `cloudfront.protected_endpoints.web_client.block` | [cloudfront.protected_endpoints.web_client.block](resources--protected_application--reference--group-004.md#canonical-3021023213323123-0132033222133130-1321231302122121-2333310022201111-2201201010233330-1121311122212320-0020120233220201-1203301013022310) |
| `cloudfront.protected_endpoints.web_client.block.body` | [cloudfront.protected_endpoints.web_client.block.body](resources--protected_application--reference--group-004.md#canonical-0110133000323011-2131012320330121-0031311203312012-1123312312220302-2202223133211012-0230201211230212-3112220003000133-3210312101023031) |
| `cloudfront.protected_endpoints.web_client.block.content_type` | [cloudfront.protected_endpoints.web_client.block.content_type](resources--protected_application--reference--group-004.md#canonical-2232003012322212-0120002333330221-1100232300322200-0303321221331313-1001023301001001-3110102232122302-2332220100113320-2210131100210012) |
| `cloudfront.protected_endpoints.web_client.block.status` | [cloudfront.protected_endpoints.web_client.block.status](resources--protected_application--reference--group-004.md#canonical-2201232320301230-0030030320031211-3002313333003131-3001333132232003-3213003131120330-2301223022201103-3113132111200010-1313113320201312) |
| `cloudfront.protected_endpoints.web_client.continue` | [cloudfront.protected_endpoints.web_client.continue](resources--protected_application--reference--group-004.md#canonical-2311103032300010-2031330231220132-1112311312130213-1230133331021112-1323321313013021-1213233020111033-0103312320230000-2313130312103310) |
| `cloudfront.protected_endpoints.web_client.continue.add_header` | [cloudfront.protected_endpoints.web_client.continue.add_header](resources--protected_application--reference--group-004.md#canonical-2031113032010112-2333221331312012-0202033200102333-1113020032032321-1332210320202300-3210023300011112-2211230133012011-2231131230021131) |
| `cloudfront.protected_endpoints.web_client.continue.no_header` | [cloudfront.protected_endpoints.web_client.continue.no_header](resources--protected_application--reference--group-004.md#canonical-0222220312223223-0311111030301312-0122122221321200-0031212223120300-1022033331232102-0010302233300300-0220012021213313-3002023330130313) |
| `cloudfront.protected_endpoints.web_client.redirect` | [cloudfront.protected_endpoints.web_client.redirect](resources--protected_application--reference--group-004.md#canonical-1131222313131103-3003102112123200-1012313320002303-2232230332223131-1333323122203120-3030233113121211-1103301032312012-1012223202312030) |
| `cloudfront.protected_endpoints.web_client.redirect.location` | [cloudfront.protected_endpoints.web_client.redirect.location](resources--protected_application--reference--group-004.md#canonical-1121321232221011-2112322122121112-2023013301322310-1112231130322201-3121130220220132-0211301022012131-2003201300020003-0033333303001200) |
| `cloudfront.protected_endpoints.web_client.redirect.status` | [cloudfront.protected_endpoints.web_client.redirect.status](resources--protected_application--reference--group-004.md#canonical-1031130333121131-3030223012220033-3223320302300312-1022200310201120-1111100330031200-1322002301113011-0330322321223230-0021003233023230) |
| `cloudfront.protected_endpoints.web_mobile_client` | [cloudfront.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-004.md#canonical-1320322121200332-2220010322220133-2330000130111113-3223210030311130-2230000200013012-0023313021311202-3021113211133321-2021032020311201) |
| `cloudfront.protected_endpoints.web_mobile_client.block_mobile` | [cloudfront.protected_endpoints.web_mobile_client.block_mobile](resources--protected_application--reference--group-004.md#canonical-0311113013311022-0101022003113001-0330301101303313-0222233331002301-0320331322223301-1100102230111213-0011321032220033-3202100220001130) |
| `cloudfront.protected_endpoints.web_mobile_client.block_mobile.body` | [cloudfront.protected_endpoints.web_mobile_client.block_mobile.body](resources--protected_application--reference--group-004.md#canonical-3020200132133013-2210311231203333-1222012122203013-0333111221330121-1011320331202310-3230320020232223-0121131123021022-1130201112002131) |
| `cloudfront.protected_endpoints.web_mobile_client.block_mobile.content_type` | [cloudfront.protected_endpoints.web_mobile_client.block_mobile.content_type](resources--protected_application--reference--group-004.md#canonical-0310030032333022-0002203130202220-2300212110101021-2213030003023300-3310301010101022-0333131320030203-2111231321112323-3103111112132200) |
| `cloudfront.protected_endpoints.web_mobile_client.block_mobile.status` | [cloudfront.protected_endpoints.web_mobile_client.block_mobile.status](resources--protected_application--reference--group-004.md#canonical-2013333210311123-3301300010013300-1032003231333330-2130102313312032-1232001022331013-2302110131100120-2013311223232332-3013203201321133) |
| `cloudfront.protected_endpoints.web_mobile_client.block_web` | [cloudfront.protected_endpoints.web_mobile_client.block_web](resources--protected_application--reference--group-004.md#canonical-2113111232331303-3320311030303130-1231123333221321-2100332300331302-1102112223112311-2002000331302232-0032003102110130-2322332201312213) |
| `cloudfront.protected_endpoints.web_mobile_client.block_web.body` | [cloudfront.protected_endpoints.web_mobile_client.block_web.body](resources--protected_application--reference--group-004.md#canonical-2320310230113023-3312123323221230-2012133321221003-3122202200322212-3202323033331303-1201132131321122-2312121011313313-3110022213233330) |
| `cloudfront.protected_endpoints.web_mobile_client.block_web.content_type` | [cloudfront.protected_endpoints.web_mobile_client.block_web.content_type](resources--protected_application--reference--group-004.md#canonical-1300302100013030-0002130133210312-2220001102311022-2333033003111211-0321031201100332-3030323102210023-1103133111200313-3300310113231013) |
| `cloudfront.protected_endpoints.web_mobile_client.block_web.status` | [cloudfront.protected_endpoints.web_mobile_client.block_web.status](resources--protected_application--reference--group-004.md#canonical-3012130022123030-1233223203113011-1333023130212021-2213312122310311-3231211222302101-3300102011022130-1010222203231320-2001220313322012) |
| `cloudfront.protected_endpoints.web_mobile_client.continue_mobile` | [cloudfront.protected_endpoints.web_mobile_client.continue_mobile](resources--protected_application--reference--group-004.md#canonical-3300230021230011-0302212320200130-0222232311113121-3322111202100210-1203023322230121-2110202310010023-0013301213311300-3320311331312232) |
| `cloudfront.protected_endpoints.web_mobile_client.continue_mobile.add_header` | [cloudfront.protected_endpoints.web_mobile_client.continue_mobile.add_header](resources--protected_application--reference--group-004.md#canonical-3033200303130003-3101111222002321-3313311211213331-1222012113313011-3231233312011221-2012100211131202-1222022022101312-1301302002013013) |
| `cloudfront.protected_endpoints.web_mobile_client.continue_mobile.no_header` | [cloudfront.protected_endpoints.web_mobile_client.continue_mobile.no_header](resources--protected_application--reference--group-004.md#canonical-2300113112202012-1203221210310031-3033323000300123-1232200223200201-2010033130322201-2330002111220121-3132201031223132-3203201232132123) |
| `cloudfront.protected_endpoints.web_mobile_client.continue_web` | [cloudfront.protected_endpoints.web_mobile_client.continue_web](resources--protected_application--reference--group-004.md#canonical-3302312310023011-0301300301131300-0223122223100233-3030030023123213-3102313310111121-3122012231230131-2332220030111122-1021330001322013) |
| `cloudfront.protected_endpoints.web_mobile_client.continue_web.add_header` | [cloudfront.protected_endpoints.web_mobile_client.continue_web.add_header](resources--protected_application--reference--group-004.md#canonical-3130133113001331-0333300001321301-1121223201102111-1002031222030100-3322110232232113-1202323112212211-3113311211201223-2320122103212120) |
| `cloudfront.protected_endpoints.web_mobile_client.continue_web.no_header` | [cloudfront.protected_endpoints.web_mobile_client.continue_web.no_header](resources--protected_application--reference--group-004.md#canonical-1000211310320030-1031331303132132-3332303230101131-3331203000121220-3110200103000300-0210222023033333-3133332310130002-0311302101111010) |
| `cloudfront.protected_endpoints.web_mobile_client.redirect_web` | [cloudfront.protected_endpoints.web_mobile_client.redirect_web](resources--protected_application--reference--group-004.md#canonical-2103232122113011-1110110210322111-3301111133220222-1130212223223132-3010222023201002-1033230122021332-0323022332220022-1202033222032232) |
| `cloudfront.protected_endpoints.web_mobile_client.redirect_web.location` | [cloudfront.protected_endpoints.web_mobile_client.redirect_web.location](resources--protected_application--reference--group-004.md#canonical-1112130100120301-1332223211231300-0322312333230133-1013213100331320-2002302011330031-0323210221032312-3131332313121101-1300333013101312) |
| `cloudfront.protected_endpoints.web_mobile_client.redirect_web.status` | [cloudfront.protected_endpoints.web_mobile_client.redirect_web.status](resources--protected_application--reference--group-004.md#canonical-2003013332213323-0133202023132331-2312012223112330-1201101100123120-1202112202223213-1231320212131001-3002220313233333-0322130133333322) |
| `cloudfront.timeout` | [cloudfront.timeout](resources--protected_application--reference--group-002.md#canonical-1021002221210131-3333001132021221-3031121310230132-3310032330123023-2111231223310301-0111013032033001-3111132201000221-2010022331122030) |
| `cloudfront.trusted_clients` | [cloudfront.trusted_clients](resources--protected_application--reference--group-004.md#canonical-1132120220202011-0133133323013122-1233021232132201-3102003232332331-1300101311013211-2330000312103021-1212031211111201-3111202301302331) |
| `cloudfront.trusted_clients.http_header` | [cloudfront.trusted_clients.http_header](resources--protected_application--reference--group-004.md#canonical-2211122232132033-2311132201232201-0020213100030330-3013302303310022-1022202113031101-0333013121313203-3022023310131233-1031330103310331) |
| `cloudfront.trusted_clients.http_header.headers` | [cloudfront.trusted_clients.http_header.headers](resources--protected_application--reference--group-004.md#canonical-0303133121023320-2312201202331213-3120103132213131-3002022320023123-3202212230032132-2200313220012013-2131230212222022-0221101123102210) |
| `cloudfront.trusted_clients.http_header.headers.exact` | [cloudfront.trusted_clients.http_header.headers.exact](resources--protected_application--reference--group-004.md#canonical-3323120122200202-0222223211020302-0013122130121330-0211121020012020-3001021300200300-0012120123203301-3301222210020202-1221232312201302) |
| `cloudfront.trusted_clients.http_header.headers.name` | [cloudfront.trusted_clients.http_header.headers.name](resources--protected_application--reference--group-004.md#canonical-0333313201131033-0223222221021321-2033110002111021-2103003031320222-2120010223033233-3201322012321232-2321331122130012-3011333222002201) |
| `cloudfront.trusted_clients.http_header.headers.regex` | [cloudfront.trusted_clients.http_header.headers.regex](resources--protected_application--reference--group-004.md#canonical-0122230232231211-2330120202132002-3223003321022231-1222023000100233-0223320131212223-2000101000211023-2210020031103101-0322100131333200) |
| `cloudfront.trusted_clients.ip_prefix` | [cloudfront.trusted_clients.ip_prefix](resources--protected_application--reference--group-004.md#canonical-0223010110330120-0012332222121001-1020131323203032-1203030232222011-1232100310120330-0310110103032012-2113311103101003-0312112222121203) |
| `cloudfront.trusted_clients.metadata` | [cloudfront.trusted_clients.metadata](resources--protected_application--reference--group-004.md#canonical-0332223122233013-2323010302030123-3200302033102213-3030032121030203-0320311320031222-1303333210101021-0313223133033303-2130210313010213) |
| `cloudfront.trusted_clients.metadata.description_spec` | [cloudfront.trusted_clients.metadata.description_spec](resources--protected_application--reference--group-004.md#canonical-2310020012220233-2012112232023231-3202313303000132-2113123233201222-2012210201311120-3230212322011333-2210231301201210-2123302023101002) |
| `cloudfront.trusted_clients.metadata.name` | [cloudfront.trusted_clients.metadata.name](resources--protected_application--reference--group-004.md#canonical-1131221133330032-1030123231013002-0032221331323021-0333320231032120-1330111210022223-3100312130230030-2113320113303002-0000113021022332) |
| `custom_connector` | [custom_connector](resources--protected_application--reference--group-004.md#canonical-1322220100112011-2321330313120110-3203101312002210-3201222022111003-1323332200302303-3320302021303002-2031221011330201-1121131222200210) |
| `description` | [description](resources--protected_application--reference--group-001.md#canonical-0310030210023000-0303300212022110-0133320002000023-3200110122200232-2310000122021122-0111300133120203-3111023303022023-2330331301302013) |
| `disable` | [disable](resources--protected_application--reference--group-001.md#canonical-3331202201311302-3100213223133201-2112211213330220-3311032210030223-0102201102011310-3013113121112300-0323103323300231-0120301311203102) |
| `f5_big_ip` | [f5_big_ip](resources--protected_application--reference--group-004.md#canonical-0020303311210203-3221033103331300-1030200133232210-3211233223033312-3120312133222121-2232232132200220-1303033323233313-1132101023330132) |
| `id` | [ID](resources--protected_application--reference--group-001.md#canonical-1013101201003121-3122123120202310-3201003322223023-1302232110012110-2123110302320103-0322130212230222-2033111012121102-1311100222120130) |
| `labels` | [labels](resources--protected_application--reference--group-001.md#canonical-0300320112320301-3332213301331101-3123220011303031-3230013032330002-3113013021201323-2221012101211323-0013133131010311-3323012221313230) |
| `name` | [name](resources--protected_application--reference--group-001.md#canonical-1130101130233101-1130013132003001-2301213010202330-0011101222321321-3310112022012321-0331102331231221-2310121003221102-3322120210310231) |
| `namespace` | [namespace](resources--protected_application--reference--group-001.md#canonical-1023330200000131-1120010020231112-1122322221012012-0230320033313213-3023301103133100-2321033301221310-3111022301201012-1301320110213121) |
| `region` | [region](resources--protected_application--reference--group-001.md#canonical-2101013102123222-3000113033310022-2331013332112101-1300310332110221-1231022001313320-0303210302300313-3220130121113002-1012302033022232) |
| `salesforce_commerce_connector` | [salesforce_commerce_connector](resources--protected_application--reference--group-004.md#canonical-0331201013122122-1222200333203031-2310102103232332-1302033332101133-1302101323312323-1302312011313211-0223222120021332-2303323100232311) |
| `timeouts` | [timeouts](resources--protected_application--reference--group-004.md#canonical-1300112031210222-3130021112013012-0211220230112202-0223230010320012-0030202320321220-3212230030013312-3231013023131020-2232021310212132) |
| `timeouts.create` | [timeouts.create](resources--protected_application--reference--group-004.md#canonical-2111200033010133-2112320001013000-2021223222113311-3033001213101000-0301001330123300-3301211303100100-2330032130330032-1231011321303211) |
| `timeouts.delete` | [timeouts.delete](resources--protected_application--reference--group-004.md#canonical-0123133032031103-0322320010021023-1203030110001011-2130221122023211-1211213203113222-0202101230210101-3203132122232232-2010212013331200) |
| `timeouts.read` | [timeouts.read](resources--protected_application--reference--group-004.md#canonical-1130010030202001-0100332322323132-1103133111012010-2211211032112003-3321031312112233-1030312021212000-2133213021213200-2030000120112201) |
| `timeouts.update` | [timeouts.update](resources--protected_application--reference--group-004.md#canonical-2100210133232210-3030312222300011-2012012231301122-2111202121202232-1200200321033103-0003202132210022-3130130313332212-2233013221210100) |

<a id="canonical-2132000231200303-1001010333212130-0200033231303303-1232223313311213-0011000320321202-1002333113231013-1120321211333110-3323321322012303"></a>

## Next pages — Property reference / 321122212010 / 13

- [adobe_commerce_connector](resources--protected_application--reference--group-001.md#canonical-3332022203110112-1213232021331323-0313203313032202-0011122310122123-2101032101331003-1001330230023123-1032122322110301-1202320102202323)
- [big_ip_iapp](resources--protected_application--reference--group-001.md#canonical-0303110002303222-0203110110101122-1011300032311021-0233330201031112-2323013212132131-0031102331100032-2101323231221123-3110022201001020)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [custom_connector](resources--protected_application--reference--group-004.md#canonical-0001110101211002-1320031003201022-1000020112012123-0310131001312003-2012011132333013-2320221010020212-3001121100313033-3301333013201232)
- [f5_big_ip](resources--protected_application--reference--group-004.md#canonical-0021222233012000-1203110331110030-2201103112320103-0232023023031221-1032102302131331-1113203023033213-3301223331110310-1100003322003121)
- [salesforce_commerce_connector](resources--protected_application--reference--group-004.md#canonical-2332033220132333-1032222222002332-0311220332202110-2323303102320312-2200130201223220-1133120130131231-2221031302220301-1011010202321210)
- [timeouts](resources--protected_application--reference--group-004.md#canonical-1321113002012002-0332322213123310-0023211222301233-1132002120201303-2223010212022303-1302330323232220-1331130002131110-3012220100132010)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-3332022203110112-1213232021331323-0313203313032202-0011122310122123-2101032101331003-1001330230023123-1032122322110301-1202320102202323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0200223113133330-2100300003103303-2213201000023323-2002331203233121-3031122033100012-0333333132013031-3101231232220003-2230001000103323"></a>

## adobe_commerce_connector — adobe_commerce_connector / 301221031133 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- adobe_commerce_connector

<a id="canonical-1010323121332303-2221100033303323-2230301120101230-1322010201212203-1313233311112130-2203323221131212-3223230222001003-1022221333123312"></a>

Type: `["object", {}]`. Optional.

\[OneOf: adobe\_commerce\_connector, big\_ip\_iapp, Cloudflare, cloudfront, custom\_connector,
f5\_big\_ip, salesforce\_commerce\_connector\] Configuration parameter for adobe commerce connector.

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

- [adobe_commerce_connector](resources--protected_application--reference--group-001.md#canonical-1010323121332303-2221100033303323-2230301120101230-1322010201212203-1313233311112130-2203323221131212-3223230222001003-1022221333123312)
- [big_ip_iapp](resources--protected_application--reference--group-001.md#canonical-1333211210302130-2202023111203221-3121230010312031-2133011210203103-1000301000230212-1303023121132113-3210201231031331-0031110003302003)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-0033200130032012-3112302102100200-3202331202111311-1301221120333011-1320010020113111-0022021003211313-0112320013122321-1213230012002011)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-0323011203311321-0002123031021111-2112223320220213-1231321321221222-3112000200110112-2331021213200101-2130200001232020-2310023233330330)
- [custom_connector](resources--protected_application--reference--group-004.md#canonical-1322220100112011-2321330313120110-3203101312002210-3201222022111003-1323332200302303-3320302021303002-2031221011330201-1121131222200210)
- [f5_big_ip](resources--protected_application--reference--group-004.md#canonical-0020303311210203-3221033103331300-1030200133232210-3211233223033312-3120312133222121-2232232132200220-1303033323233313-1132101023330132)
- [salesforce_commerce_connector](resources--protected_application--reference--group-004.md#canonical-0331201013122122-1222200333203031-2310102103232332-1302033332101133-1302101323312323-1302312011313211-0223222120021332-2303323100232311)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
adobe_commerce_connector = {}
```

<a id="canonical-3211000203222233-3330330302311032-2303302300021131-2011001121330030-3101012300201120-2112133303311323-3222302330122011-2030230122320120"></a>

## Direct properties — adobe_commerce_connector / 301221031133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2210201300202110-1020130300332231-1331211201133211-3122333202110030-0231100312013323-3122323002112111-1112331110303221-3100232323020330"></a>

## Next pages — adobe_commerce_connector / 301221031133 / 4

- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-0303110002303222-0203110110101122-1011300032311021-0233330201031112-2323013212132131-0031102331100032-2101323231221123-3110022201001020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300101233321303-1320321001001112-0101220101321001-2132231000313023-3300300121203320-2021102223212221-0231231302001003-1022022313200022"></a>

## big_ip_iapp — big_ip_iapp / 032303222233 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- big_ip_iapp

<a id="canonical-1333211210302130-2202023111203221-3121230010312031-2133011210203103-1000301000230212-1303023121132113-3210201231031331-0031110003302003"></a>

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
big_ip_iapp = {}
```

<a id="canonical-0301323302103233-0333210312300103-3031211101101023-3300330223000322-3022331021022211-0123301322130311-0022322012123302-0333032113222211"></a>

## Direct properties — big_ip_iapp / 032303222233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2102013131312313-3030221031213113-3203012233322310-0210230212311231-2133301022321013-0331000211112121-2230011131031122-2333012001223021"></a>

## Next pages — big_ip_iapp / 032303222233 / 4

- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0232131212032231-2013321322133022-0211300101033302-2302020321222023-3210233300333103-1332321302332121-1101313310031023-0030101031023230"></a>

## Cloudflare — Cloudflare / 102102012230 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- Cloudflare

<a id="canonical-0033200130032012-3112302102100200-3202331202111311-1301221120333011-1320010020113111-0022021003211313-0112320013122321-1213230012002011"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense policy configuration for Cloudflare.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("protected_endpoints"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "manual_js_insert"),
  validators.ConflictingObjectAttributes("disable_mobile_sdk",
    "mobile_sdk_config"),
  validators.ConflictingObjectAttributes("js_insertion_rules",
    "manual_js_insert")}
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
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insertion_rules\",\"manual_js_insert\"]",
  "x-ves-oneof-field-mobile_sdk_choice": "[\"disable_mobile_sdk\",\"mobile_sdk_config\"]"
}
```

Terraform syntax:

```terraform
cloudflare {
  # Configure direct properties listed below.
}
```

<a id="canonical-2122002032320001-0003313301223323-2221300113123022-2113302203013300-3333310132331311-3112232210003212-1033202231103210-0130120021223210"></a>

## Direct properties — Cloudflare / 102102012230 / 3

<a id="canonical-0232122110333012-2330231310300323-0031100101133002-3111023233022003-3130323222000002-2231321223222322-0301000031301112-2002211120110030"></a>

<a id="canonical-0303031310210313-1222110212330300-0323021023012010-1010011310201212-1020230231200213-1102323001123030-2103000022213211-0212200313021220"></a>

## continue_mitigation_action_hdr property — Cloudflare / 102102012230 / 4

Type: `"string"`. Optional.

Case-insensitive HTTP header name for Continue Mitigation Action when add header selected.

Upstream description:

A case-insensitive HTTP header name for Continue Mitigation Action when add header selected.

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
    "byteLength": {
      "max": 256
    },
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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

- [disable_js_insert](resources--protected_application--reference--group-001.md#canonical-0313222102202233-2103230331333312-1213312221310312-1012033131311232-1220121203023010-3131201013132121-0202203133033130-1100303303321102): complete subsection reference.

- [disable_mobile_sdk](resources--protected_application--reference--group-001.md#canonical-2233013312012222-0333100302122200-1001122123221000-0033222303101213-1133300131303030-3120131000012202-1103200103331031-3010302021222202): complete subsection reference.

- [js_insertion_rules](resources--protected_application--reference--group-001.md#canonical-1322002021231300-3010301333121031-1322110012103112-2020011212310320-2122221103021202-1211333331013121-3202231301310022-3131301323111101): complete subsection reference.

<a id="canonical-2012323212312333-1030111102102303-0111331221132201-2121232321313303-2311123320003123-0302230311020310-0231102111033033-0211310020222232"></a>

<a id="canonical-0320112321100023-0221023002033113-0301001010013322-2012333213122002-2300312010103201-1103210213203122-3021110203022331-2331103020322301"></a>

## loglevel property — Cloudflare / 102102012230 / 5

Type: `"string"`. Optional.

\[Enum: LOG\_UNDEFINED|LOG\_ERROR|LOG\_WARNING|LOG\_INFO|LOG\_DEBUG\] Select the level of logging
desired. Levels are cumulative (e.g. Debug includes Error, Warning, and Informational) -
LOG\_UNDEFINED: Undefined - LOG\_ERROR: Error Log only errors - LOG\_WARNING: Warning Log malicious
requests - LOG\_INFO: Info Log all requests - LOG\_DEBUG: Debug Log debugging data. Possible values
are \`LOG\_UNDEFINED\`, \`LOG\_ERROR\`, \`LOG\_WARNING\`, \`LOG\_INFO\`, \`LOG\_DEBUG\`. Defaults to
\`LOG\_UNDEFINED\`.

Upstream description:

Select the level of logging desired. Levels are cumulative (e.g. Debug includes Error, Warning, and
Informational)

&#8203;- LOG\_UNDEFINED: Undefined

&#8203;- LOG\_ERROR: Error

Log only errors &#8203;- LOG\_WARNING: Warning

Log malicious requests &#8203;- LOG\_INFO: Info

Log all requests &#8203;- LOG\_DEBUG: Debug

Log debugging data.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("LOG_UNDEFINED",
    "LOG_ERROR",
    "LOG_WARNING",
    "LOG_INFO",
    "LOG_DEBUG"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "LOG_UNDEFINED",
  "enum": [
    "LOG_UNDEFINED",
    "LOG_ERROR",
    "LOG_WARNING",
    "LOG_INFO",
    "LOG_DEBUG"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [manual_js_insert](resources--protected_application--reference--group-001.md#canonical-2302310133022000-1230130121013212-3323120233122330-2221122132220001-2122001133122302-2132010313002023-3012313120130203-3123232120220302): complete subsection reference.

- [mobile_sdk_config](resources--protected_application--reference--group-001.md#canonical-0122212303302030-3223122122222023-1210233312031201-3313210103123321-0333301122232233-2302231102220123-2301133202020201-3332033000221131): complete subsection reference.

- [protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200): complete subsection reference.

<a id="canonical-2113320001022300-3113131103021220-0133100322322202-3002011010103123-1322301023213313-2023200222010120-2231103320102112-2311131001112103"></a>

<a id="canonical-1210033332333133-1200102302003230-1011100120223010-0210113021202312-3000211023033201-2031320102201223-1122331123233202-2103223333122321"></a>

## timeout property — Cloudflare / 102102012230 / 6

Type: `"number"`. Optional.

The timeout for the inference check, in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 60000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

- [trusted_clients](resources--protected_application--reference--group-002.md#canonical-2112211221002230-0132313302003203-3023102123212203-2101231013213232-1320231230122010-0132301122202122-2323222023222001-1212210200131203): complete subsection reference.

<a id="canonical-2221013211202011-0223211331022301-1220320313012033-2222331232233031-1031013013000321-2120331131101302-0012311312012220-1133130002010120"></a>

## Next pages — Cloudflare / 102102012230 / 7

- [cloudflare.disable_js_insert](resources--protected_application--reference--group-001.md#canonical-0313222102202233-2103230331333312-1213312221310312-1012033131311232-1220121203023010-3131201013132121-0202203133033130-1100303303321102)
- [cloudflare.disable_mobile_sdk](resources--protected_application--reference--group-001.md#canonical-2233013312012222-0333100302122200-1001122123221000-0033222303101213-1133300131303030-3120131000012202-1103200103331031-3010302021222202)
- [cloudflare.js_insertion_rules](resources--protected_application--reference--group-001.md#canonical-1322002021231300-3010301333121031-1322110012103112-2020011212310320-2122221103021202-1211333331013121-3202231301310022-3131301323111101)
- [cloudflare.manual_js_insert](resources--protected_application--reference--group-001.md#canonical-2302310133022000-1230130121013212-3323120233122330-2221122132220001-2122001133122302-2132010313002023-3012313120130203-3123232120220302)
- [cloudflare.mobile_sdk_config](resources--protected_application--reference--group-001.md#canonical-0122212303302030-3223122122222023-1210233312031201-3313210103123321-0333301122232233-2302231102220123-2301133202020201-3332033000221131)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [cloudflare.trusted_clients](resources--protected_application--reference--group-002.md#canonical-2112211221002230-0132313302003203-3023102123212203-2101231013213232-1320231230122010-0132301122202122-2323222023222001-1212210200131203)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-0313222102202233-2103230331333312-1213312221310312-1012033131311232-1220121203023010-3131201013132121-0202203133033130-1100303303321102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2231023101310210-3333303030020331-3123202231031133-2001023122322310-1020333213001201-2303032132211122-0013100102332220-3222222111100320"></a>

## Cloudflare.disable_js_insert — disable_js_insert / 323322220202 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- Cloudflare.disable_js_insert

<a id="canonical-0130200133201310-1311301011120032-0213210313230222-0130312322023003-0321013213211130-0202013310210211-2212101210110210-2101231331220303"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable js insert.

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
disable_js_insert = {}
```

<a id="canonical-1223032213100001-3132201111200013-1002003023003310-3011020301312113-0302202111113032-2113100322103130-0131123021200032-1110121310311321"></a>

## Direct properties — disable_js_insert / 323322220202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3210133030211212-3101010112311311-1113132313311333-1020233322212333-0113203223303200-2112202212212312-3103010332010112-2331121313130333"></a>

## Next pages — disable_js_insert / 323322220202 / 4

- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-2233013312012222-0333100302122200-1001122123221000-0033222303101213-1133300131303030-3120131000012202-1103200103331031-3010302021222202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3300033313232113-2230030300101310-0322320202111221-3321321313313002-2031322333111323-0222003133113130-3100030023111230-1313221212230100"></a>

## Cloudflare.disable_mobile_sdk — disable_mobile_sdk / 000320200011 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- Cloudflare.disable_mobile_sdk

<a id="canonical-3020303130321110-0230231210302023-1013113131131011-3213100211131133-3021231322120222-3013310123101211-2320303333103221-0211321102020012"></a>

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
disable_mobile_sdk = {}
```

<a id="canonical-2210112210101031-2101312322300213-2133212130330313-2020232332333022-2132133123030010-3230133231012213-2333103123310132-2102131231323111"></a>

## Direct properties — disable_mobile_sdk / 000320200011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1302220213101321-1113013312202032-2113131102101002-1211101113123212-0123210303013030-0212221322312033-0311202022322200-0200021102010121"></a>

## Next pages — disable_mobile_sdk / 000320200011 / 4

- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-1322002021231300-3010301333121031-1322110012103112-2020011212310320-2122221103021202-1211333331013121-3202231301310022-3131301323111101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201031222133202-0220122323020102-2021202001300232-2332110312210033-0321003322303122-2011220120233112-2222123200103222-1212113030301201"></a>

## Cloudflare.js_insertion_rules — js_insertion_rules / 331021001323 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- Cloudflare.js_insertion_rules

<a id="canonical-0222323213013232-0020031110233020-0222022223002111-3022230120223312-3001203211213231-2031202003000333-3023112232303013-2200232000201032"></a>

Type: `"object"`. single nested block, Optional.

Defines custom JavaScript insertion rules for Bot Defense Policy.

Upstream description:

This defines custom JavaScript insertion rules for Bot Defense Policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rules")}
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
js_insertion_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-0030133031302202-3330003032102103-3000231121112112-0110031231112300-1213203123210331-1111303233230322-1101132033112200-2200132120002123"></a>

## Direct properties — js_insertion_rules / 331021001323 / 3

- [exclude_list](resources--protected_application--reference--group-001.md#canonical-1222023303122201-2323122330313330-0233010223032311-1023030333132220-1001303031220002-3322201001102320-2030023220001320-2103232200313302): complete subsection reference.

<a id="canonical-3002103322210303-1323313112022320-0120323103231022-3333012032331203-1222333231010100-1030013331210203-1310102113110211-3210002120300130"></a>

<a id="canonical-0130303331223030-2032102323332321-2220003121221133-0122103000132011-1113332133302021-0010202332003010-1030102232213221-3233021133323012"></a>

## javascript_location property — js_insertion_rules / 331021001323 / 4

Type: `"string"`. Optional.

\[Enum: JAVA\_SCRIPT\_LOCATION\_UNDEFINED|AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside
networks. - JAVA\_SCRIPT\_LOCATION\_UNDEFINED: JAVA\_SCRIPT\_LOCATION\_UNDEFINED Undefined Insert
JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript
before first tag. Possible values are \`JAVA\_SCRIPT\_LOCATION\_UNDEFINED\`, \`AFTER\_HEAD\`,
\`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to \`JAVA\_SCRIPT\_LOCATION\_UNDEFINED\`.

Upstream description:

All inside networks.

&#8203;- JAVA\_SCRIPT\_LOCATION\_UNDEFINED: JAVA\_SCRIPT\_LOCATION\_UNDEFINED

Undefined Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag.
Insert JavaScript before first tag.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("JAVA_SCRIPT_LOCATION_UNDEFINED",
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "JAVA_SCRIPT_LOCATION_UNDEFINED",
  "enum": [
    "JAVA_SCRIPT_LOCATION_UNDEFINED",
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1132331121321113-0301111132202321-0110302131322333-3020033000113103-2132212032031331-0232313203213000-0230132112010230-0010203132331222"></a>

<a id="canonical-0212311321331301-1121133221211230-2132121211201330-1103011112002130-1301303002021313-1020220011302202-1001132322200133-3111330201011101"></a>

## js_download_path property — js_insertion_rules / 331021001323 / 5

Type: `"string"`. Optional.

Web client will fetch F5 Client JavaScript from this path. This path must not conflict with any
other website/application paths. If not specified, default to ‘/common.js’.

Upstream description:

Web client will fetch F5 Client JavaScript from this path. This path must not conflict with any
other website/application paths.

If not specified, default to ‘/common.js’.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  }
}
```

- [rules](resources--protected_application--reference--group-001.md#canonical-2330032030211220-2321100333312222-2303120210101121-0201212120201330-1023331123331022-2322300220313133-1220302231122021-1001302311011021): complete subsection reference.

<a id="canonical-1133312313110303-0211032001223110-0113213011310110-1113031310020223-0131231021100132-2021333122312233-1201022230101003-2332001221302120"></a>

## Next pages — js_insertion_rules / 331021001323 / 6

- [cloudflare.js_insertion_rules.exclude_list](resources--protected_application--reference--group-001.md#canonical-1222023303122201-2323122330313330-0233010223032311-1023030333132220-1001303031220002-3322201001102320-2030023220001320-2103232200313302)
- [cloudflare.js_insertion_rules.rules](resources--protected_application--reference--group-001.md#canonical-2330032030211220-2321100333312222-2303120210101121-0201212120201330-1023331123331022-2322300220313133-1220302231122021-1001302311011021)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-1222023303122201-2323122330313330-0233010223032311-1023030333132220-1001303031220002-3322201001102320-2030023220001320-2103232200313302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3100221200131001-3003120103332303-2210310000302213-1332020122013310-1322030313300233-2233222301312222-3201110221123033-3300202032000012"></a>

## Cloudflare.js_insertion_rules.exclude_list — exclude_list / 112020313123 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.js_insertion_rules](resources--protected_application--reference--group-001.md#canonical-1322002021231300-3010301333121031-1322110012103112-2020011212310320-2122221103021202-1211333331013121-3202231301310022-3131301323111101)
- Cloudflare.js_insertion_rules.exclude_list

<a id="canonical-0330213300030033-1110030201333123-1310023130220130-0011213301322023-2220120123123033-0230103302232031-3032211322202112-1313102320333300"></a>

Type: `"object"`. list nested block, Optional.

Optional JavaScript insertions exclude list of domain and path matchers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3131031011021200-2223011000012232-2301230101020023-1032122302010300-3320231231032220-2111030233020030-0113112331021022-3003202223132010"></a>

## Direct properties — exclude_list / 112020313123 / 3

- [any_domain](resources--protected_application--reference--group-001.md#canonical-0121201313330311-1321131232102210-1012133212301332-3031002312322310-2311200101200331-3111300002311313-1223231112130023-3101021320132303): complete subsection reference.

- [domain](resources--protected_application--reference--group-001.md#canonical-0212001212012013-2013030301233213-0233202313121312-2002100223112221-3022033221122000-2211312302013103-1300022230302131-1101332223030122): complete subsection reference.

- [metadata](resources--protected_application--reference--group-001.md#canonical-2331301132110211-1013333210130302-3113310213010322-2200211101212021-2231331013311230-2003020110303201-3133210302223121-2000320220303022): complete subsection reference.

- [path](resources--protected_application--reference--group-001.md#canonical-0122200033020333-1233203121310121-3300011130303222-1021030311232233-0232031012010030-3001330100331330-1002033013202110-3303222112120231): complete subsection reference.

<a id="canonical-1021113313321232-2203221323323022-1033003121003320-1122323033032330-2202223332330012-1123322310123033-3332303230202210-3133331120001231"></a>

## Next pages — exclude_list / 112020313123 / 4

- [cloudflare.js_insertion_rules.exclude_list.any_domain](resources--protected_application--reference--group-001.md#canonical-0121201313330311-1321131232102210-1012133212301332-3031002312322310-2311200101200331-3111300002311313-1223231112130023-3101021320132303)
- [cloudflare.js_insertion_rules.exclude_list.domain](resources--protected_application--reference--group-001.md#canonical-0212001212012013-2013030301233213-0233202313121312-2002100223112221-3022033221122000-2211312302013103-1300022230302131-1101332223030122)
- [cloudflare.js_insertion_rules.exclude_list.metadata](resources--protected_application--reference--group-001.md#canonical-2331301132110211-1013333210130302-3113310213010322-2200211101212021-2231331013311230-2003020110303201-3133210302223121-2000320220303022)
- [cloudflare.js_insertion_rules.exclude_list.path](resources--protected_application--reference--group-001.md#canonical-0122200033020333-1233203121310121-3300011130303222-1021030311232233-0232031012010030-3001330100331330-1002033013202110-3303222112120231)
- [cloudflare.js_insertion_rules](resources--protected_application--reference--group-001.md#canonical-1322002021231300-3010301333121031-1322110012103112-2020011212310320-2122221103021202-1211333331013121-3202231301310022-3131301323111101)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-0121201313330311-1321131232102210-1012133212301332-3031002312322310-2311200101200331-3111300002311313-1223231112130023-3101021320132303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221320232123203-1003032212011012-1020102320313010-0202020311022220-2130013122132310-3013311223231123-2011233230013230-1022012321213313"></a>

## Cloudflare.js_insertion_rules.exclude_list.any_domain — any_domain / 010003212032 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.js_insertion_rules](resources--protected_application--reference--group-001.md#canonical-1322002021231300-3010301333121031-1322110012103112-2020011212310320-2122221103021202-1211333331013121-3202231301310022-3131301323111101)
- [cloudflare.js_insertion_rules.exclude_list](resources--protected_application--reference--group-001.md#canonical-1222023303122201-2323122330313330-0233010223032311-1023030333132220-1001303031220002-3322201001102320-2030023220001320-2103232200313302)
- Cloudflare.js_insertion_rules.exclude_list.any_domain

<a id="canonical-2031330111111121-1021113020310323-2303123310132330-3010221300300221-0001203012323122-3131102110222131-3211232113222022-1030020120221201"></a>

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
any_domain = {}
```

<a id="canonical-0020321031012220-3023330010220232-3202233230131320-2133030202220131-2330211100220033-1030232333001332-3231330100122312-2202032300003131"></a>

## Direct properties — any_domain / 010003212032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0123301002033302-1320231010331232-2302301003300320-3310123311332121-3123100022212011-1110133023120211-2112012033013312-2022200122313000"></a>

## Next pages — any_domain / 010003212032 / 4

- [cloudflare.js_insertion_rules.exclude_list](resources--protected_application--reference--group-001.md#canonical-1222023303122201-2323122330313330-0233010223032311-1023030333132220-1001303031220002-3322201001102320-2030023220001320-2103232200313302)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-0212001212012013-2013030301233213-0233202313121312-2002100223112221-3022033221122000-2211312302013103-1300022230302131-1101332223030122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023310013010110-2000133033112031-3023001012212223-1310002111200123-3111001010010210-2211131312211002-0121312112102001-3330203022301202"></a>

## Cloudflare.js_insertion_rules.exclude_list.domain — domain / 313323320101 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.js_insertion_rules](resources--protected_application--reference--group-001.md#canonical-1322002021231300-3010301333121031-1322110012103112-2020011212310320-2122221103021202-1211333331013121-3202231301310022-3131301323111101)
- [cloudflare.js_insertion_rules.exclude_list](resources--protected_application--reference--group-001.md#canonical-1222023303122201-2323122330313330-0233010223032311-1023030333132220-1001303031220002-3322201001102320-2030023220001320-2103232200313302)
- Cloudflare.js_insertion_rules.exclude_list.domain

<a id="canonical-3033003002001321-3032023022031233-2131001001023210-2111323201122213-0112100312323322-2300210213013120-3112011331010112-2120130231230001"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-2321122101200203-2230000323201033-1000132123110301-1202003232003000-0230033322113003-3122111021231322-1113203331033233-2020130012330022"></a>

## Direct properties — domain / 313323320101 / 3

<a id="canonical-1330123232333122-3031111330313232-1133303123010213-0210301001113332-1011323031110311-3222310321212001-2302131031132321-2322232301021001"></a>

<a id="canonical-0233002133020112-3302131320110311-3020102030021000-3033231322130122-0120022230031230-0323101131330233-1120112300230023-0301213233311102"></a>

## exact_value property — domain / 313323320101 / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
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
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0232003313133323-2231033303211131-0131302333221303-1100301231020030-3321311120211010-3200331132300012-2010231000013233-1301220313023031"></a>

<a id="canonical-3133121331212233-2122133231331310-0102303120022003-3131100200323010-2310322111332011-2011101233232030-1311022102231231-0022010311331012"></a>

## regex_value property — domain / 313323320101 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
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
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2312101303000212-1021131030030012-3233232023332003-3320201102130313-0212132200013233-3012212013122021-3033121032130033-1213300203100212"></a>

<a id="canonical-1333012212331111-3123103330131223-3221032130003113-3123202233003131-3310322033223132-3233023102020300-1301323133000100-2300111322113130"></a>

## suffix_value property — domain / 313323320101 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
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
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1312033323210121-3011002110113211-3230121212310223-0031300132132122-3211300123311000-3330033322322321-3303130103130003-0000012230300210"></a>

## Next pages — domain / 313323320101 / 7

- [cloudflare.js_insertion_rules.exclude_list](resources--protected_application--reference--group-001.md#canonical-1222023303122201-2323122330313330-0233010223032311-1023030333132220-1001303031220002-3322201001102320-2030023220001320-2103232200313302)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-2331301132110211-1013333210130302-3113310213010322-2200211101212021-2231331013311230-2003020110303201-3133210302223121-2000320220303022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010300321103330-1002213203331302-3123211310332012-0203012101002312-1110023321301313-0101133320201331-2121120101132320-3010313110323203"></a>

## Cloudflare.js_insertion_rules.exclude_list.metadata — metadata / 101203030120 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.js_insertion_rules](resources--protected_application--reference--group-001.md#canonical-1322002021231300-3010301333121031-1322110012103112-2020011212310320-2122221103021202-1211333331013121-3202231301310022-3131301323111101)
- [cloudflare.js_insertion_rules.exclude_list](resources--protected_application--reference--group-001.md#canonical-1222023303122201-2323122330313330-0233010223032311-1023030333132220-1001303031220002-3322201001102320-2030023220001320-2103232200313302)
- Cloudflare.js_insertion_rules.exclude_list.metadata

<a id="canonical-2321032222102220-0303103210112032-2002030320130020-1131123202111002-2123231301223231-1132113103222322-3231023202203200-2113011103300200"></a>

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

<a id="canonical-2203020001000003-3230332132310323-0212101022032312-3110021233130022-3013332302302320-3233022031213201-1032103302322033-3032221012323312"></a>

## Direct properties — metadata / 101203030120 / 3

<a id="canonical-1321001102203001-0121122003121332-2322121303031012-3110211312322202-3002101323220221-1201203000301210-2321110332120120-0020331211013123"></a>

<a id="canonical-3010001303021120-2131301113120211-0303010103302221-3323123213132003-0113110330212202-2112121323233121-1310112322211223-2032132303020202"></a>

## description_spec property — metadata / 101203030120 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-3232132001112220-0131320302111200-2301021320211022-3300310203213210-0301222323322102-0230023000030330-0212220100002120-0320122330010230"></a>

<a id="canonical-2021201121122003-0131331232301210-3110032000310011-2033002320210312-3331132013332311-3331232213300122-3030003331021113-2331020111123111"></a>

## name property — metadata / 101203030120 / 5

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

<a id="canonical-2132222121202122-1300202001132031-0232210031300102-1003302300223011-2311121332220010-0032232203011130-2232223323113033-0323023312220033"></a>

## Next pages — metadata / 101203030120 / 6

- [cloudflare.js_insertion_rules.exclude_list](resources--protected_application--reference--group-001.md#canonical-1222023303122201-2323122330313330-0233010223032311-1023030333132220-1001303031220002-3322201001102320-2030023220001320-2103232200313302)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-0122200033020333-1233203121310121-3300011130303222-1021030311232233-0232031012010030-3001330100331330-1002033013202110-3303222112120231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1130322133010301-1331200213112333-3123232201300332-3112301022221330-3221003102210303-2121221012131220-3020010120233021-1200203232001230"></a>

## Cloudflare.js_insertion_rules.exclude_list.path — path / 112323023023 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.js_insertion_rules](resources--protected_application--reference--group-001.md#canonical-1322002021231300-3010301333121031-1322110012103112-2020011212310320-2122221103021202-1211333331013121-3202231301310022-3131301323111101)
- [cloudflare.js_insertion_rules.exclude_list](resources--protected_application--reference--group-001.md#canonical-1222023303122201-2323122330313330-0233010223032311-1023030333132220-1001303031220002-3322201001102320-2030023220001320-2103232200313302)
- Cloudflare.js_insertion_rules.exclude_list.path

<a id="canonical-2132122231021221-0233010130120121-3232212003123312-1013333213103031-2212302101011010-2210111333100323-1012213330230022-3202002300113302"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-2121102333020301-2123322123101332-0003232122013002-2211110002112131-0302120100112211-0030321203230210-0102111012003130-3123323333202112"></a>

## Direct properties — path / 112323023023 / 3

<a id="canonical-2100023312320323-2112322122201022-1330021323032112-0310002001203021-0011033022220210-2300232000001000-0000332331112130-3321102221313321"></a>

<a id="canonical-0230030322100013-2022101021332200-2123320311300012-1111010322233310-0111212133221111-2103022001223013-2113332333312311-3113213002133120"></a>

## path property — path / 112323023023 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
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
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2000111232200000-1103133303121302-0033320110130323-0322033310032111-2323013032012232-3232001333123232-2312033001033302-1102103111110103"></a>

<a id="canonical-2320222031120110-1322330331230002-0020231333022103-1303003220321302-2312003323030122-0231102331010123-1211313231200100-1110030312021012"></a>

## prefix property — path / 112323023023 / 5

Type: `"string"`. Optional.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1020211323210030-2031303011231032-1133223020022030-2213220313221310-1220100312123022-3200131023202101-0301333231023110-3031012100333210"></a>

<a id="canonical-0232110110010212-3003123120323233-1302003032132311-1310331100202002-2303203022232030-3332303110031123-3003010322310122-3131101223021022"></a>

## regex property — path / 112323023023 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Provider validators and defaults (from schema source):

```go
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
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-1101301303231002-2231331321333001-2031223200323113-0212330123030003-0100013220022031-1330112221121212-2003221112210232-3313030001220033"></a>

## Next pages — path / 112323023023 / 7

- [cloudflare.js_insertion_rules.exclude_list](resources--protected_application--reference--group-001.md#canonical-1222023303122201-2323122330313330-0233010223032311-1023030333132220-1001303031220002-3322201001102320-2030023220001320-2103232200313302)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-2330032030211220-2321100333312222-2303120210101121-0201212120201330-1023331123331022-2322300220313133-1220302231122021-1001302311011021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310111001022333-3102221001122330-2311022030020221-3102323313013012-2220022212301301-2123302023223030-0121032221222033-1231021113332032"></a>

## Cloudflare.js_insertion_rules.rules — rules / 212221133233 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.js_insertion_rules](resources--protected_application--reference--group-001.md#canonical-1322002021231300-3010301333121031-1322110012103112-2020011212310320-2122221103021202-1211333331013121-3202231301310022-3131301323111101)
- Cloudflare.js_insertion_rules.rules

<a id="canonical-0213231112230230-2001130113200021-2321313322200002-3003332300112130-3011323003230121-1312023203011201-2122321332032321-3330322333200231"></a>

Type: `"object"`. list nested block, Optional.

Required list of pages to insert Bot Defense client JavaScript.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain"),
  validators.ConflictingListObjectAttributes("exact_path",
    "glob"),
  validators.ConflictingListObjectAttributes("exact_path",
    "prefix"),
  validators.ConflictingListObjectAttributes("glob",
    "prefix")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-2110022330313012-0222302222122133-2302222233001320-1330013003321101-3213033200221201-2030011111023030-2110203303302330-0030333003121232"></a>

## Direct properties — rules / 212221133233 / 3

- [any_domain](resources--protected_application--reference--group-001.md#canonical-0023021323303103-0100112231022022-3133113301032122-1303303121031232-0232200002300233-3200130212011111-3312131120003023-2333132102202220): complete subsection reference.

- [domain](resources--protected_application--reference--group-001.md#canonical-1023321211133000-2323310121123222-0301321300100100-3203001112233101-2020021132300012-2010123323000301-3232121011223300-3313203002022131): complete subsection reference.

<a id="canonical-0333320100031011-1233001330102232-2303333220333002-1123221210333003-2120133123102020-2130133031210100-3231120311212223-3310021302103233"></a>

<a id="canonical-0030013012302213-0213011100110211-1012013131122300-0130100220113031-0233302013211113-0303021031222103-3021120202233212-2002030330110321"></a>

## exact_path property — rules / 212221133233 / 4

Type: `"string"`. Optional.

Exclusive with \[glob prefix\] Exact path value to match.

Upstream description:

Exclusive with \[glob prefix\] Exact path value to match.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2300123200201111-2303330233112130-1033013032012223-1230023133311212-2110003010200302-3232123212311013-2003113012021023-0123110011231310"></a>

<a id="canonical-2321222322313202-3332013101231000-2203120100211113-3220230212003121-3123321320201201-3013012000332200-2012331331033110-0002220203130133"></a>

## glob property — rules / 212221133233 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_path prefix\] Accepts wildcards \* to match multiple characters or ? To
match a single character.

Upstream description:

Exclusive with \[exact\_path prefix\]

Accepts wildcards \* to match multiple characters or ? To match a single character.

Provider validators and defaults (from schema source):

```go
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
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,256}$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,256}$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,256}$"
  }
}
```

- [metadata](resources--protected_application--reference--group-001.md#canonical-2223230102211023-3322202112210122-0011332033131022-0101002012232222-1201232030100323-0312022032230231-1111122010031330-0223332132113123): complete subsection reference.

<a id="canonical-1031323110301022-0033213333030012-3132132112133132-3111311220121001-0012223301312201-2000002013131110-3211330133332110-3102020022301231"></a>

<a id="canonical-0301320223202020-3232210203123031-3231103012000101-1222330230022223-3121130303232232-1012301311221233-3122230110010112-1330302122100303"></a>

## prefix property — rules / 212221133233 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_path glob\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[exact\_path glob\] Path prefix to match (e.g. The value / will match on all paths)

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3100022032332032-2101001002022031-3201223003033020-2100202120130212-2002001213100132-0000032322002103-2111301301321231-2230110323001111"></a>

## Next pages — rules / 212221133233 / 7

- [cloudflare.js_insertion_rules.rules.any_domain](resources--protected_application--reference--group-001.md#canonical-0023021323303103-0100112231022022-3133113301032122-1303303121031232-0232200002300233-3200130212011111-3312131120003023-2333132102202220)
- [cloudflare.js_insertion_rules.rules.domain](resources--protected_application--reference--group-001.md#canonical-1023321211133000-2323310121123222-0301321300100100-3203001112233101-2020021132300012-2010123323000301-3232121011223300-3313203002022131)
- [cloudflare.js_insertion_rules.rules.metadata](resources--protected_application--reference--group-001.md#canonical-2223230102211023-3322202112210122-0011332033131022-0101002012232222-1201232030100323-0312022032230231-1111122010031330-0223332132113123)
- [cloudflare.js_insertion_rules](resources--protected_application--reference--group-001.md#canonical-1322002021231300-3010301333121031-1322110012103112-2020011212310320-2122221103021202-1211333331013121-3202231301310022-3131301323111101)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-0023021323303103-0100112231022022-3133113301032122-1303303121031232-0232200002300233-3200130212011111-3312131120003023-2333132102202220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2222120330121302-1210332011322223-2021131313122230-2110330333002202-1230013121121013-1123110110003111-2001301220302331-0122001311200031"></a>

## Cloudflare.js_insertion_rules.rules.any_domain — any_domain / 200002022221 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.js_insertion_rules](resources--protected_application--reference--group-001.md#canonical-1322002021231300-3010301333121031-1322110012103112-2020011212310320-2122221103021202-1211333331013121-3202231301310022-3131301323111101)
- [cloudflare.js_insertion_rules.rules](resources--protected_application--reference--group-001.md#canonical-2330032030211220-2321100333312222-2303120210101121-0201212120201330-1023331123331022-2322300220313133-1220302231122021-1001302311011021)
- Cloudflare.js_insertion_rules.rules.any_domain

<a id="canonical-3000113222132100-3121130210002030-2012101223131123-2003331220012313-0203200322312020-0323302322233030-0300031223202313-1231302231203112"></a>

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
any_domain = {}
```

<a id="canonical-0231231233010202-3001123203300002-2010030110202330-2023200320202011-3233021032003121-0232112111220222-1310230001301323-0220312130212022"></a>

## Direct properties — any_domain / 200002022221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0121131203330311-3133132021302323-2331323010120320-1231213210330132-2132210201132131-0233113002310133-0130330202022012-1130311311233023"></a>

## Next pages — any_domain / 200002022221 / 4

- [cloudflare.js_insertion_rules.rules](resources--protected_application--reference--group-001.md#canonical-2330032030211220-2321100333312222-2303120210101121-0201212120201330-1023331123331022-2322300220313133-1220302231122021-1001302311011021)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-1023321211133000-2323310121123222-0301321300100100-3203001112233101-2020021132300012-2010123323000301-3232121011223300-3313203002022131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002013333310312-0013310100232300-3323310112021013-1030130002113210-1010010313022012-2130120101200313-0300113212123322-2013321120103001"></a>

## Cloudflare.js_insertion_rules.rules.domain — domain / 131010022200 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.js_insertion_rules](resources--protected_application--reference--group-001.md#canonical-1322002021231300-3010301333121031-1322110012103112-2020011212310320-2122221103021202-1211333331013121-3202231301310022-3131301323111101)
- [cloudflare.js_insertion_rules.rules](resources--protected_application--reference--group-001.md#canonical-2330032030211220-2321100333312222-2303120210101121-0201212120201330-1023331123331022-2322300220313133-1220302231122021-1001302311011021)
- Cloudflare.js_insertion_rules.rules.domain

<a id="canonical-1312020132101210-1200331102023112-0012222131010033-3320120112103010-2103122232003330-1212020330220102-2123233233211101-3232113223220001"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-3210010303201022-1120111210322322-3312222232000032-2223232333112301-2210320230320120-0222332133102320-2332303001232033-0031010010020311"></a>

## Direct properties — domain / 131010022200 / 3

<a id="canonical-1311113103220100-1232310110012013-2321323121011310-0021303130023012-3020031212133333-0030101030101031-0121012001223021-1312331212113301"></a>

<a id="canonical-1100212203120222-2211120101000132-1103203020313222-0012012320101022-0210212203233002-0202001221203102-0330233022222221-0101203212113233"></a>

## exact_value property — domain / 131010022200 / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
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
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3203332331033212-0000320103223322-2013232301210132-2201300103333023-0211112321120301-0231221223132111-1210322200322231-0203302312121330"></a>

<a id="canonical-0311111222220203-3111201333110313-3302123210330210-0203200032021000-1332033301313103-0313110120320202-0102211012120112-2222321311130133"></a>

## regex_value property — domain / 131010022200 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
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
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2131122113001331-1100220120212201-0003022230021123-2311302001322111-2300212023323110-2300213332331113-3100000202331302-2010200232311331"></a>

<a id="canonical-0023203000310130-2310313021102102-2030100021112330-3303221001101221-0100121130012221-1020002201311022-0210103111021000-3201011010320111"></a>

## suffix_value property — domain / 131010022200 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
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
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2221201022232121-0120313303311211-0113223033133002-2202132332003000-0012222101312012-0123231312320313-0323201311003112-3202001301300100"></a>

## Next pages — domain / 131010022200 / 7

- [cloudflare.js_insertion_rules.rules](resources--protected_application--reference--group-001.md#canonical-2330032030211220-2321100333312222-2303120210101121-0201212120201330-1023331123331022-2322300220313133-1220302231122021-1001302311011021)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-2223230102211023-3322202112210122-0011332033131022-0101002012232222-1201232030100323-0312022032230231-1111122010031330-0223332132113123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2033011130301320-2032023331001223-2021003301313330-1320300030100210-3331012301210012-3132321222120031-2233222031230120-2011332000120301"></a>

## Cloudflare.js_insertion_rules.rules.metadata — metadata / 203120122210 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.js_insertion_rules](resources--protected_application--reference--group-001.md#canonical-1322002021231300-3010301333121031-1322110012103112-2020011212310320-2122221103021202-1211333331013121-3202231301310022-3131301323111101)
- [cloudflare.js_insertion_rules.rules](resources--protected_application--reference--group-001.md#canonical-2330032030211220-2321100333312222-2303120210101121-0201212120201330-1023331123331022-2322300220313133-1220302231122021-1001302311011021)
- Cloudflare.js_insertion_rules.rules.metadata

<a id="canonical-2100202210303030-0321111231302333-0103011101332011-3213311233213333-0112013023002032-2121232322021130-2132022010331212-1221221002322301"></a>

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

<a id="canonical-2221101010202001-3131013012212332-1111223002103102-1000121033003013-1123122020320003-2331211200110212-2322233232300132-2303230122122200"></a>

## Direct properties — metadata / 203120122210 / 3

<a id="canonical-2302033303312113-1002300101221003-0001003133110012-1103232220200231-2132201033011022-1112212301011111-1210010330000103-2113232320021311"></a>

<a id="canonical-2123311302121212-3202330031120311-3301002032312212-1210101213323312-3210200300032120-3131323122001021-1220133303021311-0100023131301313"></a>

## description_spec property — metadata / 203120122210 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-3200212322210020-0300312312203002-1231232010313200-3231113201102023-0111303013220032-2210203010220110-0211020322313302-0111122030310110"></a>

<a id="canonical-3121032021032022-1311131031003022-1302220022033211-3122201312210203-1013203320330102-3030310312232323-3121001320201322-0101333312120321"></a>

## name property — metadata / 203120122210 / 5

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

<a id="canonical-3123012113110100-2132302220230232-1123200201232133-1110220202032222-3103122031222133-0121331123022220-0023121200330121-3003211203320210"></a>

## Next pages — metadata / 203120122210 / 6

- [cloudflare.js_insertion_rules.rules](resources--protected_application--reference--group-001.md#canonical-2330032030211220-2321100333312222-2303120210101121-0201212120201330-1023331123331022-2322300220313133-1220302231122021-1001302311011021)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-2302310133022000-1230130121013212-3323120233122330-2221122132220001-2122001133122302-2132010313002023-3012313120130203-3123232120220302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0102131133010132-1000200202220000-2103202230302130-0233101311233122-0313030013002022-3303002020311302-2120320101100201-3022120020231322"></a>

## Cloudflare.manual_js_insert — manual_js_insert / 031311022021 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- Cloudflare.manual_js_insert

<a id="canonical-2210200121122101-3311010033010210-2200302230103311-0033110231131020-0120321133221223-3121010023303103-2332031313022312-0001033200121120"></a>

Type: `"object"`. single nested block, Optional.

Insert JavaScript Manually. Insert JavaScript manually.

Upstream description:

Insert JavaScript manually.

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
manual_js_insert {
  # Configure direct properties listed below.
}
```

<a id="canonical-2333222221100321-1023113001330322-2101221013213130-1232212232002030-0120303312103201-1021322320103231-3110312012101010-1133130010312111"></a>

## Direct properties — manual_js_insert / 031311022021 / 3

<a id="canonical-3100232132100013-1120300013012302-0223123312311223-1033303220011020-3223322223333333-3122312211313333-0223310212013131-1203203231001110"></a>

<a id="canonical-3321001033023112-3210001231330300-1121010300123302-3002130030202312-3222301333012111-0212012311223110-1023220302130301-0021012000212132"></a>

## js_download_path property — manual_js_insert / 031311022021 / 4

Type: `"string"`. Optional.

Web client will fetch F5 Client JavaScript from this path. This path must not conflict with any
other website/application paths. If not specified, default to ‘/common.js’.

Upstream description:

Web client will fetch F5 Client JavaScript from this path. This path must not conflict with any
other website/application paths.

If not specified, default to ‘/common.js’.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  }
}
```

<a id="canonical-1210030311230213-1103303031133132-1201300112222203-1210202321213300-1033322122231232-0001233231121213-1130201023230231-0322311231123102"></a>

## Next pages — manual_js_insert / 031311022021 / 5

- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-0122212303302030-3223122122222023-1210233312031201-3313210103123321-0333301122232233-2302231102220123-2301133202020201-3332033000221131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3302021111310132-0202123220233202-0020012330213221-0002020213301111-1300312210130213-1013233000010112-1322211122322320-3112302121020012"></a>

## Cloudflare.mobile_sdk_config — mobile_sdk_config / 023221103123 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- Cloudflare.mobile_sdk_config

<a id="canonical-0213022302033113-1300012000332200-0231123131211121-1320130133331310-2312110301200210-0112300020222023-3230112310333212-0200000322102223"></a>

Type: `"object"`. single nested block, Optional.

Mobile SDK Configuration. Mobile SDK configuration.

Upstream description:

Mobile SDK configuration.

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
mobile_sdk_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-2323230122333212-1111312210003123-0320312313000330-3101020021221231-0033331201020111-1311220001011120-1212103031302030-3210331010312212"></a>

## Direct properties — mobile_sdk_config / 023221103123 / 3

- [mobile_identifier](resources--protected_application--reference--group-001.md#canonical-1110022210110133-0302122033301331-1232032102212313-1122232331103031-3032202212222021-2000030110212110-0020331021132012-0111121110230301): complete subsection reference.

<a id="canonical-2010231111003000-0322120320301210-3120001133012200-2102122321102030-2311233232322332-3222220123013303-1111020101330202-1121010301100111"></a>

## Next pages — mobile_sdk_config / 023221103123 / 4

- [cloudflare.mobile_sdk_config.mobile_identifier](resources--protected_application--reference--group-001.md#canonical-1110022210110133-0302122033301331-1232032102212313-1122232331103031-3032202212222021-2000030110212110-0020331021132012-0111121110230301)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-1110022210110133-0302122033301331-1232032102212313-1122232331103031-3032202212222021-2000030110212110-0020331021132012-0111121110230301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112133001223232-1011022132232132-2020300321302311-2020111020233303-0311232112120331-2023030201102121-2312000103301220-1200301110003133"></a>

## Cloudflare.mobile_sdk_config.mobile_identifier — mobile_identifier / 330222301032 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.mobile_sdk_config](resources--protected_application--reference--group-001.md#canonical-0122212303302030-3223122122222023-1210233312031201-3313210103123321-0333301122232233-2302231102220123-2301133202020201-3332033000221131)
- Cloudflare.mobile_sdk_config.mobile_identifier

<a id="canonical-1021312012012320-1202220212211212-0211020132233221-2201101130021112-1312323302233001-0231200323000220-1031001132101231-0303203312331133"></a>

Type: `"object"`. single nested block, Optional.

Mobile Traffic Identifier. Mobile traffic identifier type.

Upstream description:

Mobile traffic identifier type.

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
mobile_identifier {
  # Configure direct properties listed below.
}
```

<a id="canonical-3231321101030311-2012001022122002-2231032133320130-2320110210113013-1333013122310013-1003111333103110-1100220102223030-3030331110310302"></a>

## Direct properties — mobile_identifier / 330222301032 / 3

- [headers](resources--protected_application--reference--group-001.md#canonical-1202033122231203-3111301221002203-0031301013310331-0310102032013233-3203323203232100-1101210013301232-2103032112020031-0021100023033332): complete subsection reference.

<a id="canonical-1103103332031020-3133332033223112-3302221122010310-3323130320321203-2112001011102123-1112213332133222-2222213111203120-0101021013313110"></a>

## Next pages — mobile_identifier / 330222301032 / 4

- [cloudflare.mobile_sdk_config.mobile_identifier.headers](resources--protected_application--reference--group-001.md#canonical-1202033122231203-3111301221002203-0031301013310331-0310102032013233-3203323203232100-1101210013301232-2103032112020031-0021100023033332)
- [cloudflare.mobile_sdk_config](resources--protected_application--reference--group-001.md#canonical-0122212303302030-3223122122222023-1210233312031201-3313210103123321-0333301122232233-2302231102220123-2301133202020201-3332033000221131)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-1202033122231203-3111301221002203-0031301013310331-0310102032013233-3203323203232100-1101210013301232-2103032112020031-0021100023033332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010020123100021-0311212221220321-1320301013031333-1300312103220023-3311303310131203-3113122223320212-0220002030210112-3310102303203310"></a>

## Cloudflare.mobile_sdk_config.mobile_identifier.headers — headers / 201021221110 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.mobile_sdk_config](resources--protected_application--reference--group-001.md#canonical-0122212303302030-3223122122222023-1210233312031201-3313210103123321-0333301122232233-2302231102220123-2301133202020201-3332033000221131)
- [cloudflare.mobile_sdk_config.mobile_identifier](resources--protected_application--reference--group-001.md#canonical-1110022210110133-0302122033301331-1232032102212313-1122232331103031-3032202212222021-2000030110212110-0020331021132012-0111121110230301)
- Cloudflare.mobile_sdk_config.mobile_identifier.headers

<a id="canonical-1201231030100212-0233131131010302-0201231221233300-1311223223132012-0212101030030202-1010230221303313-1002012310123232-3023223311223131"></a>

Type: `"object"`. list nested block, Optional.

List of headers that can be used to identify mobile traffic.

Upstream description:

A list of headers that can be used to identify mobile traffic.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("exact",
    "regex")}
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
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-3202001222011222-2220131122000021-3003123210013010-2202210111311322-1321300202100303-2213020102233222-1221233012003020-2101030012313200"></a>

## Direct properties — headers / 201021221110 / 3

<a id="canonical-3010133000331023-2333120330031212-0220201232301132-0033023232320100-2101303113331231-0012222121322201-1333011102213210-1121101132123321"></a>

<a id="canonical-0331001230213023-0011232132123300-3201122011311311-0333003122333112-1321311203012331-0321111101333102-2013230000033131-2313011021322212"></a>

## exact property — headers / 201021221110 / 4

Type: `"string"`. Optional.

Exclusive with \[regex\] Header value to match exactly.

Upstream description:

Exclusive with \[regex\] Header value to match exactly.

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
    "byteLength": {
      "max": 256
    },
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-3321100210233001-1003301231333331-0312230022111013-1031203230302321-3300113312313111-0011003300022303-0102223122130212-0023001233011313"></a>

<a id="canonical-3330123221001201-0212113123200232-0220033121312020-0301231313013313-0223200031231023-2012222203230032-0211003212103312-2013302113102301"></a>

## name property — headers / 201021221110 / 5

Type: `"string"`. Optional.

Name. Name of the header.

Upstream description:

Name of the header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-2113222220113023-2310021120233033-2023020310320113-3121032233102123-0133020001002303-0300300002020201-2202010232200113-2113300102320212"></a>

<a id="canonical-1022221313102220-3310222230131023-0113333223023123-1131333222321013-1302222323031031-1302103010031231-0100033301010103-2102230220202301"></a>

## regex property — headers / 201021221110 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\] Regex match of the header value in re2 format.

Upstream description:

Exclusive with \[exact\] Regex match of the header value in re2 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-1232310311222001-3122132030233213-2221323313203133-1222122030220201-1022330022202332-0232223103223021-3303221102032220-0210232110001023"></a>

## Next pages — headers / 201021221110 / 7

- [cloudflare.mobile_sdk_config.mobile_identifier](resources--protected_application--reference--group-001.md#canonical-1110022210110133-0302122033301331-1232032102212313-1122232331103031-3032202212222021-2000030110212110-0020331021132012-0111121110230301)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011030223333212-0033203200322013-2300303222230033-1300022012322111-3230322312100003-2202113123110303-2313333002022231-1131201232002133"></a>

## Cloudflare.protected_endpoints — protected_endpoints / 130210002022 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- Cloudflare.protected_endpoints

<a id="canonical-3212033021020010-1201311130333123-0323230221231130-1210101103312231-2101031010300201-2333230031223101-1030211303013203-0120330110233201"></a>

Type: `"object"`. list nested block, Optional.

List of protected endpoints (max 128 items).

Upstream description:

List of protected endpoints (max 128 items)

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("http_methods"),
  validators.ConflictingListObjectAttributes("any_domain",
    "domain"),
  validators.ConflictingListObjectAttributes("mobile_client",
    "web_client"),
  validators.ConflictingListObjectAttributes("mobile_client",
    "web_mobile_client"),
  validators.ConflictingListObjectAttributes("web_client",
    "web_mobile_client")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
protected_endpoints {
  # Configure direct properties listed below.
}
```

<a id="canonical-2103210002122131-2020001211131033-3130131023033311-0213122110022122-0310320110131030-3221321222221222-0233322031111030-0032021132120230"></a>

## Direct properties — protected_endpoints / 130210002022 / 3

- [any_domain](resources--protected_application--reference--group-001.md#canonical-1033311200213201-3332033221131111-2122101202012030-0220332233322202-2120211003233311-3213303210333212-2203003000022020-1313000303312232): complete subsection reference.

- [domain](resources--protected_application--reference--group-002.md#canonical-3022333301023332-2020221211122132-2220331300033002-3210230122323122-0301223130033002-1110232320103200-2132211310223201-0323000301133331): complete subsection reference.

<a id="canonical-2030333122003212-2320213012110221-0000211203030033-2010333203201321-1221013033311311-1113331000113300-3110200320331133-1032223033121312"></a>

<a id="canonical-0131332011303121-3332303112120113-3030211012222103-1132300012012133-1212311203123220-1023133201132123-1031313212003212-2100102022002221"></a>

## http_methods property — protected_endpoints / 130210002022 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
METHOD\_ANY|METHOD\_GET|METHOD\_POST|METHOD\_PUT|METHOD\_PATCH|METHOD\_DELETE|METHOD\_GET\_DOCUMENT\]
HTTP Methods. List of HTTP methods. Possible values are \`METHOD\_ANY\`, \`METHOD\_GET\`,
\`METHOD\_POST\`, \`METHOD\_PUT\`, \`METHOD\_PATCH\`, \`METHOD\_DELETE\`, \`METHOD\_GET\_DOCUMENT\`.
Defaults to \`METHOD\_ANY\`.

Upstream description:

List of HTTP methods.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[1,3,4]",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[1,3,4]",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [metadata](resources--protected_application--reference--group-002.md#canonical-0313013122021223-2132102100230101-1111201212133020-3103013102302032-0010021223220331-1202010100113100-3003121222312133-2303311000323323): complete subsection reference.

- [mobile_client](resources--protected_application--reference--group-002.md#canonical-2231203002122332-0001110313310323-3123313122330323-0133123113300301-0302002231331030-1132302003033313-1122331331111221-2320131231013033): complete subsection reference.

- [path](resources--protected_application--reference--group-002.md#canonical-1001300212102331-3212330102111323-3100201020003211-0232102000201230-2211133203031111-3132133032303233-3312123120323022-0303222121321212): complete subsection reference.

<a id="canonical-3033211113312003-0001210311213230-2200020200322332-3222012331003121-3000201230112203-1222030102113213-1132313012322123-0111001313321323"></a>

<a id="canonical-2212213323033113-1013322130111132-3332233211330202-2200101010001202-1323232331210222-0022023020201132-2002230110111201-3213103213202223"></a>

## query property — protected_endpoints / 130210002022 / 5

Type: `"string"`. Optional.

Enter a regular expression to match your query parameters of interest.

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
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

- [web_client](resources--protected_application--reference--group-002.md#canonical-0213202232233011-0000312212000013-2113100002231110-2301000001213232-0031000313130131-1333121310313202-2201321230031021-0123023212011332): complete subsection reference.

- [web_mobile_client](resources--protected_application--reference--group-002.md#canonical-1113311232320333-0300120101320203-2110223102112213-2013121133211100-3321120132301211-2020003231031133-1323311232301303-1022020213311220): complete subsection reference.

<a id="canonical-1122012122300310-2330011123221232-0013121330310230-3323032021000310-2203213312100010-3212111100233033-2033220110131102-0121320032120230"></a>

## Next pages — protected_endpoints / 130210002022 / 6

- [cloudflare.protected_endpoints.any_domain](resources--protected_application--reference--group-001.md#canonical-1033311200213201-3332033221131111-2122101202012030-0220332233322202-2120211003233311-3213303210333212-2203003000022020-1313000303312232)
- [cloudflare.protected_endpoints.domain](resources--protected_application--reference--group-002.md#canonical-3022333301023332-2020221211122132-2220331300033002-3210230122323122-0301223130033002-1110232320103200-2132211310223201-0323000301133331)
- [cloudflare.protected_endpoints.metadata](resources--protected_application--reference--group-002.md#canonical-0313013122021223-2132102100230101-1111201212133020-3103013102302032-0010021223220331-1202010100113100-3003121222312133-2303311000323323)
- [cloudflare.protected_endpoints.mobile_client](resources--protected_application--reference--group-002.md#canonical-2231203002122332-0001110313310323-3123313122330323-0133123113300301-0302002231331030-1132302003033313-1122331331111221-2320131231013033)
- [cloudflare.protected_endpoints.path](resources--protected_application--reference--group-002.md#canonical-1001300212102331-3212330102111323-3100201020003211-0232102000201230-2211133203031111-3132133032303233-3312123120323022-0303222121321212)
- [cloudflare.protected_endpoints.web_client](resources--protected_application--reference--group-002.md#canonical-0213202232233011-0000312212000013-2113100002231110-2301000001213232-0031000313130131-1333121310313202-2201321230031021-0123023212011332)
- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-1113311232320333-0300120101320203-2110223102112213-2013121133211100-3321120132301211-2020003231031133-1323311232301303-1022020213311220)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-1033311200213201-3332033221131111-2122101202012030-0220332233322202-2120211003233311-3213303210333212-2203003000022020-1313000303312232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
