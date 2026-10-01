---
page_title: "xcsh_protected_application reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_application reference."
---

# xcsh_protected_application reference

<a id="canonical-0123323202200001-0012112201110323-3000302310113000-0123203312123113-0320220133101203-0230301232310110-2013110300322131-2330133313210000"></a>

## Cloudflare.protected_endpoints.any_domain — any_domain / 122023332222 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- Cloudflare.protected_endpoints.any_domain

<a id="canonical-0113102131232031-1312323012002322-1132303201201221-3112033330031330-2322123200120211-3010223033010213-2202201103221202-0300322302303120"></a>

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
any_domain = {}
```

<a id="canonical-3322030132202300-0200312212021301-1211303201123200-3000212033030231-3312312221010202-1211132333301102-1031003013210113-2200101012223202"></a>

## Direct properties — any_domain / 122023332222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3001010233032013-0002221132313202-3301032232023002-3030331223313310-2013100300013132-0133220200103213-0020311020121301-3010201101022320"></a>

## Next pages — any_domain / 122023332222 / 4

- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-3022333301023332-2020221211122132-2220331300033002-3210230122323122-0301223130033002-1110232320103200-2132211310223201-0323000301133331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103333332232330-3203312020032023-3330222213102133-1230312230102212-3131222113112201-0132131023103310-2211122322010301-3312011223203101"></a>

## Cloudflare.protected_endpoints.domain — domain / 002231132023 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- Cloudflare.protected_endpoints.domain

<a id="canonical-1333302022232233-2120300310110333-2220012103022002-2010100130123120-1323012012121011-0002302021300122-0223033233300002-2210302213130322"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-3332001220011212-2003331100113010-3000320231120103-2103320121311130-3201000222033030-0313000213023020-3131202002321122-0202330322333101"></a>

## Direct properties — domain / 002231132023 / 3

<a id="canonical-3101013123133023-0012221223200320-0201132031322323-2123211232320231-3110303232233201-2132030232231303-2300033200030011-0000203012112223"></a>

<a id="canonical-2321013221130032-2232001323101002-2210313030221110-2023113333123222-0102310113332120-0321033301303013-3232323311311120-3302231211322111"></a>

## exact_value property — domain / 002231132023 / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3211203212221110-1033223001113201-1312122322021113-2103102030012110-1103130323232123-3211311230102313-1321113302001013-3101133021221013"></a>

<a id="canonical-3303301313300232-3230132212112003-3320111220111010-0322222113213330-1223111003303203-2231321332322123-3002211312102230-0200031300331223"></a>

## regex_value property — domain / 002231132023 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-0330230120020223-0220212123012221-2333301210133030-3112330000330020-1101312101300311-2303201222213032-2230033312101023-1023123130231212"></a>

<a id="canonical-0311322003222000-3211233130323012-2301120121102311-1230131303110232-0112322201221230-2300122213020323-2300323322311202-1303213203323320"></a>

## suffix_value property — domain / 002231132023 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0023021323330120-3230331012310201-0122312203231310-1102232032023122-0213023123122031-2233220110312220-3223102331202001-3232032321032212"></a>

## Next pages — domain / 002231132023 / 7

- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-0313013122021223-2132102100230101-1111201212133020-3103013102302032-0010021223220331-1202010100113100-3003121222312133-2303311000323323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320032303100211-3210311132132010-0123020312210230-2101000103001302-0200202330021032-2232022110101220-2201101002223121-1103222203003232"></a>

## Cloudflare.protected_endpoints.metadata — metadata / 100011321311 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- Cloudflare.protected_endpoints.metadata

<a id="canonical-0012232122120332-0212123200120100-3203013212301131-1331123202113111-0303102111122322-1220121123011023-3033001020132013-0202020101220212"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-2001032101300013-1332000223121201-0210313013111202-3112113330122312-1301211133133122-3332332331310220-1012322302032212-0101230321210200"></a>

## Direct properties — metadata / 100011321311 / 3

<a id="canonical-3221122320311101-3102002030133112-3203201021321111-2022133332000202-3330213320012312-2120022311120320-2230320003003331-1110033312203132"></a>

<a id="canonical-3012222133211301-0323100230212301-1203201203321032-1312002121320113-0230223003131011-3022020101333113-1233311131120002-1112303001021033"></a>

## description_spec property — metadata / 100011321311 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-2030311322033323-0021123213031010-1021020331002310-3001030102130330-1001003230032023-2021311000322220-0203121303211201-2200220221113023"></a>

<a id="canonical-3200320010302333-0230132312003331-0132200331321203-0322220022311133-1121001103203010-2132013131002122-1111231121201020-3333133003231201"></a>

## name property — metadata / 100011321311 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-1031332012302022-2311301212102110-0331302231301311-3203323210312120-1001030313000103-0003011030133330-0121002221332031-0320211210102023"></a>

## Next pages — metadata / 100011321311 / 6

- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-2231203002122332-0001110313310323-3123313122330323-0133123113300301-0302002231331030-1132302003033313-1122331331111221-2320131231013033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1321110012200030-1020032132012133-2302021123233233-3122313221111120-3122110023303201-2332000203121011-0310321333320130-0023133131103321"></a>

## Cloudflare.protected_endpoints.mobile_client — mobile_client / 222110021212 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- Cloudflare.protected_endpoints.mobile_client

<a id="canonical-2133232322011301-2100323220222002-3110231121020313-1033120023012213-1122032001200220-0122213120003103-3223303323201123-3332133221231032"></a>

Type: `"object"`. single nested block, Optional.

Mobile Client. Mobile client configuration OPTIONS.

Upstream description:

Mobile client configuration OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("block",
    "continue")}
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
  "x-ves-oneof-field-mitigation": "[\"block\",\"continue\"]"
}
```

Terraform syntax:

```terraform
mobile_client {
  # Configure direct properties listed below.
}
```

<a id="canonical-2102110301213233-1200030300213120-0130311130121011-3003303331210110-0320222032020202-2303113300312030-0033100213133222-3322013111123020"></a>

## Direct properties — mobile_client / 222110021212 / 3

