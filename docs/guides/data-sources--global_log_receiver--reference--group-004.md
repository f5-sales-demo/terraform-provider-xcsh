---
page_title: "xcsh_global_log_receiver reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_global_log_receiver reference."
---

# xcsh_global_log_receiver reference

<a id="canonical-0123001111011103-1013010100301101-1212323323100111-3331023101321021-3022311012301132-0321303310102333-0001021212221210-2010110203012101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `new_relic_receiver.api_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [new_relic_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-1222203022011122-1213111132203013-1201312120332200-3233301331202013-3012133313300230-3320033231321122-3201112101101230-3232031212020130)
- [new_relic_receiver.api_key](data-sources--global_log_receiver--reference--group-003.md#canonical-2211310101013202-2223001023310323-3031310222032320-3233233120202231-0101322122330110-1132232033003313-0230330322030212-0213101112010303)
- new_relic_receiver.api_key.blindfold_secret_info

<a id="canonical-1231103320103313-0102121031223030-2020310333300022-2323333332210101-0002332301310230-0123212010130013-1211123003310003-1132310233032003"></a>

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

<a id="canonical-3130001030221312-2101200311301121-1032003021131113-1033033302313002-3312110210102120-0021221133032032-0301002013220302-2201121123122020"></a>

### Direct properties for `new_relic_receiver.api_key.blindfold_secret_info`

<a id="canonical-3132200230032303-0233301133220223-3032032223020331-2023113011302323-0012322003311013-2222230022123322-1101213120232013-1220321010031022"></a>

#### `new_relic_receiver.api_key.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0023020303322230-0331201222021023-1130131230003033-3003300320320332-1302012331211332-1101001013122122-1303121302022012-2210332312003021"></a>

<a id="canonical-2133032322111103-2221201203012230-3011331313032111-3033223321031330-3112213320000130-1101203132131033-3303330213030200-1110231303310213"></a>

#### `new_relic_receiver.api_key.blindfold_secret_info.location` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2303332202301320-2010103010213131-1311021030310322-0302330302103011-0302313123130101-0112023013123310-3110122123330021-3211210013102210"></a>

<a id="canonical-1203132010311032-0032331112321223-1211120031121212-1133122022003030-1221033320001102-2100323202202203-1010110133103130-3312313321203001"></a>

#### `new_relic_receiver.api_key.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1320221211133021-1033300103012120-1112030330032201-1031201332123101-1323200101300013-2222012210121020-2200130103311203-2032232322030331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `new_relic_receiver.api_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [new_relic_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-1222203022011122-1213111132203013-1201312120332200-3233301331202013-3012133313300230-3320033231321122-3201112101101230-3232031212020130)
- [new_relic_receiver.api_key](data-sources--global_log_receiver--reference--group-003.md#canonical-2211310101013202-2223001023310323-3031310222032320-3233233120202231-0101322122330110-1132232033003313-0230330322030212-0213101112010303)
- new_relic_receiver.api_key.clear_secret_info

<a id="canonical-2323323320223303-3313101021211333-2301212210131022-2201131020203033-3112310311333132-3213130123202303-2110301011331223-1312030212011001"></a>

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

<a id="canonical-0123231002001221-3031113202331203-1333100230103111-2002123203012030-3310231320131302-2233012221130003-3100012312022002-1120230332311132"></a>

### Direct properties for `new_relic_receiver.api_key.clear_secret_info`

<a id="canonical-0021210323230210-0122010012121103-2132130022322213-2231300003033233-0201110023113311-3123033000133323-2023013122330331-2232103200220021"></a>

#### `new_relic_receiver.api_key.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1211122201032320-0213133231011120-3332212030010222-3030013103103303-0231033203102122-3231023133011302-3311231202321332-0333230323212212"></a>

<a id="canonical-3320211223000031-2223020021011331-1231003010101203-3020010011303302-0221301120201313-0113001221211012-0000223102210132-1131331122021030"></a>

#### `new_relic_receiver.api_key.clear_secret_info.url` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3300010212112002-3213231331131321-2331110230220010-2012032031333333-0321230230113030-2111100323020013-3113212012233230-0110103010023113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `new_relic_receiver.eu` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [new_relic_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-1222203022011122-1213111132203013-1201312120332200-3233301331202013-3012133313300230-3320033231321122-3201112101101230-3232031212020130)
- new_relic_receiver.eu

<a id="canonical-2221130203203000-3332120110030003-2022323311313023-0133113023001023-1211033233002101-3201012312233330-0301012221233111-1210121212320222"></a>

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

<a id="canonical-3012303210122113-2313212321310023-3132111332302132-0011323210220302-3103132011013301-3021100321221320-2203201323101102-2002031320002012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `new_relic_receiver.us` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [new_relic_receiver](data-sources--global_log_receiver--reference--group-003.md#canonical-1222203022011122-1213111132203013-1201312120332200-3233301331202013-3012133313300230-3320033231321122-3201112101101230-3232031212020130)
- new_relic_receiver.us

<a id="canonical-3302121330210213-0222031021210020-0021100300123132-2333202101020311-2103330133110230-2021332010033231-2132332113232213-2200020021210010"></a>

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

<a id="canonical-3222301230023030-0001112112122320-3100220123111200-3201120222111311-1330032121023233-1201222312320311-1111331322201210-0133312301312020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ns_all` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- ns_all

<a id="canonical-1030120112010331-2230012313100203-3132133213000333-1232233330012101-3233322300200100-0011332212003301-2110203211000302-1101000123003013"></a>

Type: `["object", {}]`. Computed.

\[OneOf: ns\_all, ns\_current, ns\_list\] Enable this option

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

- [ns_all](data-sources--global_log_receiver--reference--group-004.md#canonical-1030120112010331-2230012313100203-3132133213000333-1232233330012101-3233322300200100-0011332212003301-2110203211000302-1101000123003013)
- [ns_current](data-sources--global_log_receiver--reference--group-004.md#canonical-2200320110300010-3121021332001011-2113033313031130-3212321003321130-2223103133200223-0313021133310311-2200213300131303-3022101333300002)
- [ns_list](data-sources--global_log_receiver--reference--group-004.md#canonical-3011000200300333-3300013030233023-1333110132103131-0000012030212002-1311112102002200-3110313322002301-2002302330131301-1220022231112010)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3213312020222012-2002220201233002-0031302031132022-2311120330010033-0031032202323202-0022102301032121-0321103001233122-0001120123201310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ns_current` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- ns_current

<a id="canonical-2200320110300010-3121021332001011-2113033313031130-3212321003321130-2223103133200223-0313021133310311-2200213300131303-3022101333300002"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-3303120033132301-0202221321231331-0130123312332131-0202000200222232-3021322301011113-1000013103130021-2133210101223030-0302330311022122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ns_list` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- ns_list

<a id="canonical-3011000200300333-3300013030233023-1333110132103131-0000012030212002-1311112102002200-3110313322002301-2002302330131301-1220022231112010"></a>

Type: `"single"`. Computed.

Namespace List. Namespace List.

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

<a id="canonical-0003000010302032-0310300030113233-1220121332321331-1322311121023303-0331323120230221-3323321133211233-2303320210013321-1110012211330102"></a>

### Direct properties for `ns_list`

<a id="canonical-2313233010230233-3102101333201310-0211211210103002-2031201102321311-2323001203320232-0300122212203320-2222220003322323-3130132302111012"></a>

#### `ns_list.namespaces` property

Type: `["list", "string"]`. Computed.

Namespaces. List of namespaces to stream logs for.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-2321021321022332-2312230013122032-1100112310202133-0112002300022123-2320312003031221-1122032002223023-3311231012221300-1220003101032203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- qradar_receiver

<a id="canonical-1203320123322112-0003220233010101-0221021023303230-1311333320232133-2123212302200211-3330303321331111-3000310130102210-0323203022331310"></a>

Type: `"single"`. Computed.

Configuration parameter for qradar receiver.

Additional upstream details:

Configuration for IBM QRadar endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

<a id="canonical-1303131222320021-3230010111332203-2021233332033231-0201213121320112-1130122221032232-0113112320210002-2200331203232113-2212131213131033"></a>

### Direct properties for `qradar_receiver`

- [batch](data-sources--global_log_receiver--reference--group-004.md#canonical-3322102130323103-1103332202002221-0010102101231230-1122230223022300-2103102121313033-1201330301311321-1311300322122211-2310302202321201): complete subsection reference.

- [compression](data-sources--global_log_receiver--reference--group-004.md#canonical-0133213013222302-3031001203313320-2133001120310023-1222022232203032-3010013010213131-2112130321123312-3132123030112013-2130321233302210): complete subsection reference.

- [no_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-1111323021212322-2212331021332333-0210102220332333-3123222111000123-2131031111023130-0021122202312123-3301311202111330-0212010322032212): complete subsection reference.

<a id="canonical-1312101322230223-1133303032220313-2202202302223232-2100022232122321-2210100103132030-3120331221320220-2000210101030023-0023130122110010"></a>

<a id="canonical-1003223232323233-2032202110013303-3213033230320330-0023220132013130-1320132332211331-3222220302133030-3033331030223023-2020123013310002"></a>

#### `qradar_receiver.uri` property

Type: `"string"`. Computed.

Log Source Collector URL is the URL of the IBM QRadar Log Source Collector to send logs to,.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-2211202023111213-2221103322110030-2331033132202121-2012102021320112-1030223320201110-0000011010011333-0011001031321121-2233020113012311): complete subsection reference.

<a id="canonical-3322102130323103-1103332202002221-0010102101231230-1122230223022300-2103102121313033-1201330301311321-1311300322122211-2310302202321201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.batch` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-2321021321022332-2312230013122032-1100112310202133-0112002300022123-2320312003031221-1122032002223023-3311231012221300-1220003101032203)
- qradar_receiver.batch

<a id="canonical-0120311201302311-3111210120211230-2122322000322312-3121321102322130-0201003100023323-1031310221123110-1113000112023303-1112322133203001"></a>

Type: `"single"`. Computed.

Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-batch_bytes": "[\"max_bytes\",\"max_bytes_disabled\"]",
  "x-ves-oneof-field-batch_events": "[\"max_events\",\"max_events_disabled\"]",
  "x-ves-oneof-field-batch_timeout": "[\"timeout_seconds\",\"timeout_seconds_default\"]"
}
```

<a id="canonical-0312023101003111-3100322321213030-0231021303021131-3321032011122023-2320103323332221-0322222312131332-2313213202210023-0030230330102321"></a>

### Direct properties for `qradar_receiver.batch`

<a id="canonical-3320011112101313-0001012302000200-2131121221330311-1333213212103130-1301023010220101-0001212321120222-3020201322000230-0213003113230003"></a>

#### `qradar_receiver.batch.max_bytes` property

Type: `"number"`. Computed.

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10485760,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 4096
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

- [max_bytes_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-2200132301231210-0332202300230101-1121030321210011-0330130023223313-0100001113121002-1123232200121023-2212111021302031-3133112210010032): complete subsection reference.

<a id="canonical-0113221213010220-3233322030121102-1210122003310100-1231231221201231-0003210210030030-2313311101322011-3221211230120202-1113020023013122"></a>

<a id="canonical-3221211202200211-3032133320232133-2332121113320023-2100333121310111-1020010122332213-3232033001100110-2030020010031133-0230320031100021"></a>

#### `qradar_receiver.batch.max_events` property

Type: `"number"`. Computed.

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 32
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  }
}
```

- [max_events_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-1021103113033123-2102331111033223-2231230310312330-0311210103301023-1320113210023031-0030223003232112-3323321213032103-3312322300110210): complete subsection reference.

<a id="canonical-1101001232300231-3202223232301000-2322123000330201-0333302321032133-1010322103121022-2021031201332232-2322210001131320-2100321320111120"></a>

<a id="canonical-1220022013223030-3010013230003233-2100100023103233-2211330213212302-2320321310201322-1003223232131232-3133112202202130-1233222002202201"></a>

#### `qradar_receiver.batch.timeout_seconds` property

Type: `"string"`. Computed.

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "uint64",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  }
}
```

- [timeout_seconds_default](data-sources--global_log_receiver--reference--group-004.md#canonical-1011233013203000-1221133102312110-3003203110130230-2130221213200310-3111113320011202-3201200100313310-0222201313021301-2312300112113231): complete subsection reference.

<a id="canonical-2200132301231210-0332202300230101-1121030321210011-0330130023223313-0100001113121002-1123232200121023-2212111021302031-3133112210010032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.batch.max_bytes_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-2321021321022332-2312230013122032-1100112310202133-0112002300022123-2320312003031221-1122032002223023-3311231012221300-1220003101032203)
- [qradar_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-3322102130323103-1103332202002221-0010102101231230-1122230223022300-2103102121313033-1201330301311321-1311300322122211-2310302202321201)
- qradar_receiver.batch.max_bytes_disabled

<a id="canonical-3202322303333020-3303230033331023-1201223030333100-3012213120120320-2200000201122112-1032223133233331-0123331313110312-0010320310011321"></a>

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

<a id="canonical-1021103113033123-2102331111033223-2231230310312330-0311210103301023-1320113210023031-0030223003232112-3323321213032103-3312322300110210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.batch.max_events_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-2321021321022332-2312230013122032-1100112310202133-0112002300022123-2320312003031221-1122032002223023-3311231012221300-1220003101032203)
- [qradar_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-3322102130323103-1103332202002221-0010102101231230-1122230223022300-2103102121313033-1201330301311321-1311300322122211-2310302202321201)
- qradar_receiver.batch.max_events_disabled

<a id="canonical-3101000111023132-2300323013100103-2131022300011113-2032323020123303-1100332032023321-1103303321201100-3112200011110133-0300302313213231"></a>

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

<a id="canonical-1011233013203000-1221133102312110-3003203110130230-2130221213200310-3111113320011202-3201200100313310-0222201313021301-2312300112113231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.batch.timeout_seconds_default` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-2321021321022332-2312230013122032-1100112310202133-0112002300022123-2320312003031221-1122032002223023-3311231012221300-1220003101032203)
- [qradar_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-3322102130323103-1103332202002221-0010102101231230-1122230223022300-2103102121313033-1201330301311321-1311300322122211-2310302202321201)
- qradar_receiver.batch.timeout_seconds_default

<a id="canonical-2013111300311330-1132223122303330-3211202100302101-3302133113020101-2220213212112020-2223201233000102-3320020330002220-0101103133132313"></a>

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

<a id="canonical-0133213013222302-3031001203313320-2133001120310023-1222022232203032-3010013010213131-2112130321123312-3132123030112013-2130321233302210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.compression` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-2321021321022332-2312230013122032-1100112310202133-0112002300022123-2320312003031221-1122032002223023-3311231012221300-1220003101032203)
- qradar_receiver.compression

<a id="canonical-0303032320311320-0122021323231033-0230220210031231-1100133131230133-2332121013033202-2023113010300300-0000121330111301-0032023111332002"></a>

Type: `"single"`. Computed.

Configuration parameter for compression.

Additional upstream details:

Compression Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-compression_choice": "[\"compression_default\",\"compression_gzip\",\"compression_none\"]"
}
```

<a id="canonical-2133122332221222-1020120232303331-3022301313011100-2033023010231113-3120303032123213-3023230331213001-0002030323002121-1132312203312131"></a>

### Direct properties for `qradar_receiver.compression`

- [compression_default](data-sources--global_log_receiver--reference--group-004.md#canonical-2000320120121120-0112110202122300-1330131332321202-0133101320032010-1311102023301213-1310233032121122-1033313111132120-3020330120013003): complete subsection reference.

- [compression_gzip](data-sources--global_log_receiver--reference--group-004.md#canonical-2021221123300320-3213123120231213-1233322013302011-0123311130313303-3022301210010301-3113213030223032-3133222203021201-2030220200110111): complete subsection reference.

- [compression_none](data-sources--global_log_receiver--reference--group-004.md#canonical-3023302011231113-2020311331130023-1211211120210031-0000111011130232-1203002223332021-2032100003232100-3323301311010202-3332122213300103): complete subsection reference.

<a id="canonical-2000320120121120-0112110202122300-1330131332321202-0133101320032010-1311102023301213-1310233032121122-1033313111132120-3020330120013003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.compression.compression_default` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-2321021321022332-2312230013122032-1100112310202133-0112002300022123-2320312003031221-1122032002223023-3311231012221300-1220003101032203)
- [qradar_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-0133213013222302-3031001203313320-2133001120310023-1222022232203032-3010013010213131-2112130321123312-3132123030112013-2130321233302210)
- qradar_receiver.compression.compression_default

