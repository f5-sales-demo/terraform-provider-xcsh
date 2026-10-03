---
page_title: "xcsh_workload reference"
subcategory: "Container"
description: "Complete grouped canonical reference for xcsh_workload reference."
---

# xcsh_workload reference

<a id="canonical-2321030321303331-2100120302001222-1220131301031303-2211122021013110-0302231301230110-2321010223333012-0231332313131012-3203331000120212"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca — trusted_ca / 011011220020 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-006.md#canonical-1201123022002000-0122023033012312-0210013220112133-1220130111331213-0321022230013202-3213120123303001-1002300100311330-1023002312220123)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-0230223311012223-2220223212212231-0010330210301133-3011222311112222-1233110102130100-3102032012020332-1321001020103330-3123321320231032"></a>

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-2233023323120100-2101210323212112-0301111300313010-0033020330122121-3303031313232201-1132303011013231-3022113011112210-0301013111212020"></a>

## Direct properties — trusted_ca / 011011220020 / 3

<a id="canonical-2011221001201222-2233231031202211-3112103013222202-2322031210313131-1213321103323300-1102031200021013-0112320202331011-1320122021311323"></a>

<a id="canonical-2230223331122012-0023121021222102-0300111211232002-3203231010200303-0101012233300132-3113222212023223-1323200130312311-3000310122011011"></a>

## name property — trusted_ca / 011011220020 / 4

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

<a id="canonical-1102312333200312-2033102230013122-2202000103323132-1303210121202300-3201213100103033-1303220103011221-1232220120012331-0113220023103201"></a>

<a id="canonical-0022012010020023-0131102301001211-1020101333013033-3230311331001332-3111332022323111-3112201031110012-0230323303101023-3230030331013032"></a>

## namespace property — trusted_ca / 011011220020 / 5

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

<a id="canonical-1232011102210301-0210202231320021-1132322221122133-1020211302130331-1311013222222323-1210322023133020-3201233020030320-2021332113030222"></a>

<a id="canonical-1321213201020201-0011233122220222-1323203131210010-2333232012211130-0322120130003221-3001101022013013-1203203002033001-2010030012212003"></a>

## tenant property — trusted_ca / 011011220020 / 6

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

<a id="canonical-0123101313122110-3031202123312130-1312333320121332-1210101002131301-1020230232122023-2312312121233111-0203022301332101-3012002301232031"></a>

## Next pages — trusted_ca / 011011220020 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-006.md#canonical-1201123022002000-0122023033012312-0210013220112133-1220130111331213-0321022230013202-3213120123303001-1002300100311330-1023002312220123)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3101111211011003-0322113122333131-1320332000301323-2021001001032003-2302030313012323-0033322030001123-0200012121320010-1233222220022003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202300003132233-2100103011013002-1220003131302031-1222100232122113-3211120230213013-1010130003010112-2111333202230221-2032102212301031"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled — xfcc_disabled / 023302112231 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-006.md#canonical-1201123022002000-0122023033012312-0210013220112133-1220130111331213-0321022230013202-3213120123303001-1002300100311330-1023002312220123)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-2202112102302132-0332113313012333-0120330020131030-2131332323301100-0032332332201100-0111232011030002-1000231213010100-0332212331211300"></a>

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
xfcc_disabled = {}
```

<a id="canonical-3110201322232001-1331132121133221-1320303232011110-0232221320220110-3132022233302001-0220113201312223-2101333200310200-3302211230123322"></a>

## Direct properties — xfcc_disabled / 023302112231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2111020003000230-2002210023233023-1120201133103113-0010200330200133-1233332123131121-3132001223011202-1200200112232011-0132010322303202"></a>

## Next pages — xfcc_disabled / 023302112231 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-006.md#canonical-1201123022002000-0122023033012312-0210013220112133-1220130111331213-0321022230013202-3213120123303001-1002300100311330-1023002312220123)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2002102111223013-1130130103332311-0013000222311033-3323112200211201-3132200112231120-0221003231221002-3331103111300223-3222022211030303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0032123110330301-2200220223110103-1222001310122123-0201100202033103-1123100212301223-3313022101112131-1103320133230033-0110201032232320"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options — xfcc_options / 333223203033 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https](resources--workload--reference--group-005.md#canonical-2330031030101220-3210010323021133-3322231130101013-1101131102102332-3121022130002321-1233200100132332-0213220230130033-0110233032331011)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters](resources--workload--reference--group-006.md#canonical-3023301020330331-1221233122321220-2011202113013003-2331003313323320-0122021000022233-0121202301001222-1011313113011301-2103023311112320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-006.md#canonical-1201123022002000-0122023033012312-0210013220112133-1220130111331213-0321022230013202-3213120123303001-1002300100311330-1023002312220123)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls.xfcc_options

<a id="canonical-3033121121130101-0301331231112001-2130120032022331-0232112200131030-1001222001133232-0002200031321201-1321023032121001-3202022321321131"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("xfcc_header_elements")}
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
xfcc_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-3020221211330301-2120321133022323-3320313110132013-1331303001011123-1012031123010120-0033200300212202-3012301112331322-2032000312002330"></a>

## Direct properties — xfcc_options / 333223203033 / 3

<a id="canonical-3012103332330331-1323133201302212-1231010321232132-0131111110023322-2100220003202321-0130232120033211-2111331013112103-1031123123113132"></a>

<a id="canonical-3230313233133100-2001013310023200-3313203010221211-2321303330300120-0210201322320232-0031133331013001-2310301310332331-3022122303201000"></a>

## xfcc_header_elements property — xfcc_options / 333223203033 / 4

Type: `["list", "string"]`. Optional.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

Upstream description:

X-Forwarded-Client-Cert header elements to be added to requests.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  }
}
```

<a id="canonical-2300021001333121-3002203320030300-2313011103333210-2111030131300122-0121202111213301-0311131113320032-0220222300313331-3023213302310330"></a>

## Next pages — xfcc_options / 333223203033 / 5

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https.tls_parameters.use_mtls](resources--workload--reference--group-006.md#canonical-1201123022002000-0122023033012312-0210013220112133-1220130111331213-0321022230013202-3213120123303001-1002300100311330-1023002312220123)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300021231003002-2300111103320323-2100200313001013-1132300232212000-0120202320130211-1011212220203231-3202000102013222-1122003213002111"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert — https_auto_cert / 013012212102 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert

<a id="canonical-1123333213031103-0323222223113223-1300011122313222-0320302331210200-2222213323221312-1010222233323310-1303021110130013-3100032113003123"></a>

Type: `"object"`. single nested block, Optional.

Choice for selecting HTTP proxy with bring your own certificates.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("append_server_name",
    "default_header"),
  validators.ConflictingObjectAttributes("append_server_name",
    "pass_through"),
  validators.ConflictingObjectAttributes("append_server_name",
    "server_name"),
  validators.ConflictingObjectAttributes("default_header",
    "pass_through"),
  validators.ConflictingObjectAttributes("default_header",
    "server_name"),
  validators.ConflictingObjectAttributes("default_loadbalancer",
    "non_default_loadbalancer"),
  validators.ConflictingObjectAttributes("disable_path_normalize",
    "enable_path_normalize"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls"),
  validators.ConflictingObjectAttributes("pass_through",
    "server_name"),
  validators.ConflictingObjectAttributes("port",
    "port_ranges")}
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
  "x-ves-oneof-field-default_lb_choice": "[\"default_loadbalancer\",\"non_default_loadbalancer\"]",
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]",
  "x-ves-oneof-field-server_header_choice": "[\"append_server_name\",\"default_header\",\"pass_through\",\"server_name\"]"
}
```

Terraform syntax:

```terraform
https_auto_cert {
  # Configure direct properties listed below.
}
```

<a id="canonical-3032112303320221-2100231110200202-2122221101130032-0110212132113122-2101011033320203-1212102131233210-1320201333311212-3023133010333013"></a>

## Direct properties — https_auto_cert / 013012212102 / 3

<a id="canonical-1110220120202122-1123002323020321-0210110122120210-0022112001013322-1202212102320021-0311223232130233-1111211323223133-0303333023331330"></a>

<a id="canonical-1221011221221110-3031323010000302-0212331231231033-0330031010232213-2221020323013301-3031331023030133-2112000120011203-0231122130311202"></a>

## add_hsts property — https_auto_cert / 013012212102 / 4

Type: `"bool"`. Optional.

Add HTTP Strict-Transport-Security response header.

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

<a id="canonical-1133230122301300-1322220020221222-0203333011320203-3221112323222030-1333311300321020-1222100002103333-1012030330311223-2331030213011110"></a>

<a id="canonical-1200232303011030-0001211011203022-2233132322130302-0122211323330032-0300103001230221-2102210332200322-2300200100312201-3110120023021203"></a>

## append_server_name property — https_auto_cert / 013012212102 / 5

Type: `"string"`. Optional.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Upstream description:

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [coalescing_options](resources--workload--reference--group-007.md#canonical-1201200112321103-3013230121102322-0012022213022222-1032023322033313-3031230201201311-3033332211223310-0002133222021031-1330320101302130): complete subsection reference.

<a id="canonical-1200233332130313-0222221300300001-1031122032303033-3133132220232302-2321203300200111-3010321003320000-3300322121331320-2321011133010230"></a>

<a id="canonical-1010300021331132-0320002000202332-0321030322021123-1032123313023032-2302120223023012-3233020203312313-2231112033221111-0311031021110002"></a>

## connection_idle_timeout property — https_auto_cert / 013012212102 / 6

Type: `"number"`. Optional.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed.

Upstream description:

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [default_header](resources--workload--reference--group-007.md#canonical-0000231210020333-3302022311033333-0031231210123232-1203021320310130-3112121213201103-0010302000101103-3022103033320101-1023211322212211): complete subsection reference.

- [default_loadbalancer](resources--workload--reference--group-007.md#canonical-0020020301313213-3202033111301102-3330131131101221-2333123113210202-1000201032331321-3010202212333313-1221010200321030-1313031222112030): complete subsection reference.

- [disable_path_normalize](resources--workload--reference--group-007.md#canonical-2003030013221102-1101111123001021-3301102031002322-2131202112003232-0002112021331221-3332211211210100-2131003132001310-1332001113100131): complete subsection reference.

- [enable_path_normalize](resources--workload--reference--group-007.md#canonical-1211320032121322-3012022032220200-3130312112122033-3023221212122220-1323331123100310-1133232013113110-3221303012333303-3131223001201200): complete subsection reference.

- [http_protocol_options](resources--workload--reference--group-007.md#canonical-0100301311001121-0212203031223021-3012203333033211-2031100322132300-0312110002023220-1122212122111123-2123133203323212-3321012330030100): complete subsection reference.

<a id="canonical-3213030000201020-1010222013032203-2102330031120220-3101231203101102-0322013320320003-1223230320310312-1113312313003103-0103133302320322"></a>

<a id="canonical-0101223123333010-3003301333132001-2332223223003212-0132033003301011-0301330230212003-1203100203232311-2220320131010023-0102011213112320"></a>

## http_redirect property — https_auto_cert / 013012212102 / 7

Type: `"bool"`. Optional.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS.

Upstream description:

Redirect HTTP traffic to HTTPS.

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

- [no_mtls](resources--workload--reference--group-007.md#canonical-1113231223121000-0113312003012323-2213330232332121-2023023023331203-0311021213011213-1200310122010223-1211132000023000-2011000112321000): complete subsection reference.

- [non_default_loadbalancer](resources--workload--reference--group-007.md#canonical-1221310002232230-1033301211202133-1323122212110332-0003301010033211-0332333222200221-0210221011033212-3121203122012200-1312200202221211): complete subsection reference.

- [pass_through](resources--workload--reference--group-007.md#canonical-2033131323032122-0021123203121230-1020021333133302-0011033212330330-2202322301030300-2030023313303332-2131301203303312-2100233102111203): complete subsection reference.

<a id="canonical-2332322302112100-1231203022302023-1213302002322110-0233010100300100-2033312132311020-0010111222032010-0312132030203112-3120330023032101"></a>

<a id="canonical-2221013330032111-2001332022112201-3020020120311311-2302001010033321-1313121200232221-3131113131001102-1203022112023303-1313033200030302"></a>

## port property — https_auto_cert / 013012212102 / 8

Type: `"number"`. Optional.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1120231103123203-0221112020222020-3310313233233312-3321211303100230-2122300230120122-3233002121032331-0210113330010233-0200221011330311"></a>

<a id="canonical-1030203022210213-2010303211310311-2230012201021001-2102130202232201-1333120333130022-2311222022222310-1031320213222001-0312323211010310"></a>

## port_ranges property — https_auto_cert / 013012212102 / 9

Type: `"string"`. Optional.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

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
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.max_ports": "64",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.unique_port_range_list": "true"
  }
}
```

