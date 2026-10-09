---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-0211012312122312-1333133112322302-0301320110110110-1223302122303033-2112230030100311-2320010212233132-1113002320100321-1300010211220330"></a>

## Direct properties for `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules`

- [action_block](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2131223020301322-3003100021032330-1210101322232112-2113130012123200-2221213000203223-0303311110301331-3022022323312133-3132102301332331): complete subsection reference.

- [action_report](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0311030022230303-2110230221032132-0332031000200302-0201303011301232-1023000300322031-0310002101013313-3233013001313112-1122003202013102): complete subsection reference.

- [action_skip](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3203133033020130-2001121020123320-3332012003003030-2212112013111102-3121301133003313-2321233232221020-3012313221112133-1032013022202113): complete subsection reference.

- [api_endpoint](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3003231311130233-0223001212230013-0212113331312032-3013112131111230-2003131012233303-0202122333013000-1213303012102100-1023210112102301): complete subsection reference.

<a id="canonical-0333102001201213-3121323232033322-2321023220031320-0300312132000202-2211232303102210-3303333311213033-3131230230030331-2101330311023230"></a>

<a id="canonical-2022223012111311-0303022201000001-1120033013020020-1322033323322021-3321001330001031-0211032230033010-0331300231222221-2102013322301003"></a>

### `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_group` property

Type: `"string"`. Computed.

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-1031231313332300-0021230131031003-2100312010111321-0113223113033301-0231023303003233-2010111332032231-2232222102003130-0020302312020100"></a>

<a id="canonical-2032103033023311-3023013231331033-3221110212313033-2003331330000101-0130303202101103-1203020130221211-1122032130103333-3323313220203330"></a>

### `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.base_path` property

Type: `"string"`. Computed.

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [metadata](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0012222112331100-0303231311123222-2221002000301331-3133032301313022-3303300021003313-0332320133322232-0311001301010103-3231133003323001): complete subsection reference.

<a id="canonical-2131223020301322-3003100021032330-1210101322232112-2113130012123200-2221213000203223-0303311110301331-3022022323312133-3132102301332331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0223210132303111-0231131220101323-2132302032203111-0001030031010001-0200012101002122-1010002100322310-2000321102222021-3310311101031120)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3031331123301203-0010030212320033-1122302301331112-2002302010330130-1112002103213113-3312020210112012-2130212121313012-0130332132201332)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3200002103223002-1030201100021001-1100022310131001-0302203032313302-1201303123122100-0332211102220210-2312331123311300-3100120131120210)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block

<a id="canonical-1011301033030331-0030313301312313-1320211300303221-1121332223101020-0000200233133302-2210332303211322-0011300300312110-2021300010011012"></a>

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

<a id="canonical-0311030022230303-2110230221032132-0332031000200302-0201303011301232-1023000300322031-0310002101013313-3233013001313112-1122003202013102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0223210132303111-0231131220101323-2132302032203111-0001030031010001-0200012101002122-1010002100322310-2000321102222021-3310311101031120)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3031331123301203-0010030212320033-1122302301331112-2002302010330130-1112002103213113-3312020210112012-2130212121313012-0130332132201332)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3200002103223002-1030201100021001-1100022310131001-0302203032313302-1201303123122100-0332211102220210-2312331123311300-3100120131120210)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report

<a id="canonical-2320332220300133-2022112102101312-3321101133133123-2102313230033313-0203132222010201-1202200100122031-3003022121002010-3110300320200032"></a>

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

<a id="canonical-3203133033020130-2001121020123320-3332012003003030-2212112013111102-3121301133003313-2321233232221020-3012313221112133-1032013022202113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0223210132303111-0231131220101323-2132302032203111-0001030031010001-0200012101002122-1010002100322310-2000321102222021-3310311101031120)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3031331123301203-0010030212320033-1122302301331112-2002302010330130-1112002103213113-3312020210112012-2130212121313012-0130332132201332)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3200002103223002-1030201100021001-1100022310131001-0302203032313302-1201303123122100-0332211102220210-2312331123311300-3100120131120210)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip

<a id="canonical-0112022101130002-2003210300233310-2231202030201101-1103102332000002-1132113330132033-3231330033123323-0121111330032030-0223313211020223"></a>

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

<a id="canonical-3003231311130233-0223001212230013-0212113331312032-3013112131111230-2003131012233303-0202122333013000-1213303012102100-1023210112102301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0223210132303111-0231131220101323-2132302032203111-0001030031010001-0200012101002122-1010002100322310-2000321102222021-3310311101031120)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3031331123301203-0010030212320033-1122302301331112-2002302010330130-1112002103213113-3312020210112012-2130212121313012-0130332132201332)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3200002103223002-1030201100021001-1100022310131001-0302203032313302-1201303123122100-0332211102220210-2312331123311300-3100120131120210)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint

<a id="canonical-2231010133121310-1213202332330220-3131313012012322-2102330131213232-3033103032012002-2021120000022022-3220122031001301-1010323010030002"></a>

Type: `"single"`. Computed.

API Endpoint. This defines API endpoint.

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

<a id="canonical-1010131213122012-3233011011200323-1232112130132103-1111020211332022-1102213120210321-3020323230333300-0200030322012011-1101110333303202"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint`

<a id="canonical-0320303303011133-3223022321011132-0220030233310103-3323100302200013-3011031202222223-1330102022221120-1031302220202122-3011013320020212"></a>

#### `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint.methods` property

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0100012310323110-2033032033332021-2110201022021010-2001320022330222-3212223023030013-0313122201210320-2321331202000031-0300220031100032"></a>

<a id="canonical-0230221133131230-1101310033101100-2312000233123303-1203211201223103-2220103310011332-2323321232300202-1100133221111210-0233003312022333"></a>

#### `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint.path` property

Type: `"string"`. Computed.

Path. Path to be matched.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

<a id="canonical-0012222112331100-0303231311123222-2221002000301331-3133032301313022-3303300021003313-0332320133322232-0311001301010103-3231133003323001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0223210132303111-0231131220101323-2132302032203111-0001030031010001-0200012101002122-1010002100322310-2000321102222021-3310311101031120)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3031331123301203-0010030212320033-1122302301331112-2002302010330130-1112002103213113-3312020210112012-2130212121313012-0130332132201332)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-3200002103223002-1030201100021001-1100022310131001-0302203032313302-1201303123122100-0332211102220210-2312331123311300-3100120131120210)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata

<a id="canonical-1210120100332311-1102322230101003-0323332011000321-3211100001212023-2202023223321033-2013333123111231-1220303010020203-1313211133132111"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-3202230030332003-2322203210132301-0320320133011122-1210031302121102-3323213302212122-2031313031032302-3002223210303101-2133301213133121"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata`

<a id="canonical-2323033323300131-3211213022333330-3120313031103010-1223022011312211-2332310123322322-1220031333110132-0302031012002131-3102310311032211"></a>

#### `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-3012111110001310-0320332122300121-1123003131000000-3023312012020212-2011011233323223-2301031221302122-3333321131001112-1103011221300101"></a>

<a id="canonical-3032223012123233-3123230220333223-2113200001023202-0201121211033220-0213132112011231-1020001231131312-2000221002322311-0323312313312330"></a>

#### `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata.name` property

Type: `"string"`. Computed.

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-0313030322231232-2030011003123221-3111133030113003-0311131002031330-2123113110210002-0100121123121212-0222223012000321-0322000222331112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.settings` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- api_specification.validation_all_spec_endpoints.settings

<a id="canonical-2132231120232323-3102332213300103-0332011323311313-0131201112131010-2032123333330201-2111330101221222-3231130223201002-2110232312220113"></a>

Type: `"single"`. Computed.

OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom
list' enforcement.

Additional upstream details:

OpenAPI specification validation settings relevant for "API Inventory" enforcement and for "Custom
list" enforcement.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-fail_configuration": "[]",
  "x-ves-oneof-field-oversized_body_choice": "[\"oversized_body_fail_validation\",\"oversized_body_skip_validation\"]",
  "x-ves-oneof-field-property_validation_settings_choice": "[\"property_validation_settings_custom\",\"property_validation_settings_default\"]"
}
```

<a id="canonical-3001011011100223-1320111012132220-3320323111211000-1331123232221020-0012231223011032-2101000323203223-3122220302222131-1322301113020012"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.settings`

- [oversized_body_fail_validation](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2032123213203300-3012321101203231-3031322012132330-2312112230303101-1031030110332030-3131132312022323-2022002033121210-2032113331022230): complete subsection reference.

- [oversized_body_skip_validation](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1322123231323112-0302010010111030-3222232133100230-0302132111101110-2322103311313233-2012201321013331-0021022112030220-0030013322321220): complete subsection reference.

- [property_validation_settings_custom](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3122223131200201-1103022231023032-3220222100222030-2032030233023233-1233300210331120-2030212211003130-2003213120012111-0313311330013111): complete subsection reference.

- [property_validation_settings_default](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1231323101123001-1202002133333233-2213212021020013-3220133302210223-0011013301202230-1002310131033101-2030310111110301-1111211200312123): complete subsection reference.

<a id="canonical-2032123213203300-3012321101203231-3031322012132330-2312112230303101-1031030110332030-3131132312022323-2022002033121210-2032113331022230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_validation` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0313030322231232-2030011003123221-3111133030113003-0311131002031330-2123113110210002-0100121123121212-0222223012000321-0322000222331112)
- api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_validation

<a id="canonical-1122110330103303-2131131331232203-1100122021233121-2321131321103122-2013213013022220-3211231201321130-2302001110333230-3210330220220223"></a>

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

<a id="canonical-1322123231323112-0302010010111030-3222232133100230-0302132111101110-2322103311313233-2012201321013331-0021022112030220-0030013322321220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_validation` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0313030322231232-2030011003123221-3111133030113003-0311131002031330-2123113110210002-0100121123121212-0222223012000321-0322000222331112)
- api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_validation

<a id="canonical-2130301011103132-1333231212023113-2120212020233303-0112102120333003-2110221122310133-0101023231031330-0321302320021213-3321011202300002"></a>

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

<a id="canonical-3122223131200201-1103022231023032-3220222100222030-2032030233023233-1233300210331120-2030212211003130-2003213120012111-0313311330013111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0313030322231232-2030011003123221-3111133030113003-0311131002031330-2123113110210002-0100121123121212-0222223012000321-0322000222331112)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom

<a id="canonical-2203313003112300-3012323202322301-1023313021232301-1121202203011202-0322100121031121-3100002030310021-2310130232131221-2003022131233133"></a>

Type: `"single"`. Computed.

Configuration parameter for property validation settings custom.

Additional upstream details:

Custom property validation settings.

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

<a id="canonical-1333132210322213-0102111023332121-0330201032020301-1000201202332221-3023133011022221-2013220320210011-1133022223032203-3120013110111320"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom`

- [query_parameters](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0012220213011221-1010012002013200-3331110210130300-3231321203112010-0113030300132131-1323213200321221-2222210302120213-1221120123100301): complete subsection reference.

<a id="canonical-0012220213011221-1010012002013200-3331110210130300-3231321203112010-0113030300132131-1323213200321221-2222210302120213-1221120123100301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0313030322231232-2030011003123221-3111133030113003-0311131002031330-2123113110210002-0100121123121212-0222223012000321-0322000222331112)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3122223131200201-1103022231023032-3220222100222030-2032030233023233-1233300210331120-2030212211003130-2003213120012111-0313311330013111)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters

<a id="canonical-1001233013133113-2030033130203122-2133013101011203-3322020333310012-2211031332013012-2100231322333320-0023133010230000-1332222303202321"></a>

Type: `"single"`. Computed.

Custom settings for query parameters validation.

<a id="canonical-3133232000003222-2112301300222110-2010111013301113-0030003103331110-0233213131113121-2300332101332300-2120130231332110-3222012300030213"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters`

- [allow_additional_parameters](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1122112121302313-3212100313322333-2222130132201001-0312330030031111-2013212001222202-1230003112332233-1031100222222031-1100120313323010): complete subsection reference.

- [disallow_additional_parameters](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1023303001112001-3112323002021021-3133132303230112-3200210000301311-3213212113330000-2303332100211220-0333122031211001-1321012101000131): complete subsection reference.

