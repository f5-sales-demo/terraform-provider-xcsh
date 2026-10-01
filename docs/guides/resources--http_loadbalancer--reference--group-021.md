---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-0312330222100312-2211133223330102-1233232112213010-1210230130222312-1333313312023233-1003020032000231-0331220220221332-3020313110212331"></a>

## cookie_expiry property — clientside_action_captcha_challenge / 321230320200 / 4

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 86400),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-3301012132023002-2031313133022300-0011032233023022-2022332310020333-0333002322032032-0132321321101313-3222200000200023-1333023101100012"></a>

<a id="canonical-1023031100302310-1013233122001302-2302220330012022-1221130321212002-3012032320230132-1121201131202213-1310010100213022-2322221301313221"></a>

## custom_page property — clientside_action_captcha_challenge / 321230320200 / 5

Type: `"string"`. Optional.

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format.

Upstream description:

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". Base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-2111202101001213-2100302203211130-0113332032122133-3320200103321103-2321022233003123-3011121322121113-1112031102312211-0320320022021002"></a>

## Next pages — clientside_action_captcha_challenge / 321230320200 / 6

- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-2211021201313113-1021310021013033-1302002202111010-3312000330210330-2322112331020032-2233321200122022-2021112031123310-3213230222300211)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0333120302000032-2111010020220213-3001222100021202-3303111122030223-3301012233231223-2023303300030103-1333103133330100-2013231002002123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020210230032130-0031012331311311-1330022031221031-2012103332110332-1310001212113330-1123300320100332-0330012020220311-2323110300033321"></a>

## l7_ddos_protection.clientside_action_js_challenge — clientside_action_js_challenge / 031303000322 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-2211021201313113-1021310021013033-1302002202111010-3312000330210330-2322112331020032-2233321200122022-2021112031123310-3213230222300211)
- l7_ddos_protection.clientside_action_js_challenge

<a id="canonical-1002123331011332-0120233123332130-1132301100121330-0322323221331331-2320302021300213-0311123102121223-1020232112112213-1311132103003323"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
JavaScript. With this feature enabled, only clients that are capable of executing JavaScript(mostly
browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do..

Upstream description:

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
JavaScript.

With this feature enabled, only clients that are capable of executing JavaScript(mostly browsers)
will be allowed to complete the HTTP request.

When loadbalancer is configured to do JavaScript Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have JavaScript embedded in it. Loadbalancer
chooses a set of random numbers for every new client and sends these numbers along with an encrypted
answer with the request such that it embed these numbers as input in the JavaScript. JavaScript will
run on the requester browser and perform a complex Math operation. Script will submit the answer to
loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the
decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the
upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie
to avoid JavaScript challenge for subsequent requests.

JavaScript challenge serves following purposes \* Validate that the request is coming via a browser
that is capable for running JavaScript \* Force the browser to run a complex operation, f(X), that
requires it to spend a large number of CPU cycles. This is to slow down a potential DoS attacker by
making it difficult to launch a large request flood without having to spend even larger CPU cost at
their end.

You can enable either JavaScript challenge or Captcha challenge on a virtual host.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry",
    "js_script_delay")}
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
clientside_action_js_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-3222233230033332-1102221231013311-3120131130333202-0212310133031131-0023222002113110-1023313010223220-2031002103313200-0320231331331221"></a>

## Direct properties — clientside_action_js_challenge / 031303000322 / 3

<a id="canonical-1130330300313000-2222122322011021-0300022102312223-3323100302030303-2213012220002120-3120302120212213-1030302101312220-2103313112120001"></a>

<a id="canonical-3300300123000303-2111111333300232-2103131320110221-0110003002133022-3112212231321320-2122330332212122-0203022331230123-2322221323100300"></a>

## cookie_expiry property — clientside_action_js_challenge / 031303000322 / 4

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 86400),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-0221220121011300-1002313101300121-0211102222321213-0201232030321000-1001311012001030-1213203321203222-3033230033003012-2311331122331311"></a>

<a id="canonical-0232200303110110-2031323033023220-1320031010312223-0331223202011321-3130012121031112-3320101132133211-2123233013003212-3222212002002001"></a>

## custom_page property — clientside_action_js_challenge / 031303000322 / 5

Type: `"string"`. Optional.

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format.

Upstream description:

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". Base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3120133331312301-3010330013330133-2123130112023331-0220031111231102-0323032003300130-3022320121200321-2200100303010023-2010121123231212"></a>

<a id="canonical-3321003023132133-2233332301101032-2010131130133202-2221130022210112-3012112012200103-0323101102123311-2213302013322320-0110312031022321"></a>

## js_script_delay property — clientside_action_js_challenge / 031303000322 / 6

Type: `"number"`. Optional.

Delay introduced by JavaScript, in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1000, 60000),
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
    "minimum": 1000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-2011123313212301-0220010100323320-1231000332310031-2103202023120320-3132233203203133-1131212003222303-3310113112111212-0122332222021031"></a>

## Next pages — clientside_action_js_challenge / 031303000322 / 7

- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-2211021201313113-1021310021013033-1302002202111010-3312000330210330-2322112331020032-2233321200122022-2021112031123310-3213230222300211)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0232020023213211-0010012133112133-3201232132101111-0122310232212020-2332312100133201-3002012000121223-2333311311203210-2200010031330232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320032311220312-1133220122131321-0130201321211300-1003301021112312-1203212301212022-0110110133131030-3113022120013120-0122020100131230"></a>

## l7_ddos_protection.clientside_action_none — clientside_action_none / 022132222000 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-2211021201313113-1021310021013033-1302002202111010-3312000330210330-2322112331020032-2233321200122022-2021112031123310-3213230222300211)
- l7_ddos_protection.clientside_action_none

<a id="canonical-2232303310220223-2121002300221020-3223110321223121-2020311130310110-0313211200232231-3232133320113212-3312231203201003-2131321132222333"></a>

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
clientside_action_none = {}
```

<a id="canonical-2012030320313122-3111213132111033-3111221222211130-1301022030101212-3223030112211123-1220002131301301-2203232201000202-2121003302011002"></a>

## Direct properties — clientside_action_none / 022132222000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2113300313222232-0012133201013220-0303001213330333-2232302012222202-3101022120001022-0321001332200203-3030233231330213-3130221310300321"></a>

## Next pages — clientside_action_none / 022132222000 / 4

- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-2211021201313113-1021310021013033-1302002202111010-3312000330210330-2322112331020032-2233321200122022-2021112031123310-3213230222300211)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2032302331132210-0023203122022322-3032330330303001-2020223202033300-2030111322330001-0312200322113211-3000033301113203-3303002121321102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003103211101021-1033030222130023-1332120222012020-0111033233101102-0203302110103121-1132123122301023-0123331103222201-2203310201002311"></a>

## l7_ddos_protection.ddos_policy_custom — ddos_policy_custom / 213011230222 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-2211021201313113-1021310021013033-1302002202111010-3312000330210330-2322112331020032-2233321200122022-2021112031123310-3213230222300211)
- l7_ddos_protection.ddos_policy_custom

<a id="canonical-2023112333321212-2312212002203210-3330322233210023-0111300222230213-3130333002130102-2232322023303133-2313112231313113-0310221300132313"></a>

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
ddos_policy_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-1200220033330223-0203200220211032-2101221312222331-3000222231112133-3001110330210320-3112133003201030-1102120103320211-3211110233033100"></a>

## Direct properties — ddos_policy_custom / 213011230222 / 3

<a id="canonical-0330031133112302-3000220200103000-3110231312300030-1020012000213102-0101013233333113-1301322210230233-2032033001132021-0211230121313310"></a>

<a id="canonical-0031001100213120-1302000133313013-0220010111002022-3201033021000221-1231310322131213-3100223333300323-1111332233110131-1203232210200122"></a>

## name property — ddos_policy_custom / 213011230222 / 4

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

<a id="canonical-2103002231123130-1213230101000033-0002130112221212-2110001232211100-3333332111002003-2202010110331332-3203233112233113-0112312021230030"></a>

<a id="canonical-0303333202012310-2003321101332120-2213232301222011-1333233203000301-2223332223122010-1300011200300230-2320331222001000-0301203322013113"></a>

## namespace property — ddos_policy_custom / 213011230222 / 5

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

<a id="canonical-2313231032032102-3233331232313203-3222203321231312-2022203011101212-1323332001322310-1130010312302311-2231001122220312-0210123001113010"></a>

<a id="canonical-3111231210323310-3211313122202112-2321230132202203-3023003103112113-1200320100330000-3022221231312211-1321101311111122-0331022033230033"></a>

## tenant property — ddos_policy_custom / 213011230222 / 6

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

<a id="canonical-3032321120010311-2213112023232011-2101301332010012-3132212113111123-3012213330113311-2202130100031330-1223110022012223-2022112202310033"></a>

## Next pages — ddos_policy_custom / 213011230222 / 7

- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-2211021201313113-1021310021013033-1302002202111010-3312000330210330-2322112331020032-2233321200122022-2021112031123310-3213230222300211)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2300010132330322-2103003000322002-3021312112020131-2303000032213033-3013331212212202-0213203100113003-2201112331020132-0300213200302321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2033323202321123-1020220003302032-2220201210310301-1333030231133012-0212223321221010-2233212313023020-3110021222020332-2012202220023231"></a>

## l7_ddos_protection.ddos_policy_none — ddos_policy_none / 113103330323 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-2211021201313113-1021310021013033-1302002202111010-3312000330210330-2322112331020032-2233321200122022-2021112031123310-3213230222300211)
- l7_ddos_protection.ddos_policy_none

<a id="canonical-1030211220130021-0203311002323303-3232020002112332-0101123321211012-0001200001303032-3033202122103211-2231313012322020-1012222232213021"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ddos policy none.

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
ddos_policy_none = {}
```

<a id="canonical-0323202010112122-1311102010323013-3300333313221210-1102012203130133-2113122300221130-3310000112312110-2111211200211211-2121211122123220"></a>

## Direct properties — ddos_policy_none / 113103330323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3031023121002000-3213121133221202-0230333122230020-3320121313102332-0220320100102311-0122110210322001-1112032213313100-3132203022330002"></a>

## Next pages — ddos_policy_none / 113103330323 / 4

- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-2211021201313113-1021310021013033-1302002202111010-3312000330210330-2322112331020032-2233321200122022-2021112031123310-3213230222300211)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0301032010330233-2032013210320130-3210101323222232-0012021320302330-0223331231012320-2131320332211100-2001103200321213-1330302002102102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0311302303121112-2110133212133102-3020112212111221-0101321030033320-0202010011133332-1301113100031121-3133332312322201-1032031022111023"></a>

## l7_ddos_protection.default_rps_threshold — default_rps_threshold / 130313201011 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-2211021201313113-1021310021013033-1302002202111010-3312000330210330-2322112331020032-2233321200122022-2021112031123310-3213230222300211)
- l7_ddos_protection.default_rps_threshold

<a id="canonical-1222233123130302-3312313101101003-3032322213223300-3313002231231311-0011303300213303-3120011303321131-3231322021330020-1323122023001311"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default rps threshold.

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
default_rps_threshold = {}
```

<a id="canonical-3031110100022210-1303122233132223-0230030233322102-1102320220303000-3223110033200113-1211013002130312-2201232211231121-0320203311332100"></a>

## Direct properties — default_rps_threshold / 130313201011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3233200333112312-1120211202121233-0011112111230002-2003023131230021-1033321002112133-3003010021100111-1031213032332132-1210100231310102"></a>

## Next pages — default_rps_threshold / 130313201011 / 4

- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-2211021201313113-1021310021013033-1302002202111010-3312000330210330-2322112331020032-2233321200122022-2021112031123310-3213230222300211)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3102200321321113-1320220022300332-1021202111011303-2133113221331112-1002233311030020-3000303031300302-1122210131112013-1111010212132100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0332303110131031-1230210110022010-3131222021211120-2131321022323330-1302221133332233-0033111010113321-1012022310122033-0203020002030003"></a>

## l7_ddos_protection.mitigation_block — mitigation_block / 113301031220 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-2211021201313113-1021310021013033-1302002202111010-3312000330210330-2322112331020032-2233321200122022-2021112031123310-3213230222300211)
- l7_ddos_protection.mitigation_block

<a id="canonical-1302232310102120-3321021001223231-2300000031130123-3202303131232133-2010332011100130-0213311302321221-0323230323313212-2103123231230133"></a>

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
mitigation_block = {}
```

<a id="canonical-1202021002032133-1310222103200020-2020023212320002-3312031122311330-1310212223202202-2001023001133110-0033030103212210-1102233230321111"></a>

## Direct properties — mitigation_block / 113301031220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1333002320301222-3001123023213223-0112123013022200-2130233112000022-3122303220311333-2303202120103313-2112013012303030-0013131003313010"></a>

## Next pages — mitigation_block / 113301031220 / 4

- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-2211021201313113-1021310021013033-1302002202111010-3312000330210330-2322112331020032-2233321200122022-2021112031123310-3213230222300211)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2311312103021000-2231011313233313-0000032301302200-1110230312213123-3030113001001113-1103000230122301-3102220030133023-1103012032130232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3103003102021100-1100303230102101-0200221230122111-1031210331013210-0031001331312301-3122303020011002-1113133121332221-1122020311311210"></a>

## l7_ddos_protection.mitigation_captcha_challenge — mitigation_captcha_challenge / 123121122001 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-2211021201313113-1021310021013033-1302002202111010-3312000330210330-2322112331020032-2233321200122022-2021112031123310-3213230222300211)
- l7_ddos_protection.mitigation_captcha_challenge

<a id="canonical-2103221032121123-0022001013003020-2022320011012222-0100103012123123-1022202201202322-2120223133101130-1122000312303023-0111010011020312"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google
Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed
to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will
redirect..

Upstream description:

Enables loadbalancer to perform captcha challenge

Captcha challenge will be based on Google Recaptcha.

With this feature enabled, only clients that pass the captcha challenge will be allowed to complete
the HTTP request.

When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client
will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will
tag response header with a cookie to avoid Captcha challenge for subsequent requests.

CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally,
computers or bots are not capable of solving a captcha.

You can enable either JavaScript challenge or Captcha challenge on a virtual host.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry")}
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
mitigation_captcha_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-1312322133323101-0332213031303210-0011100321312133-1232130011220221-3032232001203020-0303033300300232-3133230002113323-1103330023202323"></a>

## Direct properties — mitigation_captcha_challenge / 123121122001 / 3

<a id="canonical-3203200311032323-2001300020100202-2311102030030231-1323223202333033-0110113301032002-2221201000022311-2202202223102211-3020312333003213"></a>

<a id="canonical-1321031011223322-0021010013003031-0032200223013310-0320120131311212-1202203322103003-1220213231221223-1223302111101300-2132200311133321"></a>

## cookie_expiry property — mitigation_captcha_challenge / 123121122001 / 4

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 86400),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-0032302013030130-1000033113220121-0202312013223001-3020312230131213-0321312312121230-3122131323032300-2320002330112121-1103223230121333"></a>

<a id="canonical-2013223133132033-3230010113001002-1320303323102010-3322111312211200-2020233113203031-2221332300221020-1332313313033310-1201220203010320"></a>

## custom_page property — mitigation_captcha_challenge / 123121122001 / 5

Type: `"string"`. Optional.

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format.

Upstream description:

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". Base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1032001233321200-0022121323011000-0321112022300302-3103001022330023-0033131022320031-2231110302033103-3212313320122332-3232121112011130"></a>

## Next pages — mitigation_captcha_challenge / 123121122001 / 6

- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-2211021201313113-1021310021013033-1302002202111010-3312000330210330-2322112331020032-2233321200122022-2021112031123310-3213230222300211)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1100330131220011-2223323303122231-3000301210002122-0002323302333200-2222221113020033-0312010003231303-0010020200123030-3130033132221212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0230320300313223-1002120200320213-2103020330323311-3230223213321211-3301031030032302-1212203300132032-3003122002121220-1000232310131202"></a>

## l7_ddos_protection.mitigation_js_challenge — mitigation_js_challenge / 001120122231 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-2211021201313113-1021310021013033-1302002202111010-3312000330210330-2322112331020032-2233321200122022-2021112031123310-3213230222300211)
- l7_ddos_protection.mitigation_js_challenge

<a id="canonical-1100012301120112-2202032112113213-3303220332023333-1011032322131031-0333322112122232-0001212113322233-3223212123022221-1010203133103031"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
JavaScript. With this feature enabled, only clients that are capable of executing JavaScript(mostly
browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do..

Upstream description:

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
JavaScript.

With this feature enabled, only clients that are capable of executing JavaScript(mostly browsers)
will be allowed to complete the HTTP request.

When loadbalancer is configured to do JavaScript Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have JavaScript embedded in it. Loadbalancer
chooses a set of random numbers for every new client and sends these numbers along with an encrypted
answer with the request such that it embed these numbers as input in the JavaScript. JavaScript will
run on the requester browser and perform a complex Math operation. Script will submit the answer to
loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the
decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the
upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie
to avoid JavaScript challenge for subsequent requests.

JavaScript challenge serves following purposes \* Validate that the request is coming via a browser
that is capable for running JavaScript \* Force the browser to run a complex operation, f(X), that
requires it to spend a large number of CPU cycles. This is to slow down a potential DoS attacker by
making it difficult to launch a large request flood without having to spend even larger CPU cost at
their end.

You can enable either JavaScript challenge or Captcha challenge on a virtual host.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry",
    "js_script_delay")}
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
mitigation_js_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-0203330321122021-2001022332020233-3030333102213233-1100013121232331-1210133200230112-2010102230233122-2002332321333232-0100001300301211"></a>

## Direct properties — mitigation_js_challenge / 001120122231 / 3

<a id="canonical-3120202110231203-0312120011330330-3323303123222302-2022101112002113-3122302020210000-0201220332103330-3031031220302121-3100303012230130"></a>

<a id="canonical-1033122232223112-0200122011313111-3103200201111113-3223112301231021-3103003310131212-3211031111120111-2121010011000331-3321200203310302"></a>

## cookie_expiry property — mitigation_js_challenge / 001120122231 / 4

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 86400),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-3130011210203130-1231200123133311-3112101002120320-1211011323312020-2111123010332100-3120232102323200-0013112132110313-0312330030003320"></a>

<a id="canonical-0332301332100322-2020231031210223-3322101023003000-0003121121203300-2200311130132133-0212233200330013-2103131303113323-2203102223223300"></a>

## custom_page property — mitigation_js_challenge / 001120122231 / 5

Type: `"string"`. Optional.

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format.

Upstream description:

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". Base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1233300201023201-3220122320031202-3010123233301300-3022210321203212-1233222013012202-3020233333001230-0203333131310113-0312120132020003"></a>

<a id="canonical-0110103001031301-3013223130132111-1301212303021100-0311133011302023-3333322111223332-3133213032223231-3223311132022120-0322200020311022"></a>

## js_script_delay property — mitigation_js_challenge / 001120122231 / 6

Type: `"number"`. Optional.

Delay introduced by JavaScript, in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1000, 60000),
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
    "minimum": 1000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-1123111033032101-3211132301331100-0323131121000123-0222100312312011-0000130202110302-2002011022011212-3022013321230100-2312332002212031"></a>

## Next pages — mitigation_js_challenge / 001120122231 / 7

- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-2211021201313113-1021310021013033-1302002202111010-3312000330210330-2322112331020032-2233321200122022-2021112031123310-3213230222300211)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3100013201201320-1030023222021000-1020231012221032-2100130020230332-3130310121332223-2110221222102103-3201213033132002-0230230101332320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002033132212230-2210131010302222-0210203011200020-1120103333202202-2331233231102311-3031133233023322-1120323311132032-3203210123221220"></a>

## least_active — least_active / 320202220301 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- least_active

<a id="canonical-0012031220101203-1331010113102002-1112333232131023-3022102022010211-3120131001301000-1311230213201110-2003321233300023-2003020103022211"></a>

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
least_active = {}
```

<a id="canonical-2223031130310221-2332012300211202-0031122320212313-2311123002232023-0102220123331031-1211312000332312-1012320320203230-3221130023313201"></a>

## Direct properties — least_active / 320202220301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1200312132313230-0220130120332122-3103202133001130-2032330200100220-2300133301201203-0213200301031020-1211332022310221-3222120122123120"></a>

## Next pages — least_active / 320202220301 / 4

- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2002001101312212-3322201121023103-1300033312332032-2332131302121221-0222023030002220-1013231333110113-1001302002032133-0122132213332322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2303222130003012-0020023231330333-2022230212123302-2010130300202011-2322003310302230-3030220031111220-0110123200330112-2120122300033203"></a>

## malware_protection_settings — malware_protection_settings / 312220021001 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- malware_protection_settings

<a id="canonical-2131132110220032-0231001231223002-1030113021221223-1220220323132330-0320302100310030-1022231121013021-1003021302003300-2223310300300221"></a>

Type: `"object"`. single nested block, Optional.

Malware Protection protects Web Apps and APIs, from malicious file uploads by scanning files in
real-time.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("malware_protection_rules")}
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
malware_protection_settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-3002332232121202-3203222232102223-2121211113232310-3002221000131323-0330113023001002-0312002300033111-1112001200332312-0231300002110001"></a>

## Direct properties — malware_protection_settings / 312220021001 / 3

- [malware_protection_rules](resources--http_loadbalancer--reference--group-021.md#canonical-0033303300121001-2012010221323132-3002001000201123-0213312213330122-0302333011220110-0323123202200322-1113312001322233-2332230322023002): complete subsection reference.

<a id="canonical-2002102321212323-1303032302222013-1231103130313301-1313003131203132-1012230331332023-1112132020313132-0100310020122332-0202120333213222"></a>

## Next pages — malware_protection_settings / 312220021001 / 4

- [malware_protection_settings.malware_protection_rules](resources--http_loadbalancer--reference--group-021.md#canonical-0033303300121001-2012010221323132-3002001000201123-0213312213330122-0302333011220110-0323123202200322-1113312001322233-2332230322023002)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0033303300121001-2012010221323132-3002001000201123-0213312213330122-0302333011220110-0323123202200322-1113312001322233-2332230322023002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320122323321302-0320020131300120-3212331330001010-2121012030303131-3303201220123033-3001201331223300-0203132131202120-3103331333313101"></a>

## malware_protection_settings.malware_protection_rules — malware_protection_rules / 013212021233 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [malware_protection_settings](resources--http_loadbalancer--reference--group-021.md#canonical-2002001101312212-3322201121023103-1300033312332032-2332131302121221-0222023030002220-1013231333110113-1001302002032133-0122132213332322)
- malware_protection_settings.malware_protection_rules

<a id="canonical-2003313323213130-3200012223333213-2123023201133202-1021313232303312-1323230003030111-3211121322032311-2213132022312010-1003333123212030"></a>

Type: `"object"`. list nested block, Optional.

Configure the match criteria to trigger Malware Protection Scan.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
malware_protection_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-3232100202010220-0123323122033302-2330333321201030-2112212201301313-1120222321000302-2020012020011011-0003311032133001-2213003331203322"></a>

## Direct properties — malware_protection_rules / 013212021233 / 3

- [action](resources--http_loadbalancer--reference--group-021.md#canonical-1103131210031313-3311231021030110-0032103021313030-3111311021013102-2030230131230231-0312110100211132-0221333000302223-0213201313023331): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-021.md#canonical-1111301311211300-2021032002103123-2030010313113220-0322312002103113-1323320333032031-3311300130232232-2120101201300121-2120202200222202): complete subsection reference.

<a id="canonical-1123211222000221-3330110210331021-3100302111013330-1320012313300220-2001132002000300-3021000111301002-1223013203011020-2312001330202201"></a>

<a id="canonical-2011300311220233-2220022121022023-1011313113323002-2323003231211321-1013011130203001-3321302130123212-0122032321200033-2101310301231202"></a>

## http_methods property — malware_protection_rules / 013212021233 / 4

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] HTTP Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

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

- [metadata](resources--http_loadbalancer--reference--group-021.md#canonical-0223322011233333-2023010130231330-0010321002000323-0122011111220022-3003332303022132-2110332221322031-3030112211110311-0332132113211233): complete subsection reference.

- [path](resources--http_loadbalancer--reference--group-021.md#canonical-3322003033222212-3013222122332232-2330201330212203-1131100100112220-3111212202333120-2013103011123302-3212033033303033-1113102122300202): complete subsection reference.

<a id="canonical-2212230111032220-1232333113233132-3001032000203021-1213122222030121-1132023032000213-1003000223002130-3222211011232103-2222120122021001"></a>

## Next pages — malware_protection_rules / 013212021233 / 5

- [malware_protection_settings.malware_protection_rules.action](resources--http_loadbalancer--reference--group-021.md#canonical-1103131210031313-3311231021030110-0032103021313030-3111311021013102-2030230131230231-0312110100211132-0221333000302223-0213201313023331)
- [malware_protection_settings.malware_protection_rules.domain](resources--http_loadbalancer--reference--group-021.md#canonical-1111301311211300-2021032002103123-2030010313113220-0322312002103113-1323320333032031-3311300130232232-2120101201300121-2120202200222202)
- [malware_protection_settings.malware_protection_rules.metadata](resources--http_loadbalancer--reference--group-021.md#canonical-0223322011233333-2023010130231330-0010321002000323-0122011111220022-3003332303022132-2110332221322031-3030112211110311-0332132113211233)
- [malware_protection_settings.malware_protection_rules.path](resources--http_loadbalancer--reference--group-021.md#canonical-3322003033222212-3013222122332232-2330201330212203-1131100100112220-3111212202333120-2013103011123302-3212033033303033-1113102122300202)
- [malware_protection_settings](resources--http_loadbalancer--reference--group-021.md#canonical-2002001101312212-3322201121023103-1300033312332032-2332131302121221-0222023030002220-1013231333110113-1001302002032133-0122132213332322)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1103131210031313-3311231021030110-0032103021313030-3111311021013102-2030230131230231-0312110100211132-0221333000302223-0213201313023331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120312102110020-2112303203133210-3202032321312202-2131113002300231-3233231330111223-2123123130101222-3101330331031031-2233211330003103"></a>

## malware_protection_settings.malware_protection_rules.action — action / 222320130033 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [malware_protection_settings](resources--http_loadbalancer--reference--group-021.md#canonical-2002001101312212-3322201121023103-1300033312332032-2332131302121221-0222023030002220-1013231333110113-1001302002032133-0122132213332322)
- [malware_protection_settings.malware_protection_rules](resources--http_loadbalancer--reference--group-021.md#canonical-0033303300121001-2012010221323132-3002001000201123-0213312213330122-0302333011220110-0323123202200322-1113312001322233-2332230322023002)
- malware_protection_settings.malware_protection_rules.action

<a id="canonical-2210212321221320-0001032121031230-1220020331000013-2320312003013120-0221222100102310-3331001130213223-2322210130011310-1302131003222021"></a>

Type: `"object"`. single nested block, Optional.

Action

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("block",
    "report")}
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
  "x-ves-oneof-field-action_choice": "[\"block\",\"report\"]"
}
```

Terraform syntax:

```terraform
action {
  # Configure direct properties listed below.
}
```

<a id="canonical-3301122032333311-2310301322031211-2321033300130120-1320310311230330-2200221112330332-3333021211110333-0020233312301223-0210131011122212"></a>

## Direct properties — action / 222320130033 / 3

