---
page_title: "xcsh_crl reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_crl reference."
---

# xcsh_crl reference

<a id="canonical-0022012330331322-3131111112231003-3101103003200200-3211113200221201-0100030210222013-3122120001201231-2010332202003001-0013313221310300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113000010131323-1330203112033303-3123212232013212-1200333300312231-2233003122200321-2013010212301303-1233103331330202-2202020331301233"></a>

## Property reference — Property reference / 022020203031 / 2

Breadcrumbs:

- [xcsh_crl](../resources/crl.md#canonical-2112221203101103-2313113220020333-1301313220130233-1101313202220020-1222121022300210-0123222110132200-2221301201331022-0001302303100010)
- Property reference

<a id="canonical-2131320133311110-1210112103103111-1303111232033302-0312021213312333-1033021313313200-0012130033012010-2200001223011122-1123012122032122"></a>

## Direct properties — Property reference / 022020203031 / 3

<a id="canonical-2223212220100113-0031301333213013-3213212123201233-2111333202300320-3313030120233201-0021001221322301-0300230123301320-2011121032113012"></a>

<a id="canonical-0103013130121211-0030300122122131-3311212120121202-1100121230231311-3221230332030330-2032221021030222-2301031022232223-0033213010102001"></a>

## annotations property — Property reference / 022020203031 / 4

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

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

<a id="canonical-2321122301001323-3020203233101203-2121212101112011-0231012113133000-3033022122031331-2300121110323022-2002121123010300-1220320021301333"></a>

<a id="canonical-0002123112333213-2021212012022303-0033120322113032-1200033111332013-1232031132021101-3312113121333001-1103231113310212-0022320033121100"></a>

## description property — Property reference / 022020203031 / 5

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

<a id="canonical-3311231113222100-2200323310131313-1130331110232110-3030131000130220-2301231210023302-1120121032103200-0000312131033000-3323131300022203"></a>

<a id="canonical-0122322102200301-1132332323212230-1033331133331131-1003303330221301-0213300131223233-0322303323023303-1210332312312322-2323230113321321"></a>

## disable property — Property reference / 022020203031 / 6

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

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

- [http_access](resources--crl--reference--group-001.md#canonical-3321030000232202-1110003002322111-3110232113133210-0120020233031003-2100102002203222-0101012223111031-0331330123323332-3220011030031300): complete subsection reference.

<a id="canonical-1233112321132201-0122020233301131-0313132013101201-0100030213323220-2132313331133331-2020333000301221-0120203030001013-2010113001111303"></a>

<a id="canonical-3001311311213132-2103333120033112-0003130232022013-0112123223120020-1332113120232201-1313111011310210-3203023032122313-2021120131202332"></a>

## ID property — Property reference / 022020203031 / 7

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0232210111323011-3322212120230120-0300131102332030-0020301110030100-1300300031301120-0002100123110102-0020110230100030-2021302302211130"></a>

<a id="canonical-2202330303311010-2032200202121213-2032332112023010-3211202303011031-3102220210003100-0310212313313331-1321132002220100-2112331120120221"></a>

## labels property — Property reference / 022020203031 / 8

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

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

<a id="canonical-3012212221311213-0333221212012101-1213131323131332-0021221010033302-2111313003333310-1011111100012232-3323333231213012-0022020012302202"></a>

<a id="canonical-0011223321121233-1112300302221213-3132320121022031-3311110013311311-2013130313012113-3112332323021111-0031331212232122-3320200033201133"></a>

## name property — Property reference / 022020203031 / 9

Type: `"string"`. Required.

Name of the CRL. Must be unique within the namespace.

Upstream description:

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

<a id="canonical-2101200011320033-2203011111331300-0002111200130130-0011223002012233-1303120130002330-1121220231231320-1001321002130310-1031002133313010"></a>

<a id="canonical-1212330303231130-3110310023020210-3111101111121130-1201310113221221-0120032001312313-2221131311023212-2312230212111232-1020333320110021"></a>

## namespace property — Property reference / 022020203031 / 10

Type: `"string"`. Required.

Namespace where the CRL is created.

Upstream description:

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

<a id="canonical-3212002102003230-0102323320310330-2111110120223131-3100322221123010-1331121223313333-1302000011320301-3122031000220013-0301102000111302"></a>

<a id="canonical-1011322023021001-3331200130103213-3212311031022113-3200210113210200-2330020202110101-3031031230000032-3311113010301303-3112203333103330"></a>

## refresh_interval property — Property reference / 022020203031 / 11

Type: `"number"`. Required.

CRL Refresh interval. CRL refresh interval, in hours.

Upstream description:

CRL refresh interval, in hours.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(6, 168),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 168,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 6
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "6",
    "ves.io.schema.rules.uint32.lte": "168"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "6",
    "ves.io.schema.rules.uint32.lte": "168"
  }
}
```

<a id="canonical-3110330032120331-2023011000312002-0233103321002000-3011200130031232-3112100021231200-0233321323023332-1223230003102101-0231312200011123"></a>

<a id="canonical-1133322303022301-2111303222020223-3113320030310100-3101111320222121-3002032310200133-2113023212013022-0332013230121230-2023330300233231"></a>

## server_address property — Property reference / 022020203031 / 12

Type: `"string"`. Required.

CRL Server address. CRL server address or hostname.

Upstream description:

CRL server address or hostname.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(255),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
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
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname_or_ip": "true",
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

<a id="canonical-3300212301011103-2313311330203120-3333222223121131-0010301012332130-1201333230213102-0331012000033313-2133222013330333-0233201223023212"></a>

<a id="canonical-0211032313031121-2232021322222212-1212010230020312-2100132101321110-2003121113103123-2010303110302022-0230203213030130-1120132001302012"></a>

## server_port property — Property reference / 022020203031 / 13

Type: `"number"`. Required.

CRL Server Port. Set CRL Server port number.

Upstream description:

Set CRL Server port number.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
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

<a id="canonical-2012102131013303-2132130023330111-1302131210020123-3322300323100310-1103313031333103-0301320311202103-3321323201200321-0330200000311120"></a>

<a id="canonical-2001230031100102-2331030122231031-2101213133233333-2303230210220030-3113023003310323-1122131122103021-2233111202003332-2023103013323302"></a>

## timeout property — Property reference / 022020203031 / 14

Type: `"number"`. Required.

CRL download timeout. CRL download wait time, in seconds.

Upstream description:

CRL download wait time, in seconds.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 180),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 180,
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "180"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "180"
  }
}
```

