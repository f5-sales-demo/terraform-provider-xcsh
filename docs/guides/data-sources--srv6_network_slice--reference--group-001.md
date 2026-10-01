---
page_title: "xcsh_srv6_network_slice reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_srv6_network_slice reference."
---

# xcsh_srv6_network_slice reference

<a id="canonical-3032223200202320-2101231011221211-2032122223212321-3111010311120030-0120223232313121-1201110023330122-3333233103131300-1301103122200231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102210212312321-2012011311003232-2201032121012001-3301221030222313-0112100301102311-2331023311123212-2320032223103223-1132102321331021"></a>

## Property reference — Property reference / 233130131010 / 2

Breadcrumbs:

- [xcsh_srv6_network_slice](../data-sources/srv6_network_slice.md#canonical-0001322321013132-3303233021210322-1322230110201110-3330310302210111-3123030332122130-1230012021313000-1121212210223211-3123021213301223)
- Property reference

<a id="canonical-1002020313031321-3231320201103200-1322231202002002-0113002020333001-3021133113010233-2222313022331221-3213100110200111-3010012120030332"></a>

## Direct properties — Property reference / 233130131010 / 3

<a id="canonical-2200002003030021-2230101012132231-1303021101223122-1130020233233211-1001330233233202-1101023112101220-3011012331120230-1020102120202112"></a>

<a id="canonical-1011232321221010-3231322112310020-2010231222333201-0110230210211002-1111031110323132-0323321131212220-2222120220130112-2031203001330131"></a>

## annotations property — Property reference / 233130131010 / 4

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
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

<a id="canonical-2302233023222232-3203221000211023-2323301103121330-2110210213320111-1312331100131031-2012313110103131-3312100310331200-3233003111303002"></a>

<a id="canonical-2112201330120100-1200223033300123-0101200221122210-2012202120333333-3321031030133100-0300311110110210-0300233010303311-2210010131132102"></a>

## connect_to_access_networks property — Property reference / 233130131010 / 5

Type: `"bool"`. Computed.

Connect all SRv6 Virtual Networks in this slice to their corresponding access networks by importing
route targets specified in the virtual network.

Upstream description:

Connect all SRv6 Virtual Networks in this slice to their corresponding access networks by importing
route targets specified in the virtual network.

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

<a id="canonical-3231321222110310-3022330023203013-0012030120111212-1012202131103103-2233322112332302-0321123223302012-1131311022032130-1303123000222212"></a>

<a id="canonical-2230010011301221-3331230020211020-3203301120133231-1111331332311211-0131201322222222-2132223332131321-2132201102332002-0331131223310112"></a>

## connect_to_enterprise_networks property — Property reference / 233130131010 / 6

Type: `"bool"`. Computed.

Connect all SRv6 Virtual Networks in this slice to their corresponding enterprise networks by
importing route targets specified in the virtual network.

Upstream description:

Connect all SRv6 Virtual Networks in this slice to their corresponding enterprise networks by
importing route targets specified in the virtual network.

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

<a id="canonical-0231130213032103-2101232221200101-0112012000231132-1133331112223022-3020303113030212-0001100013322100-2001233122312232-1232111032230201"></a>

<a id="canonical-2300301033310302-2232211303000320-3110302211321100-3202233330120320-1212201220310033-1103302221032112-2331300211010222-0020203012322200"></a>

## connect_to_internet property — Property reference / 233130131010 / 7

Type: `"bool"`. Computed.

Connect all SRv6 Virtual Networks in this slice to the Internet by importing route targets specified
in the virtual network.

Upstream description:

Connect all SRv6 Virtual Networks in this slice to the Internet by importing route targets specified
in the virtual network.

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

<a id="canonical-1323023221020132-2333310112333023-3301313120210100-2111033113312110-3231133331132013-1012122220230300-3333012120211121-3201310200211020"></a>

<a id="canonical-3312312130312300-0333010002132130-3103033112331301-0213031332033233-2303301223110330-2000210202030301-2133033322023321-0011031221310332"></a>

## description property — Property reference / 233130131010 / 8

Type: `"string"`. Computed.

Description of the Srv6NetworkSlice.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1211220222030013-3110201212203332-0323013311003101-0202022121111322-0000310133110310-0331220233231331-2222320330120232-3030021011022021"></a>

<a id="canonical-3333312300030120-1223211120023230-2231200123003100-0231033332331003-1003031112333103-2233013332133021-0132100321012132-3120110311111001"></a>

## ID property — Property reference / 233130131010 / 9

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2313303023303132-1030003032323233-0301020110010113-3122321003122123-3200330113131230-0032220312130002-1001201033233323-1213303232001003"></a>

<a id="canonical-1330023003320113-1030212323101022-3312023223011121-1300200200201321-3112231220330013-3003030002210110-2100201030131131-1223211101122011"></a>

## labels property — Property reference / 233130131010 / 10

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

<a id="canonical-3312200023233020-2123100322031212-3001100110313332-0112310300100102-2010303311201231-1200311021102230-1210301130013121-0002131231021223"></a>

<a id="canonical-3020020130021120-3223312102022212-2001201220022020-2102300302120123-1031101110030323-0111000300032131-2331023101301122-3031032021232031"></a>

## name property — Property reference / 233130131010 / 11

Type: `"string"`. Required.

Name of the Srv6NetworkSlice.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2122210222020321-0212000212102031-3023300200003220-3120121312010012-3102320221310110-0222131233130222-0131212211231130-1213201210003122"></a>

<a id="canonical-1102301231223203-0103110202300032-3221012300012331-3133130233122202-3232233113100022-3131131212020210-1201033300201202-2011310231313013"></a>

## namespace property — Property reference / 233130131010 / 12

Type: `"string"`. Optional, Computed.

Namespace where the Srv6NetworkSlice exists.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3001121331023302-2211033312120331-2101232102210132-2103301312323030-1321132100000331-2021113130132232-1132300120210123-3300211210303113"></a>

<a id="canonical-2132013300002321-2201232020021200-1121003202203011-2021133130330012-1110312203023203-2103222123002231-3030210321313131-1102100220032322"></a>

## sid_prefixes property — Property reference / 233130131010 / 13

Type: `["list", "string"]`. Computed.

SID Locator from the prefix is allocated automatically for each node in each site.

Upstream description:

A SID Locator from the prefix is allocated automatically for each node in each site.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.max_ip_prefix_length": "32",
    "ves.io.schema.rules.repeated.items.string.min_ip_prefix_length": "32",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.max_ip_prefix_length": "32",
    "ves.io.schema.rules.repeated.items.string.min_ip_prefix_length": "32",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2203203233100133-1303030001103010-3223313112101011-1010002233212301-1211313302130312-1201312323202121-0022331132111321-3122121233111200"></a>

## All schema paths — Property reference / 233130131010 / 14

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--srv6_network_slice--reference--group-001.md#canonical-2200002003030021-2230101012132231-1303021101223122-1130020233233211-1001330233233202-1101023112101220-3011012331120230-1020102120202112) |
| `connect_to_access_networks` | [connect_to_access_networks](data-sources--srv6_network_slice--reference--group-001.md#canonical-2302233023222232-3203221000211023-2323301103121330-2110210213320111-1312331100131031-2012313110103131-3312100310331200-3233003111303002) |
| `connect_to_enterprise_networks` | [connect_to_enterprise_networks](data-sources--srv6_network_slice--reference--group-001.md#canonical-3231321222110310-3022330023203013-0012030120111212-1012202131103103-2233322112332302-0321123223302012-1131311022032130-1303123000222212) |
| `connect_to_internet` | [connect_to_internet](data-sources--srv6_network_slice--reference--group-001.md#canonical-0231130213032103-2101232221200101-0112012000231132-1133331112223022-3020303113030212-0001100013322100-2001233122312232-1232111032230201) |
| `description` | [description](data-sources--srv6_network_slice--reference--group-001.md#canonical-1323023221020132-2333310112333023-3301313120210100-2111033113312110-3231133331132013-1012122220230300-3333012120211121-3201310200211020) |
| `id` | [id](data-sources--srv6_network_slice--reference--group-001.md#canonical-1211220222030013-3110201212203332-0323013311003101-0202022121111322-0000310133110310-0331220233231331-2222320330120232-3030021011022021) |
| `labels` | [labels](data-sources--srv6_network_slice--reference--group-001.md#canonical-2313303023303132-1030003032323233-0301020110010113-3122321003122123-3200330113131230-0032220312130002-1001201033233323-1213303232001003) |
| `name` | [name](data-sources--srv6_network_slice--reference--group-001.md#canonical-3312200023233020-2123100322031212-3001100110313332-0112310300100102-2010303311201231-1200311021102230-1210301130013121-0002131231021223) |
| `namespace` | [namespace](data-sources--srv6_network_slice--reference--group-001.md#canonical-2122210222020321-0212000212102031-3023300200003220-3120121312010012-3102320221310110-0222131233130222-0131212211231130-1213201210003122) |
| `sid_prefixes` | [sid_prefixes](data-sources--srv6_network_slice--reference--group-001.md#canonical-3001121331023302-2211033312120331-2101232102210132-2103301312323030-1321132100000331-2021113130132232-1132300120210123-3300211210303113) |

<a id="canonical-1003301030011021-2222203222231113-1320111020200322-0320302123233201-1221302121100221-2000313112211013-1023110301123323-1131133012201112"></a>

## Next pages — Property reference / 233130131010 / 15

- [xcsh_srv6_network_slice](../data-sources/srv6_network_slice.md#canonical-0001322321013132-3303233021210322-1322230110201110-3330310302210111-3123030332122130-1230012021313000-1121212210223211-3123021213301223)
