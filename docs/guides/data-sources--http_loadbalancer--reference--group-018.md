---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-3132323100013201-3210103000323202-0120303213201211-3200221003122313-0313121001120132-1231133311121230-1021201002210111-3031132030331113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_api_discovery.enable_learn_from_redirect_traffic` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-3102231103303332-0011102023001333-1001101113313230-0310003203030002-0230021033330011-0020301130222310-2002330010130011-0220313000312223)
- enable_api_discovery.enable_learn_from_redirect_traffic

<a id="canonical-2021020330130301-0230203110102213-1230333213203223-0223022222212020-0011323332110112-0221002323330220-2033033023322232-2331100200002130"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable learn from redirect traffic.

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

<a id="canonical-0011200023030232-1103032103102003-2020010023111201-0113013102133101-1133013020111332-3130102202020210-0213300311210300-1123323322102222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_challenge` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- enable_challenge

<a id="canonical-2230132011120203-0011333021033130-1010233313001013-1100313220213320-3231313122231210-2033232202103112-3100002200122100-0000223330000031"></a>

Type: `"single"`. Computed.

Configure auto mitigation i.e risk based challenges for malicious users.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-captcha_challenge_parameters_choice": "[\"captcha_challenge_parameters\",\"default_captcha_challenge_parameters\"]",
  "x-ves-oneof-field-js_challenge_parameters_choice": "[\"default_js_challenge_parameters\",\"js_challenge_parameters\"]",
  "x-ves-oneof-field-malicious_user_mitigation_choice": "[\"default_mitigation_settings\",\"malicious_user_mitigation\"]"
}
```

<a id="canonical-2130120103002232-3322011322003101-3220121132133012-1012203203032013-3123121122021201-0301211032322001-2023130233101330-3301230132003313"></a>

### Direct properties for `enable_challenge`

- [captcha_challenge_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-0121030213210321-1322302210313013-3112332123323020-1321123131333121-2131113211232020-3223111112222001-0011102330021020-3101111331133011): complete subsection reference.

- [default_captcha_challenge_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-3320320211033302-3033120013010222-0121120301332033-2233122323233320-1001302333013233-0310332112021003-3332022112202301-1223121323133021): complete subsection reference.

- [default_js_challenge_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-3102030212122002-0233012113210031-3000311301112132-3002110312202303-0121122210100000-1221122020103022-0302100222232211-3112000303123131): complete subsection reference.

- [default_mitigation_settings](data-sources--http_loadbalancer--reference--group-018.md#canonical-3033312103102312-3220031122103201-3003103002210011-1102312201122202-1320110311222313-3313112222220202-2223313011001322-2111030302232012): complete subsection reference.

- [js_challenge_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-2102131322013321-3103000102211223-2203223000013002-0030321001003323-3123330210130223-1111302200033320-2030022100133123-0333011120000212): complete subsection reference.

- [malicious_user_mitigation](data-sources--http_loadbalancer--reference--group-018.md#canonical-1320102211302230-1331211303022203-3011223102123110-3012311303213230-3233312022320330-0130203012003103-0020012330231303-3212030013231120): complete subsection reference.

<a id="canonical-0121030213210321-1322302210313013-3112332123323020-1321123131333121-2131113211232020-3223111112222001-0011102330021020-3101111331133011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_challenge.captcha_challenge_parameters` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_challenge](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011200023030232-1103032103102003-2020010023111201-0113013102133101-1133013020111332-3130102202020210-0213300311210300-1123323322102222)
- enable_challenge.captcha_challenge_parameters

<a id="canonical-3032100223311200-0221203213210113-1011331133000232-1130230313120232-2303131100131332-1120112133323303-0010201123012031-2301221131103322"></a>

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

<a id="canonical-2121233233103301-3112003011110122-2000210023303320-3313211302120301-0222321101003001-0013003113332301-3330022100013113-2213303233122303"></a>

### Direct properties for `enable_challenge.captcha_challenge_parameters`

<a id="canonical-3202123230132221-1330023031201010-0030232332310133-3202202111022123-3100221022010032-1302013311221100-1212330320123023-1320221020213102"></a>

#### `enable_challenge.captcha_challenge_parameters.cookie_expiry` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3132121102220233-2030221330210210-1301232021013322-0323023132121032-0103303103210133-1301131312012313-0201001222103232-1110023022100310"></a>

<a id="canonical-0333302031110233-0022020320031302-2032231010112231-2232322213212110-2223112100132323-0002110201133030-2200223231102322-3232320001000011"></a>

#### `enable_challenge.captcha_challenge_parameters.custom_page` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3320320211033302-3033120013010222-0121120301332033-2233122323233320-1001302333013233-0310332112021003-3332022112202301-1223121323133021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_challenge.default_captcha_challenge_parameters` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_challenge](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011200023030232-1103032103102003-2020010023111201-0113013102133101-1133013020111332-3130102202020210-0213300311210300-1123323322102222)
- enable_challenge.default_captcha_challenge_parameters

<a id="canonical-1311212210201333-3310002323313231-3032331230311333-1033221302322211-3312323302012013-3222132213221130-0110000133330100-1122111220210130"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default captcha challenge parameters.

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

<a id="canonical-3102030212122002-0233012113210031-3000311301112132-3002110312202303-0121122210100000-1221122020103022-0302100222232211-3112000303123131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_challenge.default_js_challenge_parameters` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_challenge](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011200023030232-1103032103102003-2020010023111201-0113013102133101-1133013020111332-3130102202020210-0213300311210300-1123323322102222)
- enable_challenge.default_js_challenge_parameters

<a id="canonical-3020201203201200-2231312222121222-0210332232000300-1221220323011002-2223000220201322-2030132011122120-2133113113213100-0320333100201231"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default js challenge parameters.

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

<a id="canonical-3033312103102312-3220031122103201-3003103002210011-1102312201122202-1320110311222313-3313112222220202-2223313011001322-2111030302232012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_challenge.default_mitigation_settings` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_challenge](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011200023030232-1103032103102003-2020010023111201-0113013102133101-1133013020111332-3130102202020210-0213300311210300-1123323322102222)
- enable_challenge.default_mitigation_settings

<a id="canonical-1011011120021113-1220233020132201-0311312110031312-2330003233133123-3103102211211231-0100313033133322-1033132133130301-2213222232020200"></a>

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

<a id="canonical-2102131322013321-3103000102211223-2203223000013002-0030321001003323-3123330210130223-1111302200033320-2030022100133123-0333011120000212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_challenge.js_challenge_parameters` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_challenge](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011200023030232-1103032103102003-2020010023111201-0113013102133101-1133013020111332-3130102202020210-0213300311210300-1123323322102222)
- enable_challenge.js_challenge_parameters

<a id="canonical-0112022001032023-3012003032213000-3102330113232230-3110012131320122-0033013133210200-1023323100101031-3313321131332302-0002132233000130"></a>

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

<a id="canonical-1232322301223133-0232002033003112-3231122222123013-0120312211032220-0232300130110033-1101132002122121-1121121231331221-2231132201030130"></a>

### Direct properties for `enable_challenge.js_challenge_parameters`

<a id="canonical-3200331223320103-3201222111012220-0322120130303032-0223222312202200-0332332303301333-1110220311013032-2302313131222131-2120210333001000"></a>

#### `enable_challenge.js_challenge_parameters.cookie_expiry` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0302033301101101-0133210200223121-0312331030323012-2230320331321203-2010320312210001-1222232201203133-1201021013010110-1300220033200113"></a>

<a id="canonical-3101310112232122-3111303223132333-2321221202011211-0113021013100321-3202203311111023-0200220103322123-2333301211130322-1002331310232311"></a>

#### `enable_challenge.js_challenge_parameters.custom_page` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1130132210022130-3323002222032031-0322210310110213-3200212000101011-3202001131221122-1321222100003313-2311310101121101-3003012230230031"></a>

<a id="canonical-0020032011111031-2221303330123100-2123031211121021-2113310211213100-0202301320203132-0102010112013003-2301011113301300-0221131301113101"></a>

#### `enable_challenge.js_challenge_parameters.js_script_delay` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1320102211302230-1331211303022203-3011223102123110-3012311303213230-3233312022320330-0130203012003103-0020012330231303-3212030013231120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_challenge.malicious_user_mitigation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_challenge](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011200023030232-1103032103102003-2020010023111201-0113013102133101-1133013020111332-3130102202020210-0213300311210300-1123323322102222)
- enable_challenge.malicious_user_mitigation

<a id="canonical-1223221130111102-0330010312211300-2232203330222213-3131013211320030-1302220223230013-1320333032212221-0313111033023212-3131213312101012"></a>

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

<a id="canonical-2131121100220011-1033322123013020-1101131201021110-1110323122303203-0120320010102232-3213202323211133-1330102203122022-0101010100101200"></a>

### Direct properties for `enable_challenge.malicious_user_mitigation`

<a id="canonical-3012221002001300-3003303023032311-2300000032010213-3033101213210332-2213303321321112-1301301203110200-2103102332120001-0232303113212013"></a>

#### `enable_challenge.malicious_user_mitigation.name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3331120333300000-0210220022321232-2212031122321230-1100303321002000-3321132232333330-3201230310100210-3103030000100311-2023002102020023"></a>

<a id="canonical-2301020112313021-3011212131100221-0220232203203030-3123211103220220-0310331323000221-1313333123301200-0332232233232331-1000133223212302"></a>

#### `enable_challenge.malicious_user_mitigation.namespace` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2223021212333312-1332213102123031-3000023132103203-2330311123203103-0203122232233131-3013303003110120-2301210032203030-1001333110330023"></a>

<a id="canonical-1121000311023310-3221022023321103-3322310203323112-1102130122321231-3110000333022211-0313210301310223-2311301203321303-1301022101222222"></a>

#### `enable_challenge.malicious_user_mitigation.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2222220232120233-1130313100210223-1110323001300223-2031010002311213-3032120333310102-3022103310121023-0103023331030110-0220211011320211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_ip_reputation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- enable_ip_reputation

<a id="canonical-1303123101122023-3103220103322001-2320130011231101-2030202013121330-2011333132210133-2222302130220303-2222103112211003-3201100223033010"></a>

Type: `"single"`. Computed.

IP Threat Category List. List of IP threat categories.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2123313103013301-1313033233311031-1101322033333312-1010030003331020-3223212110310321-3330011012201233-2100330202221313-0120330001333230"></a>

### Direct properties for `enable_ip_reputation`

<a id="canonical-1012112323213010-0121011310311223-0133232321330210-1032112020223022-0211133113012222-0330031303122203-1103303321001021-3120000302000011"></a>

#### `enable_ip_reputation.ip_threat_categories` property

Type: `["list", "string"]`. Computed.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
If the source IP matches on atleast one of the enabled IP threat categories, the request will be
denied. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`, \`WEB\_ATTACKS\`, \`BOTNETS\`,
\`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`, \`MOBILE\_THREATS\`, \`TOR\_PROXY\`,
\`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to \`SPAM\_SOURCES\`.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3112321013223011-0232120223301113-2123300000310001-2303013212022230-1020121313200102-3113021131130312-3132303103233220-0130203210202122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_malicious_user_detection` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- enable_malicious_user_detection

<a id="canonical-2110332331011132-0112220203021031-2310200300312113-0000113031331323-0022111100132113-1010312111223221-0011212101030212-2221110322223112"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable malicious user detection.

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

<a id="canonical-1111231022131133-1331313200113310-3310302011012003-0030332010313023-0230102231302212-2013001002111223-2120221113320103-0133233131213321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_threat_mesh` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- enable_threat_mesh

<a id="canonical-3202002102313212-2121133022323302-2332320333313223-1312303022323030-3030133001022000-0320122330221310-0003032232310211-2110113111030212"></a>

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

<a id="canonical-0223302231131311-1212022331012012-1201212100221020-3222300100201333-2302211010002001-3333320223120100-2001323010222333-0112102133220100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `enable_trust_client_ip_headers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- enable_trust_client_ip_headers

<a id="canonical-2200022210121310-3030222023303302-1130102212110230-2323302120331110-2331233111002022-1020110033203002-0231333331120233-3013121212303110"></a>

Type: `"single"`. Computed.

Trust Client IP Headers List. List of Client IP Headers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3303100231010101-2310222120103021-0123300130020201-1320223220321320-0231201003303011-2332003222332330-3011302020122000-3201330123021023"></a>

### Direct properties for `enable_trust_client_ip_headers`

<a id="canonical-2311033000302102-1230312213303000-0310232120202323-2122022102222021-1231201030320330-1333331122333200-2222021213313203-1231222012123122"></a>

#### `enable_trust_client_ip_headers.client_ip_headers` property

Type: `["list", "string"]`. Computed.

