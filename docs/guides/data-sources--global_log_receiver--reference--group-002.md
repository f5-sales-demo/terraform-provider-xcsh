---
page_title: "xcsh_global_log_receiver reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_global_log_receiver reference."
---

# xcsh_global_log_receiver reference

<a id="canonical-1221133333023232-1001211300011310-2323022322033321-2300202212100011-1133032003010323-2011001001302110-0310231223330121-0133310030030223"></a>

## `datadog_receiver.endpoint` property

Type: `"string"`. Computed.

Exclusive with \[site\] Datadog Endpoint,.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.9,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^https?://[^\\s/$.?#].[^\\s]*$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [no_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-0211013002312300-2111321312100022-2323331132220211-0232030001110203-1030030332113232-0131023003123313-1031233103033031-3102210222121030): complete subsection reference.

<a id="canonical-1133000031211321-2030321133010031-3322103330312201-0030223220233312-0121112011322112-0313111033200130-3310320331323122-1122220233121032"></a>

<a id="canonical-0311113200232233-1103112012001013-0121301333112010-0000013111110320-1302100011330300-1223323203032213-2101101010331200-1032222010121132"></a>

## `datadog_receiver.site` property

Type: `"string"`. Computed.

Exclusive with \[endpoint\] Datadog Site,.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname_or_ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname_or_ip": "true"
  }
}
```

- [use_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-1333322331131132-2322313302313103-3312322033322032-0213223031200131-0212001122020310-3330022121123203-0233301003231333-1112233201100333): complete subsection reference.

<a id="canonical-0313003013113333-0322321012101113-2232210232232131-1123112312130301-0323123203211301-0211311312331331-3100013213032221-3320312202113210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.batch` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1302100131001122-2223211332130100-1000320312110323-3013111102331023-3013223003223020-1013311211012003-3231113123123120-2132312002001102)
- datadog_receiver.batch

<a id="canonical-2032303320333220-1010111220121000-3120101202200030-2303220030321331-2330002010311101-0010222032033221-3300013021300232-2311301300230111"></a>

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

<a id="canonical-2321112220103200-3220203322302231-0130011222031330-3030001031221213-2322302102113010-2021013121211211-2012003030012120-3000220303321322"></a>

### Direct properties for `datadog_receiver.batch`

<a id="canonical-2201201300013211-1130123201032202-3033013131133212-2231013233030130-0012101030300122-0331201133230020-0301100331001202-0131010101323132"></a>

#### `datadog_receiver.batch.max_bytes` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [max_bytes_disabled](data-sources--global_log_receiver--reference--group-002.md#canonical-3320200222102201-0231311133320201-3020030300220032-2020003230130203-1333302233232231-1120232222303021-2123322233103322-1222312032301002): complete subsection reference.

<a id="canonical-3031303110033232-3020311223201120-2221322331021332-1110100023000312-1301111232002132-0321001310120011-2120021321131111-3232030121112303"></a>

<a id="canonical-0311301003330233-0100332310210223-0022302131001003-3130333311102332-2331210302102002-3002010011010331-0300132010122300-2323001300010000"></a>

#### `datadog_receiver.batch.max_events` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [max_events_disabled](data-sources--global_log_receiver--reference--group-002.md#canonical-0030020101103203-3310210130030232-3130110102102132-0103133222222011-2120230030021012-3001011303012011-1321200220112320-2020332022330130): complete subsection reference.

<a id="canonical-1121313132330311-3331030031303121-3002230102101032-3231333101101330-1310201232211033-3132113131323311-3133100132103112-3111011223323330"></a>

<a id="canonical-3020023123011221-2003312332000112-3331031102120103-2030223022033002-3101233132212322-3311302132031013-0022133133332101-2310230231223123"></a>

