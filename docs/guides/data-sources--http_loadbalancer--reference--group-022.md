---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-0133101222230110-3330132132320300-0030332032133332-2201321222202100-0210021221302113-2003130022332300-0003200022121103-1220013203023303"></a>

## Direct properties for `origin_server_subset_rule_list.origin_server_subset_rules.metadata`

<a id="canonical-3203310101131030-0112123223313210-0233122201320021-3011113310230213-1000102113222211-3112302223123013-3113110322012230-1303113123013013"></a>

### `origin_server_subset_rule_list.origin_server_subset_rules.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-0221011220020100-0010102100300023-2320320123032101-0011130120021001-3302003202133213-2323021212110231-0320310120323332-3200202302220113"></a>

<a id="canonical-0313130002120332-0331213311221033-0301121333233030-3202320020031310-2230132323113010-0202111231332120-2330100331110232-3220103001013122"></a>

### `origin_server_subset_rule_list.origin_server_subset_rules.metadata.name` property

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

<a id="canonical-0111233133120120-0323011111302022-2230023031221212-1230110330012322-2211321231001201-2002321333213021-1021202022222123-0122221101002230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_server_subset_rule_list.origin_server_subset_rules.none` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [origin_server_subset_rule_list](data-sources--http_loadbalancer--reference--group-021.md#canonical-1122320203011213-2200313323321113-2233212203221013-0302201330202203-1033230323313230-2002132032302120-2000012232112303-3133330100203312)
- [origin_server_subset_rule_list.origin_server_subset_rules](data-sources--http_loadbalancer--reference--group-021.md#canonical-1133030303102313-3120203001320331-2302113230222021-2220303003220112-2322131112222021-3032222103030203-0211313010002121-2333101020031113)
- origin_server_subset_rule_list.origin_server_subset_rules.none

<a id="canonical-3220000002130022-1013102213221233-0020133331001223-2221211223130013-3103032131101232-2102232223203103-2301220023222223-2331231202200323"></a>

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

<a id="canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- policy_based_challenge

<a id="canonical-0000010110233333-3321322033031133-3333123333332012-3000223230222011-1100121210223232-2221310130103120-2032020232112022-0210222331102001"></a>

Type: `"single"`. Computed.

Specifies the settings for policy rule based challenge.

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
  "x-ves-oneof-field-challenge_choice": "[\"always_enable_captcha_challenge\",\"always_enable_js_challenge\",\"no_challenge\"]",
  "x-ves-oneof-field-js_challenge_parameters_choice": "[\"default_js_challenge_parameters\",\"js_challenge_parameters\"]",
  "x-ves-oneof-field-malicious_user_mitigation_choice": "[\"default_mitigation_settings\",\"malicious_user_mitigation\"]",
  "x-ves-oneof-field-temporary_blocking_parameters_choice": "[\"default_temporary_blocking_parameters\",\"temporary_user_blocking\"]"
}
```

<a id="canonical-1122312030310313-1310321231321303-3023122000310030-3110231122300131-0321210321323122-2001222023113023-0123101113300210-1313221313100331"></a>

### Direct properties for `policy_based_challenge`

- [always_enable_captcha_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-2200212312300130-0003330022130010-2311321111322302-2110000112113100-1213102032030312-1303002201233032-3003233221121221-1333312030100221): complete subsection reference.

- [always_enable_js_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1133222000131230-0023032200003112-1111033230311311-3113100333330212-1131301231310001-0331222323011202-2222330220133121-3311322232103110): complete subsection reference.

- [captcha_challenge_parameters](data-sources--http_loadbalancer--reference--group-022.md#canonical-0031232123230311-0211102100302300-0122202222231323-1011002121100331-1103000211323133-0100003031322133-3132130021030323-0102120103131112): complete subsection reference.

- [default_captcha_challenge_parameters](data-sources--http_loadbalancer--reference--group-022.md#canonical-3033122122132230-3012313221123333-2130013103000000-2130320102210233-2202310122202323-0320010320033221-3010302311012312-1221131320330202): complete subsection reference.

- [default_js_challenge_parameters](data-sources--http_loadbalancer--reference--group-022.md#canonical-3210233223003011-1123023021123130-0201102132101200-0232230000032100-1303300132120223-2103303302313301-3111233202020030-1333221121021200): complete subsection reference.

- [default_mitigation_settings](data-sources--http_loadbalancer--reference--group-022.md#canonical-3001110002120201-1020301121231202-0010013333031201-0123200211120221-0211033021233231-0012210122330320-3220310110222103-2012203210302011): complete subsection reference.

- [default_temporary_blocking_parameters](data-sources--http_loadbalancer--reference--group-022.md#canonical-2313131231020200-2210311011111232-0132022020301311-2203023202030330-1210133110122201-3131320010032300-0101231220213330-1321102212211301): complete subsection reference.

- [js_challenge_parameters](data-sources--http_loadbalancer--reference--group-022.md#canonical-1312032111102320-0120213331303133-3332231110231301-3123320201021121-0100212113022232-2312321322331120-1131130322332031-0221220001021103): complete subsection reference.

- [malicious_user_mitigation](data-sources--http_loadbalancer--reference--group-022.md#canonical-0301011213301031-1302102201002310-1003103333312111-2320002321312333-1213202200312033-3203123032001123-1313300002102203-0212131202331300): complete subsection reference.

- [no_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1120331230101123-3020020300310231-3110322023000131-2130202113002001-2202123002003002-0100220202023003-2230213230101000-2323120133200032): complete subsection reference.

- [rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212): complete subsection reference.

- [temporary_user_blocking](data-sources--http_loadbalancer--reference--group-023.md#canonical-3113023112302332-0333303302133233-0010212202113330-1313232030220213-3320210011012203-0101333020233000-0132210002222132-1003010103121113): complete subsection reference.

<a id="canonical-2200212312300130-0003330022130010-2311321111322302-2110000112113100-1213102032030312-1303002201233032-3003233221121221-1333312030100221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.always_enable_captcha_challenge` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- policy_based_challenge.always_enable_captcha_challenge

