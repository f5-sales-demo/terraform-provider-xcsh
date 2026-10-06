---
page_title: "xcsh_bigip_http_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bigip_http_proxy reference."
---

# xcsh_bigip_http_proxy reference

<a id="canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- Property reference

<a id="canonical-1220223210111010-1031200203230101-1210122233310322-0023120322212302-3102230112332313-3213110113222311-0100203230312021-3122222301001230"></a>

### Direct properties for `xcsh_bigip_http_proxy`

- [advanced_profile](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0211003021332023-0013000312333003-0331301233132013-2000113303313022-0231112311002001-1300321212312100-1303030031223022-3332202130110223): complete subsection reference.

<a id="canonical-2111022303333002-3312130233330332-2202031233212321-0202233310102001-3121131131323122-0213211100103102-0211303303302123-2310232032221213"></a>

<a id="canonical-0332203213010113-2202203300211300-0302003232232120-0303221110222201-3033030331301232-0032301113223123-1133333232010131-0121303333011301"></a>

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

- [ddos_profile](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0202231310212321-0320003212210123-2220233331201313-2321032221213302-3030203022233130-3201133323102200-2331221232230102-3010231100102222): complete subsection reference.

<a id="canonical-2022301223002031-0031202113002123-2212323213221022-0302113130013203-3221032002312230-2031002102321023-2321330113011112-0332213322001001"></a>

<a id="canonical-0103202102013010-0213311333312220-2023331113011303-2213133031110111-1311032220031203-3031021102012210-0300133232210331-0202313202120020"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the BigIPHTTPProxy.

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

<a id="canonical-1223113323023231-0133133321321313-1011022201213110-3220310221230002-3332333101021203-0023321323233100-0012110111232232-3102121000223203"></a>

<a id="canonical-2210312022331121-2002313101133233-2033322230011031-1311223003312023-3121033332323113-3113112322330203-3313023132011130-0010010310020011"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [irules](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0122030321111211-2222212201112202-1210111200130212-1030302210311120-3100110333000313-3320233100320332-2010120111203312-1031100312332303): complete subsection reference.

<a id="canonical-1333002300112122-0012201332030023-2320233002003233-3223121231033310-0310201200300122-1021022302120300-2112230003301012-0002102332222201"></a>

<a id="canonical-0200333101113102-3112032111230320-3330200313213101-0313101013321331-1223201220123321-0002311232201311-1300112001312331-0131232312212211"></a>

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

- [lb_algorithm](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032310323023122-2220201111120120-3332302022102211-3213221130321010-2302211131303223-2203232003221211-2033103110122123-2130230031113001): complete subsection reference.

<a id="canonical-2232210122222301-1211201233022320-0030102113330003-3330030100033223-0133213211323330-1132120212213130-2203301230220312-1031331100300120"></a>

<a id="canonical-1011332200011302-1100103133321120-1001132100210112-0332020201331312-3013121331020233-2123010030023123-1021131112010012-0113002320213020"></a>

#### `name` property

Type: `"string"`. Required.

Name of the BigIPHTTPProxy.

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

<a id="canonical-2231003103321133-3100332332030313-2102102120120023-1303023121032111-2130331222100102-2132112120023201-1201003220113030-3302002303113221"></a>

<a id="canonical-1230100303202203-2033130002001011-0110020212201031-2112130100212032-0031131321023133-0222230130011330-2323031130120023-3313211100302233"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the BigIPHTTPProxy exists.

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

- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0231332111103211-2012003310023331-2332323010101232-0330133310111212-2010210001000022-3233033013001220-3131212121202013-0003211311200232): complete subsection reference.

- [proxy_advertisement](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3002122313302023-0112221032312211-1331302003112211-1311312120212303-0230110230330223-1103003130221111-0132323003002022-3002103103013123): complete subsection reference.

- [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2122312232201311-0321313133122010-1201332120301102-0201322123231112-1110023021131120-1003212310023132-2311131101031333-2330230131010322): complete subsection reference.

<a id="canonical-3012201301200003-2223022232333113-0320021000013102-0203213020130010-3222233130102132-2332230322311103-0020100312201321-0120202223323131"></a>

