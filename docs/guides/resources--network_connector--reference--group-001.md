---
page_title: "xcsh_network_connector reference"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_network_connector reference."
---

# xcsh_network_connector reference

<a id="canonical-3230201320313300-3100103222213301-3222301220332000-2102123311113022-2103213133210102-1103110232333032-2123011010231323-2310223103003331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0030231112200120-2310101313112331-0112221320311330-1010030020110100-0320030301321031-0201122000110313-2302331223321203-3130322032000100)
- Property reference

<a id="canonical-0013212300313103-0213213313123222-1220113211321100-3300320012111021-2212212121131102-1210130110010011-0112001113100130-0111022111313003"></a>

### Direct properties for `xcsh_network_connector`

<a id="canonical-0110203231220223-0021112201100213-1122230013332112-1313323302330323-2021213301220323-2331000202331231-3001233322300232-1210121031302130"></a>

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

<a id="canonical-2110120100120321-1331202333002231-2023113113103130-0330322301333202-3122003000102210-3012022311333033-3323212123300223-1301032100303103"></a>

<a id="canonical-3130012332212003-2312300203120130-3022112113223231-0120033212323000-3310333301322202-0301101222323020-1230030233120203-1212302300121213"></a>

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

<a id="canonical-0033232221300000-2001231101231000-3032223011330300-2020031112212200-0000012302321111-2221230020220211-2110100313230111-3030131020131201"></a>

<a id="canonical-1201300311023231-1121322331001100-0301231300213133-1223122003133023-0303101331010022-1333230110301123-0303330300323011-1223013032230321"></a>

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

- [disable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-2111333330030000-1002031120031312-0303212231133031-0100102111232123-2322101011031230-0120112023330111-3230013222002110-2132322230321121): complete subsection reference.

- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-1101110030313313-2102212303300111-0320130013022202-3311211123122310-2200103232211301-3223222300011312-2202123000112213-2220311201321020): complete subsection reference.

<a id="canonical-3233310021210230-1112303210122213-0321101223313322-1111220213110011-3200300330331103-0333033030311133-3031213231000121-0112310303030002"></a>

<a id="canonical-2012310201023322-1001012233112230-0032233020132001-3231232001302233-3302110323312213-0333032330031223-1121131123312002-2202220012230003"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0320003013312013-2010211023230310-3233202301122033-2102221302111033-3130230310313001-3221002120221321-0231113112221031-1120110332230301"></a>

<a id="canonical-1122230232313232-3010332311120320-1032031303232331-1331323021103321-0011122003332113-2231201221223001-1103023220312212-0330301310023011"></a>

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

<a id="canonical-1012110312332110-0012200112303020-3122323230101101-3032002131130112-3021321200232133-3213110013320110-3303231200221223-2103213010301033"></a>

<a id="canonical-1311003223332033-2103020000212333-3013303132130222-2010112301130023-3020002130103330-2101012332002310-2122333300213233-2133302020333231"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Network Connector. Must be unique within the namespace.

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

<a id="canonical-2231303130230220-1020303123101123-3331303210031230-0003231030122021-2232233111133103-0203131031030003-0012011032230313-1221203232333220"></a>

<a id="canonical-2020020131200000-1031022222333132-2223121112001023-1213301320111320-0132212202233231-0100120201223100-3202112331331333-0202120012133110"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Network Connector is created.

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