<a id="canonical-1122112121302313-3212100313322333-2222130132201001-0312330030031111-2013212001222202-1230003112332233-1031100222222031-1100120313323010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0313030322231232-2030011003123221-3111133030113003-0311131002031330-2123113110210002-0100121123121212-0222223012000321-0322000222331112)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3122223131200201-1103022231023032-3220222100222030-2032030233023233-1233300210331120-2030212211003130-2003213120012111-0313311330013111)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0012220213011221-1010012002013200-3331110210130300-3231321203112010-0113030300132131-1323213200321221-2222210302120213-1221120123100301)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters

<a id="canonical-3221113331301323-2033032311330201-2003311133333120-1020122310010311-0121102002333222-0032303233221032-1110123232321323-2013321101030002"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for allow additional parameters.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1023303001112001-3112323002021021-3133132303230112-3200210000301311-3213212113330000-2303332100211220-0333122031211001-1321012101000131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0313030322231232-2030011003123221-3111133030113003-0311131002031330-2123113110210002-0100121123121212-0222223012000321-0322000222331112)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3122223131200201-1103022231023032-3220222100222030-2032030233023233-1233300210331120-2030212211003130-2003213120012111-0313311330013111)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0012220213011221-1010012002013200-3331110210130300-3231321203112010-0113030300132131-1323213200321221-2222210302120213-1221120123100301)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters

<a id="canonical-3013012202221300-0002312202031113-0013232232310331-3110301010000133-3001013303233310-2230212230233212-2111122210113322-3312222132112203"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disallow additional parameters.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1231323101123001-1202002133333233-2213212021020013-3220133302210223-0011013301202230-1002310131033101-2030310111110301-1111211200312123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.settings.property_validation_settings_default` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [api_specification.validation_all_spec_endpoints.settings](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0313030322231232-2030011003123221-3111133030113003-0311131002031330-2123113110210002-0100121123121212-0222223012000321-0322000222331112)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_default

<a id="canonical-0001013311331122-0002303013011122-2020303022130220-2132010120011103-3232321023213230-3021101100012221-0123233022013313-0101030222303123"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for property validation settings default.

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

<a id="canonical-2110331212310200-0230113121313233-2221022221120013-2022103111323202-2110132022301012-2303032212120032-1203022022022212-1000231023221221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- api_specification.validation_all_spec_endpoints.validation_mode

<a id="canonical-1013232131312123-0130222201101031-2213010332020020-2302312021232231-2023103233011301-0330002123203132-0322002210320111-1113333330212032"></a>

Type: `"single"`. Computed.

Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of
the endpoints listed on the OpenAPI specification file (a.k.a. Swagger).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-response_validation_mode_choice": "[\"response_validation_mode_active\",\"skip_response_validation\"]",
  "x-ves-oneof-field-validation_mode_choice": "[\"skip_validation\",\"validation_mode_active\"]"
}
```

<a id="canonical-2303031203203302-1010200011320010-0322003200021103-3030312000133013-1200223022221310-1121221122331221-0223222013232122-1201322113020222"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.validation_mode`

- [response_validation_mode_active](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0322202301303102-0000302013121031-1122103012210103-1011103202201320-0112302233310011-3131022302132203-3322320222012333-1032310332232333): complete subsection reference.

- [skip_response_validation](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3110201230222330-2300300333023303-2112210211003132-1300303231123313-1110111002132202-3220222320110111-1102001023330132-0323121210300123): complete subsection reference.

- [skip_validation](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1301222223133203-1302011033211220-0222033212003301-0313121212211213-0132330131111002-3312310020212120-2301121021210202-2210010121200202): complete subsection reference.

- [validation_mode_active](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2100231133123132-2222330230220300-2313122202112311-2011023231121313-1111023321220203-1203323030303230-3110010231013301-2130111231020133): complete subsection reference.

<a id="canonical-0322202301303102-0000302013121031-1122103012210103-1011103202201320-0112302233310011-3131022302132203-3322320222012333-1032310332232333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2110331212310200-0230113121313233-2221022221120013-2022103111323202-2110132022301012-2303032212120032-1203022022022212-1000231023221221)
- api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active

<a id="canonical-0310023102122112-1123213230303312-2220100233102132-2021123212332022-2301000323213022-0203122023123323-0300002002211222-3321312220210101"></a>

Type: `"single"`. Computed.

Open API Validation Mode Active. Validation mode properties of response.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-validation_enforcement_type": "[\"enforcement_block\",\"enforcement_report\"]"
}
```

<a id="canonical-2312120232102233-0101320031231011-1332220003013310-1301122202203221-3023222311231030-2030123310301002-0013322223023323-1022330110113303"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active`

- [enforcement_block](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1322313223211001-3313221310331112-1210203033322023-1311131330101230-1310123303031200-1002233232131233-0023022221112012-1013222020303012): complete subsection reference.

- [enforcement_report](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1112222323011300-2123101110312022-2220333301303220-0233323100020002-0010322023202303-3031333303220110-2232101203002021-0001000020132223): complete subsection reference.

<a id="canonical-3321013131213030-1323310330220020-3020223003311102-3121003211201101-3002202320311010-3030033301031312-2122120311102233-1221310301222203"></a>

<a id="canonical-2223223223313023-3333230312230121-1011222232313012-2000213320333203-1031022132300303-2210131113001132-1320121101321102-1123020111023010"></a>

#### `api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.response_validation_properties` property

Type: `["list", "string"]`. Computed.

\[Enum:
PROPERTY\_QUERY\_PARAMETERS|PROPERTY\_PATH\_PARAMETERS|PROPERTY\_CONTENT\_TYPE|PROPERTY\_COOKIE\_PARAMETERS|PROPERTY\_HTTP\_HEADERS|PROPERTY\_HTTP\_BODY|PROPERTY\_SECURITY\_SCHEMA|PROPERTY\_RESPONSE\_CODE\]
List of properties of the response to validate according to the OpenAPI specification file (a.k.a.
Swagger). Possible values are \`PROPERTY\_QUERY\_PARAMETERS\`, \`PROPERTY\_PATH\_PARAMETERS\`,
\`PROPERTY\_CONTENT\_TYPE\`, \`PROPERTY\_COOKIE\_PARAMETERS\`, \`PROPERTY\_HTTP\_HEADERS\`,
\`PROPERTY\_HTTP\_BODY\`, \`PROPERTY\_SECURITY\_SCHEMA\`, \`PROPERTY\_RESPONSE\_CODE\`. Defaults to
\`PROPERTY\_QUERY\_PARAMETERS\`.

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[2,4,5,7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[2,4,5,7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1322313223211001-3313221310331112-1210203033322023-1311131330101230-1310123303031200-1002233232131233-0023022221112012-1013222020303012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_block` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2110331212310200-0230113121313233-2221022221120013-2022103111323202-2110132022301012-2303032212120032-1203022022022212-1000231023221221)
- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0322202301303102-0000302013121031-1122103012210103-1011103202201320-0112302233310011-3131022302132203-3322320222012333-1032310332232333)
- api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_block

<a id="canonical-0210310310112021-3320332012330221-1303301202122200-0131110111232011-3220103232000323-1322010311330020-0002001210223020-2321123113322211"></a>

Type: `["object", {}]`. Computed.

Blocking validation: reject traffic that violates the selected OpenAPI validation properties.
Invalid requests are returned as HTTP 403.

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

<a id="canonical-1112222323011300-2123101110312022-2220333301303220-0233323100020002-0010322023202303-3031333303220110-2232101203002021-0001000020132223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_report` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2110331212310200-0230113121313233-2221022221120013-2022103111323202-2110132022301012-2303032212120032-1203022022022212-1000231023221221)
- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0322202301303102-0000302013121031-1122103012210103-1011103202201320-0112302233310011-3131022302132203-3322320222012333-1032310332232333)
- api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_report

<a id="canonical-2322201131211202-1301000133210322-3230030332330303-3010322020230311-1012202313203110-2123010321123322-2322111323230111-3003301300323213"></a>

Type: `["object", {}]`. Computed.

Report-only validation: record OpenAPI violations while allowing the request or response to
continue.

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

<a id="canonical-3110201230222330-2300300333023303-2112210211003132-1300303231123313-1110111002132202-3220222320110111-1102001023330132-0323121210300123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode.skip_response_validation` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2110331212310200-0230113121313233-2221022221120013-2022103111323202-2110132022301012-2303032212120032-1203022022022212-1000231023221221)
- api_specification.validation_all_spec_endpoints.validation_mode.skip_response_validation

<a id="canonical-3010300131022200-1102010221213121-2333032113100101-2031122210031311-1032113322211220-0131122010232120-0102231230101010-1303222112122122"></a>

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

<a id="canonical-1301222223133203-1302011033211220-0222033212003301-0313121212211213-0132330131111002-3312310020212120-2301121021210202-2210010121200202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode.skip_validation` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2110331212310200-0230113121313233-2221022221120013-2022103111323202-2110132022301012-2303032212120032-1203022022022212-1000231023221221)
- api_specification.validation_all_spec_endpoints.validation_mode.skip_validation

<a id="canonical-1133323113022223-1033031320331033-3100010133203112-3311133210003323-3023011213202111-0233012112302222-1122332302311000-3331110021112220"></a>

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

<a id="canonical-2100231133123132-2222330230220300-2313122202112311-2011023231121313-1111023321220203-1203323030303230-3110010231013301-2130111231020133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2110331212310200-0230113121313233-2221022221120013-2022103111323202-2110132022301012-2303032212120032-1203022022022212-1000231023221221)
- api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active

<a id="canonical-1030231133302332-1202230323233032-3121110102123222-3103102123020200-2301301200230030-3111201222300101-1232110202203111-0132200221211312"></a>

Type: `"single"`. Computed.

Enable OpenAPI validation and explicitly select enforcement\_report to allow and log invalid
traffic, or enforcement\_block to reject invalid requests with HTTP 403.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-validation_enforcement_type": "[\"enforcement_block\",\"enforcement_report\"]"
}
```

<a id="canonical-0232332023221330-1112303101213223-0330131233111301-2121123033021220-2122103030230130-3012121012211110-3133312213123231-1102133312202310"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active`

- [enforcement_block](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0312320212133103-2212102103312231-2001231103001033-0033123301203110-2300322333311031-1132303202020100-3233010311131211-3322301101003112): complete subsection reference.

- [enforcement_report](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0313311012111011-3032013003230330-1212000122022311-0300000002032031-2030031302100122-0102032312303202-1213221222120201-1321010331013331): complete subsection reference.

<a id="canonical-0313010012302330-0133133111303220-2020331130210013-1211331003101121-0021232131020112-2311210230130233-1133021232221100-2311202333022310"></a>

<a id="canonical-2323211032102303-0130031101313111-3313010100200312-1213211211123032-2013301331002012-1013212231323002-2111031123010100-0133231123123331"></a>

#### `api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.request_validation_properties` property

Type: `["list", "string"]`. Computed.

\[Enum:
PROPERTY\_QUERY\_PARAMETERS|PROPERTY\_PATH\_PARAMETERS|PROPERTY\_CONTENT\_TYPE|PROPERTY\_COOKIE\_PARAMETERS|PROPERTY\_HTTP\_HEADERS|PROPERTY\_HTTP\_BODY|PROPERTY\_SECURITY\_SCHEMA|PROPERTY\_RESPONSE\_CODE\]
List of properties of the request to validate according to the OpenAPI specification file (a.k.a.
Swagger). Possible values are \`PROPERTY\_QUERY\_PARAMETERS\`, \`PROPERTY\_PATH\_PARAMETERS\`,
\`PROPERTY\_CONTENT\_TYPE\`, \`PROPERTY\_COOKIE\_PARAMETERS\`, \`PROPERTY\_HTTP\_HEADERS\`,
\`PROPERTY\_HTTP\_BODY\`, \`PROPERTY\_SECURITY\_SCHEMA\`, \`PROPERTY\_RESPONSE\_CODE\`. Defaults to
\`PROPERTY\_QUERY\_PARAMETERS\`.

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0312320212133103-2212102103312231-2001231103001033-0033123301203110-2300322333311031-1132303202020100-3233010311131211-3322301101003112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_block` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2110331212310200-0230113121313233-2221022221120013-2022103111323202-2110132022301012-2303032212120032-1203022022022212-1000231023221221)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2100231133123132-2222330230220300-2313122202112311-2011023231121313-1111023321220203-1203323030303230-3110010231013301-2130111231020133)
- api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_block

