---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-0201100210222302-1320321202001120-2223231021210210-3231122230232110-3223220030301001-2211022303201010-2033203231003330-3232320103121031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `l7_ddos_protection.ddos_policy_custom` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [l7_ddos_protection](data-sources--http_loadbalancer--reference--group-019.md#canonical-3033222321011321-1022100112201200-2011110213223303-0111130210011233-3303003112011322-3001020230211012-0110113133233100-1303211333203133)
- l7_ddos_protection.ddos_policy_custom

<a id="canonical-3102300311012023-1123322001333110-1110012110320323-1100133332030121-1220233213001113-2303203311301120-0030021330033313-2131330310302210"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-2222003131221023-2333310232230033-3100302111311200-0220220001231313-1121333030001021-1300002210300200-1021331330030333-1001123130222222"></a>

### Direct properties for `l7_ddos_protection.ddos_policy_custom`

<a id="canonical-3332013230310302-3312000103330132-0032130113011313-0103213000001112-2231310110021103-3121232111130331-1323213313322220-2332011023332010"></a>

#### `l7_ddos_protection.ddos_policy_custom.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2022123330131122-0333121110123213-2232203211231301-1231210220001231-3210202303001301-2012101020200310-3130230121311023-2113121311233300"></a>

<a id="canonical-2112111123301111-0002331001331123-3311103030232213-1333100003122333-1031011023330011-3303311010122202-3133002323132323-3100202313302200"></a>

#### `l7_ddos_protection.ddos_policy_custom.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0300312133332123-2233300313323112-3321231331210021-2222220223313113-0331113012232332-3331211302200111-1131213322222213-3331012000202300"></a>

<a id="canonical-2022100230020111-2113213222033220-0031213010333213-3022031113200131-0200111220011023-3132300223303111-0013202320010001-1032200211311133"></a>

#### `l7_ddos_protection.ddos_policy_custom.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0311010100033111-3130301232000303-3123210031221223-2200133211303302-1023320231231332-2003020031323103-0310102103012133-1030013001300302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `l7_ddos_protection.ddos_policy_none` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [l7_ddos_protection](data-sources--http_loadbalancer--reference--group-019.md#canonical-3033222321011321-1022100112201200-2011110213223303-0111130210011233-3303003112011322-3001020230211012-0110113133233100-1303211333203133)
- l7_ddos_protection.ddos_policy_none

<a id="canonical-1313221331302113-1323330230302232-0101200013012231-2021132121301123-0202122313311131-1102013220102233-1121203103103122-1211131020132020"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for ddos policy none.

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

<a id="canonical-2202000110133032-0220032020200212-1223231112002203-0201130110202112-2232221312311303-1312032332310122-3203310033223303-1103003102131201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `l7_ddos_protection.default_rps_threshold` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [l7_ddos_protection](data-sources--http_loadbalancer--reference--group-019.md#canonical-3033222321011321-1022100112201200-2011110213223303-0111130210011233-3303003112011322-3001020230211012-0110113133233100-1303211333203133)
- l7_ddos_protection.default_rps_threshold

<a id="canonical-2013302223320233-2210302113230310-1312132130130130-3012321231301232-2132200120133033-1323211210322031-3033332020112113-3003310120132323"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default rps threshold.

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

<a id="canonical-3211023313203213-1030213203333230-2310010132231333-2013231321100132-3201202111303100-3220311000313013-0221210332333132-1002302313012102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `l7_ddos_protection.mitigation_block` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [l7_ddos_protection](data-sources--http_loadbalancer--reference--group-019.md#canonical-3033222321011321-1022100112201200-2011110213223303-0111130210011233-3303003112011322-3001020230211012-0110113133233100-1303211333203133)
- l7_ddos_protection.mitigation_block

<a id="canonical-0303001330330301-1331303323002323-1022003030330322-1213231202012122-3130203220301012-0100333120202101-2312011020101112-3113222113210222"></a>

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

<a id="canonical-1031133210330210-0121031133213322-2003130133110221-3010213022121020-2230331320300312-2210230200312121-0333133221121311-3022102322112202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `l7_ddos_protection.mitigation_captcha_challenge` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [l7_ddos_protection](data-sources--http_loadbalancer--reference--group-019.md#canonical-3033222321011321-1022100112201200-2011110213223303-0111130210011233-3303003112011322-3001020230211012-0110113133233100-1303211333203133)
- l7_ddos_protection.mitigation_captcha_challenge

<a id="canonical-0322120003130210-1101230302102312-1002320223033101-2110112221322122-1313313131100032-3100020220222013-0310201330112122-3222213110111110"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3330220002013330-1030012300222233-1123311132332301-3201123020100002-3323200033232032-1332021302330233-0122230012200200-2210330113122200"></a>

### Direct properties for `l7_ddos_protection.mitigation_captcha_challenge`

<a id="canonical-2002320100021303-0102233100130112-3020121200122033-1232323122312222-2111330102100223-1022103020332200-1303003122203221-0213301020300333"></a>

#### `l7_ddos_protection.mitigation_captcha_challenge.cookie_expiry` property

Type: `"number"`. Computed.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1212000203011132-0123011030022312-1222301112300100-0030123032112200-2021202122302301-3310231223110231-1023123233331321-3233223323113020"></a>

<a id="canonical-3222203310132000-3232131123221122-2110020312103200-1301112100130033-2032110301300000-1230122133210310-2033301100232223-2201303022032313"></a>

#### `l7_ddos_protection.mitigation_captcha_challenge.custom_page` property

Type: `"string"`. Computed.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1131120333131103-3112111212112313-0222023003231031-3001211131031033-3313030101332323-0200222313320002-1101222122023332-3101023232110130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `l7_ddos_protection.mitigation_js_challenge` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [l7_ddos_protection](data-sources--http_loadbalancer--reference--group-019.md#canonical-3033222321011321-1022100112201200-2011110213223303-0111130210011233-3303003112011322-3001020230211012-0110113133233100-1303211333203133)
- l7_ddos_protection.mitigation_js_challenge

<a id="canonical-0031322103201301-0300023011020310-2202032200133223-1220233132213320-3222300011021220-2313230103330230-2000233113230120-0231322300220313"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1023301132033000-3233232223232032-0130010320211111-1132012100010310-2110033113013320-0200213113030313-3230201300332100-2323010023130112"></a>

### Direct properties for `l7_ddos_protection.mitigation_js_challenge`

<a id="canonical-1303202023100211-2203111210032201-0301003233030230-0322000310122002-2230232232201130-3132300213330100-1232023300210011-0022011000020312"></a>

#### `l7_ddos_protection.mitigation_js_challenge.cookie_expiry` property

Type: `"number"`. Computed.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3223021311212110-3133123300201212-1330301223322331-2222230310101030-2220011231000303-3320211011220121-0121330330330313-1013233211100311"></a>

<a id="canonical-0012203100013132-3002202211010011-0333010323020233-0000021122313030-3323101223311220-2001333033223332-2020113230211221-2132311302202131"></a>

#### `l7_ddos_protection.mitigation_js_challenge.custom_page` property

Type: `"string"`. Computed.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2131220113102100-3103002331021213-3212210130300323-3301103013131222-2022331022202102-2003220100301323-1322033230210001-2020211213221033"></a>

<a id="canonical-2203021201233132-3123100022132100-1012113230121130-3303100223010231-1122230332001220-2111331222330101-2232233010133111-1012313230220332"></a>

#### `l7_ddos_protection.mitigation_js_challenge.js_script_delay` property

Type: `"number"`. Computed.

Delay introduced by JavaScript, in milliseconds.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3021012323102120-3202203000032032-2133022123330310-0310112201022300-0323033113303211-2210322220030333-2103012033230122-3232312030230201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `least_active` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- least_active

<a id="canonical-0320323021313031-3212211223022031-3213230131132232-3310022021123333-3021230311220030-1312120111233112-3301232133131302-2303211223003103"></a>

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

<a id="canonical-3323002031133002-1001022201210013-2112022022231332-1120011002023112-2131123021322022-3320133021232120-0332023100123222-1301012211222113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `malware_protection_settings` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- malware_protection_settings

<a id="canonical-2310131001002223-0110011100012011-0122222211112013-2012303223103012-0223003310301012-3210222301130330-2331303223010332-2213231331220230"></a>

Type: `"single"`. Computed.

Malware Protection protects Web Apps and APIs, from malicious file uploads by scanning files in
real-time.

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

<a id="canonical-1122000220020322-2022021133331310-0230333020121022-0000001113111033-3200133130101102-3201002010132233-1330222233300110-0211111212320300"></a>

### Direct properties for `malware_protection_settings`

- [malware_protection_rules](data-sources--http_loadbalancer--reference--group-020.md#canonical-1022201020312233-3201103310230203-0233212230022213-1123110002302233-3310310010110230-2123030101301210-0203311113132213-3022213311033301): complete subsection reference.

<a id="canonical-1022201020312233-3201103310230203-0233212230022213-1123110002302233-3310310010110230-2123030101301210-0203311113132213-3022213311033301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `malware_protection_settings.malware_protection_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [malware_protection_settings](data-sources--http_loadbalancer--reference--group-020.md#canonical-3323002031133002-1001022201210013-2112022022231332-1120011002023112-2131123021322022-3320133021232120-0332023100123222-1301012211222113)
- malware_protection_settings.malware_protection_rules

<a id="canonical-2330010132310031-0101233201131330-1300310210020023-1223320220303121-0001133131130212-1310133311322000-1302233213313220-1011210031330321"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2033201021131213-3303200212300122-0102210210303132-3321203313300310-1312211213323323-3213323020223231-2311333233232003-0101022120000111"></a>

### Direct properties for `malware_protection_settings.malware_protection_rules`

- [action](data-sources--http_loadbalancer--reference--group-020.md#canonical-3232221301230133-1233322012120131-1310201111231120-3122133303310233-0323023010322200-3300210131113330-2020132222110023-3112120112230332): complete subsection reference.

- [domain](data-sources--http_loadbalancer--reference--group-020.md#canonical-0302302312030220-0110010203123002-2220113101212111-0213000002030112-1032301130023001-1300300232120232-0200032001333131-2033202223000323): complete subsection reference.

<a id="canonical-3311213002300102-2310121121223011-0032001020213213-1210221332232232-2133200110103023-2001211131001122-0321021322000123-0210131231013220"></a>

<a id="canonical-2303233203322223-1002100210330032-1220110203320030-1300211030122331-3202223233333313-0022112321021333-3121002212021031-2013110122201202"></a>

#### `malware_protection_settings.malware_protection_rules.http_methods` property

Type: `["list", "string"]`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] HTTP Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [metadata](data-sources--http_loadbalancer--reference--group-020.md#canonical-3131322320120101-3232120210310200-3020023303200330-1332013100211013-3130013223121330-0011230130300202-3322130330320320-3120332102212233): complete subsection reference.

- [path](data-sources--http_loadbalancer--reference--group-020.md#canonical-3030233220103213-2202213103102311-3010120103022100-2033300010121312-2231033100012323-1332130013123211-0023010301330223-1302032103313022): complete subsection reference.

<a id="canonical-3232221301230133-1233322012120131-1310201111231120-3122133303310233-0323023010322200-3300210131113330-2020132222110023-3112120112230332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `malware_protection_settings.malware_protection_rules.action` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [malware_protection_settings](data-sources--http_loadbalancer--reference--group-020.md#canonical-3323002031133002-1001022201210013-2112022022231332-1120011002023112-2131123021322022-3320133021232120-0332023100123222-1301012211222113)
- [malware_protection_settings.malware_protection_rules](data-sources--http_loadbalancer--reference--group-020.md#canonical-1022201020312233-3201103310230203-0233212230022213-1123110002302233-3310310010110230-2123030101301210-0203311113132213-3022213311033301)
- malware_protection_settings.malware_protection_rules.action

<a id="canonical-1232133310201113-1123210311203301-3332210321102220-1110200231131032-1111013220111332-1300231300302023-1301011200003212-3201222122301031"></a>

Type: `"single"`. Computed.

Action

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

<a id="canonical-3031130100010211-2302123032321302-3330212322032331-3032010132230103-0132332322101132-3221132312132222-2330003102013223-1031313301322001"></a>

### Direct properties for `malware_protection_settings.malware_protection_rules.action`

- [block](data-sources--http_loadbalancer--reference--group-020.md#canonical-3203231133033130-2321020212300102-1303121033130023-1103112033033023-3012001223232131-1110320202003222-2212223213321113-2310031220223120): complete subsection reference.

- [report](data-sources--http_loadbalancer--reference--group-020.md#canonical-0121233223310001-0101301200330001-2303131303213331-3311212333231002-3323120003112111-1023111122211312-0120320100022010-2231102102001213): complete subsection reference.

<a id="canonical-3203231133033130-2321020212300102-1303121033130023-1103112033033023-3012001223232131-1110320202003222-2212223213321113-2310031220223120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `malware_protection_settings.malware_protection_rules.action.block` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [malware_protection_settings](data-sources--http_loadbalancer--reference--group-020.md#canonical-3323002031133002-1001022201210013-2112022022231332-1120011002023112-2131123021322022-3320133021232120-0332023100123222-1301012211222113)
- [malware_protection_settings.malware_protection_rules](data-sources--http_loadbalancer--reference--group-020.md#canonical-1022201020312233-3201103310230203-0233212230022213-1123110002302233-3310310010110230-2123030101301210-0203311113132213-3022213311033301)
- [malware_protection_settings.malware_protection_rules.action](data-sources--http_loadbalancer--reference--group-020.md#canonical-3232221301230133-1233322012120131-1310201111231120-3122133303310233-0323023010322200-3300210131113330-2020132222110023-3112120112230332)
- malware_protection_settings.malware_protection_rules.action.block

<a id="canonical-3232031300322220-0120311112303003-0103113330121313-3212032332322203-2103003011133113-0131002311130200-3320220233010122-3212322023310330"></a>

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

<a id="canonical-0121233223310001-0101301200330001-2303131303213331-3311212333231002-3323120003112111-1023111122211312-0120320100022010-2231102102001213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `malware_protection_settings.malware_protection_rules.action.report` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [malware_protection_settings](data-sources--http_loadbalancer--reference--group-020.md#canonical-3323002031133002-1001022201210013-2112022022231332-1120011002023112-2131123021322022-3320133021232120-0332023100123222-1301012211222113)
- [malware_protection_settings.malware_protection_rules](data-sources--http_loadbalancer--reference--group-020.md#canonical-1022201020312233-3201103310230203-0233212230022213-1123110002302233-3310310010110230-2123030101301210-0203311113132213-3022213311033301)
- [malware_protection_settings.malware_protection_rules.action](data-sources--http_loadbalancer--reference--group-020.md#canonical-3232221301230133-1233322012120131-1310201111231120-3122133303310233-0323023010322200-3300210131113330-2020132222110023-3112120112230332)
- malware_protection_settings.malware_protection_rules.action.report

<a id="canonical-0022100302301211-2333213121100331-0333102200310223-0230333122301223-2301013322131313-3110313111221333-3122003323321232-2202313220301322"></a>

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

<a id="canonical-0302302312030220-0110010203123002-2220113101212111-0213000002030112-1032301130023001-1300300232120232-0200032001333131-2033202223000323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `malware_protection_settings.malware_protection_rules.domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [malware_protection_settings](data-sources--http_loadbalancer--reference--group-020.md#canonical-3323002031133002-1001022201210013-2112022022231332-1120011002023112-2131123021322022-3320133021232120-0332023100123222-1301012211222113)
- [malware_protection_settings.malware_protection_rules](data-sources--http_loadbalancer--reference--group-020.md#canonical-1022201020312233-3201103310230203-0233212230022213-1123110002302233-3310310010110230-2123030101301210-0203311113132213-3022213311033301)
- malware_protection_settings.malware_protection_rules.domain

<a id="canonical-2011213221022020-0211310322321023-1032203101320012-2320032021020331-3111122210210133-0030122112120233-0321302003103222-0021133123130322"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Additional upstream details:

Domain to be matched.

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

<a id="canonical-0001020102021021-3111233110101221-2033001130221303-3111303103023120-2302232011221223-1012112122223022-3013223232222102-3210111313313023"></a>

### Direct properties for `malware_protection_settings.malware_protection_rules.domain`

- [any_domain](data-sources--http_loadbalancer--reference--group-020.md#canonical-1100111330321030-0313000102323232-0100201221110121-1023002132212333-3010213103231330-1210013001033131-1311000311020321-0012113110313121): complete subsection reference.

- [domain](data-sources--http_loadbalancer--reference--group-020.md#canonical-1013030110003130-3020022312231232-0022222000012011-1200203032302221-1112123303132312-3021103132330230-1131102201311111-3022103213113122): complete subsection reference.

<a id="canonical-1100111330321030-0313000102323232-0100201221110121-1023002132212333-3010213103231330-1210013001033131-1311000311020321-0012113110313121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `malware_protection_settings.malware_protection_rules.domain.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [malware_protection_settings](data-sources--http_loadbalancer--reference--group-020.md#canonical-3323002031133002-1001022201210013-2112022022231332-1120011002023112-2131123021322022-3320133021232120-0332023100123222-1301012211222113)
- [malware_protection_settings.malware_protection_rules](data-sources--http_loadbalancer--reference--group-020.md#canonical-1022201020312233-3201103310230203-0233212230022213-1123110002302233-3310310010110230-2123030101301210-0203311113132213-3022213311033301)
- [malware_protection_settings.malware_protection_rules.domain](data-sources--http_loadbalancer--reference--group-020.md#canonical-0302302312030220-0110010203123002-2220113101212111-0213000002030112-1032301130023001-1300300232120232-0200032001333131-2033202223000323)
- malware_protection_settings.malware_protection_rules.domain.any_domain

<a id="canonical-1311130310201011-0323010033021133-2302000330210133-3301133332332012-1032130333230331-3131003230112113-3133331102203112-2111110103222210"></a>

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

<a id="canonical-1013030110003130-3020022312231232-0022222000012011-1200203032302221-1112123303132312-3021103132330230-1131102201311111-3022103213113122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `malware_protection_settings.malware_protection_rules.domain.domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [malware_protection_settings](data-sources--http_loadbalancer--reference--group-020.md#canonical-3323002031133002-1001022201210013-2112022022231332-1120011002023112-2131123021322022-3320133021232120-0332023100123222-1301012211222113)
- [malware_protection_settings.malware_protection_rules](data-sources--http_loadbalancer--reference--group-020.md#canonical-1022201020312233-3201103310230203-0233212230022213-1123110002302233-3310310010110230-2123030101301210-0203311113132213-3022213311033301)
- [malware_protection_settings.malware_protection_rules.domain](data-sources--http_loadbalancer--reference--group-020.md#canonical-0302302312030220-0110010203123002-2220113101212111-0213000002030112-1032301130023001-1300300232120232-0200032001333131-2033202223000323)
- malware_protection_settings.malware_protection_rules.domain.domain

<a id="canonical-2333123111221200-2321002203312121-0303121221100120-1332331311121303-3111010130103023-3223222332010133-3310232231100201-0301210202310321"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Additional upstream details:

Domains names.

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

<a id="canonical-0103220012002320-0033010300021123-0321023201322101-0200301032330113-3111030011203033-3113101103221212-3213300011313200-3202302221002203"></a>

### Direct properties for `malware_protection_settings.malware_protection_rules.domain.domain`

<a id="canonical-2003330130002310-2031320221210013-1222233130020232-3301221012311122-0032120303032002-0112312323333121-1311100100031113-2211030230100101"></a>

#### `malware_protection_settings.malware_protection_rules.domain.domain.exact_value` property

Type: `"string"`. Computed.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3220312311021330-2310123233320003-0310032013002030-3131100012200112-0121010112310101-3003101331101030-3002202021000332-2302320332002223"></a>

<a id="canonical-0011210231201032-0013001331120301-2303101202100330-3202133301311023-2321030322112313-3103321033232111-3001002132323213-0201230332003302"></a>

#### `malware_protection_settings.malware_protection_rules.domain.domain.regex_value` property

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0102012222333030-2233222233011202-3131323313310203-2002210333121031-2103231203200131-0203202311012013-0111321300001030-3302210002100111"></a>

<a id="canonical-0000013301311130-1232211222002011-0012111213220012-2211220300122020-1311010210213221-0011123303232310-1201233201022203-0301112322022221"></a>

#### `malware_protection_settings.malware_protection_rules.domain.domain.suffix_value` property

Type: `"string"`. Computed.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3131322320120101-3232120210310200-3020023303200330-1332013100211013-3130013223121330-0011230130300202-3322130330320320-3120332102212233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `malware_protection_settings.malware_protection_rules.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [malware_protection_settings](data-sources--http_loadbalancer--reference--group-020.md#canonical-3323002031133002-1001022201210013-2112022022231332-1120011002023112-2131123021322022-3320133021232120-0332023100123222-1301012211222113)
- [malware_protection_settings.malware_protection_rules](data-sources--http_loadbalancer--reference--group-020.md#canonical-1022201020312233-3201103310230203-0233212230022213-1123110002302233-3310310010110230-2123030101301210-0203311113132213-3022213311033301)
- malware_protection_settings.malware_protection_rules.metadata

<a id="canonical-1222030113201122-3310112303301320-2311213303001302-3312121220212113-3013033011311110-0001000300012033-0200323201230001-2003213221213123"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-2331303221311333-0120122222103310-3033021030232130-1232000131023322-0222130303011102-0013203311133101-3120113133123212-3110320113220022"></a>

### Direct properties for `malware_protection_settings.malware_protection_rules.metadata`

<a id="canonical-3103212110223333-3222102110122230-0330310230010003-3113330222312002-2223010313001132-3212332130221323-1311222032311230-1223133221111023"></a>

#### `malware_protection_settings.malware_protection_rules.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-0302320013011020-2010203033202000-2301113320011000-1021230200123210-2100123101202123-2313013121320112-3320330301201210-3320330313301203"></a>

<a id="canonical-3222331113000201-2332310002303132-0212303301320221-0023100131223121-1121020213120331-2231100012210030-1111132030230123-1220020323030121"></a>

#### `malware_protection_settings.malware_protection_rules.metadata.name` property

Type: `"string"`. Computed.

This is the name of the message. The value of name has to follow DNS-1035 format.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3030233220103213-2202213103102311-3010120103022100-2033300010121312-2231033100012323-1332130013123211-0023010301330223-1302032103313022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `malware_protection_settings.malware_protection_rules.path` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [malware_protection_settings](data-sources--http_loadbalancer--reference--group-020.md#canonical-3323002031133002-1001022201210013-2112022022231332-1120011002023112-2131123021322022-3320133021232120-0332023100123222-1301012211222113)
- [malware_protection_settings.malware_protection_rules](data-sources--http_loadbalancer--reference--group-020.md#canonical-1022201020312233-3201103310230203-0233212230022213-1123110002302233-3310310010110230-2123030101301210-0203311113132213-3022213311033301)
- malware_protection_settings.malware_protection_rules.path

<a id="canonical-2323233320232103-3002201233013323-2122122303212221-1123031032232132-0312200103302331-1110002131110031-0102322022022022-1323010033113230"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

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

<a id="canonical-1301211132330303-0311101123100213-2120131220223121-1203221202121331-0022023132332101-3022201011103031-3032032333100001-1321111210320230"></a>

### Direct properties for `malware_protection_settings.malware_protection_rules.path`

<a id="canonical-1110121003222301-2001223012120231-2312303003110131-2320223322002232-0033312312010123-2322303011111012-1000311302210110-2333311102031212"></a>

#### `malware_protection_settings.malware_protection_rules.path.path` property

Type: `"string"`. Computed.

Exclusive with \[prefix regular expression\] Exact path value to match.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2321000033202312-2032202031113211-0210203102111231-0323003021131222-1002210110121320-2211020002021203-2231023200332112-0113110033123003"></a>

<a id="canonical-3023131331133002-0001000001222012-3223203222120313-0331132301220012-2132203022333020-0233022002012313-0101220123123303-2222311233331110"></a>

#### `malware_protection_settings.malware_protection_rules.path.prefix` property

Type: `"string"`. Computed.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1302301213211332-3223222003011023-3122331202122021-2300211200012011-3330121323232023-3121103032310123-0023223131133310-2312102301001230"></a>

<a id="canonical-1103121130210032-1021113012330210-1133201022333332-1112332313132300-3121210220131310-0232013131212122-3022232302210203-3220303100221021"></a>

#### `malware_protection_settings.malware_protection_rules.path.regex` property

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- more_option

<a id="canonical-0322010010132332-2313220303020322-1130211011303130-3331221333130000-0202230021031103-0331020223012321-1220320220023203-0102313001330302"></a>

Type: `"single"`. Computed.

This defines various OPTIONS to define a route.

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

<a id="canonical-2301210000231212-2120133212102122-2330320200320103-1310022011211310-2212132322131323-3103021123132012-2203013121331000-2310320300001232"></a>

### Direct properties for `more_option`

- [buffer_policy](data-sources--http_loadbalancer--reference--group-020.md#canonical-1301020110201313-1231312103022303-3222012010010123-1202223101100321-3332310101312323-0232113113113013-2322230233132212-2132300021120211): complete subsection reference.

- [compression_params](data-sources--http_loadbalancer--reference--group-020.md#canonical-2302123223010202-3133131332011030-0130020210301221-1333222231121000-2323300012213030-2223330323031320-2300211101031301-2233331322213123): complete subsection reference.

<a id="canonical-3003103003203202-3222132303121311-3212211332111213-0232200023121332-0203211213300213-0120221003111103-2102223213003000-1020100011221032"></a>

<a id="canonical-1021310112232313-3133322223131300-2211223032021203-2012310220221233-2313013030032110-1200113303212320-1033110130113223-1330321002332332"></a>

#### `more_option.custom_errors` property

Type: `["map", "string"]`. Computed.

Map of integer error codes as keys and string values that can be used to provide custom HTTP pages
for each error code. Key of the map can be either response code class or HTTP Error code. Response
code classes for key is configured as follows 3 -- for 3xx response code class 4 -- for 4xx response
code class 5 -- for 5xx response code class Value of the map is string which represents custom HTTP
responses. Specific response code takes preference when both response code and response code class
matches for a request.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 16
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "ranges": [
        [
          3,
          3
        ],
        [
          4,
          4
        ],
        [
          5,
          5
        ],
        [
          300,
          599
        ]
      ],
      "type": "uint32-string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.uint32.ranges": "3,4,5,300-599",
      "ves.io.schema.rules.map.max_pairs": "16",
      "ves.io.schema.rules.map.values.string.max_len": "65536",
      "ves.io.schema.rules.map.values.string.uri_ref": "true"
    },
    "values": {
      "format": "uri-reference",
      "maxLength": 65536,
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

<a id="canonical-3100220221300131-3333201302313230-1023110021031321-3202202003200120-0033103220211232-0333232232003020-1133030013000313-1303310132310021"></a>

<a id="canonical-2011210203300210-0032002001002222-2133330212130323-1331112311322102-3212322303123031-1022101103230120-3110211221201232-0303112213100033"></a>

#### `more_option.disable_default_error_pages` property

Type: `"bool"`. Computed.

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

- [disable_path_normalize](data-sources--http_loadbalancer--reference--group-020.md#canonical-1000101022331013-2022233100203213-0200120223103201-2013122230231330-2312032010322001-2002323312310133-2233203203030012-3222213232102310): complete subsection reference.

- [enable_path_normalize](data-sources--http_loadbalancer--reference--group-020.md#canonical-2111133200111332-0320132002122001-0200103333213002-2130232322120021-0220010200230002-2120333133200030-1223021111230102-0330111320203103): complete subsection reference.

<a id="canonical-0210220310323310-0032311203120212-1333101033111003-3300331200201032-3123023120133033-3313203213313111-1013000213222122-0323332213000332"></a>

<a id="canonical-2132303010132311-0003233210300333-2011132121030012-2330313221212032-3133232030022100-2020132103232202-1030201111022302-3212213233120330"></a>

#### `more_option.idle_timeout` property

Type: `"number"`. Computed.

The amount of time that a stream can exist without upstream or downstream activity, in milliseconds.
The stream is terminated with an HTTP 504 (Gateway Timeout) error code if no upstream response
header has been received, otherwise the stream is reset.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2220201023201013-0010133303311120-0230101333203211-2232311321003223-0301330000032200-1120101200331110-0201313102201222-0033122211332103"></a>

<a id="canonical-2120223111003311-1022213333221212-1113022031002012-0023133023003020-1323121131100031-3030111231231103-1100001020223130-2333132103330120"></a>

#### `more_option.max_request_header_size` property

Type: `"number"`. Computed.

The maximum request header size for downstream connections, in KiB. An HTTP 431 (Request Header
Fields Too Large) error code is sent for requests that exceed this size.

If multiple load balancers share the same advertise\_policy, the highest value configured across all
such load balancers is used for all the load balancers in question.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2320100201223223-2112022231231031-0322301023321333-2033203332122030-3031032223100110-2030131012220320-0112221313010330-0123013010223313"></a>

<a id="canonical-0211110130130221-0210002300201012-0202033311112100-1202320111131220-3311221200103121-3130001022113333-0202300003000232-2031213121232301"></a>

#### `more_option.max_requests_per_connection` property

Type: `"number"`. Computed.

Exclusive with \[no\_request\_limit\_per\_connection\] Sets the maximum number of requests a
downstream client can send over a single connection to Envoy. Enter a value &gt;=1 to define the
request limit per connection.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [no_request_limit_per_connection](data-sources--http_loadbalancer--reference--group-020.md#canonical-0032302113322132-3201320132031121-2231322111001201-1031122030221133-2010001303320212-0301101322120010-2311332220230022-2101002112303133): complete subsection reference.

- [request_cookies_to_add](data-sources--http_loadbalancer--reference--group-020.md#canonical-1102320100210023-1210033001301330-1020000111230213-0103120220001333-0032231333013111-2231032020310030-2132123001321222-0233303131323331): complete subsection reference.

<a id="canonical-0200132312102233-2313303200101232-1130321301013012-0323133312133230-2000300003012120-0330031121103333-1200003223131230-1300010223323031"></a>

<a id="canonical-2110012003320123-3122223003231310-0132121121331320-1223132303323032-2000000101230303-2301033231010100-2200321321100220-2330233123103330"></a>

#### `more_option.request_cookies_to_remove` property

Type: `["list", "string"]`. Computed.

List of keys of Cookies to be removed from the HTTP request being sent towards upstream.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [request_headers_to_add](data-sources--http_loadbalancer--reference--group-020.md#canonical-3131323333013211-1120301133201010-1211330033032130-3121231320313312-1312311200203221-3112223310131110-3201331320102000-1013203220202110): complete subsection reference.

<a id="canonical-3301211301320200-2110123000211211-3131300213310322-3121332322213120-0201010330302010-0201223110233112-1220301232013102-2310310332122022"></a>

<a id="canonical-3212333221033001-0032300220233223-2100110332323110-2023000312313130-0333330223330130-2120230311200330-2110132013213021-1332202100330013"></a>

#### `more_option.request_headers_to_remove` property

Type: `["list", "string"]`. Computed.

List of keys of Headers to be removed from the HTTP request being sent towards upstream.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [response_cookies_to_add](data-sources--http_loadbalancer--reference--group-020.md#canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132): complete subsection reference.

<a id="canonical-0311312202320231-0230101101301112-1023033101123003-3103010310101302-1033000122331301-1222032212111311-1101123013013221-0032110112311131"></a>

<a id="canonical-1012133222200332-3021223232310123-0320121330200112-1231002302031032-3111003220032110-1112001201223120-3211332200121101-3200023110123230"></a>

#### `more_option.response_cookies_to_remove` property

Type: `["list", "string"]`. Computed.

List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire
set-cookie header will be removed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [response_headers_to_add](data-sources--http_loadbalancer--reference--group-021.md#canonical-1002330210300303-2230220311120303-1212313221313221-2322232302001122-2301211303310033-3003133330013332-2202232123020003-3010130310012121): complete subsection reference.

<a id="canonical-0332103121101133-0232230013001230-2301031011231313-3321333122113002-0230303201011211-1023000030013320-3331323213333011-2313122110322101"></a>

<a id="canonical-2033331221231221-1223101132221311-0003003313110231-2000132012000221-1022031112220321-3301331021011201-1323313302311311-1110021023102333"></a>

#### `more_option.response_headers_to_remove` property

Type: `["list", "string"]`. Computed.

List of keys of Headers to be removed from the HTTP response being sent towards downstream.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1301020110201313-1231312103022303-3222012010010123-1202223101100321-3332310101312323-0232113113113013-2322230233132212-2132300021120211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.buffer_policy` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- more_option.buffer_policy

<a id="canonical-1122012233223120-0332303022232120-0333231312201303-3303211013130102-0133301022211122-0100133203200131-2121100131301032-1200212312020112"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1331231300222223-1301203223331123-0301230002230200-0121122030231333-2111212120300001-0333111323121100-0201211133212110-1022121032131131"></a>

### Direct properties for `more_option.buffer_policy`

<a id="canonical-2230003323332232-1121230322010222-1131311112233113-3202021132030331-1333132210331331-1211013302321233-1212111002023021-2020200020023212"></a>

#### `more_option.buffer_policy.disabled` property

Type: `"bool"`. Computed.

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

<a id="canonical-0130010102101022-3122310312001121-2222213302322013-2323231123011121-2122223122322321-3002301313302023-3133030223122233-3030010320211302"></a>

<a id="canonical-0130223321121302-1102231132310210-3332012332200122-3302231312113002-3030321120133313-0300111332122202-3103133323031213-1212112000033333"></a>

#### `more_option.buffer_policy.max_request_bytes` property

Type: `"number"`. Computed.

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2302123223010202-3133131332011030-0130020210301221-1333222231121000-2323300012213030-2223330323031320-2300211101031301-2233331322213123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.compression_params` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- more_option.compression_params

<a id="canonical-0010333100010002-3212312213203031-2020212202132122-1232023323113301-0112001210201221-0222303120001233-2201013110310112-0112100100103112"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3131021203211102-1221013031023022-3230111332113100-0131011300123320-1300121300012113-0323321103101113-0312312123221010-1000023131201011"></a>

### Direct properties for `more_option.compression_params`

<a id="canonical-1032011010120133-2111232130100310-3323301333100121-1202010101233202-0031333310111310-0033120310120000-1330122101123032-0312122203301120"></a>

#### `more_option.compression_params.content_length` property

Type: `"number"`. Computed.

Minimum response length, in bytes, which will trigger compression. The. Defaults to \`30\`.

Additional upstream details:

The default value is 30.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0100322223202131-2201100111101331-3331331201321231-1220113331122230-2320202001010131-0320133311102202-2233112311330032-1313330300110110"></a>

<a id="canonical-2131222331021112-1313133010003302-1310101210333133-3311321010122202-1212131312102121-1312102231332011-2102110213112202-0320023311223323"></a>

#### `more_option.compression_params.content_type` property

Type: `["list", "string"]`. Computed.

Set of strings that allows specifying which mime-types yield compression When this field is not
defined, compression will be applied to the following mime-types: 'application/JavaScript'
'application/JSON', 'application/xhtml+XML' 'image/svg+XML' 'text/CSS' 'text/HTML' 'text/plain'
'text/XML'.

Additional upstream details:

Set of strings that allows specifying which mime-types yield compression When this field is not
defined, compression will be applied to the following mime-types: "application/JavaScript"
"application/JSON", "application/xhtml+XML" "image/svg+XML" "text/CSS" "text/HTML" "text/plain"
"text/XML"

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3111022210302221-0130233103110222-3310201032012100-0131103131212122-2223033021030212-3011220211322122-0030301100132303-0213230120303303"></a>

<a id="canonical-3030201200310013-0210002030321123-1332033121310301-3000322130203001-3212202020213130-0201022322102020-3100313213322200-3313103112023032"></a>

#### `more_option.compression_params.disable_on_etag_header` property

Type: `"bool"`. Computed.

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

<a id="canonical-1130203200030210-1200030120031201-2120121222332100-0030011210030032-0112103313120013-0103120313120023-3301333013230012-1321102000311002"></a>

<a id="canonical-0022002120313130-3311012233312112-2231033312333203-0312112220123232-3021313121121112-0001120012231031-0311122133210020-3222022130123103"></a>

#### `more_option.compression_params.remove_accept_encoding_header` property

Type: `"bool"`. Computed.

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

<a id="canonical-1000101022331013-2022233100203213-0200120223103201-2013122230231330-2312032010322001-2002323312310133-2233203203030012-3222213232102310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.disable_path_normalize` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- more_option.disable_path_normalize

<a id="canonical-2000301103010330-2321011030200210-0301012321231121-2102300321332331-2223323120323223-3320132123312310-1001020300210132-0203202102211130"></a>

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

<a id="canonical-2111133200111332-0320132002122001-0200103333213002-2130232322120021-0220010200230002-2120333133200030-1223021111230102-0330111320203103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.enable_path_normalize` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- more_option.enable_path_normalize

<a id="canonical-0130203330033313-0023310333200001-2103312233131322-3021213232203130-2101301221112000-0200200131012220-0302003210023012-2221311320110333"></a>

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

<a id="canonical-0032302113322132-3201320132031121-2231322111001201-1031122030221133-2010001303320212-0301101322120010-2311332220230022-2101002112303133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.no_request_limit_per_connection` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- more_option.no_request_limit_per_connection

<a id="canonical-0133233333210110-0332312101013111-3012010123330132-0032012332131213-0100110130322121-3303021301300331-1130123023023310-1211333300210102"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no request limit per connection.

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

<a id="canonical-1102320100210023-1210033001301330-1020000111230213-0103120220001333-0032231333013111-2231032020310030-2132123001321222-0233303131323331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.request_cookies_to_add` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- more_option.request_cookies_to_add

<a id="canonical-1002110233021011-1032012130223212-2101210303011031-3202330100000110-1333233113322202-3231211003102221-1001133232213210-2221000230220221"></a>

Type: `"list"`. Computed.

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0023002032312310-2210203002120202-0221320032103110-2001003300112233-2332012223203311-1231022100323213-2322302311210311-3102130210103132"></a>

### Direct properties for `more_option.request_cookies_to_add`

<a id="canonical-0033002313122022-0121023311303323-0333331032233333-2233111313330210-2211132313203112-2203330332220032-2133022231300200-3000020301001032"></a>

#### `more_option.request_cookies_to_add.name` property

Type: `"string"`. Computed.

Name. Name of the cookie in Cookie header.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3022021011221000-0100210211332031-2033233010021031-1220202322330212-3100232232111332-2133122213320021-1310012312210222-1232110303231321"></a>

<a id="canonical-1311032120223221-3131021201303320-2311203102033221-1010000133100133-1133310030100100-2102220210121213-3322001232201301-1001002011000132"></a>

#### `more_option.request_cookies_to_add.overwrite` property

Type: `"bool"`. Computed.

Should the value be overwritten? If true, the value is overwritten to existing values. not
overwrite. Defaults to \`do\`.

Additional upstream details:

If true, the value is overwritten to existing values. Default value is do not overwrite.

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

- [secret_value](data-sources--http_loadbalancer--reference--group-020.md#canonical-0111002202312212-2132132020211101-0112120020122220-2103133021213313-2120310210030320-3020210011023232-2300313003000332-1322232001013302): complete subsection reference.

<a id="canonical-0211010223332200-0321001022000112-2202232022323232-2213322101302010-2002112133302222-3333301201203230-0202332311123323-1121110311313321"></a>

<a id="canonical-2120000233211100-2320302012100031-3121132203133302-1113002012320221-3112111011023002-1133000123123030-0120110311320112-0003223122233131"></a>

#### `more_option.request_cookies_to_add.value` property

Type: `"string"`. Computed.

Exclusive with \[secret\_value\] Value of the Cookie header.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0111002202312212-2132132020211101-0112120020122220-2103133021213313-2120310210030320-3020210011023232-2300313003000332-1322232001013302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.request_cookies_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.request_cookies_to_add](data-sources--http_loadbalancer--reference--group-020.md#canonical-1102320100210023-1210033001301330-1020000111230213-0103120220001333-0032231333013111-2231032020310030-2132123001321222-0233303131323331)
- more_option.request_cookies_to_add.secret_value

<a id="canonical-3103313320013123-1302212011022320-2001302133222101-0023131213300203-2332113213120120-1112212223201110-1211332132123013-0201120203113002"></a>

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

<a id="canonical-2121300133222330-2101211002231220-0013001001023333-0230230311031233-1211200001233102-1122130210120012-0210311322230121-3211233012132100"></a>

### Direct properties for `more_option.request_cookies_to_add.secret_value`

- [blindfold_secret_info](data-sources--http_loadbalancer--reference--group-020.md#canonical-2201331032111322-3333310102332102-3101101023023112-3231123100131311-2121301132110123-0030123323133300-0231230032203332-1011200113133222): complete subsection reference.

- [clear_secret_info](data-sources--http_loadbalancer--reference--group-020.md#canonical-3332230303110210-3021211311020310-1031000031303222-3332300122002002-0020112222330203-3113013231211201-3211120123032013-1032110331020312): complete subsection reference.

<a id="canonical-2201331032111322-3333310102332102-3101101023023112-3231123100131311-2121301132110123-0030123323133300-0231230032203332-1011200113133222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.request_cookies_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.request_cookies_to_add](data-sources--http_loadbalancer--reference--group-020.md#canonical-1102320100210023-1210033001301330-1020000111230213-0103120220001333-0032231333013111-2231032020310030-2132123001321222-0233303131323331)
- [more_option.request_cookies_to_add.secret_value](data-sources--http_loadbalancer--reference--group-020.md#canonical-0111002202312212-2132132020211101-0112120020122220-2103133021213313-2120310210030320-3020210011023232-2300313003000332-1322232001013302)
- more_option.request_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-3230021112233220-2322111101233233-3231021210320320-2002103200222330-2103323200132121-0210030122022231-3031202230312002-1200102332032021"></a>

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

<a id="canonical-0320231031213311-0230101210223020-0222120031020110-2113302303320301-0333311213201302-1022122222231311-1301111202201121-0023302201313312"></a>

### Direct properties for `more_option.request_cookies_to_add.secret_value.blindfold_secret_info`

<a id="canonical-1013203301213310-2322022220000033-2103311033211200-1330232031302112-0232031121213022-2112131212210012-3113032222220122-2112212131032331"></a>

#### `more_option.request_cookies_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1002330022030020-2030311320030332-0020020032302002-3331111300020022-0023210330120300-1301120203221323-0212201310011313-1002310032002323"></a>

<a id="canonical-1331003023302130-2031113131112200-2123323212310231-1223010110121332-3121031321333231-1300330011103110-2110231023011112-0310223002032130"></a>

#### `more_option.request_cookies_to_add.secret_value.blindfold_secret_info.location` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1130111303031303-0323312102330223-0003201131230212-0313303022221120-2331130332223202-2123101101321112-0101320302032201-3223101033333330"></a>

<a id="canonical-0222130033110212-1202322301102131-3202012331013021-0310112130131013-2321303010203003-0130020312101111-1132132112002022-0221010112103200"></a>

#### `more_option.request_cookies_to_add.secret_value.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3332230303110210-3021211311020310-1031000031303222-3332300122002002-0020112222330203-3113013231211201-3211120123032013-1032110331020312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.request_cookies_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.request_cookies_to_add](data-sources--http_loadbalancer--reference--group-020.md#canonical-1102320100210023-1210033001301330-1020000111230213-0103120220001333-0032231333013111-2231032020310030-2132123001321222-0233303131323331)
- [more_option.request_cookies_to_add.secret_value](data-sources--http_loadbalancer--reference--group-020.md#canonical-0111002202312212-2132132020211101-0112120020122220-2103133021213313-2120310210030320-3020210011023232-2300313003000332-1322232001013302)
- more_option.request_cookies_to_add.secret_value.clear_secret_info

<a id="canonical-0033021013231130-3302231110112000-2121020222332010-1112320123023002-0313311332213222-2113013102122132-3300101230133012-2302103332131331"></a>

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

<a id="canonical-1313323333213130-2103021023201312-2001113121132102-2130221032003311-1111122313211233-0121323222000031-1302103113113233-0310101311331322"></a>

### Direct properties for `more_option.request_cookies_to_add.secret_value.clear_secret_info`

<a id="canonical-3112213213133120-0100320131002101-1120220130103300-0222121112100321-3300031321021231-0022112102330313-3011231023320322-1101031211011131"></a>

#### `more_option.request_cookies_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3323222032032221-3110001332310133-0020010122122132-3220020131301200-2002023311330023-0300012001222202-2011223012103022-3202110331301122"></a>

<a id="canonical-0132012303300311-3302120000011011-1032303031110103-2223223121221000-2322222203313320-1130021331013131-0230033100011301-1330102313002313"></a>

#### `more_option.request_cookies_to_add.secret_value.clear_secret_info.url` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3131323333013211-1120301133201010-1211330033032130-3121231320313312-1312311200203221-3112223310131110-3201331320102000-1013203220202110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.request_headers_to_add` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- more_option.request_headers_to_add

<a id="canonical-0321032021212033-0210322223131322-2231000020302232-1303232131233331-0021111011201303-2332201032213120-3002033121231021-3021030102100313"></a>

Type: `"list"`. Computed.

Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers
specified at this level are applied after headers from matched Route are applied.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0333001323030031-2322003010131130-2132120330203002-0010313233112212-1210130033010213-0013010012330122-3023013033320332-2113023232103122"></a>

### Direct properties for `more_option.request_headers_to_add`

<a id="canonical-2123201113023022-1022210101033122-0220100103012112-3303011310120322-3132111313132120-0310113203131003-0310222001022232-0010231303021300"></a>

#### `more_option.request_headers_to_add.append` property

Type: `"bool"`. Computed.

Should the value be appended? If true, the value is appended to existing values. not append.
Defaults to \`do\`.

Additional upstream details:

If true, the value is appended to existing values. Default value is do not append.

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

<a id="canonical-1020223100213021-3220011001223022-3120130002212211-0010232112201020-0232320101221202-2302221222121010-0103231332033102-1322313011101011"></a>

<a id="canonical-3002110220300130-0222012323331101-0121223121001322-2332011022323313-3002001233021010-3202012031231133-1131301001013300-1330000221300221"></a>

#### `more_option.request_headers_to_add.name` property

Type: `"string"`. Computed.

Name. Name of the HTTP header.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [secret_value](data-sources--http_loadbalancer--reference--group-020.md#canonical-1131121000130320-1202130021200221-1000000220302321-0200030221313113-3002002321113003-3013132122301312-3130102313111120-1023001113122030): complete subsection reference.

<a id="canonical-3311002101301113-2000133133121021-2010132212010003-3222122323313333-1000223210021313-2102332023131221-0020210011231201-3100001300012001"></a>

<a id="canonical-1302231220212120-2232120222002130-1012221233101222-1003332030121013-2111010023320032-1231302131030301-0332303002010021-1123301123233221"></a>

#### `more_option.request_headers_to_add.value` property

Type: `"string"`. Computed.

Exclusive with \[secret\_value\] Value of the HTTP header.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1131121000130320-1202130021200221-1000000220302321-0200030221313113-3002002321113003-3013132122301312-3130102313111120-1023001113122030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.request_headers_to_add.secret_value` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.request_headers_to_add](data-sources--http_loadbalancer--reference--group-020.md#canonical-3131323333013211-1120301133201010-1211330033032130-3121231320313312-1312311200203221-3112223310131110-3201331320102000-1013203220202110)
- more_option.request_headers_to_add.secret_value

<a id="canonical-3102310320321223-0002031100103003-1230132232331013-2033011001232011-3102200232110113-0331132100311233-2220101321322211-2130010012103101"></a>

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

<a id="canonical-3200020131331221-0202333131113303-0131331010210331-1323200221323323-3002102000320210-2202231200002333-3121212220021332-3003013332022212"></a>

### Direct properties for `more_option.request_headers_to_add.secret_value`

- [blindfold_secret_info](data-sources--http_loadbalancer--reference--group-020.md#canonical-2303110233123102-0223032021010001-3330233323001203-2010010032222233-1202312110113333-3312320030110333-3120023020231122-1322301300320002): complete subsection reference.

- [clear_secret_info](data-sources--http_loadbalancer--reference--group-020.md#canonical-3333033032103310-0321020023231321-2331031212121033-0330213321121131-2123032230333122-1111202213202312-1211212221013132-3311312323321102): complete subsection reference.

<a id="canonical-2303110233123102-0223032021010001-3330233323001203-2010010032222233-1202312110113333-3312320030110333-3120023020231122-1322301300320002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.request_headers_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.request_headers_to_add](data-sources--http_loadbalancer--reference--group-020.md#canonical-3131323333013211-1120301133201010-1211330033032130-3121231320313312-1312311200203221-3112223310131110-3201331320102000-1013203220202110)
- [more_option.request_headers_to_add.secret_value](data-sources--http_loadbalancer--reference--group-020.md#canonical-1131121000130320-1202130021200221-1000000220302321-0200030221313113-3002002321113003-3013132122301312-3130102313111120-1023001113122030)
- more_option.request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-3220321331223321-1122030132112030-2102001031332301-2311233120300003-0031331321020010-0212221212230003-0000331310103221-1322300233123021"></a>

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

<a id="canonical-1010100131212232-2001221030033123-2332233010121110-2300022021021323-3113021031010131-0323330112321223-3332333011330101-1011033003323202"></a>

### Direct properties for `more_option.request_headers_to_add.secret_value.blindfold_secret_info`

<a id="canonical-3113211213011303-2331331030030013-1230232131313330-0030230012031310-0021301333030302-2321022003121100-3212323031110320-0123211232301023"></a>

#### `more_option.request_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1011220203112322-2122122110322021-2201230313231222-0200232110203210-1000020220301101-0002121223331002-1222311233021130-1301322023330213"></a>

<a id="canonical-2123303133031001-1231111131211012-2231021233331101-2120102133332023-1321022113330121-2012023321310212-3130313330102320-0322022110021111"></a>

#### `more_option.request_headers_to_add.secret_value.blindfold_secret_info.location` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1200201102111100-3210010001130102-2022020023000101-3232011310123011-2032033321031021-1003011221303223-1320101222322123-3132001130000302"></a>

<a id="canonical-0123011131230221-1232233010031030-2013232023302111-0312100312133311-2030022200021020-3123132301202330-3130300310002222-1200202200303121"></a>

#### `more_option.request_headers_to_add.secret_value.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3333033032103310-0321020023231321-2331031212121033-0330213321121131-2123032230333122-1111202213202312-1211212221013132-3311312323321102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.request_headers_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- [more_option.request_headers_to_add](data-sources--http_loadbalancer--reference--group-020.md#canonical-3131323333013211-1120301133201010-1211330033032130-3121231320313312-1312311200203221-3112223310131110-3201331320102000-1013203220202110)
- [more_option.request_headers_to_add.secret_value](data-sources--http_loadbalancer--reference--group-020.md#canonical-1131121000130320-1202130021200221-1000000220302321-0200030221313113-3002002321113003-3013132122301312-3130102313111120-1023001113122030)
- more_option.request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-0032223111303303-3323103102122302-0213221210313210-0310220113112321-3332330102030202-3103313110113021-1212210312231213-3033112012322023"></a>

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

<a id="canonical-3210200002031022-2321321331032332-2212120333002010-2123011311323120-2313122212123313-1312202001320020-1321133331003110-0132330311030221"></a>

### Direct properties for `more_option.request_headers_to_add.secret_value.clear_secret_info`

<a id="canonical-0100032101231311-1211100122121223-3002311003332023-1212001121031111-3201012203133130-0030122130131333-0202023112321132-1230032032012300"></a>

#### `more_option.request_headers_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3110311303122221-2202310102233130-0112023222223122-0131213302312212-3303321010222330-1332211302213110-1223202322300303-0003303123310101"></a>

<a id="canonical-0222222122230213-0332222113320113-0302031132201230-3213031102332110-2013020020321110-2223213103223223-1022333123002200-2320113100112222"></a>

#### `more_option.request_headers_to_add.secret_value.clear_secret_info.url` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3211332311023321-0003133002230231-2301220231333002-0203020131002230-0223211100102320-1330323222012202-3312210301331303-3032231011333132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `more_option.response_cookies_to_add` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [more_option](data-sources--http_loadbalancer--reference--group-020.md#canonical-3020122111213230-3123110331301013-0212331102130112-1130121311213313-0102332203020303-2031330330302330-2131212323210012-3212203312120101)
- more_option.response_cookies_to_add

<a id="canonical-2101101032203133-3212123103302333-1000313321122021-1111222312312303-1221213323120033-2211011210010322-1210003322020023-3311220302023212"></a>

Type: `"list"`. Computed.

Cookies are name-value pairs along with optional attribute parameters to be added to HTTP response
being sent towards downstream. Cookies specified at this level are applied after cookies from
matched Route are applied.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0131023102123001-3330202032112311-2322223331121023-2121011233012320-3322222230121021-1032202203131201-2321112032223102-0323203311310211"></a>

### Direct properties for `more_option.response_cookies_to_add`

<a id="canonical-1301110313211001-0220322220320233-3112123012222112-2002310120110100-1300010323021320-3000233021233302-1201110122210313-2121002321201301"></a>

#### `more_option.response_cookies_to_add.add_domain` property

Type: `"string"`. Computed.

Exclusive with \[ignore\_domain\] Add domain attribute.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1032120010211121-1313120021010212-2032210122310010-0020110213311231-3003113103011010-2030100311312003-1232313111331131-0321032031312320"></a>

<a id="canonical-0023303330003000-0023130033131122-2232231222020210-3102210123210302-1003133033030022-1132010202220030-2321003120230011-1133312020230231"></a>

#### `more_option.response_cookies_to_add.add_expiry` property

Type: `"string"`. Computed.

Exclusive with \[ignore\_expiry\] Add expiry attribute.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [add_httponly](data-sources--http_loadbalancer--reference--group-021.md#canonical-0300223001322130-2030032331032033-1003212011020133-2233320201331011-0030003131232030-2312222132220200-0001123123322032-2313020031031221): complete subsection reference.

- [add_partitioned](data-sources--http_loadbalancer--reference--group-021.md#canonical-2330103311303032-3130031021333331-1202200000210223-1010020332113233-3213302010103030-1032313211131110-1222033112230303-2200112031020203): complete subsection reference.

<a id="canonical-3001002123132133-3131310312330312-0130103221233000-3130210011032022-2023223111312112-1103023132013022-1001133202012223-1030000030323101"></a>

<a id="canonical-2221121311011331-2320022101001012-1001100021011200-3231013210320111-1012211311302031-1122330100320103-3301003002011333-2212322232301101"></a>

#### `more_option.response_cookies_to_add.add_path` property

Type: `"string"`. Computed.

Exclusive with \[ignore\_path\] Add path attribute.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [add_secure](data-sources--http_loadbalancer--reference--group-021.md#canonical-1120032113222132-2213222211032003-1322223121013023-2312012011220113-2000301120202213-0212000001233232-0103230330300101-1111330203321121): complete subsection reference.

- [ignore_domain](data-sources--http_loadbalancer--reference--group-021.md#canonical-0312220213101021-3030212313010302-3202012101032100-2111031101023331-3120133123002032-2223302102000101-2330220300103220-0111233100003032): complete subsection reference.

- [ignore_expiry](data-sources--http_loadbalancer--reference--group-021.md#canonical-2022131312122120-0313230201310123-1101001012033031-3210112120013101-3230021320030132-3013030010002033-3312223102022221-3000132002223012): complete subsection reference.

- [ignore_httponly](data-sources--http_loadbalancer--reference--group-021.md#canonical-2330101032030000-0303210323310033-3012130033223300-0002321010133212-1002323223130123-3212131223020023-3310023220121103-3323022002320332): complete subsection reference.

- [ignore_max_age](data-sources--http_loadbalancer--reference--group-021.md#canonical-1020012011300002-0212210002102001-2332232100130203-1310112320331220-0321213103033111-3321203221331232-3233103020213010-2231210113023131): complete subsection reference.

- [ignore_partitioned](data-sources--http_loadbalancer--reference--group-021.md#canonical-1112032110322101-0323332202012031-1111232121233213-0100112000110132-1312301300330230-0203012013200310-2202212102303321-0201113021011112): complete subsection reference.

- [ignore_path](data-sources--http_loadbalancer--reference--group-021.md#canonical-1123231131001201-3301320030020311-0021020221303001-1213311000213321-1133011030330202-1122013213211230-2033103300113223-2230232203220113): complete subsection reference.

- [ignore_samesite](data-sources--http_loadbalancer--reference--group-021.md#canonical-3310323001033201-2132300030302321-0131212102122102-2311300130333121-1021223233022010-0110233321213022-3003311000212200-1100013303321211): complete subsection reference.

- [ignore_secure](data-sources--http_loadbalancer--reference--group-021.md#canonical-1032301223020021-1112010033311012-1021301212331103-2133132101232310-2120231003021210-1022132222130213-1003021013201232-2031020120332230): complete subsection reference.

- [ignore_value](data-sources--http_loadbalancer--reference--group-021.md#canonical-3323210223311103-2002213230221021-1310310300003112-1331323212212201-2232002302031111-3003312203213313-0020013010022100-1310301302011320): complete subsection reference.

<a id="canonical-2023100221122231-3100023321123130-2232220200130011-3303030000313231-1201032133133301-3001212220020233-0101111331021320-0233313132103010"></a>