### All schema paths for `xcsh_bigip_http_proxy`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `advanced_profile` | [advanced_profile](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1232021002110000-1312021212200231-0100300321102031-3032103331120332-1232331022211012-1010302112200103-0301220023311130-0312211200203310) |
| `advanced_profile.disable_spec` | [advanced_profile.disable_spec](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3021130012210331-0132101331131103-3111010021000021-2300232303101102-3223211303320110-2201232301113010-2011021333332010-0331310203033331) |
| `advanced_profile.enable_default_profile` | [advanced_profile.enable_default_profile](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3311013220231033-3111213113301122-2000233031303302-0011232223301303-1030221331202330-2012133022101020-1000201210223012-3222100313321200) |
| `annotations` | [annotations](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2111022303333002-3312130233330332-2202031233212321-0202233310102001-3121131131323122-0213211100103102-0211303303302123-2310232032221213) |
| `ddos_profile` | [ddos_profile](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3301030302312122-2222031033132311-2233221120231211-3222213031313210-0212311333000132-0010333312312230-1131011113012012-3012021133230013) |
| `ddos_profile.disable_ddos_mitigation` | [ddos_profile.disable_ddos_mitigation](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0320311121132203-0313213102331230-0131302210021021-3311120022300122-2212330002220011-3020331212231231-1120131220300311-3213300222121233) |
| `ddos_profile.enable_ddos_mitigation` | [ddos_profile.enable_ddos_mitigation](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3213231312201031-3312022121302211-3003102201302000-1003332133201202-0030223112003303-2132332303201000-2131330210130223-1222223201001122) |
| `description` | [description](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2022301223002031-0031202113002123-2212323213221022-0302113130013203-3221032002312230-2031002102321023-2321330113011112-0332213322001001) |
| `id` | [ID](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1223113323023231-0133133321321313-1011022201213110-3220310221230002-3332333101021203-0023321323233100-0012110111232232-3102121000223203) |
| `irules` | [irules](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3230220120322103-3032030022322332-2111033130130321-3133321120333020-0323102131112111-2230012103201133-0012232333030000-2232101000120123) |
| `irules.irules` | [irules.irules](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0021310310131000-0221202202130232-2122202112123322-2030221302021132-1022223020020122-2001010122330221-1100320120213300-0131132232200212) |
| `irules.irules.name` | [irules.irules.name](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2110212020321323-3333231302120132-0233130012130313-1122032313212101-1331103203333222-1200331332001202-0020212001301001-1001213103013131) |
| `irules.irules.namespace` | [irules.irules.namespace](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3333302313122103-2211330333130003-2333100213022302-3101033011302133-0202013332012002-3003001102301303-2220133230131223-1000311131221002) |
| `irules.irules.tenant` | [irules.irules.tenant](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2201002031202320-3110230103310323-0130311220210013-2310330303230232-0321332221000330-3013213131331031-1331221113010121-3323000013133202) |
| `labels` | [labels](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1333002300112122-0012201332030023-2320233002003233-3223121231033310-0310201200300122-1021022302120300-2112230003301012-0002102332222201) |
| `lb_algorithm` | [lb_algorithm](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1031230221331200-1313130100120310-3023300003233212-3121113220003201-3222332021232030-2133200301100211-0302132310023202-2131230132323312) |
| `lb_algorithm.round_robin` | [lb_algorithm.round_robin](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2202111331133010-2021130131223121-2200211132230120-3000001302123330-3313001131222130-2122302331100031-0210332203000231-1210232110011231) |
| `name` | [name](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2232210122222301-1211201233022320-0030102113330003-3330030100033223-0133213211323330-1132120212213130-2203301230220312-1031331100300120) |
| `namespace` | [namespace](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2231003103321133-3100332332030313-2102102120120023-1303023121032111-2130331222100102-2132112120023201-1201003220113030-3302002303113221) |
| `origin_pools` | [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3121202312031132-1020131321031032-1013111002322033-3132203333121112-1002322221210303-1103202200030311-2221310012303110-1032133321131001) |
| `origin_pools.pools` | [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1200323311003121-1322221220032011-3011022131300231-0220110130311033-1213102123123013-3220312123302033-2010223300311300-1213301021122033) |
| `origin_pools.pools.name` | [origin_pools.pools.name](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0100213132033302-1330302001101110-1120333111331002-3012233331020210-1103031333210311-3333332230221202-0202202033302003-0202103132203121) |
| `origin_pools.pools.origin_servers` | [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1030300033303001-1023233112101101-1222031112331221-2333111022110121-3100030210230122-2013332332111330-3120123210101331-0120333031203113) |
| `origin_pools.pools.origin_servers.automatic_port` | [origin_pools.pools.origin_servers.automatic_port](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2321111210202201-2023132322211103-1033131213100312-3111020332122102-1320100011021100-2323002320320203-2013101012300111-3330020113300210) |
| `origin_pools.pools.origin_servers.health_checks` | [origin_pools.pools.origin_servers.health_checks](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1331211032200131-1000332003312321-2300313023002213-1121310130013230-0200010031221313-2230330302130213-0021330202223212-0202012201013130) |
| `origin_pools.pools.origin_servers.health_checks.health_check` | [origin_pools.pools.origin_servers.health_checks.health_check](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3320322030231333-1302211102023202-2033201212313022-2211300130031233-0123311100211330-2131033111020033-3311302132103313-2002021110021033) |
| `origin_pools.pools.origin_servers.health_checks.health_check.icmp_health_check` | [origin_pools.pools.origin_servers.health_checks.health_check.icmp_health_check](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1012002023222010-3102323321100213-1003002131110130-0301121011302103-2233032301230123-1002112311323303-0030210221313311-3112200030321030) |
| `origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check` | [origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0201032232322221-0321013213033021-2110220213022231-2310230233010232-0112211112302232-1113130132020032-1001031111220013-0230131113023333) |
| `origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check.expected_response` | [origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check.expected_response](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3030312310330121-1300000231320300-1301320020103131-3203233311112310-3312120303012000-3320103233312110-1321310223221210-2323132301132010) |
| `origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check.send_payload` | [origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check.send_payload](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0002103322123231-1012230322110203-3230231323220000-2003322110022133-3331101211003033-3303310333210022-0123120213313003-3301100022012330) |
| `origin_pools.pools.origin_servers.health_checks.healthy_threshold` | [origin_pools.pools.origin_servers.health_checks.healthy_threshold](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0111031230130022-0332002222000011-3012302311002131-3301020133310332-3313210102210303-0302013000001222-2201130202130330-0212000301111313) |
| `origin_pools.pools.origin_servers.health_checks.interval` | [origin_pools.pools.origin_servers.health_checks.interval](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2322101201101201-1100130221100312-3321201133201023-0323222030121010-1230133111313023-1321223100031110-0313102232312033-2210221131110301) |
| `origin_pools.pools.origin_servers.health_checks.timeout` | [origin_pools.pools.origin_servers.health_checks.timeout](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2021202332131331-1113121330103103-3133313232213021-1331100320032011-1220203232133333-2030302032312113-3010111211331211-1233323123102200) |
| `origin_pools.pools.origin_servers.health_checks.unhealthy_threshold` | [origin_pools.pools.origin_servers.health_checks.unhealthy_threshold](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1312030133313030-2332331011031130-1032212311000133-3321303301220032-2213100003011202-1121330130301021-0312233211221101-1302212111100111) |
| `origin_pools.pools.origin_servers.lb_port` | [origin_pools.pools.origin_servers.lb_port](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2223303302123221-0301022330000013-3202232221023020-3320110113212120-1321100010032020-2212121130003222-0222031223223323-3033113303300323) |
| `origin_pools.pools.origin_servers.origin_servers` | [origin_pools.pools.origin_servers.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3323110213211100-0133232221103303-1333122212102302-1102021212113212-2230012013201010-0023120230123120-3020331323232131-3100021011210330) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service` | [origin_pools.pools.origin_servers.origin_servers.k8s_service](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0031010331320310-3321323231110011-2200210333231100-2332210030112311-0010222312233103-0102212131220313-0321303133033310-2020023123202020) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.inside_network` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.inside_network](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1300002131321021-2221113211102111-2002101303003032-0231102100311002-0132100110300300-0210313010223110-0332333323030213-1322310221300122) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.outside_network` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.outside_network](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0221231222003311-3200303030200000-3031033113201322-0311223000301033-3330311013132123-2201200213330220-2223322231310002-2320201213302312) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.protocol` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.protocol](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3303320021121202-2123303232200020-2200102002012302-1213033232230232-3232202110221330-1233312001232210-2131130013233300-2121033103103211) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.service_name` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.service_name](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3210120330332121-3032011023230202-0221331100103133-1011312022122021-0221222300201122-0133310233200131-3322212122311233-0232300020310201) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3121201200322000-2312120223232310-3121102302012021-1313230330100330-2012001101311301-1330130002312112-1132213201001221-0110030300032220) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2202322301232233-0102000322020323-2202030223013111-1002202203303210-1022100030011123-0220010302022223-1330100231311000-1330020222102022) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.name` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.name](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3331033123211113-1233113001013020-1110132331020232-0223232200231220-0213012100112302-2211032010103000-3130001121133113-1222312320321100) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.namespace` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.namespace](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3301033332332030-2033231131012300-3011012203012200-2203300303133312-0003323232011320-1032123010000312-1230203110131012-3111322000302002) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.tenant` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.tenant](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1021032003000233-2233211222230012-3103022013023223-3110011110003303-2212221310003033-0203222203100310-3330222212203031-0333101233322011) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3330321300301311-0203010320200033-0022312112331323-3011021121022231-2103113220302222-0331302130331301-2203202021223120-0113112001013013) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.name` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.name](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1000212123111013-3333002213012211-2201331310122012-0023332233023223-2112010103333010-2230023133303110-0303223323130022-1011230332113101) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.namespace` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.namespace](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0310230010230131-2010011231123302-1021102013110032-3123201001001031-2230200033222103-2200010013032113-0031012200200232-3231230033130131) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.tenant` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.tenant](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2103211112110021-2200122301000213-3021210032010301-2133100011131010-0112201302101130-0322030222220110-0223113201023120-1311332233020222) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1002303133331100-0103322110233111-3232300032302233-1031231122333011-2320121123011311-2002310211102303-2010222113331203-2121001302132120) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0202110131032323-1110103332323310-0230131012100131-2331212213103213-2120301312011031-0001200110021012-2312113210033021-3112001301010222) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3020201232130102-3021202210032021-3120030303010100-1313121221223322-3331102331210303-3122201230133212-0320313320233221-2003223203323120) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool.prefixes` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool.prefixes](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0031103033222232-0300330100211312-2212002002222200-0202002112211100-3023211130021303-2222333101001310-1230310200222201-1313032100300312) |
| `origin_pools.pools.origin_servers.origin_servers.k8s_service.vk8s_networks` | [origin_pools.pools.origin_servers.origin_servers.k8s_service.vk8s_networks](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2331312332101113-2003323101322000-0003311303131123-3302031223210003-1032132313032223-2102113020031231-3023211023030031-2210020112000101) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip` | [origin_pools.pools.origin_servers.origin_servers.private_ip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2331310112133313-1233200310301033-1121111223333013-0111322001120320-0230101230300101-1030212320011122-3332133011103032-2031301020000210) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.inside_network` | [origin_pools.pools.origin_servers.origin_servers.private_ip.inside_network](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0010322203003011-1212000123103300-2332230032003330-3330122221000113-0333332013323210-1003212202111121-2201123321021011-1303022132102113) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.ip` | [origin_pools.pools.origin_servers.origin_servers.private_ip.ip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2023303032331010-0201031130103000-2300100330010320-0233110302321003-3000023320233310-3121310121031013-3103231001301120-3212333203103303) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.outside_network` | [origin_pools.pools.origin_servers.origin_servers.private_ip.outside_network](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0121202032233212-1310133303311002-1310130233232333-0003102320123300-1131320332320223-0322312122223232-3231023202202133-0112123122103213) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.segment` | [origin_pools.pools.origin_servers.origin_servers.private_ip.segment](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1113313212013321-2022103022132332-0032330333130230-3311132113102000-3222221321202022-2330311213033030-1113103122221320-1003210200211030) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.segment.name` | [origin_pools.pools.origin_servers.origin_servers.private_ip.segment.name](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3220032020132013-0132031230030132-2223212012011002-3122311313001123-1010131023032120-0311132223321321-0111312001003300-3130323201020300) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.segment.namespace` | [origin_pools.pools.origin_servers.origin_servers.private_ip.segment.namespace](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3303033122022030-3133133213132323-2312033013100033-0300010022122202-3113000322232023-3221320032121100-0221103133211202-2002313221021020) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.segment.tenant` | [origin_pools.pools.origin_servers.origin_servers.private_ip.segment.tenant](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0102231011123312-0012030313030201-2212232331330112-1113132203121213-0223321202110311-3203011231300321-0200010100000021-0230331323100122) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3302002313311110-3230032223330121-1312310222223032-3011101321100013-0022202013211300-2100322123120212-0231022112103031-2300221213323312) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0000123300013210-1211120333113203-0132010312033211-0123310021003030-3033020122010021-3103132022121030-3001230301210123-0202231310311312) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.name` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.name](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0313020003320111-0030011133102310-0130200300032333-3212313301233232-0013323130032213-2223012113323022-1122012123201223-1211223200131121) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.namespace` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.namespace](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1103212311122202-3100101300231132-1213123111130221-2130200012301022-3022003203331312-0320231210122122-0133103100021113-0231002210000012) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.tenant` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.tenant](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1320303210231232-0010311302100132-0201331030213113-1120331311230223-1322031001313313-1320112013021201-0021313322031300-1223212031320320) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2112221121330022-0013300222002200-1122221331020100-1000022112113312-1330221130121003-2111210130201323-2111001201123302-1321120130103120) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.name` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.name](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1333300110133201-1101131001201030-2310230320222123-3000312000313121-0321203032311103-3333221320311212-3313031003333011-2233110101011230) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.namespace` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.namespace](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3100120002320023-2201100013231313-0133312011212232-2123330313321021-1232201102312101-2131102320221010-3130120333003301-1031020132212333) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.tenant` | [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.tenant](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0121000310011132-0031213203011330-1331222321021311-2100121222130212-1010000010003200-3222033311112322-2200321210110032-3223011012020333) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool` | [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2022030121202102-3010303020221303-3031332323133112-2311002033033021-1031130301330013-2320002000223111-0011123303012200-0300010003300002) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.no_snat_pool` | [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.no_snat_pool](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2213212310113002-3223303121311130-0010010332230020-3013332232332123-3010212321003100-2333301001313133-0110303321101222-2031120201231110) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool` | [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0230133210232311-2100300232330102-0122002303213112-3302032123211221-0212332202221330-3113022301332122-3133120302233230-3323020220230311) |
| `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool.prefixes` | [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool.prefixes](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2321012201211133-3031331122000113-1102222302131222-3100133210211002-1233213212322323-0032331313211112-1121000120103032-3121333001022033) |
| `origin_pools.pools.origin_servers.origin_servers.public_ip` | [origin_pools.pools.origin_servers.origin_servers.public_ip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2113011032311101-3113133101120122-3300030131200130-1332301310320033-3032230321113110-1313130001012132-1101331012131011-3001000200022333) |
| `origin_pools.pools.origin_servers.origin_servers.public_ip.ip` | [origin_pools.pools.origin_servers.origin_servers.public_ip.ip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3030103313031333-3200032100310102-1023120101222313-3332121012131323-2332301000331033-0103021131032203-3321001101202030-0010330131333000) |
| `origin_pools.pools.origin_servers.origin_servers.public_name` | [origin_pools.pools.origin_servers.origin_servers.public_name](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0130012011320102-1021302121303132-3000331123211113-1233222001210101-3110110331121312-3201330033101033-1021233233330221-2312201031320100) |
| `origin_pools.pools.origin_servers.origin_servers.public_name.dns_name` | [origin_pools.pools.origin_servers.origin_servers.public_name.dns_name](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3010022033023220-0212200300302231-2333033001203303-1101131022203213-0120111201020332-1103011121030231-3130130310212121-1331131010322112) |
| `origin_pools.pools.origin_servers.origin_servers.public_name.refresh_interval` | [origin_pools.pools.origin_servers.origin_servers.public_name.refresh_interval](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2133222010200210-0321012220311021-2220013112223300-0011220100222333-0030033312013010-2323221313212203-0202100102113113-3222232121023001) |
| `origin_pools.pools.origin_servers.port` | [origin_pools.pools.origin_servers.port](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1212203013123032-2230020310102220-3201032203301121-2020002332132010-2310130312311112-2330321322031012-3322130111103203-0120122030122101) |
| `origin_pools.pools.priority` | [origin_pools.pools.priority](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2123032312331103-2000031232321322-1312012320110113-2301330313322232-3210133202320202-2301100320201001-0201122121223132-1113233000232300) |
| `origin_pools.pools.weight` | [origin_pools.pools.weight](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1101221200022232-0213212101011100-2232102222123010-2223132211101011-3202011131331310-1032020012233100-1313212201330101-1110331121323113) |
| `proxy_advertisement` | [proxy_advertisement](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0020211011012111-0120022332000031-3102023112001321-1320200010231012-1001021323010201-0103220210122213-3111221320232220-3312130330033222) |
| `proxy_advertisement.advertise_custom` | [proxy_advertisement.advertise_custom](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3230111111231132-0222333311021112-1001120322312330-3322110123213122-1013131302001031-2210311212320310-0101132322010300-0201103231102332) |
| `proxy_advertisement.advertise_custom.advertise_where` | [proxy_advertisement.advertise_custom.advertise_where](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3002113210120210-3211302001002102-2300301120211221-0310002123213201-2303002203121313-2110103301022331-0323313211302221-0112101002130331) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2222321231312023-3013132311311232-0123210320120020-2013123023303103-2133232133200313-2012020110111022-1031222130313220-3022331231033222) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3033212131211122-1301202010102021-2230312312030312-1123330211301010-2011231010220133-3202012302201222-2100330021103003-0010320203222103) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2203313303323333-1330121131232232-3322200120012322-0232201330120001-2321003101100323-2103110122323031-2021313320303123-3020023233130223) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0311010010310101-0201103201330323-2011331210113233-0210133300233203-2022121022001133-1301110122023120-0002020233112203-3213231000101323) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1330321132220330-1233121030122223-2313003333022020-3302323212132320-0311212301032013-1103101031110220-3021131113323302-1131013303331302) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1021311111323300-0013012010022031-2332332102301323-1202212213132322-2010310011223103-0103133221101201-1320023103220212-3200320302321300) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2311332011002103-2310100023230001-3021233202013323-0013221212303003-0133131131332300-1203011111120032-3231003301101123-2000013330133110) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.name](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0120031020013213-1113030133110333-2112310020113310-2120103131023302-1302013321001223-3330030103302103-1212233202001130-2323302133332120) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.namespace](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1021130023233012-1210132332231110-3002213212201312-0121210011001321-3131023123200003-2210300132331332-2331231310233110-0310210001021303) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.tenant](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2310211011102302-2223023121303223-2203213132303130-3202102201002102-3131211133113113-2230132012122333-0332310102330210-2101003222223310) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2320321012110212-0332312031122030-3013300233111331-1100130120200322-0230233101233331-2222010100313230-1103233223222013-2031102113133213) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2221300230303120-3112123123031131-2011031031022201-2100230122012013-0111322211231323-1121330230110101-1103300012122302-1330013311233010) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1323102020221011-3110002302203023-2110320222212200-3233000011310023-3001030213232212-0121130213111231-1213311310330223-0303031301103232) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0201311110001120-3223101102213311-2133311000300331-3310203103200020-1232300000302122-1030211113101001-2021210120001003-3011010102032210) |
| `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant` | [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2100012120313313-0122321310030132-0131003022311213-2303123033120022-2222111210302111-1003230203001033-0021130322210331-2021030020203012) |
| `proxy_advertisement.advertise_custom.advertise_where.port` | [proxy_advertisement.advertise_custom.advertise_where.port](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3321320330323133-0021011223112110-0002313300312010-2033120313322033-3020023300110020-1122131220032213-3132333201201320-1112022213132002) |
| `proxy_advertisement.advertise_custom.advertise_where.port_ranges` | [proxy_advertisement.advertise_custom.advertise_where.port_ranges](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2111031010330232-0320320002220222-3233300022103210-2230012203321031-0313221321132203-1320001102102100-1030102003100000-3131221112122100) |
| `proxy_advertisement.advertise_custom.advertise_where.site` | [proxy_advertisement.advertise_custom.advertise_where.site](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1032132212302120-1033321020322003-0123333302220030-0212233323201211-2030111103230132-1120321031200220-0122031320020232-1233021133313113) |
| `proxy_advertisement.advertise_custom.advertise_where.site.ip` | [proxy_advertisement.advertise_custom.advertise_where.site.ip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1311331100011200-0321010203333033-0321122233122202-3300003133032301-1232322022200022-3330202323231030-3110000331031330-2333320123130011) |
| `proxy_advertisement.advertise_custom.advertise_where.site.network` | [proxy_advertisement.advertise_custom.advertise_where.site.network](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2121000013313202-0322100310132301-0003100222330132-2310232230033002-2100120313000333-2132313202320213-1020012110030303-3133031102130011) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site` | [proxy_advertisement.advertise_custom.advertise_where.site.site](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0321131201020030-3232000222232030-2321121301033213-2130001202231023-2322322331131210-0012003300311002-2211330101310333-0013333033331003) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.name` | [proxy_advertisement.advertise_custom.advertise_where.site.site.name](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3320320123200122-0230313301010001-0321002102033302-1122130112300123-3312122313211333-2210313102121212-3203300030000332-2330031021320201) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.site.site.namespace](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3000002132313002-2111302021003322-1121213230020322-1320102010202303-0111020233331321-0200030201013010-1021223130232123-3021102200201310) |
| `proxy_advertisement.advertise_custom.advertise_where.site.site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.site.site.tenant](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1121010132021202-0321002022320232-2222022203012233-2212200032021123-1231231101202012-3023231313013232-2130321232323223-3323131031322300) |
| `proxy_advertisement.advertise_custom.advertise_where.use_default_port` | [proxy_advertisement.advertise_custom.advertise_where.use_default_port](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0032031112220102-1111202020000200-2003020231023031-1000211011330131-2223313023313020-3012220300231320-0232300002032201-2231303131013101) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1132223020322110-3321111012121321-3130022301311123-1030221002130330-0211321003123333-0310233231213100-2300323000033200-1031003021331320) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1201120322223300-3233201301321300-1023111220232233-0310012102101332-2032133200320100-2212101323301231-3323331033120233-0021213031322223) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2333110310022301-1100133133120101-2112210211131231-1333123013300030-1110210232222011-0311300133103122-2211132000310111-2022131322121202) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_v6_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_v6_vip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0003311311132021-2100300331313220-2301020320220301-0103311013231123-1310032212330221-1310323212320322-1100303222022030-0010202133032212) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_vip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2111303010010002-3123330330211132-1110202200320330-3320133302102230-3231133013132220-2223011102221302-2322332213131220-3130313322220230) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0231231003211002-0213012112320221-1030220202313203-0221300133331021-1022120200020012-3032033233112331-3221033322100322-1100212023210122) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.name](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0113021200132133-1313133120221011-1110210000311021-0232200213212210-0313020332210200-2000303222130023-3120300100333223-2332301202223221) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.namespace](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3011102120230122-3030100113220220-3123011233123300-1023021102312321-1210201311023010-0133301200222201-1320330101220133-3201331100010313) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network.tenant](data-sources--bigip_http_proxy--reference--group-002.md#canonical-1212120121300332-0332121320202322-0203022010131310-1313233031333330-1021102201300121-0331102232222331-3213003031222200-3113112122320123) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0323223020130111-2311111313332201-2032032012212131-0102302031313020-1230332113013331-1033001213103213-0323033000110222-3310212220022301) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.network](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1321100312313220-0232330212120230-1313330122022330-3222112023031123-2201212121300331-2120202202132210-2001110322211321-3013102133110020) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0203212213312200-1213301303110312-3203122112033013-1311011201332031-0330022110310021-2213302111013020-3122231002013020-3210300213122333) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.name](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2133213131130210-0110202300300232-2213213020120332-3001330130001000-1102023111212123-3130312003003102-3320022113131001-1231122222230230) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.namespace](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3213132102033000-2032330310332303-2002220111022100-3222231032013221-3102222203310220-2101303210301200-2031320212020123-1222131012100001) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site.tenant](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2201310332223313-1331031201211000-2113310101220020-0013203001103302-1201321302101010-0031330210322232-3120110030210102-3120030303302111) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1122132330121232-1132313130031103-3101320012332120-0221203002210100-2023000232031100-3103301003112101-0320222303023330-1022332223232310) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.ip` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.ip](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0303222303201033-0323200223210132-1012320323021000-0223222131311223-1312303203013102-0101230122232010-2202113021212312-1330322200032111) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.network` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.network](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2130223322111333-3111322112131022-2003330010333310-1003022022202232-2231332323312012-1021102232021212-0113323001200201-1011222232013322) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1100303310300221-1313121023300131-2201030032333231-0310000230030031-1311123203011031-1333332011023112-0113013221000310-3311001321113330) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.name](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1132302322102021-2211132030220312-0112201013010021-2133013122002221-0301331323113110-0000320310302223-3021220123101320-2201030320103110) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.namespace](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3230113001133133-0013333203311312-2113131221232212-3322112122300132-2030322011302102-3023333011220032-2210222201231322-0321110031313000) |
| `proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site.tenant](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0110011202131232-0123222120301122-2022202123023313-2310203021303331-1123300113322132-1123121032111223-2331020230323213-1012100033100303) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3302102031323013-1031023301100102-2030020221113121-1332322010121203-2333330330001313-2223100110003332-3321112320300222-1222012101311032) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0012200233112331-3223031211012323-2331102223220223-1322020003301010-2012103302300332-1110310023233310-1331212311133120-0303300231320202) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.name` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.name](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1223201301030011-1232002321002123-2313000302320031-1200030032310321-1020131123101211-3322021200122220-2230303313331012-2021002111213322) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.namespace](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3331311100213120-0311223332022213-0023022122101011-3111211133030201-0332011010021211-1203232000002030-1300312233101112-2322123003332203) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site.tenant](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1022320212003032-2021131202222131-3012120331333010-0012031330112213-3023333230332212-3201303201023220-3022333120330232-0201020213113033) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2300231330123212-3103010303300101-0213100331032030-3311201123122011-3110131301020021-1313311120013120-2113123033032310-0000233300122230) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.name` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.name](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3131132331120122-3223303301123301-2101232000130133-0110213331001212-1023113003302020-2323320003000001-1113122310320223-3113313331132212) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.namespace` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.namespace](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2320302213133330-2003233320333010-3120322322022011-3100302220212012-0032213221203032-0321231123031230-2003331231121121-1310331311310020) |
| `proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.tenant` | [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site.tenant](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3103322221002010-3112120013320302-0112122121211133-0303332320303230-2220021313200122-2301030113111032-3133020232120111-2212201211220030) |
| `proxy_advertisement.do_not_advertise` | [proxy_advertisement.do_not_advertise](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3031312323231020-0113200313213210-0033130220231123-1200131232031233-2103303311133022-0312121011122032-3321311301333313-0022033223202221) |
| `proxy_config` | [proxy_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1231210333102133-2313202213322022-2310312222033130-2121022032332201-0122120101110200-0120012230031221-1133223213031110-3132012221102322) |
| `proxy_config.domains` | [proxy_config.domains](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3230213210313020-2020303033023122-2323101130120200-1313312022031113-1112113100102300-0022132002300312-3311322232313120-3133303303320000) |
| `proxy_config.http` | [proxy_config.http](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0201322302313330-2111332301120033-2311111103100321-2223001033000122-2303030300200010-0123002001300030-1231120123232113-3331133131313001) |
| `proxy_config.http.dns_volterra_managed` | [proxy_config.http.dns_volterra_managed](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1210201310033332-1003311010320030-3133110110022310-0002213030132210-0202003123220012-1131033011023200-1333202133031030-0022132022202320) |
| `proxy_config.http.port` | [proxy_config.http.port](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1023000321232100-3033123133210311-0033302131012302-3330311301123220-2133120132310030-0233033313132110-2221113122320312-2110022001032332) |
| `proxy_config.http.port_ranges` | [proxy_config.http.port_ranges](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3232003033211322-0311310003213102-2003220133301100-3120331330312313-1010033232123131-0211002112220210-0311000132030102-2222131220211202) |
| `proxy_config.https` | [proxy_config.https](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3131210001331111-1202123021322110-2023023311130330-0230100302301332-3210232200200233-3231011021323203-0223301322310112-2131200112121210) |
| `proxy_config.https.add_hsts` | [proxy_config.https.add_hsts](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1033332002033311-0310122232131012-3000123133001122-1133033130213120-1012202021300310-1113210103231132-3120110302221133-1230033211020032) |
| `proxy_config.https.append_server_name` | [proxy_config.https.append_server_name](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0223203231112032-2102010131131012-0010222220012032-2320121200110231-0201231312120210-3313312210212133-2103233013121220-2322213331032110) |
| `proxy_config.https.coalescing_options` | [proxy_config.https.coalescing_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0223231011323112-3000232123323222-0332233310120310-3001030300330231-3333200112321222-2333101032331223-3233012033000113-1013110120100322) |
| `proxy_config.https.coalescing_options.default_coalescing` | [proxy_config.https.coalescing_options.default_coalescing](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3132130121020200-1300333001210012-2311212022321233-3322322020003031-1330023133132101-0212313130211000-3131112022021112-2110313130020002) |
| `proxy_config.https.coalescing_options.strict_coalescing` | [proxy_config.https.coalescing_options.strict_coalescing](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3102333313000021-0230033211232111-3120020123023330-3113203323133023-1311301301013100-3203302311230302-0300221320311203-0322000303303103) |
| `proxy_config.https.connection_idle_timeout` | [proxy_config.https.connection_idle_timeout](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2233330131102310-0213111020102010-3103212210033110-1323001302233332-3030301321022101-1100131331000023-2301102112310332-2120310102200102) |
| `proxy_config.https.default_header` | [proxy_config.https.default_header](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3311022003010330-0320020132023301-2110032022032111-0123300032111330-3101032303020332-0330110132002333-3130312030302322-3303222033330030) |
| `proxy_config.https.default_loadbalancer` | [proxy_config.https.default_loadbalancer](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2303231011310021-2121212100000201-2231210333303212-0101131301121120-3300202303301302-2030222012131020-3003202301032210-0111002003101313) |
| `proxy_config.https.disable_path_normalize` | [proxy_config.https.disable_path_normalize](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1111030201013103-3030332220030332-1303232300120201-1230223320032332-2013002313202013-1020001333213210-3333010100231210-2001123300333111) |
| `proxy_config.https.enable_path_normalize` | [proxy_config.https.enable_path_normalize](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2231231313011132-3322020313011102-2121111311033203-1203133323122223-2312030121223312-3300012011023103-1100122221111221-3331121110132233) |
| `proxy_config.https.http_protocol_options` | [proxy_config.https.http_protocol_options](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2020232033203000-0122333001010013-1013233013122011-2123211103311130-0331303022202031-2321020020100120-3032203031221001-1332313132211202) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0312330010131033-2113021130013333-3311222321101023-3301101332031012-0223120330321002-0013222321131022-1003303300133320-1110030102223220) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1310032333303232-2300012323233112-2132020103112102-0122213120020001-3331112213011302-0103110211202313-0100020123221012-1031031121033320) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3002202220301331-1023230122223303-1123331023010202-2311332132123212-0301202300121011-2312320323210203-2312122310011313-3000112322000020) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2132310332210102-2230333101302022-1301321113101011-0303313201231311-2232010310031303-2301212233322300-3112211303013012-3031333211021333) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0022233102323321-2202130121021321-3022113221120012-2133100300302121-3312300100323020-2331301011203112-2210113103100301-3331220311022212) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2` | [proxy_config.https.http_protocol_options.http_protocol_enable_v1_v2](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2213131012001100-0113012110132301-2032101031231233-0300102231313021-0033312202320102-0032110130132220-0320333001100211-2312221321320313) |
| `proxy_config.https.http_protocol_options.http_protocol_enable_v2_only` | [proxy_config.https.http_protocol_options.http_protocol_enable_v2_only](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2232132303313112-0200000020322011-2312131010032230-1133220100302323-2313203231132131-1100312111222313-2020013221221200-3212233222133223) |
| `proxy_config.https.http_redirect` | [proxy_config.https.http_redirect](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0200210020210213-3213310120131313-3221320311112110-2100001202031232-2311022013033230-0211323002212202-2131212311012203-1311323331322023) |
| `proxy_config.https.non_default_loadbalancer` | [proxy_config.https.non_default_loadbalancer](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0210223233031120-1000222222333132-2232011202202123-1323202312132210-3220221210321110-2121233222203100-1002020110113313-3320223323010312) |
| `proxy_config.https.pass_through` | [proxy_config.https.pass_through](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0332022220001232-1320033131013103-3110333312203221-3201111230332103-1020312003332212-2100300020002312-2332002321120110-3312010012132321) |
| `proxy_config.https.port` | [proxy_config.https.port](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1101220131221202-1212023323213220-0230133033002030-1330101003300333-2310032112031011-1213110221132113-1113030113321013-0102010112230020) |
| `proxy_config.https.port_ranges` | [proxy_config.https.port_ranges](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1022101111013211-1231330120003213-2212133223010003-0230311311112201-1121323200233003-2302011111230103-0121010203203021-0301122101221311) |
| `proxy_config.https.server_name` | [proxy_config.https.server_name](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1332220233223222-1031003231203133-0201002030200123-0320021221312111-3331330023012203-0332310332313210-2011132312021320-2210220112111222) |
| `proxy_config.https.tls_cert_params` | [proxy_config.https.tls_cert_params](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3122310321332130-3023333001301322-0022230033301010-0131301212333330-1122301332230332-1223002012120301-2023012112110013-2230001312331233) |
| `proxy_config.https.tls_cert_params.certificates` | [proxy_config.https.tls_cert_params.certificates](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0012010113200032-1223311313032332-0112311213223312-3313231133311231-1011023100130310-2232130112302203-1220310022321221-3013131330301021) |
| `proxy_config.https.tls_cert_params.certificates.name` | [proxy_config.https.tls_cert_params.certificates.name](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3113212031330030-0033131021032003-3121120032313332-3130033230202101-1222231220312122-2012331330023023-1003121031121310-1012133132010220) |
| `proxy_config.https.tls_cert_params.certificates.namespace` | [proxy_config.https.tls_cert_params.certificates.namespace](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2300202210323110-2321000122221301-2331100300223320-2032010101023033-1103011233022101-2333131112110021-1302333322130131-1111332030123303) |
| `proxy_config.https.tls_cert_params.certificates.tenant` | [proxy_config.https.tls_cert_params.certificates.tenant](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0210320223023102-3312023001031210-1030223113322213-2200022333222200-1202302303102002-2222121323122113-1100203230102033-2323120132021133) |
| `proxy_config.https.tls_cert_params.no_mtls` | [proxy_config.https.tls_cert_params.no_mtls](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2102303322033200-1222032222111111-3222021312212332-1203120020123222-0203120132001203-0133032000211022-1012212201022013-2302332220120212) |
| `proxy_config.https.tls_cert_params.tls_config` | [proxy_config.https.tls_cert_params.tls_config](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3331311330100132-3321222311203013-0322132311102103-3111101333133032-3022112121202201-1133332222201330-2310300210303020-2210033103123021) |
| `proxy_config.https.tls_cert_params.tls_config.custom_security` | [proxy_config.https.tls_cert_params.tls_config.custom_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0202303133113310-3000002010303332-0013312210203002-2121222322131221-2122320102211002-0022222132233030-3322310113211333-1303120103013122) |
| `proxy_config.https.tls_cert_params.tls_config.custom_security.cipher_suites` | [proxy_config.https.tls_cert_params.tls_config.custom_security.cipher_suites](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2103322232331000-1130011110013132-2210120302100021-0220231022100110-2331011130223100-3102113300330002-1312221131033221-1110230120232001) |
| `proxy_config.https.tls_cert_params.tls_config.custom_security.max_version` | [proxy_config.https.tls_cert_params.tls_config.custom_security.max_version](data-sources--bigip_http_proxy--reference--group-003.md#canonical-1230012312302133-3033300223120113-2233333230333120-3211212030211321-1131010311300111-0032213312220111-2033033300033311-3230113110323301) |
| `proxy_config.https.tls_cert_params.tls_config.custom_security.min_version` | [proxy_config.https.tls_cert_params.tls_config.custom_security.min_version](data-sources--bigip_http_proxy--reference--group-003.md#canonical-3233021320101312-3322002332302312-1103332310202200-0300002232300102-0112232333131330-2302333312133123-1032201012301110-3332200322023011) |
| `proxy_config.https.tls_cert_params.tls_config.default_security` | [proxy_config.https.tls_cert_params.tls_config.default_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2011131020303231-0123323201330332-0100220102332203-2121210113201232-0320230333022103-0303310022103201-1003323223010120-0022110033210301) |
| `proxy_config.https.tls_cert_params.tls_config.low_security` | [proxy_config.https.tls_cert_params.tls_config.low_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-0100201211030102-1220202111320022-2130213102221113-1230132000300222-2231111201020112-1001223011203132-3321022013321111-2213020030331301) |
| `proxy_config.https.tls_cert_params.tls_config.medium_security` | [proxy_config.https.tls_cert_params.tls_config.medium_security](data-sources--bigip_http_proxy--reference--group-003.md#canonical-2210202220021222-0010200321103210-3100212230303022-0212233233102310-2211021233303230-0101011113203123-1213302300221002-2300113312231300) |
| `proxy_config.https.tls_cert_params.use_mtls` | [proxy_config.https.tls_cert_params.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3323333331121130-0101330023022013-2001210012100000-2022000100121200-0310131122310202-0101031010201313-1203001013330301-3133120100203002) |
| `proxy_config.https.tls_cert_params.use_mtls.client_certificate_optional` | [proxy_config.https.tls_cert_params.use_mtls.client_certificate_optional](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0223021201013221-2100222321101211-0212033301233020-3300102302231001-3020001133101233-2200013120303212-0212310202313002-1311213003002001) |
| `proxy_config.https.tls_cert_params.use_mtls.crl` | [proxy_config.https.tls_cert_params.use_mtls.crl](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2013230130231331-3111213220120122-1313330120031121-3122113321103211-3032213332212102-2211011113012213-2011330023021230-1231311103321013) |
| `proxy_config.https.tls_cert_params.use_mtls.crl.name` | [proxy_config.https.tls_cert_params.use_mtls.crl.name](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1000220020200012-1003223023321320-0103022332311222-2032202232021003-3222030213011123-0020020213120103-3110002330030210-1302002311000023) |
| `proxy_config.https.tls_cert_params.use_mtls.crl.namespace` | [proxy_config.https.tls_cert_params.use_mtls.crl.namespace](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2031100130301331-3300031331121033-2200220010131121-3122323222123120-1303201210332213-3121202201032201-2300200301030033-0021221312101001) |
| `proxy_config.https.tls_cert_params.use_mtls.crl.tenant` | [proxy_config.https.tls_cert_params.use_mtls.crl.tenant](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3303213323110123-3232230130022133-0322003010110203-1213330313232020-2121013100320312-3313122301211103-1031003131311232-1003100213220001) |
| `proxy_config.https.tls_cert_params.use_mtls.no_crl` | [proxy_config.https.tls_cert_params.use_mtls.no_crl](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3220230133002012-3023132332222313-3220100102330013-2300120223232302-0020132312130331-1013233322312203-1313301012222031-2021220011013031) |
| `proxy_config.https.tls_cert_params.use_mtls.trusted_ca` | [proxy_config.https.tls_cert_params.use_mtls.trusted_ca](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1320322012231103-1013213030132322-0101012012332221-3102031120102301-0033003121300312-1021302301331323-2311302232320021-2320011202013231) |
| `proxy_config.https.tls_cert_params.use_mtls.trusted_ca.name` | [proxy_config.https.tls_cert_params.use_mtls.trusted_ca.name](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3321312102123103-1121310202312232-3032023022101320-3300103312022211-3112323132131321-0023200103112102-3121113030303321-1120101230302221) |
| `proxy_config.https.tls_cert_params.use_mtls.trusted_ca.namespace` | [proxy_config.https.tls_cert_params.use_mtls.trusted_ca.namespace](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0133201231011233-1110333033023312-0221110313013023-0003201011032021-1112232010112123-1312221202310113-3213020213100002-0331201030101132) |
| `proxy_config.https.tls_cert_params.use_mtls.trusted_ca.tenant` | [proxy_config.https.tls_cert_params.use_mtls.trusted_ca.tenant](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1110030112223112-0310110120121200-3330030101000020-1130311213301203-0213323213001013-1131212121112123-1032231210320302-1023111232323032) |
| `proxy_config.https.tls_cert_params.use_mtls.trusted_ca_url` | [proxy_config.https.tls_cert_params.use_mtls.trusted_ca_url](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3110001310323220-3130221312113321-0010222103233203-2221133032231331-2013020333121313-2033130232023133-2203102323220113-0211111022323322) |
| `proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled` | [proxy_config.https.tls_cert_params.use_mtls.xfcc_disabled](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2330113221112113-1320323320130300-3313132313032302-3312003333210001-2220001302302011-3003022122212333-2121133122130332-0210201001032031) |
| `proxy_config.https.tls_cert_params.use_mtls.xfcc_options` | [proxy_config.https.tls_cert_params.use_mtls.xfcc_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0333001012300132-2333031203202200-2011203013312100-1320310020101130-3201013330100303-0311122213123113-3101332233200123-2223302310111222) |
| `proxy_config.https.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements` | [proxy_config.https.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0131321223303031-0133121131121120-2230320210110233-2202112121130211-3310323101220121-0231121311121001-0200301133333020-0122310320232103) |
| `proxy_config.https.tls_parameters` | [proxy_config.https.tls_parameters](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1201123003111312-3320313130233200-2321331213113100-0003203220311322-1320311011310223-1033012213322111-1232231132122133-0030211221111110) |
| `proxy_config.https.tls_parameters.no_mtls` | [proxy_config.https.tls_parameters.no_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1100013033220213-2233310102033130-2110202321202203-0212132030023300-3002210011323221-1030212023013000-1233010010213232-3022110103121021) |
| `proxy_config.https.tls_parameters.tls_certificates` | [proxy_config.https.tls_parameters.tls_certificates](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3320323100113312-2220312130011303-0123030113323211-0031222103100213-1032213130131310-1222222303020021-3221203123122013-0031331013033103) |
| `proxy_config.https.tls_parameters.tls_certificates.certificate_url` | [proxy_config.https.tls_parameters.tls_certificates.certificate_url](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0303030323131121-1333000132000120-1111313300303203-1033022133313022-3300212333011333-0323130003312211-1023202110212233-0133202001000132) |
| `proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms` | [proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1121011121201310-3022001113102121-1313021322231002-0023013031111212-1122011302123011-1131302130233032-2222311120022111-2013003023101212) |
| `proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms` | [proxy_config.https.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2200020112131100-2031330330022003-1030303013233020-1222233033031010-2232303233121010-1132220102013113-3023303112330023-2231011330123321) |
| `proxy_config.https.tls_parameters.tls_certificates.description_spec` | [proxy_config.https.tls_parameters.tls_certificates.description_spec](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0100230010000020-3010103313122030-3131230202203330-1131102001303213-1320021022133103-2232121013030232-0132212131231212-2100003332323300) |
| `proxy_config.https.tls_parameters.tls_certificates.disable_ocsp_stapling` | [proxy_config.https.tls_parameters.tls_certificates.disable_ocsp_stapling](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3112011131032311-0111312110200132-2202003321313111-2101210131233203-1213033313000333-2132032011010300-1131132021300230-1113230332332201) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key` | [proxy_config.https.tls_parameters.tls_certificates.private_key](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0032223323211231-2113320120213000-0111230213022012-1220010202312311-0313123311113020-3013021103123033-1111221300031323-3232030331233301) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info` | [proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2200023321001131-3311220312221020-1111210312231103-2011132202220311-3120213113122313-2020312010101211-2223220022021331-1300312002130220) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3203210032121330-0230123312002310-3132321111223331-2132203301333223-0000322102001232-2333100313111131-2110320130312331-0230331103121012) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location` | [proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3133203300202213-2010110021031120-2310111103022013-1323330032110333-2212011002120023-3032333113111010-1331223321113011-1100132330230011) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider` | [proxy_config.https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0003033313021112-1122233222001103-2123023031212310-2312121130212102-0112122323100230-3100021033101012-2022331230120120-2323320112301022) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info` | [proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2203023003102211-1300330303212210-1122003002201130-3121000201231202-1130012303321003-0313301130211310-2232130102233210-3023203302023122) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref` | [proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0212003223201121-3213222022213233-3101233023311001-3230213033130132-2300211102020031-3130000231303130-1201202331123231-1131133303033012) |
| `proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info.url` | [proxy_config.https.tls_parameters.tls_certificates.private_key.clear_secret_info.url](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2133231200012200-0102101010102020-3221030200001102-2323110103000120-1010100031101002-0332010003130021-1220011022133310-0020200023312210) |
| `proxy_config.https.tls_parameters.tls_certificates.use_system_defaults` | [proxy_config.https.tls_parameters.tls_certificates.use_system_defaults](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2131012012330020-2201200032231133-2222102001211331-0131002102002102-3033021312133232-2013003331001100-3033000320101203-3030003312221022) |
| `proxy_config.https.tls_parameters.tls_config` | [proxy_config.https.tls_parameters.tls_config](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2121230122113200-0212020330310131-0310333011030132-2120020003233100-2213122010011030-3120303122121221-3023023123230012-1233211002133223) |
| `proxy_config.https.tls_parameters.tls_config.custom_security` | [proxy_config.https.tls_parameters.tls_config.custom_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3122231310302001-2030100011103000-3010312033223302-1100113123132312-1011022121100210-1221222110213221-3020301332233102-0130122203120331) |
| `proxy_config.https.tls_parameters.tls_config.custom_security.cipher_suites` | [proxy_config.https.tls_parameters.tls_config.custom_security.cipher_suites](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3031001212333300-2013213133112030-0332131201101222-1333333211220120-3121010210303133-0003220000111020-0120322013113321-2302123221231201) |
| `proxy_config.https.tls_parameters.tls_config.custom_security.max_version` | [proxy_config.https.tls_parameters.tls_config.custom_security.max_version](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0020212002313020-2303323230131023-1331203132000210-0001322003331313-1320101023232123-3130033320301020-2332303322303022-3101112300302220) |
| `proxy_config.https.tls_parameters.tls_config.custom_security.min_version` | [proxy_config.https.tls_parameters.tls_config.custom_security.min_version](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3321200203313332-3032003010003302-0322331010323230-3212201330231131-0131230201223001-3320021113230012-1202101021031310-2122132232102012) |
| `proxy_config.https.tls_parameters.tls_config.default_security` | [proxy_config.https.tls_parameters.tls_config.default_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2012022031013200-3131321000000311-0031322211310113-0010013021013330-3331320322210131-0320123231201133-3312120321001331-2301230233202320) |
| `proxy_config.https.tls_parameters.tls_config.low_security` | [proxy_config.https.tls_parameters.tls_config.low_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1220231112001133-3313001100211310-1222032222212112-1310102202331033-0112231212113212-0212203020113330-0202202232232102-3110311223112322) |
| `proxy_config.https.tls_parameters.tls_config.medium_security` | [proxy_config.https.tls_parameters.tls_config.medium_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3011212221001302-2002211233102332-3221321301030222-0320210030023333-3203313003301312-2111330333032013-0120230310133201-0230331011312221) |
| `proxy_config.https.tls_parameters.use_mtls` | [proxy_config.https.tls_parameters.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1210122230200233-0020102033331321-2100213011212200-1120102233222210-1232312111302212-3132231320130111-2233302211010130-3030021333331202) |
| `proxy_config.https.tls_parameters.use_mtls.client_certificate_optional` | [proxy_config.https.tls_parameters.use_mtls.client_certificate_optional](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1222221102032102-0301010230213203-2303011321120022-1101221113201102-2132230111113303-0001210203101302-3210222100021002-3320121231111223) |
| `proxy_config.https.tls_parameters.use_mtls.crl` | [proxy_config.https.tls_parameters.use_mtls.crl](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2230022002021313-3300233322301303-2032200013000312-2330132033211031-0033221033030130-1020332003331232-0230001121330011-0231221000230332) |
| `proxy_config.https.tls_parameters.use_mtls.crl.name` | [proxy_config.https.tls_parameters.use_mtls.crl.name](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0020203023211202-3220002003002120-2310022121033330-1313122133320201-1000012210210233-0213122230113233-3213321202132320-2330213301012000) |
| `proxy_config.https.tls_parameters.use_mtls.crl.namespace` | [proxy_config.https.tls_parameters.use_mtls.crl.namespace](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2201210323033331-0013132010202030-1231130302231132-2032200110212120-0201113012203230-1001111200131130-2231033031210103-0001303021202012) |
| `proxy_config.https.tls_parameters.use_mtls.crl.tenant` | [proxy_config.https.tls_parameters.use_mtls.crl.tenant](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2230223213021202-3301210310212002-1221122122103210-2321131130300302-0003131223033113-3211211012313213-2333023013303312-3001020322321301) |
| `proxy_config.https.tls_parameters.use_mtls.no_crl` | [proxy_config.https.tls_parameters.use_mtls.no_crl](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1231121101201303-3003131010302113-0032003033121321-1220003331002032-0032301230233303-2012103131020200-0011333213103311-3120222223101223) |
| `proxy_config.https.tls_parameters.use_mtls.trusted_ca` | [proxy_config.https.tls_parameters.use_mtls.trusted_ca](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3230120200323110-3201303223210202-2210012330010220-3200010210102232-2300332301222220-1202030213011321-2031023233012221-0131011113333200) |
| `proxy_config.https.tls_parameters.use_mtls.trusted_ca.name` | [proxy_config.https.tls_parameters.use_mtls.trusted_ca.name](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1330203203203102-2230031203023130-1320022200121021-0102001232031222-0102223121311010-2023010200020321-3232100323333121-3222021303121122) |
| `proxy_config.https.tls_parameters.use_mtls.trusted_ca.namespace` | [proxy_config.https.tls_parameters.use_mtls.trusted_ca.namespace](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3033302021110310-2302133030033030-2302002103023332-3320232221310122-2311113302231212-1032233113110201-0100010202212223-3220003133011123) |
| `proxy_config.https.tls_parameters.use_mtls.trusted_ca.tenant` | [proxy_config.https.tls_parameters.use_mtls.trusted_ca.tenant](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3222122133121300-0320032003132213-3100203220012032-1110213203000120-1201201011010132-2330000133311111-0100022211120031-3012301113130021) |
| `proxy_config.https.tls_parameters.use_mtls.trusted_ca_url` | [proxy_config.https.tls_parameters.use_mtls.trusted_ca_url](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0203331203230323-2311010210221301-1222010030032211-1020233111321232-3110103023100013-0030203010101120-2132331002022002-1232310222320301) |
| `proxy_config.https.tls_parameters.use_mtls.xfcc_disabled` | [proxy_config.https.tls_parameters.use_mtls.xfcc_disabled](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3011121222000332-0101100213001002-2002033200210003-2112312223031322-0012102320011232-1322102030303020-1121133131130132-1132301201022012) |
| `proxy_config.https.tls_parameters.use_mtls.xfcc_options` | [proxy_config.https.tls_parameters.use_mtls.xfcc_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2330222222231133-3230331301000323-2323301001012310-2111002332001313-3220323200103211-2330322001330200-0111320200021020-1300032110001330) |
| `proxy_config.https.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements` | [proxy_config.https.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2212131331211022-0100321231223020-1101223300331010-0032333231222220-0012002210132301-1103221030320313-2310303232300201-0013002132213030) |
| `proxy_config.https_auto_cert` | [proxy_config.https_auto_cert](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0332331330031223-1311030322032013-0222323322323003-2023312200110121-1133222312231302-3220010201302032-0131122200321302-1310203032310021) |
| `proxy_config.https_auto_cert.add_hsts` | [proxy_config.https_auto_cert.add_hsts](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0101323301013302-1010300003023001-0220021310320031-2023030332213003-0333033221011211-3110323012321311-1222130013010031-2031020212122100) |
| `proxy_config.https_auto_cert.append_server_name` | [proxy_config.https_auto_cert.append_server_name](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1203321231323300-3331011032212001-0202030330003012-0303322322130312-2322203311312231-1030312130000011-3000331010010132-0132003200013330) |
| `proxy_config.https_auto_cert.coalescing_options` | [proxy_config.https_auto_cert.coalescing_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2323231130323012-3210230322131303-2131020020023330-2010232203322211-1223302330003001-1313310012012201-3311112113102330-1312332233123222) |
| `proxy_config.https_auto_cert.coalescing_options.default_coalescing` | [proxy_config.https_auto_cert.coalescing_options.default_coalescing](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2200130103100121-2100303223133210-0333032310003323-3233313331023211-2221220310302200-3203221032100021-2220030013332003-3211101311033033) |
| `proxy_config.https_auto_cert.coalescing_options.strict_coalescing` | [proxy_config.https_auto_cert.coalescing_options.strict_coalescing](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3112120002223131-3032013233302013-1232300130221102-1313131310201013-0222323013223331-3210233333330120-3000130231102030-1112111032310000) |
| `proxy_config.https_auto_cert.connection_idle_timeout` | [proxy_config.https_auto_cert.connection_idle_timeout](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1003203231021222-0333313032131323-3303232231021212-2232302002012130-1203103122131012-0333220033132332-0302233313103003-1011001032311300) |
| `proxy_config.https_auto_cert.default_header` | [proxy_config.https_auto_cert.default_header](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1220221311011023-3020210333333021-2113131331233221-0203100122313022-2112230013331323-1002110332332330-3313130123112130-0030202222200202) |
| `proxy_config.https_auto_cert.default_loadbalancer` | [proxy_config.https_auto_cert.default_loadbalancer](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2322230011133301-3212102133120320-0022102002213201-3200230223231301-2310001331003030-3110323131321221-1103103311313111-3030000310022331) |
| `proxy_config.https_auto_cert.disable_path_normalize` | [proxy_config.https_auto_cert.disable_path_normalize](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3331312223213323-3320203303333303-1222130303122200-0331230031023230-0113231210131310-0120331202022003-0221233022310020-2210101021101200) |
| `proxy_config.https_auto_cert.enable_path_normalize` | [proxy_config.https_auto_cert.enable_path_normalize](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0313012103101132-2312223303001101-1000003001312033-3320213321030033-2231131202213121-2122120221210311-1332102323100322-3332033132301313) |
| `proxy_config.https_auto_cert.http_protocol_options` | [proxy_config.https_auto_cert.http_protocol_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2030203200311103-0133102323320210-0201213133220333-0020301030331220-0003331022210023-2203110200331230-1201112300311021-1202312002302310) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1331332201222002-3312323111030002-2120011202220310-0322330302202103-0312112130220220-1300311332302222-0222013121131231-2211130120301010) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0003311002101100-1112112231020212-1022212122032000-3112322331022233-2311322222102220-3331331001110222-2023221223112330-2121322022233002) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1213220132133022-1002103301110233-0111221200103023-2132212230123013-3222133222202130-1311231203310010-0111331312011013-0333102333003020) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2211132310132032-3210302312232303-2111120201120310-2110221031020002-0003323021022230-2302030222300022-0020322322210132-0303322023213032) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0303100111100220-1333222222210103-0322123122232301-0003101101231223-2020130210113323-2301121121300220-2023323003020222-3233113333120301) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1030102100121011-0012313131221311-2031323100010100-2303312101120122-2032030013033033-0331212323320203-2031333302101222-3330033203113300) |
| `proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only` | [proxy_config.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3300322133202300-2323100023112013-1011223201231202-3012011200223203-1102232032300101-1332312201212322-0200301112122030-3302330023302201) |
| `proxy_config.https_auto_cert.http_redirect` | [proxy_config.https_auto_cert.http_redirect](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3203312331212003-2300330222012020-0030232232022301-3211311033122300-0021103010012303-1131331202101032-1113220010122202-0320300131003132) |
| `proxy_config.https_auto_cert.no_mtls` | [proxy_config.https_auto_cert.no_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2330213311302231-3100230321010321-3322302022301123-2320221211201113-2220033013012303-0213130013021130-1302021320113122-0133232310311020) |
| `proxy_config.https_auto_cert.non_default_loadbalancer` | [proxy_config.https_auto_cert.non_default_loadbalancer](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2312110211233330-1111110100332323-1330333030312102-0032001312103331-1110002032230222-0202202133203001-0331231133311221-3303102222203313) |
| `proxy_config.https_auto_cert.pass_through` | [proxy_config.https_auto_cert.pass_through](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3012320311120332-1323300231031113-1122203313103032-2223223202311112-0022323113312133-1033221032323132-0300331112021322-3302211223302133) |
| `proxy_config.https_auto_cert.port` | [proxy_config.https_auto_cert.port](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2331203313130010-2313112233233323-2133212012233133-1131113202323321-1302302231312313-1032030010312012-2322210330311321-1320033130131200) |
| `proxy_config.https_auto_cert.port_ranges` | [proxy_config.https_auto_cert.port_ranges](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3213212202202300-0212020001210000-2223211300031100-2313000101013332-3330201111233103-2012033212202132-1321021131012332-0301321111213033) |
| `proxy_config.https_auto_cert.server_name` | [proxy_config.https_auto_cert.server_name](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3003332330013122-0121020201022010-1313321102130303-0003120130002232-0003201002130220-1003312220213232-0311333211210230-0022100001330011) |
| `proxy_config.https_auto_cert.tls_config` | [proxy_config.https_auto_cert.tls_config](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2203113220010212-3033312212222231-3030031010323303-0300130201113211-1213023020013202-1021320131313003-1123021110133103-0222202300112121) |
| `proxy_config.https_auto_cert.tls_config.custom_security` | [proxy_config.https_auto_cert.tls_config.custom_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3001333321120210-2332221111031121-1023011233212011-2011302311220001-0111020131102301-3020021302323110-2020200333322330-3001322213013310) |
| `proxy_config.https_auto_cert.tls_config.custom_security.cipher_suites` | [proxy_config.https_auto_cert.tls_config.custom_security.cipher_suites](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1222331312321333-3320131122110231-2133110303103211-3010332021102321-3313201032220230-1201120021113232-0010131032221120-0132212220133001) |
| `proxy_config.https_auto_cert.tls_config.custom_security.max_version` | [proxy_config.https_auto_cert.tls_config.custom_security.max_version](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3332201301220022-1021020210122221-1132223303332131-0332120102311302-3333023203122002-3230010330033320-2001101010100012-3002021011123020) |
| `proxy_config.https_auto_cert.tls_config.custom_security.min_version` | [proxy_config.https_auto_cert.tls_config.custom_security.min_version](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1221333111022133-1022311032122131-1030232203331120-1220022011021102-2022302233020333-1130102321203211-2312030322103123-3210202301232310) |
| `proxy_config.https_auto_cert.tls_config.default_security` | [proxy_config.https_auto_cert.tls_config.default_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1130311323123333-3312300113330031-2021101101210033-3012221101100113-3211013031122301-1101003131212121-3103131100321303-0232312103303001) |
| `proxy_config.https_auto_cert.tls_config.low_security` | [proxy_config.https_auto_cert.tls_config.low_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1101303112301021-3003211301010330-1331211211103300-1312221123010220-1123222332332122-0330012110022103-3223320133212100-2021210321121021) |
| `proxy_config.https_auto_cert.tls_config.medium_security` | [proxy_config.https_auto_cert.tls_config.medium_security](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2123133032132331-3121020200133302-3302122031201012-2032303132033312-0200312103233303-0112333100321030-0222123100023210-2113320213301110) |
| `proxy_config.https_auto_cert.use_mtls` | [proxy_config.https_auto_cert.use_mtls](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3100112110323131-2322332323120032-2300211010031102-2000232020301131-2121021022120332-0330023300120121-1232010210001300-2132000113320332) |
| `proxy_config.https_auto_cert.use_mtls.client_certificate_optional` | [proxy_config.https_auto_cert.use_mtls.client_certificate_optional](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2303031233311231-2200311010130121-1131131000032033-2110020320303110-0132101101323320-0120003330303132-1320101120130002-0113102232120320) |
| `proxy_config.https_auto_cert.use_mtls.crl` | [proxy_config.https_auto_cert.use_mtls.crl](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2230220111033113-2111323130131011-2211003213102120-0333230030303001-1112330011122010-2300001321200212-0112202303123022-3100003220231113) |
| `proxy_config.https_auto_cert.use_mtls.crl.name` | [proxy_config.https_auto_cert.use_mtls.crl.name](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0300300001310211-0223302321020330-3130203110122012-1012220020332232-3012031113031300-2130323113010233-3211031320233020-1112313002220021) |
| `proxy_config.https_auto_cert.use_mtls.crl.namespace` | [proxy_config.https_auto_cert.use_mtls.crl.namespace](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0000222033002333-1200323302013222-0220313032232320-1203030113113111-3131002130030033-0323200202200301-0211011312112112-1232222100122212) |
| `proxy_config.https_auto_cert.use_mtls.crl.tenant` | [proxy_config.https_auto_cert.use_mtls.crl.tenant](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1210130312000102-0201213020301100-3321211303002202-1030031220230010-1023130310133302-0212123221110133-3303200110310230-2101032013222232) |
| `proxy_config.https_auto_cert.use_mtls.no_crl` | [proxy_config.https_auto_cert.use_mtls.no_crl](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3001321301202301-0032033230313231-1220013022033001-3120332212120311-1111332201322111-1212202302232033-1233113321122111-2132023310021322) |
| `proxy_config.https_auto_cert.use_mtls.trusted_ca` | [proxy_config.https_auto_cert.use_mtls.trusted_ca](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2000202221311312-0023123202220111-2330023030332203-0211131310322113-1202213100003130-2003103030033120-0112333330023023-0002231230020130) |
| `proxy_config.https_auto_cert.use_mtls.trusted_ca.name` | [proxy_config.https_auto_cert.use_mtls.trusted_ca.name](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2100012013213033-1123120032300232-3000123112232012-1311122121223213-1120010012000011-3210113103112332-0021002101312021-0123203231220201) |
| `proxy_config.https_auto_cert.use_mtls.trusted_ca.namespace` | [proxy_config.https_auto_cert.use_mtls.trusted_ca.namespace](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3323201320002212-2101211001123202-2222112111103102-3211100102100001-3201133102231032-3122222101121022-3313011032102321-3111223023321112) |
| `proxy_config.https_auto_cert.use_mtls.trusted_ca.tenant` | [proxy_config.https_auto_cert.use_mtls.trusted_ca.tenant](data-sources--bigip_http_proxy--reference--group-004.md#canonical-0000300110313023-3110122232233002-3031222220122221-0212110330101222-3203032223223230-2001320000312010-0332200222322123-1222213231010101) |
| `proxy_config.https_auto_cert.use_mtls.trusted_ca_url` | [proxy_config.https_auto_cert.use_mtls.trusted_ca_url](data-sources--bigip_http_proxy--reference--group-004.md#canonical-1030320220010232-3333301023100212-0122212210330310-3131301312300210-2312122020022010-3110033122333220-0020332233021123-2113001100031111) |
| `proxy_config.https_auto_cert.use_mtls.xfcc_disabled` | [proxy_config.https_auto_cert.use_mtls.xfcc_disabled](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3300221003232331-3211111330231010-1133102021321221-3113233120221120-3333210120021310-2100333220120013-3013033123311221-1030103223310033) |
| `proxy_config.https_auto_cert.use_mtls.xfcc_options` | [proxy_config.https_auto_cert.use_mtls.xfcc_options](data-sources--bigip_http_proxy--reference--group-004.md#canonical-2033220231110220-3301013222133312-2233001222213200-1013323223223103-3220122113013101-1213111133212201-2020132211221323-1231231210223322) |
| `proxy_config.https_auto_cert.use_mtls.xfcc_options.xfcc_header_elements` | [proxy_config.https_auto_cert.use_mtls.xfcc_options.xfcc_header_elements](data-sources--bigip_http_proxy--reference--group-004.md#canonical-3102012303001312-1122213203200321-1121003003312322-1310032330230232-0303010320133131-0320110222223321-0131031213323210-1320301203200221) |

<a id="canonical-0211003021332023-0013000312333003-0331301233132013-2000113303313022-0231112311002001-1300321212312100-1303030031223022-3332202130110223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advanced_profile` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- advanced_profile