<a id="canonical-3113220201133332-2322312212330213-3330311012221231-3212113011023203-0022323333211332-1313311022113010-0101223010332101-2312112210201130"></a>

Type: `["object", {}]`. Computed.

Blocking validation: reject traffic that violates the selected OpenAPI validation properties.
Invalid requests are returned as HTTP 403.

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

<a id="canonical-0313311012111011-3032013003230330-1212000122022311-0300000002032031-2030031302100122-0102032312303202-1213221222120201-1321010331013331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_report` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_all_spec_endpoints](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-1230202233211323-1230321023102211-1330301320021211-3321021223232301-1111222221011203-2213012100321030-2332133001313131-1123122323121001)
- [api_specification.validation_all_spec_endpoints.validation_mode](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2110331212310200-0230113121313233-2221022221120013-2022103111323202-2110132022301012-2303032212120032-1203022022022212-1000231023221221)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2100231133123132-2222330230220300-2313122202112311-2011023231121313-1111023321220203-1203323030303230-3110010231013301-2130111231020133)
- api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_report

<a id="canonical-1223213023111300-3031033203122020-1022312121020020-0013320131010021-1311010020210212-0332313121130103-3103301101013123-0021010003212030"></a>

Type: `["object", {}]`. Computed.

Report-only validation: record OpenAPI violations while allowing the request or response to
continue.

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

<a id="canonical-3003011213100323-3200022012312312-3232210200301103-3300102033112212-0303110012333130-2220221033113122-2120120331301303-2122030120320000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- api_specification.validation_custom_list

<a id="canonical-2012020032302020-1302102120033112-3201133102002120-0312130123020302-1023310031021323-1232112330302212-1121211203120002-3132232220312230"></a>

Type: `"single"`. Computed.

Define API groups, base paths, or API endpoints and their OpenAPI validation modes. Any other
API-endpoint not listed will act according to 'Fall Through Mode'.

Additional upstream details:

Any other API-endpoint not listed will act according to "Fall Through Mode".

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-oversized_body_choice": "[]"
}
```

<a id="canonical-3121331301132002-1301000202111200-3302210010123130-2013102322223021-1223121033023213-1100320221313103-0130332320130331-1210302322113032"></a>

### Direct properties for `api_specification.validation_custom_list`

- [fall_through_mode](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0123133020003332-0303201321022202-2103000012220020-1123010020233030-1210313311213211-2331133333300023-2202123013033303-3121111020322233): complete subsection reference.

- [open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3031231222230321-1030300023112022-1013022221003000-1110012111020320-2023221202120210-1211111233202113-1322002311121130-1132231003302213): complete subsection reference.

- [settings](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3103033333323320-3313030222230102-1023130023223132-3213230031201320-3031111123223030-3002233222011101-0323033122113222-1032100312112003): complete subsection reference.

<a id="canonical-0123133020003332-0303201321022202-2103000012220020-1123010020233030-1210313311213211-2331133333300023-2202123013033303-3121111020322233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.fall_through_mode` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_custom_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3003011213100323-3200022012312312-3232210200301103-3300102033112212-0303110012333130-2220221033113122-2120120331301303-2122030120320000)
- api_specification.validation_custom_list.fall_through_mode

<a id="canonical-2033001011203312-2303001202222122-0223332001223123-1212310112102130-2013103301013200-1210020011332021-2112130033230211-3221203121120311"></a>

Type: `"single"`. Computed.

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-fall_through_mode_choice": "[\"fall_through_mode_allow\",\"fall_through_mode_custom\"]"
}
```

<a id="canonical-1132211032113302-0132232002210130-3122223301230010-0021120221031321-2223120131233211-1223303001120232-3003310102303321-0101200001220022"></a>

### Direct properties for `api_specification.validation_custom_list.fall_through_mode`

- [fall_through_mode_allow](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3023002300210330-0002020222312232-3301331330200333-1020121301020330-2002002200201110-1233320220103310-1021112031012111-3313322010002220): complete subsection reference.

- [fall_through_mode_custom](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0320310112013331-1103120232330023-0113321110212022-3223210132012000-2133033311123233-3022321012312132-3333101320112233-2223333321110111): complete subsection reference.

<a id="canonical-3023002300210330-0002020222312232-3301331330200333-1020121301020330-2002002200201110-1233320220103310-1021112031012111-3313322010002220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_allow` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_custom_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3003011213100323-3200022012312312-3232210200301103-3300102033112212-0303110012333130-2220221033113122-2120120331301303-2122030120320000)
- [api_specification.validation_custom_list.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0123133020003332-0303201321022202-2103000012220020-1123010020233030-1210313311213211-2331133333300023-2202123013033303-3121111020322233)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_allow

<a id="canonical-3013332133323303-1311200230223132-1222011101302123-0113010333012302-1202210332320232-1321031230301311-2310021013221221-2002103122023230"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for fall through mode allow.

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

<a id="canonical-0320310112013331-1103120232330023-0113321110212022-3223210132012000-2133033311123233-3022321012312132-3333101320112233-2223333321110111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_custom_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3003011213100323-3200022012312312-3232210200301103-3300102033112212-0303110012333130-2220221033113122-2120120331301303-2122030120320000)
- [api_specification.validation_custom_list.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0123133020003332-0303201321022202-2103000012220020-1123010020233030-1210313311213211-2331133333300023-2202123013033303-3121111020322233)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom

<a id="canonical-1121112333123003-0123302000132003-2103333020210332-3013310313323211-1010320210201201-2310100120320302-1300313200001222-1333221323303111"></a>

Type: `"single"`. Computed.

Configuration parameter for fall through mode custom.

Additional upstream details:

Define the fall through settings.

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

<a id="canonical-1000311221100023-3213211022211010-1210232031322311-0012221113010321-2232101131110212-3313220202332222-0123312201230000-0323021111320312"></a>

### Direct properties for `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom`

- [open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1013303303033001-3221102102102101-0332202031210121-0120230102102133-2000021212010312-3220110203030312-3001301020122122-1213013032313300): complete subsection reference.

<a id="canonical-1013303303033001-3221102102102101-0332202031210121-0120230102102133-2000021212010312-3220110203030312-3001301020122122-1213013032313300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_custom_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3003011213100323-3200022012312312-3232210200301103-3300102033112212-0303110012333130-2220221033113122-2120120331301303-2122030120320000)
- [api_specification.validation_custom_list.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0123133020003332-0303201321022202-2103000012220020-1123010020233030-1210313311213211-2331133333300023-2202123013033303-3121111020322233)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0320310112013331-1103120232330023-0113321110212022-3223210132012000-2133033311123233-3022321012312132-3333101320112233-2223333321110111)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules

<a id="canonical-1313221001300210-1312321323211313-3330320332322330-2320202202310023-2303323132311203-0130312212032203-2310312123222303-1300113313322020"></a>

Type: `"list"`. Computed.

Custom Fall Through Rule List. Rule or policy definition

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 15,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 15,
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
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-2011023012231220-0110233221133021-1311330011222220-0032322200020233-3121022013333200-3311032003201321-1131103333100221-1333130010220220"></a>

### Direct properties for `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules`

- [action_block](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2131120213330120-1220303323120111-0233013122122310-2222102011322213-3211023223012223-1022030003021120-3001013311232001-3021123010312203): complete subsection reference.

- [action_report](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0130032013110021-2033130000131302-1031113210010323-3311330232302020-0300330300020332-0002023323320030-3220131311230332-1301211213100332): complete subsection reference.

- [action_skip](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1320333032332321-1133202110230313-0020210333002130-0122133303031120-0033232000030032-1213212222112001-3111023101230111-0030111032202022): complete subsection reference.

- [api_endpoint](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0302233023033132-1210022301032300-3311020122103030-0232301102030321-3103322230021321-1112230221312000-1323030231123101-3020031322223031): complete subsection reference.

<a id="canonical-3031300322032221-2301323302300303-3200111210123123-3030112130331233-0023113010100000-2122322123023323-2212213310002200-3123301032202320"></a>

<a id="canonical-0003013230313220-0133023301230013-0100322100300001-2113320232213200-2022230331022020-1333121110113313-2211213023332232-1031200302211320"></a>

#### `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_group` property

Type: `"string"`. Computed.

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-1032001111320222-2001000232321310-3012230232130202-3032002001022333-1223331323112013-3203302003201321-0012303022021202-3030001221100032"></a>

<a id="canonical-0332112211100121-3322120332102322-0121131202120001-2031103123303130-3023102022201133-3132012323133312-2033001232112330-0223111220031020"></a>

#### `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.base_path` property

Type: `"string"`. Computed.

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [metadata](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2331221011003300-2020131102021023-2301212320210203-1101321231130021-1320221302021121-2031102303303001-1221011003111031-1110332003001111): complete subsection reference.

<a id="canonical-2131120213330120-1220303323120111-0233013122122310-2222102011322213-3211023223012223-1022030003021120-3001013311232001-3021123010312203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_custom_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3003011213100323-3200022012312312-3232210200301103-3300102033112212-0303110012333130-2220221033113122-2120120331301303-2122030120320000)
- [api_specification.validation_custom_list.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0123133020003332-0303201321022202-2103000012220020-1123010020233030-1210313311213211-2331133333300023-2202123013033303-3121111020322233)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0320310112013331-1103120232330023-0113321110212022-3223210132012000-2133033311123233-3022321012312132-3333101320112233-2223333321110111)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1013303303033001-3221102102102101-0332202031210121-0120230102102133-2000021212010312-3220110203030312-3001301020122122-1213013032313300)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block

<a id="canonical-0131213101212103-1031221000021312-1212110201222301-2113110201201312-3021222011232233-0302320131013122-1203113100322003-2201212023123322"></a>

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

<a id="canonical-0130032013110021-2033130000131302-1031113210010323-3311330232302020-0300330300020332-0002023323320030-3220131311230332-1301211213100332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_custom_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3003011213100323-3200022012312312-3232210200301103-3300102033112212-0303110012333130-2220221033113122-2120120331301303-2122030120320000)
- [api_specification.validation_custom_list.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0123133020003332-0303201321022202-2103000012220020-1123010020233030-1210313311213211-2331133333300023-2202123013033303-3121111020322233)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0320310112013331-1103120232330023-0113321110212022-3223210132012000-2133033311123233-3022321012312132-3333101320112233-2223333321110111)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1013303303033001-3221102102102101-0332202031210121-0120230102102133-2000021212010312-3220110203030312-3001301020122122-1213013032313300)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report

<a id="canonical-2312132333030303-2300231331322122-0103021221303301-1031312330222320-2103002112133302-3001033231312220-2003120100131133-2102220322213133"></a>

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

<a id="canonical-1320333032332321-1133202110230313-0020210333002130-0122133303031120-0033232000030032-1213212222112001-3111023101230111-0030111032202022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_custom_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3003011213100323-3200022012312312-3232210200301103-3300102033112212-0303110012333130-2220221033113122-2120120331301303-2122030120320000)
- [api_specification.validation_custom_list.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0123133020003332-0303201321022202-2103000012220020-1123010020233030-1210313311213211-2331133333300023-2202123013033303-3121111020322233)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0320310112013331-1103120232330023-0113321110212022-3223210132012000-2133033311123233-3022321012312132-3333101320112233-2223333321110111)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1013303303033001-3221102102102101-0332202031210121-0120230102102133-2000021212010312-3220110203030312-3001301020122122-1213013032313300)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip

<a id="canonical-0221113003122333-1002012102121000-0333311021323020-3000302202101331-2222013103302120-2002102002213131-0010123103013123-0023102323033010"></a>

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

<a id="canonical-0302233023033132-1210022301032300-3311020122103030-0232301102030321-3103322230021321-1112230221312000-1323030231123101-3020031322223031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_custom_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3003011213100323-3200022012312312-3232210200301103-3300102033112212-0303110012333130-2220221033113122-2120120331301303-2122030120320000)
- [api_specification.validation_custom_list.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0123133020003332-0303201321022202-2103000012220020-1123010020233030-1210313311213211-2331133333300023-2202123013033303-3121111020322233)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0320310112013331-1103120232330023-0113321110212022-3223210132012000-2133033311123233-3022321012312132-3333101320112233-2223333321110111)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1013303303033001-3221102102102101-0332202031210121-0120230102102133-2000021212010312-3220110203030312-3001301020122122-1213013032313300)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint

<a id="canonical-3001233002133033-1010020303012233-1112312300100120-3030003103100202-3232132100112311-3032011021003002-1033301103220231-2030001032202212"></a>

Type: `"single"`. Computed.

API Endpoint. This defines API endpoint.

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

<a id="canonical-3133311302012311-1121332330222201-3001111302202313-0223301301202132-1320202333120030-3100301003321121-3100233003021323-2221200220310112"></a>

### Direct properties for `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint`

<a id="canonical-0300333131203231-2030302300001131-2012220233322210-0333332303303313-2130010001110321-0300003032311200-1013001033010101-2210023201030220"></a>

#### `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint.methods` property

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1211000102320310-2010212030020033-3110000003312200-0223011120023122-0333111333131123-0231122002021130-1221222132311203-3103122200213203"></a>

<a id="canonical-2310300032331133-1013101113332032-0020130321031331-3212333301032231-2213101320221102-3010110212212232-3013103202113313-2232201121211112"></a>

#### `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint.path` property

Type: `"string"`. Computed.

Path. Path to be matched.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

<a id="canonical-2331221011003300-2020131102021023-2301212320210203-1101321231130021-1320221302021121-2031102303303001-1221011003111031-1110332003001111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_custom_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3003011213100323-3200022012312312-3232210200301103-3300102033112212-0303110012333130-2220221033113122-2120120331301303-2122030120320000)
- [api_specification.validation_custom_list.fall_through_mode](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0123133020003332-0303201321022202-2103000012220020-1123010020233030-1210313311213211-2331133333300023-2202123013033303-3121111020322233)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0320310112013331-1103120232330023-0113321110212022-3223210132012000-2133033311123233-3022321012312132-3333101320112233-2223333321110111)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1013303303033001-3221102102102101-0332202031210121-0120230102102133-2000021212010312-3220110203030312-3001301020122122-1213013032313300)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata

<a id="canonical-0012312031212322-2202222203323200-0101011131210131-3101202132032103-2112320031023022-2031101120301113-0002312100020321-1222322001202312"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-1211112300200230-0223010230320333-0220020201302203-3030120020010113-2023130323321101-3012230031120032-1023231001131311-1330133230001333"></a>

### Direct properties for `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata`

<a id="canonical-2303002231030201-0032133011213000-3320233212022200-2320220322102031-3120100132020200-1211323233223220-2123033223122233-1233013132213313"></a>

#### `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-2012200120101212-2300102331320022-1321000312201101-1020133020213211-0211302310322110-1033222002020131-3211200132331331-3213220000120132"></a>

<a id="canonical-3222113313030013-2103011112032120-3123133002112001-1011220132320102-0011101120011032-2232130021003102-1012212322100021-1133220112322131"></a>

#### `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata.name` property

Type: `"string"`. Computed.

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-3031231222230321-1030300023112022-1013022221003000-1110012111020320-2023221202120210-1211111233202113-1322002311121130-1132231003302213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_custom_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3003011213100323-3200022012312312-3232210200301103-3300102033112212-0303110012333130-2220221033113122-2120120331301303-2122030120320000)
- api_specification.validation_custom_list.open_api_validation_rules

<a id="canonical-2213111333231333-2223202201011212-0312333203013031-3311231222101131-0022320133023331-3120133102121010-0202033021000130-1103020202001112"></a>

Type: `"list"`. Computed.

Validation List. Rule or policy definition

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 15,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 15,
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
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-1000223213031131-1202303012031232-0123211221102011-0131110330330133-1301120222313032-3101033310113202-1322202200031010-3210303313010311"></a>

### Direct properties for `api_specification.validation_custom_list.open_api_validation_rules`

- [any_domain](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2200200220311022-3312231000122001-1102230200222303-2001220010202013-2132332301123302-3201223103201112-1013232112210013-1121302123231103): complete subsection reference.

- [api_endpoint](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3001202203330110-2023332331212003-0221113301312110-0231032102230133-3313121112033333-2230011333313333-3022213212023021-1330301330120210): complete subsection reference.

<a id="canonical-1111131300223301-0331310330120032-3310031221200212-3322201323101022-3211011233100321-0211000131300213-3123212313103220-1332322213300331"></a>

<a id="canonical-2321331001201002-3230212211313003-3311121201111222-1021312120333121-3301100033001112-1303030320031223-2012113112110133-3110301110021211"></a>

#### `api_specification.validation_custom_list.open_api_validation_rules.api_group` property

Type: `"string"`. Computed.

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-2011222330133200-1323203301121012-2112312300203001-0330013103011113-2202213332122321-1020002003020122-3310012120202000-0200133022000302"></a>

<a id="canonical-0031013130230231-0011201023122031-1110210230333232-1303203123103232-2300301100012110-1320330010101031-1032320333231331-1313213101222223"></a>

#### `api_specification.validation_custom_list.open_api_validation_rules.base_path` property

Type: `"string"`. Computed.

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [metadata](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2333221022102122-1203011221020313-3313020301200010-3021330313232301-3303131321321033-3120303103103300-0012023033320132-3123110203031332): complete subsection reference.

<a id="canonical-0130200123303230-0211233212321303-0201212213012301-1300303313232120-3230031133130131-2302131332223220-3022033301012311-3321012312133322"></a>

<a id="canonical-1320232301131320-2102031210031112-3003131023131013-2300213110303133-2131211110103301-0020330000302211-1103121333130122-2112022022223011"></a>

#### `api_specification.validation_custom_list.open_api_validation_rules.specific_domain` property

Type: `"string"`. Computed.

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

- [validation_mode](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2313121301020322-2003100123003311-0323231001102302-3310000320220003-1301331320132033-1120133030013123-0212312102013020-2102303112331002): complete subsection reference.

<a id="canonical-2200200220311022-3312231000122001-1102230200222303-2001220010202013-2132332301123302-3201223103201112-1013232112210013-1121302123231103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.any_domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_custom_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3003011213100323-3200022012312312-3232210200301103-3300102033112212-0303110012333130-2220221033113122-2120120331301303-2122030120320000)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3031231222230321-1030300023112022-1013022221003000-1110012111020320-2023221202120210-1211111233202113-1322002311121130-1132231003302213)
- api_specification.validation_custom_list.open_api_validation_rules.any_domain

<a id="canonical-2201320321300003-3021122300112112-1223300300100212-3021022132020031-1130032000003000-3300023121300102-0203332321121211-0200333213330202"></a>

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

<a id="canonical-3001202203330110-2023332331212003-0221113301312110-0231032102230133-3313121112033333-2230011333313333-3022213212023021-1330301330120210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.api_endpoint` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_custom_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3003011213100323-3200022012312312-3232210200301103-3300102033112212-0303110012333130-2220221033113122-2120120331301303-2122030120320000)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3031231222230321-1030300023112022-1013022221003000-1110012111020320-2023221202120210-1211111233202113-1322002311121130-1132231003302213)
- api_specification.validation_custom_list.open_api_validation_rules.api_endpoint

<a id="canonical-1313132322232120-3232031330302300-0030303302212120-2100233022303301-3213103323113023-0210011223102021-2131312233022110-2200002121111213"></a>

Type: `"single"`. Computed.

API Endpoint. This defines API endpoint.

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

<a id="canonical-3103100000122132-2300231202003221-1002321023123230-2113130201301202-2223000313023231-2222221233301302-1111012030001213-0011113210023101"></a>

### Direct properties for `api_specification.validation_custom_list.open_api_validation_rules.api_endpoint`

<a id="canonical-1132120001333213-1021122031203031-1130122103320011-3003113222030103-1313231211213122-1131033200230313-0011022212212103-1120133020100300"></a>

#### `api_specification.validation_custom_list.open_api_validation_rules.api_endpoint.methods` property

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2130113102233322-1301121201212110-1123220022231032-0023010233101003-0322020000210003-0132031131003223-1321130032103112-2301133133333101"></a>

<a id="canonical-1212111133202302-1221230301133122-3321112232211120-0303223303022003-0130220001301100-3312303211200013-0220000012112100-1023013221101301"></a>

#### `api_specification.validation_custom_list.open_api_validation_rules.api_endpoint.path` property

Type: `"string"`. Computed.

Path. Path to be matched.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

<a id="canonical-2333221022102122-1203011221020313-3313020301200010-3021330313232301-3303131321321033-3120303103103300-0012023033320132-3123110203031332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.metadata` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_custom_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3003011213100323-3200022012312312-3232210200301103-3300102033112212-0303110012333130-2220221033113122-2120120331301303-2122030120320000)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3031231222230321-1030300023112022-1013022221003000-1110012111020320-2023221202120210-1211111233202113-1322002311121130-1132231003302213)
- api_specification.validation_custom_list.open_api_validation_rules.metadata

<a id="canonical-0112300312321200-0131030333102223-0221232331112223-1231210013231212-1112322200032302-3311302233130130-1002100122031023-2232021010101121"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-2231223130203011-0011231030223012-1102011331200303-0032333012231103-2033211302112010-1120230231212232-0003200223023102-1031021230330100"></a>

### Direct properties for `api_specification.validation_custom_list.open_api_validation_rules.metadata`

<a id="canonical-2013120211122312-2220130130133310-0222020323112200-1212003031113212-3320132030201132-0220202203310103-0003122111122323-0210021221100213"></a>

#### `api_specification.validation_custom_list.open_api_validation_rules.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-2221200122210221-3023003001300233-2212012023200320-2111300212312111-2313030032013222-0313020320031201-1101112232310320-0231323132100023"></a>

<a id="canonical-1200103231103301-3003201313222022-3223100032102200-1131220233312132-3321120020022101-2303013031322023-1101330001031300-2032011322231030"></a>

#### `api_specification.validation_custom_list.open_api_validation_rules.metadata.name` property

Type: `"string"`. Computed.

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-2313121301020322-2003100123003311-0323231001102302-3310000320220003-1301331320132033-1120133030013123-0212312102013020-2102303112331002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.validation_mode` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_custom_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3003011213100323-3200022012312312-3232210200301103-3300102033112212-0303110012333130-2220221033113122-2120120331301303-2122030120320000)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3031231222230321-1030300023112022-1013022221003000-1110012111020320-2023221202120210-1211111233202113-1322002311121130-1132231003302213)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode

<a id="canonical-2132332103301003-3333013012003321-2101322230000112-0210013131231203-2313022310331102-0322323220010113-3032010023012120-3100231131302333"></a>

Type: `"single"`. Computed.

Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of
the endpoints listed on the OpenAPI specification file (a.k.a. Swagger).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-response_validation_mode_choice": "[\"response_validation_mode_active\",\"skip_response_validation\"]",
  "x-ves-oneof-field-validation_mode_choice": "[\"skip_validation\",\"validation_mode_active\"]"
}
```

<a id="canonical-2232110131000312-3130112022011013-1202110130122011-0322101121323200-1200101203302223-1223213202213023-2201133231231032-1220012031211330"></a>

### Direct properties for `api_specification.validation_custom_list.open_api_validation_rules.validation_mode`

- [response_validation_mode_active](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2130331200123211-2120020003012220-2232212102003002-1122311212222101-2223102330003031-2301133233312102-2223320330111022-3321021220233300): complete subsection reference.

- [skip_response_validation](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1002212131320103-3333222120000331-3332021212012322-1011223113130022-0202123033130233-2230122130110000-2321332233030201-2032020102223002): complete subsection reference.

- [skip_validation](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3202113210131012-0011333033300302-3213330003012330-3013130021303011-3113011221322033-0312322020200113-2110333211102203-2310012212302001): complete subsection reference.

- [validation_mode_active](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1122211232222212-2012320223310302-3300321010022213-3321323231023200-2311003202232320-2132212313221301-0223300333203332-0101201310002212): complete subsection reference.

<a id="canonical-2130331200123211-2120020003012220-2232212102003002-1122311212222101-2223102330003031-2301133233312102-2223320330111022-3321021220233300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_custom_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3003011213100323-3200022012312312-3232210200301103-3300102033112212-0303110012333130-2220221033113122-2120120331301303-2122030120320000)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3031231222230321-1030300023112022-1013022221003000-1110012111020320-2023221202120210-1211111233202113-1322002311121130-1132231003302213)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2313121301020322-2003100123003311-0323231001102302-3310000320220003-1301331320132033-1120133030013123-0212312102013020-2102303112331002)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active

