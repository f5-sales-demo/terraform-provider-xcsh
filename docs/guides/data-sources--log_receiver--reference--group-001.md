---
page_title: "xcsh_log_receiver reference"
subcategory: "Monitoring"
description: "Complete grouped canonical reference for xcsh_log_receiver reference."
---

# xcsh_log_receiver reference

<a id="canonical-1302313133112100-1032210321010113-2121123223230031-3320110111212311-0021101303320010-2103020131022310-0310113002022220-3112011312012233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120321303121222-2302321122033103-0010102220301122-2133130221022202-3022011031323303-2332222313110112-2331113103103110-3032002102330131"></a>

## Property reference — Property reference / 121011011001 / 2

Breadcrumbs:

- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-1010023232010022-3110032022020213-2301223122100323-3033320312300203-0103220200031300-1202020311303233-2022232121113022-3202200022031102)
- Property reference

<a id="canonical-1122001223102120-0120333013332112-0013102013230030-2323301212133033-1321033113331123-1100013020032021-2002320232220313-1010310231200122"></a>

## Direct properties — Property reference / 121011011001 / 3

<a id="canonical-0002130212313331-3001232230032120-3302220103021013-0032300212123233-2332020322311312-0223300102330023-3302220333111122-0310130220101332"></a>

<a id="canonical-0032212211123211-2232021122030102-0020010300021000-3013333332132012-0000321232232222-3011120302313220-2323220312312132-3131333333312332"></a>

## annotations property — Property reference / 121011011001 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

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

<a id="canonical-3213302301210301-3022322303013122-2121300311310323-1020231023222132-3133331231332113-2010330122220211-2123201113120322-0230012021022030"></a>

<a id="canonical-1112311300332100-0203021210003102-2110121310213230-1021331210123013-2030221330133230-0123130121210131-3323012123100020-2101112212121020"></a>

## description property — Property reference / 121011011001 / 5

Type: `"string"`. Computed.

Description of the LogReceiver.

Upstream description:

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

<a id="canonical-0233112202331032-0113233032310321-1201300112333023-3023233001013221-1021213211331022-1121110222111033-2300110001333311-3031203020112033"></a>

<a id="canonical-1213202113322210-1302323133323221-0131320032012113-3232312133311303-0203332201013212-3232213103322220-0320210002203111-2321100330101023"></a>

## ID property — Property reference / 121011011001 / 6

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0120222201222233-2312230110001233-3010313323132112-3112021021023132-1221322323113313-2122230212321201-2233133120112021-3232031003300032"></a>

<a id="canonical-1000013203011212-0330112301210001-1111330212130213-2100202131132113-2321202113203122-0112112033331222-2322232302232300-0303011230002210"></a>

## labels property — Property reference / 121011011001 / 7

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

<a id="canonical-1231102023222330-2102220221103131-2210220130012002-1010222031212120-3203130223201132-3013003021011102-2300311012100323-3032101100022213"></a>

<a id="canonical-3332013233203310-1033122222213100-2223202122332201-2103131320322232-0211230311221001-3100311222132131-3020113123232013-2210212101022211"></a>

## name property — Property reference / 121011011001 / 8

Type: `"string"`. Required.

Name of the LogReceiver.

Upstream description:

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

<a id="canonical-2101001212212231-2303221310030130-1331130331320310-2333003200022002-3203003310201233-2022020322300220-1201303302120311-2322211110330302"></a>

<a id="canonical-1231313130022132-1100012023013001-2130031313030323-3133000202230120-0010313310100203-1303211030223121-2311321012223121-2322132112021123"></a>

## namespace property — Property reference / 121011011001 / 9

Type: `"string"`. Required.

Namespace where the LogReceiver exists.

Upstream description:

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