<a id="canonical-1232021002110000-1312021212200231-0100300321102031-3032103331120332-1232331022211012-1010302112200103-0301220023311130-0312211200203310"></a>

Type: `"single"`. Computed.

This defines various advanced Profile OPTIONS for a Loadbalancer.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"disable\",\"enable_default_profile\"]"
}
```

<a id="canonical-1021103221231122-1031011201000221-0130001021201322-2132120013021130-1300211212203000-3000210100213213-2032012312300002-1332201000201032"></a>

### Direct properties for `advanced_profile`

- [disable_spec](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2011321010130202-1210133331312333-1322013013210331-2213202202000132-1323230201013120-2300011130131000-1011320113313121-2002110130032021): complete subsection reference.

- [enable_default_profile](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3021130100131212-0233312231332030-1022303100330222-2111132233301233-1310130003122022-1321302122313322-3203330313202231-3222200200202000): complete subsection reference.

<a id="canonical-2011321010130202-1210133331312333-1322013013210331-2213202202000132-1323230201013120-2300011130131000-1011320113313121-2002110130032021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advanced_profile.disable_spec` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [advanced_profile](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0211003021332023-0013000312333003-0331301233132013-2000113303313022-0231112311002001-1300321212312100-1303030031223022-3332202130110223)
- advanced_profile.disable_spec

