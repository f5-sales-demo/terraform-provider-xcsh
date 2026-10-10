---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- Property reference

<a id="canonical-1202221200022220-2102122200012320-0110233103003111-0030010230022221-2011033220122330-0013323031030000-1032020333331331-0012231233311221"></a>

### Direct properties for `xcsh_cdn_loadbalancer`

- [active_service_policies](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-0202323233031101-1323302210302322-1231003212131301-2002332233310112-1203102032311001-1101301313321102-3100320222201300-0312023130123210): complete subsection reference.

<a id="canonical-3202301032122032-3132220321221301-0013133201113202-3122022123231020-0233001332012231-0033010023001012-3203030012123010-0113120231032032"></a>

<a id="canonical-1310201001312022-1301132211032302-1302331230302132-2302102102233330-1103302303232001-1312031030323101-0110312130322333-1101031102323311"></a>

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

- [api_rate_limit](data-sources--cdn_loadbalancer--reference--group-003.md#canonical-1212231311001320-2333033311133030-2212321020003122-3130132300130123-0322230310020333-0302203033211323-0200320032322203-3201010011121023): complete subsection reference.

- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111): complete subsection reference.

- [app_firewall](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0113031001311221-2011130203032111-1002013120300301-3110122301220113-0213203302131033-1313023111332322-0211210010133302-0232301320303231): complete subsection reference.

- [blocked_clients](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1012001011110000-2201321320101011-1201112211033131-1103031022310303-3032023110111010-1101322020301222-0330111212300210-2230202303200212): complete subsection reference.

- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200): complete subsection reference.

- [captcha_challenge](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1103002330300201-1231313031223333-0303202310013222-0023112211102302-3010003113123000-2212320100321323-2032133101121303-3112310212212213): complete subsection reference.