<a id="canonical-1033113110202130-1121122202002322-3010013321332001-1323131133032231-3312203100210202-2113303332112311-1102322112221223-0001213121222003"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for compression default.

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

<a id="canonical-2021221123300320-3213123120231213-1233322013302011-0123311130313303-3022301210010301-3113213030223032-3133222203021201-2030220200110111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.compression.compression_gzip` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-2321021321022332-2312230013122032-1100112310202133-0112002300022123-2320312003031221-1122032002223023-3311231012221300-1220003101032203)
- [qradar_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-0133213013222302-3031001203313320-2133001120310023-1222022232203032-3010013010213131-2112130321123312-3132123030112013-2130321233302210)
- qradar_receiver.compression.compression_gzip

<a id="canonical-3233321202021213-1230302303101303-1311033122322023-0111302223230222-3033131101333112-0023100202311013-3031312221330122-0003132100302210"></a>

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

<a id="canonical-3023302011231113-2020311331130023-1211211120210031-0000111011130232-1203002223332021-2032100003232100-3323301311010202-3332122213300103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.compression.compression_none` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-2321021321022332-2312230013122032-1100112310202133-0112002300022123-2320312003031221-1122032002223023-3311231012221300-1220003101032203)
- [qradar_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-0133213013222302-3031001203313320-2133001120310023-1222022232203032-3010013010213131-2112130321123312-3132123030112013-2130321233302210)
- qradar_receiver.compression.compression_none

<a id="canonical-1030203010111302-0132031201202133-2130302302311311-1023013202200013-2031312021213332-3023310332121022-1201022101003101-3112332223303023"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for compression none.

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

<a id="canonical-1111323021212322-2212331021332333-0210102220332333-3123222111000123-2131031111023130-0021122202312123-3301311202111330-0212010322032212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.no_tls` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-2321021321022332-2312230013122032-1100112310202133-0112002300022123-2320312003031221-1122032002223023-3311231012221300-1220003101032203)
- qradar_receiver.no_tls

<a id="canonical-1311221010330131-0100030122201013-0123222200333120-2221022230023332-0133123010211132-0300332331311303-2303010213100302-1300031202112222"></a>

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

<a id="canonical-2211202023111213-2221103322110030-2331033132202121-2012102021320112-1030223320201110-0000011010011333-0011001031321121-2233020113012311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.use_tls` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-2321021321022332-2312230013122032-1100112310202133-0112002300022123-2320312003031221-1122032002223023-3311231012221300-1220003101032203)
- qradar_receiver.use_tls

<a id="canonical-2230220303112210-1332303332023323-0303300310322223-1021010001233001-2201000221131021-3113101123232032-1210221030233301-0323022013221302"></a>

Type: `"single"`. Computed.

TLS Parameters for client connection to the endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ca_choice": "[\"no_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-mtls_choice": "[\"mtls_disabled\",\"mtls_enable\"]",
  "x-ves-oneof-field-verify_certificate": "[\"disable_verify_certificate\",\"enable_verify_certificate\"]",
  "x-ves-oneof-field-verify_hostname": "[\"disable_verify_hostname\",\"enable_verify_hostname\"]"
}
```

<a id="canonical-0103100211120011-1220121033220122-0002110310330303-2030332020023011-2231110301122211-1002132110201113-3133002220313133-0000212033323330"></a>

### Direct properties for `qradar_receiver.use_tls`

- [disable_verify_certificate](data-sources--global_log_receiver--reference--group-004.md#canonical-3322301302112011-3033121332333012-0200331303130332-2222122221010101-0133233123032301-3122001023223333-3012013311301003-2311002323231113): complete subsection reference.

- [disable_verify_hostname](data-sources--global_log_receiver--reference--group-004.md#canonical-3033133203030300-0320022320111230-3121103002223122-2122001121220300-3311303212110211-3320101231000322-2200213311010132-1203010311323000): complete subsection reference.

- [enable_verify_certificate](data-sources--global_log_receiver--reference--group-004.md#canonical-0123123233012122-2113301330223202-2121020331223201-2212321333001230-1013330210322033-3012011223301023-1101221120323310-1121100232112320): complete subsection reference.

- [enable_verify_hostname](data-sources--global_log_receiver--reference--group-004.md#canonical-1020030030331112-0002020320132001-3321213311203211-2131010110102120-2030222032200022-1112312303103323-0211132122023330-2002320301320102): complete subsection reference.

- [mtls_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-3030223213022232-2013002132203010-1232301312001002-3012003330012130-1332310031333330-0313232120113211-0310122233131202-2332020002320113): complete subsection reference.

- [mtls_enable](data-sources--global_log_receiver--reference--group-004.md#canonical-3021121332121230-0330221320320012-3133230012110201-0303030120101330-0230220010223321-3300011031303222-0211300131003300-0131002021233113): complete subsection reference.

- [no_ca](data-sources--global_log_receiver--reference--group-004.md#canonical-1321113003023032-3133333102202200-3330133200020201-2210231001033103-1113301003303003-0302111220002111-0312310133310112-2033030022002322): complete subsection reference.

<a id="canonical-3323001230003112-2032112220213030-2223112202201301-2022033202030013-3103202312020201-0202121133200210-1023023220031012-3100220201130302"></a>

<a id="canonical-1231112300233130-3332022202122230-2212123210213313-1103202313133123-2133322223331322-0230133212233213-0102233203302021-3322301021013002"></a>

#### `qradar_receiver.use_tls.trusted_ca_url` property

Type: `"string"`. Computed.

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

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
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-3322301302112011-3033121332333012-0200331303130332-2222122221010101-0133233123032301-3122001023223333-3012013311301003-2311002323231113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.use_tls.disable_verify_certificate` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-2321021321022332-2312230013122032-1100112310202133-0112002300022123-2320312003031221-1122032002223023-3311231012221300-1220003101032203)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-2211202023111213-2221103322110030-2331033132202121-2012102021320112-1030223320201110-0000011010011333-0011001031321121-2233020113012311)
- qradar_receiver.use_tls.disable_verify_certificate