<a id="canonical-3121012012013211-3323213133230010-0121322031120011-1103130301030223-3222203222102032-2112021300023120-2013210103032210-3301132303220102"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for always enable captcha challenge.

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

<a id="canonical-1133222000131230-0023032200003112-1111033230311311-3113100333330212-1131301231310001-0331222323011202-2222330220133121-3311322232103110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.always_enable_js_challenge` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- policy_based_challenge.always_enable_js_challenge

<a id="canonical-3323022012313030-2233123221131010-3103212301010220-2311320330021102-3032213230100200-1300311310211001-0201222100022212-1312313303203102"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for always enable js challenge.

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

<a id="canonical-0031232123230311-0211102100302300-0122202222231323-1011002121100331-1103000211323133-0100003031322133-3132130021030323-0102120103131112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.captcha_challenge_parameters` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- policy_based_challenge.captcha_challenge_parameters

<a id="canonical-0132303121313313-3220013101030022-2202311122010311-3232101133202333-2223200233111313-2131200032013320-0012210131121122-0020123333020030"></a>

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

<a id="canonical-3323030110003103-1132200231132213-2230302021230333-3310130221213203-3032010211100003-0221021002332211-2223102223112212-1111302322030003"></a>

### Direct properties for `policy_based_challenge.captcha_challenge_parameters`

<a id="canonical-3331203012200121-0120010122330033-1202213213203201-1300222033103301-3333233112100233-0123313301103302-1213311210312002-2110332011331003"></a>

#### `policy_based_challenge.captcha_challenge_parameters.cookie_expiry` property

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
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-0121122113120112-1310023211120300-3313002033321031-1210211310120221-2030302223201110-2201122323202202-1031221211221213-1330103120002300"></a>

<a id="canonical-3010221100331301-3032023022032332-0001323331230133-3310231221221030-3331011111111130-2232320203102022-0320010311200321-2101330332010123"></a>

#### `policy_based_challenge.captcha_challenge_parameters.custom_page` property

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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3033122122132230-3012313221123333-2130013103000000-2130320102210233-2202310122202323-0320010320033221-3010302311012312-1221131320330202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.default_captcha_challenge_parameters` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- policy_based_challenge.default_captcha_challenge_parameters

<a id="canonical-3322101031112010-3003312211102330-3031110103022003-3303232011130333-1301313011313230-0221022133301212-0323333022322322-3213200032033013"></a>

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

<a id="canonical-3210233223003011-1123023021123130-0201102132101200-0232230000032100-1303300132120223-2103303302313301-3111233202020030-1333221121021200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.default_js_challenge_parameters` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- policy_based_challenge.default_js_challenge_parameters

<a id="canonical-2322330210031102-2200133111320321-2320012000121002-0200120130123211-0220321110202302-0213003230031330-0020320201232120-1010303313221201"></a>

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

<a id="canonical-3001110002120201-1020301121231202-0010013333031201-0123200211120221-0211033021233231-0012210122330320-3220310110222103-2012203210302011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.default_mitigation_settings` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- policy_based_challenge.default_mitigation_settings

<a id="canonical-0220230311302122-1011113122131210-0222011121310121-3331233102121210-0322323201001323-0200133133213020-0131223021001230-3130121313330233"></a>

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

<a id="canonical-2313131231020200-2210311011111232-0132022020301311-2203023202030330-1210133110122201-3131320010032300-0101231220213330-1321102212211301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.default_temporary_blocking_parameters` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- policy_based_challenge.default_temporary_blocking_parameters

<a id="canonical-0200121023310230-1111301200331020-3210230111033103-0330102303213221-3233123210101111-1330011310032202-1131330201332023-2010232133033311"></a>

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

<a id="canonical-1312032111102320-0120213331303133-3332231110231301-3123320201021121-0100212113022232-2312321322331120-1131130322332031-0221220001021103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.js_challenge_parameters` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- policy_based_challenge.js_challenge_parameters

<a id="canonical-3203002322011330-3113332212120003-3222001200220312-0023331112003013-3222203310231300-3101021132311100-2120031023013102-0023222003122331"></a>

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

<a id="canonical-2301303203313331-2132111103102132-1012123313131101-2331131223000320-1221113201133002-2121110300131322-3113100313220201-2030110311210332"></a>

### Direct properties for `policy_based_challenge.js_challenge_parameters`

<a id="canonical-0003133331312000-0011121031123221-0300213132113213-3303132102233313-1003302032233220-0222233301123102-1100123320313320-1320332312310201"></a>

#### `policy_based_challenge.js_challenge_parameters.cookie_expiry` property

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
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-1203231033310133-2331231220130123-3310212231013323-2030122130212002-2311120313300000-2122002323123320-0000333103231300-1231333102013321"></a>

<a id="canonical-1302333213221022-3212233321133023-3013002123103131-3320333310010111-0030120022111330-3230122310331120-2022123233220301-2322112232231213"></a>

#### `policy_based_challenge.js_challenge_parameters.custom_page` property

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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3210002030013222-3132222333023121-3311233312321021-0203132021330121-0213221101111323-3222033001100232-0212232023133312-0321010202120201"></a>

<a id="canonical-3011132210331333-3211132131303021-3300310120233212-1001202122113231-3201001133011112-3030011131122001-3112203201132023-3321123230223122"></a>