<a id="canonical-3021130012210331-0132101331131103-3111010021000021-2300232303101102-3223211303320110-2201232301113010-2011021333332010-0331310203033331"></a>

Type: `["object", {}]`. Computed.

Enable this option

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3021130100131212-0233312231332030-1022303100330222-2111132233301233-1310130003122022-1321302122313322-3203330313202231-3222200200202000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `advanced_profile.enable_default_profile` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [advanced_profile](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0211003021332023-0013000312333003-0331301233132013-2000113303313022-0231112311002001-1300321212312100-1303030031223022-3332202130110223)
- advanced_profile.enable_default_profile

<a id="canonical-3311013220231033-3111213113301122-2000233031303302-0011232223301303-1030221331202330-2012133022101020-1000201210223012-3222100313321200"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable default profile.

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

<a id="canonical-0202231310212321-0320003212210123-2220233331201313-2321032221213302-3030203022233130-3201133323102200-2331221232230102-3010231100102222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_profile` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- ddos_profile

<a id="canonical-3301030302312122-2222031033132311-2233221120231211-3222213031313210-0212311333000132-0010333312312230-1131011113012012-3012021133230013"></a>

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

<a id="canonical-0100223203302001-1210112111303131-1001222213220010-2032111223002320-0112033000113033-0301000023202312-2321322110120003-3020033110013333"></a>