<a id="canonical-2123312131300210-2210311301101010-0211022101331300-1122302100322200-0123323320032200-0122110022023010-2233203302030223-2132103200220311"></a>

<a id="canonical-3212122310113212-3310330031302011-1212130230312022-1101033200231023-2132003333320100-0102001303211020-2020313332201022-2031003033121331"></a>

## server_name property — https_auto_cert / 013012212102 / 10

Type: `"string"`. Optional.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Upstream description:

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [tls_config](resources--workload--reference--group-007.md#canonical-3223323112230302-1313323021222023-3333321203103201-0222210032312000-2220122000303230-0323202020101213-3132102303130333-1221000232021310): complete subsection reference.

- [use_mtls](resources--workload--reference--group-007.md#canonical-0021023101320020-0312123020002010-2321023022121310-2313220310303001-0131310002012320-1232302030303303-1232020021321313-2313013010220230): complete subsection reference.

<a id="canonical-2233330333100323-1321110122330332-3130222123122020-0212003332233333-1321030100200120-0221022102112010-0313322322131120-3013331031021031"></a>

## Next pages — https_auto_cert / 013012212102 / 11

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-007.md#canonical-1201200112321103-3013230121102322-0012022213022222-1032023322033313-3031230201201311-3033332211223310-0002133222021031-1330320101302130)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_header](resources--workload--reference--group-007.md#canonical-0000231210020333-3302022311033333-0031231210123232-1203021320310130-3112121213201103-0010302000101103-3022103033320101-1023211322212211)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_loadbalancer](resources--workload--reference--group-007.md#canonical-0020020301313213-3202033111301102-3330131131101221-2333123113210202-1000201032331321-3010202212333313-1221010200321030-1313031222112030)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.disable_path_normalize](resources--workload--reference--group-007.md#canonical-2003030013221102-1101111123001021-3301102031002322-2131202112003232-0002112021331221-3332211211210100-2131003132001310-1332001113100131)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.enable_path_normalize](resources--workload--reference--group-007.md#canonical-1211320032121322-3012022032220200-3130312112122033-3023221212122220-1323331123100310-1133232013113110-3221303012333303-3131223001201200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-007.md#canonical-0100301311001121-0212203031223021-3012203333033211-2031100322132300-0312110002023220-1122212122111123-2123133203323212-3321012330030100)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.no_mtls](resources--workload--reference--group-007.md#canonical-1113231223121000-0113312003012323-2213330232332121-2023023023331203-0311021213011213-1200310122010223-1211132000023000-2011000112321000)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.non_default_loadbalancer](resources--workload--reference--group-007.md#canonical-1221310002232230-1033301211202133-1323122212110332-0003301010033211-0332333222200221-0210221011033212-3121203122012200-1312200202221211)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.pass_through](resources--workload--reference--group-007.md#canonical-2033131323032122-0021123203121230-1020021333133302-0011033212330330-2202322301030300-2030023313303332-2131301203303312-2100233102111203)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-007.md#canonical-3223323112230302-1313323021222023-3333321203103201-0222210032312000-2220122000303230-0323202020101213-3132102303130333-1221000232021310)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-007.md#canonical-0021023101320020-0312123020002010-2321023022121310-2313220310303001-0131310002012320-1232302030303303-1232020021321313-2313013010220230)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1201200112321103-3013230121102322-0012022213022222-1032023322033313-3031230201201311-3033332211223310-0002133222021031-1330320101302130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212221233231221-1130232321213303-3023000213130322-3323322112000201-1120332030120033-3121211102021101-0230100330132102-0133020323120320"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options — coalescing_options / 210012122320 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options

<a id="canonical-0111010302130232-2002313003101133-1113322030010102-0030013330233233-2202312032331010-3123213210002210-2320333031223112-3313301103101000"></a>

Type: `"object"`. single nested block, Optional.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_coalescing",
    "strict_coalescing")}
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
  "x-ves-oneof-field-coalescing_choice": "[\"default_coalescing\",\"strict_coalescing\"]"
}
```

Terraform syntax:

```terraform
coalescing_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-2020103303003302-0331101333010321-2323332223031013-1330323221201310-2023220122012103-3213223002033203-1303103112122331-2332121202221112"></a>

## Direct properties — coalescing_options / 210012122320 / 3

