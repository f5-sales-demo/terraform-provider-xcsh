---
page_title: "xcsh_network_connector reference"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_network_connector reference."
---

# xcsh_network_connector reference

<a id="canonical-3000122203132213-2030301322323023-1302132111133302-0302223201130030-0303221223333310-2021321132132200-3100201000113230-1311223311010320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-3100100033113131-1013100333112132-1220212103012023-0211023132210233-2323121020123322-2323213203113121-1203301130023023-3202032133312301)
- Property reference

<a id="canonical-3121002211330233-1211330321312223-2132311303321102-0200032220102202-1021203103230111-0110210320212010-1311212031011303-3110320023310012"></a>

### Direct properties for `xcsh_network_connector`

<a id="canonical-2013311032020130-1110101003321112-3221230230121313-2110101213200221-0110233232121222-1311100002201020-3103022321220300-1203221201002231"></a>

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

<a id="canonical-2201233211303322-2203331203033213-1223211020331100-0131120132000000-0203220120031130-2312121202221323-1020023103210200-1220012232300103"></a>

<a id="canonical-2213303231233102-0131001032110131-3111022322113002-1211332310030133-2201031032031122-0231111022212100-3132010121223323-0010313331112102"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the NetworkConnector.

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

- [disable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-3211303131123112-2030112332333102-0020201002021110-3311100303310310-3330212213323011-0113203133331032-2211011223233020-3122112200133110): complete subsection reference.

- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-2213111031121031-3003322303011211-0020021331302310-0213312231333222-0202212313032220-0002012131031003-0322101122200102-1301210120231201): complete subsection reference.

<a id="canonical-1113003133013031-0323002212113003-1310023332031123-0333212111130313-2111003233233021-2200323103113010-0101231031322111-3112020213320011"></a>

<a id="canonical-1022230230101100-1333302012311032-0001303021000313-0112333103120333-2303221011013030-0111033311132201-0001303220013203-3001110301102010"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0220111310323203-0323002211331010-3021302011031333-3300002012202213-3123232333020120-1010330110003123-3130013212322102-3012233231211111"></a>

<a id="canonical-0333303100331221-2221300111012200-2102122130031001-3230232210201230-3131331303313301-3000312230220231-0223003101010123-2311111113131200"></a>

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

<a id="canonical-1031111202233200-3312311020312033-2130231310322113-3002231301102312-0311322102221313-0131023120023022-3003001220102123-2202021313332101"></a>

<a id="canonical-3220233111220233-2112133012322330-3030231330111232-3313331011132013-2311312331131022-1201111123213000-1102000020001031-3223330122321233"></a>

#### `name` property

Type: `"string"`. Required.

Name of the NetworkConnector.

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

<a id="canonical-2101032230001323-2211310321311200-1103123312231202-1130022100110012-0102323200113230-3211221210332301-1333330221313013-3113021030110231"></a>

<a id="canonical-0000310133221302-1100003323003320-1320021221300232-1330333120220113-2013211100032111-0023120312232223-0330330112012233-3002122330113323"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the NetworkConnector exists.

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