### Direct properties for `ddos_profile`

- [disable_ddos_mitigation](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1002113022003333-3221032123221312-0110030023033003-2313131000131003-0112113211003132-3003231201110130-3221200231132311-3112033113110213): complete subsection reference.

- [enable_ddos_mitigation](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2221220302133303-1211012100211002-2100203311010102-0301121023312332-0333313310002322-0101312310030021-1033302011102201-1023320321232233): complete subsection reference.

<a id="canonical-1002113022003333-3221032123221312-0110030023033003-2313131000131003-0112113211003132-3003231201110130-3221200231132311-3112033113110213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_profile.disable_ddos_mitigation` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [ddos_profile](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0202231310212321-0320003212210123-2220233331201313-2321032221213302-3030203022233130-3201133323102200-2331221232230102-3010231100102222)
- ddos_profile.disable_ddos_mitigation

<a id="canonical-0320311121132203-0313213102331230-0131302210021021-3311120022300122-2212330002220011-3020331212231231-1120131220300311-3213300222121233"></a>

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

<a id="canonical-2221220302133303-1211012100211002-2100203311010102-0301121023312332-0333313310002322-0101312310030021-1033302011102201-1023320321232233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ddos_profile.enable_ddos_mitigation` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [ddos_profile](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0202231310212321-0320003212210123-2220233331201313-2321032221213302-3030203022233130-3201133323102200-2331221232230102-3010231100102222)
- ddos_profile.enable_ddos_mitigation