- [site_local](data-sources--log_receiver--reference--group-001.md#canonical-1001313220212102-3303211330122322-0330202222313123-0033302330103232-0000302300201130-0123200110203033-1123202130202303-0332311100023203): complete subsection reference.

- [syslog](data-sources--log_receiver--reference--group-001.md#canonical-2233303021320321-2313223010222101-1300120331303031-0311123103232332-3201330021113323-3131012002022110-0233001002200030-3220320312021330): complete subsection reference.

<a id="canonical-1321100230212131-2121010033133102-1200330130103233-0212130313010202-0303000331011221-2223000231321331-2121321112023022-1321323121122110"></a>

## All schema paths — Property reference / 121011011001 / 10

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--log_receiver--reference--group-001.md#canonical-0002130212313331-3001232230032120-3302220103021013-0032300212123233-2332020322311312-0223300102330023-3302220333111122-0310130220101332) |
| `description` | [description](data-sources--log_receiver--reference--group-001.md#canonical-3213302301210301-3022322303013122-2121300311310323-1020231023222132-3133331231332113-2010330122220211-2123201113120322-0230012021022030) |
| `id` | [ID](data-sources--log_receiver--reference--group-001.md#canonical-0233112202331032-0113233032310321-1201300112333023-3023233001013221-1021213211331022-1121110222111033-2300110001333311-3031203020112033) |
| `labels` | [labels](data-sources--log_receiver--reference--group-001.md#canonical-0120222201222233-2312230110001233-3010313323132112-3112021021023132-1221322323113313-2122230212321201-2233133120112021-3232031003300032) |
| `name` | [name](data-sources--log_receiver--reference--group-001.md#canonical-1231102023222330-2102220221103131-2210220130012002-1010222031212120-3203130223201132-3013003021011102-2300311012100323-3032101100022213) |
| `namespace` | [namespace](data-sources--log_receiver--reference--group-001.md#canonical-2101001212212231-2303221310030130-1331130331320310-2333003200022002-3203003310201233-2022020322300220-1201303302120311-2322211110330302) |
| `site_local` | [site_local](data-sources--log_receiver--reference--group-001.md#canonical-0203122202013300-2203032101033331-0332220301112133-2112012230011032-1001202223201131-2100310223212130-3211211003310120-0113110000310102) |
| `syslog` | [syslog](data-sources--log_receiver--reference--group-001.md#canonical-2333330212033223-2211130100001111-2222200333013103-3130222321021130-3132023321023211-3030222223132203-3232112201320022-1010032100330313) |
| `syslog.syslog_rfc5424` | [syslog.syslog_rfc5424](data-sources--log_receiver--reference--group-001.md#canonical-0102122221101003-0221310213100201-0113003333230030-1313132003230312-1223002310310301-2021110012012301-3122010200301131-3223333213221311) |
| `syslog.tcp_server` | [syslog.tcp_server](data-sources--log_receiver--reference--group-001.md#canonical-3101323320022030-0002023313232332-2111323313020133-1223100020213133-2000001332103113-1323010012331311-0010213002000202-3210220310311122) |
| `syslog.tcp_server.port` | [syslog.tcp_server.port](data-sources--log_receiver--reference--group-001.md#canonical-0221302223133332-3002021033120102-2021333110330021-0310120121121103-1320133023003323-3003311023203112-3111331120001310-2213010330102220) |
| `syslog.tcp_server.server_name` | [syslog.tcp_server.server_name](data-sources--log_receiver--reference--group-001.md#canonical-2111322321130232-0222220022200032-0132310033232333-3102023133111303-2313322303332213-0012231013111311-2223130120001211-3013103030022313) |
| `syslog.tls_server` | [syslog.tls_server](data-sources--log_receiver--reference--group-001.md#canonical-2321221323310222-0122012110011321-3022220230201211-3123001310301220-2112323300133100-3330320131023302-1123222001200212-3320033010111321) |
| `syslog.tls_server.default_https_port` | [syslog.tls_server.default_https_port](data-sources--log_receiver--reference--group-001.md#canonical-1320132002110013-3130233300131131-0200112111211223-3011203330201130-3321321322013203-2110131101320323-1313313312130310-3331031312301230) |
| `syslog.tls_server.default_syslog_tls_port` | [syslog.tls_server.default_syslog_tls_port](data-sources--log_receiver--reference--group-001.md#canonical-1013031330101121-2322032003212010-0312201301231302-0110103302300333-2221333021300111-3110233203330222-1122122233230312-3302221100231230) |
| `syslog.tls_server.mtls_disabled` | [syslog.tls_server.mtls_disabled](data-sources--log_receiver--reference--group-001.md#canonical-0123330300113020-0322222330033213-2133012003311333-2120110010310100-0133112032221231-1122013120323110-0121021002031220-2012231012313330) |
| `syslog.tls_server.mtls_enable` | [syslog.tls_server.mtls_enable](data-sources--log_receiver--reference--group-001.md#canonical-1001102111203122-1201321330233213-2232020210200313-0012131130133003-3112123023133310-0222130330223231-0010111111022132-1123111211122302) |
| `syslog.tls_server.mtls_enable.certificate` | [syslog.tls_server.mtls_enable.certificate](data-sources--log_receiver--reference--group-001.md#canonical-3301310312102132-0232110101110122-2322230331113330-1012103211020130-3303200213110210-3120121323121102-1330022311110301-0033302120211102) |
| `syslog.tls_server.mtls_enable.key_url` | [syslog.tls_server.mtls_enable.key_url](data-sources--log_receiver--reference--group-001.md#canonical-2303002132102202-0012021100231023-3312003223110231-0223010231211210-0101123103231311-0303020321221033-1321222113113032-0033233010210331) |
| `syslog.tls_server.mtls_enable.key_url.blindfold_secret_info` | [syslog.tls_server.mtls_enable.key_url.blindfold_secret_info](data-sources--log_receiver--reference--group-001.md#canonical-0110003232033210-2103222203232031-2021310222322232-1201310312023210-3003201020120103-0233103231303000-3023310112002021-3322203231222133) |
| `syslog.tls_server.mtls_enable.key_url.blindfold_secret_info.decryption_provider` | [syslog.tls_server.mtls_enable.key_url.blindfold_secret_info.decryption_provider](data-sources--log_receiver--reference--group-001.md#canonical-3203133120033113-3301033233022111-2221131233311011-1323203301200021-1312301201032332-2012231033120020-1333030301300000-2131113301031330) |
| `syslog.tls_server.mtls_enable.key_url.blindfold_secret_info.location` | [syslog.tls_server.mtls_enable.key_url.blindfold_secret_info.location](data-sources--log_receiver--reference--group-001.md#canonical-1102011130231233-2003300332002300-2221001300223220-2322221232033133-2133231303023102-1220131220330133-3232020031201133-3102032331000101) |
| `syslog.tls_server.mtls_enable.key_url.blindfold_secret_info.store_provider` | [syslog.tls_server.mtls_enable.key_url.blindfold_secret_info.store_provider](data-sources--log_receiver--reference--group-001.md#canonical-1132211110010231-3031302022303233-1132021110000020-0013011211330203-0332333010231113-3010200211303200-0020210120232321-0102230032121032) |
| `syslog.tls_server.mtls_enable.key_url.clear_secret_info` | [syslog.tls_server.mtls_enable.key_url.clear_secret_info](data-sources--log_receiver--reference--group-001.md#canonical-0113203002323203-0031210321031321-3223212221313121-3111331113130202-0231312033201333-0003311121113332-3313101212130120-1222212031201111) |
| `syslog.tls_server.mtls_enable.key_url.clear_secret_info.provider_ref` | [syslog.tls_server.mtls_enable.key_url.clear_secret_info.provider_ref](data-sources--log_receiver--reference--group-001.md#canonical-0133312100031322-1210102202131323-0333002132310233-2233221213002011-0300310320023112-1133201233330321-0122023013001113-3022030033002122) |
| `syslog.tls_server.mtls_enable.key_url.clear_secret_info.url` | [syslog.tls_server.mtls_enable.key_url.clear_secret_info.url](data-sources--log_receiver--reference--group-001.md#canonical-3233013223021230-3300122202321011-0212213013013213-3003301031311100-1320333102333123-2303222220022110-0003322010010000-2120000121213132) |
| `syslog.tls_server.port` | [syslog.tls_server.port](data-sources--log_receiver--reference--group-001.md#canonical-1313332010021020-2032112013201300-1132203223130030-3020322030301110-0222330122230232-2123321013032030-2210222032233202-2213131221113313) |
| `syslog.tls_server.server_name` | [syslog.tls_server.server_name](data-sources--log_receiver--reference--group-001.md#canonical-2102313020021011-3300120223320121-0201101013221303-1031233301311131-1212000212200310-1010110012101010-1300111320211130-0000201232201101) |
| `syslog.tls_server.trusted_ca_url` | [syslog.tls_server.trusted_ca_url](data-sources--log_receiver--reference--group-001.md#canonical-1011103231030312-0321021120120120-3000023122000100-2230123322023310-1311020313200311-0003300120123102-0230102100033200-3322301232232130) |
| `syslog.tls_server.volterra_ca` | [syslog.tls_server.volterra_ca](data-sources--log_receiver--reference--group-001.md#canonical-1331130111111300-2120011132020110-2330230100223232-3313330103212010-0100020303011003-1030003301201210-0222212133002033-2110113323200132) |
| `syslog.udp_server` | [syslog.udp_server](data-sources--log_receiver--reference--group-001.md#canonical-2202330320033311-0132322102233013-1010231200211102-1113220111003320-2302300230330030-0311312313022311-1003321022103031-0031031301022001) |
| `syslog.udp_server.port` | [syslog.udp_server.port](data-sources--log_receiver--reference--group-001.md#canonical-1023231310310213-2112023302311100-0023301031313011-0310113320022130-3213132302211100-1013010311221130-2113111222330222-0003211000230120) |
| `syslog.udp_server.server_name` | [syslog.udp_server.server_name](data-sources--log_receiver--reference--group-001.md#canonical-2200202031332320-2232222323100022-2211201222032032-3023311121132003-2112201222020033-3001211032230200-2211010003221131-3031122030321002) |

<a id="canonical-2213030212222011-2200130121300303-1112020313230211-3201212130022322-1011011201103230-2310111033322212-1320201123322010-0223101031320023"></a>

## Next pages — Property reference / 121011011001 / 11

- [site_local](data-sources--log_receiver--reference--group-001.md#canonical-1001313220212102-3303211330122322-0330202222313123-0033302330103232-0000302300201130-0123200110203033-1123202130202303-0332311100023203)
- [syslog](data-sources--log_receiver--reference--group-001.md#canonical-2233303021320321-2313223010222101-1300120331303031-0311123103232332-3201330021113323-3131012002022110-0233001002200030-3220320312021330)
- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-1010023232010022-3110032022020213-2301223122100323-3033320312300203-0103220200031300-1202020311303233-2022232121113022-3202200022031102)

<a id="canonical-1001313220212102-3303211330122322-0330202222313123-0033302330103232-0000302300201130-0123200110203033-1123202130202303-0332311100023203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030211233130100-3032122001122302-1110123123322030-0030000122231000-3312031021012201-3003022120200230-3001300000222101-3112211123212213"></a>

## site_local — site_local / 201112023000 / 2

Breadcrumbs:

- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-1010023232010022-3110032022020213-2301223122100323-3033320312300203-0103220200031300-1202020311303233-2022232121113022-3202200022031102)
- [Property reference](data-sources--log_receiver--reference--group-001.md#canonical-1302313133112100-1032210321010113-2121123223230031-3320110111212311-0021101303320010-2103020131022310-0310113002022220-3112011312012233)
- site_local

<a id="canonical-0203122202013300-2203032101033331-0332220301112133-2112012230011032-1001202223201131-2100310223212130-3211211003310120-0113110000310102"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2300003022022033-1300202010213222-1130011002010303-3202231022123232-3130232302320212-0023213220110031-1331332121333212-2110211303123321"></a>

## Direct properties — site_local / 201112023000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3330213110102122-1030032330301200-0002303321003232-0202321023032303-3201130020131011-0111313030330222-0133322000302021-1321133033210112"></a>

## Next pages — site_local / 201112023000 / 4

- [Property reference](data-sources--log_receiver--reference--group-001.md#canonical-1302313133112100-1032210321010113-2121123223230031-3320110111212311-0021101303320010-2103020131022310-0310113002022220-3112011312012233)
- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-1010023232010022-3110032022020213-2301223122100323-3033320312300203-0103220200031300-1202020311303233-2022232121113022-3202200022031102)

<a id="canonical-2233303021320321-2313223010222101-1300120331303031-0311123103232332-3201330021113323-3131012002022110-0233001002200030-3220320312021330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121232233000132-3220331100103322-3202233211301110-3000300322221313-2220000001301321-1333231332231213-1323211000203302-0313102101120211"></a>

## syslog — syslog / 100112002201 / 2

Breadcrumbs:

- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-1010023232010022-3110032022020213-2301223122100323-3033320312300203-0103220200031300-1202020311303233-2022232121113022-3202200022031102)
- [Property reference](data-sources--log_receiver--reference--group-001.md#canonical-1302313133112100-1032210321010113-2121123223230031-3320110111212311-0021101303320010-2103020131022310-0310113002022220-3112011312012233)
- syslog

<a id="canonical-2333330212033223-2211130100001111-2222200333013103-3130222321021130-3132023321023211-3030222223132203-3232112201320022-1010032100330313"></a>

Type: `"single"`. Computed.

Syslog Server Configuration. Configuration for syslog server.

Upstream description:

Configuration for syslog server.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-format_choice": "[\"syslog_rfc5424\"]",
  "x-ves-oneof-field-mode_choice": "[\"tcp_server\",\"tls_server\",\"udp_server\"]"
}
```

<a id="canonical-0331220010032222-0022110102110201-0023021200233320-3101301003312032-1100330013133201-2220201003101320-2300203203030302-1310301023102000"></a>

## Direct properties — syslog / 100112002201 / 3

<a id="canonical-0102122221101003-0221310213100201-0113003333230030-1313132003230312-1223002310310301-2021110012012301-3122010200301131-3223333213221311"></a>

<a id="canonical-2012200302221113-3322103321031231-1233231201031210-0133100310011212-2233310022203301-1232022203200202-0320101120332321-0200101031000013"></a>

## syslog_rfc5424 property — syslog / 100112002201 / 4

Type: `"number"`. Computed.

Exclusive with \[\] Select RFC5424 syslog format and maximum message length.

Upstream description:

Exclusive with \[\] Select RFC5424 syslog format and maximum message length.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 268435456,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 408
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "408",
    "ves.io.schema.rules.uint32.lte": "268435456"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "408",
    "ves.io.schema.rules.uint32.lte": "268435456"
  }
}
```

- [tcp_server](data-sources--log_receiver--reference--group-001.md#canonical-1313311122032313-2121003013021201-2212023223310113-2110000111203023-1321022200012102-3030323322203000-1122021033031211-3233131131030212): complete subsection reference.

- [tls_server](data-sources--log_receiver--reference--group-001.md#canonical-2111303020120022-1313000311003211-1001033313210013-1311113003121311-2321312110131203-3303201132210200-0102001112203010-0312121302300032): complete subsection reference.

- [udp_server](data-sources--log_receiver--reference--group-001.md#canonical-0201323032231330-1002012301103102-3223221321203101-0102010213220312-3222133333221302-2300211032322110-2311110202031230-3010203331102120): complete subsection reference.

<a id="canonical-3022232110330222-0110021323120222-1101133320120202-0100222302031323-2101113311131000-1332012130132012-2020031012332112-2310232232121313"></a>

## Next pages — syslog / 100112002201 / 5

- [syslog.tcp_server](data-sources--log_receiver--reference--group-001.md#canonical-1313311122032313-2121003013021201-2212023223310113-2110000111203023-1321022200012102-3030323322203000-1122021033031211-3233131131030212)
- [syslog.tls_server](data-sources--log_receiver--reference--group-001.md#canonical-2111303020120022-1313000311003211-1001033313210013-1311113003121311-2321312110131203-3303201132210200-0102001112203010-0312121302300032)
- [syslog.udp_server](data-sources--log_receiver--reference--group-001.md#canonical-0201323032231330-1002012301103102-3223221321203101-0102010213220312-3222133333221302-2300211032322110-2311110202031230-3010203331102120)
- [Property reference](data-sources--log_receiver--reference--group-001.md#canonical-1302313133112100-1032210321010113-2121123223230031-3320110111212311-0021101303320010-2103020131022310-0310113002022220-3112011312012233)
- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-1010023232010022-3110032022020213-2301223122100323-3033320312300203-0103220200031300-1202020311303233-2022232121113022-3202200022031102)

<a id="canonical-1313311122032313-2121003013021201-2212023223310113-2110000111203023-1321022200012102-3030323322203000-1122021033031211-3233131131030212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2123130011100230-1212313211103111-2220222031131211-0301301030321033-0321323011023322-3100300031123131-1102011311331000-1002021112233123"></a>

## syslog.tcp_server — tcp_server / 003202123333 / 2

Breadcrumbs:

- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-1010023232010022-3110032022020213-2301223122100323-3033320312300203-0103220200031300-1202020311303233-2022232121113022-3202200022031102)
- [Property reference](data-sources--log_receiver--reference--group-001.md#canonical-1302313133112100-1032210321010113-2121123223230031-3320110111212311-0021101303320010-2103020131022310-0310113002022220-3112011312012233)
- [syslog](data-sources--log_receiver--reference--group-001.md#canonical-2233303021320321-2313223010222101-1300120331303031-0311123103232332-3201330021113323-3131012002022110-0233001002200030-3220320312021330)
- syslog.tcp_server

<a id="canonical-3101323320022030-0002023313232332-2111323313020133-1223100020213133-2000001332103113-1323010012331311-0010213002000202-3210220310311122"></a>

Type: `"single"`. Computed.

TCP Server name and Port Number. Name and port number for a TCP server.

Upstream description:

Name and port number for a TCP server.

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

<a id="canonical-2032201221010300-2112010022300321-0132102001222023-0312133110303301-1101013002000122-0333213033320302-3003130031222122-3232231121221020"></a>

## Direct properties — tcp_server / 003202123333 / 3

<a id="canonical-0221302223133332-3002021033120102-2021333110330021-0310120121121103-1320133023003323-3003311023203112-3111331120001310-2213010330102220"></a>

<a id="canonical-0030320303212101-2021021133020113-3303011302000113-1020231132313323-3012312221133201-2210133203331200-3002123033112110-3203212311223131"></a>

## port property — tcp_server / 003202123333 / 4

Type: `"number"`. Computed.

Port Number. Port number used for communication.

Upstream description:

Port number used for communication.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2111322321130232-0222220022200032-0132310033232333-3102023133111303-2313322303332213-0012231013111311-2223130120001211-3013103030022313"></a>

<a id="canonical-3012221101322100-1323112031220023-1312210102323203-0122322122230100-0001023233330230-0300210121022011-3232211011133303-3103311332200121"></a>

## server_name property — tcp_server / 003202123333 / 5

Type: `"string"`. Computed.

Server name is fully qualified domain name or IP address of the server.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.hostname_or_ip": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname_or_ip": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3332312322322123-1012002311310322-3320220322222201-0210301001022131-0022001332000301-2301021021121033-0130232003311322-1030102221232130"></a>

## Next pages — tcp_server / 003202123333 / 6

- [syslog](data-sources--log_receiver--reference--group-001.md#canonical-2233303021320321-2313223010222101-1300120331303031-0311123103232332-3201330021113323-3131012002022110-0233001002200030-3220320312021330)
- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-1010023232010022-3110032022020213-2301223122100323-3033320312300203-0103220200031300-1202020311303233-2022232121113022-3202200022031102)

<a id="canonical-2111303020120022-1313000311003211-1001033313210013-1311113003121311-2321312110131203-3303201132210200-0102001112203010-0312121302300032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3100321330212333-2211132200121010-3230132201232012-3033001212332331-1032130313031003-3132203331310201-3012302233232320-3213330333010210"></a>

## syslog.tls_server — tls_server / 331222131000 / 2

Breadcrumbs:

- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-1010023232010022-3110032022020213-2301223122100323-3033320312300203-0103220200031300-1202020311303233-2022232121113022-3202200022031102)
- [Property reference](data-sources--log_receiver--reference--group-001.md#canonical-1302313133112100-1032210321010113-2121123223230031-3320110111212311-0021101303320010-2103020131022310-0310113002022220-3112011312012233)
- [syslog](data-sources--log_receiver--reference--group-001.md#canonical-2233303021320321-2313223010222101-1300120331303031-0311123103232332-3201330021113323-3131012002022110-0233001002200030-3220320312021330)
- syslog.tls_server

<a id="canonical-2321221323310222-0122012110011321-3022220230201211-3123001310301220-2112323300133100-3330320131023302-1123222001200212-3320033010111321"></a>

Type: `"single"`. Computed.

TLS config for client of discovery service.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ca_choice": "[\"trusted_ca_url\",\"volterra_ca\"]",
  "x-ves-oneof-field-mtls_choice": "[\"mtls_disabled\",\"mtls_enable\"]",
  "x-ves-oneof-field-port_choice": "[\"default_https_port\",\"default_syslog_tls_port\",\"port\"]"
}
```

<a id="canonical-3311320332031212-1301132332023003-1003013331202123-2221221100312002-2301013002032323-2023230101322100-2212022130030000-3201032011303100"></a>

## Direct properties — tls_server / 331222131000 / 3

- [default_https_port](data-sources--log_receiver--reference--group-001.md#canonical-1103230313230120-1031013311212111-0320122130231302-3033223000112123-1131103300012031-1313020212300121-2222320210230013-1121101221000112): complete subsection reference.

- [default_syslog_tls_port](data-sources--log_receiver--reference--group-001.md#canonical-3330020010112322-1310001313020322-2003112112210133-2200320210321330-1112213230200233-0313202031230300-0300102002312022-0332001303131322): complete subsection reference.

- [mtls_disabled](data-sources--log_receiver--reference--group-001.md#canonical-0113003102030301-1000201223302321-3212211201131111-1201021210330220-3211003100110013-3303220312030223-3002202321323311-2202123211303301): complete subsection reference.

- [mtls_enable](data-sources--log_receiver--reference--group-001.md#canonical-0021031100311221-0200310322310210-1003132312021011-2300020313013311-1212233202221013-0321223012230213-2200330310310103-2212020332100213): complete subsection reference.

<a id="canonical-1313332010021020-2032112013201300-1132203223130030-3020322030301110-0222330122230232-2123321013032030-2210222032233202-2213131221113313"></a>

<a id="canonical-0330000311232130-1213113230222230-0313122002331212-2121232203231301-3121013222312321-1312103332100300-0100303302302000-1023210122103133"></a>

## port property — tls_server / 331222131000 / 4

Type: `"number"`. Computed.

Exclusive with \[default\_https\_port default\_syslog\_tls\_port\] Custom port number used for
communication.

Upstream description:

Exclusive with \[default\_https\_port default\_syslog\_tls\_port\] Custom port number used for
communication.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2102313020021011-3300120223320121-0201101013221303-1031233301311131-1212000212200310-1010110012101010-1300111320211130-0000201232201101"></a>

<a id="canonical-0023232023100213-2113112313213133-0031023030122200-0230113202121101-3023220233200133-1313112222311330-1022013020330033-1202003321230203"></a>

## server_name property — tls_server / 331222131000 / 5

Type: `"string"`. Computed.

ServerName is passed to the server for SNI and is used in the client to check server certificates
against.

Upstream description:

ServerName is passed to the server for SNI and is used in the client to check server certificates
against.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1011103231030312-0321021120120120-3000023122000100-2230123322023310-1311020313200311-0003300120123102-0230102100033200-3322301232232130"></a>

<a id="canonical-3013132030332223-0001223131033100-3222030300313201-1230231231300302-0230013120102201-2321200103021320-0003010001122131-2331323232001302"></a>

## trusted_ca_url property — tls_server / 331222131000 / 6

Type: `"string"`. Computed.

Exclusive with \[volterra\_ca\] The URL or value for trusted Server CA certificate or certificate
chain Certificates in PEM format including the PEM headers.

Upstream description:

Exclusive with \[volterra\_ca\] The URL or value for trusted Server CA certificate or certificate
chain Certificates in PEM format including the PEM headers.

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

- [volterra_ca](data-sources--log_receiver--reference--group-001.md#canonical-0332120322001123-1101031210121330-3321121011022121-0220131021130130-1222220131220020-0000310212233011-3101311103112303-0302222023222203): complete subsection reference.

<a id="canonical-2333301010200332-0301003010221220-2201203233103021-1303203102320012-3102211111020200-3310303033021022-3321011331320300-2031003232030032"></a>

## Next pages — tls_server / 331222131000 / 7

- [syslog.tls_server.default_https_port](data-sources--log_receiver--reference--group-001.md#canonical-1103230313230120-1031013311212111-0320122130231302-3033223000112123-1131103300012031-1313020212300121-2222320210230013-1121101221000112)
- [syslog.tls_server.default_syslog_tls_port](data-sources--log_receiver--reference--group-001.md#canonical-3330020010112322-1310001313020322-2003112112210133-2200320210321330-1112213230200233-0313202031230300-0300102002312022-0332001303131322)
- [syslog.tls_server.mtls_disabled](data-sources--log_receiver--reference--group-001.md#canonical-0113003102030301-1000201223302321-3212211201131111-1201021210330220-3211003100110013-3303220312030223-3002202321323311-2202123211303301)
- [syslog.tls_server.mtls_enable](data-sources--log_receiver--reference--group-001.md#canonical-0021031100311221-0200310322310210-1003132312021011-2300020313013311-1212233202221013-0321223012230213-2200330310310103-2212020332100213)
- [syslog.tls_server.volterra_ca](data-sources--log_receiver--reference--group-001.md#canonical-0332120322001123-1101031210121330-3321121011022121-0220131021130130-1222220131220020-0000310212233011-3101311103112303-0302222023222203)
- [syslog](data-sources--log_receiver--reference--group-001.md#canonical-2233303021320321-2313223010222101-1300120331303031-0311123103232332-3201330021113323-3131012002022110-0233001002200030-3220320312021330)
- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-1010023232010022-3110032022020213-2301223122100323-3033320312300203-0103220200031300-1202020311303233-2022232121113022-3202200022031102)

<a id="canonical-1103230313230120-1031013311212111-0320122130231302-3033223000112123-1131103300012031-1313020212300121-2222320210230013-1121101221000112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331011313012312-0322220013122201-2011031013103212-0223121200220300-3120101303022033-1311021321013033-3013230323111003-3120032223111220"></a>

## syslog.tls_server.default_https_port — default_https_port / 000212002011 / 2

Breadcrumbs:

- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-1010023232010022-3110032022020213-2301223122100323-3033320312300203-0103220200031300-1202020311303233-2022232121113022-3202200022031102)
- [Property reference](data-sources--log_receiver--reference--group-001.md#canonical-1302313133112100-1032210321010113-2121123223230031-3320110111212311-0021101303320010-2103020131022310-0310113002022220-3112011312012233)
- [syslog](data-sources--log_receiver--reference--group-001.md#canonical-2233303021320321-2313223010222101-1300120331303031-0311123103232332-3201330021113323-3131012002022110-0233001002200030-3220320312021330)
- [syslog.tls_server](data-sources--log_receiver--reference--group-001.md#canonical-2111303020120022-1313000311003211-1001033313210013-1311113003121311-2321312110131203-3303201132210200-0102001112203010-0312121302300032)
- syslog.tls_server.default_https_port

<a id="canonical-1320132002110013-3130233300131131-0200112111211223-3011203330201130-3321321322013203-2110131101320323-1313313312130310-3331031312301230"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1310321010230221-3031132221200132-0222123213111000-0103021123233123-0021201213032030-0111222300100113-2222130313302313-1320123210311122"></a>

## Direct properties — default_https_port / 000212002011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2300122233311313-3213122100220213-2113121213203002-3210111200203310-1211202221022311-3222033012223113-3201103333100222-1113203030302232"></a>

## Next pages — default_https_port / 000212002011 / 4

- [syslog.tls_server](data-sources--log_receiver--reference--group-001.md#canonical-2111303020120022-1313000311003211-1001033313210013-1311113003121311-2321312110131203-3303201132210200-0102001112203010-0312121302300032)
- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-1010023232010022-3110032022020213-2301223122100323-3033320312300203-0103220200031300-1202020311303233-2022232121113022-3202200022031102)

<a id="canonical-3330020010112322-1310001313020322-2003112112210133-2200320210321330-1112213230200233-0313202031230300-0300102002312022-0332001303131322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223003131001201-1122113302003210-3022332030321311-1132331322222112-1032312000211302-1232021012021300-3301321232032301-2113113303333000"></a>

## syslog.tls_server.default_syslog_tls_port — default_syslog_tls_port / 011200100332 / 2

Breadcrumbs:

- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-1010023232010022-3110032022020213-2301223122100323-3033320312300203-0103220200031300-1202020311303233-2022232121113022-3202200022031102)
- [Property reference](data-sources--log_receiver--reference--group-001.md#canonical-1302313133112100-1032210321010113-2121123223230031-3320110111212311-0021101303320010-2103020131022310-0310113002022220-3112011312012233)
- [syslog](data-sources--log_receiver--reference--group-001.md#canonical-2233303021320321-2313223010222101-1300120331303031-0311123103232332-3201330021113323-3131012002022110-0233001002200030-3220320312021330)
- [syslog.tls_server](data-sources--log_receiver--reference--group-001.md#canonical-2111303020120022-1313000311003211-1001033313210013-1311113003121311-2321312110131203-3303201132210200-0102001112203010-0312121302300032)
- syslog.tls_server.default_syslog_tls_port

<a id="canonical-1013031330101121-2322032003212010-0312201301231302-0110103302300333-2221333021300111-3110233203330222-1122122233230312-3302221100231230"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3003123022121000-3021021200101310-2212221313231021-3021231231202213-0021122230103102-2013111023131232-0230000112310010-2011100002011223"></a>

## Direct properties — default_syslog_tls_port / 011200100332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1010313232320213-2112230223323233-2132012100221111-1202121221003112-3102023231010131-2023332302231220-3103230103133021-3003000311212130"></a>

## Next pages — default_syslog_tls_port / 011200100332 / 4

- [syslog.tls_server](data-sources--log_receiver--reference--group-001.md#canonical-2111303020120022-1313000311003211-1001033313210013-1311113003121311-2321312110131203-3303201132210200-0102001112203010-0312121302300032)
- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-1010023232010022-3110032022020213-2301223122100323-3033320312300203-0103220200031300-1202020311303233-2022232121113022-3202200022031102)

<a id="canonical-0113003102030301-1000201223302321-3212211201131111-1201021210330220-3211003100110013-3303220312030223-3002202321323311-2202123211303301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201201102113102-0333122313112203-3223313322320113-3302300030033110-3033123021102011-0301231100321211-2211002202130321-3213233003023323"></a>

## syslog.tls_server.mtls_disabled — mtls_disabled / 300203030323 / 2

Breadcrumbs:

- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-1010023232010022-3110032022020213-2301223122100323-3033320312300203-0103220200031300-1202020311303233-2022232121113022-3202200022031102)
- [Property reference](data-sources--log_receiver--reference--group-001.md#canonical-1302313133112100-1032210321010113-2121123223230031-3320110111212311-0021101303320010-2103020131022310-0310113002022220-3112011312012233)
- [syslog](data-sources--log_receiver--reference--group-001.md#canonical-2233303021320321-2313223010222101-1300120331303031-0311123103232332-3201330021113323-3131012002022110-0233001002200030-3220320312021330)
- [syslog.tls_server](data-sources--log_receiver--reference--group-001.md#canonical-2111303020120022-1313000311003211-1001033313210013-1311113003121311-2321312110131203-3303201132210200-0102001112203010-0312121302300032)
- syslog.tls_server.mtls_disabled

<a id="canonical-0123330300113020-0322222330033213-2133012003311333-2120110010310100-0133112032221231-1122013120323110-0121021002031220-2012231012313330"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2212323123303131-1223323031320213-1202110201301210-1112320322200002-1011112321332132-1300011012220032-3320113222330222-3123032320012231"></a>

## Direct properties — mtls_disabled / 300203030323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0332331222023031-0132010121322322-0303313310220331-0113011000333120-0233213323302221-1221330031132100-3211100133110132-0311002222031023"></a>

## Next pages — mtls_disabled / 300203030323 / 4

- [syslog.tls_server](data-sources--log_receiver--reference--group-001.md#canonical-2111303020120022-1313000311003211-1001033313210013-1311113003121311-2321312110131203-3303201132210200-0102001112203010-0312121302300032)
- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-1010023232010022-3110032022020213-2301223122100323-3033320312300203-0103220200031300-1202020311303233-2022232121113022-3202200022031102)

<a id="canonical-0021031100311221-0200310322310210-1003132312021011-2300020313013311-1212233202221013-0321223012230213-2200330310310103-2212020332100213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322230011020002-0332122333001331-2102303232321231-3003022200232201-1131201220000121-1223330212301233-1222232200333133-0232100022310013"></a>

## syslog.tls_server.mtls_enable — mtls_enable / 013002003011 / 2

Breadcrumbs:

- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-1010023232010022-3110032022020213-2301223122100323-3033320312300203-0103220200031300-1202020311303233-2022232121113022-3202200022031102)
- [Property reference](data-sources--log_receiver--reference--group-001.md#canonical-1302313133112100-1032210321010113-2121123223230031-3320110111212311-0021101303320010-2103020131022310-0310113002022220-3112011312012233)
- [syslog](data-sources--log_receiver--reference--group-001.md#canonical-2233303021320321-2313223010222101-1300120331303031-0311123103232332-3201330021113323-3131012002022110-0233001002200030-3220320312021330)
- [syslog.tls_server](data-sources--log_receiver--reference--group-001.md#canonical-2111303020120022-1313000311003211-1001033313210013-1311113003121311-2321312110131203-3303201132210200-0102001112203010-0312121302300032)
- syslog.tls_server.mtls_enable

<a id="canonical-1001102111203122-1201321330233213-2232020210200313-0012131130133003-3112123023133310-0222130330223231-0010111111022132-1123111211122302"></a>

Type: `"single"`. Computed.

Configuration parameter for mtls enable.

Upstream description:

TLS config for client.

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

<a id="canonical-2102211322332223-3332123231222030-1323120111220203-3230223032111101-3220201031030113-0130312021212113-0332322123103331-1313201223123323"></a>

## Direct properties — mtls_enable / 013002003011 / 3

<a id="canonical-3301310312102132-0232110101110122-2322230331113330-1012103211020130-3303200213110210-3120121323121102-1330022311110301-0033302120211102"></a>

<a id="canonical-1033013221321111-2010231212003021-1330301301121223-2311010033323012-1101103331032233-0213010320312203-1321222033110303-3230032001310230"></a>

## certificate property — mtls_enable / 013002003011 / 4

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [key_url](data-sources--log_receiver--reference--group-001.md#canonical-1200123131030100-0220323320233122-1223033302222212-1031030132222211-2121133330333321-0210223000110021-1013213333100300-2032102211111003): complete subsection reference.

<a id="canonical-0100230112223303-2322013022103123-3222220311313201-2230220112010323-3021310001112221-1101303303300301-0333210230121022-3121311323231221"></a>

## Next pages — mtls_enable / 013002003011 / 5

- [syslog.tls_server.mtls_enable.key_url](data-sources--log_receiver--reference--group-001.md#canonical-1200123131030100-0220323320233122-1223033302222212-1031030132222211-2121133330333321-0210223000110021-1013213333100300-2032102211111003)
- [syslog.tls_server](data-sources--log_receiver--reference--group-001.md#canonical-2111303020120022-1313000311003211-1001033313210013-1311113003121311-2321312110131203-3303201132210200-0102001112203010-0312121302300032)
- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-1010023232010022-3110032022020213-2301223122100323-3033320312300203-0103220200031300-1202020311303233-2022232121113022-3202200022031102)

<a id="canonical-1200123131030100-0220323320233122-1223033302222212-1031030132222211-2121133330333321-0210223000110021-1013213333100300-2032102211111003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211022123213102-2102111122233022-0302121103122201-0221020303220131-0131011010022231-2321002223001133-3303031132031222-0023010033033011"></a>

## syslog.tls_server.mtls_enable.key_url — key_url / 121200110221 / 2

Breadcrumbs:

- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-1010023232010022-3110032022020213-2301223122100323-3033320312300203-0103220200031300-1202020311303233-2022232121113022-3202200022031102)
- [Property reference](data-sources--log_receiver--reference--group-001.md#canonical-1302313133112100-1032210321010113-2121123223230031-3320110111212311-0021101303320010-2103020131022310-0310113002022220-3112011312012233)
- [syslog](data-sources--log_receiver--reference--group-001.md#canonical-2233303021320321-2313223010222101-1300120331303031-0311123103232332-3201330021113323-3131012002022110-0233001002200030-3220320312021330)
- [syslog.tls_server](data-sources--log_receiver--reference--group-001.md#canonical-2111303020120022-1313000311003211-1001033313210013-1311113003121311-2321312110131203-3303201132210200-0102001112203010-0312121302300032)
- [syslog.tls_server.mtls_enable](data-sources--log_receiver--reference--group-001.md#canonical-0021031100311221-0200310322310210-1003132312021011-2300020313013311-1212233202221013-0321223012230213-2200330310310103-2212020332100213)
- syslog.tls_server.mtls_enable.key_url

<a id="canonical-2303002132102202-0012021100231023-3312003223110231-0223010231211210-0101123103231311-0303020321221033-1321222113113032-0033233010210331"></a>

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

<a id="canonical-0301223123132033-0131321112130001-1121011133223111-1321020202101311-1311032322130212-3011031100130212-3200130312000123-1221330203310132"></a>

## Direct properties — key_url / 121200110221 / 3

- [blindfold_secret_info](data-sources--log_receiver--reference--group-001.md#canonical-0320023003010000-0011310131021200-3233133311231122-0301332100001012-0030221213223231-1232121132210122-3120301132203120-3211033231013101): complete subsection reference.

- [clear_secret_info](data-sources--log_receiver--reference--group-001.md#canonical-1202232020131122-0012230211002111-1103000101133033-2012102113321122-1012010013223003-1112330102131321-2303312000133133-0310203233231201): complete subsection reference.

<a id="canonical-0333113012310102-1222111310032131-1013121202303032-3032010100010102-3301230213001223-3233000221322203-0020210002120201-1033221003231330"></a>

## Next pages — key_url / 121200110221 / 4

- [syslog.tls_server.mtls_enable.key_url.blindfold_secret_info](data-sources--log_receiver--reference--group-001.md#canonical-0320023003010000-0011310131021200-3233133311231122-0301332100001012-0030221213223231-1232121132210122-3120301132203120-3211033231013101)
- [syslog.tls_server.mtls_enable.key_url.clear_secret_info](data-sources--log_receiver--reference--group-001.md#canonical-1202232020131122-0012230211002111-1103000101133033-2012102113321122-1012010013223003-1112330102131321-2303312000133133-0310203233231201)
- [syslog.tls_server.mtls_enable](data-sources--log_receiver--reference--group-001.md#canonical-0021031100311221-0200310322310210-1003132312021011-2300020313013311-1212233202221013-0321223012230213-2200330310310103-2212020332100213)
- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-1010023232010022-3110032022020213-2301223122100323-3033320312300203-0103220200031300-1202020311303233-2022232121113022-3202200022031102)

<a id="canonical-0320023003010000-0011310131021200-3233133311231122-0301332100001012-0030221213223231-1232121132210122-3120301132203120-3211033231013101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133030231301030-1211223210010012-0311003203030002-1120103233013232-1230331223001011-2301221230120103-1310333113132220-2113231102132320"></a>

## syslog.tls_server.mtls_enable.key_url.blindfold_secret_info — blindfold_secret_info / 120211131232 / 2

Breadcrumbs:

- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-1010023232010022-3110032022020213-2301223122100323-3033320312300203-0103220200031300-1202020311303233-2022232121113022-3202200022031102)
- [Property reference](data-sources--log_receiver--reference--group-001.md#canonical-1302313133112100-1032210321010113-2121123223230031-3320110111212311-0021101303320010-2103020131022310-0310113002022220-3112011312012233)
- [syslog](data-sources--log_receiver--reference--group-001.md#canonical-2233303021320321-2313223010222101-1300120331303031-0311123103232332-3201330021113323-3131012002022110-0233001002200030-3220320312021330)
- [syslog.tls_server](data-sources--log_receiver--reference--group-001.md#canonical-2111303020120022-1313000311003211-1001033313210013-1311113003121311-2321312110131203-3303201132210200-0102001112203010-0312121302300032)
- [syslog.tls_server.mtls_enable](data-sources--log_receiver--reference--group-001.md#canonical-0021031100311221-0200310322310210-1003132312021011-2300020313013311-1212233202221013-0321223012230213-2200330310310103-2212020332100213)
- [syslog.tls_server.mtls_enable.key_url](data-sources--log_receiver--reference--group-001.md#canonical-1200123131030100-0220323320233122-1223033302222212-1031030132222211-2121133330333321-0210223000110021-1013213333100300-2032102211111003)
- syslog.tls_server.mtls_enable.key_url.blindfold_secret_info

<a id="canonical-0110003232033210-2103222203232031-2021310222322232-1201310312023210-3003201020120103-0233103231303000-3023310112002021-3322203231222133"></a>

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

<a id="canonical-3031320002011130-2032203133120322-0112122222233000-2102311000320201-1020031130222031-0222101130300002-2111111313123233-0111103001311133"></a>

## Direct properties — blindfold_secret_info / 120211131232 / 3

<a id="canonical-3203133120033113-3301033233022111-2221131233311011-1323203301200021-1312301201032332-2012231033120020-1333030301300000-2131113301031330"></a>

<a id="canonical-1101001002103220-3113213111230133-1122022332333122-2313110331311123-0110002002013320-2020131003000110-1311203320331212-0311012230300332"></a>

## decryption_provider property — blindfold_secret_info / 120211131232 / 4

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

<a id="canonical-1102011130231233-2003300332002300-2221001300223220-2322221232033133-2133231303023102-1220131220330133-3232020031201133-3102032331000101"></a>

<a id="canonical-2320010002220200-2131022231120213-2010023130023012-0230000023312132-2001301323321002-2310330013302332-0131220301113233-2100001330233211"></a>

## location property — blindfold_secret_info / 120211131232 / 5

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

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

<a id="canonical-1132211110010231-3031302022303233-1132021110000020-0013011211330203-0332333010231113-3010200211303200-0020210120232321-0102230032121032"></a>

<a id="canonical-0330000233121111-0303111302220002-2023031123222203-1103021101101030-0233012201321101-1021221313032101-3021322311333022-3002110021333013"></a>

## store_provider property — blindfold_secret_info / 120211131232 / 6

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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

<a id="canonical-2331232230311012-2033200302121300-3310000001211100-1203022000211312-0301203201310202-2133231122011012-3122321221133312-0123200220101000"></a>

## Next pages — blindfold_secret_info / 120211131232 / 7

- [syslog.tls_server.mtls_enable.key_url](data-sources--log_receiver--reference--group-001.md#canonical-1200123131030100-0220323320233122-1223033302222212-1031030132222211-2121133330333321-0210223000110021-1013213333100300-2032102211111003)
- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-1010023232010022-3110032022020213-2301223122100323-3033320312300203-0103220200031300-1202020311303233-2022232121113022-3202200022031102)

<a id="canonical-1202232020131122-0012230211002111-1103000101133033-2012102113321122-1012010013223003-1112330102131321-2303312000133133-0310203233231201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213103013310032-2331101101323233-0210121023123123-1220302011323331-3100022202302231-2120122203212013-1031320311223310-1303001233021222"></a>

## syslog.tls_server.mtls_enable.key_url.clear_secret_info — clear_secret_info / 013210301131 / 2

Breadcrumbs:

- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-1010023232010022-3110032022020213-2301223122100323-3033320312300203-0103220200031300-1202020311303233-2022232121113022-3202200022031102)
- [Property reference](data-sources--log_receiver--reference--group-001.md#canonical-1302313133112100-1032210321010113-2121123223230031-3320110111212311-0021101303320010-2103020131022310-0310113002022220-3112011312012233)
- [syslog](data-sources--log_receiver--reference--group-001.md#canonical-2233303021320321-2313223010222101-1300120331303031-0311123103232332-3201330021113323-3131012002022110-0233001002200030-3220320312021330)
- [syslog.tls_server](data-sources--log_receiver--reference--group-001.md#canonical-2111303020120022-1313000311003211-1001033313210013-1311113003121311-2321312110131203-3303201132210200-0102001112203010-0312121302300032)
- [syslog.tls_server.mtls_enable](data-sources--log_receiver--reference--group-001.md#canonical-0021031100311221-0200310322310210-1003132312021011-2300020313013311-1212233202221013-0321223012230213-2200330310310103-2212020332100213)
- [syslog.tls_server.mtls_enable.key_url](data-sources--log_receiver--reference--group-001.md#canonical-1200123131030100-0220323320233122-1223033302222212-1031030132222211-2121133330333321-0210223000110021-1013213333100300-2032102211111003)
- syslog.tls_server.mtls_enable.key_url.clear_secret_info

<a id="canonical-0113203002323203-0031210321031321-3223212221313121-3111331113130202-0231312033201333-0003311121113332-3313101212130120-1222212031201111"></a>

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

<a id="canonical-1021312203331211-2332202031102030-0202110303022121-2300203011223321-3133111002000202-2000212100032012-2113013302011021-3211021312111102"></a>

## Direct properties — clear_secret_info / 013210301131 / 3

<a id="canonical-0133312100031322-1210102202131323-0333002132310233-2233221213002011-0300310320023112-1133201233330321-0122023013001113-3022030033002122"></a>

<a id="canonical-3311120331320203-1012332111020323-3210300301222221-2112302203301231-0332231331213101-2111222133331203-0310122030333223-1212312023331220"></a>

## provider_ref property — clear_secret_info / 013210301131 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3233013223021230-3300122202321011-0212213013013213-3003301031311100-1320333102333123-2303222220022110-0003322010010000-2120000121213132"></a>

<a id="canonical-1102100212133003-2232023321002321-3013023323232000-2232011322322323-0121223303032231-2313220312113301-2231231003010232-2013232302000111"></a>

## URL property — clear_secret_info / 013210301131 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

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

<a id="canonical-2011003113222102-3320001322012322-1002321002332212-0331300313122121-3233322320100231-2323311232001332-1002301321223000-1230102233022212"></a>

## Next pages — clear_secret_info / 013210301131 / 6

- [syslog.tls_server.mtls_enable.key_url](data-sources--log_receiver--reference--group-001.md#canonical-1200123131030100-0220323320233122-1223033302222212-1031030132222211-2121133330333321-0210223000110021-1013213333100300-2032102211111003)
- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-1010023232010022-3110032022020213-2301223122100323-3033320312300203-0103220200031300-1202020311303233-2022232121113022-3202200022031102)

<a id="canonical-0332120322001123-1101031210121330-3321121011022121-0220131021130130-1222220131220020-0000310212233011-3101311103112303-0302222023222203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210022021303200-1121123101330031-3332202210330032-1023000102131100-2103000301300213-2301103302121023-3033331032003121-0111131031021223"></a>

## syslog.tls_server.volterra_ca — volterra_ca / 223131123002 / 2

Breadcrumbs:

- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-1010023232010022-3110032022020213-2301223122100323-3033320312300203-0103220200031300-1202020311303233-2022232121113022-3202200022031102)
- [Property reference](data-sources--log_receiver--reference--group-001.md#canonical-1302313133112100-1032210321010113-2121123223230031-3320110111212311-0021101303320010-2103020131022310-0310113002022220-3112011312012233)
- [syslog](data-sources--log_receiver--reference--group-001.md#canonical-2233303021320321-2313223010222101-1300120331303031-0311123103232332-3201330021113323-3131012002022110-0233001002200030-3220320312021330)
- [syslog.tls_server](data-sources--log_receiver--reference--group-001.md#canonical-2111303020120022-1313000311003211-1001033313210013-1311113003121311-2321312110131203-3303201132210200-0102001112203010-0312121302300032)
- syslog.tls_server.volterra_ca

<a id="canonical-1331130111111300-2120011132020110-2330230100223232-3313330103212010-0100020303011003-1030003301201210-0222212133002033-2110113323200132"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for volterra ca.

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

<a id="canonical-2323011212320213-0312010221331103-2132022110322012-0020132213313121-0330132003002120-0131323031213023-3321012002000131-3112213010211113"></a>

## Direct properties — volterra_ca / 223131123002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0303222321132022-1212121130302030-2001013203213011-2002103012120211-1121131320210103-0321212101221212-2312200303120120-1022301222000221"></a>

## Next pages — volterra_ca / 223131123002 / 4

- [syslog.tls_server](data-sources--log_receiver--reference--group-001.md#canonical-2111303020120022-1313000311003211-1001033313210013-1311113003121311-2321312110131203-3303201132210200-0102001112203010-0312121302300032)
- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-1010023232010022-3110032022020213-2301223122100323-3033320312300203-0103220200031300-1202020311303233-2022232121113022-3202200022031102)

<a id="canonical-0201323032231330-1002012301103102-3223221321203101-0102010213220312-3222133333221302-2300211032322110-2311110202031230-3010203331102120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202323122022132-0001303000012201-1212111311112033-1331021321323302-0031130312022330-0100323102211033-2102233303123310-2221210232011313"></a>

## syslog.udp_server — udp_server / 121232111302 / 2

Breadcrumbs:

- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-1010023232010022-3110032022020213-2301223122100323-3033320312300203-0103220200031300-1202020311303233-2022232121113022-3202200022031102)
- [Property reference](data-sources--log_receiver--reference--group-001.md#canonical-1302313133112100-1032210321010113-2121123223230031-3320110111212311-0021101303320010-2103020131022310-0310113002022220-3112011312012233)
- [syslog](data-sources--log_receiver--reference--group-001.md#canonical-2233303021320321-2313223010222101-1300120331303031-0311123103232332-3201330021113323-3131012002022110-0233001002200030-3220320312021330)
- syslog.udp_server

<a id="canonical-2202330320033311-0132322102233013-1010231200211102-1113220111003320-2302300230330030-0311312313022311-1003321022103031-0031031301022001"></a>

Type: `"single"`. Computed.

UDP Server Name and Port Number. Name and port number for a UDP server.

Upstream description:

Name and port number for a UDP server.

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

<a id="canonical-3032023322321121-2103332013202321-2303301302312222-3023200001302022-1301312200022012-3012322203220213-0210030200230230-2212122133003120"></a>

## Direct properties — udp_server / 121232111302 / 3

<a id="canonical-1023231310310213-2112023302311100-0023301031313011-0310113320022130-3213132302211100-1013010311221130-2113111222330222-0003211000230120"></a>

<a id="canonical-2012132202221213-2032123121302201-1013321003320322-0201132133220012-2110012030120022-2231312122313100-1000200001311200-1031030322210201"></a>

## port property — udp_server / 121232111302 / 4

Type: `"number"`. Computed.

Port Number. Port number used for communication.

Upstream description:

Port number used for communication.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2200202031332320-2232222323100022-2211201222032032-3023311121132003-2112201222020033-3001211032230200-2211010003221131-3031122030321002"></a>

<a id="canonical-0020302030200211-0003002102202222-1331232022213321-1211221312221331-3103022310213132-2122002101303303-1122203030220303-0121032222020103"></a>

## server_name property — udp_server / 121232111302 / 5

Type: `"string"`. Computed.

Server name is fully qualified domain name or IP address of the server.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.hostname_or_ip": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname_or_ip": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3112220033133131-2200120111211032-3010200330030123-2221220100010130-0033133132211100-2111202002221031-2302111001013310-3011302000201101"></a>

## Next pages — udp_server / 121232111302 / 6

- [syslog](data-sources--log_receiver--reference--group-001.md#canonical-2233303021320321-2313223010222101-1300120331303031-0311123103232332-3201330021113323-3131012002022110-0233001002200030-3220320312021330)
- [xcsh_log_receiver](../data-sources/log_receiver.md#canonical-1010023232010022-3110032022020213-2301223122100323-3033320312300203-0103220200031300-1202020311303233-2022232121113022-3202200022031102)