- [sli_to_global_dr](data-sources--network_connector--reference--group-001.md#canonical-3202003000033330-0132321012323121-2201002113322223-1301301113020232-1303210303223002-1110011130333032-0202031330113122-0100012123123232): complete subsection reference.

- [sli_to_slo_snat](data-sources--network_connector--reference--group-001.md#canonical-0121123320323123-3213322110002213-2100333022000003-2120210113312111-1331133121020232-0212331020331110-1031122203113102-0201102222130031): complete subsection reference.

- [slo_to_global_dr](data-sources--network_connector--reference--group-001.md#canonical-3202010033321000-2321111032303003-3110021002030311-2331130332311233-1122123100201221-1020301003032212-3023213002320313-1232123203201023): complete subsection reference.

<a id="canonical-3112103013223300-0102130112123102-0200331221123102-0103312230230111-2230331311312233-2230210133013101-0013102013130033-0113111011232221"></a>

### All schema paths for `xcsh_network_connector`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--network_connector--reference--group-001.md#canonical-2013311032020130-1110101003321112-3221230230121313-2110101213200221-0110233232121222-1311100002201020-3103022321220300-1203221201002231) |
| `description` | [description](data-sources--network_connector--reference--group-001.md#canonical-2201233211303322-2203331203033213-1223211020331100-0131120132000000-0203220120031130-2312121202221323-1020023103210200-1220012232300103) |
| `disable_forward_proxy` | [disable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-2300212230202233-2033333212221330-0133131003033032-2132303032213322-3023121000133232-3223010030100200-3010210113300302-1130030032201312) |
| `enable_forward_proxy` | [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-0113033202113303-3332112213100002-1230112313310120-1230103301321322-3133011001013112-3023023321123133-3012303202130201-3220121110330123) |
| `enable_forward_proxy.connection_timeout` | [enable_forward_proxy.connection_timeout](data-sources--network_connector--reference--group-001.md#canonical-0020203023310230-1010300112323110-0002220211002102-1102230231111323-2101322031003030-1331133213330313-3131313301010033-0313203000033321) |
| `enable_forward_proxy.max_connect_attempts` | [enable_forward_proxy.max_connect_attempts](data-sources--network_connector--reference--group-001.md#canonical-3321213213001130-1213023303223132-1001213111320300-2321200212321012-2301311123012012-1032210210110110-2021020023302133-0321020122221031) |
| `enable_forward_proxy.no_interception` | [enable_forward_proxy.no_interception](data-sources--network_connector--reference--group-001.md#canonical-0013332213310001-2013103333213302-2231223231023001-1333230330300230-1223313020321302-1232120213020310-2230211230131001-0010312013202210) |
| `enable_forward_proxy.tls_intercept` | [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-0113032323000303-1302332013323013-3023132023020330-3322131122120302-3021211301200022-2312003132303032-0111331002301300-3332121103220211) |
| `enable_forward_proxy.tls_intercept.custom_certificate` | [enable_forward_proxy.tls_intercept.custom_certificate](data-sources--network_connector--reference--group-001.md#canonical-2303323322330003-3301200112010113-2011300133003121-0321203222023210-2331301030200033-3232123300012133-2123002113321330-0030023222012310) |
| `enable_forward_proxy.tls_intercept.custom_certificate.certificate_url` | [enable_forward_proxy.tls_intercept.custom_certificate.certificate_url](data-sources--network_connector--reference--group-001.md#canonical-3232021001330012-3120200300103320-0013321031330221-0020132233221321-3002023330221331-1330003133103013-1101320200100020-3200020013301011) |
| `enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms` | [enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms](data-sources--network_connector--reference--group-001.md#canonical-3310113002133021-2322011133131302-1122202312100103-2023323130321200-1203321223213120-2312302033201130-2303132120310120-3032113323123002) |
| `enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms.hash_algorithms` | [enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms.hash_algorithms](data-sources--network_connector--reference--group-001.md#canonical-2201212221111023-0213323100310330-0330120311230020-2000223010123313-2000302320323210-0200011112221332-1201122300312203-0122320001220322) |
| `enable_forward_proxy.tls_intercept.custom_certificate.description_spec` | [enable_forward_proxy.tls_intercept.custom_certificate.description_spec](data-sources--network_connector--reference--group-001.md#canonical-3320322202231222-2213203222133131-0221002012212033-1202121020300120-3323122103303023-1321030112201120-3322310013223310-0213310023022311) |
| `enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling` | [enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling](data-sources--network_connector--reference--group-001.md#canonical-0003130303331010-2023121332002231-1120102020013032-3101331230313321-3101310211321123-0120102313331121-0322021310202112-1332103303303300) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key](data-sources--network_connector--reference--group-001.md#canonical-1120232100211123-3103020231032030-3010313010332001-0301102220103133-2010213300333330-2102111303211331-3012203230221030-3333212120322320) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info](data-sources--network_connector--reference--group-001.md#canonical-0133020121202030-0010301030302000-1200200330000020-0200233302130001-3123110110211030-0212200110110313-0320313103000112-1103231120110232) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.decryption_provider` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.decryption_provider](data-sources--network_connector--reference--group-001.md#canonical-1000101301312223-2220133133033221-2023001210221302-3131300301330110-1101023323300000-3102100032230110-0013103133203123-0021232202203313) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.location` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.location](data-sources--network_connector--reference--group-001.md#canonical-3101301323333101-3011102211011320-1322230012232312-0223233331120232-3131002032002021-2100321103321210-3103211202332032-0331312333302020) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.store_provider` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.store_provider](data-sources--network_connector--reference--group-001.md#canonical-1003332222222212-1010233023213031-0230223310001123-2202200303011101-1001103011202333-3210330121222123-0113122200033323-0120302231300033) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info](data-sources--network_connector--reference--group-001.md#canonical-3001313212310332-0102032111122330-1111131211232332-3203222312022101-1123330312020031-0031133213122202-3222033310233000-2123101230112121) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info.provider_ref` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info.provider_ref](data-sources--network_connector--reference--group-001.md#canonical-2322330133021011-1011011100310000-3233301310332122-2000202200020032-3023321021200022-2202312212301332-1101223000011122-0033303121313322) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info.url` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info.url](data-sources--network_connector--reference--group-001.md#canonical-0311330313230222-1023311203310132-3021131201320103-0230023032302322-0212001122220133-3103302111002203-2312200312323232-3123120203132202) |
| `enable_forward_proxy.tls_intercept.custom_certificate.use_system_defaults` | [enable_forward_proxy.tls_intercept.custom_certificate.use_system_defaults](data-sources--network_connector--reference--group-001.md#canonical-1210212231203113-3030320011301310-0021101231021033-3212301010003311-3110032003131130-3332200230021020-0233013310130103-0133211321212221) |
| `enable_forward_proxy.tls_intercept.enable_for_all_domains` | [enable_forward_proxy.tls_intercept.enable_for_all_domains](data-sources--network_connector--reference--group-001.md#canonical-3123223011301100-3202313313111321-2120001012111221-0323321131303020-3323300111312130-1221233023203011-1310322021232123-1311100332133002) |
| `enable_forward_proxy.tls_intercept.policy` | [enable_forward_proxy.tls_intercept.policy](data-sources--network_connector--reference--group-001.md#canonical-3011200031021313-2002100011313000-1323013322201310-1201322103333111-2122213210011312-1231011102330123-1020110302102313-1122211322010010) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules` | [enable_forward_proxy.tls_intercept.policy.interception_rules](data-sources--network_connector--reference--group-001.md#canonical-2101301310131001-3131130103321301-3202200033013222-3301321131323120-0232233233230302-0121122231132303-2113332213332121-0020201120200121) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.disable_interception` | [enable_forward_proxy.tls_intercept.policy.interception_rules.disable_interception](data-sources--network_connector--reference--group-001.md#canonical-0231121100030001-3210102321330103-0332203313213030-1003310323312013-0000011320220311-0131001103221231-0102211021313211-2320101001001111) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match` | [enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match](data-sources--network_connector--reference--group-001.md#canonical-2322132333223222-1112010213133210-0223121320323111-0310231011021132-1201133130302112-3212330202102123-2332201021100202-2013302002232210) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.exact_value` | [enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.exact_value](data-sources--network_connector--reference--group-001.md#canonical-1203001321200033-0110130130123202-0203110230311033-1331113202122011-0332210120020221-2303031122221122-0000133320221021-3021101330002303) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.regex_value` | [enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.regex_value](data-sources--network_connector--reference--group-001.md#canonical-3013111011332111-1010112203010133-1032332122002312-0110222000211123-2300303133003223-2201032200120121-1233202131000032-3211100112120002) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.suffix_value` | [enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.suffix_value](data-sources--network_connector--reference--group-001.md#canonical-1001320300331310-2122001000230120-1001211312023313-0210332302230122-2313200033211221-1131330132212011-3221220330030123-3111323012213233) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.enable_interception` | [enable_forward_proxy.tls_intercept.policy.interception_rules.enable_interception](data-sources--network_connector--reference--group-001.md#canonical-2012303310123122-2233002221100200-2133210301221102-0331230322313020-1233311000320120-0022101213232111-2121013220000010-0102101012311131) |
| `enable_forward_proxy.tls_intercept.trusted_ca_url` | [enable_forward_proxy.tls_intercept.trusted_ca_url](data-sources--network_connector--reference--group-001.md#canonical-2001202031300320-3311221033102021-1210323122322320-0133033110211000-3033322000212301-0330322110032001-2123102201201022-1221111303303200) |
| `enable_forward_proxy.tls_intercept.volterra_certificate` | [enable_forward_proxy.tls_intercept.volterra_certificate](data-sources--network_connector--reference--group-001.md#canonical-3023030203301212-1231321030002123-3321331001113120-3123212002331330-1222330303002311-1313232112101202-1333223103222212-0000110103123113) |
| `enable_forward_proxy.tls_intercept.volterra_trusted_ca` | [enable_forward_proxy.tls_intercept.volterra_trusted_ca](data-sources--network_connector--reference--group-001.md#canonical-1011132321322221-1230121033032102-2302113223300222-2233002231023130-1222220321223300-3320222032010013-1110021311313301-1330223121233033) |
| `enable_forward_proxy.white_listed_ports` | [enable_forward_proxy.white_listed_ports](data-sources--network_connector--reference--group-001.md#canonical-1211323230300311-0111310210130003-1321112223230032-2120012113303133-0221321320231230-0033130220113112-3023010032302031-3031333130321200) |
| `enable_forward_proxy.white_listed_prefixes` | [enable_forward_proxy.white_listed_prefixes](data-sources--network_connector--reference--group-001.md#canonical-3022232102311001-1021323221210201-0020120211320213-3303022002211311-2201012223222202-3130101331130231-2133100212300313-2102001001130232) |
| `id` | [ID](data-sources--network_connector--reference--group-001.md#canonical-1113003133013031-0323002212113003-1310023332031123-0333212111130313-2111003233233021-2200323103113010-0101231031322111-3112020213320011) |
| `labels` | [labels](data-sources--network_connector--reference--group-001.md#canonical-0220111310323203-0323002211331010-3021302011031333-3300002012202213-3123232333020120-1010330110003123-3130013212322102-3012233231211111) |
| `name` | [name](data-sources--network_connector--reference--group-001.md#canonical-1031111202233200-3312311020312033-2130231310322113-3002231301102312-0311322102221313-0131023120023022-3003001220102123-2202021313332101) |
| `namespace` | [namespace](data-sources--network_connector--reference--group-001.md#canonical-2101032230001323-2211310321311200-1103123312231202-1130022100110012-0102323200113230-3211221210332301-1333330221313013-3113021030110231) |
| `sli_to_global_dr` | [sli_to_global_dr](data-sources--network_connector--reference--group-001.md#canonical-2003020113133201-0231001120300203-3122203011233032-1030031303023020-0003023321110312-3002231130022331-1020011333223303-0122002000212111) |
| `sli_to_global_dr.global_vn` | [sli_to_global_dr.global_vn](data-sources--network_connector--reference--group-001.md#canonical-3322221220213323-3223132110311133-1300333010230113-0220020012020110-3303333333001301-0300313001220231-1313303002030212-3000301203023033) |
| `sli_to_global_dr.global_vn.name` | [sli_to_global_dr.global_vn.name](data-sources--network_connector--reference--group-001.md#canonical-0013333002021302-0020321313103112-3113122131221022-3311303011021231-3122023330221103-0033300112233121-3123001210122131-2302310100123303) |
| `sli_to_global_dr.global_vn.namespace` | [sli_to_global_dr.global_vn.namespace](data-sources--network_connector--reference--group-001.md#canonical-1203311001321310-3130201301112022-3310301032030203-0011201322113013-2221103012322102-0233000122023131-3331013222203133-1103103103001100) |
| `sli_to_global_dr.global_vn.tenant` | [sli_to_global_dr.global_vn.tenant](data-sources--network_connector--reference--group-001.md#canonical-2333223121322002-3223130311111131-1213123112033030-2101233011201132-1112322320010120-0000220000033120-0001001211023122-3231322230020001) |
| `sli_to_slo_snat` | [sli_to_slo_snat](data-sources--network_connector--reference--group-001.md#canonical-2120003113110001-3210230121130331-2222110200221200-3321331001311301-2323022032112122-0213302221020112-0323032100010023-3133110321333213) |
| `sli_to_slo_snat.default_gw_snat` | [sli_to_slo_snat.default_gw_snat](data-sources--network_connector--reference--group-001.md#canonical-2310312133323330-3323001120321100-3232020203221132-2031212132000001-2103213113312200-1311222013213102-2020121330100310-1102332221213212) |
| `sli_to_slo_snat.interface_ip` | [sli_to_slo_snat.interface_ip](data-sources--network_connector--reference--group-001.md#canonical-2231312101122202-0022301123123033-3230100330001113-3133323022201213-3302021133220330-0133210302212012-3031213103013210-3230010333213012) |
| `slo_to_global_dr` | [slo_to_global_dr](data-sources--network_connector--reference--group-001.md#canonical-3113202033322122-3023213330323003-1102322112312010-3012211202021201-1200302021133313-3132010220021033-3220102232031032-2221230302300100) |
| `slo_to_global_dr.global_vn` | [slo_to_global_dr.global_vn](data-sources--network_connector--reference--group-001.md#canonical-0012113101332002-1221230321331030-3031322022220013-3200222211133023-3132203232312111-1100003300100122-3021230213003322-0130001230132222) |
| `slo_to_global_dr.global_vn.name` | [slo_to_global_dr.global_vn.name](data-sources--network_connector--reference--group-001.md#canonical-1021100223133002-0312200130110110-2133112320321301-1022321233030231-2121100213100203-0012203110101130-1120030211100300-2302031311111000) |
| `slo_to_global_dr.global_vn.namespace` | [slo_to_global_dr.global_vn.namespace](data-sources--network_connector--reference--group-001.md#canonical-2112112321313303-2001133001303301-1131012330122320-3131002322331221-3323021233020221-0330033331111102-2233332311132122-0301233221101011) |
| `slo_to_global_dr.global_vn.tenant` | [slo_to_global_dr.global_vn.tenant](data-sources--network_connector--reference--group-001.md#canonical-3020322222031203-1022033021113110-1322131033220113-1330101311023202-3100132222233032-0121111122201023-2200302320020102-3223121003321211) |

<a id="canonical-3211303131123112-2030112332333102-0020201002021110-3311100303310310-3330212213323011-0113203133331032-2211011223233020-3122112200133110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_forward_proxy` properties

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-3100100033113131-1013100333112132-1220212103012023-0211023132210233-2323121020123322-2323213203113121-1203301130023023-3202032133312301)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-3000122203132213-2030301322323023-1302132111133302-0302223201130030-0303221223333310-2021321132132200-3100201000113230-1311223311010320)
- disable_forward_proxy

<a id="canonical-2300212230202233-2033333212221330-0133131003033032-2132303032213322-3023121000133232-3223010030100200-3010210113300302-1130030032201312"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_forward\_proxy, enable\_forward\_proxy; Default: disable\_forward\_proxy\]
Configuration parameter for disable forward proxy.

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

OneOf alternatives in this subsection:

- [disable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-2300212230202233-2033333212221330-0133131003033032-2132303032213322-3023121000133232-3223010030100200-3010210113300302-1130030032201312)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-0113033202113303-3332112213100002-1230112313310120-1230103301321322-3133011001013112-3023023321123133-3012303202130201-3220121110330123)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2213111031121031-3003322303011211-0020021331302310-0213312231333222-0202212313032220-0002012131031003-0322101122200102-1301210120231201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy` properties

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-3100100033113131-1013100333112132-1220212103012023-0211023132210233-2323121020123322-2323213203113121-1203301130023023-3202032133312301)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-3000122203132213-2030301322323023-1302132111133302-0302223201130030-0303221223333310-2021321132132200-3100201000113230-1311223311010320)
- enable_forward_proxy

<a id="canonical-0113033202113303-3332112213100002-1230112313310120-1230103301321322-3133011001013112-3023023321123133-3012303202130201-3220121110330123"></a>

Type: `"single"`. Computed.

Fine tune forward proxy behavior

Few configurations allowed are

White listed ports and IP prefixes: Forward proxy does application protocol detection and server
name(SNI) detection by peeking into the traffic on the incoming downstream connection. Few protocols
doesn't have client sending the first data. In such cases, protocol and SNI detection fails. This
configuration allows, skipping protocol and SNI detection for whitelisted IP-prefix-list and ports
connection\_timeout: The timeout for new network connections to upstream server.
Max\_connect\_attempts: Maximum number of attempts made to make new network connection to upstream
server.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-tls_interception_choice": "[\"no_interception\",\"tls_intercept\"]"
}
```

<a id="canonical-2031121302110322-0003303132011023-0030330201131133-2101000023003310-1100312212113033-1101102122231203-3022010203112122-0223223300212311"></a>

### Direct properties for `enable_forward_proxy`

<a id="canonical-0020203023310230-1010300112323110-0002220211002102-1102230231111323-2101322031003030-1331133213330313-3131313301010033-0313203000033321"></a>

#### `enable_forward_proxy.connection_timeout` property

Type: `"number"`. Computed.

The timeout for new network connections to upstream server. This is specified in milliseconds. The
(2 seconds). Defaults to \`2000\`.

Additional upstream details:

The default value is 2000 (2 seconds)

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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-3321213213001130-1213023303223132-1001213111320300-2321200212321012-2301311123012012-1032210210110110-2021020023302133-0321020122221031"></a>

<a id="canonical-3230101330031122-2320023030313113-1330101000222312-3003020001302330-2102012310032101-1121000212223113-3302003233000322-1131323202213012"></a>

#### `enable_forward_proxy.max_connect_attempts` property

Type: `"number"`. Computed.

Specifies the allowed number of retries on connect failure to upstream server. Defaults to \`1\`.

Additional upstream details:

Defaults to 1.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "8"
  }
}
```

- [no_interception](data-sources--network_connector--reference--group-001.md#canonical-3000231310101202-3310310330211003-2113332123130331-0221332323310000-3210020032202311-2131020032333131-0021231310033003-0333123303010313): complete subsection reference.

- [tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-0332312230322331-2212330302231332-0303003223311003-2133233120210113-3111113131100030-0003220131100301-3333012301333000-0230313231123320): complete subsection reference.

<a id="canonical-1211323230300311-0111310210130003-1321112223230032-2120012113303133-0221321320231230-0033130220113112-3023010032302031-3031333130321200"></a>

<a id="canonical-0110102233311221-1231322313002011-3302333321110230-0110011322333111-2330211010223311-3102003331223032-1211231122123332-0321221202033332"></a>

#### `enable_forward_proxy.white_listed_ports` property

Type: `["list", "number"]`. Computed.

Traffic to these destination TCP ports is not subjected to protocol parsing Example 'tmate' server
port.

Additional upstream details:

Traffic to these destination TCP ports is not subjected to protocol parsing Example "tmate" server
port.

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
    "ves.io.schema.rules.repeated.items.uint32.lte": "65535",
    "ves.io.schema.rules.repeated.max_items": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.uint32.lte": "65535",
    "ves.io.schema.rules.repeated.max_items": "64"
  }
}
```

<a id="canonical-3022232102311001-1021323221210201-0020120211320213-3303022002211311-2201012223222202-3130101331130231-2133100212300313-2102001001130232"></a>

<a id="canonical-3111221020023203-1211000203021110-0200022012231101-3323113002032101-2000021101202010-1000133330110012-1300122123000012-3210010130302032"></a>

#### `enable_forward_proxy.white_listed_prefixes` property

Type: `["list", "string"]`. Computed.

Traffic to these destination IP prefixes is not subjected to protocol parsing Example 'tmate' server
IP.

Additional upstream details:

Traffic to these destination IP prefixes is not subjected to protocol parsing Example "tmate" server
IP.

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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3000231310101202-3310310330211003-2113332123130331-0221332323310000-3210020032202311-2131020032333131-0021231310033003-0333123303010313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy.no_interception` properties

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-3100100033113131-1013100333112132-1220212103012023-0211023132210233-2323121020123322-2323213203113121-1203301130023023-3202032133312301)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-3000122203132213-2030301322323023-1302132111133302-0302223201130030-0303221223333310-2021321132132200-3100201000113230-1311223311010320)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-2213111031121031-3003322303011211-0020021331302310-0213312231333222-0202212313032220-0002012131031003-0322101122200102-1301210120231201)
- enable_forward_proxy.no_interception

<a id="canonical-0013332213310001-2013103333213302-2231223231023001-1333230330300230-1223313020321302-1232120213020310-2230211230131001-0010312013202210"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no interception.

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

<a id="canonical-0332312230322331-2212330302231332-0303003223311003-2133233120210113-3111113131100030-0003220131100301-3333012301333000-0230313231123320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy.tls_intercept` properties

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-3100100033113131-1013100333112132-1220212103012023-0211023132210233-2323121020123322-2323213203113121-1203301130023023-3202032133312301)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-3000122203132213-2030301322323023-1302132111133302-0302223201130030-0303221223333310-2021321132132200-3100201000113230-1311223311010320)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-2213111031121031-3003322303011211-0020021331302310-0213312231333222-0202212313032220-0002012131031003-0322101122200102-1301210120231201)
- enable_forward_proxy.tls_intercept

<a id="canonical-0113032323000303-1302332013323013-3023132023020330-3322131122120302-3021211301200022-2312003132303032-0111331002301300-3332121103220211"></a>

Type: `"single"`. Computed.

Configuration to enable TLS interception.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-interception_policy_choice": "[\"enable_for_all_domains\",\"policy\"]",
  "x-ves-oneof-field-signing_cert_choice": "[\"custom_certificate\",\"volterra_certificate\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca_url\",\"volterra_trusted_ca\"]"
}
```

<a id="canonical-1301133200313112-3321000300030120-1322322300122013-0123232321003322-3233031120313323-2023023100123132-1223213001201122-1010313213020332"></a>

### Direct properties for `enable_forward_proxy.tls_intercept`

- [custom_certificate](data-sources--network_connector--reference--group-001.md#canonical-0102033312210312-0312233113133221-2322122112201330-2220002332303101-0131133020002232-2301230012022220-1332000132313000-0032310333232013): complete subsection reference.

- [enable_for_all_domains](data-sources--network_connector--reference--group-001.md#canonical-0102120221011332-0310301021113020-0023031202030023-1120130231210232-3032111321330132-3132221203012311-1223123222111321-1013321130231130): complete subsection reference.

- [policy](data-sources--network_connector--reference--group-001.md#canonical-0322223230112333-2221330301302322-3001221023121130-1302003012133113-3010000301221221-3130111103223130-0012102030021023-1032331220111311): complete subsection reference.

<a id="canonical-2001202031300320-3311221033102021-1210323122322320-0133033110211000-3033322000212301-0330322110032001-2123102201201022-1221111303303200"></a>

<a id="canonical-0000033123201001-3203332101310220-2211113232322021-3012233123312103-2313120333333230-3322133330231232-2232011020113000-1130003103221332"></a>

#### `enable_forward_proxy.tls_intercept.trusted_ca_url` property

Type: `"string"`. Computed.

Exclusive with \[volterra\_trusted\_ca\] Custom Root CA Certificate for validating upstream server
certificate.

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
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [volterra_certificate](data-sources--network_connector--reference--group-001.md#canonical-1223100201030011-2223111320311011-1300313332133110-0030313212133302-2132032200331233-3231201303132113-2113111100133111-0030231310102320): complete subsection reference.

- [volterra_trusted_ca](data-sources--network_connector--reference--group-001.md#canonical-1022113031132322-3033322111012101-3212222000220311-3200220222111210-2313231122022122-3311031010210333-2313331133233332-1333113123033020): complete subsection reference.

<a id="canonical-0102033312210312-0312233113133221-2322122112201330-2220002332303101-0131133020002232-2301230012022220-1332000132313000-0032310333232013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy.tls_intercept.custom_certificate` properties

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-3100100033113131-1013100333112132-1220212103012023-0211023132210233-2323121020123322-2323213203113121-1203301130023023-3202032133312301)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-3000122203132213-2030301322323023-1302132111133302-0302223201130030-0303221223333310-2021321132132200-3100201000113230-1311223311010320)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-2213111031121031-3003322303011211-0020021331302310-0213312231333222-0202212313032220-0002012131031003-0322101122200102-1301210120231201)
- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-0332312230322331-2212330302231332-0303003223311003-2133233120210113-3111113131100030-0003220131100301-3333012301333000-0230313231123320)
- enable_forward_proxy.tls_intercept.custom_certificate

<a id="canonical-2303323322330003-3301200112010113-2011300133003121-0321203222023210-2331301030200033-3232123300012133-2123002113321330-0030023222012310"></a>

Type: `"single"`. Computed.

Configuration parameter for custom certificate.

Additional upstream details:

Handle to fetch certificate and key.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ocsp_stapling_choice": "[\"custom_hash_algorithms\",\"disable_ocsp_stapling\",\"use_system_defaults\"]"
}
```

<a id="canonical-0000100000123012-1132201221201132-0302221010023321-3221333300303323-1213022011113213-1301200023123230-1310132301110201-1311213320000121"></a>

### Direct properties for `enable_forward_proxy.tls_intercept.custom_certificate`

<a id="canonical-3232021001330012-3120200300103320-0013321031330221-0020132233221321-3002023330221331-1330003133103013-1101320200100020-3200020013301011"></a>

#### `enable_forward_proxy.tls_intercept.custom_certificate.certificate_url` property

Type: `"string"`. Computed.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [custom_hash_algorithms](data-sources--network_connector--reference--group-001.md#canonical-1311323231302230-0120232122032300-0031011001000101-2011200202030222-3023310022330000-0331332311301121-1012003330123301-1100232320103030): complete subsection reference.

<a id="canonical-3320322202231222-2213203222133131-0221002012212033-1202121020300120-3323122103303023-1321030112201120-3322310013223310-0213310023022311"></a>

<a id="canonical-0130101200032220-2010312231221120-0103012313010312-1201330130313201-3022032323003301-0003303130310221-1122323320031133-3321323201201112"></a>

#### `enable_forward_proxy.tls_intercept.custom_certificate.description_spec` property

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--network_connector--reference--group-001.md#canonical-2232332020203321-0102231002222112-0311033310132020-2312223212301013-1120333001312213-1110233223232202-2033333322031301-2322313121003011): complete subsection reference.

- [private_key](data-sources--network_connector--reference--group-001.md#canonical-0210121113331122-2223233132220213-3020031300302212-0323332120211011-0111011220000200-2021220120312010-2000112133312310-1130021333331212): complete subsection reference.

- [use_system_defaults](data-sources--network_connector--reference--group-001.md#canonical-2301020203110021-0321121233103112-0122032033030130-2232122200231303-0010202330332113-2113301320032330-2303213113131233-0101131102013200): complete subsection reference.

<a id="canonical-1311323231302230-0120232122032300-0031011001000101-2011200202030222-3023310022330000-0331332311301121-1012003330123301-1100232320103030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-3100100033113131-1013100333112132-1220212103012023-0211023132210233-2323121020123322-2323213203113121-1203301130023023-3202032133312301)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-3000122203132213-2030301322323023-1302132111133302-0302223201130030-0303221223333310-2021321132132200-3100201000113230-1311223311010320)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-2213111031121031-3003322303011211-0020021331302310-0213312231333222-0202212313032220-0002012131031003-0322101122200102-1301210120231201)
- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-0332312230322331-2212330302231332-0303003223311003-2133233120210113-3111113131100030-0003220131100301-3333012301333000-0230313231123320)
- [enable_forward_proxy.tls_intercept.custom_certificate](data-sources--network_connector--reference--group-001.md#canonical-0102033312210312-0312233113133221-2322122112201330-2220002332303101-0131133020002232-2301230012022220-1332000132313000-0032310333232013)
- enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms

<a id="canonical-3310113002133021-2322011133131302-1122202312100103-2023323130321200-1203321223213120-2312302033201130-2303132120310120-3032113323123002"></a>

Type: `"single"`. Computed.

Specifies the hash algorithms to be used.

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

<a id="canonical-3103210332131312-3101033212001122-3032001213121021-2323231311002320-2131101331031231-3031300022011112-1002012232201123-3003301101120103"></a>

### Direct properties for `enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms`

<a id="canonical-2201212221111023-0213323100310330-0330120311230020-2000223010123313-2000302320323210-0200011112221332-1201122300312203-0122320001220322"></a>

#### `enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms.hash_algorithms` property

Type: `["list", "string"]`. Computed.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2232332020203321-0102231002222112-0311033310132020-2312223212301013-1120333001312213-1110233223232202-2033333322031301-2322313121003011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-3100100033113131-1013100333112132-1220212103012023-0211023132210233-2323121020123322-2323213203113121-1203301130023023-3202032133312301)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-3000122203132213-2030301322323023-1302132111133302-0302223201130030-0303221223333310-2021321132132200-3100201000113230-1311223311010320)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-2213111031121031-3003322303011211-0020021331302310-0213312231333222-0202212313032220-0002012131031003-0322101122200102-1301210120231201)
- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-0332312230322331-2212330302231332-0303003223311003-2133233120210113-3111113131100030-0003220131100301-3333012301333000-0230313231123320)
- [enable_forward_proxy.tls_intercept.custom_certificate](data-sources--network_connector--reference--group-001.md#canonical-0102033312210312-0312233113133221-2322122112201330-2220002332303101-0131133020002232-2301230012022220-1332000132313000-0032310333232013)
- enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling

<a id="canonical-0003130303331010-2023121332002231-1120102020013032-3101331230313321-3101310211321123-0120102313331121-0322021310202112-1332103303303300"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable ocsp stapling.

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

<a id="canonical-0210121113331122-2223233132220213-3020031300302212-0323332120211011-0111011220000200-2021220120312010-2000112133312310-1130021333331212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy.tls_intercept.custom_certificate.private_key` properties

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-3100100033113131-1013100333112132-1220212103012023-0211023132210233-2323121020123322-2323213203113121-1203301130023023-3202032133312301)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-3000122203132213-2030301322323023-1302132111133302-0302223201130030-0303221223333310-2021321132132200-3100201000113230-1311223311010320)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-2213111031121031-3003322303011211-0020021331302310-0213312231333222-0202212313032220-0002012131031003-0322101122200102-1301210120231201)
- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-0332312230322331-2212330302231332-0303003223311003-2133233120210113-3111113131100030-0003220131100301-3333012301333000-0230313231123320)
- [enable_forward_proxy.tls_intercept.custom_certificate](data-sources--network_connector--reference--group-001.md#canonical-0102033312210312-0312233113133221-2322122112201330-2220002332303101-0131133020002232-2301230012022220-1332000132313000-0032310333232013)
- enable_forward_proxy.tls_intercept.custom_certificate.private_key

<a id="canonical-1120232100211123-3103020231032030-3010313010332001-0301102220103133-2010213300333330-2102111303211331-3012203230221030-3333212120322320"></a>

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

<a id="canonical-3121111113013132-1220031222201130-2210121132201131-0333231232333102-0331031110121301-0000211012020133-0100130203003002-1331101122102213"></a>

### Direct properties for `enable_forward_proxy.tls_intercept.custom_certificate.private_key`

- [blindfold_secret_info](data-sources--network_connector--reference--group-001.md#canonical-2133022320023030-2322113202310120-3131332113100102-1012300113322212-0012202012230102-0330201232013311-0003103032002331-1320010121021111): complete subsection reference.

- [clear_secret_info](data-sources--network_connector--reference--group-001.md#canonical-0022332323220021-1310123112330231-2030211032213032-0310001321301001-3013333001330123-2302021222303220-3010321103022110-0221013221301233): complete subsection reference.

<a id="canonical-2133022320023030-2322113202310120-3131332113100102-1012300113322212-0012202012230102-0330201232013311-0003103032002331-1320010121021111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-3100100033113131-1013100333112132-1220212103012023-0211023132210233-2323121020123322-2323213203113121-1203301130023023-3202032133312301)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-3000122203132213-2030301322323023-1302132111133302-0302223201130030-0303221223333310-2021321132132200-3100201000113230-1311223311010320)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-2213111031121031-3003322303011211-0020021331302310-0213312231333222-0202212313032220-0002012131031003-0322101122200102-1301210120231201)
- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-0332312230322331-2212330302231332-0303003223311003-2133233120210113-3111113131100030-0003220131100301-3333012301333000-0230313231123320)
- [enable_forward_proxy.tls_intercept.custom_certificate](data-sources--network_connector--reference--group-001.md#canonical-0102033312210312-0312233113133221-2322122112201330-2220002332303101-0131133020002232-2301230012022220-1332000132313000-0032310333232013)
- [enable_forward_proxy.tls_intercept.custom_certificate.private_key](data-sources--network_connector--reference--group-001.md#canonical-0210121113331122-2223233132220213-3020031300302212-0323332120211011-0111011220000200-2021220120312010-2000112133312310-1130021333331212)
- enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info

<a id="canonical-0133020121202030-0010301030302000-1200200330000020-0200233302130001-3123110110211030-0212200110110313-0320313103000112-1103231120110232"></a>

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

<a id="canonical-1011001320020130-2110100300203310-1330230211322230-3131001301021331-2123221212302121-1130102012123202-1100020222222313-1213030001023312"></a>

### Direct properties for `enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info`

<a id="canonical-1000101301312223-2220133133033221-2023001210221302-3131300301330110-1101023323300000-3102100032230110-0013103133203123-0021232202203313"></a>

#### `enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-3101301323333101-3011102211011320-1322230012232312-0223233331120232-3131002032002021-2100321103321210-3103211202332032-0331312333302020"></a>

<a id="canonical-2233032123301100-0311210103000013-0332301203122322-0320011333202321-2301002003030301-0211310332023311-2032210112212330-1001231313330111"></a>

#### `enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1003332222222212-1010233023213031-0230223310001123-2202200303011101-1001103011202333-3210330121222123-0113122200033323-0120302231300033"></a>

<a id="canonical-2001113113000122-1101230303011030-3312203130333212-1213302103301130-1200102313212112-1112003032100230-1103212210113321-3221000103202013"></a>

#### `enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-0022332323220021-1310123112330231-2030211032213032-0310001321301001-3013333001330123-2302021222303220-3010321103022110-0221013221301233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-3100100033113131-1013100333112132-1220212103012023-0211023132210233-2323121020123322-2323213203113121-1203301130023023-3202032133312301)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-3000122203132213-2030301322323023-1302132111133302-0302223201130030-0303221223333310-2021321132132200-3100201000113230-1311223311010320)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-2213111031121031-3003322303011211-0020021331302310-0213312231333222-0202212313032220-0002012131031003-0322101122200102-1301210120231201)
- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-0332312230322331-2212330302231332-0303003223311003-2133233120210113-3111113131100030-0003220131100301-3333012301333000-0230313231123320)
- [enable_forward_proxy.tls_intercept.custom_certificate](data-sources--network_connector--reference--group-001.md#canonical-0102033312210312-0312233113133221-2322122112201330-2220002332303101-0131133020002232-2301230012022220-1332000132313000-0032310333232013)
- [enable_forward_proxy.tls_intercept.custom_certificate.private_key](data-sources--network_connector--reference--group-001.md#canonical-0210121113331122-2223233132220213-3020031300302212-0323332120211011-0111011220000200-2021220120312010-2000112133312310-1130021333331212)
- enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info

<a id="canonical-3001313212310332-0102032111122330-1111131211232332-3203222312022101-1123330312020031-0031133213122202-3222033310233000-2123101230112121"></a>

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

<a id="canonical-2313121233221313-0303121130313200-1222133213013011-0300110032021010-2033230111013020-0002123101201122-3313123301323131-0303220000000013"></a>

### Direct properties for `enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info`

<a id="canonical-2322330133021011-1011011100310000-3233301310332122-2000202200020032-3023321021200022-2202312212301332-1101223000011122-0033303121313322"></a>

#### `enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0311330313230222-1023311203310132-3021131201320103-0230023032302322-0212001122220133-3103302111002203-2312200312323232-3123120203132202"></a>

<a id="canonical-0020121103133031-0030202033310022-1323010020300101-0110033203301201-2030131322002110-3021200222233210-3201100011013133-0000010211111103"></a>

#### `enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2301020203110021-0321121233103112-0122032033030130-2232122200231303-0010202330332113-2113301320032330-2303213113131233-0101131102013200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy.tls_intercept.custom_certificate.use_system_defaults` properties

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-3100100033113131-1013100333112132-1220212103012023-0211023132210233-2323121020123322-2323213203113121-1203301130023023-3202032133312301)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-3000122203132213-2030301322323023-1302132111133302-0302223201130030-0303221223333310-2021321132132200-3100201000113230-1311223311010320)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-2213111031121031-3003322303011211-0020021331302310-0213312231333222-0202212313032220-0002012131031003-0322101122200102-1301210120231201)
- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-0332312230322331-2212330302231332-0303003223311003-2133233120210113-3111113131100030-0003220131100301-3333012301333000-0230313231123320)
- [enable_forward_proxy.tls_intercept.custom_certificate](data-sources--network_connector--reference--group-001.md#canonical-0102033312210312-0312233113133221-2322122112201330-2220002332303101-0131133020002232-2301230012022220-1332000132313000-0032310333232013)
- enable_forward_proxy.tls_intercept.custom_certificate.use_system_defaults

<a id="canonical-1210212231203113-3030320011301310-0021101231021033-3212301010003311-3110032003131130-3332200230021020-0233013310130103-0133211321212221"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for use system defaults.

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

<a id="canonical-0102120221011332-0310301021113020-0023031202030023-1120130231210232-3032111321330132-3132221203012311-1223123222111321-1013321130231130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy.tls_intercept.enable_for_all_domains` properties

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-3100100033113131-1013100333112132-1220212103012023-0211023132210233-2323121020123322-2323213203113121-1203301130023023-3202032133312301)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-3000122203132213-2030301322323023-1302132111133302-0302223201130030-0303221223333310-2021321132132200-3100201000113230-1311223311010320)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-2213111031121031-3003322303011211-0020021331302310-0213312231333222-0202212313032220-0002012131031003-0322101122200102-1301210120231201)
- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-0332312230322331-2212330302231332-0303003223311003-2133233120210113-3111113131100030-0003220131100301-3333012301333000-0230313231123320)
- enable_forward_proxy.tls_intercept.enable_for_all_domains

<a id="canonical-3123223011301100-3202313313111321-2120001012111221-0323321131303020-3323300111312130-1221233023203011-1310322021232123-1311100332133002"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable for all domains.

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

<a id="canonical-0322223230112333-2221330301302322-3001221023121130-1302003012133113-3010000301221221-3130111103223130-0012102030021023-1032331220111311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy.tls_intercept.policy` properties

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-3100100033113131-1013100333112132-1220212103012023-0211023132210233-2323121020123322-2323213203113121-1203301130023023-3202032133312301)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-3000122203132213-2030301322323023-1302132111133302-0302223201130030-0303221223333310-2021321132132200-3100201000113230-1311223311010320)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-2213111031121031-3003322303011211-0020021331302310-0213312231333222-0202212313032220-0002012131031003-0322101122200102-1301210120231201)
- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-0332312230322331-2212330302231332-0303003223311003-2133233120210113-3111113131100030-0003220131100301-3333012301333000-0230313231123320)
- enable_forward_proxy.tls_intercept.policy

<a id="canonical-3011200031021313-2002100011313000-1323013322201310-1201322103333111-2122213210011312-1231011102330123-1020110302102313-1122211322010010"></a>

Type: `"single"`. Computed.

Policy to enable or disable TLS interception.

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

<a id="canonical-3221321211231310-0221332310021120-1123301302232302-3133222102221223-1311322210112211-2231303300012101-2000332222111223-3123130110310203"></a>

### Direct properties for `enable_forward_proxy.tls_intercept.policy`

- [interception_rules](data-sources--network_connector--reference--group-001.md#canonical-3100030312323221-1220233133200221-2233203233210222-0120210102323100-1212122100312022-0112333012323333-3122220022013032-2132331300021210): complete subsection reference.

<a id="canonical-3100030312323221-1220233133200221-2233203233210222-0120210102323100-1212122100312022-0112333012323333-3122220022013032-2132331300021210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy.tls_intercept.policy.interception_rules` properties

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-3100100033113131-1013100333112132-1220212103012023-0211023132210233-2323121020123322-2323213203113121-1203301130023023-3202032133312301)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-3000122203132213-2030301322323023-1302132111133302-0302223201130030-0303221223333310-2021321132132200-3100201000113230-1311223311010320)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-2213111031121031-3003322303011211-0020021331302310-0213312231333222-0202212313032220-0002012131031003-0322101122200102-1301210120231201)
- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-0332312230322331-2212330302231332-0303003223311003-2133233120210113-3111113131100030-0003220131100301-3333012301333000-0230313231123320)
- [enable_forward_proxy.tls_intercept.policy](data-sources--network_connector--reference--group-001.md#canonical-0322223230112333-2221330301302322-3001221023121130-1302003012133113-3010000301221221-3130111103223130-0012102030021023-1032331220111311)
- enable_forward_proxy.tls_intercept.policy.interception_rules

<a id="canonical-2101301310131001-3131130103321301-3202200033013222-3301321131323120-0232233233230302-0121122231132303-2113332213332121-0020201120200121"></a>

Type: `"list"`. Computed.

List of ordered rules to enable or disable for TLS interception.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1000030122113233-1211010233111212-1233211322203231-3120103202120230-1301121221110111-3001322322300211-2022233131321131-3131323003021303"></a>

### Direct properties for `enable_forward_proxy.tls_intercept.policy.interception_rules`

- [disable_interception](data-sources--network_connector--reference--group-001.md#canonical-2321300102212233-0110333122113132-0302331231230122-2201002223030001-2023130133022210-0020033112333112-3220301210333133-2120120012232322): complete subsection reference.

- [domain_match](data-sources--network_connector--reference--group-001.md#canonical-0122123032022002-1030232023330031-3103332213321003-2210031130223021-2021011311010330-0200232031130133-0112212001003313-1021222130030312): complete subsection reference.

- [enable_interception](data-sources--network_connector--reference--group-001.md#canonical-1111311222013033-3332213202233312-3201121030110021-1011223103030023-1332121202012010-0010220033222103-2112203133031120-2313313213231121): complete subsection reference.

<a id="canonical-2321300102212233-0110333122113132-0302331231230122-2201002223030001-2023130133022210-0020033112333112-3220301210333133-2120120012232322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy.tls_intercept.policy.interception_rules.disable_interception` properties

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-3100100033113131-1013100333112132-1220212103012023-0211023132210233-2323121020123322-2323213203113121-1203301130023023-3202032133312301)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-3000122203132213-2030301322323023-1302132111133302-0302223201130030-0303221223333310-2021321132132200-3100201000113230-1311223311010320)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-2213111031121031-3003322303011211-0020021331302310-0213312231333222-0202212313032220-0002012131031003-0322101122200102-1301210120231201)
- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-0332312230322331-2212330302231332-0303003223311003-2133233120210113-3111113131100030-0003220131100301-3333012301333000-0230313231123320)
- [enable_forward_proxy.tls_intercept.policy](data-sources--network_connector--reference--group-001.md#canonical-0322223230112333-2221330301302322-3001221023121130-1302003012133113-3010000301221221-3130111103223130-0012102030021023-1032331220111311)
- [enable_forward_proxy.tls_intercept.policy.interception_rules](data-sources--network_connector--reference--group-001.md#canonical-3100030312323221-1220233133200221-2233203233210222-0120210102323100-1212122100312022-0112333012323333-3122220022013032-2132331300021210)
- enable_forward_proxy.tls_intercept.policy.interception_rules.disable_interception

<a id="canonical-0231121100030001-3210102321330103-0332203313213030-1003310323312013-0000011320220311-0131001103221231-0102211021313211-2320101001001111"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable interception.

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

<a id="canonical-0122123032022002-1030232023330031-3103332213321003-2210031130223021-2021011311010330-0200232031130133-0112212001003313-1021222130030312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match` properties

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-3100100033113131-1013100333112132-1220212103012023-0211023132210233-2323121020123322-2323213203113121-1203301130023023-3202032133312301)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-3000122203132213-2030301322323023-1302132111133302-0302223201130030-0303221223333310-2021321132132200-3100201000113230-1311223311010320)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-2213111031121031-3003322303011211-0020021331302310-0213312231333222-0202212313032220-0002012131031003-0322101122200102-1301210120231201)
- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-0332312230322331-2212330302231332-0303003223311003-2133233120210113-3111113131100030-0003220131100301-3333012301333000-0230313231123320)
- [enable_forward_proxy.tls_intercept.policy](data-sources--network_connector--reference--group-001.md#canonical-0322223230112333-2221330301302322-3001221023121130-1302003012133113-3010000301221221-3130111103223130-0012102030021023-1032331220111311)
- [enable_forward_proxy.tls_intercept.policy.interception_rules](data-sources--network_connector--reference--group-001.md#canonical-3100030312323221-1220233133200221-2233203233210222-0120210102323100-1212122100312022-0112333012323333-3122220022013032-2132331300021210)
- enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match

<a id="canonical-2322132333223222-1112010213133210-0223121320323111-0310231011021132-1201133130302112-3212330202102123-2332201021100202-2013302002232210"></a>

Type: `"single"`. Computed.

Configuration parameter for domain match.

Additional upstream details:

Domains names.

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

<a id="canonical-0103033030202322-0312300223001321-0010310111223032-2033021332222310-3233121213003221-0121330222013321-3131111110303211-3113032223110110"></a>

### Direct properties for `enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match`

<a id="canonical-1203001321200033-0110130130123202-0203110230311033-1331113202122011-0332210120020221-2303031122221122-0000133320221021-3021101330002303"></a>

#### `enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.exact_value` property

Type: `"string"`. Computed.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

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

<a id="canonical-3013111011332111-1010112203010133-1032332122002312-0110222000211123-2300303133003223-2201032200120121-1233202131000032-3211100112120002"></a>

<a id="canonical-2320200020020123-0110310220010130-3321032213302223-1222123021232102-2200123231322303-1130021300202121-0232021023113330-2203131001002103"></a>

#### `enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.regex_value` property

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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

<a id="canonical-1001320300331310-2122001000230120-1001211312023313-0210332302230122-2313200033211221-1131330132212011-3221220330030123-3111323012213233"></a>

<a id="canonical-1222130002332232-1211203123010031-2120023003113032-3112110021003033-3022101012012011-1323020002101112-2132021103011211-1120221120303203"></a>

#### `enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.suffix_value` property

Type: `"string"`. Computed.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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

<a id="canonical-1111311222013033-3332213202233312-3201121030110021-1011223103030023-1332121202012010-0010220033222103-2112203133031120-2313313213231121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy.tls_intercept.policy.interception_rules.enable_interception` properties

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-3100100033113131-1013100333112132-1220212103012023-0211023132210233-2323121020123322-2323213203113121-1203301130023023-3202032133312301)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-3000122203132213-2030301322323023-1302132111133302-0302223201130030-0303221223333310-2021321132132200-3100201000113230-1311223311010320)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-2213111031121031-3003322303011211-0020021331302310-0213312231333222-0202212313032220-0002012131031003-0322101122200102-1301210120231201)
- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-0332312230322331-2212330302231332-0303003223311003-2133233120210113-3111113131100030-0003220131100301-3333012301333000-0230313231123320)
- [enable_forward_proxy.tls_intercept.policy](data-sources--network_connector--reference--group-001.md#canonical-0322223230112333-2221330301302322-3001221023121130-1302003012133113-3010000301221221-3130111103223130-0012102030021023-1032331220111311)
- [enable_forward_proxy.tls_intercept.policy.interception_rules](data-sources--network_connector--reference--group-001.md#canonical-3100030312323221-1220233133200221-2233203233210222-0120210102323100-1212122100312022-0112333012323333-3122220022013032-2132331300021210)
- enable_forward_proxy.tls_intercept.policy.interception_rules.enable_interception

<a id="canonical-2012303310123122-2233002221100200-2133210301221102-0331230322313020-1233311000320120-0022101213232111-2121013220000010-0102101012311131"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable interception.

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

<a id="canonical-1223100201030011-2223111320311011-1300313332133110-0030313212133302-2132032200331233-3231201303132113-2113111100133111-0030231310102320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy.tls_intercept.volterra_certificate` properties

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-3100100033113131-1013100333112132-1220212103012023-0211023132210233-2323121020123322-2323213203113121-1203301130023023-3202032133312301)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-3000122203132213-2030301322323023-1302132111133302-0302223201130030-0303221223333310-2021321132132200-3100201000113230-1311223311010320)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-2213111031121031-3003322303011211-0020021331302310-0213312231333222-0202212313032220-0002012131031003-0322101122200102-1301210120231201)
- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-0332312230322331-2212330302231332-0303003223311003-2133233120210113-3111113131100030-0003220131100301-3333012301333000-0230313231123320)
- enable_forward_proxy.tls_intercept.volterra_certificate

<a id="canonical-3023030203301212-1231321030002123-3321331001113120-3123212002331330-1222330303002311-1313232112101202-1333223103222212-0000110103123113"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for volterra certificate.

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

<a id="canonical-1022113031132322-3033322111012101-3212222000220311-3200220222111210-2313231122022122-3311031010210333-2313331133233332-1333113123033020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy.tls_intercept.volterra_trusted_ca` properties

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-3100100033113131-1013100333112132-1220212103012023-0211023132210233-2323121020123322-2323213203113121-1203301130023023-3202032133312301)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-3000122203132213-2030301322323023-1302132111133302-0302223201130030-0303221223333310-2021321132132200-3100201000113230-1311223311010320)
- [enable_forward_proxy](data-sources--network_connector--reference--group-001.md#canonical-2213111031121031-3003322303011211-0020021331302310-0213312231333222-0202212313032220-0002012131031003-0322101122200102-1301210120231201)
- [enable_forward_proxy.tls_intercept](data-sources--network_connector--reference--group-001.md#canonical-0332312230322331-2212330302231332-0303003223311003-2133233120210113-3111113131100030-0003220131100301-3333012301333000-0230313231123320)
- enable_forward_proxy.tls_intercept.volterra_trusted_ca

<a id="canonical-1011132321322221-1230121033032102-2302113223300222-2233002231023130-1222220321223300-3320222032010013-1110021311313301-1330223121233033"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for volterra trusted ca.

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

<a id="canonical-3202003000033330-0132321012323121-2201002113322223-1301301113020232-1303210303223002-1110011130333032-0202031330113122-0100012123123232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sli_to_global_dr` properties

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-3100100033113131-1013100333112132-1220212103012023-0211023132210233-2323121020123322-2323213203113121-1203301130023023-3202032133312301)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-3000122203132213-2030301322323023-1302132111133302-0302223201130030-0303221223333310-2021321132132200-3100201000113230-1311223311010320)
- sli_to_global_dr

<a id="canonical-2003020113133201-0231001120300203-3122203011233032-1030031303023020-0003023321110312-3002231130022331-1020011333223303-0122002000212111"></a>

Type: `"single"`. Computed.

\[OneOf: sli\_to\_global\_dr, sli\_to\_slo\_snat, slo\_to\_global\_dr\] Global network reference for
direct connection.

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

- [sli_to_global_dr](data-sources--network_connector--reference--group-001.md#canonical-2003020113133201-0231001120300203-3122203011233032-1030031303023020-0003023321110312-3002231130022331-1020011333223303-0122002000212111)
- [sli_to_slo_snat](data-sources--network_connector--reference--group-001.md#canonical-2120003113110001-3210230121130331-2222110200221200-3321331001311301-2323022032112122-0213302221020112-0323032100010023-3133110321333213)
- [slo_to_global_dr](data-sources--network_connector--reference--group-001.md#canonical-3113202033322122-3023213330323003-1102322112312010-3012211202021201-1200302021133313-3132010220021033-3220102232031032-2221230302300100)

Select alternatives according to the provider validators above.

<a id="canonical-2033133312111022-1030132310122223-1232331112223012-0311212232221320-1310011112000211-1033010323100001-3111110321103001-2020200312322203"></a>

### Direct properties for `sli_to_global_dr`

- [global_vn](data-sources--network_connector--reference--group-001.md#canonical-1322133331113313-1123132102111123-0002220312323210-1113130110303320-1112010000230122-2210120102323131-3123222011303013-1302122210303030): complete subsection reference.

<a id="canonical-1322133331113313-1123132102111123-0002220312323210-1113130110303320-1112010000230122-2210120102323131-3123222011303013-1302122210303030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sli_to_global_dr.global_vn` properties

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-3100100033113131-1013100333112132-1220212103012023-0211023132210233-2323121020123322-2323213203113121-1203301130023023-3202032133312301)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-3000122203132213-2030301322323023-1302132111133302-0302223201130030-0303221223333310-2021321132132200-3100201000113230-1311223311010320)
- [sli_to_global_dr](data-sources--network_connector--reference--group-001.md#canonical-3202003000033330-0132321012323121-2201002113322223-1301301113020232-1303210303223002-1110011130333032-0202031330113122-0100012123123232)
- sli_to_global_dr.global_vn

<a id="canonical-3322221220213323-3223132110311133-1300333010230113-0220020012020110-3303333333001301-0300313001220231-1313303002030212-3000301203023033"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2323213313100133-1323103032122030-1332111133333202-3201001303330223-1001302101322321-1231103202111230-1132020231222123-3223203032210132"></a>

### Direct properties for `sli_to_global_dr.global_vn`

<a id="canonical-0013333002021302-0020321313103112-3113122131221022-3311303011021231-3122023330221103-0033300112233121-3123001210122131-2302310100123303"></a>

#### `sli_to_global_dr.global_vn.name` property

Type: `"string"`. Computed.

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

<a id="canonical-1203311001321310-3130201301112022-3310301032030203-0011201322113013-2221103012322102-0233000122023131-3331013222203133-1103103103001100"></a>

<a id="canonical-3201300221132302-3222103112320033-2022202201213003-3022001133211213-3103032032320023-0322111323113211-2331000211320311-0103020310103001"></a>

#### `sli_to_global_dr.global_vn.namespace` property

Type: `"string"`. Computed.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2333223121322002-3223130311111131-1213123112033030-2101233011201132-1112322320010120-0000220000033120-0001001211023122-3231322230020001"></a>

<a id="canonical-1013310312122012-2323121001330123-0333132300033203-2102120011010213-1113310221323223-1031132021132210-3003211313302032-2230023232022221"></a>

#### `sli_to_global_dr.global_vn.tenant` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0121123320323123-3213322110002213-2100333022000003-2120210113312111-1331133121020232-0212331020331110-1031122203113102-0201102222130031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sli_to_slo_snat` properties

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-3100100033113131-1013100333112132-1220212103012023-0211023132210233-2323121020123322-2323213203113121-1203301130023023-3202032133312301)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-3000122203132213-2030301322323023-1302132111133302-0302223201130030-0303221223333310-2021321132132200-3100201000113230-1311223311010320)
- sli_to_slo_snat

<a id="canonical-2120003113110001-3210230121130331-2222110200221200-3321331001311301-2323022032112122-0213302221020112-0323032100010023-3133110321333213"></a>

Type: `"single"`. Computed.

Configuration parameter for sli to slo snat.

Additional upstream details:

X-example: "" description.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-pool_choice": "[\"interface_ip\"]",
  "x-ves-oneof-field-routing_choice": "[\"default_gw_snat\"]"
}
```

<a id="canonical-0232300332121010-2320120032301111-0233303122033223-2120311202130101-2230112211222122-3331203210012113-2330232313100221-0331231202023022"></a>

### Direct properties for `sli_to_slo_snat`

- [default_gw_snat](data-sources--network_connector--reference--group-001.md#canonical-1202301102133331-2101121130331302-2122022022313011-1321020211212223-2213110323123020-3230121211232230-3101020201230111-2323200111013111): complete subsection reference.

- [interface_ip](data-sources--network_connector--reference--group-001.md#canonical-1312333102103110-1010220011111103-2303223101122210-0323031121230032-2201021022220321-1223031313222003-2331021210121212-3123322001001003): complete subsection reference.

<a id="canonical-1202301102133331-2101121130331302-2122022022313011-1321020211212223-2213110323123020-3230121211232230-3101020201230111-2323200111013111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sli_to_slo_snat.default_gw_snat` properties

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-3100100033113131-1013100333112132-1220212103012023-0211023132210233-2323121020123322-2323213203113121-1203301130023023-3202032133312301)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-3000122203132213-2030301322323023-1302132111133302-0302223201130030-0303221223333310-2021321132132200-3100201000113230-1311223311010320)
- [sli_to_slo_snat](data-sources--network_connector--reference--group-001.md#canonical-0121123320323123-3213322110002213-2100333022000003-2120210113312111-1331133121020232-0212331020331110-1031122203113102-0201102222130031)
- sli_to_slo_snat.default_gw_snat

<a id="canonical-2310312133323330-3323001120321100-3232020203221132-2031212132000001-2103213113312200-1311222013213102-2020121330100310-1102332221213212"></a>

Type: `"single"`. Computed.

Configuration parameter for default gw snat.

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

<a id="canonical-1312333102103110-1010220011111103-2303223101122210-0323031121230032-2201021022220321-1223031313222003-2331021210121212-3123322001001003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sli_to_slo_snat.interface_ip` properties

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-3100100033113131-1013100333112132-1220212103012023-0211023132210233-2323121020123322-2323213203113121-1203301130023023-3202032133312301)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-3000122203132213-2030301322323023-1302132111133302-0302223201130030-0303221223333310-2021321132132200-3100201000113230-1311223311010320)
- [sli_to_slo_snat](data-sources--network_connector--reference--group-001.md#canonical-0121123320323123-3213322110002213-2100333022000003-2120210113312111-1331133121020232-0212331020331110-1031122203113102-0201102222130031)
- sli_to_slo_snat.interface_ip

<a id="canonical-2231312101122202-0022301123123033-3230100330001113-3133323022201213-3302021133220330-0133210302212012-3031213103013210-3230010333213012"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3202010033321000-2321111032303003-3110021002030311-2331130332311233-1122123100201221-1020301003032212-3023213002320313-1232123203201023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `slo_to_global_dr` properties

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-3100100033113131-1013100333112132-1220212103012023-0211023132210233-2323121020123322-2323213203113121-1203301130023023-3202032133312301)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-3000122203132213-2030301322323023-1302132111133302-0302223201130030-0303221223333310-2021321132132200-3100201000113230-1311223311010320)
- slo_to_global_dr