<a id="canonical-3213231312201031-3312022121302211-3003102201302000-1003332133201202-0030223112003303-2132332303201000-2131330210130223-1222223201001122"></a>

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

<a id="canonical-0122030321111211-2222212201112202-1210111200130212-1030302210311120-3100110333000313-3320233100320332-2010120111203312-1031100312332303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `irules` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- irules

<a id="canonical-3230220120322103-3032030022322332-2111033130130321-3133321120333020-0323102131112111-2230012103201133-0012232333030000-2232101000120123"></a>

Type: `"single"`. Computed.

IRules Configuration for downstream connections.

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

<a id="canonical-0222312131013201-3201211213330300-1133230300203223-2333333213210130-1102032122323000-0030110103112200-0123300001001231-2212003301310032"></a>

### Direct properties for `irules`

- [irules](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1022101011310213-0320013022123112-2100010030302303-2313022310001031-1321030200213230-3312020110313211-0033100322332103-0011111101230233): complete subsection reference.

<a id="canonical-1022101011310213-0320013022123112-2100010030302303-2313022310001031-1321030200213230-3312020110313211-0033100322332103-0011111101230233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `irules.irules` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [irules](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0122030321111211-2222212201112202-1210111200130212-1030302210311120-3100110333000313-3320233100320332-2010120111203312-1031100312332303)
- irules.irules

<a id="canonical-0021310310131000-0221202202130232-2122202112123322-2030221302021132-1022223020020122-2001010122330221-1100320120213300-0131132232200212"></a>

Type: `"list"`. Computed.

OPTIONS for attaching iRules to BIG-IP HTTP Proxy.

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

<a id="canonical-1211201322220133-0303200113001011-3022132133130111-3133021012321032-2000001322113123-0302232313100331-2021132111113012-3021011300333133"></a>

### Direct properties for `irules.irules`

<a id="canonical-2110212020321323-3333231302120132-0233130012130313-1122032313212101-1331103203333222-1200331332001202-0020212001301001-1001213103013131"></a>

#### `irules.irules.name` property

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

<a id="canonical-3333302313122103-2211330333130003-2333100213022302-3101033011302133-0202013332012002-3003001102301303-2220133230131223-1000311131221002"></a>

<a id="canonical-3323312030122331-0233310111313330-2323202003310010-2023102331130311-3223003102011303-3012102223120331-2123323201203303-2202230213012233"></a>

#### `irules.irules.namespace` property

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

<a id="canonical-2201002031202320-3110230103310323-0130311220210013-2310330303230232-0321332221000330-3013213131331031-1331221113010121-3323000013133202"></a>

<a id="canonical-2321232011030331-1200032210301301-2210021003000000-2233132322230131-0131011221133223-1031112300222123-1313203231300032-3302022211122130"></a>

#### `irules.irules.tenant` property

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

<a id="canonical-0032310323023122-2220201111120120-3332302022102211-3213221130321010-2302211131303223-2203232003221211-2033103110122123-2130230031113001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `lb_algorithm` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- lb_algorithm

<a id="canonical-1031230221331200-1313130100120310-3023300003233212-3121113220003201-3222332021232030-2133200301100211-0302132310023202-2131230132323312"></a>

Type: `"single"`. Computed.

Configuration parameter for lb algorithm.

Additional upstream details:

Load Balancing Algorithm Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-lb_algorithm_choice": "[\"round_robin\"]"
}
```

<a id="canonical-3101303010000310-1200301301200303-0210121310130101-1302001033113120-0320103302213003-1223222032310001-1120100301112221-3111130210030011"></a>

### Direct properties for `lb_algorithm`

- [round_robin](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2222132123102321-0310120233101311-3223132232102333-3023310231010310-0231013311302133-1200323001120021-3001230322330010-2111331112100202): complete subsection reference.

<a id="canonical-2222132123102321-0310120233101311-3223132232102333-3023310231010310-0231013311302133-1200323001120021-3001230322330010-2111331112100202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `lb_algorithm.round_robin` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [lb_algorithm](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032310323023122-2220201111120120-3332302022102211-3213221130321010-2302211131303223-2203232003221211-2033103110122123-2130230031113001)
- lb_algorithm.round_robin

<a id="canonical-2202111331133010-2021130131223121-2200211132230120-3000001302123330-3313001131222130-2122302331100031-0210332203000231-1210232110011231"></a>

Type: `"single"`. Computed.

Configuration parameter for round robin.

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

<a id="canonical-0231332111103211-2012003310023331-2332323010101232-0330133310111212-2010210001000022-3233033013001220-3131212121202013-0003211311200232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- origin_pools

<a id="canonical-3121202312031132-1020131321031032-1013111002322033-3132203333121112-1002322221210303-1103202200030311-2221310012303110-1032133321131001"></a>

Type: `"single"`. Computed.

Configuration parameter for origin pools.

Additional upstream details:

List of Origin Pools.

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

<a id="canonical-2310310113110020-3131110000333030-1202230330120122-0132103002002002-3332103333323202-0121113211012313-3022123221311112-2323131031011133"></a>

### Direct properties for `origin_pools`

- [pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3313313031230022-2312333302000103-3231103120330212-0100310323113313-3223321221332123-0201100310110123-3021300111323110-0211121122030022): complete subsection reference.

<a id="canonical-3313313031230022-2312333302000103-3231103120330212-0100310323113313-3223321221332123-0201100310110123-3021300111323110-0211121122030022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0231332111103211-2012003310023331-2332323010101232-0330133310111212-2010210001000022-3233033013001220-3131212121202013-0003211311200232)
- origin_pools.pools

<a id="canonical-1200323311003121-1322221220032011-3011022131300231-0220110130311033-1213102123123013-3220312123302033-2010223300311300-1213301021122033"></a>

Type: `"list"`. Computed.

Origin Pools. List of Origin Pools.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3312202202023013-2221311003230301-1100201000220210-2121320330230110-1012102131332200-0002122310201113-2321021122222003-1310122332330132"></a>

### Direct properties for `origin_pools.pools`

<a id="canonical-0100213132033302-1330302001101110-1120333111331002-3012233331020210-1103031333210311-3333332230221202-0202202033302003-0202103132203121"></a>

#### `origin_pools.pools.name` property

Type: `"string"`. Computed.

Name. Name of the origin pool.

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

- [origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2101133002330001-2330300231031122-0101033113031001-2030032222002131-2021022000200232-2203330321023113-0231310312021020-0313202222000031): complete subsection reference.

<a id="canonical-2123032312331103-2000031232321322-1312012320110113-2301330313322232-3210133202320202-2301100320201001-0201122121223132-1113233000232300"></a>

<a id="canonical-3133331313300222-0130130233112312-0101033210130123-2212013023223220-0200121220312232-3210033320131231-3113033322332112-1303101030001332"></a>

#### `origin_pools.pools.priority` property

Type: `"number"`. Computed.

Priority of this origin pool, valid only with multiple origin pools. Value of 0 will make the pool
as lowest priority origin pool. When active origin pool is not available, lower priority origin
pools are made active as per the increasing priority.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-1101221200022232-0213212101011100-2232102222123010-2223132211101011-3202011131331310-1032020012233100-1313212201330101-1110331121323113"></a>

<a id="canonical-0323021130003023-2033222312033321-2031303120122301-1211003330003222-1301323120033312-0100133010221110-2033311122000303-2003131303110020"></a>

#### `origin_pools.pools.weight` property

Type: `"number"`. Computed.

Weight of this origin pool, valid only with multiple origin pools. Value of 0 will disable the pool.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "load-balancing",
    "constraintType": "number",
    "maximum": 100,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2101133002330001-2330300231031122-0101033113031001-2030032222002131-2021022000200232-2203330321023113-0231310312021020-0313202222000031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0231332111103211-2012003310023331-2332323010101232-0330133310111212-2010210001000022-3233033013001220-3131212121202013-0003211311200232)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3313313031230022-2312333302000103-3231103120330212-0100310323113313-3223321221332123-0201100310110123-3021300111323110-0211121122030022)
- origin_pools.pools.origin_servers

<a id="canonical-1030300033303001-1023233112101101-1222031112331221-2333111022110121-3100030210230122-2013332332111330-3120123210101331-0120333031203113"></a>

Type: `"single"`. Computed.

List of origin Servers for the BIG-IP HTTP Proxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"automatic_port\",\"lb_port\",\"port\"]"
}
```

<a id="canonical-1310323230013132-0120123032301031-2233312010223012-1213022321300211-2300032333110131-2323223311212331-2313201013233220-2000310013303122"></a>

### Direct properties for `origin_pools.pools.origin_servers`

- [automatic_port](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0121021332201130-2333130222031301-0210122232102232-2021200332301012-3320113011132130-1103310303232221-3113221000033012-2331322103112010): complete subsection reference.

- [health_checks](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1132110223001331-0120211322102202-1000300231323121-1202113231231220-3201312203121123-3030323210021221-0103103303233321-3201232111232022): complete subsection reference.

- [lb_port](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2133130220003230-2212203332132200-2211300012110111-0030222233322100-3110013300123321-2131313332132312-2312022000222213-3301101111003323): complete subsection reference.

- [origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0311113110223030-0220031310233012-1333001103301303-3011333001300332-0201122233121333-2312110211200103-1333302013013131-2312232122010130): complete subsection reference.

<a id="canonical-1212203013123032-2230020310102220-3201032203301121-2020002332132010-2310130312311112-2330321322031012-3322130111103203-0120122030122101"></a>

<a id="canonical-2321001310101221-3203003112001303-1111313230121200-3120210113223122-2313031203330233-3111232200113230-1110032310023202-1220003120012133"></a>

#### `origin_pools.pools.origin_servers.port` property

Type: `"number"`. Computed.

Exclusive with \[automatic\_port lb\_port\] Endpoint service is available on this port.

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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0121021332201130-2333130222031301-0210122232102232-2021200332301012-3320113011132130-1103310303232221-3113221000033012-2331322103112010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.automatic_port` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0231332111103211-2012003310023331-2332323010101232-0330133310111212-2010210001000022-3233033013001220-3131212121202013-0003211311200232)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3313313031230022-2312333302000103-3231103120330212-0100310323113313-3223321221332123-0201100310110123-3021300111323110-0211121122030022)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2101133002330001-2330300231031122-0101033113031001-2030032222002131-2021022000200232-2203330321023113-0231310312021020-0313202222000031)
- origin_pools.pools.origin_servers.automatic_port

<a id="canonical-2321111210202201-2023132322211103-1033131213100312-3111020332122102-1320100011021100-2323002320320203-2013101012300111-3330020113300210"></a>

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

<a id="canonical-1132110223001331-0120211322102202-1000300231323121-1202113231231220-3201312203121123-3030323210021221-0103103303233321-3201232111232022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.health_checks` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0231332111103211-2012003310023331-2332323010101232-0330133310111212-2010210001000022-3233033013001220-3131212121202013-0003211311200232)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3313313031230022-2312333302000103-3231103120330212-0100310323113313-3223321221332123-0201100310110123-3021300111323110-0211121122030022)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2101133002330001-2330300231031122-0101033113031001-2030032222002131-2021022000200232-2203330321023113-0231310312021020-0313202222000031)
- origin_pools.pools.origin_servers.health_checks

<a id="canonical-1331211032200131-1000332003312321-2300313023002213-1121310130013230-0200010031221313-2230330302130213-0021330202223212-0202012201013130"></a>

Type: `"single"`. Computed.

Configuration parameter for health checks.

Additional upstream details:

Origin Server Health Checks.

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

<a id="canonical-1220012210130200-3113031002332031-0013210003202023-0221001201123320-2223131330132130-0302010213132322-2330231303312020-2233322203110112"></a>

### Direct properties for `origin_pools.pools.origin_servers.health_checks`

- [health_check](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1030213012310332-0113020232113013-0303131033103310-3221201312201201-2013322011000013-1313200231102201-1122313311303020-2101003000322312): complete subsection reference.

<a id="canonical-0111031230130022-0332002222000011-3012302311002131-3301020133310332-3313210102210303-0302013000001222-2201130202130330-0212000301111313"></a>

<a id="canonical-1011032103310301-1331213031301322-3010310322121330-3313001312303210-1331020100030032-2331100312333023-0302020233331220-0120222111302230"></a>

#### `origin_pools.pools.origin_servers.health_checks.healthy_threshold` property

Type: `"number"`. Computed.

Number of successful responses before declaring healthy. In other words, this is the number of
healthy health checks required before a host is marked healthy. Note that during startup, only a
single successful health check is required to mark a host healthy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-2322101201101201-1100130221100312-3321201133201023-0323222030121010-1230133111313023-1321223100031110-0313102232312033-2210221131110301"></a>

<a id="canonical-3210213333313003-0132120232213233-3201113012110311-1301233031020200-1130230301310122-0033100331223133-0013301032122112-0113323211132210"></a>