<a id="canonical-1010110110313311-3022111300213103-2021321201103031-0010220303330001-3120330212103221-3311222332331301-2313001223210133-3102112211323120"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable verify certificate.

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

<a id="canonical-3033133203030300-0320022320111230-3121103002223122-2122001121220300-3311303212110211-3320101231000322-2200213311010132-1203010311323000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.use_tls.disable_verify_hostname` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-2321021321022332-2312230013122032-1100112310202133-0112002300022123-2320312003031221-1122032002223023-3311231012221300-1220003101032203)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-2211202023111213-2221103322110030-2331033132202121-2012102021320112-1030223320201110-0000011010011333-0011001031321121-2233020113012311)
- qradar_receiver.use_tls.disable_verify_hostname

<a id="canonical-0033032130120233-2223021003101311-3101310203222221-2100201013200230-2312010311131120-1121120223111102-1210313300121101-0222021301301201"></a>

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

<a id="canonical-0123123233012122-2113301330223202-2121020331223201-2212321333001230-1013330210322033-3012011223301023-1101221120323310-1121100232112320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.use_tls.enable_verify_certificate` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-2321021321022332-2312230013122032-1100112310202133-0112002300022123-2320312003031221-1122032002223023-3311231012221300-1220003101032203)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-2211202023111213-2221103322110030-2331033132202121-2012102021320112-1030223320201110-0000011010011333-0011001031321121-2233020113012311)
- qradar_receiver.use_tls.enable_verify_certificate

<a id="canonical-3103312302111301-1231132032330333-3021000133010132-1332232332313323-1211033212323012-0210102013132320-3201021000322220-3212331133103003"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable verify certificate.

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

<a id="canonical-1020030030331112-0002020320132001-3321213311203211-2131010110102120-2030222032200022-1112312303103323-0211132122023330-2002320301320102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.use_tls.enable_verify_hostname` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-2321021321022332-2312230013122032-1100112310202133-0112002300022123-2320312003031221-1122032002223023-3311231012221300-1220003101032203)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-2211202023111213-2221103322110030-2331033132202121-2012102021320112-1030223320201110-0000011010011333-0011001031321121-2233020113012311)
- qradar_receiver.use_tls.enable_verify_hostname

<a id="canonical-2332003033202112-1303203011320320-2331011201020301-0322103001220033-0302010232213302-1223032102123333-0003313100001303-2303001121133003"></a>

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

<a id="canonical-3030223213022232-2013002132203010-1232301312001002-3012003330012130-1332310031333330-0313232120113211-0310122233131202-2332020002320113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.use_tls.mtls_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-2321021321022332-2312230013122032-1100112310202133-0112002300022123-2320312003031221-1122032002223023-3311231012221300-1220003101032203)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-2211202023111213-2221103322110030-2331033132202121-2012102021320112-1030223320201110-0000011010011333-0011001031321121-2233020113012311)
- qradar_receiver.use_tls.mtls_disabled

<a id="canonical-1031002111110322-0020113230113332-0323211303022220-0203030213301032-0003032131113013-0300110032210130-2320310122133121-1223213301132030"></a>

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

<a id="canonical-3021121332121230-0330221320320012-3133230012110201-0303030120101330-0230220010223321-3300011031303222-0211300131003300-0131002021233113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.use_tls.mtls_enable` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-2321021321022332-2312230013122032-1100112310202133-0112002300022123-2320312003031221-1122032002223023-3311231012221300-1220003101032203)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-2211202023111213-2221103322110030-2331033132202121-2012102021320112-1030223320201110-0000011010011333-0011001031321121-2233020113012311)
- qradar_receiver.use_tls.mtls_enable

<a id="canonical-1102000323013231-3011332031221221-3223322211322130-3133331323223121-3232303103233103-0331221222131112-1223101302213111-3221301213001220"></a>

Type: `"single"`. Computed.

MTLS Client config allows configuration of mTLS client OPTIONS.

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

<a id="canonical-1002333101133310-1222232103112033-2220032032130020-3020022223101332-3212210133112303-1211212311332033-1031021230232201-3310110012331002"></a>

### Direct properties for `qradar_receiver.use_tls.mtls_enable`

<a id="canonical-1023232113213120-0000322100312001-2123103300323311-1302120103233201-0201210131220111-2232202012112303-2122203032030310-1003310133231322"></a>

#### `qradar_receiver.use_tls.mtls_enable.certificate` property

Type: `"string"`. Computed.

