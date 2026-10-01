---
page_title: "xcsh_aws_vpc_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_aws_vpc_site reference."
---

# xcsh_aws_vpc_site reference

<a id="canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132102232322233-0003103031122330-3100122000122033-1322121221121010-0123202322212203-2032233112022310-1022000131323311-0020222033023331"></a>

## Property reference — Property reference / 202302133313 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- Property reference

<a id="canonical-1012122110222300-2333120211300000-3121122230320230-1100011320212223-3301121213212201-0203321332233332-1321302100301022-3210322213110321"></a>

## Direct properties — Property reference / 202302133313 / 3

<a id="canonical-1020030232300332-2102110220322033-1301311211321121-2223103310302321-2110011101212201-1001131011033000-3313112103201121-0312233011032303"></a>

<a id="canonical-0323001232031033-1233031010101113-0013003210110332-3032201212133212-3100120321023230-2323101220003120-0231000120221232-1302110231030120"></a>

## address property — Property reference / 202302133313 / 4

Type: `"string"`. Optional, Computed.

Site's geographical address that can be used to determine its latitude and longitude.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [admin_password](resources--aws_vpc_site--reference--group-001.md#canonical-3223101320122233-1021302221332223-3131232122301223-3320302320210222-3320200112020003-0220131333312212-0231210310202123-1201331102030210): complete subsection reference.

<a id="canonical-3310202000221013-1010030010213322-0320133323023200-2311121320233333-2321222030212302-0313203310320220-2221311003012120-1102001010111302"></a>

<a id="canonical-0130123302301320-2213131232033101-0332212030031100-1222113023212002-2110223231103200-0212031111131000-1212220230000133-0211020123021300"></a>

## annotations property — Property reference / 202302133313 / 5

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