#### `datadog_receiver.batch.timeout_seconds` property

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
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  }
}
```

- [timeout_seconds_default](data-sources--global_log_receiver--reference--group-002.md#canonical-3001000210033201-3323112331200102-2200020220213232-2210111202221002-3230112303101020-1210112310000023-3331312311123123-1013033011201330): complete subsection reference.

<a id="canonical-3320200222102201-0231311133320201-3020030300220032-2020003230130203-1333302233232231-1120232222303021-2123322233103322-1222312032301002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.batch.max_bytes_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1302100131001122-2223211332130100-1000320312110323-3013111102331023-3013223003223020-1013311211012003-3231113123123120-2132312002001102)
- [datadog_receiver.batch](data-sources--global_log_receiver--reference--group-002.md#canonical-0313003013113333-0322321012101113-2232210232232131-1123112312130301-0323123203211301-0211311312331331-3100013213032221-3320312202113210)
- datadog_receiver.batch.max_bytes_disabled

<a id="canonical-0300120313120120-2202221113213322-3302232200330131-3210103022220310-2313110200331111-2313101032012301-1132221333303301-3121103020212330"></a>

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

<a id="canonical-0030020101103203-3310210130030232-3130110102102132-0103133222222011-2120230030021012-3001011303012011-1321200220112320-2020332022330130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.batch.max_events_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1302100131001122-2223211332130100-1000320312110323-3013111102331023-3013223003223020-1013311211012003-3231113123123120-2132312002001102)
- [datadog_receiver.batch](data-sources--global_log_receiver--reference--group-002.md#canonical-0313003013113333-0322321012101113-2232210232232131-1123112312130301-0323123203211301-0211311312331331-3100013213032221-3320312202113210)
- datadog_receiver.batch.max_events_disabled

<a id="canonical-1203231323222302-2303123302223011-3323210123033010-2331211032103030-3103133001011320-2101110233023000-1121322003211130-3033300333330131"></a>

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

<a id="canonical-3001000210033201-3323112331200102-2200020220213232-2210111202221002-3230112303101020-1210112310000023-3331312311123123-1013033011201330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.batch.timeout_seconds_default` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1302100131001122-2223211332130100-1000320312110323-3013111102331023-3013223003223020-1013311211012003-3231113123123120-2132312002001102)
- [datadog_receiver.batch](data-sources--global_log_receiver--reference--group-002.md#canonical-0313003013113333-0322321012101113-2232210232232131-1123112312130301-0323123203211301-0211311312331331-3100013213032221-3320312202113210)
- datadog_receiver.batch.timeout_seconds_default

<a id="canonical-2311113312202122-0030001323123122-2122022203000220-2131021032111203-0303230330203033-0232112023200132-2131323222033123-0312221000120233"></a>

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

<a id="canonical-0111332233101201-3133220000000011-2101103331333201-0011011330330330-2320103123303111-1230120233123202-2133312211313103-1002223122102203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.compression` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1302100131001122-2223211332130100-1000320312110323-3013111102331023-3013223003223020-1013311211012003-3231113123123120-2132312002001102)
- datadog_receiver.compression

<a id="canonical-2111012231103312-1002111120000130-0123233201020312-3100212332012222-2000212013103302-3131111011113210-2333000321013312-1201003021302321"></a>

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

<a id="canonical-2313123123330003-1323313212111110-0301010112200320-2313210030132123-3020002100212112-2232201000110013-2132022032230300-2100120211331212"></a>

### Direct properties for `datadog_receiver.compression`

- [compression_default](data-sources--global_log_receiver--reference--group-002.md#canonical-1132021113001322-2313310200301232-2103331131322202-1122023202030202-2033200130013332-0130012021103113-3123311013303100-1302333220030310): complete subsection reference.

- [compression_gzip](data-sources--global_log_receiver--reference--group-002.md#canonical-2201212030111122-1210202223311033-1102002323121203-3320120111313212-0221002312121012-3322100312323003-1313033230011321-3220012033020302): complete subsection reference.

- [compression_none](data-sources--global_log_receiver--reference--group-002.md#canonical-2313012122002132-3033222113313212-3330322130111300-3231202133322133-2131030110300320-0030212320120310-1001223102013020-2013133110201212): complete subsection reference.

<a id="canonical-1132021113001322-2313310200301232-2103331131322202-1122023202030202-2033200130013332-0130012021103113-3123311013303100-1302333220030310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.compression.compression_default` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1302100131001122-2223211332130100-1000320312110323-3013111102331023-3013223003223020-1013311211012003-3231113123123120-2132312002001102)
- [datadog_receiver.compression](data-sources--global_log_receiver--reference--group-002.md#canonical-0111332233101201-3133220000000011-2101103331333201-0011011330330330-2320103123303111-1230120233123202-2133312211313103-1002223122102203)
- datadog_receiver.compression.compression_default

<a id="canonical-3210033122031313-2322331022312232-3111023321233320-1321212101202310-3323123103132333-1112313023101011-0132301221103301-1113011311233323"></a>

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

<a id="canonical-2201212030111122-1210202223311033-1102002323121203-3320120111313212-0221002312121012-3322100312323003-1313033230011321-3220012033020302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.compression.compression_gzip` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1302100131001122-2223211332130100-1000320312110323-3013111102331023-3013223003223020-1013311211012003-3231113123123120-2132312002001102)
- [datadog_receiver.compression](data-sources--global_log_receiver--reference--group-002.md#canonical-0111332233101201-3133220000000011-2101103331333201-0011011330330330-2320103123303111-1230120233123202-2133312211313103-1002223122102203)
- datadog_receiver.compression.compression_gzip

<a id="canonical-1003313101203011-0002121021222002-1131221011123121-0030333001332310-1311121311301031-3320120313022022-2002232330012113-2000022131221022"></a>

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

<a id="canonical-2313012122002132-3033222113313212-3330322130111300-3231202133322133-2131030110300320-0030212320120310-1001223102013020-2013133110201212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.compression.compression_none` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1302100131001122-2223211332130100-1000320312110323-3013111102331023-3013223003223020-1013311211012003-3231113123123120-2132312002001102)
- [datadog_receiver.compression](data-sources--global_log_receiver--reference--group-002.md#canonical-0111332233101201-3133220000000011-2101103331333201-0011011330330330-2320103123303111-1230120233123202-2133312211313103-1002223122102203)
- datadog_receiver.compression.compression_none

<a id="canonical-2032213033110020-3311133303310011-0233122210221130-0332302102122132-0311310132021012-3231310310020002-3133232023330331-0121231311011111"></a>

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

<a id="canonical-2321202222031112-3003110221023102-3221030130122331-0233023303332023-2331003113213011-1122310102233013-1013320232010201-3030010220220212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.datadog_api_key` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1302100131001122-2223211332130100-1000320312110323-3013111102331023-3013223003223020-1013311211012003-3231113123123120-2132312002001102)
- datadog_receiver.datadog_api_key

<a id="canonical-0323101312103313-0230000103212323-0210123202113212-2123013010210113-0231132122010111-3130122121223322-2022322330031131-0203321101313112"></a>

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

<a id="canonical-1202102111113122-1113132221213301-3312333032231203-0203233331312230-0000301303121233-0230322231301321-0320001031231330-0011133102302103"></a>

### Direct properties for `datadog_receiver.datadog_api_key`

- [blindfold_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-1213221323103302-2202000332303213-0030313002312212-1023013001232311-3130012012230312-1032122120212333-0030310331022002-2121212003121131): complete subsection reference.

- [clear_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-2030223020001203-0103321302110232-0331033032001000-1230222301233201-1032332132302203-1100110220232310-2310100312223220-1022213212301203): complete subsection reference.

<a id="canonical-1213221323103302-2202000332303213-0030313002312212-1023013001232311-3130012012230312-1032122120212333-0030310331022002-2121212003121131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.datadog_api_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1302100131001122-2223211332130100-1000320312110323-3013111102331023-3013223003223020-1013311211012003-3231113123123120-2132312002001102)
- [datadog_receiver.datadog_api_key](data-sources--global_log_receiver--reference--group-002.md#canonical-2321202222031112-3003110221023102-3221030130122331-0233023303332023-2331003113213011-1122310102233013-1013320232010201-3030010220220212)
- datadog_receiver.datadog_api_key.blindfold_secret_info

<a id="canonical-3332200330132123-0323112202222013-1313202200232322-2113122320222202-0123020313013030-0010213322023021-3210032310113030-1011123330111230"></a>

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

<a id="canonical-2323303120122032-0213322113323203-3012220223012203-3210013023320301-3303323210100030-2131102330333332-3112110311310021-3301220033033221"></a>

### Direct properties for `datadog_receiver.datadog_api_key.blindfold_secret_info`

<a id="canonical-0012132103310200-0311202103123212-0303010331133003-2303310101333101-3113210133022030-2100033203131232-0331003300211031-0332132112200002"></a>

#### `datadog_receiver.datadog_api_key.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-0121032003133111-0222121113033312-3302300001221211-1032131110320231-0112031133030203-0213111331100111-2223102131000312-1213202100232300"></a>

<a id="canonical-1001112001133212-2330231220322303-0332300111322201-3210031311002001-2011200000223202-3320323121023203-1003001322222013-3201022001133220"></a>

#### `datadog_receiver.datadog_api_key.blindfold_secret_info.location` property

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

<a id="canonical-0331010232012313-3113200031122120-3301311322230013-3032023010330220-0321310311323130-0300012203203101-1100133300231221-2233112031211102"></a>

<a id="canonical-3202021323032012-3133001321022212-3003113311200331-0032101123120101-3331132002102110-1030103011020200-3201223303321311-3333030320100232"></a>

#### `datadog_receiver.datadog_api_key.blindfold_secret_info.store_provider` property

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

<a id="canonical-2030223020001203-0103321302110232-0331033032001000-1230222301233201-1032332132302203-1100110220232310-2310100312223220-1022213212301203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.datadog_api_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1302100131001122-2223211332130100-1000320312110323-3013111102331023-3013223003223020-1013311211012003-3231113123123120-2132312002001102)
- [datadog_receiver.datadog_api_key](data-sources--global_log_receiver--reference--group-002.md#canonical-2321202222031112-3003110221023102-3221030130122331-0233023303332023-2331003113213011-1122310102233013-1013320232010201-3030010220220212)
- datadog_receiver.datadog_api_key.clear_secret_info

<a id="canonical-0133211001131000-0032320011101230-2020220311222112-2223201232332301-0200230312020303-3011003000322001-1311021320212112-3333102000303303"></a>

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

<a id="canonical-2021210121010223-1303321331102122-1210311220033210-1003033223131231-0231032230211101-0011322003133232-0333303113330010-2130221223130002"></a>

### Direct properties for `datadog_receiver.datadog_api_key.clear_secret_info`

<a id="canonical-0021202032020031-0230130000322130-1023230121203001-0323020230101100-2030020130311000-3222111131121313-2213030121031230-2131221311031020"></a>

#### `datadog_receiver.datadog_api_key.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0120221103100211-1221122300302310-1231333321110210-3001202200003003-1312322132010322-3300223311332120-1013221110301102-0210221222002213"></a>

<a id="canonical-2303212102123013-1200010333313112-0112331312122212-1301313321110310-3121111213011220-1023030011000311-1202323021122023-3230011033001222"></a>

#### `datadog_receiver.datadog_api_key.clear_secret_info.url` property

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

<a id="canonical-0211013002312300-2111321312100022-2323331132220211-0232030001110203-1030030332113232-0131023003123313-1031233103033031-3102210222121030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.no_tls` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1302100131001122-2223211332130100-1000320312110323-3013111102331023-3013223003223020-1013311211012003-3231113123123120-2132312002001102)
- datadog_receiver.no_tls

<a id="canonical-2023313133203132-2232033120000100-2200002311023200-1232130322303112-0001332320232113-2030320120121102-1203123131022000-3120301030320321"></a>

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

<a id="canonical-1333322331131132-2322313302313103-3312322033322032-0213223031200131-0212001122020310-3330022121123203-0233301003231333-1112233201100333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.use_tls` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1302100131001122-2223211332130100-1000320312110323-3013111102331023-3013223003223020-1013311211012003-3231113123123120-2132312002001102)
- datadog_receiver.use_tls

<a id="canonical-3103323002222123-0322201230323303-2323300302110331-0132103001133223-3000020103120200-1013233302203210-3230211020130010-2210013321113013"></a>

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

<a id="canonical-1220132203122000-2131331232012130-3220311020313311-1123333323221321-2233023020113233-2333321320021032-1013013312133232-3332312131122010"></a>

### Direct properties for `datadog_receiver.use_tls`

- [disable_verify_certificate](data-sources--global_log_receiver--reference--group-002.md#canonical-0333222031212201-2200310113330021-0331202231220122-2222203101030313-1333313101203131-0203322222000003-1302013033203212-2032121011221023): complete subsection reference.

- [disable_verify_hostname](data-sources--global_log_receiver--reference--group-002.md#canonical-2031030212222303-3121003222222323-2022331302032212-1303121120113221-0122033202020120-1032032223301302-3021132213300212-1002130020113033): complete subsection reference.

- [enable_verify_certificate](data-sources--global_log_receiver--reference--group-002.md#canonical-3101302232112220-0100331233201321-2231012011011301-0100011103302131-1333330312322122-2101330231200011-3000201130000330-2100233330220300): complete subsection reference.

- [enable_verify_hostname](data-sources--global_log_receiver--reference--group-002.md#canonical-2221100100221330-3130120022201123-1001110231333233-3332123310301100-2203030231110320-2011030121311320-3022221100300232-3030312331313013): complete subsection reference.

- [mtls_disabled](data-sources--global_log_receiver--reference--group-002.md#canonical-3013200112021121-1023121332212000-0333232123103221-1302130022302203-0313001000003232-0103032321301223-1213123102200031-0031003331033231): complete subsection reference.

- [mtls_enable](data-sources--global_log_receiver--reference--group-002.md#canonical-3333111300121301-1032003322023032-0123303130102300-3031122333220101-1001120203320332-2333322220112303-0202331030232000-0222121011133102): complete subsection reference.

- [no_ca](data-sources--global_log_receiver--reference--group-002.md#canonical-2322301012311030-2212020232001221-2131223231300000-2232212232012113-1313100323002213-0010103002212132-3132302113010202-0202121301210032): complete subsection reference.

<a id="canonical-3300132130321212-0032031020012231-1010030323300003-0131010102111020-0310313232132301-1022012200322313-1121232122322120-1001021321311211"></a>

<a id="canonical-1012133221031132-2000223001320303-2022023110211131-1032013211110103-1312012003033022-0202331020021320-1333231301130300-0212202300130023"></a>

#### `datadog_receiver.use_tls.trusted_ca_url` property

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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-0333222031212201-2200310113330021-0331202231220122-2222203101030313-1333313101203131-0203322222000003-1302013033203212-2032121011221023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.use_tls.disable_verify_certificate` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1302100131001122-2223211332130100-1000320312110323-3013111102331023-3013223003223020-1013311211012003-3231113123123120-2132312002001102)
- [datadog_receiver.use_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-1333322331131132-2322313302313103-3312322033322032-0213223031200131-0212001122020310-3330022121123203-0233301003231333-1112233201100333)
- datadog_receiver.use_tls.disable_verify_certificate

<a id="canonical-3211012121013331-0301122231131101-2110120121103023-1011031101221000-1001021010120120-2103333003200111-1320102111222122-3301022311320223"></a>

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

<a id="canonical-2031030212222303-3121003222222323-2022331302032212-1303121120113221-0122033202020120-1032032223301302-3021132213300212-1002130020113033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.use_tls.disable_verify_hostname` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1302100131001122-2223211332130100-1000320312110323-3013111102331023-3013223003223020-1013311211012003-3231113123123120-2132312002001102)
- [datadog_receiver.use_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-1333322331131132-2322313302313103-3312322033322032-0213223031200131-0212001122020310-3330022121123203-0233301003231333-1112233201100333)
- datadog_receiver.use_tls.disable_verify_hostname

<a id="canonical-2331332232023132-2121000312003003-1331123103010203-0031330323032003-0031003101122201-0111122311033121-3111113123020202-1213020112221011"></a>

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

<a id="canonical-3101302232112220-0100331233201321-2231012011011301-0100011103302131-1333330312322122-2101330231200011-3000201130000330-2100233330220300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.use_tls.enable_verify_certificate` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1302100131001122-2223211332130100-1000320312110323-3013111102331023-3013223003223020-1013311211012003-3231113123123120-2132312002001102)
- [datadog_receiver.use_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-1333322331131132-2322313302313103-3312322033322032-0213223031200131-0212001122020310-3330022121123203-0233301003231333-1112233201100333)
- datadog_receiver.use_tls.enable_verify_certificate

<a id="canonical-2113130303131220-1003232101033221-3210011112112330-0210233330001323-2202312212223013-0033013320113303-3222031222102322-3110012032330021"></a>

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

<a id="canonical-2221100100221330-3130120022201123-1001110231333233-3332123310301100-2203030231110320-2011030121311320-3022221100300232-3030312331313013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.use_tls.enable_verify_hostname` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1302100131001122-2223211332130100-1000320312110323-3013111102331023-3013223003223020-1013311211012003-3231113123123120-2132312002001102)
- [datadog_receiver.use_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-1333322331131132-2322313302313103-3312322033322032-0213223031200131-0212001122020310-3330022121123203-0233301003231333-1112233201100333)
- datadog_receiver.use_tls.enable_verify_hostname

<a id="canonical-2211210201230223-3001032131120313-3333322310200213-3033130020200002-1021202101100223-1321200210130301-1302103021300031-2100021212121322"></a>

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

<a id="canonical-3013200112021121-1023121332212000-0333232123103221-1302130022302203-0313001000003232-0103032321301223-1213123102200031-0031003331033231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.use_tls.mtls_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1302100131001122-2223211332130100-1000320312110323-3013111102331023-3013223003223020-1013311211012003-3231113123123120-2132312002001102)
- [datadog_receiver.use_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-1333322331131132-2322313302313103-3312322033322032-0213223031200131-0212001122020310-3330022121123203-0233301003231333-1112233201100333)
- datadog_receiver.use_tls.mtls_disabled

<a id="canonical-1132320032323322-3320222303202233-1132013301000231-1131132021300223-2311113310023321-1011313203113212-3013231233022333-3331321201112021"></a>

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

<a id="canonical-3333111300121301-1032003322023032-0123303130102300-3031122333220101-1001120203320332-2333322220112303-0202331030232000-0222121011133102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.use_tls.mtls_enable` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1302100131001122-2223211332130100-1000320312110323-3013111102331023-3013223003223020-1013311211012003-3231113123123120-2132312002001102)
- [datadog_receiver.use_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-1333322331131132-2322313302313103-3312322033322032-0213223031200131-0212001122020310-3330022121123203-0233301003231333-1112233201100333)
- datadog_receiver.use_tls.mtls_enable

<a id="canonical-2222101303102021-1330333221023011-3021301212011331-3101212213021203-3011213303333330-2203201203321312-3031023011212013-0002230013302210"></a>

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

<a id="canonical-1111323121302202-3113033301210130-0321313002311203-1302112132222111-0033320103120010-1330132010110322-3000212012001210-1212101122320111"></a>

### Direct properties for `datadog_receiver.use_tls.mtls_enable`

<a id="canonical-0110312333011230-2112002111001111-1222131111302300-3320113102330100-2202230132100120-0101132310121201-1322230032212303-2323101023020012"></a>

#### `datadog_receiver.use_tls.mtls_enable.certificate` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [key_url](data-sources--global_log_receiver--reference--group-002.md#canonical-0300002020110133-0233110100022230-3212002123010013-3200231212311312-2033203000022311-3302001213032330-1032112333301132-1020231313201010): complete subsection reference.

<a id="canonical-0300002020110133-0233110100022230-3212002123010013-3200231212311312-2033203000022311-3302001213032330-1032112333301132-1020231313201010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.use_tls.mtls_enable.key_url` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1302100131001122-2223211332130100-1000320312110323-3013111102331023-3013223003223020-1013311211012003-3231113123123120-2132312002001102)
- [datadog_receiver.use_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-1333322331131132-2322313302313103-3312322033322032-0213223031200131-0212001122020310-3330022121123203-0233301003231333-1112233201100333)
- [datadog_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-002.md#canonical-3333111300121301-1032003322023032-0123303130102300-3031122333220101-1001120203320332-2333322220112303-0202331030232000-0222121011133102)
- datadog_receiver.use_tls.mtls_enable.key_url

<a id="canonical-2301132200033012-0310231230230211-0030103012130300-3322133320233213-3101132133220232-0333223211211311-0112012013122331-1213200322121112"></a>

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

<a id="canonical-3222032203033123-3223200121232202-3022002223130200-1313202012001332-1131302312223211-0212010120031131-0302130102123013-0212002130200001"></a>

### Direct properties for `datadog_receiver.use_tls.mtls_enable.key_url`

- [blindfold_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-1120203013131001-3211103031313320-0213230301021322-0230003123131320-1312223011033333-0103120002232222-3202230121221102-2201323110203130): complete subsection reference.

- [clear_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-3221212213122110-0232321330000113-0321331001003220-2330321002101001-3333310310003120-2203012313031112-1332232110222223-1330022122320010): complete subsection reference.

<a id="canonical-1120203013131001-3211103031313320-0213230301021322-0230003123131320-1312223011033333-0103120002232222-3202230121221102-2201323110203130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1302100131001122-2223211332130100-1000320312110323-3013111102331023-3013223003223020-1013311211012003-3231113123123120-2132312002001102)
- [datadog_receiver.use_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-1333322331131132-2322313302313103-3312322033322032-0213223031200131-0212001122020310-3330022121123203-0233301003231333-1112233201100333)
- [datadog_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-002.md#canonical-3333111300121301-1032003322023032-0123303130102300-3031122333220101-1001120203320332-2333322220112303-0202331030232000-0222121011133102)
- [datadog_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-002.md#canonical-0300002020110133-0233110100022230-3212002123010013-3200231212311312-2033203000022311-3302001213032330-1032112333301132-1020231313201010)
- datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info

<a id="canonical-3323322220333303-1200100303213031-0100030231211102-1020322013000131-1330201233032101-2311201203012032-2002120030130110-3011110003020320"></a>

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

<a id="canonical-0320111233100031-3113302123223032-1330323332120201-3130311111103123-2000223100100101-3001313210200221-3012321232122232-1230013023012000"></a>

### Direct properties for `datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info`

<a id="canonical-1301121212310113-0000202230113010-0100303032311111-2030001213220222-1222203010203000-1221300101002233-2113311331120302-2320113103221013"></a>

#### `datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-2331203333033013-2323202032112313-3030111120131022-0332203131012001-0300321102001031-3131000223222231-0031303100031033-1113102213011322"></a>

<a id="canonical-0200123033302121-0000311001100233-1322231200010211-1200211130023220-2113223123231300-3032333230110213-3122101010322223-3011022031123000"></a>

#### `datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.location` property

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

<a id="canonical-3130331223111032-1121100132313231-2333302332011030-0031331310002211-0312103221110310-2021123110022223-3120010122010311-3232110021210220"></a>

<a id="canonical-0320330222013032-0303022122332002-3232322001031012-1000220002233212-0311112211101103-2103300233220130-2223123101012233-3132231033222200"></a>

#### `datadog_receiver.use_tls.mtls_enable.key_url.blindfold_secret_info.store_provider` property

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

<a id="canonical-3221212213122110-0232321330000113-0321331001003220-2330321002101001-3333310310003120-2203012313031112-1332232110222223-1330022122320010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1302100131001122-2223211332130100-1000320312110323-3013111102331023-3013223003223020-1013311211012003-3231113123123120-2132312002001102)
- [datadog_receiver.use_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-1333322331131132-2322313302313103-3312322033322032-0213223031200131-0212001122020310-3330022121123203-0233301003231333-1112233201100333)
- [datadog_receiver.use_tls.mtls_enable](data-sources--global_log_receiver--reference--group-002.md#canonical-3333111300121301-1032003322023032-0123303130102300-3031122333220101-1001120203320332-2333322220112303-0202331030232000-0222121011133102)
- [datadog_receiver.use_tls.mtls_enable.key_url](data-sources--global_log_receiver--reference--group-002.md#canonical-0300002020110133-0233110100022230-3212002123010013-3200231212311312-2033203000022311-3302001213032330-1032112333301132-1020231313201010)
- datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info

<a id="canonical-2001133223133332-1130001323002101-0301011013332111-1322232223210310-2133221213333231-0013030333033013-0213130011303232-2101121303310230"></a>

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

<a id="canonical-3213303103200000-3023133202101310-1103210110202132-2130232122301211-0111113201113213-2201021002300013-2321103033111121-2113213331311110"></a>

### Direct properties for `datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info`

<a id="canonical-1113320222230323-1112003133223110-0103033220221302-3232123323020202-2123231020311221-0120121301302120-3121111113232123-0211132120231030"></a>

#### `datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2113222020302301-1032010030311033-2333302101310020-1201322122312321-2330200112211321-1201033211020121-2232031330303323-0331222113131312"></a>

<a id="canonical-2000120032033123-2102130213033120-3113133210232302-3330323120112010-3013101231022203-1301132000111321-2211200003231002-1312101020120123"></a>

#### `datadog_receiver.use_tls.mtls_enable.key_url.clear_secret_info.url` property

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

<a id="canonical-2322301012311030-2212020232001221-2131223231300000-2232212232012113-1313100323002213-0010103002212132-3132302113010202-0202121301210032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `datadog_receiver.use_tls.no_ca` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [datadog_receiver](data-sources--global_log_receiver--reference--group-001.md#canonical-1302100131001122-2223211332130100-1000320312110323-3013111102331023-3013223003223020-1013311211012003-3231113123123120-2132312002001102)
- [datadog_receiver.use_tls](data-sources--global_log_receiver--reference--group-002.md#canonical-1333322331131132-2322313302313103-3312322033322032-0213223031200131-0212001122020310-3330022121123203-0233301003231333-1112233201100333)
- datadog_receiver.use_tls.no_ca

<a id="canonical-2022202132121131-2222102202302010-2300230011332133-0302301131133300-3000223000323123-1102103203301233-1022101323122310-2110133113133101"></a>

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

<a id="canonical-2313012101200221-3103103122211233-0010031133112232-3010013110321112-0132120212211113-2313113031330310-0333130212030033-3300322103021000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dns_logs` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- dns_logs

<a id="canonical-0330003033103033-2232132203133203-1222002120211210-2331231322301220-1302122322201100-2332330003001322-1212032300002303-2110022032130313"></a>

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

<a id="canonical-1120122010230312-1032002320321100-0101031101303121-2303120002011020-2310102000121031-0120030130012202-1331001231331332-1200111011022302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp_bucket_receiver` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- gcp_bucket_receiver

<a id="canonical-0101022013031032-0020202300031330-1120302112000333-1303132333330203-3020123133110130-2002133033301033-0121030230213303-3033113231031210"></a>

Type: `"single"`. Computed.

GCP Bucket Configuration for Global Log Receiver.

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

<a id="canonical-3021012322200012-1222130122111333-3203213311220331-2002212200220331-1100002311201212-0023321323101310-0010311010100311-0310233212321000"></a>

### Direct properties for `gcp_bucket_receiver`

- [batch](data-sources--global_log_receiver--reference--group-002.md#canonical-2110231102121320-1222333203301312-3212221122300313-0202222111203100-0110212223003233-3000312330201000-3333113320323300-3231302320312231): complete subsection reference.

<a id="canonical-3012012110231221-3311321310103320-2303310200020310-3220223130120220-0132001223220011-3221323202223011-3330133220102021-1230332012102102"></a>

<a id="canonical-2301101102002220-0003312002230011-3100331003333212-1113110220312113-0013032330103030-0021012021231231-0312321222210212-2023033003133023"></a>

#### `gcp_bucket_receiver.bucket` property

Type: `"string"`. Computed.

GCP Bucket Name. GCP Bucket Name.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 3,
    "pattern": "^[a-z0-9]+[a-z0-9_\\\\.-]+[a-z0-9]$"
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
    "ves.io.schema.rules.string.pattern": "^[a-z0-9]+[a-z0-9_\\\\.-]+[a-z0-9]$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9]+[a-z0-9_\\\\.-]+[a-z0-9]$"
  }
}
```

- [compression](data-sources--global_log_receiver--reference--group-002.md#canonical-3013323021010312-3320211300110001-1210210123210213-2221203030323102-2321321123013302-3023023132013303-1201120033200332-3130122131013001): complete subsection reference.

- [filename_options](data-sources--global_log_receiver--reference--group-002.md#canonical-0111020321332333-3122322332001232-0020322202302102-1112032003200032-0000000030003212-3000313132001101-3003032022010100-1111212111123200): complete subsection reference.

- [gcp_cred](data-sources--global_log_receiver--reference--group-002.md#canonical-3021322211133022-2233031012111010-1021102202133022-1231302233112010-3301031021221022-0222300000122010-2221322103201020-0200012302203332): complete subsection reference.

<a id="canonical-2110231102121320-1222333203301312-3212221122300313-0202222111203100-0110212223003233-3000312330201000-3333113320323300-3231302320312231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp_bucket_receiver.batch` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-1120122010230312-1032002320321100-0101031101303121-2303120002011020-2310102000121031-0120030130012202-1331001231331332-1200111011022302)
- gcp_bucket_receiver.batch

<a id="canonical-2312002313333013-0021201301213130-2323001123310120-3030331123100220-0202321312222103-0101000131231011-2130322202301230-2102133103222121"></a>

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

<a id="canonical-3013023032112222-2031033232202010-0103011030221023-1011113311013223-3000310121002021-3122023003122120-1011210311210121-0302300332320322"></a>

### Direct properties for `gcp_bucket_receiver.batch`

<a id="canonical-2001223103202020-2032201133000230-2201101013112100-2311122030332100-0322131120301110-1210220032012322-0313220132032121-0112313011213330"></a>

#### `gcp_bucket_receiver.batch.max_bytes` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [max_bytes_disabled](data-sources--global_log_receiver--reference--group-002.md#canonical-1312023320231133-1213123000332303-1021300202131222-1212021313333210-0012121113002323-0000113023231102-2113022103201103-1223231121101120): complete subsection reference.

<a id="canonical-1223210221011102-1321201133120302-2212022120012022-1231213221233013-3303323002033333-1113110303332301-1300131303221131-0331201203122322"></a>

<a id="canonical-3320212012201032-3323131130301022-0321202031313011-0103103322332132-1122320022002010-0113303001132313-2001121100210323-3033011303201133"></a>

#### `gcp_bucket_receiver.batch.max_events` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [max_events_disabled](data-sources--global_log_receiver--reference--group-002.md#canonical-1222101130013222-1231213023133101-0121232302133010-3012233131102120-1233011211001122-0100203303031311-2321202113213021-1231321113211313): complete subsection reference.

<a id="canonical-3010323300330332-3110121103102301-1323312211010021-1003322112110302-3111211301303110-0201313020212032-2113232302311030-3030330133120003"></a>

<a id="canonical-3330321122033130-1122010201322213-0022313333210112-2230102023133202-0130123031020233-3313203120032131-1300202102113203-1033222221023301"></a>

#### `gcp_bucket_receiver.batch.timeout_seconds` property

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
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint64.gte": "300",
    "ves.io.schema.rules.uint64.lte": "3600"
  }
}
```

- [timeout_seconds_default](data-sources--global_log_receiver--reference--group-002.md#canonical-0002030120012132-2102031211320313-1133231322311113-1121233303030031-0300332102313100-1121130233023101-0220103311102320-2012022201122231): complete subsection reference.

<a id="canonical-1312023320231133-1213123000332303-1021300202131222-1212021313333210-0012121113002323-0000113023231102-2113022103201103-1223231121101120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp_bucket_receiver.batch.max_bytes_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-1120122010230312-1032002320321100-0101031101303121-2303120002011020-2310102000121031-0120030130012202-1331001231331332-1200111011022302)
- [gcp_bucket_receiver.batch](data-sources--global_log_receiver--reference--group-002.md#canonical-2110231102121320-1222333203301312-3212221122300313-0202222111203100-0110212223003233-3000312330201000-3333113320323300-3231302320312231)
- gcp_bucket_receiver.batch.max_bytes_disabled

<a id="canonical-3231311101021003-0131011021303012-3120311310222110-0013012000103201-3000201333110121-0220112210333110-3122212122130213-3233312201002010"></a>

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

<a id="canonical-1222101130013222-1231213023133101-0121232302133010-3012233131102120-1233011211001122-0100203303031311-2321202113213021-1231321113211313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp_bucket_receiver.batch.max_events_disabled` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-1120122010230312-1032002320321100-0101031101303121-2303120002011020-2310102000121031-0120030130012202-1331001231331332-1200111011022302)
- [gcp_bucket_receiver.batch](data-sources--global_log_receiver--reference--group-002.md#canonical-2110231102121320-1222333203301312-3212221122300313-0202222111203100-0110212223003233-3000312330201000-3333113320323300-3231302320312231)
- gcp_bucket_receiver.batch.max_events_disabled

<a id="canonical-1001021320320233-3302113230123030-3233201232200001-1323102112100102-1230320311313211-3011102010110032-1302303320110222-0330231200203003"></a>

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

<a id="canonical-0002030120012132-2102031211320313-1133231322311113-1121233303030031-0300332102313100-1121130233023101-0220103311102320-2012022201122231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp_bucket_receiver.batch.timeout_seconds_default` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-1120122010230312-1032002320321100-0101031101303121-2303120002011020-2310102000121031-0120030130012202-1331001231331332-1200111011022302)
- [gcp_bucket_receiver.batch](data-sources--global_log_receiver--reference--group-002.md#canonical-2110231102121320-1222333203301312-3212221122300313-0202222111203100-0110212223003233-3000312330201000-3333113320323300-3231302320312231)
- gcp_bucket_receiver.batch.timeout_seconds_default

<a id="canonical-0202100310333002-1303301303213110-0330030123301010-2112113010203021-3033231102023333-3100010332311130-2213113110233132-2321223222231303"></a>

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

<a id="canonical-3013323021010312-3320211300110001-1210210123210213-2221203030323102-2321321123013302-3023023132013303-1201120033200332-3130122131013001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp_bucket_receiver.compression` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-1120122010230312-1032002320321100-0101031101303121-2303120002011020-2310102000121031-0120030130012202-1331001231331332-1200111011022302)
- gcp_bucket_receiver.compression

<a id="canonical-3323333000331233-2320302013202111-1000121130313223-1011213303010133-0200020212012232-1300232133322013-1213202133212333-1220003100130011"></a>

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

<a id="canonical-2231113303023331-0301112210120231-3333110313221310-0121203030121013-3321110002002013-2021021200202110-1121000201232201-0022312333001220"></a>

### Direct properties for `gcp_bucket_receiver.compression`

- [compression_default](data-sources--global_log_receiver--reference--group-002.md#canonical-1103233100220300-2013122213131101-0202122021101313-1232200111302222-1001201221201003-1310230221221032-1232310130313112-2203323322100033): complete subsection reference.

- [compression_gzip](data-sources--global_log_receiver--reference--group-002.md#canonical-1221230300032021-3000332132001211-3112223012133312-2222111020130221-3033311021310321-3103132211212033-2012013123300011-1320030110120033): complete subsection reference.

- [compression_none](data-sources--global_log_receiver--reference--group-002.md#canonical-0230202332301212-2132132201112323-0223013303000000-2213101130332303-0030130220022312-2110211031233200-3210133233220300-0233011331210310): complete subsection reference.

<a id="canonical-1103233100220300-2013122213131101-0202122021101313-1232200111302222-1001201221201003-1310230221221032-1232310130313112-2203323322100033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp_bucket_receiver.compression.compression_default` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-1120122010230312-1032002320321100-0101031101303121-2303120002011020-2310102000121031-0120030130012202-1331001231331332-1200111011022302)
- [gcp_bucket_receiver.compression](data-sources--global_log_receiver--reference--group-002.md#canonical-3013323021010312-3320211300110001-1210210123210213-2221203030323102-2321321123013302-3023023132013303-1201120033200332-3130122131013001)
- gcp_bucket_receiver.compression.compression_default

<a id="canonical-2010133311333111-2100103031220313-3312010123330022-0221003033330332-3102301133232230-1311322303312112-1303213031223031-1203223120321033"></a>

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

<a id="canonical-1221230300032021-3000332132001211-3112223012133312-2222111020130221-3033311021310321-3103132211212033-2012013123300011-1320030110120033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp_bucket_receiver.compression.compression_gzip` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-1120122010230312-1032002320321100-0101031101303121-2303120002011020-2310102000121031-0120030130012202-1331001231331332-1200111011022302)
- [gcp_bucket_receiver.compression](data-sources--global_log_receiver--reference--group-002.md#canonical-3013323021010312-3320211300110001-1210210123210213-2221203030323102-2321321123013302-3023023132013303-1201120033200332-3130122131013001)
- gcp_bucket_receiver.compression.compression_gzip

<a id="canonical-0033111002032020-1331320011122111-2111101132311211-1000121121020301-2233311021003102-3310233123211003-1223212211000231-3012321331221033"></a>

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

<a id="canonical-0230202332301212-2132132201112323-0223013303000000-2213101130332303-0030130220022312-2110211031233200-3210133233220300-0233011331210310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp_bucket_receiver.compression.compression_none` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-1120122010230312-1032002320321100-0101031101303121-2303120002011020-2310102000121031-0120030130012202-1331001231331332-1200111011022302)
- [gcp_bucket_receiver.compression](data-sources--global_log_receiver--reference--group-002.md#canonical-3013323021010312-3320211300110001-1210210123210213-2221203030323102-2321321123013302-3023023132013303-1201120033200332-3130122131013001)
- gcp_bucket_receiver.compression.compression_none

<a id="canonical-1001113101011332-3331230122033022-3111211010132233-2320102021111111-3232222233021022-1213130133210132-2113002331221000-0120002302310330"></a>

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

<a id="canonical-0111020321332333-3122322332001232-0020322202302102-1112032003200032-0000000030003212-3000313132001101-3003032022010100-1111212111123200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp_bucket_receiver.filename_options` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-1120122010230312-1032002320321100-0101031101303121-2303120002011020-2310102000121031-0120030130012202-1331001231331332-1200111011022302)
- gcp_bucket_receiver.filename_options

<a id="canonical-3310321102231203-3310202132031210-2121323231300210-0212210123331000-0321321022120002-0121321113120013-0313002011201212-0102201203003101"></a>

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

<a id="canonical-2220203203101120-2210210101111210-2031210011221312-1022133320333210-1330122320310333-3030231210202101-1002133133303303-0213112332001030"></a>

### Direct properties for `gcp_bucket_receiver.filename_options`

<a id="canonical-1331332032221222-1111131202033310-2331321111210302-0001201322301322-1033110031323101-3121202021213211-0333101231300120-3230301123313112"></a>

#### `gcp_bucket_receiver.filename_options.custom_folder` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [log_type_folder](data-sources--global_log_receiver--reference--group-002.md#canonical-1032222303003012-2333123212212323-2213201013312011-1330212312330231-1023321110223210-0101311210021030-3203311011121033-0331133020131311): complete subsection reference.

- [no_folder](data-sources--global_log_receiver--reference--group-002.md#canonical-3331332100310101-2022322333222030-1121332300022031-2233011130323232-1300031203001011-0023103013100023-2231011303112030-1201033333002222): complete subsection reference.

<a id="canonical-1032222303003012-2333123212212323-2213201013312011-1330212312330231-1023321110223210-0101311210021030-3203311011121033-0331133020131311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp_bucket_receiver.filename_options.log_type_folder` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-1120122010230312-1032002320321100-0101031101303121-2303120002011020-2310102000121031-0120030130012202-1331001231331332-1200111011022302)
- [gcp_bucket_receiver.filename_options](data-sources--global_log_receiver--reference--group-002.md#canonical-0111020321332333-3122322332001232-0020322202302102-1112032003200032-0000000030003212-3000313132001101-3003032022010100-1111212111123200)
- gcp_bucket_receiver.filename_options.log_type_folder

<a id="canonical-2002012322213130-1122310113111002-0003232313322032-3203002320231020-3020220232200331-0221323001320301-2331121010221002-3230202032000231"></a>

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

<a id="canonical-3331332100310101-2022322333222030-1121332300022031-2233011130323232-1300031203001011-0023103013100023-2231011303112030-1201033333002222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp_bucket_receiver.filename_options.no_folder` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-1120122010230312-1032002320321100-0101031101303121-2303120002011020-2310102000121031-0120030130012202-1331001231331332-1200111011022302)
- [gcp_bucket_receiver.filename_options](data-sources--global_log_receiver--reference--group-002.md#canonical-0111020321332333-3122322332001232-0020322202302102-1112032003200032-0000000030003212-3000313132001101-3003032022010100-1111212111123200)
- gcp_bucket_receiver.filename_options.no_folder

<a id="canonical-1023312011011003-0111202123001222-3012020322130122-0120333200322313-3011123003321003-2103231001330033-0333220002000030-1213132300022121"></a>

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

<a id="canonical-3021322211133022-2233031012111010-1021102202133022-1231302233112010-3301031021221022-0222300000122010-2221322103201020-0200012302203332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `gcp_bucket_receiver.gcp_cred` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [gcp_bucket_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-1120122010230312-1032002320321100-0101031101303121-2303120002011020-2310102000121031-0120030130012202-1331001231331332-1200111011022302)
- gcp_bucket_receiver.gcp_cred

<a id="canonical-2222200312302320-1222013032102102-1033313111212022-0100201013101200-1101122101201333-3133012021222021-3203221302333022-2232033003032221"></a>

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

<a id="canonical-1133230212102102-1031011313312232-3330323120202010-3310012320303011-2112221112230202-1300230000001101-1203222211311321-1132122312012021"></a>

### Direct properties for `gcp_bucket_receiver.gcp_cred`

<a id="canonical-0122221231323100-0200002202331002-3013120110200300-0331121001210220-0323330222211330-1220223323120233-2302002213331312-2202331020033103"></a>

#### `gcp_bucket_receiver.gcp_cred.name` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1032201213232210-1213311003003223-2011101312101021-0212010232122023-2210221001033103-1223122322330132-2300010110302213-3132233223312232"></a>

<a id="canonical-0032221321122113-2133122033302103-1101012213120100-1113011101221212-0330101130221311-1110313122121303-1321201133113312-1211022130211323"></a>

#### `gcp_bucket_receiver.gcp_cred.namespace` property

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0101101300330022-2233133122212233-3032110013220322-1311010232333311-1020222133202112-0312112212313201-1023022200332230-0012031222303121"></a>

<a id="canonical-2301201322001303-3201221212103020-3323033221110113-3001231112031003-3011121101200012-0002333333113000-1333333313101230-3002202321230033"></a>

#### `gcp_bucket_receiver.gcp_cred.tenant` property

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3002110111321220-1221133112112010-0201220110211203-1222310310010201-2223203011101013-2311130130220133-2202213221131000-1102223113102220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- http_receiver

<a id="canonical-0101022033231131-0331011310223210-1113000100233231-0202310200201130-3201032123020231-1122203030212220-0223101203210223-1100223030332223"></a>

Type: `"single"`. Computed.

Configuration parameter for http receiver.

Additional upstream details:

Configuration for HTTP endpoint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-auth_choice": "[\"auth_basic\",\"auth_none\",\"auth_token\"]",
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

<a id="canonical-2132200103131311-3012122123200011-0112230102120133-2220133221001323-3213303030301110-3130312213031010-1132220322020023-3120003232303213"></a>

### Direct properties for `http_receiver`

- [auth_basic](data-sources--global_log_receiver--reference--group-002.md#canonical-0331030002301033-1331111130220320-2131103333113321-1130132020013001-3320201230123331-0332123332212033-0123212331233323-1123020203001011): complete subsection reference.

- [auth_none](data-sources--global_log_receiver--reference--group-002.md#canonical-3302311211322121-3131111100123312-2203000223121212-0000212303333012-2230333221312022-1301230013011230-0031201220011120-0230033221302033): complete subsection reference.

- [auth_token](data-sources--global_log_receiver--reference--group-003.md#canonical-3323032303112012-3022221132311311-2010023233121102-3130133130021112-2211023322231311-1131213033021222-3220020013010221-3010030100310012): complete subsection reference.

- [batch](data-sources--global_log_receiver--reference--group-003.md#canonical-1323033310020132-0332013212020031-0212133332203131-1102212322332103-0133202002022001-1321122021112302-0003122230101102-1103000322223201): complete subsection reference.

- [compression](data-sources--global_log_receiver--reference--group-003.md#canonical-1131223203223130-2000122203303022-1023002322213322-2120332113101311-3003321112120110-2022230313200233-1310131000132331-0100333201211202): complete subsection reference.

- [no_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-3203222233032313-0213102110222011-2300132012120130-2123230023320010-0202222331102302-2230031332332231-2220302333122321-2010013030330222): complete subsection reference.

<a id="canonical-2220110333311010-1121230323311311-3213322022222323-0120312332330312-3120302201133203-2131013211122023-2002130221002223-3122111322121030"></a>

<a id="canonical-3032102120312312-1112203203021202-3130323111222113-3203321033313302-2003202013032322-2100110003213310-2201023320210113-3112321120022133"></a>

#### `http_receiver.uri` property

Type: `"string"`. Computed.

HTTP URI is the URI of the HTTP endpoint to send logs to,.

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
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [use_tls](data-sources--global_log_receiver--reference--group-003.md#canonical-3032101331313202-0222233102113001-2122201030333033-1030300200032230-3113223102221213-1101102032232110-0300220211303010-0112200301102330): complete subsection reference.

<a id="canonical-0331030002301033-1331111130220320-2131103333113321-1130132020013001-3320201230123331-0332123332212033-0123212331233323-1123020203001011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.auth_basic` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-3002110111321220-1221133112112010-0201220110211203-1222310310010201-2223203011101013-2311130130220133-2202213221131000-1102223113102220)
- http_receiver.auth_basic

<a id="canonical-1120223022332122-0033233120010022-1032231222332313-3321313103213103-0111123031201311-0223133220311332-3222011212123320-1221222102023232"></a>

Type: `"single"`. Computed.

Authentication parameters to access HTPP Log Receiver Endpoint.

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

<a id="canonical-3232301020131110-2233303223212203-0222103213221221-3322103013303121-0313012103021122-0110110113123022-1101112011022332-0101202201223011"></a>

### Direct properties for `http_receiver.auth_basic`

- [password](data-sources--global_log_receiver--reference--group-002.md#canonical-2333331331030201-0113223030310320-3113222020131330-0103333222333131-1222100023021112-2121030323222113-2331030033212122-3022100120123123): complete subsection reference.

<a id="canonical-0330300300311121-1020331111221131-3003120230031300-0222311210320101-2113003211011132-2323032021223201-1132210023233233-1130213130221233"></a>

<a id="canonical-3012100121000231-0210211020201101-3023201131003221-3101130030331033-1120301102022003-3302002023233231-3210103001011212-2221023232310000"></a>

#### `http_receiver.auth_basic.user_name` property

Type: `"string"`. Computed.

username. HTTP Basic Auth username.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-2333331331030201-0113223030310320-3113222020131330-0103333222333131-1222100023021112-2121030323222113-2331030033212122-3022100120123123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.auth_basic.password` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-3002110111321220-1221133112112010-0201220110211203-1222310310010201-2223203011101013-2311130130220133-2202213221131000-1102223113102220)
- [http_receiver.auth_basic](data-sources--global_log_receiver--reference--group-002.md#canonical-0331030002301033-1331111130220320-2131103333113321-1130132020013001-3320201230123331-0332123332212033-0123212331233323-1123020203001011)
- http_receiver.auth_basic.password

<a id="canonical-3000010211001211-3230331311332221-3301313131300221-3120112001212131-0021033220312333-2121202222202231-2302131232220213-1212313203312001"></a>

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

<a id="canonical-0331031310020121-3033110033130111-0113030233102002-3022121120322001-2223213133321223-2230030001310103-2120132323012212-1100312001033131"></a>

### Direct properties for `http_receiver.auth_basic.password`

- [blindfold_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-1310022101123301-2022020131213210-0002310331210222-0102031120211111-0120032130001323-0102001301210003-3312312321311313-3011023101313000): complete subsection reference.

- [clear_secret_info](data-sources--global_log_receiver--reference--group-002.md#canonical-0201311321120021-3211322331123000-2311032031333123-1000202112003213-1221300231020123-2310031232312101-3233330310222001-1333320033301131): complete subsection reference.

<a id="canonical-1310022101123301-2022020131213210-0002310331210222-0102031120211111-0120032130001323-0102001301210003-3312312321311313-3011023101313000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.auth_basic.password.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-3002110111321220-1221133112112010-0201220110211203-1222310310010201-2223203011101013-2311130130220133-2202213221131000-1102223113102220)
- [http_receiver.auth_basic](data-sources--global_log_receiver--reference--group-002.md#canonical-0331030002301033-1331111130220320-2131103333113321-1130132020013001-3320201230123331-0332123332212033-0123212331233323-1123020203001011)
- [http_receiver.auth_basic.password](data-sources--global_log_receiver--reference--group-002.md#canonical-2333331331030201-0113223030310320-3113222020131330-0103333222333131-1222100023021112-2121030323222113-2331030033212122-3022100120123123)
- http_receiver.auth_basic.password.blindfold_secret_info

<a id="canonical-0201302301101032-0302310111120020-0333231112312302-1033133130030201-0332233020113002-3123231100322202-0100112223332112-1233032031133011"></a>

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

<a id="canonical-1012113333020022-2123303002130123-0023323100102321-0101001203032232-2022022001102022-1302210302213111-1231033210122220-2330201330130113"></a>

### Direct properties for `http_receiver.auth_basic.password.blindfold_secret_info`

<a id="canonical-0330110232311231-3023100020333321-2201013032002213-0003330332303111-0323020230221010-1113030301011320-3101101102022222-0012302302303100"></a>

#### `http_receiver.auth_basic.password.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-1100202133232122-1333132122230223-0110223322011223-0300201302130213-0212002113213200-1130331211323232-2133231222003123-2230123332120233"></a>

<a id="canonical-0202323023300230-2010233013203032-0202232312321212-3222132221323203-3011333001220130-3102322222301210-2123033000022101-2133202213023022"></a>

#### `http_receiver.auth_basic.password.blindfold_secret_info.location` property

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

<a id="canonical-3022322202113323-2202322323311321-2010323203022130-2320030032303313-0101101122211220-3200302311210220-0001022003011022-2100222111010302"></a>

<a id="canonical-3222332312330321-3323213230103131-3110033121321221-2200330022332013-0130000022310013-2020200213331213-2311302200010002-1012100120130300"></a>

#### `http_receiver.auth_basic.password.blindfold_secret_info.store_provider` property

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

<a id="canonical-0201311321120021-3211322331123000-2311032031333123-1000202112003213-1221300231020123-2310031232312101-3233330310222001-1333320033301131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.auth_basic.password.clear_secret_info` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-3002110111321220-1221133112112010-0201220110211203-1222310310010201-2223203011101013-2311130130220133-2202213221131000-1102223113102220)
- [http_receiver.auth_basic](data-sources--global_log_receiver--reference--group-002.md#canonical-0331030002301033-1331111130220320-2131103333113321-1130132020013001-3320201230123331-0332123332212033-0123212331233323-1123020203001011)
- [http_receiver.auth_basic.password](data-sources--global_log_receiver--reference--group-002.md#canonical-2333331331030201-0113223030310320-3113222020131330-0103333222333131-1222100023021112-2121030323222113-2331030033212122-3022100120123123)
- http_receiver.auth_basic.password.clear_secret_info

<a id="canonical-0013120022102030-3021321003030022-1132330013130103-1131211233311112-1120033103030020-3031010001130221-3221321103103000-2310211103010311"></a>

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

<a id="canonical-3201122121023322-0322223131210022-2220333210333000-3210300302032020-1212302222102023-1030202213003313-3311123330132310-1202201331213003"></a>

### Direct properties for `http_receiver.auth_basic.password.clear_secret_info`

<a id="canonical-2313313001130133-1310000221230212-2212323222200201-2013211231023311-1131100233022210-3220011213101100-1220112011011232-1320311302103331"></a>

#### `http_receiver.auth_basic.password.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1222312013202132-1303300111321233-2031012213130203-0321211123131321-2300332102102323-3103000123332322-0011202111223300-3303002000320101"></a>

<a id="canonical-2212211123013231-2002001001133201-0222223112301332-2312111301101122-3202022203231313-0303211300130213-0131232103311032-0313313302031001"></a>

#### `http_receiver.auth_basic.password.clear_secret_info.url` property

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

<a id="canonical-3302311211322121-3131111100123312-2203000223121212-0000212303333012-2230333221312022-1301230013011230-0031201220011120-0230033221302033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_receiver.auth_none` properties

Breadcrumbs:

- [xcsh_global_log_receiver](../data-sources/global_log_receiver.md#canonical-2323320303310003-2020233210110100-2302013311201021-3011012223132222-3230012021230210-1210312103103112-3000110300302123-2103022323231330)
- [Property reference](data-sources--global_log_receiver--reference--group-001.md#canonical-1301322323030333-2002133101301233-0213101202103030-1023010223023313-1031330331332322-1212200100011013-2110333022013330-1120323212103332)
- [http_receiver](data-sources--global_log_receiver--reference--group-002.md#canonical-3002110111321220-1221133112112010-0201220110211203-1222310310010201-2223203011101013-2311130130220133-2202213221131000-1102223113102220)
- http_receiver.auth_none

<a id="canonical-0001110103223111-1331200203113303-2333322331102123-3330222033232200-0002200002303333-2031300333230221-0120122223103023-1131311120220211"></a>

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
