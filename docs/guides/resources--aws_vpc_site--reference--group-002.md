---
page_title: "xcsh_aws_vpc_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_aws_vpc_site reference."
---

# xcsh_aws_vpc_site reference

<a id="canonical-3121313111230200-2320033110320033-2102332223122100-0330231311203321-0011311232331231-1001303121301113-1201023100223201-0123121312212012"></a>

## blocked_services.blocked_service.web_user_interface — web_user_interface / 132012210121 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [blocked_services](resources--aws_vpc_site--reference--group-001.md#canonical-3323032222210001-3302120102203102-3210021101231303-3023332013331112-3200032312323130-3223301102110203-3100201032112113-1310123003122330)
- [blocked_services.blocked_service](resources--aws_vpc_site--reference--group-001.md#canonical-0102322030032221-1312020110103322-0312312320331312-2323203222312233-3130330322220220-3322320221222010-0311102231203132-2213033023012032)
- blocked_services.blocked_service.web_user_interface

<a id="canonical-2002121123302203-2022311031003130-2333123222102332-3212030132122221-1330122220112320-3112333332130321-3212300123122031-2221320123111211"></a>

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

<a id="canonical-3321221223103111-2201333303113012-3302010121300310-1100220300210111-1100332222331332-3001001101112112-0130030030300221-1313222301123313"></a>

## Direct properties — web_user_interface / 132012210121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2102123312131023-0002021212213120-1131131222211302-1220030033202113-0110023333312000-1131203122231031-0332033313131222-2120011030230030"></a>

## Next pages — web_user_interface / 132012210121 / 4

- [blocked_services.blocked_service](resources--aws_vpc_site--reference--group-001.md#canonical-0102322030032221-1312020110103322-0312312320331312-2323203222312233-3130330322220220-3322320221222010-0311102231203132-2213033023012032)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3103301232022012-1202210120113233-3231000231313300-1032103030032313-3210112321011201-2112332333203011-1330133221022222-3221223133312113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133100301011021-1223321130333011-2200023002222330-2313122030221213-3200131120132323-0010300223211211-2231220102313230-2031112100303330"></a>

## coordinates — coordinates / 031130321203 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- coordinates

<a id="canonical-1021323130312312-0303132312300222-1201203202331331-2212211200002230-0032110323202321-3223123021132120-1023332000333220-0212132333123010"></a>

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

<a id="canonical-0101212133333323-0321303120000123-2201013103310302-3331102323003033-1232303133200233-2221013311002012-1003223023133102-2011013220120120"></a>

## Direct properties — coordinates / 031130321203 / 3

<a id="canonical-2201333302232201-0130202210333331-1303012210232311-0111233330113021-2322003013200320-2332123302011311-1013033110222331-3213201113131032"></a>

<a id="canonical-2121230131011011-3300122112231210-2211112231232122-1023132331323222-2320131022011120-1003003130303300-1303230203011010-1211021310132132"></a>

## latitude property — coordinates / 031130321203 / 4

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

<a id="canonical-1112330333201332-1032112333323302-1102300111220221-3123031033120011-3313110102331321-1331303320322311-0213232012012320-2320132212111011"></a>

<a id="canonical-3021332323102111-0032312121232021-1022131133101310-1123233320233310-3230220123213300-3012220032000212-3122301032320111-2011121331010223"></a>

## longitude property — coordinates / 031130321203 / 5

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

<a id="canonical-0332221203313332-3000030302310210-3020003131232032-2020231202231023-1000002111021301-2222130313120030-2313230310320311-0300310020230211"></a>

## Next pages — coordinates / 031130321203 / 6

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0021301331020010-2103113020103011-0313112212033023-3001103200122333-1101111311000032-3120100021023322-0133210203122003-3032032330221333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003113231200322-0111230101031203-3123100333311232-3031221221211303-2101203203303130-0233112211000322-0001012333121002-0002311221202101"></a>

## custom_dns — custom_dns / 131321322202 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- custom_dns

<a id="canonical-3013213332301333-0002013200023233-3110132301123323-2133202030131233-2130113111122330-0332112100001230-1301220122302320-1321333130313223"></a>

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

<a id="canonical-0031031312023130-2222300221123322-0321112011213001-1123312011200211-2023001320012123-1023323001113210-3102113310111332-3222003100313230"></a>

## Direct properties — custom_dns / 131321322202 / 3

<a id="canonical-1322201011331112-2033222120120211-2331122210111222-1100332212010100-1303332013131010-1213102001130120-3003030313003122-1130000112333233"></a>

<a id="canonical-3120322000131130-0020233213333330-1223032023031012-2201012113320013-1000233100130100-3232131122332102-1013123112122121-1311202331102311"></a>

## inside_nameserver property — custom_dns / 131321322202 / 4

Type: `"string"`. Optional.

Optional DNS server IP to be used for name resolution in inside network.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-1332000310230002-1023100102120310-3001310223313103-2113311020132102-2201000300100112-3133223220220232-0321223130012033-0320023102000201"></a>

<a id="canonical-3003312233021213-1232320123020020-0032322301121330-2220013112232003-1333033103333332-1303211032312223-0201233021112222-2022220012122132"></a>

## outside_nameserver property — custom_dns / 131321322202 / 5

Type: `"string"`. Optional.

Optional DNS server IP to be used for name resolution in outside network.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-1220023121210001-3321030232113020-2101033003312323-1110030021032120-2021201101230202-3233010333212313-1030103033100201-1010303131202021"></a>

## Next pages — custom_dns / 131321322202 / 6

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1123132203210311-2301200320231021-3232023100201023-0221020320203100-1321322223023130-1310113030201022-1201331103332123-0222223131322032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223321333121132-1131213330010210-1333221230112112-3033013303322120-1101120120022112-2221300203303121-2130320102012120-2120033311210021"></a>

## custom_security_group — custom_security_group / 112011132312 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- custom_security_group

<a id="canonical-2102131022332001-1020300130220322-2312211301003113-0121002312130122-3213002220113120-3310120221011300-0001100321110230-3323330010321011"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: custom\_security\_group, f5xc\_security\_group\] Enter pre created security groups for
slo(Site Local Outside) and sli(Site Local Inside) interface. Supported only for sites deployed on
existing VPC.

Upstream description:

Enter pre created security groups for slo(Site Local Outside) and sli(Site Local Inside) interface.
Supported only for sites deployed on existing VPC.

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

- [custom_security_group](resources--aws_vpc_site--reference--group-002.md#canonical-2102131022332001-1020300130220322-2312211301003113-0121002312130122-3213002220113120-3310120221011300-0001100321110230-3323330010321011)
- [f5xc_security_group](resources--aws_vpc_site--reference--group-002.md#canonical-3111232222322323-0103012011313123-2133002030330003-2311220201201222-3222332113322223-3021300030333002-3002021233321301-2220013302333233)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
custom_security_group {
  # Configure direct properties listed below.
}
```

<a id="canonical-2322131102230230-3003220200122023-0312103033131310-2002322221232332-0200013033023332-2320023300321230-2200002332321333-1101132132322211"></a>

## Direct properties — custom_security_group / 112011132312 / 3

<a id="canonical-3002101003133330-2002321230103233-2330313223120231-2012001312233020-0230202032003113-2132323103212301-1121213300103103-3310032131001313"></a>

<a id="canonical-0021133021031232-1020201221212302-0300200030122132-3102021022000012-1013111131323010-2103000011123201-1232310003133330-0220030222203222"></a>

## inside_security_group_id property — custom_security_group / 112011132312 / 4

Type: `"string"`. Optional.

Security Group ID to be attached to SLI(Site Local Inside) Interface.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  }
}
```

<a id="canonical-1211003022030111-2130100102100230-2022123001332111-1000300010112113-0020033113120233-1322220030113112-1021021112210120-1102030001131323"></a>

<a id="canonical-2202110110300010-2110113113301222-3331003010221303-0110213132320301-1103332030012130-3201020130122330-1003131332113221-3013123332322002"></a>

## outside_security_group_id property — custom_security_group / 112011132312 / 5

Type: `"string"`. Optional.

Security Group ID to be attached to SLO(Site Local Outside) Interface.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20",
    "ves.io.schema.rules.string.pattern": "^(sg-)([a-z0-9]{8}|[a-z0-9]{17})$|^$"
  }
}
```

<a id="canonical-3232303202023020-2302131132233333-3301032200332330-0011132112303131-2133230102223221-0130300221120130-3203011003102213-1112232123000100"></a>

## Next pages — custom_security_group / 112011132312 / 6

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1023131220002132-1001131221323100-0123021223123002-2100211102320210-1312301332031313-1133210130030320-0102012300311331-0022310331100331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111201023211030-0132202201331011-0223100332031330-2120120133022231-1102220301320322-0012323112332011-1332132120021110-2210103310022203"></a>

## default_blocked_services — default_blocked_services / 320320020132 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- default_blocked_services

<a id="canonical-0100133120312222-0222002310232213-0233220031220330-1112322001310302-3112301002102313-2311111022202021-3000323303112021-0022130113232302"></a>

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

<a id="canonical-3330232100022101-1221303303221232-2032323231223221-0103003110123231-2221023121033203-0023103231333012-0131012101213203-1121232221123302"></a>

## Direct properties — default_blocked_services / 320320020132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3030021212200020-1031203032333200-1003320112011321-0110303012101032-3002222323131013-1101210103110130-0232313210033223-3002133230012310"></a>

## Next pages — default_blocked_services / 320320020132 / 4

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1203233103200121-2313321031323013-0001112100021212-1302133112121133-1031101231223200-2232021210002031-1311100030200230-3200132220200320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3022011101210233-3102011031111020-0122001123100112-1302032030101233-2332212301333122-0031013012213010-0010033210203302-0130112121330233"></a>

## direct_connect_disabled — direct_connect_disabled / 013230210012 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- direct_connect_disabled

<a id="canonical-3320101330223301-0133313121033102-0021020033131303-1310100231330013-0322313322230132-1300222213333202-1223300002202120-3120211300333012"></a>

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

- [direct_connect_disabled](resources--aws_vpc_site--reference--group-002.md#canonical-3320101330223301-0133313121033102-0021020033131303-1310100231330013-0322313322230132-1300222213333202-1223300002202120-3120211300333012)
- [direct_connect_enabled](resources--aws_vpc_site--reference--group-002.md#canonical-3203330010110310-3212301011331311-0013233121301300-2022013203202310-1021123100233010-1001121303020030-2302020110113302-0110010113332001)
- [private_connectivity](resources--aws_vpc_site--reference--group-004.md#canonical-3302002132231203-3030202210010120-1310320012103133-1302011000212230-3111302233111200-0123000123331023-3012122210102203-0012021311213222)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
direct_connect_disabled = {}
```

<a id="canonical-0131220331123212-1312331132333123-1113330203120111-2122313113111103-2120022023101121-0002033131322111-0230001212221330-2012012200030322"></a>

## Direct properties — direct_connect_disabled / 013230210012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1201023221200110-2101003132303120-1301300000110223-2120112331323311-2123012313212123-0220031220323223-2331110123102033-0221031001333313"></a>

## Next pages — direct_connect_disabled / 013230210012 / 4

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0130012221000232-1311333022303301-0020012113300213-0322011110120320-0313333031011211-1310113213303030-0201300310003031-3011331001021112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311202210132321-3020231020101203-1201303303012000-3310011011011020-0332031311021123-3211201213221222-0002112020101301-2231102003300000"></a>

## direct_connect_enabled — direct_connect_enabled / 313131000130 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- direct_connect_enabled

<a id="canonical-3203330010110310-3212301011331311-0013233121301300-2022013203202310-1021123100233010-1001121303020030-2302020110113302-0110010113332001"></a>

Type: `"object"`. single nested block, Optional.

Direct Connect Configuration. Direct Connect Configuration.

Upstream description:

Direct Connect Configuration.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1312333003100003-0313232121001220-1103302311320033-2310103203321210-1221313312110320-3002220001201112-1112212232213200-2230232313103101"></a>

## Direct properties — direct_connect_enabled / 313131000130 / 3

- [auto_asn](resources--aws_vpc_site--reference--group-002.md#canonical-3301032211332211-2332111122333233-0322110230233321-2012131131230323-3232102020110011-2233301103203101-1233031120211222-3221021100002200): complete subsection reference.

<a id="canonical-0010313113130101-1211300303113110-3220230211101013-0223231111332230-2230220321232032-1230013320110032-3332323301302212-1033121111121120"></a>

<a id="canonical-0022003012223003-0310003230111130-2203011313210111-1102220312000202-3201233323110203-2211122111221333-0213032012021130-2113313313130201"></a>

## custom_asn property — direct_connect_enabled / 313131000130 / 4

Type: `"number"`. Optional.

Exclusive with \[auto\_asn\] Custom Autonomous System Number.

Upstream description:

Exclusive with \[auto\_asn\] Custom Autonomous System Number.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [hosted_vifs](resources--aws_vpc_site--reference--group-002.md#canonical-1122123323310331-2133300100123223-2100211323201032-2222233011220302-1300003201112031-0000221033021113-3213222101010022-0121311322002320): complete subsection reference.

- [standard_vifs](resources--aws_vpc_site--reference--group-002.md#canonical-3220022212033112-0320203202221210-0131000010330021-1213200323030110-3310310031011301-0323210131313230-3321022313332331-3303231113000022): complete subsection reference.

<a id="canonical-2133333001113013-0113003000311000-2303210313110301-0102003220020002-0001300032123301-0200212032001232-0012000010202123-3301120322202320"></a>

## Next pages — direct_connect_enabled / 313131000130 / 5

- [direct_connect_enabled.auto_asn](resources--aws_vpc_site--reference--group-002.md#canonical-3301032211332211-2332111122333233-0322110230233321-2012131131230323-3232102020110011-2233301103203101-1233031120211222-3221021100002200)
- [direct_connect_enabled.hosted_vifs](resources--aws_vpc_site--reference--group-002.md#canonical-1122123323310331-2133300100123223-2100211323201032-2222233011220302-1300003201112031-0000221033021113-3213222101010022-0121311322002320)
- [direct_connect_enabled.standard_vifs](resources--aws_vpc_site--reference--group-002.md#canonical-3220022212033112-0320203202221210-0131000010330021-1213200323030110-3310310031011301-0323210131313230-3321022313332331-3303231113000022)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3301032211332211-2332111122333233-0322110230233321-2012131131230323-3232102020110011-2233301103203101-1233031120211222-3221021100002200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3020113333210030-3231333312121202-2112211032013111-0221123322111313-1010033000333232-1221020122123132-3333332301203212-2230211120133130"></a>

## direct_connect_enabled.auto_asn — auto_asn / 312212011121 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [direct_connect_enabled](resources--aws_vpc_site--reference--group-002.md#canonical-0130012221000232-1311333022303301-0020012113300213-0322011110120320-0313333031011211-1310113213303030-0201300310003031-3011331001021112)
- direct_connect_enabled.auto_asn

<a id="canonical-3331131303110202-3321011330320200-0210131122230211-0202201003310001-0321100133213100-2232031302130210-1112030321333122-1311203022130311"></a>

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

<a id="canonical-2221213130123232-3222230023232203-0102313001330020-0032123301010100-2022312202020230-0000033231203031-1231132300301112-1032002233331023"></a>

## Direct properties — auto_asn / 312212011121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3232320013111132-3033201031001023-3200100111123220-2022232122132120-2333101200220113-2031203122201003-0233302330211013-3031312320100310"></a>

## Next pages — auto_asn / 312212011121 / 4

- [direct_connect_enabled](resources--aws_vpc_site--reference--group-002.md#canonical-0130012221000232-1311333022303301-0020012113300213-0322011110120320-0313333031011211-1310113213303030-0201300310003031-3011331001021112)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1122123323310331-2133300100123223-2100211323201032-2222233011220302-1300003201112031-0000221033021113-3213222101010022-0121311322002320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111320322011030-1231232012302332-1230020303230110-2320021203113110-3221221101132333-0102200312201023-1310023032122000-2320321133100201"></a>

## direct_connect_enabled.hosted_vifs — hosted_vifs / 321303011030 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [direct_connect_enabled](resources--aws_vpc_site--reference--group-002.md#canonical-0130012221000232-1311333022303301-0020012113300213-0322011110120320-0313333031011211-1310113213303030-0201300310003031-3011331001021112)
- direct_connect_enabled.hosted_vifs

<a id="canonical-1101031223230312-1111322331131323-0200300203113130-0311213121330021-2022222131122201-1322231313300121-3300232030303003-2003003330333030"></a>

Type: `"object"`. single nested block, Optional.

AWS Direct Connect Hosted VIF Configuration.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0111323310121130-2201101221311023-0031330101301023-0001003211322130-3213021000332111-3332203201213001-0302000320200313-0210332212202221"></a>

## Direct properties — hosted_vifs / 321303011030 / 3

- [site_registration_over_direct_connect](resources--aws_vpc_site--reference--group-002.md#canonical-2323323033030122-3030212303121120-0322010122103022-1200023201002022-1101001303322122-2331201121233122-3333100121111200-0313211200332200): complete subsection reference.

- [site_registration_over_internet](resources--aws_vpc_site--reference--group-002.md#canonical-2220100020100323-1312321011301011-0123000001331013-1030221010200330-2230302223123003-3032333102013030-3303230230301320-3201310030021213): complete subsection reference.

- [vif_list](resources--aws_vpc_site--reference--group-002.md#canonical-2313113013330132-1310322012212213-0200313101320333-3130221111032113-3213020303311231-3022121200002000-3013113012113121-3021311210233332): complete subsection reference.

<a id="canonical-2311201130122010-0230300323123220-1012013300211131-0203303121232223-0112003232311032-2333010313110213-2113111112030330-3110233210001200"></a>

## Next pages — hosted_vifs / 321303011030 / 4

- [direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect](resources--aws_vpc_site--reference--group-002.md#canonical-2323323033030122-3030212303121120-0322010122103022-1200023201002022-1101001303322122-2331201121233122-3333100121111200-0313211200332200)
- [direct_connect_enabled.hosted_vifs.site_registration_over_internet](resources--aws_vpc_site--reference--group-002.md#canonical-2220100020100323-1312321011301011-0123000001331013-1030221010200330-2230302223123003-3032333102013030-3303230230301320-3201310030021213)
- [direct_connect_enabled.hosted_vifs.vif_list](resources--aws_vpc_site--reference--group-002.md#canonical-2313113013330132-1310322012212213-0200313101320333-3130221111032113-3213020303311231-3022121200002000-3013113012113121-3021311210233332)
- [direct_connect_enabled](resources--aws_vpc_site--reference--group-002.md#canonical-0130012221000232-1311333022303301-0020012113300213-0322011110120320-0313333031011211-1310113213303030-0201300310003031-3011331001021112)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2323323033030122-3030212303121120-0322010122103022-1200023201002022-1101001303322122-2331201121233122-3333100121111200-0313211200332200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0311122131301130-0001200311323103-0221130032313312-1133011321203011-1302203320122232-2131320112301133-0312232031113001-0010110312100113"></a>

## direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect — site_registration_over_direct_connect / 033000321223 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [direct_connect_enabled](resources--aws_vpc_site--reference--group-002.md#canonical-0130012221000232-1311333022303301-0020012113300213-0322011110120320-0313333031011211-1310113213303030-0201300310003031-3011331001021112)
- [direct_connect_enabled.hosted_vifs](resources--aws_vpc_site--reference--group-002.md#canonical-1122123323310331-2133300100123223-2100211323201032-2222233011220302-1300003201112031-0000221033021113-3213222101010022-0121311322002320)
- direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect

<a id="canonical-0233113213130003-3222002332120202-0202020132012303-2112003001231221-1112110201301330-3030002333031002-2221100322013030-3131131322121102"></a>

Type: `"object"`. single nested block, Optional.

CloudLink ADN Network Config.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1022121312333312-3302010310302322-0013030321022300-2032202232123323-0213231101010100-2003311131321303-0313212312232110-1033231113020023"></a>

## Direct properties — site_registration_over_direct_connect / 033000321223 / 3

<a id="canonical-0112122122011031-3323022131331213-0303331313022013-1223321032001211-3322231132311031-1231332133311000-1000033120331200-3210000010220103"></a>

<a id="canonical-3200023200123213-0102022313112312-2100230002201313-0021103012333010-3032132213302223-3213100201311302-3130200203110031-3233133213310231"></a>

## cloudlink_network_name property — site_registration_over_direct_connect / 033000321223 / 4

Type: `"string"`. Optional.

Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN
network. To provision a Private ADN network, please contact F5 Distributed Cloud support.

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

<a id="canonical-1320302001321100-1212303322232233-3203220222033331-2132131203131202-1322130020023123-0221132312313010-0112331032211033-1013223103022233"></a>

## Next pages — site_registration_over_direct_connect / 033000321223 / 5

- [direct_connect_enabled.hosted_vifs](resources--aws_vpc_site--reference--group-002.md#canonical-1122123323310331-2133300100123223-2100211323201032-2222233011220302-1300003201112031-0000221033021113-3213222101010022-0121311322002320)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2220100020100323-1312321011301011-0123000001331013-1030221010200330-2230302223123003-3032333102013030-3303230230301320-3201310030021213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0102230110220033-1122000123133330-3301122003131000-3230133202102321-0011212021133030-3121220010112123-2121230302002131-0220113030222301"></a>

## direct_connect_enabled.hosted_vifs.site_registration_over_internet — site_registration_over_internet / 121220121013 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [direct_connect_enabled](resources--aws_vpc_site--reference--group-002.md#canonical-0130012221000232-1311333022303301-0020012113300213-0322011110120320-0313333031011211-1310113213303030-0201300310003031-3011331001021112)
- [direct_connect_enabled.hosted_vifs](resources--aws_vpc_site--reference--group-002.md#canonical-1122123323310331-2133300100123223-2100211323201032-2222233011220302-1300003201112031-0000221033021113-3213222101010022-0121311322002320)
- direct_connect_enabled.hosted_vifs.site_registration_over_internet

<a id="canonical-2121223101203210-1230300023011221-1331300033232302-3232103023031232-3103302021222130-0323233123032021-1131200101202221-0300233202321130"></a>

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

<a id="canonical-3120313130230013-3101133332000200-0203332321322302-2223102301003322-1202013003312323-3102323310211110-1321313013333112-2111230322330030"></a>

## Direct properties — site_registration_over_internet / 121220121013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0030211231031233-3100323110103103-3210213203000020-3233021111023231-1031201332111130-1002301321101120-2020221031033023-3002133111011211"></a>

## Next pages — site_registration_over_internet / 121220121013 / 4

- [direct_connect_enabled.hosted_vifs](resources--aws_vpc_site--reference--group-002.md#canonical-1122123323310331-2133300100123223-2100211323201032-2222233011220302-1300003201112031-0000221033021113-3213222101010022-0121311322002320)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2313113013330132-1310322012212213-0200313101320333-3130221111032113-3213020303311231-3022121200002000-3013113012113121-3021311210233332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210312313210213-1213101112020101-0223302210232103-3232110031220321-1323102000020130-2001310311232223-3201301221333201-0232111333012320"></a>

## direct_connect_enabled.hosted_vifs.vif_list — vif_list / 111223102323 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [direct_connect_enabled](resources--aws_vpc_site--reference--group-002.md#canonical-0130012221000232-1311333022303301-0020012113300213-0322011110120320-0313333031011211-1310113213303030-0201300310003031-3011331001021112)
- [direct_connect_enabled.hosted_vifs](resources--aws_vpc_site--reference--group-002.md#canonical-1122123323310331-2133300100123223-2100211323201032-2222233011220302-1300003201112031-0000221033021113-3213222101010022-0121311322002320)
- direct_connect_enabled.hosted_vifs.vif_list

<a id="canonical-2333201102213210-0202101120310320-1132312311023331-1103032023211132-2312201301021022-0201013313100110-0132103220301112-1133323331212220"></a>

Type: `"object"`. list nested block, Optional.

List of Hosted VIF Config. List of Hosted VIF Config.

Upstream description:

List of Hosted VIF Config.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2122003110313223-2201010021000312-3312133002132231-3331110011110223-0221021103103013-2212332200020321-0013200213010022-0331301032102213"></a>

## Direct properties — vif_list / 111223102323 / 3

<a id="canonical-1130002011103222-3333033110002220-1120123332111221-0212103001232233-0223231010201322-1133010312233222-3323112220012112-3030312301332033"></a>

<a id="canonical-2110212130233321-3213131323210122-1210223203320320-0221323021332233-3100101003012002-2311210112210121-2311332100102000-1331030322322311"></a>

## other_region property — vif_list / 111223102323 / 4

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
    "ves.io.schema.rules.string.in": "[\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-northeast-1\\\",\\\"ap-northeast-2\\\",\\\"ap-south-1\\\",\\\"ap-southeast-1\\\",\\\"ap-southeast-2\\\",\\\"ap-southeast-3\\\",\\\"ca-central-1\\\",\\\"eu-central-1\\\",\\\"eu-north-1\\\",\\\"eu-south-1\\\",\\\"eu-west-1\\\",\\\"eu-west-2\\\",\\\"eu-west-3\\\",\\\"me-south-1\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-1\\\",\\\"us-west-2\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-northeast-1\\\",\\\"ap-northeast-2\\\",\\\"ap-south-1\\\",\\\"ap-southeast-1\\\",\\\"ap-southeast-2\\\",\\\"ap-southeast-3\\\",\\\"ca-central-1\\\",\\\"eu-central-1\\\",\\\"eu-north-1\\\",\\\"eu-south-1\\\",\\\"eu-west-1\\\",\\\"eu-west-2\\\",\\\"eu-west-3\\\",\\\"me-south-1\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-1\\\",\\\"us-west-2\\\"]"
  }
}
```

- [same_as_site_region](resources--aws_vpc_site--reference--group-002.md#canonical-2300201201022101-2320221201033210-0332003230233133-2113021123002023-3301003032123313-2100000013120330-0310202122212231-3330232103310012): complete subsection reference.

<a id="canonical-0013013012120010-2333112211111230-1110111120030303-1113200101311301-0313032311131112-1320311220110310-0020011202102222-2202031200010201"></a>

<a id="canonical-2230332023301011-0100131331203003-1200130131333122-2031303133331012-2313013112122230-1201112111003313-2123100120320030-1202023301102100"></a>

## vif_id property — vif_list / 111223102323 / 5

Type: `"string"`. Optional.

AWS Direct Connect VIF ID that needs to be connected to the site.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0322121230022003-3300301232102033-1111321221021100-0032303121212312-3311210213101333-2321123333222132-1310222201122030-1233122322302303"></a>

## Next pages — vif_list / 111223102323 / 6

- [direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region](resources--aws_vpc_site--reference--group-002.md#canonical-2300201201022101-2320221201033210-0332003230233133-2113021123002023-3301003032123313-2100000013120330-0310202122212231-3330232103310012)
- [direct_connect_enabled.hosted_vifs](resources--aws_vpc_site--reference--group-002.md#canonical-1122123323310331-2133300100123223-2100211323201032-2222233011220302-1300003201112031-0000221033021113-3213222101010022-0121311322002320)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2300201201022101-2320221201033210-0332003230233133-2113021123002023-3301003032123313-2100000013120330-0310202122212231-3330232103310012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1132330233130222-0010022132032023-0020322010233200-1013000333033310-3300132321033131-0231310312332122-1012112011030212-3103131220231102"></a>

## direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region — same_as_site_region / 131222202220 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [direct_connect_enabled](resources--aws_vpc_site--reference--group-002.md#canonical-0130012221000232-1311333022303301-0020012113300213-0322011110120320-0313333031011211-1310113213303030-0201300310003031-3011331001021112)
- [direct_connect_enabled.hosted_vifs](resources--aws_vpc_site--reference--group-002.md#canonical-1122123323310331-2133300100123223-2100211323201032-2222233011220302-1300003201112031-0000221033021113-3213222101010022-0121311322002320)
- [direct_connect_enabled.hosted_vifs.vif_list](resources--aws_vpc_site--reference--group-002.md#canonical-2313113013330132-1310322012212213-0200313101320333-3130221111032113-3213020303311231-3022121200002000-3013113012113121-3021311210233332)
- direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region

<a id="canonical-1232211021130022-3030013003322300-0332111200031210-1130302122022213-0233132203311312-2231311303110010-0031111112311012-2131031030311331"></a>

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

<a id="canonical-3023231321130113-1111331132331333-2112100212002121-0202310310231222-3000023133221133-0313232000121222-3311010003122000-1131322023022320"></a>

## Direct properties — same_as_site_region / 131222202220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1132011100300230-2123111131133102-2122133301213223-1132133011303201-2312110110003332-3310132022303021-3021030320122303-2030022230301033"></a>

## Next pages — same_as_site_region / 131222202220 / 4

- [direct_connect_enabled.hosted_vifs.vif_list](resources--aws_vpc_site--reference--group-002.md#canonical-2313113013330132-1310322012212213-0200313101320333-3130221111032113-3213020303311231-3022121200002000-3013113012113121-3021311210233332)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3220022212033112-0320203202221210-0131000010330021-1213200323030110-3310310031011301-0323210131313230-3321022313332331-3303231113000022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112310223102013-2221112332231203-1212311001303103-3223221123313121-0032330013213210-0011021113103001-3213110010130002-3010201301302101"></a>

## direct_connect_enabled.standard_vifs — standard_vifs / 132102032332 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [direct_connect_enabled](resources--aws_vpc_site--reference--group-002.md#canonical-0130012221000232-1311333022303301-0020012113300213-0322011110120320-0313333031011211-1310113213303030-0201300310003031-3011331001021112)
- direct_connect_enabled.standard_vifs

<a id="canonical-3133133203221301-1012003031223023-2323213222312002-0201022003100310-0320000333212313-2130200020000101-2320301100330200-3112312010202022"></a>

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

<a id="canonical-1220131032012202-3331311321202233-2231222230101003-2222223013312103-1131120120022203-3200321321322110-3103332121031120-1013031312033013"></a>

## Direct properties — standard_vifs / 132102032332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1310210230202110-2120001333001100-1110320210310233-2301013031233212-0303103320133223-2203122030011003-3110020013102211-0020230031120232"></a>

## Next pages — standard_vifs / 132102032332 / 4

- [direct_connect_enabled](resources--aws_vpc_site--reference--group-002.md#canonical-0130012221000232-1311333022303301-0020012113300213-0322011110120320-0313333031011211-1310113213303030-0201300310003031-3011331001021112)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1202020332313000-1201233220303020-0132102003011220-1001221330331021-2013110031202200-3000002230300223-3313121333113121-0301333223301031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1333322122210011-1302122331333302-3201103123203033-2100211110110002-1230113302222302-0320123032013011-3131012122030322-3113222323321123"></a>

## disable_encryption — disable_encryption / 203033133302 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- disable_encryption

<a id="canonical-0323231202220003-0101221322133103-3212333233323131-1313300023012021-3011111012121102-1232330310112321-0303111313121311-2213330112232132"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_encryption, enable\_encryption; Default: disable\_encryption\] Configuration
parameter for disable encryption.

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

- [disable_encryption](resources--aws_vpc_site--reference--group-002.md#canonical-0323231202220003-0101221322133103-3212333233323131-1313300023012021-3011111012121102-1232330310112321-0303111313121311-2213330112232132)
- [enable_encryption](resources--aws_vpc_site--reference--group-002.md#canonical-2030311130233011-3303201321301001-3102011303313330-3232233031320132-3000002203131313-1332111113111132-3210221230202200-2310130331201220)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_encryption = {}
```

<a id="canonical-1303013111021311-1211002222013122-3111111123322120-0310123322211030-3111301332300031-1122230210032001-1132231030123223-3131023222131222"></a>

## Direct properties — disable_encryption / 203033133302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2231331001002201-0311100331202223-0312032121112001-1013301001313000-0211320321231133-0030011102032121-0122212331303011-1230213201011132"></a>

## Next pages — disable_encryption / 203033133302 / 4

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1000122022313012-0001101010202330-0013311322231203-1310123033123032-2000220333300012-0323120330133221-0332121010210123-3232302323030130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123301132211000-3101103332313110-3232333010130131-2220033100130213-0033222000111130-3023311212321213-2333310031311101-1012201313312220"></a>

## disable_internet_vip — disable_internet_vip / 322220231202 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- disable_internet_vip

<a id="canonical-1310233311202123-3210112112211322-1310221310030110-1100312013001000-1201111100100111-0013321222221030-1010101033002311-0313011231000023"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_internet\_vip, enable\_internet\_vip; Default: disable\_internet\_vip\] Enable
this option

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

- [disable_internet_vip](resources--aws_vpc_site--reference--group-002.md#canonical-1310233311202123-3210112112211322-1310221310030110-1100312013001000-1201111100100111-0013321222221030-1010101033002311-0313011231000023)
- [enable_internet_vip](resources--aws_vpc_site--reference--group-002.md#canonical-3321320320220100-0223310032031223-2222200212223211-0111032202012321-0331211130002101-1200202222101320-3120232022003230-3333020010213203)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_internet_vip = {}
```

<a id="canonical-3213231032223311-3321100310132111-1100303321003023-3200202223020012-0032022223200100-2221313230201332-1013201212211221-2310322232112300"></a>

## Direct properties — disable_internet_vip / 322220231202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1203030302030033-0002230020012232-2032331330200233-2210232311330311-2020220122322302-0221300033020302-2131312323103110-3230001203111132"></a>

## Next pages — disable_internet_vip / 322220231202 / 4

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1212023023013013-1212222311211002-3121310102313131-0302032221222312-3030030032230033-3321120132111211-1111123202013032-2223010303001021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102131311001322-0210300330000123-3300003231031302-0031020132033330-2203010332031103-3233210102233130-1322022121303310-3032232010311130"></a>

## egress_gateway_default — egress_gateway_default / 020312000203 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- egress_gateway_default

<a id="canonical-0223132030103101-2211331202122112-3323230320200103-0323133301320000-0132102121321000-3012123233012131-3223212203022131-1303123221003130"></a>

Type: `["object", {}]`. Optional.

\[OneOf: egress\_gateway\_default, egress\_nat\_gw, egress\_virtual\_private\_gateway; Default:
egress\_gateway\_default\] Configuration parameter for egress gateway default.

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

- [egress_gateway_default](resources--aws_vpc_site--reference--group-002.md#canonical-0223132030103101-2211331202122112-3323230320200103-0323133301320000-0132102121321000-3012123233012131-3223212203022131-1303123221003130)
- [egress_nat_gw](resources--aws_vpc_site--reference--group-002.md#canonical-2310103003033300-3011023033302013-2310233303210222-2022221301112130-3031200233302223-3213330130131303-0313213120103231-3332112010330020)
- [egress_virtual_private_gateway](resources--aws_vpc_site--reference--group-002.md#canonical-2122031301111101-2033231230130122-3331333013121230-3302201001302102-2220031020200021-1010122013112100-1233131032210133-3212033312022132)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
egress_gateway_default = {}
```

<a id="canonical-1231202203333332-3113001200310330-2112331020002211-1230230232230213-0022212102222333-1123020032313100-2103023222002131-0310233012110301"></a>

## Direct properties — egress_gateway_default / 020312000203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3222001200211302-1310020110020232-1313320200111022-2333013131033022-3020323223020021-2110223332321010-2320002200213310-0301112211000112"></a>

## Next pages — egress_gateway_default / 020312000203 / 4

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2132333123021112-2121130110202202-3232330110112120-1023112123222122-2013332222002322-0212030223100002-1131133120313012-3111311320301310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300332111231330-2122330221333001-3111320022111221-0321031221120102-2123203310131322-0010120113021102-1010231020332323-0222031202020232"></a>

## egress_nat_gw — egress_nat_gw / 232333113211 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- egress_nat_gw

<a id="canonical-2310103003033300-3011023033302013-2310233303210222-2022221301112130-3031200233302223-3213330130131303-0313213120103231-3332112010330020"></a>

Type: `"object"`. single nested block, Optional.

With this option, egress site traffic will be routed through an Network Address Translation(NAT)
Gateway.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"nat_gw_id\"]"
}
```

Terraform syntax:

```terraform
egress_nat_gw {
  # Configure direct properties listed below.
}
```

<a id="canonical-1233312131323333-1331213303312113-0021300223111103-2003102122220212-2232233032012002-3331232002132220-3210010331032231-0330112103310132"></a>

## Direct properties — egress_nat_gw / 232333113211 / 3

<a id="canonical-0223310102302221-2313322230132331-3131100132302033-1011200300000110-1033120100131311-1301111013002321-3130133111320130-2113130323222212"></a>

<a id="canonical-0230130231230312-1201031321021211-3102332120230212-1021310300020122-2110131313133012-2332033002101200-2332011320303200-1111031110132101"></a>

## nat_gw_id property — egress_nat_gw / 232333113211 / 4

Type: `"string"`. Optional.

Existing NAT Gateway ID. Exclusive with \[\]

Upstream description:

Exclusive with \[\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(21),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 21,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 21,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^(nat-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "21",
    "ves.io.schema.rules.string.pattern": "^(nat-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "21",
    "ves.io.schema.rules.string.pattern": "^(nat-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

<a id="canonical-2330322000022000-3300013330303102-0100011020212201-3133112011231312-0001222313122320-0102220312232101-0220013212220023-2232132333201300"></a>

## Next pages — egress_nat_gw / 232333113211 / 5

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0113110013033313-2210303123332123-0133200222112102-1120033121100103-1132200021011211-3010030013230211-0210122122123323-0310322011033303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3013333111010213-2333132101333002-3101211301330022-0000023321020123-0111013112300032-2202202030033013-3121110311011110-2310332031022320"></a>

## egress_virtual_private_gateway — egress_virtual_private_gateway / 213330233303 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- egress_virtual_private_gateway

<a id="canonical-2122031301111101-2033231230130122-3331333013121230-3302201001302102-2220031020200021-1010122013112100-1233131032210133-3212033312022132"></a>

Type: `"object"`. single nested block, Optional.

With this option, egress site traffic will be routed through an Virtual Private Gateway.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"vgw_id\"]"
}
```

Terraform syntax:

```terraform
egress_virtual_private_gateway {
  # Configure direct properties listed below.
}
```

<a id="canonical-1100303112210123-3331321221323301-0001002331300113-2031301220320322-1012232202100003-1000120123202010-2203002302323111-0220000310200330"></a>

## Direct properties — egress_virtual_private_gateway / 213330233303 / 3

<a id="canonical-1323032211110000-1000013300221002-2223312013223311-2020231313110032-2131320102302220-0310222202112311-2012032201313001-1302000123020212"></a>

<a id="canonical-0213021003312102-2102320123301321-1212322002220223-3333111203231122-3133322233212323-2022230332123212-3132023332000101-2331112311003213"></a>

## vgw_id property — egress_virtual_private_gateway / 213330233303 / 4

Type: `"string"`. Optional.

Existing Virtual Private Gateway ID. Exclusive with \[\]

Upstream description:

Exclusive with \[\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(21),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 21,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 21,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^(vgw-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "21",
    "ves.io.schema.rules.string.pattern": "^(vgw-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "21",
    "ves.io.schema.rules.string.pattern": "^(vgw-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

<a id="canonical-1021111323012320-1010122100322021-1013233131333033-2233102002110202-1100111213103321-3110230012223030-1300133322013312-2021003231100130"></a>

## Next pages — egress_virtual_private_gateway / 213330233303 / 5

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3122102032303111-0233131232202220-1001102333222013-2102002002231000-3013103312333313-2003211130120332-1010200300312332-3210020213022122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0221233132021132-1233231122101101-0212222111332322-1100303323003122-0102210211023022-3330203032033032-3033110113220331-0101211323330320"></a>

## enable_encryption — enable_encryption / 202212131220 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- enable_encryption

<a id="canonical-2030311130233011-3303201321301001-3102011303313330-3232233031320132-3000002203131313-1332111113111132-3210221230202200-2310130331201220"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for enable encryption.

Upstream description:

Information related to disk encryption.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("kms_key_id")}
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
enable_encryption {
  # Configure direct properties listed below.
}
```

<a id="canonical-2231031110112001-1112102221311023-3201233221130103-3301000023011313-3213212331220031-2213321322321123-3030313231313002-3232023311321231"></a>

## Direct properties — enable_encryption / 202212131220 / 3

<a id="canonical-3223130210331021-3121212000201300-1312210020210320-3221033031132001-3302200010202033-0203223202133133-2233011133020100-0113200001111031"></a>

<a id="canonical-0203320002210203-3101220012000320-2133333231010312-2320332232233213-0200123133132122-2210132333212213-2201220301001001-2010211321313023"></a>

## kms_key_id property — enable_encryption / 202212131220 / 4

Type: `"string"`. Optional.

AWS KMS Key to be used to encrypt the disk attached to the VM.

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

<a id="canonical-0300213020022233-2023222302023232-0030231203030023-1203233123231012-1313330223203111-1210320100033011-2123320101212201-2011302210122222"></a>

## Next pages — enable_encryption / 202212131220 / 5

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3302122201112333-3332110213202311-1021212130302332-1120320210103133-1101020320200200-2103031301023231-1200332010031321-0213300302110322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301301311120113-2122210231222123-2321210112122013-1131133020111322-0030011013201112-0220002022011311-2233110311123100-2130003120310303"></a>

## enable_internet_vip — enable_internet_vip / 220021311201 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- enable_internet_vip

<a id="canonical-3321320320220100-0223310032031223-2222200212223211-0111032202012321-0331211130002101-1200202222101320-3120232022003230-3333020010213203"></a>

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
enable_internet_vip = {}
```

<a id="canonical-1003331130213010-1033320001021121-1233010023121300-3033220222310211-1211230320321130-1323220221100110-1301022112012303-2133111210003332"></a>

## Direct properties — enable_internet_vip / 220021311201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1301100031301302-3222123000322223-1213010020123201-3012130313112101-0013311000212203-3313330023221303-2312331021200210-1332122332320002"></a>

## Next pages — enable_internet_vip / 220021311201 / 4

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0233100333233221-1330000203320002-0303020331312332-0103212301123033-1132021203111231-2112210122320010-2313131333302212-2132032101022030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321200130010103-0323310321100232-3212112002212230-2312320130132223-0313333002330301-1102013302231031-0032313031013323-3220213323320003"></a>

## f5_orchestrated_routing — f5_orchestrated_routing / 302212313213 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- f5_orchestrated_routing

<a id="canonical-1200010232012020-1233301121032220-0303220310331300-3022031312111003-2212022021232121-0223020131131120-1330100021330313-0013231021111023"></a>

Type: `["object", {}]`. Optional.

\[OneOf: f5\_orchestrated\_routing, manual\_routing\] Enable this option

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

- [f5_orchestrated_routing](resources--aws_vpc_site--reference--group-002.md#canonical-1200010232012020-1233301121032220-0303220310331300-3022031312111003-2212022021232121-0223020131131120-1330100021330313-0013231021111023)
- [manual_routing](resources--aws_vpc_site--reference--group-004.md#canonical-3020121031032231-1103013030230103-3031133330101100-2301130023123222-1233013021101120-3103232331131332-0022200132303212-1113101010212212)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
f5_orchestrated_routing = {}
```

<a id="canonical-0033331003020122-3000302222013102-0303110110032111-3023231302220002-3232320312101112-2312223121313303-3111023133010230-2020210320201322"></a>

## Direct properties — f5_orchestrated_routing / 302212313213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0013332011223122-3331212130210013-2011100111201003-1231120120213231-2231300110010010-0202330022123322-1321213133213231-3303323022331300"></a>

## Next pages — f5_orchestrated_routing / 302212313213 / 4

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0122030011102010-3122030030122331-3301230100222300-0332133333222221-0201113331023131-0102122321211202-2132001103112213-0230321112200201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1331221122020203-2131302021131032-1110212223202300-3121132102332313-0100133322332112-2200111330130302-3212121211231223-0103301102013022"></a>

## f5xc_security_group — f5xc_security_group / 320331033010 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- f5xc_security_group

<a id="canonical-3111232222322323-0103012011313123-2133002030330003-2311220201201222-3222332113322223-3021300030333002-3002021233321301-2220013302333233"></a>

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
f5xc_security_group = {}
```

<a id="canonical-2301212111211030-1223212330211102-3213122311320220-2131231123313112-3321221133221331-3010002330122023-0210120212322312-2331321001003011"></a>

## Direct properties — f5xc_security_group / 320331033010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3233231333121323-2020333210201010-3101022322033303-1223212030033102-2102113121331201-3210103331303322-1200232023013000-3211012101311322"></a>

## Next pages — f5xc_security_group / 320331033010 / 4

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133201032020301-2111032121111203-2101331123101303-0212211022210333-2112200110213323-2233321230312331-0300110133320231-2233033212300132"></a>

## ingress_egress_gw — ingress_egress_gw / 213022110121 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- ingress_egress_gw

<a id="canonical-3113010322002120-3232103130233031-3232200003001203-3320332313320331-2032332113301301-0323301030223020-1201010320302120-2133302011211122"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: ingress\_egress\_gw, ingress\_gw, voltstack\_cluster\] Configuration parameter for ingress
egress gw.

Upstream description:

Two interface AWS ingress/egress site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("aws_certified_hw",
    "az_nodes"),
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
  validators.ConflictingObjectAttributes("dc_cluster_group_inside_vn",
    "dc_cluster_group_outside_vn"),
  validators.ConflictingObjectAttributes("dc_cluster_group_inside_vn",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("dc_cluster_group_outside_vn",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("forward_proxy_allow_all",
    "no_forward_proxy"),
  validators.ConflictingObjectAttributes("global_network_list",
    "no_global_network"),
  validators.ConflictingObjectAttributes("inside_static_routes",
    "no_inside_static_routes"),
  validators.ConflictingObjectAttributes("no_outside_static_routes",
    "outside_static_routes"),
  validators.ConflictingObjectAttributes("sm_connection_public_ip",
    "sm_connection_pvt_ip")}
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
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group_inside_vn\",\"dc_cluster_group_outside_vn\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-inside_static_route_choice": "[\"inside_static_routes\",\"no_inside_static_routes\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]",
  "x-ves-oneof-field-outside_static_route_choice": "[\"no_outside_static_routes\",\"outside_static_routes\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]"
}
```

OneOf alternatives in this subsection:

- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-3113010322002120-3232103130233031-3232200003001203-3320332313320331-2032332113301301-0323301030223020-1201010320302120-2133302011211122)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-1213030012102200-2132122202121220-2010120201212311-1321031001321231-3302202232203113-2220103102300303-3322131013213033-1113012313110132)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-1001100130213133-3230331301033031-0331333111030200-0210320013203113-1123312203213202-1023022223312332-2230321323121003-3130130333013112)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
ingress_egress_gw {
  # Configure direct properties listed below.
}
```

<a id="canonical-0310013202120321-0130201211013301-1300212321023021-2103230112331032-1200212203303130-2230301120130001-0310330202101130-3021131230112000"></a>

## Direct properties — ingress_egress_gw / 213022110121 / 3

- [active_enhanced_firewall_policies](resources--aws_vpc_site--reference--group-002.md#canonical-1033121032210230-1220122323032323-3323321200222202-1110200022212200-1132302331001233-2022303302023312-2110120130120221-0302312113013202): complete subsection reference.

- [active_forward_proxy_policies](resources--aws_vpc_site--reference--group-002.md#canonical-2032010023210310-0333120203113312-2103312311112232-3321203111222313-2231210132332030-2321303012031232-1312103001310131-1330122232003213): complete subsection reference.

- [active_network_policies](resources--aws_vpc_site--reference--group-002.md#canonical-2132102132020200-1223012023300303-2203302223002220-3033032020232223-2101123323233323-2302223031001002-0302203010321310-2113300121321300): complete subsection reference.

- [allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-0200133132302023-1321231123110331-1303133003320320-2302223311301000-0123113020122311-1231122311003302-3211310321102310-0123330100231003): complete subsection reference.

- [allowed_vip_port_sli](resources--aws_vpc_site--reference--group-002.md#canonical-0321200102110013-2221113222102210-3011001233023112-0002332023300211-0312212112213020-1212311131213112-3011010202130131-3002322100221021): complete subsection reference.

<a id="canonical-2312333230212132-0230232310213233-3110100031120222-2331322232312031-0013213202230020-1121001102102312-3211103232102123-1020313032212210"></a>

<a id="canonical-3133312221002230-0330013123220110-0321123213322032-0303123112300330-1111031200000321-2020212332313311-1233033331201331-1101111213310131"></a>

## aws_certified_hw property — ingress_egress_gw / 213022110121 / 4

Type: `"string"`. Optional.

\[Enum: aws-byol-multi-nic-voltmesh\] AWS Certified Hardware. Name for AWS certified hardware. The
only possible value is \`aws-byol-multi-nic-voltmesh\`.

Upstream description:

Name for AWS certified hardware.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
  stringvalidator.OneOf("aws-byol-multi-nic-voltmesh"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "aws-byol-multi-nic-voltmesh"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"aws-byol-multi-nic-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"aws-byol-multi-nic-voltmesh\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [az_nodes](resources--aws_vpc_site--reference--group-002.md#canonical-3233231103113313-1010330103013321-3013311222003130-2100031033021221-1113023133322203-0220023312322322-3302003000213032-2123300021210202): complete subsection reference.

- [dc_cluster_group_inside_vn](resources--aws_vpc_site--reference--group-002.md#canonical-1202110032001010-2121110311111232-0033110201121022-0202332122110012-3010302301220221-3302211103130233-2331332211132321-3101203131220211): complete subsection reference.

- [dc_cluster_group_outside_vn](resources--aws_vpc_site--reference--group-002.md#canonical-0011011033131033-3121313210300020-2031133213011231-0312232322123203-2220013213011302-3230300023111123-3300201220331232-1211100120113031): complete subsection reference.

- [forward_proxy_allow_all](resources--aws_vpc_site--reference--group-002.md#canonical-2331012323131020-2122213031302230-3102311133021323-2212313023122230-2032000330333110-1302122010311300-2133130321011200-1112320101103201): complete subsection reference.

- [global_network_list](resources--aws_vpc_site--reference--group-002.md#canonical-0110112001213332-2233201102333212-1321232102201102-3213031203300330-1210310211301200-3133010302300211-1313222203301301-0033103200013221): complete subsection reference.

- [inside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-0122332302112321-3223212101323020-1120103012312002-2002013101203201-1312030030031113-2030131330301001-3032023323011313-1030122223331032): complete subsection reference.

- [no_dc_cluster_group](resources--aws_vpc_site--reference--group-003.md#canonical-2102113123102003-2220310101303013-0300012200101230-2331120032031002-2020320223320023-1023022111122321-3001131033200032-2103330230232000): complete subsection reference.

- [no_forward_proxy](resources--aws_vpc_site--reference--group-003.md#canonical-0300100013213131-2213313111311121-2320032011010033-1121103120121321-3332103121230003-0311021312130233-1200313321213133-3212230332110001): complete subsection reference.

- [no_global_network](resources--aws_vpc_site--reference--group-003.md#canonical-2303130302120032-0011211313313230-1011003323201320-2200003100222230-1300022023002331-3030200233303333-3232203312323132-0003221101001120): complete subsection reference.

- [no_inside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-1333002030033131-2021313011200121-0222321022222323-0132313232113010-1221003222232002-0113220201003323-2012030231320232-2301303310222211): complete subsection reference.

- [no_network_policy](resources--aws_vpc_site--reference--group-003.md#canonical-0333203101111010-1010131002112103-3313120201301303-3330022000222010-2110330132312202-3231120103313130-2213103003331033-2202002022021232): complete subsection reference.

- [no_outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-3212221300032003-1100311223312301-0322102032310120-2122130010202011-0310231012101011-1233013031023021-3230010032221120-2221220331311101): complete subsection reference.

- [outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-0132220013120133-3210001010133330-0310012213211212-0132000030002010-0032200333312300-0203100023213033-1003303232322110-3220022313232002): complete subsection reference.

- [performance_enhancement_mode](resources--aws_vpc_site--reference--group-003.md#canonical-2133010232202133-3021221010321021-2013032211210330-3301312300002132-3000011030113011-2311302233333003-2311302030320122-0211301220311201): complete subsection reference.

- [sm_connection_public_ip](resources--aws_vpc_site--reference--group-003.md#canonical-2002101132330123-2132301310023301-2122003302320220-0032010113032131-2000321312020231-2331331002032220-0312033010032022-3001123020212102): complete subsection reference.

- [sm_connection_pvt_ip](resources--aws_vpc_site--reference--group-003.md#canonical-2022002132013302-3133212203112321-3201130121023311-1321200003012130-3121321120003222-0120032003022210-3110123100000102-3112320312011021): complete subsection reference.

<a id="canonical-3232003222313132-3321112012112111-2120113310310112-3333001333302221-2010102311231213-1222022233121222-3112302232230123-3321132302311012"></a>

## Next pages — ingress_egress_gw / 213022110121 / 5

- [ingress_egress_gw.active_enhanced_firewall_policies](resources--aws_vpc_site--reference--group-002.md#canonical-1033121032210230-1220122323032323-3323321200222202-1110200022212200-1132302331001233-2022303302023312-2110120130120221-0302312113013202)
- [ingress_egress_gw.active_forward_proxy_policies](resources--aws_vpc_site--reference--group-002.md#canonical-2032010023210310-0333120203113312-2103312311112232-3321203111222313-2231210132332030-2321303012031232-1312103001310131-1330122232003213)
- [ingress_egress_gw.active_network_policies](resources--aws_vpc_site--reference--group-002.md#canonical-2132102132020200-1223012023300303-2203302223002220-3033032020232223-2101123323233323-2302223031001002-0302203010321310-2113300121321300)
- [ingress_egress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-0200133132302023-1321231123110331-1303133003320320-2302223311301000-0123113020122311-1231122311003302-3211310321102310-0123330100231003)
- [ingress_egress_gw.allowed_vip_port_sli](resources--aws_vpc_site--reference--group-002.md#canonical-0321200102110013-2221113222102210-3011001233023112-0002332023300211-0312212112213020-1212311131213112-3011010202130131-3002322100221021)
- [ingress_egress_gw.az_nodes](resources--aws_vpc_site--reference--group-002.md#canonical-3233231103113313-1010330103013321-3013311222003130-2100031033021221-1113023133322203-0220023312322322-3302003000213032-2123300021210202)
- [ingress_egress_gw.dc_cluster_group_inside_vn](resources--aws_vpc_site--reference--group-002.md#canonical-1202110032001010-2121110311111232-0033110201121022-0202332122110012-3010302301220221-3302211103130233-2331332211132321-3101203131220211)
- [ingress_egress_gw.dc_cluster_group_outside_vn](resources--aws_vpc_site--reference--group-002.md#canonical-0011011033131033-3121313210300020-2031133213011231-0312232322123203-2220013213011302-3230300023111123-3300201220331232-1211100120113031)
- [ingress_egress_gw.forward_proxy_allow_all](resources--aws_vpc_site--reference--group-002.md#canonical-2331012323131020-2122213031302230-3102311133021323-2212313023122230-2032000330333110-1302122010311300-2133130321011200-1112320101103201)
- [ingress_egress_gw.global_network_list](resources--aws_vpc_site--reference--group-002.md#canonical-0110112001213332-2233201102333212-1321232102201102-3213031203300330-1210310211301200-3133010302300211-1313222203301301-0033103200013221)
- [ingress_egress_gw.inside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-0122332302112321-3223212101323020-1120103012312002-2002013101203201-1312030030031113-2030131330301001-3032023323011313-1030122223331032)
- [ingress_egress_gw.no_dc_cluster_group](resources--aws_vpc_site--reference--group-003.md#canonical-2102113123102003-2220310101303013-0300012200101230-2331120032031002-2020320223320023-1023022111122321-3001131033200032-2103330230232000)
- [ingress_egress_gw.no_forward_proxy](resources--aws_vpc_site--reference--group-003.md#canonical-0300100013213131-2213313111311121-2320032011010033-1121103120121321-3332103121230003-0311021312130233-1200313321213133-3212230332110001)
- [ingress_egress_gw.no_global_network](resources--aws_vpc_site--reference--group-003.md#canonical-2303130302120032-0011211313313230-1011003323201320-2200003100222230-1300022023002331-3030200233303333-3232203312323132-0003221101001120)
- [ingress_egress_gw.no_inside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-1333002030033131-2021313011200121-0222321022222323-0132313232113010-1221003222232002-0113220201003323-2012030231320232-2301303310222211)
- [ingress_egress_gw.no_network_policy](resources--aws_vpc_site--reference--group-003.md#canonical-0333203101111010-1010131002112103-3313120201301303-3330022000222010-2110330132312202-3231120103313130-2213103003331033-2202002022021232)
- [ingress_egress_gw.no_outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-3212221300032003-1100311223312301-0322102032310120-2122130010202011-0310231012101011-1233013031023021-3230010032221120-2221220331311101)
- [ingress_egress_gw.outside_static_routes](resources--aws_vpc_site--reference--group-003.md#canonical-0132220013120133-3210001010133330-0310012213211212-0132000030002010-0032200333312300-0203100023213033-1003303232322110-3220022313232002)
- [ingress_egress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-003.md#canonical-2133010232202133-3021221010321021-2013032211210330-3301312300002132-3000011030113011-2311302233333003-2311302030320122-0211301220311201)
- [ingress_egress_gw.sm_connection_public_ip](resources--aws_vpc_site--reference--group-003.md#canonical-2002101132330123-2132301310023301-2122003302320220-0032010113032131-2000321312020231-2331331002032220-0312033010032022-3001123020212102)
- [ingress_egress_gw.sm_connection_pvt_ip](resources--aws_vpc_site--reference--group-003.md#canonical-2022002132013302-3133212203112321-3201130121023311-1321200003012130-3121321120003222-0120032003022210-3110123100000102-3112320312011021)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1033121032210230-1220122323032323-3323321200222202-1110200022212200-1132302331001233-2022303302023312-2110120130120221-0302312113013202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0220023002120213-0300020311310111-2031233222002111-1100032302000120-3001300032201010-3023023233033132-1321002010311031-3322301230120012"></a>

## ingress_egress_gw.active_enhanced_firewall_policies — active_enhanced_firewall_policies / 100020031322 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- ingress_egress_gw.active_enhanced_firewall_policies

<a id="canonical-0302032122011233-1222313111320012-3203112003101123-1002200201121311-0001303311103010-1122211102110021-2131010122330120-2200230213213022"></a>

Type: `"object"`. single nested block, Optional.

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Upstream description:

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0133300322020003-2131332022311232-3211301121023111-0120231131101030-1010233100120221-2200002233103032-2033121302333010-0323011233001233"></a>

## Direct properties — active_enhanced_firewall_policies / 100020031322 / 3

- [enhanced_firewall_policies](resources--aws_vpc_site--reference--group-002.md#canonical-2331322311231020-3322020201300313-3030000030231333-1111030030220011-2133131323330101-1201230222332132-2320203101222013-2322322232320111): complete subsection reference.

<a id="canonical-1101200323121001-1133200213123003-2220321312201100-0112231202100222-0321123322300130-1122003130123130-1132130113101122-3100300000331100"></a>

## Next pages — active_enhanced_firewall_policies / 100020031322 / 4

- [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies](resources--aws_vpc_site--reference--group-002.md#canonical-2331322311231020-3322020201300313-3030000030231333-1111030030220011-2133131323330101-1201230222332132-2320203101222013-2322322232320111)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2331322311231020-3322020201300313-3030000030231333-1111030030220011-2133131323330101-1201230222332132-2320203101222013-2322322232320111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031123023031003-2303003311102122-0310101121220123-2002012320220020-2133203331132011-1202222120320011-3021223200302031-1230133320100001"></a>

## ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies — enhanced_firewall_policies / 322321202221 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.active_enhanced_firewall_policies](resources--aws_vpc_site--reference--group-002.md#canonical-1033121032210230-1220122323032323-3323321200222202-1110200022212200-1132302331001233-2022303302023312-2110120130120221-0302312113013202)
- ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies

<a id="canonical-1200320122131210-1002030002311211-3023331333230300-2220303202013230-2233120020102331-1132101032002311-3011022120120132-0211330202133032"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Enhanced Firewall Policies active.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2002130020332023-0102231301211313-2320203033003302-1121102220002301-1231212133133131-1320223302122031-3311223322111113-2323301001130022"></a>

## Direct properties — enhanced_firewall_policies / 322321202221 / 3

<a id="canonical-2031120332230023-2122333110000210-2031301202232313-0112022003301320-2321203301233131-0321022033331233-1202330300133013-3311113232212011"></a>

<a id="canonical-3230201112103322-3311311313130233-2203323033211333-2020311121120232-3012120223300212-2330332123233013-0133011202010222-3112310330202021"></a>

## name property — enhanced_firewall_policies / 322321202221 / 4

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

<a id="canonical-0123202001002311-0330333230212211-2201300131001221-0330221031330023-3323310132222031-0102211300031331-3021233333030312-3321102013221123"></a>

<a id="canonical-1012010113023103-2210312121113020-1010010120000110-3231101310132132-2100010320333023-2013202103320111-1101103110113201-1031131312111202"></a>

## namespace property — enhanced_firewall_policies / 322321202221 / 5

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

<a id="canonical-0301322020031310-3333112111033020-0012330331100123-2222200133100033-0103221021133021-3103302201230330-0332133202223332-1320121330203012"></a>

<a id="canonical-1003232222123110-0103311001312101-2001023120012112-0230220330020323-0012012320013112-1312310012012002-0012203323231320-0003321231302000"></a>

## tenant property — enhanced_firewall_policies / 322321202221 / 6

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

<a id="canonical-0232020330010311-3112332020013311-2033321130213313-1300020300331021-0322330203323322-2323301000331221-1210022110131010-2230221133002132"></a>

## Next pages — enhanced_firewall_policies / 322321202221 / 7

- [ingress_egress_gw.active_enhanced_firewall_policies](resources--aws_vpc_site--reference--group-002.md#canonical-1033121032210230-1220122323032323-3323321200222202-1110200022212200-1132302331001233-2022303302023312-2110120130120221-0302312113013202)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2032010023210310-0333120203113312-2103312311112232-3321203111222313-2231210132332030-2321303012031232-1312103001310131-1330122232003213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123032011333030-3231133321222010-0001132301122332-1033233100102033-2021202221232332-1011032201100103-3321002211132212-3333330130102332"></a>

## ingress_egress_gw.active_forward_proxy_policies — active_forward_proxy_policies / 123321133031 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- ingress_egress_gw.active_forward_proxy_policies

<a id="canonical-0133021223021010-3333110210310313-3022202231231010-1003032200022322-0122101201221313-1211233320010232-2331331033102323-1302312330113301"></a>

Type: `"object"`. single nested block, Optional.

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0112122102011211-3013130020103223-2000211303030232-1023131020102003-3110103103231201-3010103030303200-1010301301110020-0021322310002232"></a>

## Direct properties — active_forward_proxy_policies / 123321133031 / 3

- [forward_proxy_policies](resources--aws_vpc_site--reference--group-002.md#canonical-0220311133120120-0310112031320303-1113313032101231-1311002133032323-2301021123010321-0322002101331331-1011212233121120-1302230100030020): complete subsection reference.

<a id="canonical-2032200110330101-0130021300330310-1202010301000133-0330213223302121-2110022100032031-0011221023321112-2201230302202030-2021333212320321"></a>

## Next pages — active_forward_proxy_policies / 123321133031 / 4

- [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies](resources--aws_vpc_site--reference--group-002.md#canonical-0220311133120120-0310112031320303-1113313032101231-1311002133032323-2301021123010321-0322002101331331-1011212233121120-1302230100030020)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0220311133120120-0310112031320303-1113313032101231-1311002133032323-2301021123010321-0322002101331331-1011212233121120-1302230100030020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301110021131020-0311223303010221-3013322111221100-3312311211121203-2100121132330321-2101112231022313-2330321222232023-0003131312203103"></a>

## ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies — forward_proxy_policies / 022310003112 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.active_forward_proxy_policies](resources--aws_vpc_site--reference--group-002.md#canonical-2032010023210310-0333120203113312-2103312311112232-3321203111222313-2231210132332030-2321303012031232-1312103001310131-1330122232003213)
- ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-3133323030310111-0320312110121121-1031201030031222-3232302312022200-3023203020323323-3330313302321120-1301302000130232-3123222233330131"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3130200111102330-1001013130321210-2122331010203212-0322013221332333-1103230330112031-0332112322221133-2012213203012212-0022002103323222"></a>

## Direct properties — forward_proxy_policies / 022310003112 / 3

<a id="canonical-3310001203321022-2200222100130001-0113321232300121-0330213023202102-1330113302020103-0101230121012230-3021311220223003-1020322313331033"></a>

<a id="canonical-1201212030023022-0312010113202332-2130000203301021-3200203220312020-2212030212111020-2133001212222033-2130233100032130-0303220330301122"></a>

## name property — forward_proxy_policies / 022310003112 / 4

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

<a id="canonical-3203031123212100-3210010003211122-2200302030331201-0000120112001333-3332011013311102-2103232012322033-3330320231212113-2002030111212333"></a>

<a id="canonical-1020313311223110-2221320331133131-2213132102020213-2102203313333121-3203021220323210-0113031000203020-2030323233202202-2102122010000103"></a>

## namespace property — forward_proxy_policies / 022310003112 / 5

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

<a id="canonical-3332011122133323-1213202010302110-0022102121220000-1332233211231313-1122330211300211-3320111233120132-2132221221333231-2203323021222330"></a>

<a id="canonical-0312130322111002-2302031232233330-0030323311223220-3312103122133220-0020013320120131-0121000130330232-1112010020132311-2222323022313013"></a>

## tenant property — forward_proxy_policies / 022310003112 / 6

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

<a id="canonical-1232103302001132-0212332203132302-2000003002121131-2113031022103233-3232200231130102-2202232230001130-1320120033123300-0021031302120022"></a>

## Next pages — forward_proxy_policies / 022310003112 / 7

- [ingress_egress_gw.active_forward_proxy_policies](resources--aws_vpc_site--reference--group-002.md#canonical-2032010023210310-0333120203113312-2103312311112232-3321203111222313-2231210132332030-2321303012031232-1312103001310131-1330122232003213)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2132102132020200-1223012023300303-2203302223002220-3033032020232223-2101123323233323-2302223031001002-0302203010321310-2113300121321300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1231230230233212-3302311011323103-2101231103022232-0202111023113303-2033312220312322-1323123333333103-2133130122230330-2332122102323000"></a>

## ingress_egress_gw.active_network_policies — active_network_policies / 133221313323 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- ingress_egress_gw.active_network_policies

<a id="canonical-2323300230300130-3102112000222001-0102133210002323-2112003012203301-2020102230032112-2302301322310200-3230132221103101-2010111333001123"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for active network policies.

Upstream description:

List of firewall policy views.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2310303233103131-0101122322332301-1331313200112230-3010331103000232-2120003312321112-1311031303332130-2103330100210001-0211002331122200"></a>

## Direct properties — active_network_policies / 133221313323 / 3

- [network_policies](resources--aws_vpc_site--reference--group-002.md#canonical-2021301202033300-0011233110303100-2212322032023112-0122003112211001-0000130103103203-2102131232210112-0011030200312010-2033003001110113): complete subsection reference.

<a id="canonical-2313000333100131-3231112001300313-1320232123013033-3220302103300000-3322110232010023-2123111022010310-1330100331213332-3202032203110000"></a>

## Next pages — active_network_policies / 133221313323 / 4

- [ingress_egress_gw.active_network_policies.network_policies](resources--aws_vpc_site--reference--group-002.md#canonical-2021301202033300-0011233110303100-2212322032023112-0122003112211001-0000130103103203-2102131232210112-0011030200312010-2033003001110113)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2021301202033300-0011233110303100-2212322032023112-0122003112211001-0000130103103203-2102131232210112-0011030200312010-2033003001110113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213201201000121-0230000202122331-2221203133021000-3220310123133100-0313110001203310-1012232332012130-2331230323212101-1123201101113123"></a>

## ingress_egress_gw.active_network_policies.network_policies — network_policies / 101222020133 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.active_network_policies](resources--aws_vpc_site--reference--group-002.md#canonical-2132102132020200-1223012023300303-2203302223002220-3033032020232223-2101123323233323-2302223031001002-0302203010321310-2113300121321300)
- ingress_egress_gw.active_network_policies.network_policies

<a id="canonical-3030020211202331-0131130333002003-0102121113023322-0313013112312121-2133322032323220-2013032023030033-2033233030312321-1111332313330133"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Firewall Policies active for this network firewall.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2130232313131120-2112323100303101-1311020112023021-1101222030323301-1203223301031233-3133330100311001-3113311010111321-1203020221132020"></a>

## Direct properties — network_policies / 101222020133 / 3

<a id="canonical-3002021112223022-0233102233201312-2111132021333202-3020321223230210-1230130220230231-2232310023103202-3033013212221022-2231321021020221"></a>

<a id="canonical-1103131110200310-3002131000001311-1102000001020232-1201310332231001-0213322203003102-3031332202223002-1021130023212011-2300002211332202"></a>

## name property — network_policies / 101222020133 / 4

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

<a id="canonical-3133101201022212-3030020200021221-3213202021011103-3132111110321031-3031323201213003-3210301133232013-0033222102113113-2013111033123322"></a>

<a id="canonical-3211213131031033-0110133311101032-0030200032223211-2331020030123121-2003033222223332-0312313320321202-0132302331221113-2203001131022220"></a>

## namespace property — network_policies / 101222020133 / 5

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

<a id="canonical-2132131332223200-3231103323221132-2310113020312232-2011111123230303-3002013330131300-0322311213033001-0202220310310323-3001011133103223"></a>

<a id="canonical-1313333111222031-3221112130002303-1122231113213012-2231000102213311-3203010230033012-3223011313131122-3031120321331310-3201323130331302"></a>

## tenant property — network_policies / 101222020133 / 6

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

<a id="canonical-1033003011022213-3310321222123220-2221212131132212-2221111123232211-2122012100030113-0321002202332221-1003322333120303-1321231230010133"></a>

## Next pages — network_policies / 101222020133 / 7

- [ingress_egress_gw.active_network_policies](resources--aws_vpc_site--reference--group-002.md#canonical-2132102132020200-1223012023300303-2203302223002220-3033032020232223-2101123323233323-2302223031001002-0302203010321310-2113300121321300)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0200133132302023-1321231123110331-1303133003320320-2302223311301000-0123113020122311-1231122311003302-3211310321102310-0123330100231003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311210221120011-1113113302300113-2303030331202113-0131133202111021-0103330132213023-2320303120101022-2210103233111333-0123230230303320"></a>

## ingress_egress_gw.allowed_vip_port — allowed_vip_port / 223201130321 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- ingress_egress_gw.allowed_vip_port

<a id="canonical-2031303023331303-1230210220212332-3122131021131200-3023120011030231-1132313303133330-2112021000313003-2300313123322013-1210120111122002"></a>

Type: `"object"`. single nested block, Optional.

Defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use
the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Upstream description:

This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client
can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_ports",
    "disable_allowed_vip_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_http_https_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_http_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_https_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_http_https_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_http_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_https_port"),
  validators.ConflictingObjectAttributes("use_http_https_port",
    "use_http_port"),
  validators.ConflictingObjectAttributes("use_http_https_port",
    "use_https_port"),
  validators.ConflictingObjectAttributes("use_http_port",
    "use_https_port")}
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
  "x-ves-oneof-field-port_choice": "[\"custom_ports\",\"disable_allowed_vip_port\",\"use_http_https_port\",\"use_http_port\",\"use_https_port\"]"
}
```

Terraform syntax:

```terraform
allowed_vip_port {
  # Configure direct properties listed below.
}
```

<a id="canonical-1030133212131200-0300122020310101-2121312202330220-2011002002210133-3230100332030231-0223011121130123-2231012312113232-0222130121003122"></a>

## Direct properties — allowed_vip_port / 223201130321 / 3

- [custom_ports](resources--aws_vpc_site--reference--group-002.md#canonical-2331003330203100-0132312230000030-0132230132133211-1130133323013323-0102201101011310-2230311030001002-0320321131303313-2111312211113223): complete subsection reference.

- [disable_allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-3201031200031002-0312131232213230-2013103031131333-2333230121220222-3311011013302133-2302200201030033-1331102230320110-3311320031022332): complete subsection reference.

- [use_http_https_port](resources--aws_vpc_site--reference--group-002.md#canonical-3232200031312120-1132310320103333-3131120230200113-3120311102021132-1010231221210302-3102310131321130-1032100221233200-0302322030200310): complete subsection reference.

- [use_http_port](resources--aws_vpc_site--reference--group-002.md#canonical-2033132201323033-1331232012022301-2131312232221001-2012111202211200-2220132113201330-0122000101102302-1113233113020230-2330031301310333): complete subsection reference.

- [use_https_port](resources--aws_vpc_site--reference--group-002.md#canonical-3121101102110133-3103111202111010-2200011313000113-2302111201131323-3330111310112123-3110223103221320-2203231321331232-0103101223320201): complete subsection reference.

<a id="canonical-1003131231113220-0333111230232132-3230130230203020-0133110301103332-0200113313210301-1031331320230222-1031121211203022-3030021122003323"></a>

## Next pages — allowed_vip_port / 223201130321 / 4

- [ingress_egress_gw.allowed_vip_port.custom_ports](resources--aws_vpc_site--reference--group-002.md#canonical-2331003330203100-0132312230000030-0132230132133211-1130133323013323-0102201101011310-2230311030001002-0320321131303313-2111312211113223)
- [ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-3201031200031002-0312131232213230-2013103031131333-2333230121220222-3311011013302133-2302200201030033-1331102230320110-3311320031022332)
- [ingress_egress_gw.allowed_vip_port.use_http_https_port](resources--aws_vpc_site--reference--group-002.md#canonical-3232200031312120-1132310320103333-3131120230200113-3120311102021132-1010231221210302-3102310131321130-1032100221233200-0302322030200310)
- [ingress_egress_gw.allowed_vip_port.use_http_port](resources--aws_vpc_site--reference--group-002.md#canonical-2033132201323033-1331232012022301-2131312232221001-2012111202211200-2220132113201330-0122000101102302-1113233113020230-2330031301310333)
- [ingress_egress_gw.allowed_vip_port.use_https_port](resources--aws_vpc_site--reference--group-002.md#canonical-3121101102110133-3103111202111010-2200011313000113-2302111201131323-3330111310112123-3110223103221320-2203231321331232-0103101223320201)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2331003330203100-0132312230000030-0132230132133211-1130133323013323-0102201101011310-2230311030001002-0320321131303313-2111312211113223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023013013301130-1013023330010032-1123331313203102-0303210302200013-1212130023200133-2001201302213312-3100020103131130-1321301111222000"></a>

## ingress_egress_gw.allowed_vip_port.custom_ports — custom_ports / 310303311212 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-0200133132302023-1321231123110331-1303133003320320-2302223311301000-0123113020122311-1231122311003302-3211310321102310-0123330100231003)
- ingress_egress_gw.allowed_vip_port.custom_ports

<a id="canonical-3312033321200302-3000333213300130-1332120100200300-3110101022001230-1120330003233030-2101103300311322-3331112002012123-1310032003231323"></a>

Type: `"object"`. single nested block, Optional.

Custom Ports. List of Custom port.

Upstream description:

List of Custom port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("port_ranges")}
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
custom_ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-2031311101303311-1313330133121030-1201203210203031-3332220010000203-2020123132330311-0033230030300001-1333232001131133-0210113312300022"></a>

## Direct properties — custom_ports / 310303311212 / 3

<a id="canonical-0023311212111330-2302201313301202-2212132330320203-0212220300303000-0030321111113320-3000320302033230-3131122321232211-3313132113133112"></a>

<a id="canonical-0222110201321323-3122030031012013-1123331303232101-1313020111122210-3331202330023312-1133212133211002-3300002332212133-3011332133212103"></a>

## port_ranges property — custom_ports / 310303311212 / 4

Type: `"string"`. Optional.

Port Ranges. Port Ranges.

Upstream description:

Port Ranges.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  }
}
```

<a id="canonical-2200133310331220-2331111330113032-3121203003013322-0130332032213122-2030012223102200-1123132102223023-3203110110223022-2310221000003131"></a>

## Next pages — custom_ports / 310303311212 / 5

- [ingress_egress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-0200133132302023-1321231123110331-1303133003320320-2302223311301000-0123113020122311-1231122311003302-3211310321102310-0123330100231003)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3201031200031002-0312131232213230-2013103031131333-2333230121220222-3311011013302133-2302200201030033-1331102230320110-3311320031022332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033332111132221-0230233210301112-2231121330223200-1301002110013321-0211133212003023-1223032122210010-3122100003121233-2212032102022103"></a>

## ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port — disable_allowed_vip_port / 311312023321 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-0200133132302023-1321231123110331-1303133003320320-2302223311301000-0123113020122311-1231122311003302-3211310321102310-0123330100231003)
- ingress_egress_gw.allowed_vip_port.disable_allowed_vip_port

<a id="canonical-2300002120112230-3023300032230002-2300200211011203-0121333110122100-1023030012201331-2322232123320001-1121132113222132-1021310212112023"></a>

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
disable_allowed_vip_port = {}
```

<a id="canonical-1021101013131102-1321102132323303-2303030310211131-0010203131221100-1221231301312133-0212101121123110-3003013202030012-3303031013002011"></a>

## Direct properties — disable_allowed_vip_port / 311312023321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0203113200323233-2010113122021132-0111011032121033-3002131003320123-0020303312323211-0301333121311331-1131022110323203-2022210301133032"></a>

## Next pages — disable_allowed_vip_port / 311312023321 / 4

- [ingress_egress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-0200133132302023-1321231123110331-1303133003320320-2302223311301000-0123113020122311-1231122311003302-3211310321102310-0123330100231003)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3232200031312120-1132310320103333-3131120230200113-3120311102021132-1010231221210302-3102310131321130-1032100221233200-0302322030200310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301201020120012-2013111202200233-3012010303320101-2301220203313323-1210312113212033-2022333012322100-2012132230321132-3302101233112011"></a>

## ingress_egress_gw.allowed_vip_port.use_http_https_port — use_http_https_port / 012030330131 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-0200133132302023-1321231123110331-1303133003320320-2302223311301000-0123113020122311-1231122311003302-3211310321102310-0123330100231003)
- ingress_egress_gw.allowed_vip_port.use_http_https_port

<a id="canonical-2203003213120012-3310030112101211-3020020222012331-0231010332313120-0021120330122321-3132310103123303-2120210112303211-0222320110302001"></a>

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
use_http_https_port = {}
```

<a id="canonical-2132212001121100-1301120202022002-2301310121033230-3031000002102331-1123002320123201-3113300302313023-1321100220210121-1301332031123313"></a>

## Direct properties — use_http_https_port / 012030330131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1320021221221101-2223100221222233-1123310203113310-1123322333233112-2023003123000302-0330202333023323-1030113323102223-3220232003202323"></a>

## Next pages — use_http_https_port / 012030330131 / 4

- [ingress_egress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-0200133132302023-1321231123110331-1303133003320320-2302223311301000-0123113020122311-1231122311003302-3211310321102310-0123330100231003)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2033132201323033-1331232012022301-2131312232221001-2012111202211200-2220132113201330-0122000101102302-1113233113020230-2330031301310333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030132223233212-0130323230310220-0002101001220233-0033311201110010-0302301023121113-0311022333032223-1210212022302121-3001220122001332"></a>

## ingress_egress_gw.allowed_vip_port.use_http_port — use_http_port / 002331210000 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-0200133132302023-1321231123110331-1303133003320320-2302223311301000-0123113020122311-1231122311003302-3211310321102310-0123330100231003)
- ingress_egress_gw.allowed_vip_port.use_http_port

<a id="canonical-1001031022320320-3002300120223201-0210200302130202-3223221230202313-3300001231032312-2020100332220002-0001030322002300-3010213011022113"></a>

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
use_http_port = {}
```

<a id="canonical-3212122202023233-1001321202302131-0101311232103011-2323033133002013-0331232323231331-1200101110232112-2311203230032011-2303011320102023"></a>

## Direct properties — use_http_port / 002331210000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1300202012100123-0123100132312312-2121023210012031-0322201113231020-3031100310011212-0103321130031213-3022113033230203-3203000030002210"></a>

## Next pages — use_http_port / 002331210000 / 4

- [ingress_egress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-0200133132302023-1321231123110331-1303133003320320-2302223311301000-0123113020122311-1231122311003302-3211310321102310-0123330100231003)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3121101102110133-3103111202111010-2200011313000113-2302111201131323-3330111310112123-3110223103221320-2203231321331232-0103101223320201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0230212023302100-3323121131222120-2211031032033031-2232313023133203-2210132202330212-0000001220230200-1101010110310121-3221220323331120"></a>

## ingress_egress_gw.allowed_vip_port.use_https_port — use_https_port / 000230233101 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-0200133132302023-1321231123110331-1303133003320320-2302223311301000-0123113020122311-1231122311003302-3211310321102310-0123330100231003)
- ingress_egress_gw.allowed_vip_port.use_https_port

<a id="canonical-3101200230133122-2203333223232300-1013032311113120-3220003120222331-3320313201121322-3102000200130103-0330111112033320-0302022202321323"></a>

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
use_https_port = {}
```

<a id="canonical-1001332332311210-3112313011023210-3132122223022311-2312231023313132-1033020333122330-2022320312231022-0100020202320310-1232112210013110"></a>

## Direct properties — use_https_port / 000230233101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0213321131230131-1200122033021301-1013122230032100-2230021331113320-2012303311102013-0220133330010101-2022001233311202-0320322302322002"></a>

## Next pages — use_https_port / 000230233101 / 4

- [ingress_egress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-0200133132302023-1321231123110331-1303133003320320-2302223311301000-0123113020122311-1231122311003302-3211310321102310-0123330100231003)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0321200102110013-2221113222102210-3011001233023112-0002332023300211-0312212112213020-1212311131213112-3011010202130131-3002322100221021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013230203022003-3202213233211100-2310113001222003-1022232321221302-0001301013033223-2011213310333301-1232210121221122-2103010123101120"></a>

## ingress_egress_gw.allowed_vip_port_sli — allowed_vip_port_sli / 301031120100 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- ingress_egress_gw.allowed_vip_port_sli

<a id="canonical-0121103020231110-0333011313232031-3301113121232323-1031031120002011-2102232000003001-3221132200232012-1122303210321001-2110002002201103"></a>

Type: `"object"`. single nested block, Optional.

Defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use
the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Upstream description:

This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client
can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_ports",
    "disable_allowed_vip_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_http_https_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_http_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_https_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_http_https_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_http_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_https_port"),
  validators.ConflictingObjectAttributes("use_http_https_port",
    "use_http_port"),
  validators.ConflictingObjectAttributes("use_http_https_port",
    "use_https_port"),
  validators.ConflictingObjectAttributes("use_http_port",
    "use_https_port")}
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
  "x-ves-oneof-field-port_choice": "[\"custom_ports\",\"disable_allowed_vip_port\",\"use_http_https_port\",\"use_http_port\",\"use_https_port\"]"
}
```

Terraform syntax:

```terraform
allowed_vip_port_sli {
  # Configure direct properties listed below.
}
```

<a id="canonical-0231023130002131-3022000310222210-3123013331121330-2203101011313111-2313200020132320-2213211220032210-1121120023111322-2111230003312002"></a>

## Direct properties — allowed_vip_port_sli / 301031120100 / 3

- [custom_ports](resources--aws_vpc_site--reference--group-002.md#canonical-3122133222012102-0021200231323322-1120133023230031-3030110022211130-1301323331203130-3322220112203132-0100123003110132-2023111212111132): complete subsection reference.

- [disable_allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-1311301001222100-2131312201022222-0030100003233322-3330123303231212-3231331201222013-0330213301102211-1223103130030022-2232013120322100): complete subsection reference.

- [use_http_https_port](resources--aws_vpc_site--reference--group-002.md#canonical-2301010232200032-1303333023212100-2133031000303202-1011320313101323-3022110320300302-0223310210021230-2323323202110203-2303302213332001): complete subsection reference.

- [use_http_port](resources--aws_vpc_site--reference--group-002.md#canonical-3200332021010020-0232203023123022-1120331002123013-2122233123300332-2310323101322223-3033312332033320-3230030212311210-3103110210232302): complete subsection reference.

- [use_https_port](resources--aws_vpc_site--reference--group-002.md#canonical-0112131330002200-2310221213121231-0333002102233220-1013001321032221-2101311231033211-2312100031210021-1031213302121300-1103330230010121): complete subsection reference.

<a id="canonical-1123012210233100-3132330122221021-1313202212112030-2233323320220330-1210221010103301-1230220321102100-2211130331131033-1222132011301121"></a>

## Next pages — allowed_vip_port_sli / 301031120100 / 4

- [ingress_egress_gw.allowed_vip_port_sli.custom_ports](resources--aws_vpc_site--reference--group-002.md#canonical-3122133222012102-0021200231323322-1120133023230031-3030110022211130-1301323331203130-3322220112203132-0100123003110132-2023111212111132)
- [ingress_egress_gw.allowed_vip_port_sli.disable_allowed_vip_port](resources--aws_vpc_site--reference--group-002.md#canonical-1311301001222100-2131312201022222-0030100003233322-3330123303231212-3231331201222013-0330213301102211-1223103130030022-2232013120322100)
- [ingress_egress_gw.allowed_vip_port_sli.use_http_https_port](resources--aws_vpc_site--reference--group-002.md#canonical-2301010232200032-1303333023212100-2133031000303202-1011320313101323-3022110320300302-0223310210021230-2323323202110203-2303302213332001)
- [ingress_egress_gw.allowed_vip_port_sli.use_http_port](resources--aws_vpc_site--reference--group-002.md#canonical-3200332021010020-0232203023123022-1120331002123013-2122233123300332-2310323101322223-3033312332033320-3230030212311210-3103110210232302)
- [ingress_egress_gw.allowed_vip_port_sli.use_https_port](resources--aws_vpc_site--reference--group-002.md#canonical-0112131330002200-2310221213121231-0333002102233220-1013001321032221-2101311231033211-2312100031210021-1031213302121300-1103330230010121)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3122133222012102-0021200231323322-1120133023230031-3030110022211130-1301323331203130-3322220112203132-0100123003110132-2023111212111132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320201320203001-1200033302321303-1113200300121131-2000032200100202-2320331123123020-2012203010321101-1001023120301333-3203102212333010"></a>

## ingress_egress_gw.allowed_vip_port_sli.custom_ports — custom_ports / 011303232310 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.allowed_vip_port_sli](resources--aws_vpc_site--reference--group-002.md#canonical-0321200102110013-2221113222102210-3011001233023112-0002332023300211-0312212112213020-1212311131213112-3011010202130131-3002322100221021)
- ingress_egress_gw.allowed_vip_port_sli.custom_ports

<a id="canonical-3013102000233231-2213331300300003-3031112213300032-3321021333103022-2001310001310221-1122132221133213-2232013201232120-0100103331200013"></a>

Type: `"object"`. single nested block, Optional.

Custom Ports. List of Custom port.

Upstream description:

List of Custom port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("port_ranges")}
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
custom_ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-0311001322030331-2002202121323130-3321101121131202-1030011121132110-1330310203030311-1233311012022023-3101011323320030-0302303221300012"></a>

## Direct properties — custom_ports / 011303232310 / 3

<a id="canonical-1221211132120223-1212111231021312-1311333212202121-0032200333020321-1232322231022311-3003122323133102-0212100310232013-2330000122213011"></a>

<a id="canonical-0123120302002130-1033031302110103-3012312301301333-3333021220133210-2003231201321001-0312310022133130-2120331033032223-1133300020101102"></a>

## port_ranges property — custom_ports / 011303232310 / 4

Type: `"string"`. Optional.

Port Ranges. Port Ranges.

Upstream description:

Port Ranges.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

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
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  }
}
```

<a id="canonical-2310121030323302-2020000213330131-1010313232200301-2303333213212331-1212232120131023-0123213130222031-2221031003113202-3112231110011013"></a>

## Next pages — custom_ports / 011303232310 / 5

- [ingress_egress_gw.allowed_vip_port_sli](resources--aws_vpc_site--reference--group-002.md#canonical-0321200102110013-2221113222102210-3011001233023112-0002332023300211-0312212112213020-1212311131213112-3011010202130131-3002322100221021)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1311301001222100-2131312201022222-0030100003233322-3330123303231212-3231331201222013-0330213301102211-1223103130030022-2232013120322100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220200012021222-2302231102122321-2110321232211003-1320003320130310-0331131313320102-3302100130122203-3301002100123313-0213321010130213"></a>

## ingress_egress_gw.allowed_vip_port_sli.disable_allowed_vip_port — disable_allowed_vip_port / 130210322201 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.allowed_vip_port_sli](resources--aws_vpc_site--reference--group-002.md#canonical-0321200102110013-2221113222102210-3011001233023112-0002332023300211-0312212112213020-1212311131213112-3011010202130131-3002322100221021)
- ingress_egress_gw.allowed_vip_port_sli.disable_allowed_vip_port

<a id="canonical-3231312210113133-0232113002203113-2011203121101312-1123020111302232-0120221103332012-1011102322200022-0123120110311103-3211103113032200"></a>

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
disable_allowed_vip_port = {}
```

<a id="canonical-2321023313300222-1300301323330212-2131210101302112-2201223010123102-3311032300030112-1323030201330131-0112102010331100-0201000020030022"></a>

## Direct properties — disable_allowed_vip_port / 130210322201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0202332201000230-0310000211123120-0320233122131211-1123032103322323-2321032001300132-1120122203221311-2012013131210211-3233133320233203"></a>

## Next pages — disable_allowed_vip_port / 130210322201 / 4

- [ingress_egress_gw.allowed_vip_port_sli](resources--aws_vpc_site--reference--group-002.md#canonical-0321200102110013-2221113222102210-3011001233023112-0002332023300211-0312212112213020-1212311131213112-3011010202130131-3002322100221021)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2301010232200032-1303333023212100-2133031000303202-1011320313101323-3022110320300302-0223310210021230-2323323202110203-2303302213332001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213301301000113-0220133130302000-3210131000023213-1102313312300120-3333232000033330-2123200322130113-0301133332130102-0212301021320030"></a>

## ingress_egress_gw.allowed_vip_port_sli.use_http_https_port — use_http_https_port / 223101010320 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.allowed_vip_port_sli](resources--aws_vpc_site--reference--group-002.md#canonical-0321200102110013-2221113222102210-3011001233023112-0002332023300211-0312212112213020-1212311131213112-3011010202130131-3002322100221021)
- ingress_egress_gw.allowed_vip_port_sli.use_http_https_port

<a id="canonical-0022002012310112-2133310302132230-1321130111201122-3300303330121122-2012231121333100-2103312010330013-3012211333130312-1030213300213010"></a>

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
use_http_https_port = {}
```

<a id="canonical-3210133213131332-3302013203213100-1003002211111021-0130332222013030-2232030321203313-3020021000001313-2212011003321100-1101131022312213"></a>

## Direct properties — use_http_https_port / 223101010320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1122002201310133-1221201002003023-2220011210333300-0231213130311302-1323303231100200-0020333320101233-0022021110003321-1233220212333030"></a>

## Next pages — use_http_https_port / 223101010320 / 4

- [ingress_egress_gw.allowed_vip_port_sli](resources--aws_vpc_site--reference--group-002.md#canonical-0321200102110013-2221113222102210-3011001233023112-0002332023300211-0312212112213020-1212311131213112-3011010202130131-3002322100221021)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3200332021010020-0232203023123022-1120331002123013-2122233123300332-2310323101322223-3033312332033320-3230030212311210-3103110210232302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121312110003132-2111033311330301-2312202333200023-1000223301301101-0303321100111331-2203011100311230-1321033011312332-2311213202311303"></a>

## ingress_egress_gw.allowed_vip_port_sli.use_http_port — use_http_port / 002112223210 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.allowed_vip_port_sli](resources--aws_vpc_site--reference--group-002.md#canonical-0321200102110013-2221113222102210-3011001233023112-0002332023300211-0312212112213020-1212311131213112-3011010202130131-3002322100221021)
- ingress_egress_gw.allowed_vip_port_sli.use_http_port

<a id="canonical-1323012030210213-2201001130002012-3023122230021002-0131122023020110-2022320330211321-1023020112303031-1333030111330101-3220123331211120"></a>

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
use_http_port = {}
```

<a id="canonical-3201101010113333-3301112312202130-3103322120330233-3302102013223231-3231100211313102-0031123210332131-3201322000333231-0231212200031333"></a>

## Direct properties — use_http_port / 002112223210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0020333003301222-3002320230230123-3231313302230300-2212332222313131-2230333102021101-3033002320122210-3030132302332311-2110332323202001"></a>

## Next pages — use_http_port / 002112223210 / 4

- [ingress_egress_gw.allowed_vip_port_sli](resources--aws_vpc_site--reference--group-002.md#canonical-0321200102110013-2221113222102210-3011001233023112-0002332023300211-0312212112213020-1212311131213112-3011010202130131-3002322100221021)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0112131330002200-2310221213121231-0333002102233220-1013001321032221-2101311231033211-2312100031210021-1031213302121300-1103330230010121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023023102331331-3012002023323231-0020333222311030-1330232233101111-0333030330210323-2233003120333002-2232001000231212-0132133322111202"></a>

## ingress_egress_gw.allowed_vip_port_sli.use_https_port — use_https_port / 011211321213 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.allowed_vip_port_sli](resources--aws_vpc_site--reference--group-002.md#canonical-0321200102110013-2221113222102210-3011001233023112-0002332023300211-0312212112213020-1212311131213112-3011010202130131-3002322100221021)
- ingress_egress_gw.allowed_vip_port_sli.use_https_port

<a id="canonical-2310132321102031-0202112311103322-1002022033003130-0022201132101101-2101230022311033-0110102323003000-1102222013000323-3111331010330131"></a>

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
use_https_port = {}
```

<a id="canonical-0333313010112322-1110013120012101-1223203220103231-1011033213121300-3312132021300001-2111203022223103-3330010123302131-3110233000232121"></a>

## Direct properties — use_https_port / 011211321213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2131123312331103-3012130311200123-0003303003031310-3023313203230010-1002313332002030-1301113101110221-1220223110130013-0302303010131213"></a>

## Next pages — use_https_port / 011211321213 / 4

- [ingress_egress_gw.allowed_vip_port_sli](resources--aws_vpc_site--reference--group-002.md#canonical-0321200102110013-2221113222102210-3011001233023112-0002332023300211-0312212112213020-1212311131213112-3011010202130131-3002322100221021)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3233231103113313-1010330103013321-3013311222003130-2100031033021221-1113023133322203-0220023312322322-3302003000213032-2123300021210202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121030021302303-1111301212121212-1210213031010102-2301000302330230-1003000212213000-1231102030222121-1330313033331102-2200302001202301"></a>

## ingress_egress_gw.az_nodes — az_nodes / 003233201030 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- ingress_egress_gw.az_nodes

<a id="canonical-3003133301012313-3032033301300111-1200200203021020-2112002110210322-1000310013300202-1311202321210312-0323123230330211-0322312123021303"></a>

Type: `"object"`. list nested block, Optional.

Only Single AZ or Three AZ(s) nodes are supported currently.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("aws_az_name"),
  validators.ConflictingListObjectAttributes("inside_subnet",
    "reserved_inside_subnet")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1,3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1,3"
  }
}
```

Terraform syntax:

```terraform
az_nodes {
  # Configure direct properties listed below.
}
```

<a id="canonical-0210222320111103-1111100332122303-0213213332010221-1312021010023013-1220011021032033-1100110023112130-3013220123113112-0302301212133202"></a>

## Direct properties — az_nodes / 003233201030 / 3

<a id="canonical-1321312020101020-3231202032331200-3223012231031213-2000113231220321-0211320020131023-0202021032131033-3300013002202101-1120201221121110"></a>

<a id="canonical-3122103210221112-2100302322130320-3322310221231213-3031230303122100-3211121320312220-1113231220002031-2033333211121222-2023322331000323"></a>

## aws_az_name property — az_nodes / 003233201030 / 4

Type: `"string"`. Optional.

AWS availability zone, must be consistent with the selected AWS region.

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

- [inside_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-3333313031011133-1311220301031012-1031202100120003-0200332010120102-3113033211331030-3301100301001212-0320011313120011-0021220301020301): complete subsection reference.

- [outside_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-1200131200113202-2321312020132300-3212113021130221-2212021110301113-2130302030311311-1112320201300033-2130133331111212-1032332330112001): complete subsection reference.

- [reserved_inside_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-2322033231033311-1111320233013230-1332133031312332-2232312031020113-3311312102121102-1121123312131111-2321203203100212-0030120003131202): complete subsection reference.

- [workload_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-2202302230110211-0330310033231000-3201133201230223-0131122332313033-2311203211103010-3213033331200021-3201203101030303-1103110033130130): complete subsection reference.

<a id="canonical-2002333302321221-1211333313033003-3320110232311332-0311301300210202-0123033323231300-2203321303231030-3311011030201002-0321111303221221"></a>

## Next pages — az_nodes / 003233201030 / 5

- [ingress_egress_gw.az_nodes.inside_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-3333313031011133-1311220301031012-1031202100120003-0200332010120102-3113033211331030-3301100301001212-0320011313120011-0021220301020301)
- [ingress_egress_gw.az_nodes.outside_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-1200131200113202-2321312020132300-3212113021130221-2212021110301113-2130302030311311-1112320201300033-2130133331111212-1032332330112001)
- [ingress_egress_gw.az_nodes.reserved_inside_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-2322033231033311-1111320233013230-1332133031312332-2232312031020113-3311312102121102-1121123312131111-2321203203100212-0030120003131202)
- [ingress_egress_gw.az_nodes.workload_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-2202302230110211-0330310033231000-3201133201230223-0131122332313033-2311203211103010-3213033331200021-3201203101030303-1103110033130130)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3333313031011133-1311220301031012-1031202100120003-0200332010120102-3113033211331030-3301100301001212-0320011313120011-0021220301020301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0021100001203131-1320330011102213-0310033200331011-2332031221330013-2212331223031311-2030133100111330-0203122332100121-2300321303212120"></a>

## ingress_egress_gw.az_nodes.inside_subnet — inside_subnet / 112011223203 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.az_nodes](resources--aws_vpc_site--reference--group-002.md#canonical-3233231103113313-1010330103013321-3013311222003130-2100031033021221-1113023133322203-0220023312322322-3302003000213032-2123300021210202)
- ingress_egress_gw.az_nodes.inside_subnet

<a id="canonical-0233131333123222-0113223302223232-0112013031113333-2311110201213111-2122001020033100-1101102220122231-1030223332220231-1213021012033203"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for inside subnet.

Upstream description:

Parameters for AWS subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_subnet_id",
    "subnet_param")}
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
  "x-ves-oneof-field-choice": "[\"existing_subnet_id\",\"subnet_param\"]"
}
```

Terraform syntax:

```terraform
inside_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-1322010301201023-3011001131132312-0321320302311102-3233002033320201-2000320110320210-2213033022211332-3320022330031203-0130020133112203"></a>

## Direct properties — inside_subnet / 112011223203 / 3

<a id="canonical-0011210001122133-0200201101003023-1133030321211220-2120202333210321-0130010230023221-2321220102201321-1303323220023121-3103212231323223"></a>

<a id="canonical-3222303200202331-2231103313223211-2112201011123103-2201111132321321-0230230003133102-2332000231112113-2312332102002001-0012022221030223"></a>

## existing_subnet_id property — inside_subnet / 112011223203 / 4

Type: `"string"`. Optional.

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Upstream description:

Exclusive with \[subnet\_param\] Information about existing subnet ID.

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
    },
    "pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

- [subnet_param](resources--aws_vpc_site--reference--group-002.md#canonical-0022223021120122-1132023300112000-0302312311211020-3313203131202122-1233201202001102-0323020203103133-1332012212301021-0110232130101313): complete subsection reference.

<a id="canonical-3303322003231101-3203112210220332-0310021333333313-3111021011310121-1032023230130000-1312203320212013-0031311201202110-3331133120001033"></a>

## Next pages — inside_subnet / 112011223203 / 5

- [ingress_egress_gw.az_nodes.inside_subnet.subnet_param](resources--aws_vpc_site--reference--group-002.md#canonical-0022223021120122-1132023300112000-0302312311211020-3313203131202122-1233201202001102-0323020203103133-1332012212301021-0110232130101313)
- [ingress_egress_gw.az_nodes](resources--aws_vpc_site--reference--group-002.md#canonical-3233231103113313-1010330103013321-3013311222003130-2100031033021221-1113023133322203-0220023312322322-3302003000213032-2123300021210202)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0022223021120122-1132023300112000-0302312311211020-3313203131202122-1233201202001102-0323020203103133-1332012212301021-0110232130101313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313030012203001-0002202021101113-3032101330300111-0133222231322320-2320202220133003-2302002021330311-1233320011111222-3121331100010022"></a>

## ingress_egress_gw.az_nodes.inside_subnet.subnet_param — subnet_param / 303313313131 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.az_nodes](resources--aws_vpc_site--reference--group-002.md#canonical-3233231103113313-1010330103013321-3013311222003130-2100031033021221-1113023133322203-0220023312322322-3302003000213032-2123300021210202)
- [ingress_egress_gw.az_nodes.inside_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-3333313031011133-1311220301031012-1031202100120003-0200332010120102-3113033211331030-3301100301001212-0320011313120011-0021220301020301)
- ingress_egress_gw.az_nodes.inside_subnet.subnet_param

<a id="canonical-3100022321220330-3133331002031211-2231200112000112-0130121030132002-2302122020102222-3221322200110103-2002021300020212-1110332201112333"></a>

Type: `"object"`. single nested block, Optional.

Parameters for creating a new cloud subnet.

Provider validators and defaults (from schema source):

```go
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
subnet_param {
  # Configure direct properties listed below.
}
```

<a id="canonical-3312320120010332-0331003023021320-0220300323202200-1200022312201210-0300010003332102-2113311213010023-1121103320332110-3033320320200102"></a>

## Direct properties — subnet_param / 303313313131 / 3

<a id="canonical-1111223303020210-2012330001231302-0130300212101131-0001210311310303-3110003100331232-0110211310003111-1221110232000131-1120332132213321"></a>

<a id="canonical-0302313001332233-0113020232323212-1112131213121100-3002122202332123-2013111330311311-2222031112133232-0121220021121121-0003203310031013"></a>

## IPv4 property — subnet_param / 303313313131 / 4

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

<a id="canonical-0012231101313210-2110111320120021-2310011231303022-2203013302331022-2132031021021330-1020120322301102-0212220011323132-0233032100300332"></a>

## Next pages — subnet_param / 303313313131 / 5

- [ingress_egress_gw.az_nodes.inside_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-3333313031011133-1311220301031012-1031202100120003-0200332010120102-3113033211331030-3301100301001212-0320011313120011-0021220301020301)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1200131200113202-2321312020132300-3212113021130221-2212021110301113-2130302030311311-1112320201300033-2130133331111212-1032332330112001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101023130220202-0002121232310030-3303022113032211-3011313311022001-1103332012131201-1132120322030210-0010221231032112-2223203231021321"></a>

## ingress_egress_gw.az_nodes.outside_subnet — outside_subnet / 223322131303 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.az_nodes](resources--aws_vpc_site--reference--group-002.md#canonical-3233231103113313-1010330103013321-3013311222003130-2100031033021221-1113023133322203-0220023312322322-3302003000213032-2123300021210202)
- ingress_egress_gw.az_nodes.outside_subnet

<a id="canonical-1322311123002012-2113223002130112-3112103212302111-0110332121010032-2010302300132020-3331023033201222-0022003132301333-0120211130202331"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for outside subnet.

Upstream description:

Parameters for AWS subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_subnet_id",
    "subnet_param")}
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
  "x-ves-oneof-field-choice": "[\"existing_subnet_id\",\"subnet_param\"]"
}
```

Terraform syntax:

```terraform
outside_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-2200333220310331-0222331330013230-0020320230100333-2310301120032212-0123312203110121-0013010031311002-3033030330333200-0130303313332032"></a>

## Direct properties — outside_subnet / 223322131303 / 3

<a id="canonical-0110120312333332-2131001012113313-0100212223013022-1003013220103010-2303111333021120-2031231300230300-1233010110232110-1122313010013222"></a>

<a id="canonical-0321222133232122-1300130112310011-1000033123221222-0130220322332132-0232300133301203-2020303230232302-3321302120003222-1300032032012111"></a>

## existing_subnet_id property — outside_subnet / 223322131303 / 4

Type: `"string"`. Optional.

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Upstream description:

Exclusive with \[subnet\_param\] Information about existing subnet ID.

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
    },
    "pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

- [subnet_param](resources--aws_vpc_site--reference--group-002.md#canonical-0320321210323023-2111321311303013-2313222232231100-0032211332213123-2120023032202302-0211303002131130-1030312212321312-0210033330210213): complete subsection reference.

<a id="canonical-1213320020113002-0122031123310311-0022221302130101-2033010131021100-1132121100200313-0212212221312202-3232210100231233-3132101101103313"></a>

## Next pages — outside_subnet / 223322131303 / 5

- [ingress_egress_gw.az_nodes.outside_subnet.subnet_param](resources--aws_vpc_site--reference--group-002.md#canonical-0320321210323023-2111321311303013-2313222232231100-0032211332213123-2120023032202302-0211303002131130-1030312212321312-0210033330210213)
- [ingress_egress_gw.az_nodes](resources--aws_vpc_site--reference--group-002.md#canonical-3233231103113313-1010330103013321-3013311222003130-2100031033021221-1113023133322203-0220023312322322-3302003000213032-2123300021210202)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0320321210323023-2111321311303013-2313222232231100-0032211332213123-2120023032202302-0211303002131130-1030312212321312-0210033330210213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102122211012212-2301101130102220-0130300021311103-2321202003212112-3222213320232122-1013000330112102-0103311212123131-3302203332331011"></a>

## ingress_egress_gw.az_nodes.outside_subnet.subnet_param — subnet_param / 300333033331 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.az_nodes](resources--aws_vpc_site--reference--group-002.md#canonical-3233231103113313-1010330103013321-3013311222003130-2100031033021221-1113023133322203-0220023312322322-3302003000213032-2123300021210202)
- [ingress_egress_gw.az_nodes.outside_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-1200131200113202-2321312020132300-3212113021130221-2212021110301113-2130302030311311-1112320201300033-2130133331111212-1032332330112001)
- ingress_egress_gw.az_nodes.outside_subnet.subnet_param

<a id="canonical-2130311332023001-1121031021323303-1312012120030203-0123001231131330-0303030012231000-2020023312032213-1103112301102123-3033120211203223"></a>

Type: `"object"`. single nested block, Optional.

Parameters for creating a new cloud subnet.

Provider validators and defaults (from schema source):

```go
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
subnet_param {
  # Configure direct properties listed below.
}
```

<a id="canonical-3320103121123001-0012322301032131-1132202321000210-3203103000122020-2102100023033303-2113132103020123-1013030032210130-3022221013032101"></a>

## Direct properties — subnet_param / 300333033331 / 3

<a id="canonical-3122001330131102-3011301230010322-3223111232233120-1130131220212022-0332203310210212-1130013011131211-1011231012103222-3233312103103322"></a>

<a id="canonical-1232203332211002-0113313321323121-2331200101111012-0313003213300211-0301030312101331-0312032012011203-3220310103021231-3330132033333323"></a>

## IPv4 property — subnet_param / 300333033331 / 4

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

<a id="canonical-1022331330313132-1331331133201101-3002330103223221-1311002003313033-3221113333212201-1222112213201322-0120211122030323-3221210333321030"></a>

## Next pages — subnet_param / 300333033331 / 5

- [ingress_egress_gw.az_nodes.outside_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-1200131200113202-2321312020132300-3212113021130221-2212021110301113-2130302030311311-1112320201300033-2130133331111212-1032332330112001)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2322033231033311-1111320233013230-1332133031312332-2232312031020113-3311312102121102-1121123312131111-2321203203100212-0030120003131202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122220100333321-0132210221303213-2123223232212300-3013331001120023-1302131231011210-3001103111212203-2313121200123100-3220003233002003"></a>

## ingress_egress_gw.az_nodes.reserved_inside_subnet — reserved_inside_subnet / 201311301300 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.az_nodes](resources--aws_vpc_site--reference--group-002.md#canonical-3233231103113313-1010330103013321-3013311222003130-2100031033021221-1113023133322203-0220023312322322-3302003000213032-2123300021210202)
- ingress_egress_gw.az_nodes.reserved_inside_subnet

<a id="canonical-1112231123312333-1231213100031111-3322312323033120-2120021303210323-1333013023021223-1233300120002131-1002000231311123-2111210311111233"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for reserved inside subnet.

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
reserved_inside_subnet = {}
```

<a id="canonical-3130220221221233-2322132130113003-1310010123201221-0121101332130011-1102011032103212-3122230023211312-3213111011311312-1113231021100320"></a>

## Direct properties — reserved_inside_subnet / 201311301300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3302232231132232-1310310303331122-2003300211101200-2330221003013330-3001222012323032-3122211321333212-1300203030133332-3020203220113202"></a>

## Next pages — reserved_inside_subnet / 201311301300 / 4

- [ingress_egress_gw.az_nodes](resources--aws_vpc_site--reference--group-002.md#canonical-3233231103113313-1010330103013321-3013311222003130-2100031033021221-1113023133322203-0220023312322322-3302003000213032-2123300021210202)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2202302230110211-0330310033231000-3201133201230223-0131122332313033-2311203211103010-3213033331200021-3201203101030303-1103110033130130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0200331010112301-0320113313313133-2223131321332231-1000322211133031-3223030013100221-0023333130200210-0013321001332203-3222000023132012"></a>

## ingress_egress_gw.az_nodes.workload_subnet — workload_subnet / 112110112100 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.az_nodes](resources--aws_vpc_site--reference--group-002.md#canonical-3233231103113313-1010330103013321-3013311222003130-2100031033021221-1113023133322203-0220023312322322-3302003000213032-2123300021210202)
- ingress_egress_gw.az_nodes.workload_subnet

<a id="canonical-3033001111311232-0120220011102100-2331133202223112-3230033012012220-3030101301011231-0112320123310133-2130021001002200-0222212122101300"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for workload subnet.

Upstream description:

Parameters for AWS subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_subnet_id",
    "subnet_param")}
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
  "x-ves-oneof-field-choice": "[\"existing_subnet_id\",\"subnet_param\"]"
}
```

Terraform syntax:

```terraform
workload_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-3213212222123010-0112210201023003-3000231302021332-1132321200122002-3112022333213032-0131012203312203-3111202123121003-3332232031212013"></a>

## Direct properties — workload_subnet / 112110112100 / 3

<a id="canonical-3122012230113233-0030311221100131-2133312013110013-2113222233312310-2323002323320320-0212112122120002-2302231030212030-2032311030331100"></a>

<a id="canonical-3012220113111222-3100311312022033-2012033200311330-3311313212230322-3201012130213201-3221323220010300-3221122100222123-3202022013012021"></a>

## existing_subnet_id property — workload_subnet / 112110112100 / 4

Type: `"string"`. Optional.

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Upstream description:

Exclusive with \[subnet\_param\] Information about existing subnet ID.

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
    },
    "pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

- [subnet_param](resources--aws_vpc_site--reference--group-002.md#canonical-1321220313100131-2032211330230213-3013111033232200-0002223021311032-0020222221312312-3303221300010002-3232131133130330-0033131101312120): complete subsection reference.

<a id="canonical-1310130122200203-0231102002010200-1321030212311322-3030223111030230-1300010112133211-3332332031110031-2220003333120013-0332122303122133"></a>

## Next pages — workload_subnet / 112110112100 / 5

- [ingress_egress_gw.az_nodes.workload_subnet.subnet_param](resources--aws_vpc_site--reference--group-002.md#canonical-1321220313100131-2032211330230213-3013111033232200-0002223021311032-0020222221312312-3303221300010002-3232131133130330-0033131101312120)
- [ingress_egress_gw.az_nodes](resources--aws_vpc_site--reference--group-002.md#canonical-3233231103113313-1010330103013321-3013311222003130-2100031033021221-1113023133322203-0220023312322322-3302003000213032-2123300021210202)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1321220313100131-2032211330230213-3013111033232200-0002223021311032-0020222221312312-3303221300010002-3232131133130330-0033131101312120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101302021130213-1003200233130100-3131200030201023-1113121210212222-3320013031131220-3122022201122332-1033121301320000-1211303032232220"></a>

## ingress_egress_gw.az_nodes.workload_subnet.subnet_param — subnet_param / 322133222200 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.az_nodes](resources--aws_vpc_site--reference--group-002.md#canonical-3233231103113313-1010330103013321-3013311222003130-2100031033021221-1113023133322203-0220023312322322-3302003000213032-2123300021210202)
- [ingress_egress_gw.az_nodes.workload_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-2202302230110211-0330310033231000-3201133201230223-0131122332313033-2311203211103010-3213033331200021-3201203101030303-1103110033130130)
- ingress_egress_gw.az_nodes.workload_subnet.subnet_param

<a id="canonical-3011023113110333-0102311101311231-1203312203110203-0320321030212301-3301332021031211-3113312211012011-3203131101030310-2100300233100232"></a>

Type: `"object"`. single nested block, Optional.

Parameters for creating a new cloud subnet.

Provider validators and defaults (from schema source):

```go
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
subnet_param {
  # Configure direct properties listed below.
}
```

<a id="canonical-3021103213332112-0221230122210210-1330333022113003-2100221232030031-1333313000200012-3311312132102212-2221330101332231-1311202002311022"></a>

## Direct properties — subnet_param / 322133222200 / 3

<a id="canonical-2333222121010331-0002230033111032-2311133323200020-3123331311031321-2001203233320133-1023200231132320-3211101111010221-1011132132121031"></a>

<a id="canonical-0202311000310301-2230200302320123-2202232321111312-2010012102011100-0221202030101203-0132230111310131-1133103302223302-2100122101112301"></a>

## IPv4 property — subnet_param / 322133222200 / 4

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

<a id="canonical-2220212331323011-2221030230331000-2310312002300322-1110320131321022-0321302211322112-1003132330002023-0113212031022220-0312311232122120"></a>

## Next pages — subnet_param / 322133222200 / 5

- [ingress_egress_gw.az_nodes.workload_subnet](resources--aws_vpc_site--reference--group-002.md#canonical-2202302230110211-0330310033231000-3201133201230223-0131122332313033-2311203211103010-3213033331200021-3201203101030303-1103110033130130)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1202110032001010-2121110311111232-0033110201121022-0202332122110012-3010302301220221-3302211103130233-2331332211132321-3101203131220211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2031103301232133-3131021122130223-1210221021213022-2130332102100102-1032112121312112-3302103020202321-0111032002012123-3002300131023211"></a>

## ingress_egress_gw.dc_cluster_group_inside_vn — dc_cluster_group_inside_vn / 000310021122 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- ingress_egress_gw.dc_cluster_group_inside_vn

<a id="canonical-1123030020110311-1001020101123103-3131130020110131-0310233132131131-0131333203030332-3003133030123321-3303120103131103-0303313112202111"></a>

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
dc_cluster_group_inside_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-2331122023230131-3010201020122231-1230213020030200-3131223112131110-2323032220220301-2321101320001100-0330031123113110-1223122001022101"></a>

## Direct properties — dc_cluster_group_inside_vn / 000310021122 / 3

<a id="canonical-1212211102121000-2031322321220113-2011120320122030-3110112131213113-0333112011013001-3330100131020231-3330301322013323-0333323111113123"></a>

<a id="canonical-1133011330031232-1333201033332203-2222012011322012-3111130102032302-0123111013023120-3321213231301013-3202201220323023-1231210121233121"></a>

## name property — dc_cluster_group_inside_vn / 000310021122 / 4

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

<a id="canonical-2012123330101223-3012201232002203-1313212112332021-1302230310220103-2222011313320303-3333000202121000-1031201223223033-3231020132102032"></a>

<a id="canonical-1213310312221330-1101303230100100-3213120303020203-1301013123013123-1113112012333122-2031133230210331-1233203311231230-0322200111210202"></a>

## namespace property — dc_cluster_group_inside_vn / 000310021122 / 5

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

<a id="canonical-0222313121302203-3011100013100212-2001023202113021-0300333012112303-0103332200103301-1001111133020032-1310212322011230-1031003312130321"></a>

<a id="canonical-2133300332212020-2311310032122220-3021030021013213-2013111331012123-1020130123011100-0020110100220221-0311113100201121-3313020322011222"></a>

## tenant property — dc_cluster_group_inside_vn / 000310021122 / 6

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

<a id="canonical-3132012122011301-0232122321303231-0000333133320202-1000220313021331-2311231110310031-3220002323203110-3130312210313101-3131311210002002"></a>

## Next pages — dc_cluster_group_inside_vn / 000310021122 / 7

- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0011011033131033-3121313210300020-2031133213011231-0312232322123203-2220013213011302-3230300023111123-3300201220331232-1211100120113031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102222321010311-0000322333021313-1220220130321011-1232322033100221-1222033101311310-1133003221011002-1200130100223333-0301300310200010"></a>

## ingress_egress_gw.dc_cluster_group_outside_vn — dc_cluster_group_outside_vn / 112112210122 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- ingress_egress_gw.dc_cluster_group_outside_vn

<a id="canonical-3101013212321113-3221311032212131-2002302112313131-3011121210013113-3221213110111000-3231322212003002-0102233101202100-2212111100223232"></a>

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
dc_cluster_group_outside_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-3103000010033220-2021010123112321-2000111030103002-1303203011333230-3103103232221211-3030022200200213-3022122231321301-0022132200000011"></a>

## Direct properties — dc_cluster_group_outside_vn / 112112210122 / 3

<a id="canonical-0232013021320202-3010330120331032-3030033300231000-2232302020321303-3002302230132122-2203211103203120-0320020032013320-3111001002010033"></a>

<a id="canonical-0211233311311013-1102030020102120-3023312332200330-3132112321120311-2110301231013002-1330203111122123-2220022230033033-1303103121311201"></a>

## name property — dc_cluster_group_outside_vn / 112112210122 / 4

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

<a id="canonical-1101122112313212-0301131231211021-0213120020113210-2122123201201123-2312231221223133-1131131331001030-2333103323300223-2312303003113132"></a>

<a id="canonical-0301211223331211-3201311232323033-3103320013222013-2120000311111123-0202301100201300-0120121322231233-2112113202320213-3232320213202301"></a>

## namespace property — dc_cluster_group_outside_vn / 112112210122 / 5

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

<a id="canonical-3222013300030021-3100321122022212-1111221130132320-0010300101310002-3233000200323111-0313020020022330-3112213100020302-0113330210232003"></a>

<a id="canonical-1112221202111231-3323011132033100-3033002213223002-3032203223131031-2303321000230321-1233010202130301-3100102101320121-3300200121113100"></a>

## tenant property — dc_cluster_group_outside_vn / 112112210122 / 6

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

<a id="canonical-2330011331331210-2321101203201033-3301011110030122-1122322112322033-3333122210113320-1301023032033003-2202323333113002-3311213120323222"></a>

## Next pages — dc_cluster_group_outside_vn / 112112210122 / 7

- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2331012323131020-2122213031302230-3102311133021323-2212313023122230-2032000330333110-1302122010311300-2133130321011200-1112320101103201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3110110313303300-2011210131111222-1320220020210233-0010310322310310-0221232302333202-3030321113221112-0133123313130020-3212223132230330"></a>

## ingress_egress_gw.forward_proxy_allow_all — forward_proxy_allow_all / 103012233033 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- ingress_egress_gw.forward_proxy_allow_all

<a id="canonical-3000031011310203-1320221003003120-2013002211321033-2110323001121100-3033323231020113-2201112003221123-1013202132131220-1200202313332110"></a>

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

<a id="canonical-3311213030021001-3330203013230231-3222310001223333-2112321202212300-2212202002120221-2212303332112301-3312333302323010-3221201333303310"></a>

## Direct properties — forward_proxy_allow_all / 103012233033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0233202111330321-2133321311211302-0303133133011212-1220002111112021-1321221332210002-2311323030033001-2231301111212211-0203223303222031"></a>

## Next pages — forward_proxy_allow_all / 103012233033 / 4

- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0110112001213332-2233201102333212-1321232102201102-3213031203300330-1210310211301200-3133010302300211-1313222203301301-0033103200013221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203233221300010-3100303023332231-2211031211001030-0222311212201210-2312331220121033-2112221300230112-0120220102203322-1020311221312103"></a>

## ingress_egress_gw.global_network_list — global_network_list / 102320321001 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- ingress_egress_gw.global_network_list

<a id="canonical-2311300032211111-2101331100232333-3311330113202313-2122130120312102-1000322021122230-0111020223311000-1200200223012300-0231303013202013"></a>

Type: `"object"`. single nested block, Optional.

Global Network Connection List. List of global network connections.

Upstream description:

List of global network connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("global_network_connections")}
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
global_network_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1100011212300322-0311213333133312-0201013331032312-3021203032102331-0121012120103131-2101221312102210-2313033300331312-1231311213220332"></a>

## Direct properties — global_network_list / 102320321001 / 3

- [global_network_connections](resources--aws_vpc_site--reference--group-002.md#canonical-3333302222033232-1233003332200333-3010013003100320-2301133331212231-2202321233313221-1031311123033011-1031232031331103-3022132100300121): complete subsection reference.

<a id="canonical-2332110023133201-0021021020223033-0133310000112213-3021113212322220-2230223111002001-1332300012223022-2101330130320211-1300103333003212"></a>

## Next pages — global_network_list / 102320321001 / 4

- [ingress_egress_gw.global_network_list.global_network_connections](resources--aws_vpc_site--reference--group-002.md#canonical-3333302222033232-1233003332200333-3010013003100320-2301133331212231-2202321233313221-1031311123033011-1031232031331103-3022132100300121)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3333302222033232-1233003332200333-3010013003100320-2301133331212231-2202321233313221-1031311123033011-1031232031331103-3022132100300121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303110213223101-0130030220311220-3011322003303100-0213121111212313-1231300301232113-0212112020201202-0130102001030231-1003303001202123"></a>

## ingress_egress_gw.global_network_list.global_network_connections — global_network_connections / 333000023010 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.global_network_list](resources--aws_vpc_site--reference--group-002.md#canonical-0110112001213332-2233201102333212-1321232102201102-3213031203300330-1210310211301200-3133010302300211-1313222203301301-0033103200013221)
- ingress_egress_gw.global_network_list.global_network_connections

<a id="canonical-2212001122030120-3331222201023201-1232020303220000-0233333013213102-2103012012113120-0233033012210010-0102311221121311-2113031310032130"></a>

Type: `"object"`. list nested block, Optional.

Global Network Connections. Global network connections.

Upstream description:

Global network connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("sli_to_global_dr",
    "slo_to_global_dr")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
global_network_connections {
  # Configure direct properties listed below.
}
```

<a id="canonical-0300231322121021-2020211223212220-1202103201001131-3232320103022213-1200220001131210-2230332321302302-1303103032303332-0003321103123231"></a>

## Direct properties — global_network_connections / 333000023010 / 3

- [sli_to_global_dr](resources--aws_vpc_site--reference--group-002.md#canonical-2210203120222300-0302312202231103-0301231011130312-3010133212000303-0030212033332212-3133000122020130-2302222112203213-3111200230232232): complete subsection reference.

- [slo_to_global_dr](resources--aws_vpc_site--reference--group-003.md#canonical-3333112200022311-2133223212010101-3230211221110023-3121110323000232-2133113222312032-0100112123230121-0030203222120223-0322323111333131): complete subsection reference.

<a id="canonical-0111311201132023-0122302122123131-0231021032311133-0313311121223230-2331322000033200-1330301122002301-1332221102220000-1023131303113310"></a>

## Next pages — global_network_connections / 333000023010 / 4

- [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](resources--aws_vpc_site--reference--group-002.md#canonical-2210203120222300-0302312202231103-0301231011130312-3010133212000303-0030212033332212-3133000122020130-2302222112203213-3111200230232232)
- [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](resources--aws_vpc_site--reference--group-003.md#canonical-3333112200022311-2133223212010101-3230211221110023-3121110323000232-2133113222312032-0100112123230121-0030203222120223-0322323111333131)
- [ingress_egress_gw.global_network_list](resources--aws_vpc_site--reference--group-002.md#canonical-0110112001213332-2233201102333212-1321232102201102-3213031203300330-1210310211301200-3133010302300211-1313222203301301-0033103200013221)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2210203120222300-0302312202231103-0301231011130312-3010133212000303-0030212033332212-3133000122020130-2302222112203213-3111200230232232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301103303022122-2100220210031222-0332022112332211-1031320232233322-2131022120120222-2123303110211310-2010102233010123-0000133321203032"></a>

## ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr — sli_to_global_dr / 023020232020 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.global_network_list](resources--aws_vpc_site--reference--group-002.md#canonical-0110112001213332-2233201102333212-1321232102201102-3213031203300330-1210310211301200-3133010302300211-1313222203301301-0033103200013221)
- [ingress_egress_gw.global_network_list.global_network_connections](resources--aws_vpc_site--reference--group-002.md#canonical-3333302222033232-1233003332200333-3010013003100320-2301133331212231-2202321233313221-1031311123033011-1031232031331103-3022132100300121)
- ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr

<a id="canonical-1101333022333030-0121123212103002-2112223330232113-2300122301012321-1030112022120133-0012322013221232-0220022012333331-1230312100012223"></a>

Type: `"object"`. single nested block, Optional.

Global network reference for direct connection.

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
sli_to_global_dr {
  # Configure direct properties listed below.
}
```

<a id="canonical-3330012132011310-3220103332221011-2012232313113312-3210231033200213-0233313031031212-0230113000333313-2030113332112100-0120023130102130"></a>

## Direct properties — sli_to_global_dr / 023020232020 / 3

- [global_vn](resources--aws_vpc_site--reference--group-002.md#canonical-1200212120123310-3023023230121101-1123020032102110-2003133000312103-1122102131010233-2223212212101120-3010023033301120-1002001111130103): complete subsection reference.

<a id="canonical-2130021111030002-3220210232100223-3131002203122100-2310110022322023-1331133220031333-3020101220303030-3011230312201313-2322323313113002"></a>

## Next pages — sli_to_global_dr / 023020232020 / 4

- [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn](resources--aws_vpc_site--reference--group-002.md#canonical-1200212120123310-3023023230121101-1123020032102110-2003133000312103-1122102131010233-2223212212101120-3010023033301120-1002001111130103)
- [ingress_egress_gw.global_network_list.global_network_connections](resources--aws_vpc_site--reference--group-002.md#canonical-3333302222033232-1233003332200333-3010013003100320-2301133331212231-2202321233313221-1031311123033011-1031232031331103-3022132100300121)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1200212120123310-3023023230121101-1123020032102110-2003133000312103-1122102131010233-2223212212101120-3010023033301120-1002001111130103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233222232110112-0213221030001223-3202301012133320-0223101022012210-0013031222113301-1030211001203112-2000311012212031-3123011130000223"></a>

## ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn — global_vn / 321213021022 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_egress_gw](resources--aws_vpc_site--reference--group-002.md#canonical-0301111030201010-2122031112010203-3213021330013203-0131020323011303-2213012321130231-1013323103213030-3223111212031102-2311112122002022)
- [ingress_egress_gw.global_network_list](resources--aws_vpc_site--reference--group-002.md#canonical-0110112001213332-2233201102333212-1321232102201102-3213031203300330-1210310211301200-3133010302300211-1313222203301301-0033103200013221)
- [ingress_egress_gw.global_network_list.global_network_connections](resources--aws_vpc_site--reference--group-002.md#canonical-3333302222033232-1233003332200333-3010013003100320-2301133331212231-2202321233313221-1031311123033011-1031232031331103-3022132100300121)
- [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](resources--aws_vpc_site--reference--group-002.md#canonical-2210203120222300-0302312202231103-0301231011130312-3010133212000303-0030212033332212-3133000122020130-2302222112203213-3111200230232232)
- ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn

<a id="canonical-0220201301001303-1003022013300213-1012313012110233-2231331200122103-0122221301330020-0320321220320020-1211210323210233-2302300000332330"></a>

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
global_vn {
  # Configure direct properties listed below.
}
```

<a id="canonical-0203332331233033-2200233110122020-3003113101122013-1312013322323200-3023110021311103-1102310213321010-3311210122220212-0311000331010231"></a>

## Direct properties — global_vn / 321213021022 / 3

<a id="canonical-0113032030220201-3330010102331213-0100333200232011-1023100320120211-0033300021102303-3121103131213103-0203103111132210-2010233200332001"></a>

<a id="canonical-1200202221222111-0231121031022132-0000122000031231-3011200311331201-2213001031220020-3133203222002101-3212232031322331-2120330003021130"></a>

## name property — global_vn / 321213021022 / 4

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

<a id="canonical-2112103002012102-1131101103202123-3101103131102101-3213323223130321-2313100032312030-2030132113112331-2200100010210110-3202203230031302"></a>