- [aws_cred](resources--aws_vpc_site--reference--group-001.md#canonical-3302022301010023-1332313033022123-3031323333133121-0322223302010331-0111002200002201-0313231300333023-0112200310110233-3023230320012221): complete subsection reference.

<a id="canonical-3212211012011003-0232210023311133-2030323011320113-2122032032031311-3221132001102323-0201133221022130-3302031103032303-2023310031222320"></a>

<a id="canonical-2102332212300121-0023023111202220-3012302110102331-0223323202110220-0231031220031332-1121123302112012-3131220133200030-1101112112303222"></a>

## aws_region property — Property reference / 202302133313 / 6

Type: `"string"`. Required.

AWS Region. Name for AWS Region.

Upstream description:

Name for AWS Region.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [block_all_services](resources--aws_vpc_site--reference--group-001.md#canonical-1000312122020333-1132302021220132-2220121211221022-3100303102333101-0020200023202000-3031020300221121-2210322002001020-3133330030320221): complete subsection reference.

- [blocked_services](resources--aws_vpc_site--reference--group-001.md#canonical-3323032222210001-3302120102203102-3210021101231303-3023332013331112-3200032312323130-3223301102110203-3100201032112113-1310123003122330): complete subsection reference.

- [coordinates](resources--aws_vpc_site--reference--group-002.md#canonical-3103301232022012-1202210120113233-3231000231313300-1032103030032313-3210112321011201-2112332333203011-1330133221022222-3221223133312113): complete subsection reference.

- [custom_dns](resources--aws_vpc_site--reference--group-002.md#canonical-0021301331020010-2103113020103011-0313112212033023-3001103200122333-1101111311000032-3120100021023322-0133210203122003-3032032330221333): complete subsection reference.

- [custom_security_group](resources--aws_vpc_site--reference--group-002.md#canonical-1123132203210311-2301200320231021-3232023100201023-0221020320203100-1321322223023130-1310113030201022-1201331103332123-0222223131322032): complete subsection reference.

- [default_blocked_services](resources--aws_vpc_site--reference--group-002.md#canonical-1023131220002132-1001131221323100-0123021223123002-2100211102320210-1312301332031313-1133210130030320-0102012300311331-0022310331100331): complete subsection reference.

<a id="canonical-1311131223100223-0212331222213323-3133000201130020-2330230122331010-0222033302010301-0320122200303001-2131030113131213-0033021011000201"></a>

<a id="canonical-0000231322313303-2102313221001202-0010233212110313-0132113201121023-0032111221321221-2120132313012233-1123212312200230-0122011021221203"></a>

## description property — Property reference / 202302133313 / 7

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

- [direct_connect_disabled](resources--aws_vpc_site--reference--group-002.md#canonical-1203233103200121-2313321031323013-0001112100021212-1302133112121133-1031101231223200-2232021210002031-1311100030200230-3200132220200320): complete subsection reference.

- [direct_connect_enabled](resources--aws_vpc_site--reference--group-002.md#canonical-0130012221000232-1311333022303301-0020012113300213-0322011110120320-0313333031011211-1310113213303030-0201300310003031-3011331001021112): complete subsection reference.

<a id="canonical-1121111313220230-3121032123310033-2030100131311122-1200020031010020-1013332011100212-2323120010233311-3310000321201012-1322212110113300"></a>

<a id="canonical-0103313312133321-0031203200212222-3313323202222332-1220330101202322-1211120320021110-1123130110303212-0211222301032110-1321101112001003"></a>

## disable property — Property reference / 202302133313 / 8

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

- [disable_encryption](resources--aws_vpc_site--reference--group-002.md#canonical-1202020332313000-1201233220303020-0132102003011220-1001221330331021-2013110031202200-3000002230300223-3313121333113121-0301333223301031): complete subsection reference.

- [disable_internet_vip](resources--aws_vpc_site--reference--group-002.md#canonical-1000122022313012-0001101010202330-0013311322231203-1310123033123032-2000220333300012-0323120330133221-0332121010210123-3232302323030130): complete subsection reference.

<a id="canonical-0222222312000030-3230102212302203-0023333112220132-1133202010212320-1121231001132133-1030033221013230-1222310021322321-3323131231030132"></a>

<a id="canonical-2333230310123001-2023213012122101-2120031130322322-2002103023111122-3121003322210300-0202313123312231-2012300312220233-1100122200122231"></a>

## disk_size property — Property reference / 202302133313 / 9

Type: `"number"`. Optional, Computed.

Disk size to be used for this instance in GiB. 80 is 80 GiB.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(2048),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "2048"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "2048"
  }
}
```

- [egress_gateway_default](resources--aws_vpc_site--reference--group-002.md#canonical-1212023023013013-1212222311211002-3121310102313131-0302032221222312-3030030032230033-3321120132111211-1111123202013032-2223010303001021): complete subsection reference.

- [egress_nat_gw](resources--aws_vpc_site--reference--group-002.md#canonical-2132333123021112-2121130110202202-3232330110112120-1023112123222122-2013332222002322-0212030223100002-1131133120313012-3111311320301310): complete subsection reference.

- [egress_virtual_private_gateway](resources--aws_vpc_site--reference--group-002.md#canonical-0113110013033313-2210303123332123-0133200222112102-1120033121100103-1132200021011211-3010030013230211-0210122122123323-0310322011033303): complete subsection reference.

- [enable_encryption](resources--aws_vpc_site--reference--group-002.md#canonical-3122102032303111-0233131232202220-1001102333222013-2102002002231000-3013103312333313-2003211130120332-1010200300312332-3210020213022122): complete subsection reference.

- [enable_internet_vip](resources--aws_vpc_site--reference--group-002.md#canonical-3302122201112333-3332110213202311-1021212130302332-1120320210103133-1101020320200200-2103031301023231-1200332010031321-0213300302110322): complete subsection reference.

- [f5_orchestrated_routing](resources--aws_vpc_site--reference--group-002.md#canonical-0233100333233221-1330000203320002-0303020331312332-0103212301123033-1132021203111231-2112210122320010-2313131333302212-2132032101022030): complete subsection reference.

- [f5xc_security_group](resources--aws_vpc_site--reference--group-002.md#canonical-0122030011102010-3122030030122331-3301230100222300-0332133333222221-0201113331023131-0102122321211202-2132001103112213-0230321112200201): complete subsection reference.

<a id="canonical-3223201000111001-2121312111313200-1313012202310001-2130113212033313-0113033102022110-0023112213122210-1112323023102223-0210131231212010"></a>

<a id="canonical-0212123033230233-0313212323333123-3232320130200230-2331113120331222-3303322223103112-2220110123301313-0102130233121101-0113030212001233"></a>

## ID property — Property reference / 202302133313 / 10

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022): complete subsection reference.

- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-0333332022003321-0130333001220102-3300331112302201-3222032002202120-2012011102331101-3213012110100103-2022011031100003-2201121222303323): complete subsection reference.

<a id="canonical-0213322100120220-2001001232303031-3230223212230030-2201102133330000-1220202030301322-2212210122333313-0020101233010300-0322111311321210"></a>

<a id="canonical-3110030113023130-1232233121310003-1022203203333231-0201120320013333-2002100321233123-2213112310233320-2021220033031130-3300112021300020"></a>

## instance_type property — Property reference / 202302133313 / 11

Type: `"string"`. Required.

Select Instance size based on performance needed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [kubernetes_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-3013301332331333-2211212210321032-3013100011321103-0221022231110303-0303221221101223-1102301300233020-0322322133203301-3113001020112122): complete subsection reference.

<a id="canonical-2303200000223303-1331320202232200-0022302322102123-2213032132011203-1300103331100231-1002300102321213-3310230202310012-2103123133113302"></a>

<a id="canonical-0313301030131203-0021333212333323-0200013233100202-1110021022032113-0331220003023322-0121033203021110-3323212322022331-1020231101200323"></a>

## labels property — Property reference / 202302133313 / 12

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

- [log_receiver](resources--aws_vpc_site--reference--group-004.md#canonical-3220233101332201-0221100213122133-2100201232213330-1302330031123033-1031032112102021-3331322021100121-1222111221211231-3331002202001201): complete subsection reference.

- [logs_streaming_disabled](resources--aws_vpc_site--reference--group-004.md#canonical-2032020211212221-2302310202003310-0312213011010303-2213221031020102-0022211032123131-0011101032031322-0302123012333120-1132100112302201): complete subsection reference.

- [manual_routing](resources--aws_vpc_site--reference--group-004.md#canonical-2332333023332232-3133011010233103-0013221302323103-2213132121331123-0331101321320212-1032203101301301-2232301020321122-2200320322313121): complete subsection reference.

<a id="canonical-0322301223312321-0030101200221233-1122222133302012-0233321210122233-2000131322032020-0012012322112220-0103302211202222-2303131023103233"></a>

<a id="canonical-0322112132311112-0103022323020101-1230330103033113-2132232110000333-2232212220112231-0110321203200330-3201300023133013-2202121122232031"></a>

## name property — Property reference / 202302133313 / 13

Type: `"string"`. Required.

Name of the AWS VPC Site. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1013131012130110-2022021211311101-3300003102330301-3121013302310023-3101312200130312-3310202012321130-2230102301320211-2133330123321133"></a>

<a id="canonical-1210103021021213-0333322300213120-0112323330220031-1121211100021002-1023013022011212-2122201132111112-2002000210302302-2111230130301332"></a>

## namespace property — Property reference / 202302133313 / 14

Type: `"string"`. Required.

Namespace where the AWS VPC Site is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
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

- [no_worker_nodes](resources--aws_vpc_site--reference--group-004.md#canonical-0323111213120210-3330110331133331-1123320133233101-0112302203111100-3323311230111302-2322033010012210-0303221110112120-1110302200110131): complete subsection reference.

<a id="canonical-3232210103132020-3020113020222012-3122131000321021-1031203301330033-2020130333332332-2213031031300311-0132132001223132-0013031200000010"></a>

<a id="canonical-3113330020021313-2210012203332112-3313313121301002-2112310130103120-0033130233223331-0002123030011302-3003012330110333-2020300101132000"></a>

## nodes_per_az property — Property reference / 202302133313 / 15

Type: `"number"`. Optional, Computed.

Exclusive with \[no\_worker\_nodes total\_nodes\] Desired Worker Nodes Per AZ. Max limit is up to
21.

Upstream description:

Exclusive with \[no\_worker\_nodes total\_nodes\] Desired Worker Nodes Per AZ. Max limit is up to
21.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 21),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 21,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "21"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "21"
  }
}
```

- [offline_survivability_mode](resources--aws_vpc_site--reference--group-004.md#canonical-2230331302030233-1231312011331200-0330011122330231-1003002230012012-3322102030022203-0101330012200101-1001211032031010-0121023120133132): complete subsection reference.

- [os](resources--aws_vpc_site--reference--group-004.md#canonical-3121332323311033-1213202021203320-2022333333100031-2121201303231310-0210311231023010-0313032021023311-3331022231220210-2100330111200133): complete subsection reference.

- [private_connectivity](resources--aws_vpc_site--reference--group-004.md#canonical-1033321313213300-3312102121203101-2023112202031312-2002230233131030-3320121000233132-3120322133021211-2332131021231000-2321013121233001): complete subsection reference.

<a id="canonical-1122121130130212-1311012323033320-0301223031313322-0111121131313332-1221301300231331-0001200100130330-2012112012303313-0220300300031003"></a>

<a id="canonical-3112301203110001-1000313321323000-3003002222010012-0313223310203230-3332000120020131-1210322023313222-0303021001311303-0100010330002213"></a>

## ssh_key property — Property reference / 202302133313 / 16

Type: `"string"`. Required.

Public SSH key. Public SSH key for accessing the site.

Upstream description:

Public SSH key for accessing the site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [sw](resources--aws_vpc_site--reference--group-004.md#canonical-1320213313311232-0331111200330103-3031121112123021-0310133212030023-0022311002322333-1033131021322001-3013223002132330-3033332120222200): complete subsection reference.

<a id="canonical-1011103100112331-1331132320312123-0033031220130201-3111121133202100-0021122122233222-0301031013223322-1303302132031113-0112223320103223"></a>

<a id="canonical-1231010212301232-3233122030200200-2300300210003000-2302133223033003-1101000033231333-2321013131100133-2122200101000023-0311201322222120"></a>

## tags property — Property reference / 202302133313 / 17

Type: `["map", "string"]`. Optional.

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console.

Upstream description:

AWS Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in AWS console.

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
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  }
}
```

- [timeouts](resources--aws_vpc_site--reference--group-004.md#canonical-3220110111301131-0033121322200310-0033001321230012-3131113012020201-3201213331302322-0033003230131202-1221000010112302-2320021200023132): complete subsection reference.

<a id="canonical-1332013110110103-0030221222310310-2233102023323102-0331312112001230-2223011133323200-1031133321333233-2230331022333211-0320212023101300"></a>

<a id="canonical-3030303121112210-2322121102232222-1022333110001100-1223210100012131-3133130020312230-3303322323121311-0102313221320022-2032010211002013"></a>

## total_nodes property — Property reference / 202302133313 / 18

Type: `"number"`. Optional, Computed.

Exclusive with \[no\_worker\_nodes nodes\_per\_az\] Total number of worker nodes to be deployed
across all AZ's used in the Site.

Upstream description:

Exclusive with \[no\_worker\_nodes nodes\_per\_az\] Total number of worker nodes to be deployed
across all AZ's used in the Site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 61),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 61,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "61"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "61"
  }
}
```

- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003): complete subsection reference.

- [vpc](resources--aws_vpc_site--reference--group-005.md#canonical-2003301201220030-2310111030110321-1023333120123110-3020131220223010-3230003013133033-1023220013310123-3332200230333020-2220002022320010): complete subsection reference.

- [waf_signatures](resources--aws_vpc_site--reference--group-005.md#canonical-2110022231300120-3130210123131302-1331212201020032-2333303013221203-0003211002203112-2311312323322011-1330000301331202-1312022031112120): complete subsection reference.

<a id="canonical-0101132320221313-2033101313312130-2011132231233011-1033123033000203-3223131220330322-0033100321000020-2200012333132210-0210031000303203"></a>

## All schema paths — Property reference / 202302133313 / 19

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address` | [address](resources--aws_vpc_site--reference--group-001.md#canonical-1020030232300332-2102110220322033-1301311211321121-2223103310302321-2110011101212201-1001131011033000-3313112103201121-0312233011032303) |
| `admin_password` | [admin_password](resources--aws_vpc_site--reference--group-001.md#canonical-1121100212113323-2122111130233300-1113011333002112-2033330023221223-0101032210331233-2020132321112001-2131131021202010-0030210212111220) |
| `admin_password.blindfold_secret_info` | [admin_password.blindfold_secret_info](resources--aws_vpc_site--reference--group-001.md#canonical-3001120211221021-1233302130003102-1123201111021011-0100011310012201-2203112132013203-1203213111110302-1230211102030122-3312111302203112) |
| `admin_password.blindfold_secret_info.decryption_provider` | [admin_password.blindfold_secret_info.decryption_provider](resources--aws_vpc_site--reference--group-001.md#canonical-2221201122112212-0001302021220113-0300333211032330-1023311000310013-2023320330321022-3230230033333013-0133201003322120-2320120211013231) |
| `admin_password.blindfold_secret_info.location` | [admin_password.blindfold_secret_info.location](resources--aws_vpc_site--reference--group-001.md#canonical-1203120311311111-0221223100121013-0211020203201013-3133323130302101-2132113322332321-3302122323100133-3002211002022200-2301312312333310) |
| `admin_password.blindfold_secret_info.store_provider` | [admin_password.blindfold_secret_info.store_provider](resources--aws_vpc_site--reference--group-001.md#canonical-2220112002202221-2112230212031223-3212312020300203-3011301003131112-1223122230312121-2312222220033232-1102031300122233-0203303202230213) |
| `admin_password.clear_secret_info` | [admin_password.clear_secret_info](resources--aws_vpc_site--reference--group-001.md#canonical-3221023332230001-1020003300323020-3202213112102111-2202032103330033-1112013320300101-3121113223213000-2321130003100211-3103023023220102) |
| `admin_password.clear_secret_info.provider_ref` | [admin_password.clear_secret_info.provider_ref](resources--aws_vpc_site--reference--group-001.md#canonical-1120000211110301-1123212222223320-0331010111303112-2312131220122330-2120120231213123-3333313311332012-1030221233123321-1030221212223320) |
| `admin_password.clear_secret_info.url` | [admin_password.clear_secret_info.url](resources--aws_vpc_site--reference--group-001.md#canonical-2331131323113233-2100331232300103-0312230302131200-1200310320002003-1110300223321211-1303131323210233-2322312223111203-0112101033020322) |
| `annotations` | [annotations](resources--aws_vpc_site--reference--group-001.md#canonical-3310202000221013-1010030010213322-0320133323023200-2311121320233333-2321222030212302-0313203310320220-2221311003012120-1102001010111302) |
| `aws_cred` | [aws_cred](resources--aws_vpc_site--reference--group-001.md#canonical-1132012231002000-1220330121113220-2303000101331110-3323123213011031-2223212131223321-2133030301301233-3133022000103202-1332330203111123) |
| `aws_cred.name` | [aws_cred.name](resources--aws_vpc_site--reference--group-001.md#canonical-3233102033232011-1122323112030322-3131333110133330-1230033212322232-1330210231322313-1013002220230102-1333323221203232-2223203113310320) |
| `aws_cred.namespace` | [aws_cred.namespace](resources--aws_vpc_site--reference--group-001.md#canonical-3110303213133131-2310003200220233-1303321300301203-0102110101013232-0113300201012100-2231010320311103-0300310313012121-0333313032100311) |
| `aws_cred.tenant` | [aws_cred.tenant](resources--aws_vpc_site--reference--group-001.md#canonical-3212301100203213-3300311132221233-1103133022011303-0021303201212200-3202301321123330-3210301233322110-1300203001202213-1031103132320103) |
| `aws_region` | [aws_region](resources--aws_vpc_site--reference--group-001.md#canonical-3212211012011003-0232210023311133-2030323011320113-2122032032031311-3221132001102323-0201133221022130-3302031103032303-2023310031222320) |
| `block_all_services` | [block_all_services](resources--aws_vpc_site--reference--group-001.md#canonical-1110123031032010-1113032032313120-1120000201103311-1100232123310211-0010220102213133-2010201012003033-1122010100320133-3003030022133303) |
| `blocked_services` | [blocked_services](resources--aws_vpc_site--reference--group-001.md#canonical-1101130131202131-3221200333012230-2011223031200331-1121301112001313-1033012302112113-0212212000031022-0021201022330101-3333033310132202) |
| `blocked_services.blocked_service` | [blocked_services.blocked_service](resources--aws_vpc_site--reference--group-001.md#canonical-0023123302030330-0322311221201333-0000103202310321-2202231033222030-2030001131331300-1321021221221011-3330302000010000-3032213320032013) |
| `blocked_services.blocked_service.dns` | [blocked_services.blocked_service.dns](resources--aws_vpc_site--reference--group-001.md#canonical-3122022230100220-0223310313020131-1012321033201200-0301233132331021-1131323023120033-1223323132222300-3332312232212031-3200130011221202) |
| `blocked_services.blocked_service.network_type` | [blocked_services.blocked_service.network_type](resources--aws_vpc_site--reference--group-001.md#canonical-1323211210131301-0201000123323213-0002130022223303-0220220131300311-2313001130112200-0323032231013102-0313020013232000-1013201032022223) |
| `blocked_services.blocked_service.ssh` | [blocked_services.blocked_service.ssh](resources--aws_vpc_site--reference--group-001.md#canonical-2333331022213023-3023111001033211-3132330212121203-0031210012120101-0201032220210222-1333211021022122-0110322103023030-2323202111320312) |
| `blocked_services.blocked_service.web_user_interface` | [blocked_services.blocked_service.web_user_interface](resources--aws_vpc_site--reference--group-002.md#canonical-2002121123302203-2022311031003130-2333123222102332-3212030132122221-1330122220112320-3112333332130321-3212300123122031-2221320123111211) |
| `coordinates` | [coordinates](resources--aws_vpc_site--reference--group-002.md#canonical-1021323130312312-0303132312300222-1201203202331331-2212211200002230-0032110323202321-3223123021132120-1023332000333220-0212132333123010) |
| `coordinates.latitude` | [coordinates.latitude](resources--aws_vpc_site--reference--group-002.md#canonical-2201333302232201-0130202210333331-1303012210232311-0111233330113021-2322003013200320-2332123302011311-1013033110222331-3213201113131032) |
| `coordinates.longitude` | [coordinates.longitude](resources--aws_vpc_site--reference--group-002.md#canonical-1112330333201332-1032112333323302-1102300111220221-3123031033120011-3313110102331321-1331303320322311-0213232012012320-2320132212111011) |
| `custom_dns` | [custom_dns](resources--aws_vpc_site--reference--group-002.md#canonical-3013213332301333-0002013200023233-3110132301123323-2133202030131233-2130113111122330-0332112100001230-1301220122302320-1321333130313223) |
| `custom_dns.inside_nameserver` | [custom_dns.inside_nameserver](resources--aws_vpc_site--reference--group-002.md#canonical-1322201011331112-2033222120120211-2331122210111222-1100332212010100-1303332013131010-1213102001130120-3003030313003122-1130000112333233) |
| `custom_dns.outside_nameserver` | [custom_dns.outside_nameserver](resources--aws_vpc_site--reference--group-002.md#canonical-1332000310230002-1023100102120310-3001310223313103-2113311020132102-2201000300100112-3133223220220232-0321223130012033-0320023102000201) |
| `custom_security_group` | [custom_security_group](resources--aws_vpc_site--reference--group-002.md#canonical-2102131022332001-1020300130220322-2312211301003113-0121002312130122-3213002220113120-3310120221011300-0001100321110230-3323330010321011) |
| `custom_security_group.inside_security_group_id` | [custom_security_group.inside_security_group_id](resources--aws_vpc_site--reference--group-002.md#canonical-3002101003133330-2002321230103233-2330313223120231-2012001312233020-0230202032003113-2132323103212301-1121213300103103-3310032131001313) |
| `custom_security_group.outside_security_group_id` | [custom_security_group.outside_security_group_id](resources--aws_vpc_site--reference--group-002.md#canonical-1211003022030111-2130100102100230-2022123001332111-1000300010112113-0020033113120233-1322220030113112-1021021112210120-1102030001131323) |
| `default_blocked_services` | [default_blocked_services](resources--aws_vpc_site--reference--group-002.md#canonical-0100133120312222-0222002310232213-0233220031220330-1112322001310302-3112301002102313-2311111022202021-3000323303112021-0022130113232302) |
| `description` | [description](resources--aws_vpc_site--reference--group-001.md#canonical-1311131223100223-0212331222213323-3133000201130020-2330230122331010-0222033302010301-0320122200303001-2131030113131213-0033021011000201) |
| `direct_connect_disabled` | [direct_connect_disabled](resources--aws_vpc_site--reference--group-002.md#canonical-3320101330223301-0133313121033102-0021020033131303-1310100231330013-0322313322230132-1300222213333202-1223300002202120-3120211300333012) |
| `direct_connect_enabled` | [direct_connect_enabled](resources--aws_vpc_site--reference--group-002.md#canonical-3203330010110310-3212301011331311-0013233121301300-2022013203202310-1021123100233010-1001121303020030-2302020110113302-0110010113332001) |
| `direct_connect_enabled.auto_asn` | [direct_connect_enabled.auto_asn](resources--aws_vpc_site--reference--group-002.md#canonical-3331131303110202-3321011330320200-0210131122230211-0202201003310001-0321100133213100-2232031302130210-1112030321333122-1311203022130311) |
| `direct_connect_enabled.custom_asn` | [direct_connect_enabled.custom_asn](resources--aws_vpc_site--reference--group-002.md#canonical-0010313113130101-1211300303113110-3220230211101013-0223231111332230-2230220321232032-1230013320110032-3332323301302212-1033121111121120) |
| `direct_connect_enabled.hosted_vifs` | [direct_connect_enabled.hosted_vifs](resources--aws_vpc_site--reference--group-002.md#canonical-1101031223230312-1111322331131323-0200300203113130-0311213121330021-2022222131122201-1322231313300121-3300232030303003-2003003330333030) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect` | [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect](resources--aws_vpc_site--reference--group-002.md#canonical-0233113213130003-3222002332120202-0202020132012303-2112003001231221-1112110201301330-3030002333031002-2221100322013030-3131131322121102) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect.cloudlink_network_name` | [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect.cloudlink_network_name](resources--aws_vpc_site--reference--group-002.md#canonical-0112122122011031-3323022131331213-0303331313022013-1223321032001211-3322231132311031-1231332133311000-1000033120331200-3210000010220103) |
| `direct_connect_enabled.hosted_vifs.site_registration_over_internet` | [direct_connect_enabled.hosted_vifs.site_registration_over_internet](resources--aws_vpc_site--reference--group-002.md#canonical-2121223101203210-1230300023011221-1331300033232302-3232103023031232-3103302021222130-0323233123032021-1131200101202221-0300233202321130) |
| `direct_connect_enabled.hosted_vifs.vif_list` | [direct_connect_enabled.hosted_vifs.vif_list](resources--aws_vpc_site--reference--group-002.md#canonical-2333201102213210-0202101120310320-1132312311023331-1103032023211132-2312201301021022-0201013313100110-0132103220301112-1133323331212220) |
| `direct_connect_enabled.hosted_vifs.vif_list.other_region` | [direct_connect_enabled.hosted_vifs.vif_list.other_region](resources--aws_vpc_site--reference--group-002.md#canonical-1130002011103222-3333033110002220-1120123332111221-0212103001232233-0223231010201322-1133010312233222-3323112220012112-3030312301332033) |
| `direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region` | [direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region](resources--aws_vpc_site--reference--group-002.md#canonical-1232211021130022-3030013003322300-0332111200031210-1130302122022213-0233132203311312-2231311303110010-0031111112311012-2131031030311331) |
| `direct_connect_enabled.hosted_vifs.vif_list.vif_id` | [direct_connect_enabled.hosted_vifs.vif_list.vif_id](resources--aws_vpc_site--reference--group-002.md#canonical-0013013012120010-2333112211111230-1110111120030303-1113200101311301-0313032311131112-1320311220110310-0020011202102222-2202031200010201) |
| `direct_connect_enabled.standard_vifs` | [direct_connect_enabled.standard_vifs](resources--aws_vpc_site--reference--group-002.md#canonical-3133133203221301-1012003031223023-2323213222312002-0201022003100310-0320000333212313-2130200020000101-2320301100330200-3112312010202022) |
| `disable` | [disable](resources--aws_vpc_site--reference--group-001.md#canonical-1121111313220230-3121032123310033-2030100131311122-1200020031010020-1013332011100212-2323120010233311-3310000321201012-1322212110113300) |
| `disable_encryption` | [disable_encryption](resources--aws_vpc_site--reference--group-002.md#canonical-0323231202220003-0101221322133103-3212333233323131-1313300023012021-3011111012121102-1232330310112321-0303111313121311-2213330112232132) |
| `disable_internet_vip` | [disable_internet_vip](resources--aws_vpc_site--reference--group-002.md#canonical-1310233311202123-3210112112211322-1310221310030110-1100312013001000-1201111100100111-0013321222221030-1010101033002311-0313011231000023) |
| `disk_size` | [disk_size](resources--aws_vpc_site--reference--group-001.md#canonical-0222222312000030-3230102212302203-0023333112220132-1133202010212320-1121231001132133-1030033221013230-1222310021322321-3323131231030132) |
| `egress_gateway_default` | [egress_gateway_default](resources--aws_vpc_site--reference--group-002.md#canonical-0223132030103101-2211331202122112-3323230320200103-0323133301320000-0132102121321000-3012123233012131-3223212203022131-1303123221003130) |
| `egress_nat_gw` | [egress_nat_gw](resources--aws_vpc_site--reference--group-002.md#canonical-2310103003033300-3011023033302013-2310233303210222-2022221301112130-3031200233302223-3213330130131303-0313213120103231-3332112010330020) |
| `egress_nat_gw.nat_gw_id` | [egress_nat_gw.nat_gw_id](resources--aws_vpc_site--reference--group-002.md#canonical-0223310102302221-2313322230132331-3131100132302033-1011200300000110-1033120100131311-1301111013002321-3130133111320130-2113130323222212) |
| `egress_virtual_private_gateway` | [egress_virtual_private_gateway](resources--aws_vpc_site--reference--group-002.md#canonical-2122031301111101-2033231230130122-3331333013121230-3302201001302102-2220031020200021-1010122013112100-1233131032210133-3212033312022132) |
| `egress_virtual_private_gateway.vgw_id` | [egress_virtual_private_gateway.vgw_id](resources--aws_vpc_site--reference--group-002.md#canonical-1323032211110000-1000013300221002-2223312013223311-2020231313110032-2131320102302220-0310222202112311-2012032201313001-1302000123020212) |
| `enable_encryption` | [enable_encryption](resources--aws_vpc_site--reference--group-002.md#canonical-2030311130233011-3303201321301001-3102011303313330-3232233031320132-3000002203131313-1332111113111132-3210221230202200-2310130331201220) |
| `enable_encryption.kms_key_id` | [enable_encryption.kms_key_id](resources--aws_vpc_site--reference--group-002.md#canonical-3223130210331021-3121212000201300-1312210020210320-3221033031132001-3302200010202033-0203223202133133-2233011133020100-0113200001111031) |
| `enable_internet_vip` | [enable_internet_vip](resources--aws_vpc_site--reference--group-002.md#canonical-3321320320220100-0223310032031223-2222200212223211-0111032202012321-0331211130002101-1200202222101320-3120232022003230-3333020010213203) |
| `f5_orchestrated_routing` | [f5_orchestrated_routing](resources--aws_vpc_site--reference--group-002.md#canonical-1200010232012020-1233301121032220-0303220310331300-3022031312111003-2212022021232121-0223020131131120-1330100021330313-0013231021111023) |
| `f5xc_security_group` | [f5xc_security_group](resources--aws_vpc_site--reference--group-002.md#canonical-3111232222322323-0103012011313123-2133002030330003-2311220201201222-3222332113322223-3021300030333002-3002021233321301-2220013302333233) |
| `id` | [ID](resources--aws_vpc_site--reference--group-001.md#canonical-3223201000111001-2121312111313200-1313012202310001-2130113212033313-0113033102022110-0023112213122210-1112323023102223-0210131231212010) |
| `ingress_egress_gw` | [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3113010322002120-3232103130233031-3232200003001203-3320332313320331-2032332113301301-0323301030223020-1201010320302120-2133302011211122) |
| `ingress_egress_gw.active_enhanced_firewall_policies` | [ingress_egress_gw.active_enhanced_firewall_policies](resources--aws_vpc_site--reference--group-002.md#canonical-0302032122011233-1222313111320012-3203112003101123-1002200201121311-0001303311103010-1122211102110021-2131010122330120-2200230213213022) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies](resources--aws_vpc_site--reference--group-002.md#canonical-1200320122131210-1002030002311211-3023331333230300-2220303202013230-2233120020102331-1132101032002311-3011022120120132-0211330202133032) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.name](resources--aws_vpc_site--reference--group-002.md#canonical-2031120332230023-2122333110000210-2031301202232313-0112022003301320-2321203301233131-0321022033331233-1202330300133013-3311113232212011) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](resources--aws_vpc_site--reference--group-002.md#canonical-0123202001002311-0330333230212211-2201300131001221-0330221031330023-3323310132222031-0102211300031331-3021233333030312-3321102013221123) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](resources--aws_vpc_site--reference--group-002.md#canonical-0301322020031310-3333112111033020-0012330331100123-2222200133100033-0103221021133021-3103302201230330-0332133202223332-1320121330203012) |
| `ingress_egress_gw.active_forward_proxy_policies` | [ingress_egress_gw.active_forward_proxy_policies](resources--aws_vpc_site--reference--group-002.md#canonical-0133021223021010-3333110210310313-3022202231231010-1003032200022322-0122101201221313-1211233320010232-2331331033102323-1302312330113301) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies](resources--aws_vpc_site--reference--group-002.md#canonical-3133323030310111-0320312110121121-1031201030031222-3232302312022200-3023203020323323-3330313302321120-1301302000130232-3123222233330131) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.name` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.name](resources--aws_vpc_site--reference--group-002.md#canonical-3310001203321022-2200222100130001-0113321232300121-0330213023202102-1330113302020103-0101230121012230-3021311220223003-1020322313331033) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.namespace` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.namespace](resources--aws_vpc_site--reference--group-002.md#canonical-3203031123212100-3210010003211122-2200302030331201-0000120112001333-3332011013311102-2103232012322033-3330320231212113-2002030111212333) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.tenant` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.tenant](resources--aws_vpc_site--reference--group-002.md#canonical-3332011122133323-1213202010302110-0022102121220000-1332233211231313-1122330211300211-3320111233120132-2132221221333231-2203323021222330) |
| `ingress_egress_gw.active_network_policies` | [ingress_egress_gw.active_network_policies](resources--aws_vpc_site--reference--group-002.md#canonical-2323300230300130-3102112000222001-0102133210002323-2112003012203301-2020102230032112-2302301322310200-3230132221103101-2010111333001123) |
| `ingress_egress_gw.active_network_policies.network_policies` | [ingress_egress_gw.active_network_policies.network_policies](resources--aws_vpc_site--reference--group-002.md#canonical-3030020211202331-0131130333002003-0102121113023322-0313013112312121-2133322032323220-2013032023030033-2033233030312321-1111332313330133) |
| `ingress_egress_gw.active_network_policies.network_policies.name` | [ingress_egress_gw.active_network_policies.network_policies.name](resources--aws_vpc_site--reference--group-002.md#canonical-3002021112223022-0233102233201312-2111132021333202-3020321223230210-1230130220230231-2232310023103202-3033013212221022-2231321021020221) |
| `ingress_egress_gw.active_network_policies.network_policies.namespace` | [ingress_egress_gw.active_network_policies.network_policies.namespace](resources--aws_vpc_site--reference--group-002.md#canonical-3133101201022212-3030020200021221-3213202021011103-3132111110321031-3031323201213003-3210301133232013-0033222102113113-2013111033123322) |
| `ingress_egress_gw.active_network_policies.network_policies.tenant` | [ingress_egress_gw.active_network_policies.network_policies.tenant](resources--aws_vpc_site--reference--group-002.md#canonical-2132131332223200-3231103323221132-2310113020312232-2011111123230303-3002013330131300-0322311213033001-0202220310310323-3001011133103223) |
| `ingress_egress_gw.allowed_vip_port` | [ingress_egress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-2031303023331303-1230210220212332-3122131021131200-3023120011030231-1132313303133330-2112021000313003-2300313123322013-1210120111122002) |
| `ingress_egress_gw.allowed_vip_port.custom_ports` | [ingress_egress_gw.allowed_vip_port.custom_ports](resources--aws_vpc_site--reference--group-002.md#canonical-3312033321200302-3000333213300130-1332120100200300-3110101022001230-1120330003233030-2101103300311322-3331112002012123-1310032003231323) |
| `ingress_egress_gw.allowed_vip_port.custom_ports.port_ranges` | [ingress_egress_gw.allowed_vip_port.custom_ports.port_ranges](resources--aws_vpc_site--reference--group-002.md#canonical-0023311212111330-2302201313301202-2212132330320203-0212220300303000-0030321111113320-3000320302033230-3131122321232211-3313132113133112) |
| `ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port` | [ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-2300002120112230-3023300032230002-2300200211011203-0121333110122100-1023030012201331-2322232123320001-1121132113222132-1021310212112023) |
| `ingress_egress_gw.allowed_vip_port.use_http_https_port` | [ingress_egress_gw.allowed_vip_port.use_http_https_port](resources--aws_vpc_site--reference--group-002.md#canonical-2203003213120012-3310030112101211-3020020222012331-0231010332313120-0021120330122321-3132310103123303-2120210112303211-0222320110302001) |
| `ingress_egress_gw.allowed_vip_port.use_http_port` | [ingress_egress_gw.allowed_vip_port.use_http_port](resources--aws_vpc_site--reference--group-002.md#canonical-1001031022320320-3002300120223201-0210200302130202-3223221230202313-3300001231032312-2020100332220002-0001030322002300-3010213011022113) |
| `ingress_egress_gw.allowed_vip_port.use_https_port` | [ingress_egress_gw.allowed_vip_port.use_https_port](resources--aws_vpc_site--reference--group-002.md#canonical-3101200230133122-2203333223232300-1013032311113120-3220003120222331-3320313201121322-3102000200130103-0330111112033320-0302022202321323) |
| `ingress_egress_gw.allowed_vip_port_sli` | [ingress_egress_gw.allowed_vip_port_sli](resources--aws_vpc_site--reference--group-002.md#canonical-0121103020231110-0333011313232031-3301113121232323-1031031120002011-2102232000003001-3221132200232012-1122303210321001-2110002002201103) |
| `ingress_egress_gw.allowed_vip_port_sli.custom_ports` | [ingress_egress_gw.allowed_vip_port_sli.custom_ports](resources--aws_vpc_site--reference--group-002.md#canonical-3013102000233231-2213331300300003-3031112213300032-3321021333103022-2001310001310221-1122132221133213-2232013201232120-0100103331200013) |
| `ingress_egress_gw.allowed_vip_port_sli.custom_ports.port_ranges` | [ingress_egress_gw.allowed_vip_port_sli.custom_ports.port_ranges](resources--aws_vpc_site--reference--group-002.md#canonical-1221211132120223-1212111231021312-1311333212202121-0032200333020321-1232322231022311-3003122323133102-0212100310232013-2330000122213011) |
| `ingress_egress_gw.allowed_vip_port_sli.disable_allowed_vip_port` | [ingress_egress_gw.allowed_vip_port_sli.disable_allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-3231312210113133-0232113002203113-2011203121101312-1123020111302232-0120221103332012-1011102322200022-0123120110311103-3211103113032200) |
| `ingress_egress_gw.allowed_vip_port_sli.use_http_https_port` | [ingress_egress_gw.allowed_vip_port_sli.use_http_https_port](resources--aws_vpc_site--reference--group-002.md#canonical-0022002012310112-2133310302132230-1321130111201122-3300303330121122-2012231121333100-2103312010330013-3012211333130312-1030213300213010) |
| `ingress_egress_gw.allowed_vip_port_sli.use_http_port` | [ingress_egress_gw.allowed_vip_port_sli.use_http_port](resources--aws_vpc_site--reference--group-002.md#canonical-1323012030210213-2201001130002012-3023122230021002-0131122023020110-2022320330211321-1023020112303031-1333030111330101-3220123331211120) |
| `ingress_egress_gw.allowed_vip_port_sli.use_https_port` | [ingress_egress_gw.allowed_vip_port_sli.use_https_port](resources--aws_vpc_site--reference--group-002.md#canonical-2310132321102031-0202112311103322-1002022033003130-0022201132101101-2101230022311033-0110102323003000-1102222013000323-3111331010330131) |
| `ingress_egress_gw.aws_certified_hw` | [ingress_egress_gw.aws_certified_hw](resources--aws_vpc_site--reference--group-002.md#canonical-2312333230212132-0230232310213233-3110100031120222-2331322232312031-0013213202230020-1121001102102312-3211103232102123-1020313032212210) |
| `ingress_egress_gw.az_nodes` | [ingress_egress_gw.az_nodes](resources--aws_vpc_site--reference--group-002.md#canonical-3003133301012313-3032033301300111-1200200203021020-2112002110210322-1000310013300202-1311202321210312-0323123230330211-0322312123021303) |
| `ingress_egress_gw.az_nodes.aws_az_name` | [ingress_egress_gw.az_nodes.aws_az_name](resources--aws_vpc_site--reference--group-002.md#canonical-1321312020101020-3231202032331200-3223012231031213-2000113231220321-0211320020131023-0202021032131033-3300013002202101-1120201221121110) |
| `ingress_egress_gw.az_nodes.inside_subnet` | [ingress_egress_gw.az_nodes.inside_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-0233131333123222-0113223302223232-0112013031113333-2311110201213111-2122001020033100-1101102220122231-1030223332220231-1213021012033203) |
| `ingress_egress_gw.az_nodes.inside_subnet.existing_subnet_id` | [ingress_egress_gw.az_nodes.inside_subnet.existing_subnet_id](resources--aws_vpc_site--reference--group-002.md#canonical-0011210001122133-0200201101003023-1133030321211220-2120202333210321-0130010230023221-2321220102201321-1303323220023121-3103212231323223) |
| `ingress_egress_gw.az_nodes.inside_subnet.subnet_param` | [ingress_egress_gw.az_nodes.inside_subnet.subnet_param](resources--aws_vpc_site--reference--group-002.md#canonical-3100022321220330-3133331002031211-2231200112000112-0130121030132002-2302122020102222-3221322200110103-2002021300020212-1110332201112333) |
| `ingress_egress_gw.az_nodes.inside_subnet.subnet_param.ipv4` | [ingress_egress_gw.az_nodes.inside_subnet.subnet_param.ipv4](resources--aws_vpc_site--reference--group-002.md#canonical-1111223303020210-2012330001231302-0130300212101131-0001210311310303-3110003100331232-0110211310003111-1221110232000131-1120332132213321) |
| `ingress_egress_gw.az_nodes.outside_subnet` | [ingress_egress_gw.az_nodes.outside_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-1322311123002012-2113223002130112-3112103212302111-0110332121010032-2010302300132020-3331023033201222-0022003132301333-0120211130202331) |
| `ingress_egress_gw.az_nodes.outside_subnet.existing_subnet_id` | [ingress_egress_gw.az_nodes.outside_subnet.existing_subnet_id](resources--aws_vpc_site--reference--group-002.md#canonical-0110120312333332-2131001012113313-0100212223013022-1003013220103010-2303111333021120-2031231300230300-1233010110232110-1122313010013222) |
| `ingress_egress_gw.az_nodes.outside_subnet.subnet_param` | [ingress_egress_gw.az_nodes.outside_subnet.subnet_param](resources--aws_vpc_site--reference--group-002.md#canonical-2130311332023001-1121031021323303-1312012120030203-0123001231131330-0303030012231000-2020023312032213-1103112301102123-3033120211203223) |
| `ingress_egress_gw.az_nodes.outside_subnet.subnet_param.ipv4` | [ingress_egress_gw.az_nodes.outside_subnet.subnet_param.ipv4](resources--aws_vpc_site--reference--group-002.md#canonical-3122001330131102-3011301230010322-3223111232233120-1130131220212022-0332203310210212-1130013011131211-1011231012103222-3233312103103322) |
| `ingress_egress_gw.az_nodes.reserved_inside_subnet` | [ingress_egress_gw.az_nodes.reserved_inside_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-1112231123312333-1231213100031111-3322312323033120-2120021303210323-1333013023021223-1233300120002131-1002000231311123-2111210311111233) |
| `ingress_egress_gw.az_nodes.workload_subnet` | [ingress_egress_gw.az_nodes.workload_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-3033001111311232-0120220011102100-2331133202223112-3230033012012220-3030101301011231-0112320123310133-2130021001002200-0222212122101300) |
| `ingress_egress_gw.az_nodes.workload_subnet.existing_subnet_id` | [ingress_egress_gw.az_nodes.workload_subnet.existing_subnet_id](resources--aws_vpc_site--reference--group-002.md#canonical-3122012230113233-0030311221100131-2133312013110013-2113222233312310-2323002323320320-0212112122120002-2302231030212030-2032311030331100) |
| `ingress_egress_gw.az_nodes.workload_subnet.subnet_param` | [ingress_egress_gw.az_nodes.workload_subnet.subnet_param](resources--aws_vpc_site--reference--group-002.md#canonical-3011023113110333-0102311101311231-1203312203110203-0320321030212301-3301332021031211-3113312211012011-3203131101030310-2100300233100232) |
| `ingress_egress_gw.az_nodes.workload_subnet.subnet_param.ipv4` | [ingress_egress_gw.az_nodes.workload_subnet.subnet_param.ipv4](resources--aws_vpc_site--reference--group-002.md#canonical-2333222121010331-0002230033111032-2311133323200020-3123331311031321-2001203233320133-1023200231132320-3211101111010221-1011132132121031) |
| `ingress_egress_gw.dc_cluster_group_inside_vn` | [ingress_egress_gw.dc_cluster_group_inside_vn](resources--aws_vpc_site--reference--group-002.md#canonical-1123030020110311-1001020101123103-3131130020110131-0310233132131131-0131333203030332-3003133030123321-3303120103131103-0303313112202111) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.name` | [ingress_egress_gw.dc_cluster_group_inside_vn.name](resources--aws_vpc_site--reference--group-002.md#canonical-1212211102121000-2031322321220113-2011120320122030-3110112131213113-0333112011013001-3330100131020231-3330301322013323-0333323111113123) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.namespace` | [ingress_egress_gw.dc_cluster_group_inside_vn.namespace](resources--aws_vpc_site--reference--group-002.md#canonical-2012123330101223-3012201232002203-1313212112332021-1302230310220103-2222011313320303-3333000202121000-1031201223223033-3231020132102032) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.tenant` | [ingress_egress_gw.dc_cluster_group_inside_vn.tenant](resources--aws_vpc_site--reference--group-002.md#canonical-0222313121302203-3011100013100212-2001023202113021-0300333012112303-0103332200103301-1001111133020032-1310212322011230-1031003312130321) |
| `ingress_egress_gw.dc_cluster_group_outside_vn` | [ingress_egress_gw.dc_cluster_group_outside_vn](resources--aws_vpc_site--reference--group-002.md#canonical-3101013212321113-3221311032212131-2002302112313131-3011121210013113-3221213110111000-3231322212003002-0102233101202100-2212111100223232) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.name` | [ingress_egress_gw.dc_cluster_group_outside_vn.name](resources--aws_vpc_site--reference--group-002.md#canonical-0232013021320202-3010330120331032-3030033300231000-2232302020321303-3002302230132122-2203211103203120-0320020032013320-3111001002010033) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.namespace` | [ingress_egress_gw.dc_cluster_group_outside_vn.namespace](resources--aws_vpc_site--reference--group-002.md#canonical-1101122112313212-0301131231211021-0213120020113210-2122123201201123-2312231221223133-1131131331001030-2333103323300223-2312303003113132) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.tenant` | [ingress_egress_gw.dc_cluster_group_outside_vn.tenant](resources--aws_vpc_site--reference--group-002.md#canonical-3222013300030021-3100321122022212-1111221130132320-0010300101310002-3233000200323111-0313020020022330-3112213100020302-0113330210232003) |
| `ingress_egress_gw.forward_proxy_allow_all` | [ingress_egress_gw.forward_proxy_allow_all](resources--aws_vpc_site--reference--group-002.md#canonical-3000031011310203-1320221003003120-2013002211321033-2110323001121100-3033323231020113-2201112003221123-1013202132131220-1200202313332110) |
| `ingress_egress_gw.global_network_list` | [ingress_egress_gw.global_network_list](resources--aws_vpc_site--reference--group-002.md#canonical-2311300032211111-2101331100232333-3311330113202313-2122130120312102-1000322021122230-0111020223311000-1200200223012300-0231303013202013) |
| `ingress_egress_gw.global_network_list.global_network_connections` | [ingress_egress_gw.global_network_list.global_network_connections](resources--aws_vpc_site--reference--group-002.md#canonical-2212001122030120-3331222201023201-1232020303220000-0233333013213102-2103012012113120-0233033012210010-0102311221121311-2113031310032130) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](resources--aws_vpc_site--reference--group-002.md#canonical-1101333022333030-0121123212103002-2112223330232113-2300122301012321-1030112022120133-0012322013221232-0220022012333331-1230312100012223) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn](resources--aws_vpc_site--reference--group-002.md#canonical-0220201301001303-1003022013300213-1012313012110233-2231331200122103-0122221301330020-0320321220320020-1211210323210233-2302300000332330) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](resources--aws_vpc_site--reference--group-002.md#canonical-0113032030220201-3330010102331213-0100333200232011-1023100320120211-0033300021102303-3121103131213103-0203103111132210-2010233200332001) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](resources--aws_vpc_site--reference--group-002.md#canonical-2112103002012102-1131101103202123-3101103131102101-3213323223130321-2313100032312030-2030132113112331-2200100010210110-3202203230031302) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](resources--aws_vpc_site--reference--group-003.md#canonical-1113222021111113-3102303102131233-0200300331120030-1113030321333030-1303031101130202-0120301021322020-1331320020320230-1030100330211311) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](resources--aws_vpc_site--reference--group-003.md#canonical-1201302112113130-3321213332021112-3031103301023000-1023320011311021-2230331230203120-3211322130132001-2131101312002300-3113001033110121) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn](resources--aws_vpc_site--reference--group-003.md#canonical-2013321103011120-0232123132221100-1031303132213121-0201221230302330-1220012010030213-2313222101213222-0022010113110103-2203301210122010) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](resources--aws_vpc_site--reference--group-003.md#canonical-0210120203320230-0020321300330230-1130300331033100-0001211131201301-1200122231320012-1313302031102031-3000203331210132-3021101230103011) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](resources--aws_vpc_site--reference--group-003.md#canonical-3222202120220213-1221312102222323-3000201222323101-0022221000023001-3210300002003233-2020200020112231-3312313222200133-3010312321001001) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](resources--aws_vpc_site--reference--group-003.md#canonical-1011110031000132-1201333210033132-1103332233301032-0020232033232313-0221323212331132-3102133323230032-2103310323203233-0103310102202012) |
| `ingress_egress_gw.inside_static_routes` | [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-2211120012230130-2021312110103100-1232302222233020-3310010021022220-0101012133111113-1033020011230112-1130311223020211-3002320203331133) |
| `ingress_egress_gw.inside_static_routes.static_route_list` | [ingress_egress_gw.inside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-3011003101000011-0121321013230302-0220200302231211-1032221011110130-0111201322300111-3023130020130203-3210301020012221-2201011022210032) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-1230301210303223-3233223003233321-2230133222323333-0222322010230102-0323102110130210-2002301020021330-1022220323201032-0033211330233231) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.attrs` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.attrs](resources--aws_vpc_site--reference--group-003.md#canonical-1301322101003303-2013220201332201-3032212202300123-3130323302013312-3013130233121010-3333310121032030-1323122230232001-2133101200121332) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels](resources--aws_vpc_site--reference--group-003.md#canonical-2121223100111322-0013010220233301-2221023003001313-2123330002310132-3202121312321333-3313120311122112-0102320303003110-0032331010003323) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-0120022331002332-3200102312222132-1301301323130221-1011200020103211-2122120121022333-2111121220123113-2111131233311030-2332011101211201) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--aws_vpc_site--reference--group-003.md#canonical-3232102000110311-0122111333100300-2230311101320131-0120000121202121-1300213330311310-2010032322201010-1231313021310212-0031323112021333) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](resources--aws_vpc_site--reference--group-003.md#canonical-1213332030311030-2111212221332031-2200202003233100-1131121100213120-0100010120113312-1022132131212221-0130023220330200-2011133233210100) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](resources--aws_vpc_site--reference--group-003.md#canonical-2331001321100220-3033030020113302-1123320010113300-1032330313120123-1101020200302112-2132003121133113-2011031310110211-2220331311320013) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](resources--aws_vpc_site--reference--group-003.md#canonical-2330313321301323-3023313310223013-3201223021300330-0100210330221310-0323030133302132-3211111020121001-2012210021300031-3232022122202113) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](resources--aws_vpc_site--reference--group-003.md#canonical-0112021300202200-2213021330010311-0102101031100131-0112333322200010-2113213021203232-2211002021100203-1220130023302233-0321110221013111) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](resources--aws_vpc_site--reference--group-003.md#canonical-0301032210323221-3132231001113302-1022000101231031-2003220301203310-1203000223203200-0201221210102320-2231111013033201-1233031002232000) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-2332300033323031-2322332230002220-2002113303330123-1221010301211222-1310112133210021-3220121102331012-0001131103120233-3100112223001200) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_vpc_site--reference--group-003.md#canonical-0311312110101002-1321223301311121-0230322023211212-0100310233301210-1021003111122201-2302213201303002-2333123212001113-3202310100003313) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--aws_vpc_site--reference--group-003.md#canonical-2221021012202230-3222222003100203-3231030121221132-0012200001020202-1010322003030322-3232221131121031-3003201032231033-2112230330033122) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](resources--aws_vpc_site--reference--group-003.md#canonical-2201233021121013-3202322313011130-3123101230300331-0001311112022233-3121032302210012-1102012313030330-0331013003302000-1323202222322312) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--aws_vpc_site--reference--group-003.md#canonical-0320013231233130-3230333203011101-1013023120110212-3031331222103332-1302002103130113-1221210013031231-3022030312220223-0220113331213003) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](resources--aws_vpc_site--reference--group-003.md#canonical-2201223221321132-2101022133023131-2310032312131022-1103200300201010-0032101001301021-0021331123031212-1103222220012210-1210002333121020) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--aws_vpc_site--reference--group-003.md#canonical-0301230120002131-3011033331002333-3332112332131221-0123202310010033-3220102121033323-2231132231032322-2112112323203020-0103233112002131) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](resources--aws_vpc_site--reference--group-003.md#canonical-1312132320100300-2232231131331120-1321001310322210-0333130330301233-0311011333312321-0232200012032030-0010202301203232-0322311310022333) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--aws_vpc_site--reference--group-003.md#canonical-1332101123202301-1023133322222210-1122013333233013-2321333220223101-1120232023021113-3012101201233003-1020011130111220-2112331321130110) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](resources--aws_vpc_site--reference--group-003.md#canonical-0331103323112320-3102232013003231-0303323323310322-3002031103130100-0203230013031113-2302021023100130-3120013102230210-3310101301330012) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.type` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.type](resources--aws_vpc_site--reference--group-003.md#canonical-1323111222111230-2102202323021313-1101132131232320-2301223022213223-0122012313030132-2102110233322303-0013111210113332-0310212303000131) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_vpc_site--reference--group-003.md#canonical-2110030202203003-1323213220311123-1310303103222132-3032020233111320-2231220012100201-3232100211302102-2110020200213022-0002211000113123) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--aws_vpc_site--reference--group-003.md#canonical-2031110202033201-0221220122323003-1203213021021031-1220222202221333-1310111011010103-3302331233303031-1023211313310122-3333030101212231) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](resources--aws_vpc_site--reference--group-003.md#canonical-0001100101112133-2130222202031113-2030320333103312-1131212032033113-0121030302133112-1301111213303010-0111310011333232-0100320103233113) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](resources--aws_vpc_site--reference--group-003.md#canonical-0110003333001000-0003033333220021-2332130233202311-1323030023110220-3032213332330211-0033230030010030-1202121313032133-0321113012331022) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--aws_vpc_site--reference--group-003.md#canonical-2022201002000120-0103323202020331-2210133202313032-3113110011211101-2010100332122221-1132113203231210-2100221121331212-0200021131301311) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](resources--aws_vpc_site--reference--group-003.md#canonical-2012131202330201-3030000032313310-1233311233102203-1020001320021110-2101001222103001-1020101010122231-3130330303121311-1121200331211221) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](resources--aws_vpc_site--reference--group-003.md#canonical-1121122111011021-1213022132303303-1130100221032312-1333121202210133-0023203122323100-2020103121012222-1122101230211231-1022213220113222) |
| `ingress_egress_gw.inside_static_routes.static_route_list.simple_static_route` | [ingress_egress_gw.inside_static_routes.static_route_list.simple_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-2323201333001302-2212031010110322-0002013222301323-2313201331300003-0322312330032200-2010133032010031-0003013321123030-1132121110122133) |
| `ingress_egress_gw.no_dc_cluster_group` | [ingress_egress_gw.no_dc_cluster_group](resources--aws_vpc_site--reference--group-003.md#canonical-3112131303332121-3110023232023132-0223013220102031-2200102333233130-3011020103012223-1201302310333023-1200122123032023-2010130221213023) |
| `ingress_egress_gw.no_forward_proxy` | [ingress_egress_gw.no_forward_proxy](resources--aws_vpc_site--reference--group-003.md#canonical-0212303303301021-1222112100211312-2002330312303032-2123232200130232-2300020112303122-0011310221101123-1100030130023220-3101011022113121) |
| `ingress_egress_gw.no_global_network` | [ingress_egress_gw.no_global_network](resources--aws_vpc_site--reference--group-003.md#canonical-1320200012001122-1020212303120121-3032021211011331-3101301100210233-3322013002112303-0222103200103220-0323000223111212-1120233120330133) |
| `ingress_egress_gw.no_inside_static_routes` | [ingress_egress_gw.no_inside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-0121003122231131-2302013320231322-0111013002111230-0011012203311312-1232331203110333-2312203211022231-2103220301311321-3010333030123303) |
| `ingress_egress_gw.no_network_policy` | [ingress_egress_gw.no_network_policy](resources--aws_vpc_site--reference--group-003.md#canonical-2010231202123310-0010000131231313-2233120101013100-0323233133201011-3003322133003211-3203223103103303-1002110212201231-2312003311000302) |
| `ingress_egress_gw.no_outside_static_routes` | [ingress_egress_gw.no_outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-2202010003222333-2313221121130112-3313213333100131-2312102021020010-0321212002033020-1113100230101001-0220121023312132-2130202210111323) |
| `ingress_egress_gw.outside_static_routes` | [ingress_egress_gw.outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-3333311300120330-0302331112121013-1203321331012102-2100011313302000-2110203231220211-2302113002302233-1323022133113331-1111032201223120) |
| `ingress_egress_gw.outside_static_routes.static_route_list` | [ingress_egress_gw.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-003.md#canonical-0331330322000311-2112233202303212-2023221320003110-2331323220210310-0030333002212030-2202210332021033-0000223113101303-0232301321230102) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-3102101030021011-1113230232321010-0000310212303300-2330330200211202-1213132110112301-1331121312331313-3111013100233002-0302321220220032) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.attrs` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.attrs](resources--aws_vpc_site--reference--group-003.md#canonical-1211021121310010-2121100222033130-1222212313130223-0122332320010110-0010231231233032-0201320002013133-3020003111122121-1122333230203320) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels](resources--aws_vpc_site--reference--group-003.md#canonical-3132222113232031-1210030102001021-1313120230303110-3030133323122233-2332030221200303-3223022112203023-2032032003302230-3130312020000303) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-003.md#canonical-3000123131123132-3002031021000133-1011310032302130-3321012112012102-0202200022112222-0013321311001032-1211021021003030-3110131230202112) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--aws_vpc_site--reference--group-003.md#canonical-1011332320112322-1203212303031311-2003321201132223-0201311223121113-0123122231133321-0111312202103203-3002113303132003-2321222100132310) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](resources--aws_vpc_site--reference--group-003.md#canonical-2112113130013313-0313021002133122-0232302322103203-2233212110103122-0213210303023200-0030010133100210-2033220321320203-0032332002311012) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](resources--aws_vpc_site--reference--group-003.md#canonical-0202320000302332-0021032221310211-1122033300330333-1231020102331203-3321123122033002-1111221130221102-1010320200121313-1220011100210211) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](resources--aws_vpc_site--reference--group-003.md#canonical-2121012332110301-3223313022123323-1013130101101211-3213003221201231-3332230220232002-1011310101012002-3202003030202132-0313302033021301) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](resources--aws_vpc_site--reference--group-003.md#canonical-1031002120323231-3133200013132312-1000130003230021-1310012113221201-3131030033300220-1220120032220112-0102021301012230-3223003301012101) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](resources--aws_vpc_site--reference--group-003.md#canonical-1023200020333112-3130330100300003-3122111112312102-3222233010120012-1321010010303033-3030011331010011-2030110301300200-2222023323211113) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-003.md#canonical-0201113313030200-1033011223223020-1331012321231122-2330322322200012-2022013010020103-1332031030311011-2102211121311202-0113331310320101) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_vpc_site--reference--group-003.md#canonical-2202331123211333-0332132330001021-3121222030022010-1120200002120110-3210121112211231-2202312122311333-2023310300223010-0021303210233220) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--aws_vpc_site--reference--group-003.md#canonical-0323311131213332-1030220001123033-1211330122332301-3030220013020003-3200030131122320-3121313310233120-3300220200332230-1133020213230030) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](resources--aws_vpc_site--reference--group-003.md#canonical-3301303210301203-0102122031111322-2113320230333130-1213112221321321-2002320313133031-2130220232031033-0201231012110113-2110211332100200) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--aws_vpc_site--reference--group-003.md#canonical-1121133120210221-2233213332213230-1131331010000013-0122013300321223-0213213012012031-3320120333002333-0233100031130131-1133000330332233) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](resources--aws_vpc_site--reference--group-003.md#canonical-1330030313303230-1311322202131231-0020120232130123-2211200202130022-0023311321231301-3321211110213202-1020110311300320-2232030210210233) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--aws_vpc_site--reference--group-003.md#canonical-0131202201220212-1301201111303122-2332030302113233-2103222011012122-2233022213120001-1330200232033313-3002213130303133-3200303301221130) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](resources--aws_vpc_site--reference--group-003.md#canonical-3312021213233231-2332300102222011-3003302302033023-2030331003202001-0122110030332000-2123221203121330-0112031203111003-2100200232130300) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--aws_vpc_site--reference--group-003.md#canonical-2030103132332031-0320022111203313-3022011021122330-3030221121312322-3001331022331223-0103123301330111-0112211330311113-1300113310300223) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](resources--aws_vpc_site--reference--group-003.md#canonical-1212322011011013-0011111101200010-2120000130020023-2310002032320113-1322332013110220-2213223320320112-3301103012212211-3232303123130202) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.type](resources--aws_vpc_site--reference--group-003.md#canonical-2030123031033313-1321311310312330-1111301123202302-1303312013123202-2312300101330001-2100310020232302-1333210032210020-2202100210130132) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_vpc_site--reference--group-003.md#canonical-0332122021213301-3100200102310321-3132232123202321-1102302002332213-1310320212230110-1311222122031103-1202130002210331-3022010333231133) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--aws_vpc_site--reference--group-003.md#canonical-1302203031321013-1022131022033022-1312122333232011-3003131320003130-1123030000120301-1222000210312312-3321121302133030-1222033013301312) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](resources--aws_vpc_site--reference--group-003.md#canonical-2103321022033000-2321031012112012-2130111330101002-0201003230220331-3022202033012301-1111112220320122-3120122331333311-0020221320000123) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](resources--aws_vpc_site--reference--group-003.md#canonical-3312111302001003-2013200330200122-3202013210013312-0131223331113210-1113222333113131-3002302132201232-3012030101302031-0200331023223113) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--aws_vpc_site--reference--group-003.md#canonical-3200323122000220-0313030100020102-3320232013110222-2220213020132221-3322220223011220-3321213230112022-2323232312311321-1313302023222103) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](resources--aws_vpc_site--reference--group-003.md#canonical-0300112001113023-1221110321202330-2032000300211230-1111332323132302-2020300311113332-3032131311332133-2030330211300100-2201323300201221) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](resources--aws_vpc_site--reference--group-003.md#canonical-1120212231232132-0323122221123132-2210020301232122-1312031201201212-3120210333201310-2332322203321312-3300323211232202-1132201311233011) |
| `ingress_egress_gw.outside_static_routes.static_route_list.simple_static_route` | [ingress_egress_gw.outside_static_routes.static_route_list.simple_static_route](resources--aws_vpc_site--reference--group-003.md#canonical-3300202110200021-3131200222231130-0003120031103123-1233000102231313-3011110121110131-0331023323203310-2302000313203113-0311000021101032) |
| `ingress_egress_gw.performance_enhancement_mode` | [ingress_egress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-003.md#canonical-2300123131010232-1123113031113220-1233211112103020-3311002221012221-0201023123110022-3123121012330220-2220002222220033-3022131101101332) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_vpc_site--reference--group-003.md#canonical-3110202022021311-3033310031210102-2331222030010020-0301031332221131-1122120032322022-0332121321232210-1222122031023221-0300323213133323) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--aws_vpc_site--reference--group-003.md#canonical-1212203211201222-1013313312303320-0000122330233310-1320333103201021-3010301132131223-1013020101230122-3333300100212011-1300301230012120) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--aws_vpc_site--reference--group-003.md#canonical-3023310121033201-2221111012031201-1112231102320022-0110211123032131-0231112332330003-1202022013132023-0302102221212220-1130100323120222) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_vpc_site--reference--group-003.md#canonical-2101012121300333-3100303213101213-3113232130023123-0203322012022211-1302132211112201-1320200212312002-1012021211031323-1100302030313321) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--aws_vpc_site--reference--group-003.md#canonical-2012131113022302-1130301202212012-2323221120112033-0331232311123332-1203112121033212-1033223032010223-3112112210031302-1102231203102003) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--aws_vpc_site--reference--group-003.md#canonical-0303232330123302-0210311322222022-2331223230312301-0301000333001230-0001122032100212-2112232233023211-0301123201210322-2110323333120210) |
| `ingress_egress_gw.sm_connection_public_ip` | [ingress_egress_gw.sm_connection_public_ip](resources--aws_vpc_site--reference--group-003.md#canonical-1033100121131301-2310222101030100-1221302231012301-2311201210222023-3221213230003131-0222223011023203-1220133201200000-3222201031320003) |
| `ingress_egress_gw.sm_connection_pvt_ip` | [ingress_egress_gw.sm_connection_pvt_ip](resources--aws_vpc_site--reference--group-003.md#canonical-3320332030031303-2001133233031333-2102032102000021-0023131223031122-3020232030130230-2132212100330232-2330200201030012-0222230112003110) |
| `ingress_gw` | [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-1213030012102200-2132122202121220-2010120201212311-1321031001321231-3302202232203113-2220103102300303-3322131013213033-1113012313110132) |
| `ingress_gw.allowed_vip_port` | [ingress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-3021003322112210-0202103121023320-0020300302212320-0232323021231221-2220311223120001-1332303313000102-2133132301011102-1001023220112003) |
| `ingress_gw.allowed_vip_port.custom_ports` | [ingress_gw.allowed_vip_port.custom_ports](resources--aws_vpc_site--reference--group-004.md#canonical-0231222022313320-3001322101013221-2133021133011331-2331233313120130-0133122312220310-1220030023223311-0203230303321221-1231232033121210) |
| `ingress_gw.allowed_vip_port.custom_ports.port_ranges` | [ingress_gw.allowed_vip_port.custom_ports.port_ranges](resources--aws_vpc_site--reference--group-004.md#canonical-0030221222012223-2323322033223131-2210020330210313-3000011333202022-1012031102131021-0333320033313210-3020002303013113-2001001131012002) |
| `ingress_gw.allowed_vip_port.disable_allowed_vip_port` | [ingress_gw.allowed_vip_port.disable_allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-0033221013102302-2323202333303221-2123211322133132-1021321220012310-3312211002323322-3201332212130010-1013333313100220-0100022332002012) |
| `ingress_gw.allowed_vip_port.use_http_https_port` | [ingress_gw.allowed_vip_port.use_http_https_port](resources--aws_vpc_site--reference--group-004.md#canonical-3313132220133323-2312312013333203-1211323133120300-0201330320103130-0310220220012032-2202112322121311-0120233333023133-1021110031032233) |
| `ingress_gw.allowed_vip_port.use_http_port` | [ingress_gw.allowed_vip_port.use_http_port](resources--aws_vpc_site--reference--group-004.md#canonical-2023330210030203-2121310211230221-1223202132221021-2032232032121323-0301301022022101-1313003211013310-2332123023022033-2033312020103200) |
| `ingress_gw.allowed_vip_port.use_https_port` | [ingress_gw.allowed_vip_port.use_https_port](resources--aws_vpc_site--reference--group-004.md#canonical-2310010113002121-2200012112303100-1232232230311301-3330031301102200-0301101322111230-1200232202203032-3222130031030202-0213233130230130) |
| `ingress_gw.aws_certified_hw` | [ingress_gw.aws_certified_hw](resources--aws_vpc_site--reference--group-003.md#canonical-0121232011233123-1311031102022212-3110300113222212-1122233132213201-0031110231313021-0103223102300312-0011300112310320-2130001131311123) |
| `ingress_gw.az_nodes` | [ingress_gw.az_nodes](resources--aws_vpc_site--reference--group-004.md#canonical-2020303030020202-0123132102230322-2002003002120201-2301021023332233-3220322333301210-2100012332111133-2333120022222301-1321010112101302) |
| `ingress_gw.az_nodes.aws_az_name` | [ingress_gw.az_nodes.aws_az_name](resources--aws_vpc_site--reference--group-004.md#canonical-3312202212100103-0100332331231020-1320111311222200-0002201111203133-1023121123200311-1331232032010230-0302323131312020-0333033232000302) |
| `ingress_gw.az_nodes.local_subnet` | [ingress_gw.az_nodes.local_subnet](resources--aws_vpc_site--reference--group-004.md#canonical-2303321002201331-1000001003002002-3011310312121331-3002310000030201-3101320321220023-1211232222302311-1110231102332323-0012210002300200) |
| `ingress_gw.az_nodes.local_subnet.existing_subnet_id` | [ingress_gw.az_nodes.local_subnet.existing_subnet_id](resources--aws_vpc_site--reference--group-004.md#canonical-3031210300322212-1222203230101121-0220033321332010-0101110033022113-3312230112101102-1323102012120201-0132221232232302-0303010132123311) |
| `ingress_gw.az_nodes.local_subnet.subnet_param` | [ingress_gw.az_nodes.local_subnet.subnet_param](resources--aws_vpc_site--reference--group-004.md#canonical-2133230210023102-1112012233010110-3213013231223000-3200310333122213-1331233001231020-1033023021223202-0010232000111213-1000220312312212) |
| `ingress_gw.az_nodes.local_subnet.subnet_param.ipv4` | [ingress_gw.az_nodes.local_subnet.subnet_param.ipv4](resources--aws_vpc_site--reference--group-004.md#canonical-2122120030230310-0113330112121332-3201303101022232-3332012103123200-2021113130310131-0110332202312132-0313120202233123-3222130001030320) |
| `ingress_gw.performance_enhancement_mode` | [ingress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-004.md#canonical-0002123232300200-1312212131002311-2100031001332320-3201100111203201-3332021031333101-0022011013231210-2102200030301322-2131033232322200) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_vpc_site--reference--group-004.md#canonical-1220023032023012-3030212213313303-0111201100320021-3123220010311013-1111332302211331-3132120321121133-0202020221310131-3200101313002010) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--aws_vpc_site--reference--group-004.md#canonical-0310233322110300-3002101000102021-2202011322322313-3303112201131311-2330123132020333-3211311010102003-0123222323012000-0320121300122012) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--aws_vpc_site--reference--group-004.md#canonical-2331200023332332-0211110030122211-0102033102122122-1112003203131222-1230120111321001-0232121113213102-2033332021221031-2322031320032203) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_vpc_site--reference--group-004.md#canonical-2132031013320230-1213330322323312-3010101333120300-1312212312013310-3322000203122012-3023123331210213-2022302111113220-2333212000113011) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--aws_vpc_site--reference--group-004.md#canonical-3131023312330111-1100000010201220-3123012203021221-1220010122231130-3112120313201331-3210330303220101-2001013021230302-0221300201221303) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--aws_vpc_site--reference--group-004.md#canonical-3202130023332213-3203201333233122-3300300100020322-3302012303122230-2102202312000122-2322231211202123-1130332102003131-2232201100211300) |
| `instance_type` | [instance_type](resources--aws_vpc_site--reference--group-001.md#canonical-0213322100120220-2001001232303031-3230223212230030-2201102133330000-1220202030301322-2212210122333313-0020101233010300-0322111311321210) |
| `kubernetes_upgrade_drain` | [kubernetes_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-2323233201000331-1113320001333220-2031013132230301-3132223121102100-0330233010300233-0333111313130310-3202020010020123-3302123301012121) |
| `kubernetes_upgrade_drain.disable_upgrade_drain` | [kubernetes_upgrade_drain.disable_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-2323122121323231-2302211031220333-2203121013110321-3002223321313120-0320320213331121-3222213203221030-0102233230303331-3003101000131312) |
| `kubernetes_upgrade_drain.enable_upgrade_drain` | [kubernetes_upgrade_drain.enable_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-3333013332001101-3220020012112110-0011012231011020-1001001222003321-3222203110031232-0202201331112110-1203002021121311-0222003101212030) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](resources--aws_vpc_site--reference--group-004.md#canonical-0032130310032231-1011131131010010-0122231201000322-3111333220233332-0131111230122103-0300021330120302-3301321211000003-1232103313332200) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count](resources--aws_vpc_site--reference--group-004.md#canonical-0303121222001030-3033100000120212-2223323213200300-2221032230120332-0213102013233311-2003313101232212-2003033331031023-0302211120313302) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage](resources--aws_vpc_site--reference--group-004.md#canonical-3320110011203230-3210303121233103-2012131122020030-2200010103002112-3101332110101332-1322132201220002-1220000323213113-1302133213110020) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout](resources--aws_vpc_site--reference--group-004.md#canonical-3020120111210331-0332312210323203-1021213011100312-3100022330131311-3331010013110203-3200223133301201-3031211021103102-0100223300200330) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](resources--aws_vpc_site--reference--group-004.md#canonical-3120110303302200-2001111223332112-1111330300002331-3201302311022113-2332201023301120-0320232201313332-2013031023103033-0312132102212133) |
| `labels` | [labels](resources--aws_vpc_site--reference--group-001.md#canonical-2303200000223303-1331320202232200-0022302322102123-2213032132011203-1300103331100231-1002300102321213-3310230202310012-2103123133113302) |
| `log_receiver` | [log_receiver](resources--aws_vpc_site--reference--group-004.md#canonical-2003213030032201-0000122320031311-3313011233322130-0100210100222012-1130101021232210-2300012013022331-3332320223010333-1031330223022230) |
| `log_receiver.name` | [log_receiver.name](resources--aws_vpc_site--reference--group-004.md#canonical-0333301110103030-1300100213013101-2131001212321332-0202033023120331-1300033220022233-1111102013120232-2120003311132120-2323112302031201) |
| `log_receiver.namespace` | [log_receiver.namespace](resources--aws_vpc_site--reference--group-004.md#canonical-1023231133123331-2221222212021130-0111112320030012-1033212331010200-2220333331013300-3311133302301130-3020022011231312-3223320120103121) |
| `log_receiver.tenant` | [log_receiver.tenant](resources--aws_vpc_site--reference--group-004.md#canonical-3332300312231233-1102220300233231-3300021312001011-3023132011030000-2323330201301200-0202211101202331-1220200220123110-3132220101033111) |
| `logs_streaming_disabled` | [logs_streaming_disabled](resources--aws_vpc_site--reference--group-004.md#canonical-2002222101023311-2002312201100012-2211123211220012-0110012100322233-0322231103202303-2230110301300200-0023321323023233-2213312131012103) |
| `manual_routing` | [manual_routing](resources--aws_vpc_site--reference--group-004.md#canonical-3020121031032231-1103013030230103-3031133330101100-2301130023123222-1233013021101120-3103232331131332-0022200132303212-1113101010212212) |
| `name` | [name](resources--aws_vpc_site--reference--group-001.md#canonical-0322301223312321-0030101200221233-1122222133302012-0233321210122233-2000131322032020-0012012322112220-0103302211202222-2303131023103233) |
| `namespace` | [namespace](resources--aws_vpc_site--reference--group-001.md#canonical-1013131012130110-2022021211311101-3300003102330301-3121013302310023-3101312200130312-3310202012321130-2230102301320211-2133330123321133) |
| `no_worker_nodes` | [no_worker_nodes](resources--aws_vpc_site--reference--group-004.md#canonical-0100330200230123-0211130123221033-1210312202010201-0333000212310210-2100131231233213-2113223121221203-2010213310032330-2231102112200001) |
| `nodes_per_az` | [nodes_per_az](resources--aws_vpc_site--reference--group-001.md#canonical-3232210103132020-3020113020222012-3122131000321021-1031203301330033-2020130333332332-2213031031300311-0132132001223132-0013031200000010) |
| `offline_survivability_mode` | [offline_survivability_mode](resources--aws_vpc_site--reference--group-004.md#canonical-1303201030303311-0220223111201320-1001213033122111-3331321022302321-0002002200100133-0000223102013323-1100133322130311-1100010331230000) |
| `offline_survivability_mode.enable_offline_survivability_mode` | [offline_survivability_mode.enable_offline_survivability_mode](resources--aws_vpc_site--reference--group-004.md#canonical-0023302230310002-0203332110132210-1322031001210032-2123102311212103-0101010112002223-2333120002200022-0203120331120103-0211130322313202) |
| `offline_survivability_mode.no_offline_survivability_mode` | [offline_survivability_mode.no_offline_survivability_mode](resources--aws_vpc_site--reference--group-004.md#canonical-0013000122232111-3213212311201302-3123020313132121-3133022302321030-0311333130031030-0231033030301012-1122331201233002-2230332103102231) |
| `os` | [os](resources--aws_vpc_site--reference--group-004.md#canonical-1211113230123003-2211300003302201-0200223302302020-3122123101012232-1101132332203223-0132213003031002-2332103112101303-0300002233221120) |
| `os.default_os_version` | [os.default_os_version](resources--aws_vpc_site--reference--group-004.md#canonical-0000022220020313-0000100001111130-1200120233120300-0213312332333122-0223331311111022-3323321003301121-0310203233203323-0000333312333030) |
| `os.operating_system_version` | [os.operating_system_version](resources--aws_vpc_site--reference--group-004.md#canonical-1233102132003121-3300310020011310-1111200213223221-2310103120323133-3112211120303020-0032200101233211-3223303101331033-2212021221302123) |
| `private_connectivity` | [private_connectivity](resources--aws_vpc_site--reference--group-004.md#canonical-3302002132231203-3030202210010120-1310320012103133-1302011000212230-3111302233111200-0123000123331023-3012122210102203-0012021311213222) |
| `private_connectivity.cloud_link` | [private_connectivity.cloud_link](resources--aws_vpc_site--reference--group-004.md#canonical-2312120313020133-0130312103011212-0120320013133101-1330101002232300-2203112210001000-1022210013230131-2310121211111312-2101302230321303) |
| `private_connectivity.cloud_link.name` | [private_connectivity.cloud_link.name](resources--aws_vpc_site--reference--group-004.md#canonical-3113030212323203-3301122231311001-2321223122330000-1322122323103202-1021010333121320-2003132103103010-0112311120231330-1231003301011012) |
| `private_connectivity.cloud_link.namespace` | [private_connectivity.cloud_link.namespace](resources--aws_vpc_site--reference--group-004.md#canonical-0111311013132000-0222013002022200-3311302100100133-3230321211102021-3011033131032221-2033123233031110-1132010031121330-1203220002120013) |
| `private_connectivity.cloud_link.tenant` | [private_connectivity.cloud_link.tenant](resources--aws_vpc_site--reference--group-004.md#canonical-1221013102031232-3213313003112021-1030131121301301-2120131321001321-0222201101021032-2310233313033031-2301132113322030-1303112102202030) |
| `private_connectivity.inside` | [private_connectivity.inside](resources--aws_vpc_site--reference--group-004.md#canonical-2123333130123022-1110110312001211-2310133230123132-2001002022120321-1230232130202110-3113022303100132-1103301210213130-1210220301310013) |
| `private_connectivity.outside` | [private_connectivity.outside](resources--aws_vpc_site--reference--group-004.md#canonical-1012102111312103-0022202011202232-0330021101113010-2111023031313032-1113100320032301-3112020202000333-1233012010113022-0313132202113100) |
| `ssh_key` | [ssh_key](resources--aws_vpc_site--reference--group-001.md#canonical-1122121130130212-1311012323033320-0301223031313322-0111121131313332-1221301300231331-0001200100130330-2012112012303313-0220300300031003) |
| `sw` | [sw](resources--aws_vpc_site--reference--group-004.md#canonical-3302121131223112-0323203102333122-3221222133232322-1031230120131200-3321012133111130-1021110211210001-2110122103301100-1303100122332200) |
| `sw.default_sw_version` | [sw.default_sw_version](resources--aws_vpc_site--reference--group-004.md#canonical-2221122213230120-0302230232210121-3022010132000012-1211133002223323-3110221233001021-1203302320013021-0311302310232123-3323110223322313) |
| `sw.volterra_software_version` | [sw.volterra_software_version](resources--aws_vpc_site--reference--group-004.md#canonical-0330331133313120-3323002110200203-1203132201001201-2003303011032223-1030220121330102-2223303020332310-0201332230130233-1022101131001203) |
| `tags` | [tags](resources--aws_vpc_site--reference--group-001.md#canonical-1011103100112331-1331132320312123-0033031220130201-3111121133202100-0021122122233222-0301031013223322-1303302132031113-0112223320103223) |
| `timeouts` | [timeouts](resources--aws_vpc_site--reference--group-004.md#canonical-3222102300020111-0333230030321122-2312130111020003-2000220130021003-3322312012111022-1323111133121021-0100312012230200-3012203030200003) |
| `timeouts.create` | [timeouts.create](resources--aws_vpc_site--reference--group-004.md#canonical-1110033210210111-3121330113031222-0211332322112222-1013012010133300-1022230022131000-1113222211333330-2130320311021120-3022022220303203) |
| `timeouts.delete` | [timeouts.delete](resources--aws_vpc_site--reference--group-004.md#canonical-0021020010033011-1113002212302021-2100102111330100-1103330300211122-1033311200023132-3200012303121133-3101331333223302-0200211210120311) |
| `timeouts.read` | [timeouts.read](resources--aws_vpc_site--reference--group-004.md#canonical-2031103132012131-0320103230022321-0110303331023210-1333133312210020-1012333233110123-2310321202033022-1023232030211320-1203013201012332) |
| `timeouts.update` | [timeouts.update](resources--aws_vpc_site--reference--group-004.md#canonical-0132331020113323-1203233103021011-0101221000332312-0321212000321321-1121031320300331-0311310101013321-3202012222312121-3232232221222113) |
| `total_nodes` | [total_nodes](resources--aws_vpc_site--reference--group-001.md#canonical-1332013110110103-0030221222310310-2233102023323102-0331312112001230-2223011133323200-1031133321333233-2230331022333211-0320212023101300) |
| `voltstack_cluster` | [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-1001100130213133-3230331301033031-0331333111030200-0210320013203113-1123312203213202-1023022223312332-2230321323121003-3130130333013112) |
| `voltstack_cluster.active_enhanced_firewall_policies` | [voltstack_cluster.active_enhanced_firewall_policies](resources--aws_vpc_site--reference--group-004.md#canonical-0012010130332031-1220002330320330-0202000120331303-2230122311323211-0312320312330212-1221331013231012-1020020230121321-2210100130202202) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies](resources--aws_vpc_site--reference--group-004.md#canonical-2032020002310320-2003132002332300-2221222300002220-0203022013223110-0000230021103023-3011303012132323-0102122213021121-3110022013311101) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.name](resources--aws_vpc_site--reference--group-004.md#canonical-3331332330313032-0302200211213333-0333013231302001-1312013213100313-0111223200020023-1023110221200111-0331221213323001-0301103230332320) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](resources--aws_vpc_site--reference--group-004.md#canonical-3203103210301012-2023332331113232-2101030303210220-1233220320230121-2230101221321012-3213022202303213-2020331013211110-2030331020300223) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](resources--aws_vpc_site--reference--group-004.md#canonical-3221211120213221-2131002232201031-3130103213210231-0110100102312300-2213210233223130-2223021320300300-2200301112213121-3032030322302330) |
| `voltstack_cluster.active_forward_proxy_policies` | [voltstack_cluster.active_forward_proxy_policies](resources--aws_vpc_site--reference--group-004.md#canonical-2102312320002120-2303022231223332-2231302200021331-3303030100001013-2331323002131100-3132302320002232-1130131021113312-0311233203212010) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies](resources--aws_vpc_site--reference--group-004.md#canonical-3113211010301112-3031120113213322-3203222223102200-1023013233111110-2133123312202101-1013310210122011-2232230103203200-1002021233230320) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.name` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.name](resources--aws_vpc_site--reference--group-004.md#canonical-2220020322201133-2320031100023201-3202013030013022-2111110220332231-3313323222101202-1100323010332012-3230212202333012-1211031012030101) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.namespace` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.namespace](resources--aws_vpc_site--reference--group-004.md#canonical-1301200023131231-0233130202331223-1223030332321203-0132201012031322-2200113330030212-3130012231221222-2333132122203022-2132101001100213) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.tenant` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.tenant](resources--aws_vpc_site--reference--group-004.md#canonical-2033301321132322-0013231112002230-1031013201202022-1310202330223123-3232121112023223-3110233300322121-3320033022121031-1123032031332120) |
| `voltstack_cluster.active_network_policies` | [voltstack_cluster.active_network_policies](resources--aws_vpc_site--reference--group-004.md#canonical-1121320310123312-2323102223033202-0330321310303233-0030000201021020-2123111132021200-1120211203020323-0100211312102321-2000020121002011) |
| `voltstack_cluster.active_network_policies.network_policies` | [voltstack_cluster.active_network_policies.network_policies](resources--aws_vpc_site--reference--group-004.md#canonical-0221220331222331-0202220101213311-3001332000200031-3212112002220013-3113220021302020-0333223013223232-2210232111010320-3200311203223201) |
| `voltstack_cluster.active_network_policies.network_policies.name` | [voltstack_cluster.active_network_policies.network_policies.name](resources--aws_vpc_site--reference--group-004.md#canonical-3000310031301300-3133221121212313-1301021100131121-2010202211031002-2332132330200200-0211211223100330-3100110221203101-3012323103130001) |
| `voltstack_cluster.active_network_policies.network_policies.namespace` | [voltstack_cluster.active_network_policies.network_policies.namespace](resources--aws_vpc_site--reference--group-004.md#canonical-0132201110223302-3202310013313031-1221111220023023-0111110011100203-3010231000222201-3222321323103332-2122321010312303-0110020102100322) |
| `voltstack_cluster.active_network_policies.network_policies.tenant` | [voltstack_cluster.active_network_policies.network_policies.tenant](resources--aws_vpc_site--reference--group-004.md#canonical-3233233030223032-1131330320010130-3113000332333232-0232303003101032-1322222212121221-3332210022010123-3312123011212022-1312010120230001) |
| `voltstack_cluster.allowed_vip_port` | [voltstack_cluster.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-2023222222123202-2120200011223002-3221210103222203-2100302221221212-3101221132001122-1100112210212012-1202322103310021-2223033321220211) |
| `voltstack_cluster.allowed_vip_port.custom_ports` | [voltstack_cluster.allowed_vip_port.custom_ports](resources--aws_vpc_site--reference--group-004.md#canonical-2000230121300102-2211113300131022-0300121110333222-3021132323020002-2200311020103233-2320231121302300-1111100023012031-0300200030230012) |
| `voltstack_cluster.allowed_vip_port.custom_ports.port_ranges` | [voltstack_cluster.allowed_vip_port.custom_ports.port_ranges](resources--aws_vpc_site--reference--group-004.md#canonical-3020121322131231-1222012313010101-1100330231031001-2010221332301033-0032023131201221-2213012033303021-2120101323100231-2323233100203322) |
| `voltstack_cluster.allowed_vip_port.disable_allowed_vip_port` | [voltstack_cluster.allowed_vip_port.disable_allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-2321200301302100-1121002122033212-1333223002310000-3333313300202103-0211332100232022-1203322113102113-1202012310202213-1323331032203210) |
| `voltstack_cluster.allowed_vip_port.use_http_https_port` | [voltstack_cluster.allowed_vip_port.use_http_https_port](resources--aws_vpc_site--reference--group-004.md#canonical-0301022122123023-2122210122132022-1232101212201301-1121001032323112-0322323301232312-0021023300301113-2032111021221010-3330030032102331) |
| `voltstack_cluster.allowed_vip_port.use_http_port` | [voltstack_cluster.allowed_vip_port.use_http_port](resources--aws_vpc_site--reference--group-004.md#canonical-2330221001303312-3220002102032100-0322001313012310-3031031033233010-3123302032233000-2232302323110323-2100001330100033-3301300331333132) |
| `voltstack_cluster.allowed_vip_port.use_https_port` | [voltstack_cluster.allowed_vip_port.use_https_port](resources--aws_vpc_site--reference--group-004.md#canonical-2333222321312300-3203000012121311-0331032230020110-2002020002332321-0232210103011301-2112320112023110-2222020101022202-2311203220020330) |
| `voltstack_cluster.aws_certified_hw` | [voltstack_cluster.aws_certified_hw](resources--aws_vpc_site--reference--group-004.md#canonical-0012212101111223-3213131022302031-2212321030322322-2023013222033333-3030121120031232-3013110230221033-1122102303311121-3332230211202022) |
| `voltstack_cluster.az_nodes` | [voltstack_cluster.az_nodes](resources--aws_vpc_site--reference--group-004.md#canonical-2120331132021020-0003101210232312-2300120213022310-0132130120312233-0033020102012102-2332222233300301-3312223212201033-2203332003012212) |
| `voltstack_cluster.az_nodes.aws_az_name` | [voltstack_cluster.az_nodes.aws_az_name](resources--aws_vpc_site--reference--group-004.md#canonical-0333033002000230-1233201120110312-3022300323000120-2032001233003100-2230133013102023-0033000312220112-0030011031013031-3020210301133011) |
| `voltstack_cluster.az_nodes.local_subnet` | [voltstack_cluster.az_nodes.local_subnet](resources--aws_vpc_site--reference--group-004.md#canonical-1030322133020120-3322322010013332-2303323320221003-0022103031331210-3300011202232223-0133103221001303-1033320311033203-0312110103220032) |
| `voltstack_cluster.az_nodes.local_subnet.existing_subnet_id` | [voltstack_cluster.az_nodes.local_subnet.existing_subnet_id](resources--aws_vpc_site--reference--group-004.md#canonical-1322021221023321-2101323323310022-3201332010311303-3013323033303021-3013311103302213-1220301310132113-3123000120223033-2023103301203030) |
| `voltstack_cluster.az_nodes.local_subnet.subnet_param` | [voltstack_cluster.az_nodes.local_subnet.subnet_param](resources--aws_vpc_site--reference--group-004.md#canonical-3111030101101210-3103133023000121-3023102023012032-0021001121001312-1123333212321001-2312001033020212-2300332013100023-2200220321032330) |
| `voltstack_cluster.az_nodes.local_subnet.subnet_param.ipv4` | [voltstack_cluster.az_nodes.local_subnet.subnet_param.ipv4](resources--aws_vpc_site--reference--group-004.md#canonical-3212013122231123-0002231331010313-0211100302022320-3222200203231203-2332333021030313-0021002220133221-3032100222111103-3111203332111100) |
| `voltstack_cluster.dc_cluster_group` | [voltstack_cluster.dc_cluster_group](resources--aws_vpc_site--reference--group-004.md#canonical-1021210323103133-0230302223310022-2113131130310211-3120302030300022-3220110320221211-1112132120120320-2300100300220033-3103200103233001) |
| `voltstack_cluster.dc_cluster_group.name` | [voltstack_cluster.dc_cluster_group.name](resources--aws_vpc_site--reference--group-004.md#canonical-1300301123101013-1012102301312003-3122110131332121-0102210201311332-3133001020222000-1203103300313030-3222023300001011-1020023031321002) |
| `voltstack_cluster.dc_cluster_group.namespace` | [voltstack_cluster.dc_cluster_group.namespace](resources--aws_vpc_site--reference--group-004.md#canonical-0313330011021321-1321103030120220-0100000303000330-3232310103333012-2311002101013221-2011113132231220-1203001230310300-2220331013200220) |
| `voltstack_cluster.dc_cluster_group.tenant` | [voltstack_cluster.dc_cluster_group.tenant](resources--aws_vpc_site--reference--group-004.md#canonical-3001233110333101-1222033022230033-3032010010220112-3223122313202201-0110031122001001-1002302030203123-1313123323130221-1011330230303302) |
| `voltstack_cluster.default_storage` | [voltstack_cluster.default_storage](resources--aws_vpc_site--reference--group-004.md#canonical-2032232203333220-3201331221322023-1120113200013131-3303100230201332-1033102032122333-0023221233011232-1130223222031202-0302221332221110) |
| `voltstack_cluster.forward_proxy_allow_all` | [voltstack_cluster.forward_proxy_allow_all](resources--aws_vpc_site--reference--group-004.md#canonical-2111332012000301-0123321313211202-3211200031003332-0030111112010031-3030211331130001-0300123332031112-0103210230230321-2031011331122230) |
| `voltstack_cluster.global_network_list` | [voltstack_cluster.global_network_list](resources--aws_vpc_site--reference--group-004.md#canonical-3201311021111031-2230113102203230-1030323011111311-1001121321231032-1110113310300220-0320301231212111-0101323303311211-3201332200010110) |
| `voltstack_cluster.global_network_list.global_network_connections` | [voltstack_cluster.global_network_list.global_network_connections](resources--aws_vpc_site--reference--group-004.md#canonical-1312113030303303-2303320113121112-0232020023132112-3023020031120303-2301123130001102-1333311231220302-3031233023023210-3000012220311032) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr](resources--aws_vpc_site--reference--group-005.md#canonical-3103233002033212-3220130333020022-2302013102032221-3133320202121031-0103212130213211-2121331232032312-2320222123123113-2130201320000202) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn](resources--aws_vpc_site--reference--group-005.md#canonical-3310102320113003-0320200231202130-1103003101113301-3023001332123103-1003211020313101-0103311221303320-3300023032121000-0301013210220331) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](resources--aws_vpc_site--reference--group-005.md#canonical-2201320031333101-0323301000011122-1111110013120300-2302212021200210-2013201033230110-1320123011223303-1331031133212121-0022200110213313) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](resources--aws_vpc_site--reference--group-005.md#canonical-3010100231033211-0033112022320232-3122020101311100-1310122230312322-3303300333321122-1112312311020023-3021312000033321-3302212113021312) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](resources--aws_vpc_site--reference--group-005.md#canonical-1300323210200231-2000033010322033-0203230220320030-1021313330230023-2301122001223221-2321300002001222-2221121310233120-2311302223313231) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr](resources--aws_vpc_site--reference--group-005.md#canonical-1021210001323230-0333331033200001-3101102200212011-2320031320102113-0031330133320013-1132230013203103-2020020200032013-3330122212023111) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn](resources--aws_vpc_site--reference--group-005.md#canonical-0100120013312331-3303301321011232-1232031032303230-1200113033302200-2312233103133122-1222313102031323-1303302131022311-2222331032003222) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](resources--aws_vpc_site--reference--group-005.md#canonical-3001023330221210-1331320010020032-2102321332033312-1033111130212013-2203112233331203-2030312210323121-0333111323312021-3110100022003012) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](resources--aws_vpc_site--reference--group-005.md#canonical-0010101331103001-3201023330003312-0100122333201200-1220022312111322-0110233101131213-2221133102332213-3020003221130202-2020022032100010) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](resources--aws_vpc_site--reference--group-005.md#canonical-0212301031230321-1003322123220133-3300303030023302-2123000120103000-2023100102301221-3002013213300101-0022332232332011-0120221033013120) |
| `voltstack_cluster.k8s_cluster` | [voltstack_cluster.k8s_cluster](resources--aws_vpc_site--reference--group-005.md#canonical-2123202230030310-0123012200333320-1031321230333201-3011110030313003-1302113123302221-3300132303102333-3013033123130313-0320013002203123) |
| `voltstack_cluster.k8s_cluster.name` | [voltstack_cluster.k8s_cluster.name](resources--aws_vpc_site--reference--group-005.md#canonical-0201002032002311-2231031322322032-1010332201113031-1333120203331000-0202232021100011-0301313313131121-3210113222010113-0120123110023200) |
| `voltstack_cluster.k8s_cluster.namespace` | [voltstack_cluster.k8s_cluster.namespace](resources--aws_vpc_site--reference--group-005.md#canonical-0010300201200211-1212121011011312-0231031333110121-0221232133120323-0321311220111331-2110300022300222-3300112013003020-3001310322220230) |
| `voltstack_cluster.k8s_cluster.tenant` | [voltstack_cluster.k8s_cluster.tenant](resources--aws_vpc_site--reference--group-005.md#canonical-3111001302231333-3213313233011100-2122003201033123-1200211112010330-1033310221302102-2323012231201230-0220000211123130-1310111323223213) |
| `voltstack_cluster.no_dc_cluster_group` | [voltstack_cluster.no_dc_cluster_group](resources--aws_vpc_site--reference--group-005.md#canonical-0012000103221302-3212103230310311-2330220222213012-2121020331202232-3001013212031233-3111021333303221-1321113321132102-2222331302210033) |
| `voltstack_cluster.no_forward_proxy` | [voltstack_cluster.no_forward_proxy](resources--aws_vpc_site--reference--group-005.md#canonical-0213132213020101-1033031001202033-1230320011210002-1101022103113013-0001102122221201-1302321021003203-3112333223133200-2102220233012101) |
| `voltstack_cluster.no_global_network` | [voltstack_cluster.no_global_network](resources--aws_vpc_site--reference--group-005.md#canonical-1203022323003133-1102132303213003-1032022321301002-2030023213113032-1332320012330003-2330333301123132-3311122121021010-1112003130133332) |
| `voltstack_cluster.no_k8s_cluster` | [voltstack_cluster.no_k8s_cluster](resources--aws_vpc_site--reference--group-005.md#canonical-3230112232022112-2322012221221023-0100100210310002-3222233302100133-3321232111021103-1000302302032220-2231321211231032-3332312112323002) |
| `voltstack_cluster.no_network_policy` | [voltstack_cluster.no_network_policy](resources--aws_vpc_site--reference--group-005.md#canonical-1201331233331310-1013313333302311-2213331000000230-1232323003022202-2200032231322221-2011032321031310-1201303202023231-1020111302013202) |
| `voltstack_cluster.no_outside_static_routes` | [voltstack_cluster.no_outside_static_routes](resources--aws_vpc_site--reference--group-005.md#canonical-3301220100020221-0310000332113223-1000211331311223-1233010230321200-0310010223230122-1310333101121002-0011013210220323-3031023323101021) |
| `voltstack_cluster.outside_static_routes` | [voltstack_cluster.outside_static_routes](resources--aws_vpc_site--reference--group-005.md#canonical-3211321313320003-1033012003231211-1123003010011310-3300012121203221-3230132320220010-1310300312022331-2123311220121233-0112120010320202) |
| `voltstack_cluster.outside_static_routes.static_route_list` | [voltstack_cluster.outside_static_routes.static_route_list](resources--aws_vpc_site--reference--group-005.md#canonical-2202311120130012-0200203031312311-3322113232111202-2312133323331232-1122023330233233-1333022113100110-3221213303310133-3031002130131131) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](resources--aws_vpc_site--reference--group-005.md#canonical-2202131110133103-3313033200133300-2201120022232313-2030202000103121-2101110121031112-2111133232003312-0313223102133201-0111022003131331) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.attrs` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.attrs](resources--aws_vpc_site--reference--group-005.md#canonical-1103003000210321-3010233320323133-3301000113001122-3123201112230330-1012100002122023-0002031332022200-0010012121002120-3123122313223112) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels](resources--aws_vpc_site--reference--group-005.md#canonical-3022202212133321-0002022300202111-3112100230112300-0023203330323303-1212202321232012-2032031230231120-1012200321110313-0202121310202302) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_vpc_site--reference--group-005.md#canonical-0010020131221232-0220310222122012-1123012121123022-0102101100302200-2132113003103312-2221220120031023-1013130001111221-2223003330101003) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](resources--aws_vpc_site--reference--group-005.md#canonical-3100032023013033-2210010102230001-3132321221031222-1011233200232000-3311302101310330-0330100223100311-2201032212333231-0110230230020210) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](resources--aws_vpc_site--reference--group-005.md#canonical-3032331120122000-0111203213220331-0312331333103211-3133120303222112-3132000020021013-3323303223032120-0303330033231121-0330133301002103) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](resources--aws_vpc_site--reference--group-005.md#canonical-2202030200323120-0103321200213320-2230033233122112-0112210220132230-2223123212233133-0033221111103021-2012332202032203-2231001020232322) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](resources--aws_vpc_site--reference--group-005.md#canonical-0131310132233110-3113002320220331-2021121011111302-0200211210020102-0001303022201322-2230313323000310-3220022111131000-2010012132313201) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](resources--aws_vpc_site--reference--group-005.md#canonical-1011203121000022-2300300022320122-3031311232100023-3320220200101310-0311233123103111-2130102001012322-1302201131032323-1233113132233121) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](resources--aws_vpc_site--reference--group-005.md#canonical-0331023011302013-2323120212031000-0112223210122230-2032001332312203-0021233300131333-1230102021021020-0101113102330023-0102212113020202) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_vpc_site--reference--group-005.md#canonical-1331232020010233-1000110301303320-2320223002033300-0333113231331103-2300311002031013-0033301202221321-0113102130031103-3210002001301133) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_vpc_site--reference--group-005.md#canonical-3230220222300301-3301010232323303-1232211001310123-2333331033001212-3211031002123211-3220313301011133-1202112000110021-1131022220101212) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](resources--aws_vpc_site--reference--group-005.md#canonical-1312100002230131-3332322102022120-2033130020033230-1332011011113131-2111201000310330-0122322013102023-2210120200002231-3002210100310111) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](resources--aws_vpc_site--reference--group-005.md#canonical-0201001023100021-2300300233122211-3302111020210333-2200120303033011-3100303323203301-0121123113023220-3110303122223003-1000323332120120) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](resources--aws_vpc_site--reference--group-005.md#canonical-3333033013031030-1213222202321321-3013010323032303-0310033302013202-2311303012311210-0003102102022113-0222213012131102-0210332230031020) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](resources--aws_vpc_site--reference--group-005.md#canonical-2201330320130021-0100013123333003-3313102102032213-0002331133300120-0333310222100133-0122132302311312-0211210100312230-0111310233312033) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](resources--aws_vpc_site--reference--group-005.md#canonical-0212333220001010-1320323311000022-2120203303112100-2110033102321320-0213132113311122-1013210020220211-1213112103300122-0010002300200112) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](resources--aws_vpc_site--reference--group-005.md#canonical-1302010223032020-3022330320031113-1133122021300020-0221211303331303-3220032120222002-1131331231321010-2313020213203013-2030032130021231) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](resources--aws_vpc_site--reference--group-005.md#canonical-2120023203201311-0133112132031300-2021103003201011-1013202321001211-0122000213330320-1321011302212120-2110022030203122-1212200302321310) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](resources--aws_vpc_site--reference--group-005.md#canonical-0110131231012110-3031010122321333-2103102031103023-3023033121130201-2210130021013231-0110010202133323-3113000331102110-0213121222122300) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.type](resources--aws_vpc_site--reference--group-005.md#canonical-0002103221323031-2031121012132200-0132211210330030-1313223213131320-2232011121010011-3300132231022130-0212120300211301-0303003210200032) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_vpc_site--reference--group-005.md#canonical-1120023002032200-3013212021321121-3022202233210111-3310122213020011-3212011201302213-3313200220300321-3302323133313102-2022133333331121) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](resources--aws_vpc_site--reference--group-005.md#canonical-3213111220201203-0220320100333312-2120113102120302-1203100330223202-2210321100132202-0331221200010230-1113210113322123-0223032000133233) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](resources--aws_vpc_site--reference--group-005.md#canonical-0330212212013130-1131033112130203-1100121230211311-2123010300023310-1212321311303330-3220301001002322-1013102223301033-1101033201010101) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](resources--aws_vpc_site--reference--group-005.md#canonical-3302330000200201-3313223201222120-3310222222212230-3200323100113021-0221333330211032-0302022001120011-1101113211320000-2213020202320311) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](resources--aws_vpc_site--reference--group-005.md#canonical-3200302132120023-3113033101203011-1032320300321223-3323333202021022-0220012101123002-2300200213201122-3333000332130012-2003200011321033) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](resources--aws_vpc_site--reference--group-005.md#canonical-0210112233310320-3010303123133221-1112321003331031-2323200201003032-3132203230023130-1230201103132020-3323202021123323-2323210232321011) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](resources--aws_vpc_site--reference--group-005.md#canonical-0300132103112012-2121211311021302-0311310023323200-1202133110321031-2313220121312112-0121000021202331-3020232011030221-0222001121032110) |
| `voltstack_cluster.outside_static_routes.static_route_list.simple_static_route` | [voltstack_cluster.outside_static_routes.static_route_list.simple_static_route](resources--aws_vpc_site--reference--group-005.md#canonical-0213022122200132-2023132032130110-2332313100023102-3103330211232122-3022231303102323-0103013213302322-2113002332331120-2222311113312132) |
| `voltstack_cluster.sm_connection_public_ip` | [voltstack_cluster.sm_connection_public_ip](resources--aws_vpc_site--reference--group-005.md#canonical-2233320320000123-2022222122310320-1301101311033300-2230110312012030-2032012003111101-2223313113002032-0202111312212132-0323111130231213) |
| `voltstack_cluster.sm_connection_pvt_ip` | [voltstack_cluster.sm_connection_pvt_ip](resources--aws_vpc_site--reference--group-005.md#canonical-2010321100122113-0021303010321202-2302133200012312-0311200330003020-0320132112333130-0011013133021020-3310021230323012-2202113232321313) |
| `voltstack_cluster.storage_class_list` | [voltstack_cluster.storage_class_list](resources--aws_vpc_site--reference--group-005.md#canonical-1202200302021313-3001210103330001-1311222112020303-0323321233012300-2332031131311111-1131212300210002-2103223120103110-1103023021221013) |
| `voltstack_cluster.storage_class_list.storage_classes` | [voltstack_cluster.storage_class_list.storage_classes](resources--aws_vpc_site--reference--group-005.md#canonical-1113001233300000-3313130222322033-0102222333223002-1331332310231000-0110101100212303-1201003301100030-3322130313230102-0010311010211301) |
| `voltstack_cluster.storage_class_list.storage_classes.default_storage_class` | [voltstack_cluster.storage_class_list.storage_classes.default_storage_class](resources--aws_vpc_site--reference--group-005.md#canonical-1332321120120123-0232331322012121-2102111220032133-3103000123211222-0310203011102003-2230011223102132-0303001232323100-0031311133222030) |
| `voltstack_cluster.storage_class_list.storage_classes.storage_class_name` | [voltstack_cluster.storage_class_list.storage_classes.storage_class_name](resources--aws_vpc_site--reference--group-005.md#canonical-1231010113012131-2120112221012212-0322022000311331-0001233133333333-1300022012132321-2122131300000100-1313003331122003-0230320112013003) |
| `vpc` | [vpc](resources--aws_vpc_site--reference--group-005.md#canonical-3001132322310212-1213202111313031-3011022301100112-2022102222011033-1030003311232103-1100203221121201-0321330223103211-1022312112121100) |
| `vpc.new_vpc` | [vpc.new_vpc](resources--aws_vpc_site--reference--group-005.md#canonical-1000101300233111-3223231002330120-1123102332102220-0003311210200011-2013002230201230-2101020023011232-1132302230012211-2131122032031112) |
| `vpc.new_vpc.autogenerate` | [vpc.new_vpc.autogenerate](resources--aws_vpc_site--reference--group-005.md#canonical-3222022302202232-2122210101110310-0322020222301021-1111013212220322-0020312013122220-1310013221321110-0022131311122000-2130021131131001) |
| `vpc.new_vpc.name_tag` | [vpc.new_vpc.name_tag](resources--aws_vpc_site--reference--group-005.md#canonical-3323300211112332-0220003000010331-0003032311232003-0201122313321222-2213012213012101-3032020303120012-3330300202101112-2332003311130032) |
| `vpc.new_vpc.primary_ipv4` | [vpc.new_vpc.primary_ipv4](resources--aws_vpc_site--reference--group-005.md#canonical-1231200203213232-3200330111111322-1000213312300220-0212200101222003-1022231223322212-0210230001100100-1201011130100202-0001300021000123) |
| `vpc.vpc_id` | [vpc.vpc_id](resources--aws_vpc_site--reference--group-005.md#canonical-0233333131322132-0213020303330231-2122332110003311-0030211233222303-0011332322101002-0021133112211132-0303030321223101-1300100200201020) |
| `waf_signatures` | [waf_signatures](resources--aws_vpc_site--reference--group-005.md#canonical-2011112213031031-1333203331230322-0303120322210030-0022200112003000-1013232101022021-1222320213122133-0303033102333320-1001301201100120) |
| `waf_signatures.automatic` | [waf_signatures.automatic](resources--aws_vpc_site--reference--group-005.md#canonical-0021300122232232-3232320311330200-2021330203000220-0301120101002231-3113032301220302-1301001301210003-0012333220110100-1211022203031201) |
| `waf_signatures.manual` | [waf_signatures.manual](resources--aws_vpc_site--reference--group-005.md#canonical-3022031212312110-0313023132123022-3102200212221131-3102102201012013-2303233212002130-1103022123321133-0102222020000202-3220333123032123) |

<a id="canonical-1230210132232122-2213002030313220-2323212310303331-0012212022123011-2212003032202203-0331211200112023-0131312122231122-1203011212131233"></a>

## Next pages — Property reference / 202302133313 / 20

- [admin_password](resources--aws_vpc_site--reference--group-001.md#canonical-3223101320122233-1021302221332223-3131232122301223-3320302320210222-3320200112020003-0220131333312212-0231210310202123-1201331102030210)
- [aws_cred](resources--aws_vpc_site--reference--group-001.md#canonical-3302022301010023-1332313033022123-3031323333133121-0322223302010331-0111002200002201-0313231300333023-0112200310110233-3023230320012221)
- [block_all_services](resources--aws_vpc_site--reference--group-001.md#canonical-1000312122020333-1132302021220132-2220121211221022-3100303102333101-0020200023202000-3031020300221121-2210322002001020-3133330030320221)
- [blocked_services](resources--aws_vpc_site--reference--group-001.md#canonical-3323032222210001-3302120102203102-3210021101231303-3023332013331112-3200032312323130-3223301102110203-3100201032112113-1310123003122330)
- [coordinates](resources--aws_vpc_site--reference--group-002.md#canonical-3103301232022012-1202210120113233-3231000231313300-1032103030032313-3210112321011201-2112332333203011-1330133221022222-3221223133312113)
- [custom_dns](resources--aws_vpc_site--reference--group-002.md#canonical-0021301331020010-2103113020103011-0313112212033023-3001103200122333-1101111311000032-3120100021023322-0133210203122003-3032032330221333)
- [custom_security_group](resources--aws_vpc_site--reference--group-002.md#canonical-1123132203210311-2301200320231021-3232023100201023-0221020320203100-1321322223023130-1310113030201022-1201331103332123-0222223131322032)
- [default_blocked_services](resources--aws_vpc_site--reference--group-002.md#canonical-1023131220002132-1001131221323100-0123021223123002-2100211102320210-1312301332031313-1133210130030320-0102012300311331-0022310331100331)
- [direct_connect_disabled](resources--aws_vpc_site--reference--group-002.md#canonical-1203233103200121-2313321031323013-0001112100021212-1302133112121133-1031101231223200-2232021210002031-1311100030200230-3200132220200320)
- [direct_connect_enabled](resources--aws_vpc_site--reference--group-002.md#canonical-0130012221000232-1311333022303301-0020012113300213-0322011110120320-0313333031011211-1310113213303030-0201300310003031-3011331001021112)
- [disable_encryption](resources--aws_vpc_site--reference--group-002.md#canonical-1202020332313000-1201233220303020-0132102003011220-1001221330331021-2013110031202200-3000002230300223-3313121333113121-0301333223301031)
- [disable_internet_vip](resources--aws_vpc_site--reference--group-002.md#canonical-1000122022313012-0001101010202330-0013311322231203-1310123033123032-2000220333300012-0323120330133221-0332121010210123-3232302323030130)
- [egress_gateway_default](resources--aws_vpc_site--reference--group-002.md#canonical-1212023023013013-1212222311211002-3121310102313131-0302032221222312-3030030032230033-3321120132111211-1111123202013032-2223010303001021)
- [egress_nat_gw](resources--aws_vpc_site--reference--group-002.md#canonical-2132333123021112-2121130110202202-3232330110112120-1023112123222122-2013332222002322-0212030223100002-1131133120313012-3111311320301310)
- [egress_virtual_private_gateway](resources--aws_vpc_site--reference--group-002.md#canonical-0113110013033313-2210303123332123-0133200222112102-1120033121100103-1132200021011211-3010030013230211-0210122122123323-0310322011033303)
- [enable_encryption](resources--aws_vpc_site--reference--group-002.md#canonical-3122102032303111-0233131232202220-1001102333222013-2102002002231000-3013103312333313-2003211130120332-1010200300312332-3210020213022122)
- [enable_internet_vip](resources--aws_vpc_site--reference--group-002.md#canonical-3302122201112333-3332110213202311-1021212130302332-1120320210103133-1101020320200200-2103031301023231-1200332010031321-0213300302110322)
- [f5_orchestrated_routing](resources--aws_vpc_site--reference--group-002.md#canonical-0233100333233221-1330000203320002-0303020331312332-0103212301123033-1132021203111231-2112210122320010-2313131333302212-2132032101022030)
- [f5xc_security_group](resources--aws_vpc_site--reference--group-002.md#canonical-0122030011102010-3122030030122331-3301230100222300-0332133333222221-0201113331023131-0102122321211202-2132001103112213-0230321112200201)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-0333332022003321-0130333001220102-3300331112302201-3222032002202120-2012011102331101-3213012110100103-2022011031100003-2201121222303323)
- [kubernetes_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-3013301332331333-2211212210321032-3013100011321103-0221022231110303-0303221221101223-1102301300233020-0322322133203301-3113001020112122)
- [log_receiver](resources--aws_vpc_site--reference--group-004.md#canonical-3220233101332201-0221100213122133-2100201232213330-1302330031123033-1031032112102021-3331322021100121-1222111221211231-3331002202001201)
- [logs_streaming_disabled](resources--aws_vpc_site--reference--group-004.md#canonical-2032020211212221-2302310202003310-0312213011010303-2213221031020102-0022211032123131-0011101032031322-0302123012333120-1132100112302201)
- [manual_routing](resources--aws_vpc_site--reference--group-004.md#canonical-2332333023332232-3133011010233103-0013221302323103-2213132121331123-0331101321320212-1032203101301301-2232301020321122-2200320322313121)
- [no_worker_nodes](resources--aws_vpc_site--reference--group-004.md#canonical-0323111213120210-3330110331133331-1123320133233101-0112302203111100-3323311230111302-2322033010012210-0303221110112120-1110302200110131)
- [offline_survivability_mode](resources--aws_vpc_site--reference--group-004.md#canonical-2230331302030233-1231312011331200-0330011122330231-1003002230012012-3322102030022203-0101330012200101-1001211032031010-0121023120133132)
- [os](resources--aws_vpc_site--reference--group-004.md#canonical-3121332323311033-1213202021203320-2022333333100031-2121201303231310-0210311231023010-0313032021023311-3331022231220210-2100330111200133)
- [private_connectivity](resources--aws_vpc_site--reference--group-004.md#canonical-1033321313213300-3312102121203101-2023112202031312-2002230233131030-3320121000233132-3120322133021211-2332131021231000-2321013121233001)
- [sw](resources--aws_vpc_site--reference--group-004.md#canonical-1320213313311232-0331111200330103-3031121112123021-0310133212030023-0022311002322333-1033131021322001-3013223002132330-3033332120222200)
- [timeouts](resources--aws_vpc_site--reference--group-004.md#canonical-3220110111301131-0033121322200310-0033001321230012-3131113012020201-3201213331302322-0033003230131202-1221000010112302-2320021200023132)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [vpc](resources--aws_vpc_site--reference--group-005.md#canonical-2003301201220030-2310111030110321-1023333120123110-3020131220223010-3230003013133033-1023220013310123-3332200230333020-2220002022320010)
- [waf_signatures](resources--aws_vpc_site--reference--group-005.md#canonical-2110022231300120-3130210123131302-1331212201020032-2333303013221203-0003211002203112-2311312323322011-1330000301331202-1312022031112120)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3223101320122233-1021302221332223-3131232122301223-3320302320210222-3320200112020003-0220131333312212-0231210310202123-1201331102030210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3012333213100213-1201103111001011-1331022031001020-0302321331211133-2132312220131032-3302201301021312-0113013123323212-2133313132210111"></a>

## admin_password — admin_password / 010220333330 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- admin_password

<a id="canonical-1121100212113323-2122111130233300-1113011333002112-2033330023221223-0101032210331233-2020132321112001-2131131021202010-0030210212111220"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
```

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

Terraform syntax:

```terraform
admin_password {
  # Configure direct properties listed below.
}
```

<a id="canonical-3003233322310121-2322102012202030-3000302012201123-1123022202113313-0020301233122133-1032201023331312-0001100132232022-1301020213203113"></a>

## Direct properties — admin_password / 010220333330 / 3

- [blindfold_secret_info](resources--aws_vpc_site--reference--group-001.md#canonical-2000101120200303-2231101231031330-3032221130001010-2203112120022113-2222110032120203-2132312203032100-0003302303312031-1201031111301321): complete subsection reference.

- [clear_secret_info](resources--aws_vpc_site--reference--group-001.md#canonical-1330320001101123-3103301223301021-3300013300003133-3002101120023230-2130210221010323-3000022312011023-3200233122012021-1321111122102123): complete subsection reference.

<a id="canonical-0132301031313331-0232030213303330-2100020301120000-2311130112011031-3023211221002102-1101203122010210-0320033101223121-3122110011000133"></a>

## Next pages — admin_password / 010220333330 / 4

- [admin_password.blindfold_secret_info](resources--aws_vpc_site--reference--group-001.md#canonical-2000101120200303-2231101231031330-3032221130001010-2203112120022113-2222110032120203-2132312203032100-0003302303312031-1201031111301321)
- [admin_password.clear_secret_info](resources--aws_vpc_site--reference--group-001.md#canonical-1330320001101123-3103301223301021-3300013300003133-3002101120023230-2130210221010323-3000022312011023-3200233122012021-1321111122102123)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2000101120200303-2231101231031330-3032221130001010-2203112120022113-2222110032120203-2132312203032100-0003302303312031-1201031111301321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123332030213101-0001003231020210-0322033202133230-3200002201203020-0222001302322231-0301232331101303-3310031132130233-1122302133301332"></a>

## admin_password.blindfold_secret_info — blindfold_secret_info / 122013000120 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [admin_password](resources--aws_vpc_site--reference--group-001.md#canonical-3223101320122233-1021302221332223-3131232122301223-3320302320210222-3320200112020003-0220131333312212-0231210310202123-1201331102030210)
- admin_password.blindfold_secret_info

<a id="canonical-3001120211221021-1233302130003102-1123201111021011-0100011310012201-2203112132013203-1203213111110302-1230211102030122-3312111302203112"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
```

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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1211220000012013-0113020133130000-1121322331300100-0112100222230023-1111001031002300-0013103131213101-3320031220030323-0033200133320233"></a>

## Direct properties — blindfold_secret_info / 122013000120 / 3

<a id="canonical-2221201122112212-0001302021220113-0300333211032330-1023311000310013-2023320330321022-3230230033333013-0133201003322120-2320120211013231"></a>

<a id="canonical-1333020313101132-2132033231302332-0222111012322001-2311333100322011-3010012230030220-0132121322311301-0223320012131210-3320010002202132"></a>

## decryption_provider property — blindfold_secret_info / 122013000120 / 4

Type: `"string"`. Optional.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1203120311311111-0221223100121013-0211020203201013-3133323130302101-2132113322332321-3302122323100133-3002211002022200-2301312312333310"></a>

<a id="canonical-2000033332121230-0203330321230110-1301123002333023-1211200013131032-2013333303310320-3013211133013332-0313030120003333-0313100100102131"></a>

## location property — blindfold_secret_info / 122013000120 / 5

Type: `"string"`. Optional, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2220112002202221-2112230212031223-3212312020300203-3011301003131112-1223122230312121-2312222220033232-1102031300122233-0203303202230213"></a>

<a id="canonical-2010022310221001-0121012311303031-0032110200220101-1211303133021020-1213133320231231-2322221201333203-3322233222122102-1311032311330001"></a>

## store_provider property — blindfold_secret_info / 122013000120 / 6

Type: `"string"`. Optional.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0122030131201000-0221223202103000-3110033321313231-3021220123110131-1101000103233131-2233311202323303-0113120311323323-0330200102030022"></a>

## Next pages — blindfold_secret_info / 122013000120 / 7

- [admin_password](resources--aws_vpc_site--reference--group-001.md#canonical-3223101320122233-1021302221332223-3131232122301223-3320302320210222-3320200112020003-0220131333312212-0231210310202123-1201331102030210)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1330320001101123-3103301223301021-3300013300003133-3002101120023230-2130210221010323-3000022312011023-3200233122012021-1321111122102123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023221122102121-0033012322113022-3121123210000112-2333231023213301-1212013003023301-2113012213113121-2030013320012202-3233303322113133"></a>

## admin_password.clear_secret_info — clear_secret_info / 013100131212 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [admin_password](resources--aws_vpc_site--reference--group-001.md#canonical-3223101320122233-1021302221332223-3131232122301223-3320302320210222-3320200112020003-0220131333312212-0231210310202123-1201331102030210)
- admin_password.clear_secret_info

<a id="canonical-3221023332230001-1020003300323020-3202213112102111-2202032103330033-1112013320300101-3121113223213000-2321130003100211-3103023023220102"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
```

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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1121111230201213-1023311003301223-3133102020022330-1110311223332123-0010233001130313-1100322103302301-1030023000003131-2331310101122232"></a>

## Direct properties — clear_secret_info / 013100131212 / 3

<a id="canonical-1120000211110301-1123212222223320-0331010111303112-2312131220122330-2120120231213123-3333313311332012-1030221233123321-1030221212223320"></a>

<a id="canonical-0100120133321332-2333213131131001-2120111320123201-2003302132233033-3201330210102113-1323120133323033-2113302020201321-1000221111320110"></a>

## provider_ref property — clear_secret_info / 013100131212 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2331131323113233-2100331232300103-0312230302131200-1200310320002003-1110300223321211-1303131323210233-2322312223111203-0112101033020322"></a>

<a id="canonical-2213211322321330-3332310323121022-0131320230231200-2012101103013030-3301310303231311-2033022022103201-0200012332211230-1231123001013011"></a>

## URL property — clear_secret_info / 013100131212 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0213320103120001-0220231033212103-0310011130021233-3223320031001100-0232302200003222-0312222212020201-0100201322210230-3011321013011301"></a>

## Next pages — clear_secret_info / 013100131212 / 6

- [admin_password](resources--aws_vpc_site--reference--group-001.md#canonical-3223101320122233-1021302221332223-3131232122301223-3320302320210222-3320200112020003-0220131333312212-0231210310202123-1201331102030210)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3302022301010023-1332313033022123-3031323333133121-0322223302010331-0111002200002201-0313231300333023-0112200310110233-3023230320012221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233100211223132-0123200131033332-0313021121313302-2301201320311302-1332101201003002-0233311032122001-1332133230011222-3133011000000131"></a>

## aws_cred — aws_cred / 021121200202 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- aws_cred

<a id="canonical-1132012231002000-1220330121113220-2303000101331110-3323123213011031-2223212131223321-2133030301301233-3133022000103202-1332330203111123"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
```

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
aws_cred {
  # Configure direct properties listed below.
}
```

<a id="canonical-3003121022133120-1333202230202010-2133233201231021-0022211233203132-1231230212312223-2123332101101033-1310313213333110-3132032010122022"></a>

## Direct properties — aws_cred / 021121200202 / 3

<a id="canonical-3233102033232011-1122323112030322-3131333110133330-1230033212322232-1330210231322313-1013002220230102-1333323221203232-2223203113310320"></a>

<a id="canonical-1112122000310323-1301131203132100-1202122000310332-1121022022023111-1322013002120213-1220130330123323-0030103320312102-0103211212212322"></a>

## name property — aws_cred / 021121200202 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3110303213133131-2310003200220233-1303321300301203-0102110101013232-0113300201012100-2231010320311103-0300310313012121-0333313032100311"></a>

<a id="canonical-0112013133230213-2132233112032333-1221022322201020-3023112103130313-0332310003000103-2223010221110331-2232112203002221-0111003000011133"></a>

## namespace property — aws_cred / 021121200202 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3212301100203213-3300311132221233-1103133022011303-0021303201212200-3202301321123330-3210301233322110-1300203001202213-1031103132320103"></a>

<a id="canonical-0300103012031213-2301003220223003-3310200122102313-1212230110302023-2220021120220103-1122323223131302-2122222013112230-1212222301221030"></a>

## tenant property — aws_cred / 021121200202 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2321133300312130-1132030120132032-3030022200023221-0313310221203113-3111300031123312-1031312223323232-0100122110021131-1312220132010322"></a>

## Next pages — aws_cred / 021121200202 / 7

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1000312122020333-1132302021220132-2220121211221022-3100303102333101-0020200023202000-3031020300221121-2210322002001020-3133330030320221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010132123212221-2100202012120321-1330332212201100-3001220111120100-0201111323003223-3123211230202321-1100331100003103-0001232031221100"></a>

## block_all_services — block_all_services / 102032330131 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- block_all_services

<a id="canonical-1110123031032010-1113032032313120-1120000201103311-1100232123310211-0010220102213133-2010201012003033-1122010100320133-3003030022133303"></a>

Type: `["object", {}]`. Optional.

\[OneOf: block\_all\_services, blocked\_services, default\_blocked\_services; Default:
default\_blocked\_services\] Enable this option

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

OneOf alternatives in this subsection:

- [block_all_services](resources--aws_vpc_site--reference--group-001.md#canonical-1110123031032010-1113032032313120-1120000201103311-1100232123310211-0010220102213133-2010201012003033-1122010100320133-3003030022133303)
- [blocked_services](resources--aws_vpc_site--reference--group-001.md#canonical-1101130131202131-3221200333012230-2011223031200331-1121301112001313-1033012302112113-0212212000031022-0021201022330101-3333033310132202)
- [default_blocked_services](resources--aws_vpc_site--reference--group-002.md#canonical-0100133120312222-0222002310232213-0233220031220330-1112322001310302-3112301002102313-2311111022202021-3000323303112021-0022130113232302)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
block_all_services = {}
```

<a id="canonical-0302323223121201-1321121301111322-0331031010231230-2133231223010200-2133020000111001-1103223022322113-1002310020313011-2220131231303132"></a>

## Direct properties — block_all_services / 102032330131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0231133233023131-0031333101313133-0111203320222001-0103011121201013-2311001110201323-1211311202233212-3303102113022311-1102201122221221"></a>

## Next pages — block_all_services / 102032330131 / 4

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3323032222210001-3302120102203102-3210021101231303-3023332013331112-3200032312323130-3223301102110203-3100201032112113-1310123003122330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232120200301322-0021101013033311-3100112031121222-0030012323120211-0130323200231231-2332123030023003-1010103231132310-0021320313011203"></a>

## blocked_services — blocked_services / 210233321223 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- blocked_services

<a id="canonical-1101130131202131-3221200333012230-2011223031200331-1121301112001313-1033012302112113-0212212000031022-0021201022330101-3333033310132202"></a>

Type: `"object"`. single nested block, Optional.

Disable node local services on this site.

Upstream description:

Disable node local services on this site. Note: The chosen services will GET disabled on all nodes
in the site.

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
blocked_services {
  # Configure direct properties listed below.
}
```

<a id="canonical-3233012101311103-1112021132113023-3230203012001021-1201111320311031-3130333111112302-3000112300302200-2322031330110333-0222321113311130"></a>

## Direct properties — blocked_services / 210233321223 / 3

- [blocked_service](resources--aws_vpc_site--reference--group-001.md#canonical-0102322030032221-1312020110103322-0312312320331312-2323203222312233-3130330322220220-3322320221222010-0311102231203132-2213033023012032): complete subsection reference.

<a id="canonical-1010333033312200-2332331212323120-0333303211130213-3131103300031323-2100321113210100-1102122003300022-0222102300012222-3021233001213022"></a>

## Next pages — blocked_services / 210233321223 / 4

- [blocked_services.blocked_service](resources--aws_vpc_site--reference--group-001.md#canonical-0102322030032221-1312020110103322-0312312320331312-2323203222312233-3130330322220220-3322320221222010-0311102231203132-2213033023012032)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0102322030032221-1312020110103322-0312312320331312-2323203222312233-3130330322220220-3322320221222010-0311102231203132-2213033023012032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133313231032021-1001000303003110-2101102100212220-2332003103100320-0321221202303133-3200123303133213-2230211022030222-3113131132000311"></a>

## blocked_services.blocked_service — blocked_service / 003023021300 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [blocked_services](resources--aws_vpc_site--reference--group-001.md#canonical-3323032222210001-3302120102203102-3210021101231303-3023332013331112-3200032312323130-3223301102110203-3100201032112113-1310123003122330)
- blocked_services.blocked_service

<a id="canonical-0023123302030330-0322311221201333-0000103202310321-2202231033222030-2030001131331300-1321021221221011-3330302000010000-3032213320032013"></a>

Type: `"object"`. list nested block, Optional.

Disable Node Local Services. Blocking or denial configuration

Upstream description:

Blocking or denial configuration

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("dns",
    "ssh"),
  validators.ConflictingListObjectAttributes("dns",
    "web_user_interface"),
  validators.ConflictingListObjectAttributes("ssh",
    "web_user_interface")}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
blocked_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-3330320201221102-3321303313012332-3320011122103120-0323123221132333-1211020130301003-2302031121013321-1132313232220113-2223010110302031"></a>

## Direct properties — blocked_service / 003023021300 / 3

- [DNS](resources--aws_vpc_site--reference--group-001.md#canonical-3302211022120303-3310020210201301-1310311112010112-2132303333212001-0330111323312131-0002120333010010-3130331231110332-0032020013230131): complete subsection reference.

<a id="canonical-1323211210131301-0201000123323213-0002130022223303-0220220131300311-2313001130112200-0323032231013102-0313020013232000-1013201032022223"></a>

<a id="canonical-0221231000032210-1032321123221201-3023020132121010-1212021031113002-3103203333313101-3110233331003303-0132103002012133-2321012213311323"></a>

## network_type property — blocked_service / 003023021300 / 4

Type: `"string"`. Optional.

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

Upstream description:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created automatically and present on all sites
Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE is a private network inside site. It
is a secure network and is not connected to public network. Virtual-network of this type is local to
every site. Two virtual networks of this type on different sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created during provisioning of site User defined per-site
virtual network. Scope of this virtual network is limited to the site. This is not yet supported
Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC directly connects to the public internet.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on RE sites only It is an internally created by the system. They must not be created by
user Virtual Networks with global scope across different sites in F5XC domain. An example global
virtual-network called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VIRTUAL_NETWORK_SITE_LOCAL",
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
    "VIRTUAL_NETWORK_MANAGEMENT"),
}
```

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

- [ssh](resources--aws_vpc_site--reference--group-001.md#canonical-2012223122113100-3010320001230313-1131221113323201-1133122300002322-2311022123133110-1310310313223000-3123222222003223-0221131220300023): complete subsection reference.

- [web_user_interface](resources--aws_vpc_site--reference--group-001.md#canonical-2330332230332100-3130113332112023-3212301120231310-3232122113032212-3112320213102031-1021333200300212-3103201330330331-0222211030023223): complete subsection reference.

<a id="canonical-0211220012231322-0301001132313033-0203210030010313-1302200310321112-1311211220311320-3103010130200331-0011120031022233-2112133200222312"></a>

## Next pages — blocked_service / 003023021300 / 5

- [blocked_services.blocked_service.dns](resources--aws_vpc_site--reference--group-001.md#canonical-3302211022120303-3310020210201301-1310311112010112-2132303333212001-0330111323312131-0002120333010010-3130331231110332-0032020013230131)
- [blocked_services.blocked_service.ssh](resources--aws_vpc_site--reference--group-001.md#canonical-2012223122113100-3010320001230313-1131221113323201-1133122300002322-2311022123133110-1310310313223000-3123222222003223-0221131220300023)
- [blocked_services.blocked_service.web_user_interface](resources--aws_vpc_site--reference--group-001.md#canonical-2330332230332100-3130113332112023-3212301120231310-3232122113032212-3112320213102031-1021333200300212-3103201330330331-0222211030023223)
- [blocked_services](resources--aws_vpc_site--reference--group-001.md#canonical-3323032222210001-3302120102203102-3210021101231303-3023332013331112-3200032312323130-3223301102110203-3100201032112113-1310123003122330)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3302211022120303-3310020210201301-1310311112010112-2132303333212001-0330111323312131-0002120333010010-3130331231110332-0032020013230131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022020333323012-1301203320311201-1131321301322002-1031330221113201-0301223132021212-2122222132100021-0130211033113023-1301103201210002"></a>

## blocked_services.blocked_service.DNS — DNS / 103000213122 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [blocked_services](resources--aws_vpc_site--reference--group-001.md#canonical-3323032222210001-3302120102203102-3210021101231303-3023332013331112-3200032312323130-3223301102110203-3100201032112113-1310123003122330)
- [blocked_services.blocked_service](resources--aws_vpc_site--reference--group-001.md#canonical-0102322030032221-1312020110103322-0312312320331312-2323203222312233-3130330322220220-3322320221222010-0311102231203132-2213033023012032)
- blocked_services.blocked_service.DNS

<a id="canonical-3122022230100220-0223310313020131-1012321033201200-0301233132331021-1131323023120033-1223323132222300-3332312232212031-3200130011221202"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
dns = {}
```

<a id="canonical-1210221321033023-1300223313210001-3030011123121122-1223021130111022-0113000000211131-2302032120221332-1233100100220311-1031101313120322"></a>

## Direct properties — DNS / 103000213122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1323232332320022-1102212013221020-0210123323211111-1222112131303100-1123122301231221-3112203332200311-2023233112231201-2021020132132020"></a>

## Next pages — DNS / 103000213122 / 4

- [blocked_services.blocked_service](resources--aws_vpc_site--reference--group-001.md#canonical-0102322030032221-1312020110103322-0312312320331312-2323203222312233-3130330322220220-3322320221222010-0311102231203132-2213033023012032)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2012223122113100-3010320001230313-1131221113323201-1133122300002322-2311022123133110-1310310313223000-3123222222003223-0221131220300023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0010022331022002-1112311212313222-0003021122101303-3013213323100001-1133100122312213-1311001302111311-2003013123022212-2330111320202000"></a>

## blocked_services.blocked_service.ssh — ssh / 222002301231 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [blocked_services](resources--aws_vpc_site--reference--group-001.md#canonical-3323032222210001-3302120102203102-3210021101231303-3023332013331112-3200032312323130-3223301102110203-3100201032112113-1310123003122330)
- [blocked_services.blocked_service](resources--aws_vpc_site--reference--group-001.md#canonical-0102322030032221-1312020110103322-0312312320331312-2323203222312233-3130330322220220-3322320221222010-0311102231203132-2213033023012032)
- blocked_services.blocked_service.ssh

<a id="canonical-2333331022213023-3023111001033211-3132330212121203-0031210012120101-0201032220210222-1333211021022122-0110322103023030-2323202111320312"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
ssh = {}
```

<a id="canonical-0311202211000113-0122313112112000-1021313113303333-3331220303330032-3020212113230110-2032232201010120-1100223101222010-0113213103330230"></a>

## Direct properties — ssh / 222002301231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3331100012023011-2211301300303230-0203231300222010-0222210302300323-1331332111220311-0301001012022311-1021012312120203-3200230001300311"></a>

## Next pages — ssh / 222002301231 / 4

- [blocked_services.blocked_service](resources--aws_vpc_site--reference--group-001.md#canonical-0102322030032221-1312020110103322-0312312320331312-2323203222312233-3130330322220220-3322320221222010-0311102231203132-2213033023012032)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2330332230332100-3130113332112023-3212301120231310-3232122113032212-3112320213102031-1021333200300212-3103201330330331-0222211030023223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