Client certificate is PEM-encoded certificate or certificate-chain.

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
    "formatDescription": "PEM-encoded X.509 certificate, max 5MB",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 100,
    "pattern": "^-----BEGIN CERTIFICATE-----\\n.*\\n-----END CERTIFICATE-----$",
    "validation": {
      "standard": "PEM"
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

- [key_url](data-sources--global_log_receiver--reference--group-004.md#canonical-3100301230130111-3222101131032033-3112311121211223-1312030001232203-0022300010321222-0102213201302321-2112300103211310-2101023102301302): complete subsection reference.

<a id="canonical-3100301230130111-3222101131032033-3112311121211223-1312030001232203-0022300010321222-0102213201302321-2112300103211310-2101023102301302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.use_tls.mtls_enable.key_url` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-2321021321022332-2312230013122032-1100112310202133-0112002300022123-2320312003031221-1122032002223023-3311231012221300-1220003101032203)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-2211202023111213-2221103322110030-2331033132202121-2012102021320112-1030223320201110-0000011010011333-0011001031321121-2233020113012311)
- [qradar_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-004.md#canonical-3021121332121230-0330221320320012-3133230012110201-0303030120101330-0230220010223321-3300011031303222-0211300131003300-0131002021233113)
- qradar_receiver.use_tls.mtls_enable.key_url

<a id="canonical-3013222211230321-0300321101212300-0131201130103201-0110310111322300-1322020301322001-0002300130110132-1012112222220011-1002132110023013"></a>

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

<a id="canonical-1231133113310031-3213302313301300-3313110333020022-2332203100231332-1021301001130003-1113311233113002-3030013010020130-1002311133012020"></a>

### Direct properties for `qradar_receiver.use_tls.mtls_enable.key_url`

- [blindfold_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-3200100200210113-0202200100333101-0210100122110030-0203022302323032-1231112033230013-3023002002031003-0202102301122020-3021323223312111): complete subsection reference.

- [clear_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-1332033023123302-2031003131031113-2323230312311202-0023312132113111-0000230321201033-3213103321131233-1032100022300220-2131301020321220): complete subsection reference.

<a id="canonical-3200100200210113-0202200100333101-0210100122110030-0203022302323032-1231112033230013-3023002002031003-0202102301122020-3021323223312111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-2321021321022332-2312230013122032-1100112310202133-0112002300022123-2320312003031221-1122032002223023-3311231012221300-1220003101032203)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-2211202023111213-2221103322110030-2331033132202121-2012102021320112-1030223320201110-0000011010011333-0011001031321121-2233020113012311)
- [qradar_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-004.md#canonical-3021121332121230-0330221320320012-3133230012110201-0303030120101330-0230220010223321-3300011031303222-0211300131003300-0131002021233113)
- [qradar_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-004.md#canonical-3100301230130111-3222101131032033-3112311121211223-1312030001232203-0022300010321222-0102213201302321-2112300103211310-2101023102301302)
- qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info

<a id="canonical-0033312112311302-2211223000112211-3020201001322220-0313222322010323-2310213203232110-0301332033031013-0211220200102303-0310133302001210"></a>

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

<a id="canonical-2211320312230020-2001013102310031-2212130222221123-1330212013021301-0101112320303310-0232121102223230-2223322032232210-3330120233203112"></a>

### Direct properties for `qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info`

<a id="canonical-0322330321221331-1013132123001032-2101210332311030-0013320031333000-2032313230201312-0030310231233321-1331013030230103-3302233212110333"></a>

#### `qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3221232030221232-3013021202033023-2322132321313123-2200002001232131-0301103100332021-3022330111022110-0100131320022201-0002231122112022"></a>

<a id="canonical-2000010310320312-0320110210021012-1012212210221302-1233003132201111-1110022103312302-1122010032123222-2013320311201122-0031111013230202"></a>

#### `qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1202122332200302-2222131032210313-1003230301231330-0123303132113322-2122113331120122-2311012020003120-2221001130132223-3032000321201012"></a>

<a id="canonical-3212013012103203-2333203001332310-0123101222203221-3302213031221110-0223211123232321-2033100030023022-2013330101102102-2103311120013023"></a>

#### `qradar_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1332033023123302-2031003131031113-2323230312311202-0023312132113111-0000230321201033-3213103321131233-1032100022300220-2131301020321220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-2321021321022332-2312230013122032-1100112310202133-0112002300022123-2320312003031221-1122032002223023-3311231012221300-1220003101032203)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-2211202023111213-2221103322110030-2331033132202121-2012102021320112-1030223320201110-0000011010011333-0011001031321121-2233020113012311)
- [qradar_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-004.md#canonical-3021121332121230-0330221320320012-3133230012110201-0303030120101330-0230220010223321-3300011031303222-0211300131003300-0131002021233113)
- [qradar_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-004.md#canonical-3100301230130111-3222101131032033-3112311121211223-1312030001232203-0022300010321222-0102213201302321-2112300103211310-2101023102301302)
- qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info

<a id="canonical-2201210312331202-0101132022131022-2222010023212102-2023013013203102-3312023322011122-3033102011133110-2031032213123303-3132011310320302"></a>

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

<a id="canonical-1210033330212023-2221030311030133-2033222220123110-1310300233221331-3113331211301211-0212112010312200-2130202001310123-1031313010203002"></a>

### Direct properties for `qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info`

<a id="canonical-0133230231030101-2220032310201011-3301300030132332-3121222122111013-2002203001121223-0100022131033210-1112000031321221-1230202330333311"></a>

#### `qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3133211121033330-0233013321011211-0020022001231222-3330023303233130-0322030301001130-2223303333312030-1110033220130211-2000002233032102"></a>

<a id="canonical-2112211110102312-0022130012220101-0302230003130223-0221301230031303-2030330132221302-0332321022112130-2023020032123122-0012110010310331"></a>

#### `qradar_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1321113003023032-3133333102202200-3330133200020201-2210231001033103-1113301003303003-0302111220002111-0312310133310112-2033030022002322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `qradar_receiver.use_tls.no_ca` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [qradar_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-2321021321022332-2312230013122032-1100112310202133-0112002300022123-2320312003031221-1122032002223023-3311231012221300-1220003101032203)
- [qradar_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-2211202023111213-2221103322110030-2331033132202121-2012102021320112-1030223320201110-0000011010011333-0011001031321121-2233020113012311)
- qradar_receiver.use_tls.no_ca

<a id="canonical-2022312312103013-1230211221100032-2332102120220202-2023002223022230-3202111112303012-1120332111111212-0022023231311011-0001333322320123"></a>

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

<a id="canonical-1320112013330301-1233300003131021-1031032331312212-0332111011101012-0110303200211110-0302300232010303-1133222333311232-3200003320210123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_logs` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- request_logs

<a id="canonical-1032022001202310-0203211112112032-3223133020101320-1122331023331223-0011221222002323-0101233233111212-1002013321121320-1323300223110231"></a>

Type: `"single"`. Computed.

Configuration for request logs with sampling choice. Allows selection between sampled (default) or
unsampled (full) request logs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-sampling_choice": "[\"sampled\",\"unsampled\"]"
}
```

<a id="canonical-1110121300221033-0102231131001221-3123001021310210-3201032000132233-3310203021301202-1103030220232011-3013110311331032-3033332220310033"></a>

### Direct properties for `request_logs`

- [sampled](data-sources--global_log_receiver--reference--group-004.md#canonical-3200100232122302-0211213120103021-3202321000131311-3212012110223110-1203323122230000-3020110233311101-1212230131233030-0233113222101002): complete subsection reference.

- [unsampled](data-sources--global_log_receiver--reference--group-004.md#canonical-1233211210202013-1333120132123330-0331113102232220-0202021333003133-1203103032313102-2012031101003210-2213123200023122-0233131311333203): complete subsection reference.

<a id="canonical-3200100232122302-0211213120103021-3202321000131311-3212012110223110-1203323122230000-3020110233311101-1212230131233030-0233113222101002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_logs.sampled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [request_logs](data-sources--global_log_receiver--reference--group-004.md#canonical-1320112013330301-1233300003131021-1031032331312212-0332111011101012-0110303200211110-0302300232010303-1133222333311232-3200003320210123)
- request_logs.sampled

<a id="canonical-0213000030003023-1302322002320330-1101232212233301-1333011030312222-0232332313312220-3002233121133303-1131323201303233-0201300101121120"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-1233211210202013-1333120132123330-0331113102232220-0202021333003133-1203103032313102-2012031101003210-2213123200023122-0233131311333203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `request_logs.unsampled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [request_logs](data-sources--global_log_receiver--reference--group-004.md#canonical-1320112013330301-1233300003131021-1031032331312212-0332111011101012-0110303200211110-0302300232010303-1133222333311232-3200003320210123)
- request_logs.unsampled

<a id="canonical-2302132121301121-0112133011310222-1221312213031332-3202011003220203-1221000010321133-3012123231311323-1213012233300221-1012211031332330"></a>

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

<a id="canonical-3122321231311212-1111122011230100-1332330302001111-0100123100212313-2310312223130013-0122212111212000-1312212133311330-2113220302103202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `s3_receiver` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- s3_receiver

<a id="canonical-3212111221313232-1220032220002121-1013323021001210-2102032312000332-1100102021232231-2013023201232020-2333203332302002-2022012110300003"></a>

Type: `"single"`. Computed.

S3 Configuration for Global Log Receiver.

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

<a id="canonical-3133231103322102-1213002210131131-0201302000222022-0320332113033310-3123333100301001-1020322133013213-1321032131221031-2003033321230120"></a>

### Direct properties for `s3_receiver`

- [aws_cred](data-sources--global_log_receiver--reference--group-004.md#canonical-0003321233311003-3301113211012122-0103013321010222-3300113211001131-0130113110200303-1032313231101020-3131112000301032-3333012230303003): complete subsection reference.

<a id="canonical-3130112022220123-0322010213101030-2030001220321011-3222113132131011-3021021310310021-2112332203022012-2013312231100122-0030030133210210"></a>

<a id="canonical-2100000002021300-1022133020231232-3023121121110203-2110000123232133-0212030311011103-1022303320123012-1130103321123021-2112311100100102"></a>

#### `s3_receiver.aws_region` property

Type: `"string"`. Computed.

\[Enum:
ap-northeast-1|ap-southeast-1|eu-central-1|eu-west-1|eu-west-3|sa-east-1|us-east-1|us-east-2|us-west-2|ca-central-1|af-south-1|ap-east-1|ap-south-1|ap-northeast-2|ap-southeast-2|eu-south-1|eu-north-1|eu-west-2|me-south-1|us-west-1|ap-southeast-3\]
AWS Region. AWS Region Name. Possible values are \`ap-northeast-1\`, \`ap-southeast-1\`,
\`eu-central-1\`, \`eu-west-1\`, \`eu-west-3\`, \`sa-east-1\`, \`us-east-1\`, \`us-east-2\`,
\`us-west-2\`, \`ca-central-1\`, \`af-south-1\`, \`ap-east-1\`, \`ap-south-1\`, \`ap-northeast-2\`,
\`ap-southeast-2\`, \`eu-south-1\`, \`eu-north-1\`, \`eu-west-2\`, \`me-south-1\`, \`us-west-1\`,
\`ap-southeast-3\`.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ap-northeast-1",
    "ap-southeast-1",
    "eu-central-1",
    "eu-west-1",
    "eu-west-3",
    "sa-east-1",
    "us-east-1",
    "us-east-2",
    "us-west-2",
    "ca-central-1",
    "af-south-1",
    "ap-east-1",
    "ap-south-1",
    "ap-northeast-2",
    "ap-southeast-2",
    "eu-south-1",
    "eu-north-1",
    "eu-west-2",
    "me-south-1",
    "us-west-1",
    "ap-southeast-3"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.in": "[\\\"ap-northeast-1\\\",\\\"ap-southeast-1\\\",\\\"eu-central-1\\\",\\\"eu-west-1\\\",\\\"eu-west-3\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-2\\\",\\\"ca-central-1\\\",\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-south-1\\\",\\\"ap-northeast-2\\\",\\\"ap-southeast-2\\\",\\\"eu-south-1\\\",\\\"eu-north-1\\\",\\\"eu-west-2\\\",\\\"me-south-1\\\",\\\"us-west-1\\\",\\\"ap-southeast-3\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ap-northeast-1\\\",\\\"ap-southeast-1\\\",\\\"eu-central-1\\\",\\\"eu-west-1\\\",\\\"eu-west-3\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-2\\\",\\\"ca-central-1\\\",\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-south-1\\\",\\\"ap-northeast-2\\\",\\\"ap-southeast-2\\\",\\\"eu-south-1\\\",\\\"eu-north-1\\\",\\\"eu-west-2\\\",\\\"me-south-1\\\",\\\"us-west-1\\\",\\\"ap-southeast-3\\\"]"
  }
}
```

- [batch](data-sources--global_log_receiver--reference--group-004.md#canonical-0030311313112002-2000231131211301-2332010033022313-3320203012212213-0313301323012200-0332113133321022-0221021300003020-2011321313332022): complete subsection reference.

<a id="canonical-0010213232000030-1120300322233310-2021103020303311-3220221331323312-0210232131011112-3122323121311010-1032300103222012-2032010201003301"></a>

<a id="canonical-2200123331133032-2233110200330301-3310200020213203-1021223210233132-1222032312321330-0210000111101001-2313220332011103-2303121111221000"></a>

#### `s3_receiver.bucket` property

Type: `"string"`. Computed.

S3 Bucket Name. S3 Bucket Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 3,
    "pattern": "^[a-z0-9]+[a-z0-9\\\\.-]+[a-z0-9]$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9]+[a-z0-9\\\\.-]+[a-z0-9]$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9]+[a-z0-9\\\\.-]+[a-z0-9]$"
  }
}
```

- [compression](data-sources--global_log_receiver--reference--group-004.md#canonical-1211010102101212-1331203202320022-2311302113132203-2212210101102333-0103101100000212-0213212213031303-1011232103321221-2101202313300203): complete subsection reference.

- [filename_options](data-sources--global_log_receiver--reference--group-004.md#canonical-0100132201122032-3320200231020022-1220020303222320-3103120123303222-1231301303201313-1212033130022310-1122103123011330-1203031120313330): complete subsection reference.

<a id="canonical-0003321233311003-3301113211012122-0103013321010222-3300113211001131-0130113110200303-1032313231101020-3131112000301032-3333012230303003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `s3_receiver.aws_cred` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-3122321231311212-1111122011230100-1332330302001111-0100123100212313-2310312223130013-0122212111212000-1312212133311330-2113220302103202)
- s3_receiver.aws_cred

<a id="canonical-3202012023230310-1232321333123111-2110301020021221-1222030102300112-1022302333122211-0002023120303122-3310223101003302-1232311223201111"></a>

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

<a id="canonical-3120022313303002-3003003312003221-1212231211210200-0103202003102130-1212332322130012-3221233130232333-0000310233112302-2022010123210010"></a>

### Direct properties for `s3_receiver.aws_cred`

<a id="canonical-0322202223031112-3032032230302130-2330322222210210-2212330001231202-0213232332031201-0130003313021110-1121222111233103-0212002031023222"></a>

#### `s3_receiver.aws_cred.name` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1300113000103321-2023233201122030-3113320102221202-2113021312022100-2011012330213333-3132130100213212-2010022121233111-2321133221202022"></a>

<a id="canonical-1030310332111233-0331331331301321-2333133022131212-3312120102121021-0223020311020222-1011332002122000-0030302120121013-2101222130213212"></a>

#### `s3_receiver.aws_cred.namespace` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1313220200133133-0303012021322010-2322333212132001-0113210103003000-0102000301112103-3231322330303333-2231133012130130-0103231211302020"></a>

<a id="canonical-2130110301002013-0011333132020320-3321301031012011-0211300313313120-3133002111002233-2102030010122111-2022120033212122-2220322330123212"></a>

#### `s3_receiver.aws_cred.tenant` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0030311313112002-2000231131211301-2332010033022313-3320203012212213-0313301323012200-0332113133321022-0221021300003020-2011321313332022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `s3_receiver.batch` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-3122321231311212-1111122011230100-1332330302001111-0100123100212313-2310312223130013-0122212111212000-1312212133311330-2113220302103202)
- s3_receiver.batch

<a id="canonical-3020230212103033-2133200232032131-3030102231220220-0212110022322010-1113213223103133-1201013033203131-3323110321302031-2023031103332131"></a>

Type: `"single"`. Computed.

Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-batch_bytes": "[\"max_bytes\",\"max_bytes_disabled\"]",
  "x-ves-oneof-field-batch_events": "[\"max_events\",\"max_events_disabled\"]",
  "x-ves-oneof-field-batch_timeout": "[\"timeout_seconds\",\"timeout_seconds_default\"]"
}
```

<a id="canonical-2002300102312030-2032311020112032-3100100033121213-2022010310031211-1231133013233121-3112303200221112-0213121011322220-2300102032333303"></a>

### Direct properties for `s3_receiver.batch`

<a id="canonical-1232003320120121-1030132103012101-3300312103030100-3012333211001311-1020133010333002-2030223013330000-3302300302123300-1231222033030221"></a>

#### `s3_receiver.batch.max_bytes` property

Type: `"number"`. Computed.

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10485760,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 4096
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

- [max_bytes_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-0300002321302223-0023223112101120-0013133312201111-0221001213120312-2003102311100230-2300032303120102-0220101020333010-3102313200012130): complete subsection reference.

<a id="canonical-1221000313212123-3110311020011320-1131332220203311-0011331333033001-1102111323231121-0112300121002313-1121202210323000-1222032103221003"></a>

<a id="canonical-1313030213203333-1120330033031333-0233302123122233-0232003202021122-1121230333331223-2033200123101120-0322130303003320-1312103030330201"></a>

#### `s3_receiver.batch.max_events` property

Type: `"number"`. Computed.

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 32
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  }
}
```

- [max_events_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-2311021311132023-1303002132300010-1100031113232002-2030101123030330-1232000102221200-1201130232132300-2001310112203113-0302210121200131): complete subsection reference.

<a id="canonical-1200121311230210-3122321333020103-0120222203010002-3232310321210101-0120232232013331-1121233123022031-1013302232323233-3231203010011013"></a>

<a id="canonical-3311000000132133-2130312330103033-0213123023123233-1130302323030122-0233003200332323-3333200131222030-2023102212002223-3113301302312123"></a>

#### `s3_receiver.batch.timeout_seconds` property

Type: `"string"`. Computed.

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "uint64",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  }
}
```

- [timeout_seconds_default](data-sources--global_log_receiver--reference--group-004.md#canonical-3102212213233231-2103011002112312-1332332211233203-3303222102232333-0311012131110002-0301232231020321-0032033130331320-1010223023231111): complete subsection reference.

<a id="canonical-0300002321302223-0023223112101120-0013133312201111-0221001213120312-2003102311100230-2300032303120102-0220101020333010-3102313200012130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `s3_receiver.batch.max_bytes_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-3122321231311212-1111122011230100-1332330302001111-0100123100212313-2310312223130013-0122212111212000-1312212133311330-2113220302103202)
- [s3_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-0030311313112002-2000231131211301-2332010033022313-3320203012212213-0313301323012200-0332113133321022-0221021300003020-2011321313332022)
- s3_receiver.batch.max_bytes_disabled

<a id="canonical-1031312331231133-2030212110221311-2021111101203223-2110103313222202-0013031132001221-1011313230222130-1223031002230130-2312311031102020"></a>

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

<a id="canonical-2311021311132023-1303002132300010-1100031113232002-2030101123030330-1232000102221200-1201130232132300-2001310112203113-0302210121200131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `s3_receiver.batch.max_events_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-3122321231311212-1111122011230100-1332330302001111-0100123100212313-2310312223130013-0122212111212000-1312212133311330-2113220302103202)
- [s3_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-0030311313112002-2000231131211301-2332010033022313-3320203012212213-0313301323012200-0332113133321022-0221021300003020-2011321313332022)
- s3_receiver.batch.max_events_disabled

<a id="canonical-3123332003223121-2213310010110002-1023023102211100-2232123322210112-2133113023311001-2021021032101303-0001130231320313-1213021231002320"></a>

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

<a id="canonical-3102212213233231-2103011002112312-1332332211233203-3303222102232333-0311012131110002-0301232231020321-0032033130331320-1010223023231111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `s3_receiver.batch.timeout_seconds_default` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-3122321231311212-1111122011230100-1332330302001111-0100123100212313-2310312223130013-0122212111212000-1312212133311330-2113220302103202)
- [s3_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-0030311313112002-2000231131211301-2332010033022313-3320203012212213-0313301323012200-0332113133321022-0221021300003020-2011321313332022)
- s3_receiver.batch.timeout_seconds_default

<a id="canonical-1223311312003032-2320022133231021-1221010003222300-1223300001213203-3311231021211012-0232320301211112-0310102303102010-3220030131330313"></a>

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

<a id="canonical-1211010102101212-1331203202320022-2311302113132203-2212210101102333-0103101100000212-0213212213031303-1011232103321221-2101202313300203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `s3_receiver.compression` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-3122321231311212-1111122011230100-1332330302001111-0100123100212313-2310312223130013-0122212111212000-1312212133311330-2113220302103202)
- s3_receiver.compression

<a id="canonical-3001333213310201-2102302233230011-3032230023101000-3232110002022301-1102001231322133-2030103332321002-3002320222013201-0312013121300213"></a>

Type: `"single"`. Computed.

Configuration parameter for compression.

Additional upstream details:

Compression Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-compression_choice": "[\"compression_default\",\"compression_gzip\",\"compression_none\"]"
}
```

<a id="canonical-2212103303100211-1032200322232111-1103011211033122-0022331321122332-0000221012031323-3233213013320211-3131000331300020-1301031233221132"></a>

### Direct properties for `s3_receiver.compression`

- [compression_default](data-sources--global_log_receiver--reference--group-004.md#canonical-2002302113301221-1031131313000313-2101220213231330-3113100123132011-3301030111033313-2221213230213223-3302132123230021-0321301201331112): complete subsection reference.

- [compression_gzip](data-sources--global_log_receiver--reference--group-004.md#canonical-3212132120322200-2302331133101021-2111211212223010-3022331020113303-2320201110202002-1320122310202113-0303332031320331-3200222123310302): complete subsection reference.

- [compression_none](data-sources--global_log_receiver--reference--group-004.md#canonical-3000302301201303-3102312231000030-2230122033111202-1100210120321131-3311002022330203-3223331330220131-2211300321013303-1221310330001320): complete subsection reference.

<a id="canonical-2002302113301221-1031131313000313-2101220213231330-3113100123132011-3301030111033313-2221213230213223-3302132123230021-0321301201331112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `s3_receiver.compression.compression_default` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-3122321231311212-1111122011230100-1332330302001111-0100123100212313-2310312223130013-0122212111212000-1312212133311330-2113220302103202)
- [s3_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-1211010102101212-1331203202320022-2311302113132203-2212210101102333-0103101100000212-0213212213031303-1011232103321221-2101202313300203)
- s3_receiver.compression.compression_default

<a id="canonical-2101011132302010-1030032330223332-2331120003000330-0021322010212233-2103201031201120-1030331312213122-1321222100100313-2202013222120010"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for compression default.

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

<a id="canonical-3212132120322200-2302331133101021-2111211212223010-3022331020113303-2320201110202002-1320122310202113-0303332031320331-3200222123310302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `s3_receiver.compression.compression_gzip` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-3122321231311212-1111122011230100-1332330302001111-0100123100212313-2310312223130013-0122212111212000-1312212133311330-2113220302103202)
- [s3_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-1211010102101212-1331203202320022-2311302113132203-2212210101102333-0103101100000212-0213212213031303-1011232103321221-2101202313300203)
- s3_receiver.compression.compression_gzip

<a id="canonical-1220000312113133-2210122210013121-3002312323100031-2332022331021101-1020122102103021-0322231130303011-2012323132101103-0020110213121231"></a>

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

<a id="canonical-3000302301201303-3102312231000030-2230122033111202-1100210120321131-3311002022330203-3223331330220131-2211300321013303-1221310330001320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `s3_receiver.compression.compression_none` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-3122321231311212-1111122011230100-1332330302001111-0100123100212313-2310312223130013-0122212111212000-1312212133311330-2113220302103202)
- [s3_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-1211010102101212-1331203202320022-2311302113132203-2212210101102333-0103101100000212-0213212213031303-1011232103321221-2101202313300203)
- s3_receiver.compression.compression_none

<a id="canonical-0222021302001020-1002323122223230-2121022131132310-2000023221200211-1311330112010322-3011000012030021-3011202100113013-0200222130110312"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for compression none.

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

<a id="canonical-0100132201122032-3320200231020022-1220020303222320-3103120123303222-1231301303201313-1212033130022310-1122103123011330-1203031120313330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `s3_receiver.filename_options` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-3122321231311212-1111122011230100-1332330302001111-0100123100212313-2310312223130013-0122212111212000-1312212133311330-2113220302103202)
- s3_receiver.filename_options

<a id="canonical-1032230031213311-2300200122322002-0302230112310111-2120100032200020-2301110213210003-2021332303201111-1322333322030322-3113032031023012"></a>

Type: `"single"`. Computed.

Filename OPTIONS allow customization of filename and folder paths used by a destination endpoint
bucket or file.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-folder": "[\"custom_folder\",\"log_type_folder\",\"no_folder\"]"
}
```

<a id="canonical-1003000111200200-3210103200303003-3300223012311113-0331221222213221-0311013133301213-1223300023033331-0212220333200110-2311103130303300"></a>

### Direct properties for `s3_receiver.filename_options`

<a id="canonical-3230322132020033-0122030010132232-1220200103000203-3133103023310103-0333033321222301-3212011323001322-3230000102021202-0103220122222322"></a>

#### `s3_receiver.filename_options.custom_folder` property

Type: `"string"`. Computed.

Exclusive with \[log\_type\_folder no\_folder\] Use your own folder name as the name of the folder
in the endpoint bucket or file The folder name must match \`/^\[a-z\_\]\[a-z0-9\\\\-\\\\.\_\]\*$/i\`

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "pattern": "^[A-Za-z_][A-Za-z0-9\\\\-\\\\._]*$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^[A-Za-z_][A-Za-z0-9\\\\-\\\\._]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^[A-Za-z_][A-Za-z0-9\\\\-\\\\._]*$"
  }
}
```

- [log_type_folder](data-sources--global_log_receiver--reference--group-004.md#canonical-0123203103002230-1021003313212300-1020132113232121-0123320120202112-3230033000000220-2331231122332210-1233332333022011-2023130223020331): complete subsection reference.

- [no_folder](data-sources--global_log_receiver--reference--group-004.md#canonical-1310231121310111-1212121112320110-1322001223001031-2110333310003030-0213212011333213-2311022001120122-1010123230030300-1132113233200221): complete subsection reference.

<a id="canonical-0123203103002230-1021003313212300-1020132113232121-0123320120202112-3230033000000220-2331231122332210-1233332333022011-2023130223020331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `s3_receiver.filename_options.log_type_folder` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-3122321231311212-1111122011230100-1332330302001111-0100123100212313-2310312223130013-0122212111212000-1312212133311330-2113220302103202)
- [s3_receiver.filename_options](data-sources--global_log_receiver--reference--group-004.md#canonical-0100132201122032-3320200231020022-1220020303222320-3103120123303222-1231301303201313-1212033130022310-1122103123011330-1203031120313330)
- s3_receiver.filename_options.log_type_folder

<a id="canonical-2322233312020231-2000120120233333-3113002301321110-0000132302223232-3132211033310322-3232132330332321-2202312231021022-3322112001213221"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for log type folder.

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

<a id="canonical-1310231121310111-1212121112320110-1322001223001031-2110333310003030-0213212011333213-2311022001120122-1010123230030300-1132113233200221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `s3_receiver.filename_options.no_folder` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [s3_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-3122321231311212-1111122011230100-1332330302001111-0100123100212313-2310312223130013-0122212111212000-1312212133311330-2113220302103202)
- [s3_receiver.filename_options](data-sources--global_log_receiver--reference--group-004.md#canonical-0100132201122032-3320200231020022-1220020303222320-3103120123303222-1231301303201313-1212033130022310-1122103123011330-1203031120313330)
- s3_receiver.filename_options.no_folder

<a id="canonical-3211122120332313-3121320303230030-0222131322213202-2002322003202101-3222031120323101-1030111030312000-2110032103222302-1132220333031010"></a>

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

<a id="canonical-1002031033031310-1233001100003213-2003102123102332-1202233033221123-1021111002311202-0121010220310110-1211132100130223-2001223323310023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `security_events` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- security_events

<a id="canonical-2121201001133031-1031120231001131-1132311112122102-3321300232323210-0223013202122312-3113023102021320-3231333333203113-1212202033021010"></a>

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

<a id="canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- splunk_receiver

<a id="canonical-2130322203222220-0110323213010302-1233212331033202-0013332333331213-2311200220221030-0313332130213320-3232333202032213-1331302023310301"></a>

Type: `"single"`. Computed.

Configuration for Splunk HEC Logs endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

<a id="canonical-1211032111113110-3223123101201030-1212323121312111-3133022103220213-3000001110130110-2222303123001123-3011202310020303-3231231203211312"></a>

### Direct properties for `splunk_receiver`

- [batch](data-sources--global_log_receiver--reference--group-004.md#canonical-1133003202003200-2330033311332012-1300112221110122-0121103123212002-0232200101212202-0002103000223132-1300212111101210-2320013320233332): complete subsection reference.

- [compression](data-sources--global_log_receiver--reference--group-004.md#canonical-3303120223230322-1112030213201012-1320223022301031-3130210101020332-0110022100312332-1230001103023212-1021220230001013-1001003233133012): complete subsection reference.

<a id="canonical-0122310212212110-3112223030311123-2232032230233300-3030032021332100-0032011231230113-2103000003021332-0101003333210312-0013311330230221"></a>

<a id="canonical-1331101103303233-0210012003033010-0322103102013101-0130102002301330-0311332312031330-2033232310300210-2010100312332332-3013033131133233"></a>

#### `splunk_receiver.endpoint` property

Type: `"string"`. Computed.

Splunk HEC Logs Endpoint. Splunk HEC Logs Endpoint, (Note: must not contain \`/services/collector\`)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1,
    "pattern": "^https?://[^\\s/$.?#].[^\\s]*$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
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

- [no_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-1233030330011113-1231121311123123-0322000010211013-3222020113322211-3022231031332102-3133330303211311-2020013231023221-3011331223022132): complete subsection reference.

- [splunk_hec_token](data-sources--global_log_receiver--reference--group-004.md#canonical-1003122200023000-0033320321101300-3222221101223110-0202121312031221-2102112313212003-2000213001302222-3312032201013200-0220313312013022): complete subsection reference.

- [use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-3030211232321010-0002332312301211-2032203223320101-1230130023032103-1312001002030231-3021212332003121-3330131103001221-2330113233101023): complete subsection reference.

<a id="canonical-1133003202003200-2330033311332012-1300112221110122-0121103123212002-0232200101212202-0002103000223132-1300212111101210-2320013320233332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.batch` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- splunk_receiver.batch

<a id="canonical-0310220231331303-0331002120123031-0301111231212212-1023313103123122-0000330132030210-3020131210312123-1200013032123101-0001321123123300"></a>

Type: `"single"`. Computed.

Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-batch_bytes": "[\"max_bytes\",\"max_bytes_disabled\"]",
  "x-ves-oneof-field-batch_events": "[\"max_events\",\"max_events_disabled\"]",
  "x-ves-oneof-field-batch_timeout": "[\"timeout_seconds\",\"timeout_seconds_default\"]"
}
```

<a id="canonical-0122311110321300-3231030002131320-0123203202321100-2303103102020122-2310201230303200-3322330133313313-3231332323230213-3330233230232202"></a>

### Direct properties for `splunk_receiver.batch`

<a id="canonical-1313123330123213-0003011032101233-3100023111111030-1132130103203023-0023311121132023-0211032310032013-2223222211222223-0103022301200332"></a>

#### `splunk_receiver.batch.max_bytes` property

Type: `"number"`. Computed.

Exclusive with \[max\_bytes\_disabled\] Send batch to endpoint after the batch is equal to or larger
than this many bytes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10485760,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 4096
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "4096",
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

- [max_bytes_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-3111222321000321-0120131323312113-1033322130301101-0000302330020000-1113012130231021-1311210222120231-1120130030002310-1200111131321321): complete subsection reference.

<a id="canonical-0332001220230311-1010200201130011-2030132303020020-3211320112201213-1023222330200313-2110312100010333-2322322023310323-2132301201311211"></a>

<a id="canonical-1011103320300302-1220003033212003-0330030301013220-1330332010202203-3022120200103233-1330123223111321-2113023310112113-1010333201213332"></a>

#### `splunk_receiver.batch.max_events` property

Type: `"number"`. Computed.

Exclusive with \[max\_events\_disabled\] Send batch to endpoint after this many log messages are in
the batch.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 32
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "32",
    "ves.io.schema.rules.uint32.lte": "2000"
  }
}
```

- [max_events_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-1033000022022201-1103122311103101-1321000023212222-0311310033112000-2321233131230033-3030001232122222-1021112131301103-3310231220311003): complete subsection reference.

<a id="canonical-0210302133331323-2002310321031133-2310301132233113-2223020302300102-2000213122103310-1100222200330101-3210300131221100-3123021020202220"></a>

<a id="canonical-0000303110021201-0102030121033201-1323223121100023-2301122013231130-1103121210130332-2323122002101220-2220010133112023-0002032220032101"></a>

#### `splunk_receiver.batch.timeout_seconds` property

Type: `"string"`. Computed.

Exclusive with \[timeout\_seconds\_default\] Send batch to the endpoint after this many seconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "uint64",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  }
}
```

- [timeout_seconds_default](data-sources--global_log_receiver--reference--group-004.md#canonical-1323132023122003-2323333011320033-1001313000033020-2232133013103301-2010201300322203-0012110111123301-3222312301131202-3213321233310032): complete subsection reference.

<a id="canonical-3111222321000321-0120131323312113-1033322130301101-0000302330020000-1113012130231021-1311210222120231-1120130030002310-1200111131321321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.batch.max_bytes_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-1133003202003200-2330033311332012-1300112221110122-0121103123212002-0232200101212202-0002103000223132-1300212111101210-2320013320233332)
- splunk_receiver.batch.max_bytes_disabled

<a id="canonical-1232023303133101-3020111132101123-0102313223233313-2300012130123310-2122102302312132-2312110110030130-0023122000311322-2103232313120211"></a>

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

<a id="canonical-1033000022022201-1103122311103101-1321000023212222-0311310033112000-2321233131230033-3030001232122222-1021112131301103-3310231220311003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.batch.max_events_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-1133003202003200-2330033311332012-1300112221110122-0121103123212002-0232200101212202-0002103000223132-1300212111101210-2320013320233332)
- splunk_receiver.batch.max_events_disabled

<a id="canonical-1311200213112200-3231012032111330-2021101211203231-1033333300110112-0230312323032022-1223132213220201-3133212203021103-0200330313121323"></a>

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

<a id="canonical-1323132023122003-2323333011320033-1001313000033020-2232133013103301-2010201300322203-0012110111123301-3222312301131202-3213321233310032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.batch.timeout_seconds_default` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.batch](data-sources--global_log_receiver--reference--group-004.md#canonical-1133003202003200-2330033311332012-1300112221110122-0121103123212002-0232200101212202-0002103000223132-1300212111101210-2320013320233332)
- splunk_receiver.batch.timeout_seconds_default

<a id="canonical-0102310221311100-3321300201103310-1120201321320102-2200031123021023-2313311331312130-2013200120021233-0020310121011101-3213322313303103"></a>

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

<a id="canonical-3303120223230322-1112030213201012-1320223022301031-3130210101020332-0110022100312332-1230001103023212-1021220230001013-1001003233133012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.compression` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- splunk_receiver.compression

<a id="canonical-1132323220133001-0201123332100322-1033310330132112-3213213013322122-0300012121231123-3020321013321232-3323213333203302-1030032221133032"></a>

Type: `"single"`. Computed.

Configuration parameter for compression.

Additional upstream details:

Compression Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-compression_choice": "[\"compression_default\",\"compression_gzip\",\"compression_none\"]"
}
```

<a id="canonical-3131001121202000-3131000000323211-2031100010002231-1313033332012111-3321301011321313-1212233032320102-0203311132001002-3320211103300330"></a>

### Direct properties for `splunk_receiver.compression`

- [compression_default](data-sources--global_log_receiver--reference--group-004.md#canonical-0220032330333000-1120010102301122-3020221220131131-3012330331312122-0213033222201331-3102203201332302-3101102300003212-0130332101101030): complete subsection reference.

- [compression_gzip](data-sources--global_log_receiver--reference--group-004.md#canonical-2212130212220231-3321010011123321-3102001102331131-0000211121213231-2131203130322331-2022003133031302-1012132212333122-3311202213132110): complete subsection reference.

- [compression_none](data-sources--global_log_receiver--reference--group-004.md#canonical-1132213011323301-0203201233202312-0103201103333001-1032323212333223-2331133103322031-2132321132233012-0223222222321301-0101121130022001): complete subsection reference.

<a id="canonical-0220032330333000-1120010102301122-3020221220131131-3012330331312122-0213033222201331-3102203201332302-3101102300003212-0130332101101030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.compression.compression_default` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-3303120223230322-1112030213201012-1320223022301031-3130210101020332-0110022100312332-1230001103023212-1021220230001013-1001003233133012)
- splunk_receiver.compression.compression_default

<a id="canonical-3311331031122302-0212312022022233-0312032020232210-2132203310211031-0103311001021322-1002222320000023-3211103312123213-0221312003323302"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for compression default.

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

<a id="canonical-2212130212220231-3321010011123321-3102001102331131-0000211121213231-2131203130322331-2022003133031302-1012132212333122-3311202213132110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.compression.compression_gzip` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-3303120223230322-1112030213201012-1320223022301031-3130210101020332-0110022100312332-1230001103023212-1021220230001013-1001003233133012)
- splunk_receiver.compression.compression_gzip

<a id="canonical-2132130303021302-2101200321303130-0221321111211022-2233003020031223-3323311320110311-1021220110220012-1303032031122203-3211100221123100"></a>

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

<a id="canonical-1132213011323301-0203201233202312-0103201103333001-1032323212333223-2331133103322031-2132321132233012-0223222222321301-0101121130022001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.compression.compression_none` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.compression](data-sources--global_log_receiver--reference--group-004.md#canonical-3303120223230322-1112030213201012-1320223022301031-3130210101020332-0110022100312332-1230001103023212-1021220230001013-1001003233133012)
- splunk_receiver.compression.compression_none

<a id="canonical-2221100001033010-3102322311003213-2110202223011002-3112322101332001-0301201130313111-1322300310132002-2120010331112000-2012303132000132"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for compression none.

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

<a id="canonical-1233030330011113-1231121311123123-0322000010211013-3222020113322211-3022231031332102-3133330303211311-2020013231023221-3011331223022132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.no_tls` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- splunk_receiver.no_tls

<a id="canonical-1323320233332023-2333133131013212-0300231313001332-3002202330132300-0231200131032222-3223322211102210-2011000310220322-1312000232220302"></a>

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

<a id="canonical-1003122200023000-0033320321101300-3222221101223110-0202121312031221-2102112313212003-2000213001302222-3312032201013200-0220313312013022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.splunk_hec_token` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- splunk_receiver.splunk_hec_token

<a id="canonical-0010103202313023-3011132320331202-0102212312100212-3001200030210211-2311320200030020-3322033230203133-3023103200120212-2003331010321021"></a>

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

<a id="canonical-1300020222003203-0202021220032203-1220023102323221-3121231100033011-1103210131332120-0020012333122123-1132022011123233-2311332200310021"></a>

### Direct properties for `splunk_receiver.splunk_hec_token`

- [blindfold_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-1121031112302011-3332331311022312-3100203321103031-3310001103223212-3131020002020022-1312123023032000-1221330110000113-3222000201010332): complete subsection reference.

- [clear_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-1021211032223203-1332333030312012-0210130231331232-3022003021333011-0000133333131223-2333213130002232-0311122212130021-0302011123212222): complete subsection reference.

<a id="canonical-1121031112302011-3332331311022312-3100203321103031-3310001103223212-3131020002020022-1312123023032000-1221330110000113-3222000201010332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.splunk_hec_token.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.splunk_hec_token](data-sources--global_log_receiver--reference--group-004.md#canonical-1003122200023000-0033320321101300-3222221101223110-0202121312031221-2102112313212003-2000213001302222-3312032201013200-0220313312013022)
- splunk_receiver.splunk_hec_token.blindfold_secret_info

<a id="canonical-3132210202011312-0201101331120113-1222232030200210-1230121333300323-1230101102232033-0020302320112213-2303000333023102-3032303211133120"></a>

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

<a id="canonical-2012001130102302-2212112312220002-1322301322103210-1101233010312102-3321322003030202-2232300330200030-0012300303013233-0321032233123120"></a>

### Direct properties for `splunk_receiver.splunk_hec_token.blindfold_secret_info`

<a id="canonical-2303101132032203-0211112123202323-3211110030200203-0131223211012130-1110332210123011-2332310301332103-0210313012320213-1111132032323103"></a>

#### `splunk_receiver.splunk_hec_token.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0103031121311131-3323313022123223-3203111111300213-2300231302300202-0232120303332032-1101033203231301-1103202232102110-3013020000113203"></a>

<a id="canonical-1221212102133203-3311332310233121-3012023000003321-3032133322213120-1203223011311332-1312113011200232-2023302001211012-2000013113301130"></a>

#### `splunk_receiver.splunk_hec_token.blindfold_secret_info.location` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1311131020313032-1133313011023101-3023031222233011-0100300020213101-1311132232113303-2121100212213012-1122100101202033-3133123220013330"></a>

<a id="canonical-2333223332030002-2332021010201230-1323021202100312-0300110130220103-3000110333321310-3303232030211011-0120200013202102-1122311021220001"></a>

#### `splunk_receiver.splunk_hec_token.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1021211032223203-1332333030312012-0210130231331232-3022003021333011-0000133333131223-2333213130002232-0311122212130021-0302011123212222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.splunk_hec_token.clear_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.splunk_hec_token](data-sources--global_log_receiver--reference--group-004.md#canonical-1003122200023000-0033320321101300-3222221101223110-0202121312031221-2102112313212003-2000213001302222-3312032201013200-0220313312013022)
- splunk_receiver.splunk_hec_token.clear_secret_info

<a id="canonical-0200023233103100-3011023230000130-0100301222313331-3121101032301320-0231302002301322-3101002221230110-1032013322310103-2122320232021113"></a>

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

<a id="canonical-2213122333302231-2020303233231130-0213033121001222-0222201332231332-0021000212010330-2001322031202333-2100133112132103-0021020220023002"></a>

### Direct properties for `splunk_receiver.splunk_hec_token.clear_secret_info`

<a id="canonical-1333313132331203-1003130333221001-3222233022303012-3223123031320111-2321100003122123-1122303121202133-0033232202131203-2101121321213203"></a>

#### `splunk_receiver.splunk_hec_token.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0231130312323030-1130312230200320-2320111130021202-2131123210301103-1320111220031212-0201110010111301-2122203012333333-3321312303021333"></a>

<a id="canonical-3003121120202202-2022302230031302-0033220032303203-2202031010323322-0032222011103211-1201202032013323-2100023030300210-0331203333013103"></a>

#### `splunk_receiver.splunk_hec_token.clear_secret_info.url` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3030211232321010-0002332312301211-2032203223320101-1230130023032103-1312001002030231-3021212332003121-3330131103001221-2330113233101023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.use_tls` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- splunk_receiver.use_tls

<a id="canonical-1031212113011132-2011333230211332-2123311222333320-3331202123213021-1113210222203122-3231203212213130-1301320311201323-1313312332103102"></a>

Type: `"single"`. Computed.

TLS Parameters for client connection to the endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ca_choice": "[\"no_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-mtls_choice": "[\"mtls_disabled\",\"mtls_enable\"]",
  "x-ves-oneof-field-verify_certificate": "[\"disable_verify_certificate\",\"enable_verify_certificate\"]",
  "x-ves-oneof-field-verify_hostname": "[\"disable_verify_hostname\",\"enable_verify_hostname\"]"
}
```

<a id="canonical-2013233201323222-3233002200220302-3101002110021103-0021330131001102-0012200321330000-2231002100210210-2323011033023200-0031113033131311"></a>

### Direct properties for `splunk_receiver.use_tls`

- [disable_verify_certificate](data-sources--global_log_receiver--reference--group-004.md#canonical-0113132010013332-1122203011101010-0110033323000110-0101112021112202-2001131212021302-1002001101133213-1332223012130120-3301023201332102): complete subsection reference.

- [disable_verify_hostname](data-sources--global_log_receiver--reference--group-004.md#canonical-2120013112032130-2201222001203321-0130203200120210-1312300211012122-3313331301223223-3233100331213300-3320320312310122-1032101010001213): complete subsection reference.

- [enable_verify_certificate](data-sources--global_log_receiver--reference--group-004.md#canonical-1232330203130112-1311330323130000-1300003021122033-1222320003112203-0333031110331212-0200201112212310-3202313032312331-1112020123301121): complete subsection reference.

- [enable_verify_hostname](data-sources--global_log_receiver--reference--group-004.md#canonical-3333222331323333-2030013030332323-0331321210300120-2212123201100313-0112313231221131-3023202031103100-1331233133033201-3300102121112321): complete subsection reference.

- [mtls_disabled](data-sources--global_log_receiver--reference--group-004.md#canonical-1131330031301212-3222002001020010-1223131122021130-3220131332023103-2120101201231023-2113023101202200-1113032233303001-1233100102003200): complete subsection reference.

- [mtls_enable](data-sources--global_log_receiver--reference--group-004.md#canonical-0121112313322202-2310211222220013-2323010331330003-2212110021230101-1233303212112301-3032023223223001-1320331112213332-0001332102120320): complete subsection reference.

- [no_ca](data-sources--global_log_receiver--reference--group-004.md#canonical-3222321113000113-2231233223133012-2032013200210220-0201021000323102-1331023333031133-0122211311330020-0200010321023212-3111003231323233): complete subsection reference.

<a id="canonical-1033200111011313-3230202023223222-2101030323123302-1003323303012031-2123100222300223-0111100333233002-2231311332002310-0220231120112202"></a>

<a id="canonical-2200211330232313-0333000100020220-0023231233200002-0212303030132223-0000333201011200-0021221132130100-3213221332210101-1302132221202122"></a>

#### `splunk_receiver.use_tls.trusted_ca_url` property

Type: `"string"`. Computed.

Exclusive with \[no\_ca\] The URL or value for trusted Server CA certificate or certificate chain
Certificates in PEM format including the PEM headers.

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
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-0113132010013332-1122203011101010-0110033323000110-0101112021112202-2001131212021302-1002001101133213-1332223012130120-3301023201332102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.use_tls.disable_verify_certificate` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-3030211232321010-0002332312301211-2032203223320101-1230130023032103-1312001002030231-3021212332003121-3330131103001221-2330113233101023)
- splunk_receiver.use_tls.disable_verify_certificate

<a id="canonical-2223231332002313-0032022223310312-1231223133230100-2123121200320213-3033232222021103-2322032320322231-1123100133131031-3003321022111031"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable verify certificate.

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

<a id="canonical-2120013112032130-2201222001203321-0130203200120210-1312300211012122-3313331301223223-3233100331213300-3320320312310122-1032101010001213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.use_tls.disable_verify_hostname` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-3030211232321010-0002332312301211-2032203223320101-1230130023032103-1312001002030231-3021212332003121-3330131103001221-2330113233101023)
- splunk_receiver.use_tls.disable_verify_hostname

<a id="canonical-2220022321101202-3023033101013113-3113332220001112-2111312011023122-2301330200013003-3113232311121320-3130132320023333-2002100320323032"></a>

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

<a id="canonical-1232330203130112-1311330323130000-1300003021122033-1222320003112203-0333031110331212-0200201112212310-3202313032312331-1112020123301121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.use_tls.enable_verify_certificate` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-3030211232321010-0002332312301211-2032203223320101-1230130023032103-1312001002030231-3021212332003121-3330131103001221-2330113233101023)
- splunk_receiver.use_tls.enable_verify_certificate

<a id="canonical-0230322230221202-2020112013000230-3123100122122031-1133022021313122-2031120331313201-1320223023301333-1021003222302010-2012313213103321"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable verify certificate.

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

<a id="canonical-3333222331323333-2030013030332323-0331321210300120-2212123201100313-0112313231221131-3023202031103100-1331233133033201-3300102121112321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.use_tls.enable_verify_hostname` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-3030211232321010-0002332312301211-2032203223320101-1230130023032103-1312001002030231-3021212332003121-3330131103001221-2330113233101023)
- splunk_receiver.use_tls.enable_verify_hostname

<a id="canonical-2310123001312230-0021301000230220-1033231002320021-2222230121103303-0033022312001232-2131331000033020-3300210221033122-2323000113022022"></a>

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

<a id="canonical-1131330031301212-3222002001020010-1223131122021130-3220131332023103-2120101201231023-2113023101202200-1113032233303001-1233100102003200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.use_tls.mtls_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-3030211232321010-0002332312301211-2032203223320101-1230130023032103-1312001002030231-3021212332003121-3330131103001221-2330113233101023)
- splunk_receiver.use_tls.mtls_disabled

<a id="canonical-3012003232022201-3102013121131022-0232111011323021-2000213212120120-3133132033333333-0323100303331113-1110320011033330-1103000320110013"></a>

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

<a id="canonical-0121112313322202-2310211222220013-2323010331330003-2212110021230101-1233303212112301-3032023223223001-1320331112213332-0001332102120320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.use_tls.mtls_enable` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-3030211232321010-0002332312301211-2032203223320101-1230130023032103-1312001002030231-3021212332003121-3330131103001221-2330113233101023)
- splunk_receiver.use_tls.mtls_enable

<a id="canonical-2010302101030012-1210121202101231-1111311133202220-2021132103100332-3012001222103222-1023303233013021-2101111032233310-0101130022223233"></a>

Type: `"single"`. Computed.

MTLS Client config allows configuration of mTLS client OPTIONS.

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

<a id="canonical-2133232103203220-3322133221320303-1302230111133201-0020231220030223-2130232130121200-3103210223301020-2222322301001213-3003321320302011"></a>

### Direct properties for `splunk_receiver.use_tls.mtls_enable`

<a id="canonical-0120010122112010-0002132022130120-2032300031113301-3100221130213203-2120213321103212-0323302233311113-3330022223033103-2032233232232220"></a>

#### `splunk_receiver.use_tls.mtls_enable.certificate` property

Type: `"string"`. Computed.

Client certificate is PEM-encoded certificate or certificate-chain.

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
    "formatDescription": "PEM-encoded X.509 certificate, max 5MB",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 100,
    "pattern": "^-----BEGIN CERTIFICATE-----\\n.*\\n-----END CERTIFICATE-----$",
    "validation": {
      "standard": "PEM"
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

- [key_url](data-sources--global_log_receiver--reference--group-004.md#canonical-3321002133010021-1100120300031012-1033302121131210-1220023323122012-0221111031311111-3003201330220311-0203301220112323-2320313232003200): complete subsection reference.

<a id="canonical-3321002133010021-1100120300031012-1033302121131210-1220023323122012-0221111031311111-3003201330220311-0203301220112323-2320313232003200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.use_tls.mtls_enable.key_url` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-3030211232321010-0002332312301211-2032203223320101-1230130023032103-1312001002030231-3021212332003121-3330131103001221-2330113233101023)
- [splunk_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-004.md#canonical-0121112313322202-2310211222220013-2323010331330003-2212110021230101-1233303212112301-3032023223223001-1320331112213332-0001332102120320)
- splunk_receiver.use_tls.mtls_enable.key_url

<a id="canonical-3113201231201020-0013002203113121-0333123202230231-3113112321232032-1132112323111003-1211202112312232-3121010132010333-3012103323301131"></a>

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

<a id="canonical-2120333113020023-0312023302102011-2300103220202222-0003122222332311-0122101311222130-2232032311032132-2231323322211033-2110022321331123"></a>

### Direct properties for `splunk_receiver.use_tls.mtls_enable.key_url`

- [blindfold_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-3202132321221132-0221233313210303-2120321211022013-3002031233031122-2113033031231200-2133201021312221-0102301201110301-2200130033132223): complete subsection reference.

- [clear_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-1031033333021002-2013213002003321-1010311121211023-3132121203200222-1133012003000201-1103133202223011-3211002103301121-1221103121103021): complete subsection reference.

<a id="canonical-3202132321221132-0221233313210303-2120321211022013-3002031233031122-2113033031231200-2133201021312221-0102301201110301-2200130033132223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-3030211232321010-0002332312301211-2032203223320101-1230130023032103-1312001002030231-3021212332003121-3330131103001221-2330113233101023)
- [splunk_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-004.md#canonical-0121112313322202-2310211222220013-2323010331330003-2212110021230101-1233303212112301-3032023223223001-1320331112213332-0001332102120320)
- [splunk_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-004.md#canonical-3321002133010021-1100120300031012-1033302121131210-1220023323122012-0221111031311111-3003201330220311-0203301220112323-2320313232003200)
- splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info

<a id="canonical-2120321122323330-0221123321123322-0012200100202233-0213210301302221-0120131120112130-0133201332300111-2003321133012323-3233130130300122"></a>

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

<a id="canonical-2023320133321331-3332211302213123-0100101103001210-1222332023210102-3101311201332212-3302303030001221-3002112113020122-2112211203210301"></a>

### Direct properties for `splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info`

<a id="canonical-3102322113210211-3303301020102000-0201200131101202-1120120113220313-1032223012312300-3130033323031032-2321220201232212-0110200311032301"></a>

#### `splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0322110130022031-1030130230320232-2101300030000133-1323310321002003-2233010231323330-3003321030223122-1230303112011132-0003030330130130"></a>

<a id="canonical-1102103210113321-2300302010133113-1310300322033012-2110000102232300-2223123313130020-2122012302323313-1132212100003103-1121302203303233"></a>

#### `splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1321131122030021-0010103220020013-0003002011330200-2320030102300313-3002321222230111-1333001301231323-1323130202223303-0011323200102022"></a>

<a id="canonical-0001321232210320-1003323121303220-2323103232021003-1301121233121223-1012131320033313-1230000213231030-1312331023302131-1110130210333311"></a>

#### `splunk_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1031033333021002-2013213002003321-1010311121211023-3132121203200222-1133012003000201-1103133202223011-3211002103301121-1221103121103021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-3030211232321010-0002332312301211-2032203223320101-1230130023032103-1312001002030231-3021212332003121-3330131103001221-2330113233101023)
- [splunk_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-004.md#canonical-0121112313322202-2310211222220013-2323010331330003-2212110021230101-1233303212112301-3032023223223001-1320331112213332-0001332102120320)
- [splunk_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-004.md#canonical-3321002133010021-1100120300031012-1033302121131210-1220023323122012-0221111031311111-3003201330220311-0203301220112323-2320313232003200)
- splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info

<a id="canonical-3232103321112020-0023111312201010-0222130313012211-2213300313122012-0100033123321213-0233220100131002-3321211203112113-0003320232222233"></a>

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

<a id="canonical-3032232100001222-3203131310101220-3100020132132121-1321332033233023-2010311203212000-3223203320322002-3322022030322000-0323103321231132"></a>

### Direct properties for `splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info`

<a id="canonical-0301133101332311-1330020030200330-2300331102002000-1021120321000120-1001022003301003-1330300323101302-2121313320110321-3121220131213222"></a>

#### `splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3012023233110300-2100323311133121-2320013322321211-0321233302300023-3323203232303231-1103221001020023-2200121332021022-2233331333210333"></a>

<a id="canonical-3101200033232210-0113331302121033-0102103211323232-3011313033133102-0031232202022113-3130211220210332-0220120210223221-3012220312131220"></a>

#### `splunk_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3222321113000113-2231233223133012-2032013200210220-0201021000323102-1331023333031133-0122211311330020-0200010321023212-3111003231323233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `splunk_receiver.use_tls.no_ca` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [splunk_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1232310221300222-2201230233110021-2102300111113113-3000302000333013-2211233301300200-1312131021121130-3013213100020322-2132133200001300)
- [splunk_receiver.use_tls](data-sources--global_log_receiver--reference--group-004.md#canonical-3030211232321010-0002332312301211-2032203223320101-1230130023032103-1312001002030231-3021212332003121-3330131103001221-2330113233101023)
- splunk_receiver.use_tls.no_ca

<a id="canonical-0121120030121223-3311221203133322-0331001300031213-2123200101302232-2330321212013011-1212222223213010-0333131113101330-2311122221133123"></a>

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

<a id="canonical-1200020203302302-1231220330101021-0122330230312102-1301232213001301-0102021020111200-2000011222121120-3221012030320313-3201023130003001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sumo_logic_receiver` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- sumo_logic_receiver

<a id="canonical-2111201203021003-1010212310123022-2023221323331032-2212330321202303-1123222120131130-2113333103311120-3110301220010121-0200302122002201"></a>

Type: `"single"`. Computed.

Configuration parameter for sumo logic receiver.

Additional upstream details:

Configuration for SumoLogic endpoint.

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

<a id="canonical-1223002010101021-3213031103223213-0300103130100012-3021300122001011-3320120211110000-2110302322022102-1011210321233020-1112003100021000"></a>

### Direct properties for `sumo_logic_receiver`

- [URL](data-sources--global_log_receiver--reference--group-004.md#canonical-3123202213122310-2130203030003301-1321103002013223-1120331211313210-0210302300210210-3200013122032031-3102321331301213-0001313201010221): complete subsection reference.

<a id="canonical-3123202213122310-2130203030003301-1321103002013223-1120331211313210-0210302300210210-3200013122032031-3102321331301213-0001313201010221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sumo_logic_receiver.url` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [sumo_logic_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1200020203302302-1231220330101021-0122330230312102-1301232213001301-0102021020111200-2000011222121120-3221012030320313-3201023130003001)
- sumo_logic_receiver.URL

<a id="canonical-3031113231133321-2013020123113033-3211311322001231-0323202201000203-0230303230001011-1131213213200200-0200233011132202-2110303323102013"></a>

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

<a id="canonical-3031320133212111-2130212130031101-0313230021103012-1023012202102120-3233200200213333-1012302213003103-3122002322021101-3111203200122001"></a>

### Direct properties for `sumo_logic_receiver.url`

- [blindfold_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-0221110123003133-2102020331223201-3023110001211023-0202223323331111-0320021323323112-2331022323002200-3011132120211200-0120032000201002): complete subsection reference.

- [clear_secret_info](data-sources--global_log_receiver--reference--group-004.md#canonical-3022112033113130-3110121321311232-1122311012323133-0110011323101211-1021100232110232-0303123002223213-3012011201032230-0103120122031222): complete subsection reference.

<a id="canonical-0221110123003133-2102020331223201-3023110001211023-0202223323331111-0320021323323112-2331022323002200-3011132120211200-0120032000201002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sumo_logic_receiver.url.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [sumo_logic_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1200020203302302-1231220330101021-0122330230312102-1301232213001301-0102021020111200-2000011222121120-3221012030320313-3201023130003001)
- [sumo_logic_receiver.url](data-sources--global_log_receiver--reference--group-004.md#canonical-3123202213122310-2130203030003301-1321103002013223-1120331211313210-0210302300210210-3200013122032031-3102321331301213-0001313201010221)
- sumo_logic_receiver.URL.blindfold_secret_info

<a id="canonical-2323021002101033-1133310322210301-3120301223233331-3200112230302133-2031000301200123-0302300201133113-1211233300003302-2312300100313001"></a>

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

<a id="canonical-0211100210333011-3330013302002101-0202312013011101-3322012312303132-3200100011101320-3102303102010300-3030002121100332-2202301211320130"></a>

### Direct properties for `sumo_logic_receiver.url.blindfold_secret_info`

<a id="canonical-1121232332123313-0120222303303113-2311222222302330-2122130020010333-1302110313023030-2122021101322213-3103010122102113-0110210101310032"></a>

#### `sumo_logic_receiver.url.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2300330231310021-1302012222330123-1310131131321132-2111021102122021-0020030123013003-1133032123120100-0102301331101023-1112121223002011"></a>

<a id="canonical-2330122321220021-3332230120021103-1300132322221032-1011211203200323-1121211212010200-3100201201201300-0311122022011132-2012212303113022"></a>

#### `sumo_logic_receiver.url.blindfold_secret_info.location` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0301223313113202-1011312230223030-2022203332200310-1023102133313112-1120213221223033-2202331111013232-3322221021130132-3013323030231013"></a>

<a id="canonical-0123032133200310-3231200302302212-1010213032120322-0110320230332001-1110132303131210-2301010022003103-0221323312032101-3220331311132302"></a>

#### `sumo_logic_receiver.url.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3022112033113130-3110121321311232-1122311012323133-0110011323101211-1021100232110232-0303123002223213-3012011201032230-0103120122031222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sumo_logic_receiver.url.clear_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [sumo_logic_receiver](data-sources--global_log_receiver--reference--group-004.md#canonical-1200020203302302-1231220330101021-0122330230312102-1301232213001301-0102021020111200-2000011222121120-3221012030320313-3201023130003001)
- [sumo_logic_receiver.url](data-sources--global_log_receiver--reference--group-004.md#canonical-3123202213122310-2130203030003301-1321103002013223-1120331211313210-0210302300210210-3200013122032031-3102321331301213-0001313201010221)
- sumo_logic_receiver.URL.clear_secret_info

<a id="canonical-0021203223313310-2211132223321003-3203301313223033-2031021001330323-0123202023020213-0021323022133130-1023113021001123-1313121020213002"></a>

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

<a id="canonical-0001333232322200-1022300011103110-1230212301132020-2202232232022020-3012130130312020-0313120000321102-0112213033330002-0323220133002121"></a>

### Direct properties for `sumo_logic_receiver.url.clear_secret_info`

<a id="canonical-1303300110310232-2312131212230133-3211332002003123-1012111001131130-3201302200320122-3231113033320302-0320333210202031-3311112222222012"></a>

#### `sumo_logic_receiver.url.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1213131123010212-2301033211021332-0222321211113120-0202330202030231-0320200010301032-1102121133202132-2110231032122033-1320213100220203"></a>

<a id="canonical-2322332231103320-0112311202033322-1131111001331103-3113210312322200-1323230202032300-1313102323001003-0223220032003321-3222011333313002"></a>

#### `sumo_logic_receiver.url.clear_secret_info.url` property

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
