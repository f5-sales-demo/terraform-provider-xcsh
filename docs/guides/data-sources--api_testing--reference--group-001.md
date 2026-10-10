---
page_title: "xcsh_api_testing reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_api_testing reference."
---

# xcsh_api_testing reference

<a id="canonical-3002213323001033-3300321333010022-1122133200032233-1213022031201310-3110232103211233-0103222003311303-0323133133201020-2103123031102103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-1103130220132111-1121033102212030-2032323111203320-0101012101310232-3300331302031232-1332131330231320-0212001110131013-1010030122002230)
- Property reference

<a id="canonical-0122002203023313-1032231211321333-2113031323310202-0032133321321202-0313022310202301-2020212213230310-2320102323230223-1110123213023211"></a>

### Direct properties for `xcsh_api_testing`

<a id="canonical-1103020021031211-1012120010012332-0332232311033102-1120120130323221-3131122001013200-3001111012100010-1211302220023020-3321131110111122"></a>

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

<a id="canonical-3201001321020223-2123230312311003-3131002302220011-2022310133110302-0013110012233230-2300231332120333-3010301133211000-1221110203032133"></a>

<a id="canonical-3313112202220032-2302312031302213-0303021232130330-3332220123121132-3203003330010220-0023001100211100-3033130131120220-2333001032033302"></a>

#### `custom_header_value` property

Type: `"string"`. Computed.

Add x-F5-API-testing-identifier header value to prevent security flags on API testing traffic.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-0303131212201020-1221031301022320-3113102313330322-1313220202312003-0013113032333122-3032311231133323-2111120313213023-2331302013123031"></a>

<a id="canonical-3300133311113131-2330233201001102-0221013331211232-3213021001102100-2022112133200032-2313031201122012-3030003303312301-1113002123200002"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the APITesting.

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