<a id="canonical-3113202033322122-3023213330323003-1102322112312010-3012211202021201-1200302021133313-3132010220021033-3220102232031032-2221230302300100"></a>

Type: `"single"`. Computed.

Global network reference for direct connection.

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

<a id="canonical-1301331220300011-3130233332210120-0200332000230213-0012330222133001-0022022203110302-3011123312132211-0331210332023321-2211021232223133"></a>

### Direct properties for `slo_to_global_dr`

- [global_vn](data-sources--network_connector--reference--group-001.md#canonical-3031000230020233-0320112012122210-0100113210120232-2103130120301030-1301200022022323-2210332302230212-1013313312330300-1322333210212302): complete subsection reference.

<a id="canonical-3031000230020233-0320112012122210-0100113210120232-2103130120301030-1301200022022323-2210332302230212-1013313312330300-1322333210212302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `slo_to_global_dr.global_vn` properties

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-3100100033113131-1013100333112132-1220212103012023-0211023132210233-2323121020123322-2323213203113121-1203301130023023-3202032133312301)
- [Property reference](data-sources--network_connector--reference--group-001.md#canonical-3000122203132213-2030301322323023-1302132111133302-0302223201130030-0303221223333310-2021321132132200-3100201000113230-1311223311010320)
- [slo_to_global_dr](data-sources--network_connector--reference--group-001.md#canonical-3202010033321000-2321111032303003-3110021002030311-2331130332311233-1122123100201221-1020301003032212-3023213002320313-1232123203201023)
- slo_to_global_dr.global_vn

<a id="canonical-0012113101332002-1221230321331030-3031322022220013-3200222211133023-3132203232312111-1100003300100122-3021230213003322-0130001230132222"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2213301131032203-0003011320311232-0330330120131002-1001303102211000-0023321311323131-1122033330020032-0131103030021313-0111111211221020"></a>

### Direct properties for `slo_to_global_dr.global_vn`

<a id="canonical-1021100223133002-0312200130110110-2133112320321301-1022321233030231-2121100213100203-0012203110101130-1120030211100300-2302031311111000"></a>

#### `slo_to_global_dr.global_vn.name` property

Type: `"string"`. Computed.

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

<a id="canonical-2112112321313303-2001133001303301-1131012330122320-3131002322331221-3323021233020221-0330033331111102-2233332311132122-0301233221101011"></a>

<a id="canonical-1112113003210201-2300112100121300-2210013320311223-2202330103201103-1131030231021103-2010032210020002-0020200112310030-3202230200311102"></a>

#### `slo_to_global_dr.global_vn.namespace` property

Type: `"string"`. Computed.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3020322222031203-1022033021113110-1322131033220113-1330101311023202-3100132222233032-0121111122201023-2200302320020102-3223121003321211"></a>

<a id="canonical-3032331001231010-0312223323113112-1223230130033330-1300033101203102-0022112130113131-2202203013132201-0210211003103020-1101231211130021"></a>

#### `slo_to_global_dr.global_vn.tenant` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```
