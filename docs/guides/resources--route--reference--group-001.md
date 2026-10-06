---
page_title: "xcsh_route reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_route reference."
---

# xcsh_route reference

<a id="canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- Property reference

<a id="canonical-1303023313102113-1022213100200030-3230112303001012-3021200031122103-0231311203133003-3012210033022213-1110101323211030-1323201111022002"></a>

### Direct properties for `xcsh_route`

<a id="canonical-0302123133012110-0003331310132110-0111231112010023-3112210312011113-2100012000013201-1033032301312220-2201323201332301-0000023211300132"></a>

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

<a id="canonical-1000331032312212-1333121103301222-1231202030233122-3020011101301303-2311022101022132-2311311332130002-1212221013303020-2123120211332213"></a>

<a id="canonical-2111311020323123-2323200220131012-0131121012012221-2220002000113022-0213111112200312-1222212303012110-2002212132021330-0011332323131010"></a>

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

<a id="canonical-0123222233122102-3320332300122320-2231223222202031-1332100113033322-3233212232333213-3222002032020322-0010231023003101-3230231210023322"></a>

<a id="canonical-0101001113121231-2113303003210013-2221323122010100-2113232230030012-0003312221300303-3100013312202333-0130221133101130-2210320203102323"></a>

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

<a id="canonical-2322110002100111-3123103111333231-0123133011222112-1220322100320200-0220222121102213-0111221330212032-2012102322010032-1301011122120031"></a>

<a id="canonical-1100133310312311-2213221021300301-0132031022301011-3202212002001133-0112331100123113-3221012320322020-0310130312111300-3322032131030213"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3123021322012221-2011110021121023-3333221222120213-2313330210123202-2100132031331303-2301133031011021-2102023322131231-0330312302221332"></a>

<a id="canonical-0011232110032210-1212233311023000-2210323022111322-0223011230300223-3232010231223123-0212120121022100-0002321221122101-2330100032212320"></a>

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

<a id="canonical-3321133000112113-3231100212211132-2123023012321202-0101111200333021-1210312103331302-1001021101300010-0203301103330032-2111323022213100"></a>

<a id="canonical-2030320030210302-3333020232330201-2000233212132020-3302113233323120-3120031133110301-1032032301120233-3001120021222132-1331012312133301"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Route. Must be unique within the namespace.

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

<a id="canonical-3200012311000202-2131301313301233-2223301312230001-0213211213201321-3300001000221212-3320021313122010-3332111312021021-0013001321013322"></a>

<a id="canonical-1030102030120020-2321010131123013-1030201001232300-2133210202023201-1120311311330223-1133123131011332-0223102312301121-3220030022321103"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Route is created.

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

- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130): complete subsection reference.