- [sli_to_global_dr](resources--network_connector--reference--group-001.md#canonical-1323220001110322-1200211010233233-0233223320201012-0231301213210211-1002133133033023-0030103033202111-0202301000230323-0310330231012320): complete subsection reference.

- [sli_to_slo_snat](resources--network_connector--reference--group-001.md#canonical-0311230213310203-1002131223302201-1301011222023101-2212122300321021-2103313222121202-2010010213310323-3111100010201210-0122233201102021): complete subsection reference.

- [slo_to_global_dr](resources--network_connector--reference--group-001.md#canonical-3330311132311200-2332201021230031-1213333312000213-3123313301103312-3120201323231333-3303231003321122-2323202202011332-1322123201030111): complete subsection reference.

- [timeouts](resources--network_connector--reference--group-001.md#canonical-3221303012220021-3113332122010110-2233110102322321-1221102001133222-3210331030002022-1131133021000020-0202000213031021-3012033020001021): complete subsection reference.

<a id="canonical-1011120001111121-0011013320030023-1221002003330233-2303311233133020-1130011112021030-3230312212100131-0113220003331231-3220210311332023"></a>

### All schema paths for `xcsh_network_connector`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--network_connector--reference--group-001.md#canonical-0110203231220223-0021112201100213-1122230013332112-1313323302330323-2021213301220323-2331000202331231-3001233322300232-1210121031302130) |
| `description` | [description](resources--network_connector--reference--group-001.md#canonical-2110120100120321-1331202333002231-2023113113103130-0330322301333202-3122003000102210-3012022311333033-3323212123300223-1301032100303103) |
| `disable` | [disable](resources--network_connector--reference--group-001.md#canonical-0033232221300000-2001231101231000-3032223011330300-2020031112212200-0000012302321111-2221230020220211-2110100313230111-3030131020131201) |
| `disable_forward_proxy` | [disable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-2000230210231113-3000132000003230-2033030131301323-2211212303020222-0330301011103001-0323331132000022-1231220030300230-2230122120112101) |
| `enable_forward_proxy` | [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-2210030333123111-2200302103331221-3131002310330210-3132133222321121-2321021301103223-0020332033021132-1320100202211220-1330111113203132) |
| `enable_forward_proxy.connection_timeout` | [enable_forward_proxy.connection_timeout](resources--network_connector--reference--group-001.md#canonical-2200130333330330-2230203002100330-1232300301021331-3110323202020312-1122021030331223-1000222231213321-1203103130103223-0300110021120003) |
| `enable_forward_proxy.max_connect_attempts` | [enable_forward_proxy.max_connect_attempts](resources--network_connector--reference--group-001.md#canonical-0300122010102131-1032222221111210-3100123031301113-3003001100313233-2200222022311031-0123313322111132-3130311330021111-1010101021212033) |
| `enable_forward_proxy.no_interception` | [enable_forward_proxy.no_interception](resources--network_connector--reference--group-001.md#canonical-2000111123312222-0331321120201221-1121212112300012-2122110331120010-3332012313313103-0003023232101002-2210300223133212-0210222311033002) |
| `enable_forward_proxy.tls_intercept` | [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-0303331031333300-2230333010210221-0201333322133022-0320213231111020-2012300010320211-0110032320011333-3131310012201330-3302322201023222) |
| `enable_forward_proxy.tls_intercept.custom_certificate` | [enable_forward_proxy.tls_intercept.custom_certificate](resources--network_connector--reference--group-001.md#canonical-0323332303131301-0113201123331210-2211102311030000-3200013231033012-2202100322230113-0112320333332033-2010032320322003-0203113310000320) |
| `enable_forward_proxy.tls_intercept.custom_certificate.certificate_url` | [enable_forward_proxy.tls_intercept.custom_certificate.certificate_url](resources--network_connector--reference--group-001.md#canonical-0023132223121321-0102220203123222-2200230020211311-1332033322001121-3333023003330321-3323000132122103-3131111013013300-2130220221300102) |
| `enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms` | [enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms](resources--network_connector--reference--group-001.md#canonical-0212213333330132-2133122011211111-0133032312022313-2120002301311112-0333032010122112-0302300330133330-3112203221323223-3231001230233222) |
| `enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms.hash_algorithms` | [enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms.hash_algorithms](resources--network_connector--reference--group-001.md#canonical-3212102131032310-0110130002211132-3030232303033023-1222100030233310-3131330213221201-2122101332311112-1330020022330113-1330213133100013) |
| `enable_forward_proxy.tls_intercept.custom_certificate.description_spec` | [enable_forward_proxy.tls_intercept.custom_certificate.description_spec](resources--network_connector--reference--group-001.md#canonical-0021032021230222-0110333130321030-2202211003323201-0021032110233023-0130030123103013-0320200100303332-2233132202312223-3121002312332212) |
| `enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling` | [enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling](resources--network_connector--reference--group-001.md#canonical-0200303211001232-1331133233122310-1203120322330320-3032202220321110-0120232130113032-2112121203110031-1133031333231012-1023000120130211) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key](resources--network_connector--reference--group-001.md#canonical-3111302120022111-0001311022011332-1231302330331300-2202333003203310-1003022003130301-2333322322100223-3330221110031033-3133302002303231) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info](resources--network_connector--reference--group-001.md#canonical-0330320232110013-1202031222323001-3201000302211013-0333132200200000-1333312233321111-1033300011112311-2001003313011333-2331311210232201) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.decryption_provider` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.decryption_provider](resources--network_connector--reference--group-001.md#canonical-0112221222002223-3300033002311202-3210331023203200-1321300113133032-2020002130221331-3120021201022210-2003222000230101-1002033211120311) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.location` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.location](resources--network_connector--reference--group-001.md#canonical-3123130213130331-1013021232000013-3321222300003030-3202211132100302-2202022000210101-2200132312132001-1321212123321313-0030231012313311) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.store_provider` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.store_provider](resources--network_connector--reference--group-001.md#canonical-3003032213220323-1121133132112300-1301330203201101-2131330131021031-1000122311313230-3210021203300132-2213101203201030-3202101301030212) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info](resources--network_connector--reference--group-001.md#canonical-3221001220310013-2321331132320213-1231333213331023-2220100200302012-2330201001000122-1330102301013320-1233102130232201-3123332110002211) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info.provider_ref` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info.provider_ref](resources--network_connector--reference--group-001.md#canonical-2313302022203312-2113313232033203-0212323020232323-0022331211321223-0010320230132211-1101330102333231-1322200232131333-3210312101312331) |
| `enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info.url` | [enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info.url](resources--network_connector--reference--group-001.md#canonical-0111212311333333-0101310020221031-0303313123012200-1013131322010113-1003011120112220-3112203200212333-1203321133130100-3122323121002212) |
| `enable_forward_proxy.tls_intercept.custom_certificate.use_system_defaults` | [enable_forward_proxy.tls_intercept.custom_certificate.use_system_defaults](resources--network_connector--reference--group-001.md#canonical-2020013313310113-1313121120001122-3310123003323112-2111001032212133-1021103201221122-3322002122101332-2223311032201010-0023121313121301) |
| `enable_forward_proxy.tls_intercept.enable_for_all_domains` | [enable_forward_proxy.tls_intercept.enable_for_all_domains](resources--network_connector--reference--group-001.md#canonical-2113022331212100-0301123001120211-1233133001101103-0002110123210330-1201233301131213-3201030312100203-2231220111101012-3113231103203303) |
| `enable_forward_proxy.tls_intercept.policy` | [enable_forward_proxy.tls_intercept.policy](resources--network_connector--reference--group-001.md#canonical-0000023031010323-1033013132330100-1320202330223321-2222030322112011-3002313133332103-0313233333222210-2102023031223213-2203130211033012) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules` | [enable_forward_proxy.tls_intercept.policy.interception_rules](resources--network_connector--reference--group-001.md#canonical-1310122333123121-0201013321303213-2133211311032201-2113223323222220-0110333100201202-0303312230213122-1212100211321112-2031122133300122) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.disable_interception` | [enable_forward_proxy.tls_intercept.policy.interception_rules.disable_interception](resources--network_connector--reference--group-001.md#canonical-3133213013032213-0220110010322121-2103221212231233-1203131221232103-1021223213132312-1023022233130210-1123312333313330-2022020000312033) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match` | [enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match](resources--network_connector--reference--group-001.md#canonical-1310013133111323-0100212120232123-2011112211102221-1113101233113103-2030210102232023-2001301310332110-0033300002130203-1032023212103123) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.exact_value` | [enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.exact_value](resources--network_connector--reference--group-001.md#canonical-3020202321010200-0332133133022311-3000201210030200-1012322002220213-2110032212323133-1121131132013321-0212111231203301-3202130210223223) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.regex_value` | [enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.regex_value](resources--network_connector--reference--group-001.md#canonical-3213211321302230-1111102131223002-2301310011021112-2320203320320321-0111100101013132-1000121133330021-2123221311102303-3203313212012222) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.suffix_value` | [enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.suffix_value](resources--network_connector--reference--group-001.md#canonical-2101223011101010-3111110222221231-0000132111222110-1310123121221000-1001130333320023-3333022333013112-3133110002131233-0030333301033022) |
| `enable_forward_proxy.tls_intercept.policy.interception_rules.enable_interception` | [enable_forward_proxy.tls_intercept.policy.interception_rules.enable_interception](resources--network_connector--reference--group-001.md#canonical-0001300021011133-2102320133220210-2303101121101321-1130021231130301-3123310121031201-1230120121320311-0131032100013223-2312322311211330) |
| `enable_forward_proxy.tls_intercept.trusted_ca_url` | [enable_forward_proxy.tls_intercept.trusted_ca_url](resources--network_connector--reference--group-001.md#canonical-0131013330232000-1313122220213201-1201310030112102-0223022012233311-2220003210123120-3110333021232033-0221310221001110-3321302103103101) |
| `enable_forward_proxy.tls_intercept.volterra_certificate` | [enable_forward_proxy.tls_intercept.volterra_certificate](resources--network_connector--reference--group-001.md#canonical-1233300032300231-1221310230311322-2213233211101311-1201121232131323-1303120230301202-3123023322303311-0311130313211320-2320001130333311) |
| `enable_forward_proxy.tls_intercept.volterra_trusted_ca` | [enable_forward_proxy.tls_intercept.volterra_trusted_ca](resources--network_connector--reference--group-001.md#canonical-2212331102231220-0330303122320011-2320323223122112-1112312022112200-3233013320232133-1012313203222302-2222022011332230-1233301000120002) |
| `enable_forward_proxy.white_listed_ports` | [enable_forward_proxy.white_listed_ports](resources--network_connector--reference--group-001.md#canonical-0013133011210032-3102213021203113-0133033102122230-2023023330200102-2131020011202223-2332321201312011-2302020003113221-0312303102110113) |
| `enable_forward_proxy.white_listed_prefixes` | [enable_forward_proxy.white_listed_prefixes](resources--network_connector--reference--group-001.md#canonical-1232231032223033-3301110020222033-1123333203301103-2302001323130233-3011321011212201-2302203200231202-2111222011030332-2231222320230013) |
| `id` | [ID](resources--network_connector--reference--group-001.md#canonical-3233310021210230-1112303210122213-0321101223313322-1111220213110011-3200300330331103-0333033030311133-3031213231000121-0112310303030002) |
| `labels` | [labels](resources--network_connector--reference--group-001.md#canonical-0320003013312013-2010211023230310-3233202301122033-2102221302111033-3130230310313001-3221002120221321-0231113112221031-1120110332230301) |
| `name` | [name](resources--network_connector--reference--group-001.md#canonical-1012110312332110-0012200112303020-3122323230101101-3032002131130112-3021321200232133-3213110013320110-3303231200221223-2103213010301033) |
| `namespace` | [namespace](resources--network_connector--reference--group-001.md#canonical-2231303130230220-1020303123101123-3331303210031230-0003231030122021-2232233111133103-0203131031030003-0012011032230313-1221203232333220) |
| `sli_to_global_dr` | [sli_to_global_dr](resources--network_connector--reference--group-001.md#canonical-3131313111101222-3033030230012301-1222310033313330-0021231222210301-2102321321231230-3121210102111133-1113322013231123-3100200331112202) |
| `sli_to_global_dr.global_vn` | [sli_to_global_dr.global_vn](resources--network_connector--reference--group-001.md#canonical-2201233121011332-1332331331023223-2202011123122301-2112112320301001-0223233311121202-3322232130312331-1223021113201110-2013231011220232) |
| `sli_to_global_dr.global_vn.name` | [sli_to_global_dr.global_vn.name](resources--network_connector--reference--group-001.md#canonical-3331131021302130-0213121113333111-0013110230131032-3231200310032333-1202203211121011-2102202001130033-2033331202023030-2121110333220002) |
| `sli_to_global_dr.global_vn.namespace` | [sli_to_global_dr.global_vn.namespace](resources--network_connector--reference--group-001.md#canonical-1033113332311303-0301330013312031-0301223303323233-2003301113233012-3022302121220121-2110200200331011-1032222013231200-2313010001310000) |
| `sli_to_global_dr.global_vn.tenant` | [sli_to_global_dr.global_vn.tenant](resources--network_connector--reference--group-001.md#canonical-0102212121111101-0110013030101310-2213330032232000-1032120330100332-1201202303202120-0312122022222021-2301211022323323-1013230131132330) |
| `sli_to_slo_snat` | [sli_to_slo_snat](resources--network_connector--reference--group-001.md#canonical-2301302022122012-1000012020121231-0330232303123223-1213230322230002-1032300010103112-2131000211233032-2301211000213210-2330002013211232) |
| `sli_to_slo_snat.default_gw_snat` | [sli_to_slo_snat.default_gw_snat](resources--network_connector--reference--group-001.md#canonical-1103133032301313-3011333333032320-0202023010010133-2121013001201100-1123103132110233-0002132123202203-0231011313101332-0330232122331122) |
| `sli_to_slo_snat.interface_ip` | [sli_to_slo_snat.interface_ip](resources--network_connector--reference--group-001.md#canonical-0132220202122132-3330031121002010-0231221311003000-0122033021221131-2311020202302333-0011202020203030-2220120230220312-1132111001210032) |
| `slo_to_global_dr` | [slo_to_global_dr](resources--network_connector--reference--group-001.md#canonical-0322021212311130-2232200200011310-3031133210012000-2232031310233133-1232232110131100-3122323013032031-0122221132332013-1032011112221203) |
| `slo_to_global_dr.global_vn` | [slo_to_global_dr.global_vn](resources--network_connector--reference--group-001.md#canonical-2031130013113210-1011001010123112-0023021111023332-2231220011032212-0123003002320010-3310123312002230-1132330331321310-1313233332021030) |
| `slo_to_global_dr.global_vn.name` | [slo_to_global_dr.global_vn.name](resources--network_connector--reference--group-001.md#canonical-2122330030232031-0030223330033020-1330010323132023-0222013311121031-2023202233102211-2133022011122122-1303032300031130-1112222120123001) |
| `slo_to_global_dr.global_vn.namespace` | [slo_to_global_dr.global_vn.namespace](resources--network_connector--reference--group-001.md#canonical-2033210323023322-0132123332311330-0123233331231131-3033331131310133-0001202200212213-1021300101222030-2213311301031102-0101003310130013) |
| `slo_to_global_dr.global_vn.tenant` | [slo_to_global_dr.global_vn.tenant](resources--network_connector--reference--group-001.md#canonical-2300002321011312-0231322120220213-3010031003313202-1232033310101032-0113302210303323-1201230131220301-1233000023221122-1302122031211111) |
| `timeouts` | [timeouts](resources--network_connector--reference--group-001.md#canonical-1212011300100100-2331233233333101-0123323232132222-1102313101023011-3323230121000012-1312223001100131-1002132111203021-0233023221330222) |
| `timeouts.create` | [timeouts.create](resources--network_connector--reference--group-001.md#canonical-3022023301211101-2102012113220031-3201230232301021-3311202121000032-1113023131210201-3312332301213012-3022231203233001-3323022001220133) |
| `timeouts.delete` | [timeouts.delete](resources--network_connector--reference--group-001.md#canonical-1303031002212220-2313133130310230-1022311300130000-3231102123001300-0103132012110233-1102003132330111-1201123131321231-0313020301310101) |
| `timeouts.read` | [timeouts.read](resources--network_connector--reference--group-001.md#canonical-1120202311333313-0223211200330102-3000030332300012-1003123112012201-3032212122120001-0323200012332003-3023200221031222-1121111100122111) |
| `timeouts.update` | [timeouts.update](resources--network_connector--reference--group-001.md#canonical-3120232202000313-0012322203300320-3313000102010031-2101101113211100-2103133100211121-2121312203322120-2302303030223000-0111231032333232) |

<a id="canonical-2111333330030000-1002031120031312-0303212231133031-0100102111232123-2322101011031230-0120112023330111-3230013222002110-2132322230321121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `disable_forward_proxy` properties

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0030231112200120-2310101313112331-0112221320311330-1010030020110100-0320030301321031-0201122000110313-2302331223321203-3130322032000100)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-3230201320313300-3100103222213301-3222301220332000-2102123311113022-2103213133210102-1103110232333032-2123011010231323-2310223103003331)
- disable_forward_proxy

<a id="canonical-2000230210231113-3000132000003230-2033030131301323-2211212303020222-0330301011103001-0323331132000022-1231220030300230-2230122120112101"></a>

Type: `["object", {}]`. Optional.

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

- [disable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-2000230210231113-3000132000003230-2033030131301323-2211212303020222-0330301011103001-0323331132000022-1231220030300230-2230122120112101)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-2210030333123111-2200302103331221-3131002310330210-3132133222321121-2321021301103223-0020332033021132-1320100202211220-1330111113203132)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_forward_proxy = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1101110030313313-2102212303300111-0320130013022202-3311211123122310-2200103232211301-3223222300011312-2202123000112213-2220311201321020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy` properties

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0030231112200120-2310101313112331-0112221320311330-1010030020110100-0320030301321031-0201122000110313-2302331223321203-3130322032000100)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-3230201320313300-3100103222213301-3222301220332000-2102123311113022-2103213133210102-1103110232333032-2123011010231323-2310223103003331)
- enable_forward_proxy

<a id="canonical-2210030333123111-2200302103331221-3131002310330210-3132133222321121-2321021301103223-0020332033021132-1320100202211220-1330111113203132"></a>

Type: `"object"`. single nested block, Optional.

Fine tune forward proxy behavior

Few configurations allowed are

White listed ports and IP prefixes: Forward proxy does application protocol detection and server
name(SNI) detection by peeking into the traffic on the incoming downstream connection. Few protocols
doesn't have client sending the first data. In such cases, protocol and SNI detection fails. This
configuration allows, skipping protocol and SNI detection for whitelisted IP-prefix-list and ports
connection\_timeout: The timeout for new network connections to upstream server.
Max\_connect\_attempts: Maximum number of attempts made to make new network connection to upstream
server.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_interception",
    "tls_intercept")}
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
  "x-ves-oneof-field-tls_interception_choice": "[\"no_interception\",\"tls_intercept\"]"
}
```

Terraform syntax:

```terraform
enable_forward_proxy {
  # Configure direct properties listed below.
}
```

<a id="canonical-3202111302312232-0101323200133331-2023002303313003-0301332231210310-2303203022203211-3211120103303032-1203122000230002-1013322221200110"></a>

### Direct properties for `enable_forward_proxy`

<a id="canonical-2200130333330330-2230203002100330-1232300301021331-3110323202020312-1122021030331223-1000222231213321-1203103130103223-0300110021120003"></a>

#### `enable_forward_proxy.connection_timeout` property

Type: `"number"`. Optional.

The timeout for new network connections to upstream server. This is specified in milliseconds. The
(2 seconds). Defaults to \`2000\`.

Additional upstream details:

The default value is 2000 (2 seconds)

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtMost(600000),
}
```

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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-0300122010102131-1032222221111210-3100123031301113-3003001100313233-2200222022311031-0123313322111132-3130311330021111-1010101021212033"></a>

<a id="canonical-2202313222222201-2301231302213330-3231112300300103-0130203200121303-3112110113321003-0013123212031201-0302311113303120-3333001132232331"></a>

#### `enable_forward_proxy.max_connect_attempts` property

Type: `"number"`. Optional.

Specifies the allowed number of retries on connect failure to upstream server. Defaults to \`1\`.

Additional upstream details:

Defaults to 1.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtMost(8),
}
```

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
    "ves.io.schema.rules.uint32.lte": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "8"
  }
}
```

- [no_interception](resources--network_connector--reference--group-001.md#canonical-1220011120312221-0033121031212311-1130212312100101-3302021232121031-1231021000210332-2311003113002230-3002313213221102-2110133322032333): complete subsection reference.

- [tls_intercept](resources--network_connector--reference--group-001.md#canonical-1103300020233001-3233102132113202-0100231103312111-2133120322011110-3203012000032323-3033230123333013-2000221112111311-2302123032122113): complete subsection reference.

<a id="canonical-0013133011210032-3102213021203113-0133033102122230-2023023330200102-2131020011202223-2332321201312011-2302020003113221-0312303102110113"></a>

<a id="canonical-3330110030121211-0311022022120021-3222121100131000-2201112110213301-2321202302031011-1030001021230303-0223201032023022-1131023032030310"></a>

#### `enable_forward_proxy.white_listed_ports` property

Type: `["list", "number"]`. Optional.

Traffic to these destination TCP ports is not subjected to protocol parsing Example 'tmate' server
port.

Additional upstream details:

Traffic to these destination TCP ports is not subjected to protocol parsing Example "tmate" server
port.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
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
    "ves.io.schema.rules.repeated.items.uint32.lte": "65535",
    "ves.io.schema.rules.repeated.max_items": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.uint32.lte": "65535",
    "ves.io.schema.rules.repeated.max_items": "64"
  }
}
```

<a id="canonical-1232231032223033-3301110020222033-1123333203301103-2302001323130233-3011321011212201-2302203200231202-2111222011030332-2231222320230013"></a>

<a id="canonical-1022032013033200-1022221002033201-3311212232202322-1123102102322210-2023122310221122-2233123211013112-0122121332023233-3311133110100021"></a>

#### `enable_forward_proxy.white_listed_prefixes` property

Type: `["list", "string"]`. Optional.

Traffic to these destination IP prefixes is not subjected to protocol parsing Example 'tmate' server
IP.

Additional upstream details:

Traffic to these destination IP prefixes is not subjected to protocol parsing Example "tmate" server
IP.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
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

<a id="canonical-1220011120312221-0033121031212311-1130212312100101-3302021232121031-1231021000210332-2311003113002230-3002313213221102-2110133322032333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy.no_interception` properties

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0030231112200120-2310101313112331-0112221320311330-1010030020110100-0320030301321031-0201122000110313-2302331223321203-3130322032000100)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-3230201320313300-3100103222213301-3222301220332000-2102123311113022-2103213133210102-1103110232333032-2123011010231323-2310223103003331)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-1101110030313313-2102212303300111-0320130013022202-3311211123122310-2200103232211301-3223222300011312-2202123000112213-2220311201321020)
- enable_forward_proxy.no_interception

<a id="canonical-2000111123312222-0331321120201221-1121212112300012-2122110331120010-3332012313313103-0003023232101002-2210300223133212-0210222311033002"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_interception = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1103300020233001-3233102132113202-0100231103312111-2133120322011110-3203012000032323-3033230123333013-2000221112111311-2302123032122113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy.tls_intercept` properties

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0030231112200120-2310101313112331-0112221320311330-1010030020110100-0320030301321031-0201122000110313-2302331223321203-3130322032000100)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-3230201320313300-3100103222213301-3222301220332000-2102123311113022-2103213133210102-1103110232333032-2123011010231323-2310223103003331)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-1101110030313313-2102212303300111-0320130013022202-3311211123122310-2200103232211301-3223222300011312-2202123000112213-2220311201321020)
- enable_forward_proxy.tls_intercept

<a id="canonical-0303331031333300-2230333010210221-0201333322133022-0320213231111020-2012300010320211-0110032320011333-3131310012201330-3302322201023222"></a>

Type: `"object"`. single nested block, Optional.

Configuration to enable TLS interception.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_certificate",
    "volterra_certificate"),
  validators.ConflictingObjectAttributes("enable_for_all_domains",
    "policy"),
  validators.ConflictingObjectAttributes("trusted_ca_url",
    "volterra_trusted_ca")}
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
  "x-ves-oneof-field-interception_policy_choice": "[\"enable_for_all_domains\",\"policy\"]",
  "x-ves-oneof-field-signing_cert_choice": "[\"custom_certificate\",\"volterra_certificate\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca_url\",\"volterra_trusted_ca\"]"
}
```

Terraform syntax:

```terraform
tls_intercept {
  # Configure direct properties listed below.
}
```

<a id="canonical-0223222120313131-1133003221213033-1310102010211102-0112321023101013-3231010113321111-2002002131203200-0323212100322032-3222231301131032"></a>

### Direct properties for `enable_forward_proxy.tls_intercept`

- [custom_certificate](resources--network_connector--reference--group-001.md#canonical-3001233111130311-2332001231310032-0012121313112300-2201122031203313-0033000103222321-0201010200101102-1232030300133322-3030101321223233): complete subsection reference.

- [enable_for_all_domains](resources--network_connector--reference--group-001.md#canonical-0310220322102112-3023130112222311-0212011213102203-1123003212100212-3112023222123021-2100130210320200-3131321232102211-3300011210112203): complete subsection reference.

- [policy](resources--network_connector--reference--group-001.md#canonical-3211212111320201-2222122020320213-1320133223221323-0001311130111112-2100023002132002-0313322211212211-2210010112011212-3000210003122113): complete subsection reference.

<a id="canonical-0131013330232000-1313122220213201-1201310030112102-0223022012233311-2220003210123120-3110333021232033-0221310221001110-3321302103103101"></a>

<a id="canonical-3223132012133001-2330323023212312-2020202333330300-1310013223203011-0221110331001220-1011110012102333-3310131223100012-1330301003120022"></a>

#### `enable_forward_proxy.tls_intercept.trusted_ca_url` property

Type: `"string"`. Optional.

Exclusive with \[volterra\_trusted\_ca\] Custom Root CA Certificate for validating upstream server
certificate.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(131072),
}
```

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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [volterra_certificate](resources--network_connector--reference--group-001.md#canonical-0130030220200330-1333222021330210-1133300012112200-0323222022203223-2213302000230231-2102101222222303-1311030030322213-2021112003332203): complete subsection reference.

- [volterra_trusted_ca](resources--network_connector--reference--group-001.md#canonical-2321122203131312-2122202320313213-1022101332303201-3100332222312032-2011110010303333-3303033030030002-3113230312013221-1310020002020330): complete subsection reference.

<a id="canonical-3001233111130311-2332001231310032-0012121313112300-2201122031203313-0033000103222321-0201010200101102-1232030300133322-3030101321223233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy.tls_intercept.custom_certificate` properties

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0030231112200120-2310101313112331-0112221320311330-1010030020110100-0320030301321031-0201122000110313-2302331223321203-3130322032000100)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-3230201320313300-3100103222213301-3222301220332000-2102123311113022-2103213133210102-1103110232333032-2123011010231323-2310223103003331)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-1101110030313313-2102212303300111-0320130013022202-3311211123122310-2200103232211301-3223222300011312-2202123000112213-2220311201321020)
- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-1103300020233001-3233102132113202-0100231103312111-2133120322011110-3203012000032323-3033230123333013-2000221112111311-2302123032122113)
- enable_forward_proxy.tls_intercept.custom_certificate

<a id="canonical-0323332303131301-0113201123331210-2211102311030000-3200013231033012-2202100322230113-0112320333332033-2010032320322003-0203113310000320"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for custom certificate.

Additional upstream details:

Handle to fetch certificate and key.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("certificate_url"),
  validators.ConflictingObjectAttributes("custom_hash_algorithms",
    "disable_ocsp_stapling"),
  validators.ConflictingObjectAttributes("custom_hash_algorithms",
    "use_system_defaults"),
  validators.ConflictingObjectAttributes("disable_ocsp_stapling",
    "use_system_defaults")}
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
  "x-ves-oneof-field-ocsp_stapling_choice": "[\"custom_hash_algorithms\",\"disable_ocsp_stapling\",\"use_system_defaults\"]"
}
```

Terraform syntax:

```terraform
custom_certificate {
  # Configure direct properties listed below.
}
```

<a id="canonical-2303313330100202-2332113332013220-2300310022232330-2123222120203220-3220300210322232-1012213313023111-1310323202311231-3302020311022010"></a>

### Direct properties for `enable_forward_proxy.tls_intercept.custom_certificate`

<a id="canonical-0023132223121321-0102220203123222-2200230020211311-1332033322001121-3333023003330321-3323000132122103-3131111013013300-2130220221300102"></a>

#### `enable_forward_proxy.tls_intercept.custom_certificate.certificate_url` property

Type: `"string"`. Optional.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

- [custom_hash_algorithms](resources--network_connector--reference--group-001.md#canonical-1222220221121321-2112110330102031-2011212032031302-3300212212120110-1123322011321313-2300011030023332-0132130133221002-0300232100200203): complete subsection reference.

<a id="canonical-0021032021230222-0110333130321030-2202211003323201-0021032110233023-0130030123103013-0320200100303332-2233132202312223-3121002312332212"></a>

<a id="canonical-1202102312132302-2023210131023233-2112322033200022-0211333021211012-1333323212322322-1231133200111210-0110121312030302-1000321122332232"></a>

#### `enable_forward_proxy.tls_intercept.custom_certificate.description_spec` property

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--network_connector--reference--group-001.md#canonical-3021203321310222-3123232223313100-2312011031021130-0113310311212011-1031322113213221-1011320213122300-0012121210212103-1210003130312211): complete subsection reference.

- [private_key](resources--network_connector--reference--group-001.md#canonical-2221211102221322-0121012130323032-2120331032320113-1022031001212310-3230231103002311-3202302003110303-0331101203120033-0203030111001221): complete subsection reference.

- [use_system_defaults](resources--network_connector--reference--group-001.md#canonical-2302323012133000-3023322330122001-2111022313231223-0013221120232101-0322010231231012-3231011213212001-1033002001331323-1201200222112102): complete subsection reference.

<a id="canonical-1222220221121321-2112110330102031-2011212032031302-3300212212120110-1123322011321313-2300011030023332-0132130133221002-0300232100200203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0030231112200120-2310101313112331-0112221320311330-1010030020110100-0320030301321031-0201122000110313-2302331223321203-3130322032000100)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-3230201320313300-3100103222213301-3222301220332000-2102123311113022-2103213133210102-1103110232333032-2123011010231323-2310223103003331)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-1101110030313313-2102212303300111-0320130013022202-3311211123122310-2200103232211301-3223222300011312-2202123000112213-2220311201321020)
- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-1103300020233001-3233102132113202-0100231103312111-2133120322011110-3203012000032323-3033230123333013-2000221112111311-2302123032122113)
- [enable_forward_proxy.tls_intercept.custom_certificate](resources--network_connector--reference--group-001.md#canonical-3001233111130311-2332001231310032-0012121313112300-2201122031203313-0033000103222321-0201010200101102-1232030300133322-3030101321223233)
- enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms

<a id="canonical-0212213333330132-2133122011211111-0133032312022313-2120002301311112-0333032010122112-0302300330133330-3112203221323223-3231001230233222"></a>

Type: `"object"`. single nested block, Optional.

Specifies the hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("hash_algorithms")}
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
custom_hash_algorithms {
  # Configure direct properties listed below.
}
```

<a id="canonical-0020033201022322-3330321112231103-0012102031113210-3103332321323220-1113303301000213-1333200202331132-3102220113300003-3003121033101123"></a>

### Direct properties for `enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms`

<a id="canonical-3212102131032310-0110130002211132-3030232303033023-1222100030233310-3131330213221201-2122101332311112-1330020022330113-1330213133100013"></a>

#### `enable_forward_proxy.tls_intercept.custom_certificate.custom_hash_algorithms.hash_algorithms` property

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3021203321310222-3123232223313100-2312011031021130-0113310311212011-1031322113213221-1011320213122300-0012121210212103-1210003130312211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0030231112200120-2310101313112331-0112221320311330-1010030020110100-0320030301321031-0201122000110313-2302331223321203-3130322032000100)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-3230201320313300-3100103222213301-3222301220332000-2102123311113022-2103213133210102-1103110232333032-2123011010231323-2310223103003331)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-1101110030313313-2102212303300111-0320130013022202-3311211123122310-2200103232211301-3223222300011312-2202123000112213-2220311201321020)
- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-1103300020233001-3233102132113202-0100231103312111-2133120322011110-3203012000032323-3033230123333013-2000221112111311-2302123032122113)
- [enable_forward_proxy.tls_intercept.custom_certificate](resources--network_connector--reference--group-001.md#canonical-3001233111130311-2332001231310032-0012121313112300-2201122031203313-0033000103222321-0201010200101102-1232030300133322-3030101321223233)
- enable_forward_proxy.tls_intercept.custom_certificate.disable_ocsp_stapling

<a id="canonical-0200303211001232-1331133233122310-1203120322330320-3032202220321110-0120232130113032-2112121203110031-1133031333231012-1023000120130211"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_ocsp_stapling = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2221211102221322-0121012130323032-2120331032320113-1022031001212310-3230231103002311-3202302003110303-0331101203120033-0203030111001221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy.tls_intercept.custom_certificate.private_key` properties

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0030231112200120-2310101313112331-0112221320311330-1010030020110100-0320030301321031-0201122000110313-2302331223321203-3130322032000100)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-3230201320313300-3100103222213301-3222301220332000-2102123311113022-2103213133210102-1103110232333032-2123011010231323-2310223103003331)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-1101110030313313-2102212303300111-0320130013022202-3311211123122310-2200103232211301-3223222300011312-2202123000112213-2220311201321020)
- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-1103300020233001-3233102132113202-0100231103312111-2133120322011110-3203012000032323-3033230123333013-2000221112111311-2302123032122113)
- [enable_forward_proxy.tls_intercept.custom_certificate](resources--network_connector--reference--group-001.md#canonical-3001233111130311-2332001231310032-0012121313112300-2201122031203313-0033000103222321-0201010200101102-1232030300133322-3030101321223233)
- enable_forward_proxy.tls_intercept.custom_certificate.private_key

<a id="canonical-3111302120022111-0001311022011332-1231302330331300-2202333003203310-1003022003130301-2333322322100223-3330221110031033-3133302002303231"></a>

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
private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-1332113233130333-1021232200312133-2201001313213130-1232133002322333-2313201121010010-2003031010331133-1323330031000120-0312110010203112"></a>

### Direct properties for `enable_forward_proxy.tls_intercept.custom_certificate.private_key`

- [blindfold_secret_info](resources--network_connector--reference--group-001.md#canonical-0301211101221003-2120223030311111-3023111113023002-1213000012321121-3102013112320113-2112103203102303-3333111333020110-1220300300101323): complete subsection reference.

- [clear_secret_info](resources--network_connector--reference--group-001.md#canonical-0032033013102232-3303100213113021-2233201201133120-2101321230020221-3000032021020022-1111303102333230-3311212113232320-1313231203200032): complete subsection reference.

<a id="canonical-0301211101221003-2120223030311111-3023111113023002-1213000012321121-3102013112320113-2112103203102303-3333111333020110-1220300300101323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0030231112200120-2310101313112331-0112221320311330-1010030020110100-0320030301321031-0201122000110313-2302331223321203-3130322032000100)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-3230201320313300-3100103222213301-3222301220332000-2102123311113022-2103213133210102-1103110232333032-2123011010231323-2310223103003331)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-1101110030313313-2102212303300111-0320130013022202-3311211123122310-2200103232211301-3223222300011312-2202123000112213-2220311201321020)
- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-1103300020233001-3233102132113202-0100231103312111-2133120322011110-3203012000032323-3033230123333013-2000221112111311-2302123032122113)
- [enable_forward_proxy.tls_intercept.custom_certificate](resources--network_connector--reference--group-001.md#canonical-3001233111130311-2332001231310032-0012121313112300-2201122031203313-0033000103222321-0201010200101102-1232030300133322-3030101321223233)
- [enable_forward_proxy.tls_intercept.custom_certificate.private_key](resources--network_connector--reference--group-001.md#canonical-2221211102221322-0121012130323032-2120331032320113-1022031001212310-3230231103002311-3202302003110303-0331101203120033-0203030111001221)
- enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info

<a id="canonical-0330320232110013-1202031222323001-3201000302211013-0333132200200000-1333312233321111-1033300011112311-2001003313011333-2331311210232201"></a>

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

<a id="canonical-1121223210131200-2020221210212233-2032200323033001-0031101323332223-1120332233131010-2321032123222112-3122111123301313-3023123022110020"></a>

### Direct properties for `enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info`

<a id="canonical-0112221222002223-3300033002311202-3210331023203200-1321300113133032-2020002130221331-3120021201022210-2003222000230101-1002033211120311"></a>

#### `enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-3123130213130331-1013021232000013-3321222300003030-3202211132100302-2202022000210101-2200132312132001-1321212123321313-0030231012313311"></a>

<a id="canonical-3031023021212201-2103020030223303-3220001201301011-1311030130100233-2110032002331003-1133301323131302-1332230320230010-2210112000233200"></a>

#### `enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3003032213220323-1121133132112300-1301330203201101-2131330131021031-1000122311313230-3210021203300132-2213101203201030-3202101301030212"></a>

<a id="canonical-2223001031320210-0203230100213212-1203313003200232-1012113320321311-0312321311331232-0003301321200120-3122213321211213-3210200312021121"></a>

#### `enable_forward_proxy.tls_intercept.custom_certificate.private_key.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

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

<a id="canonical-0032033013102232-3303100213113021-2233201201133120-2101321230020221-3000032021020022-1111303102333230-3311212113232320-1313231203200032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0030231112200120-2310101313112331-0112221320311330-1010030020110100-0320030301321031-0201122000110313-2302331223321203-3130322032000100)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-3230201320313300-3100103222213301-3222301220332000-2102123311113022-2103213133210102-1103110232333032-2123011010231323-2310223103003331)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-1101110030313313-2102212303300111-0320130013022202-3311211123122310-2200103232211301-3223222300011312-2202123000112213-2220311201321020)
- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-1103300020233001-3233102132113202-0100231103312111-2133120322011110-3203012000032323-3033230123333013-2000221112111311-2302123032122113)
- [enable_forward_proxy.tls_intercept.custom_certificate](resources--network_connector--reference--group-001.md#canonical-3001233111130311-2332001231310032-0012121313112300-2201122031203313-0033000103222321-0201010200101102-1232030300133322-3030101321223233)
- [enable_forward_proxy.tls_intercept.custom_certificate.private_key](resources--network_connector--reference--group-001.md#canonical-2221211102221322-0121012130323032-2120331032320113-1022031001212310-3230231103002311-3202302003110303-0331101203120033-0203030111001221)
- enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info

<a id="canonical-3221001220310013-2321331132320213-1231333213331023-2220100200302012-2330201001000122-1330102301013320-1233102130232201-3123332110002211"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2130221022201130-1102123112012233-0113222020331133-3202203033113233-3210103233113120-3321021033122333-0321313320220020-2210230223013232"></a>

### Direct properties for `enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info`

<a id="canonical-2313302022203312-2113313232033203-0212323020232323-0022331211321223-0010320230132211-1101330102333231-1322200232131333-3210312101312331"></a>

#### `enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0111212311333333-0101310020221031-0303313123012200-1013131322010113-1003011120112220-3112203200212333-1203321133130100-3122323121002212"></a>

<a id="canonical-1212021302030011-1122211122221311-1203221213223221-2220010310210330-0232223333120002-3221231031213212-2021031001033232-2333223331200323"></a>

#### `enable_forward_proxy.tls_intercept.custom_certificate.private_key.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2302323012133000-3023322330122001-2111022313231223-0013221120232101-0322010231231012-3231011213212001-1033002001331323-1201200222112102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy.tls_intercept.custom_certificate.use_system_defaults` properties

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0030231112200120-2310101313112331-0112221320311330-1010030020110100-0320030301321031-0201122000110313-2302331223321203-3130322032000100)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-3230201320313300-3100103222213301-3222301220332000-2102123311113022-2103213133210102-1103110232333032-2123011010231323-2310223103003331)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-1101110030313313-2102212303300111-0320130013022202-3311211123122310-2200103232211301-3223222300011312-2202123000112213-2220311201321020)
- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-1103300020233001-3233102132113202-0100231103312111-2133120322011110-3203012000032323-3033230123333013-2000221112111311-2302123032122113)
- [enable_forward_proxy.tls_intercept.custom_certificate](resources--network_connector--reference--group-001.md#canonical-3001233111130311-2332001231310032-0012121313112300-2201122031203313-0033000103222321-0201010200101102-1232030300133322-3030101321223233)
- enable_forward_proxy.tls_intercept.custom_certificate.use_system_defaults

<a id="canonical-2020013313310113-1313121120001122-3310123003323112-2111001032212133-1021103201221122-3322002122101332-2223311032201010-0023121313121301"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
use_system_defaults = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0310220322102112-3023130112222311-0212011213102203-1123003212100212-3112023222123021-2100130210320200-3131321232102211-3300011210112203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy.tls_intercept.enable_for_all_domains` properties

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0030231112200120-2310101313112331-0112221320311330-1010030020110100-0320030301321031-0201122000110313-2302331223321203-3130322032000100)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-3230201320313300-3100103222213301-3222301220332000-2102123311113022-2103213133210102-1103110232333032-2123011010231323-2310223103003331)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-1101110030313313-2102212303300111-0320130013022202-3311211123122310-2200103232211301-3223222300011312-2202123000112213-2220311201321020)
- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-1103300020233001-3233102132113202-0100231103312111-2133120322011110-3203012000032323-3033230123333013-2000221112111311-2302123032122113)
- enable_forward_proxy.tls_intercept.enable_for_all_domains

<a id="canonical-2113022331212100-0301123001120211-1233133001101103-0002110123210330-1201233301131213-3201030312100203-2231220111101012-3113231103203303"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
enable_for_all_domains = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3211212111320201-2222122020320213-1320133223221323-0001311130111112-2100023002132002-0313322211212211-2210010112011212-3000210003122113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy.tls_intercept.policy` properties

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0030231112200120-2310101313112331-0112221320311330-1010030020110100-0320030301321031-0201122000110313-2302331223321203-3130322032000100)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-3230201320313300-3100103222213301-3222301220332000-2102123311113022-2103213133210102-1103110232333032-2123011010231323-2310223103003331)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-1101110030313313-2102212303300111-0320130013022202-3311211123122310-2200103232211301-3223222300011312-2202123000112213-2220311201321020)
- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-1103300020233001-3233102132113202-0100231103312111-2133120322011110-3203012000032323-3033230123333013-2000221112111311-2302123032122113)
- enable_forward_proxy.tls_intercept.policy

<a id="canonical-0000023031010323-1033013132330100-1320202330223321-2222030322112011-3002313133332103-0313233333222210-2102023031223213-2203130211033012"></a>

Type: `"object"`. single nested block, Optional.

Policy to enable or disable TLS interception.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("interception_rules")}
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
policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-2001000120010100-1312103332320001-3301331221133000-2332011031221110-2233030011000313-2133133013013233-3012121131333222-1332332213022112"></a>

### Direct properties for `enable_forward_proxy.tls_intercept.policy`

- [interception_rules](resources--network_connector--reference--group-001.md#canonical-2212132222031223-2022113131220322-3333032311223130-1203021131321033-1301233321012022-0112103110300102-3133032102230313-3321030102212031): complete subsection reference.

<a id="canonical-2212132222031223-2022113131220322-3333032311223130-1203021131321033-1301233321012022-0112103110300102-3133032102230313-3321030102212031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy.tls_intercept.policy.interception_rules` properties

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0030231112200120-2310101313112331-0112221320311330-1010030020110100-0320030301321031-0201122000110313-2302331223321203-3130322032000100)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-3230201320313300-3100103222213301-3222301220332000-2102123311113022-2103213133210102-1103110232333032-2123011010231323-2310223103003331)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-1101110030313313-2102212303300111-0320130013022202-3311211123122310-2200103232211301-3223222300011312-2202123000112213-2220311201321020)
- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-1103300020233001-3233102132113202-0100231103312111-2133120322011110-3203012000032323-3033230123333013-2000221112111311-2302123032122113)
- [enable_forward_proxy.tls_intercept.policy](resources--network_connector--reference--group-001.md#canonical-3211212111320201-2222122020320213-1320133223221323-0001311130111112-2100023002132002-0313322211212211-2210010112011212-3000210003122113)
- enable_forward_proxy.tls_intercept.policy.interception_rules

<a id="canonical-1310122333123121-0201013321303213-2133211311032201-2113223323222220-0110333100201202-0303312230213122-1212100211321112-2031122133300122"></a>

Type: `"object"`. list nested block, Optional.

List of ordered rules to enable or disable for TLS interception.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("disable_interception",
    "enable_interception")}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

Terraform syntax:

```terraform
interception_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-3322222212020100-0311303212023313-3303333011212332-3332222211021300-1011301032201122-1112211332020113-1110033330223010-0322112131011220"></a>

### Direct properties for `enable_forward_proxy.tls_intercept.policy.interception_rules`

- [disable_interception](resources--network_connector--reference--group-001.md#canonical-3302100112031320-3032312131031033-0321102200002110-3031331311112110-1100020213030010-2210233003210202-1000300032111103-0230122331133103): complete subsection reference.

- [domain_match](resources--network_connector--reference--group-001.md#canonical-0310200011100112-3220212010300122-2323023020302311-0223333330110122-1302230013001222-0230322030120110-3332001030223301-2230211211230221): complete subsection reference.

- [enable_interception](resources--network_connector--reference--group-001.md#canonical-2233321120031103-2221330033023220-3303231331113011-2032300030102113-2320002001121322-1320031132001330-3122221121133001-2221002123022000): complete subsection reference.

<a id="canonical-3302100112031320-3032312131031033-0321102200002110-3031331311112110-1100020213030010-2210233003210202-1000300032111103-0230122331133103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy.tls_intercept.policy.interception_rules.disable_interception` properties

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0030231112200120-2310101313112331-0112221320311330-1010030020110100-0320030301321031-0201122000110313-2302331223321203-3130322032000100)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-3230201320313300-3100103222213301-3222301220332000-2102123311113022-2103213133210102-1103110232333032-2123011010231323-2310223103003331)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-1101110030313313-2102212303300111-0320130013022202-3311211123122310-2200103232211301-3223222300011312-2202123000112213-2220311201321020)
- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-1103300020233001-3233102132113202-0100231103312111-2133120322011110-3203012000032323-3033230123333013-2000221112111311-2302123032122113)
- [enable_forward_proxy.tls_intercept.policy](resources--network_connector--reference--group-001.md#canonical-3211212111320201-2222122020320213-1320133223221323-0001311130111112-2100023002132002-0313322211212211-2210010112011212-3000210003122113)
- [enable_forward_proxy.tls_intercept.policy.interception_rules](resources--network_connector--reference--group-001.md#canonical-2212132222031223-2022113131220322-3333032311223130-1203021131321033-1301233321012022-0112103110300102-3133032102230313-3321030102212031)
- enable_forward_proxy.tls_intercept.policy.interception_rules.disable_interception

<a id="canonical-3133213013032213-0220110010322121-2103221212231233-1203131221232103-1021223213132312-1023022233130210-1123312333313330-2022020000312033"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_interception = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0310200011100112-3220212010300122-2323023020302311-0223333330110122-1302230013001222-0230322030120110-3332001030223301-2230211211230221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match` properties

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0030231112200120-2310101313112331-0112221320311330-1010030020110100-0320030301321031-0201122000110313-2302331223321203-3130322032000100)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-3230201320313300-3100103222213301-3222301220332000-2102123311113022-2103213133210102-1103110232333032-2123011010231323-2310223103003331)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-1101110030313313-2102212303300111-0320130013022202-3311211123122310-2200103232211301-3223222300011312-2202123000112213-2220311201321020)
- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-1103300020233001-3233102132113202-0100231103312111-2133120322011110-3203012000032323-3033230123333013-2000221112111311-2302123032122113)
- [enable_forward_proxy.tls_intercept.policy](resources--network_connector--reference--group-001.md#canonical-3211212111320201-2222122020320213-1320133223221323-0001311130111112-2100023002132002-0313322211212211-2210010112011212-3000210003122113)
- [enable_forward_proxy.tls_intercept.policy.interception_rules](resources--network_connector--reference--group-001.md#canonical-2212132222031223-2022113131220322-3333032311223130-1203021131321033-1301233321012022-0112103110300102-3133032102230313-3321030102212031)
- enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match

<a id="canonical-1310013133111323-0100212120232123-2011112211102221-1113101233113103-2030210102232023-2001301310332110-0033300002130203-1032023212103123"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for domain match.

Additional upstream details:

Domains names.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
domain_match {
  # Configure direct properties listed below.
}
```

<a id="canonical-0123122022222322-0323110223213103-3003012010313313-1123321200000030-3323331312220121-1101300001020213-3202031331122022-2003203331010223"></a>

### Direct properties for `enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match`

<a id="canonical-3020202321010200-0332133133022311-3000201210030200-1012322002220213-2110032212323133-1121131132013321-0212111231203301-3202130210223223"></a>

#### `enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.exact_value` property

Type: `"string"`. Optional.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

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
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3213211321302230-1111102131223002-2301310011021112-2320203320320321-0111100101013132-1000121133330021-2123221311102303-3203313212012222"></a>

<a id="canonical-3102233111021111-0113332331123113-2300122333233223-3010011120311011-1020101103331121-2121220210132033-3113202100003031-2022031210311201"></a>

#### `enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.regex_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2101223011101010-3111110222221231-0000132111222110-1310123121221000-1001130333320023-3333022333013112-3133110002131233-0030333301033022"></a>

<a id="canonical-1001121313020221-0330113010022213-1330312300131003-0122100200320320-0213330330003201-2333131330223021-2322032111000310-1200032210313121"></a>

#### `enable_forward_proxy.tls_intercept.policy.interception_rules.domain_match.suffix_value` property

Type: `"string"`. Optional.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2233321120031103-2221330033023220-3303231331113011-2032300030102113-2320002001121322-1320031132001330-3122221121133001-2221002123022000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy.tls_intercept.policy.interception_rules.enable_interception` properties

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0030231112200120-2310101313112331-0112221320311330-1010030020110100-0320030301321031-0201122000110313-2302331223321203-3130322032000100)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-3230201320313300-3100103222213301-3222301220332000-2102123311113022-2103213133210102-1103110232333032-2123011010231323-2310223103003331)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-1101110030313313-2102212303300111-0320130013022202-3311211123122310-2200103232211301-3223222300011312-2202123000112213-2220311201321020)
- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-1103300020233001-3233102132113202-0100231103312111-2133120322011110-3203012000032323-3033230123333013-2000221112111311-2302123032122113)
- [enable_forward_proxy.tls_intercept.policy](resources--network_connector--reference--group-001.md#canonical-3211212111320201-2222122020320213-1320133223221323-0001311130111112-2100023002132002-0313322211212211-2210010112011212-3000210003122113)
- [enable_forward_proxy.tls_intercept.policy.interception_rules](resources--network_connector--reference--group-001.md#canonical-2212132222031223-2022113131220322-3333032311223130-1203021131321033-1301233321012022-0112103110300102-3133032102230313-3321030102212031)
- enable_forward_proxy.tls_intercept.policy.interception_rules.enable_interception

<a id="canonical-0001300021011133-2102320133220210-2303101121101321-1130021231130301-3123310121031201-1230120121320311-0131032100013223-2312322311211330"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
enable_interception = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0130030220200330-1333222021330210-1133300012112200-0323222022203223-2213302000230231-2102101222222303-1311030030322213-2021112003332203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy.tls_intercept.volterra_certificate` properties

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0030231112200120-2310101313112331-0112221320311330-1010030020110100-0320030301321031-0201122000110313-2302331223321203-3130322032000100)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-3230201320313300-3100103222213301-3222301220332000-2102123311113022-2103213133210102-1103110232333032-2123011010231323-2310223103003331)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-1101110030313313-2102212303300111-0320130013022202-3311211123122310-2200103232211301-3223222300011312-2202123000112213-2220311201321020)
- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-1103300020233001-3233102132113202-0100231103312111-2133120322011110-3203012000032323-3033230123333013-2000221112111311-2302123032122113)
- enable_forward_proxy.tls_intercept.volterra_certificate

<a id="canonical-1233300032300231-1221310230311322-2213233211101311-1201121232131323-1303120230301202-3123023322303311-0311130313211320-2320001130333311"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
volterra_certificate = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321122203131312-2122202320313213-1022101332303201-3100332222312032-2011110010303333-3303033030030002-3113230312013221-1310020002020330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_forward_proxy.tls_intercept.volterra_trusted_ca` properties

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0030231112200120-2310101313112331-0112221320311330-1010030020110100-0320030301321031-0201122000110313-2302331223321203-3130322032000100)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-3230201320313300-3100103222213301-3222301220332000-2102123311113022-2103213133210102-1103110232333032-2123011010231323-2310223103003331)
- [enable_forward_proxy](resources--network_connector--reference--group-001.md#canonical-1101110030313313-2102212303300111-0320130013022202-3311211123122310-2200103232211301-3223222300011312-2202123000112213-2220311201321020)
- [enable_forward_proxy.tls_intercept](resources--network_connector--reference--group-001.md#canonical-1103300020233001-3233102132113202-0100231103312111-2133120322011110-3203012000032323-3033230123333013-2000221112111311-2302123032122113)
- enable_forward_proxy.tls_intercept.volterra_trusted_ca

<a id="canonical-2212331102231220-0330303122320011-2320323223122112-1112312022112200-3233013320232133-1012313203222302-2222022011332230-1233301000120002"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
volterra_trusted_ca = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1323220001110322-1200211010233233-0233223320201012-0231301213210211-1002133133033023-0030103033202111-0202301000230323-0310330231012320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sli_to_global_dr` properties

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0030231112200120-2310101313112331-0112221320311330-1010030020110100-0320030301321031-0201122000110313-2302331223321203-3130322032000100)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-3230201320313300-3100103222213301-3222301220332000-2102123311113022-2103213133210102-1103110232333032-2123011010231323-2310223103003331)
- sli_to_global_dr

<a id="canonical-3131313111101222-3033030230012301-1222310033313330-0021231222210301-2102321321231230-3121210102111133-1113322013231123-3100200331112202"></a>

Type: `"object"`. single nested block, Optional.

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

- [sli_to_global_dr](resources--network_connector--reference--group-001.md#canonical-3131313111101222-3033030230012301-1222310033313330-0021231222210301-2102321321231230-3121210102111133-1113322013231123-3100200331112202)
- [sli_to_slo_snat](resources--network_connector--reference--group-001.md#canonical-2301302022122012-1000012020121231-0330232303123223-1213230322230002-1032300010103112-2131000211233032-2301211000213210-2330002013211232)
- [slo_to_global_dr](resources--network_connector--reference--group-001.md#canonical-0322021212311130-2232200200011310-3031133210012000-2232031310233133-1232232110131100-3122323013032031-0122221132332013-1032011112221203)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
sli_to_global_dr {
  # Configure direct properties listed below.
}
```

<a id="canonical-2113022002102131-1220212331131001-2000111331102233-0332300323322123-2133222212113020-0330311133310230-2100112311110130-3101120111230132"></a>

### Direct properties for `sli_to_global_dr`

- [global_vn](resources--network_connector--reference--group-001.md#canonical-1220323023210322-2311312111021101-0233212121321210-2132100302120312-1020212033321133-1200231233223101-3212323103313120-3231101232000022): complete subsection reference.

<a id="canonical-1220323023210322-2311312111021101-0233212121321210-2132100302120312-1020212033321133-1200231233223101-3212323103313120-3231101232000022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sli_to_global_dr.global_vn` properties

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0030231112200120-2310101313112331-0112221320311330-1010030020110100-0320030301321031-0201122000110313-2302331223321203-3130322032000100)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-3230201320313300-3100103222213301-3222301220332000-2102123311113022-2103213133210102-1103110232333032-2123011010231323-2310223103003331)
- [sli_to_global_dr](resources--network_connector--reference--group-001.md#canonical-1323220001110322-1200211010233233-0233223320201012-0231301213210211-1002133133033023-0030103033202111-0202301000230323-0310330231012320)
- sli_to_global_dr.global_vn

<a id="canonical-2201233121011332-1332331331023223-2202011123122301-2112112320301001-0223233311121202-3322232130312331-1223021113201110-2013231011220232"></a>

Type: `"object"`. single nested block, Optional.

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
global_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-1303210331110311-2300301223113312-0221021202230000-2022122322320122-2221123201213312-3000001302130211-3133210203331110-1011012301212133"></a>

### Direct properties for `sli_to_global_dr.global_vn`

<a id="canonical-3331131021302130-0213121113333111-0013110230131032-3231200310032333-1202203211121011-2102202001130033-2033331202023030-2121110333220002"></a>

#### `sli_to_global_dr.global_vn.name` property

Type: `"string"`. Optional.

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

<a id="canonical-1033113332311303-0301330013312031-0301223303323233-2003301113233012-3022302121220121-2110200200331011-1032222013231200-2313010001310000"></a>

<a id="canonical-2021122022020002-0003033212001130-0231003021210030-2222312010313213-0312112210123212-0031303312112320-3021011012222200-3123021010313212"></a>

#### `sli_to_global_dr.global_vn.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-0102212121111101-0110013030101310-2213330032232000-1032120330100332-1201202303202120-0312122022222021-2301211022323323-1013230131132330"></a>

<a id="canonical-3333233322332312-3130331010022020-3332010220121102-0301132233302321-3303102311000023-2012022000132010-1303113003323000-1303003103112131"></a>

#### `sli_to_global_dr.global_vn.tenant` property

Type: `"string"`. Computed.

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

<a id="canonical-0311230213310203-1002131223302201-1301011222023101-2212122300321021-2103313222121202-2010010213310323-3111100010201210-0122233201102021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sli_to_slo_snat` properties

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0030231112200120-2310101313112331-0112221320311330-1010030020110100-0320030301321031-0201122000110313-2302331223321203-3130322032000100)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-3230201320313300-3100103222213301-3222301220332000-2102123311113022-2103213133210102-1103110232333032-2123011010231323-2310223103003331)
- sli_to_slo_snat

<a id="canonical-2301302022122012-1000012020121231-0330232303123223-1213230322230002-1032300010103112-2131000211233032-2301211000213210-2330002013211232"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
sli_to_slo_snat {
  # Configure direct properties listed below.
}
```

<a id="canonical-0100212312101111-3112021211101200-1103013002012203-0202000103210112-0232321112121113-0130331230201030-2012131032010120-2300102302111100"></a>

### Direct properties for `sli_to_slo_snat`

- [default_gw_snat](resources--network_connector--reference--group-001.md#canonical-0330111330011302-2120023101223202-2001031011130121-3022130131313032-1011210130323300-3220003110300102-0210321233001002-0331210133210032): complete subsection reference.

- [interface_ip](resources--network_connector--reference--group-001.md#canonical-0033200233313320-1002013322001132-2011213031023330-3003301101112222-1121221311132020-2201013322033303-1220103211230210-2133203010203011): complete subsection reference.

<a id="canonical-0330111330011302-2120023101223202-2001031011130121-3022130131313032-1011210130323300-3220003110300102-0210321233001002-0331210133210032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sli_to_slo_snat.default_gw_snat` properties

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0030231112200120-2310101313112331-0112221320311330-1010030020110100-0320030301321031-0201122000110313-2302331223321203-3130322032000100)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-3230201320313300-3100103222213301-3222301220332000-2102123311113022-2103213133210102-1103110232333032-2123011010231323-2310223103003331)
- [sli_to_slo_snat](resources--network_connector--reference--group-001.md#canonical-0311230213310203-1002131223302201-1301011222023101-2212122300321021-2103313222121202-2010010213310323-3111100010201210-0122233201102021)
- sli_to_slo_snat.default_gw_snat

<a id="canonical-1103133032301313-3011333333032320-0202023010010133-2121013001201100-1123103132110233-0002132123202203-0231011313101332-0330232122331122"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
default_gw_snat {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0033200233313320-1002013322001132-2011213031023330-3003301101112222-1121221311132020-2201013322033303-1220103211230210-2133203010203011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sli_to_slo_snat.interface_ip` properties

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0030231112200120-2310101313112331-0112221320311330-1010030020110100-0320030301321031-0201122000110313-2302331223321203-3130322032000100)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-3230201320313300-3100103222213301-3222301220332000-2102123311113022-2103213133210102-1103110232333032-2123011010231323-2310223103003331)
- [sli_to_slo_snat](resources--network_connector--reference--group-001.md#canonical-0311230213310203-1002131223302201-1301011222023101-2212122300321021-2103313222121202-2010010213310323-3111100010201210-0122233201102021)
- sli_to_slo_snat.interface_ip

<a id="canonical-0132220202122132-3330031121002010-0231221311003000-0122033021221131-2311020202302333-0011202020203030-2220120230220312-1132111001210032"></a>

Type: `"object"`. single nested block, Optional.

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
interface_ip {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3330311132311200-2332201021230031-1213333312000213-3123313301103312-3120201323231333-3303231003321122-2323202202011332-1322123201030111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `slo_to_global_dr` properties

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0030231112200120-2310101313112331-0112221320311330-1010030020110100-0320030301321031-0201122000110313-2302331223321203-3130322032000100)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-3230201320313300-3100103222213301-3222301220332000-2102123311113022-2103213133210102-1103110232333032-2123011010231323-2310223103003331)
- slo_to_global_dr

<a id="canonical-0322021212311130-2232200200011310-3031133210012000-2232031310233133-1232232110131100-3122323013032031-0122221132332013-1032011112221203"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
slo_to_global_dr {
  # Configure direct properties listed below.
}
```

<a id="canonical-1022220133002212-3302033300331011-2130032203103312-1003321030121220-0113310300002010-3333130231120032-2333102322010023-2213132020133210"></a>

### Direct properties for `slo_to_global_dr`

- [global_vn](resources--network_connector--reference--group-001.md#canonical-2213133222102121-1010330323120101-1203112033013023-0202300133120113-1303203330032113-0301310100313001-1323302001210333-1232031223232012): complete subsection reference.

<a id="canonical-2213133222102121-1010330323120101-1203112033013023-0202300133120113-1303203330032113-0301310100313001-1323302001210333-1232031223232012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `slo_to_global_dr.global_vn` properties

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0030231112200120-2310101313112331-0112221320311330-1010030020110100-0320030301321031-0201122000110313-2302331223321203-3130322032000100)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-3230201320313300-3100103222213301-3222301220332000-2102123311113022-2103213133210102-1103110232333032-2123011010231323-2310223103003331)
- [slo_to_global_dr](resources--network_connector--reference--group-001.md#canonical-3330311132311200-2332201021230031-1213333312000213-3123313301103312-3120201323231333-3303231003321122-2323202202011332-1322123201030111)
- slo_to_global_dr.global_vn

<a id="canonical-2031130013113210-1011001010123112-0023021111023332-2231220011032212-0123003002320010-3310123312002230-1132330331321310-1313233332021030"></a>

Type: `"object"`. single nested block, Optional.

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
global_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-3211123003300301-3332333231003013-1123030221213030-1010111003110122-2003103013003020-2023312230321210-3003332232101111-3221001233133203"></a>

### Direct properties for `slo_to_global_dr.global_vn`

<a id="canonical-2122330030232031-0030223330033020-1330010323132023-0222013311121031-2023202233102211-2133022011122122-1303032300031130-1112222120123001"></a>

#### `slo_to_global_dr.global_vn.name` property

Type: `"string"`. Optional.

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

<a id="canonical-2033210323023322-0132123332311330-0123233331231131-3033331131310133-0001202200212213-1021300101222030-2213311301031102-0101003310130013"></a>

<a id="canonical-2110330023213002-0032212233312032-2123103320323011-3130233130300033-2120011213001202-2300020121130232-0033033130130231-2212133131211302"></a>

#### `slo_to_global_dr.global_vn.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-2300002321011312-0231322120220213-3010031003313202-1232033310101032-0113302210303323-1201230131220301-1233000023221122-1302122031211111"></a>

<a id="canonical-2022323212222312-2333230321201212-0232330012013011-2221011100012200-3332110123110231-2300333012013333-2200310012033120-0320322331033221"></a>

#### `slo_to_global_dr.global_vn.tenant` property

Type: `"string"`. Computed.

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

<a id="canonical-3221303012220021-3113332122010110-2233110102322321-1221102001133222-3210331030002022-1131133021000020-0202000213031021-3012033020001021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0030231112200120-2310101313112331-0112221320311330-1010030020110100-0320030301321031-0201122000110313-2302331223321203-3130322032000100)
- [Property reference](resources--network_connector--reference--group-001.md#canonical-3230201320313300-3100103222213301-3222301220332000-2102123311113022-2103213133210102-1103110232333032-2123011010231323-2310223103003331)
- timeouts

<a id="canonical-1212011300100100-2331233233333101-0123323232132222-1102313101023011-3323230121000012-1312223001100131-1002132111203021-0233023221330222"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3003020221301322-2132113302201300-1031303231211110-3230212312321001-1331010001201102-0223202210320233-2033221203322303-2303113331323102"></a>

### Direct properties for `timeouts`

<a id="canonical-3022023301211101-2102012113220031-3201230232301021-3311202121000032-1113023131210201-3312332301213012-3022231203233001-3323022001220133"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1303031002212220-2313133130310230-1022311300130000-3231102123001300-0103132012110233-1102003132330111-1201123131321231-0313020301310101"></a>

<a id="canonical-0021302023311120-0001233231300023-2003210121213331-2101210220312112-2302020202002020-3213011121321023-3102200203110131-3130112112222122"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1120202311333313-0223211200330102-3000030332300012-1003123112012201-3032212122120001-0323200012332003-3023200221031222-1121111100122111"></a>

<a id="canonical-3211301032202013-3023223022212103-2022222233013110-0032123002230023-2321313301302132-1000121021322222-2201011321003003-1123201320303100"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3120232202000313-0012322203300320-3313000102010031-2101101113211100-2103133100211121-2121312203322120-2302303030223000-0111231032333232"></a>

<a id="canonical-1022230002030013-3232230102322010-3110210231223130-3213013201001103-1221300113232130-1122130313001131-0102022302211021-1021133033013212"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
