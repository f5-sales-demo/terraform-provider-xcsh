---
page_title: "xcsh_advertise_policy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_advertise_policy reference."
---

# xcsh_advertise_policy reference

<a id="canonical-3010222011101133-3131011331220330-3123033233202132-3132111131310333-1021002301231222-0132012231031132-0333012123300203-0102212303112111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-1313323203033333-1300230112313332-2100102121033233-0222102011113211-1021300002113302-1302230112320222-2132203330001002-0321121130220222)
- Property reference

<a id="canonical-2000003332110201-3102011221233012-0230232010201233-2033011313212220-3222330310221003-3203201111122000-2210132203200322-2221303113001212"></a>

### Direct properties for `xcsh_advertise_policy`

<a id="canonical-0312132232231001-3213120311000332-3321232121113110-3310333000120320-1130013200113323-1100032021221202-1303020132222300-1311322200101322"></a>

#### `address` property

Type: `"string"`. Computed.

Optional. VIP to advertise. This VIP can be either V4/V6 address You can not specify this if where
contains a site or virtual site of type REGIONAL\_EDGE or public network If not specified and
'where' is specified with site or virtual site option, inside\_vip or outside\_vip specified in the
site..

Additional upstream details:

This VIP can be either V4/V6 address You can not specify this if where contains a site or virtual
site of type REGIONAL\_EDGE or public network If not specified and "where" is specified with site or
virtual site option, inside\_vip or outside\_vip specified in the site object will be used based on
the network type. If inside\_vip/outside\_vip is not configured in the site object, system use
interface IP in the respected networks.

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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-3031120030101113-2312130001002312-2020221221200012-0322122302222200-2201332003202210-1102002210322133-1031111331202201-2202312322330231"></a>

<a id="canonical-1220132313133113-3032133120102021-2031320120021312-0311001030313232-2221103120231211-3202221313332222-0310120323030222-0200103310023123"></a>

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

<a id="canonical-0200122303110123-1123231300223003-3211301310113100-0323211311212331-1213003023001113-3020100030012121-2231200323012312-3031121000202030"></a>

<a id="canonical-0111020332300322-1232322331112211-3323230011013031-0332012103202223-1322111021100123-1313223022013211-1322200032323202-2011230133131303"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the AdvertisePolicy.

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