#### `origin_pools.pools.origin_servers.health_checks.interval` property

Type: `"number"`. Computed.

Time interval in seconds between two health check requests.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-2021202332131331-1113121330103103-3133313232213021-1331100320032011-1220203232133333-2030302032312113-3010111211331211-1233323123102200"></a>

<a id="canonical-2001122233010113-3101112030013220-1130220220101202-1123100311020120-3033010202310330-0133230223002311-0330030012212023-2032003332203211"></a>

#### `origin_pools.pools.origin_servers.health_checks.timeout` property

Type: `"number"`. Computed.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "600"
  }
}
```

<a id="canonical-1312030133313030-2332331011031130-1032212311000133-3321303301220032-2213100003011202-1121330130301021-0312233211221101-1302212111100111"></a>

<a id="canonical-2120033012333021-0301100232013232-0010333300103112-1102102222230103-1122133202310330-3122200203313310-0110002003011222-0000220320321130"></a>

#### `origin_pools.pools.origin_servers.health_checks.unhealthy_threshold` property

Type: `"number"`. Computed.

Number of failed responses before declaring unhealthy. In other words, this is the number of
unhealthy health checks required before a host is marked unhealthy. Note that for HTTP health check
if a host responds with 503 this threshold is ignored and the host is considered unhealthy
immediately.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-1030213012310332-0113020232113013-0303131033103310-3221201312201201-2013322011000013-1313200231102201-1122313311303020-2101003000322312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.health_checks.health_check` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0231332111103211-2012003310023331-2332323010101232-0330133310111212-2010210001000022-3233033013001220-3131212121202013-0003211311200232)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3313313031230022-2312333302000103-3231103120330212-0100310323113313-3223321221332123-0201100310110123-3021300111323110-0211121122030022)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2101133002330001-2330300231031122-0101033113031001-2030032222002131-2021022000200232-2203330321023113-0231310312021020-0313202222000031)
- [origin_pools.pools.origin_servers.health_checks](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1132110223001331-0120211322102202-1000300231323121-1202113231231220-3201312203121123-3030323210021221-0103103303233321-3201232111232022)
- origin_pools.pools.origin_servers.health_checks.health_check

<a id="canonical-3320322030231333-1302211102023202-2033201212313022-2211300130031233-0123311100211330-2131033111020033-3311302132103313-2002021110021033"></a>

Type: `"list"`. Computed.

List of Health Checks. List of Health Checks.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minItems": 0,
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "0",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "0",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3122113121312321-0311202101210232-2101322331120110-3130213100310110-2020020213322303-1112020311122213-0321020123000220-1310130222333021"></a>

### Direct properties for `origin_pools.pools.origin_servers.health_checks.health_check`

- [icmp_health_check](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1103331213020200-0312202130130201-2230233033210023-2311033112220210-3011110030320332-3110222212201333-2332031333020200-1000122313233213): complete subsection reference.

- [tcp_health_check](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3032210133202030-2201100123023021-2102213122012212-3120200012030330-0232323313110103-0031022120133112-1223211222001222-2020112103020203): complete subsection reference.

<a id="canonical-1103331213020200-0312202130130201-2230233033210023-2311033112220210-3011110030320332-3110222212201333-2332031333020200-1000122313233213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.health_checks.health_check.icmp_health_check` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0231332111103211-2012003310023331-2332323010101232-0330133310111212-2010210001000022-3233033013001220-3131212121202013-0003211311200232)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3313313031230022-2312333302000103-3231103120330212-0100310323113313-3223321221332123-0201100310110123-3021300111323110-0211121122030022)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2101133002330001-2330300231031122-0101033113031001-2030032222002131-2021022000200232-2203330321023113-0231310312021020-0313202222000031)
- [origin_pools.pools.origin_servers.health_checks](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1132110223001331-0120211322102202-1000300231323121-1202113231231220-3201312203121123-3030323210021221-0103103303233321-3201232111232022)
- [origin_pools.pools.origin_servers.health_checks.health_check](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1030213012310332-0113020232113013-0303131033103310-3221201312201201-2013322011000013-1313200231102201-1122313311303020-2101003000322312)
- origin_pools.pools.origin_servers.health_checks.health_check.icmp_health_check

<a id="canonical-1012002023222010-3102323321100213-1003002131110130-0301121011302103-2233032301230123-1002112311323303-0030210221313311-3112200030321030"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for icmp health check.

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

<a id="canonical-3032210133202030-2201100123023021-2102213122012212-3120200012030330-0232323313110103-0031022120133112-1223211222001222-2020112103020203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0231332111103211-2012003310023331-2332323010101232-0330133310111212-2010210001000022-3233033013001220-3131212121202013-0003211311200232)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3313313031230022-2312333302000103-3231103120330212-0100310323113313-3223321221332123-0201100310110123-3021300111323110-0211121122030022)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2101133002330001-2330300231031122-0101033113031001-2030032222002131-2021022000200232-2203330321023113-0231310312021020-0313202222000031)
- [origin_pools.pools.origin_servers.health_checks](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1132110223001331-0120211322102202-1000300231323121-1202113231231220-3201312203121123-3030323210021221-0103103303233321-3201232111232022)
- [origin_pools.pools.origin_servers.health_checks.health_check](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1030213012310332-0113020232113013-0303131033103310-3221201312201201-2013322011000013-1313200231102201-1122313311303020-2101003000322312)
- origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check

<a id="canonical-0201032232322221-0321013213033021-2110220213022231-2310230233010232-0112211112302232-1113130132020032-1001031111220013-0230131113023333"></a>

Type: `"single"`. Computed.

Monitor reports healthy status if UDP connection is successful and response payload matches expected
response pattern.

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

<a id="canonical-3030112231110112-3300112230002131-2332301020022120-3313100302020222-1011002030002022-3201320033021303-0001200321133330-0032021020201130"></a>

### Direct properties for `origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check`

<a id="canonical-3030312310330121-1300000231320300-1301320020103131-3203233311112310-3312120303012000-3320103233312110-1321310223221210-2323132301132010"></a>

#### `origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check.expected_response` property

Type: `"string"`. Computed.

Specifies a regular expression pattern which will be matched against response payload.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-0002103322123231-1012230322110203-3230231323220000-2003322110022133-3331101211003033-3303310333210022-0123120213313003-3301100022012330"></a>

<a id="canonical-1010123302223130-3000201230301102-3322002130230201-0000030202323020-2322133200330213-0120011202110130-0313133000210332-3011132122222331"></a>

#### `origin_pools.pools.origin_servers.health_checks.health_check.tcp_health_check.send_payload` property

Type: `"string"`. Computed.

Send string. Text string sent in the request.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
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
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-2133130220003230-2212203332132200-2211300012110111-0030222233322100-3110013300123321-2131313332132312-2312022000222213-3301101111003323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.lb_port` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0231332111103211-2012003310023331-2332323010101232-0330133310111212-2010210001000022-3233033013001220-3131212121202013-0003211311200232)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3313313031230022-2312333302000103-3231103120330212-0100310323113313-3223321221332123-0201100310110123-3021300111323110-0211121122030022)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2101133002330001-2330300231031122-0101033113031001-2030032222002131-2021022000200232-2203330321023113-0231310312021020-0313202222000031)
- origin_pools.pools.origin_servers.lb_port

<a id="canonical-2223303302123221-0301022330000013-3202232221023020-3320110113212120-1321100010032020-2212121130003222-0222031223223323-3033113303300323"></a>

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

<a id="canonical-0311113110223030-0220031310233012-1333001103301303-3011333001300332-0201122233121333-2312110211200103-1333302013013131-2312232122010130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0231332111103211-2012003310023331-2332323010101232-0330133310111212-2010210001000022-3233033013001220-3131212121202013-0003211311200232)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3313313031230022-2312333302000103-3231103120330212-0100310323113313-3223321221332123-0201100310110123-3021300111323110-0211121122030022)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2101133002330001-2330300231031122-0101033113031001-2030032222002131-2021022000200232-2203330321023113-0231310312021020-0313202222000031)
- origin_pools.pools.origin_servers.origin_servers

<a id="canonical-3323110213211100-0133232221103303-1333122212102302-1102021212113212-2230012013201010-0023120230123120-3020331323232131-3100021011210330"></a>

Type: `"list"`. Computed.

List of Origin Servers. List of origin servers for Proxy.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3320120311213131-1120122030030223-1200013202312113-3003021301221011-3120231130132133-1030021213030101-1212231222102220-0230013200313311"></a>

### Direct properties for `origin_pools.pools.origin_servers.origin_servers`

- [k8s_service](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2231203012210011-1322031220221013-1110230232223022-3222311333310101-3232032100312211-3132022100132221-3313331010303100-3112223101221133): complete subsection reference.

- [private_ip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-2003022110023131-2020123013032313-0310123231223001-1223322211211210-0312230131030133-0311310220103133-1330133030212311-2123132120033110): complete subsection reference.

- [public_ip](data-sources--bigip_http_proxy--reference--group-002.md#canonical-0301302103000331-2321000012100202-2033332133332322-1332223222321202-0032201121100013-0302222333032221-3311002011230333-1233131303200331): complete subsection reference.

- [public_name](data-sources--bigip_http_proxy--reference--group-002.md#canonical-3131332003012130-3200013310120130-0300132000112002-1220331130310023-3213322212232121-2031131233110313-1200331011110213-1110323201210131): complete subsection reference.

<a id="canonical-2231203012210011-1322031220221013-1110230232223022-3222311333310101-3232032100312211-3132022100132221-3313331010303100-3112223101221133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.k8s_service` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0231332111103211-2012003310023331-2332323010101232-0330133310111212-2010210001000022-3233033013001220-3131212121202013-0003211311200232)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3313313031230022-2312333302000103-3231103120330212-0100310323113313-3223321221332123-0201100310110123-3021300111323110-0211121122030022)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2101133002330001-2330300231031122-0101033113031001-2030032222002131-2021022000200232-2203330321023113-0231310312021020-0313202222000031)
- [origin_pools.pools.origin_servers.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0311113110223030-0220031310233012-1333001103301303-3011333001300332-0201122233121333-2312110211200103-1333302013013131-2312232122010130)
- origin_pools.pools.origin_servers.origin_servers.k8s_service

<a id="canonical-0031010331320310-3321323231110011-2200210333231100-2332210030112311-0010222312233103-0102212131220313-0321303133033310-2020023123202020"></a>

Type: `"single"`. Computed.

Specify origin server with K8s service name and site information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"vk8s_networks\"]",
  "x-ves-oneof-field-service_info": "[\"service_name\"]"
}
```

<a id="canonical-2202133231120201-0213310101111330-3113021103130323-3310200101321322-0221020230323103-1130210233323322-0011332101210022-2112331322223131"></a>

### Direct properties for `origin_pools.pools.origin_servers.origin_servers.k8s_service`

- [inside_network](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3221000200101023-2113101201323212-1022033002103110-3213321322001323-1121233020013011-0202111201332223-2012331000323112-0313231003122312): complete subsection reference.

- [outside_network](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1212000031033213-3000110301001212-1023103212012201-2211220332030201-3211112012233323-2313311311000221-2303331331210313-3300031231121130): complete subsection reference.

<a id="canonical-3303320021121202-2123303232200020-2200102002012302-1213033232230232-3232202110221330-1233312001232210-2131130013233300-2121033103103211"></a>

<a id="canonical-1030313211111012-0311323200200103-2102222203022221-2121323231331112-2121032101300033-0213033112010021-0303013123031333-1302321320221103"></a>

#### `origin_pools.pools.origin_servers.origin_servers.k8s_service.protocol` property

Type: `"string"`. Computed.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_UDP\] Type of protocol - PROTOCOL\_TCP: TCP - PROTOCOL\_UDP: UDP.
Possible values are \`PROTOCOL\_TCP\`, \`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "PROTOCOL_TCP",
  "enum": [
    "PROTOCOL_TCP",
    "PROTOCOL_UDP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3210120330332121-3032011023230202-0221331100103133-1011312022122021-0221222300201122-0133310233200131-3322212122311233-0232300020310201"></a>

<a id="canonical-1030110131023131-3033222013020023-2002003232223312-3131222232321301-3212311112102101-0120033230211011-2003013133113132-1120022202012232"></a>

#### `origin_pools.pools.origin_servers.origin_servers.k8s_service.service_name` property

Type: `"string"`. Computed.

Exclusive with \[\] K8s service name of the origin server will be listed, including the namespace
and cluster-ID. For vK8s services, you need to enter a string with the format
servicename.namespace:example-namespace'frontend', namespace is 'speedtest' and cluster-ID is
'prod', then you will enter..

Additional upstream details:

For vK8s services, you need to enter a string with the format
servicename.namespace:example-namespace"frontend", namespace is "speedtest" and cluster-ID is
"prod", then you will enter "frontend.speedtest:prod". Both namespace and cluster-ID are optional.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ves_service_namespace_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ves_service_namespace_name": "true"
  }
}
```

- [site_locator](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3013001020131302-0220020310303001-3321101311030100-1301133221301211-1212123100213221-1130232123312010-0100332110301323-3321302302101333): complete subsection reference.

- [snat_pool](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2012311320203233-3122103123002220-1211033112101203-1232213112000121-2202113102331300-3233012203121102-2320113111213102-0222100132223212): complete subsection reference.

- [vk8s_networks](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0333332022122303-2030010211001122-2111322322110022-1130112012001023-2030303122131312-2102001102200302-3121123330201332-1111103023223111): complete subsection reference.

<a id="canonical-3221000200101023-2113101201323212-1022033002103110-3213321322001323-1121233020013011-0202111201332223-2012331000323112-0313231003122312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.k8s_service.inside_network` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0231332111103211-2012003310023331-2332323010101232-0330133310111212-2010210001000022-3233033013001220-3131212121202013-0003211311200232)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3313313031230022-2312333302000103-3231103120330212-0100310323113313-3223321221332123-0201100310110123-3021300111323110-0211121122030022)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2101133002330001-2330300231031122-0101033113031001-2030032222002131-2021022000200232-2203330321023113-0231310312021020-0313202222000031)
- [origin_pools.pools.origin_servers.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0311113110223030-0220031310233012-1333001103301303-3011333001300332-0201122233121333-2312110211200103-1333302013013131-2312232122010130)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2231203012210011-1322031220221013-1110230232223022-3222311333310101-3232032100312211-3132022100132221-3313331010303100-3112223101221133)
- origin_pools.pools.origin_servers.origin_servers.k8s_service.inside_network