<a id="canonical-3111231031032130-1113201020111221-3302021221021322-2101312121102313-3211312310130323-0302300301033011-1201301221133222-0222303322302122"></a>

Type: `"single"`. Computed.

Open API Validation Mode Active. Validation mode properties of response.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-validation_enforcement_type": "[\"enforcement_block\",\"enforcement_report\"]"
}
```

<a id="canonical-2101101011200003-2130130113003331-3103021020112200-2203130231130100-3131103132233302-0101101011003011-3010311011202301-1210203310000033"></a>

### Direct properties for `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active`

- [enforcement_block](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1211130000313123-3200022223110123-3332303231103323-2210003110111332-3000021001131222-3210310331233110-3333110033302200-2212320303211110): complete subsection reference.

- [enforcement_report](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3212020322130001-2310111230032233-1123021111122323-1331033212010201-2112103213333021-2313321102331011-0232010120213023-0210033012031221): complete subsection reference.

<a id="canonical-2220211312100223-0231222200100221-2013212010111010-1320201020110022-0212331132023122-3101003303121300-3313112030103312-1121230202021102"></a>

<a id="canonical-1232133203312021-3323001331201021-3233320001201300-0233311201321233-3122022312013010-2102101110032310-0232212030133020-3320130030130132"></a>

#### `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.response_validation_properties` property

Type: `["list", "string"]`. Computed.

\[Enum:
PROPERTY\_QUERY\_PARAMETERS|PROPERTY\_PATH\_PARAMETERS|PROPERTY\_CONTENT\_TYPE|PROPERTY\_COOKIE\_PARAMETERS|PROPERTY\_HTTP\_HEADERS|PROPERTY\_HTTP\_BODY|PROPERTY\_SECURITY\_SCHEMA|PROPERTY\_RESPONSE\_CODE\]
List of properties of the response to validate according to the OpenAPI specification file (a.k.a.
Swagger). Possible values are \`PROPERTY\_QUERY\_PARAMETERS\`, \`PROPERTY\_PATH\_PARAMETERS\`,
\`PROPERTY\_CONTENT\_TYPE\`, \`PROPERTY\_COOKIE\_PARAMETERS\`, \`PROPERTY\_HTTP\_HEADERS\`,
\`PROPERTY\_HTTP\_BODY\`, \`PROPERTY\_SECURITY\_SCHEMA\`, \`PROPERTY\_RESPONSE\_CODE\`. Defaults to
\`PROPERTY\_QUERY\_PARAMETERS\`.

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[2,4,5,7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[2,4,5,7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1211130000313123-3200022223110123-3332303231103323-2210003110111332-3000021001131222-3210310331233110-3333110033302200-2212320303211110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_block` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_custom_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3003011213100323-3200022012312312-3232210200301103-3300102033112212-0303110012333130-2220221033113122-2120120331301303-2122030120320000)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3031231222230321-1030300023112022-1013022221003000-1110012111020320-2023221202120210-1211111233202113-1322002311121130-1132231003302213)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2313121301020322-2003100123003311-0323231001102302-3310000320220003-1301331320132033-1120133030013123-0212312102013020-2102303112331002)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2130331200123211-2120020003012220-2232212102003002-1122311212222101-2223102330003031-2301133233312102-2223320330111022-3321021220233300)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_block

<a id="canonical-1002213322102320-3202020031003010-1132212302013102-3012220031012113-0331301321130132-2120101201021103-0213032113332113-1222122012311211"></a>

Type: `["object", {}]`. Computed.

Blocking validation: reject traffic that violates the selected OpenAPI validation properties.
Invalid requests are returned as HTTP 403.

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

<a id="canonical-3212020322130001-2310111230032233-1123021111122323-1331033212010201-2112103213333021-2313321102331011-0232010120213023-0210033012031221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_report` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_custom_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3003011213100323-3200022012312312-3232210200301103-3300102033112212-0303110012333130-2220221033113122-2120120331301303-2122030120320000)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3031231222230321-1030300023112022-1013022221003000-1110012111020320-2023221202120210-1211111233202113-1322002311121130-1132231003302213)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2313121301020322-2003100123003311-0323231001102302-3310000320220003-1301331320132033-1120133030013123-0212312102013020-2102303112331002)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2130331200123211-2120020003012220-2232212102003002-1122311212222101-2223102330003031-2301133233312102-2223320330111022-3321021220233300)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_report

<a id="canonical-2101030020000303-0122130003231210-1222331023330121-1220033320132101-3320203202210013-1300300331103032-1022110100113030-2112210031202011"></a>

Type: `["object", {}]`. Computed.

Report-only validation: record OpenAPI violations while allowing the request or response to
continue.

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

<a id="canonical-1002212131320103-3333222120000331-3332021212012322-1011223113130022-0202123033130233-2230122130110000-2321332233030201-2032020102223002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_response_validation` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_custom_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3003011213100323-3200022012312312-3232210200301103-3300102033112212-0303110012333130-2220221033113122-2120120331301303-2122030120320000)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3031231222230321-1030300023112022-1013022221003000-1110012111020320-2023221202120210-1211111233202113-1322002311121130-1132231003302213)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2313121301020322-2003100123003311-0323231001102302-3310000320220003-1301331320132033-1120133030013123-0212312102013020-2102303112331002)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_response_validation

<a id="canonical-3303121200113030-2033210212023223-2231220313023132-3323221001012111-2300203011010102-1123123233010303-3113323233223020-3030011133211102"></a>

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

<a id="canonical-3202113210131012-0011333033300302-3213330003012330-3013130021303011-3113011221322033-0312322020200113-2110333211102203-2310012212302001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_validation` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_custom_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3003011213100323-3200022012312312-3232210200301103-3300102033112212-0303110012333130-2220221033113122-2120120331301303-2122030120320000)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3031231222230321-1030300023112022-1013022221003000-1110012111020320-2023221202120210-1211111233202113-1322002311121130-1132231003302213)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2313121301020322-2003100123003311-0323231001102302-3310000320220003-1301331320132033-1120133030013123-0212312102013020-2102303112331002)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_validation

<a id="canonical-2203302001122203-0201312201223002-0220011300202010-1213210002330311-2130023202122202-1103103230211021-1000221322100312-2323121011303331"></a>

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

<a id="canonical-1122211232222212-2012320223310302-3300321010022213-3321323231023200-2311003202232320-2132212313221301-0223300333203332-0101201310002212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_custom_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3003011213100323-3200022012312312-3232210200301103-3300102033112212-0303110012333130-2220221033113122-2120120331301303-2122030120320000)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3031231222230321-1030300023112022-1013022221003000-1110012111020320-2023221202120210-1211111233202113-1322002311121130-1132231003302213)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2313121301020322-2003100123003311-0323231001102302-3310000320220003-1301331320132033-1120133030013123-0212312102013020-2102303112331002)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active

<a id="canonical-0331311021222203-3221102011102310-1322001020210020-3303330213232100-1103123102333030-0101213131111133-2102030302010203-3213020311121200"></a>

Type: `"single"`. Computed.

Enable OpenAPI validation and explicitly select enforcement\_report to allow and log invalid
traffic, or enforcement\_block to reject invalid requests with HTTP 403.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-validation_enforcement_type": "[\"enforcement_block\",\"enforcement_report\"]"
}
```

<a id="canonical-0123211201011022-3212302223133021-2331122103023012-1113113310232232-2230020021210003-3112323102333303-1102020310333032-2313020221300310"></a>

### Direct properties for `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active`

- [enforcement_block](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1031323313311320-1030231230302311-3211310011210200-1132020113320010-0313321213223231-3021022311203031-2301031320121030-3211100033200100): complete subsection reference.

- [enforcement_report](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1031310310312011-1331002113203012-0021231023122031-3303301301313220-0002303032000033-0003102313322233-0331012100103331-3213313301313212): complete subsection reference.

<a id="canonical-0121222022330131-2000333123021133-2222000233131333-2003033112130303-2103032210122001-1103330323333130-3202103110222030-0212113123332222"></a>

<a id="canonical-1331312012231112-2113002320202012-3010333113123213-0110300100221023-0100322332023001-1213331003111013-1232301110322232-1130002323201212"></a>

#### `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.request_validation_properties` property

Type: `["list", "string"]`. Computed.

\[Enum:
PROPERTY\_QUERY\_PARAMETERS|PROPERTY\_PATH\_PARAMETERS|PROPERTY\_CONTENT\_TYPE|PROPERTY\_COOKIE\_PARAMETERS|PROPERTY\_HTTP\_HEADERS|PROPERTY\_HTTP\_BODY|PROPERTY\_SECURITY\_SCHEMA|PROPERTY\_RESPONSE\_CODE\]
List of properties of the request to validate according to the OpenAPI specification file (a.k.a.
Swagger). Possible values are \`PROPERTY\_QUERY\_PARAMETERS\`, \`PROPERTY\_PATH\_PARAMETERS\`,
\`PROPERTY\_CONTENT\_TYPE\`, \`PROPERTY\_COOKIE\_PARAMETERS\`, \`PROPERTY\_HTTP\_HEADERS\`,
\`PROPERTY\_HTTP\_BODY\`, \`PROPERTY\_SECURITY\_SCHEMA\`, \`PROPERTY\_RESPONSE\_CODE\`. Defaults to
\`PROPERTY\_QUERY\_PARAMETERS\`.

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1031323313311320-1030231230302311-3211310011210200-1132020113320010-0313321213223231-3021022311203031-2301031320121030-3211100033200100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_block` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_custom_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3003011213100323-3200022012312312-3232210200301103-3300102033112212-0303110012333130-2220221033113122-2120120331301303-2122030120320000)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3031231222230321-1030300023112022-1013022221003000-1110012111020320-2023221202120210-1211111233202113-1322002311121130-1132231003302213)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2313121301020322-2003100123003311-0323231001102302-3310000320220003-1301331320132033-1120133030013123-0212312102013020-2102303112331002)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1122211232222212-2012320223310302-3300321010022213-3321323231023200-2311003202232320-2132212313221301-0223300333203332-0101201310002212)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_block

<a id="canonical-3310202032120232-2101311322202230-0310022323213130-2231133311311333-1201303210031233-0102033130120300-3100330200113330-0131322132303131"></a>

Type: `["object", {}]`. Computed.

Blocking validation: reject traffic that violates the selected OpenAPI validation properties.
Invalid requests are returned as HTTP 403.

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

<a id="canonical-1031310310312011-1331002113203012-0021231023122031-3303301301313220-0002303032000033-0003102313322233-0331012100103331-3213313301313212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_report` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_custom_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3003011213100323-3200022012312312-3232210200301103-3300102033112212-0303110012333130-2220221033113122-2120120331301303-2122030120320000)
- [api_specification.validation_custom_list.open_api_validation_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3031231222230321-1030300023112022-1013022221003000-1110012111020320-2023221202120210-1211111233202113-1322002311121130-1132231003302213)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2313121301020322-2003100123003311-0323231001102302-3310000320220003-1301331320132033-1120133030013123-0212312102013020-2102303112331002)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1122211232222212-2012320223310302-3300321010022213-3321323231023200-2311003202232320-2132212313221301-0223300333203332-0101201310002212)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_report

<a id="canonical-1120302023212311-0011322221033023-0232002231311220-1311022133123232-0010112002220231-2110211130302333-1000230103013333-0000100002230321"></a>

Type: `["object", {}]`. Computed.

Report-only validation: record OpenAPI violations while allowing the request or response to
continue.

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

<a id="canonical-3103033333323320-3313030222230102-1023130023223132-3213230031201320-3031111123223030-3002233222011101-0323033122113222-1032100312112003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.settings` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_custom_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3003011213100323-3200022012312312-3232210200301103-3300102033112212-0303110012333130-2220221033113122-2120120331301303-2122030120320000)
- api_specification.validation_custom_list.settings

<a id="canonical-1301301030313021-2310022332330230-1223300320223203-2013312233002100-3210310303213112-0030212300223202-2023023010123303-0121133100223023"></a>