- [timeouts](resources--crl--reference--group-001.md#canonical-0220030102113332-1101303210023103-1222122101230020-2033222222320013-2030332300202012-2100200130323302-1102311303330322-0322323212130002): complete subsection reference.

<a id="canonical-3330300302012232-1213312110121032-0311222011121210-0222002201112230-3033222101031110-1300021001112223-3121212131010103-1303312102201213"></a>

## All schema paths — Property reference / 022020203031 / 15

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--crl--reference--group-001.md#canonical-2223212220100113-0031301333213013-3213212123201233-2111333202300320-3313030120233201-0021001221322301-0300230123301320-2011121032113012) |
| `description` | [description](resources--crl--reference--group-001.md#canonical-2321122301001323-3020203233101203-2121212101112011-0231012113133000-3033022122031331-2300121110323022-2002121123010300-1220320021301333) |
| `disable` | [disable](resources--crl--reference--group-001.md#canonical-3311231113222100-2200323310131313-1130331110232110-3030131000130220-2301231210023302-1120121032103200-0000312131033000-3323131300022203) |
| `http_access` | [http_access](resources--crl--reference--group-001.md#canonical-3120233022203211-0122120132112100-1102030033210010-0322332133230130-3223011102212301-2030122113111330-1202112323211320-1231133311223302) |
| `http_access.path` | [http_access.path](resources--crl--reference--group-001.md#canonical-3322030022312011-0212130032221113-0313203223302111-0132301332311332-2303333332302211-1322321033231211-0230131311112023-1131202331133300) |
| `id` | [ID](resources--crl--reference--group-001.md#canonical-1233112321132201-0122020233301131-0313132013101201-0100030213323220-2132313331133331-2020333000301221-0120203030001013-2010113001111303) |
| `labels` | [labels](resources--crl--reference--group-001.md#canonical-0232210111323011-3322212120230120-0300131102332030-0020301110030100-1300300031301120-0002100123110102-0020110230100030-2021302302211130) |
| `name` | [name](resources--crl--reference--group-001.md#canonical-3012212221311213-0333221212012101-1213131323131332-0021221010033302-2111313003333310-1011111100012232-3323333231213012-0022020012302202) |
| `namespace` | [namespace](resources--crl--reference--group-001.md#canonical-2101200011320033-2203011111331300-0002111200130130-0011223002012233-1303120130002330-1121220231231320-1001321002130310-1031002133313010) |
| `refresh_interval` | [refresh_interval](resources--crl--reference--group-001.md#canonical-3212002102003230-0102323320310330-2111110120223131-3100322221123010-1331121223313333-1302000011320301-3122031000220013-0301102000111302) |
| `server_address` | [server_address](resources--crl--reference--group-001.md#canonical-3110330032120331-2023011000312002-0233103321002000-3011200130031232-3112100021231200-0233321323023332-1223230003102101-0231312200011123) |
| `server_port` | [server_port](resources--crl--reference--group-001.md#canonical-3300212301011103-2313311330203120-3333222223121131-0010301012332130-1201333230213102-0331012000033313-2133222013330333-0233201223023212) |
| `timeout` | [timeout](resources--crl--reference--group-001.md#canonical-2012102131013303-2132130023330111-1302131210020123-3322300323100310-1103313031333103-0301320311202103-3321323201200321-0330200000311120) |
| `timeouts` | [timeouts](resources--crl--reference--group-001.md#canonical-2201033001031002-0203000121033033-3112022330031210-0112322230121303-2133233120001222-1032302213333213-3131333323110003-2013122200030031) |
| `timeouts.create` | [timeouts.create](resources--crl--reference--group-001.md#canonical-0120213213300011-1020011033202201-0102202301022200-2231121000302233-1200200202211032-3213021330222232-2002200201330212-2100001032033310) |
| `timeouts.delete` | [timeouts.delete](resources--crl--reference--group-001.md#canonical-3213202133322013-0201203010033102-1123233000202030-3330223220033022-2110121120233110-3311111010232230-2203002130012213-0033031002303222) |
| `timeouts.read` | [timeouts.read](resources--crl--reference--group-001.md#canonical-3220000030300200-1100331230332010-0100301322021021-0111222230321300-2130222202101332-2012030333020103-1203203013101120-2230023021001110) |
| `timeouts.update` | [timeouts.update](resources--crl--reference--group-001.md#canonical-3130031231121203-1232002031132130-2103133102212222-0302233222300233-2333311131221101-1122221301122202-0301230022100300-0311223312300222) |

<a id="canonical-3022200233310333-1100023302011331-3112131313311212-0111330133102203-0212113032023220-1002021311103321-3203030113313230-0210203012010103"></a>

## Next pages — Property reference / 022020203031 / 16

- [http_access](resources--crl--reference--group-001.md#canonical-3321030000232202-1110003002322111-3110232113133210-0120020233031003-2100102002203222-0101012223111031-0331330123323332-3220011030031300)
- [timeouts](resources--crl--reference--group-001.md#canonical-0220030102113332-1101303210023103-1222122101230020-2033222222320013-2030332300202012-2100200130323302-1102311303330322-0322323212130002)
- [xcsh_crl](../resources/crl.md#canonical-2112221203101103-2313113220020333-1301313220130233-1101313202220020-1222121022300210-0123222110132200-2221301201331022-0001302303100010)

<a id="canonical-3321030000232202-1110003002322111-3110232113133210-0120020233031003-2100102002203222-0101012223111031-0331330123323332-3220011030031300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203230212003110-1323131132212233-3110000203211222-1111111001312112-1222011220231210-2331011100223203-2013003022120333-3011012030021220"></a>

## http_access — http_access / 000301112113 / 2

Breadcrumbs:

- [xcsh_crl](../resources/crl.md#canonical-2112221203101103-2313113220020333-1301313220130233-1101313202220020-1222121022300210-0123222110132200-2221301201331022-0001302303100010)
- [Property reference](resources--crl--reference--group-001.md#canonical-0022012330331322-3131111112231003-3101103003200200-3211113200221201-0100030210222013-3122120001201231-2010332202003001-0013313221310300)
- http_access

<a id="canonical-3120233022203211-0122120132112100-1102030033210010-0322332133230130-3223011102212301-2030122113111330-1202112323211320-1231133311223302"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http access.

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

Terraform syntax:

```terraform
http_access {
  # Configure direct properties listed below.
}
```

<a id="canonical-1233231333300102-0300123301112000-1201233333023132-0211000312210213-1312010300323131-0112021001012122-2200202322113112-1012211021020010"></a>

## Direct properties — http_access / 000301112113 / 3

<a id="canonical-3322030022312011-0212130032221113-0313203223302111-0132301332311332-2303333332302211-1322321033231211-0230131311112023-1131202331133300"></a>

<a id="canonical-1300120030313332-3320112202231010-0000103010323211-1303212223012023-0022011303113323-2012330112321331-1011130023132010-2012102133322222"></a>

## path property — http_access / 000301112113 / 4

Type: `"string"`. Optional.

CRL File path. CRL file location.

Upstream description:

CRL file location.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3110001330001030-3332013011132033-3012210010310320-0311101210212033-0300203310032130-3213000321312030-2310011210313111-0202021021123102"></a>

## Next pages — http_access / 000301112113 / 5

- [Property reference](resources--crl--reference--group-001.md#canonical-0022012330331322-3131111112231003-3101103003200200-3211113200221201-0100030210222013-3122120001201231-2010332202003001-0013313221310300)
- [xcsh_crl](../resources/crl.md#canonical-2112221203101103-2313113220020333-1301313220130233-1101313202220020-1222121022300210-0123222110132200-2221301201331022-0001302303100010)

<a id="canonical-0220030102113332-1101303210023103-1222122101230020-2033222222320013-2030332300202012-2100200130323302-1102311303330322-0322323212130002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030320012120211-0202032210223313-0023012120213131-2211213303210030-1331002232010302-1220201312232030-2231201031231233-3023132211223110"></a>

## timeouts — timeouts / 303301330203 / 2

Breadcrumbs:

- [xcsh_crl](../resources/crl.md#canonical-2112221203101103-2313113220020333-1301313220130233-1101313202220020-1222121022300210-0123222110132200-2221301201331022-0001302303100010)
- [Property reference](resources--crl--reference--group-001.md#canonical-0022012330331322-3131111112231003-3101103003200200-3211113200221201-0100030210222013-3122120001201231-2010332202003001-0013313221310300)
- timeouts

<a id="canonical-2201033001031002-0203000121033033-3112022330031210-0112322230121303-2133233120001222-1032302213333213-3131333323110003-2013122200030031"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3030002001230122-2122022211002130-1021131302210332-0003100130131031-1311100320003102-0213233331200223-0232233113323232-3133112002221131"></a>

## Direct properties — timeouts / 303301330203 / 3

<a id="canonical-0120213213300011-1020011033202201-0102202301022200-2231121000302233-1200200202211032-3213021330222232-2002200201330212-2100001032033310"></a>

<a id="canonical-2332321130030032-3001331230303213-1212330130213322-3333313202011030-0222131132221301-1203203033231102-1113301111120333-0031301221032200"></a>

## create property — timeouts / 303301330203 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3213202133322013-0201203010033102-1123233000202030-3330223220033022-2110121120233110-3311111010232230-2203002130012213-0033031002303222"></a>

<a id="canonical-0231002000123302-3333232023110131-2101010222130022-2011003100322120-3032200121131310-1010111030131312-0121113312000221-1301321203231320"></a>

## delete property — timeouts / 303301330203 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-3220000030300200-1100331230332010-0100301322021021-0111222230321300-2130222202101332-2012030333020103-1203203013101120-2230023021001110"></a>

<a id="canonical-1011020001303030-0022233202022131-0332210233321201-1020222010302123-0031123320223222-3333111331233202-0232112001300331-1120310112000330"></a>

## read property — timeouts / 303301330203 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3130031231121203-1232002031132130-2103133102212222-0302233222300233-2333311131221101-1122221301122202-0301230022100300-0311223312300222"></a>

<a id="canonical-0012023111020022-2122213112010013-1232112031313002-2312232023103300-0000100001032202-1101131001312301-3232131003301030-3012312030213322"></a>

## update property — timeouts / 303301330203 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3332302110011000-1011310003011033-2332022231300303-1231031000002320-2021301033323221-1001123012001133-1333021022121121-1300210131200321"></a>

## Next pages — timeouts / 303301330203 / 8

- [Property reference](resources--crl--reference--group-001.md#canonical-0022012330331322-3131111112231003-3101103003200200-3211113200221201-0100030210222013-3122120001201231-2010332202003001-0013313221310300)
- [xcsh_crl](../resources/crl.md#canonical-2112221203101103-2313113220020333-1301313220130233-1101313202220020-1222121022300210-0123222110132200-2221301201331022-0001302303100010)
