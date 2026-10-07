---
page_title: "xcsh_endpoint reference"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_endpoint reference."
---

# xcsh_endpoint reference

<a id="canonical-0311233330020132-3211102313010111-3331121211032203-3003300232023323-3312100201200322-2231223210001201-2130221313103221-2202022323120223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-1231112200323301-2223312220101323-0112113120202210-3011323201012220-2310321320210210-1103010223223311-2111322133311112-0132102021033230)
- Property reference

<a id="canonical-0231231030110001-2311032323202002-2331313331012130-1333131221323012-2322321220321101-2230112321102033-0233122030300202-0210322331333210"></a>

### Direct properties for `xcsh_endpoint`

<a id="canonical-3321321110012021-3110021011211113-2010113220303303-1132211032012320-2221300330233311-2203023122101322-0203332003232311-0222113230220023"></a>

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

<a id="canonical-2301132330000111-2220112323113030-1102330122010210-3121320333220120-2000303212313020-3200311200203131-3320111211301212-1130002213133321"></a>

<a id="canonical-3230130231220023-0311022223020101-2110113112030230-1233000201133110-1021031323102033-0310131003213023-2230121321331301-0203231313332301"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the Endpoint.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2032002011112010-0312130003311320-2230331302021002-3112023302220313-2210233331001100-0122130223210303-3101010210013102-1202011101231213"></a>

<a id="canonical-0101123121033121-3201122003122323-2110232201300132-2002211031332303-3211122102102222-0203221321120220-3220033123233232-3000100111013021"></a>

#### `dns_name` property

Type: `"string"`. Computed.

\[OneOf: DNS\_name, DNS\_name\_advanced, ip, service\_info\] Exclusive with \[DNS\_name\_advanced IP
service\_info\] Endpoint's IP address is discovered using DNS name resolution. The name given here
is fully qualified domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

OneOf alternatives in this subsection:

- [dns_name](data-sources--endpoint--reference--group-001.md#canonical-2032002011112010-0312130003311320-2230331302021002-3112023302220313-2210233331001100-0122130223210303-3101010210013102-1202011101231213)
- [dns_name_advanced](data-sources--endpoint--reference--group-001.md#canonical-0332130232120123-2301332122301220-2022010103120201-0333102111213220-0132320311213133-0323001212032020-2303223131211332-3022222231033030)
- [ip](data-sources--endpoint--reference--group-001.md#canonical-3031331332120331-2012203203023020-3310230313021122-0111202121321102-0231011100101301-3222221310233101-2230310332210332-2212221122103100)
- [service_info](data-sources--endpoint--reference--group-001.md#canonical-3000020313023211-2112212320113122-0102313020012112-2023233312101233-3311022203132211-1210210302010201-1112123321301021-3003322013221000)

Select alternatives according to the provider validators above.

- [dns_name_advanced](data-sources--endpoint--reference--group-001.md#canonical-2012123113222233-3233120330213010-2000031000100133-0113201221232212-3303012201020021-3213302333103112-1001012322323113-3100023101323323): complete subsection reference.

<a id="canonical-1201002012310112-1200312102123001-2121002331002131-1312222132323003-1133020330301310-0130232212213232-1231000101022331-2302223032021131"></a>

<a id="canonical-3321203120222010-3301110113023111-2201130211010012-3300323023311130-0110320102231232-3133230220013030-1131323211310020-0300322122200031"></a>

#### `health_check_port` property

Type: `"number"`. Computed.

By default the health check port of an endpoint is the same as the endpoint’s port. This option
provides an alternative health check port. Setting this with a non-zero value allows an endpoint to
have different health check port.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3123033120301223-1332301232323210-1133030133301121-0031132132003201-0311301311301303-1032032000113313-0103230231313333-0320210330313210"></a>

<a id="canonical-0213021233212131-0013000010123023-2323003110310330-3020201220110301-0122301120023221-3323222011221002-3100003303123212-3331302220110132"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3031331332120331-2012203203023020-3310230313021122-0111202121321102-0231011100101301-3222221310233101-2230310332210332-2212221122103100"></a>

<a id="canonical-3121131033032112-2213332222313100-3011201333313001-1310100223230303-1120221120221010-1301203022231122-3131302302132333-3012333100002230"></a>

#### `ip` property

Type: `"string"`. Computed.

Exclusive with \[DNS\_name DNS\_name\_advanced service\_info\] Endpoint is reachable at the given
IPv4/IPv6 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-3233301333203201-2013321321122302-1310311232000011-1010301110212000-3113313332130001-2232002313002201-2012301331331301-0211011231331022"></a>

<a id="canonical-0030311133312313-3100201100000033-0103030200303331-0331113302200300-0110221020203222-1100312003320123-1131232003330011-3311223012113221"></a>

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

<a id="canonical-2222113010310222-1331022323300033-1303130301130210-2130112313320022-1012001000133220-3123032102012032-3321110303201201-1212221232331102"></a>

<a id="canonical-2113310121312300-1301101320010301-3211010210122331-0000301301222133-1202032133311320-1323313101202231-0110021122130332-0033231101303213"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Endpoint.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1003020132303200-2111033111021231-2111122313302332-2011233311101312-1333033212320232-3212101221120331-1023200100020200-2312302023021112"></a>

<a id="canonical-0011232033310322-3003103120333130-2333132300320203-1032120010022320-1032022123331322-3021202031102020-2320221102203310-3223013223333131"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Endpoint exists.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1213010331102112-0230202301132321-3033121012312302-2133121113310310-1200230310323331-3132112322220302-0223020310110223-2110010213123233"></a>

<a id="canonical-1131110201222210-0200020120201333-3010012300011100-0321222102231230-3232010333123120-2310302333120113-2303211123300312-3312110111321303"></a>

#### `port` property

Type: `"number"`. Computed.

Endpoint service is available on this port.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0013200321221112-0322300300313033-0213033210313222-3323221021113201-3001101032020302-2313002030311000-3212212322110103-1310303001212202"></a>

<a id="canonical-2312233301333131-2111103201003132-1301331323123012-1103333200012330-3030232212133010-3211222330230330-3330121000001233-2120023101111000"></a>

#### `protocol` property

Type: `"string"`. Computed.

\[Enum: TCP|UDP\] Protocol. Endpoint protocol. Default is TCP. Both TCP and UDP protocols are
supported. Possible values are \`TCP\`, \`UDP\`.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "TCP",
    "UDP"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"TCP\\\",\\\"UDP\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"TCP\\\",\\\"UDP\\\"]"
  }
}
```

- [service_info](data-sources--endpoint--reference--group-001.md#canonical-2320321032020331-2133003222333220-3030332320301211-3130203301232322-2230031021333133-3303310130333212-1121103320212022-3100021322323222): complete subsection reference.

- [snat_pool](data-sources--endpoint--reference--group-001.md#canonical-1023321130330223-3012133131021311-2313201210232103-2201111323311030-0110100021311312-3133330013320201-0131111221101333-1312112211221301): complete subsection reference.

- [where](data-sources--endpoint--reference--group-001.md#canonical-0221230300321030-0123210220103320-1333323101102300-1113101332130112-3103232113030330-3202211311031020-0100111010231112-1330212113000110): complete subsection reference.

<a id="canonical-1113322320311130-3112211133203231-2021322222022032-0021022002321221-3000003012301331-3211200120021230-1332113101112000-1311221101033131"></a>

### All schema paths for `xcsh_endpoint`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--endpoint--reference--group-001.md#canonical-3321321110012021-3110021011211113-2010113220303303-1132211032012320-2221300330233311-2203023122101322-0203332003232311-0222113230220023) |
| `description` | [description](data-sources--endpoint--reference--group-001.md#canonical-2301132330000111-2220112323113030-1102330122010210-3121320333220120-2000303212313020-3200311200203131-3320111211301212-1130002213133321) |
| `dns_name` | [dns_name](data-sources--endpoint--reference--group-001.md#canonical-2032002011112010-0312130003311320-2230331302021002-3112023302220313-2210233331001100-0122130223210303-3101010210013102-1202011101231213) |
| `dns_name_advanced` | [dns_name_advanced](data-sources--endpoint--reference--group-001.md#canonical-0332130232120123-2301332122301220-2022010103120201-0333102111213220-0132320311213133-0323001212032020-2303223131211332-3022222231033030) |
| `dns_name_advanced.name` | [dns_name_advanced.name](data-sources--endpoint--reference--group-001.md#canonical-3020222133010232-1013112313212202-0212223132323303-3122303312130223-0012200201131223-2303331102332022-3302232300233012-3330330103202233) |
| `dns_name_advanced.refresh_interval` | [dns_name_advanced.refresh_interval](data-sources--endpoint--reference--group-001.md#canonical-3300233001000101-2321123201233321-3323120233210322-2020322001112322-0221311130101121-1322011330321230-3010032311320120-2332233200303120) |
| `health_check_port` | [health_check_port](data-sources--endpoint--reference--group-001.md#canonical-1201002012310112-1200312102123001-2121002331002131-1312222132323003-1133020330301310-0130232212213232-1231000101022331-2302223032021131) |
| `id` | [ID](data-sources--endpoint--reference--group-001.md#canonical-3123033120301223-1332301232323210-1133030133301121-0031132132003201-0311301311301303-1032032000113313-0103230231313333-0320210330313210) |
| `ip` | [ip](data-sources--endpoint--reference--group-001.md#canonical-3031331332120331-2012203203023020-3310230313021122-0111202121321102-0231011100101301-3222221310233101-2230310332210332-2212221122103100) |
| `labels` | [labels](data-sources--endpoint--reference--group-001.md#canonical-3233301333203201-2013321321122302-1310311232000011-1010301110212000-3113313332130001-2232002313002201-2012301331331301-0211011231331022) |
| `name` | [name](data-sources--endpoint--reference--group-001.md#canonical-2222113010310222-1331022323300033-1303130301130210-2130112313320022-1012001000133220-3123032102012032-3321110303201201-1212221232331102) |
| `namespace` | [namespace](data-sources--endpoint--reference--group-001.md#canonical-1003020132303200-2111033111021231-2111122313302332-2011233311101312-1333033212320232-3212101221120331-1023200100020200-2312302023021112) |
| `port` | [port](data-sources--endpoint--reference--group-001.md#canonical-1213010331102112-0230202301132321-3033121012312302-2133121113310310-1200230310323331-3132112322220302-0223020310110223-2110010213123233) |
| `protocol` | [protocol](data-sources--endpoint--reference--group-001.md#canonical-0013200321221112-0322300300313033-0213033210313222-3323221021113201-3001101032020302-2313002030311000-3212212322110103-1310303001212202) |
| `service_info` | [service_info](data-sources--endpoint--reference--group-001.md#canonical-3000020313023211-2112212320113122-0102313020012112-2023233312101233-3311022203132211-1210210302010201-1112123321301021-3003322013221000) |
| `service_info.discovery_type` | [service_info.discovery_type](data-sources--endpoint--reference--group-001.md#canonical-2302120023030331-2011000020211213-1310222013230000-2112030020010113-3301202212230130-2213120002230022-3222303003031133-0312233030310013) |
| `service_info.service_name` | [service_info.service_name](data-sources--endpoint--reference--group-001.md#canonical-1010330202332213-2020223123302200-1231020133121021-3323331301223033-2103020213323233-0203322210203132-2110013320120020-3120223313033123) |
| `service_info.service_selector` | [service_info.service_selector](data-sources--endpoint--reference--group-001.md#canonical-1010301003102313-0020322010122121-0113331131122132-1302330210100022-1321131002112210-3131023321331111-1103221101013312-3210230313333301) |
| `service_info.service_selector.expressions` | [service_info.service_selector.expressions](data-sources--endpoint--reference--group-001.md#canonical-2133222000030332-1010100202331220-2000330012313121-0211131121130000-0133323022311221-0032101303013131-2001232321022312-1202210003200122) |
| `snat_pool` | [snat_pool](data-sources--endpoint--reference--group-001.md#canonical-2100321203102210-0113210120131110-0010103112113011-2011221213102103-3102110202113231-0320200130023301-2031331331233033-1122332211113313) |
| `snat_pool.no_snat_pool` | [snat_pool.no_snat_pool](data-sources--endpoint--reference--group-001.md#canonical-3031003021301220-2203130211231112-1331210011233002-2221231122022312-3331033201221123-3330331100000020-1330120022310223-3020321310010303) |
| `snat_pool.snat_pool` | [snat_pool.snat_pool](data-sources--endpoint--reference--group-001.md#canonical-0230132031233222-2230233302022302-3001112331123303-0211122212333310-2102011103213310-1310132001213100-0320131230132210-1131123233323013) |
| `snat_pool.snat_pool.prefixes` | [snat_pool.snat_pool.prefixes](data-sources--endpoint--reference--group-001.md#canonical-0212212102121322-0130201321111010-2112132122333110-1022003122303120-3113100032232100-1300333032133322-0332303303102020-3003302301122001) |
| `where` | [where](data-sources--endpoint--reference--group-001.md#canonical-2121333000032020-1022131113332203-2312232011303003-1312110132310100-2012022001310121-1333220130230103-3023030001131210-1110231013023211) |
| `where.site` | [where.site](data-sources--endpoint--reference--group-001.md#canonical-0320013312233311-1021011023213201-0321103133231210-1232201233121313-2200031033323213-3300022221302001-0010212231112302-1020000121012222) |
| `where.site.disable_internet_vip` | [where.site.disable_internet_vip](data-sources--endpoint--reference--group-001.md#canonical-3333202212222322-3301110020210323-1133223222023202-3330321320321112-2110113103313033-0011203313230221-0130320320130303-3023133122331202) |
| `where.site.enable_internet_vip` | [where.site.enable_internet_vip](data-sources--endpoint--reference--group-001.md#canonical-0022121012000020-3211011032023103-2320301110031230-2232311313011000-1110200310321011-3220003133210223-0002230121323023-1302001002012013) |
| `where.site.network_type` | [where.site.network_type](data-sources--endpoint--reference--group-001.md#canonical-1101203201311213-1300320322200023-0103000100312312-0130103313010012-1232331020222131-0220230212131311-0021122222132130-1202012000201300) |
| `where.site.ref` | [where.site.ref](data-sources--endpoint--reference--group-001.md#canonical-0001303133233300-1023111121022320-1222133202113212-1330200231333122-2110132020213030-3321013133330301-3201132322231323-2110030112120321) |
| `where.site.ref.kind` | [where.site.ref.kind](data-sources--endpoint--reference--group-001.md#canonical-1232202022333011-0112330101012010-0020301313330302-0133100010122123-1222310203311022-1010102212132031-1001103003002333-1203031011013103) |
| `where.site.ref.name` | [where.site.ref.name](data-sources--endpoint--reference--group-001.md#canonical-2030302020210023-2132233103231233-2121032221112210-0123013201121112-1023303300303100-1230132131113001-3130223101210120-1323011231022230) |
| `where.site.ref.namespace` | [where.site.ref.namespace](data-sources--endpoint--reference--group-001.md#canonical-3112102303011130-0021012200103100-3310011012300132-0132201111300211-1101210323331312-2111021131033233-1333001001333233-1333230202001133) |
| `where.site.ref.tenant` | [where.site.ref.tenant](data-sources--endpoint--reference--group-001.md#canonical-0031321330331032-3032202022203310-3320322021020122-2211001010332312-0012330202301123-1212013013003221-3102213003301032-0230011202021220) |
| `where.site.ref.uid` | [where.site.ref.uid](data-sources--endpoint--reference--group-001.md#canonical-2021131211122230-3202203101230211-1220210333121301-1232213121313132-2031101330313331-3102313221320220-3131102022300013-2203123302133131) |
| `where.virtual_network` | [where.virtual_network](data-sources--endpoint--reference--group-001.md#canonical-1233113102112333-3203200313332003-3311001010230313-3132210121122210-1320031332203010-1211112131111311-1113131133130111-3020212330120010) |
| `where.virtual_network.ref` | [where.virtual_network.ref](data-sources--endpoint--reference--group-001.md#canonical-0203022033231002-2110023313001001-1012203200231102-2112303130132203-2121201111302231-0131331113001212-0331122133101223-3200201002031023) |
| `where.virtual_network.ref.kind` | [where.virtual_network.ref.kind](data-sources--endpoint--reference--group-001.md#canonical-2310122311221112-2133330002013300-3301323203320333-3221320022001100-3102310023313013-2211001332011121-1311132010221331-2101311012010213) |
| `where.virtual_network.ref.name` | [where.virtual_network.ref.name](data-sources--endpoint--reference--group-001.md#canonical-3001330013300101-1011021131333200-0330011010330220-3213232010213123-0022230302311211-3033001122031313-2210023222112200-3022111331201330) |
| `where.virtual_network.ref.namespace` | [where.virtual_network.ref.namespace](data-sources--endpoint--reference--group-001.md#canonical-3333213010200013-3113131311320330-0331122323322102-0131020312111111-3303230022203221-1333322310311121-2230221301130311-1022311203222332) |
| `where.virtual_network.ref.tenant` | [where.virtual_network.ref.tenant](data-sources--endpoint--reference--group-001.md#canonical-2333200310000132-2321300101322321-3121233103302320-3033212323332010-2232221101300011-2303013323231131-1001311321222010-0201330330021302) |
| `where.virtual_network.ref.uid` | [where.virtual_network.ref.uid](data-sources--endpoint--reference--group-001.md#canonical-2033023230212130-3111322002133313-2222231320302102-0013111132201301-1013123033202000-0020230020123200-0330000033322230-2133323201200303) |
| `where.virtual_site` | [where.virtual_site](data-sources--endpoint--reference--group-001.md#canonical-3123032122111323-3300021020011002-0210132202321132-2022132123310321-0003021303312201-0313100122133012-0313332123233021-2010123113013303) |
| `where.virtual_site.disable_internet_vip` | [where.virtual_site.disable_internet_vip](data-sources--endpoint--reference--group-001.md#canonical-1220133220313000-2032101220230102-3222110231223000-1112020112103332-2111223011230033-3333023031011133-3232213212330300-2121133321320220) |
| `where.virtual_site.enable_internet_vip` | [where.virtual_site.enable_internet_vip](data-sources--endpoint--reference--group-001.md#canonical-0220030200121213-2111302332022002-0023311232030213-2210131221120011-1110310222111100-2120300102122213-0112130112222220-3020221103302013) |
| `where.virtual_site.network_type` | [where.virtual_site.network_type](data-sources--endpoint--reference--group-001.md#canonical-3100033032232300-2222120220132032-1132312030113032-3233103131120202-2201133132013122-3210202120112312-3032133011011221-0301132200113232) |
| `where.virtual_site.ref` | [where.virtual_site.ref](data-sources--endpoint--reference--group-001.md#canonical-0132213111131032-3121201312203322-3111331200313023-2211302023321211-2211132223020300-1313212333001330-2130020000023330-3321231312222320) |
| `where.virtual_site.ref.kind` | [where.virtual_site.ref.kind](data-sources--endpoint--reference--group-001.md#canonical-2210320101120303-1102101113131223-3013022113233231-2031011122022222-2331313023132101-3023330203000213-2321220223031002-1302031022113230) |
| `where.virtual_site.ref.name` | [where.virtual_site.ref.name](data-sources--endpoint--reference--group-001.md#canonical-1331220123211133-1300101231300230-3103323223303321-3020000330332022-3203133203130122-2320101220221303-0223311112320322-0300133131320121) |
| `where.virtual_site.ref.namespace` | [where.virtual_site.ref.namespace](data-sources--endpoint--reference--group-001.md#canonical-3021230023120201-0000010100322023-3313302022313123-1310302220223122-2220031021331313-1201301333330302-1022032003302012-3100003221122102) |
| `where.virtual_site.ref.tenant` | [where.virtual_site.ref.tenant](data-sources--endpoint--reference--group-001.md#canonical-2200130001012330-2303130312231023-3130101111221113-1020010031133001-3020213302130023-3030320011231032-2030220222020331-2232120100132012) |
| `where.virtual_site.ref.uid` | [where.virtual_site.ref.uid](data-sources--endpoint--reference--group-001.md#canonical-1110033211210032-1301113113222133-2001132320321101-2030130221221023-2303000032212100-3120321312001300-3012312311010303-3023132333112010) |

<a id="canonical-2012123113222233-3233120330213010-2000031000100133-0113201221232212-3303012201020021-3213302333103112-1001012322323113-3100023101323323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dns_name_advanced` properties

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-1231112200323301-2223312220101323-0112113120202210-3011323201012220-2310321320210210-1103010223223311-2111322133311112-0132102021033230)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-0311233330020132-3211102313010111-3331121211032203-3003300232023323-3312100201200322-2231223210001201-2130221313103221-2202022323120223)
- dns_name_advanced

<a id="canonical-0332130232120123-2301332122301220-2022010103120201-0333102111213220-0132320311213133-0323001212032020-2303223131211332-3022222231033030"></a>

Type: `"single"`. Computed.

Specifies name and TTL used for DNS resolution.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ttl_choice": "[\"refresh_interval\"]"
}
```

<a id="canonical-3323001301233012-2211131133103112-3322033103210310-1211101333212132-0330331132222222-1331231303200131-0122022212021121-3120120133313131"></a>

### Direct properties for `dns_name_advanced`

<a id="canonical-3020222133010232-1013112313212202-0212223132323303-3122303312130223-0012200201131223-2303331102332022-3302232300233012-3330330103202233"></a>

#### `dns_name_advanced.name` property

Type: `"string"`. Computed.

Endpoint's IP address is discovered using DNS name resolution. The name given here is fully
qualified domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "format": "hostname",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3300233001000101-2321123201233321-3323120233210322-2020322001112322-0221311130101121-1322011330321230-3010032311320120-2332233200303120"></a>

<a id="canonical-1200132122123320-0230013221203102-0001022200210100-2011320132123201-0132130123033311-0030210310122230-1330012333032231-0322301112332011"></a>

#### `dns_name_advanced.refresh_interval` property

Type: `"number"`. Computed.

Exclusive with \[\] Interval for DNS refresh in seconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 604800,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "minimum": 10
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "10",
    "ves.io.schema.rules.uint32.lte": "604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "10",
    "ves.io.schema.rules.uint32.lte": "604800"
  }
}
```

<a id="canonical-2320321032020331-2133003222333220-3030332320301211-3130203301232322-2230031021333133-3303310130333212-1121103320212022-3100021322323222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service_info` properties

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-1231112200323301-2223312220101323-0112113120202210-3011323201012220-2310321320210210-1103010223223311-2111322133311112-0132102021033230)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-0311233330020132-3211102313010111-3331121211032203-3003300232023323-3312100201200322-2231223210001201-2130221313103221-2202022323120223)
- service_info

<a id="canonical-3000020313023211-2112212320113122-0102313020012112-2023233312101233-3311022203132211-1210210302010201-1112123321301021-3003322013221000"></a>

Type: `"single"`. Computed.

Specifies whether endpoint service is discovered by name or labels.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-service_info": "[\"service_name\",\"service_selector\"]"
}
```

<a id="canonical-0310233131001010-3212130032120321-3111330230322322-0200212113100211-2133213201003131-1221313030212022-1031232332330002-2102000023230000"></a>

### Direct properties for `service_info`

<a id="canonical-2302120023030331-2011000020211213-1310222013230000-2112030020010113-3301202212230130-2213120002230022-3222303003031133-0312233030310013"></a>

#### `service_info.discovery_type` property

Type: `"string"`. Computed.

\[Enum: INVALID\_DISCOVERY|K8S|CONSUL|CLASSIC\_BIGIP|THIRD\_PARTY|NGINX\_ONE\] Specifies the type of
discovery Invalid Discovery mechanism Discover from Kubernetes cluster Discover from Consul service
Discover from Classic BIG-IP Clusters Discover for Third Party Application Discover from NGINX One.
Possible values are \`INVALID\_DISCOVERY\`, \`K8S\`, \`CONSUL\`, \`CLASSIC\_BIGIP\`,
\`THIRD\_PARTY\`, \`NGINX\_ONE\`. Defaults to \`INVALID\_DISCOVERY\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "INVALID_DISCOVERY",
  "enum": [
    "INVALID_DISCOVERY",
    "K8S",
    "CONSUL",
    "CLASSIC_BIGIP",
    "THIRD_PARTY",
    "NGINX_ONE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1010330202332213-2020223123302200-1231020133121021-3323331301223033-2103020213323233-0203322210203132-2110013320120020-3120223313033123"></a>

<a id="canonical-3023031210113030-1322021102122211-2110221213120331-0030221321010201-0322132201220331-2203032201300313-3021200321112030-3102211123311303"></a>

#### `service_info.service_name` property

Type: `"string"`. Computed.

Exclusive with \[service\_selector\] Name of the service to discover with an optional namespace and
cluster identifier. The format is service\_name.namespace\_name:cluster\_identifier for K8s and
service\_name:cluster\_identifier for Consul Endpoint will be discovered in all discovery objects
where the cluster identifier matches. If cluster identifier is not specified then discovery will be
done in all discovery objects of the site.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

- [service_selector](data-sources--endpoint--reference--group-001.md#canonical-1003130010111223-3211011322331202-0102103333321011-3031122032113103-2121310131332211-0220310311101110-3011223301022202-0232210211211221): complete subsection reference.

<a id="canonical-1003130010111223-3211011322331202-0102103333321011-3031122032113103-2121310131332211-0220310311101110-3011223301022202-0232210211211221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `service_info.service_selector` properties

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-1231112200323301-2223312220101323-0112113120202210-3011323201012220-2310321320210210-1103010223223311-2111322133311112-0132102021033230)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-0311233330020132-3211102313010111-3331121211032203-3003300232023323-3312100201200322-2231223210001201-2130221313103221-2202022323120223)
- [service_info](data-sources--endpoint--reference--group-001.md#canonical-2320321032020331-2133003222333220-3030332320301211-3130203301232322-2230031021333133-3303310130333212-1121103320212022-3100021322323222)
- service_info.service_selector

<a id="canonical-1010301003102313-0020322010122121-0113331131122132-1302330210100022-1321131002112210-3131023321331111-1103221101013312-3210230313333301"></a>

Type: `"single"`. Computed.

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

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

<a id="canonical-1333112120313300-1030100212221111-3221213323002222-0010112310113002-2100100123131222-2332322101030301-3023223002023111-0021033113221300"></a>

### Direct properties for `service_info.service_selector`

<a id="canonical-2133222000030332-1010100202331220-2000330012313121-0211131121130000-0133323022311221-0032101303013131-2001232321022312-1202210003200122"></a>

#### `service_info.service_selector.expressions` property

Type: `["list", "string"]`. Computed.

Expressions contains the Kubernetes style label expression for selections.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-1023321130330223-3012133131021311-2313201210232103-2201111323311030-0110100021311312-3133330013320201-0131111221101333-1312112211221301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `snat_pool` properties

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-1231112200323301-2223312220101323-0112113120202210-3011323201012220-2310321320210210-1103010223223311-2111322133311112-0132102021033230)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-0311233330020132-3211102313010111-3331121211032203-3003300232023323-3312100201200322-2231223210001201-2130221313103221-2202022323120223)
- snat_pool

<a id="canonical-2100321203102210-0113210120131110-0010103112113011-2011221213102103-3102110202113231-0320200130023301-2031331331233033-1122332211113313"></a>

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

<a id="canonical-3302230122011122-0302323132033212-0121103320021233-3131120101232020-0132112300113313-1000122131003130-1012331212331322-0100331230010121"></a>

### Direct properties for `snat_pool`

- [no_snat_pool](data-sources--endpoint--reference--group-001.md#canonical-3302031112312132-1033312321310003-2210120022312202-1233211321032013-1023231132120320-3221210000333202-1013331213020122-0202221221233021): complete subsection reference.

- [snat_pool](data-sources--endpoint--reference--group-001.md#canonical-0033030131220320-2320010312023110-0111110003101122-3012111122313310-1331101312313310-2333001033202013-1030332231000201-1320003031332130): complete subsection reference.

<a id="canonical-3302031112312132-1033312321310003-2210120022312202-1233211321032013-1023231132120320-3221210000333202-1013331213020122-0202221221233021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `snat_pool.no_snat_pool` properties

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-1231112200323301-2223312220101323-0112113120202210-3011323201012220-2310321320210210-1103010223223311-2111322133311112-0132102021033230)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-0311233330020132-3211102313010111-3331121211032203-3003300232023323-3312100201200322-2231223210001201-2130221313103221-2202022323120223)
- [snat_pool](data-sources--endpoint--reference--group-001.md#canonical-1023321130330223-3012133131021311-2313201210232103-2201111323311030-0110100021311312-3133330013320201-0131111221101333-1312112211221301)
- snat_pool.no_snat_pool

<a id="canonical-3031003021301220-2203130211231112-1331210011233002-2221231122022312-3331033201221123-3330331100000020-1330120022310223-3020321310010303"></a>

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

<a id="canonical-0033030131220320-2320010312023110-0111110003101122-3012111122313310-1331101312313310-2333001033202013-1030332231000201-1320003031332130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `snat_pool.snat_pool` properties

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-1231112200323301-2223312220101323-0112113120202210-3011323201012220-2310321320210210-1103010223223311-2111322133311112-0132102021033230)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-0311233330020132-3211102313010111-3331121211032203-3003300232023323-3312100201200322-2231223210001201-2130221313103221-2202022323120223)
- [snat_pool](data-sources--endpoint--reference--group-001.md#canonical-1023321130330223-3012133131021311-2313201210232103-2201111323311030-0110100021311312-3133330013320201-0131111221101333-1312112211221301)
- snat_pool.snat_pool

<a id="canonical-0230132031233222-2230233302022302-3001112331123303-0211122212333310-2102011103213310-1310132001213100-0320131230132210-1131123233323013"></a>

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

<a id="canonical-3031121113211133-3300331332333031-3313213100102303-1110321213133233-2000302311031000-2221323030121110-2130101201312220-1233032132301311"></a>

### Direct properties for `snat_pool.snat_pool`

<a id="canonical-0212212102121322-0130201321111010-2112132122333110-1022003122303120-3113100032232100-1300333032133322-0332303303102020-3003302301122001"></a>

#### `snat_pool.snat_pool.prefixes` property

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0221230300321030-0123210220103320-1333323101102300-1113101332130112-3103232113030330-3202211311031020-0100111010231112-1330212113000110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where` properties

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-1231112200323301-2223312220101323-0112113120202210-3011323201012220-2310321320210210-1103010223223311-2111322133311112-0132102021033230)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-0311233330020132-3211102313010111-3331121211032203-3003300232023323-3312100201200322-2231223210001201-2130221313103221-2202022323120223)
- where

<a id="canonical-2121333000032020-1022131113332203-2312232011303003-1312110132310100-2012022001310121-1333220130230103-3023030001131210-1110231013023211"></a>

Type: `"single"`. Computed.

NetworkSiteRefSelector defines a union of reference to site or reference to virtual\_network or
reference to virtual\_site It is used to determine virtual network using following rules \* Direct
reference to virtual\_network object \* Site local network when referring to site object \* All site
local networks for sites selected by referring to virtual\_site object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ref_or_selector": "[\"site\",\"virtual_network\",\"virtual_site\"]"
}
```

<a id="canonical-0233030032313200-3101120332321013-1113003203330300-2300122121133132-0022213121311020-3010012303300321-3322222022231021-3212010200330232"></a>

### Direct properties for `where`

- [site](data-sources--endpoint--reference--group-001.md#canonical-2013023103301120-0233000012101023-3302323120100310-3121300103311102-2002121301130012-2101333220032202-1123211111022122-2111220201032300): complete subsection reference.

- [virtual_network](data-sources--endpoint--reference--group-001.md#canonical-0200122230021200-0212011333032223-1231122012211012-0230111211311230-3231020311310323-3112130130123100-2021233003100303-3120302132022221): complete subsection reference.

- [virtual_site](data-sources--endpoint--reference--group-001.md#canonical-3322032102011232-0100112103230000-0220221203331231-2320213211123223-1102212100122022-1132302102113023-3031033002013200-1002103320102200): complete subsection reference.

<a id="canonical-2013023103301120-0233000012101023-3302323120100310-3121300103311102-2002121301130012-2101333220032202-1123211111022122-2111220201032300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.site` properties

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-1231112200323301-2223312220101323-0112113120202210-3011323201012220-2310321320210210-1103010223223311-2111322133311112-0132102021033230)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-0311233330020132-3211102313010111-3331121211032203-3003300232023323-3312100201200322-2231223210001201-2130221313103221-2202022323120223)
- [where](data-sources--endpoint--reference--group-001.md#canonical-0221230300321030-0123210220103320-1333323101102300-1113101332130112-3103232113030330-3202211311031020-0100111010231112-1330212113000110)
- where.site

<a id="canonical-0320013312233311-1021011023213201-0321103133231210-1232201233121313-2200031033323213-3300022221302001-0010212231112302-1020000121012222"></a>

Type: `"single"`. Computed.

This specifies a direct reference to a site configuration object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]"
}
```

<a id="canonical-3112023313021123-2313122022331231-3022031013333111-3203110003120322-1000003230300213-0212321022020030-0332330200220331-0323123211211001"></a>

### Direct properties for `where.site`

- [disable_internet_vip](data-sources--endpoint--reference--group-001.md#canonical-3102122321231130-1310303230120130-1203320123031130-0203100333121202-1122103013021201-0010103101202002-3222131102112301-3030101312322131): complete subsection reference.

- [enable_internet_vip](data-sources--endpoint--reference--group-001.md#canonical-3231222310002010-0302033211120221-1100331130301232-1300231112111300-3223130130233113-1321302101203312-0202321132330030-2033013123030331): complete subsection reference.

<a id="canonical-1101203201311213-1300320322200023-0103000100312312-0130103313010012-1232331020222131-0220230212131311-0021122222132130-1202012000201300"></a>

<a id="canonical-0200133022210202-3202113302220332-2201123203021202-2121233202302211-0312230321112011-3313131213303120-1220333002033103-1323000201212030"></a>

#### `where.site.network_type` property

Type: `"string"`. Computed.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Additional upstream details:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected. Constraints: There can be atmost one virtual network of
this type in a given site. This network type is supported on CE sites. This network is created
automatically and present on all sites Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE
is a private network inside site. It is a secure network and is not connected to public network.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected. Constraints: There can be atmost one virtual network of
this type in a given site. This network type is supported on CE sites. This network is created
during provisioning of site User defined per-site virtual network. Scope of this virtual network is
limited to the site. This is not yet supported Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC
directly connects to the public internet. Virtual-network of this type is local to every site. Two
virtual networks of this type on different sites are neither related nor connected. Constraints:
There can be atmost one virtual network of this type in a given site. This network type is supported
on RE sites only It is an internally created by the system. They must not be created by user Virtual
Networks with global scope across different sites in F5XC domain. An example global virtual-network
called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [ref](data-sources--endpoint--reference--group-001.md#canonical-2011120121300023-2033201132203033-1110110100012300-1001120221001321-1131300010023232-2023120012103313-0122013032312321-3200120133111221): complete subsection reference.

<a id="canonical-3102122321231130-1310303230120130-1203320123031130-0203100333121202-1122103013021201-0010103101202002-3222131102112301-3030101312322131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.site.disable_internet_vip` properties

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-1231112200323301-2223312220101323-0112113120202210-3011323201012220-2310321320210210-1103010223223311-2111322133311112-0132102021033230)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-0311233330020132-3211102313010111-3331121211032203-3003300232023323-3312100201200322-2231223210001201-2130221313103221-2202022323120223)
- [where](data-sources--endpoint--reference--group-001.md#canonical-0221230300321030-0123210220103320-1333323101102300-1113101332130112-3103232113030330-3202211311031020-0100111010231112-1330212113000110)
- [where.site](data-sources--endpoint--reference--group-001.md#canonical-2013023103301120-0233000012101023-3302323120100310-3121300103311102-2002121301130012-2101333220032202-1123211111022122-2111220201032300)
- where.site.disable_internet_vip

<a id="canonical-3333202212222322-3301110020210323-1133223222023202-3330321320321112-2110113103313033-0011203313230221-0130320320130303-3023133122331202"></a>

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

<a id="canonical-3231222310002010-0302033211120221-1100331130301232-1300231112111300-3223130130233113-1321302101203312-0202321132330030-2033013123030331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.site.enable_internet_vip` properties

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-1231112200323301-2223312220101323-0112113120202210-3011323201012220-2310321320210210-1103010223223311-2111322133311112-0132102021033230)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-0311233330020132-3211102313010111-3331121211032203-3003300232023323-3312100201200322-2231223210001201-2130221313103221-2202022323120223)
- [where](data-sources--endpoint--reference--group-001.md#canonical-0221230300321030-0123210220103320-1333323101102300-1113101332130112-3103232113030330-3202211311031020-0100111010231112-1330212113000110)
- [where.site](data-sources--endpoint--reference--group-001.md#canonical-2013023103301120-0233000012101023-3302323120100310-3121300103311102-2002121301130012-2101333220032202-1123211111022122-2111220201032300)
- where.site.enable_internet_vip

<a id="canonical-0022121012000020-3211011032023103-2320301110031230-2232311313011000-1110200310321011-3220003133210223-0002230121323023-1302001002012013"></a>

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

<a id="canonical-2011120121300023-2033201132203033-1110110100012300-1001120221001321-1131300010023232-2023120012103313-0122013032312321-3200120133111221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.site.ref` properties

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-1231112200323301-2223312220101323-0112113120202210-3011323201012220-2310321320210210-1103010223223311-2111322133311112-0132102021033230)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-0311233330020132-3211102313010111-3331121211032203-3003300232023323-3312100201200322-2231223210001201-2130221313103221-2202022323120223)
- [where](data-sources--endpoint--reference--group-001.md#canonical-0221230300321030-0123210220103320-1333323101102300-1113101332130112-3103232113030330-3202211311031020-0100111010231112-1330212113000110)
- [where.site](data-sources--endpoint--reference--group-001.md#canonical-2013023103301120-0233000012101023-3302323120100310-3121300103311102-2002121301130012-2101333220032202-1123211111022122-2111220201032300)
- where.site.ref

<a id="canonical-0001303133233300-1023111121022320-1222133202113212-1330200231333122-2110132020213030-3321013133330301-3201132322231323-2110030112120321"></a>

Type: `"list"`. Computed.

Reference. A site direct reference.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-2201202021113313-1012000331212231-3021320032031131-0123301110220103-3131312221110221-0300201221111100-1322220303103211-2032030231133212"></a>

### Direct properties for `where.site.ref`

<a id="canonical-1232202022333011-0112330101012010-0020301313330302-0133100010122123-1222310203311022-1010102212132031-1001103003002333-1203031011013103"></a>

#### `where.site.ref.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2030302020210023-2132233103231233-2121032221112210-0123013201121112-1023303300303100-1230132131113001-3130223101210120-1323011231022230"></a>

<a id="canonical-3002123312233030-3120012222000222-2111222300133022-1321303112121101-3200011133332000-0231121203332022-3010210223321101-2123320233232320"></a>

#### `where.site.ref.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3112102303011130-0021012200103100-3310011012300132-0132201111300211-1101210323331312-2111021131033233-1333001001333233-1333230202001133"></a>

<a id="canonical-0321312311300313-0231332302012033-3300320220202112-2121203321330012-3001302000220023-3232321130001101-1131303100222102-2110001322013321"></a>

#### `where.site.ref.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0031321330331032-3032202022203310-3320322021020122-2211001010332312-0012330202301123-1212013013003221-3102213003301032-0230011202021220"></a>

<a id="canonical-0230313100013032-0031332302131113-2231000010213333-3122332021312100-1003211203323220-2133023133110030-0133220221123303-1320320232112120"></a>

#### `where.site.ref.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2021131211122230-3202203101230211-1220210333121301-1232213121313132-2031101330313331-3102313221320220-3131102022300013-2203123302133131"></a>

<a id="canonical-2203333102012301-2103131133231233-3112031111003001-0131231323130121-0322220030212130-0311212233220301-2212021303110320-2300123301110201"></a>

#### `where.site.ref.uid` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0200122230021200-0212011333032223-1231122012211012-0230111211311230-3231020311310323-3112130130123100-2021233003100303-3120302132022221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_network` properties

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-1231112200323301-2223312220101323-0112113120202210-3011323201012220-2310321320210210-1103010223223311-2111322133311112-0132102021033230)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-0311233330020132-3211102313010111-3331121211032203-3003300232023323-3312100201200322-2231223210001201-2130221313103221-2202022323120223)
- [where](data-sources--endpoint--reference--group-001.md#canonical-0221230300321030-0123210220103320-1333323101102300-1113101332130112-3103232113030330-3202211311031020-0100111010231112-1330212113000110)
- where.virtual_network

<a id="canonical-1233113102112333-3203200313332003-3311001010230313-3132210121122210-1320031332203010-1211112131111311-1113131133130111-3020212330120010"></a>

Type: `"single"`. Computed.

This specifies a direct reference to a network configuration object.

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

<a id="canonical-3013230021303200-1201131000132020-1211003113310200-2203133321330301-3020313232303322-0310223103122113-0202011230211332-2210320002300302"></a>

### Direct properties for `where.virtual_network`

- [ref](data-sources--endpoint--reference--group-001.md#canonical-1303321222313333-2200331112230130-3212120302121012-0203331102231233-1213023120331010-2001200300312113-3121120020113020-1321131310333110): complete subsection reference.

<a id="canonical-1303321222313333-2200331112230130-3212120302121012-0203331102231233-1213023120331010-2001200300312113-3121120020113020-1321131310333110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_network.ref` properties

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-1231112200323301-2223312220101323-0112113120202210-3011323201012220-2310321320210210-1103010223223311-2111322133311112-0132102021033230)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-0311233330020132-3211102313010111-3331121211032203-3003300232023323-3312100201200322-2231223210001201-2130221313103221-2202022323120223)
- [where](data-sources--endpoint--reference--group-001.md#canonical-0221230300321030-0123210220103320-1333323101102300-1113101332130112-3103232113030330-3202211311031020-0100111010231112-1330212113000110)
- [where.virtual_network](data-sources--endpoint--reference--group-001.md#canonical-0200122230021200-0212011333032223-1231122012211012-0230111211311230-3231020311310323-3112130130123100-2021233003100303-3120302132022221)
- where.virtual_network.ref

<a id="canonical-0203022033231002-2110023313001001-1012203200231102-2112303130132203-2121201111302231-0131331113001212-0331122133101223-3200201002031023"></a>

Type: `"list"`. Computed.

Reference. A virtual network direct reference.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-1313123113223020-0101122003231330-1032100221111000-1311121320211220-0202230232032102-2231321133213113-1032331030200332-0223233023202230"></a>

### Direct properties for `where.virtual_network.ref`

<a id="canonical-2310122311221112-2133330002013300-3301323203320333-3221320022001100-3102310023313013-2211001332011121-1311132010221331-2101311012010213"></a>

#### `where.virtual_network.ref.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3001330013300101-1011021131333200-0330011010330220-3213232010213123-0022230302311211-3033001122031313-2210023222112200-3022111331201330"></a>

<a id="canonical-2113303230222122-3303220330232332-2103211330102123-1013301033101020-2322130330100221-2010330331101202-3101012323112313-1012100322030312"></a>

#### `where.virtual_network.ref.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3333213010200013-3113131311320330-0331122323322102-0131020312111111-3303230022203221-1333322310311121-2230221301130311-1022311203222332"></a>

<a id="canonical-3321322201331122-2203023123122222-3210223013001122-2210030233031323-2021211030130321-1120100000001213-3212001023333001-2221232202120132"></a>

#### `where.virtual_network.ref.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2333200310000132-2321300101322321-3121233103302320-3033212323332010-2232221101300011-2303013323231131-1001311321222010-0201330330021302"></a>

<a id="canonical-1323122312310000-2311220123203013-0100233001010131-0200031321110231-3203112022213223-3210000123101132-0112023210020021-3032323331301221"></a>

#### `where.virtual_network.ref.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2033023230212130-3111322002133313-2222231320302102-0013111132201301-1013123033202000-0020230020123200-0330000033322230-2133323201200303"></a>

<a id="canonical-0133033032220012-0333223022130002-2221010133101223-3222011331221013-3010100112202032-3121203331023110-1211130323120233-2210103330202211"></a>

#### `where.virtual_network.ref.uid` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3322032102011232-0100112103230000-0220221203331231-2320213211123223-1102212100122022-1132302102113023-3031033002013200-1002103320102200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_site` properties

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-1231112200323301-2223312220101323-0112113120202210-3011323201012220-2310321320210210-1103010223223311-2111322133311112-0132102021033230)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-0311233330020132-3211102313010111-3331121211032203-3003300232023323-3312100201200322-2231223210001201-2130221313103221-2202022323120223)
- [where](data-sources--endpoint--reference--group-001.md#canonical-0221230300321030-0123210220103320-1333323101102300-1113101332130112-3103232113030330-3202211311031020-0100111010231112-1330212113000110)
- where.virtual_site

<a id="canonical-3123032122111323-3300021020011002-0210132202321132-2022132123310321-0003021303312201-0313100122133012-0313332123233021-2010123113013303"></a>

Type: `"single"`. Computed.

Virtual Site. A reference to virtual\_site object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-internet_vip_choice": "[\"disable_internet_vip\",\"enable_internet_vip\"]"
}
```

<a id="canonical-1233331200200220-1220020200330233-1102302120032230-0200022112322312-2031032330211010-2200100011022230-3333013332212123-2112101331031101"></a>

### Direct properties for `where.virtual_site`

- [disable_internet_vip](data-sources--endpoint--reference--group-001.md#canonical-2223030302220123-0030202221021303-3023332012213133-1121210022120032-2003213021113113-3110221103213133-0202022232200110-0311231210130313): complete subsection reference.

- [enable_internet_vip](data-sources--endpoint--reference--group-001.md#canonical-1111320021233022-3321201202321110-2300312202022013-3210233303331201-0001330010331332-0231132111003100-2110302300201032-0101300131120231): complete subsection reference.

<a id="canonical-3100033032232300-2222120220132032-1132312030113032-3233103131120202-2201133132013122-3210202120112312-3032133011011221-0301132200113232"></a>

<a id="canonical-0303131323200033-3130323112203102-2003310300213103-2122231221123003-0212331121123320-2201330001222023-0110032223101102-0200000102322213"></a>

#### `where.virtual_site.network_type` property

Type: `"string"`. Computed.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Additional upstream details:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected. Constraints: There can be atmost one virtual network of
this type in a given site. This network type is supported on CE sites. This network is created
automatically and present on all sites Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE
is a private network inside site. It is a secure network and is not connected to public network.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected. Constraints: There can be atmost one virtual network of
this type in a given site. This network type is supported on CE sites. This network is created
during provisioning of site User defined per-site virtual network. Scope of this virtual network is
limited to the site. This is not yet supported Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC
directly connects to the public internet. Virtual-network of this type is local to every site. Two
virtual networks of this type on different sites are neither related nor connected. Constraints:
There can be atmost one virtual network of this type in a given site. This network type is supported
on RE sites only It is an internally created by the system. They must not be created by user Virtual
Networks with global scope across different sites in F5XC domain. An example global virtual-network
called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [ref](data-sources--endpoint--reference--group-001.md#canonical-2022122201222210-2213131322210201-3021100033302022-1200322220310323-1323221020031330-0020031110113301-3112223002102003-2321110323102301): complete subsection reference.

<a id="canonical-2223030302220123-0030202221021303-3023332012213133-1121210022120032-2003213021113113-3110221103213133-0202022232200110-0311231210130313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_site.disable_internet_vip` properties

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-1231112200323301-2223312220101323-0112113120202210-3011323201012220-2310321320210210-1103010223223311-2111322133311112-0132102021033230)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-0311233330020132-3211102313010111-3331121211032203-3003300232023323-3312100201200322-2231223210001201-2130221313103221-2202022323120223)
- [where](data-sources--endpoint--reference--group-001.md#canonical-0221230300321030-0123210220103320-1333323101102300-1113101332130112-3103232113030330-3202211311031020-0100111010231112-1330212113000110)
- [where.virtual_site](data-sources--endpoint--reference--group-001.md#canonical-3322032102011232-0100112103230000-0220221203331231-2320213211123223-1102212100122022-1132302102113023-3031033002013200-1002103320102200)
- where.virtual_site.disable_internet_vip

<a id="canonical-1220133220313000-2032101220230102-3222110231223000-1112020112103332-2111223011230033-3333023031011133-3232213212330300-2121133321320220"></a>

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

<a id="canonical-1111320021233022-3321201202321110-2300312202022013-3210233303331201-0001330010331332-0231132111003100-2110302300201032-0101300131120231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_site.enable_internet_vip` properties

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-1231112200323301-2223312220101323-0112113120202210-3011323201012220-2310321320210210-1103010223223311-2111322133311112-0132102021033230)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-0311233330020132-3211102313010111-3331121211032203-3003300232023323-3312100201200322-2231223210001201-2130221313103221-2202022323120223)
- [where](data-sources--endpoint--reference--group-001.md#canonical-0221230300321030-0123210220103320-1333323101102300-1113101332130112-3103232113030330-3202211311031020-0100111010231112-1330212113000110)
- [where.virtual_site](data-sources--endpoint--reference--group-001.md#canonical-3322032102011232-0100112103230000-0220221203331231-2320213211123223-1102212100122022-1132302102113023-3031033002013200-1002103320102200)
- where.virtual_site.enable_internet_vip

<a id="canonical-0220030200121213-2111302332022002-0023311232030213-2210131221120011-1110310222111100-2120300102122213-0112130112222220-3020221103302013"></a>

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

<a id="canonical-2022122201222210-2213131322210201-3021100033302022-1200322220310323-1323221020031330-0020031110113301-3112223002102003-2321110323102301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_site.ref` properties

Breadcrumbs:

- [xcsh_endpoint](../data-sources/endpoint.md#canonical-1231112200323301-2223312220101323-0112113120202210-3011323201012220-2310321320210210-1103010223223311-2111322133311112-0132102021033230)
- [Property reference](data-sources--endpoint--reference--group-001.md#canonical-0311233330020132-3211102313010111-3331121211032203-3003300232023323-3312100201200322-2231223210001201-2130221313103221-2202022323120223)
- [where](data-sources--endpoint--reference--group-001.md#canonical-0221230300321030-0123210220103320-1333323101102300-1113101332130112-3103232113030330-3202211311031020-0100111010231112-1330212113000110)
- [where.virtual_site](data-sources--endpoint--reference--group-001.md#canonical-3322032102011232-0100112103230000-0220221203331231-2320213211123223-1102212100122022-1132302102113023-3031033002013200-1002103320102200)
- where.virtual_site.ref

<a id="canonical-0132213111131032-3121201312203322-3111331200313023-2211302023321211-2211132223020300-1313212333001330-2130020000023330-3321231312222320"></a>

Type: `"list"`. Computed.

Reference. A virtual\_site direct reference.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-0212000213220220-1101113312113011-1212103010131210-0300122012221301-3131330121102103-2110121001001123-3022200111133223-2203230212232031"></a>

### Direct properties for `where.virtual_site.ref`

<a id="canonical-2210320101120303-1102101113131223-3013022113233231-2031011122022222-2331313023132101-3023330203000213-2321220223031002-1302031022113230"></a>

#### `where.virtual_site.ref.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1331220123211133-1300101231300230-3103323223303321-3020000330332022-3203133203130122-2320101220221303-0223311112320322-0300133131320121"></a>

<a id="canonical-1022031221220111-1230133210222213-3231031301210022-1320121331230200-3112201013002120-3002031211322222-3121231232222302-3202012232012011"></a>

#### `where.virtual_site.ref.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3021230023120201-0000010100322023-3313302022313123-1310302220223122-2220031021331313-1201301333330302-1022032003302012-3100003221122102"></a>

<a id="canonical-0333021000113223-2202133120013122-1230000130003323-3330000021021333-2231130002013012-3330131133103131-3033011221020033-0130121123013200"></a>

#### `where.virtual_site.ref.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2200130001012330-2303130312231023-3130101111221113-1020010031133001-3020213302130023-3030320011231032-2030220222020331-2232120100132012"></a>

<a id="canonical-3101202032013231-2203111100221312-1313333320110130-3030320212031020-1002100011132102-1202312013210110-0220031100012020-0022231003332022"></a>

#### `where.virtual_site.ref.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1110033211210032-1301113113222133-2001132320321101-2030130221221023-2303000032212100-3120321312001300-3012312311010303-3023132333112010"></a>

<a id="canonical-2311333010322113-2303100102031030-2332302332331103-3011021211310231-1122223211323201-3233003323032003-0102233013330213-0013211122230222"></a>

#### `where.virtual_site.ref.uid` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