- [block](resources--protected_application--reference--group-002.md#canonical-2011123200003121-0030302210222020-2122020332212021-3210220311102131-0230112002200311-0203313000002221-1312023133331312-0131023301211131): complete subsection reference.

- [continue](resources--protected_application--reference--group-002.md#canonical-0112302133133110-0322331130001100-2111220203133221-0003231332303003-3123310333300212-0323311001332010-0033111200012101-1333131131321120): complete subsection reference.

<a id="canonical-1113212113201312-2013310131001302-3110122121230223-1022233130210112-1102112331031312-0202210332213233-3013301100133220-2232122222213320"></a>

## Next pages — mobile_client / 222110021212 / 4

- [cloudflare.protected_endpoints.mobile_client.block](resources--protected_application--reference--group-002.md#canonical-2011123200003121-0030302210222020-2122020332212021-3210220311102131-0230112002200311-0203313000002221-1312023133331312-0131023301211131)
- [cloudflare.protected_endpoints.mobile_client.continue](resources--protected_application--reference--group-002.md#canonical-0112302133133110-0322331130001100-2111220203133221-0003231332303003-3123310333300212-0323311001332010-0033111200012101-1333131131321120)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-2011123200003121-0030302210222020-2122020332212021-3210220311102131-0230112002200311-0203313000002221-1312023133331312-0131023301211131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003010012222112-2013021011010030-0132033030131122-3211101220220311-2102222112113111-2313331013310122-1323031121021223-3113313112231301"></a>

## Cloudflare.protected_endpoints.mobile_client.block — block / 300003313303 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [cloudflare.protected_endpoints.mobile_client](resources--protected_application--reference--group-002.md#canonical-2231203002122332-0001110313310323-3123313122330323-0133123113300301-0302002231331030-1132302003033313-1122331331111221-2320131231013033)
- Cloudflare.protected_endpoints.mobile_client.block

<a id="canonical-3313112003013223-3123220113312120-0222223320111203-0311211123120220-3032123231230233-1010213222120213-3101112232233223-0200202012131201"></a>

Type: `"object"`. single nested block, Optional.

Block Response for Mobile. Block Response.

Upstream description:

Block Response.

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
block {
  # Configure direct properties listed below.
}
```

<a id="canonical-3332223120121132-1222020021101302-0203121221101220-1013313022303121-3211112122020031-1022303120331311-2012212211320201-0332201021031330"></a>

## Direct properties — block / 300003313303 / 3

<a id="canonical-3132003201111223-3100312001230332-3001310012210021-3233101110033021-0220323100011111-1001333131121032-3131223320211101-1133202132303032"></a>

<a id="canonical-1221311200030022-3322101111202232-0033032210102103-0310323011221033-3021223001112211-1033221203222023-2210021102022020-0110013010023311"></a>

## body property — block / 300003313303 / 4

Type: `"string"`. Optional.

Body. Custom body message.

Upstream description:

Custom body message.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(4096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 4096,
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
    "ves.io.schema.rules.string.max_len": "4096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096"
  }
}
```

<a id="canonical-0333303101323332-3230313221321033-2102011022011123-3211330320230323-3110310202322030-2231100112030131-0002120031203110-1133112303333213"></a>

<a id="canonical-1331321313032220-3332313220330220-3200311113320101-1333210320011001-1301002203210023-3032032322132000-0122201321012013-2130130102213020"></a>

## content_type property — block / 300003313303 / 5

Type: `"string"`. Optional.

Content type to use in a block response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-1213312213002123-0001030201320001-0032122233220021-2230020030020001-2022112201321110-2012220331332131-2303101223021100-1003213111013230"></a>

<a id="canonical-1211311110312123-2302023013213313-3210003122112301-2211020030211232-3310333210002211-2021232113112002-3313202103323323-1223013131232031"></a>

## status property — block / 300003313303 / 6

Type: `"string"`. Optional.

\[Enum:
EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired\]
HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status
code OK status code Created status code Accepted status code Non Authoritative Information status
code No Content status code Reset Content status code Partial Content status code Multi Status..
Possible values are \`EmptyStatusCode\`, \`Continue\`, \`OK\`, \`Created\`, \`Accepted\`,
\`NonAuthoritativeInformation\`, \`NoContent\`, \`ResetContent\`, \`PartialContent\`,
\`MultiStatus\`, \`AlreadyReported\`, \`IMUsed\`, \`MultipleChoices\`, \`MovedPermanently\`,
\`Found\`, \`SeeOther\`, \`NotModified\`, \`UseProxy\`, \`TemporaryRedirect\`,
\`PermanentRedirect\`, \`BadRequest\`, \`Unauthorized\`, \`PaymentRequired\`, \`Forbidden\`,
\`NotFound\`, \`MethodNotAllowed\`, \`NotAcceptable\`, \`ProxyAuthenticationRequired\`,
\`RequestTimeout\`, \`Conflict\`, \`Gone\`, \`LengthRequired\`, \`PreconditionFailed\`,
\`PayloadTooLarge\`, \`URITooLong\`, \`UnsupportedMediaType\`, \`RangeNotSatisfiable\`,
\`ExpectationFailed\`, \`MisdirectedRequest\`, \`UnprocessableEntity\`, \`Locked\`,
\`FailedDependency\`, \`UpgradeRequired\`, \`PreconditionRequired\`, \`TooManyRequests\`,
\`RequestHeaderFieldsTooLarge\`, \`InternalServerError\`, \`NotImplemented\`, \`BadGateway\`,
\`ServiceUnavailable\`, \`GatewayTimeout\`, \`HTTPVersionNotSupported\`, \`VariantAlsoNegotiates\`,
\`InsufficientStorage\`, \`LoopDetected\`, \`NotExtended\`, \`NetworkAuthenticationRequired\`.
Defaults to \`EmptyStatusCode\`.

Upstream description:

HTTP response status codes

EmptyStatusCode response codes means it is not specified Continue status code OK status code Created
status code Accepted status code Non Authoritative Information status code No Content status code
Reset Content status code Partial Content status code Multi Status status code Already Reported
status code Im Used status code Multiple Choices status code Moved Permanently status code Found
status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect
status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment
Required status code Forbidden status code Not Found status code Method Not Allowed status code Not
Acceptable status code Proxy Authentication Required status code Request Timeout status code
Conflict status code Gone status code Length Required status code Precondition Failed status code
Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not
Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable
Entity status code Locked status code Failed Dependency status code Upgrade Required status code
Precondition Required status code Too Many Requests status code Request Header Fields Too Large
status code Internal Server Error status code Not Implemented status code Bad Gateway status code
Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code
Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not
Extended status code Network Authentication Required status code.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "EmptyStatusCode",
  "enum": [
    "EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1102012213123330-0110321331012332-1030110303313131-1121110233220121-1133313220301301-2120310133131011-0101102023111220-2010312133003132"></a>

## Next pages — block / 300003313303 / 7

- [cloudflare.protected_endpoints.mobile_client](resources--protected_application--reference--group-002.md#canonical-2231203002122332-0001110313310323-3123313122330323-0133123113300301-0302002231331030-1132302003033313-1122331331111221-2320131231013033)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-0112302133133110-0322331130001100-2111220203133221-0003231332303003-3123310333300212-0323311001332010-0033111200012101-1333131131321120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0120002233013213-3320022203100220-1322210112112001-3010313331220233-0210322130102130-3121030212000002-1213311102320230-0023222231122102"></a>

## Cloudflare.protected_endpoints.mobile_client.continue — continue / 200311123013 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [cloudflare.protected_endpoints.mobile_client](resources--protected_application--reference--group-002.md#canonical-2231203002122332-0001110313310323-3123313122330323-0133123113300301-0302002231331030-1132302003033313-1122331331111221-2320131231013033)
- Cloudflare.protected_endpoints.mobile_client.continue

<a id="canonical-1112321121120010-3300310011221011-2002002331301012-3101101332033303-3030301112310132-2030003032113333-0021100020032323-2320013312333332"></a>

Type: `"object"`. single nested block, Optional.

Select Continue Bot Mitigation Action. Continue mitigation action.

Upstream description:

Continue mitigation action.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("add_header",
    "no_header")}
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
  "x-ves-oneof-field-add_header_choice": "[\"add_header\",\"no_header\"]"
}
```

Terraform syntax:

```terraform
continue {
  # Configure direct properties listed below.
}
```

<a id="canonical-2322002101232211-0321200030101121-0111123013110111-0211333102010311-3313323103021000-0221201312232131-0303133002020130-3231132321030323"></a>

## Direct properties — continue / 200311123013 / 3

- [add_header](resources--protected_application--reference--group-002.md#canonical-1210200112200210-2102002030111232-0121311303301303-3122332120220223-3133113322321031-3022320011302312-3030332210313101-1001100030310200): complete subsection reference.

- [no_header](resources--protected_application--reference--group-002.md#canonical-2030231300103221-0112230220333132-3121300220003120-3310033131210020-0302313201321110-3200323013003200-1021220000012131-3230303321010330): complete subsection reference.

<a id="canonical-3321201332322011-0012201000000231-3120212112000023-2132123330101121-1101322210201231-1302213212003213-0101121223211011-3102233231131202"></a>

## Next pages — continue / 200311123013 / 4

- [cloudflare.protected_endpoints.mobile_client.continue.add_header](resources--protected_application--reference--group-002.md#canonical-1210200112200210-2102002030111232-0121311303301303-3122332120220223-3133113322321031-3022320011302312-3030332210313101-1001100030310200)
- [cloudflare.protected_endpoints.mobile_client.continue.no_header](resources--protected_application--reference--group-002.md#canonical-2030231300103221-0112230220333132-3121300220003120-3310033131210020-0302313201321110-3200323013003200-1021220000012131-3230303321010330)
- [cloudflare.protected_endpoints.mobile_client](resources--protected_application--reference--group-002.md#canonical-2231203002122332-0001110313310323-3123313122330323-0133123113300301-0302002231331030-1132302003033313-1122331331111221-2320131231013033)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-1210200112200210-2102002030111232-0121311303301303-3122332120220223-3133113322321031-3022320011302312-3030332210313101-1001100030310200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3211122000003333-2020200020031311-3032303132332320-2222011332333330-3200232131320123-0102131212030333-3030232002320020-3321022333111332"></a>

## Cloudflare.protected_endpoints.mobile_client.continue.add_header — add_header / 030211321331 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [cloudflare.protected_endpoints.mobile_client](resources--protected_application--reference--group-002.md#canonical-2231203002122332-0001110313310323-3123313122330323-0133123113300301-0302002231331030-1132302003033313-1122331331111221-2320131231013033)
- [cloudflare.protected_endpoints.mobile_client.continue](resources--protected_application--reference--group-002.md#canonical-0112302133133110-0322331130001100-2111220203133221-0003231332303003-3123310333300212-0323311001332010-0033111200012101-1333131131321120)
- Cloudflare.protected_endpoints.mobile_client.continue.add_header

<a id="canonical-3111123130103230-3300331220322230-2001231120311310-3013132112223312-3311103111123023-3012230110321233-1332000000012110-2230013333212331"></a>

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
add_header = {}
```

<a id="canonical-0202333023213212-1013010001121120-3020011000121122-2323110332233213-1101201302321022-1102103312312110-1211302213130212-2311332102011210"></a>

## Direct properties — add_header / 030211321331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1313332012311102-1312012111200302-2230322300013110-2300120113231221-2100111023000311-1131000333020323-2103102122032301-3300231121213303"></a>

## Next pages — add_header / 030211321331 / 4

- [cloudflare.protected_endpoints.mobile_client.continue](resources--protected_application--reference--group-002.md#canonical-0112302133133110-0322331130001100-2111220203133221-0003231332303003-3123310333300212-0323311001332010-0033111200012101-1333131131321120)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-2030231300103221-0112230220333132-3121300220003120-3310033131210020-0302313201321110-3200323013003200-1021220000012131-3230303321010330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303301233210313-1102322332322123-3220022120030123-1113200310232231-3210230032013330-1301202033200131-0303011210202333-2123310101213023"></a>

## Cloudflare.protected_endpoints.mobile_client.continue.no_header — no_header / 333133132331 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [cloudflare.protected_endpoints.mobile_client](resources--protected_application--reference--group-002.md#canonical-2231203002122332-0001110313310323-3123313122330323-0133123113300301-0302002231331030-1132302003033313-1122331331111221-2320131231013033)
- [cloudflare.protected_endpoints.mobile_client.continue](resources--protected_application--reference--group-002.md#canonical-0112302133133110-0322331130001100-2111220203133221-0003231332303003-3123310333300212-0323311001332010-0033111200012101-1333131131321120)
- Cloudflare.protected_endpoints.mobile_client.continue.no_header

<a id="canonical-1330223321223301-3333020230312200-1110313331002033-1113033303130100-2220030013133031-2000010323313013-2223311110103110-3233230000210011"></a>

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
no_header = {}
```

<a id="canonical-1203203010300230-1211333110230230-3313220102032330-0321333011222222-2211212133122001-0232330330323201-3111323200012330-3301221100021110"></a>

## Direct properties — no_header / 333133132331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2123231132213303-2211032323110102-1111232301111221-2202003113123012-2310100231303310-2221010102203233-1303222113022212-3101321300201232"></a>

## Next pages — no_header / 333133132331 / 4

- [cloudflare.protected_endpoints.mobile_client.continue](resources--protected_application--reference--group-002.md#canonical-0112302133133110-0322331130001100-2111220203133221-0003231332303003-3123310333300212-0323311001332010-0033111200012101-1333131131321120)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-1001300212102331-3212330102111323-3100201020003211-0232102000201230-2211133203031111-3132133032303233-3312123120323022-0303222121321212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020210122102000-0320233003221330-3231001220122331-0110312200220003-0013101112101210-0233233312311330-2030210323133121-2212001111033300"></a>

## Cloudflare.protected_endpoints.path — path / 002002011213 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- Cloudflare.protected_endpoints.path

<a id="canonical-0333011131231110-0012122202023320-2001223310020313-3303022331330010-2202301330113223-3210000332102332-0211010130002132-1231332222231331"></a>

Type: `"object"`. single nested block, Optional.

Path. URI Path

Upstream description:

URI Path

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("path")}
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
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-0203013310121302-0330331130030312-2102023220311003-1111223100113300-2333103231010120-3011330202111001-1020020310220101-2011322200111022"></a>

## Direct properties — path / 002002011213 / 3

<a id="canonical-0333032133120321-2231201221112312-1001122231233123-2212022222233311-3321030130330130-0011100311301320-3311203021310221-0000111100130200"></a>

<a id="canonical-1311300022313310-0233103121012211-1110023132213122-2121232322223122-1212021111100133-1332220222120212-1300332302313001-2021000020103231"></a>

## caseinsensitive property — path / 002002011213 / 4

Type: `"bool"`. Optional.

Should path be searched case insensitive;.

Upstream description:

Should path be searched case insensitive;

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

<a id="canonical-1031231002332013-0103102322020032-2233331220213111-3323103000020321-1330320110130011-2012003031221313-2102220222323013-2121132200303001"></a>

<a id="canonical-2333022010320311-1020320011210112-3003021023201220-3002003113102132-3131100321233213-0223221132031223-3333121002002032-0131312103220130"></a>

## path property — path / 002002011213 / 5

Type: `"string"`. Optional.

Path. URI Path

Upstream description:

URI Path

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,999}$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,999}$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,999}$"
  }
}
```

<a id="canonical-0023330010200112-0000132332323220-1110221100332323-0110213100212332-3310133233322131-3121020222020321-2113323101101013-0022233201021223"></a>

## Next pages — path / 002002011213 / 6

- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-0213202232233011-0000312212000013-2113100002231110-2301000001213232-0031000313130131-1333121310313202-2201321230031021-0123023212011332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1013110231203301-3300113310332011-2020230022120113-1331301121301300-1021210032003332-0030220331323113-3112010100120323-3133023133203232"></a>

## Cloudflare.protected_endpoints.web_client — web_client / 220200131011 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- Cloudflare.protected_endpoints.web_client

<a id="canonical-3121310223103111-0301010320012013-3120110030132220-3202231230330102-1023232201201213-0320310200133333-1122201101321010-0333232030022222"></a>

Type: `"object"`. single nested block, Optional.

Web Client. Web client configuration OPTIONS.

Upstream description:

Web client configuration OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("block",
    "continue"),
  validators.ConflictingObjectAttributes("block",
    "redirect"),
  validators.ConflictingObjectAttributes("continue",
    "redirect")}
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
  "x-ves-oneof-field-mitigation": "[\"block\",\"continue\",\"redirect\"]"
}
```

Terraform syntax:

```terraform
web_client {
  # Configure direct properties listed below.
}
```

<a id="canonical-1130013232133100-0312011232121310-1223232001320113-1102113010212133-3222130220121321-2303103020321021-3003001023310131-3110121112211023"></a>

## Direct properties — web_client / 220200131011 / 3

- [block](resources--protected_application--reference--group-002.md#canonical-3112012233110303-1213332313103120-2021013130012232-1213300123323111-0033233013221231-2213110231312112-1122201022120013-3213300223331123): complete subsection reference.

- [continue](resources--protected_application--reference--group-002.md#canonical-2332001013223300-1121102021023132-2023200303321003-1132233223312202-0333211011333032-1321233331330202-2001010011100030-3023312112102220): complete subsection reference.

- [redirect](resources--protected_application--reference--group-002.md#canonical-0331132031332122-3211021101303101-1130201233222223-2210221222322200-0212103121001101-0022333331320010-1333322112111011-1111113123101332): complete subsection reference.

<a id="canonical-3233123220302210-1121302121223023-3200301321232110-2033011112003003-3311010331020233-1031001233303203-2211020101001112-2211321313222012"></a>

## Next pages — web_client / 220200131011 / 4

- [cloudflare.protected_endpoints.web_client.block](resources--protected_application--reference--group-002.md#canonical-3112012233110303-1213332313103120-2021013130012232-1213300123323111-0033233013221231-2213110231312112-1122201022120013-3213300223331123)
- [cloudflare.protected_endpoints.web_client.continue](resources--protected_application--reference--group-002.md#canonical-2332001013223300-1121102021023132-2023200303321003-1132233223312202-0333211011333032-1321233331330202-2001010011100030-3023312112102220)
- [cloudflare.protected_endpoints.web_client.redirect](resources--protected_application--reference--group-002.md#canonical-0331132031332122-3211021101303101-1130201233222223-2210221222322200-0212103121001101-0022333331320010-1333322112111011-1111113123101332)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-3112012233110303-1213332313103120-2021013130012232-1213300123323111-0033233013221231-2213110231312112-1122201022120013-3213300223331123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1230212022013021-3131013122120303-0011231033130212-2011103012131333-2330313033210011-1020033221302312-3011030331230022-1232323210113013"></a>

## Cloudflare.protected_endpoints.web_client.block — block / 321312202132 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [cloudflare.protected_endpoints.web_client](resources--protected_application--reference--group-002.md#canonical-0213202232233011-0000312212000013-2113100002231110-2301000001213232-0031000313130131-1333121310313202-2201321230031021-0123023212011332)
- Cloudflare.protected_endpoints.web_client.block

<a id="canonical-3113230031312132-1013333030110310-1103201101313101-2010020112300212-3221033123032102-0033012120323122-2132012332211031-1132012021211023"></a>

Type: `"object"`. single nested block, Optional.

Block Response. Block Response.

Upstream description:

Block Response.

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
block {
  # Configure direct properties listed below.
}
```

<a id="canonical-3110011011320011-2112011013210202-2310100122013301-3311103132300211-2003212131023212-3303101010321221-2330102021020003-2312011303101201"></a>

## Direct properties — block / 321312202132 / 3

<a id="canonical-1000303100301120-1313211001111013-3203111203313013-2010223223103301-0300221221100200-3013101311313000-1030032320321303-3200212200122001"></a>

<a id="canonical-3201113301003221-3301222233011000-2010013003101333-3303123120303233-2000130221113220-0221011103100033-1303033331321300-1232223301021022"></a>

## body property — block / 321312202132 / 4

Type: `"string"`. Optional.

Body. Custom body message.

Upstream description:

Custom body message.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(4096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 4096,
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
    "ves.io.schema.rules.string.max_len": "4096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096"
  }
}
```

<a id="canonical-1033011230312201-2203003033100320-3223111330102330-1201133210210331-1033202101120213-3213323121201311-2100221113300310-1201122212320333"></a>

<a id="canonical-1202321031323121-1000300232323121-0001330232203203-1130100132102210-1303123010210230-2010212011323212-3000111220000321-0013120133210201"></a>

## content_type property — block / 321312202132 / 5

Type: `"string"`. Optional.

Content type to use in a block response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-2211320110302003-3110131210202212-3111200323130010-1201030132030103-0220012122300213-3022212210331323-1303223030333102-3323030012231101"></a>

<a id="canonical-1212302212301321-0202030221013310-0200333032331232-0221311031212301-0223020111311101-0122023312102212-2333303312212331-2312313032223322"></a>

## status property — block / 321312202132 / 6

Type: `"string"`. Optional.

\[Enum:
EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired\]
HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status
code OK status code Created status code Accepted status code Non Authoritative Information status
code No Content status code Reset Content status code Partial Content status code Multi Status..
Possible values are \`EmptyStatusCode\`, \`Continue\`, \`OK\`, \`Created\`, \`Accepted\`,
\`NonAuthoritativeInformation\`, \`NoContent\`, \`ResetContent\`, \`PartialContent\`,
\`MultiStatus\`, \`AlreadyReported\`, \`IMUsed\`, \`MultipleChoices\`, \`MovedPermanently\`,
\`Found\`, \`SeeOther\`, \`NotModified\`, \`UseProxy\`, \`TemporaryRedirect\`,
\`PermanentRedirect\`, \`BadRequest\`, \`Unauthorized\`, \`PaymentRequired\`, \`Forbidden\`,
\`NotFound\`, \`MethodNotAllowed\`, \`NotAcceptable\`, \`ProxyAuthenticationRequired\`,
\`RequestTimeout\`, \`Conflict\`, \`Gone\`, \`LengthRequired\`, \`PreconditionFailed\`,
\`PayloadTooLarge\`, \`URITooLong\`, \`UnsupportedMediaType\`, \`RangeNotSatisfiable\`,
\`ExpectationFailed\`, \`MisdirectedRequest\`, \`UnprocessableEntity\`, \`Locked\`,
\`FailedDependency\`, \`UpgradeRequired\`, \`PreconditionRequired\`, \`TooManyRequests\`,
\`RequestHeaderFieldsTooLarge\`, \`InternalServerError\`, \`NotImplemented\`, \`BadGateway\`,
\`ServiceUnavailable\`, \`GatewayTimeout\`, \`HTTPVersionNotSupported\`, \`VariantAlsoNegotiates\`,
\`InsufficientStorage\`, \`LoopDetected\`, \`NotExtended\`, \`NetworkAuthenticationRequired\`.
Defaults to \`EmptyStatusCode\`.

Upstream description:

HTTP response status codes

EmptyStatusCode response codes means it is not specified Continue status code OK status code Created
status code Accepted status code Non Authoritative Information status code No Content status code
Reset Content status code Partial Content status code Multi Status status code Already Reported
status code Im Used status code Multiple Choices status code Moved Permanently status code Found
status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect
status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment
Required status code Forbidden status code Not Found status code Method Not Allowed status code Not
Acceptable status code Proxy Authentication Required status code Request Timeout status code
Conflict status code Gone status code Length Required status code Precondition Failed status code
Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not
Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable
Entity status code Locked status code Failed Dependency status code Upgrade Required status code
Precondition Required status code Too Many Requests status code Request Header Fields Too Large
status code Internal Server Error status code Not Implemented status code Bad Gateway status code
Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code
Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not
Extended status code Network Authentication Required status code.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "EmptyStatusCode",
  "enum": [
    "EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2223023320111331-1300313220112301-3033020022101130-2010122103300320-3231302220231020-0202212311030013-2113203101111112-2332103202222130"></a>

## Next pages — block / 321312202132 / 7

- [cloudflare.protected_endpoints.web_client](resources--protected_application--reference--group-002.md#canonical-0213202232233011-0000312212000013-2113100002231110-2301000001213232-0031000313130131-1333121310313202-2201321230031021-0123023212011332)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-2332001013223300-1121102021023132-2023200303321003-1132233223312202-0333211011333032-1321233331330202-2001010011100030-3023312112102220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311113010003322-1120033313231102-2212210113112331-3333032220321313-0023121331001302-3221112013333123-1220032123022331-1203110132232311"></a>

## Cloudflare.protected_endpoints.web_client.continue — continue / 131220233330 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [cloudflare.protected_endpoints.web_client](resources--protected_application--reference--group-002.md#canonical-0213202232233011-0000312212000013-2113100002231110-2301000001213232-0031000313130131-1333121310313202-2201321230031021-0123023212011332)
- Cloudflare.protected_endpoints.web_client.continue

<a id="canonical-0103301000301232-2102202133001231-3101010102310102-0132321123013312-2233113022313221-2322021020021312-2320120323301222-3010203201332102"></a>

Type: `"object"`. single nested block, Optional.

Select Continue Bot Mitigation Action. Continue mitigation action.

Upstream description:

Continue mitigation action.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("add_header",
    "no_header")}
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
  "x-ves-oneof-field-add_header_choice": "[\"add_header\",\"no_header\"]"
}
```

Terraform syntax:

```terraform
continue {
  # Configure direct properties listed below.
}
```

<a id="canonical-2203222322002020-2203221332023303-1211221111101123-3233312002123101-0223332021333002-3320333202301201-3011311130030132-3210100002032220"></a>

## Direct properties — continue / 131220233330 / 3

- [add_header](resources--protected_application--reference--group-002.md#canonical-1212021322320303-3230022010211001-0330201213203001-0122013012130331-2012233102113331-2212023101132322-2021101030321122-2223321213202312): complete subsection reference.

- [no_header](resources--protected_application--reference--group-002.md#canonical-1301121201301132-1110113100101213-1133303312330101-2030211022222133-0102033221132121-1001133101222112-1131311333302220-3230120013030303): complete subsection reference.

<a id="canonical-0221023332121313-0302222233231231-2211311023210012-1310333103032102-2322321031013211-1032230113132132-2311133321003033-2333233230132320"></a>

## Next pages — continue / 131220233330 / 4

- [cloudflare.protected_endpoints.web_client.continue.add_header](resources--protected_application--reference--group-002.md#canonical-1212021322320303-3230022010211001-0330201213203001-0122013012130331-2012233102113331-2212023101132322-2021101030321122-2223321213202312)
- [cloudflare.protected_endpoints.web_client.continue.no_header](resources--protected_application--reference--group-002.md#canonical-1301121201301132-1110113100101213-1133303312330101-2030211022222133-0102033221132121-1001133101222112-1131311333302220-3230120013030303)
- [cloudflare.protected_endpoints.web_client](resources--protected_application--reference--group-002.md#canonical-0213202232233011-0000312212000013-2113100002231110-2301000001213232-0031000313130131-1333121310313202-2201321230031021-0123023212011332)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-1212021322320303-3230022010211001-0330201213203001-0122013012130331-2012233102113331-2212023101132322-2021101030321122-2223321213202312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330121102130230-1301210203010223-1231021212302121-2202322201003003-1201210211012102-0200303310233303-3032333033230203-3201221330301111"></a>

## Cloudflare.protected_endpoints.web_client.continue.add_header — add_header / 130012221102 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [cloudflare.protected_endpoints.web_client](resources--protected_application--reference--group-002.md#canonical-0213202232233011-0000312212000013-2113100002231110-2301000001213232-0031000313130131-1333121310313202-2201321230031021-0123023212011332)
- [cloudflare.protected_endpoints.web_client.continue](resources--protected_application--reference--group-002.md#canonical-2332001013223300-1121102021023132-2023200303321003-1132233223312202-0333211011333032-1321233331330202-2001010011100030-3023312112102220)
- Cloudflare.protected_endpoints.web_client.continue.add_header

<a id="canonical-2321303303011121-3102333113320202-0102303311323302-1220232333322011-2222131000233023-3323030232210031-2230332230102130-0000021200033300"></a>

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
add_header = {}
```

<a id="canonical-3201303312202223-3020223120233303-3020300020213020-3011323122030120-0322013120120202-3210112030101331-0302022332012313-2003111333021323"></a>

## Direct properties — add_header / 130012221102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0120311323311010-3330001002222111-3131212322000113-2021020111032331-3112102000203303-3031031223232021-3123312331310310-3213102311132301"></a>

## Next pages — add_header / 130012221102 / 4

- [cloudflare.protected_endpoints.web_client.continue](resources--protected_application--reference--group-002.md#canonical-2332001013223300-1121102021023132-2023200303321003-1132233223312202-0333211011333032-1321233331330202-2001010011100030-3023312112102220)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-1301121201301132-1110113100101213-1133303312330101-2030211022222133-0102033221132121-1001133101222112-1131311333302220-3230120013030303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233020112212130-0033211211332002-1012322232011222-1023203220120130-1102330221121202-0301013220221003-3331123130123023-0121220032312230"></a>

## Cloudflare.protected_endpoints.web_client.continue.no_header — no_header / 023301130233 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [cloudflare.protected_endpoints.web_client](resources--protected_application--reference--group-002.md#canonical-0213202232233011-0000312212000013-2113100002231110-2301000001213232-0031000313130131-1333121310313202-2201321230031021-0123023212011332)
- [cloudflare.protected_endpoints.web_client.continue](resources--protected_application--reference--group-002.md#canonical-2332001013223300-1121102021023132-2023200303321003-1132233223312202-0333211011333032-1321233331330202-2001010011100030-3023312112102220)
- Cloudflare.protected_endpoints.web_client.continue.no_header

<a id="canonical-0311232032311101-3201331022132222-0311230333032322-0300332202203301-2201021130112012-3312200211101303-3123032023133000-1231202233203030"></a>

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
no_header = {}
```

<a id="canonical-0010001123310120-3310310101110310-0123021211230210-1121300301030231-1210001003210111-3211300203333023-0233110301232103-2201033111120231"></a>

## Direct properties — no_header / 023301130233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0011133100013000-3010011030113333-1220030321023232-1331133230111131-2033210221231030-0323102021200001-1300102100213102-0302123013132131"></a>

## Next pages — no_header / 023301130233 / 4

- [cloudflare.protected_endpoints.web_client.continue](resources--protected_application--reference--group-002.md#canonical-2332001013223300-1121102021023132-2023200303321003-1132233223312202-0333211011333032-1321233331330202-2001010011100030-3023312112102220)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-0331132031332122-3211021101303101-1130201233222223-2210221222322200-0212103121001101-0022333331320010-1333322112111011-1111113123101332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203213120022310-2201302132300023-2133323313321033-2121202323323102-0322031232232120-1202102031202003-2112031103023313-2111032223303310"></a>

## Cloudflare.protected_endpoints.web_client.redirect — redirect / 123322101110 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [cloudflare.protected_endpoints.web_client](resources--protected_application--reference--group-002.md#canonical-0213202232233011-0000312212000013-2113100002231110-2301000001213232-0031000313130131-1333121310313202-2201321230031021-0123023212011332)
- Cloudflare.protected_endpoints.web_client.redirect

<a id="canonical-2330331123021232-2310211101130032-3220122333323223-1212121020101113-3120103311003003-2322300032133210-2011013003012000-2331103030003211"></a>

Type: `"object"`. single nested block, Optional.

Redirect. Redirect.

Upstream description:

Redirect.

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
redirect {
  # Configure direct properties listed below.
}
```

<a id="canonical-0220323103022332-0333012311331010-1223121221213301-3130020132331302-1002213212312230-3323103120011131-3012301310012203-2231212102012122"></a>

## Direct properties — redirect / 123322101110 / 3

<a id="canonical-3303322023003220-0220321223113110-2132311020120323-3102233031222011-2022222132023003-1130120000030212-2200011212301301-1302323303310203"></a>

<a id="canonical-1203000031132101-3031333121203221-3310130322110230-2010331313230301-2102120303331211-0310230322012011-2322233300130310-3032130221331223"></a>

## location property — redirect / 123322101110 / 4

Type: `"string"`. Optional.

Location. URI location for redirect response.

Upstream description:

URI location for redirect response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 512
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_bytes": "512",
    "ves.io.schema.rules.string.url_or_uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "512",
    "ves.io.schema.rules.string.url_or_uri_ref": "true"
  }
}
```

<a id="canonical-3200300102212301-2232031211301231-2210311321000113-1133101333331113-3223322220001323-0332001323333222-3313001313000320-2232123133132102"></a>

<a id="canonical-0303330203110030-3111223333010302-1330032012131100-0003032213231232-1203333100332301-1010310002031131-1302231103313010-1112221230200122"></a>

## status property — redirect / 123322101110 / 5

Type: `"string"`. Optional.

\[Enum:
EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired\]
HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status
code OK status code Created status code Accepted status code Non Authoritative Information status
code No Content status code Reset Content status code Partial Content status code Multi Status..
Possible values are \`EmptyStatusCode\`, \`Continue\`, \`OK\`, \`Created\`, \`Accepted\`,
\`NonAuthoritativeInformation\`, \`NoContent\`, \`ResetContent\`, \`PartialContent\`,
\`MultiStatus\`, \`AlreadyReported\`, \`IMUsed\`, \`MultipleChoices\`, \`MovedPermanently\`,
\`Found\`, \`SeeOther\`, \`NotModified\`, \`UseProxy\`, \`TemporaryRedirect\`,
\`PermanentRedirect\`, \`BadRequest\`, \`Unauthorized\`, \`PaymentRequired\`, \`Forbidden\`,
\`NotFound\`, \`MethodNotAllowed\`, \`NotAcceptable\`, \`ProxyAuthenticationRequired\`,
\`RequestTimeout\`, \`Conflict\`, \`Gone\`, \`LengthRequired\`, \`PreconditionFailed\`,
\`PayloadTooLarge\`, \`URITooLong\`, \`UnsupportedMediaType\`, \`RangeNotSatisfiable\`,
\`ExpectationFailed\`, \`MisdirectedRequest\`, \`UnprocessableEntity\`, \`Locked\`,
\`FailedDependency\`, \`UpgradeRequired\`, \`PreconditionRequired\`, \`TooManyRequests\`,
\`RequestHeaderFieldsTooLarge\`, \`InternalServerError\`, \`NotImplemented\`, \`BadGateway\`,
\`ServiceUnavailable\`, \`GatewayTimeout\`, \`HTTPVersionNotSupported\`, \`VariantAlsoNegotiates\`,
\`InsufficientStorage\`, \`LoopDetected\`, \`NotExtended\`, \`NetworkAuthenticationRequired\`.
Defaults to \`EmptyStatusCode\`.

Upstream description:

HTTP response status codes

EmptyStatusCode response codes means it is not specified Continue status code OK status code Created
status code Accepted status code Non Authoritative Information status code No Content status code
Reset Content status code Partial Content status code Multi Status status code Already Reported
status code Im Used status code Multiple Choices status code Moved Permanently status code Found
status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect
status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment
Required status code Forbidden status code Not Found status code Method Not Allowed status code Not
Acceptable status code Proxy Authentication Required status code Request Timeout status code
Conflict status code Gone status code Length Required status code Precondition Failed status code
Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not
Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable
Entity status code Locked status code Failed Dependency status code Upgrade Required status code
Precondition Required status code Too Many Requests status code Request Header Fields Too Large
status code Internal Server Error status code Not Implemented status code Bad Gateway status code
Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code
Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not
Extended status code Network Authentication Required status code.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "EmptyStatusCode",
  "enum": [
    "EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0111200332111311-3211112313010011-3022033220231333-1022022220112200-0310201002323021-3302121330222113-1212210302130233-0200222103233133"></a>

## Next pages — redirect / 123322101110 / 6

- [cloudflare.protected_endpoints.web_client](resources--protected_application--reference--group-002.md#canonical-0213202232233011-0000312212000013-2113100002231110-2301000001213232-0031000313130131-1333121310313202-2201321230031021-0123023212011332)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-1113311232320333-0300120101320203-2110223102112213-2013121133211100-3321120132301211-2020003231031133-1323311232301303-1022020213311220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3110110010002220-1323122012003120-2331232111111103-0232312122032202-0331010101010230-3322130322331222-3312003302330233-3330323112021211"></a>

## Cloudflare.protected_endpoints.web_mobile_client — web_mobile_client / 332013100032 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- Cloudflare.protected_endpoints.web_mobile_client

<a id="canonical-2123131020201201-3112310030020222-1000120231311011-2121330012333300-0322020211303031-3333221011223102-3030032332030222-0120232031011120"></a>

Type: `"object"`. single nested block, Optional.

Web and Mobile client configuration OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("block_mobile",
    "continue_mobile"),
  validators.ConflictingObjectAttributes("block_web",
    "continue_web"),
  validators.ConflictingObjectAttributes("block_web",
    "redirect_web"),
  validators.ConflictingObjectAttributes("continue_web",
    "redirect_web")}
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
  "x-ves-oneof-field-mobile_mitigation": "[\"block_mobile\",\"continue_mobile\"]",
  "x-ves-oneof-field-web_mitigation": "[\"block_web\",\"continue_web\",\"redirect_web\"]"
}
```

Terraform syntax:

```terraform
web_mobile_client {
  # Configure direct properties listed below.
}
```

<a id="canonical-1013120311313331-3021232111222023-2120220122020112-1010231323321122-3011110031210113-1223021002212132-2020010120223230-2022112323331223"></a>

## Direct properties — web_mobile_client / 332013100032 / 3

- [block_mobile](resources--protected_application--reference--group-002.md#canonical-3011122022300310-2132312002222213-0001223030312021-1101320311023132-2133103130122310-3311010113322001-3312323102002103-3332020000013313): complete subsection reference.

- [block_web](resources--protected_application--reference--group-002.md#canonical-0300020103002300-3330212321220103-3000022131130233-3202203013212133-1031103010321203-0023213133313221-2131121330233012-2232323300132001): complete subsection reference.

- [continue_mobile](resources--protected_application--reference--group-002.md#canonical-1221231003113011-0022303133210330-2332321100311113-1033320032232322-3222122030013201-1202220202230113-0332003233200111-2213102220011131): complete subsection reference.

- [continue_web](resources--protected_application--reference--group-002.md#canonical-0102201103120331-1320010313110222-0323100032123012-2333110230202132-0110220130102010-0200322230322022-0213211311303332-1132200123313312): complete subsection reference.

- [redirect_web](resources--protected_application--reference--group-002.md#canonical-3223112032330233-3211013131233032-3321031030031233-1211001113012013-3010223031232211-2123312332113311-3100023303213011-0010021220103331): complete subsection reference.

<a id="canonical-1321103212010101-2213022123011201-0210303010231331-2330223121021000-0303031330110002-1003133222210103-2103311203310013-3311020331112111"></a>

## Next pages — web_mobile_client / 332013100032 / 4

- [cloudflare.protected_endpoints.web_mobile_client.block_mobile](resources--protected_application--reference--group-002.md#canonical-3011122022300310-2132312002222213-0001223030312021-1101320311023132-2133103130122310-3311010113322001-3312323102002103-3332020000013313)
- [cloudflare.protected_endpoints.web_mobile_client.block_web](resources--protected_application--reference--group-002.md#canonical-0300020103002300-3330212321220103-3000022131130233-3202203013212133-1031103010321203-0023213133313221-2131121330233012-2232323300132001)
- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile](resources--protected_application--reference--group-002.md#canonical-1221231003113011-0022303133210330-2332321100311113-1033320032232322-3222122030013201-1202220202230113-0332003233200111-2213102220011131)
- [cloudflare.protected_endpoints.web_mobile_client.continue_web](resources--protected_application--reference--group-002.md#canonical-0102201103120331-1320010313110222-0323100032123012-2333110230202132-0110220130102010-0200322230322022-0213211311303332-1132200123313312)
- [cloudflare.protected_endpoints.web_mobile_client.redirect_web](resources--protected_application--reference--group-002.md#canonical-3223112032330233-3211013131233032-3321031030031233-1211001113012013-3010223031232211-2123312332113311-3100023303213011-0010021220103331)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-3011122022300310-2132312002222213-0001223030312021-1101320311023132-2133103130122310-3311010113322001-3312323102002103-3332020000013313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0022011200221121-1130130023013003-0100133032132320-1331031123200022-1213332121122300-1223010021120333-1002011200100201-3331033311321032"></a>

## Cloudflare.protected_endpoints.web_mobile_client.block_mobile — block_mobile / 231312131001 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-1113311232320333-0300120101320203-2110223102112213-2013121133211100-3321120132301211-2020003231031133-1323311232301303-1022020213311220)
- Cloudflare.protected_endpoints.web_mobile_client.block_mobile

<a id="canonical-0133011313331013-1123020310131212-1220023213112000-2210111021201323-1011301000302311-0112002121211201-1331110331131223-3232113110202010"></a>

Type: `"object"`. single nested block, Optional.

Block Response for Mobile. Block Response.

Upstream description:

Block Response.

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
block_mobile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0113133100330011-2103310231322102-2123022131211330-2213330102200122-2213030303022002-3021132103230003-2330020123002113-2230212300201121"></a>

## Direct properties — block_mobile / 231312131001 / 3

<a id="canonical-0303122013133012-1311200103102331-1033010130233000-1033133110301332-3002110223131000-3001310232330030-0203323300323320-1330213020332333"></a>

<a id="canonical-2133123133222310-1031010031333333-0132033033013233-2303003222232211-1322023132033132-3133333332023003-2033100203212130-2133113110231301"></a>

## body property — block_mobile / 231312131001 / 4

Type: `"string"`. Optional.

Body. Custom body message.

Upstream description:

Custom body message.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(4096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 4096,
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
    "ves.io.schema.rules.string.max_len": "4096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096"
  }
}
```

<a id="canonical-0003011101233020-1102323323011113-0211200332231122-3333330221323123-1101333130213100-2103323322113101-0103021003113131-2031213203313112"></a>

<a id="canonical-2101131300021011-3312231331221031-2111302333232312-3313330333013100-3213200001232233-3131010022133110-0321300221303331-1213320022012313"></a>

## content_type property — block_mobile / 231312131001 / 5

Type: `"string"`. Optional.

Content type to use in a block response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-1111102302303232-0302333100310210-0102323001312231-3031110122222102-0312230133002222-0103321311201202-2022122321031121-3321220222120131"></a>

<a id="canonical-0011121311113121-1012302021201311-0103333010100012-0201110223301001-1200033010030330-2110201232131220-3112113330300200-1211301112121322"></a>

## status property — block_mobile / 231312131001 / 6

Type: `"string"`. Optional.

\[Enum:
EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired\]
HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status
code OK status code Created status code Accepted status code Non Authoritative Information status
code No Content status code Reset Content status code Partial Content status code Multi Status..
Possible values are \`EmptyStatusCode\`, \`Continue\`, \`OK\`, \`Created\`, \`Accepted\`,
\`NonAuthoritativeInformation\`, \`NoContent\`, \`ResetContent\`, \`PartialContent\`,
\`MultiStatus\`, \`AlreadyReported\`, \`IMUsed\`, \`MultipleChoices\`, \`MovedPermanently\`,
\`Found\`, \`SeeOther\`, \`NotModified\`, \`UseProxy\`, \`TemporaryRedirect\`,
\`PermanentRedirect\`, \`BadRequest\`, \`Unauthorized\`, \`PaymentRequired\`, \`Forbidden\`,
\`NotFound\`, \`MethodNotAllowed\`, \`NotAcceptable\`, \`ProxyAuthenticationRequired\`,
\`RequestTimeout\`, \`Conflict\`, \`Gone\`, \`LengthRequired\`, \`PreconditionFailed\`,
\`PayloadTooLarge\`, \`URITooLong\`, \`UnsupportedMediaType\`, \`RangeNotSatisfiable\`,
\`ExpectationFailed\`, \`MisdirectedRequest\`, \`UnprocessableEntity\`, \`Locked\`,
\`FailedDependency\`, \`UpgradeRequired\`, \`PreconditionRequired\`, \`TooManyRequests\`,
\`RequestHeaderFieldsTooLarge\`, \`InternalServerError\`, \`NotImplemented\`, \`BadGateway\`,
\`ServiceUnavailable\`, \`GatewayTimeout\`, \`HTTPVersionNotSupported\`, \`VariantAlsoNegotiates\`,
\`InsufficientStorage\`, \`LoopDetected\`, \`NotExtended\`, \`NetworkAuthenticationRequired\`.
Defaults to \`EmptyStatusCode\`.

Upstream description:

HTTP response status codes

EmptyStatusCode response codes means it is not specified Continue status code OK status code Created
status code Accepted status code Non Authoritative Information status code No Content status code
Reset Content status code Partial Content status code Multi Status status code Already Reported
status code Im Used status code Multiple Choices status code Moved Permanently status code Found
status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect
status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment
Required status code Forbidden status code Not Found status code Method Not Allowed status code Not
Acceptable status code Proxy Authentication Required status code Request Timeout status code
Conflict status code Gone status code Length Required status code Precondition Failed status code
Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not
Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable
Entity status code Locked status code Failed Dependency status code Upgrade Required status code
Precondition Required status code Too Many Requests status code Request Header Fields Too Large
status code Internal Server Error status code Not Implemented status code Bad Gateway status code
Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code
Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not
Extended status code Network Authentication Required status code.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "EmptyStatusCode",
  "enum": [
    "EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2001300211032322-1200203010021331-0121013312012222-2013220011301231-0213331312211311-0222010330303223-2230203312300210-2320330321122010"></a>

## Next pages — block_mobile / 231312131001 / 7

- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-1113311232320333-0300120101320203-2110223102112213-2013121133211100-3321120132301211-2020003231031133-1323311232301303-1022020213311220)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-0300020103002300-3330212321220103-3000022131130233-3202203013212133-1031103010321203-0023213133313221-2131121330233012-2232323300132001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202000310330003-3113313003211323-2111220033230000-1010231333120301-0303020333021201-3200022322223310-2213321331023333-3303010202323133"></a>

## Cloudflare.protected_endpoints.web_mobile_client.block_web — block_web / 332002032020 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-1113311232320333-0300120101320203-2110223102112213-2013121133211100-3321120132301211-2020003231031133-1323311232301303-1022020213311220)
- Cloudflare.protected_endpoints.web_mobile_client.block_web

<a id="canonical-2121130001221211-0013111003301112-2133210303201331-1133310022221301-1000321321310303-2311123121010010-2210200222132000-3030230201230212"></a>

Type: `"object"`. single nested block, Optional.

Block Response. Block Response.

Upstream description:

Block Response.

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
block_web {
  # Configure direct properties listed below.
}
```

<a id="canonical-1123023222111213-3033230003012011-2013002133013032-0003020222033123-3130231120003320-1202331033313032-1331230222232312-2132032332132012"></a>

## Direct properties — block_web / 332002032020 / 3

<a id="canonical-2302110221123123-3332211302310103-0121133000032001-2021002300103130-2002032303031301-1011021112301100-1111303211110012-2303302031001000"></a>

<a id="canonical-1021113332311023-3231002310113331-2300121223302113-0220010030202112-3302302010131003-2020120030011201-1023013122203223-3211321122013133"></a>

## body property — block_web / 332002032020 / 4

Type: `"string"`. Optional.

Body. Custom body message.

Upstream description:

Custom body message.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(4096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 4096,
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
    "ves.io.schema.rules.string.max_len": "4096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096"
  }
}
```

<a id="canonical-0230230222313103-3122022003233301-0100003013301013-1013112130311330-2123020110012301-1332123333202032-2133230201212221-2013231302223121"></a>

<a id="canonical-2112303031130131-2320322223323230-0101001313012202-2130032310230010-0300310221310302-2113302320320003-2233110231130321-2310101102331213"></a>

## content_type property — block_web / 332002032020 / 5

Type: `"string"`. Optional.

Content type to use in a block response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-3123331322020200-0030320132330200-3200132200233223-0130333321201231-0323233332223322-3013131003120010-3001232120301002-2303220332333203"></a>

<a id="canonical-3322020110330311-2312313331111003-3132000120223022-2102032312121310-0233202032103130-1012001301211002-3230002322113122-1201002310222001"></a>

## status property — block_web / 332002032020 / 6

Type: `"string"`. Optional.

\[Enum:
EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired\]
HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status
code OK status code Created status code Accepted status code Non Authoritative Information status
code No Content status code Reset Content status code Partial Content status code Multi Status..
Possible values are \`EmptyStatusCode\`, \`Continue\`, \`OK\`, \`Created\`, \`Accepted\`,
\`NonAuthoritativeInformation\`, \`NoContent\`, \`ResetContent\`, \`PartialContent\`,
\`MultiStatus\`, \`AlreadyReported\`, \`IMUsed\`, \`MultipleChoices\`, \`MovedPermanently\`,
\`Found\`, \`SeeOther\`, \`NotModified\`, \`UseProxy\`, \`TemporaryRedirect\`,
\`PermanentRedirect\`, \`BadRequest\`, \`Unauthorized\`, \`PaymentRequired\`, \`Forbidden\`,
\`NotFound\`, \`MethodNotAllowed\`, \`NotAcceptable\`, \`ProxyAuthenticationRequired\`,
\`RequestTimeout\`, \`Conflict\`, \`Gone\`, \`LengthRequired\`, \`PreconditionFailed\`,
\`PayloadTooLarge\`, \`URITooLong\`, \`UnsupportedMediaType\`, \`RangeNotSatisfiable\`,
\`ExpectationFailed\`, \`MisdirectedRequest\`, \`UnprocessableEntity\`, \`Locked\`,
\`FailedDependency\`, \`UpgradeRequired\`, \`PreconditionRequired\`, \`TooManyRequests\`,
\`RequestHeaderFieldsTooLarge\`, \`InternalServerError\`, \`NotImplemented\`, \`BadGateway\`,
\`ServiceUnavailable\`, \`GatewayTimeout\`, \`HTTPVersionNotSupported\`, \`VariantAlsoNegotiates\`,
\`InsufficientStorage\`, \`LoopDetected\`, \`NotExtended\`, \`NetworkAuthenticationRequired\`.
Defaults to \`EmptyStatusCode\`.

Upstream description:

HTTP response status codes

EmptyStatusCode response codes means it is not specified Continue status code OK status code Created
status code Accepted status code Non Authoritative Information status code No Content status code
Reset Content status code Partial Content status code Multi Status status code Already Reported
status code Im Used status code Multiple Choices status code Moved Permanently status code Found
status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect
status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment
Required status code Forbidden status code Not Found status code Method Not Allowed status code Not
Acceptable status code Proxy Authentication Required status code Request Timeout status code
Conflict status code Gone status code Length Required status code Precondition Failed status code
Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not
Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable
Entity status code Locked status code Failed Dependency status code Upgrade Required status code
Precondition Required status code Too Many Requests status code Request Header Fields Too Large
status code Internal Server Error status code Not Implemented status code Bad Gateway status code
Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code
Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not
Extended status code Network Authentication Required status code.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "EmptyStatusCode",
  "enum": [
    "EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2320231133121311-2210230101222311-3001311231113133-0332300213203001-3230313001211113-1332301223102012-3113132011123112-3201312102002113"></a>

## Next pages — block_web / 332002032020 / 7

- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-1113311232320333-0300120101320203-2110223102112213-2013121133211100-3321120132301211-2020003231031133-1323311232301303-1022020213311220)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-1221231003113011-0022303133210330-2332321100311113-1033320032232322-3222122030013201-1202220202230113-0332003233200111-2213102220011131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333133222320233-1102322133322200-0021310303331001-3123202232201230-3020101333030221-0023131031132003-0211100013323123-3311111102212233"></a>

## Cloudflare.protected_endpoints.web_mobile_client.continue_mobile — continue_mobile / 330230201230 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-1113311232320333-0300120101320203-2110223102112213-2013121133211100-3321120132301211-2020003231031133-1323311232301303-1022020213311220)
- Cloudflare.protected_endpoints.web_mobile_client.continue_mobile

<a id="canonical-0020312020023110-0332102220321200-3230300020103030-1132122301212201-0233213121010131-3203022302000111-0123111032103312-2102120101012013"></a>

Type: `"object"`. single nested block, Optional.

Select Continue Bot Mitigation Action. Continue mitigation action.

Upstream description:

Continue mitigation action.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("add_header",
    "no_header")}
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
  "x-ves-oneof-field-add_header_choice": "[\"add_header\",\"no_header\"]"
}
```

Terraform syntax:

```terraform
continue_mobile {
  # Configure direct properties listed below.
}
```

<a id="canonical-1320213123302323-1013212203320310-3111222010110010-1001020032211201-1022230013203332-2101010132130002-3002102122011222-2030120013301203"></a>

## Direct properties — continue_mobile / 330230201230 / 3

- [add_header](resources--protected_application--reference--group-002.md#canonical-2102330112200223-3232310120303122-0221020231032121-0203100321232220-3201010221220023-2231023201203310-2220102323301002-1302333012122101): complete subsection reference.

- [no_header](resources--protected_application--reference--group-002.md#canonical-3211131321032221-0023223022233123-1200201002323332-0330121120210022-1313323322220013-0303302010313103-0123032330210210-2011221122101301): complete subsection reference.

<a id="canonical-1332203310233111-2102331231331320-0200302202212222-0010223022323031-2131313123332103-0233103023102113-1111230230301111-3020102223132302"></a>

## Next pages — continue_mobile / 330230201230 / 4

- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile.add_header](resources--protected_application--reference--group-002.md#canonical-2102330112200223-3232310120303122-0221020231032121-0203100321232220-3201010221220023-2231023201203310-2220102323301002-1302333012122101)
- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile.no_header](resources--protected_application--reference--group-002.md#canonical-3211131321032221-0023223022233123-1200201002323332-0330121120210022-1313323322220013-0303302010313103-0123032330210210-2011221122101301)
- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-1113311232320333-0300120101320203-2110223102112213-2013121133211100-3321120132301211-2020003231031133-1323311232301303-1022020213311220)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-2102330112200223-3232310120303122-0221020231032121-0203100321232220-3201010221220023-2231023201203310-2220102323301002-1302333012122101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0312320101320202-0220321102121121-0333021232313012-0310221203332221-1212001002333310-1130121220123231-2030100301330211-0323103001012021"></a>

## Cloudflare.protected_endpoints.web_mobile_client.continue_mobile.add_header — add_header / 333133322031 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-1113311232320333-0300120101320203-2110223102112213-2013121133211100-3321120132301211-2020003231031133-1323311232301303-1022020213311220)
- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile](resources--protected_application--reference--group-002.md#canonical-1221231003113011-0022303133210330-2332321100311113-1033320032232322-3222122030013201-1202220202230113-0332003233200111-2213102220011131)
- Cloudflare.protected_endpoints.web_mobile_client.continue_mobile.add_header

<a id="canonical-1202022210201011-0123213303012133-0021013221331133-1130233300313303-3012203120122012-1022313201231031-3133230211232202-2211212333113303"></a>

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
add_header = {}
```

<a id="canonical-0022201031202233-1322031322333233-3313030022313301-3213133203201300-0132231101132322-0202320202111212-1300312310033202-0333212311003222"></a>

## Direct properties — add_header / 333133322031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1111012102222111-0020230030130310-1331302332302310-1122101032133213-2003212103120333-0112333021302230-1101203213131020-0320100112231310"></a>

## Next pages — add_header / 333133322031 / 4

- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile](resources--protected_application--reference--group-002.md#canonical-1221231003113011-0022303133210330-2332321100311113-1033320032232322-3222122030013201-1202220202230113-0332003233200111-2213102220011131)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-3211131321032221-0023223022233123-1200201002323332-0330121120210022-1313323322220013-0303302010313103-0123032330210210-2011221122101301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0113120200132202-3023201200321130-2110213113001310-3321223110103202-0223220133220233-1301111212212131-1122223102032310-0123323213031101"></a>

## Cloudflare.protected_endpoints.web_mobile_client.continue_mobile.no_header — no_header / 320213113201 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-1113311232320333-0300120101320203-2110223102112213-2013121133211100-3321120132301211-2020003231031133-1323311232301303-1022020213311220)
- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile](resources--protected_application--reference--group-002.md#canonical-1221231003113011-0022303133210330-2332321100311113-1033320032232322-3222122030013201-1202220202230113-0332003233200111-2213102220011131)
- Cloudflare.protected_endpoints.web_mobile_client.continue_mobile.no_header

<a id="canonical-3303302132232022-1331033131311302-2121311121201222-2030211302222123-3302123132132333-0131003203330201-3020330222201202-0302132133330303"></a>

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
no_header = {}
```

<a id="canonical-3232320231110330-0131201020123013-0321121012020223-1122131101013012-0223212201302012-1011222203121123-2301333223123332-0123211232112122"></a>

## Direct properties — no_header / 320213113201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2022222023103023-1033033223021333-0222021321330213-0103301001312222-2301120300013110-0032331232333333-2130002201221222-2102330003330221"></a>

## Next pages — no_header / 320213113201 / 4

- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile](resources--protected_application--reference--group-002.md#canonical-1221231003113011-0022303133210330-2332321100311113-1033320032232322-3222122030013201-1202220202230113-0332003233200111-2213102220011131)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-0102201103120331-1320010313110222-0323100032123012-2333110230202132-0110220130102010-0200322230322022-0213211311303332-1132200123313312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011033011101122-1233113031310022-2013233110010310-3120002200011331-2321233303311030-2103311300122113-1100033010120012-2031122322231130"></a>

## Cloudflare.protected_endpoints.web_mobile_client.continue_web — continue_web / 121333220021 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-1113311232320333-0300120101320203-2110223102112213-2013121133211100-3321120132301211-2020003231031133-1323311232301303-1022020213311220)
- Cloudflare.protected_endpoints.web_mobile_client.continue_web

<a id="canonical-1231113030022102-2000200322232122-2321312123020001-0113033202330212-0330111012330220-0303213222231103-1010020132330221-2102320221012112"></a>

Type: `"object"`. single nested block, Optional.

Select Continue Bot Mitigation Action. Continue mitigation action.

Upstream description:

Continue mitigation action.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("add_header",
    "no_header")}
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
  "x-ves-oneof-field-add_header_choice": "[\"add_header\",\"no_header\"]"
}
```

Terraform syntax:

```terraform
continue_web {
  # Configure direct properties listed below.
}
```

<a id="canonical-0332310012331000-3302223031333130-1230021022212012-3013333323233300-0302202033110312-2113212302101131-2311121333020312-1000100232302012"></a>

## Direct properties — continue_web / 121333220021 / 3

- [add_header](resources--protected_application--reference--group-002.md#canonical-3232112211331123-0232322102013112-2011011120101230-1332323131200220-3300113010212303-3022010100230100-0212222101122000-3002112312002200): complete subsection reference.

- [no_header](resources--protected_application--reference--group-002.md#canonical-2123113300302133-0133103220000233-1323120313122230-2303121030233303-1111203210320201-0133130331321320-3331221103003212-3323132312222133): complete subsection reference.

<a id="canonical-2301311330202211-0331112332333101-2231320203132222-3123222221331132-2003202302232002-2210332232121010-0203133210313223-3013233011233300"></a>

## Next pages — continue_web / 121333220021 / 4

- [cloudflare.protected_endpoints.web_mobile_client.continue_web.add_header](resources--protected_application--reference--group-002.md#canonical-3232112211331123-0232322102013112-2011011120101230-1332323131200220-3300113010212303-3022010100230100-0212222101122000-3002112312002200)
- [cloudflare.protected_endpoints.web_mobile_client.continue_web.no_header](resources--protected_application--reference--group-002.md#canonical-2123113300302133-0133103220000233-1323120313122230-2303121030233303-1111203210320201-0133130331321320-3331221103003212-3323132312222133)
- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-1113311232320333-0300120101320203-2110223102112213-2013121133211100-3321120132301211-2020003231031133-1323311232301303-1022020213311220)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-3232112211331123-0232322102013112-2011011120101230-1332323131200220-3300113010212303-3022010100230100-0212222101122000-3002112312002200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211030120022222-1220220231321230-3323301330100113-2010113123303030-3022033010201223-3221132031120030-0003321133220100-3321202200113312"></a>

## Cloudflare.protected_endpoints.web_mobile_client.continue_web.add_header — add_header / 133300321232 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-1113311232320333-0300120101320203-2110223102112213-2013121133211100-3321120132301211-2020003231031133-1323311232301303-1022020213311220)
- [cloudflare.protected_endpoints.web_mobile_client.continue_web](resources--protected_application--reference--group-002.md#canonical-0102201103120331-1320010313110222-0323100032123012-2333110230202132-0110220130102010-0200322230322022-0213211311303332-1132200123313312)
- Cloudflare.protected_endpoints.web_mobile_client.continue_web.add_header

<a id="canonical-3112123101221322-0223303102103030-3230202330103013-1233213330203102-1010233232312332-0001122023102003-3021210221121023-3001311310310102"></a>

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
add_header = {}
```

<a id="canonical-1313011111222123-0120212131133231-1132231233010013-1030001010200221-3021203231120000-1123212000320230-0331313322112322-1031011223010022"></a>

## Direct properties — add_header / 133300321232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0102223131323200-0123020032203131-3333133022321001-3133310003323201-2113130033112121-1201200310210330-0222133212201012-1232030112123112"></a>

## Next pages — add_header / 133300321232 / 4

- [cloudflare.protected_endpoints.web_mobile_client.continue_web](resources--protected_application--reference--group-002.md#canonical-0102201103120331-1320010313110222-0323100032123012-2333110230202132-0110220130102010-0200322230322022-0213211311303332-1132200123313312)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-2123113300302133-0133103220000233-1323120313122230-2303121030233303-1111203210320201-0133130331321320-3331221103003212-3323132312222133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112232030331220-0333112311223131-1331020212320220-2200121331123010-1301310232103012-3010323323313310-3210030023233203-0330223020211031"></a>

## Cloudflare.protected_endpoints.web_mobile_client.continue_web.no_header — no_header / 230001103330 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-1113311232320333-0300120101320203-2110223102112213-2013121133211100-3321120132301211-2020003231031133-1323311232301303-1022020213311220)
- [cloudflare.protected_endpoints.web_mobile_client.continue_web](resources--protected_application--reference--group-002.md#canonical-0102201103120331-1320010313110222-0323100032123012-2333110230202132-0110220130102010-0200322230322022-0213211311303332-1132200123313312)
- Cloudflare.protected_endpoints.web_mobile_client.continue_web.no_header

<a id="canonical-0211301332003020-1133323021030330-0010121000200202-2310012110103130-0000210211103002-0331032303232133-3220200013231022-0002333302010130"></a>

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
no_header = {}
```

<a id="canonical-1110022032323220-0320013313100021-1202230333100001-0230300330023331-3201221231122213-0021013303302311-0131312301211321-3001200022121330"></a>

## Direct properties — no_header / 230001103330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3032130132232331-1033333232001123-1233003330330122-0202232003302210-2000010110132122-2220120033112213-0321202233023003-0120103012311021"></a>

## Next pages — no_header / 230001103330 / 4

- [cloudflare.protected_endpoints.web_mobile_client.continue_web](resources--protected_application--reference--group-002.md#canonical-0102201103120331-1320010313110222-0323100032123012-2333110230202132-0110220130102010-0200322230322022-0213211311303332-1132200123313312)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-3223112032330233-3211013131233032-3321031030031233-1211001113012013-3010223031232211-2123312332113311-3100023303213011-0010021220103331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011301202031032-3023020213012030-3303131131123313-1021003323111320-1120301331301100-0130133323003131-2022220031001321-0021010131230012"></a>

## Cloudflare.protected_endpoints.web_mobile_client.redirect_web — redirect_web / 321323100301 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.protected_endpoints](resources--protected_application--reference--group-001.md#canonical-3122300021323100-2230320020332020-3300130320322201-0303020102021123-1120021320220130-3133000132211113-0133313003222210-2112203222032200)
- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-1113311232320333-0300120101320203-2110223102112213-2013121133211100-3321120132301211-2020003231031133-1323311232301303-1022020213311220)
- Cloudflare.protected_endpoints.web_mobile_client.redirect_web

<a id="canonical-3321121013300203-2221300202100130-2001121130003103-2223110132233113-0030113330210212-3033221230011011-3021233113031113-3001101201201123"></a>

Type: `"object"`. single nested block, Optional.

Redirect. Redirect.

Upstream description:

Redirect.

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
redirect_web {
  # Configure direct properties listed below.
}
```

<a id="canonical-3333223202020201-3130101233212033-3332103023130303-1103301302022122-0010032330330023-1330113310110322-3133032300110211-0133323303330222"></a>

## Direct properties — redirect_web / 321323100301 / 3

<a id="canonical-1033102313130121-1212111200320312-2103010022303313-0320211001010201-2231122222210312-1130002201302220-1022010303201320-1022331001203312"></a>

<a id="canonical-1301233231010010-0230333013030132-3310033133213030-2112031031102132-3020023220222321-1021121112310331-2231022103003012-1213211113312111"></a>

## location property — redirect_web / 321323100301 / 4

Type: `"string"`. Optional.

Location. URI location for redirect response.

Upstream description:

URI location for redirect response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 512
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_bytes": "512",
    "ves.io.schema.rules.string.url_or_uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "512",
    "ves.io.schema.rules.string.url_or_uri_ref": "true"
  }
}
```

<a id="canonical-1231212322331113-2002101310332100-3002103011131103-0233313001232011-0113123001313021-2301311123033110-3200212033223130-2311023300212001"></a>

<a id="canonical-2122323330213232-2201133130021020-2133023010310013-3232302303012303-0132100000300002-1003300300220000-0031231102112200-3030221332131221"></a>

## status property — redirect_web / 321323100301 / 5

Type: `"string"`. Optional.

\[Enum:
EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired\]
HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status
code OK status code Created status code Accepted status code Non Authoritative Information status
code No Content status code Reset Content status code Partial Content status code Multi Status..
Possible values are \`EmptyStatusCode\`, \`Continue\`, \`OK\`, \`Created\`, \`Accepted\`,
\`NonAuthoritativeInformation\`, \`NoContent\`, \`ResetContent\`, \`PartialContent\`,
\`MultiStatus\`, \`AlreadyReported\`, \`IMUsed\`, \`MultipleChoices\`, \`MovedPermanently\`,
\`Found\`, \`SeeOther\`, \`NotModified\`, \`UseProxy\`, \`TemporaryRedirect\`,
\`PermanentRedirect\`, \`BadRequest\`, \`Unauthorized\`, \`PaymentRequired\`, \`Forbidden\`,
\`NotFound\`, \`MethodNotAllowed\`, \`NotAcceptable\`, \`ProxyAuthenticationRequired\`,
\`RequestTimeout\`, \`Conflict\`, \`Gone\`, \`LengthRequired\`, \`PreconditionFailed\`,
\`PayloadTooLarge\`, \`URITooLong\`, \`UnsupportedMediaType\`, \`RangeNotSatisfiable\`,
\`ExpectationFailed\`, \`MisdirectedRequest\`, \`UnprocessableEntity\`, \`Locked\`,
\`FailedDependency\`, \`UpgradeRequired\`, \`PreconditionRequired\`, \`TooManyRequests\`,
\`RequestHeaderFieldsTooLarge\`, \`InternalServerError\`, \`NotImplemented\`, \`BadGateway\`,
\`ServiceUnavailable\`, \`GatewayTimeout\`, \`HTTPVersionNotSupported\`, \`VariantAlsoNegotiates\`,
\`InsufficientStorage\`, \`LoopDetected\`, \`NotExtended\`, \`NetworkAuthenticationRequired\`.
Defaults to \`EmptyStatusCode\`.

Upstream description:

HTTP response status codes

EmptyStatusCode response codes means it is not specified Continue status code OK status code Created
status code Accepted status code Non Authoritative Information status code No Content status code
Reset Content status code Partial Content status code Multi Status status code Already Reported
status code Im Used status code Multiple Choices status code Moved Permanently status code Found
status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect
status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment
Required status code Forbidden status code Not Found status code Method Not Allowed status code Not
Acceptable status code Proxy Authentication Required status code Request Timeout status code
Conflict status code Gone status code Length Required status code Precondition Failed status code
Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not
Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable
Entity status code Locked status code Failed Dependency status code Upgrade Required status code
Precondition Required status code Too Many Requests status code Request Header Fields Too Large
status code Internal Server Error status code Not Implemented status code Bad Gateway status code
Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code
Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not
Extended status code Network Authentication Required status code.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "EmptyStatusCode",
  "enum": [
    "EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2221330130001201-0302110033100111-3211130233131231-0220301030220301-0000203130220332-0313200023231133-3321022101211230-0030300001212031"></a>

## Next pages — redirect_web / 321323100301 / 6

- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--reference--group-002.md#canonical-1113311232320333-0300120101320203-2110223102112213-2013121133211100-3321120132301211-2020003231031133-1323311232301303-1022020213311220)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-2112211221002230-0132313302003203-3023102123212203-2101231013213232-1320231230122010-0132301122202122-2323222023222001-1212210200131203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323232003320130-2203231113131100-2312232131321130-0203303302130101-3321001000023201-2322023220323103-2213213032030103-3320331201301013"></a>

## Cloudflare.trusted_clients — trusted_clients / 202103331203 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- Cloudflare.trusted_clients

<a id="canonical-1131333132113202-2301303320003333-2310101333111312-1111122101112132-0233321030333333-2010121213002123-3233022020022211-3220210000021010"></a>

Type: `"object"`. list nested block, Optional.

Define your allowlists to skip Bot Defense inference processing.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("http_header",
    "ip_prefix")}
```

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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
trusted_clients {
  # Configure direct properties listed below.
}
```

<a id="canonical-2111201130311002-2311112230110123-0221230232013111-0111000223231301-2111223103301122-0321220110102310-0311113210302220-2132303311230102"></a>

## Direct properties — trusted_clients / 202103331203 / 3

- [http_header](resources--protected_application--reference--group-002.md#canonical-0332132102022001-1122202233120232-0022011102002233-2201111233120232-0222033002222233-2101103130212233-3310121222321213-2030023011111303): complete subsection reference.

<a id="canonical-2201222032231023-1100111113322000-0310001022010001-0102110101131113-1102200003213313-0102301312330023-1222313303121203-2221110231322320"></a>

<a id="canonical-1003310332300123-2230330131113210-1023223111031111-2002222231000333-0323132121331303-1321322233032302-3303330130200223-3222002123101202"></a>

## ip_prefix property — trusted_clients / 202103331203 / 4

Type: `"string"`. Optional.

Exclusive with \[http\_header\] IP prefix string.

Upstream description:

Exclusive with \[http\_header\] IP prefix string.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.CIDRValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
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
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

- [metadata](resources--protected_application--reference--group-002.md#canonical-3110220033333323-0312031131030203-3323100222312233-0022112032113311-1113103231330102-2133330301210221-0020222332120210-2130011333123321): complete subsection reference.

<a id="canonical-0123013000203303-3120312323223003-2001003112022212-2230222121130001-1330320102010110-3103303031321323-2133122211022323-2301302132110311"></a>

## Next pages — trusted_clients / 202103331203 / 5

- [cloudflare.trusted_clients.http_header](resources--protected_application--reference--group-002.md#canonical-0332132102022001-1122202233120232-0022011102002233-2201111233120232-0222033002222233-2101103130212233-3310121222321213-2030023011111303)
- [cloudflare.trusted_clients.metadata](resources--protected_application--reference--group-002.md#canonical-3110220033333323-0312031131030203-3323100222312233-0022112032113311-1113103231330102-2133330301210221-0020222332120210-2130011333123321)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-0332132102022001-1122202233120232-0022011102002233-2201111233120232-0222033002222233-2101103130212233-3310121222321213-2030023011111303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013002211120101-3333302013331221-2001100312120012-1120011302020022-3032322331021031-3000011130112202-2112033301123001-2020032320310220"></a>

## Cloudflare.trusted_clients.http_header — http_header / 213032223020 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.trusted_clients](resources--protected_application--reference--group-002.md#canonical-2112211221002230-0132313302003203-3023102123212203-2101231013213232-1320231230122010-0132301122202122-2323222023222001-1212210200131203)
- Cloudflare.trusted_clients.http_header

<a id="canonical-2213020301213230-2223211032033033-0202320213002123-3210032202222300-3202103221203122-3213131331002000-1012201022100103-3200023101103113"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http header.

Upstream description:

Request header name and value pairs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("headers")}
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
http_header {
  # Configure direct properties listed below.
}
```

<a id="canonical-0200300221220022-0323132111002201-0213100302021303-2121113030031320-3132121302221113-0302023132213202-0301300223002310-2022012320111211"></a>

## Direct properties — http_header / 213032223020 / 3

- [headers](resources--protected_application--reference--group-002.md#canonical-2030020201011302-3313320113133233-0030311112300302-2300121302030020-3030211001313012-3023223310212313-3022233213231302-3120013020211031): complete subsection reference.

<a id="canonical-1223133203312310-2213332323012221-2333033333332300-2323320202200120-3231021001301021-2332211231321220-1111233333321231-1222310210002010"></a>

## Next pages — http_header / 213032223020 / 4

- [cloudflare.trusted_clients.http_header.headers](resources--protected_application--reference--group-002.md#canonical-2030020201011302-3313320113133233-0030311112300302-2300121302030020-3030211001313012-3023223310212313-3022233213231302-3120013020211031)
- [cloudflare.trusted_clients](resources--protected_application--reference--group-002.md#canonical-2112211221002230-0132313302003203-3023102123212203-2101231013213232-1320231230122010-0132301122202122-2323222023222001-1212210200131203)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-2030020201011302-3313320113133233-0030311112300302-2300121302030020-3030211001313012-3023223310212313-3022233213231302-3120013020211031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111330011203032-2011102201030001-3222233011311202-1002202203102121-2321230020132111-3022101123002030-2300333220223332-0332210300213321"></a>

## Cloudflare.trusted_clients.http_header.headers — headers / 320011233123 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.trusted_clients](resources--protected_application--reference--group-002.md#canonical-2112211221002230-0132313302003203-3023102123212203-2101231013213232-1320231230122010-0132301122202122-2323222023222001-1212210200131203)
- [cloudflare.trusted_clients.http_header](resources--protected_application--reference--group-002.md#canonical-0332132102022001-1122202233120232-0022011102002233-2201111233120232-0222033002222233-2101103130212233-3310121222321213-2030023011111303)
- Cloudflare.trusted_clients.http_header.headers

<a id="canonical-3330231022313213-3223033010203031-0313200320102002-0022023110303003-0302212332312310-2113001220001210-3233023111233123-0002131211310310"></a>

Type: `"object"`. list nested block, Optional.

List of HTTP header name and value pairs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("exact",
    "regex")}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-3032232121310111-0030112031300320-0003230012032032-0110202311302103-2021031032202132-1233121030103013-3111030211313030-3130312221213313"></a>

## Direct properties — headers / 320011233123 / 3

<a id="canonical-1333300212333232-3301012321211102-1103332103020112-3213011313013220-1223202100120100-1331223101203002-1302302210112010-0022002331123111"></a>

<a id="canonical-1212312313302301-2101132330120301-2302010033120323-3020220110233013-3202022020120210-2131020020031032-3000123003302000-0023301202020121"></a>

## exact property — headers / 320011233123 / 4

Type: `"string"`. Optional.

Exclusive with \[regex\] Header value to match exactly.

Upstream description:

Exclusive with \[regex\] Header value to match exactly.

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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-2321222303030310-2210030221113123-2032101101310213-1003332013310120-0020312120133031-2303222032133101-3331133000332130-0223201100032231"></a>

<a id="canonical-2100013000210002-0002010302301301-2113321012230303-2000122110133011-1230012222231312-1330003323110300-3303012102033233-3310021121231022"></a>

## name property — headers / 320011233123 / 5

Type: `"string"`. Optional.

Name. Name of the header.

Upstream description:

Name of the header.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-2230121101211233-0032331121321033-0300113110313301-0121212211202000-1322132012300033-3320013310323303-0022113313233133-3111203311012013"></a>

<a id="canonical-0110320211333232-0010321332010101-1301033002320313-1120201023020312-3020220301001103-0101112112022021-1023210000031211-0030310101233223"></a>

## regex property — headers / 320011233123 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\] Regex match of the header value in re2 format.

Upstream description:

Exclusive with \[exact\] Regex match of the header value in re2 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.not_empty": "true",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3130330302312032-0131213110233000-2032023211012211-0300013222222030-2011020121110020-0300000132023313-1221330333311003-1133230310010332"></a>

## Next pages — headers / 320011233123 / 7

- [cloudflare.trusted_clients.http_header](resources--protected_application--reference--group-002.md#canonical-0332132102022001-1122202233120232-0022011102002233-2201111233120232-0222033002222233-2101103130212233-3310121222321213-2030023011111303)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-3110220033333323-0312031131030203-3323100222312233-0022112032113311-1113103231330102-2133330301210221-0020222332120210-2130011333123321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3111203322210231-1121231212331210-3132333131231130-1113030312231011-1013001312321311-1033012323113313-3020312221333221-0313332301111021"></a>

## Cloudflare.trusted_clients.metadata — metadata / 002220311112 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [Cloudflare](resources--protected_application--reference--group-001.md#canonical-1300213123131122-1030311223302010-1123221002212300-1123122213010200-1110331233223310-3211033220031200-3333302300120030-0020033300031111)
- [cloudflare.trusted_clients](resources--protected_application--reference--group-002.md#canonical-2112211221002230-0132313302003203-3023102123212203-2101231013213232-1320231230122010-0132301122202122-2323222023222001-1212210200131203)
- Cloudflare.trusted_clients.metadata

<a id="canonical-2201322213223112-0221310303320031-0310103213131102-3310301311330211-3030033213020102-3313111313313131-1033300012313331-2023323332223203"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-0103332033220320-0333210220223300-2213110032100332-1121233100103330-1101112321113102-2032300010211030-3303030310313232-1003120231110033"></a>

## Direct properties — metadata / 002220311112 / 3

<a id="canonical-0111211300022002-3211121013120131-2301311331223031-2012223102013002-1301321232223013-0030102321103232-3100220330330301-2123320202310321"></a>

<a id="canonical-2001200212211332-0033333033223310-0001213121132211-1123113221322132-2002330332023301-0202001330210002-3122121110113311-1112302313011110"></a>

## description_spec property — metadata / 002220311112 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-1323133113332330-1010301320001032-3201230023321303-3132322322330310-3332130233322011-1203321023322012-2203311333222201-0010112022030203"></a>

<a id="canonical-0031012031332233-2023210322013112-1310231031111222-3213133203023132-0021212321132122-2322232211301122-3132110332100322-3110022113110030"></a>

## name property — metadata / 002220311112 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-3023120301233111-1122013110031100-1012030303310310-3123202010031102-2233220002233322-1202210130033120-3220320200031203-1103232002121331"></a>

## Next pages — metadata / 002220311112 / 6

- [cloudflare.trusted_clients](resources--protected_application--reference--group-002.md#canonical-2112211221002230-0132313302003203-3023102123212203-2101231013213232-1320231230122010-0132301122202122-2323222023222001-1212210200131203)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301323032122310-1023012122002120-0312303102202223-3322211012022132-0321110213011231-1210320302313031-0221221103013013-0111311121310030"></a>

## cloudfront — cloudfront / 002210331013 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- cloudfront

<a id="canonical-0323011203311321-0002123031021111-2112223320220213-1231321321221222-3112000200110112-2331021213200101-2130200001232020-2310023233330330"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense policy configuration for AWS Cloudfront.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("protected_endpoints"),
  validators.ConflictingObjectAttributes("aws_configuration_id_selector",
    "aws_configuration_tag_selector"),
  validators.ConflictingObjectAttributes("aws_configuration_id_selector",
    "disable_aws_configuration"),
  validators.ConflictingObjectAttributes("aws_configuration_tag_selector",
    "disable_aws_configuration"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "js_insertion_rules"),
  validators.ConflictingObjectAttributes("disable_js_insert",
    "manual_js_insert"),
  validators.ConflictingObjectAttributes("disable_mobile_sdk",
    "mobile_sdk_config"),
  validators.ConflictingObjectAttributes("js_insertion_rules",
    "manual_js_insert")}
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
  "x-ves-oneof-field-aws_configuration_type_choice": "[\"aws_configuration_id_selector\",\"aws_configuration_tag_selector\",\"disable_aws_configuration\"]",
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insertion_rules\",\"manual_js_insert\"]",
  "x-ves-oneof-field-mobile_sdk_choice": "[\"disable_mobile_sdk\",\"mobile_sdk_config\"]"
}
```

Terraform syntax:

```terraform
cloudfront {
  # Configure direct properties listed below.
}
```

<a id="canonical-0012302202210201-1320222031110213-3131020310123133-2100222221300311-1131110321303311-0101103312102232-3232210201303230-3312210012000312"></a>

## Direct properties — cloudfront / 002210331013 / 3

- [aws_configuration_id_selector](resources--protected_application--reference--group-002.md#canonical-2001311023332132-2323212030212103-0022332300333333-3003021123320310-0021023320310331-3000001132213121-1331211303103221-2312121021022113): complete subsection reference.

- [aws_configuration_tag_selector](resources--protected_application--reference--group-002.md#canonical-2301003033010132-2011002322022103-1032023300113333-3203030110102133-1233003321210103-1010302203112203-1332321303311222-0210232022202322): complete subsection reference.

<a id="canonical-3023102100102300-2020110030232123-2220003032101000-0131310021331123-3232113012330302-1231031333101223-0311223313303212-2202020002221333"></a>

<a id="canonical-3303232113322331-3330210000232103-2110333030032212-3021030103220131-3030101221223223-2313233022222220-0110212030322231-3003130311111030"></a>

## continue_mitigation_action_hdr property — cloudfront / 002210331013 / 4

Type: `"string"`. Optional.

Case-insensitive HTTP header name for Continue Mitigation Action when add header selected.

Upstream description:

A case-insensitive HTTP header name for Continue Mitigation Action when add header selected.

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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-2122303003123230-3100332312320120-0100202001332333-3031321110321312-3123012101232213-0112120220231323-2003032103203231-0300310223013010"></a>

<a id="canonical-0322203331031123-1330123133021210-2313232020233000-3211002231200112-3030021101102131-1011120210003102-3233212312100230-1012120132000230"></a>

## data_sample property — cloudfront / 002210331013 / 5

Type: `"number"`. Optional.

Limit on amount of request-body data (other than F5 telemetry) to send for analysis (limit 1,048,576
== 1 MiByte).

Upstream description:

Limit on amount of request-body data (other than F5 telemetry) to send for analysis (limit 1,048,576
== 1 MiByte)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 1048576),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1048576,
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
    "ves.io.schema.rules.uint32.lte": "1048576"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "1048576"
  }
}
```

- [disable_aws_configuration](resources--protected_application--reference--group-002.md#canonical-2010231000113321-3132021233310322-1232323033122132-2011123130301301-1003110013021021-2110131123331323-0202112031020112-0133123112200100): complete subsection reference.

- [disable_js_insert](resources--protected_application--reference--group-002.md#canonical-2021213232222233-1330301300310122-1020102001112303-2300100211113202-0323221313013012-1033210213010223-1122330233301300-1222030203220202): complete subsection reference.

- [disable_mobile_sdk](resources--protected_application--reference--group-002.md#canonical-0021320220121330-1123001003212312-0103101010131121-2011033121210101-0313032101220000-3302332203012213-3111202202113020-2302220232120222): complete subsection reference.

- [js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-0231330302202122-3333330110232010-0233303200130220-0301122123331132-2200201002333200-2020222112331210-1313021231112202-3002213230012210): complete subsection reference.

<a id="canonical-0200213222021013-2012212303233100-0013222033230112-2123103232303031-3332032132212033-2213310230003002-0123211310101320-2332102230131121"></a>

<a id="canonical-1313133102222013-3003010202031131-1320302310002323-0303031300123000-2232323130031033-0323031002330223-0030132333311212-0121000003000030"></a>

## loglevel property — cloudfront / 002210331013 / 6

Type: `"string"`. Optional.

\[Enum: LOG\_UNDEFINED|LOG\_ERROR|LOG\_WARNING|LOG\_INFO|LOG\_DEBUG\] Select the level of logging
desired. Levels are cumulative (e.g. Debug includes Error, Warning, and Informational) -
LOG\_UNDEFINED: Undefined - LOG\_ERROR: Error Log only errors - LOG\_WARNING: Warning Log malicious
requests - LOG\_INFO: Info Log all requests - LOG\_DEBUG: Debug Log debugging data. Possible values
are \`LOG\_UNDEFINED\`, \`LOG\_ERROR\`, \`LOG\_WARNING\`, \`LOG\_INFO\`, \`LOG\_DEBUG\`. Defaults to
\`LOG\_UNDEFINED\`.

Upstream description:

Select the level of logging desired. Levels are cumulative (e.g. Debug includes Error, Warning, and
Informational)

&#8203;- LOG\_UNDEFINED: Undefined

&#8203;- LOG\_ERROR: Error

Log only errors &#8203;- LOG\_WARNING: Warning

Log malicious requests &#8203;- LOG\_INFO: Info

Log all requests &#8203;- LOG\_DEBUG: Debug

Log debugging data.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("LOG_UNDEFINED",
    "LOG_ERROR",
    "LOG_WARNING",
    "LOG_INFO",
    "LOG_DEBUG"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "LOG_UNDEFINED",
  "enum": [
    "LOG_UNDEFINED",
    "LOG_ERROR",
    "LOG_WARNING",
    "LOG_INFO",
    "LOG_DEBUG"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [manual_js_insert](resources--protected_application--reference--group-003.md#canonical-3313102332000123-2222313311313211-0013120212332103-3201031211032333-2332132333320123-1122123222002123-2323011023213133-1102033113123031): complete subsection reference.

- [mobile_sdk_config](resources--protected_application--reference--group-003.md#canonical-3033020132301333-0003231320203312-0230003322013233-2133203312003202-3300000233233110-3001200323232120-3021132321131110-0111230321020021): complete subsection reference.

- [protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232): complete subsection reference.

<a id="canonical-1021002221210131-3333001132021221-3031121310230132-3310032330123023-2111231223310301-0111013032033001-3111132201000221-2010022331122030"></a>

<a id="canonical-2013123123300300-3200202300321222-2102201230010131-0013322023011213-3203030223312133-1220032021223221-3120202032020021-3213201211020320"></a>

## timeout property — cloudfront / 002210331013 / 7

Type: `"number"`. Optional.

The timeout for the inference check, in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 60000),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [trusted_clients](resources--protected_application--reference--group-004.md#canonical-1331313303213213-1122122223112222-3001222033333102-3201121101201113-1303223000202010-0120230110212212-1113101221303311-3220103301020330): complete subsection reference.

<a id="canonical-0302020222331312-0021112323223230-3003233301320201-0232021321311202-0331320330222333-0200033321100121-2210030202201121-0122230232031220"></a>

## Next pages — cloudfront / 002210331013 / 8

- [cloudfront.aws_configuration_id_selector](resources--protected_application--reference--group-002.md#canonical-2001311023332132-2323212030212103-0022332300333333-3003021123320310-0021023320310331-3000001132213121-1331211303103221-2312121021022113)
- [cloudfront.aws_configuration_tag_selector](resources--protected_application--reference--group-002.md#canonical-2301003033010132-2011002322022103-1032023300113333-3203030110102133-1233003321210103-1010302203112203-1332321303311222-0210232022202322)
- [cloudfront.disable_aws_configuration](resources--protected_application--reference--group-002.md#canonical-2010231000113321-3132021233310322-1232323033122132-2011123130301301-1003110013021021-2110131123331323-0202112031020112-0133123112200100)
- [cloudfront.disable_js_insert](resources--protected_application--reference--group-002.md#canonical-2021213232222233-1330301300310122-1020102001112303-2300100211113202-0323221313013012-1033210213010223-1122330233301300-1222030203220202)
- [cloudfront.disable_mobile_sdk](resources--protected_application--reference--group-002.md#canonical-0021320220121330-1123001003212312-0103101010131121-2011033121210101-0313032101220000-3302332203012213-3111202202113020-2302220232120222)
- [cloudfront.js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-0231330302202122-3333330110232010-0233303200130220-0301122123331132-2200201002333200-2020222112331210-1313021231112202-3002213230012210)
- [cloudfront.manual_js_insert](resources--protected_application--reference--group-003.md#canonical-3313102332000123-2222313311313211-0013120212332103-3201031211032333-2332132333320123-1122123222002123-2323011023213133-1102033113123031)
- [cloudfront.mobile_sdk_config](resources--protected_application--reference--group-003.md#canonical-3033020132301333-0003231320203312-0230003322013233-2133203312003202-3300000233233110-3001200323232120-3021132321131110-0111230321020021)
- [cloudfront.protected_endpoints](resources--protected_application--reference--group-003.md#canonical-2331031331211001-2203303321110110-3203202010032323-1311020101113303-1331221311131220-0322003310233220-3231300302101310-1323132122212232)
- [cloudfront.trusted_clients](resources--protected_application--reference--group-004.md#canonical-1331313303213213-1122122223112222-3001222033333102-3201121101201113-1303223000202010-0120230110212212-1113101221303311-3220103301020330)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-2001311023332132-2323212030212103-0022332300333333-3003021123320310-0021023320310331-3000001132213121-1331211303103221-2312121021022113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011302202020303-3231230010302202-3112310123122023-3330100303123232-0031311311132212-1103131321112002-0102010301302002-3102313003030203"></a>

## cloudfront.aws_configuration_id_selector — aws_configuration_id_selector / 231230020203 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- cloudfront.aws_configuration_id_selector

<a id="canonical-3203022033232033-3300020132202320-1102030032203330-2003000231210111-3312100132300310-3123020102121330-3011301311010131-0202031002313021"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for aws configuration ID selector.

Upstream description:

List of CloudFront distributions.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ids")}
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
aws_configuration_id_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-3012222111113001-1002230030001111-2102212231011021-0131222310213202-3321233202021211-1113331030203231-0101113230331331-0233012032132233"></a>

## Direct properties — aws_configuration_id_selector / 231230020203 / 3

<a id="canonical-3303221233233132-0320333202103230-3223102230210232-1233010332032202-2100132030001101-1333323132221330-2013021033103023-1332331330212200"></a>

<a id="canonical-0213233203310322-1113032310333220-2223031013221020-2212313321103202-1132021213312031-0023001213023303-0032231313311110-0212213113002011"></a>

## ids property — aws_configuration_id_selector / 231230020203 / 4

Type: `["list", "string"]`. Optional.

Add AWS CloudFront distribution ID, e.g. ABCDEFGHI0JKLM.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "32",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.pattern": "^[A-Z0-9]+$",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "32",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.pattern": "^[A-Z0-9]+$",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2012223123200112-1332200002210011-2023310133032330-3131110313131003-0230022123101020-2112030301323321-3331311221030003-1302221221200001"></a>

## Next pages — aws_configuration_id_selector / 231230020203 / 5

- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-2301003033010132-2011002322022103-1032023300113333-3203030110102133-1233003321210103-1010302203112203-1332321303311222-0210232022202322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2013130111331212-3311012232100021-1021301022032202-0132133133120130-1113322303132123-1133210112323203-0131223000033313-0222131032113201"></a>

## cloudfront.aws_configuration_tag_selector — aws_configuration_tag_selector / 100300102032 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- cloudfront.aws_configuration_tag_selector

<a id="canonical-3111200223233223-1200320210211330-0230003313032323-1021100122121321-3103023312000001-3333002130312022-1220101211330211-3303323330333200"></a>

Type: `"object"`. single nested block, Optional.

Distribution Tag List. CloudFront distribution tag list.

Upstream description:

CloudFront distribution tag list.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("tags")}
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
aws_configuration_tag_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-1200110110312201-2213113303201320-1213001203203332-2330300201230110-3322122022330200-0113102002022031-1001020323200323-0230123221003203"></a>

## Direct properties — aws_configuration_tag_selector / 100300102032 / 3

<a id="canonical-3020023302330110-2111212012031200-1033132330322123-0311203023233002-0302330302031012-1002012103021021-2220302301233330-2330202313221023"></a>

<a id="canonical-0031112333131032-0311203322112231-0221012233020000-1222322301020233-3230220321132010-3120003222221102-3311103202133121-3021013300313322"></a>

## tags property — aws_configuration_tag_selector / 100300102032 / 4

Type: `["map", "string"]`. Optional.

List contains the Cloudfront distribution selection by tags key is a AWS tag name, and the value is
regular expression to match.

Upstream description:

List contains the Cloudfront distribution selection by tags key is a AWS tag name, and the value is
regular expression to match.

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
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.min_pairs": "1",
    "ves.io.schema.rules.map.values.string.max_len": "256",
    "ves.io.schema.rules.map.values.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.regex": "true",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.min_pairs": "1",
    "ves.io.schema.rules.map.values.string.max_len": "256",
    "ves.io.schema.rules.map.values.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.regex": "true",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-2211233002220123-1000232002121313-1000012121233002-0310103331121332-0120033332032001-0110212010300111-1133132313302212-3111301002203220"></a>

## Next pages — aws_configuration_tag_selector / 100300102032 / 5

- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-2010231000113321-3132021233310322-1232323033122132-2011123130301301-1003110013021021-2110131123331323-0202112031020112-0133123112200100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1311313310100201-0033321113033231-3020103333113013-3210212011030110-0031220101230101-2223220003203103-1112012311223112-0021120113331022"></a>

## cloudfront.disable_aws_configuration — disable_aws_configuration / 230230133001 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- cloudfront.disable_aws_configuration

<a id="canonical-3010332231112002-0212022212200200-0102313303312133-3321122002011231-0313303023203122-3313003310233133-3302302230013120-0123113001011332"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable aws configuration.

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
disable_aws_configuration = {}
```

<a id="canonical-1320200210200211-2002033112020330-1022123100030302-3103032101233321-3331321233012322-2310201302212213-3011132032330123-1302123102200113"></a>

## Direct properties — disable_aws_configuration / 230230133001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2200312110011003-0230131101002303-1201023303202300-1002211121220123-2300012200123213-1312301022203230-2112101210002302-0101003011322230"></a>

## Next pages — disable_aws_configuration / 230230133001 / 4

- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-2021213232222233-1330301300310122-1020102001112303-2300100211113202-0323221313013012-1033210213010223-1122330233301300-1222030203220202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0313112330220221-2331122102231200-3030120123300000-0331333011321020-3200213223221222-3201120300130202-2103111302030203-1320233131110123"></a>

## cloudfront.disable_js_insert — disable_js_insert / 211302313311 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- cloudfront.disable_js_insert

<a id="canonical-1230102223233022-0213201032200202-2121230130310300-0201101222312303-3200023102330103-3110320331321110-2320112112030201-0211233121022103"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable js insert.

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
disable_js_insert = {}
```

<a id="canonical-1221201311231123-1021312111113330-1000201320001210-0301030110023300-0331100301011123-2201211031210033-1311101332313321-0113232222133121"></a>

## Direct properties — disable_js_insert / 211302313311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3013033302022131-0101310002123032-3032223323021221-1220110002211233-2031122210321211-0301320103022121-3113021112221230-0310133102332111"></a>

## Next pages — disable_js_insert / 211302313311 / 4

- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-0021320220121330-1123001003212312-0103101010131121-2011033121210101-0313032101220000-3302332203012213-3111202202113020-2302220232120222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110320313111303-2122220020121112-1213102330220110-1202023010333311-0222130031200202-3101112010003321-0132003333103303-2201113312032201"></a>

## cloudfront.disable_mobile_sdk — disable_mobile_sdk / 232332102032 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- cloudfront.disable_mobile_sdk

<a id="canonical-0023020223203211-3222322130010222-1303112203011322-0002030320131221-3231003133223110-1013210213202031-3320203303220120-2110212223330220"></a>

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
disable_mobile_sdk = {}
```

<a id="canonical-1110323301301011-1302332320313331-3003030301011001-0301112033030203-0112312331311311-0110220221113202-2303330132103200-1230302011230301"></a>

## Direct properties — disable_mobile_sdk / 232332102032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1330010210021220-2333120310032110-3013210330213233-2012321130100320-1310023032030201-1231120033312222-0011023311231330-2320013031300313"></a>

## Next pages — disable_mobile_sdk / 232332102032 / 4

- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-0231330302202122-3333330110232010-0233303200130220-0301122123331132-2200201002333200-2020222112331210-1313021231112202-3002213230012210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030023123232121-1031322331020022-2231212210012123-1201213200231323-2321222302320303-3020021213211000-2011302101102022-0012011232103130"></a>

## cloudfront.js_insertion_rules — js_insertion_rules / 331011130110 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- cloudfront.js_insertion_rules

<a id="canonical-3210233031332321-1322302011001003-1221210302212332-2320120111323113-0031000121030112-2102222233000002-3200320002220101-0130303100011223"></a>

Type: `"object"`. single nested block, Optional.

Defines custom JavaScript insertion rules for Bot Defense Policy.

Upstream description:

This defines custom JavaScript insertion rules for Bot Defense Policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rules")}
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
js_insertion_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-0221200312013222-3011110100113033-1300002113020231-0223130122032113-3222200220233331-3031110212111321-0321230012010120-1030102300023031"></a>

## Direct properties — js_insertion_rules / 331011130110 / 3

- [exclude_list](resources--protected_application--reference--group-002.md#canonical-3030120312130131-3321220122223102-0002133300030211-3213101323223323-3213011023030033-1101103231223220-2013022123312013-2102222320232103): complete subsection reference.

<a id="canonical-2202113333201211-0202033333110130-3122112132112320-0102033330112232-0120131210132022-3323330202111303-0032112220023001-0230103022020231"></a>

<a id="canonical-3110201222120110-1021002102002203-0012131211102012-2113100120101232-1113121023112013-0233111101332112-0230210131131102-2120000221132112"></a>

## javascript_location property — js_insertion_rules / 331011130110 / 4

Type: `"string"`. Optional.

\[Enum: JAVA\_SCRIPT\_LOCATION\_UNDEFINED|AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside
networks. - JAVA\_SCRIPT\_LOCATION\_UNDEFINED: JAVA\_SCRIPT\_LOCATION\_UNDEFINED Undefined Insert
JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript
before first tag. Possible values are \`JAVA\_SCRIPT\_LOCATION\_UNDEFINED\`, \`AFTER\_HEAD\`,
\`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to \`JAVA\_SCRIPT\_LOCATION\_UNDEFINED\`.

Upstream description:

All inside networks.

&#8203;- JAVA\_SCRIPT\_LOCATION\_UNDEFINED: JAVA\_SCRIPT\_LOCATION\_UNDEFINED

Undefined Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag.
Insert JavaScript before first tag.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("JAVA_SCRIPT_LOCATION_UNDEFINED",
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "JAVA_SCRIPT_LOCATION_UNDEFINED",
  "enum": [
    "JAVA_SCRIPT_LOCATION_UNDEFINED",
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

<a id="canonical-1130113100230033-3222100303210203-2033033033211020-2231210332003110-1021200220210211-3022213130222300-0011333102002313-3021131033130331"></a>

<a id="canonical-2023100011023113-2301213212211013-2211230200313231-3202332111112312-2002110221200130-0013000113231022-2302332022212020-1112010020321111"></a>

## javascript_mode property — js_insertion_rules / 331011130110 / 5

Type: `"string"`. Optional.

\[Enum: ASYNC\_JS\_NO\_CACHING|ASYNC\_JS\_CACHING|SYNC\_JS\_NO\_CACHING|SYNC\_JS\_CACHING\] Web
Client JavaScript Mode. Bot Defense JavaScript for telemetry collection is requested asynchronously,
and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested asynchronously,
and it is cacheable Bot Defense JavaScript for telemetry collection is requested.. Possible values
are \`ASYNC\_JS\_NO\_CACHING\`, \`ASYNC\_JS\_CACHING\`, \`SYNC\_JS\_NO\_CACHING\`,
\`SYNC\_JS\_CACHING\`. Defaults to \`ASYNC\_JS\_NO\_CACHING\`.

Upstream description:

Web Client JavaScript Mode.

Bot Defense JavaScript for telemetry collection is requested asynchronously, and it is non-cacheable
Bot Defense JavaScript for telemetry collection is requested asynchronously, and it is cacheable Bot
Defense JavaScript for telemetry collection is requested synchronously, and it is non-cacheable Bot
Defense JavaScript for telemetry collection is requested synchronously, and it is cacheable.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ASYNC_JS_NO_CACHING",
    "ASYNC_JS_CACHING",
    "SYNC_JS_NO_CACHING",
    "SYNC_JS_CACHING"),
}
```

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

<a id="canonical-1220121003202302-1221023020103302-3333133331011023-2113030202202010-3120020131211223-3110113030200310-3210233230210213-1321322232010110"></a>

<a id="canonical-3122103220221302-2323303002100333-1231013302122203-0321131121301221-2320203133302202-0331222231333012-1303033301233202-3303000333023123"></a>

## js_download_path property — js_insertion_rules / 331011130110 / 6

Type: `"string"`. Optional.

Web client will fetch F5 Client JavaScript from this path. This path must not conflict with any
other website/application paths. If not specified, default to ‘/common.js’.

Upstream description:

Web client will fetch F5 Client JavaScript from this path. This path must not conflict with any
other website/application paths.

If not specified, default to ‘/common.js’.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  }
}
```

- [rules](resources--protected_application--reference--group-002.md#canonical-1111120223012133-1010232232112011-0002033223202121-2202200003200333-0330223310201233-0110011213302203-0110111123233202-0232200123131110): complete subsection reference.

<a id="canonical-3220132103003000-2323011202312020-1031021020111321-1202230221000030-1123303222103323-2231000201202000-0121220010120111-0332311030232323"></a>

## Next pages — js_insertion_rules / 331011130110 / 7

- [cloudfront.js_insertion_rules.exclude_list](resources--protected_application--reference--group-002.md#canonical-3030120312130131-3321220122223102-0002133300030211-3213101323223323-3213011023030033-1101103231223220-2013022123312013-2102222320232103)
- [cloudfront.js_insertion_rules.rules](resources--protected_application--reference--group-002.md#canonical-1111120223012133-1010232232112011-0002033223202121-2202200003200333-0330223310201233-0110011213302203-0110111123233202-0232200123131110)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-3030120312130131-3321220122223102-0002133300030211-3213101323223323-3213011023030033-1101103231223220-2013022123312013-2102222320232103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321210221010021-1110200202232310-1322233333221003-3122120213203332-0112002322032132-3003312112130311-1300303332200123-1130123130113313"></a>

## cloudfront.js_insertion_rules.exclude_list — exclude_list / 000331103303 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-0231330302202122-3333330110232010-0233303200130220-0301122123331132-2200201002333200-2020222112331210-1313021231112202-3002213230012210)
- cloudfront.js_insertion_rules.exclude_list

<a id="canonical-1101033233220033-1113023000031000-0030302122212213-2022303303330232-2112200022103331-3100233310302212-1303323211233213-0031011213222123"></a>

Type: `"object"`. list nested block, Optional.

Optional JavaScript insertions exclude list of domain and path matchers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain")}
```

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
exclude_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3131003331130003-0332322300320220-1033131300302223-1101011231223312-3302232302333000-2110130000102013-1213023103002013-2232120133003023"></a>

## Direct properties — exclude_list / 000331103303 / 3

- [any_domain](resources--protected_application--reference--group-002.md#canonical-0022030103303102-2202331303322220-3022121202010330-1031302103233020-2112220330332130-0010031203101232-2212333311021031-2113120120013133): complete subsection reference.

- [domain](resources--protected_application--reference--group-002.md#canonical-0223132131303030-1021321002133031-1322031020210023-3021012231212331-3323212220200021-3333310133301130-1033301220110133-0232301203331323): complete subsection reference.

- [metadata](resources--protected_application--reference--group-002.md#canonical-2010012310003231-3020100210200133-2210333023120033-1323211001220022-1201221122322313-0012223022121231-2311120202101323-3322032232313013): complete subsection reference.

- [path](resources--protected_application--reference--group-002.md#canonical-2021213021001211-1331122120210300-3310333130003330-2200033112003033-1101020130331102-3033121122222031-1213332313232000-1103013002303331): complete subsection reference.

<a id="canonical-1331323022122312-3023021131030310-1102310331211333-2121232000121223-0030100331301311-1220230113020223-1020213020333201-3003032313323230"></a>

## Next pages — exclude_list / 000331103303 / 4

- [cloudfront.js_insertion_rules.exclude_list.any_domain](resources--protected_application--reference--group-002.md#canonical-0022030103303102-2202331303322220-3022121202010330-1031302103233020-2112220330332130-0010031203101232-2212333311021031-2113120120013133)
- [cloudfront.js_insertion_rules.exclude_list.domain](resources--protected_application--reference--group-002.md#canonical-0223132131303030-1021321002133031-1322031020210023-3021012231212331-3323212220200021-3333310133301130-1033301220110133-0232301203331323)
- [cloudfront.js_insertion_rules.exclude_list.metadata](resources--protected_application--reference--group-002.md#canonical-2010012310003231-3020100210200133-2210333023120033-1323211001220022-1201221122322313-0012223022121231-2311120202101323-3322032232313013)
- [cloudfront.js_insertion_rules.exclude_list.path](resources--protected_application--reference--group-002.md#canonical-2021213021001211-1331122120210300-3310333130003330-2200033112003033-1101020130331102-3033121122222031-1213332313232000-1103013002303331)
- [cloudfront.js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-0231330302202122-3333330110232010-0233303200130220-0301122123331132-2200201002333200-2020222112331210-1313021231112202-3002213230012210)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-0022030103303102-2202331303322220-3022121202010330-1031302103233020-2112220330332130-0010031203101232-2212333311021031-2113120120013133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210313023220013-3113033202010230-2103230033123200-3300231233202300-0222003113203300-3331200020020222-1011202331310221-3011321231303112"></a>

## cloudfront.js_insertion_rules.exclude_list.any_domain — any_domain / 302122103012 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-0231330302202122-3333330110232010-0233303200130220-0301122123331132-2200201002333200-2020222112331210-1313021231112202-3002213230012210)
- [cloudfront.js_insertion_rules.exclude_list](resources--protected_application--reference--group-002.md#canonical-3030120312130131-3321220122223102-0002133300030211-3213101323223323-3213011023030033-1101103231223220-2013022123312013-2102222320232103)
- cloudfront.js_insertion_rules.exclude_list.any_domain

<a id="canonical-3302113222033311-2330222021330030-1232232230111230-3132310211021313-0303002021030313-2302010301330002-3302331002201222-1100003210233000"></a>

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
any_domain = {}
```

<a id="canonical-0130111310330321-3333013130110012-1021302023322123-0030122233022202-1230231202003213-2002033223322332-3310331220300332-1010001222031313"></a>

## Direct properties — any_domain / 302122103012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1100310001120202-2012202220021131-3120313302203232-0303113010032111-0103100201010112-2200202313323322-1233010223022323-1112111322022103"></a>

## Next pages — any_domain / 302122103012 / 4

- [cloudfront.js_insertion_rules.exclude_list](resources--protected_application--reference--group-002.md#canonical-3030120312130131-3321220122223102-0002133300030211-3213101323223323-3213011023030033-1101103231223220-2013022123312013-2102222320232103)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-0223132131303030-1021321002133031-1322031020210023-3021012231212331-3323212220200021-3333310133301130-1033301220110133-0232301203331323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321103001302133-2233102201310302-2213300223210120-1301102103321002-0012331213013301-0112102300201103-0021220311212012-0223212231110113"></a>

## cloudfront.js_insertion_rules.exclude_list.domain — domain / 331130022203 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-0231330302202122-3333330110232010-0233303200130220-0301122123331132-2200201002333200-2020222112331210-1313021231112202-3002213230012210)
- [cloudfront.js_insertion_rules.exclude_list](resources--protected_application--reference--group-002.md#canonical-3030120312130131-3321220122223102-0002133300030211-3213101323223323-3213011023030033-1101103231223220-2013022123312013-2102222320232103)
- cloudfront.js_insertion_rules.exclude_list.domain

<a id="canonical-0110023101111330-2100003101201210-3220210301313112-3122012123320131-2203001331203003-0000122002021020-1212112020000021-0032011100222031"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-1201220123121330-1301021201313320-2011220030331130-3330112312231323-3111101110310031-3303031331211031-0112121022303001-0131210200130201"></a>

## Direct properties — domain / 331130022203 / 3

<a id="canonical-3333212031332210-1233001202012210-1000012210011312-0211203132313301-1030032121202301-0202332231300320-0210222211000213-2130312303020320"></a>

<a id="canonical-0221301003322033-2310011123133031-3303131101032111-3303223323200102-1213000131123113-3001123003100131-2130111311302123-1313310001322233"></a>

## exact_value property — domain / 331130022203 / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1212312013033113-3312110132332300-0231320133000112-0012322322210203-3020120112230002-3112222233213232-2233232100203321-0331231103333122"></a>

<a id="canonical-2122010033001033-3012120221122200-1200313000312201-0222320213012010-3203030221230310-1321100013330202-2312212232333232-1301202131313230"></a>

## regex_value property — domain / 331130022203 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3010312020320321-3021213012322230-3031023033233131-2321330012130212-0011120322101223-1232122311121032-1312213211013311-1033232103200021"></a>

<a id="canonical-0101320311212030-3313001021022322-2131333033023210-2002000113213022-0123133001212231-1021130013120131-0122003303221332-1112003113102111"></a>

## suffix_value property — domain / 331130022203 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0300332332032300-0001331133301331-0313302200030101-1011300122202003-1220232023131311-0223201230222211-2103301300122112-3022112121303300"></a>

## Next pages — domain / 331130022203 / 7

- [cloudfront.js_insertion_rules.exclude_list](resources--protected_application--reference--group-002.md#canonical-3030120312130131-3321220122223102-0002133300030211-3213101323223323-3213011023030033-1101103231223220-2013022123312013-2102222320232103)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-2010012310003231-3020100210200133-2210333023120033-1323211001220022-1201221122322313-0012223022121231-2311120202101323-3322032232313013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110201111021230-0001310021221121-2130120233221113-1320010131031113-3111320213231321-3103013322020133-3100321131023312-1200221003110213"></a>

## cloudfront.js_insertion_rules.exclude_list.metadata — metadata / 320010212203 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-0231330302202122-3333330110232010-0233303200130220-0301122123331132-2200201002333200-2020222112331210-1313021231112202-3002213230012210)
- [cloudfront.js_insertion_rules.exclude_list](resources--protected_application--reference--group-002.md#canonical-3030120312130131-3321220122223102-0002133300030211-3213101323223323-3213011023030033-1101103231223220-2013022123312013-2102222320232103)
- cloudfront.js_insertion_rules.exclude_list.metadata

<a id="canonical-0331131003120202-3321122131300201-2120330321122301-3103100301232002-0000223312033102-3121300220220003-1112331321203033-1100312310231133"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-3032203310301330-1230120323213211-3220001132100330-2011013000330111-3302323010321131-1303131321010302-0011321200100223-3333003313032032"></a>

## Direct properties — metadata / 320010212203 / 3

<a id="canonical-2100000131003312-2022111030200000-3111123023030023-2000323110003021-3132300120130122-3032321112010011-0230111020001131-1333320233222210"></a>

<a id="canonical-0232301320130230-2133330133302120-1132231130002023-1113223101122122-2321321000031322-2103121300233020-0013010323231301-1231132211021213"></a>

## description_spec property — metadata / 320010212203 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-2231111332221102-1022111101313032-1023111230022112-0233102111121013-3330311302003012-0212321203211303-3132222012130122-2203321011313212"></a>

<a id="canonical-2221021113231201-1330333123323212-3001123013210023-1320023210023332-3013001232020221-0201011101020132-0003221302013131-1100301120230111"></a>

## name property — metadata / 320010212203 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-3101002122033111-1200003332130200-0131223030200231-2233123122101301-1111001101210030-1201202331211021-0133101022002102-1302103022212032"></a>

## Next pages — metadata / 320010212203 / 6

- [cloudfront.js_insertion_rules.exclude_list](resources--protected_application--reference--group-002.md#canonical-3030120312130131-3321220122223102-0002133300030211-3213101323223323-3213011023030033-1101103231223220-2013022123312013-2102222320232103)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-2021213021001211-1331122120210300-3310333130003330-2200033112003033-1101020130331102-3033121122222031-1213332313232000-1103013002303331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130020213101230-0321202100002330-0002113203110010-2213220213132132-0333333033122230-3230013213001200-1212023233100201-3311013111220322"></a>

## cloudfront.js_insertion_rules.exclude_list.path — path / 103033113330 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-0231330302202122-3333330110232010-0233303200130220-0301122123331132-2200201002333200-2020222112331210-1313021231112202-3002213230012210)
- [cloudfront.js_insertion_rules.exclude_list](resources--protected_application--reference--group-002.md#canonical-3030120312130131-3321220122223102-0002133300030211-3213101323223323-3213011023030033-1101103231223220-2013022123312013-2102222320232103)
- cloudfront.js_insertion_rules.exclude_list.path

<a id="canonical-1212333221130231-2003301110201313-2232333231213322-1023001213302123-1301011201211301-2220213111211321-3222130333003200-3221103010032323"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-0203200230102301-3203230313200320-3033033001031333-1021000002332123-3320300302221333-1010310023311101-3233012013233013-0111210122020313"></a>

## Direct properties — path / 103033113330 / 3

<a id="canonical-1332230000021302-3310310103030120-1023000022321332-3310201023011011-2102301103111110-0033112122000113-3113220300033031-0322111222120020"></a>

<a id="canonical-1311313301133111-2012023312230102-3123232321031333-1332231023301323-3303120102333110-2021222102033223-0200330222210301-1021113230012213"></a>

## path property — path / 103033113330 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0312021110333212-0133023012130221-3313130013330000-0223010212101331-3111002101320321-2230203003231321-3322011120020212-0203202302032220"></a>

<a id="canonical-2233021100030323-1033233101110010-1323222101123320-0322002010202220-0101213320032032-3011010032000320-2000330031100200-0332220220222013"></a>

## prefix property — path / 103033113330 / 5

Type: `"string"`. Optional.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1330210002002220-3200033311123313-1313330130330212-2020313310103020-1011330330212113-3120023031110102-1102112211311022-0233010132000313"></a>

<a id="canonical-0121122303200221-0123003123200032-0131012011010020-2030011221232302-0201200300132233-1122023021322102-2033111203122030-3222331313122013"></a>

## regex property — path / 103033113330 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3332320220333300-1203020213022203-2113200103123011-2203022213233323-1222133312013313-1032313033223213-3233320230211101-0301020210100323"></a>

## Next pages — path / 103033113330 / 7

- [cloudfront.js_insertion_rules.exclude_list](resources--protected_application--reference--group-002.md#canonical-3030120312130131-3321220122223102-0002133300030211-3213101323223323-3213011023030033-1101103231223220-2013022123312013-2102222320232103)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-1111120223012133-1010232232112011-0002033223202121-2202200003200333-0330223310201233-0110011213302203-0110111123233202-0232200123131110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332220013312213-2233231223211231-1030020213300122-2013201311023122-0120223130323113-3122220120220300-0023003202222220-3033012122201302"></a>

## cloudfront.js_insertion_rules.rules — rules / 301131112033 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-0231330302202122-3333330110232010-0233303200130220-0301122123331132-2200201002333200-2020222112331210-1313021231112202-3002213230012210)
- cloudfront.js_insertion_rules.rules

<a id="canonical-2232301211223321-3221213223300310-2233200033002102-0010210130210031-1102112312112332-1322001133222310-3302100211330323-2001211333132331"></a>

Type: `"object"`. list nested block, Optional.

Required list of pages to insert Bot Defense client JavaScript.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "domain"),
  validators.ConflictingListObjectAttributes("exact_path",
    "glob"),
  validators.ConflictingListObjectAttributes("exact_path",
    "prefix"),
  validators.ConflictingListObjectAttributes("glob",
    "prefix")}
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-1031000332000132-0011030021021113-1011230112021022-2310210112301313-1332013103331220-1001223031031020-0000031002002303-2032113002113200"></a>

## Direct properties — rules / 301131112033 / 3

- [any_domain](resources--protected_application--reference--group-002.md#canonical-3113132303113221-1313322211002210-1201130223003310-2203023013231322-2212023120011313-1030032202011113-3201110011121201-2030233302323303): complete subsection reference.

- [domain](resources--protected_application--reference--group-002.md#canonical-0023212320322203-0213121232200112-2220021122213223-3223312330031332-2202130332023231-1020232012200112-3022013203103133-1303001233321310): complete subsection reference.

<a id="canonical-0323311013030222-3112110323330323-0233023332333221-2331300003222313-3013231300321133-2211110231220200-3103011023320312-1000212103101020"></a>

<a id="canonical-3330131120302223-0221013101212300-1012120133023031-3311103111123221-2211310132030123-1310120232213120-1303333320113013-2200102120101322"></a>

## exact_path property — rules / 301131112033 / 4

Type: `"string"`. Optional.

Exclusive with \[glob prefix\] Exact path value to match.

Upstream description:

Exclusive with \[glob prefix\] Exact path value to match.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2023020032320232-1121333112011120-0031322303211220-1001303331022202-1133300331031122-2330202221130231-2123101233211112-1012213100011300"></a>

<a id="canonical-1001010131301222-1132200112120101-1113233011322333-3111122120130211-1012223020312233-0132301113011202-0133211301001023-1323333111330330"></a>

## glob property — rules / 301131112033 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_path prefix\] Accepts wildcards \* to match multiple characters or ? To
match a single character.

Upstream description:

Exclusive with \[exact\_path prefix\]

Accepts wildcards \* to match multiple characters or ? To match a single character.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,256}$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,256}$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.pattern": "^[\\\\\\\"$&'*+./0-9:?@A-Z_a-z~-]{1,256}$"
  }
}
```

- [metadata](resources--protected_application--reference--group-002.md#canonical-1223022223311200-0231213311321001-3003302023112203-3232220111011133-3011123103101123-0333000231303333-3023031110022220-2000201223130122): complete subsection reference.

<a id="canonical-3223031203233032-2003012201202212-2232133100013233-3202123130232202-2330313013202000-1013211002131002-2302122321032232-0311121222322110"></a>

<a id="canonical-1330133223000303-2200003230010033-0002212220213333-2110130223203112-0102202220100220-2130213000002201-1331333012112133-3302223122102113"></a>

## prefix property — rules / 301131112033 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_path glob\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[exact\_path glob\] Path prefix to match (e.g. The value / will match on all paths)

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3133121310303202-3131020301333221-0333131121302112-3023123331022233-3200023330010222-0200313001011230-3023021221221113-2031320332011022"></a>

## Next pages — rules / 301131112033 / 7

- [cloudfront.js_insertion_rules.rules.any_domain](resources--protected_application--reference--group-002.md#canonical-3113132303113221-1313322211002210-1201130223003310-2203023013231322-2212023120011313-1030032202011113-3201110011121201-2030233302323303)
- [cloudfront.js_insertion_rules.rules.domain](resources--protected_application--reference--group-002.md#canonical-0023212320322203-0213121232200112-2220021122213223-3223312330031332-2202130332023231-1020232012200112-3022013203103133-1303001233321310)
- [cloudfront.js_insertion_rules.rules.metadata](resources--protected_application--reference--group-002.md#canonical-1223022223311200-0231213311321001-3003302023112203-3232220111011133-3011123103101123-0333000231303333-3023031110022220-2000201223130122)
- [cloudfront.js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-0231330302202122-3333330110232010-0233303200130220-0301122123331132-2200201002333200-2020222112331210-1313021231112202-3002213230012210)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-3113132303113221-1313322211002210-1201130223003310-2203023013231322-2212023120011313-1030032202011113-3201110011121201-2030233302323303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0011000131311230-0123002010132031-1132211131021333-2023132002201131-2211313222221223-1103203200132303-2210121303013003-2031102010121022"></a>

## cloudfront.js_insertion_rules.rules.any_domain — any_domain / 132321301100 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-0231330302202122-3333330110232010-0233303200130220-0301122123331132-2200201002333200-2020222112331210-1313021231112202-3002213230012210)
- [cloudfront.js_insertion_rules.rules](resources--protected_application--reference--group-002.md#canonical-1111120223012133-1010232232112011-0002033223202121-2202200003200333-0330223310201233-0110011213302203-0110111123233202-0232200123131110)
- cloudfront.js_insertion_rules.rules.any_domain

<a id="canonical-2132311030110110-1123210001112230-1120313320130332-3013011220000113-3030012003101131-1121000303311021-2113331331203333-3322333030110332"></a>

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
any_domain = {}
```

<a id="canonical-1211233130203312-0122032132120030-2021201223003233-3011003123011222-1223332230013030-3113111101110301-3111021121203113-0312130122000021"></a>

## Direct properties — any_domain / 132321301100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0222213223121312-0101331332213331-3323111000132032-0230002122000103-3132201003323032-0022012113303002-0221033000211011-3231320332102203"></a>

## Next pages — any_domain / 132321301100 / 4

- [cloudfront.js_insertion_rules.rules](resources--protected_application--reference--group-002.md#canonical-1111120223012133-1010232232112011-0002033223202121-2202200003200333-0330223310201233-0110011213302203-0110111123233202-0232200123131110)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-0023212320322203-0213121232200112-2220021122213223-3223312330031332-2202130332023231-1020232012200112-3022013203103133-1303001233321310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232120221130201-2200211330131212-1303120203212332-1000022000021101-3010031303220310-1002201330212012-3032033201331102-1332302203132031"></a>

## cloudfront.js_insertion_rules.rules.domain — domain / 032022231101 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-0231330302202122-3333330110232010-0233303200130220-0301122123331132-2200201002333200-2020222112331210-1313021231112202-3002213230012210)
- [cloudfront.js_insertion_rules.rules](resources--protected_application--reference--group-002.md#canonical-1111120223012133-1010232232112011-0002033223202121-2202200003200333-0330223310201233-0110011213302203-0110111123233202-0232200123131110)
- cloudfront.js_insertion_rules.rules.domain

<a id="canonical-0013100022002003-0302333331023221-2001211133322011-3033012111230113-0122131231331000-2302133312331300-1210000301330103-0112000111012303"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domains names.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("exact_value",
    "regex_value"),
  validators.ConflictingObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingObjectAttributes("regex_value",
    "suffix_value")}
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
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-3200120210133203-0020033003232021-3101032001333032-1333122332110000-1111011201012222-1020110200210320-2223103210222223-3021013003203202"></a>

## Direct properties — domain / 032022231101 / 3

<a id="canonical-0313103222213221-2320320302100212-1203013322231112-1203012021312131-1020203233213233-3002120033100003-2130303013103211-1310320113202121"></a>

<a id="canonical-3033012322020302-0020210230003300-2310202001122020-2000322313122322-1230013022123102-0102001302201332-2111301300033231-1103120101122110"></a>

## exact_value property — domain / 032022231101 / 4

Type: `"string"`. Optional.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1200233221202121-0322022332012233-1123223020103322-3223110021030232-1113113123120113-2112030201033010-3310211110332131-3133120321230221"></a>

<a id="canonical-3032001122020111-2311010220110301-3303003332123320-3113112212020231-0333310202010003-2223201003322333-3222323001133113-1311313333033121"></a>

## regex_value property — domain / 032022231101 / 5

Type: `"string"`. Optional.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-3210320010130001-0220111310231323-3000133310023222-0133103323012132-1122323030322200-1030012132302010-1020122231303030-2323210032031102"></a>

<a id="canonical-3322202130012122-2303101302133220-0130302132211113-2230312210033301-1303202003231212-3322021222232301-0030200201230122-0221200202013000"></a>

## suffix_value property — domain / 032022231101 / 6

Type: `"string"`. Optional.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3231003010111032-2111322202233203-0223301320021101-2310113331231020-0320120330223323-1330303232221313-2231022331011012-2220102111311311"></a>

## Next pages — domain / 032022231101 / 7

- [cloudfront.js_insertion_rules.rules](resources--protected_application--reference--group-002.md#canonical-1111120223012133-1010232232112011-0002033223202121-2202200003200333-0330223310201233-0110011213302203-0110111123233202-0232200123131110)
- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)

<a id="canonical-1223022223311200-0231213311321001-3003302023112203-3232220111011133-3011123103101123-0333000231303333-3023031110022220-2000201223130122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301101000223210-0020112331320221-0130300032302213-0313333133101000-0332311210112011-3003323031010020-2311313212101113-0111203200223100"></a>

## cloudfront.js_insertion_rules.rules.metadata — metadata / 223202103302 / 2

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md#canonical-0002322312001300-2321122230000322-1132302132120131-2310022322102121-3300111330013030-3223012202311310-1021020212202332-0020010133311002)
- [Property reference](resources--protected_application--reference--group-001.md#canonical-1112011310232010-1130212322220320-1323121033211100-0313202312023232-3332212122311300-1331303332321012-0100301021312031-1320300331013311)
- [cloudfront](resources--protected_application--reference--group-002.md#canonical-3023230230223213-0113310010212111-0023030201111012-2212202220232211-0013230310132300-3132331133230002-2331013021302322-2211302231010320)
- [cloudfront.js_insertion_rules](resources--protected_application--reference--group-002.md#canonical-0231330302202122-3333330110232010-0233303200130220-0301122123331132-2200201002333200-2020222112331210-1313021231112202-3002213230012210)
- [cloudfront.js_insertion_rules.rules](resources--protected_application--reference--group-002.md#canonical-1111120223012133-1010232232112011-0002033223202121-2202200003200333-0330223310201233-0110011213302203-0110111123233202-0232200123131110)
- cloudfront.js_insertion_rules.rules.metadata

<a id="canonical-1031010132131112-1022320021232131-3131213322032030-1122331330333012-2122113332003112-2220333131220202-0022122112001111-3303133012331302"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-0330030032123311-0113001112321121-2101302133120203-3201200132101021-2100000010030023-3223232020002321-0133130320013221-1003202132231020"></a>

## Direct properties — metadata / 223202103302 / 3

<a id="canonical-1012322111320232-0210033131222102-3113200102233330-3221112010312322-1123233302313332-1200001101020331-2231022303312001-2101130001202201"></a>

<a id="canonical-3321101333030232-1030221022331131-1030010101313013-3030233232102302-0011200101231233-0222001130301312-0003331000012233-0132013202312323"></a>

## description_spec property — metadata / 223202103302 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-1101222332210323-2303302130210230-0113003103232131-0330223020030202-3200221223103311-3303020310000332-3212333020211031-0221311202310323"></a>