#### `policy_based_challenge.js_challenge_parameters.js_script_delay` property

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0301011213301031-1302102201002310-1003103333312111-2320002321312333-1213202200312033-3203123032001123-1313300002102203-0212131202331300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.malicious_user_mitigation` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- policy_based_challenge.malicious_user_mitigation

<a id="canonical-3302122032032323-0332311103021220-0132023113133220-3003130332001212-3111220003211231-1110003110321010-0003221332013300-1001111321300033"></a>

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

<a id="canonical-2122332230331111-1311302011111131-0121223330301011-2213120100222133-2132020113213201-3302203012121211-3020330021200220-3133133221331120"></a>

### Direct properties for `policy_based_challenge.malicious_user_mitigation`

<a id="canonical-3330201221322233-2133222010300220-1100200311221330-0313222022323022-1201310113103303-2200313211221012-3231222222203011-1220213100113032"></a>

#### `policy_based_challenge.malicious_user_mitigation.name` property

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

<a id="canonical-3332302130211120-1121101230201213-0320201321001023-0211121111021023-3231233123303320-2331021100212230-3323112312033131-0130222313003100"></a>

<a id="canonical-0020212133010222-1203012132113013-2123231121303312-3020130032203302-1303311130310222-0233103130012032-2312000132222023-2223013221321101"></a>

#### `policy_based_challenge.malicious_user_mitigation.namespace` property

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

<a id="canonical-0002030331333323-2000210011222100-2120332221021100-3032121320203300-2303322220203310-3312021112001220-0312121230132122-2023102032132110"></a>

<a id="canonical-3311200300333221-3123332220000310-0210130032311012-2303111032021331-3102230333221301-2212313332103013-3113021100023003-0332322113131011"></a>

#### `policy_based_challenge.malicious_user_mitigation.tenant` property

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

<a id="canonical-1120331230101123-3020020300310231-3110322023000131-2130202113002001-2202123002003002-0100220202023003-2230213230101000-2323120133200032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.no_challenge` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- policy_based_challenge.no_challenge

<a id="canonical-2320130130323121-1010221101023232-0101312303131301-3231300231100303-0112001120310333-1120122121200301-3300103322312213-0200222212323132"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no challenge.

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

<a id="canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- policy_based_challenge.rule_list

<a id="canonical-0112033020000303-3130130031301111-1031310202013211-1310223133303330-2212102323023212-1220110221030312-0132121103133023-2131123033033021"></a>

Type: `"single"`. Computed.

List of challenge rules to be used in policy based challenge.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1102223020023111-3023202121320022-1003213121120210-1331011330330202-3002203132120110-1302031321231103-2221121232022010-3331323211111013"></a>

### Direct properties for `policy_based_challenge.rule_list`

- [rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003): complete subsection reference.

<a id="canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- policy_based_challenge.rule_list.rules

<a id="canonical-0002011213112120-2203023133222132-2112033212121333-0313333330000000-1002333110220330-2302313333221300-2223121032312303-0222110022133313"></a>

Type: `"list"`. Computed.

Rules that specify the match conditions and challenge type to be launched. When a challenge type is
selected to be always enabled, these rules can be used to disable challenge or launch a different
challenge for requests that match the specified conditions.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
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

<a id="canonical-1102303221202220-1111220031122222-1223330230233210-2020101033330201-0212330203210331-2311223030130111-2203332212321103-0201010321231231"></a>

### Direct properties for `policy_based_challenge.rule_list.rules`