- [default_coalescing](resources--workload--reference--group-007.md#canonical-2033100212322113-0030113210213312-3133203001101010-3012133101203130-1131220111120022-1312121312100300-3021123213311110-1001321022032212): complete subsection reference.

- [strict_coalescing](resources--workload--reference--group-007.md#canonical-2231332113233303-2301213033203312-0320332322103222-0033131331302213-3012321233100333-1133121231020310-0302311220010310-1211113312320303): complete subsection reference.

<a id="canonical-1320232322320301-1133321112221322-1222232233211220-1230102223130131-2023033331211330-1101111201201211-0012132002323020-1332331221110312"></a>

## Next pages — coalescing_options / 210012122320 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing](resources--workload--reference--group-007.md#canonical-2033100212322113-0030113210213312-3133203001101010-3012133101203130-1131220111120022-1312121312100300-3021123213311110-1001321022032212)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing](resources--workload--reference--group-007.md#canonical-2231332113233303-2301213033203312-0320332322103222-0033131331302213-3012321233100333-1133121231020310-0302311220010310-1211113312320303)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2033100212322113-0030113210213312-3133203001101010-3012133101203130-1131220111120022-1312121312100300-3021123213311110-1001321022032212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211333220131333-2031001230233233-0120332003130322-2330030303020222-0311301223123332-3100013331021332-3311013021023133-2003123020120223"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing — default_coalescing / 111211120222 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-007.md#canonical-1201200112321103-3013230121102322-0012022213022222-1032023322033313-3031230201201311-3033332211223310-0002133222021031-1330320101302130)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.default_coalescing

<a id="canonical-0120110123010023-3101331200220210-2131220213111101-2021020302123302-2121001100202330-3120022022232023-1202303031112332-0100311121031032"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default coalescing.

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
default_coalescing = {}
```

<a id="canonical-3122211033123132-3122112211310311-2323012133311313-0100100232011013-3321013023130112-0123113130100122-1033301112033033-3021231322121110"></a>

## Direct properties — default_coalescing / 111211120222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1223022010122320-3200020002122222-3002230030332033-0023300223212303-0102313102102322-2021033310023030-2232001000101233-2001031130002310"></a>

## Next pages — default_coalescing / 111211120222 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-007.md#canonical-1201200112321103-3013230121102322-0012022213022222-1032023322033313-3031230201201311-3033332211223310-0002133222021031-1330320101302130)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2231332113233303-2301213033203312-0320332322103222-0033131331302213-3012321233100333-1133121231020310-0302311220010310-1211113312320303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1302133302032030-3312032023333311-2232133231020330-1222310333021330-2210100023012020-2311203030120020-0311121203000132-1011112113013301"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing — strict_coalescing / 020213203002 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-007.md#canonical-1201200112321103-3013230121102322-0012022213022222-1032023322033313-3031230201201311-3033332211223310-0002133222021031-1330320101302130)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options.strict_coalescing

<a id="canonical-3022000202010221-3301122122311230-0222230033012120-2231300011123022-0313332131130331-0133131102231023-2003311033300331-3113330202230302"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for strict coalescing.

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
strict_coalescing = {}
```

<a id="canonical-2320300131020032-3331133200202203-1022211120031300-1001030101330312-0333302120032012-0333112013113132-1000212332213212-2103320101013223"></a>

## Direct properties — strict_coalescing / 020213203002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0302311212330013-3131212121013011-2231101220033202-3020031123030302-3310233202030221-2111201101120023-1223003011001330-2103203010112002"></a>

## Next pages — strict_coalescing / 020213203002 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.coalescing_options](resources--workload--reference--group-007.md#canonical-1201200112321103-3013230121102322-0012022213022222-1032023322033313-3031230201201311-3033332211223310-0002133222021031-1330320101302130)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0000231210020333-3302022311033333-0031231210123232-1203021320310130-3112121213201103-0010302000101103-3022103033320101-1023211322212211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203221003113123-2331333122122033-3301332130000212-2110210001110011-1211230302023310-3131222010003122-2020023100010221-1012033110111020"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_header — default_header / 303203210323 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_header

<a id="canonical-1012132203023033-3030011310223012-3013031010233332-1201120332311332-2232020233012113-3110130121010202-1023230221300011-1220120122310013"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default header.

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
default_header = {}
```

<a id="canonical-0202133300120030-1312212222331103-2011101213300233-3120102123023030-0230020120122133-3202003201130132-0012010102323302-1201003320000222"></a>

## Direct properties — default_header / 303203210323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0121333122112001-3300031201102033-0333300113213131-0001020200300100-1130023330012230-2103311223021233-3122012021232203-0312333022113013"></a>

## Next pages — default_header / 303203210323 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0020020301313213-3202033111301102-3330131131101221-2333123113210202-1000201032331321-3010202212333313-1221010200321030-1313031222112030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220102030013232-2121300001120223-3001322003320303-3223023221032031-3330311332230320-3333333102112112-1323323110110113-1113210110311312"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_loadbalancer — default_loadbalancer / 021201222122 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.default_loadbalancer

<a id="canonical-2310230321323202-3132011231201302-2321121300310222-1313002110020332-2301330010130113-1302201303203030-0333302123012223-0232110301231300"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default loadbalancer.

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
default_loadbalancer = {}
```

<a id="canonical-0303210231031101-3110032011233222-2131023320121303-0223013313320303-2113112333333220-0233033200233020-3201112013222200-3220323023011232"></a>

## Direct properties — default_loadbalancer / 021201222122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2000122200220013-3200133331310222-3121310130113113-2212001121020111-2210210311031120-0000332113111301-2000013222320122-3220111221301323"></a>

## Next pages — default_loadbalancer / 021201222122 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2003030013221102-1101111123001021-3301102031002322-2131202112003232-0002112021331221-3332211211210100-2131003132001310-1332001113100131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101230313131200-3300101113113303-2230222122020202-3033021031310233-0112301111112233-2311021203122032-3103111313211301-1331130113113231"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.disable_path_normalize — disable_path_normalize / 031101300111 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.disable_path_normalize

<a id="canonical-0011301330220003-2302013222103110-3021301303200313-1011231010123002-1200320111002122-1301203222330102-0313020111031130-1132323020002111"></a>

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
disable_path_normalize = {}
```

<a id="canonical-1323000123123112-0230021220002323-1203002323232323-3030301321003203-0002202031210110-1000303023033120-2102110301002301-2312233230112033"></a>

## Direct properties — disable_path_normalize / 031101300111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1301102322132001-3212221222202233-0333210111132101-1212011120211103-1120010303231132-3102202123010100-2333210310032000-3200320321113022"></a>

## Next pages — disable_path_normalize / 031101300111 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1211320032121322-3012022032220200-3130312112122033-3023221212122220-1323331123100310-1133232013113110-3221303012333303-3131223001201200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1013022111213133-2111020321111233-3311031331031220-2110020102002002-1130100321002220-2233201303313132-1202132100223203-3021122030323013"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.enable_path_normalize — enable_path_normalize / 133230222032 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.enable_path_normalize

<a id="canonical-0233333212123021-0101303232130232-3302122320111102-3320120133131000-3003101033013220-1301201131223101-1020221121112033-3310031302223303"></a>

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
enable_path_normalize = {}
```

<a id="canonical-3122013001313121-0010000102233013-1110321213232003-3220100232200112-1230111110202130-1103001131122221-0330300003031322-0022331310113033"></a>

## Direct properties — enable_path_normalize / 133230222032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2113303011003001-0301332003130022-3331133001120002-3222010210033302-0000121333213112-1322013010132331-2321003213130112-0002331200303210"></a>

## Next pages — enable_path_normalize / 133230222032 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0100301311001121-0212203031223021-3012203333033211-2031100322132300-0312110002023220-1122212122111123-2123133203323212-3321012330030100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032022311021313-2212233130122121-0012320131001210-3300102102330101-2321002222032132-3103001133000210-3032103110000101-3320010211312212"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options — http_protocol_options / 131010010312 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options

<a id="canonical-2131031222333210-2010000212032131-2021000311130101-3233132320202020-3013312211312033-1212010323221212-1100102130333112-1001301021221202"></a>

Type: `"object"`. single nested block, Optional.

HTTP protocol configuration OPTIONS for downstream connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("http_protocol_enable_v1_only",
    "http_protocol_enable_v1_v2"),
  validators.ConflictingObjectAttributes("http_protocol_enable_v1_only",
    "http_protocol_enable_v2_only"),
  validators.ConflictingObjectAttributes("http_protocol_enable_v1_v2",
    "http_protocol_enable_v2_only")}
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
  "x-ves-oneof-field-http_protocol_choice": "[\"http_protocol_enable_v1_only\",\"http_protocol_enable_v1_v2\",\"http_protocol_enable_v2_only\"]"
}
```

Terraform syntax:

```terraform
http_protocol_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-2220031030222222-1022300130322000-3232103301021112-1033202213230121-2111110323311011-2231123132000003-0001323333022013-1213230000031033"></a>

## Direct properties — http_protocol_options / 131010010312 / 3