Define the list of one or more Client IP Headers. Headers will be used in order from top to bottom,
meaning if the first header is not present in the request, the system will proceed to check for the
second header, and so on, until one of the listed headers is found. If none of the defined headers
exist, or the value is not an IP address, then the system will use the source IP of the packet. If
multiple defined headers with different names are present in the request, the value of the first
header name in the configuration will be used. If multiple defined headers with the same name are
present in the request, values of all those headers will be combined. The system will read the
right-most IP address from header, if there are multiple IP addresses in the header value. For
X-Forwarded-For header, the system will read the IP address(rightmost - 1), as the client IP.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0103130023112220-0323202202232201-0100110021021230-0011021230100313-0322001310012013-3211213332301231-0231033133330111-0023122020322232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- graphql_rules

<a id="canonical-3003300301222323-0112102102301002-0113203300032213-3032201212021111-3200003212221223-0030112023020101-3221203102131101-2313232223103031"></a>

Type: `"list"`. Computed.

GraphQL is a query language and server-side runtime for APIs which provides a complete and
understandable description of the data in API. GraphQL gives clients the power to ask for exactly
what they need, makes it easier to evolve APIs over time, and enables powerful developer tools.
Policy configuration to analyze GraphQL queries and prevent GraphQL tailored attacks.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-1203231330111332-1223101331302103-2332023222001122-0332123231101030-1233103133331020-2222021021032021-2223211020230100-1222311310321311"></a>

### Direct properties for `graphql_rules`