- [metadata](data-sources--http_loadbalancer--reference--group-022.md#canonical-2220202321033213-0130010312331133-2002322003111230-2100213323113320-3321130302213322-1013132322011023-2110203123300232-2230122032313120): complete subsection reference.

- [spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020): complete subsection reference.

<a id="canonical-2220202321033213-0130010312331133-2002322003111230-2100213323113320-3321130302213322-1013132322011023-2110203123300232-2230122032313120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.metadata` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- policy_based_challenge.rule_list.rules.metadata

<a id="canonical-1320122212031230-2012200322011311-2102303113202222-2033330331302223-3232121112232121-1233310221131031-1223210002112232-2312031113302112"></a>

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

<a id="canonical-2030003311221103-2132102310302111-1020010233302023-3310310033113002-3233311001033120-1313111221311123-3030300131300223-0202102320112231"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.metadata`

<a id="canonical-3232102331323011-3133222313322001-0320330110130102-0331230312030210-0322212330020023-3100100102203211-1133210112331212-1231301123322203"></a>

#### `policy_based_challenge.rule_list.rules.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-0203322113230030-2330031122113012-0232320213002322-2011233312000222-3131100200223222-2312312222021021-0210301030220213-0103011030011032"></a>

<a id="canonical-1021003313212012-0322011223023000-1320302002221031-3113033111102112-0310303103013321-0332232231221101-3111022122300220-2201313031232221"></a>

#### `policy_based_challenge.rule_list.rules.metadata.name` property

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

<a id="canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- policy_based_challenge.rule_list.rules.spec

<a id="canonical-0323031210333222-1233122213321110-2200220321203222-1332132121002211-3200103321213011-1011000132321300-0103131020221130-2213133001000221"></a>

Type: `"single"`. Computed.

A Challenge Rule consists of an unordered list of predicates and an action. The predicates are
evaluated against a set of input fields that are extracted from or derived from an L7 request API. A
request API is considered to match the rule if all predicates in the rule evaluate to true for that
request. Any predicates that are not specified in a rule are implicitly considered to be true. If a
request API matches a challenge rule, the configured challenge is enforced.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-asn_choice": "[\"any_asn\",\"asn_list\",\"asn_matcher\"]",
  "x-ves-oneof-field-challenge_action": "[\"disable_challenge\",\"enable_captcha_challenge\",\"enable_javascript_challenge\"]",
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_selector\"]",
  "x-ves-oneof-field-ip_choice": "[\"any_ip\",\"ip_matcher\",\"ip_prefix_list\"]",
  "x-ves-oneof-field-tls_fingerprint_choice": "[\"tls_fingerprint_matcher\"]"
}
```

<a id="canonical-3302301330221333-1121010303022122-0231022211323031-0003100013311031-1103321113030101-1000200200131202-1300322111112210-3332032210021321"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec`

- [any_asn](data-sources--http_loadbalancer--reference--group-022.md#canonical-1110112001201001-2303203033000203-2130231020022321-0301000311232001-3303110003001103-3213333301100031-3233111003330311-2233200300222312): complete subsection reference.

- [any_client](data-sources--http_loadbalancer--reference--group-022.md#canonical-2131033212110203-1003133220231010-0120221011210230-2010103232223323-2201102302313333-1112211012012310-2021013231000133-1110323231322001): complete subsection reference.

- [any_ip](data-sources--http_loadbalancer--reference--group-022.md#canonical-2221231100300023-2210020303223021-0031320102333202-2102200030032311-1121023002003232-1230320120100333-3030311210313130-0320312022112313): complete subsection reference.

- [arg_matchers](data-sources--http_loadbalancer--reference--group-022.md#canonical-0110033001130101-2302000012330021-1032020000013231-1010313010121321-2002213032330122-2301101210112121-1323313301211320-0322322111123321): complete subsection reference.

- [asn_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-2102013312122211-0130022121103002-0113331230002130-1101210113331010-0303031332122311-1011021310002013-3013123220013223-3303110032133201): complete subsection reference.

- [asn_matcher](data-sources--http_loadbalancer--reference--group-022.md#canonical-1210333023331302-1002320301033103-0323331220200302-3131012022231022-1030210300202110-2222323112233201-2100120123121101-2231120113310221): complete subsection reference.

- [body_matcher](data-sources--http_loadbalancer--reference--group-022.md#canonical-2130222302101313-2233302001002321-3322010333311213-1103213000211301-2133320123330332-2302210210112022-2321133102321132-2323201001222012): complete subsection reference.

- [client_selector](data-sources--http_loadbalancer--reference--group-022.md#canonical-2000323223223323-2213103122032232-2032012302000332-1103023033311001-3303310000122131-2110300101110211-3021302032013102-2223201132322222): complete subsection reference.

- [cookie_matchers](data-sources--http_loadbalancer--reference--group-022.md#canonical-0210300122220212-1033203223102220-1333302020031120-3310022303300010-3023022233223333-2310311011132310-0312022113220300-3331113000211000): complete subsection reference.

- [disable_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-3020133121101200-2222013212131112-3111032310221132-2213303201101321-1300132123112201-3131333230313112-1220001302231030-2202110030031102): complete subsection reference.

- [domain_matcher](data-sources--http_loadbalancer--reference--group-022.md#canonical-3311122103211312-0002003200333203-1210023030112201-2031031301112330-3220213200303010-1323003002301022-1230010220321100-3120321133102132): complete subsection reference.

- [enable_captcha_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-2322120012012001-0300113331321331-0300333021300022-2012123332131133-3100110230212230-0001321233021031-1103301000033302-2331302011012102): complete subsection reference.

- [enable_javascript_challenge](data-sources--http_loadbalancer--reference--group-023.md#canonical-0130320232310300-1003033303212310-1030202000201111-3022320032203012-1110200010232110-0320022322130201-0102110002122103-1202210113103203): complete subsection reference.

<a id="canonical-3033112312333232-0322231223101310-0020003010130031-2022222232313103-2022211111013020-0123221323323002-2203121000301023-0132100030302332"></a>

<a id="canonical-3200310313002133-0010311010212123-2312102101331200-0220220330102301-2110012303211210-2303022333112300-0031100002131010-2121220122333113"></a>

#### `policy_based_challenge.rule_list.rules.spec.expiration_timestamp` property

Type: `"string"`. Computed.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Additional upstream details:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
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
  }
}
```

- [headers](data-sources--http_loadbalancer--reference--group-023.md#canonical-1113203301223010-0101302010202223-2303331001210210-3312001032331310-0000232302322131-2020012100112010-0301301230311022-0313021023200101): complete subsection reference.

- [http_method](data-sources--http_loadbalancer--reference--group-023.md#canonical-3020310301331122-1212000113011100-1023213212330331-3122103113210321-2002122203320320-2120121220311201-1333220030000032-3120130221023203): complete subsection reference.

- [ip_matcher](data-sources--http_loadbalancer--reference--group-023.md#canonical-3022123003013200-3203320321330332-0121233201103200-2331212133333301-0201213111233230-2211313020321111-1111213132220212-0330232000202000): complete subsection reference.

- [ip_prefix_list](data-sources--http_loadbalancer--reference--group-023.md#canonical-2102212001013000-1331112303230010-2332311101103003-1203312030331231-1110322311233210-0133212103001332-1212233031202211-1222020223230121): complete subsection reference.

- [path](data-sources--http_loadbalancer--reference--group-023.md#canonical-1130112213311222-2212120101303321-2312021201330112-1010000001222113-3331020212132323-1321320012001100-3330022120222112-2310110330301333): complete subsection reference.

- [query_params](data-sources--http_loadbalancer--reference--group-023.md#canonical-2212313110313101-3103020010021203-3333032032000212-0231022312231220-3100321030220223-2310330231030300-0030223113223002-1002022001130332): complete subsection reference.

- [tls_fingerprint_matcher](data-sources--http_loadbalancer--reference--group-023.md#canonical-0022113021303313-2123000230233031-0231021111110113-2111313001202330-1220202120123310-1300312201020212-0013033111222322-0130112033123023): complete subsection reference.

<a id="canonical-1110112001201001-2303203033000203-2130231020022321-0301000311232001-3303110003001103-3213333301100031-3233111003330311-2233200300222312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.any_asn` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.any_asn

<a id="canonical-2220023100303032-0201020023121322-1020213323310013-3331211221230310-0230321033103003-2133322231303300-2002322302221203-3013030202323010"></a>

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

<a id="canonical-2131033212110203-1003133220231010-0120221011210230-2010103232223323-2201102302313333-1112211012012310-2021013231000133-1110323231322001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.any_client` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.any_client

<a id="canonical-1320332133220011-1122301231023312-0212022112130303-3201123103131022-3032003100032233-3112231132220031-3302202301131321-2230113321023123"></a>

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

<a id="canonical-2221231100300023-2210020303223021-0031320102333202-2102200030032311-1121023002003232-1230320120100333-3030311210313130-0320312022112313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.any_ip` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.any_ip

<a id="canonical-0020203310301203-0222332010332030-0333232231123232-2132212121232310-0302011223212231-1303200332210123-1112210021003002-0231110321210021"></a>

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

<a id="canonical-0110033001130101-2302000012330021-1032020000013231-1010313010121321-2002213032330122-2301101210112121-1323313301211320-0322322111123321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.arg_matchers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.arg_matchers

<a id="canonical-1303310122220111-1310201020311311-0022202211303133-0113322003103112-3000220033011012-3023310130301333-3330210133121100-1112101332002300"></a>

Type: `"list"`. Computed.

A list of predicates for all POST args that need to be matched. The criteria for matching each arg
are described in individual instances of ArgMatcherType. The actual arg values are extracted from
the request API as a list of strings for each arg selector name. Note that all specified arg matcher
predicates must evaluate to true. A request body greater than 64KB will not be evaluated.

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-0303200301301121-2121010210322130-1102202200223211-1220221231323032-2210221220320223-1112322113022111-3332233133030211-3111101301223222"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.arg_matchers`

- [check_not_present](data-sources--http_loadbalancer--reference--group-022.md#canonical-3133333330030322-2220000010213013-3110010100310230-3123211332302122-1120321320202223-0032130303003232-0133213021333212-3313220200332210): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-022.md#canonical-0131230030203320-2020323201102121-0131233223113130-1220130231031001-2132221112102023-0303101232333333-1121320312132301-0332103121303333): complete subsection reference.

<a id="canonical-1300033201121020-3121120032303230-0300201201300211-2220110233022121-2302201331021310-3131220330122320-3101103312320132-3330122331231022"></a>

<a id="canonical-3322320331333102-0312312201321220-0022312322010031-1020201302021203-1031222123021331-3330230301012322-1333323220102101-1002122022103000"></a>

#### `policy_based_challenge.rule_list.rules.spec.arg_matchers.invert_matcher` property

Type: `"bool"`. Computed.

Invert Matcher. Invert Match of the expression defined.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [item](data-sources--http_loadbalancer--reference--group-022.md#canonical-1103003032012030-2123231301121231-0303231223113312-2020201323002031-3102033021301203-0201210303120020-2231301110220133-2201020120000100): complete subsection reference.

<a id="canonical-2031110211220022-1302122011131010-0130201210101002-3300310203213300-0132223112030312-0110111231210230-0010112210232231-2122321321010303"></a>

<a id="canonical-3020011133011323-2330103310112000-2022111201213000-1313332113300020-3113130010110101-2100033302010033-2202313102333011-1103110210100133"></a>

#### `policy_based_challenge.rule_list.rules.spec.arg_matchers.name` property

Type: `"string"`. Computed.

A case-sensitive JSON path in the HTTP request body.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.json_path": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.json_path": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-3133333330030322-2220000010213013-3110010100310230-3123211332302122-1120321320202223-0032130303003232-0133213021333212-3313220200332210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](data-sources--http_loadbalancer--reference--group-022.md#canonical-0110033001130101-2302000012330021-1032020000013231-1010313010121321-2002213032330122-2301101210112121-1323313301211320-0322322111123321)
- policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present

<a id="canonical-3132110003131111-0312212112113120-2123100000223032-2112222012130030-3210212013002310-0220001332011111-0212120110300132-0122010221331310"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

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

<a id="canonical-0131230030203320-2020323201102121-0131233223113130-1220130231031001-2132221112102023-0303101232333333-1121320312132301-0332103121303333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](data-sources--http_loadbalancer--reference--group-022.md#canonical-0110033001130101-2302000012330021-1032020000013231-1010313010121321-2002213032330122-2301101210112121-1323313301211320-0322322111123321)
- policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present

<a id="canonical-0331100002103000-2110300110112210-1123331221202330-2020030102120200-0212332011330000-0300311321000133-1020311210322232-2310123323211011"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

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

<a id="canonical-1103003032012030-2123231301121231-0303231223113312-2020201323002031-3102033021301203-0201210303120020-2231301110220133-2201020120000100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.arg_matchers.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](data-sources--http_loadbalancer--reference--group-022.md#canonical-0110033001130101-2302000012330021-1032020000013231-1010313010121321-2002213032330122-2301101210112121-1323313301211320-0322322111123321)
- policy_based_challenge.rule_list.rules.spec.arg_matchers.item

<a id="canonical-3022020131222231-2031302300231130-0211213323132330-1323110100220210-1331032231022312-0011001122211311-3001012022223202-0323212232003001"></a>

Type: `"single"`. Computed.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3331210313213101-1231021323203320-3013130102133012-2321110313312110-1231303033103033-0020010021003111-3031232022100310-3111213130020213"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.arg_matchers.item`

<a id="canonical-3301103100301020-3212311120100233-3100103331333012-2330322100220003-3131102233022132-1320123023013013-0203023131301122-3330202100022320"></a>

#### `policy_based_challenge.rule_list.rules.spec.arg_matchers.item.exact_values` property

Type: `["list", "string"]`. Computed.

A list of exact values to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0203203023012221-0201011032122302-0021203220110220-2123021010103330-1233201011321320-1310210130311222-3122202023101100-1312202023102302"></a>

<a id="canonical-3030022212033111-0120222230221013-0213312330303121-1010122313100130-1001130023102230-3102020331203001-0230020032102323-3323212113030330"></a>

#### `policy_based_challenge.rule_list.rules.spec.arg_matchers.item.regex_values` property

Type: `["list", "string"]`. Computed.

A list of regular expressions to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0131021111021310-0332330311322210-2030032021111133-1113033310133001-3131020312010303-0221112132331010-3332032232312131-3022223311233122"></a>

<a id="canonical-2231110020102221-0301211110323223-1322213012011031-1133012032132332-0303020331211012-3031201013232101-3221123130102203-1302231100011020"></a>

#### `policy_based_challenge.rule_list.rules.spec.arg_matchers.item.transformers` property

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2102013312122211-0130022121103002-0113331230002130-1101210113331010-0303031332122311-1011021310002013-3013123220013223-3303110032133201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.asn_list` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.asn_list

<a id="canonical-3232022000013312-3200233001020001-3320123322111222-1231113310203313-2101333021013102-2311020101201001-1112031030323333-1200013222133010"></a>

Type: `"single"`. Computed.

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1200011310113202-0012312100101001-2123320123133123-3310211030223310-3100202221023032-0122321333030322-3210331000121012-0230021211113001"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.asn_list`

<a id="canonical-2133013201110012-3333233112010203-2020320002010323-2103112101333100-3220231003303111-3203032021010330-1132312223313103-1301122032301323"></a>

#### `policy_based_challenge.rule_list.rules.spec.asn_list.as_numbers` property

Type: `["list", "number"]`. Computed.

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1210333023331302-1002320301033103-0323331220200302-3131012022231022-1030210300202110-2222323112233201-2100120123121101-2231120113310221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.asn_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.asn_matcher

<a id="canonical-3220001330130201-0032103022002012-1123331232101010-1221001120032011-1020123223322103-1001023023201332-0032132302300111-3002211101233033"></a>

Type: `"single"`. Computed.

Match any AS number contained in the list of bgp\_asn\_sets.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1233333131133332-2033121022032332-0010031032132203-0202330313302010-0231003303230311-1112101130332210-0021102220203032-3032001300031133"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.asn_matcher`

- [asn_sets](data-sources--http_loadbalancer--reference--group-022.md#canonical-2330133203333230-3320312011011120-0003000002223301-2313202131110302-1230200012313310-1220131111110131-3210202332000010-0033330023301131): complete subsection reference.

<a id="canonical-2330133203333230-3320312011011120-0003000002223301-2313202131110302-1230200012313310-1220131111110131-3210202332000010-0033330023301131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [policy_based_challenge.rule_list.rules.spec.asn_matcher](data-sources--http_loadbalancer--reference--group-022.md#canonical-1210333023331302-1002320301033103-0323331220200302-3131012022231022-1030210300202110-2222323112233201-2100120123121101-2231120113310221)
- policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets

<a id="canonical-0220113231231311-3230023011213210-0021020131323322-1222330131133220-2003100300122330-1321123131022010-1022210032032331-3223123101303302"></a>

Type: `"list"`. Computed.

A list of references to bgp\_asn\_set objects.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

<a id="canonical-2111232103100230-2230330110202230-0033232121210313-3121031201003210-2012001033213101-3021233332032130-2103122202013212-3023002113233013"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets`

<a id="canonical-3002322311301132-0302222011030120-1313222213123112-3033301030032031-0003302330202212-3202232333102110-0112301333032320-0202202321032221"></a>

#### `policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2020202301133111-3133123023032310-2013223133022230-1111122223211020-0212201230123211-2021033200012320-0113012132200031-0102013013233303"></a>

<a id="canonical-3132121221221232-3332120020302100-0312123322223303-1122102313331013-2012130320320003-2003301003322010-1213212323222022-3202320110131310"></a>

#### `policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3200123230113132-0102113213020330-3300231123001022-0011201230133100-3101311002202213-3010220102322013-0233200310020020-1013321031330201"></a>

<a id="canonical-2033022123213030-2133132232223203-3302303102032123-3300000223031323-3221013103230132-2100213211130103-3103021312303230-1003213021302220"></a>

#### `policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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

<a id="canonical-2133232101110300-0202330332101011-0011110203220021-3132013232313322-3220323303233233-3003013201122120-1300210321103202-3012331002101020"></a>

<a id="canonical-2033120231230213-3333302023003022-1033232132003012-2213323002130202-0133121023320223-1202320013210020-3103222113323333-1333311321330123"></a>

#### `policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3311301110223232-3312330003002330-1013303203123120-0230323233103210-0110212303010021-1300031323030222-0333223321120221-0011220132113313"></a>

<a id="canonical-1200020010330132-2110331110322201-1020123032220030-1203002112333122-1201000102120222-2103200231332201-2010220122112220-1221033031322031"></a>

#### `policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets.uid` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2130222302101313-2233302001002321-3322010333311213-1103213000211301-2133320123330332-2302210210112022-2321133102321132-2323201001222012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.body_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.body_matcher

<a id="canonical-2320203100130322-1320213032223100-2221301312313033-1311212031313012-2201313001110203-3111313110310312-2023302023023112-2332232001223133"></a>

Type: `"single"`. Computed.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2312330013020012-0101133201031001-2201102303322012-3200121110322213-1303033321210111-2200300022101123-0223321033030233-0302030122123231"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.body_matcher`

<a id="canonical-2203303212001303-0332132311332132-2203203200013223-2111300323110220-0300310333030032-2312020031011331-2132301113020221-1333120120311131"></a>

#### `policy_based_challenge.rule_list.rules.spec.body_matcher.exact_values` property

Type: `["list", "string"]`. Computed.

A list of exact values to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0030032112313111-0303221323313010-1301110103030221-2313330003011220-2100301220000332-1112032032300133-2031031200233333-1121320313213301"></a>

<a id="canonical-1100221212132333-1102322010033030-2300111101003211-0102200003212002-2033223031012222-0203321113302010-1030320110012322-1220010022200001"></a>

#### `policy_based_challenge.rule_list.rules.spec.body_matcher.regex_values` property

Type: `["list", "string"]`. Computed.

A list of regular expressions to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1101300010012111-2000033020112020-1120001113123222-1133030333101103-3331121321112312-1021321332102130-1112322113300131-0112102121002100"></a>

<a id="canonical-3102103333233101-1100022010121232-1321120000111301-0213132112311303-0120230103312113-1201220203133211-1312130231113302-2213001320102313"></a>

#### `policy_based_challenge.rule_list.rules.spec.body_matcher.transformers` property

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2000323223223323-2213103122032232-2032012302000332-1103023033311001-3303310000122131-2110300101110211-3021302032013102-2223201132322222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.client_selector` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.client_selector

<a id="canonical-2130302001110200-3113200111030130-3220302221122103-1222333022020330-1320201122303212-3120011121132110-3111011231213133-2313013013030333"></a>

Type: `"single"`. Computed.

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0332322303120011-1223102300331313-0111230323101322-3112313011030333-1121123022132123-0131221212223023-2220302020121222-2112213311031131"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.client_selector`

<a id="canonical-2223022131233120-1023113210212333-0002310200010310-1030230112003130-2010000000220222-3213011021113011-3210031123231321-1302230020020233"></a>

#### `policy_based_challenge.rule_list.rules.spec.client_selector.expressions` property

Type: `["list", "string"]`. Computed.

Expressions contains the Kubernetes style label expression for selections.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-0210300122220212-1033203223102220-1333302020031120-3310022303300010-3023022233223333-2310311011132310-0312022113220300-3331113000211000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.cookie_matchers` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers

<a id="canonical-0330220220022312-2113132102033122-2331130300201321-0300222122130130-0003031012231323-1300320231332032-3133302033001032-2013313121203232"></a>

Type: `"list"`. Computed.

A list of predicates for all cookies that need to be matched. The criteria for matching each cookie
is described in individual instances of CookieMatcherType. The actual cookie values are extracted
from the request API as a list of strings for each cookie name. Note that all specified cookie
matcher predicates must evaluate to true.

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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

<a id="canonical-3020133021002310-3031222333111321-2030203201310132-3102220102020110-3210021012310320-3102132222132033-1210223203003231-1021203123131022"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.cookie_matchers`

- [check_not_present](data-sources--http_loadbalancer--reference--group-022.md#canonical-0031222222021030-1123230322221101-2312222211332320-2313313311300321-0111310300121311-0022223222230002-3211121110101033-1102303030211321): complete subsection reference.

- [check_present](data-sources--http_loadbalancer--reference--group-022.md#canonical-0001102211232200-3033001330012222-2322132003313031-0323333322311333-0122022010012011-3300102132131101-1213222013232013-0032112232313211): complete subsection reference.

<a id="canonical-1311312002320120-0322333130000233-3231301101210100-0313023100330230-2310100223112133-3213232312110012-0201021211220222-0232332113103010"></a>

<a id="canonical-2332023333023232-2110322122130220-3232020131322321-2103002210303112-0020311012312303-1301331002100112-0110101233031321-1001202102220323"></a>

#### `policy_based_challenge.rule_list.rules.spec.cookie_matchers.invert_matcher` property

Type: `"bool"`. Computed.

Invert Matcher. Invert Match of the expression defined.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [item](data-sources--http_loadbalancer--reference--group-022.md#canonical-3202023032200013-2112000112203020-3322321231230211-2221300323112312-0013131000131023-0033130310220021-1002032030203000-0000000320001032): complete subsection reference.

<a id="canonical-0012322202002230-1222120203333030-1331213032011032-3112332030013321-1320231323110212-0301322013311311-2321023123012333-1300231130313132"></a>

<a id="canonical-0013202321101121-0203102110210203-2132222211333200-1213302232333111-2110231010021002-2201210101001000-0313310211123201-2010110030102230"></a>

#### `policy_based_challenge.rule_list.rules.spec.cookie_matchers.name` property

Type: `"string"`. Computed.

Cookie Name. A case-sensitive cookie name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-0031222222021030-1123230322221101-2312222211332320-2313313311300321-0111310300121311-0022223222230002-3211121110101033-1102303030211321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](data-sources--http_loadbalancer--reference--group-022.md#canonical-0210300122220212-1033203223102220-1333302020031120-3310022303300010-3023022233223333-2310311011132310-0312022113220300-3331113000211000)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present

<a id="canonical-0002322133132002-0221313101213013-2212122203310001-3213133121313322-0121022322220301-2133103110203131-3130123000210023-1022001302200130"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

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

<a id="canonical-0001102211232200-3033001330012222-2322132003313031-0323333322311333-0122022010012011-3300102132131101-1213222013232013-0032112232313211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](data-sources--http_loadbalancer--reference--group-022.md#canonical-0210300122220212-1033203223102220-1333302020031120-3310022303300010-3023022233223333-2310311011132310-0312022113220300-3331113000211000)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present

<a id="canonical-2010030212220022-2033103212331323-3231203213032133-0110202010331300-2330201311033121-2001030133011311-1222011310300030-0233022312322120"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

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

<a id="canonical-3202023032200013-2112000112203020-3322321231230211-2221300323112312-0013131000131023-0033130310220021-1002032030203000-0000000320001032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.cookie_matchers.item` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](data-sources--http_loadbalancer--reference--group-022.md#canonical-0210300122220212-1033203223102220-1333302020031120-3310022303300010-3023022233223333-2310311011132310-0312022113220300-3331113000211000)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers.item

<a id="canonical-3233123310212100-2013012010023022-3112121231332231-0213331010110013-3133311230303120-3011121033211132-1031213120010201-1311132222230202"></a>

Type: `"single"`. Computed.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3011231310301322-3022200102121321-1323100222111011-0032303123113100-2001302310303112-2300013302132110-1313230331231101-2010212320120000"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.cookie_matchers.item`

<a id="canonical-2023130201232133-2333122012301232-2223131333100213-2201130130023212-0002012013310200-1332303111022302-1112113332123131-2000310331331010"></a>

#### `policy_based_challenge.rule_list.rules.spec.cookie_matchers.item.exact_values` property

Type: `["list", "string"]`. Computed.

A list of exact values to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2102100023332302-1032201012310102-1102333020210133-0112030021322001-2323003332321331-0100323312331202-2202022301200020-3002302203011330"></a>

<a id="canonical-3112121000223303-2001311013012032-2100231300330131-3200020103101312-1113320221010011-0101321111100200-1231000310313333-3030231102230013"></a>

#### `policy_based_challenge.rule_list.rules.spec.cookie_matchers.item.regex_values` property

Type: `["list", "string"]`. Computed.

A list of regular expressions to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1031332203333332-2133100332103003-1002212300200200-3002023021023033-2303301323101003-0212132031310210-3120123100321033-2222023113001310"></a>

<a id="canonical-2213312311200032-3102121113330203-2122021002221111-0003230011100033-1122033133320331-1323003202333301-1132013131331120-2321213011302132"></a>

#### `policy_based_challenge.rule_list.rules.spec.cookie_matchers.item.transformers` property

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Additional upstream details:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3020133121101200-2222013212131112-3111032310221132-2213303201101321-1300132123112201-3131333230313112-1220001302231030-2202110030031102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.disable_challenge` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.disable_challenge

<a id="canonical-0110103111021102-0323322010203133-2132313210011330-2321020002111130-1032201130303311-1002302322210030-2323010220013020-0012031033232200"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable challenge.

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

<a id="canonical-3311122103211312-0002003200333203-1210023030112201-2031031301112330-3220213200303010-1323003002301022-1230010220321100-3120321133102132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.domain_matcher` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.domain_matcher

<a id="canonical-0022003301132120-2012330310002231-2003200321021121-0201323002001010-3303201301200300-3332022133331000-0201032010212100-3032311131202203"></a>

Type: `"single"`. Computed.

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0113330120023001-3302101002132233-1333010110333303-3111211203220113-0020023301122202-0001301220131123-0122133123020223-1301231012011000"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.domain_matcher`

<a id="canonical-3323313332033222-1132112031222230-3013120021313123-1100231331031322-1332222121113013-2032330110013203-3211323000202112-2111110303212302"></a>

#### `policy_based_challenge.rule_list.rules.spec.domain_matcher.exact_values` property

Type: `["list", "string"]`. Computed.

A list of exact values to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2022021002231003-3233011220312233-0232022301113232-3133120131133312-3332332010002120-3310000222111202-0333131011213302-3220010121300322"></a>

<a id="canonical-3213031223102201-1332020020230213-1202121030302113-3131003002120230-0213231233001020-3113300213111303-0033222231221231-3220213310301022"></a>

#### `policy_based_challenge.rule_list.rules.spec.domain_matcher.regex_values` property

Type: `["list", "string"]`. Computed.

A list of regular expressions to match the input against.

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2322120012012001-0300113331321331-0300333021300022-2012123332131133-3100110230212230-0001321233021031-1103301000033302-2331302011012102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge` properties

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [policy_based_challenge](data-sources--http_loadbalancer--reference--group-022.md#canonical-1023133103012222-1300011200232022-1030230202310131-0213212001121312-1033331230101001-1003330132311111-1021320231332331-2113303231031310)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--reference--group-022.md#canonical-3010011021321102-0333202331311212-0021003010321123-1313203103120123-0131220132222100-3110010320233302-2212033110100322-2030230110231212)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--reference--group-022.md#canonical-2120030031313333-0313330033322030-0030333201322122-0332313320330200-3201101222001200-3311033010030022-2201020323310123-1000302221011003)
- [policy_based_challenge.rule_list.rules.spec](data-sources--http_loadbalancer--reference--group-022.md#canonical-1033032311323213-2221032202001110-3210122211031203-3121302020200110-2133320132321022-1133011300020100-3101332212212131-2112000012001020)
- policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge

<a id="canonical-1200021313030121-0123022221030023-2231322200300223-1033113202131013-3111100201002213-3012110212103020-1313132130011032-1003003223003230"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable captcha challenge.

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