Type: `"single"`. Computed.

OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom
list' enforcement.

Additional upstream details:

OpenAPI specification validation settings relevant for "API Inventory" enforcement and for "Custom
list" enforcement.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-fail_configuration": "[]",
  "x-ves-oneof-field-oversized_body_choice": "[\"oversized_body_fail_validation\",\"oversized_body_skip_validation\"]",
  "x-ves-oneof-field-property_validation_settings_choice": "[\"property_validation_settings_custom\",\"property_validation_settings_default\"]"
}
```

<a id="canonical-2203031011210121-3231023000112330-2201112312331323-2110010033110301-1120320330300302-3023333010003011-2320320310220133-1332230310112210"></a>

### Direct properties for `api_specification.validation_custom_list.settings`

- [oversized_body_fail_validation](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2332031101300023-0232120002302031-3031332223220023-0323211220202323-1010201112000123-1011010113220321-3021232333000031-0300222121312022): complete subsection reference.

- [oversized_body_skip_validation](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1330133313020213-2111330223121333-3032031111121020-2230322033132102-0222002213013131-3213202331032002-2223320123323112-0333210130333313): complete subsection reference.

- [property_validation_settings_custom](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2201103103321303-1110332012011022-1322311310322203-0103201322210120-0110301003313200-0030310003002012-0203203113322212-3233031010220113): complete subsection reference.

- [property_validation_settings_default](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1332010003310000-0021333333013233-2211330123200333-2001200002303111-0221210032220010-2112031130110021-1203131223211223-1322310122020000): complete subsection reference.

<a id="canonical-2332031101300023-0232120002302031-3031332223220023-0323211220202323-1010201112000123-1011010113220321-3021232333000031-0300222121312022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.settings.oversized_body_fail_validation` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_custom_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3003011213100323-3200022012312312-3232210200301103-3300102033112212-0303110012333130-2220221033113122-2120120331301303-2122030120320000)
- [api_specification.validation_custom_list.settings](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3103033333323320-3313030222230102-1023130023223132-3213230031201320-3031111123223030-3002233222011101-0323033122113222-1032100312112003)
- api_specification.validation_custom_list.settings.oversized_body_fail_validation

<a id="canonical-1321002302320032-1322013231200023-0201110030321120-1001002212020301-0222130311303230-3212112330232232-1132233012211011-3120021023023130"></a>

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

<a id="canonical-1330133313020213-2111330223121333-3032031111121020-2230322033132102-0222002213013131-3213202331032002-2223320123323112-0333210130333313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.settings.oversized_body_skip_validation` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_custom_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3003011213100323-3200022012312312-3232210200301103-3300102033112212-0303110012333130-2220221033113122-2120120331301303-2122030120320000)
- [api_specification.validation_custom_list.settings](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3103033333323320-3313030222230102-1023130023223132-3213230031201320-3031111123223030-3002233222011101-0323033122113222-1032100312112003)
- api_specification.validation_custom_list.settings.oversized_body_skip_validation

<a id="canonical-0122330210203010-0010223110302302-0131223002001120-1030222132330312-1313301122102000-2100123302130113-3200331121123121-3112233110023321"></a>

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

<a id="canonical-2201103103321303-1110332012011022-1322311310322203-0103201322210120-0110301003313200-0030310003002012-0203203113322212-3233031010220113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.settings.property_validation_settings_custom` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_custom_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3003011213100323-3200022012312312-3232210200301103-3300102033112212-0303110012333130-2220221033113122-2120120331301303-2122030120320000)
- [api_specification.validation_custom_list.settings](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3103033333323320-3313030222230102-1023130023223132-3213230031201320-3031111123223030-3002233222011101-0323033122113222-1032100312112003)
- api_specification.validation_custom_list.settings.property_validation_settings_custom

<a id="canonical-2120223330203112-2203012201331001-1122110013102100-3131123113300312-3002330133120130-0211132333201023-1312111112201111-0312123101231011"></a>

Type: `"single"`. Computed.

Configuration parameter for property validation settings custom.

Additional upstream details:

Custom property validation settings.

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

<a id="canonical-0011011230322033-1232032310013012-1230210322021310-3233301112220012-2010223120300332-0112021123131111-3202332203321311-1321320001203201"></a>

### Direct properties for `api_specification.validation_custom_list.settings.property_validation_settings_custom`

- [query_parameters](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3122031132003002-0112311312030231-0330300020030230-0333231313121011-0322300233223101-3132011020310010-1303113300203223-1320120200202232): complete subsection reference.

<a id="canonical-3122031132003002-0112311312030231-0330300020030230-0333231313121011-0322300233223101-3132011020310010-1303113300203223-1320120200202232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_custom_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3003011213100323-3200022012312312-3232210200301103-3300102033112212-0303110012333130-2220221033113122-2120120331301303-2122030120320000)
- [api_specification.validation_custom_list.settings](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3103033333323320-3313030222230102-1023130023223132-3213230031201320-3031111123223030-3002233222011101-0323033122113222-1032100312112003)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2201103103321303-1110332012011022-1322311310322203-0103201322210120-0110301003313200-0030310003002012-0203203113322212-3233031010220113)
- api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters

<a id="canonical-0102222220203000-2300022220022111-1211231113331132-2032101101001232-3333230221332033-1313103120222302-2310103023211231-0123202210112211"></a>

Type: `"single"`. Computed.

Custom settings for query parameters validation.

<a id="canonical-3222212120103200-2112013322303023-3122131210103013-3020120211101202-2313001300223123-2322100230130012-0100222333102312-1133010122301003"></a>

### Direct properties for `api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters`

- [allow_additional_parameters](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0012231103000333-0100321000310233-1203330122133013-2332103032232220-1001321123200231-2023033301010010-3023012313212313-0331133211003200): complete subsection reference.

- [disallow_additional_parameters](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2232123212321030-3110111103331220-1321330203013210-0133321100130221-3033323232321201-1021023103321123-2202312013311201-3101133230132231): complete subsection reference.

<a id="canonical-0012231103000333-0100321000310233-1203330122133013-2332103032232220-1001321123200231-2023033301010010-3023012313212313-0331133211003200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_custom_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3003011213100323-3200022012312312-3232210200301103-3300102033112212-0303110012333130-2220221033113122-2120120331301303-2122030120320000)
- [api_specification.validation_custom_list.settings](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3103033333323320-3313030222230102-1023130023223132-3213230031201320-3031111123223030-3002233222011101-0323033122113222-1032100312112003)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2201103103321303-1110332012011022-1322311310322203-0103201322210120-0110301003313200-0030310003002012-0203203113322212-3233031010220113)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3122031132003002-0112311312030231-0330300020030230-0333231313121011-0322300233223101-3132011020310010-1303113300203223-1320120200202232)
- api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters

<a id="canonical-3100120012121021-0312032212232200-2230121010030122-0100012012322100-1110330123011012-0132210011302233-3012303332213101-0312203322122323"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for allow additional parameters.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2232123212321030-3110111103331220-1321330203013210-0133321100130221-3033323232321201-1021023103321123-2202312013311201-3101133230132231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_custom_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3003011213100323-3200022012312312-3232210200301103-3300102033112212-0303110012333130-2220221033113122-2120120331301303-2122030120320000)
- [api_specification.validation_custom_list.settings](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3103033333323320-3313030222230102-1023130023223132-3213230031201320-3031111123223030-3002233222011101-0323033122113222-1032100312112003)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2201103103321303-1110332012011022-1322311310322203-0103201322210120-0110301003313200-0030310003002012-0203203113322212-3233031010220113)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3122031132003002-0112311312030231-0330300020030230-0333231313121011-0322300233223101-3132011020310010-1303113300203223-1320120200202232)
- api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters

<a id="canonical-2321313303132220-0322111213102213-3331132332321121-1223023001223223-0232122330311232-2001233201123300-1003030301130033-2120302203111321"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disallow additional parameters.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1332010003310000-0021333333013233-2211330123200333-2001200002303111-0221210032220010-2112031130110021-1203131223211223-1322310122020000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.settings.property_validation_settings_default` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- [api_specification.validation_custom_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3003011213100323-3200022012312312-3232210200301103-3300102033112212-0303110012333130-2220221033113122-2120120331301303-2122030120320000)
- [api_specification.validation_custom_list.settings](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3103033333323320-3313030222230102-1023130023223132-3213230031201320-3031111123223030-3002233222011101-0323033122113222-1032100312112003)
- api_specification.validation_custom_list.settings.property_validation_settings_default

<a id="canonical-0021333113313121-1120212323021001-1321100132110312-0021031121021231-0010001222330020-0011022232200022-2320322132212211-3012312031110311"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for property validation settings default.

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

<a id="canonical-2200322100312233-1331332231131222-2201003130302321-0312322000110013-0131003103300102-0103311103220223-2320201221023111-2120100012312210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_disabled` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [api_specification](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-0312023312231331-3233323101302200-0303011303013312-3001210323313032-0323300100221230-2302012222103202-1313023310123112-2311001320021111)
- api_specification.validation_disabled

<a id="canonical-2232110123131233-1331020120332301-1312012022203131-1013312011202201-3030330113202302-3210021122020230-3002203320030210-3101100220322302"></a>

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

<a id="canonical-0113031001311221-2011130203032111-1002013120300301-3110122301220113-0213203302131033-1313023111332322-0211210010133302-0232301320303231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_firewall` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- app_firewall

<a id="canonical-1311110300332332-2133320331021102-0032303300011320-2202321033023003-0232332302313222-2232123202112131-0333231113210200-2023331003030033"></a>

Type: `"single"`. Computed.

\[OneOf: app\_firewall, disable\_waf; Default: disable\_waf\] Type establishes a direct reference
from one object(the referrer) to another(the referred). Such a reference is in form of
tenant/namespace/name.

Additional upstream details:

This type establishes a direct reference from one object(the referrer) to another(the referred).

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