<a id="canonical-1300002131321021-2221113211102111-2002101303003032-0231102100311002-0132100110300300-0210313010223110-0332333323030213-1322310221300122"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inside network.

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

<a id="canonical-1212000031033213-3000110301001212-1023103212012201-2211220332030201-3211112012233323-2313311311000221-2303331331210313-3300031231121130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.k8s_service.outside_network` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0231332111103211-2012003310023331-2332323010101232-0330133310111212-2010210001000022-3233033013001220-3131212121202013-0003211311200232)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3313313031230022-2312333302000103-3231103120330212-0100310323113313-3223321221332123-0201100310110123-3021300111323110-0211121122030022)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2101133002330001-2330300231031122-0101033113031001-2030032222002131-2021022000200232-2203330321023113-0231310312021020-0313202222000031)
- [origin_pools.pools.origin_servers.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0311113110223030-0220031310233012-1333001103301303-3011333001300332-0201122233121333-2312110211200103-1333302013013131-2312232122010130)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2231203012210011-1322031220221013-1110230232223022-3222311333310101-3232032100312211-3132022100132221-3313331010303100-3112223101221133)
- origin_pools.pools.origin_servers.origin_servers.k8s_service.outside_network

<a id="canonical-0221231222003311-3200303030200000-3031033113201322-0311223000301033-3330311013132123-2201200213330220-2223322231310002-2320201213302312"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for outside network.

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

<a id="canonical-3013001020131302-0220020310303001-3321101311030100-1301133221301211-1212123100213221-1130232123312010-0100332110301323-3321302302101333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0231332111103211-2012003310023331-2332323010101232-0330133310111212-2010210001000022-3233033013001220-3131212121202013-0003211311200232)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3313313031230022-2312333302000103-3231103120330212-0100310323113313-3223321221332123-0201100310110123-3021300111323110-0211121122030022)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2101133002330001-2330300231031122-0101033113031001-2030032222002131-2021022000200232-2203330321023113-0231310312021020-0313202222000031)
- [origin_pools.pools.origin_servers.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0311113110223030-0220031310233012-1333001103301303-3011333001300332-0201122233121333-2312110211200103-1333302013013131-2312232122010130)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2231203012210011-1322031220221013-1110230232223022-3222311333310101-3232032100312211-3132022100132221-3313331010303100-3112223101221133)
- origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator

<a id="canonical-3121201200322000-2312120223232310-3121102302012021-1313230330100330-2012001101311301-1330130002312112-1132213201001221-0110030300032220"></a>

Type: `"single"`. Computed.

This message defines a reference to a site or virtual site object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

<a id="canonical-0330102302230213-1320010231220230-2011230321300031-3132121013103211-0223221112330202-2303213301211213-1301302131023103-1120303121303230"></a>

### Direct properties for `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator`

- [site](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1203112022111213-2312032113112212-3331321330113230-2132030100311233-0333111030201112-2113110112102323-2333030031232131-2010201031313211): complete subsection reference.

- [virtual_site](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1233330111112002-1330102002011302-1122323003312222-0111101102333002-2301133331223033-0320311330211312-1102212002002223-2213231211322120): complete subsection reference.

<a id="canonical-1203112022111213-2312032113112212-3331321330113230-2132030100311233-0333111030201112-2113110112102323-2333030031232131-2010201031313211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0231332111103211-2012003310023331-2332323010101232-0330133310111212-2010210001000022-3233033013001220-3131212121202013-0003211311200232)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3313313031230022-2312333302000103-3231103120330212-0100310323113313-3223321221332123-0201100310110123-3021300111323110-0211121122030022)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2101133002330001-2330300231031122-0101033113031001-2030032222002131-2021022000200232-2203330321023113-0231310312021020-0313202222000031)
- [origin_pools.pools.origin_servers.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0311113110223030-0220031310233012-1333001103301303-3011333001300332-0201122233121333-2312110211200103-1333302013013131-2312232122010130)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2231203012210011-1322031220221013-1110230232223022-3222311333310101-3232032100312211-3132022100132221-3313331010303100-3112223101221133)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3013001020131302-0220020310303001-3321101311030100-1301133221301211-1212123100213221-1130232123312010-0100332110301323-3321302302101333)
- origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site

<a id="canonical-2202322301232233-0102000322020323-2202030223013111-1002202203303210-1022100030011123-0220010302022223-1330100231311000-1330020222102022"></a>

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

<a id="canonical-0020103230223121-0233022332303232-2031202212201300-2323333303203200-0033303303133023-1222231230000222-3101300123322023-0100001321311323"></a>

### Direct properties for `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site`

<a id="canonical-3331033123211113-1233113001013020-1110132331020232-0223232200231220-0213012100112302-2211032010103000-3130001121133113-1222312320321100"></a>

#### `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.name` property

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

<a id="canonical-3301033332332030-2033231131012300-3011012203012200-2203300303133312-0003323232011320-1032123010000312-1230203110131012-3111322000302002"></a>

<a id="canonical-3322102203233201-1320330132223323-0313302323321111-2232020230233211-2213213303003021-3311213310120003-2311123321100202-1001301302320121"></a>

#### `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.namespace` property

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

<a id="canonical-1021032003000233-2233211222230012-3103022013023223-3110011110003303-2212221310003033-0203222203100310-3330222212203031-0333101233322011"></a>

<a id="canonical-0001130003330001-1010300010101133-2011031232130010-3123001001202122-2022121021330033-2100131110022021-2012333311311122-1332112003022211"></a>

#### `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.site.tenant` property

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

<a id="canonical-1233330111112002-1330102002011302-1122323003312222-0111101102333002-2301133331223033-0320311330211312-1102212002002223-2213231211322120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0231332111103211-2012003310023331-2332323010101232-0330133310111212-2010210001000022-3233033013001220-3131212121202013-0003211311200232)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3313313031230022-2312333302000103-3231103120330212-0100310323113313-3223321221332123-0201100310110123-3021300111323110-0211121122030022)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2101133002330001-2330300231031122-0101033113031001-2030032222002131-2021022000200232-2203330321023113-0231310312021020-0313202222000031)
- [origin_pools.pools.origin_servers.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0311113110223030-0220031310233012-1333001103301303-3011333001300332-0201122233121333-2312110211200103-1333302013013131-2312232122010130)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2231203012210011-1322031220221013-1110230232223022-3222311333310101-3232032100312211-3132022100132221-3313331010303100-3112223101221133)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3013001020131302-0220020310303001-3321101311030100-1301133221301211-1212123100213221-1130232123312010-0100332110301323-3321302302101333)
- origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site

<a id="canonical-3330321300301311-0203010320200033-0022312112331323-3011021121022231-2103113220302222-0331302130331301-2203202021223120-0113112001013013"></a>

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

<a id="canonical-3123031312230132-2001312203221012-1223302023032332-3332210220211131-3013301012003232-2322203001330302-2310031330323221-0230123103200031"></a>

### Direct properties for `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site`

<a id="canonical-1000212123111013-3333002213012211-2201331310122012-0023332233023223-2112010103333010-2230023133303110-0303223323130022-1011230332113101"></a>

#### `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.name` property

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

<a id="canonical-0310230010230131-2010011231123302-1021102013110032-3123201001001031-2230200033222103-2200010013032113-0031012200200232-3231230033130131"></a>

<a id="canonical-3130330133312300-3301000033310002-0332310302231020-1031120111232303-1213320202301003-2223303220223133-3132221332003112-2001203110231323"></a>

#### `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.namespace` property

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

<a id="canonical-2103211112110021-2200122301000213-3021210032010301-2133100011131010-0112201302101130-0322030222220110-0223113201023120-1311332233020222"></a>

<a id="canonical-3112000221221301-0220233013212013-1322203121331311-3300311222110332-3102311222212103-1020030330022032-0032203230330010-1320102001332231"></a>

#### `origin_pools.pools.origin_servers.origin_servers.k8s_service.site_locator.virtual_site.tenant` property

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

<a id="canonical-2012311320203233-3122103123002220-1211033112101203-1232213112000121-2202113102331300-3233012203121102-2320113111213102-0222100132223212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0231332111103211-2012003310023331-2332323010101232-0330133310111212-2010210001000022-3233033013001220-3131212121202013-0003211311200232)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3313313031230022-2312333302000103-3231103120330212-0100310323113313-3223321221332123-0201100310110123-3021300111323110-0211121122030022)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2101133002330001-2330300231031122-0101033113031001-2030032222002131-2021022000200232-2203330321023113-0231310312021020-0313202222000031)
- [origin_pools.pools.origin_servers.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0311113110223030-0220031310233012-1333001103301303-3011333001300332-0201122233121333-2312110211200103-1333302013013131-2312232122010130)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2231203012210011-1322031220221013-1110230232223022-3222311333310101-3232032100312211-3132022100132221-3313331010303100-3112223101221133)
- origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool

<a id="canonical-1002303133331100-0103322110233111-3232300032302233-1031231122333011-2320121123011311-2002310211102303-2010222113331203-2121001302132120"></a>

Type: `"single"`. Computed.

SNAT Pool. SNAT Pool configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

<a id="canonical-1113020121101211-1122210321022001-0323330021231300-3300301310211222-0220122312020301-0010232010012100-1302202123031131-1220030220331002"></a>

### Direct properties for `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool`

- [no_snat_pool](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1120220110223131-2230332301321311-2120211010323212-2100121310333320-1021333311110310-1310303203001100-0113103212231101-0132113302212322): complete subsection reference.

- [snat_pool](data-sources--bigip_http_proxy--reference--group-001.md#canonical-1313213220013222-0223033121012202-3103220100120033-1013232031031330-3210122032310102-1102300333101312-0201002020030202-3001031321112020): complete subsection reference.

<a id="canonical-1120220110223131-2230332301321311-2120211010323212-2100121310333320-1021333311110310-1310303203001100-0113103212231101-0132113302212322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0231332111103211-2012003310023331-2332323010101232-0330133310111212-2010210001000022-3233033013001220-3131212121202013-0003211311200232)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3313313031230022-2312333302000103-3231103120330212-0100310323113313-3223321221332123-0201100310110123-3021300111323110-0211121122030022)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2101133002330001-2330300231031122-0101033113031001-2030032222002131-2021022000200232-2203330321023113-0231310312021020-0313202222000031)
- [origin_pools.pools.origin_servers.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0311113110223030-0220031310233012-1333001103301303-3011333001300332-0201122233121333-2312110211200103-1333302013013131-2312232122010130)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2231203012210011-1322031220221013-1110230232223022-3222311333310101-3232032100312211-3132022100132221-3313331010303100-3112223101221133)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2012311320203233-3122103123002220-1211033112101203-1232213112000121-2202113102331300-3233012203121102-2320113111213102-0222100132223212)
- origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool

<a id="canonical-0202110131032323-1110103332323310-0230131012100131-2331212213103213-2120301312011031-0001200110021012-2312113210033021-3112001301010222"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no snat pool.

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

<a id="canonical-1313213220013222-0223033121012202-3103220100120033-1013232031031330-3210122032310102-1102300333101312-0201002020030202-3001031321112020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0231332111103211-2012003310023331-2332323010101232-0330133310111212-2010210001000022-3233033013001220-3131212121202013-0003211311200232)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3313313031230022-2312333302000103-3231103120330212-0100310323113313-3223321221332123-0201100310110123-3021300111323110-0211121122030022)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2101133002330001-2330300231031122-0101033113031001-2030032222002131-2021022000200232-2203330321023113-0231310312021020-0313202222000031)
- [origin_pools.pools.origin_servers.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0311113110223030-0220031310233012-1333001103301303-3011333001300332-0201122233121333-2312110211200103-1333302013013131-2312232122010130)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2231203012210011-1322031220221013-1110230232223022-3222311333310101-3232032100312211-3132022100132221-3313331010303100-3112223101221133)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2012311320203233-3122103123002220-1211033112101203-1232213112000121-2202113102331300-3233012203121102-2320113111213102-0222100132223212)
- origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool

<a id="canonical-3020201232130102-3021202210032021-3120030303010100-1313121221223322-3331102331210303-3122201230133212-0320313320233221-2003223203323120"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2222223111002121-2210123133233203-3122213320330210-1302121231300110-3212110023022330-0100221111220031-2201123333003310-1202011223300112"></a>

### Direct properties for `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool`

<a id="canonical-0031103033222232-0300330100211312-2212002002222200-0202002112211100-3023211130021303-2222333101001310-1230310200222201-1313032100300312"></a>

#### `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool.prefixes` property

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

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

<a id="canonical-0333332022122303-2030010211001122-2111322322110022-1130112012001023-2030303122131312-2102001102200302-3121123330201332-1111103023223111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.k8s_service.vk8s_networks` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../data-sources/bigip_http_proxy.md#canonical-0023331032232330-1000103122121111-1312200212021131-0121303000221120-3202200300130220-2113130013223101-2121132022100113-3123313211100112)
- [Property reference](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0032303031101201-2032022330022010-1302301210002232-2111021301322320-0310111332211013-3303200232022211-2333023210300102-1210223202311002)
- [origin_pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0231332111103211-2012003310023331-2332323010101232-0330133310111212-2010210001000022-3233033013001220-3131212121202013-0003211311200232)
- [origin_pools.pools](data-sources--bigip_http_proxy--reference--group-001.md#canonical-3313313031230022-2312333302000103-3231103120330212-0100310323113313-3223321221332123-0201100310110123-3021300111323110-0211121122030022)
- [origin_pools.pools.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2101133002330001-2330300231031122-0101033113031001-2030032222002131-2021022000200232-2203330321023113-0231310312021020-0313202222000031)
- [origin_pools.pools.origin_servers.origin_servers](data-sources--bigip_http_proxy--reference--group-001.md#canonical-0311113110223030-0220031310233012-1333001103301303-3011333001300332-0201122233121333-2312110211200103-1333302013013131-2312232122010130)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](data-sources--bigip_http_proxy--reference--group-001.md#canonical-2231203012210011-1322031220221013-1110230232223022-3222311333310101-3232032100312211-3132022100132221-3313331010303100-3112223101221133)
- origin_pools.pools.origin_servers.origin_servers.k8s_service.vk8s_networks

<a id="canonical-2331312332101113-2003323101322000-0003311303131123-3302031223210003-1032132313032223-2102113020031231-3023211023030031-2210020112000101"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for vk8s networks.

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
