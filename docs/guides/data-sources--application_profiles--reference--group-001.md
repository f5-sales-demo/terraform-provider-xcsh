---
page_title: "xcsh_application_profiles reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_application_profiles reference."
---

# xcsh_application_profiles reference

<a id="canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- Property reference

<a id="canonical-2312003133231201-0130313320222022-2110132132333001-1123030002031103-2213102232301110-3200132030123120-1020231012020232-3110333223201210"></a>

### Direct properties for `xcsh_application_profiles`

- [advanced_tcp_profile](data-sources--application_profiles--reference--group-001.md#canonical-3021122313210131-2313033223133231-0330112103210330-1123322311012132-1112331223212022-3311220221212131-2333232332331221-1001103023222323): complete subsection reference.

<a id="canonical-1111033000233032-2010110330112021-3020021200333133-0122202223103122-3322232222132030-1223321130112120-3201010111311332-0333222100102331"></a>

<a id="canonical-1003010121301333-3023020211033211-0120301001220113-3330301133332012-3213003103222231-2313133031023033-2212020000000002-2012021201220011"></a>

#### `annotations` property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Additional upstream details:

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

- [ddos_profile](data-sources--application_profiles--reference--group-001.md#canonical-0221310331001010-1231120233222200-3103210111100320-0121030023132112-2012102312231220-0311210333123322-0133000211232001-0223320121031332): complete subsection reference.

<a id="canonical-1000130003032211-2302200201022221-3322022302310031-2222001113012331-2312200112003203-1023132331330033-0130030333121112-1321213211010320"></a>

<a id="canonical-0022101013120133-3223320002210223-3330300133300230-2002121130330033-1223013220131213-2322213033121033-2121230133201203-0331110111132302"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the ApplicationProfiles.

Additional upstream details:

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3102322000133300-0032022312310313-3313102210210131-3110312201300220-0302303023020232-2330302020202110-3132221223210232-1123103203010211"></a>

<a id="canonical-0230021023011233-1021031122301030-2121322030322122-3321012000221120-0330022303323131-3301002330112111-2322213320310112-1002220302223012"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [irules](data-sources--application_profiles--reference--group-001.md#canonical-2300032011023220-0003113120202231-2113023300123123-0013230121220321-0310120231310032-0012232113223333-2122301113222011-3212010010132123): complete subsection reference.

<a id="canonical-3203311222131321-3130221030332011-1011012120203121-1312033211210200-1213203211001011-2302333012332123-3113002313013033-2302311021022221"></a>

<a id="canonical-2020100213101333-1330310113011332-3121120230102013-2012022320000111-0112003232220123-3122230302123213-1103013320032013-0003110120210113"></a>

#### `labels` property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

<a id="canonical-3222211203100022-2033032210200233-0011013010121130-3131301033231301-3211103000312102-1021132201210103-1001110312133310-1021213022003302"></a>

<a id="canonical-3320020032032030-2312230033130220-2033113323002212-1132230012112113-1312131311323001-1101001233110213-3123313200331232-3222031123100323"></a>

#### `name` property

Type: `"string"`. Required.

Name of the ApplicationProfiles.

Additional upstream details:

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1100123023333110-2313212212121013-1321230011322301-1213233121030200-2221231210331123-2013123312301333-1301230100320202-3312111233120031"></a>

<a id="canonical-2223023030023213-1122120223332112-0231131132220011-1102211122003122-3222003031000022-2332122121032013-2212302011320033-1102112231312120"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the ApplicationProfiles exists.

Additional upstream details:

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130): complete subsection reference.

<a id="canonical-3231131223003100-2330121320111330-3002232131211030-0300010212311313-2023303200310331-2230032321132132-1101033223033332-3232300131233231"></a>

### All schema paths for `xcsh_application_profiles`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `advanced_tcp_profile` | [advanced_tcp_profile](data-sources--application_profiles--reference--group-001.md#canonical-3103131213032020-2232220002323000-3022110002112000-2132202321322211-0103100312231112-0011131333112302-0131132001303202-2001112033100223) |
| `advanced_tcp_profile.disable_tcp_advanced_profile` | [advanced_tcp_profile.disable_tcp_advanced_profile](data-sources--application_profiles--reference--group-001.md#canonical-0031021223020121-0313201222331322-2213213200312303-1200002202321001-2020122300120332-3232201110311310-2222112231313233-1212332103200323) |
| `advanced_tcp_profile.enable_tcp_advanced_profile` | [advanced_tcp_profile.enable_tcp_advanced_profile](data-sources--application_profiles--reference--group-001.md#canonical-2132123223321020-3220331133132202-2132112203103122-3030222322022111-0112222320032010-0213023031331001-1032132003323303-2101312322033331) |
| `annotations` | [annotations](data-sources--application_profiles--reference--group-001.md#canonical-1111033000233032-2010110330112021-3020021200333133-0122202223103122-3322232222132030-1223321130112120-3201010111311332-0333222100102331) |
| `ddos_profile` | [ddos_profile](data-sources--application_profiles--reference--group-001.md#canonical-2320300300230322-1211321231211200-2130113001113201-0022333322320010-0321020212300310-1202032132303002-1133321112002032-3320101300333020) |
| `ddos_profile.disable_ddos_mitigation` | [ddos_profile.disable_ddos_mitigation](data-sources--application_profiles--reference--group-001.md#canonical-2322002221332203-2201123333331111-2111331012010001-0003122232130211-3213000131220033-2020210012320323-3330123030122303-3311333022131021) |
| `ddos_profile.enable_ddos_mitigation` | [ddos_profile.enable_ddos_mitigation](data-sources--application_profiles--reference--group-001.md#canonical-1201301110212113-3212113310322200-2101012321233320-3310312202101201-0012310112100332-0312121210221332-0100311332322101-0201211232113212) |
| `description` | [description](data-sources--application_profiles--reference--group-001.md#canonical-1000130003032211-2302200201022221-3322022302310031-2222001113012331-2312200112003203-1023132331330033-0130030333121112-1321213211010320) |
| `id` | [ID](data-sources--application_profiles--reference--group-001.md#canonical-3102322000133300-0032022312310313-3313102210210131-3110312201300220-0302303023020232-2330302020202110-3132221223210232-1123103203010211) |
| `irules` | [irules](data-sources--application_profiles--reference--group-001.md#canonical-2133110011210120-3221112323330121-0301000132001231-3020023210120022-3213032011312201-3103230032102013-0212012120003223-3330103002332200) |
| `irules.kind` | [irules.kind](data-sources--application_profiles--reference--group-001.md#canonical-2231003121311202-0213213112332303-0030001011123023-0202323320301000-2222332313302223-0110031013300313-1003021123332320-3110011302033013) |
| `irules.name` | [irules.name](data-sources--application_profiles--reference--group-001.md#canonical-3032133331010100-0130211001111000-3013131111201000-0203101213010032-0101102112312322-3000330001202100-0310312321112210-1231230202312022) |
| `irules.namespace` | [irules.namespace](data-sources--application_profiles--reference--group-001.md#canonical-3011031323202031-2003130132103122-1022010132132101-1332333031231233-2113023203202320-2330031200031011-0302100133331233-3332022003111310) |
| `irules.tenant` | [irules.tenant](data-sources--application_profiles--reference--group-001.md#canonical-0113113302233222-0203333101022230-3110022201223300-3111200220033013-3232002213031321-3112221330302101-2221032120231323-3122010002331010) |
| `irules.uid` | [irules.uid](data-sources--application_profiles--reference--group-001.md#canonical-2203323300131223-1213221320020123-1022210120132113-3130023331202232-0231320232213032-0323103100113100-0123100233211303-0002003310201103) |
| `labels` | [labels](data-sources--application_profiles--reference--group-001.md#canonical-3203311222131321-3130221030332011-1011012120203121-1312033211210200-1213203211001011-2302333012332123-3113002313013033-2302311021022221) |
| `name` | [name](data-sources--application_profiles--reference--group-001.md#canonical-3222211203100022-2033032210200233-0011013010121130-3131301033231301-3211103000312102-1021132201210103-1001110312133310-1021213022003302) |
| `namespace` | [namespace](data-sources--application_profiles--reference--group-001.md#canonical-1100123023333110-2313212212121013-1321230011322301-1213233121030200-2221231210331123-2013123312301333-1301230100320202-3312111233120031) |
| `virtual_server` | [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-1231202302200330-2103201112010223-1211212111221222-0330313120012320-2323003313200213-2022013313322102-1230030320022013-0231202313021200) |
| `virtual_server.access_profile` | [virtual_server.access_profile](data-sources--application_profiles--reference--group-001.md#canonical-3301203010312110-0311022301323111-1330231332111100-2011320331000132-2222130021103133-1331303222310033-0030323201201010-1103232313122231) |
| `virtual_server.access_profile.kind` | [virtual_server.access_profile.kind](data-sources--application_profiles--reference--group-001.md#canonical-2231313103100123-1012223020200132-2030112321133131-0220311323301323-0220203233301110-2311102001232002-0233002222132323-3322222210121110) |
| `virtual_server.access_profile.name` | [virtual_server.access_profile.name](data-sources--application_profiles--reference--group-001.md#canonical-1001030002331213-1200220200021021-3201311020122131-2212313223210133-3122313001223220-1120130322012302-2333031103020103-2011310110121310) |
| `virtual_server.access_profile.namespace` | [virtual_server.access_profile.namespace](data-sources--application_profiles--reference--group-001.md#canonical-1202313233310023-2303312323211101-2001001313003333-2120023100022020-1122211022133301-1210121203322231-2313120300100021-0322132303033111) |
| `virtual_server.access_profile.tenant` | [virtual_server.access_profile.tenant](data-sources--application_profiles--reference--group-001.md#canonical-3213211300012100-2020000321031102-3033213323022002-2210033232122303-2333200332120001-2311200303030333-0303332210131321-2021223102113003) |
| `virtual_server.access_profile.uid` | [virtual_server.access_profile.uid](data-sources--application_profiles--reference--group-001.md#canonical-2233312101133310-1111313233121311-0123002303310202-3332331102213102-1331030031323222-0302203002201213-2310321013022120-0310323012123132) |
| `virtual_server.address_translation` | [virtual_server.address_translation](data-sources--application_profiles--reference--group-001.md#canonical-3111302313230133-2200011021210113-0023331013233300-3132021121122301-2010330222212230-1300230031112223-0210130111031012-1012132131033200) |
| `virtual_server.address_translation.address_translation_disable` | [virtual_server.address_translation.address_translation_disable](data-sources--application_profiles--reference--group-001.md#canonical-0333100012003221-3200321123211200-3303231003000021-0012100202222013-1111303330313130-3331303221230020-2113320030021002-3232202123121300) |
| `virtual_server.address_translation.address_translation_enable` | [virtual_server.address_translation.address_translation_enable](data-sources--application_profiles--reference--group-001.md#canonical-2310003010031330-1103021321220101-2110200112013000-1010330212232203-1110223010001323-1323203330230203-2321210132033320-3330122102130101) |
| `virtual_server.auto_last_hop` | [virtual_server.auto_last_hop](data-sources--application_profiles--reference--group-001.md#canonical-2103132212111120-2202120331300033-1000221003232103-0010001000132330-3301310123322201-1332000011200121-2121211201322111-1000123012312330) |
| `virtual_server.auto_last_hop.auto_last_hop_default` | [virtual_server.auto_last_hop.auto_last_hop_default](data-sources--application_profiles--reference--group-001.md#canonical-0033033102121111-1003132132313300-3313002022121131-3131133122300133-3012033310320320-2313312211100312-1112302200233220-0100001011011301) |
| `virtual_server.auto_last_hop.auto_last_hop_disable` | [virtual_server.auto_last_hop.auto_last_hop_disable](data-sources--application_profiles--reference--group-001.md#canonical-3310122322320331-1011030232231333-3303321213223210-2323130100230010-3212012213201132-0113312033230321-3322202001332131-2311100302011303) |
| `virtual_server.auto_last_hop.auto_last_hop_enable` | [virtual_server.auto_last_hop.auto_last_hop_enable](data-sources--application_profiles--reference--group-001.md#canonical-2002020012330323-0233012113113230-1000323110103211-1133012112003332-2112110100330132-3021310202232111-1110121030233113-0000021223223323) |
| `virtual_server.clone_pool_client` | [virtual_server.clone_pool_client](data-sources--application_profiles--reference--group-001.md#canonical-1133020132233211-2310110101220011-0012011332002321-1013213132313033-2101320220311103-3111012322101121-0030230000333132-3232002121221313) |
| `virtual_server.clone_pool_client.kind` | [virtual_server.clone_pool_client.kind](data-sources--application_profiles--reference--group-001.md#canonical-3213032121001202-0030120300231013-1122132212012331-0213111222032330-0231101222112200-0021300203313331-2021101022032200-3003233111333003) |
| `virtual_server.clone_pool_client.name` | [virtual_server.clone_pool_client.name](data-sources--application_profiles--reference--group-001.md#canonical-1321303301323112-3232331030023023-3022333311103010-3213223013313322-0301310112122133-0213021320332213-1321011203112101-3020013230120201) |
| `virtual_server.clone_pool_client.namespace` | [virtual_server.clone_pool_client.namespace](data-sources--application_profiles--reference--group-001.md#canonical-2211231222223111-2013313312230333-2230101323231222-3310030032031122-1311333213212023-3001220311021101-2201003010330230-3331023313321232) |
| `virtual_server.clone_pool_client.tenant` | [virtual_server.clone_pool_client.tenant](data-sources--application_profiles--reference--group-001.md#canonical-2010332203320110-3211023123321223-2220230012123321-3220030210210222-1120021002013323-3013300001130013-1222022213222021-1102013303321332) |
| `virtual_server.clone_pool_client.uid` | [virtual_server.clone_pool_client.uid](data-sources--application_profiles--reference--group-001.md#canonical-1233111210320002-1132132001131321-3331021232002313-3213030003001313-0100233223131310-0101033102133302-1212313112300010-3321322311312100) |
| `virtual_server.clone_pool_server` | [virtual_server.clone_pool_server](data-sources--application_profiles--reference--group-001.md#canonical-0221313111003203-0222221132312111-2113212313313121-3203223020221012-0120231221323112-2032123110022222-2323131330111332-3112223023201220) |
| `virtual_server.clone_pool_server.kind` | [virtual_server.clone_pool_server.kind](data-sources--application_profiles--reference--group-001.md#canonical-1233112111102023-0311022121101203-1103103012100323-3002011000102131-2303210030131133-0232113013121110-0210132133322212-3222003200112331) |
| `virtual_server.clone_pool_server.name` | [virtual_server.clone_pool_server.name](data-sources--application_profiles--reference--group-001.md#canonical-2303031121211220-1001230213300030-1230003332311032-2020100332231213-2002002231020011-1210302022002120-3131131320031001-0031001203220231) |
| `virtual_server.clone_pool_server.namespace` | [virtual_server.clone_pool_server.namespace](data-sources--application_profiles--reference--group-001.md#canonical-2101320312013231-0223313033130100-2303200330132312-1032011011320301-3201333231101301-0303122123311321-0010020112230332-3320033113030101) |
| `virtual_server.clone_pool_server.tenant` | [virtual_server.clone_pool_server.tenant](data-sources--application_profiles--reference--group-001.md#canonical-0032333022002101-1013203220103012-2001022001111112-1223202313132122-0230230300111130-2222232131103130-3002310012201213-0113010322130101) |
| `virtual_server.clone_pool_server.uid` | [virtual_server.clone_pool_server.uid](data-sources--application_profiles--reference--group-001.md#canonical-1213233212200103-0311132031101102-1310100010321021-2230013220132230-1202213133313003-3222020310320312-3333231200201320-3203202321303212) |
| `virtual_server.connection_limit` | [virtual_server.connection_limit](data-sources--application_profiles--reference--group-001.md#canonical-3003311010303322-2010201320203020-2022302001310330-3022301301000020-1332202330032100-3120230332203201-3101010322033203-1000331103200032) |
| `virtual_server.connection_rate_limit` | [virtual_server.connection_rate_limit](data-sources--application_profiles--reference--group-001.md#canonical-3012230232332103-2322330122210131-1311320333112003-1131223222023210-0013122003222112-0201211222130320-1003001021101031-3103322232233031) |
| `virtual_server.connection_rate_limit_mode` | [virtual_server.connection_rate_limit_mode](data-sources--application_profiles--reference--group-001.md#canonical-3131301032310212-2101213001120020-1230322302233022-3233223330230310-3120121330213030-1101103200302113-0112333012303230-0113213203012130) |
| `virtual_server.connection_rate_limit_mode.per_destination_address` | [virtual_server.connection_rate_limit_mode.per_destination_address](data-sources--application_profiles--reference--group-001.md#canonical-3232212303310231-3321200231301300-3113023011023311-3313110003322033-3033323201110311-1020200001001221-2000001212031120-1222302330103002) |
| `virtual_server.connection_rate_limit_mode.per_destination_address.destination_mask` | [virtual_server.connection_rate_limit_mode.per_destination_address.destination_mask](data-sources--application_profiles--reference--group-001.md#canonical-1001012133322200-0221033223121133-2331121003312111-2331010112133332-2213231123033330-2303322333300222-3202213011302123-2332321302033223) |
| `virtual_server.connection_rate_limit_mode.per_source_address` | [virtual_server.connection_rate_limit_mode.per_source_address](data-sources--application_profiles--reference--group-001.md#canonical-0222220332310221-3031030011221323-1011011023211200-3200123213223121-3232010100021321-1301303022330301-2133102203321113-1013302001231120) |
| `virtual_server.connection_rate_limit_mode.per_source_address.source_mask` | [virtual_server.connection_rate_limit_mode.per_source_address.source_mask](data-sources--application_profiles--reference--group-001.md#canonical-3231311131311102-0231202211132103-2303012002111113-2330022230301011-0030001031300231-1310220303210001-3022302002212130-3121310330023200) |
| `virtual_server.connection_rate_limit_mode.per_source_destination_address` | [virtual_server.connection_rate_limit_mode.per_source_destination_address](data-sources--application_profiles--reference--group-001.md#canonical-1031201331123121-2000310300021002-3311103012020003-0212333312222020-0230130220020202-2233323312133112-0233212223003233-1102110103323223) |
| `virtual_server.connection_rate_limit_mode.per_source_destination_address.destination_mask` | [virtual_server.connection_rate_limit_mode.per_source_destination_address.destination_mask](data-sources--application_profiles--reference--group-001.md#canonical-2012310123231112-1103313101030133-1301112222002121-2033330023011110-1022331020330012-3231001300131133-3032133330112032-1212002012211230) |
| `virtual_server.connection_rate_limit_mode.per_source_destination_address.source_mask` | [virtual_server.connection_rate_limit_mode.per_source_destination_address.source_mask](data-sources--application_profiles--reference--group-001.md#canonical-2032010011100310-3330311012133302-2122031000123320-0321121311231120-3121120001110123-0103102332131030-1120003131321032-3300220100223020) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server` | [virtual_server.connection_rate_limit_mode.per_virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-1331330110210211-0333212231000230-0110321023330032-0102230332332312-3201232310221002-2210232211210120-2112233221213103-0120230211021302) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address` | [virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address](data-sources--application_profiles--reference--group-002.md#canonical-3220030303033011-3012301301201201-0101030020012131-1020011223121000-3211332310323201-1001300032010220-0220233121023231-1222131010302332) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address.destination_mask` | [virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address.destination_mask](data-sources--application_profiles--reference--group-002.md#canonical-2010001302211202-1100010010013311-1233103233003303-1033120112232121-1103132203003312-3312101021011210-3301301231332313-2120331213111232) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_source_address` | [virtual_server.connection_rate_limit_mode.per_virtual_server_source_address](data-sources--application_profiles--reference--group-002.md#canonical-2200310022132310-0022320033303201-1221121303212313-0023032311121021-0132213113023313-0203322021121123-0310121123003232-2203111130300323) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_source_address.source_mask` | [virtual_server.connection_rate_limit_mode.per_virtual_server_source_address.source_mask](data-sources--application_profiles--reference--group-002.md#canonical-0210301212220123-3231221303212121-0200023030013231-2010121112031133-1210122320102032-0221312232012133-2032100003022101-3322203332220022) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address` | [virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address](data-sources--application_profiles--reference--group-002.md#canonical-2101111302302031-3303033332212322-2221033113310323-2212330000212022-0031231213311210-1013012320030012-2123011313130230-2121022213132100) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address.destination_mask` | [virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address.destination_mask](data-sources--application_profiles--reference--group-002.md#canonical-3331131121102303-2113213123122300-2100102030220210-0100231313300302-0333313102223213-3203320303213332-1223313022313131-0033133122300323) |
| `virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address.source_mask` | [virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address.source_mask](data-sources--application_profiles--reference--group-002.md#canonical-2201301222220233-0323210012201101-1011301021320331-1202310001012032-3201303223330020-3112323200233000-0000031021123301-3203230213203121) |
| `virtual_server.default_persistence_profile` | [virtual_server.default_persistence_profile](data-sources--application_profiles--reference--group-002.md#canonical-0132132103200311-2033001011211102-0100003330302202-0233230330203011-2233203102120101-2103030030021102-3123032023321232-2003112322313011) |
| `virtual_server.default_persistence_profile.kind` | [virtual_server.default_persistence_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-1003033222132011-1212223230221031-2122020121312023-3200220110320323-2202033120032100-2221200130121120-1111130032013303-1323321123211022) |
| `virtual_server.default_persistence_profile.name` | [virtual_server.default_persistence_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-0323022333203032-0212101021130100-2302201310301013-0222311202013221-0131201220112030-2000231322101020-3112210220032112-0010110110013033) |
| `virtual_server.default_persistence_profile.namespace` | [virtual_server.default_persistence_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-2121032020200303-2110122310333003-0230110011331321-3021013132210133-2120003102102211-3111120101022303-0111322101212100-1123201200120132) |
| `virtual_server.default_persistence_profile.tenant` | [virtual_server.default_persistence_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-3330133112002113-2311000132033033-2131202220220302-3320331331211023-1002010111320320-2323120222230320-2222203303200211-3121233133002222) |
| `virtual_server.default_persistence_profile.uid` | [virtual_server.default_persistence_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-2001112003110002-0113210100223020-0212001313320330-0201231332103133-3122033311100110-3012310202321213-0030000331123200-3203111200313103) |
| `virtual_server.default_pool` | [virtual_server.default_pool](data-sources--application_profiles--reference--group-002.md#canonical-0032003130213122-2132210021133011-3210323101123201-1331222302230120-0032201321121012-3110301121223230-0003122221011010-1032022120330101) |
| `virtual_server.default_pool.kind` | [virtual_server.default_pool.kind](data-sources--application_profiles--reference--group-002.md#canonical-2001322211031033-3012213203300203-0332021203222310-2022130233030313-0211030233313101-2311132012310131-1031331002122321-2211100200300020) |
| `virtual_server.default_pool.name` | [virtual_server.default_pool.name](data-sources--application_profiles--reference--group-002.md#canonical-3032213320312323-1222330100023233-0203210311022012-2031300302130321-2113001100321111-1031330222223021-0022211120231000-3010311211301130) |
| `virtual_server.default_pool.namespace` | [virtual_server.default_pool.namespace](data-sources--application_profiles--reference--group-002.md#canonical-2213113003303231-3012211330331322-2322232121021123-0301301300313033-2213011220003202-2202123002313023-1110221223322320-2201113301211122) |
| `virtual_server.default_pool.tenant` | [virtual_server.default_pool.tenant](data-sources--application_profiles--reference--group-002.md#canonical-2101231233003012-1022330031222033-3333232013003103-3000200131232022-3211121313031313-3210202200021021-1023111210130210-1330012202211330) |
| `virtual_server.default_pool.uid` | [virtual_server.default_pool.uid](data-sources--application_profiles--reference--group-002.md#canonical-1102111030133221-2033112123021013-0103232100123020-2101023303300120-2303133300032021-1132011233300002-1120020300312300-0203030222333133) |
| `virtual_server.fallback_persistence_profile` | [virtual_server.fallback_persistence_profile](data-sources--application_profiles--reference--group-002.md#canonical-3031222100321330-0030112200331121-0332312032101230-1311302311301132-3331230021310111-3211231210330021-0020121210000121-0002230321230031) |
| `virtual_server.fallback_persistence_profile.kind` | [virtual_server.fallback_persistence_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-0030302100303202-1212333230313111-0303102311120132-3221302203132232-0003230200203003-3113322120233110-0310222332331220-1013031233323303) |
| `virtual_server.fallback_persistence_profile.name` | [virtual_server.fallback_persistence_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-2011111330302121-0333221012103021-2300221212230222-1002311200222123-2120302201321022-0023113201310232-3222100312012030-1120032032031021) |
| `virtual_server.fallback_persistence_profile.namespace` | [virtual_server.fallback_persistence_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-0303302310133032-3032201103100331-2010020001010120-3331030321032303-0012223230010003-3321023122310122-1003202200132331-3001130133211110) |
| `virtual_server.fallback_persistence_profile.tenant` | [virtual_server.fallback_persistence_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-3201033001223330-3121021210100210-0300013013232022-2231132300230223-1312020323022223-1330123210130203-0320221230230233-2022203031303110) |
| `virtual_server.fallback_persistence_profile.uid` | [virtual_server.fallback_persistence_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-1233332121110312-2121230310203112-2122220322000230-0003010021220333-0113233222031100-1103103221322123-0021102002030130-0230131202201112) |
| `virtual_server.fix_profile` | [virtual_server.fix_profile](data-sources--application_profiles--reference--group-002.md#canonical-2333300011301210-3213230203312113-1013022110101200-3032023200322121-0233300201201000-2231022121031032-0010213221331220-2022133112030130) |
| `virtual_server.fix_profile.kind` | [virtual_server.fix_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-1122323333132002-1031133023202302-0020013310120130-2302230332230021-3113310310020301-1013332300031313-0321022323310031-1213321111323220) |
| `virtual_server.fix_profile.name` | [virtual_server.fix_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-1202213102011031-0321201101321200-1230021110201302-2213102011121020-2301120021313320-3323131033231323-3322323101322211-1002113330323223) |
| `virtual_server.fix_profile.namespace` | [virtual_server.fix_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-1022213132133231-3232100032330111-3122311330023201-0302301220113330-3112201112001013-3003130000211031-3110213122111312-3232202203300201) |
| `virtual_server.fix_profile.tenant` | [virtual_server.fix_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-2210201132131102-1311223300223201-2300123300333022-0211333211131023-3220302112010123-0231103320230133-1322301332231230-1332032320230131) |
| `virtual_server.fix_profile.uid` | [virtual_server.fix_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-0301101102332320-3101201312310111-3312202310302130-1030213111232002-0012022000032332-1201331030013301-2113312011323232-3122113333122331) |
| `virtual_server.http` | [virtual_server.http](data-sources--application_profiles--reference--group-002.md#canonical-0332210133000202-2230023030102321-1003012303213311-3321201223311210-2203110311213223-0000100332023203-2000332111113203-0210213200130102) |
| `virtual_server.http.client_ssl_profile` | [virtual_server.http.client_ssl_profile](data-sources--application_profiles--reference--group-002.md#canonical-0123212130303322-2310313033030130-0210020200201023-0021321231300010-2132332222301330-1101211203203322-3312033201032210-1130021131011132) |
| `virtual_server.http.client_ssl_profile.kind` | [virtual_server.http.client_ssl_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-0330212301111223-1331320121103011-0212221011211302-3303331313023020-1200330111331303-0102303120033001-3213310120312331-3332321033233232) |
| `virtual_server.http.client_ssl_profile.name` | [virtual_server.http.client_ssl_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-1301013003203323-3121333113231220-3122313222231100-0211133220321300-2130023302300010-3033212131012010-1101331330003003-3320203132312211) |
| `virtual_server.http.client_ssl_profile.namespace` | [virtual_server.http.client_ssl_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-3202100111313031-1033032122032101-1221112321211023-3220002113000231-1021032033321032-1110102212212031-3010331220333030-0222131333330322) |
| `virtual_server.http.client_ssl_profile.tenant` | [virtual_server.http.client_ssl_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-0011122300002311-0022113012033013-0100020111021212-3203312231221023-1331210321103232-0332220223030030-0213210003023121-0112002313221222) |
| `virtual_server.http.client_ssl_profile.uid` | [virtual_server.http.client_ssl_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-3003121022321333-2131310000203333-2000032211011121-0003102030023003-1211123100111121-1213121300012333-0230132301332200-1013033301331001) |
| `virtual_server.http.http2_client_profile` | [virtual_server.http.http2_client_profile](data-sources--application_profiles--reference--group-002.md#canonical-0012301231330132-1110003010313111-1012101002322020-2033103122023020-2112120112323313-2021302032211212-3123123233130132-1102123322120112) |
| `virtual_server.http.http2_client_profile.kind` | [virtual_server.http.http2_client_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-1121320112113233-2321000023103231-3232010021233333-0102102331321232-1031220113320312-0200000101222203-2312002111011032-1331333213011222) |
| `virtual_server.http.http2_client_profile.name` | [virtual_server.http.http2_client_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-0333113220111131-0200313300010222-1033201120321030-0333010222010011-1133132202332130-1223013300012111-1320233011230332-1332311102301102) |
| `virtual_server.http.http2_client_profile.namespace` | [virtual_server.http.http2_client_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-3011122323210220-2300230321021320-3030310102122103-1222031033020212-3000110301030203-3123311202021100-1112020220302313-0130303101003130) |
| `virtual_server.http.http2_client_profile.tenant` | [virtual_server.http.http2_client_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-0023132203130003-3100033101022030-2233220303000121-3333102120213030-0332021010310331-0320301323221122-1023123002231333-1102301202213102) |
| `virtual_server.http.http2_client_profile.uid` | [virtual_server.http.http2_client_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-3310202122033310-0333103032303210-0111200102021023-0011231000220302-1322130233221310-0331301312021102-3122101102332021-0323123001232021) |
| `virtual_server.http.http2_server_profile` | [virtual_server.http.http2_server_profile](data-sources--application_profiles--reference--group-002.md#canonical-3100221000003202-0011121223033221-0100212112312130-2332013321322131-2011132332002032-0103012333032331-0033303331120310-2322313033002211) |
| `virtual_server.http.http2_server_profile.kind` | [virtual_server.http.http2_server_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-3003201302123102-0131323320121100-0312023002330311-0213200212301013-3200323032133303-0320131331310023-3011232121133130-0333312223212110) |
| `virtual_server.http.http2_server_profile.name` | [virtual_server.http.http2_server_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-0323000333332100-0212123233220221-1013023332202330-1302132132022331-3310133030130320-2003110323303230-3111001213100333-1223200020210133) |
| `virtual_server.http.http2_server_profile.namespace` | [virtual_server.http.http2_server_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-2023133010231332-2322113012220303-3223103233230002-2220312133100221-3210111012001213-1003311203113022-3232213213102311-1112122030231232) |
| `virtual_server.http.http2_server_profile.tenant` | [virtual_server.http.http2_server_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-2230201011113020-0303123301033333-0333232122013021-1312132130100211-1333322202132023-3230002122003302-0331302021010310-3333030123012312) |
| `virtual_server.http.http2_server_profile.uid` | [virtual_server.http.http2_server_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-0221011001313102-2201313332032313-1213311303203310-1103222111020332-1131000202123010-2131012000332021-2022220302221003-2110223233200011) |
| `virtual_server.http.http_client_profile` | [virtual_server.http.http_client_profile](data-sources--application_profiles--reference--group-002.md#canonical-0122213013203311-0011121202023332-1112322313000120-3233320001311220-1000111203212220-3103101321130103-3221111101110100-2110113111101021) |
| `virtual_server.http.http_client_profile.kind` | [virtual_server.http.http_client_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-3130223010330110-3111311310000132-3130133113011032-0003311000320310-2201012310213331-0013203313230313-3220010003322101-2102202312303120) |
| `virtual_server.http.http_client_profile.name` | [virtual_server.http.http_client_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-2030010032032303-0223103012303101-1321230032221332-0331332221202220-0130331323221231-0232212311303331-2021203323123203-0101323320330013) |
| `virtual_server.http.http_client_profile.namespace` | [virtual_server.http.http_client_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-3302333001011032-2331022302123000-1002232303031221-0313332232300122-0131210102123221-1101301131232013-1212001201210220-3223303202021302) |
| `virtual_server.http.http_client_profile.tenant` | [virtual_server.http.http_client_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-2103131031223013-3213122033332101-2203321320001313-1221011301300011-3301123011020222-1202103112032001-1212233221012303-2023202102012103) |
| `virtual_server.http.http_client_profile.uid` | [virtual_server.http.http_client_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-0022030222030032-1221331331300322-1313031203311011-2223010123220202-3330002332112323-3031130300322331-0332113301120201-0121330030002222) |
| `virtual_server.http.http_server_profile` | [virtual_server.http.http_server_profile](data-sources--application_profiles--reference--group-002.md#canonical-2222000223211031-2223113121002311-0222131313111320-3033332030100010-2312312323103010-2303030133110132-2300013102311330-2312310013203000) |
| `virtual_server.http.http_server_profile.kind` | [virtual_server.http.http_server_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-0220020120332021-1320223210302022-2012100001300212-0210310322201231-3131122303321211-3031230200122331-0313230023322333-2003201111210133) |
| `virtual_server.http.http_server_profile.name` | [virtual_server.http.http_server_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-0330022211333111-2133203201302210-2023133103232103-0003112023321033-1202010012013121-0120300310022303-0022101201332130-2311001330212123) |
| `virtual_server.http.http_server_profile.namespace` | [virtual_server.http.http_server_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-0221113310311100-2303211132011110-2132010031333232-3212122120332330-3132330122031230-2213233101002201-0021100300201332-2300032210022022) |
| `virtual_server.http.http_server_profile.tenant` | [virtual_server.http.http_server_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-1123230331322200-0032200020310200-3300311113320211-3333330011121301-0032131300123023-3012111131010021-1031303003012033-3002031313020300) |
| `virtual_server.http.http_server_profile.uid` | [virtual_server.http.http_server_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-2111302002103033-0130233300201033-2311032012120212-2002332233333122-3201023331132122-1231302011033020-0021232302301031-1323113020231331) |
| `virtual_server.http.ocsp_profile` | [virtual_server.http.ocsp_profile](data-sources--application_profiles--reference--group-002.md#canonical-0312201123200303-2231230312032301-0320013132131331-2110201112323330-0303012111100303-2113311232223102-1132201032212311-0223113002312122) |
| `virtual_server.http.ocsp_profile.kind` | [virtual_server.http.ocsp_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-1230321100122233-1330030032331102-1011321132311202-1230201333001112-3020131021003002-2313032102101122-2033020111301233-3310222001333002) |
| `virtual_server.http.ocsp_profile.name` | [virtual_server.http.ocsp_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-2300120023011101-1003011311103222-1021331130111103-0110312333233121-2210110233332230-3323203223021213-3011323230123332-1113121003220113) |
| `virtual_server.http.ocsp_profile.namespace` | [virtual_server.http.ocsp_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-1013230020021112-2103023303323201-1313311132102120-1213302210323321-1231013202311330-2221011000321322-0233032101210112-2111100112223332) |
| `virtual_server.http.ocsp_profile.tenant` | [virtual_server.http.ocsp_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-0100022002001322-1021321332112223-3311221331010022-3020313303313020-0002010112210202-1021000120223003-3212212132212021-2303310033001010) |
| `virtual_server.http.ocsp_profile.uid` | [virtual_server.http.ocsp_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-0012103220212103-3133322111302033-3111110101321033-2133221330012002-0021020122032223-2111033232303201-2033022300101231-3322220221110322) |
| `virtual_server.http.server_ssl_profile` | [virtual_server.http.server_ssl_profile](data-sources--application_profiles--reference--group-002.md#canonical-1002021031003123-3301122331321210-2221332022112120-0113010120102001-2113123001232222-2032211313333321-0210030103211211-2311113023013102) |
| `virtual_server.http.server_ssl_profile.kind` | [virtual_server.http.server_ssl_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-3311012123223100-0110231323110230-0211311200221320-3221103133333001-0212032323233010-2130021033312022-3030030020121301-3002132121032210) |
| `virtual_server.http.server_ssl_profile.name` | [virtual_server.http.server_ssl_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-1030021331210101-2200002322120211-2323013310202120-1203021323310130-1210331100232330-3120031033132003-2331022130003202-0122030011013203) |
| `virtual_server.http.server_ssl_profile.namespace` | [virtual_server.http.server_ssl_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-3113120212023101-2300332332201020-0101203020000121-2201310223112231-0001201222020302-1300033002210211-0022321321223200-3122103120231000) |
| `virtual_server.http.server_ssl_profile.tenant` | [virtual_server.http.server_ssl_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-3031231301112020-0231332321123113-0300221303333200-1301012232120202-0021010332013201-3121003000301320-3121130131313122-0201320020303223) |
| `virtual_server.http.server_ssl_profile.uid` | [virtual_server.http.server_ssl_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-1122012112311002-2201110321212122-1321110010312023-0210330310203121-1020231213111131-0022303111300131-2231123123033200-0130200022003103) |
| `virtual_server.http.stream_profile` | [virtual_server.http.stream_profile](data-sources--application_profiles--reference--group-002.md#canonical-3132213201233220-2021321102011100-2220012312310301-2111100133020122-0113222003302321-0210230000211023-0212123303121012-0330303213211020) |
| `virtual_server.http.stream_profile.kind` | [virtual_server.http.stream_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-0230000210332311-3222330121211231-0313100023201302-3121323333230211-1230211122223110-0010112302220312-2313032022222302-3113110312202011) |
| `virtual_server.http.stream_profile.name` | [virtual_server.http.stream_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-2313322103002010-3131132321133021-2322122220300121-2333301230313310-2033221132021321-1022002030010210-2211220030230310-2031100202031213) |
| `virtual_server.http.stream_profile.namespace` | [virtual_server.http.stream_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-0112132222301202-1200112033033102-2321311030201032-1120322323111123-0000110320100111-1031102032000001-3221213112001131-3000133211321112) |
| `virtual_server.http.stream_profile.tenant` | [virtual_server.http.stream_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-0211303002301213-3133131212013203-2010030331002032-1321123011301201-2010003013300133-1203311212211200-1322322013130321-1311333002021333) |
| `virtual_server.http.stream_profile.uid` | [virtual_server.http.stream_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-0113231310200301-1311112021202013-3120232101113123-3013332122201322-0123131300000300-2202313030200323-3112032231221310-2210002310210331) |
| `virtual_server.http.tcp_client_profile` | [virtual_server.http.tcp_client_profile](data-sources--application_profiles--reference--group-002.md#canonical-3011012310321011-2333310101023002-0011200302330111-2000010233110013-3320100100233230-2301020320100133-0132033001200123-0113100311321001) |
| `virtual_server.http.tcp_client_profile.kind` | [virtual_server.http.tcp_client_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-1212311020300120-3223311303210313-3002012312212301-1022303013313113-3100021220010112-3222130302123310-2010120322132300-0323212023320012) |
| `virtual_server.http.tcp_client_profile.name` | [virtual_server.http.tcp_client_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-3332132222033033-2301003133213321-1213111300233311-1212232002231210-2110011220113012-0323003320000212-3111300220210113-0022001320313130) |
| `virtual_server.http.tcp_client_profile.namespace` | [virtual_server.http.tcp_client_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-3220331132322333-3230233201232213-3303322201112300-0221213302032333-2102021112002120-2301330330101033-2131311012203003-1203120030202132) |
| `virtual_server.http.tcp_client_profile.tenant` | [virtual_server.http.tcp_client_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-1133313331130332-2322311123201010-3320121211210312-3020312130211013-2323012013220031-3032311310222013-3022322131322301-2310200333131302) |
| `virtual_server.http.tcp_client_profile.uid` | [virtual_server.http.tcp_client_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-1033210200222021-1022310333120121-2330230122033303-2031002132312230-1121022000001310-2213323200123123-1211320023200301-3111011021330233) |
| `virtual_server.http.tcp_server_profile` | [virtual_server.http.tcp_server_profile](data-sources--application_profiles--reference--group-002.md#canonical-1031331212321311-2121010201011133-3302302120310303-1102130320220322-3003012010022312-1023032032222021-1022233221133022-2330033333122101) |
| `virtual_server.http.tcp_server_profile.kind` | [virtual_server.http.tcp_server_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-1323303322331223-1320320020221330-3021330232013222-2333003012232321-2231000113111220-2003132302303211-2232213011203200-1110021211323301) |
| `virtual_server.http.tcp_server_profile.name` | [virtual_server.http.tcp_server_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-3300330202103222-0100203132332121-3232202230203232-2131001120123333-3022320322211211-2233200120311222-3020100220131332-1121013213120102) |
| `virtual_server.http.tcp_server_profile.namespace` | [virtual_server.http.tcp_server_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-0320310322010003-3223133131303000-1133212111311322-2200213301033300-3312123331221330-3322033122131302-3012013231211333-0103122012211312) |
| `virtual_server.http.tcp_server_profile.tenant` | [virtual_server.http.tcp_server_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-3123302210032010-0232313203233113-1013211031331112-1333202330031020-1132302220112131-2000332303021013-0020310203102020-1232233011132113) |
| `virtual_server.http.tcp_server_profile.uid` | [virtual_server.http.tcp_server_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-0110223330032031-3001323230312213-0031231320211103-2320031033002210-1020300230111321-0223210330021113-0311332323313323-2021202311213012) |
| `virtual_server.http.websocket_client_profile` | [virtual_server.http.websocket_client_profile](data-sources--application_profiles--reference--group-002.md#canonical-2121221311233212-2120123101131321-2110032213122311-3123023313202112-2020133303031211-1032112331010320-3131021133220001-3113233032033111) |
| `virtual_server.http.websocket_client_profile.kind` | [virtual_server.http.websocket_client_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-0222222103110001-1002002221103301-3202200233023103-1310232302223113-3023320111221312-2212112111132131-1223332200310333-1113210310212102) |
| `virtual_server.http.websocket_client_profile.name` | [virtual_server.http.websocket_client_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-1311012201312322-3202230111100332-0020212011322213-1233102122013231-3303303210030131-0011023300132032-1111013302222222-1211322013233232) |
| `virtual_server.http.websocket_client_profile.namespace` | [virtual_server.http.websocket_client_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-3202101013120120-1333202000102333-2313323001303201-1221213201032333-3300220033232132-0311211233201113-0031210302211023-2133100300023003) |
| `virtual_server.http.websocket_client_profile.tenant` | [virtual_server.http.websocket_client_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-2313300013011033-2233210100003122-2232111002302210-0103030023021332-2112211101320221-3300233330011203-0103320023130202-0010122201310320) |
| `virtual_server.http.websocket_client_profile.uid` | [virtual_server.http.websocket_client_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-1001222312200023-1333330130232011-3020300133322313-2003110012100022-1213310030132202-0323303301321010-3223023103022213-2302120100230311) |
| `virtual_server.http.websocket_server_profile` | [virtual_server.http.websocket_server_profile](data-sources--application_profiles--reference--group-002.md#canonical-2201322310330012-0030203000221221-0021002000020131-2001313000332122-1212102021103111-1202323011122211-0022203212120322-2021130023213312) |
| `virtual_server.http.websocket_server_profile.kind` | [virtual_server.http.websocket_server_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-2113303302210122-1331100322212320-3002313112332101-0111033133110202-0013132123301012-2211230311200211-0021232021323020-3013020023022323) |
| `virtual_server.http.websocket_server_profile.name` | [virtual_server.http.websocket_server_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-1220330222021131-2101022210032013-2232112322030021-1003313011323211-1321333223100132-0200110333130223-2301330013113200-3023133222313032) |
| `virtual_server.http.websocket_server_profile.namespace` | [virtual_server.http.websocket_server_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-1013011332013003-0132001011101221-2102320030130101-1200201121202313-3202122102202223-2323113121102300-2121020030232022-0331322121201310) |
| `virtual_server.http.websocket_server_profile.tenant` | [virtual_server.http.websocket_server_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-1122001220100113-3013222333021331-3132110100200130-0333021013321101-0023321213021310-1222131313112121-1221212033032302-1011002133300211) |
| `virtual_server.http.websocket_server_profile.uid` | [virtual_server.http.websocket_server_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-3013331233210030-1111331300310222-2123230102000022-1331223232103330-1223113213322022-1021101222011022-2000012012102120-2132213313022113) |
| `virtual_server.http3` | [virtual_server.http3](data-sources--application_profiles--reference--group-002.md#canonical-1121301310330311-2011010320101113-1233100013121330-1212222102110332-0233311301302301-3101011232131032-2201031113231103-1111233100033111) |
| `virtual_server.http3.client_ssl_profile` | [virtual_server.http3.client_ssl_profile](data-sources--application_profiles--reference--group-002.md#canonical-0231222012023300-3000220213330321-1201232010123230-2122300030212300-3323232103310130-2020222202310003-2032120013231110-1103203210211223) |
| `virtual_server.http3.client_ssl_profile.kind` | [virtual_server.http3.client_ssl_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-3023210330100002-0301132022221320-1210111112032033-2323010021100302-1221111220121021-2013012310011322-0132212333030201-3023331310002100) |
| `virtual_server.http3.client_ssl_profile.name` | [virtual_server.http3.client_ssl_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-3102011121030132-2022113223311310-2021312011221123-1321102032013232-0001133212123102-3231120011011211-1203000133011022-3203132301322111) |
| `virtual_server.http3.client_ssl_profile.namespace` | [virtual_server.http3.client_ssl_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-2200103300303200-0032203021033103-1220023131000030-0113130133232110-2230031210213110-3222010131020210-2010221031310013-3110202013331212) |
| `virtual_server.http3.client_ssl_profile.tenant` | [virtual_server.http3.client_ssl_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-3013210231013221-2032101203312023-2020332233133332-1031301303312231-0200223320202002-3212100031100202-0220233232230203-1102132230232333) |
| `virtual_server.http3.client_ssl_profile.uid` | [virtual_server.http3.client_ssl_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-0211001030212123-1133333311333100-3201021301100023-3010002033101321-1111311221012113-0000132323310323-1223212133233113-3313201313310101) |
| `virtual_server.http3.http3_profile` | [virtual_server.http3.http3_profile](data-sources--application_profiles--reference--group-002.md#canonical-3220112303100231-3223101201022303-2333230121221013-3100012311121031-2123220033020201-2112013333213311-2220331222111132-1231020322012332) |
| `virtual_server.http3.http3_profile.kind` | [virtual_server.http3.http3_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-2212211331103123-1222030223132231-3212202301323310-0333202003032001-1330111230323301-0111003313132013-2322132033200203-1230321003122103) |
| `virtual_server.http3.http3_profile.name` | [virtual_server.http3.http3_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-1320331030103202-3222003032300311-1023223332000212-2121000221320301-2023213301210031-3112302333020303-3222130021023113-0003131122012022) |
| `virtual_server.http3.http3_profile.namespace` | [virtual_server.http3.http3_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-1220031310231311-1323022030121301-3031000223321303-0221013200110131-2221232021012213-2231301102231013-1003022122312310-0302200000312013) |
| `virtual_server.http3.http3_profile.tenant` | [virtual_server.http3.http3_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-0000311231313331-3002032321122220-1031230122133331-1233111201131021-3033112200202330-2030303232000313-3321033111212131-3322320000100201) |
| `virtual_server.http3.http3_profile.uid` | [virtual_server.http3.http3_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-1020010203221023-0322101300123233-2303221333012232-2320233011032131-3101122132123030-3123211213322113-0111130210113032-2110002311121321) |
| `virtual_server.http3.http_client_profile` | [virtual_server.http3.http_client_profile](data-sources--application_profiles--reference--group-002.md#canonical-0111303033120332-3212222230103002-1130222323000111-1122332302102222-1201303031332111-0001211231200012-0101300313323110-2011320002122311) |
| `virtual_server.http3.http_client_profile.kind` | [virtual_server.http3.http_client_profile.kind](data-sources--application_profiles--reference--group-002.md#canonical-2200202211102132-3330231311121203-0031101230223300-1112023020223203-3120232213213200-3022331010033130-0131321303031111-2210303121203000) |
| `virtual_server.http3.http_client_profile.name` | [virtual_server.http3.http_client_profile.name](data-sources--application_profiles--reference--group-002.md#canonical-1221301021233131-2123220302230000-1003133311333212-0230013231000302-2032022023213211-2302201121132103-0232112103303030-1200301302220303) |
| `virtual_server.http3.http_client_profile.namespace` | [virtual_server.http3.http_client_profile.namespace](data-sources--application_profiles--reference--group-002.md#canonical-1320213213032310-0031223033202201-2030311113223331-0212230001330133-3033003332220010-0003300001302120-0023021022000313-3233122021020312) |
| `virtual_server.http3.http_client_profile.tenant` | [virtual_server.http3.http_client_profile.tenant](data-sources--application_profiles--reference--group-002.md#canonical-2212113133011312-2110111302222332-2120213221033022-2031203211000323-1023113003030331-0322300321332121-0021000232113332-3021232022231002) |
| `virtual_server.http3.http_client_profile.uid` | [virtual_server.http3.http_client_profile.uid](data-sources--application_profiles--reference--group-002.md#canonical-1232303301322000-1231030012103001-1131120020012023-0030100113120313-2230303110102000-2020220231300302-0103300133302130-1211122220000310) |
| `virtual_server.http3.http_server_profile` | [virtual_server.http3.http_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-3221312202023033-1102223012313300-3203010330011230-1210002202322011-2313030323231301-2222112330223212-1201303211311211-3123213113100210) |
| `virtual_server.http3.http_server_profile.kind` | [virtual_server.http3.http_server_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-0101013121221132-1202201020020332-2202330300302233-3313203200201322-1002310330030222-1012311322121200-2132232133011113-1033300123331332) |
| `virtual_server.http3.http_server_profile.name` | [virtual_server.http3.http_server_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-0013313312210120-3130313333213333-1333213021111233-3213323103121202-1301102233021123-3303133120310333-2032312231023020-1323111311320313) |
| `virtual_server.http3.http_server_profile.namespace` | [virtual_server.http3.http_server_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-2213223111233123-1123303202020100-1120121101130001-1010113022330302-2310102333101101-2123223003313132-3012031111320013-3220300102203001) |
| `virtual_server.http3.http_server_profile.tenant` | [virtual_server.http3.http_server_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-2030211003122121-2230233120231211-0113032010002221-3021110031232003-0100121021033311-0230121020212022-3132123120320200-0321222313213122) |
| `virtual_server.http3.http_server_profile.uid` | [virtual_server.http3.http_server_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-1133202011021202-3203233311231300-0221232013231200-0120132201231022-2302233003031303-0000233300032303-2221030002333213-0101322032333203) |
| `virtual_server.http3.quic_profile` | [virtual_server.http3.quic_profile](data-sources--application_profiles--reference--group-003.md#canonical-2311003220220232-3230111332311321-3322012111002332-1020121133211300-1233032213310202-3311330100212020-3100011331330303-2020320213010313) |
| `virtual_server.http3.quic_profile.kind` | [virtual_server.http3.quic_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-1203331001223123-1121330223313120-1033122102032331-3330213310013021-1112032331012131-1111030133002003-0330213110200010-3231112320003111) |
| `virtual_server.http3.quic_profile.name` | [virtual_server.http3.quic_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-0111021202211021-3110031212030122-0003313012020102-1322022003100203-1313300131210223-1103332102031111-3000123132230023-1031100120212321) |
| `virtual_server.http3.quic_profile.namespace` | [virtual_server.http3.quic_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-2320111311120023-2023011231032330-2013033301023013-2020001013232033-2131302110331312-3220312130013031-0130011102002220-1203222002102111) |
| `virtual_server.http3.quic_profile.tenant` | [virtual_server.http3.quic_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-2132213323131023-1321220100310030-3313323012000221-1102232221332033-1232011223323110-0310111030032100-0121223023131133-3210001113331001) |
| `virtual_server.http3.quic_profile.uid` | [virtual_server.http3.quic_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-2023200220330321-1303111011233222-1230003323303202-1102333003332033-2022212213331330-2302133320032220-2332200210130310-2003123300223210) |
| `virtual_server.http3.server_ssl_profile` | [virtual_server.http3.server_ssl_profile](data-sources--application_profiles--reference--group-003.md#canonical-2202321321031233-2131133023103213-1033011033022203-2110233230020231-0003201231231231-1123321331123200-3032133011212211-3033330020221003) |
| `virtual_server.http3.server_ssl_profile.kind` | [virtual_server.http3.server_ssl_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-0012302313310131-1103200323031103-0120323210001020-1301132123213310-2210330331101200-1130030112120121-2031233323002311-1121011332311312) |
| `virtual_server.http3.server_ssl_profile.name` | [virtual_server.http3.server_ssl_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-1032032130033032-0001002123020200-2320330323322011-2101111011000112-0212031120113311-1121131012021232-0202202313011301-2202300323013101) |
| `virtual_server.http3.server_ssl_profile.namespace` | [virtual_server.http3.server_ssl_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-2220313001030323-0113120120312012-0312303200101103-1033001000112200-1022312000013333-0232101221333330-3110131333312022-0032212001100100) |
| `virtual_server.http3.server_ssl_profile.tenant` | [virtual_server.http3.server_ssl_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-0121333110001302-3011222323013000-2320123123203122-2232230320201010-1233200001213303-2210332201313033-1233201001131230-3031332103113033) |
| `virtual_server.http3.server_ssl_profile.uid` | [virtual_server.http3.server_ssl_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-3321212230302200-3033132130011312-0223201231333020-2211031231201202-2212202010312300-3121032210113012-0232102002202013-1301020232030002) |
| `virtual_server.http3.tcp_server_profile` | [virtual_server.http3.tcp_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-2111211220132202-1113012320222200-1013130022202123-3033202211130113-0002201233323031-3303210132013210-2321202231101210-0233111122201231) |
| `virtual_server.http3.tcp_server_profile.kind` | [virtual_server.http3.tcp_server_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-2333132203103120-3312032123020120-0320323131013211-2323021212220322-3220223202233133-1210201130022213-3010220310120233-3132313323203001) |
| `virtual_server.http3.tcp_server_profile.name` | [virtual_server.http3.tcp_server_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-2321221311323011-2003030111232032-3233130311301302-0210122112230333-3133320213120232-2030132302320031-1223012120022102-1111300111011123) |
| `virtual_server.http3.tcp_server_profile.namespace` | [virtual_server.http3.tcp_server_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-2112112320211220-1031113122132000-0213222102131010-2010111132203013-0313310023333332-1033011133002230-3003220320003310-0302312203222222) |
| `virtual_server.http3.tcp_server_profile.tenant` | [virtual_server.http3.tcp_server_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-2321221222112321-3200021332030111-3123111333232112-0303003031221302-1002332202323233-3321232232331230-3311221121312121-2301231300000023) |
| `virtual_server.http3.tcp_server_profile.uid` | [virtual_server.http3.tcp_server_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-3331011033001003-1230331200212230-1022023101120323-1303102002021121-0131323333023332-0313223021002301-0123303202321220-0103011111120322) |
| `virtual_server.http3.udp_client_profile` | [virtual_server.http3.udp_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-3110111131023211-0330232112200211-3112303113213121-0121113033113231-3120200331110202-2211221310223221-2122023033133323-3120120311021221) |
| `virtual_server.http3.udp_client_profile.kind` | [virtual_server.http3.udp_client_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-2203200122311100-1333100112021001-3012100102130012-0020003230301123-0232310230310103-3330000102320121-0122232132322310-3003121302130021) |
| `virtual_server.http3.udp_client_profile.name` | [virtual_server.http3.udp_client_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-2113002232231032-2002133030000231-1203233301211311-3333013303123110-2020111212300030-3303230223011012-2313110012013321-2201113110023200) |
| `virtual_server.http3.udp_client_profile.namespace` | [virtual_server.http3.udp_client_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-2203103302203021-2303332233323320-2303101123233113-3000030102020003-3000100100131132-3022010120133112-3000033120011332-3231211002132301) |
| `virtual_server.http3.udp_client_profile.tenant` | [virtual_server.http3.udp_client_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-0200301032120213-3100013023102132-0132212020103313-0222322033333302-1233011110233112-1030311003100221-1110230100220110-0131110212023231) |
| `virtual_server.http3.udp_client_profile.uid` | [virtual_server.http3.udp_client_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-3133021312220110-2030102332311103-2020010201220102-1333302111113001-1133310131021113-2120300102023310-2000011100002313-1211201232232031) |
| `virtual_server.http3.udp_server_profile` | [virtual_server.http3.udp_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-0003321320221213-2220121121131132-0000313303113012-3330123230300312-3310311321000131-2222322030230303-3311223213331001-2212313203101203) |
| `virtual_server.http3.udp_server_profile.kind` | [virtual_server.http3.udp_server_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-0132022331223331-3013330311300132-0133002321200022-1323113301002103-0033112113333331-2300010221011213-1222202212213110-0320112301010203) |
| `virtual_server.http3.udp_server_profile.name` | [virtual_server.http3.udp_server_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-2121223101103300-1232331220320023-1310212030013311-1321000112232033-3200020302331210-1132030023212313-1223033320233300-2221003102023221) |
| `virtual_server.http3.udp_server_profile.namespace` | [virtual_server.http3.udp_server_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-3113011111203131-0000010322121200-2030300203201222-1221011232220100-3111121130312010-0122222132121130-1213013310112133-0021002211022202) |
| `virtual_server.http3.udp_server_profile.tenant` | [virtual_server.http3.udp_server_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-2113201101102202-0111211201203112-2332130131123121-2032111133310112-0220201320123330-0013220230131001-2030312000222333-0210110323330131) |
| `virtual_server.http3.udp_server_profile.uid` | [virtual_server.http3.udp_server_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-1231221011321021-3213212321120213-1011312332333333-3221100201313303-1320021001111033-3202003320012100-3112310302203103-1231022310231323) |
| `virtual_server.https` | [virtual_server.https](data-sources--application_profiles--reference--group-003.md#canonical-3110112100311110-2102131302131202-3130311013113200-3302111333312320-0231310322001110-1202132200130320-3130103300220103-2113202112113123) |
| `virtual_server.https.client_ssl_profile` | [virtual_server.https.client_ssl_profile](data-sources--application_profiles--reference--group-003.md#canonical-1132322120022333-0312030202333221-3321011032223110-3301111312201221-1130120122132222-1002223013031310-0302311312132232-2230003122102313) |
| `virtual_server.https.client_ssl_profile.kind` | [virtual_server.https.client_ssl_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-1110121312130230-2001103131200203-1002102313133131-3303021103132200-0130032231211223-1300220112302302-1310022311212221-1003213311001200) |
| `virtual_server.https.client_ssl_profile.name` | [virtual_server.https.client_ssl_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-3213033313130003-3221300102030120-0100313221202330-2311332320111010-1102210020102132-3000121221120320-2031021232010320-3210222111211002) |
| `virtual_server.https.client_ssl_profile.namespace` | [virtual_server.https.client_ssl_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-2033311002330301-2231010213033013-0030022112033122-3332110123102221-0113120203033013-1322003033232020-0311011023132232-3103020330101200) |
| `virtual_server.https.client_ssl_profile.tenant` | [virtual_server.https.client_ssl_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-0112111010121310-3102223002232002-1203013211000031-0230311212231010-3021331010111232-0011332030130232-0100211101100303-1331233201010010) |
| `virtual_server.https.client_ssl_profile.uid` | [virtual_server.https.client_ssl_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-3102103202302101-1123202102012103-3301002133331202-2120011031320023-1331100003303233-3322102211013001-3233002212032231-3020020120022222) |
| `virtual_server.https.http2_client_profile` | [virtual_server.https.http2_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-2300311230302121-0332301030310130-1130120012100320-2113332311032112-0221030233102130-1111102222010113-0033022301021033-2213113030330110) |
| `virtual_server.https.http2_client_profile.kind` | [virtual_server.https.http2_client_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-2203030312232013-3121232202213222-0130200302212031-1332233000303222-3133213310313232-1320020223300131-3103231333212211-2131010323301103) |
| `virtual_server.https.http2_client_profile.name` | [virtual_server.https.http2_client_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-3020012020102033-3031123323232103-1101332333121031-3310203103033000-0323113313003102-2231033220100231-1120320202220231-0021100003102003) |
| `virtual_server.https.http2_client_profile.namespace` | [virtual_server.https.http2_client_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-2031100001000231-1032101003301222-1130121203102332-2132333023131010-1313101131303212-3333120331003113-1210022321033001-2310122203121332) |
| `virtual_server.https.http2_client_profile.tenant` | [virtual_server.https.http2_client_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-2102323113323032-3303022213303210-3110011201110201-3102333032000120-3200233312023010-0222202311102321-3321200121212102-2331311230020032) |
| `virtual_server.https.http2_client_profile.uid` | [virtual_server.https.http2_client_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-0113223121300303-0230023313301202-0133133133313111-2113020333313130-3011122232313030-2032023212113113-3100203330213232-0200001101003223) |
| `virtual_server.https.http2_server_profile` | [virtual_server.https.http2_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-3032313001230132-3323121103023232-0030300210011030-1130303023121331-2203133002010120-0303200302222011-0302021130001313-3122222101303101) |
| `virtual_server.https.http2_server_profile.kind` | [virtual_server.https.http2_server_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-1202301132213331-1203021030233232-3130223333223111-3232102130330231-3032012001222110-2030210230301101-1100131103020110-3201103013230121) |
| `virtual_server.https.http2_server_profile.name` | [virtual_server.https.http2_server_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-3223330231301021-2100323011202103-0303120122312233-1130131121332113-3302210111222212-0310231232121113-0001022121001002-0313100230201000) |
| `virtual_server.https.http2_server_profile.namespace` | [virtual_server.https.http2_server_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-0033232312013133-0001112331212301-2033100130301223-3023303120230001-1331320012111232-0222311310203001-1103012201231303-3102002200100102) |
| `virtual_server.https.http2_server_profile.tenant` | [virtual_server.https.http2_server_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-3121201231013100-2201323213221001-2113310322022100-0232323033123122-1232133331201202-2133003212303330-2221320121331011-0100133130003133) |
| `virtual_server.https.http2_server_profile.uid` | [virtual_server.https.http2_server_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-0002203310232131-0133132223210032-3321200332310202-3002203210020131-0203032211210212-0000210210003031-3011322201110133-0231231033200011) |
| `virtual_server.https.http_client_profile` | [virtual_server.https.http_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-3131300113300020-0011113001200100-2121210322120013-3211213213122233-0003210302021133-0330020220231003-3331332200203012-1213322003103003) |
| `virtual_server.https.http_client_profile.kind` | [virtual_server.https.http_client_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-0311231031223312-1003233110001021-1330321330102201-3310001122320102-0233220002200231-1021012000031103-0223003332101120-1133323203202031) |
| `virtual_server.https.http_client_profile.name` | [virtual_server.https.http_client_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-3210120202130120-3122200123122223-2101031210003201-1231020113131103-1120112130330211-2313010312302330-0223121021321122-1120033321301012) |
| `virtual_server.https.http_client_profile.namespace` | [virtual_server.https.http_client_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-0113320320112000-2022020213132233-3201323212030132-3212202201011120-1021323133120231-1130330123031211-3123230000032131-1133011021322211) |
| `virtual_server.https.http_client_profile.tenant` | [virtual_server.https.http_client_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-3013122132032220-3300200112313121-2133301131222123-3100233121013313-2331231100323012-3101110000001020-3023333113213212-2101030100013002) |
| `virtual_server.https.http_client_profile.uid` | [virtual_server.https.http_client_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-0212230031010012-0201111222230021-1302231113311132-3133123131023130-0021121001231123-1132301212233032-2112022201202112-0000022132301230) |
| `virtual_server.https.http_server_profile` | [virtual_server.https.http_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-3132302132312210-2020323020222112-0001200000121223-0212232330220030-0003202223331003-2221331033000110-3033303120211121-0231001312112020) |
| `virtual_server.https.http_server_profile.kind` | [virtual_server.https.http_server_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-2313123202213030-3331122332020222-0231131311022311-2030221320303210-3111033333222111-2301100010102000-1303122020202000-1201232312322311) |
| `virtual_server.https.http_server_profile.name` | [virtual_server.https.http_server_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-1033133132012220-0313302002231332-0322332232330211-3013130103131212-3133331010033001-2113322113310302-3321011000102031-2003003101331213) |
| `virtual_server.https.http_server_profile.namespace` | [virtual_server.https.http_server_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-0203121001132023-0203020112110320-3102230001231011-2233302203012112-1031230303112010-3202203103000223-0213013312323102-0101212220333122) |
| `virtual_server.https.http_server_profile.tenant` | [virtual_server.https.http_server_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-2210223023312222-3232112103301023-3000033211200110-3223131332100020-3001303110110331-0003102012120113-0102333211110310-0112210233032223) |
| `virtual_server.https.http_server_profile.uid` | [virtual_server.https.http_server_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-1012232303223130-1311121212303020-0211321020020313-3000230031033011-0123023111100110-0102021101333001-0011130133022030-1212310110012030) |
| `virtual_server.https.ocsp_profile` | [virtual_server.https.ocsp_profile](data-sources--application_profiles--reference--group-003.md#canonical-2013002101022013-2331200321301103-1321203302300002-0113011110003003-2233212212030320-2102000221123111-0330031003132102-1112012323101021) |
| `virtual_server.https.ocsp_profile.kind` | [virtual_server.https.ocsp_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-0333132003010323-1120331001331123-3310322213323130-0001032331212300-1133120131302323-3033121313311120-0211223222112230-3230022231333202) |
| `virtual_server.https.ocsp_profile.name` | [virtual_server.https.ocsp_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-2230023313120100-1330310233232010-1320231330332233-0103000301232031-0130023113011302-2232132112201320-1012220320013001-1032232031122212) |
| `virtual_server.https.ocsp_profile.namespace` | [virtual_server.https.ocsp_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-0331122031032201-2023303021010123-2333311332200113-2011112031123032-3302313123333113-0012233132033133-1030012303030221-3220113120200330) |
| `virtual_server.https.ocsp_profile.tenant` | [virtual_server.https.ocsp_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-3322233310021320-1312220233121113-1020031232221023-3201103031032002-3210033330122110-1013321200010323-0000100322132202-0013012333221102) |
| `virtual_server.https.ocsp_profile.uid` | [virtual_server.https.ocsp_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-3110303230120211-2032002231223301-1320102101230200-2002000300101110-1110113211211010-0030300222021022-1222112133013200-1300000130031200) |
| `virtual_server.https.server_ssl_profile` | [virtual_server.https.server_ssl_profile](data-sources--application_profiles--reference--group-003.md#canonical-3303311223330312-3033332311203113-2213312330332231-2300030212111132-2232201300002011-2011301203332012-3231333312121132-0112202100122202) |
| `virtual_server.https.server_ssl_profile.kind` | [virtual_server.https.server_ssl_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-2003120003221002-1223130000200002-1311221231030330-3323013323220100-2001121132220131-0102313123312332-1001220333133001-3021133133120310) |
| `virtual_server.https.server_ssl_profile.name` | [virtual_server.https.server_ssl_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-1323230312222023-1102033003233303-3033113201322003-2301312203332210-1220230321310322-2031333012333011-0321233332123113-3101220303303233) |
| `virtual_server.https.server_ssl_profile.namespace` | [virtual_server.https.server_ssl_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-0000120220131302-3312322122131013-2230331032101122-2001222121232320-3223201220221111-3231121103122201-3332022131220100-2300223323101021) |
| `virtual_server.https.server_ssl_profile.tenant` | [virtual_server.https.server_ssl_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-3203330132232233-1123323113131011-2211210121320212-0102303120111012-3200212011203330-3122333003123130-1223213321023120-3222101303211133) |
| `virtual_server.https.server_ssl_profile.uid` | [virtual_server.https.server_ssl_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-1131310220201110-0200002213220302-0211122122322001-3121111100222010-1021010002120130-2200013233333033-2011122113132300-2300020121320232) |
| `virtual_server.https.stream_profile` | [virtual_server.https.stream_profile](data-sources--application_profiles--reference--group-003.md#canonical-2101112123203110-0231221321022130-3103201012031200-2300212020012011-1230032322101301-3311322301203000-1102100102000233-0331221321313011) |
| `virtual_server.https.stream_profile.kind` | [virtual_server.https.stream_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-2121222122231220-3003121210123321-3022213100303002-2112102222210331-0122110232203201-2020110010230013-3310202321301001-1100211213332302) |
| `virtual_server.https.stream_profile.name` | [virtual_server.https.stream_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-1221211022022220-2320212233303101-1212300121102300-1321012331102123-0003011221003303-1033023031111032-3123122100121202-2200202113010331) |
| `virtual_server.https.stream_profile.namespace` | [virtual_server.https.stream_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-1323301301310000-2200120032213330-1233200232021002-3102101333111003-0032113310102103-3322013311201003-1333210302001003-3210030121102231) |
| `virtual_server.https.stream_profile.tenant` | [virtual_server.https.stream_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-3112011222113010-1100200331000332-0300133333112322-3020010302200310-0022133213200113-2022023232322223-3023223231113201-0211020312011022) |
| `virtual_server.https.stream_profile.uid` | [virtual_server.https.stream_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-0032123333333230-0013230113220010-1313011122213332-2220233113113320-2323333000130303-3202120010003011-2211030223311031-3102301001030012) |
| `virtual_server.https.tcp_client_profile` | [virtual_server.https.tcp_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-3311100221020133-0023232010132211-2010011331121213-0031000022333011-2010010233030320-1311003313310322-2210321223120132-2003213033101022) |
| `virtual_server.https.tcp_client_profile.kind` | [virtual_server.https.tcp_client_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-3000013103023010-2330212310323033-3331201020222200-2300032002011121-1030111032321030-3123233320330320-1120003212313332-3223330333223231) |
| `virtual_server.https.tcp_client_profile.name` | [virtual_server.https.tcp_client_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-1103011220011233-0321202113110021-1233000302213123-1330032120031301-1032103112102200-3213311102120310-2131001220211321-1312231201322121) |
| `virtual_server.https.tcp_client_profile.namespace` | [virtual_server.https.tcp_client_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-3232010022121310-1211101223223122-3211031002300112-2221020122011223-2223232313022020-1000003023320233-3332211113001112-1133001100310232) |
| `virtual_server.https.tcp_client_profile.tenant` | [virtual_server.https.tcp_client_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-1010202122201212-3223101002323311-3323013011303332-2201003313100330-0002331213203202-0210133211322131-2013211312212333-0010231013133300) |
| `virtual_server.https.tcp_client_profile.uid` | [virtual_server.https.tcp_client_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-0020322001231033-0020201001103000-1130013231123021-1213213011110002-2313300100001021-3202223011030302-1003323133110233-2100012113311032) |
| `virtual_server.https.tcp_server_profile` | [virtual_server.https.tcp_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-2333122203230102-3331000211011113-0031023332320321-1200000031100101-3333133213321022-0131033020022213-2301223321102130-1311030113100202) |
| `virtual_server.https.tcp_server_profile.kind` | [virtual_server.https.tcp_server_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-0013012323002312-2102331123231221-3321222103101202-3223212101113302-0001220321120300-1313303203230220-3100212321213302-2032130012332223) |
| `virtual_server.https.tcp_server_profile.name` | [virtual_server.https.tcp_server_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-0101010133201312-1311303112221210-0313101110330203-0121102123023200-0311013002111200-3200231221011232-3330311203333100-1000320201312101) |
| `virtual_server.https.tcp_server_profile.namespace` | [virtual_server.https.tcp_server_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-3213110330021022-1300303210103311-1200130123220003-2223303111133100-0022030323301221-3303203201230303-1100221103300031-3132030201310202) |
| `virtual_server.https.tcp_server_profile.tenant` | [virtual_server.https.tcp_server_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-3332321010023303-2101121321323103-3022220032202230-3223230130111301-2311310131032001-3123030223220033-3010200221002202-0203201330120022) |
| `virtual_server.https.tcp_server_profile.uid` | [virtual_server.https.tcp_server_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-1223313010313313-2333222130310231-3023011022002022-3203010033002221-0201112013000013-1102301122020330-3111031123131230-0200330302303321) |
| `virtual_server.https.websocket_client_profile` | [virtual_server.https.websocket_client_profile](data-sources--application_profiles--reference--group-003.md#canonical-3110112022222213-3031103003231113-2210033100330223-0310132300103132-0223331232220230-0023202223300301-2213232333332010-2300321331131331) |
| `virtual_server.https.websocket_client_profile.kind` | [virtual_server.https.websocket_client_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-2131032331210311-0021210032233031-2212110110323130-2003022003203300-0220211333330021-1213012133012212-2313323222131102-1200120200323202) |
| `virtual_server.https.websocket_client_profile.name` | [virtual_server.https.websocket_client_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-3330323230322020-2012031130030030-2113120330302100-0213200231332113-1013003120101320-2233132220023213-2213003311230130-3133031110103210) |
| `virtual_server.https.websocket_client_profile.namespace` | [virtual_server.https.websocket_client_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-0110203121212333-2032333011013122-2231111302201211-0023122303121011-2120131233032302-3100031303021212-2210132220213330-1231322001022102) |
| `virtual_server.https.websocket_client_profile.tenant` | [virtual_server.https.websocket_client_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-1010103130330300-3233000010210011-2300111333231231-0002203032303110-0323123333311023-3210313102301312-1222113131321323-1003020200223120) |
| `virtual_server.https.websocket_client_profile.uid` | [virtual_server.https.websocket_client_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-1303232103101303-3000223033320023-1032313332322333-2002313232312132-3023100201120110-0112023020122122-1032321302023332-3332313012112301) |
| `virtual_server.https.websocket_server_profile` | [virtual_server.https.websocket_server_profile](data-sources--application_profiles--reference--group-003.md#canonical-2012111301231322-1111322030030322-2111202200121021-1030233113133301-1101212311113022-0101312003203001-0210220000030302-3210332010000033) |
| `virtual_server.https.websocket_server_profile.kind` | [virtual_server.https.websocket_server_profile.kind](data-sources--application_profiles--reference--group-003.md#canonical-0113033232120111-1002201022313210-1030311130010023-0330023210033321-0321320202020100-0113311232230321-0011102320232210-3231213010232232) |
| `virtual_server.https.websocket_server_profile.name` | [virtual_server.https.websocket_server_profile.name](data-sources--application_profiles--reference--group-003.md#canonical-1021000021011333-1233331232122330-0220113002301313-1002332003100013-0122220032220202-1032332031021002-1332013333133021-1232312100301023) |
| `virtual_server.https.websocket_server_profile.namespace` | [virtual_server.https.websocket_server_profile.namespace](data-sources--application_profiles--reference--group-003.md#canonical-1231220022320220-3302223200010122-0200303000022131-1112222231322113-1203111323210202-0101212310121011-1110132012210213-2303211230123023) |
| `virtual_server.https.websocket_server_profile.tenant` | [virtual_server.https.websocket_server_profile.tenant](data-sources--application_profiles--reference--group-003.md#canonical-0012033122103103-3012233012033213-0001202231333010-0301331321100133-2002120120102011-2212122300310322-2123102231131002-1031133331211032) |
| `virtual_server.https.websocket_server_profile.uid` | [virtual_server.https.websocket_server_profile.uid](data-sources--application_profiles--reference--group-003.md#canonical-0200303012122222-2213200013030120-2112010231302332-3023210330121221-1212123022222013-2200211110011231-0331213032113332-3031132301122001) |
| `virtual_server.immediate_action_on_service_down` | [virtual_server.immediate_action_on_service_down](data-sources--application_profiles--reference--group-003.md#canonical-0210331121211012-2203311120332113-1110330020220232-1213101202321223-2133202200212011-3301130210013120-3000113220233111-2331332203033110) |
| `virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_drop` | [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_drop](data-sources--application_profiles--reference--group-003.md#canonical-3233003231010201-1213111003322323-2221030030111111-3203212212203322-2322001311323312-0200002223310232-2310212013022231-1232232032330123) |
| `virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_none` | [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_none](data-sources--application_profiles--reference--group-003.md#canonical-3000331313012003-1313113120221112-1321320333220223-3223310230103233-3322133303302333-0133332030023102-1120203222130030-3333112121013130) |
| `virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset` | [virtual_server.immediate_action_on_service_down.immediate_action_on_service_down_reset](data-sources--application_profiles--reference--group-003.md#canonical-3031312321022010-3333033133022131-3100131033230303-0330010323330111-3111100013013201-0013013212012130-1313113331031321-3002313232001233) |
| `virtual_server.last_hop_pool` | [virtual_server.last_hop_pool](data-sources--application_profiles--reference--group-003.md#canonical-1331332331330133-1331300103320202-1223101301311323-3321312330210100-2311310201321032-0223302031132302-3110001232223011-0133002310032201) |
| `virtual_server.last_hop_pool.kind` | [virtual_server.last_hop_pool.kind](data-sources--application_profiles--reference--group-003.md#canonical-3320230021230312-1112302231000133-3132223210220133-1333213323222103-2023332210021111-0103333110320202-2121301203201323-1310211133223332) |
| `virtual_server.last_hop_pool.name` | [virtual_server.last_hop_pool.name](data-sources--application_profiles--reference--group-003.md#canonical-1100132102022032-0321130101023003-1222000130122110-2213212022000021-3122321123032201-0033210123000313-0322011202122302-0203100331021230) |
| `virtual_server.last_hop_pool.namespace` | [virtual_server.last_hop_pool.namespace](data-sources--application_profiles--reference--group-003.md#canonical-1022223331012211-3012300010010122-1133012302110113-2111323031231331-3333221002320111-1131010121320023-2223003313103221-3213220311111132) |
| `virtual_server.last_hop_pool.tenant` | [virtual_server.last_hop_pool.tenant](data-sources--application_profiles--reference--group-003.md#canonical-0322032110112232-1120333001010313-2222111031233223-3300212301332020-0232023133213333-0010200330202303-0022102030021102-3312021210321200) |
| `virtual_server.last_hop_pool.uid` | [virtual_server.last_hop_pool.uid](data-sources--application_profiles--reference--group-003.md#canonical-0222231132230310-3110230303320121-1133232113122013-1301213010002020-3310012202021002-1233230323311331-2202231211123133-2030012323232211) |
| `virtual_server.nat64` | [virtual_server.nat64](data-sources--application_profiles--reference--group-003.md#canonical-3112222212333220-2111210012320301-1230013102331102-1113102012213010-1223310022222321-0001023002031323-0231201223233030-0321201011113133) |
| `virtual_server.nat64.nat64_disable` | [virtual_server.nat64.nat64_disable](data-sources--application_profiles--reference--group-004.md#canonical-3330021000232332-3112120221301132-1202112021311233-1002102220223223-1323103112322321-3020330031020301-1010001112301333-3210122112321303) |
| `virtual_server.nat64.nat64_enable` | [virtual_server.nat64.nat64_enable](data-sources--application_profiles--reference--group-004.md#canonical-2323330300333113-0111210302021122-3222102210220312-3120003231033321-1223002012201122-2212212232011330-0022023201212313-2213003331011323) |
| `virtual_server.port_translation` | [virtual_server.port_translation](data-sources--application_profiles--reference--group-004.md#canonical-0230012221221331-3021110120120210-2301323100032210-0112002110033330-3020100332131003-3320320221322132-3320012303020331-0002233212311013) |
| `virtual_server.port_translation.port_translation_disable` | [virtual_server.port_translation.port_translation_disable](data-sources--application_profiles--reference--group-004.md#canonical-1223333011230223-3003313223102211-3213332123113011-2011113321311333-1131201300001232-3222010100131121-3323120023222032-1210311122222011) |
| `virtual_server.port_translation.port_translation_enable` | [virtual_server.port_translation.port_translation_enable](data-sources--application_profiles--reference--group-004.md#canonical-1123021020332223-3003002223232331-2000213010302201-0232230211330103-1311212121233320-2022302310202033-1333022301001012-0122310131120012) |
| `virtual_server.request_logging_profile` | [virtual_server.request_logging_profile](data-sources--application_profiles--reference--group-004.md#canonical-1031110300301322-3011012111331131-3320322013310021-1331332303130013-0333231103331122-3012333020301221-1133212222013310-2220232130012002) |
| `virtual_server.request_logging_profile.kind` | [virtual_server.request_logging_profile.kind](data-sources--application_profiles--reference--group-004.md#canonical-3111011233010123-1322230210312302-3110220021311312-0123010212232222-1221021223020311-2233320201303112-3213221001323133-2222331001120331) |
| `virtual_server.request_logging_profile.name` | [virtual_server.request_logging_profile.name](data-sources--application_profiles--reference--group-004.md#canonical-0032131021213201-0032210012001322-1022033201203132-0221223233220321-1230311221001013-3311132303202130-1002223322112110-0211321030200212) |
| `virtual_server.request_logging_profile.namespace` | [virtual_server.request_logging_profile.namespace](data-sources--application_profiles--reference--group-004.md#canonical-2333231231202222-3000120200101013-2333122312302300-1010300001320031-1312000333330232-3020130310023312-2300030113133222-1332021211122210) |
| `virtual_server.request_logging_profile.tenant` | [virtual_server.request_logging_profile.tenant](data-sources--application_profiles--reference--group-004.md#canonical-2321321001100300-2222002201310132-0312100012233022-1110131130102021-1202103112220223-1310120111103121-1312203321021203-0323333131013303) |
| `virtual_server.request_logging_profile.uid` | [virtual_server.request_logging_profile.uid](data-sources--application_profiles--reference--group-004.md#canonical-1211021330212112-0113010233101110-3213130122212220-3122221003303033-3220003021232020-0230220003222200-0120231211301201-3100113311233200) |
| `virtual_server.source_port` | [virtual_server.source_port](data-sources--application_profiles--reference--group-004.md#canonical-2231031203231101-1320022022331310-2302331331320012-0013323110030013-0322230203211311-0231131003221110-2123031233033330-3020101331100032) |
| `virtual_server.source_port.source_port_change` | [virtual_server.source_port.source_port_change](data-sources--application_profiles--reference--group-004.md#canonical-1033230100021110-1232321310030300-0013311113121012-3131323201213223-3131121020222213-2123022303233030-0200332113031001-0112011010112101) |
| `virtual_server.source_port.source_port_preserve` | [virtual_server.source_port.source_port_preserve](data-sources--application_profiles--reference--group-004.md#canonical-3122210132111330-3213033313322121-2230323132323233-0202202102221001-2303103302013330-1203220001233003-0133322023132333-3231231023120323) |
| `virtual_server.source_port.source_port_preserve_strict` | [virtual_server.source_port.source_port_preserve_strict](data-sources--application_profiles--reference--group-004.md#canonical-2103231131103021-2223110201303233-3123113133113121-2130101121221022-2030110102132300-0102210010213033-2201112112020000-0330100113133120) |
| `virtual_server.statistics_profile` | [virtual_server.statistics_profile](data-sources--application_profiles--reference--group-004.md#canonical-2030120322312202-1311203300012321-0321320031013201-3233002122323031-2033010022232303-2110113120103221-1303322002232120-0232013212222013) |
| `virtual_server.statistics_profile.kind` | [virtual_server.statistics_profile.kind](data-sources--application_profiles--reference--group-004.md#canonical-2302010102032211-2101220322000120-1122102130000030-0331230231012103-2033003332311120-1211011002101230-2122113311020032-3233202230133301) |
| `virtual_server.statistics_profile.name` | [virtual_server.statistics_profile.name](data-sources--application_profiles--reference--group-004.md#canonical-1210213022223212-3133101111210020-3322303133122300-0123123013113330-2003233011230003-3122232322221131-1113313133211010-2201303221233131) |
| `virtual_server.statistics_profile.namespace` | [virtual_server.statistics_profile.namespace](data-sources--application_profiles--reference--group-004.md#canonical-3233300113322310-1133113133232223-1230022323013133-1211132212130013-2113321221200310-3121033032102200-0220012013130010-0221021211333210) |
| `virtual_server.statistics_profile.tenant` | [virtual_server.statistics_profile.tenant](data-sources--application_profiles--reference--group-004.md#canonical-3200032003300132-3303302103202001-2000030303131112-3100100220110121-3330100231021011-2333200103022023-0302221203020032-2031313331032301) |
| `virtual_server.statistics_profile.uid` | [virtual_server.statistics_profile.uid](data-sources--application_profiles--reference--group-004.md#canonical-3101021230321013-1312011302130111-3030332132013102-0111212322122021-3203222232012013-1112302322213300-3330331220232323-2233030123302330) |
| `virtual_server.tcp` | [virtual_server.tcp](data-sources--application_profiles--reference--group-004.md#canonical-3332310323031233-0313321330000232-3300010012123032-3321003311001223-1002000031011012-1221013001120331-2103101021011301-1010110123323100) |
| `virtual_server.tcp.client_ssl_profile` | [virtual_server.tcp.client_ssl_profile](data-sources--application_profiles--reference--group-004.md#canonical-1223323031103032-1333301010223312-0323313113233102-2211030313223003-2112002133300232-3131120003201031-3300020231132200-2200333311102021) |
| `virtual_server.tcp.client_ssl_profile.kind` | [virtual_server.tcp.client_ssl_profile.kind](data-sources--application_profiles--reference--group-004.md#canonical-1220010130112121-3220233320100101-3013123332101221-3221133232033011-0022001132301033-2210111002022001-1312132313210102-3332110020230311) |
| `virtual_server.tcp.client_ssl_profile.name` | [virtual_server.tcp.client_ssl_profile.name](data-sources--application_profiles--reference--group-004.md#canonical-1201232331100312-2231133323321300-1101222002301001-0300120320023002-3312310012202030-2312332002222021-2213223111232022-1122310001101202) |
| `virtual_server.tcp.client_ssl_profile.namespace` | [virtual_server.tcp.client_ssl_profile.namespace](data-sources--application_profiles--reference--group-004.md#canonical-0122021320113200-3302220020122013-3130330102311221-3032220300022223-2333010320122220-2033103132200010-0323310333210302-2213121032322330) |
| `virtual_server.tcp.client_ssl_profile.tenant` | [virtual_server.tcp.client_ssl_profile.tenant](data-sources--application_profiles--reference--group-004.md#canonical-1012013322002021-2211211113222003-2032221301132030-3223122332033031-3311002131012113-1300301110330220-1032312100101300-0001322103310122) |
| `virtual_server.tcp.client_ssl_profile.uid` | [virtual_server.tcp.client_ssl_profile.uid](data-sources--application_profiles--reference--group-004.md#canonical-2121012201112312-0002001312031121-3201331313001310-2031032313001331-2023333220313133-2013320223201000-2101002012103321-0323010233232122) |
| `virtual_server.tcp.ocsp_profile` | [virtual_server.tcp.ocsp_profile](data-sources--application_profiles--reference--group-004.md#canonical-3222113202322322-0331222010032111-0231200332330131-1120010211000200-1320103023320021-3020032311122013-3012321300232301-1212233311202030) |
| `virtual_server.tcp.ocsp_profile.kind` | [virtual_server.tcp.ocsp_profile.kind](data-sources--application_profiles--reference--group-004.md#canonical-2133103031213311-1023200220030223-0300213312300320-3033203323222030-3003022112303203-3322032133130302-0203212000000210-3210320201220332) |
| `virtual_server.tcp.ocsp_profile.name` | [virtual_server.tcp.ocsp_profile.name](data-sources--application_profiles--reference--group-004.md#canonical-1303302102302030-0100000232300311-2210102233131213-3023320231311100-0230322133121211-0012123230110220-1211021233111031-3332133132300322) |
| `virtual_server.tcp.ocsp_profile.namespace` | [virtual_server.tcp.ocsp_profile.namespace](data-sources--application_profiles--reference--group-004.md#canonical-3130230020320011-3103103201100321-1203031032010003-0133021210001121-0021231011103331-3233220302011033-3212223022020031-2202112331033011) |
| `virtual_server.tcp.ocsp_profile.tenant` | [virtual_server.tcp.ocsp_profile.tenant](data-sources--application_profiles--reference--group-004.md#canonical-0322021303013323-0202320023211102-3100101303010000-2210122222101322-2322233231000131-2303121110023200-1220332322133203-1222233023231233) |
| `virtual_server.tcp.ocsp_profile.uid` | [virtual_server.tcp.ocsp_profile.uid](data-sources--application_profiles--reference--group-004.md#canonical-3233130311023002-2222211223202010-1211231210210222-1320312111330103-3231201020300111-3033322213122232-0100232132220223-1321323302211202) |
| `virtual_server.tcp.server_ssl_profile` | [virtual_server.tcp.server_ssl_profile](data-sources--application_profiles--reference--group-004.md#canonical-1301210030113010-3001031000022023-1203202300132011-2320223221023320-0201113030211233-1113112203320131-0321103311130200-2313302302001132) |
| `virtual_server.tcp.server_ssl_profile.kind` | [virtual_server.tcp.server_ssl_profile.kind](data-sources--application_profiles--reference--group-004.md#canonical-3123220302001003-0221110312110212-3010023023031232-0100021213122233-0231101221102032-3010020223123303-3201231220101332-3233230022333100) |
| `virtual_server.tcp.server_ssl_profile.name` | [virtual_server.tcp.server_ssl_profile.name](data-sources--application_profiles--reference--group-004.md#canonical-3220023021330010-1022331102300302-1130310302211113-2203331201322310-0233211102323232-1103213012203022-2011230332230113-0132221121220122) |
| `virtual_server.tcp.server_ssl_profile.namespace` | [virtual_server.tcp.server_ssl_profile.namespace](data-sources--application_profiles--reference--group-004.md#canonical-3111223102012123-0203023033221220-0332303002133231-2031010210000102-1020011001012231-1102000000113011-3002330230133202-1213201203331323) |
| `virtual_server.tcp.server_ssl_profile.tenant` | [virtual_server.tcp.server_ssl_profile.tenant](data-sources--application_profiles--reference--group-004.md#canonical-2022110203301233-2233230031101202-1300223210332300-3321000123200220-1131320332213333-2202311133113200-2133233021110212-0032013030303331) |
| `virtual_server.tcp.server_ssl_profile.uid` | [virtual_server.tcp.server_ssl_profile.uid](data-sources--application_profiles--reference--group-004.md#canonical-0103200032003121-2303201330033021-0032320131330211-3131000310203303-3233202112120102-1231321133321301-3102221230101311-1003122022310022) |
| `virtual_server.tcp.tcp_client_profile` | [virtual_server.tcp.tcp_client_profile](data-sources--application_profiles--reference--group-004.md#canonical-0233130021010220-1121023203223312-3032230023001021-2100113110302023-2010233012010211-1320211021101331-2021110312323310-1030121232022110) |
| `virtual_server.tcp.tcp_client_profile.kind` | [virtual_server.tcp.tcp_client_profile.kind](data-sources--application_profiles--reference--group-004.md#canonical-3032220312231213-2300232121210230-2212123331131322-1031321232021130-3023223320313113-2021331121331212-1331312010331211-2032022210121023) |
| `virtual_server.tcp.tcp_client_profile.name` | [virtual_server.tcp.tcp_client_profile.name](data-sources--application_profiles--reference--group-004.md#canonical-0303213032221112-1100121300200200-1320222322310231-3123002211221131-1230311310232123-3131301212131303-1102221213030213-2031300131331031) |
| `virtual_server.tcp.tcp_client_profile.namespace` | [virtual_server.tcp.tcp_client_profile.namespace](data-sources--application_profiles--reference--group-004.md#canonical-2001200320213233-1120030103320132-1032112022111031-3211100010100330-2302333212033031-2133231220001010-2233111313032112-3212032333310201) |
| `virtual_server.tcp.tcp_client_profile.tenant` | [virtual_server.tcp.tcp_client_profile.tenant](data-sources--application_profiles--reference--group-004.md#canonical-1212100302003313-0331210323332030-1201303123012230-0030221332203201-3001113132301312-3101032331100222-3233223310013113-1110332333101122) |
| `virtual_server.tcp.tcp_client_profile.uid` | [virtual_server.tcp.tcp_client_profile.uid](data-sources--application_profiles--reference--group-004.md#canonical-0233221223033011-1212023121031111-0321033132333302-0003331011310333-1113023121012311-1020210031000302-1221023202000030-0220012210312210) |
| `virtual_server.tcp.tcp_server_profile` | [virtual_server.tcp.tcp_server_profile](data-sources--application_profiles--reference--group-004.md#canonical-2223101012002002-0211221222230013-2203113222222221-2000010023113232-0022101110310031-3312130211212133-3101130033321322-2112213130333221) |
| `virtual_server.tcp.tcp_server_profile.kind` | [virtual_server.tcp.tcp_server_profile.kind](data-sources--application_profiles--reference--group-004.md#canonical-2331002323321112-1113033321031011-1321232120031023-1023222002103233-2021021130003312-0301112010210103-0120110110023201-2002121230011020) |
| `virtual_server.tcp.tcp_server_profile.name` | [virtual_server.tcp.tcp_server_profile.name](data-sources--application_profiles--reference--group-004.md#canonical-2003031013322331-0100100330122133-0303211313113032-2110233133231212-3012313200011313-2031020020322202-1020332111121023-1223121200213100) |
| `virtual_server.tcp.tcp_server_profile.namespace` | [virtual_server.tcp.tcp_server_profile.namespace](data-sources--application_profiles--reference--group-004.md#canonical-3322220201010210-0203220112133211-1120123201023000-0231313101033100-1332111013202002-0310230131120331-2321312000120302-2102230102110023) |
| `virtual_server.tcp.tcp_server_profile.tenant` | [virtual_server.tcp.tcp_server_profile.tenant](data-sources--application_profiles--reference--group-004.md#canonical-0321123000011210-2010112330021131-1331101313300303-2310210120100000-1323223202001223-1130031103303331-1222032231310200-2101120201110231) |
| `virtual_server.tcp.tcp_server_profile.uid` | [virtual_server.tcp.tcp_server_profile.uid](data-sources--application_profiles--reference--group-004.md#canonical-3112202321320130-0302202312113323-0330303200001100-3323010223012021-2212010221311202-3223032212230232-1231023121312230-0110130011003100) |
| `virtual_server.udp` | [virtual_server.udp](data-sources--application_profiles--reference--group-004.md#canonical-3230102230310120-1010023200321221-3133023233101012-0120203133222230-3233210233210103-1332031102321121-2231002312110123-2021321301012232) |
| `virtual_server.udp.client_ssl_profile` | [virtual_server.udp.client_ssl_profile](data-sources--application_profiles--reference--group-004.md#canonical-2223111031211133-0101333002323221-1330302200031101-2132112223123222-2201122023213030-1020321313300011-1010133331202031-3230300310130112) |
| `virtual_server.udp.client_ssl_profile.kind` | [virtual_server.udp.client_ssl_profile.kind](data-sources--application_profiles--reference--group-004.md#canonical-0322301012103310-2211210202321302-2220100202033221-3123231121221030-2222333013100200-2223332112112210-1131021123131111-3231210003232133) |
| `virtual_server.udp.client_ssl_profile.name` | [virtual_server.udp.client_ssl_profile.name](data-sources--application_profiles--reference--group-004.md#canonical-0132103230323301-2310210101023022-3221100322330230-2020223323320212-0212212311121321-0113221020133331-2303003101010200-2122303332230211) |
| `virtual_server.udp.client_ssl_profile.namespace` | [virtual_server.udp.client_ssl_profile.namespace](data-sources--application_profiles--reference--group-004.md#canonical-1210103233123230-0011200330100300-3122332312110310-0233010331020013-1122002212112101-1212203332320202-1321202211323003-1102301221013010) |
| `virtual_server.udp.client_ssl_profile.tenant` | [virtual_server.udp.client_ssl_profile.tenant](data-sources--application_profiles--reference--group-004.md#canonical-0211311122032003-3222130111113002-3101000323101010-2021113023000331-2111000102130221-2111120203222011-1100013033011023-0203132122212300) |
| `virtual_server.udp.client_ssl_profile.uid` | [virtual_server.udp.client_ssl_profile.uid](data-sources--application_profiles--reference--group-004.md#canonical-2211100301020302-2320012111132033-1201303120022021-1220213100122220-2012021300310132-0233030333313032-2011213112313003-0231302123123110) |
| `virtual_server.udp.server_ssl_profile` | [virtual_server.udp.server_ssl_profile](data-sources--application_profiles--reference--group-004.md#canonical-0101033012301333-2330013230023300-1030200021120003-3111233103113030-1102102100332031-2131121031322222-2213131312112330-2301020232113100) |
| `virtual_server.udp.server_ssl_profile.kind` | [virtual_server.udp.server_ssl_profile.kind](data-sources--application_profiles--reference--group-004.md#canonical-2332323020202202-2301010013113311-3201203231333312-0131221313130303-2301001121121100-0303113122120330-0132303133110002-2020321130031112) |
| `virtual_server.udp.server_ssl_profile.name` | [virtual_server.udp.server_ssl_profile.name](data-sources--application_profiles--reference--group-004.md#canonical-2221303121101223-0133032033012012-3313000223203101-3331221001133121-2010132220010230-0231313311333303-1323002202323311-1301012013102131) |
| `virtual_server.udp.server_ssl_profile.namespace` | [virtual_server.udp.server_ssl_profile.namespace](data-sources--application_profiles--reference--group-004.md#canonical-1100203111120031-2120323223132010-2301303001222002-3012020030133313-1213312200202111-0210120101322223-2301011312021011-2000111012011011) |
| `virtual_server.udp.server_ssl_profile.tenant` | [virtual_server.udp.server_ssl_profile.tenant](data-sources--application_profiles--reference--group-004.md#canonical-1223122120321102-2000203333123003-1322122023312212-3000031011123100-3212313310230210-2312232300021021-2013231031020232-1302021022310010) |
| `virtual_server.udp.server_ssl_profile.uid` | [virtual_server.udp.server_ssl_profile.uid](data-sources--application_profiles--reference--group-004.md#canonical-2231130232100022-3123001103220031-2033212310210310-2000023222321131-0000302302133013-2212322120103331-0210122331000112-0102011101103103) |
| `virtual_server.udp.udp_client_profile` | [virtual_server.udp.udp_client_profile](data-sources--application_profiles--reference--group-004.md#canonical-1022022013310211-3131023311012011-0332223013233330-0310123313002302-3003030100013303-1331103101100223-0002131121201323-2223000021033100) |
| `virtual_server.udp.udp_client_profile.kind` | [virtual_server.udp.udp_client_profile.kind](data-sources--application_profiles--reference--group-004.md#canonical-3331321213300032-2002032122201223-3132221303332332-3112013011102232-2303100021103021-1210120010303210-3203033000020233-0323212333002033) |
| `virtual_server.udp.udp_client_profile.name` | [virtual_server.udp.udp_client_profile.name](data-sources--application_profiles--reference--group-004.md#canonical-1320223331022233-0211201230233120-0223002213100200-3210120200222330-2133220001131202-2031210132300313-2232233032130111-3320131033220202) |
| `virtual_server.udp.udp_client_profile.namespace` | [virtual_server.udp.udp_client_profile.namespace](data-sources--application_profiles--reference--group-004.md#canonical-2303121332220020-2200102003102322-3210233320133012-3021031022032321-2130102310302211-3320012122133032-1213132002202031-3321312330233131) |
| `virtual_server.udp.udp_client_profile.tenant` | [virtual_server.udp.udp_client_profile.tenant](data-sources--application_profiles--reference--group-004.md#canonical-1111313003330213-2112303201023030-0332303020331322-0221001122220130-3311311030311033-3231130330213022-3301032100122122-3323200303203223) |
| `virtual_server.udp.udp_client_profile.uid` | [virtual_server.udp.udp_client_profile.uid](data-sources--application_profiles--reference--group-004.md#canonical-3123323202130100-3132220203121313-3210312130122012-1311201213301201-2212033102311122-2000111122023232-1101032320213121-3232013020101233) |
| `virtual_server.udp.udp_server_profile` | [virtual_server.udp.udp_server_profile](data-sources--application_profiles--reference--group-004.md#canonical-1232210232213322-3200033101023210-1022100231121012-3311323303321202-0130032002133020-2301003122033333-3021010230323203-3032301322031133) |
| `virtual_server.udp.udp_server_profile.kind` | [virtual_server.udp.udp_server_profile.kind](data-sources--application_profiles--reference--group-004.md#canonical-0301232220333320-0212023310223321-2233100220312112-2300331223010303-2011223202200333-1133121101223323-2123120033132120-0231132012112322) |
| `virtual_server.udp.udp_server_profile.name` | [virtual_server.udp.udp_server_profile.name](data-sources--application_profiles--reference--group-004.md#canonical-1301113121300332-2011130120100301-2332131131321130-2110210213232321-3331311222230133-2321020223101121-1033201330333101-3311123323211313) |
| `virtual_server.udp.udp_server_profile.namespace` | [virtual_server.udp.udp_server_profile.namespace](data-sources--application_profiles--reference--group-004.md#canonical-1100321230312010-3023132212020002-3332131310320113-0331330002112322-1132021322313012-1211330321022323-0202220332323020-1300313033013323) |
| `virtual_server.udp.udp_server_profile.tenant` | [virtual_server.udp.udp_server_profile.tenant](data-sources--application_profiles--reference--group-004.md#canonical-0020233331120310-1320122000011301-1300121301323333-1233233113202320-3223012111320202-0231232222032033-2103203000222311-1012123333332121) |
| `virtual_server.udp.udp_server_profile.uid` | [virtual_server.udp.udp_server_profile.uid](data-sources--application_profiles--reference--group-004.md#canonical-3302132310200220-1112211012111031-0123230123211311-2131012311210121-3310321130120201-2201123213011210-1123332101010032-1222202211310232) |
| `virtual_server.virtual_server_state` | [virtual_server.virtual_server_state](data-sources--application_profiles--reference--group-004.md#canonical-2021230121300320-1123230031330120-2201000331133031-3132210011103321-1233130023200123-0213223330301201-0132022232032021-1202010313232100) |
| `virtual_server.virtual_server_state.state_disabled` | [virtual_server.virtual_server_state.state_disabled](data-sources--application_profiles--reference--group-004.md#canonical-2333230033233103-1022032110001213-1211312103321332-2213010220333300-1321201012201122-0222101313111331-2101220020032332-2311032233010201) |
| `virtual_server.virtual_server_state.state_enabled` | [virtual_server.virtual_server_state.state_enabled](data-sources--application_profiles--reference--group-004.md#canonical-2113010333112121-3313300103110012-2033000201111131-3131131322212022-2332033122133012-0222030133022310-0331101012023000-1102133031030230) |
| `virtual_server.vs_score` | [virtual_server.vs_score](data-sources--application_profiles--reference--group-001.md#canonical-1101221003132231-2302003212102002-0031303203022332-0130313122313322-3003301331000021-3023001002002021-3020022133323122-1130301100301033) |

<a id="canonical-3021122313210131-2313033223133231-0330112103210330-1123322311012132-1112331223212022-3311220221212131-2333232332331221-1001103023222323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advanced_tcp_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- advanced_tcp_profile

<a id="canonical-3103131213032020-2232220002323000-3022110002112000-2132202321322211-0103100312231112-0011131333112302-0131132001303202-2001112033100223"></a>

Type: `"single"`. Computed.

Configuration parameter for advanced tcp profile.

Additional upstream details:

BIG-IP Advanced TCP Profile.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-tcp_advanced_profile_choice": "[\"disable_tcp_advanced_profile\",\"enable_tcp_advanced_profile\"]"
}
```

<a id="canonical-2320021133323221-3220320103031102-3230223112201333-2110000202012013-1022202023302330-1312323030211230-3011321321331303-2132200120130030"></a>

### Direct properties for `advanced_tcp_profile`

- [disable_tcp_advanced_profile](data-sources--application_profiles--reference--group-001.md#canonical-1211312203022020-2210333120001231-2112233012131222-0322020312302212-3111201120012100-1103233322330232-2213012011113132-1113110021011321): complete subsection reference.

- [enable_tcp_advanced_profile](data-sources--application_profiles--reference--group-001.md#canonical-1201031020131211-1201301212331330-0323300030103000-2132033101012301-0212000023330022-2330220031030100-2102032303300011-1103132223321011): complete subsection reference.

<a id="canonical-1211312203022020-2210333120001231-2112233012131222-0322020312302212-3111201120012100-1103233322330232-2213012011113132-1113110021011321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advanced_tcp_profile.disable_tcp_advanced_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [advanced_tcp_profile](data-sources--application_profiles--reference--group-001.md#canonical-3021122313210131-2313033223133231-0330112103210330-1123322311012132-1112331223212022-3311220221212131-2333232332331221-1001103023222323)
- advanced_tcp_profile.disable_tcp_advanced_profile

<a id="canonical-0031021223020121-0313201222331322-2213213200312303-1200002202321001-2020122300120332-3232201110311310-2222112231313233-1212332103200323"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable tcp advanced profile.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1201031020131211-1201301212331330-0323300030103000-2132033101012301-0212000023330022-2330220031030100-2102032303300011-1103132223321011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advanced_tcp_profile.enable_tcp_advanced_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [advanced_tcp_profile](data-sources--application_profiles--reference--group-001.md#canonical-3021122313210131-2313033223133231-0330112103210330-1123322311012132-1112331223212022-3311220221212131-2333232332331221-1001103023222323)
- advanced_tcp_profile.enable_tcp_advanced_profile

<a id="canonical-2132123223321020-3220331133132202-2132112203103122-3030222322022111-0112222320032010-0213023031331001-1032132003323303-2101312322033331"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable tcp advanced profile.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0221310331001010-1231120233222200-3103210111100320-0121030023132112-2012102312231220-0311210333123322-0133000211232001-0223320121031332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- ddos_profile

<a id="canonical-2320300300230322-1211321231211200-2130113001113201-0022333322320010-0321020212300310-1202032132303002-1133321112002032-3320101300333020"></a>

Type: `"single"`. Computed.

Configuration parameter for ddos profile.

Additional upstream details:

BIG-IP DDoS Protection Rules.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ddos_mitigation_choice": "[\"disable_ddos_mitigation\",\"enable_ddos_mitigation\"]"
}
```

<a id="canonical-3021332012331213-3011220113013120-0121303122232302-0321013320002302-1012213120133113-1113232123002200-3202322120200310-2202303000113310"></a>

### Direct properties for `ddos_profile`

- [disable_ddos_mitigation](data-sources--application_profiles--reference--group-001.md#canonical-1301103321130332-0102030120011123-3330123021113132-1122232120300233-3312032313302030-1021322231120222-2330133133011321-2013122333211030): complete subsection reference.

- [enable_ddos_mitigation](data-sources--application_profiles--reference--group-001.md#canonical-2130333211101101-3312313133130030-3010302112310122-3220212303100023-3000212102001201-1123123122113223-2312033101232100-0303313303021010): complete subsection reference.

<a id="canonical-1301103321130332-0102030120011123-3330123021113132-1122232120300233-3312032313302030-1021322231120222-2330133133011321-2013122333211030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_profile.disable_ddos_mitigation` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [ddos_profile](data-sources--application_profiles--reference--group-001.md#canonical-0221310331001010-1231120233222200-3103210111100320-0121030023132112-2012102312231220-0311210333123322-0133000211232001-0223320121031332)
- ddos_profile.disable_ddos_mitigation

<a id="canonical-2322002221332203-2201123333331111-2111331012010001-0003122232130211-3213000131220033-2020210012320323-3330123030122303-3311333022131021"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2130333211101101-3312313133130030-3010302112310122-3220212303100023-3000212102001201-1123123122113223-2312033101232100-0303313303021010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_profile.enable_ddos_mitigation` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [ddos_profile](data-sources--application_profiles--reference--group-001.md#canonical-0221310331001010-1231120233222200-3103210111100320-0121030023132112-2012102312231220-0311210333123322-0133000211232001-0223320121031332)
- ddos_profile.enable_ddos_mitigation

<a id="canonical-1201301110212113-3212113310322200-2101012321233320-3310312202101201-0012310112100332-0312121210221332-0100311332322101-0201211232113212"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2300032011023220-0003113120202231-2113023300123123-0013230121220321-0310120231310032-0012232113223333-2122301113222011-3212010010132123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `irules` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- irules

<a id="canonical-2133110011210120-3221112323330121-0301000132001231-3020023210120022-3213032011312201-3103230032102013-0212012120003223-3330103002332200"></a>

Type: `"list"`. Computed.

OPTIONS for attaching iRules to BIG-IP Proxy.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1231011023311003-1032310310133330-3202101122130303-0033133203202032-2121120031120112-1011231021221131-0112332013200102-0313020220002023"></a>

### Direct properties for `irules`

<a id="canonical-2231003121311202-0213213112332303-0030001011123023-0202323320301000-2222332313302223-0110031013300313-1003021123332320-3110011302033013"></a>

#### `irules.kind` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3032133331010100-0130211001111000-3013131111201000-0203101213010032-0101102112312322-3000330001202100-0310312321112210-1231230202312022"></a>

<a id="canonical-3001321323331210-1202113101213111-0031303222103202-1320232123122202-0201330133021310-3120230302320110-0032202212000031-0131313102320222"></a>

#### `irules.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3011031323202031-2003130132103122-1022010132132101-1332333031231233-2113023203202320-2330031200031011-0302100133331233-3332022003111310"></a>

<a id="canonical-0203020202120232-1212200310113230-0333333131310210-0111320003120033-0211101321102302-3320002100202002-0201102103031133-3311220001321131"></a>

#### `irules.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0113113302233222-0203333101022230-3110022201223300-3111200220033013-3232002213031321-3112221330302101-2221032120231323-3122010002331010"></a>

<a id="canonical-3012110002303131-1102103330223002-3120010322231020-3112232212031032-1213302303312303-1000030103203300-3023223313333213-1203210133231123"></a>

#### `irules.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2203323300131223-1213221320020123-1022210120132113-3130023331202232-0231320232213032-0323103100113100-0123100233211303-0002003310201103"></a>

<a id="canonical-1130230021030000-1300203210213321-2202011021210201-3223122023333201-0222122212123022-3311012313101102-2223202333233130-0032030233221030"></a>

#### `irules.uid` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- virtual_server

<a id="canonical-1231202302200330-2103201112010223-1211212111221222-0330313120012320-2323003313200213-2022013313322102-1230030320022013-0231202313021200"></a>

Type: `"single"`. Computed.

Specifies configuration related to virtual server.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-virtual_server_type": "[\"http\",\"http3\",\"https\",\"tcp\",\"udp\"]"
}
```

<a id="canonical-2233203101303021-1323103100331112-2301120020222203-2301220131222301-3102131111200112-0103220022112003-0331030213321130-2332130212130312"></a>

### Direct properties for `virtual_server`

- [access_profile](data-sources--application_profiles--reference--group-001.md#canonical-0330111230021223-2331322222210021-0000010131012030-0313133000321011-2023010031313221-0130003120311221-1211320111201111-0220033231131332): complete subsection reference.

- [address_translation](data-sources--application_profiles--reference--group-001.md#canonical-1331202223132031-0030211010012330-0023312023321130-1130322230022331-1311131323321212-0122132311013200-3032201032200000-2221121310313101): complete subsection reference.

- [auto_last_hop](data-sources--application_profiles--reference--group-001.md#canonical-1213231221312103-2101111203001231-0201213300231211-0313012112021033-1012200101220121-2101220302131103-3212333133311002-3200300331321033): complete subsection reference.

- [clone_pool_client](data-sources--application_profiles--reference--group-001.md#canonical-1131130303011232-1131010213131200-0320213310313203-0333203302302103-2233202003023320-3210231020011110-2312132120031031-1033111000202213): complete subsection reference.

- [clone_pool_server](data-sources--application_profiles--reference--group-001.md#canonical-2032302013033112-3013031031321102-0211223022023110-2100203132130112-2020020003012330-2030002012033001-3131021200201022-3303102003112033): complete subsection reference.

<a id="canonical-3003311010303322-2010201320203020-2022302001310330-3022301301000020-1332202330032100-3120230332203201-3101010322033203-1000331103200032"></a>

<a id="canonical-1113303311310022-1212123231202002-2310312113001232-3000333132312132-3033000303302030-2000212031120321-0121222201103000-3202100202101322"></a>

#### `virtual_server.connection_limit` property

Type: `"number"`. Computed.

Specifies the maximum number of concurrent connections allowed for the virtual server. Setting this
to 0 turns off connection limits. The default is 0.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-3012230232332103-2322330122210131-1311320333112003-1131223222023210-0013122003222112-0201211222130320-1003001021101031-3103322232233031"></a>

<a id="canonical-1223020121130300-2302001132133323-1001113001300103-1112302010330111-1212110130201202-1123323231310130-2210323103002112-3011321200231001"></a>

#### `virtual_server.connection_rate_limit` property

Type: `"number"`. Computed.

Specifies the maximum number of connections-per-second allowed for a virtual server. When the number
of connections-per-second reaches the limit for a given virtual server, the system drops (UDP) or
resets (TCP) additional connection requests. This helps detect Denial of Service attacks, where
connection requests flood a virtual server. Setting this to 0 turns off connection limits. The
default is 0.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [connection_rate_limit_mode](data-sources--application_profiles--reference--group-001.md#canonical-0033330122013310-3003120310012330-3211000223223221-0223111311312230-0022202003132033-0321011121003210-1210122112102020-2213303203301133): complete subsection reference.

- [default_persistence_profile](data-sources--application_profiles--reference--group-002.md#canonical-0003121231013131-0212113030120031-0011021022033310-1103303313001332-1210322112031030-2212031333233022-0102311320322312-2320101203011201): complete subsection reference.

- [default_pool](data-sources--application_profiles--reference--group-002.md#canonical-0333332101230033-1320131322211033-1021010222233010-0332010100110010-2322112110200221-2223211110003312-2100223212033220-1202302003303231): complete subsection reference.

- [fallback_persistence_profile](data-sources--application_profiles--reference--group-002.md#canonical-3313233021210032-2130033220330002-1330201211103120-1320001012332020-3303101002000220-0110120200003300-2322012100222132-0123332102020230): complete subsection reference.

- [fix_profile](data-sources--application_profiles--reference--group-002.md#canonical-3332013220003223-0301321210023212-0103101231222131-1313130322011033-3331230212003132-3203221202003221-3220112221012033-3121020033101103): complete subsection reference.

- [http](data-sources--application_profiles--reference--group-002.md#canonical-0102203023322313-2333211312212213-0011231012003031-3320001333231133-2303121120302112-1333102100320323-1322213333033302-0021001000133031): complete subsection reference.

- [http3](data-sources--application_profiles--reference--group-002.md#canonical-2233213302301031-1011021100023220-0023001210122232-0332111101220031-0312121002130112-0330200230012032-0313003311311113-0322012000031213): complete subsection reference.

- [https](data-sources--application_profiles--reference--group-003.md#canonical-0321010331132100-1111331101132032-2102011312012203-1301132220133301-2212023121221300-1311331320200031-2332233222113212-0111022211020333): complete subsection reference.

- [immediate_action_on_service_down](data-sources--application_profiles--reference--group-003.md#canonical-2300123112133033-1303100333113302-3231112332130213-3210001021212130-0303103023020132-3211232333100212-1030102323131320-3023101112213320): complete subsection reference.

- [last_hop_pool](data-sources--application_profiles--reference--group-003.md#canonical-0000301021323203-1003301002002311-0310011213320313-0102003200331211-3103021001233200-3212113122100001-3320222320102111-3313223020133023): complete subsection reference.

- [nat64](data-sources--application_profiles--reference--group-003.md#canonical-3320001111230010-0321322131120330-1030210203213223-1310133111230310-1101021300130220-2333001002233033-0302123321201121-2333013113030122): complete subsection reference.

- [port_translation](data-sources--application_profiles--reference--group-004.md#canonical-1033032133123230-2203033131231333-2313120131311332-1121301201221211-1321320320100030-3030133220301332-1312323030311221-3002302202301220): complete subsection reference.

- [request_logging_profile](data-sources--application_profiles--reference--group-004.md#canonical-3232320301212213-2201233011223222-1121012211011023-1212323111231033-1300332330213333-1302020033323222-1223102122223230-1103311321323132): complete subsection reference.

- [source_port](data-sources--application_profiles--reference--group-004.md#canonical-1000133010120020-1221212220013310-3203232303131320-2201230212112001-3230003001002030-0010312331212203-2010213211021213-0221332122203203): complete subsection reference.

- [statistics_profile](data-sources--application_profiles--reference--group-004.md#canonical-0230033123202003-3221010301120301-1321302111033133-3300013133310233-0333313030302000-0200200313230110-2223320101301232-1322122301001122): complete subsection reference.

- [tcp](data-sources--application_profiles--reference--group-004.md#canonical-1332122030012230-3112023330232023-3120100021202202-2130120302110303-1003010002323020-0300001302111033-2112323013011103-3321010030001101): complete subsection reference.

- [udp](data-sources--application_profiles--reference--group-004.md#canonical-3111012112013331-0210230112200321-2133130123230231-2102200122231133-2202101332332032-3132132213230031-3232202101331130-1313232102312102): complete subsection reference.

- [virtual_server_state](data-sources--application_profiles--reference--group-004.md#canonical-0300001203202301-3231100101123312-2010220203023303-2030320120111312-3201010231210333-1312201331313222-1321111320032321-3200132220130130): complete subsection reference.

<a id="canonical-1101221003132231-2302003212102002-0031303203022332-0130313122313322-3003301331000021-3023001002002021-3020022133323122-1130301100301033"></a>

<a id="canonical-3120221113303202-2321023222213202-1123000120103100-1121002000231233-3213303321000313-2010112023323120-0330230322113013-0012003011300001"></a>

#### `virtual_server.vs_score` property

Type: `"number"`. Computed.

Specifies the virtual server score in percent. Global Traffic Manager (GTM) can rely on this value
to load balance traffic in a proportional manner. The , meaning that no additional metric is applied
for the virtual server.

Additional upstream details:

The default is 0, meaning that no additional metric is applied for the virtual server.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-0330111230021223-2331322222210021-0000010131012030-0313133000321011-2023010031313221-0130003120311221-1211320111201111-0220033231131332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.access_profile` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- virtual_server.access_profile

<a id="canonical-3301203010312110-0311022301323111-1330231332111100-2011320331000132-2222130021103133-1331303222310033-0030323201201010-1103232313122231"></a>

Type: `"list"`. Computed.

Specifies an access policy that determines the authentication rules and access controls applied to
user sessions for this virtual server.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3122310022033003-1031010100112321-1001321131312230-2122300200130020-1310022301322201-2020212101130233-3012003333300333-3332032330121313"></a>

### Direct properties for `virtual_server.access_profile`

<a id="canonical-2231313103100123-1012223020200132-2030112321133131-0220311323301323-0220203233301110-2311102001232002-0233002222132323-3322222210121110"></a>

#### `virtual_server.access_profile.kind` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1001030002331213-1200220200021021-3201311020122131-2212313223210133-3122313001223220-1120130322012302-2333031103020103-2011310110121310"></a>

<a id="canonical-2113233302331033-2132230033103301-1100103003213203-2113031233332303-1002200130020311-0123212322011101-2303002331123203-1021203311303332"></a>

#### `virtual_server.access_profile.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1202313233310023-2303312323211101-2001001313003333-2120023100022020-1122211022133301-1210121203322231-2313120300100021-0322132303033111"></a>

<a id="canonical-2211330201201310-3331321311233332-1131303021113111-1112332113012303-2102033313113003-0002330220022302-1310001221131332-2111211112031131"></a>

#### `virtual_server.access_profile.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3213211300012100-2020000321031102-3033213323022002-2210033232122303-2333200332120001-2311200303030333-0303332210131321-2021223102113003"></a>

<a id="canonical-0130030011020000-2013202022223032-1130110322032330-3232320030232320-2131003221113113-3333233330000322-1330133332301210-0110320211300030"></a>

#### `virtual_server.access_profile.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2233312101133310-1111313233121311-0123002303310202-3332331102213102-1331030031323222-0302203002201213-2310321013022120-0310323012123132"></a>

<a id="canonical-0222023300001212-3331322211201223-1320330012312212-2021333110223122-1300101031003301-1301032330322101-2132021120022110-2002022121213001"></a>

#### `virtual_server.access_profile.uid` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1331202223132031-0030211010012330-0023312023321130-1130322230022331-1311131323321212-0122132311013200-3032201032200000-2221121310313101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.address_translation` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- virtual_server.address_translation

<a id="canonical-3111302313230133-2200011021210113-0023331013233300-3132021121122301-2010330222212230-1300230031112223-0210130111031012-1012132131033200"></a>

Type: `"single"`. Computed.

Specifies, when checked (enabled), that the system translates the address of the virtual server.
When cleared (disabled), specifies that the system uses the address without translation. This option
is useful when the system is load balancing devices that have the same IP address. The default is
enabled.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-address_translation_choice": "[\"address_translation_disable\",\"address_translation_enable\"]"
}
```

<a id="canonical-1020222013221300-3001211331100332-2211111020320111-1331311120122221-0223012321023110-2311202100201012-3111311003123301-1303120212000130"></a>

### Direct properties for `virtual_server.address_translation`

- [address_translation_disable](data-sources--application_profiles--reference--group-001.md#canonical-2111011230102001-3320020103022322-0133110333221222-1230203000020200-3132321322311123-0220122330020331-2001023001122333-3030031230012122): complete subsection reference.

- [address_translation_enable](data-sources--application_profiles--reference--group-001.md#canonical-1111333211113101-0010211023033132-2012210122112210-0233012120223020-1013122223223102-3002213121301212-1321012330310103-1110033330320331): complete subsection reference.

<a id="canonical-2111011230102001-3320020103022322-0133110333221222-1230203000020200-3132321322311123-0220122330020331-2001023001122333-3030031230012122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.address_translation.address_translation_disable` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.address_translation](data-sources--application_profiles--reference--group-001.md#canonical-1331202223132031-0030211010012330-0023312023321130-1130322230022331-1311131323321212-0122132311013200-3032201032200000-2221121310313101)
- virtual_server.address_translation.address_translation_disable

<a id="canonical-0333100012003221-3200321123211200-3303231003000021-0012100202222013-1111303330313130-3331303221230020-2113320030021002-3232202123121300"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1111333211113101-0010211023033132-2012210122112210-0233012120223020-1013122223223102-3002213121301212-1321012330310103-1110033330320331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.address_translation.address_translation_enable` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.address_translation](data-sources--application_profiles--reference--group-001.md#canonical-1331202223132031-0030211010012330-0023312023321130-1130322230022331-1311131323321212-0122132311013200-3032201032200000-2221121310313101)
- virtual_server.address_translation.address_translation_enable

<a id="canonical-2310003010031330-1103021321220101-2110200112013000-1010330212232203-1110223010001323-1323203330230203-2321210132033320-3330122102130101"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1213231221312103-2101111203001231-0201213300231211-0313012112021033-1012200101220121-2101220302131103-3212333133311002-3200300331321033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.auto_last_hop` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- virtual_server.auto_last_hop

<a id="canonical-2103132212111120-2202120331300033-1000221003232103-0010001000132330-3301310123322201-1332000011200121-2121211201322111-1000123012312330"></a>

Type: `"single"`. Computed.

When enabled, allows the system to send return traffic to the MAC address that transmitted the
request, even if the routing table points to a different network or interface. As a result, the
system can send return traffic to clients even when there is no matching route. For example, if the
system does not have a default route configured and the client is located on a remote network. This
setting is also useful when the system is load balancing transparent devices that do not modify the
source IP address of the packet. Without the last hop option enabled, the system could return
connections to a different transparent node, resulting in asymmetric routing. You can configure this
setting globally and on an object level. You set the global Auto Last Hop value on the System ::
Configuration :: Local Traffic :: General screen. To configure this setting globally, retain the
Default setting. When you configure Auto Last Hop with a value other than Default at the object
level, its setting takes precedence over the global setting. This enables you to configure auto last
hop on a per-virtual server basis. The default is Default, meaning that the system uses the global
auto-lasthop setting to send back the request.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-auto_last_hop_choice": "[\"auto_last_hop_default\",\"auto_last_hop_disable\",\"auto_last_hop_enable\"]"
}
```

<a id="canonical-1202210222201202-2130211301023323-3003201311110303-1102110213133000-2313213132130230-1112231012300000-0201232300120023-2331033311313223"></a>

### Direct properties for `virtual_server.auto_last_hop`

- [auto_last_hop_default](data-sources--application_profiles--reference--group-001.md#canonical-2022202003321202-0000232233101313-1132023331313010-2010231221333133-3230131322101322-3033101310123031-3233122211121121-2210132301101013): complete subsection reference.

- [auto_last_hop_disable](data-sources--application_profiles--reference--group-001.md#canonical-3023333310000330-2311322123300120-1330121020233003-1023110322110203-3230211030203320-3333121111331332-2330233202203201-1013122232123220): complete subsection reference.

- [auto_last_hop_enable](data-sources--application_profiles--reference--group-001.md#canonical-2010131201000311-3023001133332320-2011000031201313-0300000013003111-0133020120331113-2203133021110213-2010213130332333-1302102022032200): complete subsection reference.

<a id="canonical-2022202003321202-0000232233101313-1132023331313010-2010231221333133-3230131322101322-3033101310123031-3233122211121121-2210132301101013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.auto_last_hop.auto_last_hop_default` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.auto_last_hop](data-sources--application_profiles--reference--group-001.md#canonical-1213231221312103-2101111203001231-0201213300231211-0313012112021033-1012200101220121-2101220302131103-3212333133311002-3200300331321033)
- virtual_server.auto_last_hop.auto_last_hop_default

<a id="canonical-0033033102121111-1003132132313300-3313002022121131-3131133122300133-3012033310320320-2313312211100312-1112302200233220-0100001011011301"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for auto last hop default.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3023333310000330-2311322123300120-1330121020233003-1023110322110203-3230211030203320-3333121111331332-2330233202203201-1013122232123220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.auto_last_hop.auto_last_hop_disable` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.auto_last_hop](data-sources--application_profiles--reference--group-001.md#canonical-1213231221312103-2101111203001231-0201213300231211-0313012112021033-1012200101220121-2101220302131103-3212333133311002-3200300331321033)
- virtual_server.auto_last_hop.auto_last_hop_disable

<a id="canonical-3310122322320331-1011030232231333-3303321213223210-2323130100230010-3212012213201132-0113312033230321-3322202001332131-2311100302011303"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for auto last hop disable.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2010131201000311-3023001133332320-2011000031201313-0300000013003111-0133020120331113-2203133021110213-2010213130332333-1302102022032200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.auto_last_hop.auto_last_hop_enable` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.auto_last_hop](data-sources--application_profiles--reference--group-001.md#canonical-1213231221312103-2101111203001231-0201213300231211-0313012112021033-1012200101220121-2101220302131103-3212333133311002-3200300331321033)
- virtual_server.auto_last_hop.auto_last_hop_enable

<a id="canonical-2002020012330323-0233012113113230-1000323110103211-1133012112003332-2112110100330132-3021310202232111-1110121030233113-0000021223223323"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for auto last hop enable.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1131130303011232-1131010213131200-0320213310313203-0333203302302103-2233202003023320-3210231020011110-2312132120031031-1033111000202213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.clone_pool_client` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- virtual_server.clone_pool_client

<a id="canonical-1133020132233211-2310110101220011-0012011332002321-1013213132313033-2101320220311103-3111012322101121-0030230000333132-3232002121221313"></a>

Type: `"list"`. Computed.

Replicates client-side traffic (that is, prior to address translation) to a member of the specified
pool.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1332031133212321-3333023023311323-0011320211122122-0333023322312202-0032232120323101-3310103211330101-0110102123113123-3232132320102023"></a>

### Direct properties for `virtual_server.clone_pool_client`

<a id="canonical-3213032121001202-0030120300231013-1122132212012331-0213111222032330-0231101222112200-0021300203313331-2021101022032200-3003233111333003"></a>

#### `virtual_server.clone_pool_client.kind` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1321303301323112-3232331030023023-3022333311103010-3213223013313322-0301310112122133-0213021320332213-1321011203112101-3020013230120201"></a>

<a id="canonical-1212302130220311-0022113121011020-2023230211212230-2001013030203001-2002100202302013-2031333101232132-3002322022032301-0231033312003200"></a>

#### `virtual_server.clone_pool_client.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2211231222223111-2013313312230333-2230101323231222-3310030032031122-1311333213212023-3001220311021101-2201003010330230-3331023313321232"></a>

<a id="canonical-2211123012122013-0111033320130110-2032121003202211-3100013030111120-2000202133031210-1310021013121131-3022331322333301-0332332110120013"></a>

#### `virtual_server.clone_pool_client.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2010332203320110-3211023123321223-2220230012123321-3220030210210222-1120021002013323-3013300001130013-1222022213222021-1102013303321332"></a>

<a id="canonical-3111020002022200-0232332320033032-1030220210323331-0321202120220131-3101300102201002-2230003303101211-1021331200211330-0212200220221213"></a>

#### `virtual_server.clone_pool_client.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1233111210320002-1132132001131321-3331021232002313-3213030003001313-0100233223131310-0101033102133302-1212313112300010-3321322311312100"></a>

<a id="canonical-2302031123122231-3110102321233202-1222231201103103-0111221010112220-0330211011320030-3202033211201100-1312121121330013-3103112102110000"></a>

#### `virtual_server.clone_pool_client.uid` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2032302013033112-3013031031321102-0211223022023110-2100203132130112-2020020003012330-2030002012033001-3131021200201022-3303102003112033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.clone_pool_server` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- virtual_server.clone_pool_server

<a id="canonical-0221313111003203-0222221132312111-2113212313313121-3203223020221012-0120231221323112-2032123110022222-2323131330111332-3112223023201220"></a>

Type: `"list"`. Computed.

Replicates server-side traffic (that is, prior to address translation) to a member of the specified
pool.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1330321302331311-3232320031020203-3333003332332013-0001131002313111-2301212211102001-3333232120131301-0123303102200010-0022030133032023"></a>

### Direct properties for `virtual_server.clone_pool_server`

<a id="canonical-1233112111102023-0311022121101203-1103103012100323-3002011000102131-2303210030131133-0232113013121110-0210132133322212-3222003200112331"></a>

#### `virtual_server.clone_pool_server.kind` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2303031121211220-1001230213300030-1230003332311032-2020100332231213-2002002231020011-1210302022002120-3131131320031001-0031001203220231"></a>

<a id="canonical-2220130130333121-0003032203321012-1233200010321003-0320010001312230-2003100001321101-3032303131212000-2322220310323020-0300111122330200"></a>

#### `virtual_server.clone_pool_server.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2101320312013231-0223313033130100-2303200330132312-1032011011320301-3201333231101301-0303122123311321-0010020112230332-3320033113030101"></a>

<a id="canonical-0200212003222322-0231020021130313-1131231300200232-3003221110230223-2113000021233322-2203332310110011-0200323331320030-0303102303313003"></a>

#### `virtual_server.clone_pool_server.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0032333022002101-1013203220103012-2001022001111112-1223202313132122-0230230300111130-2222232131103130-3002310012201213-0113010322130101"></a>

<a id="canonical-2322112300013010-0311333102213011-3102000113000233-1033211000203333-0332301100221100-1302030221110003-0300033332222312-3001113310020020"></a>

#### `virtual_server.clone_pool_server.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1213233212200103-0311132031101102-1310100010321021-2230013220132230-1202213133313003-3222020310320312-3333231200201320-3203202321303212"></a>

<a id="canonical-3301003013312223-0130220203320032-3123030010121003-2120033012223132-3323333332032111-1333101022101023-2021110112011133-2201300002302312"></a>

#### `virtual_server.clone_pool_server.uid` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0033330122013310-3003120310012330-3211000223223221-0223111311312230-0022202003132033-0321011121003210-1210122112102020-2213303203301133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.connection_rate_limit_mode` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- virtual_server.connection_rate_limit_mode

<a id="canonical-3131301032310212-2101213001120020-1230322302233022-3233223330230310-3120121330213030-1101103200302113-0112333012303230-0113213203012130"></a>

Type: `"single"`. Computed.

Configuration parameter for connection rate limit mode.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-connection_rate_limit_mode_choice": "[\"per_destination_address\",\"per_source_address\",\"per_source_destination_address\",\"per_virtual_server\",\"per_virtual_server_destination_address\",\"per_virtual_server_source_address\",\"per_virtual_server_source_destination_address\"]"
}
```

<a id="canonical-1111212110213200-1002333111311302-3222131220112132-2313310322313202-3223103301132023-3031322103103203-1113320232123231-0001100100302132"></a>

### Direct properties for `virtual_server.connection_rate_limit_mode`

- [per_destination_address](data-sources--application_profiles--reference--group-001.md#canonical-1002013231330223-0210311011222020-0200121110332311-3322011022001233-1013111003000130-0032320001012112-2201213212310310-3101222331022232): complete subsection reference.

- [per_source_address](data-sources--application_profiles--reference--group-001.md#canonical-1113223020303102-3013133000002220-2223331001100031-0021102011132320-3121321313300031-0222001230303320-0323021203100011-3222002200123223): complete subsection reference.

- [per_source_destination_address](data-sources--application_profiles--reference--group-001.md#canonical-3330032322333222-0102301301210031-1303211030222203-1011111011221201-0201131101002113-0322122120130331-3322302200001000-1031202013332002): complete subsection reference.

- [per_virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-1211300320020032-2332220233313232-1323310232111231-1112301311010201-0031313130123310-2313110331332002-3100233322000311-2120000323231021): complete subsection reference.

- [per_virtual_server_destination_address](data-sources--application_profiles--reference--group-002.md#canonical-0022023203222223-2233310313211013-3231110112311020-0302220001032013-2123020333021102-2010333100122021-3300132022010222-2001210102010202): complete subsection reference.

- [per_virtual_server_source_address](data-sources--application_profiles--reference--group-002.md#canonical-1203000232222321-0022300313030100-1233122111223210-1013001222300302-3030121231032212-1011223211131230-0322211311210331-2212223232333230): complete subsection reference.

- [per_virtual_server_source_destination_address](data-sources--application_profiles--reference--group-002.md#canonical-1022121021032130-3303223311200321-3023112202000010-2312212012233102-1313112211301022-3100031320321011-1110221300212230-0022333003220131): complete subsection reference.

<a id="canonical-1002013231330223-0210311011222020-0200121110332311-3322011022001233-1013111003000130-0032320001012112-2201213212310310-3101222331022232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.connection_rate_limit_mode.per_destination_address` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.connection_rate_limit_mode](data-sources--application_profiles--reference--group-001.md#canonical-0033330122013310-3003120310012330-3211000223223221-0223111311312230-0022202003132033-0321011121003210-1210122112102020-2213303203301133)
- virtual_server.connection_rate_limit_mode.per_destination_address

<a id="canonical-3232212303310231-3321200231301300-3113023011023311-3313110003322033-3033323201110311-1020200001001221-2000001212031120-1222302330103002"></a>

Type: `"single"`. Computed.

Destination Address Mask.

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

<a id="canonical-0333211210221211-2330033011021121-0212203132212321-1100023302112330-1101102021003223-1032032030232011-1332131203320230-0100023022020213"></a>

### Direct properties for `virtual_server.connection_rate_limit_mode.per_destination_address`

<a id="canonical-1001012133322200-0221033223121133-2331121003312111-2331010112133332-2213231123033330-2303322333300222-3202213011302123-2332321302033223"></a>

#### `virtual_server.connection_rate_limit_mode.per_destination_address.destination_mask` property

Type: `"number"`. Computed.

Configuration parameter for destination mask.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-1113223020303102-3013133000002220-2223331001100031-0021102011132320-3121321313300031-0222001230303320-0323021203100011-3222002200123223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.connection_rate_limit_mode.per_source_address` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.connection_rate_limit_mode](data-sources--application_profiles--reference--group-001.md#canonical-0033330122013310-3003120310012330-3211000223223221-0223111311312230-0022202003132033-0321011121003210-1210122112102020-2213303203301133)
- virtual_server.connection_rate_limit_mode.per_source_address

<a id="canonical-0222220332310221-3031030011221323-1011011023211200-3200123213223121-3232010100021321-1301303022330301-2133102203321113-1013302001231120"></a>

Type: `"single"`. Computed.

Source Address Mask.

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

<a id="canonical-0112233101230000-3112323123111023-3321233011331030-1320222313233130-1300111000332103-2322210023012213-0200003320011211-0131201120013000"></a>

### Direct properties for `virtual_server.connection_rate_limit_mode.per_source_address`

<a id="canonical-3231311131311102-0231202211132103-2303012002111113-2330022230301011-0030001031300231-1310220303210001-3022302002212130-3121310330023200"></a>

#### `virtual_server.connection_rate_limit_mode.per_source_address.source_mask` property

Type: `"number"`. Computed.

Configuration parameter for source mask.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-3330032322333222-0102301301210031-1303211030222203-1011111011221201-0201131101002113-0322122120130331-3322302200001000-1031202013332002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.connection_rate_limit_mode.per_source_destination_address` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.connection_rate_limit_mode](data-sources--application_profiles--reference--group-001.md#canonical-0033330122013310-3003120310012330-3211000223223221-0223111311312230-0022202003132033-0321011121003210-1210122112102020-2213303203301133)
- virtual_server.connection_rate_limit_mode.per_source_destination_address

<a id="canonical-1031201331123121-2000310300021002-3311103012020003-0212333312222020-0230130220020202-2233323312133112-0233212223003233-1102110103323223"></a>

Type: `"single"`. Computed.

Destination and Source Address Mask.

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

<a id="canonical-0002021111023013-0111201131030112-0322333210212302-0302101121102231-1331002303231002-0000001303200122-1203323021200131-1333233230033002"></a>

### Direct properties for `virtual_server.connection_rate_limit_mode.per_source_destination_address`

<a id="canonical-2012310123231112-1103313101030133-1301112222002121-2033330023011110-1022331020330012-3231001300131133-3032133330112032-1212002012211230"></a>

#### `virtual_server.connection_rate_limit_mode.per_source_destination_address.destination_mask` property

Type: `"number"`. Computed.

Configuration parameter for destination mask.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-2032010011100310-3330311012133302-2122031000123320-0321121311231120-3121120001110123-0103102332131030-1120003131321032-3300220100223020"></a>

<a id="canonical-1203310210132321-1302210102303202-0333312323001302-3333211223121031-2313301300313200-2100122121013131-3302030132332203-3031031030032222"></a>

#### `virtual_server.connection_rate_limit_mode.per_source_destination_address.source_mask` property

Type: `"number"`. Computed.

Configuration parameter for source mask.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-1211300320020032-2332220233313232-1323310232111231-1112301311010201-0031313130123310-2313110331332002-3100233322000311-2120000323231021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `virtual_server.connection_rate_limit_mode.per_virtual_server` properties

Breadcrumbs:

- [xcsh_application_profiles](../data-sources/application_profiles.md#canonical-1311211203002323-2003223022233300-1230021211313022-0103300013312020-2232032102012103-1003213312301301-2022210332331102-2300323031031330)
- [Property reference](data-sources--application_profiles--reference--group-001.md#canonical-3100202011322300-0311322111330123-0130031222313021-0003030210210221-2333312022200112-1223301232231311-3102021102011321-0333231203003231)
- [virtual_server](data-sources--application_profiles--reference--group-001.md#canonical-2011012130233013-1021222233222201-1010303311001211-1001312100300300-2222221213013310-2022113133020232-2330233312111312-1320022121303130)
- [virtual_server.connection_rate_limit_mode](data-sources--application_profiles--reference--group-001.md#canonical-0033330122013310-3003120310012330-3211000223223221-0223111311312230-0022202003132033-0321011121003210-1210122112102020-2213303203301133)
- virtual_server.connection_rate_limit_mode.per_virtual_server

<a id="canonical-1331330110210211-0333212231000230-0110321023330032-0102230332332312-3201232310221002-2210232211210120-2112233221213103-0120230211021302"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for per virtual server.

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

This is an empty object or choice marker. It has no direct properties.
