---
page_title: "xcsh_aws_tgw_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_aws_tgw_site reference."
---

# xcsh_aws_tgw_site reference

<a id="canonical-0103221113020220-1023032300110200-1122320003021100-3301223110103000-1232122010031200-2123003310013320-0101220201312303-1332330322113022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033303300001302-2312031200000231-1223331133013232-3313111002301320-1011031130132331-3023222212200320-3210333330131201-2021332022231223"></a>

## aws_parameters.new_tgw.system_generated — system_generated / 333230003030 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- [aws_parameters.new_tgw](resources--aws_tgw_site--reference--group-001.md#canonical-0201010233301021-1032210023013130-2031032132102301-2312330221121331-1233321022203030-1131222010323002-1133333302232200-1323232013331102)
- aws_parameters.new_tgw.system_generated

<a id="canonical-0111121102212333-0213012030310031-3030310222332333-3332300200303000-1113100212010122-1202123331013121-1321213111032012-3331003311210232"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for system generated.

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
system_generated = {}
```

<a id="canonical-3101021302323030-1202223023012200-2212300122013123-1010230122100023-2230212311300310-1012111033023212-3031100001210123-3032020221132021"></a>

## Direct properties — system_generated / 333230003030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1232300101121031-1030331101203111-2333031011230000-3300120300030322-0222231321310100-1212001130212110-3203303100132222-1102110211310121"></a>

## Next pages — system_generated / 333230003030 / 4

- [aws_parameters.new_tgw](resources--aws_tgw_site--reference--group-001.md#canonical-0201010233301021-1032210023013130-2031032132102301-2312330221121331-1233321022203030-1131222010323002-1133333302232200-1323232013331102)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-1012321201321212-3333210011321220-1300013121113023-2010222022103120-1331133131000231-2010332022011311-0221120300323130-1130322133020332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211323002213233-2021122023223313-1333013331102100-0323333032011132-3321013231320010-3313122133232232-2230011013311123-0203013103031201"></a>

## aws_parameters.new_tgw.user_assigned — user_assigned / 130321323220 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- [aws_parameters.new_tgw](resources--aws_tgw_site--reference--group-001.md#canonical-0201010233301021-1032210023013130-2031032132102301-2312330221121331-1233321022203030-1131222010323002-1133333302232200-1323232013331102)
- aws_parameters.new_tgw.user_assigned

<a id="canonical-0330001111021200-3103321232303210-3331011223323012-1300120130320102-1213002312130202-0013030121002232-2120012033130313-0131122121103022"></a>

Type: `"object"`. single nested block, Optional.

Information needed when ASNs are assigned by the user.

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
user_assigned {
  # Configure direct properties listed below.
}
```

<a id="canonical-3302023211031203-1202103230002200-3031112303222122-3313101331212203-0333020002020002-0231101320033331-3313102301313302-3100222202031330"></a>

## Direct properties — user_assigned / 130321323220 / 3

<a id="canonical-1011331022322102-3213010103332221-2130231022232013-3022211320330333-1133011200301013-3232230210123202-2111213323033120-0211033131113300"></a>

<a id="canonical-2312030310031330-0303033330132003-3032122122313022-0022100112013301-0013331212031230-1213131302102031-0000103212332102-2122330112100110"></a>

## tgw_asn property — user_assigned / 130321323220 / 4

Type: `"number"`. Optional.

TGW ASN. Allowed range for 16-bit private ASNs include 64512 to 65534.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(64513, 65534),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65534,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 64513
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "64512",
    "ves.io.schema.rules.uint32.lte": "65534"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "64512",
    "ves.io.schema.rules.uint32.lte": "65534"
  }
}
```

<a id="canonical-1232102011212333-3331120221333323-0303033233121113-2011012001312002-3303000130223213-3021030313232013-3320201231033030-2303021202332112"></a>

<a id="canonical-3203203200310310-2230312233301132-2020203033210113-2213020021333210-2101132320101213-2123120320103013-1223203301123120-2001112221101101"></a>

## volterra_site_asn property — user_assigned / 130321323220 / 5

Type: `"number"`. Optional.

Enter F5XC Site ASN. F5XC Site ASN.

Upstream description:

F5XC Site ASN.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0230021003330222-2012300003231313-1321031131222112-3211131113302100-3002002023100022-3212112130103011-1103011102022212-3002332111110101"></a>

## Next pages — user_assigned / 130321323220 / 6

- [aws_parameters.new_tgw](resources--aws_tgw_site--reference--group-001.md#canonical-0201010233301021-1032210023013130-2031032132102301-2312330221121331-1233321022203030-1131222010323002-1133333302232200-1323232013331102)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-0220020323123321-0313220113210230-0113310122202133-2213030113313113-2003203010103222-0032202230121212-2000112223330031-2120322213130322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1230120221121102-0311200001031212-3123101033013220-3013023131111203-2022303231230203-1330332323202211-2021110103201112-0002013322210111"></a>

## aws_parameters.new_vpc — new_vpc / 120313232032 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- aws_parameters.new_vpc

<a id="canonical-1310323200030310-1133330202021202-1100131023133332-3020003203323020-3211303000210113-1221203012121332-1212110132010031-0133330001030123"></a>

Type: `"object"`. single nested block, Optional.

AWS VPC Parameters. Parameters to create new AWS VPC.

Upstream description:

Parameters to create new AWS VPC.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("primary_ipv4"),
  validators.ConflictingObjectAttributes("autogenerate",
    "name_tag")}
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
  "x-ves-oneof-field-name_choice": "[\"autogenerate\",\"name_tag\"]"
}
```

Terraform syntax:

```terraform
new_vpc {
  # Configure direct properties listed below.
}
```

<a id="canonical-3311133220201132-1011110110230122-1002120232332331-2303310123123201-2321301013323132-3211000232111230-0022113003120023-0212123123222303"></a>

## Direct properties — new_vpc / 120313232032 / 3

- [autogenerate](resources--aws_tgw_site--reference--group-002.md#canonical-1202212102332103-2130320110010101-3030031132133003-0203112230312220-1031032300311121-2022201222313202-3111232110333201-1332220130120323): complete subsection reference.

<a id="canonical-1212320211313121-2000132201323233-1333203221233220-1003203212223101-3223110010000032-3110111030322022-2331113201002103-3023102001333020"></a>

<a id="canonical-2220030030012112-2013012112001201-1012231030011123-2332012001311320-1303132112200220-1231030123122003-3010120223030030-0123131003203323"></a>

## name_tag property — new_vpc / 120313232032 / 4

Type: `"string"`. Optional.

Exclusive with \[autogenerate\] Specify the VPC Name.

Upstream description:

Exclusive with \[autogenerate\] Specify the VPC Name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-0202210203032333-1011320030010133-3231013110333120-2010202203131303-2201012311021020-0101021301220021-0312020002100113-2022221030022230"></a>

<a id="canonical-0101030302013303-1021211313020202-0313112222021030-1331000212132002-3213301320220021-2012103311232023-3332133323102122-0112233001031210"></a>

## primary_ipv4 property — new_vpc / 120313232032 / 5

Type: `"string"`. Optional.

IPv4 CIDR block for this VPC. It has to be private address space. The Primary IPv4 block cannot be
modified. All subnets prefixes in this VPC must be part of this CIDR block.

Upstream description:

IPv4 CIDR block for this VPC. It has to be private address space. The Primary IPv4 block cannot be
modified. All subnets prefixes in this VPC must be part of this CIDR block.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28",
    "ves.io.schema.rules.string.min_ip_prefix_length": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28",
    "ves.io.schema.rules.string.min_ip_prefix_length": "16"
  }
}
```

<a id="canonical-0301113320122222-0311102201321121-1021220013102302-0332021302200021-0002030111300302-1101202102101303-3200013121023203-2230221231101212"></a>

## Next pages — new_vpc / 120313232032 / 6

- [aws_parameters.new_vpc.autogenerate](resources--aws_tgw_site--reference--group-002.md#canonical-1202212102332103-2130320110010101-3030031132133003-0203112230312220-1031032300311121-2022201222313202-3111232110333201-1332220130120323)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-1202212102332103-2130320110010101-3030031132133003-0203112230312220-1031032300311121-2022201222313202-3111232110333201-1332220130120323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202222103233103-0113003231313001-2313133012323123-0011330030323110-3013122312123302-1302300000033200-2223033200310223-0011002321130120"></a>

## aws_parameters.new_vpc.autogenerate — autogenerate / 310100301222 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- [aws_parameters.new_vpc](resources--aws_tgw_site--reference--group-002.md#canonical-0220020323123321-0313220113210230-0113310122202133-2213030113313113-2003203010103222-0032202230121212-2000112223330031-2120322213130322)
- aws_parameters.new_vpc.autogenerate

<a id="canonical-1030323220322200-1222101031101001-1101210132322011-2221103231200131-3130311101322132-2121213012112311-3231200030300203-1313220103131013"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for autogenerate.

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
autogenerate = {}
```

<a id="canonical-3320033130233131-2221332103311302-0132321012331121-0213200023132001-3313031310023030-3320011133023320-2320223002000003-2111031232300130"></a>

## Direct properties — autogenerate / 310100301222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1202231113231132-3330223313313212-3011122030130320-2231230211010330-2201311012131000-2030203123121302-0200311222022333-3320230013001000"></a>

## Next pages — autogenerate / 310100301222 / 4

- [aws_parameters.new_vpc](resources--aws_tgw_site--reference--group-002.md#canonical-0220020323123321-0313220113210230-0113310122202133-2213030113313113-2003203010103222-0032202230121212-2000112223330031-2120322213130322)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-2221202120101010-2021301212311032-1333222232212320-0013032102113003-0302100203333220-1333213231210112-0133010203022110-3321330223121101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303323201033221-2102211021233122-3100221302230301-1121312321311012-3000102133033002-1110000103003231-2022023301101123-1313131130322030"></a>

## aws_parameters.no_worker_nodes — no_worker_nodes / 231111002023 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- aws_parameters.no_worker_nodes

<a id="canonical-2330102331213101-2132233231211011-1133011110121003-1212332313202320-0100102233212210-2131112012221122-2232210213100033-2221323320303322"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no worker nodes.

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
no_worker_nodes = {}
```

<a id="canonical-3103000103011321-3212231202011000-1133101213131132-0222330001122122-0120102113103211-2300321212203002-3032213123121002-2302010232222200"></a>

## Direct properties — no_worker_nodes / 231111002023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1131200212321201-2231033121131121-0021232211112300-2022003320013313-0113213132222300-1000232321312333-1103133021121110-0021312010310103"></a>

## Next pages — no_worker_nodes / 231111002023 / 4

- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-3121213320330121-3012221332212132-3211100230303323-0221132120130330-0212223211102331-2031212313332200-1211230111210012-3023322021210301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201023300210300-2101023011131012-0011301330210132-3013230100033121-3312303222222321-1200010233203122-0332011010001221-2023321221022210"></a>

## aws_parameters.reserved_tgw_cidr — reserved_tgw_cidr / 130013101300 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- aws_parameters.reserved_tgw_cidr

<a id="canonical-0201133133111122-2313023300020033-2110120023221213-2212010103221203-3101130033022023-3323222110101132-0211321201221120-0223101021232022"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for reserved tgw cidr.

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
reserved_tgw_cidr = {}
```

<a id="canonical-2110311222030221-1311231113130312-0122123201220220-2312021301003323-0020110121032213-1202302201200001-1101310013200131-2311332133320320"></a>

## Direct properties — reserved_tgw_cidr / 130013101300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0031132310330323-0033230203233011-3213210211202322-1131333333201213-0303212113311031-1032211330320020-0103230200230110-3012330023030220"></a>

## Next pages — reserved_tgw_cidr / 130013101300 / 4

- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-3131221233111313-0201223031103232-3231330222212320-3200313300213211-3332131122121033-2101311133231111-0111222033131212-1030022202322120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220110012320321-3101303103031133-0003122021330302-1031031032333101-1113322233110230-2103023110110211-1311102011202130-2313012323111002"></a>

## aws_parameters.tgw_cidr — tgw_cidr / 203211312210 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- aws_parameters.tgw_cidr

<a id="canonical-1000130131132022-3233322210220011-1103322332231323-1303320000223000-1022130001331122-0001121303233233-1212203130132022-0130300223310110"></a>

Type: `"object"`. single nested block, Optional.

Parameters for creating a new cloud subnet.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("ipv4")}
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
tgw_cidr {
  # Configure direct properties listed below.
}
```

<a id="canonical-3231133110020312-0010130330200320-0033212131023220-3031003110302212-1102321121311021-0021333012230332-0232131223212013-3230321212113203"></a>

## Direct properties — tgw_cidr / 203211312210 / 3

<a id="canonical-3133113100200201-0002021100330001-3110231111202001-0231002303223302-1223112123332130-2200212313123322-1302320221101220-3020012303322023"></a>

<a id="canonical-2311133230231121-1220130203330110-3023213232203232-2202120232311130-3313302320322301-2120333310330223-2133333122301030-3011212200100223"></a>

## IPv4 property — tgw_cidr / 203211312210 / 4

Type: `"string"`. Optional.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-3333120013001200-1102213100001120-1012221220121200-3001221002011203-2300200203020200-2110031130302303-2122010030130110-0002133221102121"></a>

## Next pages — tgw_cidr / 203211312210 / 5

- [aws_parameters](resources--aws_tgw_site--reference--group-001.md#canonical-0211211321130333-3310103111003032-2331313101231121-1000033231012332-1331312012302111-2003331003101111-3231100013121220-1221232021331020)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-3303203233101111-1231211011103232-0133220301212023-0131322023033222-3032123223330012-1020121030201121-2102031230301211-3211000022222033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321211130123131-0022022100112312-1211103132210032-0313221330331000-3311131120013221-1131021200023202-1023030010313013-0102320110231310"></a>

## block_all_services — block_all_services / 323330223001 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- block_all_services

<a id="canonical-2102122013330111-1010023202001123-2202113001233212-1001103032131210-1101003121133022-1001323233130310-3210010332121103-0232332110003011"></a>

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

- [block_all_services](resources--aws_tgw_site--reference--group-002.md#canonical-2102122013330111-1010023202001123-2202113001233212-1001103032131210-1101003121133022-1001323233130310-3210010332121103-0232332110003011)
- [blocked_services](resources--aws_tgw_site--reference--group-002.md#canonical-3302233010132223-0120213300333133-0303022322031213-0031121322002022-3032132110113103-1023021232102221-0132300011023001-0112123320221203)
- [default_blocked_services](resources--aws_tgw_site--reference--group-002.md#canonical-2000132300021100-3032231133121301-3320213021031121-3000122030011111-0212331022121323-3002011020221112-0222023233121020-3032212002101311)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
block_all_services = {}
```

<a id="canonical-2321120023311232-2122001013302033-2201003322320122-0022003302231011-0232301333233323-2100001133103011-1110130112321222-1221131310220223"></a>

## Direct properties — block_all_services / 323330223001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2001203331322232-2230021133000320-2322331102120222-3110200031202202-2112113032020030-0033103222133311-1310322032132210-2203312302120122"></a>

## Next pages — block_all_services / 323330223001 / 4

- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-3220312023213323-0222323213213020-3313231022111310-1322001320032202-1021333312223203-0100222312131310-2203300011012010-3103320210301101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102230201223002-2133000100311032-1202302311220121-1313223333211332-1133211102123100-3031013131223203-1020102322020110-1013000001310213"></a>

## blocked_services — blocked_services / 133123211221 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- blocked_services

<a id="canonical-3302233010132223-0120213300333133-0303022322031213-0031121322002022-3032132110113103-1023021232102221-0132300011023001-0112123320221203"></a>

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

<a id="canonical-0100000112222223-0330333113002300-3010021223103100-2331233313231022-1033302100303310-2323311303310303-2010200202322333-0300003030013231"></a>

## Direct properties — blocked_services / 133123211221 / 3