- [http_protocol_enable_v1_only](resources--workload--reference--group-007.md#canonical-1332323302302232-1121031212311221-1212110010231110-2020131322213102-3113030320203131-1331120110303030-3012201102222321-1322110303212302): complete subsection reference.

- [http_protocol_enable_v1_v2](resources--workload--reference--group-007.md#canonical-3032211223130332-0332212123020232-1221300020022302-1212302001302122-0120100223310333-0113322103011110-1132321201023300-2331203211323221): complete subsection reference.

- [http_protocol_enable_v2_only](resources--workload--reference--group-007.md#canonical-1222102231003010-2332332223213022-1130311011110333-3330223000111102-3231333202011311-2013320111211301-2322022220310002-3102323210031331): complete subsection reference.

<a id="canonical-1012323023331012-3220113032013012-3010211032310130-1111002002123111-2302030012110321-0211031330032022-2332011221211132-3110113100001322"></a>

## Next pages — http_protocol_options / 131010010312 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-007.md#canonical-1332323302302232-1121031212311221-1212110010231110-2020131322213102-3113030320203131-1331120110303030-3012201102222321-1322110303212302)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2](resources--workload--reference--group-007.md#canonical-3032211223130332-0332212123020232-1221300020022302-1212302001302122-0120100223310333-0113322103011110-1132321201023300-2331203211323221)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only](resources--workload--reference--group-007.md#canonical-1222102231003010-2332332223213022-1130311011110333-3330223000111102-3231333202011311-2013320111211301-2322022220310002-3102323210031331)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1332323302302232-1121031212311221-1212110010231110-2020131322213102-3113030320203131-1331120110303030-3012201102222321-1322110303212302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003032300220010-0120011102322221-1220310202103222-3023213200222102-0023020013303020-2213311132130212-0232300133030210-1031220110213333"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only — http_protocol_enable_v1_only / 321313230332 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-007.md#canonical-0100301311001121-0212203031223021-3012203333033211-2031100322132300-0312110002023220-1122212122111123-2123133203323212-3321012330030100)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-0313110302122332-2223130123102232-2330230022202031-2021221110012033-2331003010300030-2130123321123223-3131312023330302-3133213313133030"></a>

Type: `"object"`. single nested block, Optional.

HTTP/1.1 Protocol OPTIONS for downstream connections.

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
http_protocol_enable_v1_only {
  # Configure direct properties listed below.
}
```

<a id="canonical-2203012332300001-2213321011013121-0133331322220003-3100133223301223-1021030013013301-1122010333103231-2013010100301012-3210001223230013"></a>

## Direct properties — http_protocol_enable_v1_only / 321313230332 / 3

- [header_transformation](resources--workload--reference--group-007.md#canonical-0322011231312032-3113311323230003-1300323213222321-2201003331111220-3011210021000132-1233221133222200-2132333323102021-1210031221002212): complete subsection reference.

<a id="canonical-2112130221313123-3123010112232103-0011222013000032-2201103330223330-1200010210121113-0003312132200333-2212032213002120-0332100213023223"></a>

## Next pages — http_protocol_enable_v1_only / 321313230332 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-007.md#canonical-0322011231312032-3113311323230003-1300323213222321-2201003331111220-3011210021000132-1233221133222200-2132333323102021-1210031221002212)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-007.md#canonical-0100301311001121-0212203031223021-3012203333033211-2031100322132300-0312110002023220-1122212122111123-2123133203323212-3321012330030100)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0322011231312032-3113311323230003-1300323213222321-2201003331111220-3011210021000132-1233221133222200-2132333323102021-1210031221002212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200230102300211-2001313210203303-2231332033013013-2230012231001132-3200321003123321-2310230220212010-0321113010111013-0132313310201323"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation — header_transformation / 133223032111 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-007.md#canonical-0100301311001121-0212203031223021-3012203333033211-2031100322132300-0312110002023220-1122212122111123-2123133203323212-3321012330030100)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-007.md#canonical-1332323302302232-1121031212311221-1212110010231110-2020131322213102-3113030320203131-1331120110303030-3012201102222321-1322110303212302)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-2213122002000331-1301030100232220-1212321121331102-3232310230103000-3230232122200201-3213121222133102-0102120232231121-2223322210000030"></a>

Type: `"object"`. single nested block, Optional.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_header_transformation",
    "preserve_case_header_transformation"),
  validators.ConflictingObjectAttributes("default_header_transformation",
    "proper_case_header_transformation"),
  validators.ConflictingObjectAttributes("preserve_case_header_transformation",
    "proper_case_header_transformation")}
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
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

Terraform syntax:

```terraform
header_transformation {
  # Configure direct properties listed below.
}
```

<a id="canonical-2010111110201210-1110100202132200-2123102321320020-0200303100031123-2121123200201101-1000110230000022-1000110232033123-3102331220212321"></a>

## Direct properties — header_transformation / 133223032111 / 3

- [default_header_transformation](resources--workload--reference--group-007.md#canonical-0203312031021013-1002020000101122-0200200102200011-0131210323221213-1022220113201212-2220013211230023-3223033030111321-0013332333032310): complete subsection reference.

- [preserve_case_header_transformation](resources--workload--reference--group-007.md#canonical-0120031012131011-3003023003031321-1001023020122332-1131101210103120-2113100303031201-1132201102030321-2010230033030133-1101101312121113): complete subsection reference.

- [proper_case_header_transformation](resources--workload--reference--group-007.md#canonical-3302121110313213-0133003200112310-2030321302223310-1331310102232010-1232122230011133-3111223321202010-0231102300321212-3013201220202211): complete subsection reference.

<a id="canonical-3312311121023033-2113020113110223-0302331203221001-1130320120303000-0312231202033101-0313012023023311-0010210120311330-2033023330201101"></a>

## Next pages — header_transformation / 133223032111 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](resources--workload--reference--group-007.md#canonical-0203312031021013-1002020000101122-0200200102200011-0131210323221213-1022220113201212-2220013211230023-3223033030111321-0013332333032310)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](resources--workload--reference--group-007.md#canonical-0120031012131011-3003023003031321-1001023020122332-1131101210103120-2113100303031201-1132201102030321-2010230033030133-1101101312121113)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](resources--workload--reference--group-007.md#canonical-3302121110313213-0133003200112310-2030321302223310-1331310102232010-1232122230011133-3111223321202010-0231102300321212-3013201220202211)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-007.md#canonical-1332323302302232-1121031212311221-1212110010231110-2020131322213102-3113030320203131-1331120110303030-3012201102222321-1322110303212302)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0203312031021013-1002020000101122-0200200102200011-0131210323221213-1022220113201212-2220013211230023-3223033030111321-0013332333032310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310322123001113-3210020013220331-0023322323212011-2312211312003312-0302013311122010-2031322110030022-3320222100002300-2021120203001110"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — default_header_transformation / 100121222102 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-007.md#canonical-0100301311001121-0212203031223021-3012203333033211-2031100322132300-0312110002023220-1122212122111123-2123133203323212-3321012330030100)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-007.md#canonical-1332323302302232-1121031212311221-1212110010231110-2020131322213102-3113030320203131-1331120110303030-3012201102222321-1322110303212302)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-007.md#canonical-0322011231312032-3113311323230003-1300323213222321-2201003331111220-3011210021000132-1233221133222200-2132333323102021-1210031221002212)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-1301023201132132-2023311113002312-1303300013102113-3030333202030003-3333013212120232-2013311131200032-2113003202030321-2000202011002202"></a>

Type: `["object", {}]`. Optional.

Use the platform's current default HTTP header transformation behavior.

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
default_header_transformation = {}
```

<a id="canonical-2232222232022310-2003032131321302-2031312213033010-2202113320131311-2203200313001001-3203313111233000-2130000113110232-3332112223001302"></a>

## Direct properties — default_header_transformation / 100121222102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321102203003210-2323220011301010-1003131230331231-2310223123013203-2121103131003212-3333222311022323-0223113112010033-0311312100301223"></a>

## Next pages — default_header_transformation / 100121222102 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-007.md#canonical-0322011231312032-3113311323230003-1300323213222321-2201003331111220-3011210021000132-1233221133222200-2132333323102021-1210031221002212)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0120031012131011-3003023003031321-1001023020122332-1131101210103120-2113100303031201-1132201102030321-2010230033030133-1101101312121113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232310211213213-3212321232001130-0010222233013200-2202311103303200-3330312012110000-2233221122211220-1323210232331032-3102031013123102"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — preserve_case_header_transformation / 202023012003 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-007.md#canonical-0100301311001121-0212203031223021-3012203333033211-2031100322132300-0312110002023220-1122212122111123-2123133203323212-3321012330030100)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-007.md#canonical-1332323302302232-1121031212311221-1212110010231110-2020131322213102-3113030320203131-1331120110303030-3012201102222321-1322110303212302)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-007.md#canonical-0322011231312032-3113311323230003-1300323213222321-2201003331111220-3011210021000132-1233221133222200-2132333323102021-1210031221002212)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-1132221200201220-3223211132231000-3233110001202222-2102202203003023-0302133213102000-3322212230012310-1030331023300222-1321003033311103"></a>

Type: `["object", {}]`. Optional.

Preserve HTTP header-name case when upstream case must remain unchanged.

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
preserve_case_header_transformation = {}
```

<a id="canonical-3203100001101201-1331132330231330-2001031323200001-3213002133222310-3012032001133000-2110120222222130-1310313312211100-3312002031301112"></a>

## Direct properties — preserve_case_header_transformation / 202023012003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0212021021123203-3231302200322302-3210113123322320-3121021101132013-0200100333211031-1121012203322312-1003003013131030-1112320220213123"></a>

## Next pages — preserve_case_header_transformation / 202023012003 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-007.md#canonical-0322011231312032-3113311323230003-1300323213222321-2201003331111220-3011210021000132-1233221133222200-2132333323102021-1210031221002212)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3302121110313213-0133003200112310-2030321302223310-1331310102232010-1232122230011133-3111223321202010-0231102300321212-3013201220202211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102013321111231-3203120113321200-2132021201333111-2100021010313331-2320230131132222-0002110102333302-0231022120331013-3210322133323320"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — proper_case_header_transformation / 112103202103 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-007.md#canonical-0100301311001121-0212203031223021-3012203333033211-2031100322132300-0312110002023220-1122212122111123-2123133203323212-3321012330030100)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only](resources--workload--reference--group-007.md#canonical-1332323302302232-1121031212311221-1212110010231110-2020131322213102-3113030320203131-1331120110303030-3012201102222321-1322110303212302)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-007.md#canonical-0322011231312032-3113311323230003-1300323213222321-2201003331111220-3011210021000132-1233221133222200-2132333323102021-1210031221002212)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-1233112022312323-0131131312102030-1022020023301300-0321310330222210-1023002233133230-2220102231121222-2220200003200013-2223100201013211"></a>

Type: `["object", {}]`. Optional.

Transform HTTP header names to proper case when explicit transformation is required.

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
proper_case_header_transformation = {}
```

<a id="canonical-3333213111010300-3212002321321231-0232213001220020-2020122233110212-2211222200311220-2110301113331201-0013233121122102-1123100112120011"></a>

## Direct properties — proper_case_header_transformation / 112103202103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0310101200332310-0101223312300203-0100100011220031-0031333101233032-1121130022313303-1131022212201101-2113230213230133-1311013223320021"></a>

## Next pages — proper_case_header_transformation / 112103202103 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_only.header_transformation](resources--workload--reference--group-007.md#canonical-0322011231312032-3113311323230003-1300323213222321-2201003331111220-3011210021000132-1233221133222200-2132333323102021-1210031221002212)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3032211223130332-0332212123020232-1221300020022302-1212302001302122-0120100223310333-0113322103011110-1132321201023300-2331203211323221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303011203313121-1232010111023200-1320231032110233-0310333132011131-3330313030300303-1213302103200112-0320221022201231-0123121023111020"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2 — http_protocol_enable_v1_v2 / 120232331301 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-007.md#canonical-0100301311001121-0212203031223021-3012203333033211-2031100322132300-0312110002023220-1122212122111123-2123133203323212-3321012330030100)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-2223323212133232-1021210321121312-3231032322022223-0111102201330032-0310201130210311-2312103003200101-1011233203030012-1031032200011003"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for http protocol enable v1 v2.

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
http_protocol_enable_v1_v2 = {}
```

<a id="canonical-1110133023113302-2002020111012211-3023321202132103-1320031113000230-2323103212233113-2303020300111000-0230030200022202-1210223111113213"></a>

## Direct properties — http_protocol_enable_v1_v2 / 120232331301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3211230013102312-3011223232121021-0322020122130122-2031112131303011-2121111311103311-2001020301322323-3313220130112103-0031103302012302"></a>

## Next pages — http_protocol_enable_v1_v2 / 120232331301 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-007.md#canonical-0100301311001121-0212203031223021-3012203333033211-2031100322132300-0312110002023220-1122212122111123-2123133203323212-3321012330030100)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1222102231003010-2332332223213022-1130311011110333-3330223000111102-3231333202011311-2013320111211301-2322022220310002-3102323210031331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122120112200033-1023223011320221-0111301202221130-2020222111223031-1313330210301200-0211323010212120-1002101133101220-3201000313232200"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only — http_protocol_enable_v2_only / 112011010130 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-007.md#canonical-0100301311001121-0212203031223021-3012203333033211-2031100322132300-0312110002023220-1122212122111123-2123133203323212-3321012330030100)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-0133030122320133-1302001233320010-3112220233333101-1112332002323212-3132121103313333-1211110200210313-0022320303100313-3310133131022213"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for http protocol enable v2 only.

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
http_protocol_enable_v2_only = {}
```

<a id="canonical-0220031200112312-2000002021202032-1121223131130111-1330211103320211-0323213300123231-1313123003030333-0302033131112222-2023130312200021"></a>

## Direct properties — http_protocol_enable_v2_only / 112011010130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2310032131212320-1033010330020022-1302321113000132-0132030313202102-2100313303322201-3220010000233113-3003203312202123-0023012230311110"></a>

## Next pages — http_protocol_enable_v2_only / 112011010130 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.http_protocol_options](resources--workload--reference--group-007.md#canonical-0100301311001121-0212203031223021-3012203333033211-2031100322132300-0312110002023220-1122212122111123-2123133203323212-3321012330030100)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1113231223121000-0113312003012323-2213330232332121-2023023023331203-0311021213011213-1200310122010223-1211132000023000-2011000112321000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2231221001000323-1000301010120313-2023222101231002-2132220012133323-0223001000232331-2320001001002310-0030020223020333-3112131102102302"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.no_mtls — no_mtls / 013210113033 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.no_mtls

<a id="canonical-3221030000033213-3302230112301212-3132313212122213-1202201313230211-1120013003331002-0220010112323000-0020322110021110-0000310012002303"></a>

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
no_mtls = {}
```

<a id="canonical-0112320003111221-1322230202101020-3220213031102233-2212313131001023-2332202130033232-3011131202030330-2221132023320103-3320030331013311"></a>

## Direct properties — no_mtls / 013210113033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3331113110320312-1110203003013201-0203031220210303-1331330233001201-3203001133221032-3333231101012230-2110111011210200-0203312021301200"></a>

## Next pages — no_mtls / 013210113033 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1221310002232230-1033301211202133-1323122212110332-0003301010033211-0332333222200221-0210221011033212-3121203122012200-1312200202221211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131120023323130-2031210221022101-3300102203001311-2320212003120223-0201233322113312-2212112220112032-1130330100130302-3222120311321301"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.non_default_loadbalancer — non_default_loadbalancer / 130023303213 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.non_default_loadbalancer

<a id="canonical-1021301102331030-2103303130301321-3032313312100011-1013302133312003-1013313323322233-2300300220120221-2120021122130013-0123231312302131"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for non default loadbalancer.

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
non_default_loadbalancer = {}
```

<a id="canonical-2021200030110020-2212020201201122-2030231100101201-3101313103021322-2022222031123122-2211020100233201-0123231323020213-3022111110032112"></a>

## Direct properties — non_default_loadbalancer / 130023303213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0122221331332130-0021010223213301-3111210201112203-1003133103020032-1123330020123210-2310300330111330-2313013230001002-0021332210002302"></a>

## Next pages — non_default_loadbalancer / 130023303213 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2033131323032122-0021123203121230-1020021333133302-0011033212330330-2202322301030300-2030023313303332-2131301203303312-2100233102111203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322122033200221-3021020112330210-2102303221310312-0113231030020300-3203233020210112-1032112232213311-2030010210231330-1122100210221211"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.pass_through — pass_through / 000203131330 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.pass_through

<a id="canonical-0010130331201022-2223011210220302-1211303222113202-3013230233132112-1120022112310331-3320120302320331-2232023033211310-2023222010223103"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for pass through.

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
pass_through = {}
```

<a id="canonical-0203012320011302-1031201003132322-2112101221110002-2220011100313302-3200102222011332-2023113100303231-2223011002113313-2231103113113223"></a>

## Direct properties — pass_through / 000203131330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3302022200100020-2103311203332322-1300100300323111-2310120103311331-2101220210031233-2033203202103211-3330233000301012-3002011221001312"></a>

## Next pages — pass_through / 000203131330 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3223323112230302-1313323021222023-3333321203103201-0222210032312000-2220122000303230-0323202020101213-3132102303130333-1221000232021310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130101002033201-2322002031323000-1221010000332231-3122212030231203-0232213333322333-1131301200201102-0011121230022102-2330333211230320"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config — tls_config / 021310323301 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config

<a id="canonical-0322213321332121-2110123101020020-3332101300211100-1230111011013120-3110222331102301-2320001010033013-0220112032321332-2320123213101223"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-2311211302212012-2003111020030230-2313301301120020-2212211110230233-2123033203121312-3323220103302032-0003000331013231-2331023020012120"></a>

## Direct properties — tls_config / 021310323301 / 3

- [custom_security](resources--workload--reference--group-007.md#canonical-2210213031030220-0331231123310231-1202102311020323-1200032103333323-3313210101213210-3200101102302323-3030323130322020-1033033033223102): complete subsection reference.

- [default_security](resources--workload--reference--group-007.md#canonical-0303012013132120-2031131031120213-0123302322222320-0323212310001233-0131210112032312-3311031330123311-0331310232231111-3221302030110232): complete subsection reference.

- [low_security](resources--workload--reference--group-007.md#canonical-0312213103031030-1103023303301112-1030123112332331-0210233330010203-3213130011120110-1212030232012012-0100232011300333-0200200100320120): complete subsection reference.

- [medium_security](resources--workload--reference--group-007.md#canonical-3213020000102211-2312312032233321-3130231023301302-2320001201111131-3001221010231300-3231001303001133-3030301330133032-2021102012000320): complete subsection reference.

<a id="canonical-0200003130021021-1020210101121223-2211223221110210-3210110122121131-0022303203001202-3012013300031201-3203233020213322-2010210210023210"></a>

## Next pages — tls_config / 021310323301 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security](resources--workload--reference--group-007.md#canonical-2210213031030220-0331231123310231-1202102311020323-1200032103333323-3313210101213210-3200101102302323-3030323130322020-1033033033223102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.default_security](resources--workload--reference--group-007.md#canonical-0303012013132120-2031131031120213-0123302322222320-0323212310001233-0131210112032312-3311031330123311-0331310232231111-3221302030110232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.low_security](resources--workload--reference--group-007.md#canonical-0312213103031030-1103023303301112-1030123112332331-0210233330010203-3213130011120110-1212030232012012-0100232011300333-0200200100320120)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.medium_security](resources--workload--reference--group-007.md#canonical-3213020000102211-2312312032233321-3130231023301302-2320001201111131-3001221010231300-3231001303001133-3030301330133032-2021102012000320)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2210213031030220-0331231123310231-1202102311020323-1200032103333323-3313210101213210-3200101102302323-3030323130322020-1033033033223102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131120032303302-3013121000131223-3320013301131110-1113333032322231-0120121103230013-2233323003030010-1201201221003110-1133021013002300"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security — custom_security / 333021002220 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-007.md#canonical-3223323112230302-1313323021222023-3333321203103201-0222210032312000-2220122000303230-0323202020101213-3132102303130333-1221000232021310)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.custom_security

<a id="canonical-1020300302123203-0300133101321211-1302132302301100-3011311222002120-1023230023122313-1301213000103210-3133311130231100-1010122310200201"></a>

Type: `"object"`. single nested block, Optional.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cipher_suites")}
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
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-3131003133000023-1202132000333230-2332132130131023-2201001003331203-1020033203133122-1010230331133013-0032323310120212-0323032313201232"></a>

## Direct properties — custom_security / 333021002220 / 3

<a id="canonical-1031131322332121-0123232220131320-0213212023101232-1212111203021110-1102303332012021-1321301121012023-3123111322032123-1302110132230322"></a>

<a id="canonical-1003003323103030-0333012303131002-3310200333031132-1111121300113233-0233223233231232-2310203323200002-2033222130001031-2231330333330321"></a>

## cipher_suites property — custom_security / 333021002220 / 4

Type: `["list", "string"]`. Optional.

The TLS listener will only support the specified cipher list.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3203103011233330-1322311001323323-0120020002212133-3212111200003223-2000311001002131-2031023020020011-1201320131133121-1010331330121223"></a>

<a id="canonical-0321100022320120-3101331021121311-3203111120312202-2202100332330132-3322212323332313-2023113010202111-1110233231020231-0203121222101033"></a>

## max_version property — custom_security / 333021002220 / 5

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

<a id="canonical-3131330200032010-0202110223102101-2030311232033010-3012233022223222-2313212212202030-0111031300233331-1310012011321032-2232000123130123"></a>

<a id="canonical-0223220023121210-1011102203220220-2300210213130111-1320332310001200-1303303230230233-2231022031203133-3020020333220110-0110232133313221"></a>

## min_version property — custom_security / 333021002220 / 6

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

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

<a id="canonical-2031130213211313-1010100111220212-1011330122301033-0021200202322222-3312112332123213-0101300323130300-2012032232132303-0233201133121232"></a>

## Next pages — custom_security / 333021002220 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-007.md#canonical-3223323112230302-1313323021222023-3333321203103201-0222210032312000-2220122000303230-0323202020101213-3132102303130333-1221000232021310)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0303012013132120-2031131031120213-0123302322222320-0323212310001233-0131210112032312-3311031330123311-0331310232231111-3221302030110232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202220221032220-3111102121213223-0103313311111101-3211323223322322-2201321121312321-3033023211332122-0210302130201230-0310331231201133"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.default_security — default_security / 103221022300 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-007.md#canonical-3223323112230302-1313323021222023-3333321203103201-0222210032312000-2220122000303230-0323202020101213-3132102303130333-1221000232021310)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.default_security

<a id="canonical-3021333200232033-2310010131011120-1200033223003311-0322102223200032-3323103103120310-0123110023302333-3201220013111222-3211220111123213"></a>

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
default_security = {}
```

<a id="canonical-2212203122203232-0221013203211100-3133212320002000-0132313120020011-0302200113211120-0301302030112132-1133202020131300-3213300020020121"></a>

## Direct properties — default_security / 103221022300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0112210222121121-1000131100001202-1122123122203120-3203122311312103-2201002203020003-3331020101013310-1310330312011331-1003320111111012"></a>

## Next pages — default_security / 103221022300 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-007.md#canonical-3223323112230302-1313323021222023-3333321203103201-0222210032312000-2220122000303230-0323202020101213-3132102303130333-1221000232021310)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0312213103031030-1103023303301112-1030123112332331-0210233330010203-3213130011120110-1212030232012012-0100232011300333-0200200100320120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310203011021223-0103031331032211-3031313322030023-0303302322130011-0223230303032323-3012233102312020-3023100322223231-1300213032020203"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.low_security — low_security / 210203210113 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-007.md#canonical-3223323112230302-1313323021222023-3333321203103201-0222210032312000-2220122000303230-0323202020101213-3132102303130333-1221000232021310)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.low_security

<a id="canonical-2232331032032332-0020330311232221-0122200122312110-1302321033100210-2333101202331033-3123301213013123-2030102232021013-2123211221313232"></a>

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
low_security = {}
```

<a id="canonical-3111000203322100-0230223333000003-2333112212120113-1021300232103223-0330112123003322-3231221201332331-1020020000221310-2212011211002131"></a>

## Direct properties — low_security / 210203210113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321102101303320-0310223200331003-1310201023213212-1310210323123113-0212121323032011-1232132113011112-1330001203233033-3333022000222333"></a>

## Next pages — low_security / 210203210113 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-007.md#canonical-3223323112230302-1313323021222023-3333321203103201-0222210032312000-2220122000303230-0323202020101213-3132102303130333-1221000232021310)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3213020000102211-2312312032233321-3130231023301302-2320001201111131-3001221010231300-3231001303001133-3030301330133032-2021102012000320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301033031123210-0332033201331122-0121131331230222-3233222003131311-3320201021011233-0012310322310221-1221122222123012-2202012310312301"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.medium_security — medium_security / 022031122001 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-007.md#canonical-3223323112230302-1313323021222023-3333321203103201-0222210032312000-2220122000303230-0323202020101213-3132102303130333-1221000232021310)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config.medium_security

<a id="canonical-3330333002231303-1130200111020221-0132332133301020-3032023233121010-1123131100301030-1011300112013102-0030302323022222-3322332132223230"></a>

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
medium_security = {}
```

<a id="canonical-3010132313203300-1212011031022322-1011233332302330-0321121322103120-3031233301311233-3203132133322203-2101331030030110-1013210302022033"></a>

## Direct properties — medium_security / 022031122001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3003231023330131-1110002300203200-1102131130222231-3311321031210330-1310102322020321-0210220122313103-3013030022231232-3332111302102103"></a>

## Next pages — medium_security / 022031122001 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.tls_config](resources--workload--reference--group-007.md#canonical-3223323112230302-1313323021222023-3333321203103201-0222210032312000-2220122000303230-0323202020101213-3132102303130333-1221000232021310)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0021023101320020-0312123020002010-2321023022121310-2313220310303001-0131310002012320-1232302030303303-1232020021321313-2313013010220230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213312103113311-2020111022213220-2130030302200123-2230230301120113-0312300130022332-1223201001330311-1322130231023100-1122332310202303"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls — use_mtls / 213023120202 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls

<a id="canonical-0112132223201103-2012210001022101-1202101022210311-0002300202133133-2223011222011021-0030310321030301-0330302031303301-0032333012312233"></a>

Type: `"object"`. single nested block, Optional.

Validation context for downstream client TLS connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("crl",
    "no_crl"),
  validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url"),
  validators.ConflictingObjectAttributes("xfcc_disabled",
    "xfcc_options")}
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
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

Terraform syntax:

```terraform
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-1121021012130101-3232100223223011-0321200030313211-1323200222020132-0231323013030310-2202200100313322-3231000123100001-3021300333111313"></a>

## Direct properties — use_mtls / 213023120202 / 3

<a id="canonical-2233113303130031-2232032203322212-0022002030121113-2313023203003220-0022023211111011-1232120130122101-2323312330321020-0123232213112133"></a>

<a id="canonical-3200301100120310-0303210021133322-0101303211020321-0103011230323120-0222020322201213-0211333101223313-3221023131121002-1331110003111011"></a>

## client_certificate_optional property — use_mtls / 213023120202 / 4

Type: `"bool"`. Optional.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated.

Upstream description:

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

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

- [crl](resources--workload--reference--group-007.md#canonical-1013302010320000-1011230322021312-3012323210131123-3013312012031311-0022332332220121-3220001222033012-2013223323201213-3301300020232222): complete subsection reference.

- [no_crl](resources--workload--reference--group-007.md#canonical-3302231003012110-2000302332123203-3313302022111322-3230310203330311-1122230012120113-2313102032030111-0213000203031203-0031301310033200): complete subsection reference.

- [trusted_ca](resources--workload--reference--group-007.md#canonical-3312221002221111-0222020203132233-1233200001302233-3002223000011203-3312232220103132-2323300230323022-3230031300330032-3121332103113222): complete subsection reference.

<a id="canonical-0210032111200002-2230323120233123-3110120112131301-2022332013113302-2220111120220221-0310111133023213-3100121122103033-3220123111211112"></a>

<a id="canonical-0022201301213000-3033121223321221-3101330312132010-2300230321302202-0333011010112102-2022103330133102-0213213102110220-1312133221011133"></a>

## trusted_ca_url property — use_mtls / 213023120202 / 5

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

- [xfcc_disabled](resources--workload--reference--group-007.md#canonical-1233303310133222-2233301300010300-0330120213311132-3312323203123231-3003031133123110-2230220233101011-0112102000310033-1130101303113222): complete subsection reference.

- [xfcc_options](resources--workload--reference--group-007.md#canonical-0313113021102333-0001033033120201-2221111132100321-0120033003012012-2022223101220022-3330230223013023-3322013300320223-3000211100123010): complete subsection reference.

<a id="canonical-3223302012122102-1023022323021211-3231303331320032-0203100103000030-3321001321030302-1003213002100013-0300233211322302-2230312032202313"></a>

## Next pages — use_mtls / 213023120202 / 6

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.crl](resources--workload--reference--group-007.md#canonical-1013302010320000-1011230322021312-3012323210131123-3013312012031311-0022332332220121-3220001222033012-2013223323201213-3301300020232222)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.no_crl](resources--workload--reference--group-007.md#canonical-3302231003012110-2000302332123203-3313302022111322-3230310203330311-1122230012120113-2313102032030111-0213000203031203-0031301310033200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca](resources--workload--reference--group-007.md#canonical-3312221002221111-0222020203132233-1233200001302233-3002223000011203-3312232220103132-2323300230323022-3230031300330032-3121332103113222)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled](resources--workload--reference--group-007.md#canonical-1233303310133222-2233301300010300-0330120213311132-3312323203123231-3003031133123110-2230220233101011-0112102000310033-1130101303113222)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options](resources--workload--reference--group-007.md#canonical-0313113021102333-0001033033120201-2221111132100321-0120033003012012-2022223101220022-3330230223013023-3322013300320223-3000211100123010)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1013302010320000-1011230322021312-3012323210131123-3013312012031311-0022332332220121-3220001222033012-2013223323201213-3301300020232222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313332201233111-1323102313120000-2330031110232101-2200120333121122-0103110103032301-2232021033000130-0122311201232313-2320313121310312"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.crl — crl / 001113130100 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-007.md#canonical-0021023101320020-0312123020002010-2321023022121310-2313220310303001-0131310002012320-1232302030303303-1232020021321313-2313013010220230)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.crl

<a id="canonical-3302100103232001-1113213111211313-1100221212220202-0130110321130212-1233010013230331-0201001232032221-1211212120020002-3003313330113230"></a>

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
crl {
  # Configure direct properties listed below.
}
```

<a id="canonical-3223020332221202-1031301001111102-1220133221310223-3102003300022330-2130010310301201-2103223300212013-0233233133320012-1221312300312102"></a>

## Direct properties — crl / 001113130100 / 3

<a id="canonical-3010133121312001-2311230113122301-3322120121200122-1110211033023330-3032130100113112-2321001123331122-2310023111123131-2020303123222011"></a>

<a id="canonical-2320223111121132-2123113021233023-2311321303113111-1131022200303011-3010231311030030-0123322120111231-2131003122321230-2323311333222123"></a>

## name property — crl / 001113130100 / 4

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

<a id="canonical-3201200331012222-0031122202001301-2021211312112033-1220232013001022-1102030130312031-2202213322310001-1231330303000112-1132120303310301"></a>

<a id="canonical-0112321030121131-3331303021200212-2210132003010031-1103321013023332-3103313100023322-3212303112331321-1321332030012333-0002003023200301"></a>

## namespace property — crl / 001113130100 / 5

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

<a id="canonical-1032332220233002-2033303230022312-2022333132322301-0223032232212301-1031221032113332-3231311002123233-3231300131002100-0302301302131313"></a>

<a id="canonical-0200213023010013-2312302302113100-0132203013113233-2200211012033130-1222320030032012-0310000201020032-0332102203021233-1002122333021032"></a>

## tenant property — crl / 001113130100 / 6

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

<a id="canonical-2320120000302101-0010112311111212-2320323330320032-2031002233110033-1003031033131002-1121301210102123-1131122003210100-3231011312223111"></a>

## Next pages — crl / 001113130100 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-007.md#canonical-0021023101320020-0312123020002010-2321023022121310-2313220310303001-0131310002012320-1232302030303303-1232020021321313-2313013010220230)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3302231003012110-2000302332123203-3313302022111322-3230310203330311-1122230012120113-2313102032030111-0213000203031203-0031301310033200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203313111303030-3023232032110223-2103002000011303-3132110213000200-3212322310111011-2031020231102100-3021110213330223-1312302310300003"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.no_crl — no_crl / 330011032200 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-007.md#canonical-0021023101320020-0312123020002010-2321023022121310-2313220310303001-0131310002012320-1232302030303303-1232020021321313-2313013010220230)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.no_crl

<a id="canonical-1003120012203210-0110302111102201-0300223220010010-0211020010000012-1220332133323012-1300202303202032-0003121031033232-2210202123030232"></a>

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
no_crl = {}
```

<a id="canonical-3320023310330122-1301221110130002-3301323213132323-2301113032302000-1213132103113332-1003100333100012-0333223120211210-2321223220310101"></a>

## Direct properties — no_crl / 330011032200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3111222331330133-3233222203112302-3010330331002023-1322002310311330-0001032312210331-2213320022030313-0011321002031313-0220030001220302"></a>

## Next pages — no_crl / 330011032200 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-007.md#canonical-0021023101320020-0312123020002010-2321023022121310-2313220310303001-0131310002012320-1232302030303303-1232020021321313-2313013010220230)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3312221002221111-0222020203132233-1233200001302233-3002223000011203-3312232220103132-2323300230323022-3230031300330032-3121332103113222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231100310313131-3310103301033311-0111211222210110-3311021220021032-1021011223111222-1220203020012000-3311331211310103-2033121310222033"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca — trusted_ca / 333332033020 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-007.md#canonical-0021023101320020-0312123020002010-2321023022121310-2313220310303001-0131310002012320-1232302030303303-1232020021321313-2313013010220230)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.trusted_ca

<a id="canonical-0011023310313320-3232231100312221-2313112022121131-0133230202123201-1231302211210033-3210313301322112-2033002011012201-3201000120232230"></a>

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-2322302003011223-3211131121111033-1013132100312103-0003102202001201-2302130203230000-1030320020002330-2233232102133200-3220310301331000"></a>

## Direct properties — trusted_ca / 333332033020 / 3

<a id="canonical-1131201020030221-2121132012000113-2120013303131220-1002322300133232-2031210010312010-0110033221122321-3210311213230233-1211303213200121"></a>

<a id="canonical-1130000122333003-0130300230321131-3133023021002030-0331120232221321-1123303331122112-2031102310202312-2300330320130313-0311321213011311"></a>

## name property — trusted_ca / 333332033020 / 4

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

<a id="canonical-3120311003232110-1103313233212022-2333031321223311-0000021030233333-0112012202321320-3323023222123300-2202323021230330-0000332113100310"></a>

<a id="canonical-2213312011110000-1312321322101313-3231002023033222-1310020212003032-3310201322231303-1331321313232103-3002100100110132-2030122001113120"></a>

## namespace property — trusted_ca / 333332033020 / 5

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

<a id="canonical-1320221100331001-2010031132021000-2232302213103202-2133320330332301-2032202212220113-0202333303133010-2112311113320013-0012233030331323"></a>

<a id="canonical-2022010310203003-0103300033211301-0112201331033031-2321211333001122-3313023120023222-3233011113303030-0033210012003002-1211033320112021"></a>

## tenant property — trusted_ca / 333332033020 / 6

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

<a id="canonical-2030103231122000-2130303202332332-1101002000033013-2013033123033200-0230003321103311-3212223323223332-3022201303221010-3022020012223030"></a>

## Next pages — trusted_ca / 333332033020 / 7

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-007.md#canonical-0021023101320020-0312123020002010-2321023022121310-2313220310303001-0131310002012320-1232302030303303-1232020021321313-2313013010220230)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-1233303310133222-2233301300010300-0330120213311132-3312323203123231-3003031133123110-2230220233101011-0112102000310033-1130101303113222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1033321212301111-2203202310033211-1111022112311222-3212210033302322-2321322131012300-1113003023122310-2313210312330022-2101310103313231"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled — xfcc_disabled / 101310220212 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-007.md#canonical-0021023101320020-0312123020002010-2321023022121310-2313220310303001-0131310002012320-1232302030303303-1232020021321313-2313013010220230)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_disabled

<a id="canonical-3212000331322202-0230222221111110-3112112003133323-1201003233203003-2100010313100202-2021021100122312-3021303020233213-0001013330320212"></a>

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
xfcc_disabled = {}
```

<a id="canonical-0133120222313110-3331310320023022-2221202233230320-2022131011000302-0223013303210301-2102133030112020-1200200223213201-3203021000131013"></a>

## Direct properties — xfcc_disabled / 101310220212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2023133231323303-2012001330122313-2200103333030330-0011121210313010-0122123203212202-2011202110233122-3303103003202021-2002012133211003"></a>

## Next pages — xfcc_disabled / 101310220212 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-007.md#canonical-0021023101320020-0312123020002010-2321023022121310-2313220310303001-0131310002012320-1232302030303303-1232020021321313-2313013010220230)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-0313113021102333-0001033033120201-2221111132100321-0120033003012012-2022223101220022-3330230223013023-3322013300320223-3000211100123010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1313100112102011-1313320232103232-0220130111102023-1021131331210121-3221013003313313-0311203010312022-1021001031310002-3230110323322303"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options — xfcc_options / 001021030122 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert](resources--workload--reference--group-007.md#canonical-3121331321113110-2211010202311133-0013031020303103-3021013112113312-1002110313333221-0013020011222112-1231032331210100-0332331122030312)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-007.md#canonical-0021023101320020-0312123020002010-2321023022121310-2313220310303001-0131310002012320-1232302030303303-1232020021321313-2313013010220230)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls.xfcc_options

<a id="canonical-2223222212002233-2310112010203311-1133332112313122-1001013321021330-1022333221301332-0121010211131122-3333211020211030-3212002303331220"></a>

Type: `"object"`. single nested block, Optional.

X-Forwarded-Client-Cert header elements to be added to requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("xfcc_header_elements")}
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
xfcc_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-0130010220122103-1122030122101113-1111100320001011-1213212033033202-0232010232333021-1312210323222211-0233113022210103-2221302011130321"></a>

## Direct properties — xfcc_options / 001021030122 / 3

<a id="canonical-1101111110132231-0330231032110211-1112023223222202-3320212212101131-3232303002320223-0332112001010300-1010120110213222-0012331321003132"></a>

<a id="canonical-3123013313102133-3231322122233202-3220000202202211-2210311232302022-0322213311100132-0100022121033132-0010021333122332-2300011030022123"></a>

## xfcc_header_elements property — xfcc_options / 001021030122 / 4

Type: `["list", "string"]`. Optional.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

Upstream description:

X-Forwarded-Client-Cert header elements to be added to requests.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]"
  }
}
```

<a id="canonical-3321310030101220-2111300133000313-3203020313021111-1303300012120112-2010031120221003-3010021120001232-2132200322320303-0311320202321032"></a>

## Next pages — xfcc_options / 001021030122 / 5

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.https_auto_cert.use_mtls](resources--workload--reference--group-007.md#canonical-0021023101320020-0312123020002010-2321023022121310-2313220310303001-0131310002012320-1232302030303303-1232020021321313-2313013010220230)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3011211302110102-1110003330130012-2330003231002123-1333212203033211-1212101000211203-3333313212202213-1122232332203332-0220110332323232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101000322210023-1012323321311133-1001221313332121-2022013022303223-3320203323232203-0200212220111202-2102110021333223-2220202203300212"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes — specific_routes / 323001323132 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes

<a id="canonical-1220233233300032-0113322210322011-0123210001203323-3232103110313331-3330130123333031-1010112103100003-2011332201212333-2120002321201123"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to define a route.

Upstream description:

This defines various OPTIONS to define a route.

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
specific_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-0113002302101331-1203330332303203-2112232302323121-1133103002121320-1011221013233230-2132131332221013-3013002021110203-2233010213211131"></a>

## Direct properties — specific_routes / 323001323132 / 3

- [routes](resources--workload--reference--group-007.md#canonical-3101223212120012-3332300223001100-0333220310012333-1031332113331120-2210213133122022-1333032223122231-2030231210333100-1103221131133311): complete subsection reference.

<a id="canonical-0222110011222220-1330210300233200-1003113030101012-2002302003210210-0302210003003230-0321012112221321-1321033022100123-0203312221020022"></a>

## Next pages — specific_routes / 323001323132 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-007.md#canonical-3101223212120012-3332300223001100-0333220310012333-1031332113331120-2210213133122022-1333032223122231-2030231210333100-1103221131133311)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3101223212120012-3332300223001100-0333220310012333-1031332113331120-2210213133122022-1333032223122231-2030231210333100-1103221131133311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310221200033313-0113301012303323-1000120311123330-3032330121123321-0311003010331312-0030322211311103-3100313303303112-1020220203323213"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes — routes / 221031120321 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-007.md#canonical-3011211302110102-1110003330130012-2330003231002123-1333212203033211-1212101000211203-3333313212202213-1122232332203332-0220110332323232)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes

<a id="canonical-0320120233011301-2200031330002110-3221121311302333-3003121232230310-3332323122323322-2022332313211203-2001121302330031-1311110103303222"></a>

Type: `"object"`. list nested block, Optional.

Routes. Routes for this loadbalancer.

Upstream description:

Routes for this loadbalancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("custom_route_object",
    "direct_response_route"),
  validators.ConflictingListObjectAttributes("custom_route_object",
    "redirect_route"),
  validators.ConflictingListObjectAttributes("custom_route_object",
    "simple_route"),
  validators.ConflictingListObjectAttributes("direct_response_route",
    "redirect_route"),
  validators.ConflictingListObjectAttributes("direct_response_route",
    "simple_route"),
  validators.ConflictingListObjectAttributes("redirect_route",
    "simple_route")}
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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

Terraform syntax:

```terraform
routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-1122330302013211-0130322300303331-0211131323021112-2112101303013333-3123320001301120-1102321112103023-0003033332231120-3302020201323130"></a>