- [app_firewall](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1311110300332332-2133320331021102-0032303300011320-2202321033023003-0232332302313222-2232123202112131-0333231113210200-2023331003030033)
- [disable_waf](data-sources--cdn_loadbalancer--reference--group-011.md#canonical-2122210121011302-0110121231301020-0200000013333201-2333003210032130-3133020131132302-0002220100021211-1112113032300322-0202210321020123)

Select alternatives according to the provider validators above.

<a id="canonical-3232001130121232-1333130001312330-0202130120201233-1222221123120010-0222221103000231-3332213132221020-2023330011311030-2332302112210301"></a>

### Direct properties for `app_firewall`

<a id="canonical-0101131302322201-2112122100001220-1011120113023130-0233203210121110-1101012123332031-2121223302232131-1032010210113012-0012232122312213"></a>

#### `app_firewall.name` property

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

<a id="canonical-2223033231310023-2300003111103120-1132333211123112-1013332213301023-0202003110300103-0222310310033012-1003033101230232-0130221212200201"></a>

<a id="canonical-1122013123301130-2130011103113331-2302032000220203-3323320013210001-0100031220010032-1300333210002230-3113331000020201-2210100102303100"></a>

#### `app_firewall.namespace` property

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

<a id="canonical-2312221032212200-2332102033331023-0303131103112022-0101030110100130-0232312031030033-2232023201223132-3303202010031130-1010020131230322"></a>

<a id="canonical-1203030333222011-2113332321320020-3030133222022130-0110233030101202-1332230221012030-2111302311112322-2031321303131113-3322122223023013"></a>

#### `app_firewall.tenant` property

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

<a id="canonical-1012001011110000-2201321320101011-1201112211033131-1103031022310303-3032023110111010-1101322020301222-0330111212300210-2230202303200212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_clients` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- blocked_clients

<a id="canonical-3031123300221013-2100221323201303-3132230313003313-3131023231122313-2103011202110313-1323020121322301-3333333233332132-0113323213130222"></a>

Type: `"list"`. Computed.

Define rules to block IP Prefixes or AS numbers.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-2313321311311020-0330210121321002-3300020010132000-1223332122212220-2202203223211232-3131002211123321-0300001023211133-1233302013232002"></a>

### Direct properties for `blocked_clients`

<a id="canonical-0121200000200230-0021230132301003-1013332210310103-1103011120232221-1131123332211123-1002023221102021-1020130310302311-1131211320300222"></a>

#### `blocked_clients.actions` property

Type: `["list", "string"]`. Computed.

\[Enum:
SKIP\_PROCESSING\_WAF|SKIP\_PROCESSING\_BOT|SKIP\_PROCESSING\_MUM|SKIP\_PROCESSING\_IP\_REPUTATION|SKIP\_PROCESSING\_API\_PROTECTION|SKIP\_PROCESSING\_OAS\_VALIDATION|SKIP\_PROCESSING\_DDOS\_PROTECTION|SKIP\_PROCESSING\_THREAT\_MESH|SKIP\_PROCESSING\_MALWARE\_PROTECTION\]
Actions that should be taken when client identifier matches the rule. Possible values are
\`SKIP\_PROCESSING\_WAF\`, \`SKIP\_PROCESSING\_BOT\`, \`SKIP\_PROCESSING\_MUM\`,
\`SKIP\_PROCESSING\_IP\_REPUTATION\`, \`SKIP\_PROCESSING\_API\_PROTECTION\`,
\`SKIP\_PROCESSING\_OAS\_VALIDATION\`, \`SKIP\_PROCESSING\_DDOS\_PROTECTION\`,
\`SKIP\_PROCESSING\_THREAT\_MESH\`, \`SKIP\_PROCESSING\_MALWARE\_PROTECTION\`. Defaults to
\`SKIP\_PROCESSING\_WAF\`.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 10,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3122101033303123-3030010103003303-2123212333013222-0122200121131112-3333233110222033-3122332200200021-2133300202133111-1123000220220013"></a>

<a id="canonical-1221332101103321-0202132332220001-2201021030013202-3000203222232313-1130302330113203-1100033332331223-1212303301102333-0023331333333312"></a>

#### `blocked_clients.as_number` property

Type: `"number"`. Computed.

Exclusive with \[http\_header ip\_prefix IPv6\_prefix user\_identifier\] RFC 6793 defined 4-byte AS
number.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 401308,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "401308"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "401308"
  }
}
```

- [bot_skip_processing](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000232312200003-1123300321221231-0333131013101212-1313332201132201-1111201212011030-2032231231010201-0330202021013002-1203000323220303): complete subsection reference.

<a id="canonical-2311021213032333-0030311123212033-0013322023012332-0232221100201103-2200103100100102-2223331022212222-1012130230020003-0222103012010233"></a>

<a id="canonical-1332132321131322-2312130322222311-3222300212132133-3031131203223010-0233020123331220-3233003022213231-3301132201020323-1313113112020111"></a>

#### `blocked_clients.expiration_timestamp` property

Type: `"string"`. Computed.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Additional upstream details:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
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
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  }
}
```

- [http_header](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3222111102020121-2200312331132301-2313123020323212-2210321013123113-3130320020110303-2000213233000120-3333101322102201-2033103200133302): complete subsection reference.

<a id="canonical-2202223000121100-2301101202311230-1022101030312301-2002032321131200-0230032032202011-0122202122223030-3121203023313302-0220321231323002"></a>

<a id="canonical-2230102300123301-3332102233022123-1022133012130120-3032131102301332-0202130303231020-1332003130333312-2020012111102230-2112221100113100"></a>

#### `blocked_clients.ip_prefix` property

Type: `"string"`. Computed.

Exclusive with \[as\_number http\_header IPv6\_prefix user\_identifier\] IPv4 prefix string.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4_prefix": "true"
  }
}
```

<a id="canonical-2231321332321233-2111030303211011-0011003123111221-2222102323332120-2110002203020201-0230333111013030-2332011031033102-0123323301100221"></a>

<a id="canonical-2233121230212021-1010322231300120-0112301113230333-1011231110112100-1101003300023301-1203200110011322-3310001301312332-1221330303312010"></a>

#### `blocked_clients.ipv6_prefix` property

Type: `"string"`. Computed.

Exclusive with \[as\_number http\_header ip\_prefix user\_identifier\] IPv6 prefix string.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true"
  }
}
```

- [metadata](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0332122220010132-3003212310132001-3200233131130203-3002000320030333-3223100311203221-2013202331301022-2002302000000323-1210201223011100): complete subsection reference.

- [skip_processing](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2213100101223122-0130122203032003-2313212221101303-3022313132331032-3300322100123003-3312130112333203-2123201210320333-3001302303003303): complete subsection reference.

<a id="canonical-0110001331132031-1000112332103122-1032110202320213-0223102103113100-1111031011201311-3001133330212300-0221301201123311-0100310222022022"></a>

<a id="canonical-3210223203020222-3133310233023122-0233211211332311-0310232110233313-0103320100010202-2012303320122202-0301003100122330-3022320123330332"></a>

#### `blocked_clients.user_identifier` property

Type: `"string"`. Computed.

Exclusive with \[as\_number http\_header ip\_prefix IPv6\_prefix\] Identify user based on user
identifier. User identifier value needs to be copied from security event.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [waf_skip_processing](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0130110100100120-2303231110321211-1013122300233201-1201111230323020-3033111300331133-3210003230300330-2000232003002333-2130012312022332): complete subsection reference.

<a id="canonical-2000232312200003-1123300321221231-0333131013101212-1313332201132201-1111201212011030-2032231231010201-0330202021013002-1203000323220303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_clients.bot_skip_processing` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [blocked_clients](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1012001011110000-2201321320101011-1201112211033131-1103031022310303-3032023110111010-1101322020301222-0330111212300210-2230202303200212)
- blocked_clients.bot_skip_processing

<a id="canonical-1221301201301230-1230002231230022-1213133023332001-2230020132133323-1320213211110233-3300301001203122-1111010130331200-1003211023211322"></a>

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

<a id="canonical-3222111102020121-2200312331132301-2313123020323212-2210321013123113-3130320020110303-2000213233000120-3333101322102201-2033103200133302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_clients.http_header` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [blocked_clients](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1012001011110000-2201321320101011-1201112211033131-1103031022310303-3032023110111010-1101322020301222-0330111212300210-2230202303200212)
- blocked_clients.http_header

<a id="canonical-3310002130213130-3232022220101001-1233311300313111-1300231232020021-2133200010012030-0131300122203311-3331112003301313-2200003333313000"></a>

Type: `"single"`. Computed.

Configuration parameter for http header.

Additional upstream details:

Request header name and value pairs.

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

<a id="canonical-3233330002321310-1010010133221302-2221123200333120-1013021123001101-0212220110331133-3110111123000121-3023312102303120-2223020011102212"></a>

### Direct properties for `blocked_clients.http_header`

- [headers](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3323103231233331-0320102122313133-0302330312020003-3200013223121301-1333323210211202-0000213201213002-3011120310102003-3100131023303111): complete subsection reference.

<a id="canonical-3323103231233331-0320102122313133-0302330312020003-3200013223121301-1333323210211202-0000213201213002-3011120310102003-3100131023303111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_clients.http_header.headers` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [blocked_clients](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1012001011110000-2201321320101011-1201112211033131-1103031022310303-3032023110111010-1101322020301222-0330111212300210-2230202303200212)
- [blocked_clients.http_header](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3222111102020121-2200312331132301-2313123020323212-2210321013123113-3130320020110303-2000213233000120-3333101322102201-2033103200133302)
- blocked_clients.http_header.headers

<a id="canonical-0010311232300321-1000303023113311-2300130110000310-2101222030033311-3233020001033021-2122203022113100-3223120012230212-3021033211302311"></a>

Type: `"list"`. Computed.

List of HTTP header name and value pairs.

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
    },
    "minItems": 0,
    "uniqueItems": false
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

<a id="canonical-2122321003032233-3012223303322333-0013031031001303-0031133312322021-3220301213121303-3100102013301321-1333023312021230-2210331121100212"></a>

### Direct properties for `blocked_clients.http_header.headers`

<a id="canonical-3312111322200333-2120323331102220-2133023032310231-2120102100222210-2020210113213131-2200121011030300-0313122000321213-0212323303103303"></a>

#### `blocked_clients.http_header.headers.exact` property

Type: `"string"`. Computed.

Exclusive with \[presence regular expression\] Header value to match exactly.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-0131132110131200-1032202231012030-0001020202000321-1100330323120202-2023031213201223-1320122010211321-0213333111331110-1202210201232311"></a>

<a id="canonical-1332330310011110-2203310132210132-1031002233133003-3203111020321010-1110100312202300-2121323310101021-3301020122220100-0102122232011200"></a>

#### `blocked_clients.http_header.headers.invert_match` property

Type: `"bool"`. Computed.

Invert the result of the match to detect missing header or non-matching value.

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

<a id="canonical-0322210211320211-0223130232120301-3331223223021020-0232103223121320-2312001120220220-2101322300322302-3233210123300322-1311201330021021"></a>

<a id="canonical-1110100222112123-1021130130213121-1320202313120030-0331132320230230-3130213012232233-2022022213112301-0212231110211112-2113120213132010"></a>

#### `blocked_clients.http_header.headers.name` property

Type: `"string"`. Computed.

Name. Name of the header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-2210020001000303-0010321320331330-1323312202013011-2001223302022320-0102112230213032-3020113011103222-2213011110113000-1230122332021311"></a>

<a id="canonical-2200230131303321-2121020103333002-3233010122213132-3313312323121221-1123120213101303-2100013231321123-0310302032002322-2030113123301112"></a>

#### `blocked_clients.http_header.headers.presence` property

Type: `"bool"`. Computed.

Exclusive with \[exact regular expression\] If true, check for presence of header.

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

<a id="canonical-1031031033332000-1321203212200133-1031313311122323-0213201110223333-0333022122101001-0011130233202222-3300022301313301-1021233233033111"></a>

<a id="canonical-2323111213212133-3013123320033110-2101333121113033-3322301202000233-2220012011330100-0312132322230101-3233123320131320-0221001121232213"></a>

#### `blocked_clients.http_header.headers.regex` property

Type: `"string"`. Computed.

Exclusive with \[exact presence\] regular expression match of the header value in re2 format.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-0332122220010132-3003212310132001-3200233131130203-3002000320030333-3223100311203221-2013202331301022-2002302000000323-1210201223011100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_clients.metadata` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [blocked_clients](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1012001011110000-2201321320101011-1201112211033131-1103031022310303-3032023110111010-1101322020301222-0330111212300210-2230202303200212)
- blocked_clients.metadata

<a id="canonical-3011112210032213-0200001301201002-0002120300233200-2101101323110203-1313222300203022-0102031021022223-1303002010121332-3200232230111221"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-3031120113213233-2322322100133302-3320220023311230-0221000313021120-0313002210113012-2321132211232000-1200121003120223-2011233000231033"></a>

### Direct properties for `blocked_clients.metadata`

<a id="canonical-3221303203300311-3310222013231211-2113210022020112-3333003232301330-0023231332112121-0232002003130203-2112303312112320-1202013111202022"></a>

#### `blocked_clients.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-3021103133232030-3120332120330301-1330333300231212-3012100102022322-0233320031231103-2113020032033112-0110233112120322-1111300001211002"></a>

<a id="canonical-2021321221232311-3111311230011200-3033222221032311-3321303210133203-1132120200230233-2213202331032101-1033322212000200-3321101133233020"></a>

#### `blocked_clients.metadata.name` property

Type: `"string"`. Computed.

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-2213100101223122-0130122203032003-2313212221101303-3022313132331032-3300322100123003-3312130112333203-2123201210320333-3001302303003303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_clients.skip_processing` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [blocked_clients](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1012001011110000-2201321320101011-1201112211033131-1103031022310303-3032023110111010-1101322020301222-0330111212300210-2230202303200212)
- blocked_clients.skip_processing

<a id="canonical-1222231102023020-2000131101313203-0200231020121332-2333121002331222-3200002310031233-0320331232322301-1210210101330221-3210233202133321"></a>

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

<a id="canonical-0130110100100120-2303231110321211-1013122300233201-1201111230323020-3033111300331133-3210003230300330-2000232003002333-2130012312022332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_clients.waf_skip_processing` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [blocked_clients](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1012001011110000-2201321320101011-1201112211033131-1103031022310303-3032023110111010-1101322020301222-0330111212300210-2230202303200212)
- blocked_clients.waf_skip_processing

<a id="canonical-2033030030301221-3122113203233233-3011013120202333-1010122012203011-0302031102232322-3323123130333010-2103010102213232-1120322110313200"></a>

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