- [block](resources--http_loadbalancer--reference--group-021.md#canonical-0112031212302003-3133023110301111-0110110200302313-0313313333322201-1122032312202132-1222020022232311-1220002101220110-0021030120230233): complete subsection reference.

- [report](resources--http_loadbalancer--reference--group-021.md#canonical-3112330000013332-1132020133131113-1130320100133102-1320011030300130-1201330121232231-1331233000130002-2301130201122113-0023013303320102): complete subsection reference.

<a id="canonical-2022033133011000-0232323012313122-2320012202100220-2003210120001133-1230013230122133-1332131201123030-2012211133230013-2211212213202011"></a>

## Next pages — action / 222320130033 / 4

- [malware_protection_settings.malware_protection_rules.action.block](resources--http_loadbalancer--reference--group-021.md#canonical-0112031212302003-3133023110301111-0110110200302313-0313313333322201-1122032312202132-1222020022232311-1220002101220110-0021030120230233)
- [malware_protection_settings.malware_protection_rules.action.report](resources--http_loadbalancer--reference--group-021.md#canonical-3112330000013332-1132020133131113-1130320100133102-1320011030300130-1201330121232231-1331233000130002-2301130201122113-0023013303320102)
- [malware_protection_settings.malware_protection_rules](resources--http_loadbalancer--reference--group-021.md#canonical-0033303300121001-2012010221323132-3002001000201123-0213312213330122-0302333011220110-0323123202200322-1113312001322233-2332230322023002)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0112031212302003-3133023110301111-0110110200302313-0313313333322201-1122032312202132-1222020022232311-1220002101220110-0021030120230233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300220031001223-0311320110110332-1310130312120233-2023030113122132-3132120130101133-0221103330300213-1003201131032200-2120123100233201"></a>

## malware_protection_settings.malware_protection_rules.action.block — block / 013023110011 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [malware_protection_settings](resources--http_loadbalancer--reference--group-021.md#canonical-2002001101312212-3322201121023103-1300033312332032-2332131302121221-0222023030002220-1013231333110113-1001302002032133-0122132213332322)
- [malware_protection_settings.malware_protection_rules](resources--http_loadbalancer--reference--group-021.md#canonical-0033303300121001-2012010221323132-3002001000201123-0213312213330122-0302333011220110-0323123202200322-1113312001322233-2332230322023002)
- [malware_protection_settings.malware_protection_rules.action](resources--http_loadbalancer--reference--group-021.md#canonical-1103131210031313-3311231021030110-0032103021313030-3111311021013102-2030230131230231-0312110100211132-0221333000302223-0213201313023331)
- malware_protection_settings.malware_protection_rules.action.block

<a id="canonical-2201020301210300-1131232200010033-0221011120232302-0002230200030322-2001031000023301-0133311301213023-1320222301120313-0333201123331111"></a>

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
block = {}
```

<a id="canonical-0210111333000021-2213202211130000-0013321031200310-3031311232023221-1231122302001113-0102223203013113-1112330030120112-3332323220202103"></a>

## Direct properties — block / 013023110011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0212313111121133-2032203121222120-2221131113111133-1121331001211022-1311333001321203-2300000130320011-1021032322010122-0102101000210212"></a>

## Next pages — block / 013023110011 / 4

- [malware_protection_settings.malware_protection_rules.action](resources--http_loadbalancer--reference--group-021.md#canonical-1103131210031313-3311231021030110-0032103021313030-3111311021013102-2030230131230231-0312110100211132-0221333000302223-0213201313023331)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3112330000013332-1132020133131113-1130320100133102-1320011030300130-1201330121232231-1331233000130002-2301130201122113-0023013303320102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312321032232221-1311000003110110-0332110031303133-3112301120021332-2301212303113113-1103020003332001-3332100031132212-2002232012210231"></a>

## malware_protection_settings.malware_protection_rules.action.report — report / 001100301231 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [malware_protection_settings](resources--http_loadbalancer--reference--group-021.md#canonical-2002001101312212-3322201121023103-1300033312332032-2332131302121221-0222023030002220-1013231333110113-1001302002032133-0122132213332322)
- [malware_protection_settings.malware_protection_rules](resources--http_loadbalancer--reference--group-021.md#canonical-0033303300121001-2012010221323132-3002001000201123-0213312213330122-0302333011220110-0323123202200322-1113312001322233-2332230322023002)
- [malware_protection_settings.malware_protection_rules.action](resources--http_loadbalancer--reference--group-021.md#canonical-1103131210031313-3311231021030110-0032103021313030-3111311021013102-2030230131230231-0312110100211132-0221333000302223-0213201313023331)
- malware_protection_settings.malware_protection_rules.action.report

<a id="canonical-2212332212203203-1212320102213321-1311013232130133-0110022201032212-1013203321232313-2320220023301231-2232032022002133-3111133300103200"></a>

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
report = {}
```

<a id="canonical-2230002131201121-1323322132021131-3033312213301102-3021201300300033-2212121003321033-2031132212320002-2013002102322321-0220310213323022"></a>

## Direct properties — report / 001100301231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3022030322020121-1222132123321232-2333110112010221-2332133230312222-0211231213123331-3111322230012302-3323222222213232-0101030212323022"></a>

## Next pages — report / 001100301231 / 4

- [malware_protection_settings.malware_protection_rules.action](resources--http_loadbalancer--reference--group-021.md#canonical-1103131210031313-3311231021030110-0032103021313030-3111311021013102-2030230131230231-0312110100211132-0221333000302223-0213201313023331)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1111301311211300-2021032002103123-2030010313113220-0322312002103113-1323320333032031-3311300130232232-2120101201300121-2120202200222202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320230111031033-0021130023313321-2231132321121202-0220212001113202-3103011102032113-2010231300310330-2303101311003131-2130221310120022"></a>

## malware_protection_settings.malware_protection_rules.domain — domain / 210010130122 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [malware_protection_settings](resources--http_loadbalancer--reference--group-021.md#canonical-2002001101312212-3322201121023103-1300033312332032-2332131302121221-0222023030002220-1013231333110113-1001302002032133-0122132213332322)
- [malware_protection_settings.malware_protection_rules](resources--http_loadbalancer--reference--group-021.md#canonical-0033303300121001-2012010221323132-3002001000201123-0213312213330122-0302333011220110-0323123202200322-1113312001322233-2332230322023002)
- malware_protection_settings.malware_protection_rules.domain

<a id="canonical-1003330200023001-2210003333123122-0122013213303230-0112311201111033-1303210113110223-1211222211210032-0000332302311112-0301313310321100"></a>

Type: `"object"`. single nested block, Optional.

Domain name for routing and identification.

Upstream description:

Domain to be matched.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("any_domain",
    "domain")}
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
  "x-ves-oneof-field-domain_matcher": "[\"any_domain\",\"domain\"]"
}
```

Terraform syntax:

```terraform
domain {
  # Configure direct properties listed below.
}
```

<a id="canonical-0033001231310201-2011211313333232-3211203332202212-2222113010332333-0030221313301203-3010200203333002-0121100101211333-2223321103331300"></a>

## Direct properties — domain / 210010130122 / 3

- [any_domain](resources--http_loadbalancer--reference--group-021.md#canonical-2131102033323300-2113201202300021-3330301302112221-1222123122210123-2210311031002332-0212011032023222-1332220103131100-3120311313223133): complete subsection reference.

- [domain](resources--http_loadbalancer--reference--group-021.md#canonical-3101221001133011-1222331203111231-3210213103012220-0303200021313013-1100223121022033-2131002210230220-2011121030102021-3301233222323300): complete subsection reference.

<a id="canonical-3330123022323233-2112130121120223-1310112123033233-2313120210131112-2313233321000102-3002130323023031-3332011211333331-0011213000202321"></a>

## Next pages — domain / 210010130122 / 4

- [malware_protection_settings.malware_protection_rules.domain.any_domain](resources--http_loadbalancer--reference--group-021.md#canonical-2131102033323300-2113201202300021-3330301302112221-1222123122210123-2210311031002332-0212011032023222-1332220103131100-3120311313223133)
- [malware_protection_settings.malware_protection_rules.domain.domain](resources--http_loadbalancer--reference--group-021.md#canonical-3101221001133011-1222331203111231-3210213103012220-0303200021313013-1100223121022033-2131002210230220-2011121030102021-3301233222323300)
- [malware_protection_settings.malware_protection_rules](resources--http_loadbalancer--reference--group-021.md#canonical-0033303300121001-2012010221323132-3002001000201123-0213312213330122-0302333011220110-0323123202200322-1113312001322233-2332230322023002)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2131102033323300-2113201202300021-3330301302112221-1222123122210123-2210311031002332-0212011032023222-1332220103131100-3120311313223133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3211213103013312-2320231301123021-1031102100123310-2003023322323122-0311210132023020-2000322211111323-2312010303210120-3221113033103302"></a>

## malware_protection_settings.malware_protection_rules.domain.any_domain — any_domain / 121200202222 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [malware_protection_settings](resources--http_loadbalancer--reference--group-021.md#canonical-2002001101312212-3322201121023103-1300033312332032-2332131302121221-0222023030002220-1013231333110113-1001302002032133-0122132213332322)
- [malware_protection_settings.malware_protection_rules](resources--http_loadbalancer--reference--group-021.md#canonical-0033303300121001-2012010221323132-3002001000201123-0213312213330122-0302333011220110-0323123202200322-1113312001322233-2332230322023002)
- [malware_protection_settings.malware_protection_rules.domain](resources--http_loadbalancer--reference--group-021.md#canonical-1111301311211300-2021032002103123-2030010313113220-0322312002103113-1323320333032031-3311300130232232-2120101201300121-2120202200222202)
- malware_protection_settings.malware_protection_rules.domain.any_domain

<a id="canonical-0322031120230010-0132231233010131-1130223030320130-0012013111312031-2102300232213321-1223332233210022-3022333032000113-0213202321122233"></a>

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

<a id="canonical-0200233230332320-2002203131113121-1021230123233012-3313303332013103-3221313323331130-3233320112130111-3232101201123323-3001331211213300"></a>

## Direct properties — any_domain / 121200202222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0011012012330232-1023112331033112-2222200301101232-2203131210100332-1203012331112003-1232031213333113-2110323213132002-0232132211033310"></a>

## Next pages — any_domain / 121200202222 / 4

- [malware_protection_settings.malware_protection_rules.domain](resources--http_loadbalancer--reference--group-021.md#canonical-1111301311211300-2021032002103123-2030010313113220-0322312002103113-1323320333032031-3311300130232232-2120101201300121-2120202200222202)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3101221001133011-1222331203111231-3210213103012220-0303200021313013-1100223121022033-2131002210230220-2011121030102021-3301233222323300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000002313312123-1202320023020000-0103001103320320-3002230111331002-1121211312131330-2020223313033110-1110200211200331-2020113201222121"></a>

## malware_protection_settings.malware_protection_rules.domain.domain — domain / 331212000021 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [malware_protection_settings](resources--http_loadbalancer--reference--group-021.md#canonical-2002001101312212-3322201121023103-1300033312332032-2332131302121221-0222023030002220-1013231333110113-1001302002032133-0122132213332322)
- [malware_protection_settings.malware_protection_rules](resources--http_loadbalancer--reference--group-021.md#canonical-0033303300121001-2012010221323132-3002001000201123-0213312213330122-0302333011220110-0323123202200322-1113312001322233-2332230322023002)
- [malware_protection_settings.malware_protection_rules.domain](resources--http_loadbalancer--reference--group-021.md#canonical-1111301311211300-2021032002103123-2030010313113220-0322312002103113-1323320333032031-3311300130232232-2120101201300121-2120202200222202)
- malware_protection_settings.malware_protection_rules.domain.domain

<a id="canonical-3013210230031230-0133322212032112-1131033220101310-1331302201011331-3201033311123113-1021310321302300-0223121103321312-0010011123020213"></a>

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

<a id="canonical-2333121013000221-1020001210111023-1202011122211212-2111131223210310-2121003121121030-3100032122232100-2302122231023302-3132302103000200"></a>

## Direct properties — domain / 331212000021 / 3

<a id="canonical-0211210032303032-3211321012301132-0032322101132302-2320133313332202-3032212000021000-0301113201220332-2103311121003120-1221321023000100"></a>

<a id="canonical-0302012021101102-0131332000013101-0023112012131030-1330112230110030-0311300301301102-0301200212130012-3101322013230223-0200003033310233"></a>

## exact_value property — domain / 331212000021 / 4

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

<a id="canonical-0022000221230100-0102203232003301-1303022013320203-2331022002321003-1133110232202033-1323300311131122-2203210301320111-1203321200033120"></a>

<a id="canonical-3021200302110120-0331313202302002-0203321220302311-1100013300321132-0331022322213332-1320133311233200-2000223200112012-1022232113133223"></a>

## regex_value property — domain / 331212000021 / 5

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

<a id="canonical-0010010030103301-2023113003223320-2301322311213132-3202000113112200-0033310003213221-3330200010123222-3123323002130032-2323001330233021"></a>

<a id="canonical-1111223110211300-3300023101230102-2302211111210233-0033233313003322-1211230201201323-3322023133232212-1010203321322311-3012201212022110"></a>

## suffix_value property — domain / 331212000021 / 6

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

<a id="canonical-2223202020103321-1311222332122213-2203130012120213-2123333023023133-1011210020010122-1103310221333111-3111322003210111-0121313233030021"></a>

## Next pages — domain / 331212000021 / 7

- [malware_protection_settings.malware_protection_rules.domain](resources--http_loadbalancer--reference--group-021.md#canonical-1111301311211300-2021032002103123-2030010313113220-0322312002103113-1323320333032031-3311300130232232-2120101201300121-2120202200222202)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0223322011233333-2023010130231330-0010321002000323-0122011111220022-3003332303022132-2110332221322031-3030112211110311-0332132113211233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123001330313320-2132202000120231-1103130003310213-1130303311032203-3110012102330023-1131012321330202-2120223130110103-3030330231212203"></a>

## malware_protection_settings.malware_protection_rules.metadata — metadata / 123102103212 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [malware_protection_settings](resources--http_loadbalancer--reference--group-021.md#canonical-2002001101312212-3322201121023103-1300033312332032-2332131302121221-0222023030002220-1013231333110113-1001302002032133-0122132213332322)
- [malware_protection_settings.malware_protection_rules](resources--http_loadbalancer--reference--group-021.md#canonical-0033303300121001-2012010221323132-3002001000201123-0213312213330122-0302333011220110-0323123202200322-1113312001322233-2332230322023002)
- malware_protection_settings.malware_protection_rules.metadata

<a id="canonical-3011002222322020-1030203011213323-0012132331130010-3300233120200330-1133300022131013-1012000101213112-1123303230102322-2311000301203013"></a>

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

<a id="canonical-2023300012012323-3012103223223223-0312200033131221-1232033103112223-2302321200333010-3203113310321001-1123202003232303-2312120113002200"></a>

## Direct properties — metadata / 123102103212 / 3

<a id="canonical-3120131232311310-3031022212001111-0301030021022322-2231111321233331-0302131330211212-2233103023202232-2023012122102030-2213301200310230"></a>

<a id="canonical-2332212011332330-0030320220200300-3312311123321010-0202313330101201-2210302301213030-1012221302100130-3333131123130313-3011123000122032"></a>

## description_spec property — metadata / 123102103212 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-1112012021101222-0213121100311203-2320323103333001-0303333110200111-0112232232101200-2110120122031331-1032012221033332-0123302000021301"></a>

<a id="canonical-3122310131200211-0130101033313202-1103331202023200-1023023230011333-2323203230221113-1130023103310330-2312332231112112-0313101112210120"></a>

## name property — metadata / 123102103212 / 5

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

<a id="canonical-2023333021033123-0202231320330320-1122123201123113-2111311223020331-3110230220130002-1021110301012332-3121111303200303-1203212131110102"></a>

## Next pages — metadata / 123102103212 / 6

- [malware_protection_settings.malware_protection_rules](resources--http_loadbalancer--reference--group-021.md#canonical-0033303300121001-2012010221323132-3002001000201123-0213312213330122-0302333011220110-0323123202200322-1113312001322233-2332230322023002)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3322003033222212-3013222122332232-2330201330212203-1131100100112220-3111212202333120-2013103011123302-3212033033303033-1113102122300202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123222031031323-0001232120022310-0003220002232311-3002030300112320-1213130330133220-3323201201313001-1323100212113010-0231130323131132"></a>

## malware_protection_settings.malware_protection_rules.path — path / 130132032001 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [malware_protection_settings](resources--http_loadbalancer--reference--group-021.md#canonical-2002001101312212-3322201121023103-1300033312332032-2332131302121221-0222023030002220-1013231333110113-1001302002032133-0122132213332322)
- [malware_protection_settings.malware_protection_rules](resources--http_loadbalancer--reference--group-021.md#canonical-0033303300121001-2012010221323132-3002001000201123-0213312213330122-0302333011220110-0323123202200322-1113312001322233-2332230322023002)
- malware_protection_settings.malware_protection_rules.path

<a id="canonical-3202311131310211-3011313221120113-2012111202020300-3220012011020032-3301301020212321-0333231121311020-1102232223332222-1131221100230002"></a>

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

<a id="canonical-0310020012313132-0102133333020013-1222230120000013-2300310031101301-3203001013202301-2113010321110113-2020202231230312-3130230112030120"></a>

## Direct properties — path / 130132032001 / 3

<a id="canonical-0133012320321310-0000331202130002-3332210113201130-3003232221023320-1330322333321321-0021031203113320-1333212002102112-0223210112310000"></a>

<a id="canonical-0202333333033322-2233112221211111-2111011221112123-0230000031002311-2203231232101110-0202103100010222-3202320230303323-2023010231300020"></a>

## path property — path / 130132032001 / 4

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

<a id="canonical-0022220321020031-3331133032103013-0103200003322011-2323001103313301-1120020211230213-0101100321132020-1112101110113332-1303030303203311"></a>

<a id="canonical-3323212212120333-3222000102110103-1100232332101113-0122321001313212-1313333211222333-0130123110113001-1112200300132331-2003302012212121"></a>

## prefix property — path / 130132032001 / 5

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

<a id="canonical-2223200020130020-3022200320323301-3012131323102100-1303210220121021-2320233332303330-3120222133032102-1010222103330223-2233221202322132"></a>

<a id="canonical-1012312231113312-3120311233001212-1131103232133232-3003321200131220-3003010211122130-1101330333221210-3321313031313311-1213331010333323"></a>

## regex property — path / 130132032001 / 6

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

<a id="canonical-0223301330331023-1332223031330121-1233220300121013-2032310132112230-2000231031211012-3033001112131322-2322232313331231-0032230133032112"></a>

## Next pages — path / 130132032001 / 7

- [malware_protection_settings.malware_protection_rules](resources--http_loadbalancer--reference--group-021.md#canonical-0033303300121001-2012010221323132-3002001000201123-0213312213330122-0302333011220110-0323123202200322-1113312001322233-2332230322023002)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202220202023023-1123321120321133-2321233102021112-0212022121012210-3002103202322130-3331231010302201-2032220001010132-2210203003200202"></a>

## more_option — more_option / 122013322012 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- more_option

<a id="canonical-0301121120121203-3222003213101121-0113110333321300-0301211131322210-1010300021322222-1320321112322013-1333031323000321-2130301001121303"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to define a route.

Upstream description:

This defines various OPTIONS to define a route.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_path_normalize",
    "enable_path_normalize"),
  validators.ConflictingObjectAttributes("max_requests_per_connection",
    "no_request_limit_per_connection")}
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
  "x-ves-oneof-field-max_requests_per_connection_choice": "[\"max_requests_per_connection\",\"no_request_limit_per_connection\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-strict_sni_host_header_check_choice": "[]"
}
```

Terraform syntax:

```terraform
more_option {
  # Configure direct properties listed below.
}
```

<a id="canonical-2323102323023010-1322000032200312-1030333302003230-0231333311000000-2300000130001220-1012302131213000-0201201312110112-3201321101311223"></a>

## Direct properties — more_option / 122013322012 / 3

- [buffer_policy](resources--http_loadbalancer--reference--group-021.md#canonical-3001230121131210-2320111332013012-2111012221211123-0201000121203103-2231133211130003-1333223331020330-3002111133222000-2333201100021120): complete subsection reference.

- [compression_params](resources--http_loadbalancer--reference--group-021.md#canonical-2210002000001212-0010203202310332-1010121210231301-3211000110202332-0022230020101320-1312133232033210-1220213012122321-0223023202033001): complete subsection reference.

<a id="canonical-1232310212203031-1011003112231311-0301012021213023-0202203302011012-3110031020213333-0221322012311302-0331132310033000-0133130012030103"></a>

<a id="canonical-2010213111320221-0230000102013223-0102112331232220-3320122300210003-0101111022023033-0223201222231020-2201211300132132-1003111030320213"></a>

## custom_errors property — more_option / 122013322012 / 4

Type: `["map", "string"]`. Optional.

Map of integer error codes as keys and string values that can be used to provide custom HTTP pages
for each error code. Key of the map can be either response code class or HTTP Error code. Response
code classes for key is configured as follows 3 -- for 3xx response code class 4 -- for 4xx..

Upstream description:

Map of integer error codes as keys and string values that can be used to provide custom HTTP pages
for each error code. Key of the map can be either response code class or HTTP Error code. Response
code classes for key is configured as follows 3 -- for 3xx response code class 4 -- for 4xx response
code class 5 -- for 5xx response code class Value of the map is string which represents custom HTTP
responses. Specific response code takes preference when both response code and response code class
matches for a request.

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
    "ves.io.schema.rules.map.keys.uint32.ranges": "3,4,5,300-599",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "65536",
    "ves.io.schema.rules.map.values.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.uint32.ranges": "3,4,5,300-599",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "65536",
    "ves.io.schema.rules.map.values.string.uri_ref": "true"
  }
}
```

<a id="canonical-0003010020000331-3022131020303020-0030122112122032-1231222310330031-2331302102321230-3303231000100222-0322021313023123-0203330321022032"></a>

<a id="canonical-2021200011132202-0133302000310103-1201132311200223-0123323223131033-3212113030102122-2031012232321112-1123210232211311-3331131321222003"></a>

## disable_default_error_pages property — more_option / 122013322012 / 5

Type: `"bool"`. Optional.

Disable the use of default F5XC error pages.

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

- [disable_path_normalize](resources--http_loadbalancer--reference--group-021.md#canonical-1033132001213130-2001323231130303-3302133101202300-2012320223111211-2113022313122123-3231031321132200-0022333302032210-1001120012321201): complete subsection reference.

- [enable_path_normalize](resources--http_loadbalancer--reference--group-021.md#canonical-0200102323330033-2123023013233122-3220321300123212-3020110001131003-2313312012331231-1030020232012220-3302210111313313-3333332202102203): complete subsection reference.

<a id="canonical-3133133130103310-2221113213321333-2100330301230320-3121330330021302-1012302031310233-1212213230000320-2313023232031322-1200013331201100"></a>

<a id="canonical-2121223031310030-1222133101220332-3311323101100233-0320301232213222-0120133002322103-0211002223233310-0310222223002011-3103313203111033"></a>

## idle_timeout property — more_option / 122013322012 / 6

Type: `"number"`. Optional.

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.
The stream is terminated with a HTTP 504 (Gateway Timeout) error code if no upstream response header
has been received, otherwise the stream is reset.

Upstream description:

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.
The stream is terminated with a HTTP 504 (Gateway Timeout) error code if no upstream response header
has been received, otherwise the stream is reset.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(3600000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 3600000,
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
    "ves.io.schema.rules.uint32.lte": "3600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "3600000"
  }
}
```

<a id="canonical-2320201121311013-3222113310331220-2022031231232100-2123200131322331-0022100121101013-2013333110300032-2010331000023223-0300102132301330"></a>

<a id="canonical-3332002333330303-1122111131131222-1301320323301303-1201031233222010-1332102101220211-0302101322311001-2120103130000001-2020230111001003"></a>

## max_request_header_size property — more_option / 122013322012 / 7

Type: `"number"`. Optional.

The maximum request header size for downstream connections, in KiB. A HTTP 431 (Request Header
Fields Too Large) error code is sent for requests that exceed this size. If multiple load balancers
share the same advertise\_policy, the highest value configured across all such load balancers is
used..

Upstream description:

The maximum request header size for downstream connections, in KiB. A HTTP 431 (Request Header
Fields Too Large) error code is sent for requests that exceed this size.

If multiple load balancers share the same advertise\_policy, the highest value configured across all
such load balancers is used for all the load balancers in question.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(96),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 96,
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
    "ves.io.schema.rules.uint32.lte": "96"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "96"
  }
}
```

<a id="canonical-3210333230300103-1110210333313003-2032322200012210-1233030221201010-1233313323033321-3011031120001131-2301022303033130-2303121213302013"></a>

<a id="canonical-2331331312212022-3101101033232301-3001200303323031-0312203332021332-0223201302320132-1123030031201003-2112212023310013-3111201102310320"></a>

## max_requests_per_connection property — more_option / 122013322012 / 8

Type: `"number"`. Optional.

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests a
downstream client can send over a single connection to Envoy. Enter a value &gt;=1 to define the
request limit per connection.

Upstream description:

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests a
downstream client can send over a single connection to Envoy. Enter a value &gt;=1 to define the
request limit per connection.

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

- [no_request_limit_per_connection](resources--http_loadbalancer--reference--group-021.md#canonical-1023102010030230-2332223133022232-3123113033013210-3231031103313320-3102130133132220-1233310023222222-3310320102211232-1232211101323103): complete subsection reference.

- [request_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-1002021132200112-1120200221201022-0220211301011220-0101022100322133-2202230030112132-3203321123003202-3201221131021020-2201202320300321): complete subsection reference.

<a id="canonical-0001203023021003-2201113102202312-1201013113011320-0023211001113321-3323330130202112-1231201112202203-0021331321032321-1311201312313220"></a>

<a id="canonical-1032312010120111-2301033231133103-1012121302112132-0301123012120223-0332310232101102-1121030101033333-3231313210322323-2033201031030311"></a>

## request_cookies_to_remove property — more_option / 122013322012 / 9

Type: `["list", "string"]`. Optional.

List of keys of Cookies to be removed from the HTTP request being sent towards upstream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [request_headers_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-0212010211023031-3133301021032123-0202022321022330-0302322100113223-2120010011130022-1322200102120032-0121222313312022-2202030212133020): complete subsection reference.

<a id="canonical-2000001000331333-3102211222113332-2222103130112201-2313221022102210-2322133333023210-2223320322213100-3300301032032203-0303020131202333"></a>

<a id="canonical-1203220123211112-0333213101021211-3032322322301333-1123002121131122-1132021012300312-0033101032210333-0003203203030123-2303201211200311"></a>

## request_headers_to_remove property — more_option / 122013322012 / 10

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP request being sent towards upstream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322): complete subsection reference.

<a id="canonical-1110233201010302-3032033303013001-2003113233111332-2123112122113013-1110033223033131-1123332232233112-1231013300301132-2321231333012323"></a>

<a id="canonical-0111231323010311-0322111133320230-0121120332301133-0203102022003032-1211012031312003-3031221012110323-2322011103011003-1033223120013330"></a>

## response_cookies_to_remove property — more_option / 122013322012 / 11

Type: `["list", "string"]`. Optional.

List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire
set-cookie header will be removed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_headers_to_add](resources--http_loadbalancer--reference--group-022.md#canonical-1323300001113123-3203203102212333-2210110210001301-2210001303013102-0333132110100211-1122113113020333-0111303223313232-2011231300210212): complete subsection reference.

<a id="canonical-2101231132233033-3100231200220211-0221311200303311-2121010320102122-3110000012120102-0222113301002220-0203111133213111-0033113233220222"></a>

<a id="canonical-2122211111013022-1223223131220202-0030113233332023-1300301132013133-2111311031202302-3020031001320102-0012331111020032-0012213302200003"></a>

## response_headers_to_remove property — more_option / 122013322012 / 12

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP response being sent towards downstream.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0231231001033210-1332020311213333-0002100131220221-1200111130212011-1303201312332101-3321201220100301-3231231201212220-2110222123213030"></a>

## Next pages — more_option / 122013322012 / 13

- [more_option.buffer_policy](resources--http_loadbalancer--reference--group-021.md#canonical-3001230121131210-2320111332013012-2111012221211123-0201000121203103-2231133211130003-1333223331020330-3002111133222000-2333201100021120)
- [more_option.compression_params](resources--http_loadbalancer--reference--group-021.md#canonical-2210002000001212-0010203202310332-1010121210231301-3211000110202332-0022230020101320-1312133232033210-1220213012122321-0223023202033001)
- [more_option.disable_path_normalize](resources--http_loadbalancer--reference--group-021.md#canonical-1033132001213130-2001323231130303-3302133101202300-2012320223111211-2113022313122123-3231031321132200-0022333302032210-1001120012321201)
- [more_option.enable_path_normalize](resources--http_loadbalancer--reference--group-021.md#canonical-0200102323330033-2123023013233122-3220321300123212-3020110001131003-2313312012331231-1030020232012220-3302210111313313-3333332202102203)
- [more_option.no_request_limit_per_connection](resources--http_loadbalancer--reference--group-021.md#canonical-1023102010030230-2332223133022232-3123113033013210-3231031103313320-3102130133132220-1233310023222222-3310320102211232-1232211101323103)
- [more_option.request_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-1002021132200112-1120200221201022-0220211301011220-0101022100322133-2202230030112132-3203321123003202-3201221131021020-2201202320300321)
- [more_option.request_headers_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-0212010211023031-3133301021032123-0202022321022330-0302322100113223-2120010011130022-1322200102120032-0121222313312022-2202030212133020)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- [more_option.response_headers_to_add](resources--http_loadbalancer--reference--group-022.md#canonical-1323300001113123-3203203102212333-2210110210001301-2210001303013102-0333132110100211-1122113113020333-0111303223313232-2011231300210212)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3001230121131210-2320111332013012-2111012221211123-0201000121203103-2231133211130003-1333223331020330-3002111133222000-2333201100021120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031331002021313-3013030210202021-3323032231300112-3320033031122030-1223102103113131-3230303000120210-3033101122000230-0202300300330312"></a>

## more_option.buffer_policy — buffer_policy / 110323101211 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- more_option.buffer_policy

<a id="canonical-0323113100202132-1003203131133102-3131011321202002-3311332311013221-1100322202112230-0032311233010220-3021203232003123-0033113313123013"></a>

Type: `"object"`. single nested block, Optional.

Some upstream applications are not capable of handling streamed data. This config enables buffering
the entire request before sending to upstream application. We can specify the maximum buffer size
and buffer interval with this config.

Upstream description:

Some upstream applications are not capable of handling streamed data. This config enables buffering
the entire request before sending to upstream application. We can specify the maximum buffer size
and buffer interval with this config.

Buffering can be enabled and disabled at VirtualHost and Route levels Route level buffer
configuration takes precedence.

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
buffer_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-1133310320321131-0130310310013203-0023222133331312-3232222232323100-0123230011331200-3003322323020320-2213113211332220-3232133222112121"></a>

## Direct properties — buffer_policy / 110323101211 / 3

<a id="canonical-1002021302011323-3111011323002210-2101002202032320-2020320032211003-0222123233302023-0211113131232313-0231313323001302-2001003201211100"></a>

<a id="canonical-2233130113211013-3220112232023312-0311203131113220-2333312122011212-1230311301012321-3201303130010301-0132210213112111-0313221212122303"></a>

## disabled property — buffer_policy / 110323101211 / 4

Type: `"bool"`. Optional.

Disable buffering for a particular route. This is useful when virtual-host has buffering, but we
need to disable it on a specific route. The value of this field is ignored for virtual-host.

Upstream description:

Disable buffering for a particular route. This is useful when virtual-host has buffering, but we
need to disable it on a specific route. The value of this field is ignored for virtual-host.

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

<a id="canonical-3101130222312100-3032320302031312-1000312310013232-1210330020221022-3012133230231330-3312031103200132-1201000100102223-2011332002030210"></a>

<a id="canonical-1012013211310222-3213330223312130-3111131313202333-2322333021023130-0312031213110303-2202303221210313-2020033213331120-1120110312121210"></a>

## max_request_bytes property — buffer_policy / 110323101211 / 5

Type: `"number"`. Optional.

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

Upstream description:

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(10485760),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10485760,
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
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

<a id="canonical-3030213132331231-3230332310220013-3012101120310202-2030221331333000-1232322110231123-0312202300213110-3133013020300132-3330312223133222"></a>

## Next pages — buffer_policy / 110323101211 / 6

- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2210002000001212-0010203202310332-1010121210231301-3211000110202332-0022230020101320-1312133232033210-1220213012122321-0223023202033001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0310101201233022-3012122022210300-2233012231032323-0211001231333130-2223022020230320-3322321323331031-2133201232123323-2021003221331211"></a>

## more_option.compression_params — compression_params / 330212233302 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- more_option.compression_params

<a id="canonical-1121131012300332-2021323232113211-0313320003030013-0323330333223223-2332101130220121-1033321203333320-0031313131030120-0102312030331010"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to compress dispatched data from an upstream service upon client request. The
content is compressed and then sent to the client with the appropriate headers if either response
and request allow. Only GZIP compression is supported.

Upstream description:

Enables loadbalancer to compress dispatched data from an upstream service upon client request. The
content is compressed and then sent to the client with the appropriate headers if either response
and request allow. Only GZIP compression is supported.

By default compression will be skipped when:

A request does NOT contain accept-encoding header. A request includes accept-encoding header, but it
does not contain “gzip” or “\*”. A request includes accept-encoding with “gzip” or “\*” with the
weight “q=0”. Note that the “gzip” will have a higher weight then “\*”. For example, if
accept-encoding is “gzip;q=0,\*;q=1”, the filter will not compress. But if the header is set to
“\*;q=0,gzip;q=1”, the filter will compress. A request whose accept-encoding header includes
“identity”. A response contains a content-encoding header. A response contains a cache-control
header whose value includes “no-transform”. A response contains a transfer-encoding header whose
value includes “gzip”. A response does not contain a content-type value that matches one of the
selected mime-types, which default to application/JavaScript, application/JSON,
application/xhtml+XML, image/svg+XML, text/CSS, text/HTML, text/plain, text/XML. Neither
content-length nor transfer-encoding headers are present in the response. Response size is smaller
than 30 bytes (only applicable when transfer-encoding is not chunked).

When compression is applied:

The content-length is removed from response headers. Response headers contain “transfer-encoding:
chunked” and do not contain “content-encoding” header. The “vary: accept-encoding” header is
inserted on every response.

GZIP Compression Level:

A value which is optimal balance between speed of compression and amount of compression is chosen.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("content_length")}
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
compression_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-2001321110120231-1012022031231210-2123022110210103-1331113232023201-3311230333131110-3303331130110210-2200321000213310-1232311302011000"></a>

## Direct properties — compression_params / 330212233302 / 3

<a id="canonical-2232222112120022-1210233201132300-2011232023322030-1333032011301300-2220023022001332-2330223031101322-3010113121012111-3202220023233122"></a>

<a id="canonical-0022100322230120-0222301011100323-3020233001100313-3101113310313300-3330211011213331-0222310200332110-0003133312102203-1233013003310103"></a>

## content_length property — compression_params / 330212233302 / 4

Type: `"number"`. Optional.

Minimum response length, in bytes, which will trigger compression. The. Defaults to \`30\`.

Upstream description:

Minimum response length, in bytes, which will trigger compression. The default value is 30.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(30),
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
    "minimum": 30
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "30"
  }
}
```

<a id="canonical-1222121131321102-0310221332100002-2112113033112111-2302221103312121-1120110032022002-2323333200210131-2123110011033322-3210330233013111"></a>

<a id="canonical-3102211330023033-0103330121010230-3103131112301132-0322320013210130-1111103001333103-0212013120300120-2002321021002310-1332111221310100"></a>

## content_type property — compression_params / 330212233302 / 5

Type: `["list", "string"]`. Optional.

Set of strings that allows specifying which mime-types yield compression When this field is not
defined, compression will be applied to the following mime-types: 'application/JavaScript'
'application/JSON', 'application/xhtml+XML' 'image/svg+XML' 'text/CSS' 'text/HTML' 'text/plain'
'text/XML'.

Upstream description:

Set of strings that allows specifying which mime-types yield compression When this field is not
defined, compression will be applied to the following mime-types: "application/JavaScript"
"application/JSON", "application/xhtml+XML" "image/svg+XML" "text/CSS" "text/HTML" "text/plain"
"text/XML"

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(50),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 50,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 50,
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "50",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3322132223311003-1013301032111003-1001203323301010-3313011021110132-0112231202332102-0221203300231221-0101112022231331-1033023331223303"></a>

<a id="canonical-2300312023221112-1033303322012301-3121102233302223-3023111123002321-3231012313121301-1023100212211210-1023200131032031-3100221123213030"></a>

## disable_on_etag_header property — compression_params / 330212233302 / 6

Type: `"bool"`. Optional.

If true, disables compression when the response contains an etag header. When it is false, weak
etags will be preserved and the ones that require strong validation will be removed.

Upstream description:

If true, disables compression when the response contains an etag header. When it is false, weak
etags will be preserved and the ones that require strong validation will be removed.

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

<a id="canonical-3003330301022113-0311323333232302-0130213220123233-2033201213121321-1002300201323222-3320133113223212-1223200200023102-3330201111220002"></a>

<a id="canonical-0210301311223022-1231301320112203-2101101212121231-1010003213211020-1013101131132033-3023013231201203-1210020100011011-1232000100200011"></a>

## remove_accept_encoding_header property — compression_params / 330212233302 / 7

Type: `"bool"`. Optional.

If true, removes accept-encoding from the request headers before dispatching it to the upstream so
that responses do not GET compressed before reaching the filter.

Upstream description:

If true, removes accept-encoding from the request headers before dispatching it to the upstream so
that responses do not GET compressed before reaching the filter.

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

<a id="canonical-2203233330203203-1312102103101212-3133231232130231-0010002003022103-2133313110020100-1323223303230201-0233123031123300-3000312331213110"></a>

## Next pages — compression_params / 330212233302 / 8

- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1033132001213130-2001323231130303-3302133101202300-2012320223111211-2113022313122123-3231031321132200-0022333302032210-1001120012321201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320011030332003-1102210130131200-2122223331010231-0013031100111102-1312310003202231-2132012020301230-1002330031121110-3323311221210120"></a>

## more_option.disable_path_normalize — disable_path_normalize / 310301122211 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- more_option.disable_path_normalize

<a id="canonical-0031312320132020-3012321002233212-2012033220311202-0220320331220201-0133312310033010-3001301122000121-1330321221230022-3210300033320232"></a>

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

<a id="canonical-1213103322213120-0233331202331010-2232232322023132-2313300103102122-0020221100012110-2321320113002202-1211233020000310-3303232032122201"></a>

## Direct properties — disable_path_normalize / 310301122211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3302201323012310-0330130122332033-1333120330231223-3120303111323021-2301133232011310-1232011122203233-2231321333332300-3132230110022030"></a>

## Next pages — disable_path_normalize / 310301122211 / 4

- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0200102323330033-2123023013233122-3220321300123212-3020110001131003-2313312012331231-1030020232012220-3302210111313313-3333332202102203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123301333103210-1120112122023133-1123113202100203-1311020012312212-2300003103331132-1320211031231031-1221001221220002-2032111010322021"></a>

## more_option.enable_path_normalize — enable_path_normalize / 023103311210 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- more_option.enable_path_normalize

<a id="canonical-2223132031010201-3001001330102303-1100302130200330-1131133312013011-1121233200131030-1233301013023331-2121112011101023-1102131102221132"></a>

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

<a id="canonical-3123231123112332-3130001110110101-2011233233320210-3220003303003131-0113133331310303-0113311020131230-3213032210102332-0101113331121202"></a>

## Direct properties — enable_path_normalize / 023103311210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3330221110011210-1330021131030301-3221000003102021-3112103321032230-0312131020012033-2310330030033002-2003102011211220-1032230122200030"></a>

## Next pages — enable_path_normalize / 023103311210 / 4

- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1023102010030230-2332223133022232-3123113033013210-3231031103313320-3102130133132220-1233310023222222-3310320102211232-1232211101323103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1302020313331021-1000012332111101-3213103012112100-1203022202300131-0301300133312302-2132011310302323-2120003320031101-0303322301030231"></a>

## more_option.no_request_limit_per_connection — no_request_limit_per_connection / 111300130312 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- more_option.no_request_limit_per_connection

<a id="canonical-1111012202212331-3101223002301022-1111133110302213-0300013131301101-3321310113223220-3110000330131212-0030333130303223-0233231320003111"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no request limit per connection.

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
no_request_limit_per_connection = {}
```

<a id="canonical-1333202311100001-2032010000222012-3102222013230121-0101223223011211-3223212212330023-3300021230032001-3113320100230101-2220021300101333"></a>

## Direct properties — no_request_limit_per_connection / 111300130312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0222022302210122-0330213020220103-0011311322123322-1223320111113123-2211223320200223-1211212230211313-3123300301321332-1020121311003001"></a>

## Next pages — no_request_limit_per_connection / 111300130312 / 4

- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1002021132200112-1120200221201022-0220211301011220-0101022100322133-2202230030112132-3203321123003202-3201221131021020-2201202320300321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0213121122032113-3101211223101220-3131022022331021-0200130013210121-0322011302323111-0002023010330011-1120322021101022-3103210111330220"></a>

## more_option.request_cookies_to_add — request_cookies_to_add / 020101010122 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- more_option.request_cookies_to_add

<a id="canonical-3032323122310110-3121030123113131-3023320332332001-3202231032213332-1120301133031303-2213133132222320-1230322332200312-3121212312110220"></a>

Type: `"object"`. list nested block, Optional.

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

Upstream description:

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
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
request_cookies_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-3202203222112113-2222100003202221-0121223222212233-3303003001131333-0211213023230312-1002131011212111-2121110210121022-2323232303203330"></a>

## Direct properties — request_cookies_to_add / 020101010122 / 3

<a id="canonical-3223031230130313-2102013213320213-0230313030120222-2010321020000032-1232311222103111-0110132331011332-0302012332130110-0022013131223010"></a>

<a id="canonical-0031321321102120-0311123230011103-1313313133131221-0102020211222201-3003131102032131-3213120232032310-2003102102202322-2302322333302203"></a>

## name property — request_cookies_to_add / 020101010122 / 4

Type: `"string"`. Optional.

Name. Name of the cookie in Cookie header.

Upstream description:

Name of the cookie in Cookie header.

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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0330001022311333-1112122103232003-2032332310213200-2003213210223022-3020300311111220-1112331331331130-2112213303223311-0130113123332002"></a>

<a id="canonical-3000023121322113-3131300330331122-3102113103310201-0220122000210023-3223031322102010-0320301310112022-0300210112020213-1231101300312111"></a>

## overwrite property — request_cookies_to_add / 020101010122 / 5

Type: `"bool"`. Optional.

Should the value be overwritten? If true, the value is overwritten to existing values. not
overwrite. Defaults to \`do\`.

Upstream description:

Should the value be overwritten? If true, the value is overwritten to existing values. Default value
is do not overwrite.

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

- [secret_value](resources--http_loadbalancer--reference--group-021.md#canonical-1200231202100131-3122103102312230-3321323121301302-0320302022310221-2130220131323003-3013203321222032-0213321020233320-0102321301003002): complete subsection reference.

<a id="canonical-0213331232201033-1313310132332030-1322013201332100-2001333030222210-3222212212313223-1332333113230120-1322011212110013-2112211320101332"></a>

<a id="canonical-3013320020101030-1212131123200310-0020232101203321-1023222203103212-3111323220200033-3312122112112230-0310301032210131-0220123312201203"></a>

## value property — request_cookies_to_add / 020101010122 / 6

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the Cookie header.

Upstream description:

Exclusive with \[secret\_value\] Value of the Cookie header.

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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-3112221030102021-3303002103322211-3131202030120200-3313333120111331-2330230213111133-2333212113322201-3330331333013103-0212202330321311"></a>

## Next pages — request_cookies_to_add / 020101010122 / 7

- [more_option.request_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-021.md#canonical-1200231202100131-3122103102312230-3321323121301302-0320302022310221-2130220131323003-3013203321222032-0213321020233320-0102321301003002)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1200231202100131-3122103102312230-3321323121301302-0320302022310221-2130220131323003-3013203321222032-0213321020233320-0102321301003002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2230330021021022-1332203313120100-3133122120010212-1302112213120311-2100022331100232-0103301330203220-0112030323010202-2330113201201022"></a>

## more_option.request_cookies_to_add.secret_value — secret_value / 210003231120 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.request_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-1002021132200112-1120200221201022-0220211301011220-0101022100322133-2202230030112132-3203321123003202-3201221131021020-2201202320300321)
- more_option.request_cookies_to_add.secret_value

<a id="canonical-0221121301010031-3113320022011033-2132220322220332-3301332130321120-1200230002022010-2230220333110030-3301003313012221-2133032013022121"></a>

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
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-2012220311300221-1222003220330320-3233303230211312-0211111212031303-2201211012121021-1303312023102133-3201310201030323-0320310112301332"></a>

## Direct properties — secret_value / 210003231120 / 3

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-021.md#canonical-3212223011311230-1101101021112121-0000001323101121-3121010323012221-1122033033202333-0113322212312222-0130101301111303-1012021321001103): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-021.md#canonical-1130230013202222-3321303302113110-0031022130203202-3021000132131211-1232111221120003-3111130031312020-2000132202000331-1213100002032323): complete subsection reference.

<a id="canonical-3021211230213022-2113123112221033-0333030233102233-3000001201031133-0013321011211333-1122100121210031-2102333101010123-0101322201221131"></a>

## Next pages — secret_value / 210003231120 / 4

- [more_option.request_cookies_to_add.secret_value.blindfold_secret_info](resources--http_loadbalancer--reference--group-021.md#canonical-3212223011311230-1101101021112121-0000001323101121-3121010323012221-1122033033202333-0113322212312222-0130101301111303-1012021321001103)
- [more_option.request_cookies_to_add.secret_value.clear_secret_info](resources--http_loadbalancer--reference--group-021.md#canonical-1130230013202222-3321303302113110-0031022130203202-3021000132131211-1232111221120003-3111130031312020-2000132202000331-1213100002032323)
- [more_option.request_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-1002021132200112-1120200221201022-0220211301011220-0101022100322133-2202230030112132-3203321123003202-3201221131021020-2201202320300321)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3212223011311230-1101101021112121-0000001323101121-3121010323012221-1122033033202333-0113322212312222-0130101301111303-1012021321001103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120332122010202-2121223311213013-1013112202212310-3222102301123103-0103330001012003-1312233100121012-1230322311300333-3123033121032220"></a>

## more_option.request_cookies_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 203211120301 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.request_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-1002021132200112-1120200221201022-0220211301011220-0101022100322133-2202230030112132-3203321123003202-3201221131021020-2201202320300321)
- [more_option.request_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-021.md#canonical-1200231202100131-3122103102312230-3321323121301302-0320302022310221-2130220131323003-3013203321222032-0213321020233320-0102321301003002)
- more_option.request_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-3022301301330213-3130223201112012-3221013302021031-0322101202130203-2132132022011133-1131203233132223-0333212000221302-2111122122033121"></a>

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

<a id="canonical-2222301032331100-0333300032212013-2211113323232010-1111220211231313-3211013230133130-0111220032210020-0122131323003031-0330233130012001"></a>

## Direct properties — blindfold_secret_info / 203211120301 / 3

<a id="canonical-3111210021221031-1333013030220011-2203233303330030-2221133010001302-1323213223311233-1103310232221313-1310131303113323-3212320013221302"></a>

<a id="canonical-3113213232301203-2001312222112132-2233303213301010-3311121233201223-3301002122302223-0010122222302300-1201311002001130-3002313202300130"></a>

## decryption_provider property — blindfold_secret_info / 203211120301 / 4

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

<a id="canonical-2231133030003110-1323321222321031-2031033330103012-2111312313131300-1203213220032302-2212230023201232-1233113210200320-2013032022033000"></a>

<a id="canonical-2322221022033023-3031222130130113-1232231003320232-3132321022320302-0121112103110103-0120130132321113-1131333032300220-1303230323322013"></a>

## location property — blindfold_secret_info / 203211120301 / 5

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

<a id="canonical-1310232213102210-3223031111321233-0020103120301303-1333001210012223-2323320323311300-0333003023022223-3133021032113203-2001230312210112"></a>

<a id="canonical-2330002030321300-2123232121332202-0333012200203033-1230232100031322-2001301012032213-2131212320201000-0223312323330230-1232010230131300"></a>

## store_provider property — blindfold_secret_info / 203211120301 / 6

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

<a id="canonical-2102213303113303-1130203303111023-3231313301112112-1202313300010303-3123112130003320-3232212133330213-0330100221012000-1101222133310232"></a>

## Next pages — blindfold_secret_info / 203211120301 / 7

- [more_option.request_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-021.md#canonical-1200231202100131-3122103102312230-3321323121301302-0320302022310221-2130220131323003-3013203321222032-0213321020233320-0102321301003002)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1130230013202222-3321303302113110-0031022130203202-3021000132131211-1232111221120003-3111130031312020-2000132202000331-1213100002032323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013332300332002-3322233323003231-3331023321130030-2011222310100212-3021120131123131-3311320313301111-3003031213212130-0120022013012120"></a>

## more_option.request_cookies_to_add.secret_value.clear_secret_info — clear_secret_info / 230123231222 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.request_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-1002021132200112-1120200221201022-0220211301011220-0101022100322133-2202230030112132-3203321123003202-3201221131021020-2201202320300321)
- [more_option.request_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-021.md#canonical-1200231202100131-3122103102312230-3321323121301302-0320302022310221-2130220131323003-3013203321222032-0213321020233320-0102321301003002)
- more_option.request_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-2121322231223222-1010122033302110-3022113010122321-3010323202233031-1001330112223313-2201300230003332-2302130233231133-2310330130230232"></a>

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

<a id="canonical-2010322033121001-3332213210030210-1023201300223010-3332132123200210-0130132021332322-0122012132220021-1330313221301103-2300301310301103"></a>

## Direct properties — clear_secret_info / 230123231222 / 3

<a id="canonical-3232021011321100-1120123303101023-2110102012131222-3222312323213100-1220100232132021-2100310122021321-1323313113300320-0230321331000123"></a>

<a id="canonical-1323032233103132-0132212013203323-2231313100011110-1202212010322020-3313310301010212-3310333223331030-1021101220201030-2120022301300022"></a>

## provider_ref property — clear_secret_info / 230123231222 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2110120230133213-1123313130111121-3113302322323220-3002013100112031-3002320031310001-3001332232032232-2310312301322203-1213100212331310"></a>

<a id="canonical-2013301010002322-0032211223300100-0120211012222233-0110210011131202-2033002301201021-1313333301111302-3230222013012101-1010121303113301"></a>

## URL property — clear_secret_info / 230123231222 / 5

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

<a id="canonical-1100202030010223-2232203020220113-1132033113020210-0303011333200130-3320231312113011-3202221120203001-1021202322222302-2210311100031111"></a>

## Next pages — clear_secret_info / 230123231222 / 6

- [more_option.request_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-021.md#canonical-1200231202100131-3122103102312230-3321323121301302-0320302022310221-2130220131323003-3013203321222032-0213321020233320-0102321301003002)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0212010211023031-3133301021032123-0202022321022330-0302322100113223-2120010011130022-1322200102120032-0121222313312022-2202030212133020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200321211200331-3133103132130011-1033200100021331-0113301023030102-0211302030220312-1201113212102212-1211013100133303-0213200310202021"></a>

## more_option.request_headers_to_add — request_headers_to_add / 202202310302 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- more_option.request_headers_to_add

<a id="canonical-2130203021120103-0323302203001320-0202002010230312-0313110122330201-2011033231233121-3000321130232300-1133032213022220-3130203220310303"></a>

Type: `"object"`. list nested block, Optional.

Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers
specified at this level are applied after headers from matched Route are applied.

Upstream description:

Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers
specified at this level are applied after headers from matched Route are applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
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
request_headers_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-1212031330310002-1323032323023212-2112333011033221-2323232000312223-3230311312202303-3200002302032132-1213230201332103-1112103203223130"></a>

## Direct properties — request_headers_to_add / 202202310302 / 3

<a id="canonical-0020212320213223-0320210010000203-1113302231302003-3201131003122113-2110013313100012-2310202012222131-3112201221221130-1210023301320113"></a>

<a id="canonical-1211021133301013-1203333213111012-0201213200332200-0120310130020113-1130303303310002-0231030211312021-1003201130220120-0020201333232332"></a>

## append property — request_headers_to_add / 202202310302 / 4

Type: `"bool"`. Optional.

Should the value be appended? If true, the value is appended to existing values. not append.
Defaults to \`do\`.

Upstream description:

Should the value be appended? If true, the value is appended to existing values. Default value is do
not append.

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

<a id="canonical-1223110022033123-2212001121200000-1320022222302310-0210001331302030-2020332012212331-2102103103133332-1312310002201120-0313011322323232"></a>

<a id="canonical-0213003100131323-3110332010330130-3023100310132223-2323321110213110-3323121333302121-1023201201220211-2233013332000120-3022303022333123"></a>

## name property — request_headers_to_add / 202202310302 / 5

Type: `"string"`. Optional.

Name. Name of the HTTP header.

Upstream description:

Name of the HTTP header.

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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [secret_value](resources--http_loadbalancer--reference--group-021.md#canonical-1221010331201200-0212121130100102-2130022130111032-3131212003010100-2223101302323110-1321202201013102-0230012322212320-2000322312131132): complete subsection reference.

<a id="canonical-1110312211102232-3310202211222101-1131310302321120-2210123012203213-1002322102201300-1013100102331202-0133033013212033-2310000202323203"></a>

<a id="canonical-2320000223312232-0210211203130103-0300033210332112-2030113022101330-2100311222231203-2031202022332133-2111000102000122-2202122123003323"></a>

## value property — request_headers_to_add / 202202310302 / 6

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the HTTP header.

Upstream description:

Exclusive with \[secret\_value\] Value of the HTTP header.

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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-0321232120032030-3131020210112100-3103200313001123-0031310000320113-1231321222130233-1020213003032102-0330201023001332-1130303120322330"></a>

## Next pages — request_headers_to_add / 202202310302 / 7

- [more_option.request_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-021.md#canonical-1221010331201200-0212121130100102-2130022130111032-3131212003010100-2223101302323110-1321202201013102-0230012322212320-2000322312131132)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1221010331201200-0212121130100102-2130022130111032-3131212003010100-2223101302323110-1321202201013102-0230012322212320-2000322312131132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300210310212123-2233222020333031-3011231230312201-0120221323130010-0210111322312111-3012023132313302-1131130300223130-0023232122122003"></a>

## more_option.request_headers_to_add.secret_value — secret_value / 013000202332 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.request_headers_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-0212010211023031-3133301021032123-0202022321022330-0302322100113223-2120010011130022-1322200102120032-0121222313312022-2202030212133020)
- more_option.request_headers_to_add.secret_value

<a id="canonical-3021003333212310-0212103131121332-0310032320203132-2300313111111310-2131230233223102-2301011101313202-1313112032232233-0200303231303300"></a>

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
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-2001102303323003-0101300010302232-3023311023020120-1013033013033313-2222223302313031-2233133020233112-2310332003212220-2300230022223021"></a>

## Direct properties — secret_value / 013000202332 / 3

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-021.md#canonical-2010103132210122-2112320332010301-2312113203113220-2020202102020110-3230111012202102-2013032220012301-2101232330313301-0103000202321321): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-021.md#canonical-3132322013300010-2321001232232230-2321232232031202-0223213322231303-1010330010213313-3322200112111211-2310313203111211-1103102212011201): complete subsection reference.

<a id="canonical-2110211030103003-2233310331001323-0021100010210302-0013022223100333-3012213013303102-2032333320132033-3102320200111300-1223113021110330"></a>

## Next pages — secret_value / 013000202332 / 4

- [more_option.request_headers_to_add.secret_value.blindfold_secret_info](resources--http_loadbalancer--reference--group-021.md#canonical-2010103132210122-2112320332010301-2312113203113220-2020202102020110-3230111012202102-2013032220012301-2101232330313301-0103000202321321)
- [more_option.request_headers_to_add.secret_value.clear_secret_info](resources--http_loadbalancer--reference--group-021.md#canonical-3132322013300010-2321001232232230-2321232232031202-0223213322231303-1010330010213313-3322200112111211-2310313203111211-1103102212011201)
- [more_option.request_headers_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-0212010211023031-3133301021032123-0202022321022330-0302322100113223-2120010011130022-1322200102120032-0121222313312022-2202030212133020)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2010103132210122-2112320332010301-2312113203113220-2020202102020110-3230111012202102-2013032220012301-2101232330313301-0103000202321321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210131302333200-0322030000113303-2333311001000220-3313032312330003-3113132003120030-3230312320022110-3103023102113332-2020101122310020"></a>

## more_option.request_headers_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 232232001320 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.request_headers_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-0212010211023031-3133301021032123-0202022321022330-0302322100113223-2120010011130022-1322200102120032-0121222313312022-2202030212133020)
- [more_option.request_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-021.md#canonical-1221010331201200-0212121130100102-2130022130111032-3131212003010100-2223101302323110-1321202201013102-0230012322212320-2000322312131132)
- more_option.request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-3011200322113002-2233131122133300-2233021121312112-2020310331002321-2233121212301323-3233212232131113-3313321031331223-0313022021203303"></a>

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

<a id="canonical-1103112203110333-0303332300013231-1202333031013320-3102110133211032-3111203221331023-1311023120231112-2031013103030321-2311333210031003"></a>

## Direct properties — blindfold_secret_info / 232232001320 / 3

<a id="canonical-3132133012011212-3333201101223033-1333200011013200-2211112103220322-1312332230210223-3123021012320221-1103233230221022-0000120120023122"></a>

<a id="canonical-3201113002300202-2233133121131130-2030030032111312-0212323030003033-1333132333322003-0002301222011103-1103123221032300-1331211120200100"></a>

## decryption_provider property — blindfold_secret_info / 232232001320 / 4

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

<a id="canonical-0003322130212200-0112132022121220-2310300120211131-0002100130321313-3203023101031311-3003322031012211-3112103032213222-2211202233033113"></a>

<a id="canonical-0232132030313021-0203030023102012-0103300022131232-1103321212213033-1100311321202101-2120310221323113-0231133012000222-2212132320003022"></a>

## location property — blindfold_secret_info / 232232001320 / 5

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

<a id="canonical-2000213321302213-3001121301010313-2111102023221122-1101120110300223-2312011102023000-1023120101300300-0101101103202332-0030122003200200"></a>

<a id="canonical-3032031120220332-0021110032201112-2223221322201133-2322121013210033-1021232033320201-0121322202011213-0323030111213003-2331300203121231"></a>

## store_provider property — blindfold_secret_info / 232232001320 / 6

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

<a id="canonical-1112030113312113-1123302303131021-1003012330301100-3130200233200032-3000123322101333-0201332031020211-3012103013321013-2110230203121023"></a>

## Next pages — blindfold_secret_info / 232232001320 / 7

- [more_option.request_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-021.md#canonical-1221010331201200-0212121130100102-2130022130111032-3131212003010100-2223101302323110-1321202201013102-0230012322212320-2000322312131132)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3132322013300010-2321001232232230-2321232232031202-0223213322231303-1010330010213313-3322200112111211-2310313203111211-1103102212011201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310330121210132-0010212300220133-2121302110111130-1232032132203312-2021223030121301-1330120033230203-1312021102123120-3230200000102133"></a>

## more_option.request_headers_to_add.secret_value.clear_secret_info — clear_secret_info / 232121121021 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.request_headers_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-0212010211023031-3133301021032123-0202022321022330-0302322100113223-2120010011130022-1322200102120032-0121222313312022-2202030212133020)
- [more_option.request_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-021.md#canonical-1221010331201200-0212121130100102-2130022130111032-3131212003010100-2223101302323110-1321202201013102-0230012322212320-2000322312131132)
- more_option.request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-2133110200023110-2120222330031220-0301100112232023-3201102122222321-0110322323222232-1202120200122332-1303112221121031-2330310211231230"></a>

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

<a id="canonical-3132120211310012-0230232233000030-2030303221233202-2000033213210211-0030313322102203-0212132212211213-3010120031233331-1231200131133103"></a>

## Direct properties — clear_secret_info / 232121121021 / 3

<a id="canonical-2011322302010102-0303102020100213-0032320013131201-2231012211021131-3103233102113033-3103221102001020-3302002011230212-2033223322213121"></a>

<a id="canonical-3122332031230301-2320201333331211-3230130110030231-0232330223302322-1111123311301110-3110102322112230-0033332121101323-1321013101131030"></a>

## provider_ref property — clear_secret_info / 232121121021 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1203311033303300-1012122233031113-1312310012103010-1300331312100102-0331221001103313-0101300212020310-0122132211301113-0002103211312103"></a>

<a id="canonical-1122100003020203-3231223213131033-2313333223303220-0021321302102211-2301012313123020-2011120220012321-0211230300213232-0111003210113233"></a>

## URL property — clear_secret_info / 232121121021 / 5

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

<a id="canonical-3333333310222023-1200212221202211-1333303020312213-0131221111230131-3303321123120323-2021331321122313-0000300333210302-2020010030110220"></a>

## Next pages — clear_secret_info / 232121121021 / 6

- [more_option.request_headers_to_add.secret_value](resources--http_loadbalancer--reference--group-021.md#canonical-1221010331201200-0212121130100102-2130022130111032-3131212003010100-2223101302323110-1321202201013102-0230012322212320-2000322312131132)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0110303323220233-2231211211333003-0020033013110001-2201312031021310-3130312212102122-0003120111313321-3311120121010320-2123220223311032"></a>

## more_option.response_cookies_to_add — response_cookies_to_add / 010132303111 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- more_option.response_cookies_to_add

<a id="canonical-2210232100010011-1111213113122000-0233322020322012-2233232312233130-2131020222212133-1023312011122300-2300021132030302-0231100333212010"></a>

Type: `"object"`. list nested block, Optional.

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream. Cookies specified at this level are applied after cookies from
matched Route are applied.

Upstream description:

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream. Cookies specified at this level are applied after cookies from
matched Route are applied.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("add_domain",
    "ignore_domain"),
  validators.ConflictingListObjectAttributes("add_expiry",
    "ignore_expiry"),
  validators.ConflictingListObjectAttributes("add_httponly",
    "ignore_httponly"),
  validators.ConflictingListObjectAttributes("add_partitioned",
    "ignore_partitioned"),
  validators.ConflictingListObjectAttributes("add_path",
    "ignore_path"),
  validators.ConflictingListObjectAttributes("add_secure",
    "ignore_secure"),
  validators.ConflictingListObjectAttributes("ignore_max_age",
    "max_age_value"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_lax"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_none"),
  validators.ConflictingListObjectAttributes("ignore_samesite",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("ignore_value",
    "secret_value"),
  validators.ConflictingListObjectAttributes("ignore_value",
    "value"),
  validators.ConflictingListObjectAttributes("samesite_lax",
    "samesite_none"),
  validators.ConflictingListObjectAttributes("samesite_lax",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("samesite_none",
    "samesite_strict"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
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
response_cookies_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-1131120132203003-0131101002023213-2322210021013102-0230121002002020-0321332011323231-0232302021200222-0323130120220031-2100003123212223"></a>

## Direct properties — response_cookies_to_add / 010132303111 / 3

<a id="canonical-0130232302120322-2223013233211002-3232031032002011-1320311111012221-0003301212020101-1121233322232321-1212102232120200-2013200120230011"></a>

<a id="canonical-1100233112330313-1322130103120331-1220311123332133-0310032003303002-3233201023303221-1110031211013331-3130122211213112-3121321030303330"></a>

## add_domain property — response_cookies_to_add / 010132303111 / 4

Type: `"string"`. Optional.

Exclusive with \[ignore\_domain\] Add domain attribute.

Upstream description:

Exclusive with \[ignore\_domain\] Add domain attribute.

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

<a id="canonical-0200300200321301-1021031133310101-0010323230003021-1020033320110112-3312313122230021-1102113210322113-0001201110022232-2331130200131112"></a>

<a id="canonical-0130001202300022-3010322013112103-2011300031220211-2020012130113033-3121022013033331-3230230121201213-1200000320131320-0030311102331123"></a>

## add_expiry property — response_cookies_to_add / 010132303111 / 5

Type: `"string"`. Optional.

Exclusive with \[ignore\_expiry\] Add expiry attribute.

Upstream description:

Exclusive with \[ignore\_expiry\] Add expiry attribute.

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

- [add_httponly](resources--http_loadbalancer--reference--group-021.md#canonical-1002223003312222-3323120332331013-2110013313012200-0310311031032200-3120310230321321-2220220120200110-2331000210131301-2200301302230222): complete subsection reference.

- [add_partitioned](resources--http_loadbalancer--reference--group-021.md#canonical-3032202112123222-1203330201103322-2201301210111202-2103213021030130-0312313201233030-0113100302023230-1320020002033131-3220030032003230): complete subsection reference.

<a id="canonical-1232102210233232-2012221103202222-0012133213010020-1212301033221013-3222312111212023-0303220001233332-0121020333132100-0303010021103311"></a>

<a id="canonical-0302103101323003-1012130233233110-1110102130301200-2331230010233222-2230223110213123-3013311211301311-0102213212122301-0203003020320123"></a>

## add_path property — response_cookies_to_add / 010132303111 / 6

Type: `"string"`. Optional.

Exclusive with \[ignore\_path\] Add path attribute.

Upstream description:

Exclusive with \[ignore\_path\] Add path attribute.

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

- [add_secure](resources--http_loadbalancer--reference--group-021.md#canonical-3122323312222233-0123020120213002-2320121333022110-1012310032330111-3110330322133130-2201003000130212-1113120210201203-1101300320213223): complete subsection reference.

- [ignore_domain](resources--http_loadbalancer--reference--group-021.md#canonical-0202333311301131-2012001021033211-3321233012130301-1030202123133202-1031123031220323-0001212200123130-2103021111331113-1323233233032022): complete subsection reference.

- [ignore_expiry](resources--http_loadbalancer--reference--group-021.md#canonical-3103210013332130-0130102020120301-1322322331232231-0001030203212102-2303200111200013-2031032121300230-1022231000230322-3112321032220321): complete subsection reference.

- [ignore_httponly](resources--http_loadbalancer--reference--group-021.md#canonical-3320123211213101-1112001020132233-0322212101320131-0030131211333122-0320223001223122-2312230112030311-0311303013321322-2110221023233003): complete subsection reference.

- [ignore_max_age](resources--http_loadbalancer--reference--group-021.md#canonical-1230032003021303-2222103313010210-3122331222222303-2003030321011123-3020302131020021-0213101112230101-0231203012321100-3011222032123121): complete subsection reference.

- [ignore_partitioned](resources--http_loadbalancer--reference--group-021.md#canonical-0020113202110313-3213001200332002-0210220031101122-1111021332233111-1200001221122212-2211013021230023-2100312112131202-2110003213010331): complete subsection reference.

- [ignore_path](resources--http_loadbalancer--reference--group-021.md#canonical-2203100133020033-3302330110301230-2033023031033333-3202331011111133-3131213212320220-0023330110030133-3312131320203330-0100021001110101): complete subsection reference.

- [ignore_samesite](resources--http_loadbalancer--reference--group-021.md#canonical-3102202303211003-2113120021232103-2130232022121211-2212031113323321-3122201000122330-0112011030100001-3100031111123032-1032210030321010): complete subsection reference.

- [ignore_secure](resources--http_loadbalancer--reference--group-021.md#canonical-0002020303200233-3101302111100323-3200302300131111-0013102123133200-3010102210020310-2123032023311011-3323300311033131-3330012021231230): complete subsection reference.

- [ignore_value](resources--http_loadbalancer--reference--group-021.md#canonical-3100111032103113-2230310102122122-0313103010101213-0223103301010022-0102110020033202-2223130032033032-3300020311131322-1201010312202110): complete subsection reference.

<a id="canonical-1221020110032102-2121230213101211-1003033010300232-1011013000303013-2332333311310100-2320201202103221-3302032312133102-2013231331212110"></a>

<a id="canonical-2123000130213023-2123320023301303-3112331002221123-0000113120033021-0213120211223130-3123121133132111-0230330300310313-1201203131001323"></a>

## max_age_value property — response_cookies_to_add / 010132303111 / 7

Type: `"number"`. Optional.

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Upstream description:

Exclusive with \[ignore\_max\_age\] Add max age attribute.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(34560000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 34560000,
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
    "ves.io.schema.rules.uint32.lte": "34560000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "34560000"
  }
}
```

<a id="canonical-1221221012100330-3301013232301110-3111033133030120-1211233022122112-2231332112103233-1003013323030330-0012230332221330-1311133020023031"></a>

<a id="canonical-3100031201002300-1022131011312021-0021121213323233-2020030103331101-0233101031111100-2033132030230130-0102321130131300-2300221100330331"></a>

## name property — response_cookies_to_add / 010132303111 / 8

Type: `"string"`. Optional.

Name. Name of the cookie in Cookie header.

Upstream description:

Name of the cookie in Cookie header.

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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0230221012220033-1001000131103220-1322330232031102-1120112321012320-3123022213001100-1330023002100303-3231222003312013-0322100310303210"></a>

<a id="canonical-0230220233200102-0233033012223311-3120312331123010-3003232200123133-1003323321211231-0013033233300320-3001020032011223-3100102033012300"></a>

## overwrite property — response_cookies_to_add / 010132303111 / 9

Type: `"bool"`. Optional.

Should the value be overwritten? If true, the value is overwritten to existing values. not
overwrite. Defaults to \`do\`.

Upstream description:

Should the value be overwritten? If true, the value is overwritten to existing values. Default value
is do not overwrite.

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

- [samesite_lax](resources--http_loadbalancer--reference--group-021.md#canonical-2330303333211221-2202033113332220-1120033202332121-0320332101312012-2322102332223022-1322012331213210-0001230321211032-1111320222330013): complete subsection reference.

- [samesite_none](resources--http_loadbalancer--reference--group-021.md#canonical-1332120002013122-1133230013132220-3110131310111110-3221333313023120-1213021213132300-2020010202123201-1232130222012003-1213113011012213): complete subsection reference.

- [samesite_strict](resources--http_loadbalancer--reference--group-021.md#canonical-3200311123303121-3223302333130102-2013211221321310-3331031100310202-1011011132210011-1311002320131212-2320111013112103-0311130332202121): complete subsection reference.

- [secret_value](resources--http_loadbalancer--reference--group-021.md#canonical-1322103020211312-3032032200003212-0100210121000112-1223112101202100-2220323300011112-2200212031011232-1301320100211112-0222303132332301): complete subsection reference.

<a id="canonical-0200102003030300-3133312330100212-0222100230323310-3032133001220023-0101003202010012-0122011133300122-2313322012311120-1132301323221132"></a>

<a id="canonical-3020322122320130-1000130122210131-1123120231123131-1103011003302213-3110201212001123-3321231222331212-2111221333310131-1212313000233010"></a>

## value property — response_cookies_to_add / 010132303111 / 10

Type: `"string"`. Optional.

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

Upstream description:

Exclusive with \[ignore\_value secret\_value\] Value of the Cookie header.

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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-2232121323001201-2312332203122331-0100220203000102-2302233202332133-1122231333223232-2121103022310022-0330031322312132-2213232122121100"></a>

## Next pages — response_cookies_to_add / 010132303111 / 11

- [more_option.response_cookies_to_add.add_httponly](resources--http_loadbalancer--reference--group-021.md#canonical-1002223003312222-3323120332331013-2110013313012200-0310311031032200-3120310230321321-2220220120200110-2331000210131301-2200301302230222)
- [more_option.response_cookies_to_add.add_partitioned](resources--http_loadbalancer--reference--group-021.md#canonical-3032202112123222-1203330201103322-2201301210111202-2103213021030130-0312313201233030-0113100302023230-1320020002033131-3220030032003230)
- [more_option.response_cookies_to_add.add_secure](resources--http_loadbalancer--reference--group-021.md#canonical-3122323312222233-0123020120213002-2320121333022110-1012310032330111-3110330322133130-2201003000130212-1113120210201203-1101300320213223)
- [more_option.response_cookies_to_add.ignore_domain](resources--http_loadbalancer--reference--group-021.md#canonical-0202333311301131-2012001021033211-3321233012130301-1030202123133202-1031123031220323-0001212200123130-2103021111331113-1323233233032022)
- [more_option.response_cookies_to_add.ignore_expiry](resources--http_loadbalancer--reference--group-021.md#canonical-3103210013332130-0130102020120301-1322322331232231-0001030203212102-2303200111200013-2031032121300230-1022231000230322-3112321032220321)
- [more_option.response_cookies_to_add.ignore_httponly](resources--http_loadbalancer--reference--group-021.md#canonical-3320123211213101-1112001020132233-0322212101320131-0030131211333122-0320223001223122-2312230112030311-0311303013321322-2110221023233003)
- [more_option.response_cookies_to_add.ignore_max_age](resources--http_loadbalancer--reference--group-021.md#canonical-1230032003021303-2222103313010210-3122331222222303-2003030321011123-3020302131020021-0213101112230101-0231203012321100-3011222032123121)
- [more_option.response_cookies_to_add.ignore_partitioned](resources--http_loadbalancer--reference--group-021.md#canonical-0020113202110313-3213001200332002-0210220031101122-1111021332233111-1200001221122212-2211013021230023-2100312112131202-2110003213010331)
- [more_option.response_cookies_to_add.ignore_path](resources--http_loadbalancer--reference--group-021.md#canonical-2203100133020033-3302330110301230-2033023031033333-3202331011111133-3131213212320220-0023330110030133-3312131320203330-0100021001110101)
- [more_option.response_cookies_to_add.ignore_samesite](resources--http_loadbalancer--reference--group-021.md#canonical-3102202303211003-2113120021232103-2130232022121211-2212031113323321-3122201000122330-0112011030100001-3100031111123032-1032210030321010)
- [more_option.response_cookies_to_add.ignore_secure](resources--http_loadbalancer--reference--group-021.md#canonical-0002020303200233-3101302111100323-3200302300131111-0013102123133200-3010102210020310-2123032023311011-3323300311033131-3330012021231230)
- [more_option.response_cookies_to_add.ignore_value](resources--http_loadbalancer--reference--group-021.md#canonical-3100111032103113-2230310102122122-0313103010101213-0223103301010022-0102110020033202-2223130032033032-3300020311131322-1201010312202110)
- [more_option.response_cookies_to_add.samesite_lax](resources--http_loadbalancer--reference--group-021.md#canonical-2330303333211221-2202033113332220-1120033202332121-0320332101312012-2322102332223022-1322012331213210-0001230321211032-1111320222330013)
- [more_option.response_cookies_to_add.samesite_none](resources--http_loadbalancer--reference--group-021.md#canonical-1332120002013122-1133230013132220-3110131310111110-3221333313023120-1213021213132300-2020010202123201-1232130222012003-1213113011012213)
- [more_option.response_cookies_to_add.samesite_strict](resources--http_loadbalancer--reference--group-021.md#canonical-3200311123303121-3223302333130102-2013211221321310-3331031100310202-1011011132210011-1311002320131212-2320111013112103-0311130332202121)
- [more_option.response_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-021.md#canonical-1322103020211312-3032032200003212-0100210121000112-1223112101202100-2220323300011112-2200212031011232-1301320100211112-0222303132332301)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1002223003312222-3323120332331013-2110013313012200-0310311031032200-3120310230321321-2220220120200110-2331000210131301-2200301302230222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133001013213110-1123031011032300-0323101021203312-2023211111212221-0320111020030001-3312312310231203-2113200132000013-3012211033030321"></a>

## more_option.response_cookies_to_add.add_httponly — add_httponly / 211332320123 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- more_option.response_cookies_to_add.add_httponly

<a id="canonical-2012231221230031-1033312213213311-1011332301011313-1000333300222102-0333312201323233-2212112013202102-3132010332021033-0130032302221213"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for add httponly.

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
add_httponly = {}
```

<a id="canonical-2001313230133020-0120201111013231-2022030231110102-1231203011032233-1333000201221321-1312211101203312-3012133321213023-3100022102100130"></a>

## Direct properties — add_httponly / 211332320123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1210232211321220-2320310001320000-2320011132111113-3330211330010100-2220232300231003-3312113230300201-3130213312212202-3123131120223101"></a>

## Next pages — add_httponly / 211332320123 / 4

- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3032202112123222-1203330201103322-2201301210111202-2103213021030130-0312313201233030-0113100302023230-1320020002033131-3220030032003230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2112222000211231-2202102230130012-3211120031121032-0311332132212231-0222203201122020-2002312322132122-0032120012201221-3312123030311113"></a>

## more_option.response_cookies_to_add.add_partitioned — add_partitioned / 331312212223 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- more_option.response_cookies_to_add.add_partitioned

<a id="canonical-3331113321300300-1231033130101032-2012230122101233-1213011313213312-2230000321323301-3030030212000033-0211033132320033-3210102312130211"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for add partitioned.

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
add_partitioned = {}
```

<a id="canonical-2302122303313331-1011020200313010-2132331333323312-1301311222330201-2131123332202310-0311013230130330-1233110131300230-3233132013113032"></a>

## Direct properties — add_partitioned / 331312212223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1111032002101300-1213220122100023-1033030322131000-3200232133020111-0323201030121111-1301030020202233-0202031232123310-3310301220101201"></a>

## Next pages — add_partitioned / 331312212223 / 4

- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3122323312222233-0123020120213002-2320121333022110-1012310032330111-3110330322133130-2201003000130212-1113120210201203-1101300320213223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333130113022001-0213001232121333-2230301222133101-2322031122102231-3322100122233223-2020312100220030-1133210312333201-3202032220112130"></a>

## more_option.response_cookies_to_add.add_secure — add_secure / 221021222031 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- more_option.response_cookies_to_add.add_secure

<a id="canonical-1302010002030101-3221023300011123-1111301000332113-0122212111312223-2231211220312113-2012222123110200-0122330122303020-0021300011301230"></a>

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
add_secure = {}
```

<a id="canonical-2133001102032002-2222230113103100-0203100331101021-1100010202233222-3031131313011212-2030030323120011-0301202103300122-2232320022232332"></a>

## Direct properties — add_secure / 221021222031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0130230003130300-3200123030130111-1011011201211022-3132103001100122-0133230021012201-2203310213213333-1221022002123123-1030322220213122"></a>

## Next pages — add_secure / 221021222031 / 4

- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0202333311301131-2012001021033211-3321233012130301-1030202123133202-1031123031220323-0001212200123130-2103021111331113-1323233233032022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2230233322322200-2131213232202012-1023321302201030-3311111010011300-1202321103220133-2203030023133133-2002023321111020-1013120322100001"></a>

## more_option.response_cookies_to_add.ignore_domain — ignore_domain / 000133101202 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- more_option.response_cookies_to_add.ignore_domain

<a id="canonical-2031120102133221-1330300213221023-3122133100201330-3102223311210032-0301321203002022-2300223321102303-3313302020133220-2021132103313131"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore domain.

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
ignore_domain = {}
```

<a id="canonical-1233333211022132-2210100320213323-3202132322333302-0021310113300210-2101210310120130-0132132312212210-2021022033110333-0200212020313020"></a>

## Direct properties — ignore_domain / 000133101202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3010101033232202-3223210202103003-1102312212113332-0001021032200310-2011031320231121-2023112130202023-2011211100111223-2322320223310301"></a>

## Next pages — ignore_domain / 000133101202 / 4

- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3103210013332130-0130102020120301-1322322331232231-0001030203212102-2303200111200013-2031032121300230-1022231000230322-3112321032220321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0101132120030222-3232220100021100-3002332130333313-0323110123302022-3133002112233131-1210011031302120-0033131210112202-3002122001013233"></a>

## more_option.response_cookies_to_add.ignore_expiry — ignore_expiry / 301023323220 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- more_option.response_cookies_to_add.ignore_expiry

<a id="canonical-0313102123210110-0230222123323330-0222301022112231-0120123023032000-3120221031001100-0131201201111203-0331022000120110-2210130102230122"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore expiry.

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
ignore_expiry = {}
```

<a id="canonical-1101103023112300-0133203022020112-2132013013213021-3330101011133311-1302233121033001-2312333033300332-2120002322023232-1113010221212301"></a>

## Direct properties — ignore_expiry / 301023323220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1320103231323112-3122032301211201-1130331230011321-0101202103210323-0321022303210200-1101333303300312-2321111131323330-2122232303231110"></a>

## Next pages — ignore_expiry / 301023323220 / 4

- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3320123211213101-1112001020132233-0322212101320131-0030131211333122-0320223001223122-2312230112030311-0311303013321322-2110221023233003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3120312230103001-0021031133211111-2201233113321200-0302033223032303-3210003310312230-1302302332200101-0331212213102032-0112101110010310"></a>

## more_option.response_cookies_to_add.ignore_httponly — ignore_httponly / 312211110213 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- more_option.response_cookies_to_add.ignore_httponly

<a id="canonical-1320132312121302-1302210133321031-1220333311001332-3133010020113231-0122222023112200-3200130003223233-1313223101023021-2031013030210303"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore httponly.

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
ignore_httponly = {}
```

<a id="canonical-1111303212221203-3331102332001211-0121301112002202-2321131022320032-2232231301320002-3023300001023100-3001103222002102-1220223332132311"></a>

## Direct properties — ignore_httponly / 312211110213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3320211031031100-1132300202333301-3332110201131130-3322321012033313-0231032002022213-3313131111000200-0121330003331230-2231302112311121"></a>

## Next pages — ignore_httponly / 312211110213 / 4

- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1230032003021303-2222103313010210-3122331222222303-2003030321011123-3020302131020021-0213101112230101-0231203012321100-3011222032123121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1020001223301001-0032020003000323-3212320032331002-0132233111023130-2002311212231003-3210201212013103-0301010111011223-0010001313222003"></a>

## more_option.response_cookies_to_add.ignore_max_age — ignore_max_age / 013221310000 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- more_option.response_cookies_to_add.ignore_max_age

<a id="canonical-1110023003123313-0231110112022103-1113103332000112-3121002030331232-1303222013032013-0001200022133313-1023312310030322-0223022002003323"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore max age.

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
ignore_max_age = {}
```

<a id="canonical-2030030211221330-0231332303030021-0232130010132131-2011121000232001-3212232311230320-2132012220333031-3000200120201023-3212013132120120"></a>

## Direct properties — ignore_max_age / 013221310000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0312033101311111-2111311232122113-3331122013110100-0301202112123121-3031110213300013-1323333213133122-1010021023310212-0122331131122230"></a>

## Next pages — ignore_max_age / 013221310000 / 4

- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0020113202110313-3213001200332002-0210220031101122-1111021332233111-1200001221122212-2211013021230023-2100312112131202-2110003213010331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1100010330020112-2111213210321223-0323130223332222-2301330003120201-2332312112303212-0323033232202210-0002002011113201-3001212033020022"></a>

## more_option.response_cookies_to_add.ignore_partitioned — ignore_partitioned / 223201103021 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- more_option.response_cookies_to_add.ignore_partitioned

<a id="canonical-1002210120110013-0121223220223223-0221100210312310-2013210233312213-1022321133101022-3312032020211011-3333231222010020-3313032323231333"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore partitioned.

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
ignore_partitioned = {}
```

<a id="canonical-3221330132222021-1131222310121220-2121012300132321-3202130321002212-2131032332123211-1033100212131103-1133223111102330-3211322323100330"></a>

## Direct properties — ignore_partitioned / 223201103021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3001023000213132-3303103030031033-0031203301211110-2011003030033103-0030002220012032-0320213233230123-0311122122002013-2000130321002230"></a>

## Next pages — ignore_partitioned / 223201103021 / 4

- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2203100133020033-3302330110301230-2033023031033333-3202331011111133-3131213212320220-0023330110030133-3312131320203330-0100021001110101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0232300031032013-2110122020113322-3322320330033031-2230202333213033-2000031313312121-3023221030310022-1103021112202100-3000030020301202"></a>

## more_option.response_cookies_to_add.ignore_path — ignore_path / 230321313212 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- more_option.response_cookies_to_add.ignore_path

<a id="canonical-0211300202122213-2111120002313001-0120113102300100-2332011121101122-2020133231000010-2311323322022232-2111222110000320-0222303133133230"></a>

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
ignore_path = {}
```

<a id="canonical-2101100010203200-2332333331233331-2322132313022010-1123012003013230-3010100012032312-0111001002110233-2013102102300333-2332101333020010"></a>

## Direct properties — ignore_path / 230321313212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3121213303303210-3230102310130221-3210303020311212-2233202301101033-2130233211030232-3210213020331220-2013332210102322-3203110220023221"></a>

## Next pages — ignore_path / 230321313212 / 4

- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3102202303211003-2113120021232103-2130232022121211-2212031113323321-3122201000122330-0112011030100001-3100031111123032-1032210030321010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3110103311122233-2003010222310021-1001123013020110-1022223223310230-1300201203000210-3220221111110221-0223301100210112-1032023120201033"></a>

## more_option.response_cookies_to_add.ignore_samesite — ignore_samesite / 233201103302 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- more_option.response_cookies_to_add.ignore_samesite

<a id="canonical-3030113033021231-2331301211302011-2202133021013201-3112123001132010-3312112233123333-3102110332022131-3210102222033111-2002121212231020"></a>

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
ignore_samesite = {}
```

<a id="canonical-0131120212301033-0021210122301122-2220132321113013-2223200010310100-1122233000130222-1330023303301000-3123011333113322-1002202320130023"></a>

## Direct properties — ignore_samesite / 233201103302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0130300230003232-2120000123211221-0310321111012101-0201331000022233-2211330131231332-0200333102130100-2001133121010023-3300312310111033"></a>

## Next pages — ignore_samesite / 233201103302 / 4

- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-0002020303200233-3101302111100323-3200302300131111-0013102123133200-3010102210020310-2123032023311011-3323300311033131-3330012021231230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3222310322221022-1130312011000020-1023032033121100-3303103021132211-0032112210221321-1223231012312032-0233101222121331-2211021110033113"></a>

## more_option.response_cookies_to_add.ignore_secure — ignore_secure / 213311033003 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- more_option.response_cookies_to_add.ignore_secure

<a id="canonical-0321323310001023-3011033333322330-2123131221013320-0002332020333000-2331102100333322-2120203133232032-1202203101133122-2332323002030313"></a>

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
ignore_secure = {}
```

<a id="canonical-2321101222031233-0322030233333132-3123200320303100-1032333220110112-3130103111001021-3011122012322131-2312002021212330-2023121201212302"></a>

## Direct properties — ignore_secure / 213311033003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1312121323303200-2202323123130113-2232122203121313-0032031332331302-1003300030301311-3033123310302210-2001330003033100-1223232302120213"></a>

## Next pages — ignore_secure / 213311033003 / 4

- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3100111032103113-2230310102122122-0313103010101213-0223103301010022-0102110020033202-2223130032033032-3300020311131322-1201010312202110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122223231300033-0231130001030102-2012012310330230-0202110200103000-3030132132231011-0301031321211102-0010101231302030-0231200300212200"></a>

## more_option.response_cookies_to_add.ignore_value — ignore_value / 012002322221 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- more_option.response_cookies_to_add.ignore_value

<a id="canonical-3303310332300023-2330121133120201-2303303032220230-0033333301212021-0301200210131103-0311102002332310-3130312122021030-0310102312132031"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore value.

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
ignore_value = {}
```

<a id="canonical-1100121333231100-0022102021101033-0213003301212030-2331330203110110-1202330022110022-1230100211211300-3133102023301130-3110301311221032"></a>

## Direct properties — ignore_value / 012002322221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1013012133001020-0202032233012201-1333211003013230-3122211220112321-3200300133002310-3001311022011112-0002333232022112-2201103033320222"></a>

## Next pages — ignore_value / 012002322221 / 4

- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-2330303333211221-2202033113332220-1120033202332121-0320332101312012-2322102332223022-1322012331213210-0001230321211032-1111320222330013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122300311012010-0312322322001301-1122223302331230-3130130323323231-0011023300032003-2010133202010123-0333223313311221-2110133323233003"></a>

## more_option.response_cookies_to_add.samesite_lax — samesite_lax / 112331321021 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- more_option.response_cookies_to_add.samesite_lax

<a id="canonical-2221220202231312-1222332333213110-0000203132332230-1323003010302323-2001111023122201-0303203220303032-0302331001033020-1203131302312030"></a>

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
samesite_lax = {}
```

<a id="canonical-2102322112332301-1321101230133321-1220121213032011-1303300230312133-3220333210003311-3210213201121311-0212111012101102-1030320013232032"></a>

## Direct properties — samesite_lax / 112331321021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3312003321210012-2212102123223310-0221331332010033-2130301223322233-2022202231301030-3000113033020022-2301200321301212-1020032331200110"></a>

## Next pages — samesite_lax / 112331321021 / 4

- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1332120002013122-1133230013132220-3110131310111110-3221333313023120-1213021213132300-2020010202123201-1232130222012003-1213113011012213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233101022233312-0101302111213033-2012021323211032-0221010222132112-0010001002212033-2313322222033231-1033121332003312-0331022333213331"></a>

## more_option.response_cookies_to_add.samesite_none — samesite_none / 213113313131 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- more_option.response_cookies_to_add.samesite_none

<a id="canonical-0030223113031120-0323122102230123-2122303023023120-2103302110032103-0132201112313110-3110011321311023-1111312010213003-2102121032321120"></a>

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
samesite_none = {}
```

<a id="canonical-1330120330103321-2111110121112230-2102300030000122-3320213010022330-2213213201120022-2310301011300130-1323113122301111-3302022323303233"></a>

## Direct properties — samesite_none / 213113313131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1010013000320330-2333231112332313-3230022313121130-0232302122303010-0003302310012211-2023000132023110-2300220210213003-1010011322020202"></a>

## Next pages — samesite_none / 213113313131 / 4

- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3200311123303121-3223302333130102-2013211221321310-3331031100310202-1011011132210011-1311002320131212-2320111013112103-0311130332202121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310320010023211-0203201223112133-3032320001321233-0303133313120330-2203103121121133-3032301222331310-0301232111333131-2302233023112032"></a>

## more_option.response_cookies_to_add.samesite_strict — samesite_strict / 312101212120 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- more_option.response_cookies_to_add.samesite_strict

<a id="canonical-3031031131230122-0133032102123202-2133220322213101-3120131322101112-0022130223021200-2030001101100222-3111213133121203-2122100330310222"></a>

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
samesite_strict = {}
```

<a id="canonical-2101022023331211-2130010023223033-3123023122032212-2333113221222203-3223022232123332-3032300032031320-1112003311012313-0301232012200303"></a>

## Direct properties — samesite_strict / 312101212120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1000010012213031-0123221120231211-3123213221030220-2113123302230000-0232311132231120-1220030111110113-3030033120331332-0332301132121223"></a>

## Next pages — samesite_strict / 312101212120 / 4

- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1322103020211312-3032032200003212-0100210121000112-1223112101202100-2220323300011112-2200212031011232-1301320100211112-0222303132332301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0321221100200112-0000211331020011-2030202002201021-2010031311123121-3231312330202311-3111321110322103-2311212233001100-2321011030100012"></a>

## more_option.response_cookies_to_add.secret_value — secret_value / 030120300223 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- more_option.response_cookies_to_add.secret_value

<a id="canonical-2120210132333100-3201302311113121-0121212310132213-2222320233211302-3221113013201033-3202232012313031-1121212033113012-0203200320032121"></a>

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
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-0132300002020302-1031120032132321-0333021202011120-3232012332101030-2222200332232230-2200002320312323-0011131112122120-3303212030331232"></a>

## Direct properties — secret_value / 030120300223 / 3

- [blindfold_secret_info](resources--http_loadbalancer--reference--group-021.md#canonical-3021231033011023-1011212223030023-2030200113003131-3322130011212001-1111133222233230-3112130301022221-0112233203021013-3000022003112333): complete subsection reference.

- [clear_secret_info](resources--http_loadbalancer--reference--group-021.md#canonical-1123201200031101-2233032321030232-1300131011211233-1133230233030101-0232031323331322-0020221323211110-3023231133130300-2112201123133020): complete subsection reference.

<a id="canonical-3321133031023002-2020220020001203-2331223121232223-0030012100002102-0030332332131022-2102001201131300-1322020020213313-2322121221200030"></a>

## Next pages — secret_value / 030120300223 / 4

- [more_option.response_cookies_to_add.secret_value.blindfold_secret_info](resources--http_loadbalancer--reference--group-021.md#canonical-3021231033011023-1011212223030023-2030200113003131-3322130011212001-1111133222233230-3112130301022221-0112233203021013-3000022003112333)
- [more_option.response_cookies_to_add.secret_value.clear_secret_info](resources--http_loadbalancer--reference--group-021.md#canonical-1123201200031101-2233032321030232-1300131011211233-1133230233030101-0232031323331322-0020221323211110-3023231133130300-2112201123133020)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-3021231033011023-1011212223030023-2030200113003131-3322130011212001-1111133222233230-3112130301022221-0112233203021013-3000022003112333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122300302201101-1210322211032123-1102310022132212-2003120002332231-2133312012020321-0210130210320011-2010022222023133-3331321332120313"></a>

## more_option.response_cookies_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 303100331130 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Property reference](resources--http_loadbalancer--reference--group-001.md#canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233)
- [more_option](resources--http_loadbalancer--reference--group-021.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232)
- [more_option.response_cookies_to_add](resources--http_loadbalancer--reference--group-021.md#canonical-3230122303302011-1112303312333322-0101311101122013-2310111003121102-0102222303200001-3332202002002202-3032200210010322-1220300032103322)
- [more_option.response_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-021.md#canonical-1322103020211312-3032032200003212-0100210121000112-1223112101202100-2220323300011112-2200212031011232-1301320100211112-0222303132332301)
- more_option.response_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-3102233100020032-0311321100223313-0200333101111312-1113130223221332-1231132233103133-3300200312302123-3202310323123000-2333013232200201"></a>

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

<a id="canonical-1213301023323302-2313101010112322-1312030312310000-1130211012122003-2222123100322011-2232302322021201-0031322100200300-3301112031130202"></a>

## Direct properties — blindfold_secret_info / 303100331130 / 3

<a id="canonical-2012010011030123-1000000122000012-3001110202022333-3321022113032110-1033113100000010-1133013021220123-1230001331112223-1232322202033001"></a>

<a id="canonical-3023221320310332-1003310331211031-0031132000313233-2120023001222131-3113200110213102-1003033030132000-3231010233012320-0130231020102320"></a>

## decryption_provider property — blindfold_secret_info / 303100331130 / 4

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

<a id="canonical-0201203233232220-0120323232113000-0221002322020123-0013112303332231-3300021332332312-2110233303101202-0132123230032030-2230212121011211"></a>

<a id="canonical-2320133202213132-3000322223000110-0020123023133333-2220301010312223-0220101101030321-2101020010233302-3320332130030210-2220233013320032"></a>

## location property — blindfold_secret_info / 303100331130 / 5

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

<a id="canonical-2232212101231013-1231330012122003-3301330122100220-1102121023233022-2100131001230300-1011302320333312-0031312232200102-1311103101221321"></a>

<a id="canonical-1020321211213010-3312131131003122-3221211011022132-1101332312102330-1123031203312300-2221222231333101-2210301213320012-2310211120122211"></a>

## store_provider property — blindfold_secret_info / 303100331130 / 6

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

<a id="canonical-1221230231021230-1013233001223201-3201130301322323-2232012121100110-2002120323120022-2010020030313213-1111102000210212-0210011110003320"></a>

## Next pages — blindfold_secret_info / 303100331130 / 7

- [more_option.response_cookies_to_add.secret_value](resources--http_loadbalancer--reference--group-021.md#canonical-1322103020211312-3032032200003212-0100210121000112-1223112101202100-2220323300011112-2200212031011232-1301320100211112-0222303132332301)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)

<a id="canonical-1123201200031101-2233032321030232-1300131011211233-1133230233030101-0232031323331322-0020221323211110-3023231133130300-2112201123133020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