- [domains](data-sources--api_testing--reference--group-001.md#canonical-3131031132021133-3010303022131200-2033203230030023-3131100130120320-1203210003222110-1020112003233203-2133130033211033-0011012200000012): complete subsection reference.

- [every_day](data-sources--api_testing--reference--group-001.md#canonical-2332330000101223-2332033331213132-1301121202210033-3103023332020211-1332031203333033-1103110000210321-0012010013012201-0313032122013212): complete subsection reference.

- [every_month](data-sources--api_testing--reference--group-001.md#canonical-0122310103331021-2020002230323013-1010102020012020-0010131103300303-2132223303320021-2313113101112011-2022122010230101-3111013330111132): complete subsection reference.

- [every_week](data-sources--api_testing--reference--group-001.md#canonical-1113210123100310-3302013201121202-0001131003030011-3012002030222033-3121033100323211-0012003202031021-0012001301310130-2133332223112101): complete subsection reference.

<a id="canonical-3100002301203022-0313032210201332-1023231200113323-1110321133121210-1131311300203011-1211322030302230-2230310122113223-2111233132330113"></a>

<a id="canonical-3312201120011021-3103310120001120-3033321032130232-3223102230213112-2311301200222331-1203123033212122-1130201331233013-2031231101202130"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3102320013100130-1213101221220300-1003212010101032-1001111122301230-2032212212330300-1123110230110223-2311232330002133-2223230311203031"></a>

<a id="canonical-0020223233322110-3201121332131020-0332212210021102-3132331033303332-1001333230012113-1231223033032112-2002021002110203-0310013212102201"></a>

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

<a id="canonical-0212133222122002-1021223203332120-0330220011033033-0231020313303330-1211300313212022-2202232100231321-2323112023203230-2321100223032100"></a>

<a id="canonical-1032010322232313-3003222100313212-3223000032133202-3001032102132232-3110113200222020-0110331121013211-2332230300112123-3223003110212000"></a>

#### `name` property

Type: `"string"`. Required.

Name of the APITesting.

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

<a id="canonical-1231313211222201-2211331331013321-2200321212331322-0030311200021232-2113002200213000-2300232330202021-1200333312100113-2002310101132232"></a>

<a id="canonical-0310321302323322-2020023322002130-3203022330313010-1003033201101002-0221323111212210-3001112200302101-2332312310300112-0113033113122322"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the APITesting exists.

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

<a id="canonical-1120200122221011-3012210102301211-2112032203202023-3010200030213031-2111313101023331-2113131223231011-0131203112010031-2113122033212032"></a>

### All schema paths for `xcsh_api_testing`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--api_testing--reference--group-001.md#canonical-1103020021031211-1012120010012332-0332232311033102-1120120130323221-3131122001013200-3001111012100010-1211302220023020-3321131110111122) |
| `custom_header_value` | [custom_header_value](data-sources--api_testing--reference--group-001.md#canonical-3201001321020223-2123230312311003-3131002302220011-2022310133110302-0013110012233230-2300231332120333-3010301133211000-1221110203032133) |
| `description` | [description](data-sources--api_testing--reference--group-001.md#canonical-0303131212201020-1221031301022320-3113102313330322-1313220202312003-0013113032333122-3032311231133323-2111120313213023-2331302013123031) |
| `domains` | [domains](data-sources--api_testing--reference--group-001.md#canonical-1013131132000222-0312311131020101-2230020302133332-2302012300012003-0322220303220233-2130100213320331-2332321302132210-0221011130133103) |
| `domains.allow_destructive_methods` | [domains.allow_destructive_methods](data-sources--api_testing--reference--group-001.md#canonical-2132023131232033-3322113010110310-3122223100323213-0233131230331211-1302213211031320-3232320130131103-3230110130122313-1121201222001002) |
| `domains.credentials` | [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-2113121020221110-2212320211211302-2110001303131031-1023132322323102-3011330130301300-0301030123033300-1131111201013113-0131100212202222) |
| `domains.credentials.admin` | [domains.credentials.admin](data-sources--api_testing--reference--group-001.md#canonical-3023103201113330-2122020023132112-2200330010032112-0310222320111020-3020013203231311-3233111032332010-1200100201022023-3203203101302321) |
| `domains.credentials.api_key` | [domains.credentials.api_key](data-sources--api_testing--reference--group-001.md#canonical-2230331013031320-3302001333122331-1301111321023220-3132212111112033-1002222303023200-0131302021200232-1130210123233120-1012131101033221) |
| `domains.credentials.api_key.key` | [domains.credentials.api_key.key](data-sources--api_testing--reference--group-001.md#canonical-3320203313320213-1303331210010331-3202120330001231-2333111300031003-3032300203310132-1001110022312030-3003300310102130-1022310102021120) |
| `domains.credentials.api_key.value` | [domains.credentials.api_key.value](data-sources--api_testing--reference--group-001.md#canonical-2232301212332330-3312211002311032-0200133010001322-1012222003033000-0321112212130210-0132300031322031-3230321011002002-1003213033221013) |
| `domains.credentials.api_key.value.blindfold_secret_info` | [domains.credentials.api_key.value.blindfold_secret_info](data-sources--api_testing--reference--group-001.md#canonical-1120010202233203-3233031020303221-0112012011210030-2113113112000200-3223300233330332-3333001013323111-3301103003022212-0221312221133113) |
| `domains.credentials.api_key.value.blindfold_secret_info.decryption_provider` | [domains.credentials.api_key.value.blindfold_secret_info.decryption_provider](data-sources--api_testing--reference--group-001.md#canonical-2021230212111110-2223233213220231-0330131133333333-2331031001300230-3030122213021303-0023021112110123-3030123230321321-2002031122221120) |
| `domains.credentials.api_key.value.blindfold_secret_info.location` | [domains.credentials.api_key.value.blindfold_secret_info.location](data-sources--api_testing--reference--group-001.md#canonical-2213223233202330-0231312121231322-1033020331101013-0011200033111333-0133323222232311-0311233232303312-2230221220022021-2331013032131121) |
| `domains.credentials.api_key.value.blindfold_secret_info.store_provider` | [domains.credentials.api_key.value.blindfold_secret_info.store_provider](data-sources--api_testing--reference--group-001.md#canonical-2123122031233302-2020232201310110-3333120233323222-1331102332302133-3000010031111201-3001133033221202-2111112131312213-0210110001001201) |
| `domains.credentials.api_key.value.clear_secret_info` | [domains.credentials.api_key.value.clear_secret_info](data-sources--api_testing--reference--group-001.md#canonical-1111232200223002-3312010012121301-1232003301211122-1002201021031332-3232331010023230-0123010211331322-3310213220102230-2101130132223210) |
| `domains.credentials.api_key.value.clear_secret_info.provider_ref` | [domains.credentials.api_key.value.clear_secret_info.provider_ref](data-sources--api_testing--reference--group-001.md#canonical-1030022331122011-0303013112120103-2302212322331122-3020303202333130-3130113212313021-2320030130301300-1323022321022323-3332123011022022) |
| `domains.credentials.api_key.value.clear_secret_info.url` | [domains.credentials.api_key.value.clear_secret_info.url](data-sources--api_testing--reference--group-001.md#canonical-0133011001031223-3001302133321002-3121200223002131-3202033113010013-1313311331100313-3223323333130010-1220103000322122-2000221031011320) |
| `domains.credentials.basic_auth` | [domains.credentials.basic_auth](data-sources--api_testing--reference--group-001.md#canonical-0310131311023121-0002222131222213-0332033220223312-3021233233223110-0302110131121311-0212202233132113-1322132331023200-1110100232213123) |
| `domains.credentials.basic_auth.password` | [domains.credentials.basic_auth.password](data-sources--api_testing--reference--group-001.md#canonical-3033320110213213-1213023013122222-1230333231233210-2332011223222122-2021102320023103-1320322033132002-3310302211212033-2220103312221231) |
| `domains.credentials.basic_auth.password.blindfold_secret_info` | [domains.credentials.basic_auth.password.blindfold_secret_info](data-sources--api_testing--reference--group-001.md#canonical-1322022233223311-3132201112330322-3213310303302000-3003310310012021-0112333233112201-3202121101332322-0221011120001302-2121203011003110) |
| `domains.credentials.basic_auth.password.blindfold_secret_info.decryption_provider` | [domains.credentials.basic_auth.password.blindfold_secret_info.decryption_provider](data-sources--api_testing--reference--group-001.md#canonical-3220332201333333-1332311220101333-1030232221321322-2032201331213313-3102200131333011-2120120200323311-2303221000101333-0011123203110232) |
| `domains.credentials.basic_auth.password.blindfold_secret_info.location` | [domains.credentials.basic_auth.password.blindfold_secret_info.location](data-sources--api_testing--reference--group-001.md#canonical-3210031122132101-2103331033012001-1301111230203222-0101230301032312-2320330113322213-3113220320220322-0123211313311211-3133113220323322) |
| `domains.credentials.basic_auth.password.blindfold_secret_info.store_provider` | [domains.credentials.basic_auth.password.blindfold_secret_info.store_provider](data-sources--api_testing--reference--group-001.md#canonical-0213122233331010-3023033000223223-2033210332131132-2111202111332321-2321212332133112-1112021123302101-3032033210233230-1031302213003233) |
| `domains.credentials.basic_auth.password.clear_secret_info` | [domains.credentials.basic_auth.password.clear_secret_info](data-sources--api_testing--reference--group-001.md#canonical-2100333031032003-3032123112211020-2031231313200212-1203111312223200-1233230310322012-0130021030311310-1233133130300320-1220231321030022) |
| `domains.credentials.basic_auth.password.clear_secret_info.provider_ref` | [domains.credentials.basic_auth.password.clear_secret_info.provider_ref](data-sources--api_testing--reference--group-001.md#canonical-1213323300002222-0311123101203030-2130313103121130-2221222120021122-1023322221331022-3100210200113113-3230011213003021-2230213212110221) |
| `domains.credentials.basic_auth.password.clear_secret_info.url` | [domains.credentials.basic_auth.password.clear_secret_info.url](data-sources--api_testing--reference--group-001.md#canonical-2102311000113002-0010012113221100-0313231030232202-0330023001231232-1111221332133333-1002002001132123-1130301212323121-0221110012021011) |
| `domains.credentials.basic_auth.user` | [domains.credentials.basic_auth.user](data-sources--api_testing--reference--group-001.md#canonical-3332010120322002-1120231332211033-2000230033333230-2310031331322310-3112220220132333-1011301033010201-2303200111223303-1003033120021200) |
| `domains.credentials.bearer_token` | [domains.credentials.bearer_token](data-sources--api_testing--reference--group-001.md#canonical-3333302033332302-0303220303222322-0020112202322002-2003130130030300-2002200130330221-0110100211003101-1023311321213023-2310122130030010) |
| `domains.credentials.bearer_token.token` | [domains.credentials.bearer_token.token](data-sources--api_testing--reference--group-001.md#canonical-1020202220311202-2100310222031310-2233320212101303-3202103223032313-1113322032222012-0110120001233212-3213220230011121-1301232112111222) |
| `domains.credentials.bearer_token.token.blindfold_secret_info` | [domains.credentials.bearer_token.token.blindfold_secret_info](data-sources--api_testing--reference--group-001.md#canonical-2021112330110020-0221310132202222-1020100233222100-3023230103023012-2021303320220211-2110012002213223-0023203333131212-2200002311331223) |
| `domains.credentials.bearer_token.token.blindfold_secret_info.decryption_provider` | [domains.credentials.bearer_token.token.blindfold_secret_info.decryption_provider](data-sources--api_testing--reference--group-001.md#canonical-3310100322211111-3021201122133030-3231233333200200-1201003103223111-1212202030301213-3102033310320303-3202130200232013-2333333003312222) |
| `domains.credentials.bearer_token.token.blindfold_secret_info.location` | [domains.credentials.bearer_token.token.blindfold_secret_info.location](data-sources--api_testing--reference--group-001.md#canonical-1200310021010233-0221121212033220-1213310322201223-2000031101222221-2120233323012030-2123212021103321-2301211300303301-1003031221020311) |
| `domains.credentials.bearer_token.token.blindfold_secret_info.store_provider` | [domains.credentials.bearer_token.token.blindfold_secret_info.store_provider](data-sources--api_testing--reference--group-001.md#canonical-0232210222032220-0130323033103022-0000323202320001-3020323312031210-2220313200222200-0233313230322202-1223100212111031-2003103103010200) |
| `domains.credentials.bearer_token.token.clear_secret_info` | [domains.credentials.bearer_token.token.clear_secret_info](data-sources--api_testing--reference--group-001.md#canonical-1210312332030300-3113220123300030-3222331000320331-2203330322010322-3013020003331222-1022130112303122-1103220003302313-3333112202312200) |
| `domains.credentials.bearer_token.token.clear_secret_info.provider_ref` | [domains.credentials.bearer_token.token.clear_secret_info.provider_ref](data-sources--api_testing--reference--group-001.md#canonical-3033123022113031-2301330232230310-0300001021112211-3201323302302131-1202311131331313-3022000023231112-1312302211101322-3123133020201222) |
| `domains.credentials.bearer_token.token.clear_secret_info.url` | [domains.credentials.bearer_token.token.clear_secret_info.url](data-sources--api_testing--reference--group-001.md#canonical-2333313001303112-0300302203302321-0033310020111000-0123011011221000-0211032303031011-0231330213110102-0020332212332230-3010012333131131) |
| `domains.credentials.credential_name` | [domains.credentials.credential_name](data-sources--api_testing--reference--group-001.md#canonical-2102032210331320-2200030321320122-1310211130303012-3113032112320303-0303233202232021-1110232301233302-1020100020331302-0201113101023332) |
| `domains.credentials.login_endpoint` | [domains.credentials.login_endpoint](data-sources--api_testing--reference--group-001.md#canonical-1312310200200210-0302220330021121-1221001012310121-0103213302033013-1221213213322011-1123313002131102-1033011322100130-2232213100132323) |
| `domains.credentials.login_endpoint.json_payload` | [domains.credentials.login_endpoint.json_payload](data-sources--api_testing--reference--group-001.md#canonical-2303122032121013-1111003101120011-1013231300002312-3232133123333310-0322320001032020-2311022132210320-2220111333121322-0130103110031113) |
| `domains.credentials.login_endpoint.json_payload.blindfold_secret_info` | [domains.credentials.login_endpoint.json_payload.blindfold_secret_info](data-sources--api_testing--reference--group-001.md#canonical-0000123313012110-1300031230023111-1233232321101312-1231303220312201-1133311320003000-0220302120013321-1001221323233123-1032011311230333) |
| `domains.credentials.login_endpoint.json_payload.blindfold_secret_info.decryption_provider` | [domains.credentials.login_endpoint.json_payload.blindfold_secret_info.decryption_provider](data-sources--api_testing--reference--group-001.md#canonical-0232221201203200-2020330122013330-2212112230022332-0100033311101223-1020212203002113-2220323323020012-2110000131033312-3031223221202202) |
| `domains.credentials.login_endpoint.json_payload.blindfold_secret_info.location` | [domains.credentials.login_endpoint.json_payload.blindfold_secret_info.location](data-sources--api_testing--reference--group-001.md#canonical-2220010131300012-2030330110120111-3111003100201022-0110021231032312-3301212103022220-3022303310000130-1113312323233330-3331121312213031) |
| `domains.credentials.login_endpoint.json_payload.blindfold_secret_info.store_provider` | [domains.credentials.login_endpoint.json_payload.blindfold_secret_info.store_provider](data-sources--api_testing--reference--group-001.md#canonical-3001132322213012-0002230110201013-0033102213331212-3103201220322032-3311321223013223-0000001322230132-2303231113232301-0213123030023101) |
| `domains.credentials.login_endpoint.json_payload.clear_secret_info` | [domains.credentials.login_endpoint.json_payload.clear_secret_info](data-sources--api_testing--reference--group-001.md#canonical-1122011112323321-3202200133103211-0131330322313321-1102320220110121-3223213122100220-1311220232020120-0120111132020321-2322301200303231) |
| `domains.credentials.login_endpoint.json_payload.clear_secret_info.provider_ref` | [domains.credentials.login_endpoint.json_payload.clear_secret_info.provider_ref](data-sources--api_testing--reference--group-001.md#canonical-2300311012231330-2302331312200010-3002303003120330-1200121023033213-1001320121113322-3032311313221033-1232221101301301-2033032303032320) |
| `domains.credentials.login_endpoint.json_payload.clear_secret_info.url` | [domains.credentials.login_endpoint.json_payload.clear_secret_info.url](data-sources--api_testing--reference--group-001.md#canonical-1301013033220122-2023313333201103-1300121013131313-1113323112222300-2221133230032200-3213213303212100-2300300010022311-3101133012012322) |
| `domains.credentials.login_endpoint.method` | [domains.credentials.login_endpoint.method](data-sources--api_testing--reference--group-001.md#canonical-0332031302022303-1301300000320132-1321211231110112-1312131033221231-3203222200332010-0232013100022200-0333011312311223-3221011213121213) |
| `domains.credentials.login_endpoint.path` | [domains.credentials.login_endpoint.path](data-sources--api_testing--reference--group-001.md#canonical-1133330331013211-1003330310012220-0331220312313132-2123233112211031-1033303112312111-2322310032322131-2030101033202300-2222330133230300) |
| `domains.credentials.login_endpoint.token_response_key` | [domains.credentials.login_endpoint.token_response_key](data-sources--api_testing--reference--group-001.md#canonical-0130221103112313-0302121130320211-0320120013203100-0211030302323021-1120302103333231-1231011012222220-1131022022012212-0213311033232022) |
| `domains.credentials.standard` | [domains.credentials.standard](data-sources--api_testing--reference--group-001.md#canonical-2001102331012111-0121132221113310-0231230111332111-3213320323120332-0130013101200330-1110213121132230-2232310030310303-0212032201111322) |
| `domains.domain` | [domains.domain](data-sources--api_testing--reference--group-001.md#canonical-1201001100231012-3113213203003213-2201232331230112-3202131033012223-0212331102030303-3110233100330212-2000311011300120-2122112210010101) |
| `every_day` | [every_day](data-sources--api_testing--reference--group-001.md#canonical-1002112133320131-1113103223332230-2103101203022311-3130200100223301-1323310023313311-1030101030222313-1301003002011112-0122033133232303) |
| `every_month` | [every_month](data-sources--api_testing--reference--group-001.md#canonical-2322320322313112-3031010302331223-0213123102313213-1131132122230130-2332120313200230-3102300213022012-2312331110302001-3021122300203310) |
| `every_week` | [every_week](data-sources--api_testing--reference--group-001.md#canonical-0122300232033003-2000113000323033-3131200131033012-0131033010312033-2320122122222023-3220232113122130-0313222121102113-0333211000302330) |
| `id` | [ID](data-sources--api_testing--reference--group-001.md#canonical-3100002301203022-0313032210201332-1023231200113323-1110321133121210-1131311300203011-1211322030302230-2230310122113223-2111233132330113) |
| `labels` | [labels](data-sources--api_testing--reference--group-001.md#canonical-3102320013100130-1213101221220300-1003212010101032-1001111122301230-2032212212330300-1123110230110223-2311232330002133-2223230311203031) |
| `name` | [name](data-sources--api_testing--reference--group-001.md#canonical-0212133222122002-1021223203332120-0330220011033033-0231020313303330-1211300313212022-2202232100231321-2323112023203230-2321100223032100) |
| `namespace` | [namespace](data-sources--api_testing--reference--group-001.md#canonical-1231313211222201-2211331331013321-2200321212331322-0030311200021232-2113002200213000-2300232330202021-1200333312100113-2002310101132232) |

<a id="canonical-3131031132021133-3010303022131200-2033203230030023-3131100130120320-1203210003222110-1020112003233203-2133130033211033-0011012200000012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains` properties

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-1103130220132111-1121033102212030-2032323111203320-0101012101310232-3300331302031232-1332131330231320-0212001110131013-1010030122002230)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-3002213323001033-3300321333010022-1122133200032233-1213022031201310-3110232103211233-0103222003311303-0323133133201020-2103123031102103)
- domains

<a id="canonical-1013131132000222-0312311131020101-2230020302133332-2302012300012003-0322220303220233-2130100213320331-2332321302132210-0221011130133103"></a>

Type: `"list"`. Computed.

Add and configure testing domains and credentials.

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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

<a id="canonical-0332012101221103-0322030323120311-3311013322330022-1003130130011023-3302011213230303-3221032321231000-3133312131332103-3103133232023002"></a>

### Direct properties for `domains`

<a id="canonical-2132023131232033-3322113010110310-3122223100323213-0233131230331211-1302213211031320-3232320130131103-3230110130122313-1121201222001002"></a>

#### `domains.allow_destructive_methods` property

Type: `"bool"`. Computed.

Enable to allow API Testing to execute against destructive methods. Use with caution as these may
modify or DELETE data.

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

- [credentials](data-sources--api_testing--reference--group-001.md#canonical-3300302021011323-1233011033232202-0001131123022203-0113010333011001-2021301130020132-0223321201313022-3300013030331020-3221333322300112): complete subsection reference.

<a id="canonical-1201001100231012-3113213203003213-2201232331230112-3202131033012223-0212331102030303-3110233100330212-2000311011300120-2122112210010101"></a>

<a id="canonical-0133201131012111-3021003033121302-2221212121332021-0123331311122323-0203100331332000-2133110002231021-1021130302302123-3323001101012032"></a>

#### `domains.domain` property

Type: `"string"`. Computed.

Add your testing environment domain. Be aware that running tests on a production domain can impact
live applications, as API testing cannot distinguish between production and testing environments.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

<a id="canonical-3300302021011323-1233011033232202-0001131123022203-0113010333011001-2021301130020132-0223321201313022-3300013030331020-3221333322300112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials` properties

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-1103130220132111-1121033102212030-2032323111203320-0101012101310232-3300331302031232-1332131330231320-0212001110131013-1010030122002230)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-3002213323001033-3300321333010022-1122133200032233-1213022031201310-3110232103211233-0103222003311303-0323133133201020-2103123031102103)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-3131031132021133-3010303022131200-2033203230030023-3131100130120320-1203210003222110-1020112003233203-2133130033211033-0011012200000012)
- domains.credentials

<a id="canonical-2113121020221110-2212320211211302-2110001303131031-1023132322323102-3011330130301300-0301030123033300-1131111201013113-0131100212202222"></a>

Type: `"list"`. Computed.

Add credentials for API testing to use in the selected environment.

Receipt-pinned upstream constraints:

```json
{
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

<a id="canonical-1211031122301113-0300121203213132-1302332232311010-3121111232232231-2113121300233231-2011210121100323-2221131300220212-3201202303213010"></a>

### Direct properties for `domains.credentials`

- [admin](data-sources--api_testing--reference--group-001.md#canonical-1323002131113221-1230221322223220-1302203131311221-2122133201003322-0011323330132323-1311321103321231-0101022313013030-0332132022112202): complete subsection reference.

- [api_key](data-sources--api_testing--reference--group-001.md#canonical-2323012032023313-2023111230121123-3012012100133321-3110102131113012-1312012331001023-3102320211110021-2123213113002301-1122312003302023): complete subsection reference.

- [basic_auth](data-sources--api_testing--reference--group-001.md#canonical-2222012103200021-2301131133131111-1001100303300203-1312321200133203-1112233030011310-1033130110010302-0020033113311323-1103302310233122): complete subsection reference.

- [bearer_token](data-sources--api_testing--reference--group-001.md#canonical-1002300133110320-3332210320021021-1111101131301023-0231213012302320-0122202302002011-0210210210210113-3032002033021311-1031201330223013): complete subsection reference.

<a id="canonical-2102032210331320-2200030321320122-1310211130303012-3113032112320303-0303233202232021-1110232301233302-1020100020331302-0201113101023332"></a>

<a id="canonical-2222033320131032-3201133203122022-1022012222130033-3012321300310012-2012332112331010-1211013212113203-3133220311310213-2322000100100322"></a>

#### `domains.credentials.credential_name` property

Type: `"string"`. Computed.

Enter a unique name for the credentials used in API testing.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [login_endpoint](data-sources--api_testing--reference--group-001.md#canonical-0221111312210123-2010030233223322-3000002131202012-3033031111303313-0012021113313032-1232111121100330-2132230211331112-2231110212122030): complete subsection reference.

- [standard](data-sources--api_testing--reference--group-001.md#canonical-3002030031233231-3203321203210311-3323002110331313-0000211312223220-2331122312020232-0300312011303021-2313102000201302-2122100300103000): complete subsection reference.

<a id="canonical-1323002131113221-1230221322223220-1302203131311221-2122133201003322-0011323330132323-1311321103321231-0101022313013030-0332132022112202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.admin` properties

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-1103130220132111-1121033102212030-2032323111203320-0101012101310232-3300331302031232-1332131330231320-0212001110131013-1010030122002230)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-3002213323001033-3300321333010022-1122133200032233-1213022031201310-3110232103211233-0103222003311303-0323133133201020-2103123031102103)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-3131031132021133-3010303022131200-2033203230030023-3131100130120320-1203210003222110-1020112003233203-2133130033211033-0011012200000012)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-3300302021011323-1233011033232202-0001131123022203-0113010333011001-2021301130020132-0223321201313022-3300013030331020-3221333322300112)
- domains.credentials.admin

<a id="canonical-3023103201113330-2122020023132112-2200330010032112-0310222320111020-3020013203231311-3233111032332010-1200100201022023-3203203101302321"></a>

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

<a id="canonical-2323012032023313-2023111230121123-3012012100133321-3110102131113012-1312012331001023-3102320211110021-2123213113002301-1122312003302023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.api_key` properties

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-1103130220132111-1121033102212030-2032323111203320-0101012101310232-3300331302031232-1332131330231320-0212001110131013-1010030122002230)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-3002213323001033-3300321333010022-1122133200032233-1213022031201310-3110232103211233-0103222003311303-0323133133201020-2103123031102103)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-3131031132021133-3010303022131200-2033203230030023-3131100130120320-1203210003222110-1020112003233203-2133130033211033-0011012200000012)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-3300302021011323-1233011033232202-0001131123022203-0113010333011001-2021301130020132-0223321201313022-3300013030331020-3221333322300112)
- domains.credentials.api_key

<a id="canonical-2230331013031320-3302001333122331-1301111321023220-3132212111112033-1002222303023200-0131302021200232-1130210123233120-1012131101033221"></a>

Type: `"single"`. Computed.

API Key

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

<a id="canonical-0301330130032332-1230302111100230-2022111103122210-2212211031212132-2022000233122222-1032031202301322-3232021121000013-1133003210030233"></a>

### Direct properties for `domains.credentials.api_key`

<a id="canonical-3320203313320213-1303331210010331-3202120330001231-2333111300031003-3032300203310132-1001110022312030-3003300310102130-1022310102021120"></a>

#### `domains.credentials.api_key.key` property

Type: `"string"`. Computed.

Key. Cryptographic key material

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [value](data-sources--api_testing--reference--group-001.md#canonical-2033133221021103-2323110122222020-0103300020123113-2231302321001321-1102320020231113-1332111033311102-3201302311123202-3101312303232320): complete subsection reference.

<a id="canonical-2033133221021103-2323110122222020-0103300020123113-2231302321001321-1102320020231113-1332111033311102-3201302311123202-3101312303232320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.api_key.value` properties

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-1103130220132111-1121033102212030-2032323111203320-0101012101310232-3300331302031232-1332131330231320-0212001110131013-1010030122002230)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-3002213323001033-3300321333010022-1122133200032233-1213022031201310-3110232103211233-0103222003311303-0323133133201020-2103123031102103)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-3131031132021133-3010303022131200-2033203230030023-3131100130120320-1203210003222110-1020112003233203-2133130033211033-0011012200000012)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-3300302021011323-1233011033232202-0001131123022203-0113010333011001-2021301130020132-0223321201313022-3300013030331020-3221333322300112)
- [domains.credentials.api_key](data-sources--api_testing--reference--group-001.md#canonical-2323012032023313-2023111230121123-3012012100133321-3110102131113012-1312012331001023-3102320211110021-2123213113002301-1122312003302023)
- domains.credentials.api_key.value

<a id="canonical-2232301212332330-3312211002311032-0200133010001322-1012222003033000-0321112212130210-0132300031322031-3230321011002002-1003213033221013"></a>

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

<a id="canonical-2210022210003111-2132222103332212-1222103220130210-1221323113112332-0222220112220020-2332233232331202-2213232323000233-1302332230203313"></a>

### Direct properties for `domains.credentials.api_key.value`

- [blindfold_secret_info](data-sources--api_testing--reference--group-001.md#canonical-3202113212311301-1233231231220212-0102033031210012-3332132201231212-3033320331302011-1022303312322132-1121032122203003-2303203202231211): complete subsection reference.

- [clear_secret_info](data-sources--api_testing--reference--group-001.md#canonical-1313320330100020-0132213202021012-0133001322202321-2222303123201132-3022202321210121-1222031323320321-3121313211332123-0201310323000102): complete subsection reference.

<a id="canonical-3202113212311301-1233231231220212-0102033031210012-3332132201231212-3033320331302011-1022303312322132-1121032122203003-2303203202231211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.api_key.value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-1103130220132111-1121033102212030-2032323111203320-0101012101310232-3300331302031232-1332131330231320-0212001110131013-1010030122002230)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-3002213323001033-3300321333010022-1122133200032233-1213022031201310-3110232103211233-0103222003311303-0323133133201020-2103123031102103)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-3131031132021133-3010303022131200-2033203230030023-3131100130120320-1203210003222110-1020112003233203-2133130033211033-0011012200000012)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-3300302021011323-1233011033232202-0001131123022203-0113010333011001-2021301130020132-0223321201313022-3300013030331020-3221333322300112)
- [domains.credentials.api_key](data-sources--api_testing--reference--group-001.md#canonical-2323012032023313-2023111230121123-3012012100133321-3110102131113012-1312012331001023-3102320211110021-2123213113002301-1122312003302023)
- [domains.credentials.api_key.value](data-sources--api_testing--reference--group-001.md#canonical-2033133221021103-2323110122222020-0103300020123113-2231302321001321-1102320020231113-1332111033311102-3201302311123202-3101312303232320)
- domains.credentials.api_key.value.blindfold_secret_info

<a id="canonical-1120010202233203-3233031020303221-0112012011210030-2113113112000200-3223300233330332-3333001013323111-3301103003022212-0221312221133113"></a>

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

<a id="canonical-1003310111232201-3023112221121132-2003222133020330-0032310110013131-1102302312320310-2132213210320120-2032303322320331-2021312013132102"></a>

### Direct properties for `domains.credentials.api_key.value.blindfold_secret_info`

<a id="canonical-2021230212111110-2223233213220231-0330131133333333-2331031001300230-3030122213021303-0023021112110123-3030123230321321-2002031122221120"></a>

#### `domains.credentials.api_key.value.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2213223233202330-0231312121231322-1033020331101013-0011200033111333-0133323222232311-0311233232303312-2230221220022021-2331013032131121"></a>

<a id="canonical-3220331002020231-0210020301230223-1302021201032212-3201300212132332-0132103110333212-3113302230102101-0311032132300001-2333232223212301"></a>

#### `domains.credentials.api_key.value.blindfold_secret_info.location` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2123122031233302-2020232201310110-3333120233323222-1331102332302133-3000010031111201-3001133033221202-2111112131312213-0210110001001201"></a>

<a id="canonical-2203331111311210-1130222010101003-3321111321230003-1102211013120210-2211002113003323-1113110321213220-2213101123220111-0133031303202302"></a>

#### `domains.credentials.api_key.value.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1313320330100020-0132213202021012-0133001322202321-2222303123201132-3022202321210121-1222031323320321-3121313211332123-0201310323000102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.api_key.value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-1103130220132111-1121033102212030-2032323111203320-0101012101310232-3300331302031232-1332131330231320-0212001110131013-1010030122002230)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-3002213323001033-3300321333010022-1122133200032233-1213022031201310-3110232103211233-0103222003311303-0323133133201020-2103123031102103)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-3131031132021133-3010303022131200-2033203230030023-3131100130120320-1203210003222110-1020112003233203-2133130033211033-0011012200000012)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-3300302021011323-1233011033232202-0001131123022203-0113010333011001-2021301130020132-0223321201313022-3300013030331020-3221333322300112)
- [domains.credentials.api_key](data-sources--api_testing--reference--group-001.md#canonical-2323012032023313-2023111230121123-3012012100133321-3110102131113012-1312012331001023-3102320211110021-2123213113002301-1122312003302023)
- [domains.credentials.api_key.value](data-sources--api_testing--reference--group-001.md#canonical-2033133221021103-2323110122222020-0103300020123113-2231302321001321-1102320020231113-1332111033311102-3201302311123202-3101312303232320)
- domains.credentials.api_key.value.clear_secret_info

<a id="canonical-1111232200223002-3312010012121301-1232003301211122-1002201021031332-3232331010023230-0123010211331322-3310213220102230-2101130132223210"></a>

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

<a id="canonical-1001011002120302-3311221100321300-1301213303030032-1230203223301032-3200101330023100-0232021201222212-0222033201330332-0310213202100230"></a>

### Direct properties for `domains.credentials.api_key.value.clear_secret_info`

<a id="canonical-1030022331122011-0303013112120103-2302212322331122-3020303202333130-3130113212313021-2320030130301300-1323022321022323-3332123011022022"></a>

#### `domains.credentials.api_key.value.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0133011001031223-3001302133321002-3121200223002131-3202033113010013-1313311331100313-3223323333130010-1220103000322122-2000221031011320"></a>

<a id="canonical-2330323121032120-2231001122230203-1121023331033120-3100111232111122-3313221202131101-2300122132330222-1311033121322120-1300201212301001"></a>

#### `domains.credentials.api_key.value.clear_secret_info.url` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2222012103200021-2301131133131111-1001100303300203-1312321200133203-1112233030011310-1033130110010302-0020033113311323-1103302310233122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.basic_auth` properties

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-1103130220132111-1121033102212030-2032323111203320-0101012101310232-3300331302031232-1332131330231320-0212001110131013-1010030122002230)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-3002213323001033-3300321333010022-1122133200032233-1213022031201310-3110232103211233-0103222003311303-0323133133201020-2103123031102103)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-3131031132021133-3010303022131200-2033203230030023-3131100130120320-1203210003222110-1020112003233203-2133130033211033-0011012200000012)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-3300302021011323-1233011033232202-0001131123022203-0113010333011001-2021301130020132-0223321201313022-3300013030331020-3221333322300112)
- domains.credentials.basic_auth

<a id="canonical-0310131311023121-0002222131222213-0332033220223312-3021233233223110-0302110131121311-0212202233132113-1322132331023200-1110100232213123"></a>

Type: `"single"`. Computed.

Basic Authentication.

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

<a id="canonical-0122203100020213-1121100131121113-0030220103133230-0010031021203123-2221010323110130-2113230012122002-0003111300032002-0331230121300212"></a>

### Direct properties for `domains.credentials.basic_auth`

- [password](data-sources--api_testing--reference--group-001.md#canonical-1110003032023122-2000201110200230-0320101020230033-2311010103202303-0021230333313321-3113200303121203-0213030333023202-1132021011011031): complete subsection reference.

<a id="canonical-3332010120322002-1120231332211033-2000230033333230-2310031331322310-3112220220132333-1011301033010201-2303200111223303-1003033120021200"></a>

<a id="canonical-1013221001201210-1100213000131220-2311101221120122-2121032311321030-1212122311303322-2112111202232232-1302301102203121-1022123322302202"></a>

#### `domains.credentials.basic_auth.user` property

Type: `"string"`. Computed.

User. Configuration parameter for user

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-1110003032023122-2000201110200230-0320101020230033-2311010103202303-0021230333313321-3113200303121203-0213030333023202-1132021011011031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.basic_auth.password` properties

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-1103130220132111-1121033102212030-2032323111203320-0101012101310232-3300331302031232-1332131330231320-0212001110131013-1010030122002230)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-3002213323001033-3300321333010022-1122133200032233-1213022031201310-3110232103211233-0103222003311303-0323133133201020-2103123031102103)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-3131031132021133-3010303022131200-2033203230030023-3131100130120320-1203210003222110-1020112003233203-2133130033211033-0011012200000012)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-3300302021011323-1233011033232202-0001131123022203-0113010333011001-2021301130020132-0223321201313022-3300013030331020-3221333322300112)
- [domains.credentials.basic_auth](data-sources--api_testing--reference--group-001.md#canonical-2222012103200021-2301131133131111-1001100303300203-1312321200133203-1112233030011310-1033130110010302-0020033113311323-1103302310233122)
- domains.credentials.basic_auth.password

<a id="canonical-3033320110213213-1213023013122222-1230333231233210-2332011223222122-2021102320023103-1320322033132002-3310302211212033-2220103312221231"></a>

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

<a id="canonical-2013332311102333-0303131330332222-2330333022132131-1203000232300112-1210003113220132-3231021121030112-3110211323030101-1331310300322132"></a>

### Direct properties for `domains.credentials.basic_auth.password`

- [blindfold_secret_info](data-sources--api_testing--reference--group-001.md#canonical-1302202121201313-1333323201223222-3333002012213321-2001131122230032-1000013230330312-0202233331023021-0333222010322101-2300122022131222): complete subsection reference.

- [clear_secret_info](data-sources--api_testing--reference--group-001.md#canonical-1131031303102011-3202301021233203-0000110333212331-0223201020103322-1302232000220332-3132133112132020-3010011202030302-3303312212213122): complete subsection reference.

<a id="canonical-1302202121201313-1333323201223222-3333002012213321-2001131122230032-1000013230330312-0202233331023021-0333222010322101-2300122022131222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.basic_auth.password.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-1103130220132111-1121033102212030-2032323111203320-0101012101310232-3300331302031232-1332131330231320-0212001110131013-1010030122002230)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-3002213323001033-3300321333010022-1122133200032233-1213022031201310-3110232103211233-0103222003311303-0323133133201020-2103123031102103)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-3131031132021133-3010303022131200-2033203230030023-3131100130120320-1203210003222110-1020112003233203-2133130033211033-0011012200000012)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-3300302021011323-1233011033232202-0001131123022203-0113010333011001-2021301130020132-0223321201313022-3300013030331020-3221333322300112)
- [domains.credentials.basic_auth](data-sources--api_testing--reference--group-001.md#canonical-2222012103200021-2301131133131111-1001100303300203-1312321200133203-1112233030011310-1033130110010302-0020033113311323-1103302310233122)
- [domains.credentials.basic_auth.password](data-sources--api_testing--reference--group-001.md#canonical-1110003032023122-2000201110200230-0320101020230033-2311010103202303-0021230333313321-3113200303121203-0213030333023202-1132021011011031)
- domains.credentials.basic_auth.password.blindfold_secret_info

<a id="canonical-1322022233223311-3132201112330322-3213310303302000-3003310310012021-0112333233112201-3202121101332322-0221011120001302-2121203011003110"></a>

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

<a id="canonical-0132311133101103-2121000131031323-0013000020030210-1001200212111200-3300300222113333-3031033311323121-3233132000202330-0232120230323202"></a>

### Direct properties for `domains.credentials.basic_auth.password.blindfold_secret_info`

<a id="canonical-3220332201333333-1332311220101333-1030232221321322-2032201331213313-3102200131333011-2120120200323311-2303221000101333-0011123203110232"></a>

#### `domains.credentials.basic_auth.password.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3210031122132101-2103331033012001-1301111230203222-0101230301032312-2320330113322213-3113220320220322-0123211313311211-3133113220323322"></a>

<a id="canonical-2111330020110102-3030003001101103-1300300133101011-0322032122111312-1201123111021333-3133331232233110-2103330211202020-0001033132333012"></a>

#### `domains.credentials.basic_auth.password.blindfold_secret_info.location` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0213122233331010-3023033000223223-2033210332131132-2111202111332321-2321212332133112-1112021123302101-3032033210233230-1031302213003233"></a>

<a id="canonical-0321302231233222-3110020130323232-1013332233320133-0203102002012333-0131131202132230-0133122023323211-1132101201312022-3100202313203022"></a>

#### `domains.credentials.basic_auth.password.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1131031303102011-3202301021233203-0000110333212331-0223201020103322-1302232000220332-3132133112132020-3010011202030302-3303312212213122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.basic_auth.password.clear_secret_info` properties

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-1103130220132111-1121033102212030-2032323111203320-0101012101310232-3300331302031232-1332131330231320-0212001110131013-1010030122002230)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-3002213323001033-3300321333010022-1122133200032233-1213022031201310-3110232103211233-0103222003311303-0323133133201020-2103123031102103)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-3131031132021133-3010303022131200-2033203230030023-3131100130120320-1203210003222110-1020112003233203-2133130033211033-0011012200000012)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-3300302021011323-1233011033232202-0001131123022203-0113010333011001-2021301130020132-0223321201313022-3300013030331020-3221333322300112)
- [domains.credentials.basic_auth](data-sources--api_testing--reference--group-001.md#canonical-2222012103200021-2301131133131111-1001100303300203-1312321200133203-1112233030011310-1033130110010302-0020033113311323-1103302310233122)
- [domains.credentials.basic_auth.password](data-sources--api_testing--reference--group-001.md#canonical-1110003032023122-2000201110200230-0320101020230033-2311010103202303-0021230333313321-3113200303121203-0213030333023202-1132021011011031)
- domains.credentials.basic_auth.password.clear_secret_info

<a id="canonical-2100333031032003-3032123112211020-2031231313200212-1203111312223200-1233230310322012-0130021030311310-1233133130300320-1220231321030022"></a>

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

<a id="canonical-1303213033002133-0101213133120201-3013031031320032-1111022013203333-2101331321320122-2132213210023330-3321230331322313-3202121033313122"></a>

### Direct properties for `domains.credentials.basic_auth.password.clear_secret_info`

<a id="canonical-1213323300002222-0311123101203030-2130313103121130-2221222120021122-1023322221331022-3100210200113113-3230011213003021-2230213212110221"></a>

#### `domains.credentials.basic_auth.password.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2102311000113002-0010012113221100-0313231030232202-0330023001231232-1111221332133333-1002002001132123-1130301212323121-0221110012021011"></a>

<a id="canonical-2130202031213132-1013133320101100-3032011013121000-3303302302211311-1103131233032303-3231213200233301-0022312003132133-2210321000222230"></a>

#### `domains.credentials.basic_auth.password.clear_secret_info.url` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1002300133110320-3332210320021021-1111101131301023-0231213012302320-0122202302002011-0210210210210113-3032002033021311-1031201330223013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.bearer_token` properties

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-1103130220132111-1121033102212030-2032323111203320-0101012101310232-3300331302031232-1332131330231320-0212001110131013-1010030122002230)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-3002213323001033-3300321333010022-1122133200032233-1213022031201310-3110232103211233-0103222003311303-0323133133201020-2103123031102103)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-3131031132021133-3010303022131200-2033203230030023-3131100130120320-1203210003222110-1020112003233203-2133130033211033-0011012200000012)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-3300302021011323-1233011033232202-0001131123022203-0113010333011001-2021301130020132-0223321201313022-3300013030331020-3221333322300112)
- domains.credentials.bearer_token

<a id="canonical-3333302033332302-0303220303222322-0020112202322002-2003130130030300-2002200130330221-0110100211003101-1023311321213023-2310122130030010"></a>

Type: `"single"`. Computed.

Configuration parameter for bearer token.

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

<a id="canonical-3302111022123010-2321021332111223-3321120031213303-2000122102313010-2321132202031000-3223211130113122-0002231222133232-0002023200130323"></a>

### Direct properties for `domains.credentials.bearer_token`

- [token](data-sources--api_testing--reference--group-001.md#canonical-2321300003011131-3322013221100103-0031022101130311-3101120310232002-0011110233101202-1132233333232000-2030130212012000-3022013123230123): complete subsection reference.

<a id="canonical-2321300003011131-3322013221100103-0031022101130311-3101120310232002-0011110233101202-1132233333232000-2030130212012000-3022013123230123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.bearer_token.token` properties

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-1103130220132111-1121033102212030-2032323111203320-0101012101310232-3300331302031232-1332131330231320-0212001110131013-1010030122002230)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-3002213323001033-3300321333010022-1122133200032233-1213022031201310-3110232103211233-0103222003311303-0323133133201020-2103123031102103)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-3131031132021133-3010303022131200-2033203230030023-3131100130120320-1203210003222110-1020112003233203-2133130033211033-0011012200000012)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-3300302021011323-1233011033232202-0001131123022203-0113010333011001-2021301130020132-0223321201313022-3300013030331020-3221333322300112)
- [domains.credentials.bearer_token](data-sources--api_testing--reference--group-001.md#canonical-1002300133110320-3332210320021021-1111101131301023-0231213012302320-0122202302002011-0210210210210113-3032002033021311-1031201330223013)
- domains.credentials.bearer_token.token

<a id="canonical-1020202220311202-2100310222031310-2233320212101303-3202103223032313-1113322032222012-0110120001233212-3213220230011121-1301232112111222"></a>

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

<a id="canonical-3010032013301231-3300112103130303-0012130330001320-2201023321201220-0230303110210022-2201221232312212-1121133303100321-0031223303230330"></a>

### Direct properties for `domains.credentials.bearer_token.token`

- [blindfold_secret_info](data-sources--api_testing--reference--group-001.md#canonical-1032133100230011-3200332110132202-0120303023012213-1032133220102220-2113120030210210-0130300010023331-2013021103301230-1132311003223220): complete subsection reference.

- [clear_secret_info](data-sources--api_testing--reference--group-001.md#canonical-3021231312020223-1333110231202332-3010232002120011-1220122332032012-3211132022132311-1323030102111321-3033333023232133-2022130033123303): complete subsection reference.

<a id="canonical-1032133100230011-3200332110132202-0120303023012213-1032133220102220-2113120030210210-0130300010023331-2013021103301230-1132311003223220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.bearer_token.token.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-1103130220132111-1121033102212030-2032323111203320-0101012101310232-3300331302031232-1332131330231320-0212001110131013-1010030122002230)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-3002213323001033-3300321333010022-1122133200032233-1213022031201310-3110232103211233-0103222003311303-0323133133201020-2103123031102103)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-3131031132021133-3010303022131200-2033203230030023-3131100130120320-1203210003222110-1020112003233203-2133130033211033-0011012200000012)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-3300302021011323-1233011033232202-0001131123022203-0113010333011001-2021301130020132-0223321201313022-3300013030331020-3221333322300112)
- [domains.credentials.bearer_token](data-sources--api_testing--reference--group-001.md#canonical-1002300133110320-3332210320021021-1111101131301023-0231213012302320-0122202302002011-0210210210210113-3032002033021311-1031201330223013)
- [domains.credentials.bearer_token.token](data-sources--api_testing--reference--group-001.md#canonical-2321300003011131-3322013221100103-0031022101130311-3101120310232002-0011110233101202-1132233333232000-2030130212012000-3022013123230123)
- domains.credentials.bearer_token.token.blindfold_secret_info

<a id="canonical-2021112330110020-0221310132202222-1020100233222100-3023230103023012-2021303320220211-2110012002213223-0023203333131212-2200002311331223"></a>

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

<a id="canonical-1111132022103203-3230011003203211-3220211212121123-3123020031112202-0222020023003031-1210132323222022-1013030302022000-3032122232023031"></a>

### Direct properties for `domains.credentials.bearer_token.token.blindfold_secret_info`

<a id="canonical-3310100322211111-3021201122133030-3231233333200200-1201003103223111-1212202030301213-3102033310320303-3202130200232013-2333333003312222"></a>

#### `domains.credentials.bearer_token.token.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1200310021010233-0221121212033220-1213310322201223-2000031101222221-2120233323012030-2123212021103321-2301211300303301-1003031221020311"></a>

<a id="canonical-0303232121033300-0223300133333011-0211030111200213-0012302221200101-1113030303220221-3303032200331233-2213202303132212-1333123232313120"></a>

#### `domains.credentials.bearer_token.token.blindfold_secret_info.location` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0232210222032220-0130323033103022-0000323202320001-3020323312031210-2220313200222200-0233313230322202-1223100212111031-2003103103010200"></a>

<a id="canonical-0111203301112101-3212220312323213-2312203122103023-0232111012023103-0203310333113101-3032333200102301-3032303000222230-2122021131133323"></a>

#### `domains.credentials.bearer_token.token.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3021231312020223-1333110231202332-3010232002120011-1220122332032012-3211132022132311-1323030102111321-3033333023232133-2022130033123303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.bearer_token.token.clear_secret_info` properties

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-1103130220132111-1121033102212030-2032323111203320-0101012101310232-3300331302031232-1332131330231320-0212001110131013-1010030122002230)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-3002213323001033-3300321333010022-1122133200032233-1213022031201310-3110232103211233-0103222003311303-0323133133201020-2103123031102103)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-3131031132021133-3010303022131200-2033203230030023-3131100130120320-1203210003222110-1020112003233203-2133130033211033-0011012200000012)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-3300302021011323-1233011033232202-0001131123022203-0113010333011001-2021301130020132-0223321201313022-3300013030331020-3221333322300112)
- [domains.credentials.bearer_token](data-sources--api_testing--reference--group-001.md#canonical-1002300133110320-3332210320021021-1111101131301023-0231213012302320-0122202302002011-0210210210210113-3032002033021311-1031201330223013)
- [domains.credentials.bearer_token.token](data-sources--api_testing--reference--group-001.md#canonical-2321300003011131-3322013221100103-0031022101130311-3101120310232002-0011110233101202-1132233333232000-2030130212012000-3022013123230123)
- domains.credentials.bearer_token.token.clear_secret_info

<a id="canonical-1210312332030300-3113220123300030-3222331000320331-2203330322010322-3013020003331222-1022130112303122-1103220003302313-3333112202312200"></a>

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

<a id="canonical-2112210121131121-1201320203320321-0301213022030223-2120312000201003-0321030303232230-1031311321210123-1330320323032010-0103012221202223"></a>

### Direct properties for `domains.credentials.bearer_token.token.clear_secret_info`

<a id="canonical-3033123022113031-2301330232230310-0300001021112211-3201323302302131-1202311131331313-3022000023231112-1312302211101322-3123133020201222"></a>

#### `domains.credentials.bearer_token.token.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2333313001303112-0300302203302321-0033310020111000-0123011011221000-0211032303031011-0231330213110102-0020332212332230-3010012333131131"></a>

<a id="canonical-3321013302022301-3110032323200313-2102320300212311-3231133112002110-3233321302312333-2101121130302212-1211220013120300-3101303002211002"></a>

#### `domains.credentials.bearer_token.token.clear_secret_info.url` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0221111312210123-2010030233223322-3000002131202012-3033031111303313-0012021113313032-1232111121100330-2132230211331112-2231110212122030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.login_endpoint` properties

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-1103130220132111-1121033102212030-2032323111203320-0101012101310232-3300331302031232-1332131330231320-0212001110131013-1010030122002230)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-3002213323001033-3300321333010022-1122133200032233-1213022031201310-3110232103211233-0103222003311303-0323133133201020-2103123031102103)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-3131031132021133-3010303022131200-2033203230030023-3131100130120320-1203210003222110-1020112003233203-2133130033211033-0011012200000012)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-3300302021011323-1233011033232202-0001131123022203-0113010333011001-2021301130020132-0223321201313022-3300013030331020-3221333322300112)
- domains.credentials.login_endpoint

<a id="canonical-1312310200200210-0302220330021121-1221001012310121-0103213302033013-1221213213322011-1123313002131102-1033011322100130-2232213100132323"></a>

Type: `"single"`. Computed.

Login Endpoint.

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

<a id="canonical-0002231311020003-3302231332331120-2200013101302001-0303303302231021-3210001203202003-1103031131323110-1331212112322230-1301233213231100"></a>

### Direct properties for `domains.credentials.login_endpoint`

- [json_payload](data-sources--api_testing--reference--group-001.md#canonical-1032020103220110-1223321110010000-0130003201003131-0201302222102212-0322033310232232-1300022101013023-2023002120231330-3211103232321133): complete subsection reference.

<a id="canonical-0332031302022303-1301300000320132-1321211231110112-1312131033221231-3203222200332010-0232013100022200-0333011312311223-3221011213121213"></a>

<a id="canonical-0133030032110111-1313300300320031-2110103212100333-1030111012232200-0013312202000011-1311210103331113-3103033103212023-3121303312231312"></a>

#### `domains.credentials.login_endpoint.method` property

Type: `"string"`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

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

<a id="canonical-1133330331013211-1003330310012220-0331220312313132-2123233112211031-1033303112312111-2322310032322131-2030101033202300-2222330133230300"></a>

<a id="canonical-1203320311023013-2113230013012011-2001212330211313-1003120020101210-3020321333111321-1120102222321310-2122002111132210-3332003110300301"></a>

#### `domains.credentials.login_endpoint.path` property

Type: `"string"`. Computed.

Path. URL path for the endpoint

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "1024"
  }
}
```

<a id="canonical-0130221103112313-0302121130320211-0320120013203100-0211030302323021-1120302103333231-1231011012222220-1131022022012212-0213311033232022"></a>

<a id="canonical-0112123032003302-2132313013002022-1220133332310020-1130030031000303-1132320233212100-0210023220113121-2103221330103013-2112310121203301"></a>

#### `domains.credentials.login_endpoint.token_response_key` property

Type: `"string"`. Computed.

Configuration parameter for token response key.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1032020103220110-1223321110010000-0130003201003131-0201302222102212-0322033310232232-1300022101013023-2023002120231330-3211103232321133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.login_endpoint.json_payload` properties

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-1103130220132111-1121033102212030-2032323111203320-0101012101310232-3300331302031232-1332131330231320-0212001110131013-1010030122002230)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-3002213323001033-3300321333010022-1122133200032233-1213022031201310-3110232103211233-0103222003311303-0323133133201020-2103123031102103)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-3131031132021133-3010303022131200-2033203230030023-3131100130120320-1203210003222110-1020112003233203-2133130033211033-0011012200000012)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-3300302021011323-1233011033232202-0001131123022203-0113010333011001-2021301130020132-0223321201313022-3300013030331020-3221333322300112)
- [domains.credentials.login_endpoint](data-sources--api_testing--reference--group-001.md#canonical-0221111312210123-2010030233223322-3000002131202012-3033031111303313-0012021113313032-1232111121100330-2132230211331112-2231110212122030)
- domains.credentials.login_endpoint.json_payload

<a id="canonical-2303122032121013-1111003101120011-1013231300002312-3232133123333310-0322320001032020-2311022132210320-2220111333121322-0130103110031113"></a>

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

<a id="canonical-2133312200002121-3231223002220320-1131021303012222-3332232300210013-1001213202123311-3030230333110210-2323231300323130-0032302110001201"></a>

### Direct properties for `domains.credentials.login_endpoint.json_payload`

- [blindfold_secret_info](data-sources--api_testing--reference--group-001.md#canonical-3331221132100012-0232102030221322-0133033003121003-3310010000231201-1132003312101232-2121032220022011-1102300212222000-1031222103012033): complete subsection reference.

- [clear_secret_info](data-sources--api_testing--reference--group-001.md#canonical-2031221303032320-3032320030331003-2113313300030331-0011031311122010-2030101312310203-1000303002033021-1232001110113120-3002033202202210): complete subsection reference.

<a id="canonical-3331221132100012-0232102030221322-0133033003121003-3310010000231201-1132003312101232-2121032220022011-1102300212222000-1031222103012033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.login_endpoint.json_payload.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-1103130220132111-1121033102212030-2032323111203320-0101012101310232-3300331302031232-1332131330231320-0212001110131013-1010030122002230)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-3002213323001033-3300321333010022-1122133200032233-1213022031201310-3110232103211233-0103222003311303-0323133133201020-2103123031102103)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-3131031132021133-3010303022131200-2033203230030023-3131100130120320-1203210003222110-1020112003233203-2133130033211033-0011012200000012)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-3300302021011323-1233011033232202-0001131123022203-0113010333011001-2021301130020132-0223321201313022-3300013030331020-3221333322300112)
- [domains.credentials.login_endpoint](data-sources--api_testing--reference--group-001.md#canonical-0221111312210123-2010030233223322-3000002131202012-3033031111303313-0012021113313032-1232111121100330-2132230211331112-2231110212122030)
- [domains.credentials.login_endpoint.json_payload](data-sources--api_testing--reference--group-001.md#canonical-1032020103220110-1223321110010000-0130003201003131-0201302222102212-0322033310232232-1300022101013023-2023002120231330-3211103232321133)
- domains.credentials.login_endpoint.json_payload.blindfold_secret_info

<a id="canonical-0000123313012110-1300031230023111-1233232321101312-1231303220312201-1133311320003000-0220302120013321-1001221323233123-1032011311230333"></a>

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

<a id="canonical-2203103113010013-0131301213213203-2002010132200300-1323020230211031-0330111002010131-3130200113220110-3301333222233203-0210333313112203"></a>

### Direct properties for `domains.credentials.login_endpoint.json_payload.blindfold_secret_info`

<a id="canonical-0232221201203200-2020330122013330-2212112230022332-0100033311101223-1020212203002113-2220323323020012-2110000131033312-3031223221202202"></a>

#### `domains.credentials.login_endpoint.json_payload.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2220010131300012-2030330110120111-3111003100201022-0110021231032312-3301212103022220-3022303310000130-1113312323233330-3331121312213031"></a>

<a id="canonical-2122022300332111-3123223223023003-0111330223022031-0211332211033110-1331311023022020-2010300231033013-3310131120023112-1013130010031112"></a>

#### `domains.credentials.login_endpoint.json_payload.blindfold_secret_info.location` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3001132322213012-0002230110201013-0033102213331212-3103201220322032-3311321223013223-0000001322230132-2303231113232301-0213123030023101"></a>

<a id="canonical-2020103030111212-2031202111203233-0311320023032331-3113212222233122-3131120322210331-2321321010020030-1232000223220013-0032311011230133"></a>

#### `domains.credentials.login_endpoint.json_payload.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2031221303032320-3032320030331003-2113313300030331-0011031311122010-2030101312310203-1000303002033021-1232001110113120-3002033202202210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.login_endpoint.json_payload.clear_secret_info` properties

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-1103130220132111-1121033102212030-2032323111203320-0101012101310232-3300331302031232-1332131330231320-0212001110131013-1010030122002230)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-3002213323001033-3300321333010022-1122133200032233-1213022031201310-3110232103211233-0103222003311303-0323133133201020-2103123031102103)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-3131031132021133-3010303022131200-2033203230030023-3131100130120320-1203210003222110-1020112003233203-2133130033211033-0011012200000012)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-3300302021011323-1233011033232202-0001131123022203-0113010333011001-2021301130020132-0223321201313022-3300013030331020-3221333322300112)
- [domains.credentials.login_endpoint](data-sources--api_testing--reference--group-001.md#canonical-0221111312210123-2010030233223322-3000002131202012-3033031111303313-0012021113313032-1232111121100330-2132230211331112-2231110212122030)
- [domains.credentials.login_endpoint.json_payload](data-sources--api_testing--reference--group-001.md#canonical-1032020103220110-1223321110010000-0130003201003131-0201302222102212-0322033310232232-1300022101013023-2023002120231330-3211103232321133)
- domains.credentials.login_endpoint.json_payload.clear_secret_info

<a id="canonical-1122011112323321-3202200133103211-0131330322313321-1102320220110121-3223213122100220-1311220232020120-0120111132020321-2322301200303231"></a>

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

<a id="canonical-0013130333320202-3103011233101112-1033031002332320-3011110022013010-3002300311332303-3112030301130333-1332030300232101-1310000020300331"></a>

### Direct properties for `domains.credentials.login_endpoint.json_payload.clear_secret_info`

<a id="canonical-2300311012231330-2302331312200010-3002303003120330-1200121023033213-1001320121113322-3032311313221033-1232221101301301-2033032303032320"></a>

#### `domains.credentials.login_endpoint.json_payload.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1301013033220122-2023313333201103-1300121013131313-1113323112222300-2221133230032200-3213213303212100-2300300010022311-3101133012012322"></a>

<a id="canonical-0122221321013223-2332111110033102-3120030001331112-1003303221202021-1212331233133213-0103233323033200-3012113003100313-1321112231011302"></a>

#### `domains.credentials.login_endpoint.json_payload.clear_secret_info.url` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3002030031233231-3203321203210311-3323002110331313-0000211312223220-2331122312020232-0300312011303021-2313102000201302-2122100300103000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `domains.credentials.standard` properties

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-1103130220132111-1121033102212030-2032323111203320-0101012101310232-3300331302031232-1332131330231320-0212001110131013-1010030122002230)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-3002213323001033-3300321333010022-1122133200032233-1213022031201310-3110232103211233-0103222003311303-0323133133201020-2103123031102103)
- [domains](data-sources--api_testing--reference--group-001.md#canonical-3131031132021133-3010303022131200-2033203230030023-3131100130120320-1203210003222110-1020112003233203-2133130033211033-0011012200000012)
- [domains.credentials](data-sources--api_testing--reference--group-001.md#canonical-3300302021011323-1233011033232202-0001131123022203-0113010333011001-2021301130020132-0223321201313022-3300013030331020-3221333322300112)
- domains.credentials.standard

<a id="canonical-2001102331012111-0121132221113310-0231230111332111-3213320323120332-0130013101200330-1110213121132230-2232310030310303-0212032201111322"></a>

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

<a id="canonical-2332330000101223-2332033331213132-1301121202210033-3103023332020211-1332031203333033-1103110000210321-0012010013012201-0313032122013212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `every_day` properties

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-1103130220132111-1121033102212030-2032323111203320-0101012101310232-3300331302031232-1332131330231320-0212001110131013-1010030122002230)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-3002213323001033-3300321333010022-1122133200032233-1213022031201310-3110232103211233-0103222003311303-0323133133201020-2103123031102103)
- every_day

<a id="canonical-1002112133320131-1113103223332230-2103101203022311-3130200100223301-1323310023313311-1030101030222313-1301003002011112-0122033133232303"></a>

Type: `["object", {}]`. Computed.

\[OneOf: every\_day, every\_month, every\_week\] Enable this option

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

- [every_day](data-sources--api_testing--reference--group-001.md#canonical-1002112133320131-1113103223332230-2103101203022311-3130200100223301-1323310023313311-1030101030222313-1301003002011112-0122033133232303)
- [every_month](data-sources--api_testing--reference--group-001.md#canonical-2322320322313112-3031010302331223-0213123102313213-1131132122230130-2332120313200230-3102300213022012-2312331110302001-3021122300203310)
- [every_week](data-sources--api_testing--reference--group-001.md#canonical-0122300232033003-2000113000323033-3131200131033012-0131033010312033-2320122122222023-3220232113122130-0313222121102113-0333211000302330)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0122310103331021-2020002230323013-1010102020012020-0010131103300303-2132223303320021-2313113101112011-2022122010230101-3111013330111132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `every_month` properties

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-1103130220132111-1121033102212030-2032323111203320-0101012101310232-3300331302031232-1332131330231320-0212001110131013-1010030122002230)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-3002213323001033-3300321333010022-1122133200032233-1213022031201310-3110232103211233-0103222003311303-0323133133201020-2103123031102103)
- every_month

<a id="canonical-2322320322313112-3031010302331223-0213123102313213-1131132122230130-2332120313200230-3102300213022012-2312331110302001-3021122300203310"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for every month.

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

<a id="canonical-1113210123100310-3302013201121202-0001131003030011-3012002030222033-3121033100323211-0012003202031021-0012001301310130-2133332223112101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `every_week` properties

Breadcrumbs:

- [xcsh_api_testing](../data-sources/api_testing.md#canonical-1103130220132111-1121033102212030-2032323111203320-0101012101310232-3300331302031232-1332131330231320-0212001110131013-1010030122002230)
- [Property reference](data-sources--api_testing--reference--group-001.md#canonical-3002213323001033-3300321333010022-1122133200032233-1213022031201310-3110232103211233-0103222003311303-0323133133201020-2103123031102103)
- every_week

<a id="canonical-0122300232033003-2000113000323033-3131200131033012-0131033010312033-2320122122222023-3220232113122130-0313222121102113-0333211000302330"></a>

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