- [timeouts](resources--route--reference--group-003.md#canonical-2000122211223303-0323203133120331-3213333010022223-1120331312123010-0233002030003331-1012023100032301-3200122303322132-3110023202032230): complete subsection reference.

<a id="canonical-3233313001133010-1310003320103001-1012210130331213-1212001220021313-1320120030031230-0001000223322320-3120302331322102-1000222133123022"></a>

### All schema paths for `xcsh_route`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--route--reference--group-001.md#canonical-0302123133012110-0003331310132110-0111231112010023-3112210312011113-2100012000013201-1033032301312220-2201323201332301-0000023211300132) |
| `description` | [description](resources--route--reference--group-001.md#canonical-1000331032312212-1333121103301222-1231202030233122-3020011101301303-2311022101022132-2311311332130002-1212221013303020-2123120211332213) |
| `disable` | [disable](resources--route--reference--group-001.md#canonical-0123222233122102-3320332300122320-2231223222202031-1332100113033322-3233212232333213-3222002032020322-0010231023003101-3230231210023322) |
| `id` | [ID](resources--route--reference--group-001.md#canonical-2322110002100111-3123103111333231-0123133011222112-1220322100320200-0220222121102213-0111221330212032-2012102322010032-1301011122120031) |
| `labels` | [labels](resources--route--reference--group-001.md#canonical-3123021322012221-2011110021121023-3333221222120213-2313330210123202-2100132031331303-2301133031011021-2102023322131231-0330312302221332) |
| `name` | [name](resources--route--reference--group-001.md#canonical-3321133000112113-3231100212211132-2123023012321202-0101111200333021-1210312103331302-1001021101300010-0203301103330032-2111323022213100) |
| `namespace` | [namespace](resources--route--reference--group-001.md#canonical-3200012311000202-2131301313301233-2223301312230001-0213211213201321-3300001000221212-3320021313122010-3332111312021021-0013001321013322) |
| `routes` | [routes](resources--route--reference--group-001.md#canonical-0133221031332213-0212022022311103-2201220331203131-1113221111013120-1003113131023013-3000203201133220-1323332302302103-3202120201221101) |
| `routes.bot_defense_javascript_injection` | [routes.bot_defense_javascript_injection](resources--route--reference--group-001.md#canonical-2110300030223323-2123122301232111-1311000003023132-3110003301012121-3311313231011333-0020321223333000-0110120022113310-1302111221112312) |
| `routes.bot_defense_javascript_injection.javascript_location` | [routes.bot_defense_javascript_injection.javascript_location](resources--route--reference--group-001.md#canonical-0333300210132320-2000231113231302-3212102231113112-0131101303213011-2233102323221302-2013100210333002-0000222222020300-3302303023012120) |
| `routes.bot_defense_javascript_injection.javascript_tags` | [routes.bot_defense_javascript_injection.javascript_tags](resources--route--reference--group-001.md#canonical-3023000021320011-1330102330112013-0223010021210211-3210231100312313-1131030032001313-2110010003112201-0233312213202313-2302232221021201) |
| `routes.bot_defense_javascript_injection.javascript_tags.javascript_url` | [routes.bot_defense_javascript_injection.javascript_tags.javascript_url](resources--route--reference--group-001.md#canonical-2131130212122122-0123330303311320-1002233332203212-1200323130130230-3002301333231131-3131223330221303-0330102200133110-0310123233103133) |
| `routes.bot_defense_javascript_injection.javascript_tags.tag_attributes` | [routes.bot_defense_javascript_injection.javascript_tags.tag_attributes](resources--route--reference--group-001.md#canonical-2321030323303300-0101232310030120-0322202202213132-0133133321033111-0120202222001333-0233302011010203-3100331002021012-1023303201231222) |
| `routes.bot_defense_javascript_injection.javascript_tags.tag_attributes.javascript_tag` | [routes.bot_defense_javascript_injection.javascript_tags.tag_attributes.javascript_tag](resources--route--reference--group-001.md#canonical-3013110013231302-1113133132231030-1112001130031201-0210330210122331-0103033231020220-3112212121001023-1233201321112032-0132310123133332) |
| `routes.bot_defense_javascript_injection.javascript_tags.tag_attributes.tag_value` | [routes.bot_defense_javascript_injection.javascript_tags.tag_attributes.tag_value](resources--route--reference--group-001.md#canonical-2003301222121222-3100123212002211-1312021012332002-2130011102011131-3202220132210221-3331322210203213-1012103133311002-3232300211312231) |
| `routes.disable_location_add` | [routes.disable_location_add](resources--route--reference--group-001.md#canonical-3123131022232020-1110122213003003-2311313103012220-1002200121010033-0132330130132202-3311011203303233-0320102110322331-1311310311221302) |
| `routes.inherited_bot_defense_javascript_injection` | [routes.inherited_bot_defense_javascript_injection](resources--route--reference--group-001.md#canonical-3031110333212100-3311303001122332-1222202322013313-3311021130111113-2203331200230023-3323131021212321-1230230101110210-2221003022200013) |
| `routes.inherited_waf_exclusion` | [routes.inherited_waf_exclusion](resources--route--reference--group-001.md#canonical-0222323203311203-3122123113231311-2030320123112331-3301222010331122-0011033212222010-2130102323010020-2132122111233102-0232133102302023) |
| `routes.match` | [routes.match](resources--route--reference--group-001.md#canonical-1212232011123331-0302020011221001-3303220003211103-1002301213133003-2001330133131302-3220123202131221-1200113300031113-1032022331013012) |
| `routes.match.headers` | [routes.match.headers](resources--route--reference--group-001.md#canonical-2102220120303012-3233212223322100-3331211131033303-0112220303333130-3333323230230111-1003133013333231-0212321102320333-3220012013131022) |
| `routes.match.headers.exact` | [routes.match.headers.exact](resources--route--reference--group-001.md#canonical-1100212133312132-2220022001031120-2233023111313322-1302203303101333-3330111110123011-3303113120002100-1311110221223103-1001112032133013) |
| `routes.match.headers.invert_match` | [routes.match.headers.invert_match](resources--route--reference--group-001.md#canonical-0033203301230103-0103020212311113-3212121200301323-3012200302122122-0011301111303101-3210302212021021-1232300033232220-1133033012332201) |
| `routes.match.headers.name` | [routes.match.headers.name](resources--route--reference--group-001.md#canonical-2303000310312022-1221011323220101-0331230010220221-1201120110100110-3002301021031011-3220301131211102-3112003231023201-2233101013233201) |
| `routes.match.headers.presence` | [routes.match.headers.presence](resources--route--reference--group-001.md#canonical-2012121223031133-1132112203323001-1113003311313120-3003030133320120-2213300303231323-3322311022023013-2222200222023001-3100003231122330) |
| `routes.match.headers.regex` | [routes.match.headers.regex](resources--route--reference--group-001.md#canonical-2302031023110023-1133211120233210-1213021003200121-3321221223031331-3110310100212021-0021300030101130-1212301320033001-2123013021113023) |
| `routes.match.http_method` | [routes.match.http_method](resources--route--reference--group-001.md#canonical-1300332101201021-3020022112002332-1300330031031032-2211100111020012-3022020311132321-0003333111211320-1300330232101233-0112200010312100) |
| `routes.match.incoming_port` | [routes.match.incoming_port](resources--route--reference--group-001.md#canonical-3322313333321233-1223203333211101-2103331320321321-0200210333032122-1011220021033223-0111100233311013-2321010323112230-1103130212110100) |
| `routes.match.incoming_port.no_port_match` | [routes.match.incoming_port.no_port_match](resources--route--reference--group-001.md#canonical-2203033002331200-1312012220203233-2211021011330232-2332003203312210-1032331010332213-0123002320002321-1121033231222030-3013002121303033) |
| `routes.match.incoming_port.port` | [routes.match.incoming_port.port](resources--route--reference--group-001.md#canonical-1101210011110003-1100320300132330-1033012213132200-0113200033003231-2230031300222121-2302031222231032-2101010123323323-2031233123301021) |
| `routes.match.incoming_port.port_ranges` | [routes.match.incoming_port.port_ranges](resources--route--reference--group-001.md#canonical-1220101020002313-3100203213201012-2013100023102301-1011232033231311-3320113033233333-2121331223310113-3331023010301222-0333213003230110) |
| `routes.match.path` | [routes.match.path](resources--route--reference--group-001.md#canonical-2310103323020213-2131213230202102-3220022223100010-1003323020132113-3033303303121231-0011312022333201-1203003130222021-3303200033310022) |
| `routes.match.path.path` | [routes.match.path.path](resources--route--reference--group-001.md#canonical-3330023032200200-1020022031021101-1131111310113112-3210302113111310-1002211103012320-1000122111121020-1013021113202301-3320003101001021) |
| `routes.match.path.prefix` | [routes.match.path.prefix](resources--route--reference--group-001.md#canonical-0331312111033232-2001311102130002-1310000302011113-1223223131311120-2133303313203112-3110012010322122-0220010332223011-1112233111320030) |
| `routes.match.path.regex` | [routes.match.path.regex](resources--route--reference--group-001.md#canonical-3212303032321121-2030031123331311-2011230211331201-1231323300103123-1030103232203011-2210001033010323-3233300322321130-0023301122123001) |
| `routes.match.query_params` | [routes.match.query_params](resources--route--reference--group-001.md#canonical-1321200231320101-1200322230113332-0300330030332322-2223102322333210-0010023310001221-3310310200232010-1200303113111120-0201223330223223) |
| `routes.match.query_params.exact` | [routes.match.query_params.exact](resources--route--reference--group-001.md#canonical-3203031022000302-3131320100133212-3331302120013023-3220010123332031-1032231133101031-0333101013031233-0200100312033302-1023310200210312) |
| `routes.match.query_params.key` | [routes.match.query_params.key](resources--route--reference--group-001.md#canonical-2120130330332313-0132123111222313-3321101221020110-1310301313323133-2110022333220312-1330123102133111-2321211132101131-0132003333322012) |
| `routes.match.query_params.regex` | [routes.match.query_params.regex](resources--route--reference--group-001.md#canonical-1222022213030212-1230200231111230-2101310131101000-2233321132003000-2132203133310003-1330332121031121-2133112112102102-2023311300230300) |
| `routes.request_cookies_to_add` | [routes.request_cookies_to_add](resources--route--reference--group-001.md#canonical-0213313220312023-3323111230001202-1010213210223121-2032303322033120-1311002121223311-0233300222211113-1131112010313100-3022232223300020) |
| `routes.request_cookies_to_add.name` | [routes.request_cookies_to_add.name](resources--route--reference--group-001.md#canonical-3223320331232223-1020203022131313-3232222002003003-1220223012103310-1323211013100113-1102233012212113-3101001312020011-0303213222012012) |
| `routes.request_cookies_to_add.overwrite` | [routes.request_cookies_to_add.overwrite](resources--route--reference--group-001.md#canonical-0103202201201002-0320013303131121-3111101302130330-2120302322110003-1232012102232103-0110322220023130-3030202202231202-3200022212213122) |
| `routes.request_cookies_to_add.secret_value` | [routes.request_cookies_to_add.secret_value](resources--route--reference--group-001.md#canonical-1000001123202002-2110322230212232-1032220013300111-2223210221320103-3311121032102200-2331013110312113-0202133131033200-0122032011023023) |
| `routes.request_cookies_to_add.secret_value.blindfold_secret_info` | [routes.request_cookies_to_add.secret_value.blindfold_secret_info](resources--route--reference--group-001.md#canonical-2310102223033013-2301333133211323-3230210002120111-0311310201030322-1212223202210211-2212222220211030-0332323222201030-2321131231013311) |
| `routes.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [routes.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--route--reference--group-001.md#canonical-1130101311130133-2011133302020333-2313032330002012-2010010221333033-0221301001022332-0112232112113011-1333102020110121-0002223331010122) |
| `routes.request_cookies_to_add.secret_value.blindfold_secret_info.location` | [routes.request_cookies_to_add.secret_value.blindfold_secret_info.location](resources--route--reference--group-001.md#canonical-3111232031032011-3123032202332103-1102202330002033-2112033131130131-0223301233203321-2212210201231213-3213220133200133-2231131220000301) |
| `routes.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [routes.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider](resources--route--reference--group-002.md#canonical-2302201103321222-2232302000123302-0223031030121111-3212223103132320-0130111302302100-1302203200002113-3010120122303201-0003302102233031) |
| `routes.request_cookies_to_add.secret_value.clear_secret_info` | [routes.request_cookies_to_add.secret_value.clear_secret_info](resources--route--reference--group-002.md#canonical-1302010322033002-0002023011311213-3003132213212323-2023130302230230-1102201002332113-1230000100121202-1011322023000122-3022311123223123) |
| `routes.request_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [routes.request_cookies_to_add.secret_value.clear_secret_info.provider_ref](resources--route--reference--group-002.md#canonical-0301322020023310-0332102310303010-0202011321331203-2133321000312211-3100220122223323-3020312100013103-2332330020113112-2232103011312203) |
| `routes.request_cookies_to_add.secret_value.clear_secret_info.url` | [routes.request_cookies_to_add.secret_value.clear_secret_info.url](resources--route--reference--group-002.md#canonical-3202200100333113-1210221333322311-0210331330100200-2331332211120320-0230102210130301-0221210122111221-0223000213112332-3000101033131221) |
| `routes.request_cookies_to_add.value` | [routes.request_cookies_to_add.value](resources--route--reference--group-001.md#canonical-1320130001223112-0300321201003011-0133200122112012-1122323331211122-2312301313110030-3302333312222232-3230011232203330-0132312313120223) |
| `routes.request_cookies_to_remove` | [routes.request_cookies_to_remove](resources--route--reference--group-001.md#canonical-3012111200222313-2033222131203000-1203100112010032-2202332323230230-2103103103201211-2032013220023311-3203211310233120-0331022213232012) |
| `routes.request_headers_to_add` | [routes.request_headers_to_add](resources--route--reference--group-002.md#canonical-1132203010013323-3020310323133333-0001320310213001-0031303131120100-0302301011123303-3313311112123211-3131210123100211-0312001013002032) |
| `routes.request_headers_to_add.append` | [routes.request_headers_to_add.append](resources--route--reference--group-002.md#canonical-1231110302120301-2213203232101212-2102222111312110-2313330002003003-1131323013322310-0222123110231133-3321112323113032-1320121020301333) |
| `routes.request_headers_to_add.name` | [routes.request_headers_to_add.name](resources--route--reference--group-002.md#canonical-0331233131112102-0233330301212023-0300020101030103-1332221133212303-0313200121312032-1100231232322130-2122313212321111-2323120023002301) |
| `routes.request_headers_to_add.secret_value` | [routes.request_headers_to_add.secret_value](resources--route--reference--group-002.md#canonical-1102202003112002-0303233213222201-0231003212122100-2133112130320232-0333332013010200-0201030201023002-3131112211310112-1130121110131012) |
| `routes.request_headers_to_add.secret_value.blindfold_secret_info` | [routes.request_headers_to_add.secret_value.blindfold_secret_info](resources--route--reference--group-002.md#canonical-0021322313122300-0321132321320300-0020313230232201-3331033012103101-3320011121023112-2111110232031330-1133102012111011-0031001110000201) |
| `routes.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [routes.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--route--reference--group-002.md#canonical-1022103121000233-2113100133221211-1230301211132003-0220321002133311-0011332011330122-2312112332110012-1102203200333223-1311110302223010) |
| `routes.request_headers_to_add.secret_value.blindfold_secret_info.location` | [routes.request_headers_to_add.secret_value.blindfold_secret_info.location](resources--route--reference--group-002.md#canonical-2202023223011203-0300033102221001-1212210230300210-3033221120233001-3030231223112300-1122020312102130-0121311122222223-1230003010320111) |
| `routes.request_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [routes.request_headers_to_add.secret_value.blindfold_secret_info.store_provider](resources--route--reference--group-002.md#canonical-3313021302012222-1332203013331332-0333002110322013-2311113122300030-3121121010100321-2313133002021320-2311131112332013-3312230011213323) |
| `routes.request_headers_to_add.secret_value.clear_secret_info` | [routes.request_headers_to_add.secret_value.clear_secret_info](resources--route--reference--group-002.md#canonical-2200231113331102-2320203122111000-2112201132211231-0110032320110223-3210033112102020-1202113221210201-0113120101300221-3231002333021303) |
| `routes.request_headers_to_add.secret_value.clear_secret_info.provider_ref` | [routes.request_headers_to_add.secret_value.clear_secret_info.provider_ref](resources--route--reference--group-002.md#canonical-1323023311113212-1330013111012322-3312322202033110-3201110322233301-0300012301333020-3130212310303220-2032203311220012-1012322333212001) |
| `routes.request_headers_to_add.secret_value.clear_secret_info.url` | [routes.request_headers_to_add.secret_value.clear_secret_info.url](resources--route--reference--group-002.md#canonical-3230023032133231-2233013032333331-0320332201203330-0020000010012131-0023000033002100-1121320030012231-2211313221230012-1210233011103113) |
| `routes.request_headers_to_add.value` | [routes.request_headers_to_add.value](resources--route--reference--group-002.md#canonical-0332212132012103-2100320021230031-1111011011101123-0331311213220302-2102100013302310-1003113213022310-2332222313311032-3323320120022121) |
| `routes.request_headers_to_remove` | [routes.request_headers_to_remove](resources--route--reference--group-001.md#canonical-2000132021120210-2123211030323333-3113030330320322-1310010033210211-0311313313231100-0100121312012110-0330313132201203-3222212320103103) |
| `routes.response_cookies_to_add` | [routes.response_cookies_to_add](resources--route--reference--group-002.md#canonical-0013333121022032-2312111032311133-0220132002003233-1123020133213121-1010302332221333-0213200110232021-2231311110031332-0101201033113330) |
| `routes.response_cookies_to_add.add_domain` | [routes.response_cookies_to_add.add_domain](resources--route--reference--group-002.md#canonical-1230301332112030-3323133021030012-0223202322110103-2033022130131332-2203220102022011-1103220023320032-2322321210330302-2122220323333131) |
| `routes.response_cookies_to_add.add_expiry` | [routes.response_cookies_to_add.add_expiry](resources--route--reference--group-002.md#canonical-2311312331013221-0010302332221202-0031211210303111-1021100023221321-3020313231031200-1310112130201302-0023201231232203-2103123231323020) |
| `routes.response_cookies_to_add.add_httponly` | [routes.response_cookies_to_add.add_httponly](resources--route--reference--group-002.md#canonical-3313121310131313-1200232200002320-3010311302311013-2122300200132233-3303133131303232-1130131310212211-0332311211221100-3201310101003223) |
| `routes.response_cookies_to_add.add_partitioned` | [routes.response_cookies_to_add.add_partitioned](resources--route--reference--group-002.md#canonical-3202102133211301-0232112122021123-2120010110011222-0030232300021313-1332102101311330-0023100303232031-2033120102210110-2230303000202322) |
| `routes.response_cookies_to_add.add_path` | [routes.response_cookies_to_add.add_path](resources--route--reference--group-002.md#canonical-2210220010112201-3300213012203231-0211333312112032-0222122321312012-0103121222331313-2230212311003123-3203230030130032-0311033212032123) |
| `routes.response_cookies_to_add.add_secure` | [routes.response_cookies_to_add.add_secure](resources--route--reference--group-002.md#canonical-2020200201210230-2101132301023322-3021310033200022-3012210121020011-3331330033132012-3133023231120033-0211312033023300-0332321220103220) |
| `routes.response_cookies_to_add.ignore_domain` | [routes.response_cookies_to_add.ignore_domain](resources--route--reference--group-002.md#canonical-0322223013231320-0112121303012213-2033222310022210-3220320033300333-3303222121233100-3011031111131100-1230212322003001-3333201222212323) |
| `routes.response_cookies_to_add.ignore_expiry` | [routes.response_cookies_to_add.ignore_expiry](resources--route--reference--group-002.md#canonical-0223311001021311-3201101231101121-2123120013223303-1122232102103232-3310102311000001-1320100011330130-3222121133222001-0100131101111233) |
| `routes.response_cookies_to_add.ignore_httponly` | [routes.response_cookies_to_add.ignore_httponly](resources--route--reference--group-002.md#canonical-1010300020330100-2003321200002321-2333210303203230-0313223021022011-0330330203323022-1302212210130212-2131312331320002-1120221002201310) |
| `routes.response_cookies_to_add.ignore_max_age` | [routes.response_cookies_to_add.ignore_max_age](resources--route--reference--group-002.md#canonical-3311130300323131-2302231310000213-0132310110303020-2302122010120332-0230332103213320-1230231211010212-1101110010012130-1121221312300001) |
| `routes.response_cookies_to_add.ignore_partitioned` | [routes.response_cookies_to_add.ignore_partitioned](resources--route--reference--group-002.md#canonical-1302212122310333-3231222220112232-0231100213102021-2102211133203000-2313310020310032-1303021323000312-2113202123021233-3320201131113302) |
| `routes.response_cookies_to_add.ignore_path` | [routes.response_cookies_to_add.ignore_path](resources--route--reference--group-002.md#canonical-3333332031312010-2100013032012213-3121312000030000-0130201031012103-2003011303111202-1301222101002223-1230310123330003-1320002102332312) |
| `routes.response_cookies_to_add.ignore_samesite` | [routes.response_cookies_to_add.ignore_samesite](resources--route--reference--group-002.md#canonical-1103213222110232-3130201200113123-1233330132111030-1310030303322311-1230231322033333-1320113203200021-1111120132330030-1323000230311022) |
| `routes.response_cookies_to_add.ignore_secure` | [routes.response_cookies_to_add.ignore_secure](resources--route--reference--group-002.md#canonical-3020212013112210-2331303102032023-3102132112200031-3233311121300111-3032111233223102-2331001120332101-2221132202000111-0000310020011030) |
| `routes.response_cookies_to_add.ignore_value` | [routes.response_cookies_to_add.ignore_value](resources--route--reference--group-002.md#canonical-3103133231111101-0121120331100222-0221023311211012-1231330311332331-1131002321023101-1110330321111032-1301211310231300-2321100301301000) |
| `routes.response_cookies_to_add.max_age_value` | [routes.response_cookies_to_add.max_age_value](resources--route--reference--group-002.md#canonical-2322130201023212-2301122001100031-2103100220020101-1212111300211311-3132110332002330-1013210333301322-2201122200003032-0213303220022122) |
| `routes.response_cookies_to_add.name` | [routes.response_cookies_to_add.name](resources--route--reference--group-002.md#canonical-2202011030200302-0022031032312002-2100203101101030-1011123220031213-3013312131203122-1120323322223113-2333111232122013-1010113232222132) |
| `routes.response_cookies_to_add.overwrite` | [routes.response_cookies_to_add.overwrite](resources--route--reference--group-002.md#canonical-2131230100112111-1223133320131030-3220222201212303-1231103121113322-1132312112213202-1003020200203002-2001311302231003-1301021011021302) |
| `routes.response_cookies_to_add.samesite_lax` | [routes.response_cookies_to_add.samesite_lax](resources--route--reference--group-002.md#canonical-0112001302022023-2323303332111330-0002201230102212-1320000221312213-2011131012232202-2010010010101311-1212121230220012-1223212110131333) |
| `routes.response_cookies_to_add.samesite_none` | [routes.response_cookies_to_add.samesite_none](resources--route--reference--group-002.md#canonical-3321302310202120-1122120133321121-1123033221223100-1122123010020133-2213003320313332-2201022130213003-1123013230223303-3222332211120031) |
| `routes.response_cookies_to_add.samesite_strict` | [routes.response_cookies_to_add.samesite_strict](resources--route--reference--group-002.md#canonical-0130210212200013-3201301113032300-1323301330210223-1323001233330321-1311113121231213-1110100031032202-2020112132000122-2132121122333102) |
| `routes.response_cookies_to_add.secret_value` | [routes.response_cookies_to_add.secret_value](resources--route--reference--group-002.md#canonical-3333012202231210-1213231003113113-3013110202200021-3130213311331223-1302020012113132-1201311223130132-2212110302232231-0222130212021331) |
| `routes.response_cookies_to_add.secret_value.blindfold_secret_info` | [routes.response_cookies_to_add.secret_value.blindfold_secret_info](resources--route--reference--group-002.md#canonical-0303113223102022-1013222213232312-0033331310201200-0033112110002322-3223123233032110-1200133110123002-0012323302012302-3130221113333333) |
| `routes.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` | [routes.response_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--route--reference--group-002.md#canonical-1320200232112221-0310101311132113-2322003033123213-0322221132003033-0123112030201200-1213120122120301-1033031003103310-3102020010312201) |
| `routes.response_cookies_to_add.secret_value.blindfold_secret_info.location` | [routes.response_cookies_to_add.secret_value.blindfold_secret_info.location](resources--route--reference--group-002.md#canonical-3332321000030013-0030022113000333-1333311132312303-0232023031120032-1231311100001033-1033002232322033-2130113231021213-2200122120313132) |
| `routes.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider` | [routes.response_cookies_to_add.secret_value.blindfold_secret_info.store_provider](resources--route--reference--group-002.md#canonical-0203311012021020-1201000121031203-1200332103120323-1203232112100133-2100211210103213-3032021203110000-3331122100100030-0010002102030023) |
| `routes.response_cookies_to_add.secret_value.clear_secret_info` | [routes.response_cookies_to_add.secret_value.clear_secret_info](resources--route--reference--group-002.md#canonical-0100330301122212-0332002202023200-1120220132321303-1310102230331011-1212333232121103-2130322020000333-2301100033200120-3122301303010320) |
| `routes.response_cookies_to_add.secret_value.clear_secret_info.provider_ref` | [routes.response_cookies_to_add.secret_value.clear_secret_info.provider_ref](resources--route--reference--group-002.md#canonical-0222311132220211-0323221102202232-2032203211123012-1220013113333130-1313213133223213-3103003212332220-2030111111332122-1020123323102010) |
| `routes.response_cookies_to_add.secret_value.clear_secret_info.url` | [routes.response_cookies_to_add.secret_value.clear_secret_info.url](resources--route--reference--group-002.md#canonical-1211100222030223-0303102213221322-1002013303011131-1033012210112222-0212001320021131-3201011233003132-1131031300002120-2300102103302212) |
| `routes.response_cookies_to_add.value` | [routes.response_cookies_to_add.value](resources--route--reference--group-002.md#canonical-1113130012332221-0201333220100130-3031212102223232-0031332010031031-3310211011231132-2231201321303003-0100221120220302-2210332332030133) |
| `routes.response_cookies_to_remove` | [routes.response_cookies_to_remove](resources--route--reference--group-001.md#canonical-3000123122132002-3003130020302303-1110003203300213-2120031120103322-3100221011232302-1212103211122100-1223133201220003-2100312311201011) |
| `routes.response_headers_to_add` | [routes.response_headers_to_add](resources--route--reference--group-002.md#canonical-3012321031120202-3202102101021221-0203033222130113-3111023000220022-1031202222231233-0033100130113123-1033113333112101-2231222023203000) |
| `routes.response_headers_to_add.append` | [routes.response_headers_to_add.append](resources--route--reference--group-002.md#canonical-0323103003213112-2030130102321021-1012120012203301-2300223320203020-0022202013333102-0000201032211230-2033100102312111-2222322003021312) |
| `routes.response_headers_to_add.name` | [routes.response_headers_to_add.name](resources--route--reference--group-002.md#canonical-1121120323000111-2321222301100202-0332221131102001-0132122300111123-0131302113302303-2332201111231203-3332030102010331-1321103131013100) |
| `routes.response_headers_to_add.secret_value` | [routes.response_headers_to_add.secret_value](resources--route--reference--group-002.md#canonical-0000022131220300-0121321101113311-3213311012133020-1033322222111313-3022201321012111-0202323131303130-0323232032320023-2322012002000131) |
| `routes.response_headers_to_add.secret_value.blindfold_secret_info` | [routes.response_headers_to_add.secret_value.blindfold_secret_info](resources--route--reference--group-003.md#canonical-1332301321300232-0033033312301202-3300011230112200-3322100333133323-3122312110121023-1323311120232300-2333020000202212-0301120130332001) |
| `routes.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` | [routes.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider](resources--route--reference--group-003.md#canonical-0133130110130102-3311201331123012-1021303033003210-2022323123022233-3103003303333220-0231233231211212-2012123121323222-1103120300310301) |
| `routes.response_headers_to_add.secret_value.blindfold_secret_info.location` | [routes.response_headers_to_add.secret_value.blindfold_secret_info.location](resources--route--reference--group-003.md#canonical-3300223022133313-0000020321001221-1032203231013200-2102212221113210-2101231013130322-2301013002332322-1032220211211321-3021213302331013) |
| `routes.response_headers_to_add.secret_value.blindfold_secret_info.store_provider` | [routes.response_headers_to_add.secret_value.blindfold_secret_info.store_provider](resources--route--reference--group-003.md#canonical-2302200200012210-2113003013310212-0133000132100012-1031033101020122-3200323021312002-3031122220023131-2132230200331132-2333230100231022) |
| `routes.response_headers_to_add.secret_value.clear_secret_info` | [routes.response_headers_to_add.secret_value.clear_secret_info](resources--route--reference--group-003.md#canonical-1331312122013013-1330201332131122-0222020212331112-2022112310002112-0130110322213010-2101223300121211-2002213200103203-3301003332123232) |
| `routes.response_headers_to_add.secret_value.clear_secret_info.provider_ref` | [routes.response_headers_to_add.secret_value.clear_secret_info.provider_ref](resources--route--reference--group-003.md#canonical-0131001300012132-3022201030022021-0123312023221113-2112131110200210-2003303100011121-3312122011002020-0032220030232300-2030100013232311) |
| `routes.response_headers_to_add.secret_value.clear_secret_info.url` | [routes.response_headers_to_add.secret_value.clear_secret_info.url](resources--route--reference--group-003.md#canonical-0223010103300110-0211201203103003-2012121323000313-3013103331332330-0211222120010132-2001231301011120-2211111322231213-0002220013303111) |
| `routes.response_headers_to_add.value` | [routes.response_headers_to_add.value](resources--route--reference--group-002.md#canonical-0331310331200301-1301313020030203-2131321332222313-1312210200103233-0321133212123303-2103101120133330-2213023130013113-0202322201100130) |
| `routes.response_headers_to_remove` | [routes.response_headers_to_remove](resources--route--reference--group-001.md#canonical-2021321232003032-1020032131113312-1010023032012230-3231233031203010-2011210022122030-1211132113222310-1020033202001210-3021131321223301) |
| `routes.route_destination` | [routes.route_destination](resources--route--reference--group-003.md#canonical-3202133011020222-0010230313321311-0131111301313033-1310221232100321-0213231323010232-3301302020002202-3103333310010032-0320130202302201) |
| `routes.route_destination.auto_host_rewrite` | [routes.route_destination.auto_host_rewrite](resources--route--reference--group-003.md#canonical-1020132210222313-3231102003332113-3110211221330212-0030001130323103-1011302201323323-2212011001130231-2231210021230010-3221133321011012) |
| `routes.route_destination.buffer_policy` | [routes.route_destination.buffer_policy](resources--route--reference--group-003.md#canonical-3332132131303110-3303021323103203-0322131323020321-0231302332022120-1022312330333011-0232101300120021-2113302103113023-1101233121302130) |
| `routes.route_destination.buffer_policy.disabled` | [routes.route_destination.buffer_policy.disabled](resources--route--reference--group-003.md#canonical-0330123133230013-1002310231131330-3110113001002321-2113303001101220-1101201022212302-1100300121113233-3011130113322220-2322310221113322) |
| `routes.route_destination.buffer_policy.max_request_bytes` | [routes.route_destination.buffer_policy.max_request_bytes](resources--route--reference--group-003.md#canonical-0133300333132100-1212310332103320-2300201203201202-1312231132213131-3111112232011112-0032203230311202-3212213102201010-1123200102023032) |
| `routes.route_destination.cors_policy` | [routes.route_destination.cors_policy](resources--route--reference--group-003.md#canonical-1002001011310010-1321122311202020-2313123222111310-3320112133233220-2230111312030313-1101313021102132-0133231303220330-3031002200232210) |
| `routes.route_destination.cors_policy.allow_credentials` | [routes.route_destination.cors_policy.allow_credentials](resources--route--reference--group-003.md#canonical-3323033113301101-0302003112302303-2131012202212330-2023310300032320-1220202323210123-2000202301120212-2322221300120022-3101211022202133) |
| `routes.route_destination.cors_policy.allow_headers` | [routes.route_destination.cors_policy.allow_headers](resources--route--reference--group-003.md#canonical-3220220220033200-2122010221110132-3233031032133231-0300322132303113-0131030302201232-1133112310311110-3022031332000103-3200131331022203) |
| `routes.route_destination.cors_policy.allow_methods` | [routes.route_destination.cors_policy.allow_methods](resources--route--reference--group-003.md#canonical-0230123213310113-3311223003320010-1000303203333030-1233202232313332-1331033031322313-0300103221131030-3002030102123100-1103102111111130) |
| `routes.route_destination.cors_policy.allow_origin` | [routes.route_destination.cors_policy.allow_origin](resources--route--reference--group-003.md#canonical-1110121302020313-0112330301222121-1230203133301011-2203333301323330-0303122302302000-0001122220100130-0013203230003312-0222030020233111) |
| `routes.route_destination.cors_policy.allow_origin_regex` | [routes.route_destination.cors_policy.allow_origin_regex](resources--route--reference--group-003.md#canonical-1111200100321101-0111023300010103-1003031313231313-3133312211302100-3030212200133132-3133112133011031-0131102031333022-0200200103201032) |
| `routes.route_destination.cors_policy.disabled` | [routes.route_destination.cors_policy.disabled](resources--route--reference--group-003.md#canonical-1120312123121223-1212222203002303-2121032100112233-0313100131003030-2330233101131213-2001200210232211-0032010010233212-0001220321323323) |
| `routes.route_destination.cors_policy.expose_headers` | [routes.route_destination.cors_policy.expose_headers](resources--route--reference--group-003.md#canonical-2223112113323322-3120132232203030-3213313322322010-0001213132102302-2123231113023320-3122300021111112-0222212020200001-3131333131202021) |
| `routes.route_destination.cors_policy.maximum_age` | [routes.route_destination.cors_policy.maximum_age](resources--route--reference--group-003.md#canonical-3331311321130311-0312233203312001-3013022312322103-1320132021201011-3231030110121121-2230133322202300-0121221033211130-1311310200220221) |
| `routes.route_destination.csrf_policy` | [routes.route_destination.csrf_policy](resources--route--reference--group-003.md#canonical-1201101113101320-2323301123110211-1200230130122233-3301301233302201-3021210021301132-1021301001313121-2100102101223223-2233302311020121) |
| `routes.route_destination.csrf_policy.all_load_balancer_domains` | [routes.route_destination.csrf_policy.all_load_balancer_domains](resources--route--reference--group-003.md#canonical-2111101213212300-0130030103322333-1332123223120002-3111103133313301-1123113333333313-1012130102030020-1113003312222022-1010012231312002) |
| `routes.route_destination.csrf_policy.custom_domain_list` | [routes.route_destination.csrf_policy.custom_domain_list](resources--route--reference--group-003.md#canonical-2011200113112300-1022220223302031-1200333312000212-3333000112000233-0200131212300031-1033212130200112-2232322101200303-1010012132201200) |
| `routes.route_destination.csrf_policy.custom_domain_list.domains` | [routes.route_destination.csrf_policy.custom_domain_list.domains](resources--route--reference--group-003.md#canonical-2113012313133103-1230112103030103-3031012103330221-0103200013312213-3202332010101333-0011232323032230-0333110031110211-2211010302131233) |
| `routes.route_destination.csrf_policy.disabled` | [routes.route_destination.csrf_policy.disabled](resources--route--reference--group-003.md#canonical-2130231331111223-1212202021021111-2310102021210323-2210221111033232-0101302130031312-1221302112222322-1333032313010221-0301310213331023) |
| `routes.route_destination.destinations` | [routes.route_destination.destinations](resources--route--reference--group-003.md#canonical-1112021013231321-0113122210020133-0133302023010120-1210312112321010-0022033012131121-3201100010031021-0301121112201221-3033121132300102) |
| `routes.route_destination.destinations.cluster` | [routes.route_destination.destinations.cluster](resources--route--reference--group-003.md#canonical-3112032321222223-1010021133011211-1113310223033131-3132130010120322-2003312221021003-0013003132130320-0033112213100023-1220313200230202) |
| `routes.route_destination.destinations.cluster.kind` | [routes.route_destination.destinations.cluster.kind](resources--route--reference--group-003.md#canonical-2023211112200200-0100113331330221-2222103221003131-2211321032130321-1300230223123331-3230202121220200-3002221233122210-1131030121010120) |
| `routes.route_destination.destinations.cluster.name` | [routes.route_destination.destinations.cluster.name](resources--route--reference--group-003.md#canonical-3100010121222333-3010001321212223-0213030111102020-3023213200303320-1120123120212221-1211333332023103-1022102322102222-1101202031332300) |
| `routes.route_destination.destinations.cluster.namespace` | [routes.route_destination.destinations.cluster.namespace](resources--route--reference--group-003.md#canonical-1211111020320010-2311111221323230-2320033033230113-1221220110023003-2112013130310000-1221303201312100-2210203201031131-3333022323130130) |
| `routes.route_destination.destinations.cluster.tenant` | [routes.route_destination.destinations.cluster.tenant](resources--route--reference--group-003.md#canonical-2331211232213031-0302233020211102-1332301332000233-3021222302012223-2000022010032202-3210120323131330-2212112213022330-1122223011100310) |
| `routes.route_destination.destinations.cluster.uid` | [routes.route_destination.destinations.cluster.uid](resources--route--reference--group-003.md#canonical-3321123113222022-2101000101122233-3110020210233011-1002101313223231-1202103133203322-3030223230221110-2120330332110203-0011013032100200) |
| `routes.route_destination.destinations.endpoint_subsets` | [routes.route_destination.destinations.endpoint_subsets](resources--route--reference--group-003.md#canonical-3023111002013112-0331013020301003-3212323002011002-1112223230232101-3133001122203222-0100001203201100-0003120121303232-1223203221010123) |
| `routes.route_destination.destinations.priority` | [routes.route_destination.destinations.priority](resources--route--reference--group-003.md#canonical-0231211220302310-3022333121201331-1003013221103033-0002202231313023-0211010332033010-3022301211332031-1301122300300022-2033121310233002) |
| `routes.route_destination.destinations.weight` | [routes.route_destination.destinations.weight](resources--route--reference--group-003.md#canonical-2211122320202232-0233131203103100-0212332313122322-0132220101021033-0130223203322103-2020101001221102-0310002000023120-3303310333030233) |
| `routes.route_destination.do_not_retract_cluster` | [routes.route_destination.do_not_retract_cluster](resources--route--reference--group-003.md#canonical-3133130213210122-1031303100213011-1111022303303330-0312012102221332-0001232112220322-0332103332122131-3133202220002222-1021130122001220) |
| `routes.route_destination.endpoint_subsets` | [routes.route_destination.endpoint_subsets](resources--route--reference--group-003.md#canonical-2133132312333130-0300231121312103-1332221322003010-2111323330203031-2131323312000232-0032222122112032-2211313312103010-3211322011201233) |
| `routes.route_destination.hash_policy` | [routes.route_destination.hash_policy](resources--route--reference--group-003.md#canonical-1013123110222103-0013032232130130-0302013230222213-0022333002111232-2111033110301233-2212332330201011-0003023310311103-2313110000010202) |
| `routes.route_destination.hash_policy.cookie` | [routes.route_destination.hash_policy.cookie](resources--route--reference--group-003.md#canonical-0121132100220221-0023323022210331-2200322300133131-3203100112223213-0020321121012322-0000303122301332-2113212000022320-2130001121312221) |
| `routes.route_destination.hash_policy.cookie.add_httponly` | [routes.route_destination.hash_policy.cookie.add_httponly](resources--route--reference--group-003.md#canonical-3212231201012130-1013300112120131-2112100201132112-1100303313013020-1202202102233233-0310023302032233-1031101300032220-2030121030223001) |
| `routes.route_destination.hash_policy.cookie.add_secure` | [routes.route_destination.hash_policy.cookie.add_secure](resources--route--reference--group-003.md#canonical-2312212130233131-3221331103033330-3023122000300132-2012111202331322-0223120102001003-0112331302010013-1112201011003033-0301333123322300) |
| `routes.route_destination.hash_policy.cookie.ignore_httponly` | [routes.route_destination.hash_policy.cookie.ignore_httponly](resources--route--reference--group-003.md#canonical-1313012031333203-1202020032331201-0301033201322333-3231001301303313-3001011323201200-0011100313023221-3303322101101012-0320321113230203) |
| `routes.route_destination.hash_policy.cookie.ignore_samesite` | [routes.route_destination.hash_policy.cookie.ignore_samesite](resources--route--reference--group-003.md#canonical-1012302222333022-1131133020203313-2232202232122013-1000020011323330-1112331233330311-0013231230132112-3101322231300311-2003101212321101) |
| `routes.route_destination.hash_policy.cookie.ignore_secure` | [routes.route_destination.hash_policy.cookie.ignore_secure](resources--route--reference--group-003.md#canonical-0000320013100001-3302102311111101-2111213111133302-2303200023233131-0230001001211200-2100030310112220-3311010000220332-2030201312100113) |
| `routes.route_destination.hash_policy.cookie.name` | [routes.route_destination.hash_policy.cookie.name](resources--route--reference--group-003.md#canonical-2302030300221023-2331030021033101-1312032003011013-0201112131021130-1021222032000212-2110110122322300-1312101322220013-3120112100120302) |
| `routes.route_destination.hash_policy.cookie.path` | [routes.route_destination.hash_policy.cookie.path](resources--route--reference--group-003.md#canonical-2202212010101230-3113131123302322-0003213230130300-0133303200122310-0322012003221001-0311112111213110-3213033230103321-2023203333023210) |
| `routes.route_destination.hash_policy.cookie.samesite_lax` | [routes.route_destination.hash_policy.cookie.samesite_lax](resources--route--reference--group-003.md#canonical-2220033133323022-3322202332100320-2113110323303100-3301321122303303-3321212030133232-2102332122321132-1021131012021213-3302332002123220) |
| `routes.route_destination.hash_policy.cookie.samesite_none` | [routes.route_destination.hash_policy.cookie.samesite_none](resources--route--reference--group-003.md#canonical-3132323303320010-2000331232023100-1232330132301001-0330031100201322-3310230111121302-0003330111320100-1221301031020112-0313031212230021) |
| `routes.route_destination.hash_policy.cookie.samesite_strict` | [routes.route_destination.hash_policy.cookie.samesite_strict](resources--route--reference--group-003.md#canonical-2001031230200211-1300113322313300-0111300330102032-0022121202310311-3003200013310030-2100313223231131-2230222032022010-1322002120123202) |
| `routes.route_destination.hash_policy.cookie.ttl` | [routes.route_destination.hash_policy.cookie.ttl](resources--route--reference--group-003.md#canonical-0333302131313121-3003220321013302-0302023310313033-3301133120302022-0013212000332001-2002310323011023-1230332031213232-1113223213333230) |
| `routes.route_destination.hash_policy.header_name` | [routes.route_destination.hash_policy.header_name](resources--route--reference--group-003.md#canonical-3030332102111320-3020113232322212-3101233203311223-1010021321003200-3322032323321223-0033112123323221-0122321001230102-0220302031201033) |
| `routes.route_destination.hash_policy.source_ip` | [routes.route_destination.hash_policy.source_ip](resources--route--reference--group-003.md#canonical-3222200112323331-0132001111011033-1031020113001202-2030021003001013-1101222102312200-1302103300113332-0203211003123333-1310233331023132) |
| `routes.route_destination.hash_policy.terminal` | [routes.route_destination.hash_policy.terminal](resources--route--reference--group-003.md#canonical-3000123202120111-1230023210113110-3023021023030021-2013210011302133-3220320302311300-1302311102233113-3121103003232213-2311122211000103) |
| `routes.route_destination.host_rewrite` | [routes.route_destination.host_rewrite](resources--route--reference--group-003.md#canonical-0213120322322211-2323332001212230-3002223111223201-1101113121101231-3333221201301321-3200122111111020-3322032111111131-0232230332313301) |
| `routes.route_destination.mirror_policy` | [routes.route_destination.mirror_policy](resources--route--reference--group-003.md#canonical-3103103220332022-3301230001230003-3012303123110101-2223301012122202-2211031023332111-2001013323223200-3231310322021311-1122013113121022) |
| `routes.route_destination.mirror_policy.cluster` | [routes.route_destination.mirror_policy.cluster](resources--route--reference--group-003.md#canonical-1030111212111010-3323120011231231-2133313023011112-1002310313231321-2221033103313322-1313011222202320-3133221322313321-1222101001303022) |
| `routes.route_destination.mirror_policy.cluster.kind` | [routes.route_destination.mirror_policy.cluster.kind](resources--route--reference--group-003.md#canonical-2221031211000201-1220033013223111-2302031011231120-3032132022122322-0232322232131131-0132132031332103-2301332122020330-3220112132002003) |
| `routes.route_destination.mirror_policy.cluster.name` | [routes.route_destination.mirror_policy.cluster.name](resources--route--reference--group-003.md#canonical-0203122202231013-3133012020002100-1311330321123202-3323010023223232-2133333123032011-0203313200303213-1123203013220030-1112311010300323) |
| `routes.route_destination.mirror_policy.cluster.namespace` | [routes.route_destination.mirror_policy.cluster.namespace](resources--route--reference--group-003.md#canonical-1021310020012002-2010233013322010-3203103112022112-0330013323120320-3132101211101130-2220110302201021-1120013110000313-0000230203122121) |
| `routes.route_destination.mirror_policy.cluster.tenant` | [routes.route_destination.mirror_policy.cluster.tenant](resources--route--reference--group-003.md#canonical-3213321013010012-0310230311312222-0303020231223313-2130300121232001-0313232023213031-2010132203100123-3130300003033200-0322103220333201) |
| `routes.route_destination.mirror_policy.cluster.uid` | [routes.route_destination.mirror_policy.cluster.uid](resources--route--reference--group-003.md#canonical-0230013212020031-1321223123332333-1031033101222221-2110030220102210-2001320220133102-1220210013100113-0201122032321313-1101130313122111) |
| `routes.route_destination.mirror_policy.percent` | [routes.route_destination.mirror_policy.percent](resources--route--reference--group-003.md#canonical-1221011213212311-0111213320333033-0200230201121021-1331233132231010-2103301020223121-0131121322020310-3113323201321122-0122000203332030) |
| `routes.route_destination.mirror_policy.percent.denominator` | [routes.route_destination.mirror_policy.percent.denominator](resources--route--reference--group-003.md#canonical-0322132122101310-3332333202133030-0000301300122213-3322300211233313-3302232213002003-0313301301320320-1320313010033320-0010221133032321) |
| `routes.route_destination.mirror_policy.percent.numerator` | [routes.route_destination.mirror_policy.percent.numerator](resources--route--reference--group-003.md#canonical-2120233110230332-0132030231313322-2002033300333200-3013302331202211-0133023223012133-0233032230132023-2133320030221022-1300010210110210) |
| `routes.route_destination.prefix_rewrite` | [routes.route_destination.prefix_rewrite](resources--route--reference--group-003.md#canonical-3300231230103211-3131012203302321-2231333032013023-2032201232302133-1110332030311313-2233313102012022-1133202032013231-1003031021221222) |
| `routes.route_destination.priority` | [routes.route_destination.priority](resources--route--reference--group-003.md#canonical-2332133011003023-2203210112332003-2220001320113301-2011311223221131-1003131023130120-1303012222331311-0112223313131223-3011301200031022) |
| `routes.route_destination.query_params` | [routes.route_destination.query_params](resources--route--reference--group-003.md#canonical-3322102330231331-3300311203201320-0000122101103111-0303021333331323-0110130202120013-1000122210000132-1311002111001011-2011230101312303) |
| `routes.route_destination.query_params.remove_all_params` | [routes.route_destination.query_params.remove_all_params](resources--route--reference--group-003.md#canonical-2322332332011010-0233233011332323-2123011002233030-0130232223102323-3322121130223020-0320133222311302-0101120212121121-0202330323113232) |
| `routes.route_destination.query_params.replace_params` | [routes.route_destination.query_params.replace_params](resources--route--reference--group-003.md#canonical-2310033101202202-0311102331302133-0032200113103022-1322210031100130-1031022113100101-0023233313122120-3011011001112213-0030102233031330) |
| `routes.route_destination.query_params.retain_all_params` | [routes.route_destination.query_params.retain_all_params](resources--route--reference--group-003.md#canonical-0311130221111002-3330300100223330-0211020210230110-0020221103200130-3123102032212302-0333311320201113-0033322201013121-2032330220201100) |
| `routes.route_destination.regex_rewrite` | [routes.route_destination.regex_rewrite](resources--route--reference--group-003.md#canonical-1013121130310212-2012233002003333-0120310130032131-2002312220330201-0333322120201030-2113031003131210-0013120202131000-3232130213013232) |
| `routes.route_destination.regex_rewrite.pattern` | [routes.route_destination.regex_rewrite.pattern](resources--route--reference--group-003.md#canonical-0311210110212320-0230203213210212-3310002103200221-3332232332322132-1203123301012110-0011212131023023-1223023120320113-0321331210203331) |
| `routes.route_destination.regex_rewrite.substitution` | [routes.route_destination.regex_rewrite.substitution](resources--route--reference--group-003.md#canonical-3132122230211331-0031333022032033-2113100322230131-3210001100100032-2220222021132230-2100332022103211-1123302122113223-1011103302200213) |
| `routes.route_destination.retract_cluster` | [routes.route_destination.retract_cluster](resources--route--reference--group-003.md#canonical-3133033201021013-1220100131310331-1023103000323112-0322203333003220-0111303312211200-2012200320231002-2331332302211033-1110213301231210) |
| `routes.route_destination.retry_policy` | [routes.route_destination.retry_policy](resources--route--reference--group-003.md#canonical-0332120203322330-0231001201132303-3131103033201212-1120133013223103-1202022020111333-0121203233322211-1110030301201322-0303332300101330) |
| `routes.route_destination.retry_policy.back_off` | [routes.route_destination.retry_policy.back_off](resources--route--reference--group-003.md#canonical-2322231011021330-2102132021303133-3132030223101012-2210121213010213-0200032120022221-0031302302230023-3103312203222103-2020322003020233) |
| `routes.route_destination.retry_policy.back_off.base_interval` | [routes.route_destination.retry_policy.back_off.base_interval](resources--route--reference--group-003.md#canonical-2213302023003231-3230112012222031-1101221232222011-1113223022002033-0201221202333130-3010212303332231-3311102330103111-2131202232330231) |
| `routes.route_destination.retry_policy.back_off.max_interval` | [routes.route_destination.retry_policy.back_off.max_interval](resources--route--reference--group-003.md#canonical-0312303303212133-0310220220200023-0310123303300013-2323231000102133-2113013320110101-0313111321112222-1222121111103131-2111301021121001) |
| `routes.route_destination.retry_policy.num_retries` | [routes.route_destination.retry_policy.num_retries](resources--route--reference--group-003.md#canonical-3330313310300130-3312310312010001-3130022311003221-0230220102233132-2323201321202100-2210223100300010-2230203102012202-3233211212301120) |
| `routes.route_destination.retry_policy.per_try_timeout` | [routes.route_destination.retry_policy.per_try_timeout](resources--route--reference--group-003.md#canonical-1001121203001131-0211103020313001-0301332131232112-1131011033213310-1001332013303300-2031301013131130-0111110211103001-0133002232303101) |
| `routes.route_destination.retry_policy.retriable_status_codes` | [routes.route_destination.retry_policy.retriable_status_codes](resources--route--reference--group-003.md#canonical-2203021300100032-3203111312120103-1120203200023132-3032002101313003-1112120230223212-0333313011200302-0232121123231202-3333003103222300) |
| `routes.route_destination.retry_policy.retry_condition` | [routes.route_destination.retry_policy.retry_condition](resources--route--reference--group-003.md#canonical-0212201223233101-3232113001000331-1101210313232012-3213203123232000-2221212111313322-3322311020331033-1010013330220203-1330311101111131) |
| `routes.route_destination.spdy_config` | [routes.route_destination.spdy_config](resources--route--reference--group-003.md#canonical-0030003212102300-1303002203222302-0012302330132121-3322003213202013-2302221310210203-3332130210330001-0312101133322231-0012213100003100) |
| `routes.route_destination.spdy_config.use_spdy` | [routes.route_destination.spdy_config.use_spdy](resources--route--reference--group-003.md#canonical-3123311300030233-2123110021122301-2012210221223130-3012121032221111-3330210021300003-0021203222001111-3100001231322013-2303320203332323) |
| `routes.route_destination.timeout` | [routes.route_destination.timeout](resources--route--reference--group-003.md#canonical-0023211312233333-0003323002031112-0021112010003130-1033120010002230-1222210230321312-1113203201320201-1202223102012200-3130121332201223) |
| `routes.route_destination.web_socket_config` | [routes.route_destination.web_socket_config](resources--route--reference--group-003.md#canonical-3101002232312011-1211310201000221-2323130121302321-0210310032312112-0020010111312100-1232122231102203-2230232321313322-3312113132221333) |
| `routes.route_destination.web_socket_config.use_websocket` | [routes.route_destination.web_socket_config.use_websocket](resources--route--reference--group-003.md#canonical-1113111201133330-1023102023302213-3112232113230133-2211301220013013-3222101020120300-3322023031001122-3203323200221220-0120103221113331) |
| `routes.route_direct_response` | [routes.route_direct_response](resources--route--reference--group-003.md#canonical-2302300123131210-3213103312320230-3310001122100100-2321123102113011-1202212130102111-2103211033323100-0000112122332101-2001013102102011) |
| `routes.route_direct_response.response_body_encoded` | [routes.route_direct_response.response_body_encoded](resources--route--reference--group-003.md#canonical-0222302303020322-3312111332232133-3113311202113321-0301303122100303-2101201311310001-3223302220120231-3233222213233331-1011112302133331) |
| `routes.route_direct_response.response_code` | [routes.route_direct_response.response_code](resources--route--reference--group-003.md#canonical-1121110031003012-0213231211323312-2123332020010221-0310030111300302-1300120221030231-0121002300230111-2030030021333322-0311203130333202) |
| `routes.route_redirect` | [routes.route_redirect](resources--route--reference--group-003.md#canonical-2133332303131212-2113132023333213-3232122202313313-0300313003133233-3011011131333210-0010102301013122-2011110201201133-2223120002333222) |
| `routes.route_redirect.host_redirect` | [routes.route_redirect.host_redirect](resources--route--reference--group-003.md#canonical-0101010203031003-0230330123212111-0122102202320322-0030210232212021-1120120032332120-3030210302022322-3033101303213312-3031021030031331) |
| `routes.route_redirect.path_redirect` | [routes.route_redirect.path_redirect](resources--route--reference--group-003.md#canonical-2321122203200303-2303333300122111-0011200122023233-1303101211302100-2302021111020300-3203322033120320-1100123302120223-2112200202330333) |
| `routes.route_redirect.prefix_rewrite` | [routes.route_redirect.prefix_rewrite](resources--route--reference--group-003.md#canonical-2133231130102301-0030333130120332-1203322110022230-2122132321233330-3010330213011322-2303233302103310-3303323200123100-0121211201310310) |
| `routes.route_redirect.proto_redirect` | [routes.route_redirect.proto_redirect](resources--route--reference--group-003.md#canonical-2111212001231200-3011023330002222-3100002230103332-3303130312113302-1131013220011203-3030101113312303-2333201203000013-2031123132031023) |
| `routes.route_redirect.remove_all_params` | [routes.route_redirect.remove_all_params](resources--route--reference--group-003.md#canonical-1130112110102132-0323112211230000-1232021131331331-3101130220131031-0111301330030320-0100103011021211-3213122000121013-1230222121201101) |
| `routes.route_redirect.replace_params` | [routes.route_redirect.replace_params](resources--route--reference--group-003.md#canonical-1310012213020231-0021013011110001-0001123211220323-1210231031011330-0323002121101322-1120203232133303-0102213133121203-3011222130132001) |
| `routes.route_redirect.response_code` | [routes.route_redirect.response_code](resources--route--reference--group-003.md#canonical-2130203220230202-0120301232033223-1203032010212222-1112301300010311-2012223123120031-3220330300000003-0303030302003330-0200322111102321) |
| `routes.route_redirect.retain_all_params` | [routes.route_redirect.retain_all_params](resources--route--reference--group-003.md#canonical-3233111231320202-3103213010231200-1000333130330113-0120203100231100-0322122110033111-3021201220220321-3011313123021331-3033333032201200) |
| `routes.service_policy` | [routes.service_policy](resources--route--reference--group-003.md#canonical-2023323222222023-1122101000020101-0133022030032131-3110002210200313-0221223013131021-1331312110320202-3111210010012031-2010022313202003) |
| `routes.service_policy.disable_spec` | [routes.service_policy.disable_spec](resources--route--reference--group-003.md#canonical-0212233320321030-1201131112030001-3332030223133030-2022210001322213-0213221132002203-1020301233333320-2032231013032230-2120121200230223) |
| `routes.waf_exclusion_policy` | [routes.waf_exclusion_policy](resources--route--reference--group-003.md#canonical-2002013313001200-2211230001233100-3213002023010302-0003113113213113-1220101021221323-2310310033302131-2033311313033122-3303301021200331) |
| `routes.waf_exclusion_policy.name` | [routes.waf_exclusion_policy.name](resources--route--reference--group-003.md#canonical-0323333110321132-1321231002100031-0332113312313302-2121031002221313-1132233200120001-2311323000001222-0120011311300232-2131030120330001) |
| `routes.waf_exclusion_policy.namespace` | [routes.waf_exclusion_policy.namespace](resources--route--reference--group-003.md#canonical-3320223201311312-1230332200221132-2023211032030213-2311120001232102-1330201233101010-3312101333222000-0032101301312212-3231202330011003) |
| `routes.waf_exclusion_policy.tenant` | [routes.waf_exclusion_policy.tenant](resources--route--reference--group-003.md#canonical-0112303323221002-0112010132321100-2020213300131313-2313332203100003-1110211220313310-3011301330230102-0320120110211120-0220022210011311) |
| `routes.waf_type` | [routes.waf_type](resources--route--reference--group-003.md#canonical-2001321113131313-1313201110111212-3122331232022222-2331202122321020-0221120200230002-3200330210012331-0322102302101033-3110033121012002) |
| `routes.waf_type.app_firewall` | [routes.waf_type.app_firewall](resources--route--reference--group-003.md#canonical-2231021231220033-0233233122220013-1111123320222301-3022211031003332-0300100210121302-1331002312111300-3302022222231002-0021032131203131) |
| `routes.waf_type.app_firewall.app_firewall` | [routes.waf_type.app_firewall.app_firewall](resources--route--reference--group-003.md#canonical-2230210213222132-1012121003022303-0323000302220103-3321133300210232-3211231012303323-1310203122333131-3311113033023200-0213002021322231) |
| `routes.waf_type.app_firewall.app_firewall.kind` | [routes.waf_type.app_firewall.app_firewall.kind](resources--route--reference--group-003.md#canonical-0221200322002023-1111100102000210-0022012102232211-2131103030222311-2003300313323301-1231311300220312-0010332013321010-1212012202001003) |
| `routes.waf_type.app_firewall.app_firewall.name` | [routes.waf_type.app_firewall.app_firewall.name](resources--route--reference--group-003.md#canonical-2113332133021120-1220103233200303-2230331332001002-0023322013133132-1333132123222221-1021002333312120-2033020003322030-1333321132312220) |
| `routes.waf_type.app_firewall.app_firewall.namespace` | [routes.waf_type.app_firewall.app_firewall.namespace](resources--route--reference--group-003.md#canonical-3023230311231311-0210103201300133-3023212300121331-2221021202312331-3302003200221011-1332131113221120-2022033100100023-0032333322132230) |
| `routes.waf_type.app_firewall.app_firewall.tenant` | [routes.waf_type.app_firewall.app_firewall.tenant](resources--route--reference--group-003.md#canonical-1002113013001210-3323003100130111-3313321303111123-2221132302330032-1323233312030212-3313311302122331-0112123023103011-1201311213223302) |
| `routes.waf_type.app_firewall.app_firewall.uid` | [routes.waf_type.app_firewall.app_firewall.uid](resources--route--reference--group-003.md#canonical-2121302121131022-3021111201022332-2320202231233313-0322330330202130-1202113003123103-0121310301331031-2210223120012030-2333122130310120) |
| `routes.waf_type.disable_waf` | [routes.waf_type.disable_waf](resources--route--reference--group-003.md#canonical-3011003032123110-2333031310103322-1333213021123113-0113330312130021-1111112331131121-0333010000212002-3120112133321320-2201031223002331) |
| `routes.waf_type.inherit_waf` | [routes.waf_type.inherit_waf](resources--route--reference--group-003.md#canonical-1231213121221303-0023310201031321-1202211302132031-0232203213331212-2012120322121332-0132103323132111-0233231003033032-2121200121120001) |
| `timeouts` | [timeouts](resources--route--reference--group-003.md#canonical-3111321022202013-2032120201220322-1301231002330300-0031231130003330-2001303332311331-2230212213301331-2030333231301120-3130212320122332) |
| `timeouts.create` | [timeouts.create](resources--route--reference--group-003.md#canonical-1130001230123010-1221321323133311-0332223011301032-1222031132033112-0312023111332000-1030220102100201-3331011330003112-1100012311023001) |
| `timeouts.delete` | [timeouts.delete](resources--route--reference--group-003.md#canonical-0323312333201213-3130331323123302-0301112112123302-3300000201101000-2112201012020120-0302111212132300-1013131300210210-2222020332310212) |
| `timeouts.read` | [timeouts.read](resources--route--reference--group-003.md#canonical-1101103120000211-1211330033202310-1212223212233002-1111001232202201-0121320303030010-1032322300320121-3022222303233011-3223030111021110) |
| `timeouts.update` | [timeouts.update](resources--route--reference--group-003.md#canonical-0123103313202120-1031213332201320-2333113113000030-1103221112203330-0103121100203123-3132223020113212-0000203132130333-1332013103312230) |

<a id="canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- routes

<a id="canonical-0133221031332213-0212022022311103-2201220331203131-1113221111013120-1003113131023013-3000203201133220-1323332302302103-3202120201221101"></a>

Type: `"object"`. list nested block, Optional.

List of routes to match for incoming request.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("bot_defense_javascript_injection",
    "inherited_bot_defense_javascript_injection"),
  validators.ConflictingListObjectAttributes("inherited_waf_exclusion",
    "waf_exclusion_policy"),
  validators.ConflictingListObjectAttributes("route_destination",
    "route_direct_response"),
  validators.ConflictingListObjectAttributes("route_destination",
    "route_redirect"),
  validators.ConflictingListObjectAttributes("route_direct_response",
    "route_redirect")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 257,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 257,
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
    "ves.io.schema.rules.repeated.max_items": "257"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "257"
  }
}
```

Terraform syntax:

```terraform
routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-0030110122213132-2123220212033233-0100022213101102-3131302210000011-1211232131121123-1232321203321031-1002221321220103-0200200120211233"></a>

### Direct properties for `routes`

- [bot_defense_javascript_injection](resources--route--reference--group-001.md#canonical-3012131313022222-2001213223301312-2233030002013101-2003113330213301-2010231110200010-3301033100022212-2131321012232112-0103110012233022): complete subsection reference.

<a id="canonical-3123131022232020-1110122213003003-2311313103012220-1002200121010033-0132330130132202-3311011203303233-0320102110322331-1311310311221302"></a>

<a id="canonical-2332220020032201-0321303230110110-2100231000232123-1121311122210100-0131330323203030-0100031312320123-3123320100303032-2133112011202302"></a>

#### `routes.disable_location_add` property

Type: `"bool"`. Optional.

Disables append of x-F5 Distributed Cloud-location = &lt;RE-site-name&gt; at route level, if it is
configured at virtual-host level. This configuration is ignored on CE sites.

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

- [inherited_bot_defense_javascript_injection](resources--route--reference--group-001.md#canonical-2123313302223301-3333013230331200-2233332001323113-1230333300031213-3302313023201021-3130132010232033-0222110312122303-2023210000320333): complete subsection reference.

- [inherited_waf_exclusion](resources--route--reference--group-001.md#canonical-0032302101320100-1131213221210110-1230011201023312-3132212221331233-3120213132223021-1023233212231221-0020011010101123-1030112131000013): complete subsection reference.

- [match](resources--route--reference--group-001.md#canonical-1121102331321302-0010313033102331-2333311030330220-0020233203000100-3013202201203230-1223313133131032-1223120122130112-1013213030313303): complete subsection reference.

- [request_cookies_to_add](resources--route--reference--group-001.md#canonical-1222311103011223-2023210230230212-3221022312223230-2011000303333233-2201120130230001-0322033000011011-0012020012001121-1313001212301020): complete subsection reference.

<a id="canonical-3012111200222313-2033222131203000-1203100112010032-2202332323230230-2103103103201211-2032013220023311-3203211310233120-0331022213232012"></a>

<a id="canonical-2130031033120010-0023003133033122-1003313132022100-2031232002333120-1303303031303032-3312220221213222-3323031212010300-2332310113301203"></a>

#### `routes.request_cookies_to_remove` property

Type: `["list", "string"]`. Optional.

List of keys of Cookies to be removed from the HTTP request being sent towards upstream.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [request_headers_to_add](resources--route--reference--group-002.md#canonical-3110130201033312-1313023321233221-3201132303201121-3020210012322131-2100011331031103-0301020201111323-0032131300201003-1021223233210133): complete subsection reference.

<a id="canonical-2000132021120210-2123211030323333-3113030330320322-1310010033210211-0311313313231100-0100121312012110-0330313132201203-3222212320103103"></a>

<a id="canonical-3232112200312233-0113013112233102-0002123023101221-2202100230011312-3110233102220233-1122110222011123-0322211210010203-2102133101022203"></a>

#### `routes.request_headers_to_remove` property

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP request being sent towards upstream.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

- [response_cookies_to_add](resources--route--reference--group-002.md#canonical-1012103221102121-0311302110220133-3312032210222233-2321211003121320-3110000121321323-2231112320331231-3120130302021121-2113210223220100): complete subsection reference.

<a id="canonical-3000123122132002-3003130020302303-1110003203300213-2120031120103322-3100221011232302-1212103211122100-1223133201220003-2100312311201011"></a>

<a id="canonical-1332303003033103-0112311123001131-0210212303122221-3023023222320010-0021021110111301-0222320123223011-2313233030323220-1123102121312002"></a>

#### `routes.response_cookies_to_remove` property

Type: `["list", "string"]`. Optional.

List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire
set-cookie header will be removed.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_headers_to_add](resources--route--reference--group-002.md#canonical-1033033030111003-2022222120022003-1123301032332002-1313030103020021-3112122133001003-0321311203111100-2123212000030110-2132031210122220): complete subsection reference.

<a id="canonical-2021321232003032-1020032131113312-1010023032012230-3231233031203010-2011210022122030-1211132113222310-1020033202001210-3021131321223301"></a>

<a id="canonical-3113133202213203-0320220221320101-1130330010000001-1213231202122022-1311003020120020-1310230320211110-3331020132102223-1323021033202132"></a>

#### `routes.response_headers_to_remove` property

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP response being sent towards downstream.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

- [route_destination](resources--route--reference--group-003.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310): complete subsection reference.

- [route_direct_response](resources--route--reference--group-003.md#canonical-0033133232132220-1302111112032120-3000302211110332-3011333112200100-0032120221023020-1101332301231111-1032033302231223-1323030110311233): complete subsection reference.

- [route_redirect](resources--route--reference--group-003.md#canonical-3030203133200300-2232032232300021-2021302312232013-2012023332102223-2002132031320330-1223213213033033-3233302022003331-1303231221020021): complete subsection reference.

- [service_policy](resources--route--reference--group-003.md#canonical-0223310311031311-1102132201322103-0030022030223013-3213023033100232-1022023332031231-1310033332022211-3110312333301203-2212302001232030): complete subsection reference.

- [waf_exclusion_policy](resources--route--reference--group-003.md#canonical-3113313331301130-2022203212231123-3330200013112321-1303112232000301-0332001232221002-3313202122033003-1323230113203023-3011123020002001): complete subsection reference.

- [waf_type](resources--route--reference--group-003.md#canonical-1322122033023130-0213202101131021-1131300023012121-1013202313230211-0213001020021200-2003012230323221-0031201130003011-1032223101322301): complete subsection reference.

<a id="canonical-3012131313022222-2001213223301312-2233030002013101-2003113330213301-2010231110200010-3301033100022212-2131321012232112-0103110012233022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.bot_defense_javascript_injection` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- routes.bot_defense_javascript_injection

<a id="canonical-2110300030223323-2123122301232111-1311000003023132-3110003301012121-3311313231011333-0020321223333000-0110120022113310-1302111221112312"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense JavaScript Injection Configuration for inline bot defense deployments.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("javascript_tags")}
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
bot_defense_javascript_injection {
  # Configure direct properties listed below.
}
```

<a id="canonical-3022102313331321-2202312312301322-1103123200000000-2021313332100323-0323133111211202-1231233220310021-1111320000103102-3012323002112232"></a>

### Direct properties for `routes.bot_defense_javascript_injection`

<a id="canonical-0333300210132320-2000231113231302-3212102231113112-0131101303213011-2233102323221302-2013100210333002-0000222222020300-3302303023012120"></a>

#### `routes.bot_defense_javascript_injection.javascript_location` property

Type: `"string"`. Optional.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["AFTER_HEAD","AFTER_TITLE_END","BEFORE_SCRIPT"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
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

- [javascript_tags](resources--route--reference--group-001.md#canonical-3023111011302212-1023013120020122-1221033020111230-2330123312100111-2120120303302011-0100200230213023-2323312001110330-1202333223230313): complete subsection reference.

<a id="canonical-3023111011302212-1023013120020122-1221033020111230-2330123312100111-2120120303302011-0100200230213023-2323312001110330-1202333223230313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.bot_defense_javascript_injection.javascript_tags` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.bot_defense_javascript_injection](resources--route--reference--group-001.md#canonical-3012131313022222-2001213223301312-2233030002013101-2003113330213301-2010231110200010-3301033100022212-2131321012232112-0103110012233022)
- routes.bot_defense_javascript_injection.javascript_tags

<a id="canonical-3023000021320011-1330102330112013-0223010021210211-3210231100312313-1131030032001313-2110010003112201-0233312213202313-2302232221021201"></a>

Type: `"object"`. list nested block, Optional.

Select Add item to configure your JavaScript tag. If adding both Bot Adv and Fraud, the Bot
JavaScript should be added first.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("javascript_url")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
javascript_tags {
  # Configure direct properties listed below.
}
```

<a id="canonical-3122122121102123-0311113301303213-2001001311300132-2220002001331300-2103021103303200-1112013310330221-0100200311101301-2300032213333022"></a>

### Direct properties for `routes.bot_defense_javascript_injection.javascript_tags`

<a id="canonical-2131130212122122-0123330303311320-1002233332203212-1200323130130230-3002301333231131-3131223330221303-0330102200133110-0310123233103133"></a>

#### `routes.bot_defense_javascript_injection.javascript_tags.javascript_url` property

Type: `"string"`. Optional.

Please enter the full URL (include domain and path), or relative path.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 2048,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.max_bytes": "2048",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "2048",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [tag_attributes](resources--route--reference--group-001.md#canonical-1111122211123332-1130212302210202-1002122120121203-2122203320220203-2312130000203100-1210300330320313-3320032203220303-3132232130100013): complete subsection reference.

<a id="canonical-1111122211123332-1130212302210202-1002122120121203-2122203320220203-2312130000203100-1210300330320313-3320032203220303-3132232130100013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.bot_defense_javascript_injection.javascript_tags.tag_attributes` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.bot_defense_javascript_injection](resources--route--reference--group-001.md#canonical-3012131313022222-2001213223301312-2233030002013101-2003113330213301-2010231110200010-3301033100022212-2131321012232112-0103110012233022)
- [routes.bot_defense_javascript_injection.javascript_tags](resources--route--reference--group-001.md#canonical-3023111011302212-1023013120020122-1221033020111230-2330123312100111-2120120303302011-0100200230213023-2323312001110330-1202333223230313)
- routes.bot_defense_javascript_injection.javascript_tags.tag_attributes

<a id="canonical-2321030323303300-0101232310030120-0322202202213132-0133133321033111-0120202222001333-0233302011010203-3100331002021012-1023303201231222"></a>

Type: `"object"`. list nested block, Optional.

Add the tag attributes you want to include in your JavaScript tag.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
tag_attributes {
  # Configure direct properties listed below.
}
```

<a id="canonical-2033203132112022-2102312103310101-1131133122203033-0220102330031131-3100110322032331-1112211231301201-1002011222113023-1232323131132313"></a>

### Direct properties for `routes.bot_defense_javascript_injection.javascript_tags.tag_attributes`

<a id="canonical-3013110013231302-1113133132231030-1112001130031201-0210330210122331-0103033231020220-3112212121001023-1233201321112032-0132310123133332"></a>

#### `routes.bot_defense_javascript_injection.javascript_tags.tag_attributes.javascript_tag` property

Type: `"string"`. Optional.

\[Enum:
JS\_ATTR\_ID|JS\_ATTR\_CID|JS\_ATTR\_CN|JS\_ATTR\_API\_DOMAIN|JS\_ATTR\_API\_URL|JS\_ATTR\_API\_PATH|JS\_ATTR\_ASYNC|JS\_ATTR\_DEFER\]
Select from one of the predefined tag attributes. Possible values are \`JS\_ATTR\_ID\`,
\`JS\_ATTR\_CID\`, \`JS\_ATTR\_CN\`, \`JS\_ATTR\_API\_DOMAIN\`, \`JS\_ATTR\_API\_URL\`,
\`JS\_ATTR\_API\_PATH\`, \`JS\_ATTR\_ASYNC\`, \`JS\_ATTR\_DEFER\`. Defaults to \`JS\_ATTR\_ID\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["JS_ATTR_API_DOMAIN","JS_ATTR_API_PATH","JS_ATTR_API_URL","JS_ATTR_ASYNC","JS_ATTR_CID","JS_ATTR_CN","JS_ATTR_DEFER","JS_ATTR_ID"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("JS_ATTR_ID",
    "JS_ATTR_CID",
    "JS_ATTR_CN",
    "JS_ATTR_API_DOMAIN",
    "JS_ATTR_API_URL",
    "JS_ATTR_API_PATH",
    "JS_ATTR_ASYNC",
    "JS_ATTR_DEFER"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "JS_ATTR_ID",
  "enum": [
    "JS_ATTR_ID",
    "JS_ATTR_CID",
    "JS_ATTR_CN",
    "JS_ATTR_API_DOMAIN",
    "JS_ATTR_API_URL",
    "JS_ATTR_API_PATH",
    "JS_ATTR_ASYNC",
    "JS_ATTR_DEFER"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2003301222121222-3100123212002211-1312021012332002-2130011102011131-3202220132210221-3331322210203213-1012103133311002-3232300211312231"></a>

<a id="canonical-1211221221001012-1012132230320221-2330033132100102-1312021100033102-2110113332313310-3222320020330033-2123103102102012-2002211300331301"></a>

#### `routes.bot_defense_javascript_injection.javascript_tags.tag_attributes.tag_value` property

Type: `"string"`. Optional.

Value. Add the tag attribute value.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.max_bytes": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024"
  }
}
```

<a id="canonical-2123313302223301-3333013230331200-2233332001323113-1230333300031213-3302313023201021-3130132010232033-0222110312122303-2023210000320333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.inherited_bot_defense_javascript_injection` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- routes.inherited_bot_defense_javascript_injection

<a id="canonical-3031110333212100-3311303001122332-1222202322013313-3311021130111113-2203331200230023-3323131021212321-1230230101110210-2221003022200013"></a>

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
inherited_bot_defense_javascript_injection = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0032302101320100-1131213221210110-1230011201023312-3132212221331233-3120213132223021-1023233212231221-0020011010101123-1030112131000013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.inherited_waf_exclusion` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- routes.inherited_waf_exclusion

<a id="canonical-0222323203311203-3122123113231311-2030320123112331-3301222010331122-0011033212222010-2130102323010020-2132122111233102-0232133102302023"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inherited waf exclusion.

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
inherited_waf_exclusion = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1121102331321302-0010313033102331-2333311030330220-0020233203000100-3013202201203230-1223313133131032-1223120122130112-1013213030313303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.match` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- routes.match

<a id="canonical-1212232011123331-0302020011221001-3303220003211103-1002301213133003-2001330133131302-3220123202131221-1200113300031113-1032022331013012"></a>

Type: `"object"`. list nested block, Optional.

Match. Route match condition.

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
match {
  # Configure direct properties listed below.
}
```

<a id="canonical-0123302323203323-0213322232332301-3330331303211003-3101302030121212-0111022313010032-3300211322331320-0112233310210223-1220021021112120"></a>

### Direct properties for `routes.match`

- [headers](resources--route--reference--group-001.md#canonical-1231001332220321-3202001223021011-1122010000110133-0332120211030122-2020003311013111-3330112211131113-1211310101122121-2022023230321200): complete subsection reference.

<a id="canonical-1300332101201021-3020022112002332-1300330031031032-2211100111020012-3022020311132321-0003333111211320-1300330232101233-0112200010312100"></a>

<a id="canonical-1112100223112011-3330132312200202-0200132223002212-3113101212101023-1221221031032131-1320012203012111-0220031033333011-0022112302012320"></a>

#### `routes.match.http_method` property

Type: `"string"`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ANY","CONNECT","COPY","DELETE","GET","HEAD","OPTIONS","PATCH","POST","PUT","TRACE"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [incoming_port](resources--route--reference--group-001.md#canonical-0330303320202102-0131302212221113-0133311303023300-3210222200302223-1213303320301020-3122000120001300-3123032302211110-2220330211300021): complete subsection reference.

- [path](resources--route--reference--group-001.md#canonical-1002312122123222-2033203330332312-1323123113131312-2032021222022000-2220031020111332-3000201300302323-3333213122233121-0112231033102033): complete subsection reference.

- [query_params](resources--route--reference--group-001.md#canonical-1032010033330232-3103322230120112-2313232022223123-1323210231121313-2320232301003303-1100122121022210-1010012121021221-3021120211000212): complete subsection reference.

<a id="canonical-1231001332220321-3202001223021011-1122010000110133-0332120211030122-2020003311013111-3330112211131113-1211310101122121-2022023230321200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.match.headers` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.match](resources--route--reference--group-001.md#canonical-1121102331321302-0010313033102331-2333311030330220-0020233203000100-3013202201203230-1223313133131032-1223120122130112-1013213030313303)
- routes.match.headers

<a id="canonical-2102220120303012-3233212223322100-3331211131033303-0112220303333130-3333323230230111-1003133013333231-0212321102320333-3220012013131022"></a>

Type: `"object"`. list nested block, Optional.

Headers. List of (key, value) headers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("exact",
    "presence"),
  validators.ConflictingListObjectAttributes("exact",
    "regex"),
  validators.ConflictingListObjectAttributes("presence",
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minItems": 0,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-0332211312202120-1303200211331201-2331033330102323-1032313322021301-1020013003002310-1000333022201221-3222033131302013-3302300030333130"></a>

### Direct properties for `routes.match.headers`

<a id="canonical-1100212133312132-2220022001031120-2233023111313322-1302203303101333-3330111110123011-3303113120002100-1311110221223103-1001112032133013"></a>

#### `routes.match.headers.exact` property

Type: `"string"`. Optional.

Exclusive with \[presence regular expression\] Header value to match exactly.

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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-0033203301230103-0103020212311113-3212121200301323-3012200302122122-0011301111303101-3210302212021021-1232300033232220-1133033012332201"></a>

<a id="canonical-2003100123032112-0230023212112130-2330113130112021-2130113233011322-2123221203200201-3021121002231003-0102132032211233-2222230330113221"></a>

#### `routes.match.headers.invert_match` property

Type: `"bool"`. Optional.

Invert the result of the match to detect missing header or non-matching value.

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

<a id="canonical-2303000310312022-1221011323220101-0331230010220221-1201120110100110-3002301021031011-3220301131211102-3112003231023201-2233101013233201"></a>

<a id="canonical-2203211033303001-3231312310110230-3023212333230101-2122320332221123-3031331311202231-1121331123323100-3231303131020033-1230203321211023"></a>

#### `routes.match.headers.name` property

Type: `"string"`. Optional.

Name. Name of the header.

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

<a id="canonical-2012121223031133-1132112203323001-1113003311313120-3003030133320120-2213300303231323-3322311022023013-2222200222023001-3100003231122330"></a>

<a id="canonical-0130030313000212-3331201231332231-0212012123202220-1120023210110112-2330120102121033-1021032332103321-2130200111212211-3123131322312013"></a>

#### `routes.match.headers.presence` property

Type: `"bool"`. Optional.

Exclusive with \[exact regular expression\] If true, check for presence of header.

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

<a id="canonical-2302031023110023-1133211120233210-1213021003200121-3321221223031331-3110310100212021-0021300030101130-1212301320033001-2123013021113023"></a>

<a id="canonical-3011013121030002-3230201210011231-3120210021301121-3222101330102330-0223330133201212-1021313000121032-0213313201002030-0312311232111021"></a>

#### `routes.match.headers.regex` property

Type: `"string"`. Optional.

Exclusive with \[exact presence\] regular expression match of the header value in re2 format.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-0330303320202102-0131302212221113-0133311303023300-3210222200302223-1213303320301020-3122000120001300-3123032302211110-2220330211300021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.match.incoming_port` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.match](resources--route--reference--group-001.md#canonical-1121102331321302-0010313033102331-2333311030330220-0020233203000100-3013202201203230-1223313133131032-1223120122130112-1013213030313303)
- routes.match.incoming_port

<a id="canonical-3322313333321233-1223203333211101-2103331320321321-0200210333032122-1011220021033223-0111100233311013-2321010323112230-1103130212110100"></a>

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
incoming_port {
  # Configure direct properties listed below.
}
```

<a id="canonical-2231002320203113-3102000002211323-2223003133120230-0110100012230233-3122102301011322-3003101123021112-1323223233233221-2111221201013322"></a>

### Direct properties for `routes.match.incoming_port`

- [no_port_match](resources--route--reference--group-001.md#canonical-3130033300012001-2302020302203032-0000302003213101-2332122331012333-0103130103322232-3223301101200100-0120231213231201-3003032220012030): complete subsection reference.

<a id="canonical-1101210011110003-1100320300132330-1033012213132200-0113200033003231-2230031300222121-2302031222231032-2101010123323323-2031233123301021"></a>

<a id="canonical-2111022320323021-3130332303003133-1000200032213233-0031133112300010-3013022003003003-3032002223030202-2302332211023201-0111310321232011"></a>

#### `routes.match.incoming_port.port` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1220101020002313-3100203213201012-2013100023102301-1011232033231311-3320113033233333-2121331223310113-3331023010301222-0333213003230110"></a>

<a id="canonical-2233020230033131-1203112033023023-1001222032110223-2300030130332132-0230113201120321-2313220031130310-2121000011220120-1201233322213133"></a>

#### `routes.match.incoming_port.port_ranges` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3130033300012001-2302020302203032-0000302003213101-2332122331012333-0103130103322232-3223301101200100-0120231213231201-3003032220012030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.match.incoming_port.no_port_match` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.match](resources--route--reference--group-001.md#canonical-1121102331321302-0010313033102331-2333311030330220-0020233203000100-3013202201203230-1223313133131032-1223120122130112-1013213030313303)
- [routes.match.incoming_port](resources--route--reference--group-001.md#canonical-0330303320202102-0131302212221113-0133311303023300-3210222200302223-1213303320301020-3122000120001300-3123032302211110-2220330211300021)
- routes.match.incoming_port.no_port_match

<a id="canonical-2203033002331200-1312012220203233-2211021011330232-2332003203312210-1032331010332213-0123002320002321-1121033231222030-3013002121303033"></a>

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

<a id="canonical-1002312122123222-2033203330332312-1323123113131312-2032021222022000-2220031020111332-3000201300302323-3333213122233121-0112231033102033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.match.path` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.match](resources--route--reference--group-001.md#canonical-1121102331321302-0010313033102331-2333311030330220-0020233203000100-3013202201203230-1223313133131032-1223120122130112-1013213030313303)
- routes.match.path

<a id="canonical-2310103323020213-2131213230202102-3220022223100010-1003323020132113-3033303303121231-0011312022333201-1203003130222021-3303200033310022"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1310102300333312-3202113011111320-1031030112012210-0312113121121302-1320220220233033-2322003210102123-1130331132210120-0020303131203233"></a>

### Direct properties for `routes.match.path`

<a id="canonical-3330023032200200-1020022031021101-1131111310113112-3210302113111310-1002211103012320-1000122111121020-1013021113202301-3320003101001021"></a>

#### `routes.match.path.path` property

Type: `"string"`. Optional.

Exclusive with \[prefix regular expression\] Exact path value to match.

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
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0331312111033232-2001311102130002-1310000302011113-1223223131311120-2133303313203112-3110012010322122-0220010332223011-1112233111320030"></a>

<a id="canonical-2330201232133201-0230010120212231-0111012322010010-0220131303331213-3010220120131231-2301032102033002-3111210321200112-2001203221033000"></a>

#### `routes.match.path.prefix` property

Type: `"string"`. Optional.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3212303032321121-2030031123331311-2011230211331201-1231323300103123-1030103232203011-2210001033010323-3233300322321130-0023301122123001"></a>

<a id="canonical-1000112330133100-1130303330311320-0331331300000033-1213001133230322-1010213012300023-2313031001320002-0332313100301333-0121010022233323"></a>

#### `routes.match.path.regex` property

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1032010033330232-3103322230120112-2313232022223123-1323210231121313-2320232301003303-1100122121022210-1010012121021221-3021120211000212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.match.query_params` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.match](resources--route--reference--group-001.md#canonical-1121102331321302-0010313033102331-2333311030330220-0020233203000100-3013202201203230-1223313133131032-1223120122130112-1013213030313303)
- routes.match.query_params

<a id="canonical-1321200231320101-1200322230113332-0300330030332322-2223102322333210-0010023310001221-3310310200232010-1200303113111120-0201223330223223"></a>

Type: `"object"`. list nested block, Optional.

Query Parameters. List of (key, value) query parameters.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("key"),
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
query_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-2310013233223020-0011023323312220-0211032203311202-1332022313301010-0330302303201023-2202320101230222-3000023323013231-3223312102002033"></a>

### Direct properties for `routes.match.query_params`

<a id="canonical-3203031022000302-3131320100133212-3331302120013023-3220010123332031-1032231133101031-0333101013031233-0200100312033302-1023310200210312"></a>

#### `routes.match.query_params.exact` property

Type: `"string"`. Optional.

Exclusive with \[regular expression\] Exact match value for the query parameter key.

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

<a id="canonical-2120130330332313-0132123111222313-3321101221020110-1310301313323133-2110022333220312-1330123102133111-2321211132101131-0132003333322012"></a>

<a id="canonical-0303303220332203-3102100031202103-2333101031231032-0033120030122101-0013100010202013-3011202023013311-3220300223232110-0301203031212012"></a>

#### `routes.match.query_params.key` property

Type: `"string"`. Optional.

Query parameter key In the above example, assignee\_username is the key.

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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-1222022213030212-1230200231111230-2101310131101000-2233321132003000-2132203133310003-1330332121031121-2133112112102102-2023311300230300"></a>

<a id="canonical-3031303130232323-1301331320022333-3002210133212023-3332032321333332-3032202001030030-3033310022022130-1032032330201310-1321320322020233"></a>

#### `routes.match.query_params.regex` property

Type: `"string"`. Optional.

Exclusive with \[exact\] regular expression match value for the query parameter key.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1222311103011223-2023210230230212-3221022312223230-2011000303333233-2201120130230001-0322033000011011-0012020012001121-1313001212301020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.request_cookies_to_add` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- routes.request_cookies_to_add

<a id="canonical-0213313220312023-3323111230001202-1010213210223121-2032303322033120-1311002121223311-0233300222211113-1131112010313100-3022232223300020"></a>

Type: `"object"`. list nested block, Optional.

Cookies are key-value pairs to be added to HTTP request being routed towards upstream.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
request_cookies_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-0313210031011213-2021122123310322-2320321221323011-3000222211312222-2301000113332122-0303111113320232-0221220031230123-2232210000202202"></a>

### Direct properties for `routes.request_cookies_to_add`

<a id="canonical-3223320331232223-1020203022131313-3232222002003003-1220223012103310-1323211013100113-1102233012212113-3101001312020011-0303213222012012"></a>

#### `routes.request_cookies_to_add.name` property

Type: `"string"`. Optional.

Name. Name of the cookie in Cookie header.

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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0103202201201002-0320013303131121-3111101302130330-2120302322110003-1232012102232103-0110322220023130-3030202202231202-3200022212213122"></a>

<a id="canonical-2112120202333021-1132111111011122-2022033323301112-1212122203211320-0232203332221302-0031002030102122-0130033020222002-2021221211110021"></a>

#### `routes.request_cookies_to_add.overwrite` property

Type: `"bool"`. Optional.

Should the value be overwritten? If true, the value is overwritten to existing values. not
overwrite. Defaults to \`do\`.

Additional upstream details:

If true, the value is overwritten to existing values. Default value is do not overwrite.

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

- [secret_value](resources--route--reference--group-001.md#canonical-0131021101021111-2020001221122310-0130013213011131-2010320233202120-0230331100012010-0223112033311130-3122300203102331-3031102030131230): complete subsection reference.

<a id="canonical-1320130001223112-0300321201003011-0133200122112012-1122323331211122-2312301313110030-3302333312222232-3230011232203330-0132312313120223"></a>

<a id="canonical-2332310210300233-3202022220023322-1031300200232301-3023201222012003-1111012213032222-1230201101203020-3233230211010120-1332112321001130"></a>

#### `routes.request_cookies_to_add.value` property

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the Cookie header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-0131021101021111-2020001221122310-0130013213011131-2010320233202120-0230331100012010-0223112033311130-3122300203102331-3031102030131230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.request_cookies_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.request_cookies_to_add](resources--route--reference--group-001.md#canonical-1222311103011223-2023210230230212-3221022312223230-2011000303333233-2201120130230001-0322033000011011-0012020012001121-1313001212301020)
- routes.request_cookies_to_add.secret_value

<a id="canonical-1000001123202002-2110322230212232-1032220013300111-2223210221320103-3311121032102200-2331013110312113-0202133131033200-0122032011023023"></a>

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
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-3023200121011003-1323313130233033-0130301231330221-3312031133313110-0201112323303113-1312113210010210-1131010301232210-2003223300202233"></a>

### Direct properties for `routes.request_cookies_to_add.secret_value`

- [blindfold_secret_info](resources--route--reference--group-001.md#canonical-2003130033020210-0102213301131030-0030002310012301-2212221030030031-1203100013033131-3033101200132321-2202013220312032-2012011012220331): complete subsection reference.

- [clear_secret_info](resources--route--reference--group-002.md#canonical-0332020022011013-0132312011111202-2313031230311313-3200013110320001-3202013302013121-0330000023331320-0222300302312031-1103301330312022): complete subsection reference.

<a id="canonical-2003130033020210-0102213301131030-0030002310012301-2212221030030031-1203100013033131-3033101200132321-2202013220312032-2012011012220331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.request_cookies_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.request_cookies_to_add](resources--route--reference--group-001.md#canonical-1222311103011223-2023210230230212-3221022312223230-2011000303333233-2201120130230001-0322033000011011-0012020012001121-1313001212301020)
- [routes.request_cookies_to_add.secret_value](resources--route--reference--group-001.md#canonical-0131021101021111-2020001221122310-0130013213011131-2010320233202120-0230331100012010-0223112033311130-3122300203102331-3031102030131230)
- routes.request_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-2310102223033013-2301333133211323-3230210002120111-0311310201030322-1212223202210211-2212222220211030-0332323222201030-2321131231013311"></a>

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

<a id="canonical-2013030030121232-2031101313002310-0001123123122101-1113032131011023-2031213321003113-0101111010030303-1200211303030302-3121130112221132"></a>

### Direct properties for `routes.request_cookies_to_add.secret_value.blindfold_secret_info`

<a id="canonical-1130101311130133-2011133302020333-2313032330002012-2010010221333033-0221301001022332-0112232112113011-1333102020110121-0002223331010122"></a>

#### `routes.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-3111232031032011-3123032202332103-1102202330002033-2112033131130131-0223301233203321-2212210201231213-3213220133200133-2231131220000301"></a>