- [client_side_defense](data-sources--cdn_loadbalancer--reference--group-009.md#canonical-1031112122022321-3012321000333023-2232001202012132-0321133111111102-1203101100022232-0213233213130113-2121311321122232-1121123211010231): complete subsection reference.

- [cors_policy](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-3330100113233220-0200103010232013-0220221020032201-2012031211002200-0321000202211233-0202201031003232-2322233111100113-1212133131020203): complete subsection reference.

- [csrf_policy](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-1311323223301232-0300332102000101-2321001332311130-2131032223022213-0212021121211332-3030332100231313-0332301223110011-3303032003211230): complete subsection reference.

- [custom_cache_rule](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0012310100301121-0211013102110231-1333303221223123-0323131031032113-2122320213230021-2211113110211001-3231311331000132-0313231022331320): complete subsection reference.

- [data_guard_rules](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-3232111322200122-1211030013030231-2030110132200332-3301322123212301-3113011300213020-1001032223131111-1323030222021301-3223022213321211): complete subsection reference.

- [ddos_mitigation_rules](data-sources--cdn_loadbalancer--reference--group-010.md#canonical-0203332321031033-3130203001012232-3010332203030033-2133312223332032-0321102121221133-1312222023101103-1322331131210123-2310301103323130): complete subsection reference.

- [default_cache_action](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1332133310300002-2012012123231201-2303220202332322-0013322233210021-2211322201103311-3322122203213210-1031203102223230-1301100022213013): complete subsection reference.

- [default_sensitive_data_policy](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1010012002311312-2002113220132033-1221331121112232-1021133132132021-2130332101202113-0221322231220102-3102103213111300-3013223032223130): complete subsection reference.

<a id="canonical-1012113000021230-3312021213123233-3110231320033302-0100302313012230-2300321032333223-0320322200222213-0113122131101100-0132001303002000"></a>

<a id="canonical-3311333221310320-2003300202012303-1131200321122312-2033010333310323-0312211101002023-0023100210312130-2112233230331301-2132203132331010"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the CDNLoadBalancer.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [disable_api_definition](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3133121230022100-2022230330003230-3321032220303102-0222331220132321-1003003002112321-2113013000232022-2132122322130220-2202000001220303): complete subsection reference.

- [disable_api_discovery](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2210310302113200-0331303012200133-2220232021130020-3021233123102120-2002110232001123-2011310312211231-3221320011121131-2310031201232130): complete subsection reference.

- [disable_client_side_defense](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3310130022320333-2023113130232022-3001123220220101-1211301202220220-3211221311101222-0131032021311110-1330300000010302-1132112202231013): complete subsection reference.

- [disable_ip_reputation](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1230330223220032-2320012133303033-2132103221322133-1121210333222032-0231301133223331-2312100030231031-3133012220032010-0022222303103100): complete subsection reference.

- [disable_malicious_user_detection](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1200122332202011-0321013102011031-2230312120233300-1320313121323121-0112202233132321-3110032333203110-1102302030023231-1110320100000231): complete subsection reference.

- [disable_rate_limit](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2103233312101321-1012103031202211-1313020110322221-3331212202131021-3223002230233033-1202020313103331-3110132313033031-0121221031132120): complete subsection reference.

- [disable_threat_mesh](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2330031120223001-2010001102032332-1020100131030200-2020123121201321-3030200301231003-3000132310302113-1001322311010112-0120000032212121): complete subsection reference.

- [disable_waf](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1011233221222131-0010110333302333-0131112313222132-1210002303313123-1230101332030311-1003310103111003-1103032133222033-1011313002220211): complete subsection reference.

<a id="canonical-3331233232310311-2112310212231201-1331212302123221-0311001230131102-3222031111021030-1330311333212320-1220302200000233-0200230213301332"></a>

<a id="canonical-2300201213022212-3200322302101122-2232312303321301-0102320303131313-2210011211211102-1131030212021301-2030102301100022-0221112103311001"></a>

#### `domains` property

Type: `["list", "string"]`. Computed.

A list of fully qualified domain names. The CDN Distribution will be setup for these FQDN name(s).
\[This can be a domain or a sub-domain\]

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.pattern": "[\\\\.]+[A-Za-z]+",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.pattern": "[\\\\.]+[A-Za-z]+",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [enable_api_discovery](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1031021330032101-2321133230002323-2010000001130032-2132210013020010-2111220000123113-2321111233123321-0002100332010323-3322300231021213): complete subsection reference.

- [enable_challenge](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3100203021021010-2201131230132213-0121112120112123-1013013330103303-0200011323230211-3000331323222323-3133302303313021-3321010120100122): complete subsection reference.

- [enable_ip_reputation](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3313210111131130-3013223323200003-3222320132312223-1223312001120130-3200303232330111-0021013210202313-3232133003332212-2303232210322300): complete subsection reference.

- [enable_malicious_user_detection](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-1203313010102222-3212102003230230-2302023211203303-0222300301333121-1211131120202323-0002103332313100-2010320313200122-1201232211313203): complete subsection reference.

- [enable_threat_mesh](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3300213101122201-3232231121230323-2211130201313010-2120031103123202-3222331012020323-0110112112321033-1223200100130021-3333301222320233): complete subsection reference.

- [graphql_rules](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3301232011322201-3322122212320230-0212320203323203-1233020330031223-0013321230233213-1003032122130202-3223331312211021-0020303030213011): complete subsection reference.

- [http](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-0000202303000030-1313011110300110-2010231120333333-2012111302020010-2232013230200230-3211311320303033-1310131202202210-0323010122200301): complete subsection reference.

- [https](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-3230312322320110-2123202322212223-0312012332030102-0320301210203212-1331020322112130-2221022302312033-0120321223220102-3121132202012233): complete subsection reference.

- [https_auto_cert](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1022110223113311-1122120102013303-2121023120231002-0103032202213002-0201010020313211-3100130013113312-0132133320111300-3120303011211300): complete subsection reference.

<a id="canonical-1212212130311313-3101031131313132-1311113322333032-3000323320033130-0120120321200000-0310112000000033-3132112210113020-1312103012101212"></a>

<a id="canonical-0320023332132330-3013300121120020-2000131000333312-0122111110322320-2323220103211313-1121002232132122-3120023130322023-1121120203310222"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [js_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2302310323202313-1111033213121230-1303220321000110-2312133012002202-1332110200301033-3033101101203021-2210102112012310-2223033331003110): complete subsection reference.

- [jwt_validation](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-3221300002102023-2131303022302312-2013133022023320-2323211023010130-0213232211110303-2022001333033230-3112000221200233-0200013031112310): complete subsection reference.

- [l7_ddos_action_block](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1333333200120332-2013203303220313-1010312331020333-1200321100302113-0213120311020001-2122222320111122-1131233320112210-3120230011331311): complete subsection reference.

- [l7_ddos_action_default](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1200300131023022-3211332031133233-1011202032330202-3130010032100021-3311130322311121-3313332200131021-3212002210303013-2323210011000003): complete subsection reference.

- [l7_ddos_action_js_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1000121000031123-3101310311110223-2321303023200233-3323223212003123-0212021030302130-1321311333202303-0011001201301122-3311213103331020): complete subsection reference.

<a id="canonical-2312310023332211-0003310233023123-3220333322033231-2210201000033112-1022233310000302-3313300301112032-2223210222010302-2223100101311102"></a>

<a id="canonical-3133101033230331-3211020122330323-3331022100330130-0220000001313020-0322121313121203-0300031213122303-3031121001302030-1032311030130023"></a>

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

<a id="canonical-0222101301112000-3103123020131310-1220123332200221-3103300100331022-0030111320002331-1303003303320323-3111033123102031-0002333012310202"></a>

<a id="canonical-2211011100121030-2302220132113001-1313121000300000-3133101231312023-2111233003201321-1010233120212230-2133101302320220-1301322030103300"></a>

#### `name` property

Type: `"string"`. Required.

Name of the CDNLoadBalancer.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1120203320122212-1100112313103300-1111020123231101-0210120302021322-0022020023323022-0310101332231331-2000300102200203-1310321120232110"></a>

<a id="canonical-1123132133211311-0000220012323330-1300302232100133-3211130220133332-2201121100223103-1001122320033300-0021313030131131-1322331022020230"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the CDNLoadBalancer exists.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [no_challenge](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-1232110221313322-1131211111332023-1003031130302012-1213030301210330-0103133011322131-1112113122330123-0213220333222100-0202331332301210): complete subsection reference.

- [no_service_policies](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0100133220321213-2123201303210012-3220032001222313-3010303023322211-2133122331211121-0312122103223102-2233330202130123-2213221132032212): complete subsection reference.

- [origin_pool](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0223231232013221-1031133321030110-3120202101012130-1330201133233301-2333202311002001-1210313131130123-2211221120033030-3320230211121110): complete subsection reference.

- [other_settings](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0331123030101331-3110222312202112-3321300002131202-0110203121023102-3022013301122310-3122220023301103-1113321203000113-2311022130102310): complete subsection reference.

- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032): complete subsection reference.

- [protected_cookies](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2023232301302102-0212232022203201-2200202312020130-0231000031321122-3332031322021011-0211302203032230-3121202032220100-3232322100313223): complete subsection reference.

- [rate_limit](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0102233231103312-3003333110331031-2212203312231300-1123020223313130-0000231100201301-3230312232210312-1323001032103133-0332221002012303): complete subsection reference.

- [sensitive_data_policy](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-2323021210132303-1002202011101212-1233211310323332-0231102230303132-1000132101121121-1023313113211031-3200220110120300-3202302103311313): complete subsection reference.

- [service_policies_from_namespace](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-2013200020130301-1113102330210232-2012033221211111-1333232332201202-0131030031333211-0001020213120320-0310201210113323-0331122232203010): complete subsection reference.

- [slow_ddos_mitigation](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-0001301213230111-3103221002232321-2201322121023120-2011320222222231-0223121210332333-2311331313122203-3000301231032130-2101022223112121): complete subsection reference.

- [system_default_timeouts](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-3322032032102002-0320311003220131-1102031321312113-0210213301323200-2120310230130103-3321101121233321-0212132023110130-0313002323330033): complete subsection reference.

- [trusted_clients](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-1002100210320110-3201332210010310-0021103002112133-3002233130101311-0011130210203103-1302313010232122-0212020011000112-0312323220130232): complete subsection reference.

- [user_id_client_ip](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-1101232232111221-2200012223301301-3121031112233231-2313000312132330-2230301033331100-2213101132232022-2002003032201231-2132233222201102): complete subsection reference.

- [user_identification](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-3030211022123120-0133111110122030-3220101023132303-2132132233213000-3210220030203103-2301322311111110-2301202312112311-2202100000202220): complete subsection reference.

- [waf_exclusion](data-sources--cdn_loadbalancer--reference--group-015.md#canonical-0003231021200200-0111123130110021-3222203002002223-0130001133010311-2001112003301020-2221002010123132-0022333031103000-3212320223300313): complete subsection reference.