- [dualstack](data-sources--advertise_policy--reference--group-001.md#canonical-1300321132021101-2312230222232310-0310313232330230-3132121011011210-3012022120112200-0112311123223231-0322331101121122-2111101311031301): complete subsection reference.

<a id="canonical-3301302323133032-1100031013012121-3300121203312120-3133113300000010-1320131211031312-0302013213120031-0202330021313000-2220212213221200"></a>

<a id="canonical-2232203231122302-1323233313030111-0032203212233331-1220220210220211-0133031100310131-3110111032022033-0122320021331201-3320313001332033"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [IPv4](data-sources--advertise_policy--reference--group-001.md#canonical-1300230300021232-2301101201003023-1102112101111220-2033230013100203-3113130231001221-2021032021321002-0121110201113322-3013030122031113): complete subsection reference.

- [IPv6](data-sources--advertise_policy--reference--group-001.md#canonical-3030320210032120-0311231213102131-3313031321031213-2212033303022323-3232202022303210-1200012232101102-2301310012000100-2103101021323202): complete subsection reference.

<a id="canonical-2212211202303302-0000012303020100-2323223232121113-3300202332003230-1000012321200311-3203021022221123-3231303031231333-1101031203022131"></a>

<a id="canonical-3132313100031230-2232032201202313-3331211303322032-0022022303100332-2300303213213221-1120233010002122-1320111103303203-1031303121303320"></a>

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

<a id="canonical-1002010303212113-1033301222331121-2312313020331313-3101133020231223-1132330002031123-3301120221212212-0333110133001131-1033220200330001"></a>

<a id="canonical-1332200023312110-1333033131221333-1021310000230200-3100010311231101-1132021122110131-3301220100033121-2311122301322122-2213100022113120"></a>

#### `name` property

Type: `"string"`. Required.

Name of the AdvertisePolicy.

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

<a id="canonical-3331112310123001-0302130210322102-3131100311213301-1000323100011102-2100113202322332-3300303010010313-1322330001321212-0023333112312003"></a>

<a id="canonical-1131332301102202-0232312330210321-0310231022130113-2320230231113220-1131102212221223-3232130200230112-0312323200212100-0002212120302310"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the AdvertisePolicy exists.

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

<a id="canonical-2310112132003132-1322111003131131-0003202130123323-0310020312200220-2022020223323031-2231021312201312-0322320002113021-2322130030003010"></a>

<a id="canonical-3120331132023020-2211333022323120-2013022111001021-1120132322300102-3123102131010200-0011110320223131-1001303212011202-2331101332202131"></a>

#### `port` property

Type: `"number"`. Computed.

\[OneOf: port, port\_ranges\] Exclusive with \[port\_ranges\] Port to advertise.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

OneOf alternatives in this subsection:

- [port](data-sources--advertise_policy--reference--group-001.md#canonical-2310112132003132-1322111003131131-0003202130123323-0310020312200220-2022020223323031-2231021312201312-0322320002113021-2322130030003010)
- [port_ranges](data-sources--advertise_policy--reference--group-001.md#canonical-0330030032021233-2303323321232322-0013322130011333-1333333330202322-2310020310131232-3300020300021313-1210110123300012-3320210213123000)

Select alternatives according to the provider validators above.

<a id="canonical-0330030032021233-2303323321232322-0013322130011333-1333333330202322-2310020310131232-3300020300021313-1210110123300012-3320210213123000"></a>

<a id="canonical-1110030212332031-2220333210302310-2120120102233020-1310123022112331-0121232030011020-3023222201301330-3232001122231103-0330302220322113"></a>

#### `port_ranges` property

Type: `"string"`. Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Additional upstream details:

Each port range consists of a single port or two ports separated by "-".

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "1024",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "1024",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-3123101322303100-2321232122022213-0000012212222132-3103211030122010-2030232233220300-0322130032210303-2102010330102232-2313112031200120"></a>

<a id="canonical-1313221011103023-1003122221012331-3132212301012330-2002112223123312-1103123322130331-1012120113022203-2103221311122330-1133312200003131"></a>

#### `protocol` property

Type: `"string"`. Computed.

\[Enum: TCP|UDP\] Protocol. Protocol to advertise. Possible values are \`TCP\`, \`UDP\`.

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
    "ves.io.schema.rules.string.in": "[\\\"TCP\\\",\\\"UDP\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"TCP\\\",\\\"UDP\\\"]"
  }
}
```

- [public_ip](data-sources--advertise_policy--reference--group-001.md#canonical-1200131111121300-2031302003000210-2100010210320211-0330113201210113-2201133223113121-2132003101201300-2012110130310313-2132313212322132): complete subsection reference.

<a id="canonical-2233133230321002-1033223112323331-3000030013331232-1320332012012232-3212231202130200-1122220103010310-0110331203103101-0010032112011030"></a>

<a id="canonical-0111032122131132-3220020010222021-2222233313102030-2230100023003123-1211220112310212-0012032233330003-3003301100100123-0202013120132320"></a>

#### `skip_xff_append` property

Type: `"bool"`. Computed.

If set, the loadbalancer will not append the remote address to the x-forwarded-for HTTP header.

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

- [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-1333123330133002-3132202201000303-2210303001223323-0323031213133013-3032003211001213-0123132123210112-1310023301123000-2302301003333103): complete subsection reference.

- [where](data-sources--advertise_policy--reference--group-001.md#canonical-3031120312202103-3312322201132331-3211232330311113-3303330100012123-2023223003011103-2303121112322110-3133101002313302-0322223323111230): complete subsection reference.

<a id="canonical-3232303020020103-0101223122003002-3300012012322112-1102122322022000-0300131233033302-1132222233332022-3122123123303112-1112321120010200"></a>

### All schema paths for `xcsh_advertise_policy`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address` | [address](data-sources--advertise_policy--reference--group-001.md#canonical-0312132232231001-3213120311000332-3321232121113110-3310333000120320-1130013200113323-1100032021221202-1303020132222300-1311322200101322) |
| `annotations` | [annotations](data-sources--advertise_policy--reference--group-001.md#canonical-3031120030101113-2312130001002312-2020221221200012-0322122302222200-2201332003202210-1102002210322133-1031111331202201-2202312322330231) |
| `description` | [description](data-sources--advertise_policy--reference--group-001.md#canonical-0200122303110123-1123231300223003-3211301310113100-0323211311212331-1213003023001113-3020100030012121-2231200323012312-3031121000202030) |
| `dualstack` | [dualstack](data-sources--advertise_policy--reference--group-001.md#canonical-3110332331203313-1230033002233231-2021132101321131-1132101320121132-2031001023320201-3212132102001311-3011310030332233-0010330121331002) |
| `id` | [ID](data-sources--advertise_policy--reference--group-001.md#canonical-3301302323133032-1100031013012121-3300121203312120-3133113300000010-1320131211031312-0302013213120031-0202330021313000-2220212213221200) |
| `ipv4` | [IPv4](data-sources--advertise_policy--reference--group-001.md#canonical-0311001102233200-2111233203223202-0233023313003031-2333202000330031-3212310231021120-0010222210301213-3223112212313323-1101300211102033) |
| `ipv6` | [IPv6](data-sources--advertise_policy--reference--group-001.md#canonical-2223220003111223-2131102330233131-3131123332110332-1302101012211303-2331223220300011-1132132020333301-0233200113310023-1213101002231130) |
| `labels` | [labels](data-sources--advertise_policy--reference--group-001.md#canonical-2212211202303302-0000012303020100-2323223232121113-3300202332003230-1000012321200311-3203021022221123-3231303031231333-1101031203022131) |
| `name` | [name](data-sources--advertise_policy--reference--group-001.md#canonical-1002010303212113-1033301222331121-2312313020331313-3101133020231223-1132330002031123-3301120221212212-0333110133001131-1033220200330001) |
| `namespace` | [namespace](data-sources--advertise_policy--reference--group-001.md#canonical-3331112310123001-0302130210322102-3131100311213301-1000323100011102-2100113202322332-3300303010010313-1322330001321212-0023333112312003) |
| `port` | [port](data-sources--advertise_policy--reference--group-001.md#canonical-2310112132003132-1322111003131131-0003202130123323-0310020312200220-2022020223323031-2231021312201312-0322320002113021-2322130030003010) |
| `port_ranges` | [port_ranges](data-sources--advertise_policy--reference--group-001.md#canonical-0330030032021233-2303323321232322-0013322130011333-1333333330202322-2310020310131232-3300020300021313-1210110123300012-3320210213123000) |
| `protocol` | [protocol](data-sources--advertise_policy--reference--group-001.md#canonical-3123101322303100-2321232122022213-0000012212222132-3103211030122010-2030232233220300-0322130032210303-2102010330102232-2313112031200120) |
| `public_ip` | [public_ip](data-sources--advertise_policy--reference--group-001.md#canonical-0130323220103120-1130113000323300-1111022312300101-3023223013312320-0111133202030113-2231331021301101-2131221200010303-1333002302322113) |
| `public_ip.kind` | [public_ip.kind](data-sources--advertise_policy--reference--group-001.md#canonical-2132113232012102-1130300311110210-0200303013321030-3221102111100003-3120100200312313-0132301100103331-1001032313203030-0020013021103003) |
| `public_ip.name` | [public_ip.name](data-sources--advertise_policy--reference--group-001.md#canonical-0021012222210203-3001300310013003-2121021210130021-0200000210012002-1211112212313032-0010122300220303-2323211130302013-0231031332230120) |
| `public_ip.namespace` | [public_ip.namespace](data-sources--advertise_policy--reference--group-001.md#canonical-3123302312121102-1010313312132031-0120022323130013-0102232030102020-1232330130033210-0012031300331033-0123323211201200-3023312203300221) |
| `public_ip.tenant` | [public_ip.tenant](data-sources--advertise_policy--reference--group-001.md#canonical-3113322331131023-1020032322012121-2001013121212312-2301311131133312-0100033320123031-3200312102002232-0200133231001013-1200130010321201) |
| `public_ip.uid` | [public_ip.uid](data-sources--advertise_policy--reference--group-001.md#canonical-1220002110103321-3220010323130100-0322023021030133-1202221011110330-0213210233202211-3221112021013021-3012021213031112-2200033001132213) |
| `skip_xff_append` | [skip_xff_append](data-sources--advertise_policy--reference--group-001.md#canonical-2233133230321002-1033223112323331-3000030013331232-1320332012012232-3212231202130200-1122220103010310-0110331203103101-0010032112011030) |
| `tls_parameters` | [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-2030303113101301-2121131101223031-1033132312000123-2333301012321331-1023101302033332-3123100110013221-3302303102231301-1321203221022212) |
| `tls_parameters.client_certificate_optional` | [tls_parameters.client_certificate_optional](data-sources--advertise_policy--reference--group-001.md#canonical-0120233222022210-2223331321103132-2103101123211021-2131033320113010-3203323213022000-1102003312132200-2030021011001021-0222111102112100) |
| `tls_parameters.client_certificate_required` | [tls_parameters.client_certificate_required](data-sources--advertise_policy--reference--group-001.md#canonical-2023131012203303-3011110321031300-1231222211012201-0202213201333133-3010201113321023-3223332230303132-3011100233100023-0030333302313213) |
| `tls_parameters.common_params` | [tls_parameters.common_params](data-sources--advertise_policy--reference--group-001.md#canonical-2013012131003031-0211023313122120-1112233202023300-3221110003103201-2032020123102102-0122001232031030-0002203231110303-0223232130123302) |
| `tls_parameters.common_params.cipher_suites` | [tls_parameters.common_params.cipher_suites](data-sources--advertise_policy--reference--group-001.md#canonical-1202321230222022-0021232311223230-3032032210130323-3120130312232211-0000131102301003-2211203221320221-1220131303301003-1030200300113322) |
| `tls_parameters.common_params.maximum_protocol_version` | [tls_parameters.common_params.maximum_protocol_version](data-sources--advertise_policy--reference--group-001.md#canonical-2023132031212230-0301130211223110-2012301231133230-2201212023333033-1300101032011303-3012022212121121-2133031233202311-3103101203221110) |
| `tls_parameters.common_params.minimum_protocol_version` | [tls_parameters.common_params.minimum_protocol_version](data-sources--advertise_policy--reference--group-001.md#canonical-0331133332322112-1232320133300032-0021130321321221-0110301212102020-3100223320222223-0221133103123231-0100033223202131-0021032120231301) |
| `tls_parameters.common_params.tls_certificates` | [tls_parameters.common_params.tls_certificates](data-sources--advertise_policy--reference--group-001.md#canonical-1222213130200331-1130113221210331-0222333201310333-3013333203013312-1303300200003003-3211010323221230-2121110203102203-0232020101001131) |
| `tls_parameters.common_params.tls_certificates.certificate_url` | [tls_parameters.common_params.tls_certificates.certificate_url](data-sources--advertise_policy--reference--group-001.md#canonical-0322330110030212-1111323030232100-1031222333131310-2032031133322012-1012212303310013-3231230112330001-2232210132230022-2221233113201221) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms](data-sources--advertise_policy--reference--group-001.md#canonical-0022130212201333-0100321332020002-0120032201210001-1021022302312330-2120313210322032-1123103332121310-2120213132013102-0033301112022221) |
| `tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms` | [tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms](data-sources--advertise_policy--reference--group-001.md#canonical-0020231313123331-0201210013230302-1223112310002310-1211302031132111-2101201301313302-2120220202202313-3323022300022013-2320302012002210) |
| `tls_parameters.common_params.tls_certificates.description_spec` | [tls_parameters.common_params.tls_certificates.description_spec](data-sources--advertise_policy--reference--group-001.md#canonical-1333003000112301-1033200210130013-1111031010020213-1332230331312203-3120111133200130-1320202202202033-3320330311132313-3311200133020201) |
| `tls_parameters.common_params.tls_certificates.disable_ocsp_stapling` | [tls_parameters.common_params.tls_certificates.disable_ocsp_stapling](data-sources--advertise_policy--reference--group-001.md#canonical-2231133033210132-2320311001312301-3020010000231212-0103223122103213-0021313333230003-2233000223202322-2032012221333213-0003002133002200) |
| `tls_parameters.common_params.tls_certificates.private_key` | [tls_parameters.common_params.tls_certificates.private_key](data-sources--advertise_policy--reference--group-001.md#canonical-1132213023320212-2320103103001122-2013121313303021-0311113231320302-2032132202330132-3223322121131320-2033233220123003-2110132230322211) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info](data-sources--advertise_policy--reference--group-001.md#canonical-1322113120103321-1010032331131200-1211111001032220-3303100103322300-3220323130231123-2013011100300222-3112021212032213-2022112011103311) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider](data-sources--advertise_policy--reference--group-001.md#canonical-1011032011020301-2021301231132303-2300223331222113-1001333211133100-3312023031130001-3033021033130122-0303211321123211-1031220122300102) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location](data-sources--advertise_policy--reference--group-001.md#canonical-1132100033312133-1112231322303223-2030031002011131-1301100112230233-2131113330200330-3000300133021200-1233203203101032-2032200230221001) |
| `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider` | [tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider](data-sources--advertise_policy--reference--group-001.md#canonical-3301212221312233-1221321003302012-2111312111110122-1011110232120003-3320210330303130-1123023122011032-0013232021322231-1311322122211122) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info](data-sources--advertise_policy--reference--group-001.md#canonical-3211213312333320-3312223232120000-1223133101221223-2130101210130233-0111102031031121-1222113310231210-1101120121120000-0102121200102101) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref](data-sources--advertise_policy--reference--group-001.md#canonical-0101213033033303-1231222133003302-3133300113022320-1302322002310102-1301313322030101-0200323022230210-1120113232102012-1120213331201000) |
| `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url` | [tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url](data-sources--advertise_policy--reference--group-001.md#canonical-1231020221311323-0331130101123002-0212032233022022-1221001233031313-0112201303221031-1320302003300131-1122033101320021-2323102223230230) |
| `tls_parameters.common_params.tls_certificates.use_system_defaults` | [tls_parameters.common_params.tls_certificates.use_system_defaults](data-sources--advertise_policy--reference--group-001.md#canonical-1132111213223223-2223100213131222-1233213002203233-1112321103323200-0020202101112321-3232322110000000-1031320132312301-0030232220202331) |
| `tls_parameters.common_params.validation_params` | [tls_parameters.common_params.validation_params](data-sources--advertise_policy--reference--group-001.md#canonical-0222333311120112-3212233212212210-1333000211110323-3312322220223311-1111312211310000-0320332112310330-0122111210012012-2133303012201012) |
| `tls_parameters.common_params.validation_params.skip_hostname_verification` | [tls_parameters.common_params.validation_params.skip_hostname_verification](data-sources--advertise_policy--reference--group-001.md#canonical-3203003023122310-2213301032023311-1310300213303311-1323113120130201-3201000102010211-0320221212212222-0310113331230110-2203130012201302) |
| `tls_parameters.common_params.validation_params.trusted_ca` | [tls_parameters.common_params.validation_params.trusted_ca](data-sources--advertise_policy--reference--group-001.md#canonical-1201013103201003-0030220023332100-3320032211032103-0010221032231210-0113132202220201-3303220112213201-3321030130333022-1202200322221122) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list](data-sources--advertise_policy--reference--group-001.md#canonical-0011201011201202-2032202221230122-2220102122232123-1211103333000232-0003313311100323-1133201100320223-1012003231133211-3010032312000201) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind](data-sources--advertise_policy--reference--group-001.md#canonical-1323302110331303-3223010111301023-1223022222113130-3301033332302200-1232330330130132-1030012311223121-2312200322311031-2222233110133102) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name](data-sources--advertise_policy--reference--group-001.md#canonical-3113213102332332-2232001022003310-3313023012020000-1122210222132032-2203121213221022-2320120121101021-1000032010002213-1121033021122232) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace](data-sources--advertise_policy--reference--group-001.md#canonical-0113200311110322-2023123201130201-3212000013023001-3023010030323213-3301210232110131-3222223232322300-3231032121321223-1123221032010223) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant](data-sources--advertise_policy--reference--group-001.md#canonical-2102131121033132-2131323331002120-0232202203210120-1011111110302201-0022130221310200-2122323303003320-2101220121332220-0200221200020203) |
| `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid` | [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid](data-sources--advertise_policy--reference--group-001.md#canonical-1203111022012130-1202232320232333-0121132300102302-3321232333333313-1212030130131132-2133210102303120-2113313101233120-3133012022230013) |
| `tls_parameters.common_params.validation_params.trusted_ca_url` | [tls_parameters.common_params.validation_params.trusted_ca_url](data-sources--advertise_policy--reference--group-001.md#canonical-3031212213021121-1303031323023233-3213232300023320-0302331130202020-2322303023011111-3232213003331310-1002102130330102-2223011303021031) |
| `tls_parameters.common_params.validation_params.verify_subject_alt_names` | [tls_parameters.common_params.validation_params.verify_subject_alt_names](data-sources--advertise_policy--reference--group-001.md#canonical-2302320130233213-2211000201233332-3130211031220012-1222113302231030-1130320223310331-1033223133311022-2002032100200201-1110103200110031) |
| `tls_parameters.no_client_certificate` | [tls_parameters.no_client_certificate](data-sources--advertise_policy--reference--group-001.md#canonical-1212233002330210-2132223133003320-3001030323211231-1230030123010031-2200213133200320-0131113322020321-1332330231302113-2322200330311031) |
| `tls_parameters.xfcc_header_elements` | [tls_parameters.xfcc_header_elements](data-sources--advertise_policy--reference--group-001.md#canonical-0132203210121200-3211302212012212-0123001222302133-3201003021131232-3100331230220222-0322300301101012-2313122231111231-3020332302302322) |
| `where` | [where](data-sources--advertise_policy--reference--group-001.md#canonical-3103303030003230-1022132033313103-3132012201201021-2031302011003301-2203103031010133-2220311023322021-3121312013101103-1232030232300210) |
| `where.site` | [where.site](data-sources--advertise_policy--reference--group-001.md#canonical-2311021213213111-3223002110022133-2003311203122233-3213330001321111-2112123310232321-2002212213203321-3220323001311322-1120111011021001) |
| `where.site.disable_internet_vip` | [where.site.disable_internet_vip](data-sources--advertise_policy--reference--group-001.md#canonical-3110001113213022-2230033113023303-2310101101110102-2001111223133113-1330123321103212-2220300130211211-2332033323212003-3213320222320033) |
| `where.site.enable_internet_vip` | [where.site.enable_internet_vip](data-sources--advertise_policy--reference--group-001.md#canonical-3121011200101302-3213222311230032-1003332103022031-3133121301110310-0302322131133112-0312323311103313-3110032301313031-3323222200233223) |
| `where.site.network_type` | [where.site.network_type](data-sources--advertise_policy--reference--group-001.md#canonical-3001201233320113-0313210332323013-3000220320211130-2312331310113121-1311320320020101-2322332302221202-3011110332310012-2122023332101130) |
| `where.site.ref` | [where.site.ref](data-sources--advertise_policy--reference--group-001.md#canonical-0231200303100320-3320101302333100-1100022201322321-0223000031020302-0002010212111113-0212322103012011-2112100211133312-3100200301201020) |
| `where.site.ref.kind` | [where.site.ref.kind](data-sources--advertise_policy--reference--group-001.md#canonical-0111131030311220-3102132310330233-2112300203100022-1332201000312101-1311203300032030-2002210002123201-1001211103210031-1232113021331001) |
| `where.site.ref.name` | [where.site.ref.name](data-sources--advertise_policy--reference--group-001.md#canonical-3202212023223313-2230332312302122-0302112203121122-0022302301100120-0120000313210130-0033103001332233-2323100131210032-3223303033023313) |
| `where.site.ref.namespace` | [where.site.ref.namespace](data-sources--advertise_policy--reference--group-001.md#canonical-2120103012002021-2130101311303122-1232211132012201-3122233321022011-3013330231020311-3013102030110101-3111313002311010-3031222302210023) |
| `where.site.ref.tenant` | [where.site.ref.tenant](data-sources--advertise_policy--reference--group-001.md#canonical-2312231000201300-1012321032130320-3002010020023333-2011032123323010-2122120323232010-2332313201230111-0131033200123011-1000212023122300) |
| `where.site.ref.uid` | [where.site.ref.uid](data-sources--advertise_policy--reference--group-001.md#canonical-1211232112013201-2233203223313330-3220121000122003-0011022132221303-0100230000121013-1031013231220231-0121300100120120-3320130023223031) |
| `where.virtual_network` | [where.virtual_network](data-sources--advertise_policy--reference--group-001.md#canonical-0133210112013231-0123020130212101-2022313223223021-2220130202233120-3221022230023111-1112303100230320-2311133032123031-0323111201010330) |
| `where.virtual_network.ref` | [where.virtual_network.ref](data-sources--advertise_policy--reference--group-001.md#canonical-0213011101230011-2212000332021103-1113011010232000-1202032032121103-3103200111232031-3302012001201320-1313001232101333-0302223120101110) |
| `where.virtual_network.ref.kind` | [where.virtual_network.ref.kind](data-sources--advertise_policy--reference--group-001.md#canonical-0212202131032030-2100110111213200-2122011101100201-0220323020333322-2102220020012233-1320302003122113-0102012221300023-2311111323120213) |
| `where.virtual_network.ref.name` | [where.virtual_network.ref.name](data-sources--advertise_policy--reference--group-001.md#canonical-2020100100312323-2000133203311231-1012102222203223-3133211230021212-3012123200321201-1020333131302310-3212321122022313-3230333232031312) |
| `where.virtual_network.ref.namespace` | [where.virtual_network.ref.namespace](data-sources--advertise_policy--reference--group-001.md#canonical-0323110132023332-2202023201200102-1222221011132032-1003101013333110-3230111110113123-0232333121203032-1032103022120121-1113302030120203) |
| `where.virtual_network.ref.tenant` | [where.virtual_network.ref.tenant](data-sources--advertise_policy--reference--group-001.md#canonical-1303001030131103-2012301003120001-0221123330013213-1023101321012333-3100012210122211-0211202303101310-1201132101220101-0103331231211020) |
| `where.virtual_network.ref.uid` | [where.virtual_network.ref.uid](data-sources--advertise_policy--reference--group-001.md#canonical-0211033132022001-3101330231023203-2000200321030011-3202022302320113-2120123010122323-2013003102021231-1331113221201222-2210012222323220) |
| `where.virtual_site` | [where.virtual_site](data-sources--advertise_policy--reference--group-001.md#canonical-2103103210031300-3323110301123030-2210220123312033-2021011030110003-2113200102120020-1132230310311233-1323110013123332-0033033300123100) |
| `where.virtual_site.disable_internet_vip` | [where.virtual_site.disable_internet_vip](data-sources--advertise_policy--reference--group-001.md#canonical-1022031331000310-3202123331332002-1010300112203033-0012032130210200-1323311003210100-2222111320023102-1202121213323001-0121222122121211) |
| `where.virtual_site.enable_internet_vip` | [where.virtual_site.enable_internet_vip](data-sources--advertise_policy--reference--group-001.md#canonical-2102030023030022-1101333311232131-3200112310202331-0003020323023002-3222302110300111-3223022233213322-2023202013231020-0233003312231101) |
| `where.virtual_site.network_type` | [where.virtual_site.network_type](data-sources--advertise_policy--reference--group-001.md#canonical-0122202132132202-1223322021021201-2003200112013002-0332313211210230-1312210001331202-1320213232010020-0102213332113012-1300313131323311) |
| `where.virtual_site.ref` | [where.virtual_site.ref](data-sources--advertise_policy--reference--group-001.md#canonical-1112311103233103-2030020002031312-3222321132201003-3202322231323212-0121013311013112-2011310130321232-1212323002202000-1330131133120300) |
| `where.virtual_site.ref.kind` | [where.virtual_site.ref.kind](data-sources--advertise_policy--reference--group-001.md#canonical-2133222020103022-3011333111111133-1302000332123213-2203123121032321-2321131212132030-1000332112313121-3122302031213300-0013333202330200) |
| `where.virtual_site.ref.name` | [where.virtual_site.ref.name](data-sources--advertise_policy--reference--group-001.md#canonical-3010012131231323-2133300023311031-2233130102113310-1302101033310123-3220210131131332-2333302221021333-2131001132200310-2111211102230313) |
| `where.virtual_site.ref.namespace` | [where.virtual_site.ref.namespace](data-sources--advertise_policy--reference--group-001.md#canonical-3033230303220010-1021131100030301-0131123012322212-0100133131113120-0003123301000200-3011321123300000-2021010021333110-2130223101312320) |
| `where.virtual_site.ref.tenant` | [where.virtual_site.ref.tenant](data-sources--advertise_policy--reference--group-001.md#canonical-1003032210320102-2323232002032102-2112023121013112-2202213102202031-0100221003011213-3133102312123333-0003010030333230-2222123202311012) |
| `where.virtual_site.ref.uid` | [where.virtual_site.ref.uid](data-sources--advertise_policy--reference--group-001.md#canonical-3331031330212101-3320313120001010-0320323203300320-1231012212233102-1101333320132202-2221130210220132-0232012203021103-3032200201002113) |

<a id="canonical-1300321132021101-2312230222232310-0310313232330230-3132121011011210-3012022120112200-0112311123223231-0322331101121122-2111101311031301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `dualstack` properties

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-1313323203033333-1300230112313332-2100102121033233-0222102011113211-1021300002113302-1302230112320222-2132203330001002-0321121130220222)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-3010222011101133-3131011331220330-3123033233202132-3132111131310333-1021002301231222-0132012231031132-0333012123300203-0102212303112111)
- dualstack

<a id="canonical-3110332331203313-1230033002233231-2021132101321131-1132101320121132-2031001023320201-3212132102001311-3011310030332233-0010330121331002"></a>

Type: `["object", {}]`. Computed.

\[OneOf: dualstack, IPv4, IPv6\] Enable this option

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

- [dualstack](data-sources--advertise_policy--reference--group-001.md#canonical-3110332331203313-1230033002233231-2021132101321131-1132101320121132-2031001023320201-3212132102001311-3011310030332233-0010330121331002)
- [IPv4](data-sources--advertise_policy--reference--group-001.md#canonical-0311001102233200-2111233203223202-0233023313003031-2333202000330031-3212310231021120-0010222210301213-3223112212313323-1101300211102033)
- [IPv6](data-sources--advertise_policy--reference--group-001.md#canonical-2223220003111223-2131102330233131-3131123332110332-1302101012211303-2331223220300011-1132132020333301-0233200113310023-1213101002231130)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1300230300021232-2301101201003023-1102112101111220-2033230013100203-3113130231001221-2021032021321002-0121110201113322-3013030122031113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipv4` properties

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-1313323203033333-1300230112313332-2100102121033233-0222102011113211-1021300002113302-1302230112320222-2132203330001002-0321121130220222)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-3010222011101133-3131011331220330-3123033233202132-3132111131310333-1021002301231222-0132012231031132-0333012123300203-0102212303112111)
- IPv4

<a id="canonical-0311001102233200-2111233203223202-0233023313003031-2333202000330031-3212310231021120-0010222210301213-3223112212313323-1101300211102033"></a>

Type: `["object", {}]`. Computed.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

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

<a id="canonical-3030320210032120-0311231213102131-3313031321031213-2212033303022323-3232202022303210-1200012232101102-2301310012000100-2103101021323202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `ipv6` properties

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-1313323203033333-1300230112313332-2100102121033233-0222102011113211-1021300002113302-1302230112320222-2132203330001002-0321121130220222)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-3010222011101133-3131011331220330-3123033233202132-3132111131310333-1021002301231222-0132012231031132-0333012123300203-0102212303112111)
- IPv6

<a id="canonical-2223220003111223-2131102330233131-3131123332110332-1302101012211303-2331223220300011-1132132020333301-0233200113310023-1213101002231130"></a>

Type: `["object", {}]`. Computed.

IPv6 address in colon-separated hexadecimal format.

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

<a id="canonical-1200131111121300-2031302003000210-2100010210320211-0330113201210113-2201133223113121-2132003101201300-2012110130310313-2132313212322132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `public_ip` properties

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-1313323203033333-1300230112313332-2100102121033233-0222102011113211-1021300002113302-1302230112320222-2132203330001002-0321121130220222)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-3010222011101133-3131011331220330-3123033233202132-3132111131310333-1021002301231222-0132012231031132-0333012123300203-0102212303112111)
- public_ip

<a id="canonical-0130323220103120-1130113000323300-1111022312300101-3023223013312320-0111133202030113-2231331021301101-2131221200010303-1333002302322113"></a>

Type: `"list"`. Computed.

Optional. Public VIP to advertise This field is mutually exclusive with where and address fields.

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

<a id="canonical-2311100222323323-1210133032111011-0023202202101202-1221101200031300-1022220102202332-2221131121331013-2323020030222313-0110011300312100"></a>

### Direct properties for `public_ip`

<a id="canonical-2132113232012102-1130300311110210-0200303013321030-3221102111100003-3120100200312313-0132301100103331-1001032313203030-0020013021103003"></a>

#### `public_ip.kind` property

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

<a id="canonical-0021012222210203-3001300310013003-2121021210130021-0200000210012002-1211112212313032-0010122300220303-2323211130302013-0231031332230120"></a>

<a id="canonical-3003321033301213-1003321123203220-2031032102023202-1221211113313021-1003212220210331-1022202203220031-3233031003033231-0023303100120130"></a>

#### `public_ip.name` property

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

<a id="canonical-3123302312121102-1010313312132031-0120022323130013-0102232030102020-1232330130033210-0012031300331033-0123323211201200-3023312203300221"></a>

<a id="canonical-3300102102320102-2033132203021030-0120131010013202-3323221313023031-1003332131300131-1102200012200133-3332010321312012-3023030331200011"></a>

#### `public_ip.namespace` property

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

<a id="canonical-3113322331131023-1020032322012121-2001013121212312-2301311131133312-0100033320123031-3200312102002232-0200133231001013-1200130010321201"></a>

<a id="canonical-2320300323102120-1100111031131221-1132331021011232-3023311101100220-2030110231100122-2231331001331310-2223131103011023-0023030112031303"></a>

#### `public_ip.tenant` property

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

<a id="canonical-1220002110103321-3220010323130100-0322023021030133-1202221011110330-0213210233202211-3221112021013021-3012021213031112-2200033001132213"></a>

<a id="canonical-3023333101111223-2000303023213101-1030112022310003-3132013331233132-1021310003221210-2131121031101123-1103301312110022-2313330022201011"></a>

#### `public_ip.uid` property

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

<a id="canonical-1333123330133002-3132202201000303-2210303001223323-0323031213133013-3032003211001213-0123132123210112-1310023301123000-2302301003333103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters` properties

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-1313323203033333-1300230112313332-2100102121033233-0222102011113211-1021300002113302-1302230112320222-2132203330001002-0321121130220222)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-3010222011101133-3131011331220330-3123033233202132-3132111131310333-1021002301231222-0132012231031132-0333012123300203-0102212303112111)
- tls_parameters

<a id="canonical-2030303113101301-2121131101223031-1033132312000123-2333301012321331-1023101302033332-3123100110013221-3302303102231301-1321203221022212"></a>

Type: `"single"`. Computed.

TLS configuration for downstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-client_certificate_verify_choice": "[\"client_certificate_optional\",\"client_certificate_required\",\"no_client_certificate\"]"
}
```

<a id="canonical-0122201212111131-3112130331003220-1232111103313321-2331211223233102-3101012112222110-0112300001300321-3002211032100200-1322300222122013"></a>

### Direct properties for `tls_parameters`

- [client_certificate_optional](data-sources--advertise_policy--reference--group-001.md#canonical-3020301002012332-1020200011231021-3003223203122012-3312110133011330-3232221232111211-1310220133301101-0021131020210233-1020003130133013): complete subsection reference.

- [client_certificate_required](data-sources--advertise_policy--reference--group-001.md#canonical-2233231030003123-0210001130012303-3311101233201323-2210013301130132-1221113220331111-2112113023231003-3220220123331333-0301222032322302): complete subsection reference.

- [common_params](data-sources--advertise_policy--reference--group-001.md#canonical-3313122022011100-1100013303303203-1233331311300012-1313103030111210-3021010313003023-0130303131120033-0022100313131303-0202130100132002): complete subsection reference.

- [no_client_certificate](data-sources--advertise_policy--reference--group-001.md#canonical-2223212220111111-2233230113112200-3100103203001101-3030012220303303-1111203311122213-2331023301330200-2331001011111302-1221201303301111): complete subsection reference.

<a id="canonical-0132203210121200-3211302212012212-0123001222302133-3201003021131232-3100331230220222-0322300301101012-2313122231111231-3020332302302322"></a>

<a id="canonical-2033222032111011-0133103312303110-1302323101333031-2112231123101120-1031301211301010-1031330322020033-0211021130331232-1030021311303012"></a>

#### `tls_parameters.xfcc_header_elements` property

Type: `["list", "string"]`. Computed.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be set in an mTLS enabled connections. If none are
defined, the header will not be added. Possible values are \`XFCC\_NONE\`, \`XFCC\_CERT\`,
\`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to \`XFCC\_NONE\`.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3020301002012332-1020200011231021-3003223203122012-3312110133011330-3232221232111211-1310220133301101-0021131020210233-1020003130133013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.client_certificate_optional` properties

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-1313323203033333-1300230112313332-2100102121033233-0222102011113211-1021300002113302-1302230112320222-2132203330001002-0321121130220222)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-3010222011101133-3131011331220330-3123033233202132-3132111131310333-1021002301231222-0132012231031132-0333012123300203-0102212303112111)
- [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-1333123330133002-3132202201000303-2210303001223323-0323031213133013-3032003211001213-0123132123210112-1310023301123000-2302301003333103)
- tls_parameters.client_certificate_optional

<a id="canonical-0120233222022210-2223331321103132-2103101123211021-2131033320113010-3203323213022000-1102003312132200-2030021011001021-0222111102112100"></a>

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

<a id="canonical-2233231030003123-0210001130012303-3311101233201323-2210013301130132-1221113220331111-2112113023231003-3220220123331333-0301222032322302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.client_certificate_required` properties

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-1313323203033333-1300230112313332-2100102121033233-0222102011113211-1021300002113302-1302230112320222-2132203330001002-0321121130220222)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-3010222011101133-3131011331220330-3123033233202132-3132111131310333-1021002301231222-0132012231031132-0333012123300203-0102212303112111)
- [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-1333123330133002-3132202201000303-2210303001223323-0323031213133013-3032003211001213-0123132123210112-1310023301123000-2302301003333103)
- tls_parameters.client_certificate_required

<a id="canonical-2023131012203303-3011110321031300-1231222211012201-0202213201333133-3010201113321023-3223332230303132-3011100233100023-0030333302313213"></a>

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

<a id="canonical-3313122022011100-1100013303303203-1233331311300012-1313103030111210-3021010313003023-0130303131120033-0022100313131303-0202130100132002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params` properties

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-1313323203033333-1300230112313332-2100102121033233-0222102011113211-1021300002113302-1302230112320222-2132203330001002-0321121130220222)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-3010222011101133-3131011331220330-3123033233202132-3132111131310333-1021002301231222-0132012231031132-0333012123300203-0102212303112111)
- [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-1333123330133002-3132202201000303-2210303001223323-0323031213133013-3032003211001213-0123132123210112-1310023301123000-2302301003333103)
- tls_parameters.common_params

<a id="canonical-2013012131003031-0211023313122120-1112233202023300-3221110003103201-2032020123102102-0122001232031030-0002203231110303-0223232130123302"></a>

Type: `"single"`. Computed.

Information of different aspects for TLS authentication related to ciphers, certificates and trust
store.

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

<a id="canonical-2230221303213231-3112021232013331-0231013223323222-2232113122212012-0201330112232300-1213100211102121-0333222013233102-1223003131300323"></a>

### Direct properties for `tls_parameters.common_params`

<a id="canonical-1202321230222022-0021232311223230-3032032210130323-3120130312232211-0000131102301003-2211203221320221-1220131303301003-1030200300113322"></a>

#### `tls_parameters.common_params.cipher_suites` property

Type: `["list", "string"]`. Computed.

The following list specifies the supported cipher suite TLS\_AES\_128\_GCM\_SHA256
TLS\_AES\_256\_GCM\_SHA384 TLS\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384 TLS\_ECDHE\_RSA\_WITH\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_CBC\_SHA TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_CBC\_SHA
TLS\_ECDHE\_RSA\_WITH\_AES\_128\_CBC\_SHA TLS\_ECDHE\_RSA\_WITH\_AES\_256\_CBC\_SHA
TLS\_RSA\_WITH\_AES\_128\_CBC\_SHA TLS\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_RSA\_WITH\_AES\_256\_CBC\_SHA TLS\_RSA\_WITH\_AES\_256\_GCM\_SHA384

If not specified, the default list: TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_RSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384 will be used.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2023132031212230-0301130211223110-2012301231133230-2201212023333033-1300101032011303-3012022212121121-2133031233202311-3103101203221110"></a>

<a id="canonical-1311202102313311-3012033302030332-3310223002133232-1000323023311303-2002133130013331-3021122133323211-0103313321220100-1203310330202233"></a>

#### `tls_parameters.common_params.maximum_protocol_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0331133332322112-1232320133300032-0021130321321221-0110301212102020-3100223320222223-0221133103123231-0100033223202131-0021032120231301"></a>

<a id="canonical-1121000312232132-0212221030310203-2333230221111212-0211221301023000-3030331330322010-1302213030031203-3030332110032101-2210021322210113"></a>

#### `tls_parameters.common_params.minimum_protocol_version` property

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [tls_certificates](data-sources--advertise_policy--reference--group-001.md#canonical-3212201103100220-3022333332123023-1031313202031303-0022022110101322-2030230222033000-0213221202032221-0323000133123303-0222021120221130): complete subsection reference.

- [validation_params](data-sources--advertise_policy--reference--group-001.md#canonical-0322133030322321-2320230313111300-2303313113210011-0223030322312100-2333103030321112-2022322223330100-3101112313222031-0303321002201220): complete subsection reference.

<a id="canonical-3212201103100220-3022333332123023-1031313202031303-0022022110101322-2030230222033000-0213221202032221-0323000133123303-0222021120221130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates` properties

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-1313323203033333-1300230112313332-2100102121033233-0222102011113211-1021300002113302-1302230112320222-2132203330001002-0321121130220222)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-3010222011101133-3131011331220330-3123033233202132-3132111131310333-1021002301231222-0132012231031132-0333012123300203-0102212303112111)
- [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-1333123330133002-3132202201000303-2210303001223323-0323031213133013-3032003211001213-0123132123210112-1310023301123000-2302301003333103)
- [tls_parameters.common_params](data-sources--advertise_policy--reference--group-001.md#canonical-3313122022011100-1100013303303203-1233331311300012-1313103030111210-3021010313003023-0130303131120033-0022100313131303-0202130100132002)
- tls_parameters.common_params.tls_certificates

<a id="canonical-1222213130200331-1130113221210331-0222333201310333-3013333203013312-1303300200003003-3211010323221230-2121110203102203-0232020101001131"></a>

Type: `"list"`. Computed.

TLS Certificates. Set of TLS certificates.

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

<a id="canonical-3102330202303203-2303203012120111-3213131210102010-0211303113101232-0320012300213002-0000331221221230-2133001302222132-3103201212333101"></a>

### Direct properties for `tls_parameters.common_params.tls_certificates`

<a id="canonical-0322330110030212-1111323030232100-1031222333131310-2032031133322012-1012212303310013-3231230112330001-2232210132230022-2221233113201221"></a>

#### `tls_parameters.common_params.tls_certificates.certificate_url` property

Type: `"string"`. Computed.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [custom_hash_algorithms](data-sources--advertise_policy--reference--group-001.md#canonical-0321212230323001-2302231122203113-3103310200012312-3210130233010120-2223331022211233-0123031121131301-2313121310123202-0130112203012333): complete subsection reference.

<a id="canonical-1333003000112301-1033200210130013-1111031010020213-1332230331312203-3120111133200130-1320202202202033-3320330311132313-3311200133020201"></a>

<a id="canonical-3000103200213003-0310322200301113-2102303222133322-2303222100211103-3021213302210222-2310012223111323-2331010000333100-1011033300112023"></a>

#### `tls_parameters.common_params.tls_certificates.description_spec` property

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--advertise_policy--reference--group-001.md#canonical-3022332101000010-1000000103201220-0323030213212321-2101313110002231-2023130000130303-2103232331122000-3301033111311021-0312002022032103): complete subsection reference.

- [private_key](data-sources--advertise_policy--reference--group-001.md#canonical-0323030313030010-2103100132200310-0233232101002303-0003302312120300-1123101022232022-0232332232222113-1000213320213003-1301123113301303): complete subsection reference.

- [use_system_defaults](data-sources--advertise_policy--reference--group-001.md#canonical-3201311330022203-3322122202111030-3103113302222112-0203202111310032-1210021233130302-2133300011030133-0100330000102231-3203313201032233): complete subsection reference.

<a id="canonical-0321212230323001-2302231122203113-3103310200012312-3210130233010120-2223331022211233-0123031121131301-2313121310123202-0130112203012333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-1313323203033333-1300230112313332-2100102121033233-0222102011113211-1021300002113302-1302230112320222-2132203330001002-0321121130220222)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-3010222011101133-3131011331220330-3123033233202132-3132111131310333-1021002301231222-0132012231031132-0333012123300203-0102212303112111)
- [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-1333123330133002-3132202201000303-2210303001223323-0323031213133013-3032003211001213-0123132123210112-1310023301123000-2302301003333103)
- [tls_parameters.common_params](data-sources--advertise_policy--reference--group-001.md#canonical-3313122022011100-1100013303303203-1233331311300012-1313103030111210-3021010313003023-0130303131120033-0022100313131303-0202130100132002)
- [tls_parameters.common_params.tls_certificates](data-sources--advertise_policy--reference--group-001.md#canonical-3212201103100220-3022333332123023-1031313202031303-0022022110101322-2030230222033000-0213221202032221-0323000133123303-0222021120221130)
- tls_parameters.common_params.tls_certificates.custom_hash_algorithms

<a id="canonical-0022130212201333-0100321332020002-0120032201210001-1021022302312330-2120313210322032-1123103332121310-2120213132013102-0033301112022221"></a>

Type: `"single"`. Computed.

Specifies the hash algorithms to be used.

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

<a id="canonical-3201131321001210-3120131230020002-3311212031130301-0103320323222203-2032223310122112-1001221032131203-3223020122323110-3111233210120001"></a>

### Direct properties for `tls_parameters.common_params.tls_certificates.custom_hash_algorithms`

<a id="canonical-0020231313123331-0201210013230302-1223112310002310-1211302031132111-2101201301313302-2120220202202313-3323022300022013-2320302012002210"></a>

#### `tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms` property

Type: `["list", "string"]`. Computed.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3022332101000010-1000000103201220-0323030213212321-2101313110002231-2023130000130303-2103232331122000-3301033111311021-0312002022032103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-1313323203033333-1300230112313332-2100102121033233-0222102011113211-1021300002113302-1302230112320222-2132203330001002-0321121130220222)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-3010222011101133-3131011331220330-3123033233202132-3132111131310333-1021002301231222-0132012231031132-0333012123300203-0102212303112111)
- [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-1333123330133002-3132202201000303-2210303001223323-0323031213133013-3032003211001213-0123132123210112-1310023301123000-2302301003333103)
- [tls_parameters.common_params](data-sources--advertise_policy--reference--group-001.md#canonical-3313122022011100-1100013303303203-1233331311300012-1313103030111210-3021010313003023-0130303131120033-0022100313131303-0202130100132002)
- [tls_parameters.common_params.tls_certificates](data-sources--advertise_policy--reference--group-001.md#canonical-3212201103100220-3022333332123023-1031313202031303-0022022110101322-2030230222033000-0213221202032221-0323000133123303-0222021120221130)
- tls_parameters.common_params.tls_certificates.disable_ocsp_stapling

<a id="canonical-2231133033210132-2320311001312301-3020010000231212-0103223122103213-0021313333230003-2233000223202322-2032012221333213-0003002133002200"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable ocsp stapling.

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

<a id="canonical-0323030313030010-2103100132200310-0233232101002303-0003302312120300-1123101022232022-0232332232222113-1000213320213003-1301123113301303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates.private_key` properties

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-1313323203033333-1300230112313332-2100102121033233-0222102011113211-1021300002113302-1302230112320222-2132203330001002-0321121130220222)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-3010222011101133-3131011331220330-3123033233202132-3132111131310333-1021002301231222-0132012231031132-0333012123300203-0102212303112111)
- [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-1333123330133002-3132202201000303-2210303001223323-0323031213133013-3032003211001213-0123132123210112-1310023301123000-2302301003333103)
- [tls_parameters.common_params](data-sources--advertise_policy--reference--group-001.md#canonical-3313122022011100-1100013303303203-1233331311300012-1313103030111210-3021010313003023-0130303131120033-0022100313131303-0202130100132002)
- [tls_parameters.common_params.tls_certificates](data-sources--advertise_policy--reference--group-001.md#canonical-3212201103100220-3022333332123023-1031313202031303-0022022110101322-2030230222033000-0213221202032221-0323000133123303-0222021120221130)
- tls_parameters.common_params.tls_certificates.private_key

<a id="canonical-1132213023320212-2320103103001122-2013121313303021-0311113231320302-2032132202330132-3223322121131320-2033233220123003-2110132230322211"></a>

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

<a id="canonical-1012012130103321-3212301230121213-2130312330022130-2130030023310231-1110230023030220-1002320221110313-2122200332223030-2231102020202133"></a>

### Direct properties for `tls_parameters.common_params.tls_certificates.private_key`

- [blindfold_secret_info](data-sources--advertise_policy--reference--group-001.md#canonical-2022031222202320-1021203131310222-1320023103302221-1203132302323222-1230230223133121-1322200123112123-2023133021212102-0023213131331331): complete subsection reference.

- [clear_secret_info](data-sources--advertise_policy--reference--group-001.md#canonical-0332112320202101-2312133030102011-0122111101032322-0100013112131302-2132322223230313-3212312323211201-3000120301211331-1321231123300132): complete subsection reference.

<a id="canonical-2022031222202320-1021203131310222-1320023103302221-1203132302323222-1230230223133121-1322200123112123-2023133021212102-0023213131331331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-1313323203033333-1300230112313332-2100102121033233-0222102011113211-1021300002113302-1302230112320222-2132203330001002-0321121130220222)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-3010222011101133-3131011331220330-3123033233202132-3132111131310333-1021002301231222-0132012231031132-0333012123300203-0102212303112111)
- [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-1333123330133002-3132202201000303-2210303001223323-0323031213133013-3032003211001213-0123132123210112-1310023301123000-2302301003333103)
- [tls_parameters.common_params](data-sources--advertise_policy--reference--group-001.md#canonical-3313122022011100-1100013303303203-1233331311300012-1313103030111210-3021010313003023-0130303131120033-0022100313131303-0202130100132002)
- [tls_parameters.common_params.tls_certificates](data-sources--advertise_policy--reference--group-001.md#canonical-3212201103100220-3022333332123023-1031313202031303-0022022110101322-2030230222033000-0213221202032221-0323000133123303-0222021120221130)
- [tls_parameters.common_params.tls_certificates.private_key](data-sources--advertise_policy--reference--group-001.md#canonical-0323030313030010-2103100132200310-0233232101002303-0003302312120300-1123101022232022-0232332232222113-1000213320213003-1301123113301303)
- tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-1322113120103321-1010032331131200-1211111001032220-3303100103322300-3220323130231123-2013011100300222-3112021212032213-2022112011103311"></a>

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

<a id="canonical-1220321012032330-1323233322223313-3020233221121233-1030311211122212-0211322010201102-2332121321312031-2132221320311230-3312002333103020"></a>

### Direct properties for `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-1011032011020301-2021301231132303-2300223331222113-1001333211133100-3312023031130001-3033021033130122-0303211321123211-1031220122300102"></a>

#### `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

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

<a id="canonical-1132100033312133-1112231322303223-2030031002011131-1301100112230233-2131113330200330-3000300133021200-1233203203101032-2032200230221001"></a>

<a id="canonical-0121200302330122-0003321330203030-1000113023303001-2102201020202300-2223213012202023-2223011013003112-1200113320231131-0323330323131033"></a>

#### `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location` property

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

<a id="canonical-3301212221312233-1221321003302012-2111312111110122-1011110232120003-3320210330303130-1123023122011032-0013232021322231-1311322122211122"></a>

<a id="canonical-2220320300121210-2302013302113120-0000013112202332-1230023330322032-0100023222220001-3231121212022323-0200233021012012-0132023013300222"></a>

#### `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider` property

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

<a id="canonical-0332112320202101-2312133030102011-0122111101032322-0100013112131302-2132322223230313-3212312323211201-3000120301211331-1321231123300132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-1313323203033333-1300230112313332-2100102121033233-0222102011113211-1021300002113302-1302230112320222-2132203330001002-0321121130220222)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-3010222011101133-3131011331220330-3123033233202132-3132111131310333-1021002301231222-0132012231031132-0333012123300203-0102212303112111)
- [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-1333123330133002-3132202201000303-2210303001223323-0323031213133013-3032003211001213-0123132123210112-1310023301123000-2302301003333103)
- [tls_parameters.common_params](data-sources--advertise_policy--reference--group-001.md#canonical-3313122022011100-1100013303303203-1233331311300012-1313103030111210-3021010313003023-0130303131120033-0022100313131303-0202130100132002)
- [tls_parameters.common_params.tls_certificates](data-sources--advertise_policy--reference--group-001.md#canonical-3212201103100220-3022333332123023-1031313202031303-0022022110101322-2030230222033000-0213221202032221-0323000133123303-0222021120221130)
- [tls_parameters.common_params.tls_certificates.private_key](data-sources--advertise_policy--reference--group-001.md#canonical-0323030313030010-2103100132200310-0233232101002303-0003302312120300-1123101022232022-0232332232222113-1000213320213003-1301123113301303)
- tls_parameters.common_params.tls_certificates.private_key.clear_secret_info

<a id="canonical-3211213312333320-3312223232120000-1223133101221223-2130101210130233-0111102031031121-1222113310231210-1101120121120000-0102121200102101"></a>

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

<a id="canonical-0322010213021032-3122300111323031-1210303130002013-3311030130330221-2232111100120133-0200221333102123-0130120032023100-3023030312303210"></a>

### Direct properties for `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info`

<a id="canonical-0101213033033303-1231222133003302-3133300113022320-1302322002310102-1301313322030101-0200323022230210-1120113232102012-1120213331201000"></a>

#### `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1231020221311323-0331130101123002-0212032233022022-1221001233031313-0112201303221031-1320302003300131-1122033101320021-2323102223230230"></a>

<a id="canonical-1220222111103032-1120012321232300-2203101120212012-0000313020012023-1231210333001112-3322221320130011-0323132021333321-2103331321100222"></a>

#### `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url` property

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

<a id="canonical-3201311330022203-3322122202111030-3103113302222112-0203202111310032-1210021233130302-2133300011030133-0100330000102231-3203313201032233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-1313323203033333-1300230112313332-2100102121033233-0222102011113211-1021300002113302-1302230112320222-2132203330001002-0321121130220222)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-3010222011101133-3131011331220330-3123033233202132-3132111131310333-1021002301231222-0132012231031132-0333012123300203-0102212303112111)
- [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-1333123330133002-3132202201000303-2210303001223323-0323031213133013-3032003211001213-0123132123210112-1310023301123000-2302301003333103)
- [tls_parameters.common_params](data-sources--advertise_policy--reference--group-001.md#canonical-3313122022011100-1100013303303203-1233331311300012-1313103030111210-3021010313003023-0130303131120033-0022100313131303-0202130100132002)
- [tls_parameters.common_params.tls_certificates](data-sources--advertise_policy--reference--group-001.md#canonical-3212201103100220-3022333332123023-1031313202031303-0022022110101322-2030230222033000-0213221202032221-0323000133123303-0222021120221130)
- tls_parameters.common_params.tls_certificates.use_system_defaults

<a id="canonical-1132111213223223-2223100213131222-1233213002203233-1112321103323200-0020202101112321-3232322110000000-1031320132312301-0030232220202331"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for use system defaults.

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

<a id="canonical-0322133030322321-2320230313111300-2303313113210011-0223030322312100-2333103030321112-2022322223330100-3101112313222031-0303321002201220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.validation_params` properties

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-1313323203033333-1300230112313332-2100102121033233-0222102011113211-1021300002113302-1302230112320222-2132203330001002-0321121130220222)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-3010222011101133-3131011331220330-3123033233202132-3132111131310333-1021002301231222-0132012231031132-0333012123300203-0102212303112111)
- [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-1333123330133002-3132202201000303-2210303001223323-0323031213133013-3032003211001213-0123132123210112-1310023301123000-2302301003333103)
- [tls_parameters.common_params](data-sources--advertise_policy--reference--group-001.md#canonical-3313122022011100-1100013303303203-1233331311300012-1313103030111210-3021010313003023-0130303131120033-0022100313131303-0202130100132002)
- tls_parameters.common_params.validation_params

<a id="canonical-0222333311120112-3212233212212210-1333000211110323-3312322220223311-1111312211310000-0320332112310330-0122111210012012-2133303012201012"></a>

Type: `"single"`. Computed.

This includes URL for a trust store, whether SAN verification is required and list of Subject Alt
Names for verification.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]"
}
```

<a id="canonical-3230223331033010-0230012321133103-3331012131213300-1011302310312021-0032131003221000-0300132200022220-1300112111001032-2132311013200003"></a>

### Direct properties for `tls_parameters.common_params.validation_params`

<a id="canonical-3203003023122310-2213301032023311-1310300213303311-1323113120130201-3201000102010211-0320221212212222-0310113331230110-2203130012201302"></a>

#### `tls_parameters.common_params.validation_params.skip_hostname_verification` property

Type: `"bool"`. Computed.

When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to
the connecting hostname.

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

- [trusted_ca](data-sources--advertise_policy--reference--group-001.md#canonical-3123013323111022-3221213313030012-3113003210231313-3012022013330003-1001120213101111-0121112133012312-1123013223130130-1130121330300033): complete subsection reference.

<a id="canonical-3031212213021121-1303031323023233-3213232300023320-0302331130202020-2322303023011111-3232213003331310-1002102130330102-2223011303021031"></a>

<a id="canonical-3131220022100211-1233321331221030-1331203311121020-0123212121200332-0232231233310221-3213330001302320-0101021123001333-2321002113131331"></a>

#### `tls_parameters.common_params.validation_params.trusted_ca_url` property

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

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
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-2302320130233213-2211000201233332-3130211031220012-1222113302231030-1130320223310331-1033223133311022-2002032100200201-1110103200110031"></a>

<a id="canonical-0201023120323303-3111000332222032-1232110131131211-1122233000313000-0012030130311323-3231111002020302-0101310132333202-3003322123022202"></a>

#### `tls_parameters.common_params.validation_params.verify_subject_alt_names` property

Type: `["list", "string"]`. Computed.

List of acceptable Subject Alt Names/CN in the peer's certificate. When skip\_hostname\_verification
is false and verify\_subject\_alt\_names is empty, the hostname of the peer will be used for
matching against SAN/CN of peer's certificate.

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

<a id="canonical-3123013323111022-3221213313030012-3113003210231313-3012022013330003-1001120213101111-0121112133012312-1123013223130130-1130121330300033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.validation_params.trusted_ca` properties

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-1313323203033333-1300230112313332-2100102121033233-0222102011113211-1021300002113302-1302230112320222-2132203330001002-0321121130220222)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-3010222011101133-3131011331220330-3123033233202132-3132111131310333-1021002301231222-0132012231031132-0333012123300203-0102212303112111)
- [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-1333123330133002-3132202201000303-2210303001223323-0323031213133013-3032003211001213-0123132123210112-1310023301123000-2302301003333103)
- [tls_parameters.common_params](data-sources--advertise_policy--reference--group-001.md#canonical-3313122022011100-1100013303303203-1233331311300012-1313103030111210-3021010313003023-0130303131120033-0022100313131303-0202130100132002)
- [tls_parameters.common_params.validation_params](data-sources--advertise_policy--reference--group-001.md#canonical-0322133030322321-2320230313111300-2303313113210011-0223030322312100-2333103030321112-2022322223330100-3101112313222031-0303321002201220)
- tls_parameters.common_params.validation_params.trusted_ca

<a id="canonical-1201013103201003-0030220023332100-3320032211032103-0010221032231210-0113132202220201-3303220112213201-3321030130333022-1202200322221122"></a>

Type: `"single"`. Computed.

Root CA Certificate Reference. Reference to Root CA Certificate.

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

<a id="canonical-3033022203313013-1111213011003033-3303120023000211-2303031022332112-3132032320011030-3210020023201130-3031203210232132-1002020223013011"></a>

### Direct properties for `tls_parameters.common_params.validation_params.trusted_ca`

- [trusted_ca_list](data-sources--advertise_policy--reference--group-001.md#canonical-0232202223122102-2212112130322213-3030103003312101-2103023233322222-0112200221213031-0103321323102020-3123211132301113-1031323212300212): complete subsection reference.

<a id="canonical-0232202223122102-2212112130322213-3030103003312101-2103023233322222-0112200221213031-0103321323102020-3123211132301113-1031323212300212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list` properties

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-1313323203033333-1300230112313332-2100102121033233-0222102011113211-1021300002113302-1302230112320222-2132203330001002-0321121130220222)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-3010222011101133-3131011331220330-3123033233202132-3132111131310333-1021002301231222-0132012231031132-0333012123300203-0102212303112111)
- [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-1333123330133002-3132202201000303-2210303001223323-0323031213133013-3032003211001213-0123132123210112-1310023301123000-2302301003333103)
- [tls_parameters.common_params](data-sources--advertise_policy--reference--group-001.md#canonical-3313122022011100-1100013303303203-1233331311300012-1313103030111210-3021010313003023-0130303131120033-0022100313131303-0202130100132002)
- [tls_parameters.common_params.validation_params](data-sources--advertise_policy--reference--group-001.md#canonical-0322133030322321-2320230313111300-2303313113210011-0223030322312100-2333103030321112-2022322223330100-3101112313222031-0303321002201220)
- [tls_parameters.common_params.validation_params.trusted_ca](data-sources--advertise_policy--reference--group-001.md#canonical-3123013323111022-3221213313030012-3113003210231313-3012022013330003-1001120213101111-0121112133012312-1123013223130130-1130121330300033)
- tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list

<a id="canonical-0011201011201202-2032202221230122-2220102122232123-1211103333000232-0003313311100323-1133201100320223-1012003231133211-3010032312000201"></a>

Type: `"list"`. Computed.

Root CA Certificate Reference. Reference to Root CA Certificate.

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-1300003011120212-2130120223222313-3202312300221120-0320032123303020-1203003200203100-0201313131002010-3111031022332021-1321033110233210"></a>

### Direct properties for `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list`

<a id="canonical-1323302110331303-3223010111301023-1223022222113130-3301033332302200-1232330330130132-1030012311223121-2312200322311031-2222233110133102"></a>

#### `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind` property

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

<a id="canonical-3113213102332332-2232001022003310-3313023012020000-1122210222132032-2203121213221022-2320120121101021-1000032010002213-1121033021122232"></a>

<a id="canonical-3002222312002210-3232231101031111-1013032013222333-2112032121323130-0320220222320300-1312300220223031-1101101213010031-3000021303231001"></a>

#### `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name` property

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

<a id="canonical-0113200311110322-2023123201130201-3212000013023001-3023010030323213-3301210232110131-3222223232322300-3231032121321223-1123221032010223"></a>

<a id="canonical-2333203031331322-2211112231132320-2032220223123333-0011332313231001-1202003313310103-0331312033031130-3102223321011212-3120133211321013"></a>

#### `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace` property

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

<a id="canonical-2102131121033132-2131323331002120-0232202203210120-1011111110302201-0022130221310200-2122323303003320-2101220121332220-0200221200020203"></a>

<a id="canonical-1303123033211032-1120111211330220-1210021231222301-0013323302021112-3302321031333312-2303113313020203-0220131323223223-1002013223001220"></a>

#### `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant` property

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

<a id="canonical-1203111022012130-1202232320232333-0121132300102302-3321232333333313-1212030130131132-2133210102303120-2113313101233120-3133012022230013"></a>

<a id="canonical-0301231123301330-3231311132021330-0000231030223201-0002331311303232-1223012210233010-3211202132330000-0022332030033031-2302001322010001"></a>

#### `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid` property

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

<a id="canonical-2223212220111111-2233230113112200-3100103203001101-3030012220303303-1111203311122213-2331023301330200-2331001011111302-1221201303301111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.no_client_certificate` properties

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-1313323203033333-1300230112313332-2100102121033233-0222102011113211-1021300002113302-1302230112320222-2132203330001002-0321121130220222)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-3010222011101133-3131011331220330-3123033233202132-3132111131310333-1021002301231222-0132012231031132-0333012123300203-0102212303112111)
- [tls_parameters](data-sources--advertise_policy--reference--group-001.md#canonical-1333123330133002-3132202201000303-2210303001223323-0323031213133013-3032003211001213-0123132123210112-1310023301123000-2302301003333103)
- tls_parameters.no_client_certificate

<a id="canonical-1212233002330210-2132223133003320-3001030323211231-1230030123010031-2200213133200320-0131113322020321-1332330231302113-2322200330311031"></a>

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

<a id="canonical-3031120312202103-3312322201132331-3211232330311113-3303330100012123-2023223003011103-2303121112322110-3133101002313302-0322223323111230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where` properties

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-1313323203033333-1300230112313332-2100102121033233-0222102011113211-1021300002113302-1302230112320222-2132203330001002-0321121130220222)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-3010222011101133-3131011331220330-3123033233202132-3132111131310333-1021002301231222-0132012231031132-0333012123300203-0102212303112111)
- where

<a id="canonical-3103303030003230-1022132033313103-3132012201201021-2031302011003301-2203103031010133-2220311023322021-3121312013101103-1232030232300210"></a>

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

<a id="canonical-0000001233013213-3213310202033013-1330131010330210-2131102301130213-0021213003031013-0211130311330313-1032300303133213-3000231122012232"></a>

### Direct properties for `where`

- [site](data-sources--advertise_policy--reference--group-001.md#canonical-2112001213300311-2001131200322322-0303021101333320-3111001211121210-2322201100323202-2300302322003030-0333212200221122-1202013330321122): complete subsection reference.

- [virtual_network](data-sources--advertise_policy--reference--group-001.md#canonical-3022121311302203-3101121122331332-1030013301210030-1330202012133112-2103012211120022-3202303132311322-2231200013110312-3231223110331133): complete subsection reference.

- [virtual_site](data-sources--advertise_policy--reference--group-001.md#canonical-1122300101022033-1023233111033000-3200123020121203-1031111032230031-2102210032312210-0211312202230101-3021312302303020-0330211202113310): complete subsection reference.

<a id="canonical-2112001213300311-2001131200322322-0303021101333320-3111001211121210-2322201100323202-2300302322003030-0333212200221122-1202013330321122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.site` properties

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-1313323203033333-1300230112313332-2100102121033233-0222102011113211-1021300002113302-1302230112320222-2132203330001002-0321121130220222)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-3010222011101133-3131011331220330-3123033233202132-3132111131310333-1021002301231222-0132012231031132-0333012123300203-0102212303112111)
- [where](data-sources--advertise_policy--reference--group-001.md#canonical-3031120312202103-3312322201132331-3211232330311113-3303330100012123-2023223003011103-2303121112322110-3133101002313302-0322223323111230)
- where.site

<a id="canonical-2311021213213111-3223002110022133-2003311203122233-3213330001321111-2112123310232321-2002212213203321-3220323001311322-1120111011021001"></a>

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

<a id="canonical-0313203211310123-0001220131132221-2111122311223013-1123131013132120-1320011300211300-2033201322113123-1133330221112130-1033320003031323"></a>

### Direct properties for `where.site`

- [disable_internet_vip](data-sources--advertise_policy--reference--group-001.md#canonical-1021013022020010-3313001012132201-2100333122002313-1121022203303210-1132313121232000-0200012031302332-0130023302333300-3111332100232223): complete subsection reference.

- [enable_internet_vip](data-sources--advertise_policy--reference--group-001.md#canonical-2133100313033002-3110321222301102-0322220320113320-3120220113102331-1223102200031132-1102310000101011-1100100333032123-3123303200133112): complete subsection reference.

<a id="canonical-3001201233320113-0313210332323013-3000220320211130-2312331310113121-1311320320020101-2322332302221202-3011110332310012-2122023332101130"></a>

<a id="canonical-3210100320033011-1013113031303200-0012332000211312-0103101220120120-1002303123303101-1223010212132200-0210012200011020-1001333003220213"></a>

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

- [ref](data-sources--advertise_policy--reference--group-001.md#canonical-1013013220222110-1003110320330113-2232333213302211-0103001232112231-2020220310231313-0013121121333233-0230111230332101-2022013120022030): complete subsection reference.

<a id="canonical-1021013022020010-3313001012132201-2100333122002313-1121022203303210-1132313121232000-0200012031302332-0130023302333300-3111332100232223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.site.disable_internet_vip` properties

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-1313323203033333-1300230112313332-2100102121033233-0222102011113211-1021300002113302-1302230112320222-2132203330001002-0321121130220222)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-3010222011101133-3131011331220330-3123033233202132-3132111131310333-1021002301231222-0132012231031132-0333012123300203-0102212303112111)
- [where](data-sources--advertise_policy--reference--group-001.md#canonical-3031120312202103-3312322201132331-3211232330311113-3303330100012123-2023223003011103-2303121112322110-3133101002313302-0322223323111230)
- [where.site](data-sources--advertise_policy--reference--group-001.md#canonical-2112001213300311-2001131200322322-0303021101333320-3111001211121210-2322201100323202-2300302322003030-0333212200221122-1202013330321122)
- where.site.disable_internet_vip

<a id="canonical-3110001113213022-2230033113023303-2310101101110102-2001111223133113-1330123321103212-2220300130211211-2332033323212003-3213320222320033"></a>

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

<a id="canonical-2133100313033002-3110321222301102-0322220320113320-3120220113102331-1223102200031132-1102310000101011-1100100333032123-3123303200133112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.site.enable_internet_vip` properties

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-1313323203033333-1300230112313332-2100102121033233-0222102011113211-1021300002113302-1302230112320222-2132203330001002-0321121130220222)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-3010222011101133-3131011331220330-3123033233202132-3132111131310333-1021002301231222-0132012231031132-0333012123300203-0102212303112111)
- [where](data-sources--advertise_policy--reference--group-001.md#canonical-3031120312202103-3312322201132331-3211232330311113-3303330100012123-2023223003011103-2303121112322110-3133101002313302-0322223323111230)
- [where.site](data-sources--advertise_policy--reference--group-001.md#canonical-2112001213300311-2001131200322322-0303021101333320-3111001211121210-2322201100323202-2300302322003030-0333212200221122-1202013330321122)
- where.site.enable_internet_vip

<a id="canonical-3121011200101302-3213222311230032-1003332103022031-3133121301110310-0302322131133112-0312323311103313-3110032301313031-3323222200233223"></a>

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

<a id="canonical-1013013220222110-1003110320330113-2232333213302211-0103001232112231-2020220310231313-0013121121333233-0230111230332101-2022013120022030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.site.ref` properties

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-1313323203033333-1300230112313332-2100102121033233-0222102011113211-1021300002113302-1302230112320222-2132203330001002-0321121130220222)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-3010222011101133-3131011331220330-3123033233202132-3132111131310333-1021002301231222-0132012231031132-0333012123300203-0102212303112111)
- [where](data-sources--advertise_policy--reference--group-001.md#canonical-3031120312202103-3312322201132331-3211232330311113-3303330100012123-2023223003011103-2303121112322110-3133101002313302-0322223323111230)
- [where.site](data-sources--advertise_policy--reference--group-001.md#canonical-2112001213300311-2001131200322322-0303021101333320-3111001211121210-2322201100323202-2300302322003030-0333212200221122-1202013330321122)
- where.site.ref

<a id="canonical-0231200303100320-3320101302333100-1100022201322321-0223000031020302-0002010212111113-0212322103012011-2112100211133312-3100200301201020"></a>

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-1011032231000211-2102002112021003-0020320312331122-1033113333031320-3332000332010000-0310320001303110-3232113320012323-3200331133321322"></a>

### Direct properties for `where.site.ref`

<a id="canonical-0111131030311220-3102132310330233-2112300203100022-1332201000312101-1311203300032030-2002210002123201-1001211103210031-1232113021331001"></a>

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

<a id="canonical-3202212023223313-2230332312302122-0302112203121122-0022302301100120-0120000313210130-0033103001332233-2323100131210032-3223303033023313"></a>

<a id="canonical-1033321103133000-1110313112133302-3123023233003101-0032002230130332-2130130311112132-2202120020302110-2130013100110210-2120331123223113"></a>

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

<a id="canonical-2120103012002021-2130101311303122-1232211132012201-3122233321022011-3013330231020311-3013102030110101-3111313002311010-3031222302210023"></a>

<a id="canonical-1223312013331321-3121033003032002-0233132233010013-3103230013110301-2121023130112332-3103032300213231-1130130100120311-1012113211300301"></a>

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

<a id="canonical-2312231000201300-1012321032130320-3002010020023333-2011032123323010-2122120323232010-2332313201230111-0131033200123011-1000212023122300"></a>

<a id="canonical-0212003311031022-2120221320130312-3233010331122333-1211311102230102-1112001133222102-0211131012303303-2203013322122212-3132013211111112"></a>

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

<a id="canonical-1211232112013201-2233203223313330-3220121000122003-0011022132221303-0100230000121013-1031013231220231-0121300100120120-3320130023223031"></a>

<a id="canonical-3020122031312311-0331233103122010-0030130213301010-3021212211323100-1322220122101300-2222011131120122-1022131223111201-1212221333202332"></a>

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

<a id="canonical-3022121311302203-3101121122331332-1030013301210030-1330202012133112-2103012211120022-3202303132311322-2231200013110312-3231223110331133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_network` properties

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-1313323203033333-1300230112313332-2100102121033233-0222102011113211-1021300002113302-1302230112320222-2132203330001002-0321121130220222)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-3010222011101133-3131011331220330-3123033233202132-3132111131310333-1021002301231222-0132012231031132-0333012123300203-0102212303112111)
- [where](data-sources--advertise_policy--reference--group-001.md#canonical-3031120312202103-3312322201132331-3211232330311113-3303330100012123-2023223003011103-2303121112322110-3133101002313302-0322223323111230)
- where.virtual_network

<a id="canonical-0133210112013231-0123020130212101-2022313223223021-2220130202233120-3221022230023111-1112303100230320-2311133032123031-0323111201010330"></a>

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

<a id="canonical-2313101303331211-0110202330030133-3033033012020022-1020332312010310-3010213231320200-3212102000331220-0313121321220222-1120313223210120"></a>

### Direct properties for `where.virtual_network`

- [ref](data-sources--advertise_policy--reference--group-001.md#canonical-0121332123303331-0320331131031002-1333030322113132-3113032223130300-2122131122121302-1302001111200322-0322102112320031-0012312012111010): complete subsection reference.

<a id="canonical-0121332123303331-0320331131031002-1333030322113132-3113032223130300-2122131122121302-1302001111200322-0322102112320031-0012312012111010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_network.ref` properties

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-1313323203033333-1300230112313332-2100102121033233-0222102011113211-1021300002113302-1302230112320222-2132203330001002-0321121130220222)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-3010222011101133-3131011331220330-3123033233202132-3132111131310333-1021002301231222-0132012231031132-0333012123300203-0102212303112111)
- [where](data-sources--advertise_policy--reference--group-001.md#canonical-3031120312202103-3312322201132331-3211232330311113-3303330100012123-2023223003011103-2303121112322110-3133101002313302-0322223323111230)
- [where.virtual_network](data-sources--advertise_policy--reference--group-001.md#canonical-3022121311302203-3101121122331332-1030013301210030-1330202012133112-2103012211120022-3202303132311322-2231200013110312-3231223110331133)
- where.virtual_network.ref

<a id="canonical-0213011101230011-2212000332021103-1113011010232000-1202032032121103-3103200111232031-3302012001201320-1313001232101333-0302223120101110"></a>

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-3222223122011232-3331201130331103-2102023130232332-3110023021013220-1300211032103023-3111023020302201-1000232303232110-2001201011232323"></a>

### Direct properties for `where.virtual_network.ref`

<a id="canonical-0212202131032030-2100110111213200-2122011101100201-0220323020333322-2102220020012233-1320302003122113-0102012221300023-2311111323120213"></a>

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

<a id="canonical-2020100100312323-2000133203311231-1012102222203223-3133211230021212-3012123200321201-1020333131302310-3212321122022313-3230333232031312"></a>

<a id="canonical-1220131001222111-0003320223101102-3011121200302011-1123001122113322-3210232012023211-2203103232303130-1233213103213100-3301321210031322"></a>

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

<a id="canonical-0323110132023332-2202023201200102-1222221011132032-1003101013333110-3230111110113123-0232333121203032-1032103022120121-1113302030120203"></a>

<a id="canonical-1103331220202011-3202130022022021-3122011300012302-3021200111323100-3200032102201003-0220032021212120-3023220133233220-3330010002032123"></a>

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

<a id="canonical-1303001030131103-2012301003120001-0221123330013213-1023101321012333-3100012210122211-0211202303101310-1201132101220101-0103331231211020"></a>

<a id="canonical-0330013203301033-2003213211202323-2303031100322031-2211230312001122-0013311333013020-3200032110221301-2212212130232003-2001002131032110"></a>

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

<a id="canonical-0211033132022001-3101330231023203-2000200321030011-3202022302320113-2120123010122323-2013003102021231-1331113221201222-2210012222323220"></a>

<a id="canonical-1300113201200000-1123312201001322-2133111113230012-0221322301002022-3023120220001013-0323300012012030-1213311122223120-3123022301231003"></a>

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

<a id="canonical-1122300101022033-1023233111033000-3200123020121203-1031111032230031-2102210032312210-0211312202230101-3021312302303020-0330211202113310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_site` properties

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-1313323203033333-1300230112313332-2100102121033233-0222102011113211-1021300002113302-1302230112320222-2132203330001002-0321121130220222)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-3010222011101133-3131011331220330-3123033233202132-3132111131310333-1021002301231222-0132012231031132-0333012123300203-0102212303112111)
- [where](data-sources--advertise_policy--reference--group-001.md#canonical-3031120312202103-3312322201132331-3211232330311113-3303330100012123-2023223003011103-2303121112322110-3133101002313302-0322223323111230)
- where.virtual_site

<a id="canonical-2103103210031300-3323110301123030-2210220123312033-2021011030110003-2113200102120020-1132230310311233-1323110013123332-0033033300123100"></a>

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

<a id="canonical-2003122203220132-1111121301321020-3000012031220100-2020310102122102-1121310020323310-1132011210121013-3023320123010331-2220213031331332"></a>

### Direct properties for `where.virtual_site`

- [disable_internet_vip](data-sources--advertise_policy--reference--group-001.md#canonical-1310031202030210-3300302112322233-3201303220121123-0102303230330122-3322130100303022-1310332203131312-0223210312213202-0020111320223102): complete subsection reference.

- [enable_internet_vip](data-sources--advertise_policy--reference--group-001.md#canonical-3330222322123002-1200130002330023-2331321232113321-3203312001021201-3123301201222023-0020120011033003-1331211033133322-0101332221131230): complete subsection reference.

<a id="canonical-0122202132132202-1223322021021201-2003200112013002-0332313211210230-1312210001331202-1320213232010020-0102213332113012-1300313131323311"></a>

<a id="canonical-3122303120033022-0200233202302210-2033002113021233-1122231211322211-3311322112201012-0121002322303321-0213311023010202-3203112321000201"></a>

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

- [ref](data-sources--advertise_policy--reference--group-001.md#canonical-3333333023331330-1213320333031232-0003211103021132-2230331323020311-1221110211113002-0120220321221123-0010212301010122-3013032320313313): complete subsection reference.

<a id="canonical-1310031202030210-3300302112322233-3201303220121123-0102303230330122-3322130100303022-1310332203131312-0223210312213202-0020111320223102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_site.disable_internet_vip` properties

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-1313323203033333-1300230112313332-2100102121033233-0222102011113211-1021300002113302-1302230112320222-2132203330001002-0321121130220222)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-3010222011101133-3131011331220330-3123033233202132-3132111131310333-1021002301231222-0132012231031132-0333012123300203-0102212303112111)
- [where](data-sources--advertise_policy--reference--group-001.md#canonical-3031120312202103-3312322201132331-3211232330311113-3303330100012123-2023223003011103-2303121112322110-3133101002313302-0322223323111230)
- [where.virtual_site](data-sources--advertise_policy--reference--group-001.md#canonical-1122300101022033-1023233111033000-3200123020121203-1031111032230031-2102210032312210-0211312202230101-3021312302303020-0330211202113310)
- where.virtual_site.disable_internet_vip

<a id="canonical-1022031331000310-3202123331332002-1010300112203033-0012032130210200-1323311003210100-2222111320023102-1202121213323001-0121222122121211"></a>

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

<a id="canonical-3330222322123002-1200130002330023-2331321232113321-3203312001021201-3123301201222023-0020120011033003-1331211033133322-0101332221131230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_site.enable_internet_vip` properties

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-1313323203033333-1300230112313332-2100102121033233-0222102011113211-1021300002113302-1302230112320222-2132203330001002-0321121130220222)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-3010222011101133-3131011331220330-3123033233202132-3132111131310333-1021002301231222-0132012231031132-0333012123300203-0102212303112111)
- [where](data-sources--advertise_policy--reference--group-001.md#canonical-3031120312202103-3312322201132331-3211232330311113-3303330100012123-2023223003011103-2303121112322110-3133101002313302-0322223323111230)
- [where.virtual_site](data-sources--advertise_policy--reference--group-001.md#canonical-1122300101022033-1023233111033000-3200123020121203-1031111032230031-2102210032312210-0211312202230101-3021312302303020-0330211202113310)
- where.virtual_site.enable_internet_vip

<a id="canonical-2102030023030022-1101333311232131-3200112310202331-0003020323023002-3222302110300111-3223022233213322-2023202013231020-0233003312231101"></a>

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

<a id="canonical-3333333023331330-1213320333031232-0003211103021132-2230331323020311-1221110211113002-0120220321221123-0010212301010122-3013032320313313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `where.virtual_site.ref` properties

Breadcrumbs:

- [xcsh_advertise_policy](../data-sources/advertise_policy.md#canonical-1313323203033333-1300230112313332-2100102121033233-0222102011113211-1021300002113302-1302230112320222-2132203330001002-0321121130220222)
- [Property reference](data-sources--advertise_policy--reference--group-001.md#canonical-3010222011101133-3131011331220330-3123033233202132-3132111131310333-1021002301231222-0132012231031132-0333012123300203-0102212303112111)
- [where](data-sources--advertise_policy--reference--group-001.md#canonical-3031120312202103-3312322201132331-3211232330311113-3303330100012123-2023223003011103-2303121112322110-3133101002313302-0322223323111230)
- [where.virtual_site](data-sources--advertise_policy--reference--group-001.md#canonical-1122300101022033-1023233111033000-3200123020121203-1031111032230031-2102210032312210-0211312202230101-3021312302303020-0330211202113310)
- where.virtual_site.ref

<a id="canonical-1112311103233103-2030020002031312-3222321132201003-3202322231323212-0121013311013112-2011310130321232-1212323002202000-1330131133120300"></a>

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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-1113311021013300-3313322012203011-3332123111222300-2210202110230020-2132321212110003-0321320101330323-3012120010303211-0103022310230311"></a>

### Direct properties for `where.virtual_site.ref`

<a id="canonical-2133222020103022-3011333111111133-1302000332123213-2203123121032321-2321131212132030-1000332112313121-3122302031213300-0013333202330200"></a>

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

<a id="canonical-3010012131231323-2133300023311031-2233130102113310-1302101033310123-3220210131131332-2333302221021333-2131001132200310-2111211102230313"></a>

<a id="canonical-0120302113213213-0323202030113003-1312303312301121-0211223223111102-1013131331000231-2331001000330202-0022003323100123-3312120020121023"></a>

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

<a id="canonical-3033230303220010-1021131100030301-0131123012322212-0100133131113120-0003123301000200-3011321123300000-2021010021333110-2130223101312320"></a>

<a id="canonical-2211301002212213-2312021122103322-2100000131212203-1330130121133311-0311130203232303-3023130221223233-2220332220003323-1023103312210113"></a>

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

<a id="canonical-1003032210320102-2323232002032102-2112023121013112-2202213102202031-0100221003011213-3133102312123333-0003010030333230-2222123202311012"></a>

<a id="canonical-1222101230011123-0003312321113202-1033002211003312-2123022301012000-0101230002311023-3230300231320103-1000113302021331-0223221103302122"></a>

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

<a id="canonical-3331031330212101-3320313120001010-0320323203300320-1231012212233102-1101333320132202-2221130210220132-0232012203021103-3032200201002113"></a>

<a id="canonical-3222202212303021-0101101021031311-2233100332330121-2332303313333221-3002222021313203-0303122203101123-1022102100203010-1123111013023133"></a>

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