- [blocked_service](resources--aws_tgw_site--reference--group-002.md#canonical-3330321021120220-0211220102310003-2312323202103000-1212113200003211-1110303112200210-1003321013202112-3202030102311132-2210312332310310): complete subsection reference.

<a id="canonical-3301202130001012-3011222322022031-0113013200203113-3111222300323131-1003223113221300-3013020301330022-1231031202103023-3300323310030012"></a>

## Next pages — blocked_services / 133123211221 / 4

- [blocked_services.blocked_service](resources--aws_tgw_site--reference--group-002.md#canonical-3330321021120220-0211220102310003-2312323202103000-1212113200003211-1110303112200210-1003321013202112-3202030102311132-2210312332310310)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-3330321021120220-0211220102310003-2312323202103000-1212113200003211-1110303112200210-1003321013202112-3202030102311132-2210312332310310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303231220132022-2020313312313232-1120020210002323-0301222133122122-2003031111111312-0102322113200300-2030123323000312-1030020300013320"></a>

## blocked_services.blocked_service — blocked_service / 331113201102 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [blocked_services](resources--aws_tgw_site--reference--group-002.md#canonical-3220312023213323-0222323213213020-3313231022111310-1322001320032202-1021333312223203-0100222312131310-2203300011012010-3103320210301101)
- blocked_services.blocked_service

<a id="canonical-3232323021210320-2030120321131233-2200003200231311-1122202131000103-2232212322013320-0200123033321333-1000122113222002-3310110030110322"></a>

Type: `"object"`. list nested block, Optional.

Disable Node Local Services. Blocking or denial configuration

Upstream description:

Blocking or denial configuration

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3303100302131330-1023103221203212-0123111132331321-0233312310123230-3120002300210230-3010013222230310-1333201102332311-3310111130311011"></a>

## Direct properties — blocked_service / 331113201102 / 3

- [DNS](resources--aws_tgw_site--reference--group-002.md#canonical-3020320313222311-1001121100203321-0103302101312103-1110011032310223-1311003322332201-1311110320103121-0221331030321101-3303203312111111): complete subsection reference.

<a id="canonical-1311331023131303-3013131202132123-1001203333303330-0133312133021122-0230003012302303-1221003112120102-3200103113322303-1100221323121223"></a>

<a id="canonical-3301230022020302-2023332203123323-2033033300222223-2023101133301103-3231221021010133-2201211121300332-1013321012210302-3211021032121003"></a>

## network_type property — blocked_service / 331113201102 / 4

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
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["VIRTUAL_NETWORK_GLOBAL","VIRTUAL_NETWORK_IP_AUTO","VIRTUAL_NETWORK_IP_FABRIC","VIRTUAL_NETWORK_MANAGEMENT","VIRTUAL_NETWORK_PER_SITE","VIRTUAL_NETWORK_PUBLIC","VIRTUAL_NETWORK_SEGMENT","VIRTUAL_NETWORK_SITE_LOCAL","VIRTUAL_NETWORK_SITE_LOCAL_INSIDE","VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE","VIRTUAL_NETWORK_SITE_SERVICE","VIRTUAL_NETWORK_SRV6_NETWORK","VIRTUAL_NETWORK_VER_INTERNAL","VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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

- [SSH](resources--aws_tgw_site--reference--group-002.md#canonical-1213232301302012-0033110021321203-0001210010031131-0213213120033203-2020321313002311-0230123221101221-0133132203122332-1121203001001130): complete subsection reference.

- [web_user_interface](resources--aws_tgw_site--reference--group-002.md#canonical-0331001001211212-0112333223132230-1202112232020332-1000310030213223-2121121313201202-1020113330330100-1201303212013222-3322211033220310): complete subsection reference.

<a id="canonical-3210003220130313-1333000110020303-3332210321301210-1101001032011333-3102000033300301-2310131023231311-3031032212320302-3320110010331100"></a>

## Next pages — blocked_service / 331113201102 / 5

- [blocked_services.blocked_service.dns](resources--aws_tgw_site--reference--group-002.md#canonical-3020320313222311-1001121100203321-0103302101312103-1110011032310223-1311003322332201-1311110320103121-0221331030321101-3303203312111111)
- [blocked_services.blocked_service.ssh](resources--aws_tgw_site--reference--group-002.md#canonical-1213232301302012-0033110021321203-0001210010031131-0213213120033203-2020321313002311-0230123221101221-0133132203122332-1121203001001130)
- [blocked_services.blocked_service.web_user_interface](resources--aws_tgw_site--reference--group-002.md#canonical-0331001001211212-0112333223132230-1202112232020332-1000310030213223-2121121313201202-1020113330330100-1201303212013222-3322211033220310)
- [blocked_services](resources--aws_tgw_site--reference--group-002.md#canonical-3220312023213323-0222323213213020-3313231022111310-1322001320032202-1021333312223203-0100222312131310-2203300011012010-3103320210301101)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-3020320313222311-1001121100203321-0103302101312103-1110011032310223-1311003322332201-1311110320103121-0221331030321101-3303203312111111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221022211300302-1310231321212111-3302200211221221-1322211321322222-3231103130311323-0211012133113012-3023102112112203-1010220212113120"></a>

## blocked_services.blocked_service.DNS — DNS / 330302023321 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [blocked_services](resources--aws_tgw_site--reference--group-002.md#canonical-3220312023213323-0222323213213020-3313231022111310-1322001320032202-1021333312223203-0100222312131310-2203300011012010-3103320210301101)
- [blocked_services.blocked_service](resources--aws_tgw_site--reference--group-002.md#canonical-3330321021120220-0211220102310003-2312323202103000-1212113200003211-1110303112200210-1003321013202112-3202030102311132-2210312332310310)
- blocked_services.blocked_service.DNS

<a id="canonical-1202012320210102-2010111123230120-2210312110322221-1131322313023120-0220102223023011-2112231300121031-0320120011330122-0110230013010111"></a>

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

<a id="canonical-1203022231103030-2002203321211310-2021131332303120-0200201302000313-0131022230131102-1312232010212031-3023100133013122-3102312010111222"></a>

## Direct properties — DNS / 330302023321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0122121133212233-1121332332112310-1211111010012223-3130131120111102-0133323002303020-1202303233231221-3211132220312131-1213012300001321"></a>

## Next pages — DNS / 330302023321 / 4

- [blocked_services.blocked_service](resources--aws_tgw_site--reference--group-002.md#canonical-3330321021120220-0211220102310003-2312323202103000-1212113200003211-1110303112200210-1003321013202112-3202030102311132-2210312332310310)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-1213232301302012-0033110021321203-0001210010031131-0213213120033203-2020321313002311-0230123221101221-0133132203122332-1121203001001130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213120120122033-0113323323302022-2213021000011011-3110020132311230-2102023321302301-1001230131030233-3311100111010322-0221003321013101"></a>

## blocked_services.blocked_service.SSH — SSH / 010302002120 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [blocked_services](resources--aws_tgw_site--reference--group-002.md#canonical-3220312023213323-0222323213213020-3313231022111310-1322001320032202-1021333312223203-0100222312131310-2203300011012010-3103320210301101)
- [blocked_services.blocked_service](resources--aws_tgw_site--reference--group-002.md#canonical-3330321021120220-0211220102310003-2312323202103000-1212113200003211-1110303112200210-1003321013202112-3202030102311132-2210312332310310)
- blocked_services.blocked_service.SSH

<a id="canonical-2300120013121203-0200230303333200-1130003323210323-3303110301202213-0232121120311210-2123313102331133-1322000313333232-2023333110333310"></a>

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

<a id="canonical-0130221133133222-3010203202121201-0012223212333013-3010220201013122-0001032210303202-3000210333302032-0303333202012210-1120031131220230"></a>

## Direct properties — SSH / 010302002120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1013210101212203-0333110001221121-2022033131012000-2112331102131232-2020131311200033-2233201320100111-1023313223300200-0131030110202303"></a>

## Next pages — SSH / 010302002120 / 4

- [blocked_services.blocked_service](resources--aws_tgw_site--reference--group-002.md#canonical-3330321021120220-0211220102310003-2312323202103000-1212113200003211-1110303112200210-1003321013202112-3202030102311132-2210312332310310)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-0331001001211212-0112333223132230-1202112232020332-1000310030213223-2121121313201202-1020113330330100-1201303212013222-3322211033220310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122312333220201-2011233111003211-1011111033112302-0323200021111101-1332023213121033-1312231030101210-1130303031212200-3332230002121310"></a>

## blocked_services.blocked_service.web_user_interface — web_user_interface / 331222020000 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [blocked_services](resources--aws_tgw_site--reference--group-002.md#canonical-3220312023213323-0222323213213020-3313231022111310-1322001320032202-1021333312223203-0100222312131310-2203300011012010-3103320210301101)
- [blocked_services.blocked_service](resources--aws_tgw_site--reference--group-002.md#canonical-3330321021120220-0211220102310003-2312323202103000-1212113200003211-1110303112200210-1003321013202112-3202030102311132-2210312332310310)
- blocked_services.blocked_service.web_user_interface

<a id="canonical-0100000011322010-0000112302101201-1033121000101300-2230123232130323-3033123132202010-1221222212321021-0033022102211230-2200202113103033"></a>

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
web_user_interface = {}
```

<a id="canonical-3030102230311121-0023013130213300-2101122331012220-0022021002230310-0223013112011111-3333300223113231-1332130001322120-2312220212110123"></a>

## Direct properties — web_user_interface / 331222020000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3310112000000123-3232133312132003-2121022023223120-0032201223032200-0322012130221102-0323011113010133-0023010332302133-3230022030130102"></a>

## Next pages — web_user_interface / 331222020000 / 4

- [blocked_services.blocked_service](resources--aws_tgw_site--reference--group-002.md#canonical-3330321021120220-0211220102310003-2312323202103000-1212113200003211-1110303112200210-1003321013202112-3202030102311132-2210312332310310)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-3330100032212111-2023231232130113-2033223012210111-0102030223332333-1003211333301322-2022000120201201-1202112101331102-3200002321100010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212210012033210-0133203223202020-0301112213310031-0022023030300122-3230102100011102-3112331133201133-1302101312030322-3202231301032213"></a>

## coordinates — coordinates / 221030223220 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- coordinates

<a id="canonical-1322130112300230-2201202131210003-3110330212222003-3200131313331101-2031002300011121-0103333022000001-0101232323121033-3133102321032320"></a>

Type: `"object"`. single nested block, Optional.

Coordinates of the site which provides the site physical location.

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
coordinates {
  # Configure direct properties listed below.
}
```

<a id="canonical-1011323303132013-3321222230023221-3020002020122221-1223011202320302-3031021303233122-2320013322220001-0221232321110002-2100230322220111"></a>

## Direct properties — coordinates / 221030223220 / 3

<a id="canonical-2203232223332003-3121331313020231-2202132302103223-3102212232003033-1132213210200031-1130232113120103-1110031111032030-1331031331120302"></a>

<a id="canonical-2322103202322031-0022331023300212-0333203022332022-3011023333231122-0312300120131021-3113013320323322-1022031102102301-0220132102303101"></a>

## latitude property — coordinates / 221030223220 / 4

Type: `"number"`. Optional.

Latitude. Latitude of the site location.

Upstream description:

Latitude of the site location.

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
    "ves.io.schema.rules.float.gte": "-90.0",
    "ves.io.schema.rules.float.lte": "90.0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-90.0",
    "ves.io.schema.rules.float.lte": "90.0"
  }
}
```

<a id="canonical-1203223312231130-1313333013331213-1203200202332120-1130002331110103-0030033320301123-1311332210123030-0132010103230102-1233033030311110"></a>

<a id="canonical-0333121011101003-3213331233020133-1233001023232213-0020122313102212-3213333312133213-1123131000112223-0123333300303322-2120233122223110"></a>

## longitude property — coordinates / 221030223220 / 5

Type: `"number"`. Optional.

Longitude. Longitude of site location.

Upstream description:

Longitude of site location.

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
    "ves.io.schema.rules.float.gte": "-180.0",
    "ves.io.schema.rules.float.lte": "180.0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-180.0",
    "ves.io.schema.rules.float.lte": "180.0"
  }
}
```

<a id="canonical-2313323223001001-0113231213110030-2021311210322312-1023122130331103-0121211011101202-2120013222203222-1221201120311023-2201322230032322"></a>

## Next pages — coordinates / 221030223220 / 6

- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-0310130113103331-1321231230201233-0112322012222330-2020103011100211-0311022022233022-2032112020200200-0003021332313331-1212231013022131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0002031113203111-2031022232001123-2101222010011123-1023313002103231-1031213321330130-0313001132133321-1220220223121223-3321011223333121"></a>

## custom_dns — custom_dns / 113132201322 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- custom_dns

<a id="canonical-3301212133232220-2000311102031332-2002222321000123-1123032130321021-1030202011310010-3222130302212312-3301212333220031-1330000221313110"></a>

Type: `"object"`. single nested block, Optional.

Custom DNS is the configured for specify CE site.

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
custom_dns {
  # Configure direct properties listed below.
}
```

<a id="canonical-2213320010133232-1010111302302101-0131212232110101-3211221212320113-0013332322110222-0320003301030030-3311112122222020-0100122221033020"></a>

## Direct properties — custom_dns / 113132201322 / 3

<a id="canonical-3202031013002113-1220210120210232-2212103110111213-2211022200231223-1212201101222113-3301323311312122-3031311333221313-0322303120033321"></a>

<a id="canonical-1011202131000300-2331003302110323-3211111101110321-2232120300222300-3100023232112102-3221230221002211-2230300001000222-0311322113232300"></a>

## inside_nameserver property — custom_dns / 113132201322 / 4

Type: `"string"`. Optional.

Optional DNS server IP to be used for name resolution in inside network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-1002100023101022-3013312000013012-2213111001322011-3231102000002022-1202021122303312-2120121203032222-0100231320212221-2133221001303032"></a>

<a id="canonical-3202111212001200-3000231211131310-3232301312003210-0110101213212121-0301330230020002-3310013331031232-1020332232213230-0211200201111022"></a>

## outside_nameserver property — custom_dns / 113132201322 / 5

Type: `"string"`. Optional.

Optional DNS server IP to be used for name resolution in outside network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-1020223100332232-3301210223300132-1203213123331003-0232100220113111-2213310022020201-0330211133222012-1002310103332311-2200331200131030"></a>

## Next pages — custom_dns / 113132201322 / 6

- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-0232210132032302-2311321302133012-3210002210020112-0120013313103120-1000032132332201-1332320121032013-1330221230222303-1221113233030112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121011103213033-2210333302223001-3300220111220013-1310200033323222-3333302233311013-1011200132003010-1012022102002331-1113313000310013"></a>

## default_blocked_services — default_blocked_services / 231022210312 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- default_blocked_services

<a id="canonical-2000132300021100-3032231133121301-3320213021031121-3000122030011111-0212331022121323-3002011020221112-0222023233121020-3032212002101311"></a>

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
default_blocked_services = {}
```

<a id="canonical-3103332213312321-2130312233112001-3221320323123013-2230012112111123-2033331302222030-0121320331030031-1202210323122030-0320301132101120"></a>

## Direct properties — default_blocked_services / 231022210312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1211201013032223-1113121300112310-1311100213122111-0112101332201012-0321312132202010-1032100312030310-2132311033113031-1112021122231012"></a>

## Next pages — default_blocked_services / 231022210312 / 4

- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-0300022120212210-3312311111132100-2030111122121212-3323031300112111-1131201020023012-1112013201302223-0312310103022301-3030112201131321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232031221302301-3233002311333132-1133031000231202-2330103122303311-3113303113331001-0103331001323001-3113112321100110-1211001130101021"></a>

## direct_connect_disabled — direct_connect_disabled / 321210332312 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- direct_connect_disabled

<a id="canonical-3122301210323203-0212333020332103-3222231132103033-0120023200212032-1133103322211222-1022132010311032-1113200013022001-2103001020300111"></a>

Type: `["object", {}]`. Optional.

\[OneOf: direct\_connect\_disabled, direct\_connect\_enabled, private\_connectivity\] Enable this
option

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

- [direct_connect_disabled](resources--aws_tgw_site--reference--group-002.md#canonical-3122301210323203-0212333020332103-3222231132103033-0120023200212032-1133103322211222-1022132010311032-1113200013022001-2103001020300111)
- [direct_connect_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-0101210101031120-0031320100323200-0003123032213330-2111000120012310-3013021323212201-2113331121220033-1320011211103123-3030003300300023)
- [private_connectivity](resources--aws_tgw_site--reference--group-002.md#canonical-0032101031203303-1213220331310310-0100112031003321-2233001013333112-0313130131201302-3200303332021210-2121212012311110-1202302003101320)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
direct_connect_disabled = {}
```

<a id="canonical-2303312121032211-0232333123312313-2211133222310301-1332111033332310-1131313303032103-0000032213330013-0033132332200200-3123011203313230"></a>

## Direct properties — direct_connect_disabled / 321210332312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1331201010101232-2133000331230231-1133231221310131-1220322312012102-2020310223101120-2311312130221003-1232000330012300-3022121002302031"></a>

## Next pages — direct_connect_disabled / 321210332312 / 4

- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-2230200213030222-3210301232310133-3003333211130302-2032210021222010-3232300331010023-0232100221031011-3110102001002130-0222113323033123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000303002131002-0021323122023010-1130131003133212-0230033103022113-2311121211210013-3313031230200001-3011122211233030-3303212222103101"></a>

## direct_connect_enabled — direct_connect_enabled / 222312000103 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- direct_connect_enabled

<a id="canonical-0101210101031120-0031320100323200-0003123032213330-2111000120012310-3013021323212201-2113331121220033-1320011211103123-3030003300300023"></a>

Type: `"object"`. single nested block, Optional.

Direct Connect Configuration. Direct Connect Configuration.

Upstream description:

Direct Connect Configuration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto_asn",
    "custom_asn"),
  validators.ConflictingObjectAttributes("hosted_vifs",
    "standard_vifs")}
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
  "x-ves-oneof-field-asn_choice": "[\"auto_asn\",\"custom_asn\"]",
  "x-ves-oneof-field-vif_choice": "[\"hosted_vifs\",\"standard_vifs\"]"
}
```

Terraform syntax:

```terraform
direct_connect_enabled {
  # Configure direct properties listed below.
}
```

<a id="canonical-2310320131202230-2003201010332320-0003032302202002-3323313323221300-0303120331122220-1201203032313113-0111111123031220-3330010320230103"></a>

## Direct properties — direct_connect_enabled / 222312000103 / 3

- [auto_asn](resources--aws_tgw_site--reference--group-002.md#canonical-1212212231300100-3002120022310313-2111300031220131-3012121123321321-0022032223021300-0000311222223301-0131030232112302-2220110231321023): complete subsection reference.

<a id="canonical-2313213201302220-3211201102221122-3230330010001211-2002302330213032-1210133120210103-2313330332120033-2311333011210102-3211303130330200"></a>

<a id="canonical-1223302000223312-1323002313303331-3333013101300012-0001003000113233-2222000323312120-0312112100032003-0202312121203331-0033020123021112"></a>

## custom_asn property — direct_connect_enabled / 222312000103 / 4

Type: `"number"`. Optional.

Exclusive with \[auto\_asn\] Custom Autonomous System Number.

Upstream description:

Exclusive with \[auto\_asn\] Custom Autonomous System Number.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [hosted_vifs](resources--aws_tgw_site--reference--group-002.md#canonical-2332030220330220-0011312032330110-3032020120103013-2010112332321323-3113221112030013-1012122201123002-2030023302301131-0101003301122300): complete subsection reference.

- [standard_vifs](resources--aws_tgw_site--reference--group-002.md#canonical-0203322303111331-2203130323333113-1223320210331012-3020332313111013-3323101013301101-3033212221200222-2013111020230323-2103231132202000): complete subsection reference.

<a id="canonical-2230033231131020-3212220300303032-3032100212133231-3120313020330222-2022112002200101-3201021010003322-3231203013310203-1222102021131230"></a>

## Next pages — direct_connect_enabled / 222312000103 / 5

- [direct_connect_enabled.auto_asn](resources--aws_tgw_site--reference--group-002.md#canonical-1212212231300100-3002120022310313-2111300031220131-3012121123321321-0022032223021300-0000311222223301-0131030232112302-2220110231321023)
- [direct_connect_enabled.hosted_vifs](resources--aws_tgw_site--reference--group-002.md#canonical-2332030220330220-0011312032330110-3032020120103013-2010112332321323-3113221112030013-1012122201123002-2030023302301131-0101003301122300)
- [direct_connect_enabled.standard_vifs](resources--aws_tgw_site--reference--group-002.md#canonical-0203322303111331-2203130323333113-1223320210331012-3020332313111013-3323101013301101-3033212221200222-2013111020230323-2103231132202000)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-1212212231300100-3002120022310313-2111300031220131-3012121123321321-0022032223021300-0000311222223301-0131030232112302-2220110231321023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031210203330132-1001032022322103-3212010120110111-0030331222100101-0030333020332023-0031230230312223-1310012221323120-3213023100231122"></a>

## direct_connect_enabled.auto_asn — auto_asn / 113102211330 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [direct_connect_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-2230200213030222-3210301232310133-3003333211130302-2032210021222010-3232300331010023-0232100221031011-3110102001002130-0222113323033123)
- direct_connect_enabled.auto_asn

<a id="canonical-3121232211122213-2201321102313323-3313232122120001-0203232112130021-3332110312311003-1313001210220231-2133312220123231-3212233223132130"></a>

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
auto_asn = {}
```

<a id="canonical-1101000003233022-0320112211132030-3322133230001011-3021031033122122-2103123101002032-2000113001233220-0030102023020123-1031311330022331"></a>

## Direct properties — auto_asn / 113102211330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0131113021001301-1003331303231210-1223000202030302-2030332311231312-1130123120101003-3123012110102203-0023330021210121-0023203021020133"></a>

## Next pages — auto_asn / 113102211330 / 4

- [direct_connect_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-2230200213030222-3210301232310133-3003333211130302-2032210021222010-3232300331010023-0232100221031011-3110102001002130-0222113323033123)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-2332030220330220-0011312032330110-3032020120103013-2010112332321323-3113221112030013-1012122201123002-2030023302301131-0101003301122300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0022203232101200-0213010111121233-1103122232033333-2302223231011001-3310310331323232-3122033102211111-3010103021120330-0230321223001121"></a>

## direct_connect_enabled.hosted_vifs — hosted_vifs / 313132100302 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [direct_connect_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-2230200213030222-3210301232310133-3003333211130302-2032210021222010-3232300331010023-0232100221031011-3110102001002130-0222113323033123)
- direct_connect_enabled.hosted_vifs

<a id="canonical-0101213200030233-1002131032330333-3322132312031301-2222130310133211-3021123001303033-0223032300213333-0031103331211302-0033111320332320"></a>

Type: `"object"`. single nested block, Optional.

AWS Direct Connect Hosted VIF Configuration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("site_registration_over_direct_connect",
    "site_registration_over_internet")}
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
  "x-ves-oneof-field-connectivity_options": "[\"site_registration_over_direct_connect\",\"site_registration_over_internet\"]"
}
```

Terraform syntax:

```terraform
hosted_vifs {
  # Configure direct properties listed below.
}
```

<a id="canonical-0303003233300303-0013012111230033-3301103112023113-3003223000213230-2010110231121331-3233303213221023-3210010302310202-3331332230200212"></a>

## Direct properties — hosted_vifs / 313132100302 / 3

- [site_registration_over_direct_connect](resources--aws_tgw_site--reference--group-002.md#canonical-3211300021320132-0112201232011331-3011312231111300-3222102132100313-1013223121333313-0103003121102121-2220211033331200-2230300013010200): complete subsection reference.

- [site_registration_over_internet](resources--aws_tgw_site--reference--group-002.md#canonical-2233303031303323-2200031230302211-1120123022102133-0030301332213231-3313302220320011-2101210010103233-1312233312211302-3332120200130330): complete subsection reference.

- [vif_list](resources--aws_tgw_site--reference--group-002.md#canonical-0011221312000023-0333022123120120-2230132022303332-2133231000121113-0201013333223313-1013022112012032-3330202203123231-0023030233200303): complete subsection reference.

<a id="canonical-0102223231011331-1110331032120230-3012003230013010-3002032322110302-2032133102021032-1220022230122323-0233121310200230-0333211323230222"></a>

## Next pages — hosted_vifs / 313132100302 / 4

- [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect](resources--aws_tgw_site--reference--group-002.md#canonical-3211300021320132-0112201232011331-3011312231111300-3222102132100313-1013223121333313-0103003121102121-2220211033331200-2230300013010200)
- [direct_connect_enabled.hosted_vifs.site_registration_over_internet](resources--aws_tgw_site--reference--group-002.md#canonical-2233303031303323-2200031230302211-1120123022102133-0030301332213231-3313302220320011-2101210010103233-1312233312211302-3332120200130330)
- [direct_connect_enabled.hosted_vifs.vif_list](resources--aws_tgw_site--reference--group-002.md#canonical-0011221312000023-0333022123120120-2230132022303332-2133231000121113-0201013333223313-1013022112012032-3330202203123231-0023030233200303)
- [direct_connect_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-2230200213030222-3210301232310133-3003333211130302-2032210021222010-3232300331010023-0232100221031011-3110102001002130-0222113323033123)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-3211300021320132-0112201232011331-3011312231111300-3222102132100313-1013223121333313-0103003121102121-2220211033331200-2230300013010200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322012323001112-3213003120010212-2110112122111231-0101102130311130-1333201200210230-0323231131002101-2122030013211313-2212313332202233"></a>

## direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect — site_registration_over_direct_connect / 120203231000 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [direct_connect_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-2230200213030222-3210301232310133-3003333211130302-2032210021222010-3232300331010023-0232100221031011-3110102001002130-0222113323033123)
- [direct_connect_enabled.hosted_vifs](resources--aws_tgw_site--reference--group-002.md#canonical-2332030220330220-0011312032330110-3032020120103013-2010112332321323-3113221112030013-1012122201123002-2030023302301131-0101003301122300)
- direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect

<a id="canonical-1302123302001321-3300232121100232-1102223012202033-0322220000222331-0323101112003120-3210001333212202-3310100023221212-3111023100031112"></a>

Type: `"object"`. single nested block, Optional.

CloudLink ADN Network Config.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("cloudlink_network_name")}
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
site_registration_over_direct_connect {
  # Configure direct properties listed below.
}
```

<a id="canonical-1121211333303210-2321313233322133-0201200011210100-1212220132122131-3110122332003232-1212111012022331-0303010002202231-3101321113221300"></a>

## Direct properties — site_registration_over_direct_connect / 120203231000 / 3

<a id="canonical-2020113230220132-0320210213313323-2023031331130120-2003110020130301-3320231023031211-3003322212133100-1323132313203101-2100232310331321"></a>

<a id="canonical-1003332231201311-1231310132321310-1213233022232333-0131003002131222-1212122303201011-0220001311311302-0032203110031233-0133133010202200"></a>

## cloudlink_network_name property — site_registration_over_direct_connect / 120203231000 / 4

Type: `"string"`. Optional.

Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN
network. To provision a Private ADN network, please contact F5 Distributed Cloud support.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0002021332132323-2232112031303103-3133032332221232-2003303122013122-1211223200230133-3322000103120300-1030300310023321-1013233332220232"></a>

## Next pages — site_registration_over_direct_connect / 120203231000 / 5

- [direct_connect_enabled.hosted_vifs](resources--aws_tgw_site--reference--group-002.md#canonical-2332030220330220-0011312032330110-3032020120103013-2010112332321323-3113221112030013-1012122201123002-2030023302301131-0101003301122300)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-2233303031303323-2200031230302211-1120123022102133-0030301332213231-3313302220320011-2101210010103233-1312233312211302-3332120200130330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303132122230011-3222011003231013-0121331101132313-1322012232300211-2200320330310132-2310001113302030-1113202133313210-1303322000233130"></a>

## direct_connect_enabled.hosted_vifs.site_registration_over_internet — site_registration_over_internet / 310111311123 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [direct_connect_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-2230200213030222-3210301232310133-3003333211130302-2032210021222010-3232300331010023-0232100221031011-3110102001002130-0222113323033123)
- [direct_connect_enabled.hosted_vifs](resources--aws_tgw_site--reference--group-002.md#canonical-2332030220330220-0011312032330110-3032020120103013-2010112332321323-3113221112030013-1012122201123002-2030023302301131-0101003301122300)
- direct_connect_enabled.hosted_vifs.site_registration_over_internet

<a id="canonical-1220203121310223-3032330333023223-3111211023101030-1020231331311101-3033032121300022-3002112230120120-2011301131131032-1302211102000210"></a>

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
site_registration_over_internet = {}
```

<a id="canonical-1322230120003312-1332131223210320-3012001321321202-1102100333210303-2333303230111323-1102033113113122-2211023332131300-0233030310332210"></a>

## Direct properties — site_registration_over_internet / 310111311123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1300203213023133-3021220113303203-0132032102021331-3233030211030123-2210210033313221-3133211200101333-0001022321222230-2301130121320132"></a>

## Next pages — site_registration_over_internet / 310111311123 / 4

- [direct_connect_enabled.hosted_vifs](resources--aws_tgw_site--reference--group-002.md#canonical-2332030220330220-0011312032330110-3032020120103013-2010112332321323-3113221112030013-1012122201123002-2030023302301131-0101003301122300)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-0011221312000023-0333022123120120-2230132022303332-2133231000121113-0201013333223313-1013022112012032-3330202203123231-0023030233200303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123100231311023-0002033112221333-0322312313132231-0320223201202300-2232012311212022-0113311211101320-2102100210003323-0101202102013013"></a>

## direct_connect_enabled.hosted_vifs.vif_list — vif_list / 121030301321 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [direct_connect_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-2230200213030222-3210301232310133-3003333211130302-2032210021222010-3232300331010023-0232100221031011-3110102001002130-0222113323033123)
- [direct_connect_enabled.hosted_vifs](resources--aws_tgw_site--reference--group-002.md#canonical-2332030220330220-0011312032330110-3032020120103013-2010112332321323-3113221112030013-1012122201123002-2030023302301131-0101003301122300)
- direct_connect_enabled.hosted_vifs.vif_list

<a id="canonical-1102332011103332-3003102021133331-2023221300213012-1220130210320033-3211010122103223-1330212212021120-0010123001122201-1313202111312220"></a>

Type: `"object"`. list nested block, Optional.

List of Hosted VIF Config. List of Hosted VIF Config.

Upstream description:

List of Hosted VIF Config.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("vif_id"),
  validators.ConflictingListObjectAttributes("other_region",
    "same_as_site_region")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 30,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 30,
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
    "ves.io.schema.rules.repeated.max_items": "30",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "30",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
vif_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1303130330330022-1213221020330333-3300032200322031-2100203222213020-1200202220020001-1233110001213231-3201222113013203-0000310132323003"></a>

## Direct properties — vif_list / 121030301321 / 3

<a id="canonical-1230333100130313-3123210233300333-0320133111300223-1210131101221130-2211310203333230-2010002100311202-0112210212223033-1322133330111013"></a>

<a id="canonical-1023233302323232-2200300122032112-2131010113100000-0130223311100233-1231033320302301-1331013123102233-1023333331200131-1103302233313012"></a>

## other_region property — vif_list / 121030301321 / 4

Type: `"string"`. Optional.

\[Enum:
af-south-1|ap-east-1|ap-northeast-1|ap-northeast-2|ap-south-1|ap-southeast-1|ap-southeast-2|ap-southeast-3|ca-central-1|eu-central-1|eu-north-1|eu-south-1|eu-west-1|eu-west-2|eu-west-3|me-south-1|sa-east-1|us-east-1|us-east-2|us-west-1|us-west-2\]
Exclusive with \[same\_as\_site\_region\] Other Region. Possible values are \`af-south-1\`,
\`ap-east-1\`, \`ap-northeast-1\`, \`ap-northeast-2\`, \`ap-south-1\`, \`ap-southeast-1\`,
\`ap-southeast-2\`, \`ap-southeast-3\`, \`ca-central-1\`, \`eu-central-1\`, \`eu-north-1\`,
\`eu-south-1\`, \`eu-west-1\`, \`eu-west-2\`, \`eu-west-3\`, \`me-south-1\`, \`sa-east-1\`,
\`us-east-1\`, \`us-east-2\`, \`us-west-1\`, \`us-west-2\`.

Upstream description:

Exclusive with \[same\_as\_site\_region\] Other Region.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["af-south-1","ap-east-1","ap-northeast-1","ap-northeast-2","ap-south-1","ap-southeast-1","ap-southeast-2","ap-southeast-3","ca-central-1","eu-central-1","eu-north-1","eu-south-1","eu-west-1","eu-west-2","eu-west-3","me-south-1","sa-east-1","us-east-1","us-east-2","us-west-1","us-west-2"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("af-south-1",
    "ap-east-1",
    "ap-northeast-1",
    "ap-northeast-2",
    "ap-south-1",
    "ap-southeast-1",
    "ap-southeast-2",
    "ap-southeast-3",
    "ca-central-1",
    "eu-central-1",
    "eu-north-1",
    "eu-south-1",
    "eu-west-1",
    "eu-west-2",
    "eu-west-3",
    "me-south-1",
    "sa-east-1",
    "us-east-1",
    "us-east-2",
    "us-west-1",
    "us-west-2"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "af-south-1",
    "ap-east-1",
    "ap-northeast-1",
    "ap-northeast-2",
    "ap-south-1",
    "ap-southeast-1",
    "ap-southeast-2",
    "ap-southeast-3",
    "ca-central-1",
    "eu-central-1",
    "eu-north-1",
    "eu-south-1",
    "eu-west-1",
    "eu-west-2",
    "eu-west-3",
    "me-south-1",
    "sa-east-1",
    "us-east-1",
    "us-east-2",
    "us-west-1",
    "us-west-2"
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
    "ves.io.schema.rules.string.in": "[\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-northeast-1\\\",\\\"ap-northeast-2\\\",\\\"ap-south-1\\\",\\\"ap-southeast-1\\\",\\\"ap-southeast-2\\\",\\\"ap-southeast-3\\\",\\\"ca-central-1\\\",\\\"eu-central-1\\\",\\\"eu-north-1\\\",\\\"eu-south-1\\\",\\\"eu-west-1\\\",\\\"eu-west-2\\\",\\\"eu-west-3\\\",\\\"me-south-1\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-1\\\",\\\"us-west-2\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-northeast-1\\\",\\\"ap-northeast-2\\\",\\\"ap-south-1\\\",\\\"ap-southeast-1\\\",\\\"ap-southeast-2\\\",\\\"ap-southeast-3\\\",\\\"ca-central-1\\\",\\\"eu-central-1\\\",\\\"eu-north-1\\\",\\\"eu-south-1\\\",\\\"eu-west-1\\\",\\\"eu-west-2\\\",\\\"eu-west-3\\\",\\\"me-south-1\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-1\\\",\\\"us-west-2\\\"]"
  }
}
```

- [same_as_site_region](resources--aws_tgw_site--reference--group-002.md#canonical-1103203201203221-2232200121201212-3001332002021010-3011110001033223-3033312013300232-3010110101202312-2123331310303333-2203213223132331): complete subsection reference.

<a id="canonical-3311133220111202-1201130302031100-3113022103321013-2011301233030231-1313013000002320-3302022213330223-1010331010121323-0123313021031111"></a>

<a id="canonical-2003200200220221-0321111101002013-0230011223102033-1132103311233310-2221230111120233-2333221003332122-3113313001202333-0123021211300212"></a>

## vif_id property — vif_list / 121030301321 / 5

Type: `"string"`. Optional.

AWS Direct Connect VIF ID that needs to be connected to the site.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "pattern": "^(dxvif-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^(dxvif-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^(dxvif-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

<a id="canonical-1232213132203322-1100131220333032-1122311203222130-2320002102320101-2303100112123131-1020210113300120-2030310311232031-0033320111012010"></a>

## Next pages — vif_list / 121030301321 / 6

- [direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region](resources--aws_tgw_site--reference--group-002.md#canonical-1103203201203221-2232200121201212-3001332002021010-3011110001033223-3033312013300232-3010110101202312-2123331310303333-2203213223132331)
- [direct_connect_enabled.hosted_vifs](resources--aws_tgw_site--reference--group-002.md#canonical-2332030220330220-0011312032330110-3032020120103013-2010112332321323-3113221112030013-1012122201123002-2030023302301131-0101003301122300)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-1103203201203221-2232200121201212-3001332002021010-3011110001033223-3033312013300232-3010110101202312-2123331310303333-2203213223132331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310303230301013-2032202110020310-2231331331020232-1122031102023331-1011110010223332-3331231113121230-1223211321130330-3100300000212221"></a>

## direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region — same_as_site_region / 222313312100 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [direct_connect_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-2230200213030222-3210301232310133-3003333211130302-2032210021222010-3232300331010023-0232100221031011-3110102001002130-0222113323033123)
- [direct_connect_enabled.hosted_vifs](resources--aws_tgw_site--reference--group-002.md#canonical-2332030220330220-0011312032330110-3032020120103013-2010112332321323-3113221112030013-1012122201123002-2030023302301131-0101003301122300)
- [direct_connect_enabled.hosted_vifs.vif_list](resources--aws_tgw_site--reference--group-002.md#canonical-0011221312000023-0333022123120120-2230132022303332-2133231000121113-0201013333223313-1013022112012032-3330202203123231-0023030233200303)
- direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region

<a id="canonical-0203022322211121-1110020013101112-1111321310121311-3311113203210201-3233322013200211-2113003132031101-2123203202003201-1112231120003200"></a>

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
same_as_site_region = {}
```

<a id="canonical-0012000000010212-1013233021212000-2221203233200200-2112013321331332-2030021221303102-1120111223212200-1101321130021333-2202131113121012"></a>

## Direct properties — same_as_site_region / 222313312100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1222133301020231-2110312310331213-0213312332201312-0122302013002130-0012131122302321-1031133203131321-1321113333202213-3000123121321323"></a>

## Next pages — same_as_site_region / 222313312100 / 4

- [direct_connect_enabled.hosted_vifs.vif_list](resources--aws_tgw_site--reference--group-002.md#canonical-0011221312000023-0333022123120120-2230132022303332-2133231000121113-0201013333223313-1013022112012032-3330202203123231-0023030233200303)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-0203322303111331-2203130323333113-1223320210331012-3020332313111013-3323101013301101-3033212221200222-2013111020230323-2103231132202000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011031033110012-3130111212312222-3223312213203213-0222001103102132-0203331121101112-3310333331332132-1012331121313001-0131101101303020"></a>

## direct_connect_enabled.standard_vifs — standard_vifs / 221013213132 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [direct_connect_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-2230200213030222-3210301232310133-3003333211130302-2032210021222010-3232300331010023-0232100221031011-3110102001002130-0222113323033123)
- direct_connect_enabled.standard_vifs

<a id="canonical-3202231200200123-0213221330301310-1113000011111010-3012320231321102-2333300202232111-2133002311220111-3031312110232311-2110112112331202"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for standard vifs.

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
standard_vifs = {}
```

<a id="canonical-1031231131213130-0033300120232231-0322032333002120-0033110331323301-0223100031320021-1233100300033023-2102013301313031-1110233113303000"></a>

## Direct properties — standard_vifs / 221013213132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2223111033200230-1332330232030323-1033000302312101-0213311030331003-2211022030223333-3331201220213033-0220322032323020-1330113012322320"></a>

## Next pages — standard_vifs / 221013213132 / 4

- [direct_connect_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-2230200213030222-3210301232310133-3003333211130302-2032210021222010-3232300331010023-0232100221031011-3110102001002130-0222113323033123)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-1133003220233223-0103011013313311-0210231331121001-1312002230012130-1301002123001123-3100032321303113-3312011031132303-0221202103220332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1302110301031112-3020002223203021-2333213300110101-3111130102200200-1312202112212110-2131333020300320-2121122113302223-3030301020220333"></a>

## kubernetes_upgrade_drain — kubernetes_upgrade_drain / 311100313031 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- kubernetes_upgrade_drain

<a id="canonical-1033200023221322-2013021020023300-1131032000223030-2302223233233033-1002030300321221-2333213323101132-1303230010021302-1311111233311103"></a>

Type: `"object"`. single nested block, Optional.

Specify how worker nodes within a site will be upgraded.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_upgrade_drain",
    "enable_upgrade_drain")}
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
  "x-ves-oneof-field-kubernetes_upgrade_drain_enable_choice": "[\"disable_upgrade_drain\",\"enable_upgrade_drain\"]"
}
```

Terraform syntax:

```terraform
kubernetes_upgrade_drain {
  # Configure direct properties listed below.
}
```

<a id="canonical-0012010022123332-0333131102322201-0131033231222331-3033201333032320-0002133111333302-0102230310103010-3330302113122120-0101231330313123"></a>

## Direct properties — kubernetes_upgrade_drain / 311100313031 / 3

- [disable_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-0320100301132103-3223313010113130-2300120333231013-2311323302211321-1332220222003320-1222121122233230-2311200133012211-0323233000232013): complete subsection reference.

- [enable_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-2103130212033323-1230331311211111-0200302310110032-2313212322210011-1232112123132331-1122103302221313-0000323122222102-3131303021131311): complete subsection reference.

<a id="canonical-1301130313331311-0032002312223311-2003031201210333-1211112232001212-1222032003202030-3330121321303032-2323213030302122-1122221310000223"></a>

## Next pages — kubernetes_upgrade_drain / 311100313031 / 4

- [kubernetes_upgrade_drain.disable_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-0320100301132103-3223313010113130-2300120333231013-2311323302211321-1332220222003320-1222121122233230-2311200133012211-0323233000232013)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-2103130212033323-1230331311211111-0200302310110032-2313212322210011-1232112123132331-1122103302221313-0000323122222102-3131303021131311)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-0320100301132103-3223313010113130-2300120333231013-2311323302211321-1332220222003320-1222121122233230-2311200133012211-0323233000232013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233032003033232-2232201210122303-3000020212031312-1033312313022310-3130010222010133-1003110023113331-3311230102321123-3303312212122222"></a>

## kubernetes_upgrade_drain.disable_upgrade_drain — disable_upgrade_drain / 101302232022 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [kubernetes_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-1133003220233223-0103011013313311-0210231331121001-1312002230012130-1301002123001123-3100032321303113-3312011031132303-0221202103220332)
- kubernetes_upgrade_drain.disable_upgrade_drain

<a id="canonical-2321311030030323-3303130110223011-1313111023330230-2312101022101201-0311101131101120-2033221321112030-0131032131210332-0000200303113003"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable upgrade drain.

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
disable_upgrade_drain = {}
```

<a id="canonical-3221021032210100-0031113112321013-1332201030321110-1332230213013221-2310300003322322-0321001230113332-1300130310011323-3310312222200110"></a>

## Direct properties — disable_upgrade_drain / 101302232022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1200011103330322-1222331311202201-2332300000213332-3233331003320031-3123111021222231-3101033202321203-0010211133333111-0202000011311233"></a>

## Next pages — disable_upgrade_drain / 101302232022 / 4

- [kubernetes_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-1133003220233223-0103011013313311-0210231331121001-1312002230012130-1301002123001123-3100032321303113-3312011031132303-0221202103220332)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-2103130212033323-1230331311211111-0200302310110032-2313212322210011-1232112123132331-1122103302221313-0000323122222102-3131303021131311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0102323121311001-0011312030313303-1032213322211131-3100122230010113-1020013313110031-3021023031230213-1003220031003200-2001213022331220"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain — enable_upgrade_drain / 202013332220 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [kubernetes_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-1133003220233223-0103011013313311-0210231331121001-1312002230012130-1301002123001123-3100032321303113-3312011031132303-0221202103220332)
- kubernetes_upgrade_drain.enable_upgrade_drain

<a id="canonical-1200232111122311-3110100123033033-2022302120121002-3131321030211202-1010212032220013-0312220211101223-0221011123231311-0033131231302222"></a>

Type: `"object"`. single nested block, Optional.

Specify batch upgrade settings for worker nodes within a site.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("drain_node_timeout"),
  validators.ConflictingObjectAttributes("disable_vega_upgrade_mode",
    "enable_vega_upgrade_mode"),
  validators.ConflictingObjectAttributes("drain_max_unavailable_node_count",
    "drain_max_unavailable_node_percentage")}
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
  "x-ves-oneof-field-drain_max_unavailable_choice": "[\"drain_max_unavailable_node_count\", \"drain_max_unavailable_node_percentage\"]",
  "x-ves-oneof-field-vega_upgrade_mode_toggle_choice": "[\"disable_vega_upgrade_mode\",\"enable_vega_upgrade_mode\"]"
}
```

Terraform syntax:

```terraform
enable_upgrade_drain {
  # Configure direct properties listed below.
}
```

<a id="canonical-3321033101011102-0123011133303122-3202012112133312-2012000023020301-1230311201313320-2322021201020123-0233112001111132-2100130130232102"></a>

## Direct properties — enable_upgrade_drain / 202013332220 / 3

- [disable_vega_upgrade_mode](resources--aws_tgw_site--reference--group-002.md#canonical-0001000310120231-1020023311312103-1101002301112233-2100202200212013-2113201332131122-3303111011112332-3310111133032032-0133230100232302): complete subsection reference.

<a id="canonical-0100032102010030-0211110231131030-3310031320313032-3232302113201301-3311300030323300-2133012231131203-0212231330100231-0033131133102122"></a>

<a id="canonical-1233232333011123-1113313011223202-2133313121133113-0103033230030023-2133313030302222-3012121303322233-3202013022313320-3123002103012311"></a>

## drain_max_unavailable_node_count property — enable_upgrade_drain / 202013332220 / 4

Type: `"number"`. Optional.

Node Batch Size Count. Exclusive with \[\]

Upstream description:

Exclusive with \[\]

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 5000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5000"
  }
}
```

<a id="canonical-2132020032200131-0122312123002001-2133220323330323-3210113101210332-1030131322022210-2100303003010030-3201032310123020-3310211112132233"></a>

<a id="canonical-0012211333033211-3211002030222000-1010231331001013-0103332332220120-0020223000132011-2111120301113211-3330322313032113-1031002231032202"></a>

## drain_max_unavailable_node_percentage property — enable_upgrade_drain / 202013332220 / 5

Type: `"number"`. Optional.

Maximum percentage of nodes unavailable during upgrade draining.

<a id="canonical-2332333120231310-3130220320112211-3230010203111202-3233031313002012-2332311100321131-0130312002320023-0313313012301021-2322120333330213"></a>

<a id="canonical-2002032110322300-0002010001213203-1320323120031011-0020101132210301-3322332201133123-0112013023301312-3122123021122231-3232322020201221"></a>

## drain_node_timeout property — enable_upgrade_drain / 202013332220 / 6

Type: `"number"`. Optional.

Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait
indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set
of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It
is..

Upstream description:

Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait
indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set
of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It is
recommended to use the default value).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(0, 900),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 900,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "900"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "900"
  }
}
```

- [enable_vega_upgrade_mode](resources--aws_tgw_site--reference--group-002.md#canonical-2211032113230223-1131203232203132-0223033202031211-3332110322033201-3001122313123202-3101330102011322-1012022110220013-0002201210131200): complete subsection reference.

<a id="canonical-3022003202031112-2233122212002211-1002022030321200-2123103200013322-1332313301200020-2231033313330323-2001020122012313-1020001123302033"></a>

## Next pages — enable_upgrade_drain / 202013332220 / 7

- [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](resources--aws_tgw_site--reference--group-002.md#canonical-0001000310120231-1020023311312103-1101002301112233-2100202200212013-2113201332131122-3303111011112332-3310111133032032-0133230100232302)
- [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](resources--aws_tgw_site--reference--group-002.md#canonical-2211032113230223-1131203232203132-0223033202031211-3332110322033201-3001122313123202-3101330102011322-1012022110220013-0002201210131200)
- [kubernetes_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-1133003220233223-0103011013313311-0210231331121001-1312002230012130-1301002123001123-3100032321303113-3312011031132303-0221202103220332)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-0001000310120231-1020023311312103-1101002301112233-2100202200212013-2113201332131122-3303111011112332-3310111133032032-0133230100232302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103222320130203-1113321332013001-1032212223131213-2320011332022300-3011002220332011-3320111320112233-3000012322202031-1033123011023120"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode — disable_vega_upgrade_mode / 320230311330 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [kubernetes_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-1133003220233223-0103011013313311-0210231331121001-1312002230012130-1301002123001123-3100032321303113-3312011031132303-0221202103220332)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-2103130212033323-1230331311211111-0200302310110032-2313212322210011-1232112123132331-1122103302221313-0000323122222102-3131303021131311)
- kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode

<a id="canonical-1123113300100201-1232102030333230-0001112321103330-2332133220230131-3103221113110133-3233222303032031-1322310131322323-1210200331301203"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable vega upgrade mode.

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
disable_vega_upgrade_mode = {}
```

<a id="canonical-0302133311322200-0333101310203001-2113333231132130-0220122103110300-0210322120133123-2223302003012221-1030000300222312-2122010012202103"></a>

## Direct properties — disable_vega_upgrade_mode / 320230311330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2110203031133123-3010013133333031-1223012333011321-0320032000130132-1322131020221201-2020133331332131-1021200221130111-0311323030200012"></a>

## Next pages — disable_vega_upgrade_mode / 320230311330 / 4

- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-2103130212033323-1230331311211111-0200302310110032-2313212322210011-1232112123132331-1122103302221313-0000323122222102-3131303021131311)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-2211032113230223-1131203232203132-0223033202031211-3332110322033201-3001122313123202-3101330102011322-1012022110220013-0002201210131200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223120220010313-0310211003332001-3332220222012100-3102031031332222-3330113032222012-3300102033222132-3200331131210033-3201233031321100"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode — enable_vega_upgrade_mode / 213023310302 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [kubernetes_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-1133003220233223-0103011013313311-0210231331121001-1312002230012130-1301002123001123-3100032321303113-3312011031132303-0221202103220332)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-2103130212033323-1230331311211111-0200302310110032-2313212322210011-1232112123132331-1122103302221313-0000323122222102-3131303021131311)
- kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode

<a id="canonical-2301300120132201-2230011223011021-0113233331330310-2131211201312221-1310021202212130-2031000122012101-0030323213113223-0020120003031212"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable vega upgrade mode.

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
enable_vega_upgrade_mode = {}
```

<a id="canonical-1112013103010310-0032000320232222-1300113211002122-1112002123111133-2031201231333300-3022013103211300-0222211233201111-1013222323102320"></a>

## Direct properties — enable_vega_upgrade_mode / 213023310302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3101122312323002-1113323233310313-2331331313032233-3313130332223111-2113211001011011-2200200310002321-0300100002312230-1331313223013332"></a>

## Next pages — enable_vega_upgrade_mode / 213023310302 / 4

- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-2103130212033323-1230331311211111-0200302310110032-2313212322210011-1232112123132331-1122103302221313-0000323122222102-3131303021131311)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-1103022023311031-0103303110232123-3012233320211332-2012113211332023-0001320020120012-1023313201031223-2223210013123113-1123000001332222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0332120313021231-3211122212301310-2223331133133333-3023101100312313-2031020123102022-2021213203003110-0130010303013001-1301130211120333"></a>

## log_receiver — log_receiver / 233301013322 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- log_receiver

<a id="canonical-3200330213313000-2030032223302232-2201302110100212-2112331310122100-3222022123013021-3010123110322120-3312001030023232-2103001000200103"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: log\_receiver, logs\_streaming\_disabled\] Type establishes a direct reference from one
object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

OneOf alternatives in this subsection:

- [log_receiver](resources--aws_tgw_site--reference--group-002.md#canonical-3200330213313000-2030032223302232-2201302110100212-2112331310122100-3222022123013021-3010123110322120-3312001030023232-2103001000200103)
- [logs_streaming_disabled](resources--aws_tgw_site--reference--group-002.md#canonical-2002101032021020-0020310230203133-2120133233201122-2330201230310123-1230132103102232-3201030321232221-0102033001221130-1101311022321020)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
log_receiver {
  # Configure direct properties listed below.
}
```

<a id="canonical-0132211222200220-2201022331032012-2203033133233232-1330203233230302-1233021211101211-3130102203321303-1103123210300101-1210221222222210"></a>

## Direct properties — log_receiver / 233301013322 / 3

<a id="canonical-0000202320001303-3323032101300221-3112021102003223-3023110313013100-2122130223322233-1132331232300323-2211131201212232-2301322230002022"></a>

<a id="canonical-3203003310011130-1320010122312323-1220300103223212-3010102303211210-0120231101322323-1210323312222202-3100000021201313-0132322330013231"></a>

## name property — log_receiver / 233301013322 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-3110200222111133-0222012113123120-0320320200220330-1212221000020202-2121301321333221-1323200011020300-1303023130312233-3203312231230000"></a>

<a id="canonical-1010010021321202-2233211321120031-1103312311310030-2030202330030131-1121013332122132-3300110021303032-3010223100032232-0303200222133323"></a>

## namespace property — log_receiver / 233301013322 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1332012211222331-1202230002030213-1212133223320202-1310121022001131-0330213202212001-0313133221203212-3133100120003233-2032030000302123"></a>

<a id="canonical-0211131110331013-2212022001132122-1210310222121211-0310200012131102-2331101203133330-2130112113321032-0001221213231212-0121212013103311"></a>

## tenant property — log_receiver / 233301013322 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2111322001121310-1313110223103120-0200302322301223-0211111332002013-2200123212221023-2333303021030003-2333110102330232-1202022312111000"></a>

## Next pages — log_receiver / 233301013322 / 7

- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-3331230112202213-2130303012010322-2331122233103201-0333302220133100-3033003202003300-0301310322330232-2200123022312220-2023201330131100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0110023220203330-1032303222223003-3112032110233132-1112121302121011-3331013320300123-1133112302331030-3122111312031120-1313331201312121"></a>

## logs_streaming_disabled — logs_streaming_disabled / 013120321102 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- logs_streaming_disabled

<a id="canonical-2002101032021020-0020310230203133-2120133233201122-2330201230310123-1230132103102232-3201030321232221-0102033001221130-1101311022321020"></a>

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
logs_streaming_disabled = {}
```

<a id="canonical-3200333010001220-3010202130121112-2211211032000321-1131001030121113-1102033103322200-3230322222323120-0012202130322013-1212002013032221"></a>

## Direct properties — logs_streaming_disabled / 013120321102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3131113300332130-2102211313111333-3110311312012211-1102123010311211-1122231021113331-0002210033330032-2113300220002202-1310220233331032"></a>

## Next pages — logs_streaming_disabled / 013120321102 / 4

- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-3001311020203333-3221233013320302-3311320030212213-1023113012213111-0220020321311331-1032303112200020-0010110000121231-3320101333113323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001203002310022-3211330213233100-2020312201021020-1331112233113032-0033232202013102-2110110232002201-1231030131312030-0122223202102211"></a>

## offline_survivability_mode — offline_survivability_mode / 100332130323 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- offline_survivability_mode

<a id="canonical-3300322230023203-3003210213133023-2000020032102101-2302030102313121-0313002003300123-1030030132102223-0233311001201132-0322030200021100"></a>

Type: `"object"`. single nested block, Optional.

Offline Survivability allows the Site to continue functioning normally without traffic loss during
periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this
feature is enabled, a site can continue to function as is with existing configuration for upto 7..

Upstream description:

Offline Survivability allows the Site to continue functioning normally without traffic loss during
periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this
feature is enabled, a site can continue to function as is with existing configuration for upto 7
days, even when the site is offline. The certificates needed to keep the services running on this
site are signed using a local CA. Secrets would also be cached locally to handle the connectivity
loss. When the mode is toggled, services will restart and traffic disruption will be seen.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("enable_offline_survivability_mode",
    "no_offline_survivability_mode")}
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
  "x-ves-oneof-field-offline_survivability_mode_choice": "[\"enable_offline_survivability_mode\",\"no_offline_survivability_mode\"]"
}
```

Terraform syntax:

```terraform
offline_survivability_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-3203021003130322-1320333030021111-0120113131223100-3133201332212202-1323332102032100-0131332103221112-0202230022213120-1030003120022201"></a>

## Direct properties — offline_survivability_mode / 100332130323 / 3

- [enable_offline_survivability_mode](resources--aws_tgw_site--reference--group-002.md#canonical-1330200230003111-1130322033202321-0233022322200303-2120031100031312-2033330231120100-2333300203112312-1022330022220021-3300322001131333): complete subsection reference.

- [no_offline_survivability_mode](resources--aws_tgw_site--reference--group-002.md#canonical-1211231003331122-3230110231022113-0022001200110322-0310103021323011-0000133310023113-3132000331002203-0332123011312211-1123102122322213): complete subsection reference.

<a id="canonical-2312222110120100-1102320000210130-3302102300302022-2230321321130111-3301300210312020-0222101301300212-3220013333111231-3102031321213330"></a>

## Next pages — offline_survivability_mode / 100332130323 / 4

- [offline_survivability_mode.enable_offline_survivability_mode](resources--aws_tgw_site--reference--group-002.md#canonical-1330200230003111-1130322033202321-0233022322200303-2120031100031312-2033330231120100-2333300203112312-1022330022220021-3300322001131333)
- [offline_survivability_mode.no_offline_survivability_mode](resources--aws_tgw_site--reference--group-002.md#canonical-1211231003331122-3230110231022113-0022001200110322-0310103021323011-0000133310023113-3132000331002203-0332123011312211-1123102122322213)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-1330200230003111-1130322033202321-0233022322200303-2120031100031312-2033330231120100-2333300203112312-1022330022220021-3300322001131333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113112230211000-3303331201311230-0233211020013222-1322132210202011-2022022123302222-0301031231001002-3132122300200030-0020112001211111"></a>

## offline_survivability_mode.enable_offline_survivability_mode — enable_offline_survivability_mode / 022323031211 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [offline_survivability_mode](resources--aws_tgw_site--reference--group-002.md#canonical-3001311020203333-3221233013320302-3311320030212213-1023113012213111-0220020321311331-1032303112200020-0010110000121231-3320101333113323)
- offline_survivability_mode.enable_offline_survivability_mode

<a id="canonical-3102311111032021-2102131022103122-0323013130233001-2202323223033230-3033023120201323-3001221330312231-2120012003231112-1102323011130112"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable offline survivability mode.

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
enable_offline_survivability_mode = {}
```

<a id="canonical-0302203023311330-0223132033001023-3001133000311033-1133022203302002-1220311000310020-0222330230132003-3230030321300122-2312132330330320"></a>

## Direct properties — enable_offline_survivability_mode / 022323031211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3231010331033011-2132322131121222-0210233211133313-1303212122000102-1023322200033003-3200030323310222-2311120203333112-0021010030311312"></a>

## Next pages — enable_offline_survivability_mode / 022323031211 / 4

- [offline_survivability_mode](resources--aws_tgw_site--reference--group-002.md#canonical-3001311020203333-3221233013320302-3311320030212213-1023113012213111-0220020321311331-1032303112200020-0010110000121231-3320101333113323)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-1211231003331122-3230110231022113-0022001200110322-0310103021323011-0000133310023113-3132000331002203-0332123011312211-1123102122322213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0220221200222202-2231313022031210-2210011300132320-2133202321310320-1012120033020310-0312332310331031-1013100330210213-3320300331232133"></a>

## offline_survivability_mode.no_offline_survivability_mode — no_offline_survivability_mode / 100030223221 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [offline_survivability_mode](resources--aws_tgw_site--reference--group-002.md#canonical-3001311020203333-3221233013320302-3311320030212213-1023113012213111-0220020321311331-1032303112200020-0010110000121231-3320101333113323)
- offline_survivability_mode.no_offline_survivability_mode

<a id="canonical-1210333111300321-0312011223003131-1000332133031110-2021202311030122-1012131300311001-0121202101233210-1200301212300000-0222110110212313"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no offline survivability mode.

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
no_offline_survivability_mode = {}
```

<a id="canonical-3213230330210100-1131223112202331-1121132200021022-2120010102010110-2323333311101133-0232211000001121-0003021331030320-3102000222202313"></a>

## Direct properties — no_offline_survivability_mode / 100030223221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2333310213200332-1310033030221112-1130312311201312-2000310302203102-0300001111123300-3330230313033011-2313211310130132-2103002201112120"></a>

## Next pages — no_offline_survivability_mode / 100030223221 / 4

- [offline_survivability_mode](resources--aws_tgw_site--reference--group-002.md#canonical-3001311020203333-3221233013320302-3311320030212213-1023113012213111-0220020321311331-1032303112200020-0010110000121231-3320101333113323)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-2303123002201131-0222010221210120-2033022000213311-1111020303311112-2012121320122023-3033220200130232-0231103221011210-1220321132322030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032213210021031-3300301112330212-1302232313011012-1233001031032202-2110323200033311-1210222121332221-2123133213133103-0000022230220101"></a>

## os — os / 103131211230 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- os

<a id="canonical-2022233012203233-1131120110103021-2231122321223033-1110132121110001-0110233111012332-1133131122312010-0111113200211020-0321130302300020"></a>

Type: `"object"`. single nested block, Optional.

Select the F5XC Operating System Version for the site. By default, latest available OS Version will
be used. Refer to release notes to find required released OS versions.

Upstream description:

Select the F5XC Operating System Version for the site. By default, latest available OS Version will
be used. Refer to release notes to find required released OS versions.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_os_version",
    "operating_system_version")}
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
  "x-ves-oneof-field-operating_system_version_choice": "[\"default_os_version\",\"operating_system_version\"]"
}
```

Terraform syntax:

```terraform
os {
  # Configure direct properties listed below.
}
```

<a id="canonical-3100222133222021-1021202111330231-0321131011200021-2230010211312011-1230211310230310-2122032123231211-2122031320100030-0210111332223110"></a>

## Direct properties — os / 103131211230 / 3

- [default_os_version](resources--aws_tgw_site--reference--group-002.md#canonical-2331203310211212-3302313221001221-2013123300210331-0303032311322023-3102210230221323-3000303132210212-3102230122221230-1320323110323330): complete subsection reference.

<a id="canonical-0230031123322331-1223213200113230-2012100100201131-0122120131121010-2010030203111133-1302212222112301-1131003022103310-1203310311011212"></a>

<a id="canonical-1130333201131102-3202011310213212-1332000311211323-0313333302030101-3212112001101022-1002322120333010-1031322321213031-1130230121001311"></a>

## operating_system_version property — os / 103131211230 / 4

Type: `"string"`. Optional.

Exclusive with \[default\_os\_version\] Specify a OS version to be used e.g. 9.2024.6.

Upstream description:

Exclusive with \[default\_os\_version\] Specify a OS version to be used e.g. 9.2024.6.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
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
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-2221303221321223-0123323222130032-1223302113201130-2000022003333331-0021313332222011-1123030213230213-3220002031210010-1330231321013233"></a>

## Next pages — os / 103131211230 / 5

- [os.default_os_version](resources--aws_tgw_site--reference--group-002.md#canonical-2331203310211212-3302313221001221-2013123300210331-0303032311322023-3102210230221323-3000303132210212-3102230122221230-1320323110323330)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-2331203310211212-3302313221001221-2013123300210331-0303032311322023-3102210230221323-3000303132210212-3102230122221230-1320323110323330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300210001110000-0000123023123322-3023211313121311-1103222231032201-2213210303023103-0303132031223210-1002222011333312-2203231302112022"></a>

## os.default_os_version — default_os_version / 001311310330 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [os](resources--aws_tgw_site--reference--group-002.md#canonical-2303123002201131-0222010221210120-2033022000213311-1111020303311112-2012121320122023-3033220200130232-0231103221011210-1220321132322030)
- os.default_os_version

<a id="canonical-1110320223313311-0102213232221200-0121123211033220-0220300200110212-3233132233230130-1133230212030031-1211301022233233-0020232011212100"></a>

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
default_os_version = {}
```

<a id="canonical-0220221202002130-3033113302023112-2220031320230212-3202232112111013-1323130033313212-3323323323322120-2330123001012022-2201130021103121"></a>

## Direct properties — default_os_version / 001311310330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3121123002132132-2103031211010211-1201210223200231-0102022311333121-1322020333313200-0331210331213100-2030233001202011-1022022021133210"></a>

## Next pages — default_os_version / 001311310330 / 4

- [os](resources--aws_tgw_site--reference--group-002.md#canonical-2303123002201131-0222010221210120-2033022000213311-1111020303311112-2012121320122023-3033220200130232-0231103221011210-1220321132322030)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-1003000322233122-1203103022221013-2012312321203031-3020031313331223-2123222020211330-2332120211321231-1232113122132011-2222323003332113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112321222010131-1223101011113313-3131010302320213-0022332101121222-0102032231220130-2330331120323330-1121121130233012-2310021333000221"></a>

## performance_enhancement_mode — performance_enhancement_mode / 330031133122 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- performance_enhancement_mode

<a id="canonical-0033131101311310-3332202301032222-1231111021311112-2103103213010223-3221322322311112-2231101023030002-0030233130121020-0013122120101122"></a>

Type: `"object"`. single nested block, Optional.

Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("perf_mode_l3_enhanced",
    "perf_mode_l7_enhanced")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"perf_mode_l3_enhanced\",\"perf_mode_l7_enhanced\"]"
}
```

Terraform syntax:

```terraform
performance_enhancement_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-3223013211132310-0033322320022013-2332032312220023-2021133102030310-1020223022220203-1233233200330000-3010232230000302-1231131230301321"></a>

## Direct properties — performance_enhancement_mode / 330031133122 / 3

- [perf_mode_l3_enhanced](resources--aws_tgw_site--reference--group-002.md#canonical-3003132121022311-2223212333230003-0233232222111121-3131033210122022-1121301121302013-3301022130321232-1313000102211112-3322000222130123): complete subsection reference.

- [perf_mode_l7_enhanced](resources--aws_tgw_site--reference--group-002.md#canonical-3203003101111012-0201100211313210-2100200111130231-0131130201313122-0313010331200103-2030102310003103-0013100312103033-2130200133331330): complete subsection reference.

<a id="canonical-3131231112231221-2213132000331231-1013331131132033-2121231133201023-3020313001133012-2110211210203111-3210233333321223-3121310212012213"></a>

## Next pages — performance_enhancement_mode / 330031133122 / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_tgw_site--reference--group-002.md#canonical-3003132121022311-2223212333230003-0233232222111121-3131033210122022-1121301121302013-3301022130321232-1313000102211112-3322000222130123)
- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_tgw_site--reference--group-002.md#canonical-3203003101111012-0201100211313210-2100200111130231-0131130201313122-0313010331200103-2030102310003103-0013100312103033-2130200133331330)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-3003132121022311-2223212333230003-0233232222111121-3131033210122022-1121301121302013-3301022130321232-1313000102211112-3322000222130123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3030323100132130-3321321322300310-3113322122202232-0322213023220131-2220332000321223-3103112000010303-2121212201113002-0031122303330102"></a>

## performance_enhancement_mode.perf_mode_l3_enhanced — perf_mode_l3_enhanced / 011021300312 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [performance_enhancement_mode](resources--aws_tgw_site--reference--group-002.md#canonical-1003000322233122-1203103022221013-2012312321203031-3020031313331223-2123222020211330-2332120211321231-1232113122132011-2222323003332113)
- performance_enhancement_mode.perf_mode_l3_enhanced

<a id="canonical-3031131233122022-3021100000201010-0100312021101123-0302201103003120-2222022303001011-0101132111330313-2102020030122230-3201322123332330"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for perf mode l3 enhanced.

Upstream description:

L3 enhanced performance mode OPTIONS.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("jumbo",
    "no_jumbo")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo\",\"no_jumbo\"]"
}
```

Terraform syntax:

```terraform
perf_mode_l3_enhanced {
  # Configure direct properties listed below.
}
```

<a id="canonical-1121030130223023-0302100032021001-0021201313220202-2123310222231132-1031002231121123-2122313123020311-0030233023311323-1312002023103321"></a>

## Direct properties — perf_mode_l3_enhanced / 011021300312 / 3

- [jumbo](resources--aws_tgw_site--reference--group-002.md#canonical-3321002213031313-3332121033133332-2001332233010203-1100302131120132-1013002212133101-1223323301111133-1123222310033001-2331331320223313): complete subsection reference.

- [no_jumbo](resources--aws_tgw_site--reference--group-002.md#canonical-1010332012232311-3200012011213203-0100120310321002-0302223012302120-0330213001013212-3001211012002031-2230202000313123-0230321103313311): complete subsection reference.

<a id="canonical-1132210023022002-0203000201002313-0203020200130320-2332002200100013-3330220131130033-1201323110100213-1102000031123321-3210133332330012"></a>

## Next pages — perf_mode_l3_enhanced / 011021300312 / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--aws_tgw_site--reference--group-002.md#canonical-3321002213031313-3332121033133332-2001332233010203-1100302131120132-1013002212133101-1223323301111133-1123222310033001-2331331320223313)
- [performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--aws_tgw_site--reference--group-002.md#canonical-1010332012232311-3200012011213203-0100120310321002-0302223012302120-0330213001013212-3001211012002031-2230202000313123-0230321103313311)
- [performance_enhancement_mode](resources--aws_tgw_site--reference--group-002.md#canonical-1003000322233122-1203103022221013-2012312321203031-3020031313331223-2123222020211330-2332120211321231-1232113122132011-2222323003332113)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-3321002213031313-3332121033133332-2001332233010203-1100302131120132-1013002212133101-1223323301111133-1123222310033001-2331331320223313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011120123211111-1310302023132200-0322113002011221-1120320212312030-3220322022201022-0213100300330330-0122101132030230-3112103001200032"></a>

## performance_enhancement_mode.perf_mode_l3_enhanced.jumbo — jumbo / 121131123232 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [performance_enhancement_mode](resources--aws_tgw_site--reference--group-002.md#canonical-1003000322233122-1203103022221013-2012312321203031-3020031313331223-2123222020211330-2332120211321231-1232113122132011-2222323003332113)
- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_tgw_site--reference--group-002.md#canonical-3003132121022311-2223212333230003-0233232222111121-3131033210122022-1121301121302013-3301022130321232-1313000102211112-3322000222130123)
- performance_enhancement_mode.perf_mode_l3_enhanced.jumbo

<a id="canonical-1133202232121332-2001123221223200-0200201001122002-1123021220313321-3102021220212320-2233032232213203-0311313033221012-0221122210331111"></a>

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
jumbo = {}
```

<a id="canonical-3210321232121331-2232311202200200-0032100310110302-2020013201102200-1313233023112012-0021323001203011-1333201130132333-2311310113213330"></a>

## Direct properties — jumbo / 121131123232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3011032112131211-2001133322113213-1220213000321211-1310212333312203-1132123323322230-2011112103130012-1010230330002313-1020031223033331"></a>

## Next pages — jumbo / 121131123232 / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_tgw_site--reference--group-002.md#canonical-3003132121022311-2223212333230003-0233232222111121-3131033210122022-1121301121302013-3301022130321232-1313000102211112-3322000222130123)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-1010332012232311-3200012011213203-0100120310321002-0302223012302120-0330213001013212-3001211012002031-2230202000313123-0230321103313311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003113012220301-1232003011010020-0221330330322201-3031303302121111-1233101033233122-3232123200200230-3331230331210320-2000221123310303"></a>

## performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo — no_jumbo / 232331102103 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [performance_enhancement_mode](resources--aws_tgw_site--reference--group-002.md#canonical-1003000322233122-1203103022221013-2012312321203031-3020031313331223-2123222020211330-2332120211321231-1232113122132011-2222323003332113)
- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_tgw_site--reference--group-002.md#canonical-3003132121022311-2223212333230003-0233232222111121-3131033210122022-1121301121302013-3301022130321232-1313000102211112-3322000222130123)
- performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo

<a id="canonical-2102101130231123-3003330221202203-2100033210001220-1322113203212203-2002332203313302-1113321313213123-0232022331230110-3200112000003111"></a>

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
no_jumbo = {}
```

<a id="canonical-2311322202020320-1003310312213020-3012002300322120-1031211120010320-0321232312000231-0122112322013103-0033302020101302-2332332211333023"></a>

## Direct properties — no_jumbo / 232331102103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3220111033221201-1231131302023111-1310030230213301-2011103121133102-1021111002020123-1031033131013131-0013210203102202-3210300323031022"></a>

## Next pages — no_jumbo / 232331102103 / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_tgw_site--reference--group-002.md#canonical-3003132121022311-2223212333230003-0233232222111121-3131033210122022-1121301121302013-3301022130321232-1313000102211112-3322000222130123)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-3203003101111012-0201100211313210-2100200111130231-0131130201313122-0313010331200103-2030102310003103-0013100312103033-2130200133331330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0022121330222001-2131332113232023-0330312023212220-1010122220001301-0101221020200232-2102101202201131-3313132201010300-1101220221130122"></a>

## performance_enhancement_mode.perf_mode_l7_enhanced — perf_mode_l7_enhanced / 321112110202 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [performance_enhancement_mode](resources--aws_tgw_site--reference--group-002.md#canonical-1003000322233122-1203103022221013-2012312321203031-3020031313331223-2123222020211330-2332120211321231-1232113122132011-2222323003332113)
- performance_enhancement_mode.perf_mode_l7_enhanced

<a id="canonical-3012111321231111-2101101120303013-3112002220233001-0103231321103101-1301102032000303-1121230021300303-0121131013013210-3202020210213011"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for perf mode l7 enhanced.

Upstream description:

L7 enhanced performance mode OPTIONS.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("jumbo_disabled",
    "jumbo_enabled")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo_disabled\",\"jumbo_enabled\"]"
}
```

Terraform syntax:

```terraform
perf_mode_l7_enhanced {
  # Configure direct properties listed below.
}
```

<a id="canonical-1132310302330200-3003102031200231-0130133201211301-3310200212232211-0001112201333310-1020122110233002-1000233000210133-2212123331021102"></a>

## Direct properties — perf_mode_l7_enhanced / 321112110202 / 3

- [jumbo_disabled](resources--aws_tgw_site--reference--group-002.md#canonical-2312301122223030-2313002303210011-2021210333303330-0312303212220211-1013002220122022-3201002232121313-3230213311021022-0010230312132233): complete subsection reference.

- [jumbo_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-2232003102101112-0000102001013332-2012231122013202-1100331100023300-1000000121223220-3303231231011221-0112331300123033-0012201133021100): complete subsection reference.

<a id="canonical-2013232301232110-0020323313011320-1022330003230113-1002303222223300-2023132033100221-1212110333010230-1302100213002023-1111223222011121"></a>

## Next pages — perf_mode_l7_enhanced / 321112110202 / 4

- [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--aws_tgw_site--reference--group-002.md#canonical-2312301122223030-2313002303210011-2021210333303330-0312303212220211-1013002220122022-3201002232121313-3230213311021022-0010230312132233)
- [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-2232003102101112-0000102001013332-2012231122013202-1100331100023300-1000000121223220-3303231231011221-0112331300123033-0012201133021100)
- [performance_enhancement_mode](resources--aws_tgw_site--reference--group-002.md#canonical-1003000322233122-1203103022221013-2012312321203031-3020031313331223-2123222020211330-2332120211321231-1232113122132011-2222323003332113)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-2312301122223030-2313002303210011-2021210333303330-0312303212220211-1013002220122022-3201002232121313-3230213311021022-0010230312132233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2330311320303211-2211100233001330-3302211022132232-3330302230212003-2230201333101022-2223011211030103-1123011322321120-2311130101222311"></a>

## performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled — jumbo_disabled / 131222233131 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [performance_enhancement_mode](resources--aws_tgw_site--reference--group-002.md#canonical-1003000322233122-1203103022221013-2012312321203031-3020031313331223-2123222020211330-2332120211321231-1232113122132011-2222323003332113)
- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_tgw_site--reference--group-002.md#canonical-3203003101111012-0201100211313210-2100200111130231-0131130201313122-0313010331200103-2030102310003103-0013100312103033-2130200133331330)
- performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled

<a id="canonical-3232113332032210-3033020110001332-3120303332123222-1313201303211311-1323232330320102-2210133300022211-2301131313332112-2031331001121001"></a>

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
jumbo_disabled = {}
```

<a id="canonical-3333133300120320-3303020201123003-2301213121310301-3230003023101002-3220222023312310-0112232202121023-1011132202033233-1302201333233132"></a>

## Direct properties — jumbo_disabled / 131222233131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0132222113102013-3030122333320120-0133010213321230-2123221111123001-0323110230302132-2333132213302120-1033122320031322-0210020031323130"></a>

## Next pages — jumbo_disabled / 131222233131 / 4

- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_tgw_site--reference--group-002.md#canonical-3203003101111012-0201100211313210-2100200111130231-0131130201313122-0313010331200103-2030102310003103-0013100312103033-2130200133331330)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-2232003102101112-0000102001013332-2012231122013202-1100331100023300-1000000121223220-3303231231011221-0112331300123033-0012201133021100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1013203203110222-0121132113001111-2032031221311301-2013133002131121-1230003132012323-2303210110111212-3000202013202110-3202012331113131"></a>

## performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled — jumbo_enabled / 232002221102 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [performance_enhancement_mode](resources--aws_tgw_site--reference--group-002.md#canonical-1003000322233122-1203103022221013-2012312321203031-3020031313331223-2123222020211330-2332120211321231-1232113122132011-2222323003332113)
- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_tgw_site--reference--group-002.md#canonical-3203003101111012-0201100211313210-2100200111130231-0131130201313122-0313010331200103-2030102310003103-0013100312103033-2130200133331330)
- performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled

<a id="canonical-1302132332003232-2231010300212103-3220223101330203-1030103232103133-2320221032333322-0113023013130103-0211213101221213-3021311201102112"></a>

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
jumbo_enabled = {}
```

<a id="canonical-1321033112210010-0230210322103021-2110331020113013-2002030131330331-3001302120303023-2001111103213130-3132103112310121-3131333101332322"></a>

## Direct properties — jumbo_enabled / 232002221102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1110031202123131-1031313130032022-0032000032122202-2113002221203200-2201133110300021-3332130302013230-2212122131203032-1100213220313011"></a>

## Next pages — jumbo_enabled / 232002221102 / 4

- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_tgw_site--reference--group-002.md#canonical-3203003101111012-0201100211313210-2100200111130231-0131130201313122-0313010331200103-2030102310003103-0013100312103033-2130200133331330)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-0112133200322013-2300122300221313-0113120023203023-2303120023331133-1221201211013222-2113331302223210-0212130311102330-1220131330231211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313103113033131-3223232013101131-2033302110133033-0102102100101003-2122333313310103-0111000010131012-3302232020312030-0000103000223212"></a>

## private_connectivity — private_connectivity / 232201101133 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- private_connectivity

<a id="canonical-0032101031203303-1213220331310310-0100112031003321-2233001013333112-0313130131201302-3200303332021210-2121212012311110-1202302003101320"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for private connectivity.

Upstream description:

Private Connect Configuration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("inside",
    "outside")}
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
  "x-ves-oneof-field-network_options": "[\"inside\",\"outside\"]"
}
```

Terraform syntax:

```terraform
private_connectivity {
  # Configure direct properties listed below.
}
```

<a id="canonical-0121121123323230-2210313131213011-1211302220031033-3133301200231311-2022130121311221-0323301210111333-1323201110133300-1012132113231023"></a>

## Direct properties — private_connectivity / 232201101133 / 3

- [cloud_link](resources--aws_tgw_site--reference--group-002.md#canonical-1033222232330221-0100013301112211-1303202302330111-2032023131012000-0330112130302333-2202111020220333-0331321122302132-1210221002130122): complete subsection reference.

- [inside](resources--aws_tgw_site--reference--group-002.md#canonical-2023313230113323-3212330110031301-2321233013221222-0303103323200022-3120103120231201-0000330013331130-0200202102300330-2020320010112123): complete subsection reference.

- [outside](resources--aws_tgw_site--reference--group-002.md#canonical-0120301302031320-1033020223303023-1321023321022311-1101000120300312-2200002021233212-2103030102312321-1122032310022332-2312011301133033): complete subsection reference.

<a id="canonical-0130013132020302-1320133022001031-1323332122130021-1023321322310200-0300333012002320-1232101223101201-3131112200301121-1231330202230000"></a>

## Next pages — private_connectivity / 232201101133 / 4

- [private_connectivity.cloud_link](resources--aws_tgw_site--reference--group-002.md#canonical-1033222232330221-0100013301112211-1303202302330111-2032023131012000-0330112130302333-2202111020220333-0331321122302132-1210221002130122)
- [private_connectivity.inside](resources--aws_tgw_site--reference--group-002.md#canonical-2023313230113323-3212330110031301-2321233013221222-0303103323200022-3120103120231201-0000330013331130-0200202102300330-2020320010112123)
- [private_connectivity.outside](resources--aws_tgw_site--reference--group-002.md#canonical-0120301302031320-1033020223303023-1321023321022311-1101000120300312-2200002021233212-2103030102312321-1122032310022332-2312011301133033)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-1033222232330221-0100013301112211-1303202302330111-2032023131012000-0330112130302333-2202111020220333-0331321122302132-1210221002130122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301221000333002-0310303331210102-2012320303122200-3310233000222002-1101121222132032-0022213303012003-0223220332001300-1031320121322230"></a>

## private_connectivity.cloud_link — cloud_link / 011330030330 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [private_connectivity](resources--aws_tgw_site--reference--group-002.md#canonical-0112133200322013-2300122300221313-0113120023203023-2303120023331133-1221201211013222-2113331302223210-0212130311102330-1220131330231211)
- private_connectivity.cloud_link

<a id="canonical-3111321303121122-2312011022310303-3100232113220302-2131113333123132-0011231211122320-0333233213113021-0210213211331231-1220120202231003"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
cloud_link {
  # Configure direct properties listed below.
}
```

<a id="canonical-1020332302113000-0301033333131100-1313130312111201-2230130201020033-2333001012113022-3130302001200133-3003313213221112-0102123312223311"></a>

## Direct properties — cloud_link / 011330030330 / 3

<a id="canonical-0232202221200200-0021331112120002-3210031312330320-2203122013111223-1320220211000013-3030111211001013-3022223310011123-2332320112133330"></a>

<a id="canonical-3300232021130223-3131003232012320-3011210121333000-1220133210003310-2112301032021333-3313313021300211-0211311302031323-2120002003021032"></a>

## name property — cloud_link / 011330030330 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-0013230331222100-1112002313210311-0011012202232320-0120012221100220-0303121330011123-1032322012213200-1013233013232111-3110320202202133"></a>

<a id="canonical-2110033332131202-2221102222232002-2210312322223022-0222222003032010-1010101222313032-3003000312302232-3103221013033232-2113030130233022"></a>

## namespace property — cloud_link / 011330030330 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2300311023332321-0130022203130210-2103210303300010-1210211122232211-1233011212030010-1211030123132122-3132220123321130-3320330233122120"></a>

<a id="canonical-3300331201222030-1033322013220322-0130311131013302-2303010230300002-3102000203011111-1221200130001223-3231103103012302-3120202323000300"></a>

## tenant property — cloud_link / 011330030330 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2333331130233011-1101111211130130-1101200203001320-2223130300001122-2203023312210022-3130021123320101-2300002322103203-3032313200123013"></a>

## Next pages — cloud_link / 011330030330 / 7

- [private_connectivity](resources--aws_tgw_site--reference--group-002.md#canonical-0112133200322013-2300122300221313-0113120023203023-2303120023331133-1221201211013222-2113331302223210-0212130311102330-1220131330231211)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-2023313230113323-3212330110031301-2321233013221222-0303103323200022-3120103120231201-0000330013331130-0200202102300330-2020320010112123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0231233203302313-3210102112321111-3320210000302333-3122010100021102-0331032331222020-1200300330231233-0233332013021021-0123302222321220"></a>

## private_connectivity.inside — inside / 221223103133 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [private_connectivity](resources--aws_tgw_site--reference--group-002.md#canonical-0112133200322013-2300122300221313-0113120023203023-2303120023331133-1221201211013222-2113331302223210-0212130311102330-1220131330231211)
- private_connectivity.inside

<a id="canonical-3322222031113011-0022300130320002-0103000223333032-2002132022323331-3301233111322220-1021202102322010-3321201120003223-3300111221311300"></a>

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
inside = {}
```

<a id="canonical-2330101112010322-0101212313200012-3013112011211132-1000311320210012-2302302211232230-1111010020131231-2320112322020220-2112211021211200"></a>

## Direct properties — inside / 221223103133 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0121312110110310-3121000111202200-2110003012203030-3022200210302311-2323100021103130-2022322203113102-2201011320003002-2020110110323220"></a>

## Next pages — inside / 221223103133 / 4

- [private_connectivity](resources--aws_tgw_site--reference--group-002.md#canonical-0112133200322013-2300122300221313-0113120023203023-2303120023331133-1221201211013222-2113331302223210-0212130311102330-1220131330231211)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-0120301302031320-1033020223303023-1321023321022311-1101000120300312-2200002021233212-2103030102312321-1122032310022332-2312011301133033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3111212300001010-1211031221201121-3201300113202211-3323212230120231-0010330121201212-1111130021100113-0230030200200203-1310022130310133"></a>

## private_connectivity.outside — outside / 031213302302 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [private_connectivity](resources--aws_tgw_site--reference--group-002.md#canonical-0112133200322013-2300122300221313-0113120023203023-2303120023331133-1221201211013222-2113331302223210-0212130311102330-1220131330231211)
- private_connectivity.outside

<a id="canonical-3010220230021001-3111020322112312-0012310223231331-1003222101220002-2123202202201312-3230213231033031-3321013301121021-2233222320322231"></a>

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
outside = {}
```

<a id="canonical-0311001203002220-0300110321100121-3232011330130230-3010230000231331-0321313002112222-3203110313013002-2133221111120031-0032200312103102"></a>

## Direct properties — outside / 031213302302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3323023212313211-2102113021010021-3013133122211230-1010202313223223-2330221333123002-1211012102312203-3233331110022203-2223002101211012"></a>

## Next pages — outside / 031213302302 / 4

- [private_connectivity](resources--aws_tgw_site--reference--group-002.md#canonical-0112133200322013-2300122300221313-0113120023203023-2303120023331133-1221201211013222-2113331302223210-0212130311102330-1220131330231211)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-1300131232312032-1320122110210113-3211331031330230-1322221001301333-0300033311131133-2330301300022230-1020032213210011-3220200321030231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021103010311311-3211022220030213-2102012102113112-1103103100221322-3020200303022302-2032321132212121-0323321101220030-3311031213020030"></a>

## sw — sw / 210320222230 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- sw

<a id="canonical-1220203211332031-1012022012103003-2202213333310310-0003011022001030-2333301133330210-1321202102321312-3302111220323310-0220233021002132"></a>

Type: `"object"`. single nested block, Optional.

Select the F5XC Software Version for the site. By default, latest available F5XC Software Version
will be used. Refer to release notes to find required released SW versions.

Upstream description:

Select the F5XC Software Version for the site. By default, latest available F5XC Software Version
will be used. Refer to release notes to find required released SW versions.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_sw_version",
    "volterra_software_version")}
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
  "x-ves-oneof-field-volterra_sw_version_choice": "[\"default_sw_version\",\"volterra_software_version\"]"
}
```

Terraform syntax:

```terraform
sw {
  # Configure direct properties listed below.
}
```

<a id="canonical-0221323102030120-1001230110011101-3201101010022102-2302213312233001-3103010002012100-3302312311211202-2030300121112112-0131331012311033"></a>

## Direct properties — sw / 210320222230 / 3

- [default_sw_version](resources--aws_tgw_site--reference--group-002.md#canonical-2122101121313223-3322012222120200-3011303013000313-3220223332211102-3223000323012033-1321203110211113-1102203012031232-3202123010202331): complete subsection reference.

<a id="canonical-1321321203111033-1013033230200311-0123231331120201-3033201212120102-1213213211202313-0212000222010303-0123111021302231-1113023221232321"></a>

<a id="canonical-3200303033031210-2031023032221211-2323332123220312-1300113012023333-0201231033321131-1313033321100003-3232010300020303-2330213011202020"></a>

## volterra_software_version property — sw / 210320222230 / 4

Type: `"string"`. Optional.

Exclusive with \[default\_sw\_version\] Specify a F5XC Software Version to be used e.g.
Crt-20210329-1002.

Upstream description:

Exclusive with \[default\_sw\_version\] Specify a F5XC Software Version to be used e.g.
Crt-20210329-1002.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
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
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-0030031223122023-1232123332203222-3201321131123313-2031321233131330-2332222212021223-3321330122303200-2023323010120102-1003133133200310"></a>

## Next pages — sw / 210320222230 / 5

- [sw.default_sw_version](resources--aws_tgw_site--reference--group-002.md#canonical-2122101121313223-3322012222120200-3011303013000313-3220223332211102-3223000323012033-1321203110211113-1102203012031232-3202123010202331)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-2122101121313223-3322012222120200-3011303013000313-3220223332211102-3223000323012033-1321203110211113-1102203012031232-3202123010202331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323222113301033-1203323311130310-2131322321101110-3122112001212301-3011133101033322-1230200202320103-2230121111302322-1220010300311132"></a>

## sw.default_sw_version — default_sw_version / 102111022101 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [sw](resources--aws_tgw_site--reference--group-002.md#canonical-1300131232312032-1320122110210113-3211331031330230-1322221001301333-0300033311131133-2330301300022230-1020032213210011-3220200321030231)
- sw.default_sw_version

<a id="canonical-2331321303310332-1300321330320310-0001030011123100-0231302303203300-1210200211212113-3101132200131013-2113203202302100-0120030010221232"></a>

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
default_sw_version = {}
```

<a id="canonical-3320211000022211-3332232213131001-1030200331110011-3300122001111233-2231200023102103-0313130313231112-1210321213231222-2012002111333321"></a>

## Direct properties — default_sw_version / 102111022101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3003221220013221-3302232001200312-2332220320132202-3033001230333230-2131312322333003-0311022103202003-0300121312220320-1110230121103011"></a>

## Next pages — default_sw_version / 102111022101 / 4

- [sw](resources--aws_tgw_site--reference--group-002.md#canonical-1300131232312032-1320122110210113-3211331031330230-1322221001301333-0300033311131133-2330301300022230-1020032213210011-3220200321030231)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-1133330010300121-0303232012001232-1332311230303120-2211233112111030-1123201320011012-1210113320203131-3032233221100133-3213002021031030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2230121333301331-1321322110202110-0011311220322002-3320122111103103-2101123313132213-1311331110223121-1200020221201302-1212012332133030"></a>

## tgw_security — tgw_security / 333120202103 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- tgw_security

<a id="canonical-1312311100131121-3233120200020121-0002032311133320-3321112003303032-1030013003323100-3023320302103001-2101023301132013-3111100330221013"></a>

Type: `"object"`. single nested block, Optional.

Security Configuration for transit gateway.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("active_east_west_service_policies",
    "east_west_service_policy_allow_all"),
  validators.ConflictingObjectAttributes("active_east_west_service_policies",
    "no_east_west_policy"),
  validators.ConflictingObjectAttributes("active_enhanced_firewall_policies",
    "active_network_policies"),
  validators.ConflictingObjectAttributes("active_enhanced_firewall_policies",
    "no_network_policy"),
  validators.ConflictingObjectAttributes("active_forward_proxy_policies",
    "forward_proxy_allow_all"),
  validators.ConflictingObjectAttributes("active_forward_proxy_policies",
    "no_forward_proxy"),
  validators.ConflictingObjectAttributes("active_network_policies",
    "no_network_policy"),
  validators.ConflictingObjectAttributes("east_west_service_policy_allow_all",
    "no_east_west_policy"),
  validators.ConflictingObjectAttributes("forward_proxy_allow_all",
    "no_forward_proxy")}
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
  "x-ves-oneof-field-east_west_service_policy_choice": "[\"active_east_west_service_policies\",\"east_west_service_policy_allow_all\",\"no_east_west_policy\"]",
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]"
}
```

Terraform syntax:

```terraform
tgw_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-1130131203112203-3103032320123320-0311113133213332-3020021112001221-1213310120010212-3021003001301033-3222121103020313-1330210202200013"></a>

## Direct properties — tgw_security / 333120202103 / 3

- [active_east_west_service_policies](resources--aws_tgw_site--reference--group-002.md#canonical-2233103232021301-3211110120311313-1122222320100212-3011331000331302-3311002321132211-2222000130113213-0321012021212213-3312331031310030): complete subsection reference.

- [active_enhanced_firewall_policies](resources--aws_tgw_site--reference--group-002.md#canonical-2310103112323101-3232032222023122-3013133222200133-0011132033232220-2220103123223101-2110233232233231-3103023133133021-2202121332213133): complete subsection reference.

- [active_forward_proxy_policies](resources--aws_tgw_site--reference--group-002.md#canonical-1031013310020101-2212122023230013-2013203210021030-1131300210302300-3323030102103103-2203030010012100-1133111331201201-0002132330131311): complete subsection reference.

- [active_network_policies](resources--aws_tgw_site--reference--group-002.md#canonical-3313010131101201-0330013203223302-3310032201223232-3120232120111023-1323300221122300-1111001213032333-3033303213023132-0113333122021120): complete subsection reference.

- [east_west_service_policy_allow_all](resources--aws_tgw_site--reference--group-002.md#canonical-1212231303220233-2010113303122033-3011320331223131-0321033320220030-1220333031302100-2013330003103020-3132213201030302-3013202322330332): complete subsection reference.

- [forward_proxy_allow_all](resources--aws_tgw_site--reference--group-002.md#canonical-1201120002231301-0121123222011213-0222322100210132-2331203031331310-1021212323122132-2103312223133222-1200301331133103-3000320022330201): complete subsection reference.

- [no_east_west_policy](resources--aws_tgw_site--reference--group-002.md#canonical-2023022012132131-2202121200230011-1120311221213111-2232023300131020-3021323023111122-3133100110121313-0221121132110210-2010132230010101): complete subsection reference.

- [no_forward_proxy](resources--aws_tgw_site--reference--group-002.md#canonical-2321021113032211-1112301111210302-2122022013230311-2310330302301131-1212323230212112-0001210222132001-2222022202023102-3331013003033310): complete subsection reference.

- [no_network_policy](resources--aws_tgw_site--reference--group-003.md#canonical-1203031221233023-1200122111122211-2020320201211200-1312131320023211-3002302120132300-3203030331100323-3021220313112001-0230023220113030): complete subsection reference.

<a id="canonical-0221231032311231-0210011103230112-1310323223212110-3032210102300310-0220032303012302-2332211320203331-2112302110013030-0022023230010101"></a>

## Next pages — tgw_security / 333120202103 / 4

- [tgw_security.active_east_west_service_policies](resources--aws_tgw_site--reference--group-002.md#canonical-2233103232021301-3211110120311313-1122222320100212-3011331000331302-3311002321132211-2222000130113213-0321012021212213-3312331031310030)
- [tgw_security.active_enhanced_firewall_policies](resources--aws_tgw_site--reference--group-002.md#canonical-2310103112323101-3232032222023122-3013133222200133-0011132033232220-2220103123223101-2110233232233231-3103023133133021-2202121332213133)
- [tgw_security.active_forward_proxy_policies](resources--aws_tgw_site--reference--group-002.md#canonical-1031013310020101-2212122023230013-2013203210021030-1131300210302300-3323030102103103-2203030010012100-1133111331201201-0002132330131311)
- [tgw_security.active_network_policies](resources--aws_tgw_site--reference--group-002.md#canonical-3313010131101201-0330013203223302-3310032201223232-3120232120111023-1323300221122300-1111001213032333-3033303213023132-0113333122021120)
- [tgw_security.east_west_service_policy_allow_all](resources--aws_tgw_site--reference--group-002.md#canonical-1212231303220233-2010113303122033-3011320331223131-0321033320220030-1220333031302100-2013330003103020-3132213201030302-3013202322330332)
- [tgw_security.forward_proxy_allow_all](resources--aws_tgw_site--reference--group-002.md#canonical-1201120002231301-0121123222011213-0222322100210132-2331203031331310-1021212323122132-2103312223133222-1200301331133103-3000320022330201)
- [tgw_security.no_east_west_policy](resources--aws_tgw_site--reference--group-002.md#canonical-2023022012132131-2202121200230011-1120311221213111-2232023300131020-3021323023111122-3133100110121313-0221121132110210-2010132230010101)
- [tgw_security.no_forward_proxy](resources--aws_tgw_site--reference--group-002.md#canonical-2321021113032211-1112301111210302-2122022013230311-2310330302301131-1212323230212112-0001210222132001-2222022202023102-3331013003033310)
- [tgw_security.no_network_policy](resources--aws_tgw_site--reference--group-003.md#canonical-1203031221233023-1200122111122211-2020320201211200-1312131320023211-3002302120132300-3203030331100323-3021220313112001-0230023220113030)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-2233103232021301-3211110120311313-1122222320100212-3011331000331302-3311002321132211-2222000130113213-0321012021212213-3312331031310030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2033023333012323-0012223003001113-3032133300001333-1302031310012310-1131130022011013-2213032330213211-1112331011122232-0330101332033331"></a>

## tgw_security.active_east_west_service_policies — active_east_west_service_policies / 112100110122 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-1133330010300121-0303232012001232-1332311230303120-2211233112111030-1123201320011012-1210113320203131-3032233221100133-3213002021031030)
- tgw_security.active_east_west_service_policies

<a id="canonical-1110123011202323-1301012013302030-3302320031320021-3130231311231300-3300033322110302-1120032122233231-0133230032331311-3103112210102221"></a>

Type: `"object"`. single nested block, Optional.

Active service policies for the east-west proxy.

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
active_east_west_service_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-2320000222213113-0223220030331300-3233212100120311-3332011000303221-3023003201310123-1112030213010130-1331122300322212-0201100320031212"></a>

## Direct properties — active_east_west_service_policies / 112100110122 / 3

- [service_policies](resources--aws_tgw_site--reference--group-002.md#canonical-3031210222331320-1321331113102131-0302223233021222-1112100300300103-1013233221322113-0010220002021301-1121021121101311-2233010303121102): complete subsection reference.

<a id="canonical-3031122212000221-0002122231021213-3013210302212011-1322310202131220-1023233232131110-3020100010221203-3203300012110033-1123100212100220"></a>

## Next pages — active_east_west_service_policies / 112100110122 / 4

- [tgw_security.active_east_west_service_policies.service_policies](resources--aws_tgw_site--reference--group-002.md#canonical-3031210222331320-1321331113102131-0302223233021222-1112100300300103-1013233221322113-0010220002021301-1121021121101311-2233010303121102)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-1133330010300121-0303232012001232-1332311230303120-2211233112111030-1123201320011012-1210113320203131-3032233221100133-3213002021031030)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-3031210222331320-1321331113102131-0302223233021222-1112100300300103-1013233221322113-0010220002021301-1121021121101311-2233010303121102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0022201220311121-0321012200330312-0000212313132111-1113002313030122-2131102031132112-1303201121022231-0310320331023220-0112203313121033"></a>

## tgw_security.active_east_west_service_policies.service_policies — service_policies / 301231312210 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-1133330010300121-0303232012001232-1332311230303120-2211233112111030-1123201320011012-1210113320203131-3032233221100133-3213002021031030)
- [tgw_security.active_east_west_service_policies](resources--aws_tgw_site--reference--group-002.md#canonical-2233103232021301-3211110120311313-1122222320100212-3011331000331302-3311002321132211-2222000130113213-0321012021212213-3312331031310030)
- tgw_security.active_east_west_service_policies.service_policies

<a id="canonical-0032013332330212-1223121033100231-0211013010300312-3200310103111020-3330330210001323-3130022301111011-2131112133212203-2003323231002333"></a>

Type: `"object"`. list nested block, Optional.

List of references to service\_policy objects.

Upstream description:

A list of references to service\_policy objects.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
service_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-1031322230200033-1300230311232313-2100333300212112-1331332132110321-0213001313202332-1123202031200101-0321033110130302-0112220003022213"></a>

## Direct properties — service_policies / 301231312210 / 3

<a id="canonical-0132022012210012-0302103130320323-0112320303231331-0033212001222100-0032212010213300-0223332112130201-0121203032310123-3002310321102113"></a>

<a id="canonical-3230333122230032-0201101123220003-1321031133133133-3113021033201010-1223320022231100-2131112202321031-3123323102320132-2323001200222022"></a>

## name property — service_policies / 301231312210 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-2323302332212223-3030033003231002-3111111213220311-3323122230102022-3023222010320221-2132001313322021-0111323133321322-1122130323333131"></a>

<a id="canonical-0113320201000312-1330122002300010-0021031320013231-0310321321322312-0102331102021102-1130111022101123-0221113133031233-2331030012131121"></a>

## namespace property — service_policies / 301231312210 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1203310231210323-1202132101011333-3103232112031203-1123230031333202-2302211020232013-3111332330030203-1231011322112122-1002220032102320"></a>

<a id="canonical-3023301330200111-0120100020321130-0203010312010121-2031232111300332-3200013331012300-1111030310323002-2212231333132200-2310113311001021"></a>

## tenant property — service_policies / 301231312210 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0230232031230101-0103200211010230-3111132312123223-1021013310103033-0233202100103002-2003133002020330-3303000202131031-2101233102231123"></a>

## Next pages — service_policies / 301231312210 / 7

- [tgw_security.active_east_west_service_policies](resources--aws_tgw_site--reference--group-002.md#canonical-2233103232021301-3211110120311313-1122222320100212-3011331000331302-3311002321132211-2222000130113213-0321012021212213-3312331031310030)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-2310103112323101-3232032222023122-3013133222200133-0011132033232220-2220103123223101-2110233232233231-3103023133133021-2202121332213133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203221133023030-2131032102232001-0322331303121111-3201302101131321-2220021030330121-2313102222212012-2221031132013211-0111330232133132"></a>

## tgw_security.active_enhanced_firewall_policies — active_enhanced_firewall_policies / 312201320310 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-1133330010300121-0303232012001232-1332311230303120-2211233112111030-1123201320011012-1210113320203131-3032233221100133-3213002021031030)
- tgw_security.active_enhanced_firewall_policies

<a id="canonical-0223021110012232-2102232232321313-1300331100001221-2121131131222332-0101133022121232-3133330011203213-0101312232330033-2031021121112033"></a>

Type: `"object"`. single nested block, Optional.

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Upstream description:

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("enhanced_firewall_policies")}
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
active_enhanced_firewall_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-1013300332133011-0303100001030201-3112030113302223-0011331220112331-0202213012133020-2120111332210200-0033311312103002-1313222112230333"></a>

## Direct properties — active_enhanced_firewall_policies / 312201320310 / 3

- [enhanced_firewall_policies](resources--aws_tgw_site--reference--group-002.md#canonical-2331222203321201-2200122323322103-0123332223213331-2110100102323123-2220210323232113-2222232312221210-2102122211121032-3112031123130232): complete subsection reference.

<a id="canonical-2123003012312203-2123011122102123-0003103032111313-0032003203333302-2101133003203320-1323321030132131-3213233111223332-3112210022233320"></a>

## Next pages — active_enhanced_firewall_policies / 312201320310 / 4

- [tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies](resources--aws_tgw_site--reference--group-002.md#canonical-2331222203321201-2200122323322103-0123332223213331-2110100102323123-2220210323232113-2222232312221210-2102122211121032-3112031123130232)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-1133330010300121-0303232012001232-1332311230303120-2211233112111030-1123201320011012-1210113320203131-3032233221100133-3213002021031030)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-2331222203321201-2200122323322103-0123332223213331-2110100102323123-2220210323232113-2222232312221210-2102122211121032-3112031123130232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2100330023303200-0320020120310010-0030131300220210-2302323103131300-1210221230221313-1130120033220203-1331013001231030-2112312101131320"></a>

## tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies — enhanced_firewall_policies / 301303332211 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-1133330010300121-0303232012001232-1332311230303120-2211233112111030-1123201320011012-1210113320203131-3032233221100133-3213002021031030)
- [tgw_security.active_enhanced_firewall_policies](resources--aws_tgw_site--reference--group-002.md#canonical-2310103112323101-3232032222023122-3013133222200133-0011132033232220-2220103123223101-2110233232233231-3103023133133021-2202121332213133)
- tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies

<a id="canonical-0133031213232322-2011100200011223-1013312231310002-1000222322200300-0003301312122011-2310201300213302-0232033102311022-2321113013221021"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Enhanced Firewall Policies active.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
enhanced_firewall_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-3011323120303310-3021131221110002-3331111112223232-0113312231022332-2113232002313300-2332103211330221-0331031203113222-0330102121203313"></a>

## Direct properties — enhanced_firewall_policies / 301303332211 / 3

<a id="canonical-0002331331222220-3001220211102132-3123322023301321-0313002120232033-0131212313122033-0122112131201013-3332330000203003-1103220133003110"></a>

<a id="canonical-1330300021223321-2210031001101030-0212030332313122-2202203330023333-0312123220221011-3230323031210203-3203020121313333-0022002011111113"></a>

## name property — enhanced_firewall_policies / 301303332211 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-1320231312112002-3233203000120332-0120113030321030-0033213100013200-2133320232132132-3302313000202220-2031133313211033-0332320321032303"></a>

<a id="canonical-3002111120232123-0011000323131100-2033120330002011-3133113230031013-0320122201122022-3210102222123032-0230132132230132-2200000020102333"></a>

## namespace property — enhanced_firewall_policies / 301303332211 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3330302211231010-1122033211133123-2313113302203230-3113310221323120-2030332003002113-0010320011300010-1000131022333123-2131221330313230"></a>

<a id="canonical-1202203200210020-1212223101300102-2021111231033303-1321102002021212-3312221001123112-0132211330121013-3302223322131211-3331330332033211"></a>

## tenant property — enhanced_firewall_policies / 301303332211 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3121310221200321-0203211213211113-1212102222030020-0332021330031101-3330130332223301-2221001023321003-1300132213120122-1212322002222202"></a>

## Next pages — enhanced_firewall_policies / 301303332211 / 7

- [tgw_security.active_enhanced_firewall_policies](resources--aws_tgw_site--reference--group-002.md#canonical-2310103112323101-3232032222023122-3013133222200133-0011132033232220-2220103123223101-2110233232233231-3103023133133021-2202121332213133)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-1031013310020101-2212122023230013-2013203210021030-1131300210302300-3323030102103103-2203030010012100-1133111331201201-0002132330131311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323201320211130-0230022212033220-3000213303131110-2322302113031020-2203132003200321-1323330332103022-1000223132023023-1331003331330133"></a>

## tgw_security.active_forward_proxy_policies — active_forward_proxy_policies / 120120233123 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-1133330010300121-0303232012001232-1332311230303120-2211233112111030-1123201320011012-1210113320203131-3032233221100133-3213002021031030)
- tgw_security.active_forward_proxy_policies

<a id="canonical-0033133103231112-1322011022002333-1101200110203123-2211112332330230-2033111032003301-1032010021000210-3120103020011233-1312333132003201"></a>

Type: `"object"`. single nested block, Optional.

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("forward_proxy_policies")}
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
active_forward_proxy_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-1132132332131121-2130032023220133-1212331010121033-1120332321100222-1221301121201321-3231332303333100-2302330232120223-0102110301130010"></a>

## Direct properties — active_forward_proxy_policies / 120120233123 / 3

- [forward_proxy_policies](resources--aws_tgw_site--reference--group-002.md#canonical-1111110110312031-3311213033030211-0211323330230313-3211200330202220-0112122010221201-0123012131313101-3103301111201130-1130200312313302): complete subsection reference.

<a id="canonical-3131221120120012-1023012110130101-1020300330120210-2231000312102132-2313230233303011-1111012101202220-2003321320230203-3021233002302213"></a>

## Next pages — active_forward_proxy_policies / 120120233123 / 4

- [tgw_security.active_forward_proxy_policies.forward_proxy_policies](resources--aws_tgw_site--reference--group-002.md#canonical-1111110110312031-3311213033030211-0211323330230313-3211200330202220-0112122010221201-0123012131313101-3103301111201130-1130200312313302)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-1133330010300121-0303232012001232-1332311230303120-2211233112111030-1123201320011012-1210113320203131-3032233221100133-3213002021031030)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-1111110110312031-3311213033030211-0211323330230313-3211200330202220-0112122010221201-0123012131313101-3103301111201130-1130200312313302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2021020302310220-2033222231123303-2220200313331330-2110121103332203-2011230020101011-2301230112030003-2010000121020301-0011321330120302"></a>

## tgw_security.active_forward_proxy_policies.forward_proxy_policies — forward_proxy_policies / 311202002333 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-1133330010300121-0303232012001232-1332311230303120-2211233112111030-1123201320011012-1210113320203131-3032233221100133-3213002021031030)
- [tgw_security.active_forward_proxy_policies](resources--aws_tgw_site--reference--group-002.md#canonical-1031013310020101-2212122023230013-2013203210021030-1131300210302300-3323030102103103-2203030010012100-1133111331201201-0002132330131311)
- tgw_security.active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-0120012202100031-1003213003133310-0010210302111203-2322111200130301-3101221011013332-1330111233102221-1110100230121323-0022103331111122"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
forward_proxy_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-0003201213010301-0212103331231123-2123130320313222-3100111112033310-1220102111132131-1022013232112313-2002312232323231-3330300132300322"></a>

## Direct properties — forward_proxy_policies / 311202002333 / 3

<a id="canonical-3320012022220330-2301110030300010-1010001121021323-2203031022222212-3032101222133130-1210020310230212-0222030200310211-0010301131321221"></a>

<a id="canonical-2023311102221231-3002013212210312-1012010332121220-2303002312003002-1132201010311310-0312022121032103-1033111321131023-3100212320310022"></a>

## name property — forward_proxy_policies / 311202002333 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-2001130010322213-1121312320023202-0213213032323233-1222113322301333-2111020000232333-0313311213231333-3033220000002221-3010101312111130"></a>

<a id="canonical-1310203211222030-3330323021212300-0013123323012000-1122103120120200-3002300021222002-1122020312021001-0132311103133233-1231110101112102"></a>

## namespace property — forward_proxy_policies / 311202002333 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2330230101002121-0102021221322121-0313131233330030-2001221111323223-2001021222112020-0323212320312022-0222120303123231-0122011012311101"></a>

<a id="canonical-2020233123032321-0231130211101120-0313013132222213-2231121003331132-2331301101023303-3030012323011310-0312313313232312-2331311230003002"></a>

## tenant property — forward_proxy_policies / 311202002333 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2310210323231213-3233202013031103-0120113111210233-3222121313203303-1312001003221120-3012132030102223-3011000333020311-3231021130032132"></a>

## Next pages — forward_proxy_policies / 311202002333 / 7

- [tgw_security.active_forward_proxy_policies](resources--aws_tgw_site--reference--group-002.md#canonical-1031013310020101-2212122023230013-2013203210021030-1131300210302300-3323030102103103-2203030010012100-1133111331201201-0002132330131311)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-3313010131101201-0330013203223302-3310032201223232-3120232120111023-1323300221122300-1111001213032333-3033303213023132-0113333122021120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120200123021003-0222130113030311-3211020100113121-1002331131210312-1133022212032011-0222132303112331-1021202120213023-1333120022310211"></a>

## tgw_security.active_network_policies — active_network_policies / 001031333313 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-1133330010300121-0303232012001232-1332311230303120-2211233112111030-1123201320011012-1210113320203131-3032233221100133-3213002021031030)
- tgw_security.active_network_policies

<a id="canonical-1013300123232013-1133231232230201-1233133330301023-2223323200222110-3313312302312022-2032021311033220-1213111133330030-1003213120311330"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for active network policies.

Upstream description:

List of firewall policy views.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("network_policies")}
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
active_network_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-0012001202301231-1103010132221131-1003030023000100-1210210020022222-2333330100012300-1013333100322130-0311010110033030-2321012110311220"></a>

## Direct properties — active_network_policies / 001031333313 / 3

- [network_policies](resources--aws_tgw_site--reference--group-002.md#canonical-2303210302020230-0330032302331132-3033322230220123-3331232211012013-3201201131301032-3223333130233302-3210330021122201-0013111113323011): complete subsection reference.

<a id="canonical-3133003302112320-1332022233030301-2101331312102021-2123320210023120-3131020121211323-3032202212012031-0200112302133000-3120131211022321"></a>

## Next pages — active_network_policies / 001031333313 / 4

- [tgw_security.active_network_policies.network_policies](resources--aws_tgw_site--reference--group-002.md#canonical-2303210302020230-0330032302331132-3033322230220123-3331232211012013-3201201131301032-3223333130233302-3210330021122201-0013111113323011)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-1133330010300121-0303232012001232-1332311230303120-2211233112111030-1123201320011012-1210113320203131-3032233221100133-3213002021031030)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-2303210302020230-0330032302331132-3033322230220123-3331232211012013-3201201131301032-3223333130233302-3210330021122201-0013111113323011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231321012022330-2121313100102211-0222021210120321-3321203003201210-0101101100121331-0101230023010001-0331212132012201-2231101023313231"></a>

## tgw_security.active_network_policies.network_policies — network_policies / 030210203213 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-1133330010300121-0303232012001232-1332311230303120-2211233112111030-1123201320011012-1210113320203131-3032233221100133-3213002021031030)
- [tgw_security.active_network_policies](resources--aws_tgw_site--reference--group-002.md#canonical-3313010131101201-0330013203223302-3310032201223232-3120232120111023-1323300221122300-1111001213032333-3033303213023132-0113333122021120)
- tgw_security.active_network_policies.network_policies

<a id="canonical-3322031111021233-2111122223311232-1123330223011130-1113203303012103-0101212220101211-1310231030032023-0030213132100231-3030310320013233"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Firewall Policies active for this network firewall.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
network_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-3223021122232133-3113211110121332-3331303303003212-2302300110021312-2123322130032332-2002323321332311-0302312201222220-3330200312013011"></a>

## Direct properties — network_policies / 030210203213 / 3

<a id="canonical-2001010001020323-1231020032133321-1022223202113001-1222220202031303-2331032202300020-3032031333202302-3313111013102201-1013301303130000"></a>

<a id="canonical-3001121232122313-2102002221011232-2023120032211231-0213313232322213-1102313031122331-3211230021233202-0113121101002101-3123112231010333"></a>

## name property — network_policies / 030210203213 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-3231103000203331-1032023231010333-1221201311023113-0232101103212323-2230311131320313-0121222330212100-0222223220231330-1121203023213010"></a>

<a id="canonical-0203210231210123-1130313300203220-1301322001200332-3333012203311120-0323111210102303-1223112332130100-1131302211210323-0032022311120011"></a>

## namespace property — network_policies / 030210203213 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0023021330101130-3110202032110331-3321013023003222-3012322110130030-3212033111010211-0032032303213112-3023330021223122-0033030101312023"></a>

<a id="canonical-1330303102132001-1013202211113302-3011313110331320-1232312320203302-1200113331102322-3310100211202330-3232220201102123-0030002312120212"></a>

## tenant property — network_policies / 030210203213 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3033101013210131-2201202002112022-2102231233310311-2310012000233102-1221301132331132-3000312221230322-1230200102202333-3333110103131212"></a>

## Next pages — network_policies / 030210203213 / 7

- [tgw_security.active_network_policies](resources--aws_tgw_site--reference--group-002.md#canonical-3313010131101201-0330013203223302-3310032201223232-3120232120111023-1323300221122300-1111001213032333-3033303213023132-0113333122021120)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-1212231303220233-2010113303122033-3011320331223131-0321033320220030-1220333031302100-2013330003103020-3132213201030302-3013202322330332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3103010320011013-1231201122213311-0011030103021021-1233212233110133-0103300031303201-2230132200003202-1213221220300233-0103100320101011"></a>

## tgw_security.east_west_service_policy_allow_all — east_west_service_policy_allow_all / 100220120022 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-1133330010300121-0303232012001232-1332311230303120-2211233112111030-1123201320011012-1210113320203131-3032233221100133-3213002021031030)
- tgw_security.east_west_service_policy_allow_all

<a id="canonical-2230001032112130-1221022223012331-1201031222002003-3213100210232013-3121010322112012-2203012101131333-1022333201302303-0001223010310023"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for east west service policy allow all.

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
east_west_service_policy_allow_all = {}
```

<a id="canonical-3120331122113103-1031303302311132-1230331300330213-3213222012213221-2021313101130022-3110223032203332-1013310001110311-3133230212132111"></a>

## Direct properties — east_west_service_policy_allow_all / 100220120022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1210030110233201-1231201223010222-3213332010312311-2233013233132002-3111101020121121-3300012211131211-0301122333002212-1100110011011310"></a>

## Next pages — east_west_service_policy_allow_all / 100220120022 / 4

- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-1133330010300121-0303232012001232-1332311230303120-2211233112111030-1123201320011012-1210113320203131-3032233221100133-3213002021031030)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-1201120002231301-0121123222011213-0222322100210132-2331203031331310-1021212323122132-2103312223133222-1200301331133103-3000320022330201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123223300333022-1020133100003211-0002330330310120-3011230313100013-3110230222300302-1033000210303211-3032101023331031-0333103011310220"></a>

## tgw_security.forward_proxy_allow_all — forward_proxy_allow_all / 102330203233 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-1133330010300121-0303232012001232-1332311230303120-2211233112111030-1123201320011012-1210113320203131-3032233221100133-3213002021031030)
- tgw_security.forward_proxy_allow_all

<a id="canonical-2333003131113210-2222122020212330-3010313002130002-0013111320230030-0110201233302222-0313233211302130-1300131313312001-0332311230012302"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for forward proxy allow all.

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
forward_proxy_allow_all = {}
```

<a id="canonical-2321133031120021-0223131232122100-1000013113332133-0102211022303022-1302222120330210-0330002120233020-1113020013310010-2031011200201100"></a>

## Direct properties — forward_proxy_allow_all / 102330203233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0001230030212231-2000200332230333-3202031030301012-3033130210330300-0330203012013331-1022303001022031-2303103221021333-2022230100110030"></a>

## Next pages — forward_proxy_allow_all / 102330203233 / 4

- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-1133330010300121-0303232012001232-1332311230303120-2211233112111030-1123201320011012-1210113320203131-3032233221100133-3213002021031030)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-2023022012132131-2202121200230011-1120311221213111-2232023300131020-3021323023111122-3133100110121313-0221121132110210-2010132230010101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033213102311023-0102011123033101-2100210020210033-0233222321302212-2122321001321133-0200300311110213-1110322102310211-3102232210300200"></a>

## tgw_security.no_east_west_policy — no_east_west_policy / 013232001322 / 2

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-1133330010300121-0303232012001232-1332311230303120-2211233112111030-1123201320011012-1210113320203131-3032233221100133-3213002021031030)
- tgw_security.no_east_west_policy

<a id="canonical-0323023213030011-0012122230003031-3012313112310022-0233333100101312-2100201110023311-3003213002211131-3322021222131210-0123102200133322"></a>

Type: `["object", {}]`. Optional.

Policy configuration for this feature.

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
no_east_west_policy = {}
```

<a id="canonical-3013312231001102-2133322022021330-1011212202311203-2102000001221233-0120220023113320-3032320120102133-2203030101310013-2312011200130233"></a>

## Direct properties — no_east_west_policy / 013232001322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1223121130233233-3211211311313230-0320322201300312-0133123301222231-1333112213332002-2311201221231322-0021111121011331-2223303031021100"></a>

## Next pages — no_east_west_policy / 013232001322 / 4

- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-1133330010300121-0303232012001232-1332311230303120-2211233112111030-1123201320011012-1210113320203131-3032233221100133-3213002021031030)
- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)

<a id="canonical-2321021113032211-1112301111210302-2122022013230311-2310330302301131-1212323230212112-0001210222132001-2222022202023102-3331013003033310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
