---
page_title: "xcsh_healthcheck reference"
subcategory: "Monitoring"
description: "Complete grouped canonical reference for xcsh_healthcheck reference."
---

# xcsh_healthcheck reference

<a id="canonical-2233200310100332-2033003002302200-0110210311211131-2311120103323110-0232222133220131-1110332103312313-1221101021321013-1120313110032312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_healthcheck](../data-sources/healthcheck.md#canonical-1302112230101233-2202130312331232-0231113312211032-1103323302011301-3303300020333020-2302311211132100-1303020110013123-3330132302312212)
- Property reference

<a id="canonical-2203133122130221-0100222132321233-1102022223030330-2032311112110202-3102122123103032-2320220003211312-2133013313312123-3123310210213002"></a>

### Direct properties for `xcsh_healthcheck`

<a id="canonical-1232211122333122-1323103001301102-2322121103222232-1131300212113113-1303233300130132-0122333302221200-3113000323232333-3221223020130020"></a>

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

- [default_jitter](data-sources--healthcheck--reference--group-001.md#canonical-0322032010000220-3333013213012021-1212301213233322-2302110220211022-2021313303330311-0102320321001210-0030103103101313-3320111112233323): complete subsection reference.

<a id="canonical-0321223302100302-2133330103022023-3101032022300313-3103303122100321-1221322211331233-1212233001320233-3030331322201303-1133030130201303"></a>

<a id="canonical-3323101213300100-2212332332202231-0312321133211131-1032123130302101-0322330022311131-3221130300102302-3120230301031223-0322002013333232"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the Healthcheck.

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

<a id="canonical-0112320222123113-2220002313211022-0211332201200111-0103331120132221-3122112031223013-0130200300200331-0212220331221003-0300202011322132"></a>

<a id="canonical-0132002133003002-2202311130310011-0323031121332102-1122002200232220-1321213100101113-0013030212133120-0231121323331102-3011011202233230"></a>

#### `healthy_threshold` property

Type: `"number"`. Computed.

Number of successful responses before declaring healthy. In other words, this is the number of
healthy health checks required before a host is marked healthy. Note that during startup, only a
single successful health check is required to mark a host healthy. Recommended: \`3\`.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "threshold",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "category": "threshold",
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

- [http_health_check](data-sources--healthcheck--reference--group-001.md#canonical-2003303010100100-0312223313133022-2122302221230233-1022031031211221-2021003312003010-0120011031122331-0233001303013031-2103231033203121): complete subsection reference.

<a id="canonical-1101003303311010-0213102330230300-0123203321131021-0332232332202103-0113330203101300-2130010121311111-2222222231102303-0120232213321223"></a>

<a id="canonical-2331120222302321-1030022332102311-1302201213223113-1121212100311013-3330231332210303-1032310002202223-3033110323211312-3321033303030233"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2221320233212121-0221030331212211-3130310222301220-0000232231023010-0010333003210113-2123323233313330-3032000231023231-3322111003233023"></a>

<a id="canonical-2103213112312120-3001131132233230-3312112320013313-3322030330233312-1102003302302232-0223030200031303-2313101023311323-0320333333033111"></a>

#### `interval` property

Type: `"number"`. Computed.

Time interval in seconds between two healthcheck requests. Recommended: \`15\`.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3003001021322123-1022101303321202-1111321221030303-3113010212112222-0311303131010221-2302023133200110-1221031211233332-2231232130013301"></a>

<a id="canonical-1112032123123302-2211333131310200-1112223001231112-3202122332310203-0301332012123130-2303230233323220-3303012100231211-0013220310220331"></a>

#### `jitter_percent` property

Type: `"number"`. Computed.

Exclusive with \[default\_jitter\] Specify a custom jitter value as a percentage of the health check
interval. Valid values are 0 (to disable jitter) and 10 to 50. Server applies default when omitted.
Recommended: \`30\`.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "timing",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "category": "timing",
      "confidence": 0.99,
      "note": "Non-contiguous: {0} union [10, 50] — values 1-9 rejected by API",
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "multipleOf": 1,
    "ranges": [
      {
        "maximum": 0,
        "minimum": 0
      },
      {
        "maximum": 50,
        "minimum": 10
      }
    ]
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-50"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-50"
  }
}
```

<a id="canonical-0232011101001231-1323132223131320-3202303220030202-3220322223011000-0222101110233333-0222130133011102-3320121021102103-0013302323300222"></a>

<a id="canonical-2022120121310302-2100330133312332-1302230112312313-3311231331321132-2211323101111123-1303011021213132-1310110213111201-1010203323021023"></a>

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

<a id="canonical-3002222103111133-2020000010200220-0333303023312213-1331030200203300-1120010213022013-2011012033231023-1023103113223003-2101211113130312"></a>

<a id="canonical-0202133131132032-1013133210111030-2323123320311320-3321120002130122-2131300331112120-0102031010333301-1213123233031100-1012332110233012"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Healthcheck.

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

<a id="canonical-0331302212213331-3121333122121211-2103103300321110-3300013103021331-1323032311123131-1232110032103112-3030000000003213-2323002230130321"></a>

<a id="canonical-1330303012023023-1201302012023120-3330131202330101-0111003313230102-2002312221113302-2132133000131311-3113121032310113-0133020232102211"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Healthcheck exists.

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

- [tcp_health_check](data-sources--healthcheck--reference--group-001.md#canonical-1330301203333120-1033112131230233-2331230222230103-3210322303111102-0231333321033120-1333231120301121-1132033312010021-3232320102232110): complete subsection reference.

<a id="canonical-1012323120331100-0033021130112012-0311132301201023-1313133233300003-0032100003321021-3003111312332010-0323132330322013-2032130101030222"></a>

<a id="canonical-3222323312303310-3303132033211032-3033331233120330-3002333000003001-3122130303030113-0322031202123001-3013322123211300-2300231011103332"></a>

#### `timeout` property

Type: `"number"`. Computed.

Timeout in seconds to wait for successful response. In other words, it is the time to wait for a
health check response. If the timeout is reached the health check attempt will be considered a
failure. Recommended: \`3\`.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "timing",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600,
    "metadata": {
      "category": "timing",
      "confidence": 0.99,
      "note": "API rejects timeout > 600 for healthchecks (global pattern says 3600)",
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [udp_icmp_health_check](data-sources--healthcheck--reference--group-001.md#canonical-2000010223012103-0021100002113232-0202333221122111-1223003032022202-3030103302000301-2011220103210012-0020032300010103-1200022022300232): complete subsection reference.

<a id="canonical-2310201322003110-1122011110312210-0230301323122313-2321302220301300-2230121023000011-0003021030110022-1301123331002221-2212203032312221"></a>

<a id="canonical-0013003330313021-3230221100131100-0322102323113331-1332110021013223-0232131112023033-3203020302011011-3110310203300301-2120232112330201"></a>

#### `unhealthy_threshold` property

Type: `"number"`. Computed.

Number of failed responses before declaring unhealthy. In other words, this is the number of
unhealthy health checks required before a host is marked unhealthy. Note that for HTTP health
checking if a host responds with 503 this threshold is ignored and the host is considered unhealthy
immediately. Recommended: \`1\`.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "threshold",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16,
    "metadata": {
      "category": "threshold",
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.lte": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "16"
  }
}
```

<a id="canonical-2233333313021330-1320232120223200-1320103000012311-1302130132203131-3113230230112033-0203203211230120-1111220100311332-2210131312221101"></a>

### All schema paths for `xcsh_healthcheck`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--healthcheck--reference--group-001.md#canonical-1232211122333122-1323103001301102-2322121103222232-1131300212113113-1303233300130132-0122333302221200-3113000323232333-3221223020130020) |
| `default_jitter` | [default_jitter](data-sources--healthcheck--reference--group-001.md#canonical-1122120131222012-0000212202133323-1232001231232233-0030023321111012-2223032031022030-2132033000102122-3300122101033323-0332311330032301) |
| `description` | [description](data-sources--healthcheck--reference--group-001.md#canonical-0321223302100302-2133330103022023-3101032022300313-3103303122100321-1221322211331233-1212233001320233-3030331322201303-1133030130201303) |
| `healthy_threshold` | [healthy_threshold](data-sources--healthcheck--reference--group-001.md#canonical-0112320222123113-2220002313211022-0211332201200111-0103331120132221-3122112031223013-0130200300200331-0212220331221003-0300202011322132) |
| `http_health_check` | [http_health_check](data-sources--healthcheck--reference--group-001.md#canonical-2331330213023223-2011030311202130-2220033223233230-2303201323020232-0121220303101223-3230320131223322-2132202000222301-3211132112123100) |
| `http_health_check.expected_response` | [http_health_check.expected_response](data-sources--healthcheck--reference--group-001.md#canonical-2213013201333330-3331032013031302-3302023223030020-3231233311200230-2111310303121101-3033033033002133-1103310320021133-0211231122011122) |
| `http_health_check.expected_status_codes` | [http_health_check.expected_status_codes](data-sources--healthcheck--reference--group-001.md#canonical-0311300000021211-1001210000033021-0001132201023311-3120000331303333-2222103231211133-3110231031230330-1200323301113202-1001212002323031) |
| `http_health_check.headers` | [http_health_check.headers](data-sources--healthcheck--reference--group-001.md#canonical-3132202221322210-0030020003132122-0201102110312303-1301200322302111-3312102022320131-1321200031311002-0213111232311230-3213011211323111) |
| `http_health_check.host_header` | [http_health_check.host_header](data-sources--healthcheck--reference--group-001.md#canonical-2210210300320020-0011030032201203-3031222111311112-1221201110321330-1332313003313222-2111233120132123-0322020230321121-3331313321321212) |
| `http_health_check.path` | [http_health_check.path](data-sources--healthcheck--reference--group-001.md#canonical-3210333133113130-3120100302222031-3021210230322311-1122113322112030-3312222133130203-0231132232001331-2311321033300000-1300011212130311) |
| `http_health_check.request_headers_to_remove` | [http_health_check.request_headers_to_remove](data-sources--healthcheck--reference--group-001.md#canonical-0323312301210033-1233030000002202-3200110122220033-3000002131110320-1022033123110033-3000020200012002-2232212113013300-3101023331223031) |
| `http_health_check.use_http2` | [http_health_check.use_http2](data-sources--healthcheck--reference--group-001.md#canonical-1211212301222113-2232313313333001-3201020310033212-1000100103020322-3211313300012322-2331002311132210-0000103123132112-0013201112132211) |
| `http_health_check.use_origin_server_name` | [http_health_check.use_origin_server_name](data-sources--healthcheck--reference--group-001.md#canonical-3321302000310110-2012322023111023-3022100211333222-1320200233013020-0122200123230303-3213102332012320-0322100231001212-2302013022301000) |
| `id` | [ID](data-sources--healthcheck--reference--group-001.md#canonical-1101003303311010-0213102330230300-0123203321131021-0332232332202103-0113330203101300-2130010121311111-2222222231102303-0120232213321223) |
| `interval` | [interval](data-sources--healthcheck--reference--group-001.md#canonical-2221320233212121-0221030331212211-3130310222301220-0000232231023010-0010333003210113-2123323233313330-3032000231023231-3322111003233023) |
| `jitter_percent` | [jitter_percent](data-sources--healthcheck--reference--group-001.md#canonical-3003001021322123-1022101303321202-1111321221030303-3113010212112222-0311303131010221-2302023133200110-1221031211233332-2231232130013301) |
| `labels` | [labels](data-sources--healthcheck--reference--group-001.md#canonical-0232011101001231-1323132223131320-3202303220030202-3220322223011000-0222101110233333-0222130133011102-3320121021102103-0013302323300222) |
| `name` | [name](data-sources--healthcheck--reference--group-001.md#canonical-3002222103111133-2020000010200220-0333303023312213-1331030200203300-1120010213022013-2011012033231023-1023103113223003-2101211113130312) |
| `namespace` | [namespace](data-sources--healthcheck--reference--group-001.md#canonical-0331302212213331-3121333122121211-2103103300321110-3300013103021331-1323032311123131-1232110032103112-3030000000003213-2323002230130321) |
| `tcp_health_check` | [tcp_health_check](data-sources--healthcheck--reference--group-001.md#canonical-2013313121123120-2221121131000120-0010103123112000-3023323210232211-0011132200002102-3102321323310132-1103130321001131-0220331212133311) |
| `tcp_health_check.expected_response` | [tcp_health_check.expected_response](data-sources--healthcheck--reference--group-001.md#canonical-2001333000123021-3120103110220010-1013331313311321-0012222130313023-3130132213301321-3113101131023002-2033203331003203-3332303210130230) |
| `tcp_health_check.send_payload` | [tcp_health_check.send_payload](data-sources--healthcheck--reference--group-001.md#canonical-0212313312110021-1123231230203132-3031032133001121-2010303320322132-3211002133303012-3211011022022132-2322311203230123-2012110230103012) |
| `timeout` | [timeout](data-sources--healthcheck--reference--group-001.md#canonical-1012323120331100-0033021130112012-0311132301201023-1313133233300003-0032100003321021-3003111312332010-0323132330322013-2032130101030222) |
| `udp_icmp_health_check` | [udp_icmp_health_check](data-sources--healthcheck--reference--group-001.md#canonical-1033303112210301-3232331032212330-3311323033003020-0302030333303210-1020033231323000-2101300110130213-3213002221210220-2102210112011203) |
| `unhealthy_threshold` | [unhealthy_threshold](data-sources--healthcheck--reference--group-001.md#canonical-2310201322003110-1122011110312210-0230301323122313-2321302220301300-2230121023000011-0003021030110022-1301123331002221-2212203032312221) |

<a id="canonical-0322032010000220-3333013213012021-1212301213233322-2302110220211022-2021313303330311-0102320321001210-0030103103101313-3320111112233323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `default_jitter` properties

Breadcrumbs:

- [xcsh_healthcheck](../data-sources/healthcheck.md#canonical-1302112230101233-2202130312331232-0231113312211032-1103323302011301-3303300020333020-2302311211132100-1303020110013123-3330132302312212)
- [Property reference](data-sources--healthcheck--reference--group-001.md#canonical-2233200310100332-2033003002302200-0110210311211131-2311120103323110-0232222133220131-1110332103312313-1221101021321013-1120313110032312)
- default_jitter

<a id="canonical-1122120131222012-0000212202133323-1232001231232233-0030023321111012-2223032031022030-2132033000102122-3300122101033323-0332311330032301"></a>

Type: `["object", {}]`. Computed.

\[OneOf: default\_jitter, jitter\_percent; Default: default\_jitter\] Configuration parameter for
default jitter.

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

- [default_jitter](data-sources--healthcheck--reference--group-001.md#canonical-1122120131222012-0000212202133323-1232001231232233-0030023321111012-2223032031022030-2132033000102122-3300122101033323-0332311330032301)
- [jitter_percent](data-sources--healthcheck--reference--group-001.md#canonical-3003001021322123-1022101303321202-1111321221030303-3113010212112222-0311303131010221-2302023133200110-1221031211233332-2231232130013301)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2003303010100100-0312223313133022-2122302221230233-1022031031211221-2021003312003010-0120011031122331-0233001303013031-2103231033203121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_health_check` properties

Breadcrumbs:

- [xcsh_healthcheck](../data-sources/healthcheck.md#canonical-1302112230101233-2202130312331232-0231113312211032-1103323302011301-3303300020333020-2302311211132100-1303020110013123-3330132302312212)
- [Property reference](data-sources--healthcheck--reference--group-001.md#canonical-2233200310100332-2033003002302200-0110210311211131-2311120103323110-0232222133220131-1110332103312313-1221101021321013-1120313110032312)
- http_health_check

<a id="canonical-2331330213023223-2011030311202130-2220033223233230-2303201323020232-0121220303101223-3230320131223322-2132202000222301-3211132112123100"></a>

Type: `"single"`. Computed.

\[OneOf: http\_health\_check, tcp\_health\_check, udp\_icmp\_health\_check\] Healthy if 'GET' method
on URL 'HTTP(s)://&lt;host&gt;/&lt;path&gt;' with optional '&lt;header&gt;' returns success. 'host'
is not used for DNS resolution. It is used as HTTP Header in the request.

Additional upstream details:

Healthy if "GET" method on URL "HTTP(s)://&lt;host&gt;/&lt;path&gt;" with optional "&lt;header&gt;"
returns success. "host" is not used for DNS resolution.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-host_header_choice": "[\"host_header\",\"use_origin_server_name\"]"
}
```

OneOf alternatives in this subsection:

- [http_health_check](data-sources--healthcheck--reference--group-001.md#canonical-2331330213023223-2011030311202130-2220033223233230-2303201323020232-0121220303101223-3230320131223322-2132202000222301-3211132112123100)
- [tcp_health_check](data-sources--healthcheck--reference--group-001.md#canonical-2013313121123120-2221121131000120-0010103123112000-3023323210232211-0011132200002102-3102321323310132-1103130321001131-0220331212133311)
- [udp_icmp_health_check](data-sources--healthcheck--reference--group-001.md#canonical-1033303112210301-3232331032212330-3311323033003020-0302030333303210-1020033231323000-2101300110130213-3213002221210220-2102210112011203)

Select alternatives according to the provider validators above.

<a id="canonical-0021013323202221-0030201112322222-2203310012000020-0021032001231321-1201103313323030-1101201131313131-3113133110022030-0013121203330102"></a>

### Direct properties for `http_health_check`

<a id="canonical-2213013201333330-3331032013031302-3302023223030020-3231233311200230-2111310303121101-3033033033002133-1103310320021133-0211231122011122"></a>

#### `http_health_check.expected_response` property

Type: `"string"`. Computed.

Raw bytes expected in the response of HTTP health check. Input is to be given in Hex encoded format.
If left empty, then response body is not considered for evaluating health check status. Server
applies default when omitted.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-0311300000021211-1001210000033021-0001132201023311-3120000331303333-2222103231211133-3110231031230330-1200323301113202-1001212002323031"></a>

<a id="canonical-1203301133100311-0232010131220011-1022223211233200-1321021031001321-1233230021103223-2023030331113201-0211003030130323-0331213003232023"></a>

#### `http_health_check.expected_status_codes` property

Type: `["list", "string"]`. Computed.

Specifies a list of HTTP response status codes considered healthy. To treat default HTTP expected
status code 200 as healthy, user has to configure it explicitly. This is a list of strings, each of
which is single HTTP status code or a range with start and end values separated by '-'. Defaults to
\`\[\]\`. Server applies default when omitted.

Additional upstream details:

This is a list of strings, each of which is single HTTP status code or a range with start and end
values separated by "-".

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.items.string.http_status_range": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "10",
    "ves.io.schema.rules.repeated.items.string.min_len": "3",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.http_status_range": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "10",
    "ves.io.schema.rules.repeated.items.string.min_len": "3",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3132202221322210-0030020003132122-0201102110312303-1301200322302111-3312102022320131-1321200031311002-0213111232311230-3213011211323111"></a>

<a id="canonical-1201000132120223-3332011233131201-3332100103313113-3132120232333130-0330232030121023-2200020113111301-2332211011001030-2331133311222201"></a>

#### `http_health_check.headers` property

Type: `["map", "string"]`. Computed.

Specifies a list of HTTP headers that should be added to each request that is sent to the health
checked cluster. This is a list of key-value pairs. Defaults to \`map\[\]\`. Server applies default
when omitted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 16
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 256,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "256",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "16",
      "ves.io.schema.rules.map.values.string.max_len": "2048",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 2048,
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
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "256",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "2048",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-2210210300320020-0011030032201203-3031222111311112-1221201110321330-1332313003313222-2111233120132123-0322020230321121-3331313321321212"></a>

<a id="canonical-2110003031202101-2101202011222130-1330220122220333-2322010311331013-0011222313012002-0302120131131030-0202133330330103-0201100130102321"></a>

#### `http_health_check.host_header` property

Type: `"string"`. Computed.

Exclusive with \[use\_origin\_server\_name\] The value of the host header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 262,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 262,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  }
}
```

<a id="canonical-3210333133113130-3120100302222031-3021210230322311-1122113322112030-3312222133130203-0231132232001331-2311321033300000-1300011212130311"></a>

<a id="canonical-1231202133112321-3010120212033102-3332021120012102-1131100302321132-2012332332301113-1110232333133130-2311313323320002-0313210312102311"></a>

#### `http_health_check.path` property

Type: `"string"`. Computed.

Specifies the HTTP path that will be requested during health checking. Recommended: \`/\`.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-0323312301210033-1233030000002202-3200110122220033-3000002131110320-1022033123110033-3000020200012002-2232212113013300-3101023331223031"></a>

<a id="canonical-2303030232012231-2333303120111032-1223133323101012-0321012031201322-3010030221311030-3123030021300233-2210302220011300-2010100320001331"></a>

#### `http_health_check.request_headers_to_remove` property

Type: `["list", "string"]`. Computed.

Specifies a list of HTTP headers that should be removed from each request that is sent to the health
checked cluster. This is a list of keys of headers. Defaults to \`\[\]\`. Server applies default
when omitted.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-1211212301222113-2232313313333001-3201020310033212-1000100103020322-3211313300012322-2331002311132210-0000103123132112-0013201112132211"></a>

<a id="canonical-0011233223020020-0131221123221122-1131200101032032-1220332113103203-2223210003013132-2212013301232200-2101313112200311-1101310233133101"></a>

#### `http_health_check.use_http2` property

Type: `"bool"`. Computed.

If set, health checks will be made using HTTP/2. Defaults to \`false\`. Server applies default when
omitted. Recommended: \`false\`.

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

- [use_origin_server_name](data-sources--healthcheck--reference--group-001.md#canonical-2312002232323201-3031121100020211-3131233122120133-0320000132012210-3122112001203221-0211331000320303-2102201133112023-0300333213130220): complete subsection reference.

<a id="canonical-2312002232323201-3031121100020211-3131233122120133-0320000132012210-3122112001203221-0211331000320303-2102201133112023-0300333213130220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http_health_check.use_origin_server_name` properties

Breadcrumbs:

- [xcsh_healthcheck](../data-sources/healthcheck.md#canonical-1302112230101233-2202130312331232-0231113312211032-1103323302011301-3303300020333020-2302311211132100-1303020110013123-3330132302312212)
- [Property reference](data-sources--healthcheck--reference--group-001.md#canonical-2233200310100332-2033003002302200-0110210311211131-2311120103323110-0232222133220131-1110332103312313-1221101021321013-1120313110032312)
- [http_health_check](data-sources--healthcheck--reference--group-001.md#canonical-2003303010100100-0312223313133022-2122302221230233-1022031031211221-2021003312003010-0120011031122331-0233001303013031-2103231033203121)
- http_health_check.use_origin_server_name

<a id="canonical-3321302000310110-2012322023111023-3022100211333222-1320200233013020-0122200123230303-3213102332012320-0322100231001212-2302013022301000"></a>

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

<a id="canonical-1330301203333120-1033112131230233-2331230222230103-3210322303111102-0231333321033120-1333231120301121-1132033312010021-3232320102232110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tcp_health_check` properties

Breadcrumbs:

- [xcsh_healthcheck](../data-sources/healthcheck.md#canonical-1302112230101233-2202130312331232-0231113312211032-1103323302011301-3303300020333020-2302311211132100-1303020110013123-3330132302312212)
- [Property reference](data-sources--healthcheck--reference--group-001.md#canonical-2233200310100332-2033003002302200-0110210311211131-2311120103323110-0232222133220131-1110332103312313-1221101021321013-1120313110032312)
- tcp_health_check

<a id="canonical-2013313121123120-2221121131000120-0010103123112000-3023323210232211-0011132200002102-3102321323310132-1103130321001131-0220331212133311"></a>

Type: `"single"`. Computed.

Healthy if TCP connection is successful and response payload matches &lt;expected\_response&gt;.

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

<a id="canonical-3120101133101022-3303132223000301-1001213003133232-1311311201232113-0013330131032232-2222003321110211-3122333211310231-3010320313222330"></a>

### Direct properties for `tcp_health_check`

<a id="canonical-2001333000123021-3120103110220010-1013331313311321-0012222130313023-3130132213301321-3113101131023002-2033203331003203-3332303210130230"></a>

#### `tcp_health_check.expected_response` property

Type: `"string"`. Computed.

Raw bytes expected in the request. Describes the encoding of the payload bytes in the payload. Hex
encoded payload.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-0212313312110021-1123231230203132-3031032133001121-2010303320322132-3211002133303012-3211011022022132-2322311203230123-2012110230103012"></a>

<a id="canonical-3211312323303010-0221021322102323-3011001102000321-2210113030232131-1313011202001313-1022111002132013-3323131021132020-3330030203101002"></a>

#### `tcp_health_check.send_payload` property

Type: `"string"`. Computed.

Raw bytes sent in the request. Empty payloads imply a connect-only health check. Describes the
encoding of the payload bytes in the payload. Hex encoded payload.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "2048"
  }
}
```

<a id="canonical-2000010223012103-0021100002113232-0202333221122111-1223003032022202-3030103302000301-2011220103210012-0020032300010103-1200022022300232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `udp_icmp_health_check` properties

Breadcrumbs:

- [xcsh_healthcheck](../data-sources/healthcheck.md#canonical-1302112230101233-2202130312331232-0231113312211032-1103323302011301-3303300020333020-2302311211132100-1303020110013123-3330132302312212)
- [Property reference](data-sources--healthcheck--reference--group-001.md#canonical-2233200310100332-2033003002302200-0110210311211131-2311120103323110-0232222133220131-1110332103312313-1221101021321013-1120313110032312)
- udp_icmp_health_check

<a id="canonical-1033303112210301-3232331032212330-3311323033003020-0302030333303210-1020033231323000-2101300110130213-3213002221210220-2102210112011203"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for udp icmp health check.

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
