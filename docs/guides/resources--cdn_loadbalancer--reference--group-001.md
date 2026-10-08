---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- Property reference

<a id="canonical-1331110121211222-3033111223232031-1013032221321021-2132332300332101-3133220032221202-2122111211300013-0321211313201130-0303223112210133"></a>

### Direct properties for `xcsh_cdn_loadbalancer`

- [active_service_policies](resources--cdn_loadbalancer--reference--group-003.md#canonical-3130102200011303-0111031030113000-2110033302010100-0122110303103100-1302031032000121-2223322011312132-1131011012312331-2113001031103330): complete subsection reference.

<a id="canonical-0301303203110331-1122002012123120-1101332101032102-0012202003301032-3022321210112232-0223200232311021-2323330020123222-3221003210200211"></a>

<a id="canonical-1011012200313010-3310103301330321-1310112222033032-3231121323213012-2302103203302103-0113320031232320-0332222310101231-0223131100103201"></a>

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

- [api_rate_limit](resources--cdn_loadbalancer--reference--group-003.md#canonical-2133210031320010-1311202110033220-0033302203133300-0301223100233300-2031122313303032-3222000002123002-2313003121221300-1330112323321220): complete subsection reference.

- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312): complete subsection reference.

- [app_firewall](resources--cdn_loadbalancer--reference--group-006.md#canonical-0320211303002011-0303200110322030-3020012231113003-3312221223200013-2030312203122030-1312030231213021-0122121213001310-1020121102323120): complete subsection reference.

- [blocked_clients](resources--cdn_loadbalancer--reference--group-006.md#canonical-0101200112203333-2313100322333123-1133301001021123-0132230132312103-2222120211222232-1330312333203021-1213021020200201-2312220003310232): complete subsection reference.

- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100): complete subsection reference.

- [captcha_challenge](resources--cdn_loadbalancer--reference--group-008.md#canonical-0130021111111203-2032310323002002-0122300313003322-0110301130313122-2232111201113100-3201003313113133-0220131013102102-1223311220131210): complete subsection reference.

- [client_side_defense](resources--cdn_loadbalancer--reference--group-008.md#canonical-2102312123223023-1000221113330222-1301310330113210-3223103113203301-2201320003323112-0122122321013233-0222303013101221-3232122123323112): complete subsection reference.

- [cors_policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-1201021310323311-1300002333101322-1212100020212023-2033122312102030-0331223230312000-1222113322200232-2203221130113100-0233323002331212): complete subsection reference.

- [csrf_policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-1000130300033331-3002211233300130-2313130100130200-1302332133032013-0013323333213201-3223012120111001-2100331133200101-2301200003031301): complete subsection reference.

- [custom_cache_rule](resources--cdn_loadbalancer--reference--group-009.md#canonical-2210332032110200-0210201322313110-1131031023311132-0213211002011103-0233222021320322-2330011210220101-1033313321031003-0323003302010310): complete subsection reference.

- [data_guard_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-0131303033130011-1300120013211032-0221103230201231-3131000301020103-3202102131011030-0111033231131332-0212333123202100-2233232322111033): complete subsection reference.

- [ddos_mitigation_rules](resources--cdn_loadbalancer--reference--group-009.md#canonical-1113101222111100-0231333020320020-2111302302220332-3302220301331200-2213011011211300-1222011002133311-2332301203001000-3323001221330013): complete subsection reference.

- [default_cache_action](resources--cdn_loadbalancer--reference--group-009.md#canonical-2300021122222120-3103201200301030-2101110210210123-3021222103220303-1221111131320223-2202110300233100-2330133310101313-0202223013312313): complete subsection reference.

- [default_sensitive_data_policy](resources--cdn_loadbalancer--reference--group-009.md#canonical-2310301221200230-0201112131201231-1131101200303021-2130030310111230-3203022302002332-2030103200001230-3233113321131031-0233232233132333): complete subsection reference.

<a id="canonical-1220101012031021-2321220220222300-2001020033021302-3201121032100201-3310133031012210-0130213113013100-3133121313101123-2323000101212033"></a>

<a id="canonical-1323203000113100-3310210001013203-0001020333101213-2300310231031301-2220133002102022-3131213011313123-3320321033120231-1230311022212033"></a>

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

<a id="canonical-1010303000113323-3100222011102021-3232012302103322-1031012211321003-1211323013231012-2333211123310202-0100223313231332-3213012203021103"></a>

<a id="canonical-2203321311103233-1031312013303023-2012330302002233-2021222033230231-0302232211133332-2232011331033331-0221321010230123-3032321131101131"></a>

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

- [disable_api_definition](resources--cdn_loadbalancer--reference--group-009.md#canonical-3312231102230133-0002321102020331-1031013220300023-0101221003122222-2100330313000330-2331331200001221-0133311233301310-3133210321022010): complete subsection reference.

- [disable_api_discovery](resources--cdn_loadbalancer--reference--group-009.md#canonical-2130210132113322-0123113132000212-3230220223101021-1200003113211031-0113012202110012-1311330331231000-1333100002101020-1131001031331003): complete subsection reference.

- [disable_client_side_defense](resources--cdn_loadbalancer--reference--group-009.md#canonical-0322312112202232-3232301311330110-0300133101202313-0011202213000301-3113120123102301-0132223030100120-2202020201132330-2132133022103120): complete subsection reference.

- [disable_ip_reputation](resources--cdn_loadbalancer--reference--group-009.md#canonical-1031001202212121-2320213202323013-2310203020210132-3222031002011331-1231231310110332-2313123223113001-3022212102330213-3000321230000321): complete subsection reference.

- [disable_malicious_user_detection](resources--cdn_loadbalancer--reference--group-009.md#canonical-3323303332010223-3130303213321301-1233002121331100-0110203021132310-0033302313301323-2313122210131232-3213230202131103-1322110213131001): complete subsection reference.

- [disable_rate_limit](resources--cdn_loadbalancer--reference--group-009.md#canonical-2101211221210102-0323002031331010-2210220231120012-0023021303231032-1032303023320223-0003233332001012-3103103110223321-1232203302313021): complete subsection reference.

- [disable_threat_mesh](resources--cdn_loadbalancer--reference--group-009.md#canonical-0300001220010210-2011013010302113-1031230110020223-2000000323133201-2331022100022333-0022210002101020-3232010103322200-0011200213203231): complete subsection reference.

- [disable_waf](resources--cdn_loadbalancer--reference--group-009.md#canonical-1120023301221013-0010101221332312-2101301230211303-3223121323013122-2033123211312311-2201011122202012-0103112113322002-3003131323102011): complete subsection reference.

<a id="canonical-2023030233300132-1323132122102030-1032002121031230-3133210003120103-0112313233213311-2311030220230213-0301130101021012-0023120102302132"></a>

<a id="canonical-3311323312210023-2232232120312020-2102313331230203-1102322010101202-0113302303221302-3302102132222001-2120020313002333-2111312110323033"></a>

#### `domains` property

Type: `["list", "string"]`. Required.

A list of fully qualified domain names. The CDN Distribution will be setup for these FQDN name(s).
\[This can be a domain or a sub-domain\]

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 32),
}
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-009.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323): complete subsection reference.

- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-2332130232332220-3122100131132301-2112222332200300-2232023130202233-3132033002233230-3320011212210312-2203130020311011-3020303300221303): complete subsection reference.

- [enable_ip_reputation](resources--cdn_loadbalancer--reference--group-010.md#canonical-3132020110023330-0221123212220300-2303122132103021-2301032210323100-0112321332310001-2233030031121302-2100003011131031-0321200323303122): complete subsection reference.

- [enable_malicious_user_detection](resources--cdn_loadbalancer--reference--group-010.md#canonical-1303213031102030-3122302223322303-3002003203030013-0131022002010001-0002312200302222-3331123211021133-0230132330122210-3301210122110300): complete subsection reference.

- [enable_threat_mesh](resources--cdn_loadbalancer--reference--group-010.md#canonical-3301230133222123-3331320000010320-3300322323323202-1103213233103032-2103000031113100-2212030000211223-1213133100220200-0023212123220030): complete subsection reference.

- [graphql_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-0131310203002311-1002202102331210-3120001102321131-2313330213020203-3112000203032220-1033322112213220-3032123210010110-1010202111302223): complete subsection reference.

- [http](resources--cdn_loadbalancer--reference--group-010.md#canonical-3320012221332001-3023010200311133-0012202111211311-2312132120132210-3032223212022123-1212221020223222-2203232202220232-2203232310011031): complete subsection reference.

- [https](resources--cdn_loadbalancer--reference--group-010.md#canonical-2101233301010003-1200221022012203-0031300003110020-0322211021320210-2003201303102013-1022203021201030-0102222103131230-0232323033112111): complete subsection reference.

- [https_auto_cert](resources--cdn_loadbalancer--reference--group-011.md#canonical-2333321023320231-1103212001100120-1332321202002133-1230013202111101-3203220322212103-3223230223121121-1120131300100331-2001211211221323): complete subsection reference.

<a id="canonical-2203021201333211-1313012200332011-2133132221003211-1332100300200020-3012113331303011-2332013301001303-2330313111110320-0032220000321223"></a>

<a id="canonical-1202233030320323-3203103310302302-3102123201003302-2333230232112020-3122230231121302-2131312110303210-0123312213123120-3023020321022200"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [js_challenge](resources--cdn_loadbalancer--reference--group-011.md#canonical-0013121320212310-0321212022202332-3221102212101112-2133200201303301-2223113031321031-2230211002301113-1023222011212331-2310212110032321): complete subsection reference.

- [jwt_validation](resources--cdn_loadbalancer--reference--group-011.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031): complete subsection reference.

- [l7_ddos_action_block](resources--cdn_loadbalancer--reference--group-011.md#canonical-1302023000110302-3120111302312101-1232211223311220-1031332223001000-2302333000033013-1001023131000303-3232033331020302-2320323022012121): complete subsection reference.

- [l7_ddos_action_default](resources--cdn_loadbalancer--reference--group-011.md#canonical-2230330203013032-1133333031113100-0021112120003221-1230203101320213-0202130233200301-3000332131012213-3203322211203031-1131222233212002): complete subsection reference.

- [l7_ddos_action_js_challenge](resources--cdn_loadbalancer--reference--group-011.md#canonical-3230011220011130-3203232003333030-0210211220332212-1300220223001301-2310021232213132-0112010313031231-3000331122302003-2133202101030211): complete subsection reference.

<a id="canonical-2210000023122012-0301321111222332-0101323102120031-0020132310101223-0312011332033000-3100212333301120-2113033200002010-1333123011110320"></a>

<a id="canonical-0313132001133212-2233302132330221-1323030212201303-3031202120313223-0212330131123321-1210302221312333-0302222302103332-1003223310313133"></a>

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

<a id="canonical-2201200200201033-2221330001000032-3303213312232123-1123320023233301-1321111312012312-1310110313223021-3013012221221030-0221023033001333"></a>

<a id="canonical-0231130100121301-3012231313212332-2231130022331020-1330132230002312-0031101220113101-1031213111022203-2322122331120210-2312222311120313"></a>

#### `name` property

Type: `"string"`. Required.

Name of the CDN Load Balancer. Must be unique within the namespace.

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

<a id="canonical-2333201101202222-0221112010331213-1201021121131333-0211311213323110-3133013220223313-3010013213132100-3230023110312001-2331330122113121"></a>

<a id="canonical-1030001001310132-3001030230023200-1303332320022213-2021331013122320-0120310200333131-0223333300013200-1013231323123130-3222313010323012"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the CDN Load Balancer is created.

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

- [no_challenge](resources--cdn_loadbalancer--reference--group-011.md#canonical-3312000011031203-2231300221302120-2103312210100030-2210320110231313-2210213211211221-2111110102013130-3032201012110003-2023110221100031): complete subsection reference.

- [no_service_policies](resources--cdn_loadbalancer--reference--group-011.md#canonical-3123112331011010-0132220300333311-0001300132330131-1132331012130223-3311232302020312-2200131102201331-2233330000312303-0132221122110323): complete subsection reference.

- [origin_pool](resources--cdn_loadbalancer--reference--group-011.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312): complete subsection reference.

- [other_settings](resources--cdn_loadbalancer--reference--group-012.md#canonical-1202031330011101-0332311032231013-2313101330202033-1021201330201212-0201100211020002-3013210230001210-2223102333313322-3023223103202022): complete subsection reference.

- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033): complete subsection reference.

- [protected_cookies](resources--cdn_loadbalancer--reference--group-013.md#canonical-0002203113111232-3103331201000322-1000322330121231-1231211020332002-1300222301221203-3030113211312023-2111302010111232-1220233103210122): complete subsection reference.

- [rate_limit](resources--cdn_loadbalancer--reference--group-013.md#canonical-0211103021332212-1322020302031330-3011120223222013-3020020323221001-3323130330213221-2013323120323210-2013211203122002-2322130213211113): complete subsection reference.

- [sensitive_data_policy](resources--cdn_loadbalancer--reference--group-014.md#canonical-3101301311121220-1022212312223231-0103230132010000-1031303112000210-1003201331112110-2210011211032323-3113203213013300-1111103022320301): complete subsection reference.

- [service_policies_from_namespace](resources--cdn_loadbalancer--reference--group-014.md#canonical-0232201322200031-0330122100331311-3221203302302022-3021301122011232-3230010232020130-3332030103113102-0223113230122010-1011332330021133): complete subsection reference.

- [slow_ddos_mitigation](resources--cdn_loadbalancer--reference--group-014.md#canonical-1233003020202211-1121000220031121-0002002321202011-3200220102103002-3312211001303013-1111001323333210-3232211013001110-3302023302323112): complete subsection reference.

- [system_default_timeouts](resources--cdn_loadbalancer--reference--group-014.md#canonical-1302302300220120-1232133331030022-2202232123331001-1001230212313333-2002032030033031-0300122120322121-0023222330210101-0203103301320000): complete subsection reference.

- [timeouts](resources--cdn_loadbalancer--reference--group-014.md#canonical-1120320111222111-0221312232103321-2100002103330133-0121321020113111-3112222032011030-0332323220222220-0312012223333230-3320123022202133): complete subsection reference.

- [trusted_clients](resources--cdn_loadbalancer--reference--group-014.md#canonical-3322120013001311-3032003213121101-0001221002023010-2323123102111022-1020033201211233-0231122131211102-2122323232230303-2111233201201302): complete subsection reference.

- [user_id_client_ip](resources--cdn_loadbalancer--reference--group-014.md#canonical-3302200200302032-1033011220200321-0322330033221123-1310331022111310-3130333121221113-1313030021202201-1313311122312131-2320123301120231): complete subsection reference.

- [user_identification](resources--cdn_loadbalancer--reference--group-014.md#canonical-1112230101301222-3222000100021310-1113222132033220-1211130133022022-2010313030300030-0312321030200003-0211330333102113-1230112232300132): complete subsection reference.

- [waf_exclusion](resources--cdn_loadbalancer--reference--group-014.md#canonical-1021310101112110-2010221130101331-3210001303022230-0133011220020312-0113113013202010-1102203233130111-1212311201131021-2300222110021322): complete subsection reference.