- [any_domain](data-sources--http_loadbalancer--reference--group-018.md#canonical-1103202203211103-3211330223021312-1123113112120011-1233031121112012-1021100231020021-2020333210313301-0021123232303112-0300000110232202): complete subsection reference.

<a id="canonical-1211112232001330-0011213223031200-3123002002221022-1120003133100000-0220321303321130-0300321333023110-2010123312322011-2002310220001132"></a>

<a id="canonical-1303030001312201-2120022002332303-2011033001331120-3131033313313312-2303300102330132-3120323033010033-2220101112022332-1320202112000303"></a>

#### `graphql_rules.exact_path` property

Type: `"string"`. Computed.

Specifies the exact path to GraphQL endpoint. Defaults to \`/graphql\`.

Additional upstream details:

Default value is /GraphQL.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2311132103301320-3130110111011310-3121301033220113-2200220302312202-0331213331001221-2330133200320132-0332010201110230-1012211232011022"></a>

<a id="canonical-1202133120303200-2012000333212320-3310032113231200-2032201001323303-1222330323322220-3210000003100223-2001212022102201-0320023223121203"></a>

#### `graphql_rules.exact_value` property

Type: `"string"`. Computed.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [graphql_settings](data-sources--http_loadbalancer--reference--group-018.md#canonical-0123001120213220-2233121321230101-1121121021122232-2222133312301221-0002121031222210-2213300001322210-1121112013301123-0233210112333010): complete subsection reference.

- [metadata](data-sources--http_loadbalancer--reference--group-018.md#canonical-0331131331213103-3103312023301300-0303321120101113-1323132201332122-3010203210003332-0132112222323313-0300230120102230-0200300120030102): complete subsection reference.

- [method_get](data-sources--http_loadbalancer--reference--group-018.md#canonical-2213131120013323-3301221131310223-0021331301012202-2012020231122113-2320030003221211-2011022111211030-2033112121032000-0121013330103132): complete subsection reference.

- [method_post](data-sources--http_loadbalancer--reference--group-018.md#canonical-3222113113333222-1110230032332100-3230132102303033-3021033313031223-1323103130002112-1230113120111312-3022313200321012-2021113111302130): complete subsection reference.

<a id="canonical-2322012123010321-3223301302302110-0330011221003331-1333220231001233-0222122201303130-3321031333001133-0301203311100033-3112103221311203"></a>

<a id="canonical-1203213322031220-0031131333011223-3010320101111021-3212230122130001-2132203013011010-1022333001000233-1132001012100300-1001113230322330"></a>

#### `graphql_rules.suffix_value` property

Type: `"string"`. Computed.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1103202203211103-3211330223021312-1123113112120011-1233031121112012-1021100231020021-2020333210313301-0021123232303112-0300000110232202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules.any_domain` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [graphql_rules](data-sources--http_loadbalancer--reference--group-018.md#canonical-0103130023112220-0323202202232201-0100110021021230-0011021230100313-0322001310012013-3211213332301231-0231033133330111-0023122020322232)
- graphql_rules.any_domain

<a id="canonical-2312213232002023-1020230030200310-2130312302100223-0030203120121230-3311030320001021-2311211322310010-1303212222201231-2323103012123302"></a>

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

<a id="canonical-0123001120213220-2233121321230101-1121121021122232-2222133312301221-0002121031222210-2213300001322210-1121112013301123-0233210112333010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules.graphql_settings` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [graphql_rules](data-sources--http_loadbalancer--reference--group-018.md#canonical-0103130023112220-0323202202232201-0100110021021230-0011021230100313-0322001310012013-3211213332301231-0231033133330111-0023122020322232)
- graphql_rules.graphql_settings

<a id="canonical-0303023310303203-1012022212101111-1312302301131313-2222022312011131-3130103102331113-2213013223222133-2201133320312023-3220233121132301"></a>

Type: `"single"`. Computed.

Configuration parameter for GraphQL settings.

Additional upstream details:

GraphQL configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-allow_introspection_queries_choice": "[\"disable_introspection\",\"enable_introspection\"]"
}
```

<a id="canonical-1133010011333231-2121223202122113-2030021111232103-1022230011213233-0110232101021003-3003233102001133-2102330230300122-0102202223222021"></a>

### Direct properties for `graphql_rules.graphql_settings`

- [disable_introspection](data-sources--http_loadbalancer--reference--group-018.md#canonical-0012210321003021-0131310211201301-2231210000010011-3012020310011010-1321113220200330-3323113000331123-0011121210023111-0123133333310232): complete subsection reference.

- [enable_introspection](data-sources--http_loadbalancer--reference--group-018.md#canonical-2102110122303030-2330303013000003-2103130320031002-0103012222203332-2122133010233322-1101313001311033-2131321203213310-2103121010001121): complete subsection reference.

<a id="canonical-1113131222001120-2000213310301321-3003013322112102-2220000302321312-1323330131313231-2230212332121120-0010321332110221-0022121033331231"></a>

<a id="canonical-1032233322233100-1030102211223010-1312103112311232-2001203222002111-0201322222303021-2122130121033222-2031102213121123-2302020303230332"></a>

#### `graphql_rules.graphql_settings.max_batched_queries` property

Type: `"number"`. Computed.

Specify maximum number of queries in a single batched request.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```

<a id="canonical-3121321230030223-1303132121033033-3130112323313232-0010001003010012-1131001212222102-0110133223200310-0231011003122303-1000310032012000"></a>

<a id="canonical-3322200100203302-0332013210331113-2300112023101113-0110311021202030-2221221023113001-0000232030132231-3223330023213121-3122102331330103"></a>

#### `graphql_rules.graphql_settings.max_depth` property

Type: `"number"`. Computed.

Specify maximum depth for the GraphQL query.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```

<a id="canonical-0200000033130222-3332001233201213-1033222121130203-2111020203223012-2233202210313020-3202221030101003-1130223212210323-1130030123001101"></a>

<a id="canonical-3201322313312332-0001312020310213-2120210003000331-2201302201102030-1222210031210332-0033103202303023-2132000323203111-1313201221031223"></a>

#### `graphql_rules.graphql_settings.max_total_length` property

Type: `"number"`. Computed.

Specify maximum length in bytes for the GraphQL query.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16386,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.lte": "16386"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "16386"
  }
}
```

<a id="canonical-0012210321003021-0131310211201301-2231210000010011-3012020310011010-1321113220200330-3323113000331123-0011121210023111-0123133333310232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules.graphql_settings.disable_introspection` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [graphql_rules](data-sources--http_loadbalancer--reference--group-018.md#canonical-0103130023112220-0323202202232201-0100110021021230-0011021230100313-0322001310012013-3211213332301231-0231033133330111-0023122020322232)
- [graphql_rules.graphql_settings](data-sources--http_loadbalancer--reference--group-018.md#canonical-0123001120213220-2233121321230101-1121121021122232-2222133312301221-0002121031222210-2213300001322210-1121112013301123-0233210112333010)
- graphql_rules.graphql_settings.disable_introspection

<a id="canonical-1302310210100031-1123101200113103-3220003022122001-0302011303301231-3133301230200303-3303103303201031-3020112022331230-3233323012003221"></a>

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

<a id="canonical-2102110122303030-2330303013000003-2103130320031002-0103012222203332-2122133010233322-1101313001311033-2131321203213310-2103121010001121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules.graphql_settings.enable_introspection` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [graphql_rules](data-sources--http_loadbalancer--reference--group-018.md#canonical-0103130023112220-0323202202232201-0100110021021230-0011021230100313-0322001310012013-3211213332301231-0231033133330111-0023122020322232)
- [graphql_rules.graphql_settings](data-sources--http_loadbalancer--reference--group-018.md#canonical-0123001120213220-2233121321230101-1121121021122232-2222133312301221-0002121031222210-2213300001322210-1121112013301123-0233210112333010)
- graphql_rules.graphql_settings.enable_introspection

<a id="canonical-0212011322132211-0112331031021003-0021313310333023-1232300201333210-2111132120023311-2212210220310133-3103211000120111-3231333013322012"></a>

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

<a id="canonical-0331131331213103-3103312023301300-0303321120101113-1323132201332122-3010203210003332-0132112222323313-0300230120102230-0200300120030102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [graphql_rules](data-sources--http_loadbalancer--reference--group-018.md#canonical-0103130023112220-0323202202232201-0100110021021230-0011021230100313-0322001310012013-3211213332301231-0231033133330111-0023122020322232)
- graphql_rules.metadata

<a id="canonical-1013113033301002-0322300022011321-1112033212321301-3121303300023132-2032223231230021-0030213320030110-2220330200330200-1213200132201321"></a>

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

<a id="canonical-3010232220113001-1013232130222113-2310230002013133-1211131011201301-3112333211103320-2111002232102213-0233021302233133-3203300131102312"></a>

### Direct properties for `graphql_rules.metadata`

<a id="canonical-3011103202022300-2311330331122223-0003133313301230-0102000020133230-1110320312231021-2321221211220231-3113223202032020-3123002313310303"></a>

#### `graphql_rules.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-3131323021112101-3333133230112300-1230111301231301-3131231331233030-2001012102031303-0303010011011033-1123132323131230-3330001312032330"></a>

<a id="canonical-2022111113020220-3031232231023133-0311310313110220-3022200122100233-1121110013133113-3000202030133302-3012112202220331-0122031001331130"></a>

#### `graphql_rules.metadata.name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2213131120013323-3301221131310223-0021331301012202-2012020231122113-2320030003221211-2011022111211030-2033112121032000-0121013330103132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules.method_get` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [graphql_rules](data-sources--http_loadbalancer--reference--group-018.md#canonical-0103130023112220-0323202202232201-0100110021021230-0011021230100313-0322001310012013-3211213332301231-0231033133330111-0023122020322232)
- graphql_rules.method_get

<a id="canonical-3122110333003031-3301301230323030-1033002311011021-1232332313312032-3313311202212020-1222121230101132-0112213210210022-2200202033210210"></a>

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

<a id="canonical-3222113113333222-1110230032332100-3230132102303033-3021033313031223-1323103130002112-1230113120111312-3022313200321012-2021113111302130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `graphql_rules.method_post` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [graphql_rules](data-sources--http_loadbalancer--reference--group-018.md#canonical-0103130023112220-0323202202232201-0100110021021230-0011021230100313-0322001310012013-3211213332301231-0231033133330111-0023122020322232)
- graphql_rules.method_post

<a id="canonical-2102133311231221-3231102332301020-1200103133131020-3333003021112000-1020003202011201-1111202302312110-1130120110130303-1122321033203133"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for method post.

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

<a id="canonical-1221000122101103-3100213033303232-2132010133302010-2012231203222223-2220001112313332-0002130322200033-3133223330023000-2102003233113221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `http` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- http

<a id="canonical-0330312012020130-1032221113102101-2323210100123332-0323300213303303-1312010310011021-1100003013333313-1230010000210213-0202111322331322"></a>

Type: `"single"`. Computed.

HTTP Choice. Choice for selecting HTTP proxy. Changing this type selection requires recreation and
may interrupt service. Supported settings within the same selected type remain updatable.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]"
}
```

<a id="canonical-1301013210122310-3210231010213302-2002013230010311-0200020002223023-0130200031212210-1113110120201022-3323112300223130-3312322202020210"></a>

### Direct properties for `http`

<a id="canonical-2333320123122231-0113023133300110-1130333332032230-2323021310122230-2331123132221110-0100301232330020-1120323203310123-1003320220301121"></a>

#### `http.dns_volterra_managed` property

Type: `"bool"`. Computed.

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME
record should be created in your DNS provider's portal.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2300222103020010-2213302110302023-1121302331202330-1123202110201310-2323230011231322-0200231233233322-1313131212303121-3012220023002123"></a>

<a id="canonical-1201111003131033-0223321011332003-1332330130202110-3121102330222303-2030000320212130-0110202010202110-1020221311100212-1132132110211011"></a>

#### `http.port` property

Type: `"number"`. Computed.

Exclusive with \[port\_ranges\] HTTP port to Listen.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3222322022130121-3123001301122303-1010323020213032-3101302222321312-3133311303110113-2102020021203222-1102230210101010-3333220232320010"></a>

<a id="canonical-3220013032313220-1020031332123213-3213112313302313-3300120312033030-0222113302203302-0313102123222033-0112231203013132-2302131320303212"></a>

#### `http.port_ranges` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- https

<a id="canonical-1230020122213233-2010123001220332-2010332323100211-3023001312301001-0100321202103310-1030231133101202-0011031010011003-2002013130023101"></a>

Type: `"single"`. Computed.

Choice for selecting HTTP proxy with bring your own certificates. Changing this type selection
requires recreation and may interrupt service. Supported settings within the same selected type
remain updatable.

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
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]",
  "x-ves-oneof-field-server_header_choice": "[\"append_server_name\",\"default_header\",\"pass_through\",\"server_name\"]",
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_parameters\"]"
}
```

<a id="canonical-0212320231230202-3111023303203223-2220311021302312-1131131120012200-2020320033301303-3212313011131130-1103013011232032-0011102130110012"></a>

### Direct properties for `https`

<a id="canonical-2330100312310220-1130123313303213-2320002233211023-2221122020330233-1322022213302320-0200010301232131-0211201330012232-0311101032031331"></a>

#### `https.add_hsts` property

Type: `"bool"`. Computed.

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

<a id="canonical-1312321213302200-2311130132312113-2221003330200222-3013300111111203-2332101010302020-1332203121111130-1222012020200203-1102221010301231"></a>

<a id="canonical-0133021200020131-1010032132022303-1101312112312332-2013121002323023-0121122113220133-1013020023320001-1313013030012321-1323333030113003"></a>

#### `https.append_server_name` property

Type: `"string"`. Computed.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [coalescing_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-1001220203312330-3031013021133302-3130013202123010-1232020313310122-0213332300123110-3132233211103221-0220011120321111-3220032023002332): complete subsection reference.

<a id="canonical-1001301223000032-1020300012030222-1203311102100300-2030223230203003-2232111200202103-2122121330003311-1120203120321212-2301212021200331"></a>

<a id="canonical-1010310022131223-2211333003201131-3302212322131113-3101300100301331-1232123222131301-3331031130013132-1320211322221010-3130312320322330"></a>

#### `https.connection_idle_timeout` property

Type: `"number"`. Computed.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [default_header](data-sources--http_loadbalancer--reference--group-018.md#canonical-1011220220213223-0023022311131111-3321100321212103-2302012100332100-3310212111103331-2030012330210201-1100032222211300-2333103131312233): complete subsection reference.

- [default_loadbalancer](data-sources--http_loadbalancer--reference--group-018.md#canonical-1123003233230101-3321333100231312-1212220312212102-3022231131212132-1132223102003323-3030022222021003-1030200021021213-1030221101223101): complete subsection reference.

- [disable_path_normalize](data-sources--http_loadbalancer--reference--group-018.md#canonical-1313230312203121-0221130100202030-2130030031311110-0000203003102103-0330223030021321-3212202011001322-2323331303122122-3302211302001001): complete subsection reference.

- [enable_path_normalize](data-sources--http_loadbalancer--reference--group-018.md#canonical-3211010213301321-0310002022003033-2320121222112212-3322033031120213-3201231212313320-2233231203130122-3211232302302121-2012120022302233): complete subsection reference.

- [http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-0003020203103000-2333100023323311-1111213021211113-2303201112013303-2123231030312110-3202013021320132-3330331220213021-1110230322200021): complete subsection reference.

<a id="canonical-3031001212231120-3111211022232210-0331103001030120-0113311233000023-2322202320221102-3120231023311013-3010302033101002-0012201010103110"></a>

<a id="canonical-1330021220022031-0012131012301331-2321330023023201-0331310023321122-1312330011103131-1231303023320031-0331023101103301-3032212221103133"></a>

#### `https.http_redirect` property

Type: `"bool"`. Computed.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [non_default_loadbalancer](data-sources--http_loadbalancer--reference--group-018.md#canonical-1101330312012103-0233323210311221-1033012102103323-0220330330101233-1110011203231222-2231210133200221-1022100012023223-3112010202021011): complete subsection reference.

- [pass_through](data-sources--http_loadbalancer--reference--group-018.md#canonical-3210232200212003-0002112320100202-1003013010223122-1001223222303210-2322123332210310-3322112011113221-2311013202301022-0201323130200031): complete subsection reference.

<a id="canonical-2221113200201003-0032023033301302-2020120122212023-0003333331133102-1331201130310022-0033023130203113-2332302223212000-2213322012001302"></a>

<a id="canonical-0313022203230120-2223333120231230-1130202033020200-0301220033132330-3311002113333211-1032230113110131-2232213113203003-0320221232210100"></a>

#### `https.port` property

Type: `"number"`. Computed.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3001111332001110-0303213033020310-3003021121310020-0100033122330120-1320310003000230-2020013321002122-0013333330232023-1201003030132111"></a>

<a id="canonical-1001212303003123-3323030333103313-0211021300100001-3212222030011203-0101133102012322-1222100121332310-2330232000130013-1001113302121200"></a>

#### `https.port_ranges` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3203301010030220-1312311200112222-1021030233110111-0000200320200320-3031233213222101-2213233322332111-1123301030323213-3023302310030333"></a>

<a id="canonical-1312013231020220-2010021032112021-2232112310303333-3102311303021023-3110333232113022-1310232111022031-1003003301030313-2330302110002331"></a>

#### `https.server_name` property

Type: `"string"`. Computed.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011121033233220-3311202112112001-0220333111322122-2213202322313002-3213103232311233-2200102030312213-0032120231111113-3201230203322302): complete subsection reference.

- [tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031): complete subsection reference.

<a id="canonical-1001220203312330-3031013021133302-3130013202123010-1232020313310122-0213332300123110-3132233211103221-0220011120321111-3220032023002332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.coalescing_options` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- https.coalescing_options

<a id="canonical-0202223120121212-2311130013123131-0220023023210102-3100231213033113-1213022002201210-1222111202100133-2213311313000211-1112201321201201"></a>

Type: `"single"`. Computed.

TLS connection coalescing configuration (not compatible with mTLS).

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

<a id="canonical-3310001210132030-2231232213032313-1300312300020203-0113330103322022-0020301202322131-0032001310030333-3001010331231121-3113031023213203"></a>

### Direct properties for `https.coalescing_options`

- [default_coalescing](data-sources--http_loadbalancer--reference--group-018.md#canonical-1230031131322300-0203330331131212-1121001011020002-0130210100301302-0120110010120333-3220333103210232-2111032023332220-2300213123311220): complete subsection reference.

- [strict_coalescing](data-sources--http_loadbalancer--reference--group-018.md#canonical-2310100133023231-2213123132321011-2120323001032322-1320201323220323-2302223031032123-2022333320330330-3030020322303211-2010112320230212): complete subsection reference.

<a id="canonical-1230031131322300-0203330331131212-1121001011020002-0130210100301302-0120110010120333-3220333103210232-2111032023332220-2300213123311220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.coalescing_options.default_coalescing` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.coalescing_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-1001220203312330-3031013021133302-3130013202123010-1232020313310122-0213332300123110-3132233211103221-0220011120321111-3220032023002332)
- https.coalescing_options.default_coalescing

<a id="canonical-3320213130112310-2233113110331030-2133310210020131-1013021112113203-1133323000230020-3110201222201023-2010112213201332-3313233333200303"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default coalescing.

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

<a id="canonical-2310100133023231-2213123132321011-2120323001032322-1320201323220323-2302223031032123-2022333320330330-3030020322303211-2010112320230212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.coalescing_options.strict_coalescing` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.coalescing_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-1001220203312330-3031013021133302-3130013202123010-1232020313310122-0213332300123110-3132233211103221-0220011120321111-3220032023002332)
- https.coalescing_options.strict_coalescing

<a id="canonical-1001220333201210-2133213300202022-0101311012010211-0032232201131330-3320021033333111-1002111013011331-2122302023133230-2313020200302102"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for strict coalescing.

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

<a id="canonical-1011220220213223-0023022311131111-3321100321212103-2302012100332100-3310212111103331-2030012330210201-1100032222211300-2333103131312233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.default_header` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- https.default_header

<a id="canonical-0222103132323210-3210011212101022-2002102233121002-2122312321202133-3123212120210232-0221212103021010-1200311301000011-2122213002323313"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default header.

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

<a id="canonical-1123003233230101-3321333100231312-1212220312212102-3022231131212132-1132223102003323-3030022222021003-1030200021021213-1030221101223101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.default_loadbalancer` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- https.default_loadbalancer

<a id="canonical-2331031330211323-1202303110133100-2223313211012211-1313312200332113-1131210121120303-3332101000222033-2000131211313320-3201101213313133"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default loadbalancer.

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

<a id="canonical-1313230312203121-0221130100202030-2130030031311110-0000203003102103-0330223030021321-3212202011001322-2323331303122122-3302211302001001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.disable_path_normalize` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- https.disable_path_normalize

<a id="canonical-0012123310211301-2122213003032100-3212100223103311-1032223102013203-3311123320222111-3031301021130103-2303203020001233-3221333103301212"></a>

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

<a id="canonical-3211010213301321-0310002022003033-2320121222112212-3322033031120213-3201231212313320-2233231203130122-3211232302302121-2012120022302233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.enable_path_normalize` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- https.enable_path_normalize

<a id="canonical-1003023302220222-2302331233003301-2113022230130101-3221320232300130-3120320301100100-0133330010321321-0333030122200130-2032020213320331"></a>

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

<a id="canonical-0003020203103000-2333100023323311-1111213021211113-2303201112013303-2123231030312110-3202013021320132-3330331220213021-1110230322200021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.http_protocol_options` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- https.http_protocol_options

<a id="canonical-0331333210213110-1000010230021220-2202131230100023-2311311131111133-0332120100223232-2232233202201301-0032303213122333-2313333310030013"></a>

Type: `"single"`. Computed.

HTTP protocol configuration OPTIONS for downstream connections.

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

<a id="canonical-1231203331312002-1212120301222301-0030312113033121-1122303113212113-2113131033312022-3212200231333000-2301222203010320-3133023112133121"></a>

### Direct properties for `https.http_protocol_options`

- [http_protocol_enable_v1_only](data-sources--http_loadbalancer--reference--group-018.md#canonical-2131231302230011-3310322232331211-2030001330101211-0203321332010113-1210021203131212-2233023230202032-0200233320020102-0220312013203322): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--http_loadbalancer--reference--group-018.md#canonical-3123333211101233-2011211330112003-1122020031211202-0332221113321002-0200002300022131-3003033110131232-0230101212213231-1221232312332022): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--http_loadbalancer--reference--group-018.md#canonical-1130003130113323-3323300300133313-1002312300021013-1023122000132201-0231101201200133-2331101003201001-3220011220132332-2211322332112133): complete subsection reference.

<a id="canonical-2131231302230011-3310322232331211-2030001330101211-0203321332010113-1210021203131212-2233023230202032-0200233320020102-0220312013203322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.http_protocol_options.http_protocol_enable_v1_only` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-0003020203103000-2333100023323311-1111213021211113-2303201112013303-2123231030312110-3202013021320132-3330331220213021-1110230322200021)
- https.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-0021131313021203-1020002021110101-0112313000001311-0000233120133200-3331210103012213-1111013323130311-2301310230230122-1132211312003103"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3021123301203210-1033100131200121-2310203130002121-1012203211131101-3231212322301111-3110120111023301-2030301201021021-0333112011232020"></a>

### Direct properties for `https.http_protocol_options.http_protocol_enable_v1_only`

- [header_transformation](data-sources--http_loadbalancer--reference--group-018.md#canonical-0310021231311132-0310332010021331-3020123330001232-2022001313200222-2220233211302110-3230100022101233-0212202212133201-0023221100111030): complete subsection reference.

<a id="canonical-0310021231311132-0310332010021331-3020123330001232-2022001313200222-2220233211302110-3230100022101233-0212202212133201-0023221100111030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.http_protocol_options.http_protocol_enable_v1_only.header_transformation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-0003020203103000-2333100023323311-1111213021211113-2303201112013303-2123231030312110-3202013021320132-3330331220213021-1110230322200021)
- [https.http_protocol_options.http_protocol_enable_v1_only](data-sources--http_loadbalancer--reference--group-018.md#canonical-2131231302230011-3310322232331211-2030001330101211-0203321332010113-1210021203131212-2233023230202032-0200233320020102-0220312013203322)
- https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-0011103212003100-3122001323220110-0223303311021233-3131203013311210-0301031231310020-0032221130311131-3101300132321221-0100332230020232"></a>

Type: `"single"`. Computed.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

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

<a id="canonical-2113121113100030-2031222201311332-3103000231331123-0202221013131133-0031133033030110-0113030213031313-1303123221033333-2231103102300003"></a>

### Direct properties for `https.http_protocol_options.http_protocol_enable_v1_only.header_transformation`

- [default_header_transformation](data-sources--http_loadbalancer--reference--group-018.md#canonical-0212231112231303-3110332202132222-2112001213020312-0022110013323210-3111223000101213-3333210101120311-3131102223033011-0110130330303311): complete subsection reference.

- [preserve_case_header_transformation](data-sources--http_loadbalancer--reference--group-018.md#canonical-0101031213333122-1122323003122010-1212011021122202-3132110001231312-1022210133303112-2212213133001130-1112133213030021-0120100133212023): complete subsection reference.

- [proper_case_header_transformation](data-sources--http_loadbalancer--reference--group-018.md#canonical-1131012011232031-3011220131031031-2223313233001213-0020010012103312-2021303331223301-3111133222023212-3203231323332010-0212230320313133): complete subsection reference.

<a id="canonical-0212231112231303-3110332202132222-2112001213020312-0022110013323210-3111223000101213-3333210101120311-3131102223033011-0110130330303311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-0003020203103000-2333100023323311-1111213021211113-2303201112013303-2123231030312110-3202013021320132-3330331220213021-1110230322200021)
- [https.http_protocol_options.http_protocol_enable_v1_only](data-sources--http_loadbalancer--reference--group-018.md#canonical-2131231302230011-3310322232331211-2030001330101211-0203321332010113-1210021203131212-2233023230202032-0200233320020102-0220312013203322)
- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--http_loadbalancer--reference--group-018.md#canonical-0310021231311132-0310332010021331-3020123330001232-2022001313200222-2220233211302110-3230100022101233-0212202212133201-0023221100111030)
- https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-0223213122123031-1203331030133110-2033000132112012-1221333130201233-0123101221000000-1110312332122223-1220101121021311-2022331122122022"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0101031213333122-1122323003122010-1212011021122202-3132110001231312-1022210133303112-2212213133001130-1112133213030021-0120100133212023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-0003020203103000-2333100023323311-1111213021211113-2303201112013303-2123231030312110-3202013021320132-3330331220213021-1110230322200021)
- [https.http_protocol_options.http_protocol_enable_v1_only](data-sources--http_loadbalancer--reference--group-018.md#canonical-2131231302230011-3310322232331211-2030001330101211-0203321332010113-1210021203131212-2233023230202032-0200233320020102-0220312013203322)
- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--http_loadbalancer--reference--group-018.md#canonical-0310021231311132-0310332010021331-3020123330001232-2022001313200222-2220233211302110-3230100022101233-0212202212133201-0023221100111030)
- https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-0122002100132211-1000033221323131-2130201033330311-3121302110030131-0033003312013310-1033332301000031-0112110312201323-1213122321322311"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1131012011232031-3011220131031031-2223313233001213-0020010012103312-2021303331223301-3111133222023212-3203231323332010-0212230320313133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-0003020203103000-2333100023323311-1111213021211113-2303201112013303-2123231030312110-3202013021320132-3330331220213021-1110230322200021)
- [https.http_protocol_options.http_protocol_enable_v1_only](data-sources--http_loadbalancer--reference--group-018.md#canonical-2131231302230011-3310322232331211-2030001330101211-0203321332010113-1210021203131212-2233023230202032-0200233320020102-0220312013203322)
- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--http_loadbalancer--reference--group-018.md#canonical-0310021231311132-0310332010021331-3020123330001232-2022001313200222-2220233211302110-3230100022101233-0212202212133201-0023221100111030)
- https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-1132120303330223-2213302003231303-0102333312130121-3220111023230232-3123222322332333-1110033022032010-0222010002131203-1131022232012003"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3123333211101233-2011211330112003-1122020031211202-0332221113321002-0200002300022131-3003033110131232-0230101212213231-1221232312332022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.http_protocol_options.http_protocol_enable_v1_v2` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-0003020203103000-2333100023323311-1111213021211113-2303201112013303-2123231030312110-3202013021320132-3330331220213021-1110230322200021)
- https.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-2022311233313011-0323000202123222-2121021202001200-2000223030012133-2223031011323030-3023302233112231-1010333221300313-0100223032130103"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for http protocol enable v1 v2.

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

<a id="canonical-1130003130113323-3323300300133313-1002312300021013-1023122000132201-0231101201200133-2331101003201001-3220011220132332-2211322332112133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.http_protocol_options.http_protocol_enable_v2_only` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-0003020203103000-2333100023323311-1111213021211113-2303201112013303-2123231030312110-3202013021320132-3330331220213021-1110230322200021)
- https.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-2202101213113122-1101012000231012-2000033211033011-3201311103302232-3012220321002333-1113221002331313-3222231110332013-3231000112113332"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for http protocol enable v2 only.

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

<a id="canonical-1101330312012103-0233323210311221-1033012102103323-0220330330101233-1110011203231222-2231210133200221-1022100012023223-3112010202021011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.non_default_loadbalancer` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- https.non_default_loadbalancer

<a id="canonical-3022332322133201-3122111332221310-2303311112031200-3012332122103121-3330302003203013-0021310103100123-2330113103011121-3022110211012320"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for non default loadbalancer.

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

<a id="canonical-3210232200212003-0002112320100202-1003013010223122-1001223222303210-2322123332210310-3322112011113221-2311013202301022-0201323130200031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.pass_through` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- https.pass_through

<a id="canonical-1221302222302211-0011201000033021-0320010103220203-2110031312322200-3113222122102102-1220100030232100-0213212203210223-0230323001200133"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for pass through.

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

<a id="canonical-0011121033233220-3311202112112001-0220333111322122-2213202322313002-3213103232311233-2200102030312213-0032120231111113-3201230203322302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_params` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- https.tls_cert_params

<a id="canonical-1131030322122230-2321302021230200-1202302123312222-1123302122220220-0020220323233321-2013203301201210-1320212200231112-2023130102322202"></a>

Type: `"single"`. Computed.

Configuration parameter for tls cert params.

Additional upstream details:

Select TLS Parameters and Certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

<a id="canonical-2323012233332233-2300031101333001-2223210231012213-3101313023132323-1032000032120021-0312212021303021-3313201303330202-0311021001133012"></a>

### Direct properties for `https.tls_cert_params`

- [certificates](data-sources--http_loadbalancer--reference--group-018.md#canonical-2312203201012133-1132133332020110-0112332133100322-2212200200230021-0203033110100302-1301102101101330-3200320103310333-1123222313100333): complete subsection reference.

- [no_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-0223133102312020-3233203020002133-2303010022213201-3210211100122220-3221130303102112-2303220003302120-2023322301131321-1100122120021020): complete subsection reference.

- [tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-0330213002100010-3210123313013210-0111102220002103-1222200012322122-1323203130221323-1301013230233031-3110101122322333-0033031012123001): complete subsection reference.

- [use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-0100211003133012-2022100102310023-1213133030310333-1112311330321323-1113030000201010-3012132103130231-3110303333323213-0302113200332100): complete subsection reference.

<a id="canonical-2312203201012133-1132133332020110-0112332133100322-2212200200230021-0203033110100302-1301102101101330-3200320103310333-1123222313100333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_params.certificates` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011121033233220-3311202112112001-0220333111322122-2213202322313002-3213103232311233-2200102030312213-0032120231111113-3201230203322302)
- https.tls_cert_params.certificates

<a id="canonical-1120102101120123-2320123033022110-1023031020010332-3023332003031301-1303123310100213-2323131320201120-0023033312303221-2232233032203211"></a>

Type: `"list"`. Computed.

Select one or more certificates with any domain names.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2133030232220320-2202302311131022-1012221333301120-0123222322311213-3102333030331020-3332201023103300-0301033033101330-0220121301303021"></a>

### Direct properties for `https.tls_cert_params.certificates`

<a id="canonical-3231230013320110-0222020222103123-2310232020022100-2223022310330223-3220031103020012-0131313233223310-0200003033022220-3013303210132213"></a>

#### `https.tls_cert_params.certificates.name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3032321302131012-0202112021130221-0001112013212331-0311221212213313-1220300121010033-0031123320210321-0100003230223213-1322333031323203"></a>

<a id="canonical-3022330200232303-1331120011123000-3030333230011200-2110022113231322-2210200212023321-3203210101030100-1120313133130222-3201103230032123"></a>

#### `https.tls_cert_params.certificates.namespace` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1003000122023113-1322131222101211-1021321013232300-1310332101213003-2102320012232210-3121123031212112-1002011002232201-1021110222013102"></a>

<a id="canonical-1203233112120001-1100213233231120-0332021103303002-0110031010200102-0133001123221313-0202211123233331-2300030203132310-0132100033323111"></a>

#### `https.tls_cert_params.certificates.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0223133102312020-3233203020002133-2303010022213201-3210211100122220-3221130303102112-2303220003302120-2023322301131321-1100122120021020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_params.no_mtls` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011121033233220-3311202112112001-0220333111322122-2213202322313002-3213103232311233-2200102030312213-0032120231111113-3201230203322302)
- https.tls_cert_params.no_mtls

<a id="canonical-1212130101301300-1200311122200223-1102001223031000-1211100303323103-3300210313002102-3301121131021330-2222100133023331-0231210332011113"></a>

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

<a id="canonical-0330213002100010-3210123313013210-0111102220002103-1222200012322122-1323203130221323-1301013230233031-3110101122322333-0033031012123001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_params.tls_config` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011121033233220-3311202112112001-0220333111322122-2213202322313002-3213103232311233-2200102030312213-0032120231111113-3201230203322302)
- https.tls_cert_params.tls_config

<a id="canonical-3132332323332323-1330201111102230-2303222301010003-0110202022001022-3211310012032202-1022001320233031-1230021221313333-1022112121101110"></a>

Type: `"single"`. Computed.

This defines various OPTIONS to configure TLS configuration parameters.

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

<a id="canonical-0300222211131012-2300303012133011-1331233103012313-1300311301322020-2133313011333330-0130203133322321-3020321111302212-0020303021302103"></a>

### Direct properties for `https.tls_cert_params.tls_config`

- [custom_security](data-sources--http_loadbalancer--reference--group-018.md#canonical-2320123123023013-2210220122302000-0310020301201310-3312122223321030-2321012331122220-1332231220120032-2133031212023320-0230312002103003): complete subsection reference.

- [default_security](data-sources--http_loadbalancer--reference--group-018.md#canonical-1312123121000103-2131020023213321-0200320131030212-1301323003312120-3201030013320331-0321113330013220-1220223102010122-0033300221233103): complete subsection reference.

- [low_security](data-sources--http_loadbalancer--reference--group-018.md#canonical-0002220110031322-2000022330002023-1210033032322201-2222102300031303-3102020301103131-2121013132310021-0212213012103132-2201330202321122): complete subsection reference.

- [medium_security](data-sources--http_loadbalancer--reference--group-018.md#canonical-3203032113231011-0312131120020200-2200310022201122-0130021221021222-3312322001011302-1231211323221220-3320013323033222-1221033013130332): complete subsection reference.

<a id="canonical-2320123123023013-2210220122302000-0310020301201310-3312122223321030-2321012331122220-1332231220120032-2133031212023320-0230312002103003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_params.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011121033233220-3311202112112001-0220333111322122-2213202322313002-3213103232311233-2200102030312213-0032120231111113-3201230203322302)
- [https.tls_cert_params.tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-0330213002100010-3210123313013210-0111102220002103-1222200012322122-1323203130221323-1301013230233031-3110101122322333-0033031012123001)
- https.tls_cert_params.tls_config.custom_security

<a id="canonical-2132132210332122-1303031131313203-3120222101101220-3101001213332003-1113301010310222-2023211301020210-1101303131233133-2202101220133110"></a>

Type: `"single"`. Computed.

This defines TLS protocol config including min/max versions and allowed ciphers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1123301210212102-3321310130111110-2320301101113030-0110221133033131-0212120300133200-0111021030103123-3030300302221200-1010322100333221"></a>

### Direct properties for `https.tls_cert_params.tls_config.custom_security`

<a id="canonical-3111000330222101-1000023332013311-0232331332131230-2122331232111200-1020101230231002-1232230122030301-2320313100121320-1002110212003223"></a>

#### `https.tls_cert_params.tls_config.custom_security.cipher_suites` property

Type: `["list", "string"]`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1222131331031232-3010221200333230-0001302300233313-3020230203003100-1020300312120120-2123122313110223-3212230311200120-3002210310321221"></a>

<a id="canonical-2332221013030221-2103132222220330-0002233201111210-3323021020231013-0231020023100203-2122300330112031-1031331312230000-1020313111311130"></a>

#### `https.tls_cert_params.tls_config.custom_security.max_version` property

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

<a id="canonical-2221113133121120-1020032000103101-2033232003331220-1320201031211001-1121000320121300-2013232120123322-1122212322310233-1201110300023032"></a>

<a id="canonical-0130313110323103-3213220323132230-2230221011231312-2003100130333120-1330122033130302-2310030231110130-2220302022200133-1332323210111113"></a>

#### `https.tls_cert_params.tls_config.custom_security.min_version` property

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

<a id="canonical-1312123121000103-2131020023213321-0200320131030212-1301323003312120-3201030013320331-0321113330013220-1220223102010122-0033300221233103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_params.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011121033233220-3311202112112001-0220333111322122-2213202322313002-3213103232311233-2200102030312213-0032120231111113-3201230203322302)
- [https.tls_cert_params.tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-0330213002100010-3210123313013210-0111102220002103-1222200012322122-1323203130221323-1301013230233031-3110101122322333-0033031012123001)
- https.tls_cert_params.tls_config.default_security

<a id="canonical-3322200322120202-2200000121000232-3311322221132002-2121313103112033-0301332000200222-2103013300201000-3100000303211011-1020330032231131"></a>

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

<a id="canonical-0002220110031322-2000022330002023-1210033032322201-2222102300031303-3102020301103131-2121013132310021-0212213012103132-2201330202321122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_params.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011121033233220-3311202112112001-0220333111322122-2213202322313002-3213103232311233-2200102030312213-0032120231111113-3201230203322302)
- [https.tls_cert_params.tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-0330213002100010-3210123313013210-0111102220002103-1222200012322122-1323203130221323-1301013230233031-3110101122322333-0033031012123001)
- https.tls_cert_params.tls_config.low_security

<a id="canonical-0111210212313120-2012301312113313-3312310033200313-3200010002302220-3311303331131321-0013201220321203-1203103020221031-0223032022030323"></a>

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

<a id="canonical-3203032113231011-0312131120020200-2200310022201122-0130021221021222-3312322001011302-1231211323221220-3320013323033222-1221033013130332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_params.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011121033233220-3311202112112001-0220333111322122-2213202322313002-3213103232311233-2200102030312213-0032120231111113-3201230203322302)
- [https.tls_cert_params.tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-0330213002100010-3210123313013210-0111102220002103-1222200012322122-1323203130221323-1301013230233031-3110101122322333-0033031012123001)
- https.tls_cert_params.tls_config.medium_security

<a id="canonical-0333033101123201-0020322121030321-3332133321200122-3200322112101032-1210322031220320-0223221201202321-1213001202023110-3212110202300231"></a>

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

<a id="canonical-0100211003133012-2022100102310023-1213133030310333-1112311330321323-1113030000201010-3012132103130231-3110303333323213-0302113200332100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_params.use_mtls` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011121033233220-3311202112112001-0220333111322122-2213202322313002-3213103232311233-2200102030312213-0032120231111113-3201230203322302)
- https.tls_cert_params.use_mtls

<a id="canonical-1001033312230212-1211032321133220-3111300223331030-0212320310232012-0130220033221203-2211122032330020-2201201310310132-2210222321003323"></a>

Type: `"single"`. Computed.

Validation context for downstream client TLS connections.

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

<a id="canonical-3121100212103131-0022233101312033-3110312130103032-2130332020112121-1132203130011312-1113103221330122-1113332021322103-3110230031301123"></a>

### Direct properties for `https.tls_cert_params.use_mtls`

<a id="canonical-1300211232332133-1310201331032020-0121311300011030-2220311321300030-0013120202110000-2301310102121301-0033033121332313-0011223112122111"></a>

#### `https.tls_cert_params.use_mtls.client_certificate_optional` property

Type: `"bool"`. Computed.

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

- [crl](data-sources--http_loadbalancer--reference--group-018.md#canonical-2132103301102021-1132103122101230-3123022210011220-0320030213232331-2311321203000203-0221302121132220-1123231211020323-1032331010322322): complete subsection reference.

- [no_crl](data-sources--http_loadbalancer--reference--group-018.md#canonical-3213132300132122-1300233220030300-2200001200000201-3223021231022201-1031301301232330-0011113121023200-3202230312323002-0300101313321032): complete subsection reference.

- [trusted_ca](data-sources--http_loadbalancer--reference--group-018.md#canonical-1221321023303223-0313120023012131-3100323220013221-2112001021000211-2302311022133303-2200303100331113-3200311001223200-1323113213332130): complete subsection reference.

<a id="canonical-3320211123102210-0111332210001123-0231300303110011-2121031020201210-2013110210213002-2223222113022022-1231030033333330-1131230100130301"></a>

<a id="canonical-2220202300222111-1300013211223212-3122312221210233-1210331332022201-3233200231133333-1022300232031223-3011120120330202-2322331011231112"></a>

#### `https.tls_cert_params.use_mtls.trusted_ca_url` property

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [xfcc_disabled](data-sources--http_loadbalancer--reference--group-018.md#canonical-0101113332133200-2011200000230031-0032201320100301-2333000103132121-0121221113003010-3220211212032320-1023031111012212-1113123323123003): complete subsection reference.

- [xfcc_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-3013120300333313-2132201100310030-1310220122232030-3302131210133100-0133030331022131-2300210103021313-2322311123331223-0123230011202223): complete subsection reference.

<a id="canonical-2132103301102021-1132103122101230-3123022210011220-0320030213232331-2311321203000203-0221302121132220-1123231211020323-1032331010322322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_params.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011121033233220-3311202112112001-0220333111322122-2213202322313002-3213103232311233-2200102030312213-0032120231111113-3201230203322302)
- [https.tls_cert_params.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-0100211003133012-2022100102310023-1213133030310333-1112311330321323-1113030000201010-3012132103130231-3110303333323213-0302113200332100)
- https.tls_cert_params.use_mtls.crl

<a id="canonical-1210222203113202-0333012120220033-3030323113120011-1133111213321133-2001332000210130-3220311012203101-3100201010200121-0020101133001300"></a>

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

<a id="canonical-2212313103021310-0332233220021010-1230311300033200-0012110320022230-2201011331212122-0221131003021303-1112002203332111-2032222321213313"></a>

### Direct properties for `https.tls_cert_params.use_mtls.crl`

<a id="canonical-2122203201320103-1211231210303023-3313330010321213-0331002032300330-1101321222331320-2000210021102200-0211133003131131-0001200302001133"></a>

#### `https.tls_cert_params.use_mtls.crl.name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2303130111210311-1113313301301223-0310130231201121-1120311102131120-1231313232023102-0301212102032333-0031320130110030-1302322313331313"></a>

<a id="canonical-0210122031221012-1133001010213101-3010110301002232-2232223220211230-2021122333113000-1223300002133313-3200120121301313-3022100020310311"></a>

#### `https.tls_cert_params.use_mtls.crl.namespace` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3332333011331301-2203022320030220-2231132310013103-3023112023122010-3323111233213030-0121303200310220-3323122033000023-1131231203033012"></a>

<a id="canonical-1101033020301022-1323030021132012-3330023021330302-2223032213010101-2032320102022011-0233020012130320-0113202121310221-1301321232102112"></a>

#### `https.tls_cert_params.use_mtls.crl.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3213132300132122-1300233220030300-2200001200000201-3223021231022201-1031301301232330-0011113121023200-3202230312323002-0300101313321032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_params.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011121033233220-3311202112112001-0220333111322122-2213202322313002-3213103232311233-2200102030312213-0032120231111113-3201230203322302)
- [https.tls_cert_params.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-0100211003133012-2022100102310023-1213133030310333-1112311330321323-1113030000201010-3012132103130231-3110303333323213-0302113200332100)
- https.tls_cert_params.use_mtls.no_crl

<a id="canonical-1230323110030311-2121301101130301-0301130221013000-0212302022123022-1300202203001002-3002000211033201-2213131321121310-1120103211030031"></a>

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

<a id="canonical-1221321023303223-0313120023012131-3100323220013221-2112001021000211-2302311022133303-2200303100331113-3200311001223200-1323113213332130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_params.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011121033233220-3311202112112001-0220333111322122-2213202322313002-3213103232311233-2200102030312213-0032120231111113-3201230203322302)
- [https.tls_cert_params.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-0100211003133012-2022100102310023-1213133030310333-1112311330321323-1113030000201010-3012132103130231-3110303333323213-0302113200332100)
- https.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-0122021112200112-1111230123132222-0113022013033320-1232321113330213-2312210210023121-2212013113021130-0012232310021023-3111221112233201"></a>

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

<a id="canonical-0223213022203033-3123303003101000-0013021012300130-3032131123201233-0001322030201321-2033232112023003-1012222321000111-0132121221201321"></a>

### Direct properties for `https.tls_cert_params.use_mtls.trusted_ca`

<a id="canonical-0232301222302201-2131033302202020-2113011101002332-0101012022123321-3121220033213000-1303200130300110-3103002211223112-3320023210002312"></a>

#### `https.tls_cert_params.use_mtls.trusted_ca.name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0101330022303022-2213311001331012-2322000333101230-1010120003000222-2031232111101200-2311332210102221-3123023122111133-0301233003020221"></a>

<a id="canonical-3001203131301231-1332110222313031-3323131320202103-1200032331110010-2000121002222020-2132021303031112-1020321311210021-1231032221021020"></a>

#### `https.tls_cert_params.use_mtls.trusted_ca.namespace` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3223132220333130-3133220000311002-2332220013020020-1321302023020331-3223200231211223-2303222312023303-3201210331210210-2122321001122323"></a>

<a id="canonical-0132231330203223-1222113032202013-3200122232211311-3010103000300322-2112000220013210-0311111013333322-2012323013121002-2120100130122230"></a>

#### `https.tls_cert_params.use_mtls.trusted_ca.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0101113332133200-2011200000230031-0032201320100301-2333000103132121-0121221113003010-3220211212032320-1023031111012212-1113123323123003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_params.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011121033233220-3311202112112001-0220333111322122-2213202322313002-3213103232311233-2200102030312213-0032120231111113-3201230203322302)
- [https.tls_cert_params.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-0100211003133012-2022100102310023-1213133030310333-1112311330321323-1113030000201010-3012132103130231-3110303333323213-0302113200332100)
- https.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-2322010031032131-3312013130331212-0102311232120110-2330300003210030-1020233133212212-0100102011213000-0032223113221113-0230012220113302"></a>

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

<a id="canonical-3013120300333313-2132201100310030-1310220122232030-3302131210133100-0133030331022131-2300210103021313-2322311123331223-0123230011202223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_cert_params.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011121033233220-3311202112112001-0220333111322122-2213202322313002-3213103232311233-2200102030312213-0032120231111113-3201230203322302)
- [https.tls_cert_params.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-0100211003133012-2022100102310023-1213133030310333-1112311330321323-1113030000201010-3012132103130231-3110303333323213-0302113200332100)
- https.tls_cert_params.use_mtls.xfcc_options

<a id="canonical-2130001213013112-0103123210133333-3113021131112120-0212312202213003-3230133212032311-3331302202311213-0011222010101203-1302111030201321"></a>

Type: `"single"`. Computed.

X-Forwarded-Client-Cert header elements to be added to requests.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3222333212111012-3130112220010301-2302231001311132-3000023000212313-3021200223010322-2201121031012113-3103010200200100-3211133332233212"></a>

### Direct properties for `https.tls_cert_params.use_mtls.xfcc_options`

<a id="canonical-3333310231313021-0311013111032003-2033201100230010-1010022102230101-2032032112122302-2213132201001211-1100303130313101-3212111013312112"></a>

#### `https.tls_cert_params.use_mtls.xfcc_options.xfcc_header_elements` property

Type: `["list", "string"]`. Computed.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

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

<a id="canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- https.tls_parameters

<a id="canonical-0020301130102023-0302311331221000-2131120320011101-3101303223201033-1102232300211302-3020101122010120-1203012131312021-1133211332021303"></a>

Type: `"single"`. Computed.

Configuration parameter for tls parameters.

Additional upstream details:

Inline TLS parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

<a id="canonical-1332212010321010-2212110033131020-1313310210002302-1222120102120133-0131311202303123-2320202303310120-0201112302302112-1022333300313102"></a>

### Direct properties for `https.tls_parameters`

- [no_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-3201121212312000-1320103233230333-2023311223003113-0102232030012220-3101110302231321-0222013011220221-3130301332332213-3101023121210333): complete subsection reference.

- [tls_certificates](data-sources--http_loadbalancer--reference--group-018.md#canonical-1233120012132312-2323002313212002-0230233201032230-0030031113003030-0223022002211100-0030021312120301-3023220210001332-3021131311022323): complete subsection reference.

- [tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-1012111220110210-2102012230102300-3223123311002312-3303231133232111-0210100303201310-2213213133030233-1220200110120130-1010030123121230): complete subsection reference.

- [use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-3011011302300202-0330111321313003-2330103220210211-3110001210211122-3003210232200230-1123321222021031-0311313312302021-2121223323012232): complete subsection reference.

<a id="canonical-3201121212312000-1320103233230333-2023311223003113-0102232030012220-3101110302231321-0222013011220221-3130301332332213-3101023121210333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.no_mtls` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- https.tls_parameters.no_mtls

<a id="canonical-1133013132301033-0012221120320300-0131120212231022-0303100011302011-2033203022312021-3003011020332033-0000211221111010-1002332200000322"></a>

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

<a id="canonical-1233120012132312-2323002313212002-0230233201032230-0030031113003030-0223022002211100-0030021312120301-3023220210001332-3021131311022323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.tls_certificates` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- https.tls_parameters.tls_certificates

<a id="canonical-2331300233121013-2130300120030312-2333111120021031-3130330120200001-0233031132111220-3122203202120103-1123312001110310-2023111302123302"></a>

Type: `"list"`. Computed.

Users can add one or more certificates that share the same set of domains. For example, domain.com
and \*.domain.com - but use different signature algorithms.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-2303102323022111-3100013000310320-2002301302232110-2320022130313020-2112113232311110-2213020000123231-1100211101023130-0032220222001020"></a>

### Direct properties for `https.tls_parameters.tls_certificates`

<a id="canonical-0321232132310301-0132122102303120-2011302233130030-0133310100123201-2211111020003030-0230030221210201-0132331031010201-2210203031301230"></a>

#### `https.tls_parameters.tls_certificates.certificate_url` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [custom_hash_algorithms](data-sources--http_loadbalancer--reference--group-018.md#canonical-2130310001100133-1220032202322233-3023101300330010-2022213021332233-0331030322031202-0303103112320000-0111232112030131-1200220111311232): complete subsection reference.

<a id="canonical-3333203201111120-2103122200230212-0032130122301331-1212031122331022-0102211233031120-1201310332120101-3322202220211013-0231303121103310"></a>

<a id="canonical-1123330113103223-0320300230100231-1210322130002302-2210111212301201-1330002121003332-1003213300122320-1220201231003212-2111321120122331"></a>

#### `https.tls_parameters.tls_certificates.description_spec` property

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--http_loadbalancer--reference--group-018.md#canonical-3121301022103202-0020210230210201-1233201302323220-3031133020213323-0212102312212311-1010311200212101-0113231212231020-3001233332313222): complete subsection reference.

- [private_key](data-sources--http_loadbalancer--reference--group-018.md#canonical-0031331322111222-1031332100211301-2333330101002212-2330332000301300-3320220000031210-3210121222110210-1113321203310002-2022002202313003): complete subsection reference.

- [use_system_defaults](data-sources--http_loadbalancer--reference--group-018.md#canonical-1001202002030212-2020213322121313-3222120131122030-2101300003120030-1023100100102232-2023123212311020-2132011210001120-2312203113300010): complete subsection reference.

<a id="canonical-2130310001100133-1220032202322233-3023101300330010-2022213021332233-0331030322031202-0303103112320000-0111232112030131-1200220111311232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- [https.tls_parameters.tls_certificates](data-sources--http_loadbalancer--reference--group-018.md#canonical-1233120012132312-2323002313212002-0230233201032230-0030031113003030-0223022002211100-0030021312120301-3023220210001332-3021131311022323)
- https.tls_parameters.tls_certificates.custom_hash_algorithms

<a id="canonical-2301012321323123-1121333312221123-2201003201220203-3223121233031012-0102120113023023-3312101233322332-0100122131101132-2020011131320212"></a>

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

<a id="canonical-2210122031100130-1231300211220031-1332333111331201-0131222201132212-3032321232001233-1112100101010230-2010211111302323-2013302231130323"></a>

### Direct properties for `https.tls_parameters.tls_certificates.custom_hash_algorithms`

<a id="canonical-3021113100122302-0000130210020110-2001230311233233-2031303210122323-0330113132133121-2133302003012322-3200123132100203-1002100112120310"></a>

#### `https.tls_parameters.tls_certificates.custom_hash_algorithms.hash_algorithms` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3121301022103202-0020210230210201-1233201302323220-3031133020213323-0212102312212311-1010311200212101-0113231212231020-3001233332313222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- [https.tls_parameters.tls_certificates](data-sources--http_loadbalancer--reference--group-018.md#canonical-1233120012132312-2323002313212002-0230233201032230-0030031113003030-0223022002211100-0030021312120301-3023220210001332-3021131311022323)
- https.tls_parameters.tls_certificates.disable_ocsp_stapling

<a id="canonical-2303321200202001-1230303202333321-2113100211223203-3112101022220033-0022331201102130-1031012120330233-1011321112313213-1223300113201033"></a>

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

<a id="canonical-0031331322111222-1031332100211301-2333330101002212-2330332000301300-3320220000031210-3210121222110210-1113321203310002-2022002202313003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.tls_certificates.private_key` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- [https.tls_parameters.tls_certificates](data-sources--http_loadbalancer--reference--group-018.md#canonical-1233120012132312-2323002313212002-0230233201032230-0030031113003030-0223022002211100-0030021312120301-3023220210001332-3021131311022323)
- https.tls_parameters.tls_certificates.private_key

<a id="canonical-2313001132021230-2120201201202320-0203233211013102-3233102301021310-2211303202223330-2320100122013221-1101230222221112-2303123102122020"></a>

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

<a id="canonical-2331132313230331-1102331311302330-0103232030033013-2001213201300130-3231302210220311-3103021120120122-0120200133032021-2223020113111320"></a>

### Direct properties for `https.tls_parameters.tls_certificates.private_key`

- [blindfold_secret_info](data-sources--http_loadbalancer--reference--group-018.md#canonical-0101301211321131-1013101131023322-2201320033212201-3013133012210322-0103312012111132-2323212132131300-2011301221320220-1130300332223230): complete subsection reference.

- [clear_secret_info](data-sources--http_loadbalancer--reference--group-018.md#canonical-3200220310231302-0123030020201003-2102112011211023-1033132310020320-2300210300201101-0300312310233330-2220302223332121-2231230002232300): complete subsection reference.

<a id="canonical-0101301211321131-1013101131023322-2201320033212201-3013133012210322-0103312012111132-2323212132131300-2011301221320220-1130300332223230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.tls_certificates.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- [https.tls_parameters.tls_certificates](data-sources--http_loadbalancer--reference--group-018.md#canonical-1233120012132312-2323002313212002-0230233201032230-0030031113003030-0223022002211100-0030021312120301-3023220210001332-3021131311022323)
- [https.tls_parameters.tls_certificates.private_key](data-sources--http_loadbalancer--reference--group-018.md#canonical-0031331322111222-1031332100211301-2333330101002212-2330332000301300-3320220000031210-3210121222110210-1113321203310002-2022002202313003)
- https.tls_parameters.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-0202223010313203-1203021101200010-2021312021001231-0013320121302101-0012001112210030-2122001132231303-2123000103313302-3333320121001210"></a>

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

<a id="canonical-2302000113220201-2100311222220332-3313222033102020-1022132122212013-3100102231320310-2103023220312030-1123323030121031-3320230120302002"></a>

### Direct properties for `https.tls_parameters.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-1310301213302121-3001011021200303-0211031212222331-1001333133001001-0330231230322120-0021102000023001-2232321121033013-2010031222300333"></a>

#### `https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1010130230003110-1213302313020300-3130301233203002-1223110121333313-1020013331212011-0011002320223021-2331203322302033-1323301022003001"></a>

<a id="canonical-2323001001132123-3023211122202303-3013333232231130-2201130321030003-3313323203101233-1202210321203122-0212202121030200-2131120120211303"></a>

#### `https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.location` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3122001032220323-2233312033213122-1132030232232332-2030321223030112-1133121100212221-0232221002021303-1131010310021312-0013022000312313"></a>

<a id="canonical-2313010330222032-0030211031213111-3111120221301202-3211123011132312-3211230220021111-0322330211311123-2120331220122332-1333303131130302"></a>

#### `https.tls_parameters.tls_certificates.private_key.blindfold_secret_info.store_provider` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3200220310231302-0123030020201003-2102112011211023-1033132310020320-2300210300201101-0300312310233330-2220302223332121-2231230002232300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.tls_certificates.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- [https.tls_parameters.tls_certificates](data-sources--http_loadbalancer--reference--group-018.md#canonical-1233120012132312-2323002313212002-0230233201032230-0030031113003030-0223022002211100-0030021312120301-3023220210001332-3021131311022323)
- [https.tls_parameters.tls_certificates.private_key](data-sources--http_loadbalancer--reference--group-018.md#canonical-0031331322111222-1031332100211301-2333330101002212-2330332000301300-3320220000031210-3210121222110210-1113321203310002-2022002202313003)
- https.tls_parameters.tls_certificates.private_key.clear_secret_info

<a id="canonical-3003000001010212-0002002222301301-2211311001232023-1100313232000200-1100201101210121-2010133202031011-2211110202033131-0303011213002203"></a>

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

<a id="canonical-3131202201213223-0312121122102020-0220112231022230-2213131232212232-2130011102010222-2320003202013011-3330221310132012-0331312211100020"></a>

### Direct properties for `https.tls_parameters.tls_certificates.private_key.clear_secret_info`

<a id="canonical-0113313123013200-3010030103302103-3330323213022321-1323131032032201-2320021201221030-1321321111120311-1220133010101303-0031133103323112"></a>

#### `https.tls_parameters.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2111323220313311-2323121011113120-2000120101111332-2323121311230330-2331011013332032-2003011103210031-2230002223322130-0113000001002221"></a>

<a id="canonical-0200020102023212-1221222033220221-2312123100323023-1221100200030222-0213011013300320-1321121101302003-1310021100023213-2123300103100220"></a>

#### `https.tls_parameters.tls_certificates.private_key.clear_secret_info.url` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1001202002030212-2020213322121313-3222120131122030-2101300003120030-1023100100102232-2023123212311020-2132011210001120-2312203113300010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- [https.tls_parameters.tls_certificates](data-sources--http_loadbalancer--reference--group-018.md#canonical-1233120012132312-2323002313212002-0230233201032230-0030031113003030-0223022002211100-0030021312120301-3023220210001332-3021131311022323)
- https.tls_parameters.tls_certificates.use_system_defaults

<a id="canonical-2120222210031223-3021231123312333-0201330331321103-2213003223303003-3100120322310221-3232131301313012-1103010222231200-3233013132133101"></a>

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

<a id="canonical-1012111220110210-2102012230102300-3223123311002312-3303231133232111-0210100303201310-2213213133030233-1220200110120130-1010030123121230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.tls_config` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- https.tls_parameters.tls_config

<a id="canonical-2220011313112200-2112232331122113-1131221032002221-2113233011032321-3220133023133221-3213230311020232-0030101302320303-1333232302331113"></a>

Type: `"single"`. Computed.

This defines various OPTIONS to configure TLS configuration parameters.

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

<a id="canonical-3200011301021321-2030123203130011-1132302030213220-0122120231123312-3303203010203000-2332010200202122-0020112203020230-2322300302112122"></a>

### Direct properties for `https.tls_parameters.tls_config`

- [custom_security](data-sources--http_loadbalancer--reference--group-018.md#canonical-0120113003131211-1222320232220321-1103222033031223-1233013103213100-0121311111012012-1213331101001001-2101110031001332-0322012100203123): complete subsection reference.

- [default_security](data-sources--http_loadbalancer--reference--group-018.md#canonical-1301220222130220-0320332112003203-3001321233212300-0202203132012133-2333132211230130-2313200123031312-2211212302223221-1110103210022023): complete subsection reference.

- [low_security](data-sources--http_loadbalancer--reference--group-018.md#canonical-2100231223023213-2213001132101021-1332301103122232-3101232202021113-2223303200203321-2323102301200010-1200311003303100-1132213132011220): complete subsection reference.

- [medium_security](data-sources--http_loadbalancer--reference--group-018.md#canonical-1231131120130120-0312320030302212-2303001023032013-3323330022210210-3012121333203322-0200200002123120-3220001332232003-2203301212212311): complete subsection reference.

<a id="canonical-0120113003131211-1222320232220321-1103222033031223-1233013103213100-0121311111012012-1213331101001001-2101110031001332-0322012100203123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- [https.tls_parameters.tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-1012111220110210-2102012230102300-3223123311002312-3303231133232111-0210100303201310-2213213133030233-1220200110120130-1010030123121230)
- https.tls_parameters.tls_config.custom_security

<a id="canonical-3123113312122123-2323302111113131-3023320311121203-2031000301102222-2121221221122331-0023010233201012-0130012231101113-2122321302303202"></a>

Type: `"single"`. Computed.

This defines TLS protocol config including min/max versions and allowed ciphers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2031131331102011-0311002313201012-3030100131020133-0321003222222101-0011001221310323-1000113123331233-1331011013101001-2311221022233031"></a>

### Direct properties for `https.tls_parameters.tls_config.custom_security`

<a id="canonical-2022230313013222-0330033231121223-3233300001221212-3113312011301112-0322031213203113-0201101211032131-3322013123113301-1231202031210020"></a>

#### `https.tls_parameters.tls_config.custom_security.cipher_suites` property

Type: `["list", "string"]`. Computed.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0010321301201120-2313111323301212-1113322212013132-2003111333123001-0302013303010211-1231100233220101-1312311113021011-1100332323212032"></a>

<a id="canonical-1031013010003002-3120020111002103-1333200321101321-1120311023221211-1331020010211112-3021222003200021-0121303212302130-1221032312010031"></a>

#### `https.tls_parameters.tls_config.custom_security.max_version` property

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

<a id="canonical-1303213330302312-2231020120230100-0213210003012002-1203223022033000-2231303330012300-1130222311133330-1302033201122122-1022212130320221"></a>

<a id="canonical-3020023103320011-2033333003233132-3020002310110012-3101022320101330-2222233222101011-3120203221020101-3010012233011021-2010013220312310"></a>

#### `https.tls_parameters.tls_config.custom_security.min_version` property

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

<a id="canonical-1301220222130220-0320332112003203-3001321233212300-0202203132012133-2333132211230130-2313200123031312-2211212302223221-1110103210022023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- [https.tls_parameters.tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-1012111220110210-2102012230102300-3223123311002312-3303231133232111-0210100303201310-2213213133030233-1220200110120130-1010030123121230)
- https.tls_parameters.tls_config.default_security

<a id="canonical-1033320200100101-2103311100000032-3122030322010212-3030201332221302-3123301323021102-0103320211012313-1221301131230211-3110111320221031"></a>

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

<a id="canonical-2100231223023213-2213001132101021-1332301103122232-3101232202021113-2223303200203321-2323102301200010-1200311003303100-1132213132011220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- [https.tls_parameters.tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-1012111220110210-2102012230102300-3223123311002312-3303231133232111-0210100303201310-2213213133030233-1220200110120130-1010030123121230)
- https.tls_parameters.tls_config.low_security

<a id="canonical-2110311332123103-2222322021113002-0012112111021032-2133102131123302-2210002102123032-3321131232133131-0001130131121203-3001330230000032"></a>

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

<a id="canonical-1231131120130120-0312320030302212-2303001023032013-3323330022210210-3012121333203322-0200200002123120-3220001332232003-2203301212212311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- [https.tls_parameters.tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-1012111220110210-2102012230102300-3223123311002312-3303231133232111-0210100303201310-2213213133030233-1220200110120130-1010030123121230)
- https.tls_parameters.tls_config.medium_security

<a id="canonical-2212312301311031-0103111011022020-3000200100230032-2220012002212012-1203131100221032-3121010030122022-0311011000222310-2203203211020133"></a>

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

<a id="canonical-3011011302300202-0330111321313003-2330103220210211-3110001210211122-3003210232200230-1123321222021031-0311313312302021-2121223323012232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.use_mtls` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- https.tls_parameters.use_mtls

<a id="canonical-3322110332022200-3230111111213320-2220011031023301-1020312011321223-3321230122210123-0300333303233110-0330132033103022-3021132131312131"></a>

Type: `"single"`. Computed.

Validation context for downstream client TLS connections.

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

<a id="canonical-0132002212323321-1001133321012311-1202121120300330-0030311311131021-1231221110003303-2110113100022003-1011333333023120-2232213123032313"></a>

### Direct properties for `https.tls_parameters.use_mtls`

<a id="canonical-2133102121003001-0011032220122220-1121132220011223-0330321311033002-1120023012200203-3001302121321103-0132003323010313-1122320220002332"></a>

#### `https.tls_parameters.use_mtls.client_certificate_optional` property

Type: `"bool"`. Computed.

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

- [crl](data-sources--http_loadbalancer--reference--group-018.md#canonical-2120200310111303-2123002332032223-1320022103031131-2131313231021033-3223303003130213-2332200022300000-3300301003112020-3121000323201322): complete subsection reference.

- [no_crl](data-sources--http_loadbalancer--reference--group-018.md#canonical-3320113311333223-0331311201323111-3230310102122032-3110221022211232-3221133201031320-3221313011212200-0113023301232113-0213230313003003): complete subsection reference.

- [trusted_ca](data-sources--http_loadbalancer--reference--group-018.md#canonical-3130030222012122-3200001201311013-2320233200223213-2023020020303332-3100033030103312-3221133022132003-3311030330120211-3330220323230031): complete subsection reference.

<a id="canonical-0101003002023002-3233120221330120-2200002002013302-1103310210020312-0300321310112133-0211130302133323-1011233032302231-2131203122321003"></a>

<a id="canonical-2032201021323333-1200132220113000-3112200100202012-0131020120012100-2031300033210320-2133302032202212-0133120321333022-2301302330220121"></a>

#### `https.tls_parameters.use_mtls.trusted_ca_url` property

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [xfcc_disabled](data-sources--http_loadbalancer--reference--group-018.md#canonical-0211201201101012-2032030012112333-1023312213220032-3102332300203301-0211033020001201-3031133001311121-3033000210103130-2210310102021001): complete subsection reference.

- [xfcc_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-1322023323220103-0120111112033232-3301100122101000-1022023102030303-3333300212330203-2221222021232123-2032120200322113-2021012121033102): complete subsection reference.

<a id="canonical-2120200310111303-2123002332032223-1320022103031131-2131313231021033-3223303003130213-2332200022300000-3300301003112020-3121000323201322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.use_mtls.crl` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- [https.tls_parameters.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-3011011302300202-0330111321313003-2330103220210211-3110001210211122-3003210232200230-1123321222021031-0311313312302021-2121223323012232)
- https.tls_parameters.use_mtls.crl

<a id="canonical-1332020003312231-1221223330013003-2031200313212222-1113222203200011-1223122032331031-1321122102001233-2012120322120321-1011232323330132"></a>

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

<a id="canonical-0220120131101030-0211222310112301-3023233012022030-1230011211302332-2333211331032302-3300130231023210-1312123322002303-0131232320323330"></a>

### Direct properties for `https.tls_parameters.use_mtls.crl`

<a id="canonical-3331002021302301-0330322310201211-1311202303312013-0133220300202113-3130220300131031-2331221032032311-0222132032020321-0010023301021130"></a>

#### `https.tls_parameters.use_mtls.crl.name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0303312231121300-0110220233110320-2330230213030231-0122101022331330-3121001011300300-1100331000002212-1212111020020011-3023023302203200"></a>

<a id="canonical-3130002103312322-3131132130210303-3020212331201000-3301322122030210-1211023203333112-3003230212330122-1311102020133312-3111132221330100"></a>

#### `https.tls_parameters.use_mtls.crl.namespace` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2013331223121203-0132203020210001-0210122201122313-2322230223230030-1001301001310132-3100130030302013-1210312023331121-2102230131022133"></a>

<a id="canonical-0101103133320013-0233030320231122-2310011220210031-3300002332132302-1331321220223123-2030223131101332-2300301101003301-3010102322133310"></a>

#### `https.tls_parameters.use_mtls.crl.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3320113311333223-0331311201323111-3230310102122032-3110221022211232-3221133201031320-3221313011212200-0113023301232113-0213230313003003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.use_mtls.no_crl` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- [https.tls_parameters.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-3011011302300202-0330111321313003-2330103220210211-3110001210211122-3003210232200230-1123321222021031-0311313312302021-2121223323012232)
- https.tls_parameters.use_mtls.no_crl

<a id="canonical-3110223100130310-2032020003111330-3020130211313200-0312211213003311-1232311321302112-1222122032002010-1302300300133201-0233111230023213"></a>

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

<a id="canonical-3130030222012122-3200001201311013-2320233200223213-2023020020303332-3100033030103312-3221133022132003-3311030330120211-3330220323230031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.use_mtls.trusted_ca` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- [https.tls_parameters.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-3011011302300202-0330111321313003-2330103220210211-3110001210211122-3003210232200230-1123321222021031-0311313312302021-2121223323012232)
- https.tls_parameters.use_mtls.trusted_ca

<a id="canonical-0112200110031212-2210000101230001-2111323220223222-3301210332222213-1201113330201300-2031222231020231-0013220320030210-0032220303130233"></a>

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

<a id="canonical-3210323333233023-1121133223033233-3001301113333332-0023101320000021-2333301131310320-3010120203322003-1200120130130023-2032030132202103"></a>

### Direct properties for `https.tls_parameters.use_mtls.trusted_ca`

<a id="canonical-0102213021210203-1213303232001121-3310020331312320-1011031023003012-0113102322131101-1320201123120003-1031110011210220-2313211302332112"></a>

#### `https.tls_parameters.use_mtls.trusted_ca.name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1131101131120302-2021211210103200-3013310033210232-1200133230222032-0003222220331123-1232122002103313-3023203331322132-0133302312311213"></a>

<a id="canonical-0120211313312032-2121212002331030-0013202110213212-3130023232011031-1213120312133020-2221310233332121-1321211202323331-2123110113113101"></a>

#### `https.tls_parameters.use_mtls.trusted_ca.namespace` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3233011020300321-1332303333210002-2313023203122233-2123102120201100-2212011223310213-1011121121110031-2022322231301312-3000300003331111"></a>

<a id="canonical-3002131022003333-3331212122320233-3233223332033311-2202122203013103-2102323021203101-2310310331032013-3031301033211123-3313101302032111"></a>

#### `https.tls_parameters.use_mtls.trusted_ca.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0211201201101012-2032030012112333-1023312213220032-3102332300203301-0211033020001201-3031133001311121-3033000210103130-2210310102021001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.use_mtls.xfcc_disabled` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- [https.tls_parameters.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-3011011302300202-0330111321313003-2330103220210211-3110001210211122-3003210232200230-1123321222021031-0311313312302021-2121223323012232)
- https.tls_parameters.use_mtls.xfcc_disabled

<a id="canonical-1331032301023001-3233221131202213-3013203003233100-3100310233000002-1133211110313230-2303132201302232-0321001233101233-3203303001210332"></a>

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

<a id="canonical-1322023323220103-0120111112033232-3301100122101000-1022023102030303-3333300212330203-2221222021232123-2032120200322113-2021012121033102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https.tls_parameters.use_mtls.xfcc_options` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- [https.tls_parameters.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-3011011302300202-0330111321313003-2330103220210211-3110001210211122-3003210232200230-1123321222021031-0311313312302021-2121223323012232)
- https.tls_parameters.use_mtls.xfcc_options

<a id="canonical-3123120232210113-1120201001032320-2111202300231311-1233011303113113-0033112111130132-2103102323301100-2121332011213003-0221103330122030"></a>

Type: `"single"`. Computed.

X-Forwarded-Client-Cert header elements to be added to requests.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2202300213202300-0211233123302123-2311003020121030-1100101112100130-0131131122133011-3031312332220311-3301323333031233-0113023122232101"></a>

### Direct properties for `https.tls_parameters.use_mtls.xfcc_options`

<a id="canonical-1201111313033303-2301101222133031-1131123132122133-1032332010230322-3330002213022112-0323121332320112-1311202313003020-1323232323011023"></a>

#### `https.tls_parameters.use_mtls.xfcc_options.xfcc_header_elements` property

Type: `["list", "string"]`. Computed.

\[Enum: XFCC\_NONE|XFCC\_CERT|XFCC\_CHAIN|XFCC\_SUBJECT|XFCC\_URI|XFCC\_DNS\]
X-Forwarded-Client-Cert header elements to be added to requests. Possible values are \`XFCC\_NONE\`,
\`XFCC\_CERT\`, \`XFCC\_CHAIN\`, \`XFCC\_SUBJECT\`, \`XFCC\_URI\`, \`XFCC\_DNS\`. Defaults to
\`XFCC\_NONE\`.

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

<a id="canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- https_auto_cert

<a id="canonical-2110233032332000-2230100230102301-1331113121321211-2111113312132303-0303213311122331-0023321211013320-1001321133320230-1213120312211202"></a>

Type: `"single"`. Computed.

Choice for selecting HTTP proxy with bring your own certificates. Changing this type selection
requires recreation and may interrupt service. Supported settings within the same selected type
remain updatable.

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

<a id="canonical-0130023103110112-2011130000201230-0233132200200100-0131201212002100-1002133231310233-0330302200333230-3101332020123012-2010033311102203"></a>

### Direct properties for `https_auto_cert`

<a id="canonical-0311303033330303-3310230302321111-3310101311311301-0233300221203113-0203113302113003-0330110321330102-0222223303302301-2103113302231311"></a>

#### `https_auto_cert.add_hsts` property

Type: `"bool"`. Computed.

Add HTTP Strict-Transport-Security response header. Defaults to \`false\`. Server applies default
when omitted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0331322310020213-0323122013022122-2132011311010113-3333332311100003-0320221330101333-1110331123332312-3021233210201302-2230211233320122"></a>

<a id="canonical-2323002003113130-1213112321300033-1122030230200203-1220313320122123-3330022310212000-0121012101001022-2202121321121031-0330032313100300"></a>

#### `https_auto_cert.append_server_name` property

Type: `"string"`. Computed.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [coalescing_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-3330021010212031-0201302233303230-1312110320113122-3111123201223001-3030232120013221-0131032013033332-3200121332102023-3113132330033130): complete subsection reference.

<a id="canonical-1022121000013012-2210122201230101-0323231032201000-3112220222201320-1021312213330001-1312122132113121-1120021201212110-3023013230002123"></a>

<a id="canonical-3101332222001001-3023201221303112-3301320331333030-1320310302220301-2313300103010113-2233200113223102-3130320120110200-0133222001020111"></a>

#### `https_auto_cert.connection_idle_timeout` property

Type: `"number"`. Computed.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Server
applies default when omitted.

Additional upstream details:

Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [default_header](data-sources--http_loadbalancer--reference--group-018.md#canonical-3313201031301311-2102131330032210-3313230003331111-0210121112312003-1313020011212333-0021220023223203-2231232032113133-1232222301222200): complete subsection reference.

- [default_loadbalancer](data-sources--http_loadbalancer--reference--group-018.md#canonical-1103130031312003-1131202320320111-3122032303332320-0133012311002333-0120323032300311-0210331101320210-0111202033012332-2321332303211221): complete subsection reference.

- [disable_path_normalize](data-sources--http_loadbalancer--reference--group-018.md#canonical-2230032312031133-3013111001221130-3332003130301132-0033310101300322-2330212002021033-2331122100220031-1313210230130110-3033330130320201): complete subsection reference.

- [enable_path_normalize](data-sources--http_loadbalancer--reference--group-018.md#canonical-0220213122201313-3032013121031020-3011110313200301-0022211020120231-3132020001103231-0223313003010232-3011313203012023-1001002312333313): complete subsection reference.

- [http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-0223221303111212-2311310233201012-0022301212130031-0021031201210233-1012320212223102-1103101333323033-1011102123312000-1311212223230101): complete subsection reference.

<a id="canonical-2300013310330101-2231112312330110-2220123021033103-1121103012133100-0022010222230123-1010232133130323-1212203132330132-2103103120332300"></a>

<a id="canonical-3210213012000311-2333110300010302-2110001230231111-3122330011210003-1033203301220030-2233313210321123-1020121031223311-3310331221133123"></a>

#### `https_auto_cert.http_redirect` property

Type: `"bool"`. Computed.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS. Defaults to \`false\`. Server applies
default when omitted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [no_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-2102301213102300-3130211303123110-1211031100221301-0101130110032303-3233220303001111-3333311120020322-3000202300233331-2213001132011033): complete subsection reference.

- [non_default_loadbalancer](data-sources--http_loadbalancer--reference--group-019.md#canonical-3311320301322302-1303220013302033-0332111200020023-1101002220032022-2331300211113301-0130223132030323-1230211323100313-3210102020111331): complete subsection reference.

- [pass_through](data-sources--http_loadbalancer--reference--group-019.md#canonical-2221203212221022-2312321133113202-1222303012100333-2123103030313231-1012021303103113-3232300033113221-1000231113031103-1202131222211332): complete subsection reference.

<a id="canonical-1223132001313023-1132130130033013-2223113102320013-1233222230030003-3110033102033313-2313000012313310-3312300320110013-3301030113232203"></a>

<a id="canonical-1113033023111211-2202321012032103-1123002301330322-0002320120211330-3222113000101211-0333130213020211-3232220103123133-3113002330121110"></a>

#### `https_auto_cert.port` property

Type: `"number"`. Computed.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2221033321122331-2320132130221131-2103331021120121-2232302330210213-3110120001033202-3302220101023133-2011000211211333-3231030031022003"></a>

<a id="canonical-1230101321231101-1230320102100023-2002312321012101-2232201210312312-3121000003321313-3123220221110333-1110211132233022-3323231133133023"></a>

#### `https_auto_cert.port_ranges` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3233311203010311-3020131011312031-1321323001322102-0030023112213323-3033223021201001-0100303110310320-2121310001110313-3013333121201101"></a>

<a id="canonical-3312132030303123-3132210000301113-3103331320313300-0330030022331303-3213003120313000-3003032002331023-1031303332023302-0012322023313300"></a>

#### `https_auto_cert.server_name` property

Type: `"string"`. Computed.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [tls_config](data-sources--http_loadbalancer--reference--group-019.md#canonical-3100120330010023-2022113202221330-2121231130110112-1332331220230013-0321001320221130-1032222103322022-3013320033323313-3233231032313132): complete subsection reference.

- [use_mtls](data-sources--http_loadbalancer--reference--group-019.md#canonical-3112303013001300-2210312200330132-1212110332323131-2201213132010133-1223203032213302-1022302011231301-1200230200201123-3321302210331330): complete subsection reference.

<a id="canonical-3330021010212031-0201302233303230-1312110320113122-3111123201223001-3030232120013221-0131032013033332-3200121332102023-3113132330033130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.coalescing_options` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-018.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- https_auto_cert.coalescing_options

<a id="canonical-3021311112321010-2123213221032100-1022221302021303-1233121321132130-3331311112020211-0230211310322133-3302321111130103-2202233300110022"></a>

Type: `"single"`. Computed.

TLS connection coalescing configuration (not compatible with mTLS).

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

<a id="canonical-2231313032120122-0033212122033103-3013202002132130-0022202223233130-2022022121111200-0122203112020033-0103021202322132-3333303012222021"></a>

### Direct properties for `https_auto_cert.coalescing_options`

- [default_coalescing](data-sources--http_loadbalancer--reference--group-018.md#canonical-0200233202323332-2013331221132310-0011132211023212-0130321121221120-1121220203202323-0213333300220112-1130200031032132-1233230203100003): complete subsection reference.

- [strict_coalescing](data-sources--http_loadbalancer--reference--group-018.md#canonical-3111003333102233-0313220330123210-2332032101011011-3213121121321113-1213132132201201-3102132322123112-3103222003222232-2031320322011023): complete subsection reference.

<a id="canonical-0200233202323332-2013331221132310-0011132211023212-0130321121221120-1121220203202323-0213333300220112-1130200031032132-1233230203100003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.coalescing_options.default_coalescing` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-018.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.coalescing_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-3330021010212031-0201302233303230-1312110320113122-3111123201223001-3030232120013221-0131032013033332-3200121332102023-3113132330033130)
- https_auto_cert.coalescing_options.default_coalescing

<a id="canonical-3000211232211231-2211313012312213-3300122103111203-1012302301332313-1022012322010203-2101221333131013-2020100232013223-0133020302030203"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default coalescing.

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

<a id="canonical-3111003333102233-0313220330123210-2332032101011011-3213121121321113-1213132132201201-3102132322123112-3103222003222232-2031320322011023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.coalescing_options.strict_coalescing` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-018.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.coalescing_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-3330021010212031-0201302233303230-1312110320113122-3111123201223001-3030232120013221-0131032013033332-3200121332102023-3113132330033130)
- https_auto_cert.coalescing_options.strict_coalescing

<a id="canonical-1232211103222321-1213020133213102-1223032112033203-0322233012300131-1233203033210032-2102300033313232-3330113321213222-3211213231320210"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for strict coalescing.

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

<a id="canonical-3313201031301311-2102131330032210-3313230003331111-0210121112312003-1313020011212333-0021220023223203-2231232032113133-1232222301222200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.default_header` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-018.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- https_auto_cert.default_header

<a id="canonical-1332131330320330-1321132103220233-2120221220011102-2103131200010313-3002220212013211-1330003002213112-1012332033331211-2320331201030303"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default header.

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

<a id="canonical-1103130031312003-1131202320320111-3122032303332320-0133012311002333-0120323032300311-0210331101320210-0111202033012332-2321332303211221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.default_loadbalancer` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-018.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- https_auto_cert.default_loadbalancer

<a id="canonical-1310002031103010-3003232121201212-0302302330012121-0203311003230212-3103021323233100-3033313333021321-0222012030132132-0003232323333032"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default loadbalancer.

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

<a id="canonical-2230032312031133-3013111001221130-3332003130301132-0033310101300322-2330212002021033-2331122100220031-1313210230130110-3033330130320201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.disable_path_normalize` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-018.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- https_auto_cert.disable_path_normalize

<a id="canonical-2102321302210232-2332302131231010-0131312013131320-2012212301310233-0120332230332033-2112313030020221-3110202212112200-0332021003233212"></a>

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

<a id="canonical-0220213122201313-3032013121031020-3011110313200301-0022211020120231-3132020001103231-0223313003010232-3011313203012023-1001002312333313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.enable_path_normalize` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-018.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- https_auto_cert.enable_path_normalize

<a id="canonical-2022132030002222-1301101020102230-1130210022011301-0013220120003133-0213212031123332-0003110001233132-2012133123212302-0103012300312330"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

<a id="canonical-0223221303111212-2311310233201012-0022301212130031-0021031201210233-1012320212223102-1103101333323033-1011102123312000-1311212223230101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.http_protocol_options` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-018.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- https_auto_cert.http_protocol_options

<a id="canonical-1331022330003023-2221013213220111-2323330022030111-1122110020120211-0023131200032311-3103030130323100-0203233312031121-0021033122100321"></a>

Type: `"single"`. Computed.

HTTP protocol configuration OPTIONS for downstream connections.

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

<a id="canonical-3313331300031222-2013222033333222-1303033303311002-0100001131030221-2202122300121000-2110210022013001-0322300133330133-2101312223230113"></a>

### Direct properties for `https_auto_cert.http_protocol_options`

- [http_protocol_enable_v1_only](data-sources--http_loadbalancer--reference--group-018.md#canonical-3130213100020121-1321030213331201-3103012013210121-3201020002013131-0230020221301123-0111000310331322-0030300310110203-1322312112023100): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--http_loadbalancer--reference--group-019.md#canonical-1103132123031212-1312011311132313-1101103331230311-0211210100021021-0223211311301221-2201203120330131-0323223220202031-2113302232301212): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--http_loadbalancer--reference--group-019.md#canonical-0011132011130333-0210213032201022-0233112013130221-1333101332310100-3311123321233200-3302001330230303-1003013201132323-0021222312201002): complete subsection reference.

<a id="canonical-3130213100020121-1321030213331201-3103012013210121-3201020002013131-0230020221301123-0111000310331322-0030300310110203-1322312112023100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `https_auto_cert.http_protocol_options.http_protocol_enable_v1_only` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-018.md#canonical-3202333311333232-3203220122320221-0301112021032122-2113130221132111-3003312022133121-2310001002222111-3002330102201213-2112003113303131)
- [https_auto_cert.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-0223221303111212-2311310233201012-0022301212130031-0021031201210233-1012320212223102-1103101333323033-1011102123312000-1311212223230101)
- https_auto_cert.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-3122230030213302-0110123203102111-1022330022220130-2110032233112112-2121030021330223-3013320302021203-2322202231033223-2000331212220021"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2120233130210211-3023110031222322-3213002310233102-3011212222003100-3100301003013202-3201132302302323-0133331020221021-2201121231200201"></a>

### Direct properties for `https_auto_cert.http_protocol_options.http_protocol_enable_v1_only`

- [header_transformation](data-sources--http_loadbalancer--reference--group-019.md#canonical-1031030003113032-1112300103303030-1323323222230133-3030212220032313-2311022311132132-1003301332113213-3210233333230020-3022233020101003): complete subsection reference.