## Direct properties — routes / 221031120321 / 3

- [custom_route_object](resources--workload--reference--group-007.md#canonical-2000122322012020-3202320303320213-0231013331320221-2023031002230033-2112230020111131-0300200113021323-3020232330312200-0223013321012112): complete subsection reference.

- [direct_response_route](resources--workload--reference--group-008.md#canonical-3231322330202023-0011210010031213-0320313110113210-0323302203013003-1003131123333113-2031030012023021-3201100012202203-0030301312131222): complete subsection reference.

- [redirect_route](resources--workload--reference--group-008.md#canonical-0003020102311303-0233233032322332-3002231010033312-1313303221323312-0103013131011120-3102111011102132-1313020300032233-1111311233230113): complete subsection reference.

- [simple_route](resources--workload--reference--group-008.md#canonical-2332230303301330-0120223130133121-2131103033210110-0000331031322110-2102032123012022-0000121011033231-1222320000321202-3323210133011002): complete subsection reference.

<a id="canonical-3031331323210302-3220122011020233-3121312303113123-0301132101013132-0133110120131120-2102020303002132-0120120312320310-2230002103220003"></a>

## Next pages — routes / 221031120321 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-007.md#canonical-2000122322012020-3202320303320213-0231013331320221-2023031002230033-2112230020111131-0300200113021323-3020232330312200-0223013321012112)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.direct_response_route](resources--workload--reference--group-008.md#canonical-3231322330202023-0011210010031213-0320313110113210-0323302203013003-1003131123333113-2031030012023021-3201100012202203-0030301312131222)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.redirect_route](resources--workload--reference--group-008.md#canonical-0003020102311303-0233233032322332-3002231010033312-1313303221323312-0103013131011120-3102111011102132-1313020300032233-1111311233230113)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.simple_route](resources--workload--reference--group-008.md#canonical-2332230303301330-0120223130133121-2131103033210110-0000331031322110-2102032123012022-0000121011033231-1222320000321202-3323210133011002)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-007.md#canonical-3011211302110102-1110003330130012-2330003231002123-1333212203033211-1212101000211203-3333313212202213-1122232332203332-0220110332323232)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2000122322012020-3202320303320213-0231013331320221-2023031002230033-2112230020111131-0300200113021323-3020232330312200-0223013321012112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200032221212301-0002223032022301-3231323220302100-3023222222100133-2322030302031101-1320220222110013-0132110212122031-0123002222131110"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object — custom_route_object / 333310232001 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-007.md#canonical-3011211302110102-1110003330130012-2330003231002123-1333212203033211-1212101000211203-3333313212202213-1122232332203332-0220110332323232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-007.md#canonical-3101223212120012-3332300223001100-0333220310012333-1031332113331120-2210213133122022-1333032223122231-2030231210333100-1103221131133311)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object

<a id="canonical-1213010101330121-1132211303120323-2212112310000322-0211021200113210-1032301122302013-0020112303011230-0001012023302322-2121330103223001"></a>

Type: `"object"`. single nested block, Optional.

Custom route uses a route object created outside of this view.

Upstream description:

A custom route uses a route object created outside of this view.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("caching_disable",
    "caching_inherit")}
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
  "x-ves-oneof-field-caching": "[\"caching_disable\",\"caching_inherit\"]"
}
```

Terraform syntax:

```terraform
custom_route_object {
  # Configure direct properties listed below.
}
```

<a id="canonical-1102213011201212-1032231033132302-2032002213210012-3130101212011233-0222031013203202-1232002131032310-2200100011001001-0230123221313311"></a>

## Direct properties — custom_route_object / 333310232001 / 3

- [caching_disable](resources--workload--reference--group-007.md#canonical-2013012230220021-3011122033311001-0101101133003320-1011321003233233-1110102020203011-3021221203323112-2313303310222023-0000221032002120): complete subsection reference.

- [caching_inherit](resources--workload--reference--group-007.md#canonical-3230112111031231-2002111310330022-0223003233132003-2212133100121201-3032223120100312-2001123022122023-2121112123003212-3102202133020102): complete subsection reference.

- [route_ref](resources--workload--reference--group-008.md#canonical-3113101301231330-1201332333231130-2122102231230121-2033310222131210-0021302132233220-0032213131220003-2201110210212221-1233123001332220): complete subsection reference.

<a id="canonical-1123111021121332-0130000013001232-3300121212000012-2120202201133322-1230022032331311-3203313023301332-3222231121323202-0132320200132133"></a>

## Next pages — custom_route_object / 333310232001 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable](resources--workload--reference--group-007.md#canonical-2013012230220021-3011122033311001-0101101133003320-1011321003233233-1110102020203011-3021221203323112-2313303310222023-0000221032002120)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_inherit](resources--workload--reference--group-007.md#canonical-3230112111031231-2002111310330022-0223003233132003-2212133100121201-3032223120100312-2001123022122023-2121112123003212-3102202133020102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.route_ref](resources--workload--reference--group-008.md#canonical-3113101301231330-1201332333231130-2122102231230121-2033310222131210-0021302132233220-0032213131220003-2201110210212221-1233123001332220)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-007.md#canonical-3101223212120012-3332300223001100-0333220310012333-1031332113331120-2210213133122022-1333032223122231-2030231210333100-1103221131133311)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-2013012230220021-3011122033311001-0101101133003320-1011321003233233-1110102020203011-3021221203323112-2313303310222023-0000221032002120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130121311331323-0300323230303103-0332111023323013-3023302003113301-3210002212323002-1221233000001030-3233012021320022-3231223311301020"></a>

## service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable — caching_disable / 023031211002 / 2

Breadcrumbs:

- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)
- [Property reference](resources--workload--reference--group-001.md#canonical-2012112310002012-0330033330301132-3320113331023022-3021220030311012-0303221211130302-2102311022120112-1230121320333310-1232212220033130)
- [service](resources--workload--reference--group-005.md#canonical-0320132022221002-1121000220331020-2101233311220001-3200320202013031-2313301031031233-1332333320121020-1312300001112023-0212231303200011)
- [service.advertise_options](resources--workload--reference--group-005.md#canonical-2321101321310001-0022332320310221-1010303010131020-3231233110321220-1121213033111131-2021113020233203-2010000000223003-2113200202030130)
- [service.advertise_options.advertise_custom](resources--workload--reference--group-005.md#canonical-2121230312030330-3010113231323001-1233013032320322-1032203002210311-2022021002113232-1123312032112001-0012103213230102-1112311001110020)
- [service.advertise_options.advertise_custom.ports](resources--workload--reference--group-005.md#canonical-0133221101333210-0222103021001232-3212010032013103-3331120032012221-1012203003232110-1233102323001233-2233321121201101-0122313122210200)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer](resources--workload--reference--group-005.md#canonical-2311001323203002-2023333020201320-1330320322101233-1033123102030111-2013210002333120-2130122003123103-1133301323123220-2302130221012102)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes](resources--workload--reference--group-007.md#canonical-3011211302110102-1110003330130012-2330003231002123-1333212203033211-1212101000211203-3333313212202213-1122232332203332-0220110332323232)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes](resources--workload--reference--group-007.md#canonical-3101223212120012-3332300223001100-0333220310012333-1031332113331120-2210213133122022-1333032223122231-2030231210333100-1103221131133311)
- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-007.md#canonical-2000122322012020-3202320303320213-0231013331320221-2023031002230033-2112230020111131-0300200113021323-3020232330312200-0223013321012112)
- service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object.caching_disable

<a id="canonical-3101000232011003-1012310213111031-2311001000121111-2022312233213323-0002322333112000-2323211101031320-0102233320331323-2312021121110303"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for caching disable.

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
caching_disable = {}
```

<a id="canonical-2122010332230022-3313322113101232-2302322120313123-3113013321110323-3003332312233111-2311331030132300-3320111101221022-2222302121311231"></a>

## Direct properties — caching_disable / 023031211002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0011101221121000-2213220010320333-2130031102131220-1101321323021330-2210332302303311-2013230230313321-0110130130102023-2002210223221032"></a>

## Next pages — caching_disable / 023031211002 / 4

- [service.advertise_options.advertise_custom.ports.http_loadbalancer.specific_routes.routes.custom_route_object](resources--workload--reference--group-007.md#canonical-2000122322012020-3202320303320213-0231013331320221-2023031002230033-2112230020111131-0300200113021323-3020232330312200-0223013321012112)
- [xcsh_workload](../resources/workload.md#canonical-0311312121011133-1003302211333333-2022302011300003-2003003101223313-0230001021323223-1120310231002123-2102002120200113-3003130201212322)

<a id="canonical-3230112111031231-2002111310330022-0223003233132003-2212133100121201-3032223120100312-2001123022122023-2121112123003212-3102202133020102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