<a id="canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- bot_defense

<a id="canonical-2200110012102030-1101223103010130-1302102213101302-3003213131112323-1312333021000202-3112120003121220-1020101311333133-3010002001320212"></a>

Type: `"single"`. Computed.

This defines various configuration OPTIONS for Bot Defense Policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cors_support_choice": "[\"disable_cors_support\",\"enable_cors_support\"]"
}
```

<a id="canonical-2203230021003310-2212322011211122-0213321110313201-3123231032302201-3122333110021302-3123113202300033-3322320210120032-0322313300011332"></a>

### Direct properties for `bot_defense`

- [disable_cors_support](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0220321332121111-1132233033210203-0120113230232023-3020203103011323-1012212203021000-2131322303313333-0023201201123030-0203330031230112): complete subsection reference.

- [enable_cors_support](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1101323212321033-3331011001021301-0113112011232131-3121022131023122-1202202022310203-2021121130021010-1330331331102312-1203023012203233): complete subsection reference.

- [policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111): complete subsection reference.

<a id="canonical-3101012110121222-2133102131333111-1313301000210213-2031310011221222-3312300003133300-3131003212101021-0223131023213032-2232311322222011"></a>

<a id="canonical-2130133220330211-3300331313310200-2002311111230112-2231030032231122-1130130012031121-1233321322212230-1033011210231121-1331123100200013"></a>

#### `bot_defense.regional_endpoint` property

Type: `"string"`. Computed.

\[Enum: AUTO|US|EU|ASIA\] Defines a selection for Bot Defense region - AUTO: AUTO Automatic
selection based on client IP address - US: US US region - EU: EU European Union region - ASIA: ASIA
Asia region. Possible values are \`AUTO\`, \`US\`, \`EU\`, \`ASIA\`. Defaults to \`AUTO\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "AUTO",
  "enum": [
    "AUTO",
    "US",
    "EU",
    "ASIA"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1131031120203320-2220001032201102-0320230111023111-3023111233320022-0113030302323121-1113300323000211-2202203211310123-1031303111122122"></a>

<a id="canonical-2022323302002320-2322112203200232-0112021033230330-3101002001311323-2232000111130210-0011320203102001-3330332323031033-1323310210223233"></a>

#### `bot_defense.timeout` property

Type: `"number"`. Computed.

The timeout for the inference check, in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-0220321332121111-1132233033210203-0120113230232023-3020203103011323-1012212203021000-2131322303313333-0023201201123030-0203330031230112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.disable_cors_support` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- bot_defense.disable_cors_support

<a id="canonical-3122123232113000-2213122221203131-0100101002120202-1131233330301132-2331110101233001-2322233301222022-2012100320302312-1130000022121333"></a>

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

<a id="canonical-1101323212321033-3331011001021301-0113112011232131-3121022131023122-1202202022310203-2021121130021010-1330331331102312-1203023012203233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.enable_cors_support` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- bot_defense.enable_cors_support

<a id="canonical-3103132222232232-0233333210203321-1330011122322112-1322303231132201-2320133220102001-0103121112323222-2330132321223203-0011001010220320"></a>

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

<a id="canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- bot_defense.policy

<a id="canonical-2320030302020113-3201031010311320-1020022233332120-0003312032201220-3111003020001202-0130112213122223-3130002133010013-1220010032231220"></a>

Type: `"single"`. Computed.

This defines various configuration OPTIONS for Bot Defense policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insert_all_pages\",\"js_insert_all_pages_except\",\"js_insertion_rules\"]",
  "x-ves-oneof-field-mobile_sdk_choice": "[\"disable_mobile_sdk\",\"mobile_sdk_config\"]"
}
```

<a id="canonical-3103302110311230-1310330013321130-3330032312302331-0210322310112013-2121230221000232-1200320233031333-3000220022223110-1333200112123100"></a>

### Direct properties for `bot_defense.policy`

- [disable_js_insert](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2331222130233333-3233120003332120-1203210331221010-1202002200131010-0103310113310101-1023311102332330-2000310123000223-3012211331111131): complete subsection reference.

- [disable_mobile_sdk](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2210230332322302-2113322223033313-1030112123331232-0010020011013110-0222030330103332-2010230221300103-1032031001022201-0331011123330322): complete subsection reference.

<a id="canonical-1011123212022030-3031132001202030-3121202221003002-2011130113213012-2210111003221321-2022301123001121-2001023010312231-2320303120321112"></a>

<a id="canonical-1330121303011200-2320131011231132-1312323132011112-2323233331310130-1131033203223000-2311200210330122-2320021001233022-0002001013301211"></a>

#### `bot_defense.policy.javascript_mode` property

Type: `"string"`. Computed.

\[Enum: ASYNC\_JS\_NO\_CACHING|ASYNC\_JS\_CACHING|SYNC\_JS\_NO\_CACHING|SYNC\_JS\_CACHING\] Web
Client JavaScript Mode. Bot Defense JavaScript for telemetry collection is requested asynchronously,
and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested asynchronously,
and it is cacheable Bot Defense JavaScript for telemetry collection is requested.. Possible values
are \`ASYNC\_JS\_NO\_CACHING\`, \`ASYNC\_JS\_CACHING\`, \`SYNC\_JS\_NO\_CACHING\`,
\`SYNC\_JS\_CACHING\`. Defaults to \`ASYNC\_JS\_NO\_CACHING\`.

Additional upstream details:

Web Client JavaScript Mode. Bot Defense JavaScript for telemetry collection is requested
asynchronously, and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested
asynchronously, and it is cacheable Bot Defense JavaScript for telemetry collection is requested
synchronously, and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested
synchronously, and it is cacheable.

Receipt-pinned upstream constraints:

```json
{
  "default": "ASYNC_JS_NO_CACHING",
  "enum": [
    "ASYNC_JS_NO_CACHING",
    "ASYNC_JS_CACHING",
    "SYNC_JS_NO_CACHING",
    "SYNC_JS_CACHING"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1301220130213313-0032111222102223-3111212333332321-0203102111110101-3132321030132230-3003201111003311-3102210100031323-2130010123102331"></a>

<a id="canonical-2230022123030012-1203011013132122-0232103223221010-2000100200321320-3020103020201231-2101301021211031-0312332330020122-2012220111331111"></a>

#### `bot_defense.policy.js_download_path` property

Type: `"string"`. Computed.

Customize Bot Defense Client JavaScript path. If not specified, default \`/common.js\`

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  }
}
```

- [js_insert_all_pages](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0101120322313112-2121032301302102-2131103332212003-0130020330220130-1121200132333113-0323311122012212-0131023132122330-1232021023111111): complete subsection reference.

- [js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2022211013032221-0322203320020120-2233313130120312-1220111100112121-1330312300220022-0033222022020111-2110001111121102-0111001301120220): complete subsection reference.

- [js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0213332331203203-1301130331322102-1031013002231302-2220131212030113-3130331233230213-2202020332110333-1013222022322212-3221113033201331): complete subsection reference.

- [mobile_sdk_config](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-2212311100013330-3302031011003100-3121130332112301-2222033303133101-0101302002131221-3213321232023301-3022211202010011-1010112330103303): complete subsection reference.

- [protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0333103032312221-3013021303201032-1120321232101312-3113231132012221-2103103303322103-1213232220131000-1132323130120031-0331111033201123): complete subsection reference.

<a id="canonical-2331222130233333-3233120003332120-1203210331221010-1202002200131010-0103310113310101-1023311102332330-2000310123000223-3012211331111131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.disable_js_insert` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- bot_defense.policy.disable_js_insert

<a id="canonical-3300003122021331-2230023312201001-3201213321231130-2232222102032002-1123122023002012-1222123230031103-3203030022213131-2010030322120132"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable js insert.

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

<a id="canonical-2210230332322302-2113322223033313-1030112123331232-0010020011013110-0222030330103332-2010230221300103-1032031001022201-0331011123330322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.disable_mobile_sdk` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- bot_defense.policy.disable_mobile_sdk

<a id="canonical-1323002213001003-3033102011100110-3232113213200212-0031200132112021-2313212220130103-1230203232221323-3131301033230000-0023333101323011"></a>

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

<a id="canonical-0101120322313112-2121032301302102-2131103332212003-0130020330220130-1121200132333113-0323311122012212-0131023132122330-1232021023111111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insert_all_pages` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- bot_defense.policy.js_insert_all_pages

<a id="canonical-3232120300031120-1300233230000320-1000230131120223-2322120301313002-1020232231302123-2133221322131322-0123222313021011-0223000023223213"></a>

Type: `"single"`. Computed.

Insert Bot Defense JavaScript in all pages.

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

<a id="canonical-0131001311012302-3020212321221202-1300031003002302-0003010003120200-3300212131010032-2023202301320212-2203321332103213-1102332110021110"></a>

### Direct properties for `bot_defense.policy.js_insert_all_pages`

<a id="canonical-0011103323002103-2031230102232033-1010220130302223-0002102103100010-3033320320013131-1122231022313333-3303111302023212-3321303013321202"></a>

#### `bot_defense.policy.js_insert_all_pages.javascript_location` property

Type: `"string"`. Computed.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2022211013032221-0322203320020120-2233313130120312-1220111100112121-1330312300220022-0033222022020111-2110001111121102-0111001301120220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insert_all_pages_except` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- bot_defense.policy.js_insert_all_pages_except

<a id="canonical-1202100222221023-1003210212121212-0321210222120111-0312330313132030-1231332311030312-1122322223223302-3331322100202322-2002001112302322"></a>

Type: `"single"`. Computed.

Insert Bot Defense JavaScript in all pages with the exceptions.

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

<a id="canonical-0012221032332130-0212120210111303-2100321323323201-2122200101200222-0213310333123022-3312023111212021-1321231031323312-1032233122201230"></a>

### Direct properties for `bot_defense.policy.js_insert_all_pages_except`

- [exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1332331100221221-3220203213321032-1003020212222120-1210000010023112-2112102203222113-0312110010010002-1200331310033331-1301112221012310): complete subsection reference.

<a id="canonical-2013300112311010-0122100023211202-3230322130002301-3020210002212030-1210213203103121-1123330001000223-2101231320212102-0021110002300210"></a>

<a id="canonical-0302120322311130-2332332322013201-3211110201030102-2002222000031333-1110122221132220-0031033100223201-2311233302023002-1000232320212201"></a>

#### `bot_defense.policy.js_insert_all_pages_except.javascript_location` property

Type: `"string"`. Computed.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1332331100221221-3220203213321032-1003020212222120-1210000010023112-2112102203222113-0312110010010002-1200331310033331-1301112221012310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insert_all_pages_except.exclude_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2000302003002003-1302112013101323-0110300221323130-0312013122130131-3122223032023133-2111310301100130-3113012002330202-0112012313123200)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111)
- [bot_defense.policy.js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-2022211013032221-0322203320020120-2233313130120312-1220111100112121-1330312300220022-0033222022020111-2110001111121102-0111001301120220)
- bot_defense.policy.js_insert_all_pages_except.exclude_list

<a id="canonical-1312222123313202-0100030220120332-0012312333302311-3313320322031030-3002101320212312-0201332203032313-3310330010111312-2013131003231023"></a>

Type: `"list"`. Computed.

Optional JavaScript insertions exclude list of domain and path matchers.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3332123021132101-0133023113321220-3100001010023312-2031103102310101-0211001010021003-0020131002203220-2332002210101020-1103211013300122"></a>

### Direct properties for `bot_defense.policy.js_insert_all_pages_except.exclude_list`

- [any_domain](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-3221022013311210-1303333312110100-3231202033012211-1212032101120010-3012033110122123-2112021112133103-3031121010220331-2300033211120132): complete subsection reference.

- [domain](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-3302131133301200-2312333313230321-0231022102310322-2030333233012300-0330222233233023-1030020001221232-2130212320321310-1200012022223021): complete subsection reference.

- [metadata](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-3002001232030111-2330311122330111-1001301203133100-0311232213323021-1222202301002021-2123322010233023-1200033010220110-1013223312300122): complete subsection reference.

- [path](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-0303132031033301-0102301230131313-1233011001011330-3221110132312121-1303013211012200-1331301112222131-2300311201113312-3332201113113303): complete subsection reference.
