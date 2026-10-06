---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-3312020000010120-2203332210322233-3311220013302033-2322323332031303-1011331103131122-0012000131132203-0012322122202132-3123232001012113"></a>

## Direct properties for `other_settings.logging_options.client_log_options`

<a id="canonical-0131300120111001-1232112003111111-1103112220032012-0103021112133301-0313033233232222-3201122200120222-0013111223331310-2022312132020001"></a>

### `other_settings.logging_options.client_log_options.header_list` property

Type: `["list", "string"]`. Computed.

Headers. List of headers.

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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1310022203322020-1113212000111233-1003332001203230-3101200021031032-1320202113312011-1202222032211131-2200332101002333-0131103023312300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `other_settings.logging_options.origin_log_options` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [other_settings](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-0331123030101331-3110222312202112-3321300002131202-0110203121023102-3022013301122310-3122220023301103-1113321203000113-2311022130102310)
- [other_settings.logging_options](data-sources--cdn_loadbalancer--reference--group-012.md#canonical-2223103101102133-1321201122230001-0323100123120022-2010212213323231-0212002021131321-2330330122301331-0201233222123130-0011222322220210)
- other_settings.logging_options.origin_log_options

<a id="canonical-0303010130113303-3131023320201211-3230022132013033-0120110203112022-0300012002332311-2133222011002302-1333121010121201-1201212000010222"></a>

Type: `"single"`. Computed.

Configuration parameter for origin log options.

Additional upstream details:

List of headers to Log.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2113122313312110-1311311021012010-1123332121332213-0121210012323133-2313210012323213-2311111220333112-1322101201233331-2022103032112203"></a>

### Direct properties for `other_settings.logging_options.origin_log_options`

<a id="canonical-0102133032102322-0110210001023110-1220122322013032-3023310220112223-1003132021220021-3300223032303112-3223223121323132-1222032021022100"></a>

#### `other_settings.logging_options.origin_log_options.header_list` property

Type: `["list", "string"]`. Computed.

Headers. List of headers.

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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- policy_based_challenge

<a id="canonical-3320322221113222-1301022123211111-0332333001003130-3223221331301230-0032202300300213-3012213113233021-3011001102203211-3312103133001231"></a>

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

<a id="canonical-1313323303131330-0231113230323022-1332130320001113-3313320310332010-2110232220101030-2010022201000201-3323222033200113-1230320002300210"></a>

### Direct properties for `policy_based_challenge`

- [always_enable_captcha_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2030100001113330-2112130200112202-0300113310000313-0011033200330002-2001203120111112-0210111202310120-1012011312101001-1132200313323033): complete subsection reference.

- [always_enable_js_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0202231002132331-3211223301131113-3203232031001211-0123222132330301-3000202112102202-2000023102110011-2133210333203322-0130022301002031): complete subsection reference.

- [captcha_challenge_parameters](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1233022200332332-0330313332021131-2013210330331111-1131122023213203-2012210033333002-3020133112022330-3100110012233002-2333210233002322): complete subsection reference.

- [default_captcha_challenge_parameters](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0331321120033113-3101221012002312-3022320202232320-2310011132202011-2111232120123322-3010030202321212-2302113320000000-1203031132231320): complete subsection reference.

- [default_js_challenge_parameters](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-3020030331320002-2323312103120222-0233230333010101-0021201002220312-1100120123133312-1223333330123232-0330210121202121-2301300222123300): complete subsection reference.

- [default_mitigation_settings](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0013310103300131-0320102103313301-3110201221302003-0201101332013301-0110221332202122-0200223010132002-0110231220302001-0212121033102102): complete subsection reference.

- [default_temporary_blocking_parameters](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2000332003010103-1332232202222102-3023131022123002-3112110001311302-3301003202012101-0233023131213231-2121322301033201-2313020222321003): complete subsection reference.

- [js_challenge_parameters](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-3312332033330310-2102020113221113-3101100321233023-1112230300001110-2201122230122202-3102111101333130-1332220233232332-2003002330003022): complete subsection reference.

- [malicious_user_mitigation](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1330002230033301-1101210211330033-3222123112310332-2120020120220110-1331221313302032-0233330300030313-2222232003232031-2302002013120121): complete subsection reference.

- [no_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2321003023230011-0001310133230312-1321311122223130-3311300232130110-1231301132210323-0102103110013003-3132232311112002-2223233200201122): complete subsection reference.

- [rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032): complete subsection reference.

- [temporary_user_blocking](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1003222301031030-1210100221301230-3231203311030021-1323133211333010-2213032112313020-0122121033322021-0312113011113021-0102101331131010): complete subsection reference.

<a id="canonical-2030100001113330-2112130200112202-0300113310000313-0011033200330002-2001203120111112-0210111202310120-1012011312101001-1132200313323033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.always_enable_captcha_challenge` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- policy_based_challenge.always_enable_captcha_challenge

<a id="canonical-1310130023331123-3033331132230202-2033332020012300-2200033110222130-0103021032223101-3013212001211033-1232201312313002-0003301032121002"></a>

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

<a id="canonical-0202231002132331-3211223301131113-3203232031001211-0123222132330301-3000202112102202-2000023102110011-2133210333203322-0130022301002031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.always_enable_js_challenge` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- policy_based_challenge.always_enable_js_challenge

<a id="canonical-0303203022313111-3321010021001103-3121023030032011-0301021113202220-2012003313313031-0203112113123010-1311023312022132-3233012013031030"></a>

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

<a id="canonical-1233022200332332-0330313332021131-2013210330331111-1131122023213203-2012210033333002-3020133112022330-3100110012233002-2333210233002322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.captcha_challenge_parameters` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- policy_based_challenge.captcha_challenge_parameters

<a id="canonical-1202102332031323-1330320321333001-0031120030210322-3301301222213322-3211233302331113-1300322100310320-1301200233133012-0111301013123211"></a>

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

<a id="canonical-1200202211000112-2011330202212030-2030011012322022-0011001011122102-2031231211302132-0131120203303111-2110110200233200-2203102032302312"></a>

### Direct properties for `policy_based_challenge.captcha_challenge_parameters`

<a id="canonical-3320311233023302-3331230210313032-0213102132023302-2003132000213032-3001333031101111-2020132022230300-0111213221220203-0000023022312030"></a>

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

<a id="canonical-0212021211312100-1100023311213113-1221131021022203-3303313200113323-1232001111322302-3101311010000111-1303133132121123-2001300023100112"></a>

<a id="canonical-0230120222332101-2200031301121121-3302002112001312-0202031323333111-1113210120333033-1000031303320200-3010202302032313-3123203031223321"></a>

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

<a id="canonical-0331321120033113-3101221012002312-3022320202232320-2310011132202011-2111232120123322-3010030202321212-2302113320000000-1203031132231320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.default_captcha_challenge_parameters` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- policy_based_challenge.default_captcha_challenge_parameters

<a id="canonical-3301003200022032-2032000102302220-2101112003311033-3303123001031213-0030330213003310-2210132213213301-0322121231131331-0033012033120132"></a>

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

<a id="canonical-3020030331320002-2323312103120222-0233230333010101-0021201002220312-1100120123133312-1223333330123232-0330210121202121-2301300222123300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.default_js_challenge_parameters` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- policy_based_challenge.default_js_challenge_parameters

<a id="canonical-3301310220233332-0021123312313000-2110201301010023-0211321120321130-2122203301323000-0012333000301230-0001000132101010-3110323010011010"></a>

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

<a id="canonical-0013310103300131-0320102103313301-3110201221302003-0201101332013301-0110221332202122-0200223010132002-0110231220302001-0212121033102102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.default_mitigation_settings` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- policy_based_challenge.default_mitigation_settings

<a id="canonical-0301320321221233-2133022123012021-0201133310221200-0001021310002132-2211002101110103-1303032313213203-0110333003003220-1210020022003000"></a>

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

<a id="canonical-2000332003010103-1332232202222102-3023131022123002-3112110001311302-3301003202012101-0233023131213231-2121322301033201-2313020222321003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.default_temporary_blocking_parameters` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- policy_based_challenge.default_temporary_blocking_parameters

<a id="canonical-2032010310013323-1323020010330113-3301100320121200-1023002032023332-1130310323311101-0312132211312103-1333323101110003-1310022002321220"></a>

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

<a id="canonical-3312332033330310-2102020113221113-3101100321233023-1112230300001110-2201122230122202-3102111101333130-1332220233232332-2003002330003022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.js_challenge_parameters` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- policy_based_challenge.js_challenge_parameters

<a id="canonical-3121021032220013-0130132120022221-2303303022110331-2203112213000131-2100303321303211-2323312000120330-3302100102132311-3313120120313033"></a>

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

<a id="canonical-2233322102330101-0002102120233211-1321000002232313-0211123031010112-3330030223331320-2321211003010213-1220311211120233-1303030121103203"></a>

### Direct properties for `policy_based_challenge.js_challenge_parameters`

<a id="canonical-2120001301233212-0122030013212213-2101031123131333-1313000113000033-3233303223203231-3322301123303021-3101131211100131-2220010301113112"></a>

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

<a id="canonical-1123010333130110-0101300133323130-1032201011011112-2200203212103010-0333032020203332-0012300203011330-1000022130133201-1122012203210323"></a>

<a id="canonical-3213331111102220-2233031012222111-0030012130332130-1111110300203103-0013022011223220-3323320231322333-3013213011311023-3132020311131021"></a>

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

<a id="canonical-1331012310130122-3302221222321310-2122202311001111-2222011023020211-0200233031103320-3022002130131213-1211002030131330-2302101131122032"></a>

<a id="canonical-1221212201021200-1120332332030112-3000331123031110-1120320321003112-2320210101033302-0110320201321021-1212311010101320-0123211210000113"></a>

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

<a id="canonical-1330002230033301-1101210211330033-3222123112310332-2120020120220110-1331221313302032-0233330300030313-2222232003232031-2302002013120121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.malicious_user_mitigation` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- policy_based_challenge.malicious_user_mitigation

<a id="canonical-3023133122132021-1110222220201021-1021200213121201-1221020133203031-2221120103120231-0202012302030201-2220101100200212-1222013303132232"></a>

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

<a id="canonical-0222030000113123-2231322311313110-2022222303310211-0203123031322303-2203112232021232-2131232113133203-1230001022100001-3231230010303130"></a>

### Direct properties for `policy_based_challenge.malicious_user_mitigation`

<a id="canonical-0333232023202303-0023223002102012-1333121200010031-2233200322222200-0311010132012332-1023010022331330-0113101330230323-2001010301222233"></a>

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

<a id="canonical-1002301203001100-2310332101303203-2201021230211000-2231233123212300-2202211130233230-2233100231321102-3332221130011220-3010022031031003"></a>

<a id="canonical-2212001100310010-0010300320131023-0103330020223132-1232301301303333-0221311010010133-0302121231003002-3330210301320233-2333300011313302"></a>

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

<a id="canonical-0010310210103133-2200213233011303-2222230001033013-0321122323310230-2202100111330003-2321021210323331-1223312200122013-1210332022210011"></a>

<a id="canonical-2320012013002303-3012113220222001-2020313303330020-1131002310312232-3310030012231221-3222100210032001-0210022032220122-3310321010333321"></a>

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

<a id="canonical-2321003023230011-0001310133230312-1321311122223130-3311300232130110-1231301132210323-0102103110013003-3132232311112002-2223233200201122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.no_challenge` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- policy_based_challenge.no_challenge

<a id="canonical-2322021002120111-2300110033002302-2213223302123232-3030202321013200-0002210130300231-2330202002130220-0210131023023112-0322020021232013"></a>

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

<a id="canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- policy_based_challenge.rule_list

<a id="canonical-3223023233302113-2331321300311011-3030021203032020-2300022301310233-0321301120310023-1231132010003313-1333200131311323-2102222132012310"></a>

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

<a id="canonical-0322112120201113-0020023102232233-2000303122003332-0333023130220322-2003011331020010-1020030120101133-0031230203320320-0233202123000223"></a>

### Direct properties for `policy_based_challenge.rule_list`

- [rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302): complete subsection reference.

<a id="canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- policy_based_challenge.rule_list.rules

<a id="canonical-2032201021203200-2112232030123112-0321223303103001-1223332133322120-0111331112221013-2213030123122132-3232130031322130-2123210222232202"></a>

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

<a id="canonical-1230322333322232-0331123123200203-0003101321000322-3012213210323313-3001310231131220-2320022222131302-3331202123313120-0220323201230233"></a>

### Direct properties for `policy_based_challenge.rule_list.rules`

- [metadata](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2213332232012120-0000232121012231-1202321003100000-2023033121213332-1312100012202001-3200303132103321-3203131011202132-0023232002221101): complete subsection reference.

- [spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230): complete subsection reference.

<a id="canonical-2213332232012120-0000232121012231-1202321003100000-2023033121213332-1312100012202001-3200303132103321-3203131011202132-0023232002221101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.metadata` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- policy_based_challenge.rule_list.rules.metadata

<a id="canonical-1023020111301133-1000120113032133-0023001330320023-3000130132123213-3220112312032021-0200012212013222-3010201001221130-3333322001212100"></a>

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

<a id="canonical-1101312313333311-2103321320330201-2223102201103101-1231323033012030-1120321002302113-2210132033221300-1103122022110222-0112222200033332"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.metadata`

<a id="canonical-3301103031112033-2231100111121212-2122313033131122-1112223220011011-2230113313012022-2123300121022300-0311031112121100-1322220233132003"></a>

#### `policy_based_challenge.rule_list.rules.metadata.description_spec` property

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-0302210100211301-2001000223102232-2320300133211233-3112300023302310-3003222130312231-1331123031113210-1323010032030202-0213213013302003"></a>

<a id="canonical-3223100233002231-3320100331131210-3202232100011010-3011130213201123-1333300122331230-3310213111303122-1203101200322033-2120023233122221"></a>

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

<a id="canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- policy_based_challenge.rule_list.rules.spec

<a id="canonical-1022213103032223-0211022310130111-0122230103313203-0002321332030333-3000332002010300-2221301133113323-2101103223330032-1130131310222120"></a>

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

<a id="canonical-2123013133221212-1103020002223223-2210023212312313-1133222221222323-3322302023210202-1330232213202100-0031101010303313-1021102213130231"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec`

- [any_asn](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1211213323333331-3123231310232023-0102332112110100-2011112333113110-1210213012301101-0233123203311333-0212000332013310-0211301122011232): complete subsection reference.

- [any_client](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-3001032210121312-2232021320102331-2221130312230313-0322032201301033-0313120033310020-1120133211031201-2201030012031023-2021320321112111): complete subsection reference.

- [any_ip](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2132021231130112-2123321303233202-1003030103123331-3220221301231301-0320121202232322-3113302121103110-3122011123303333-3113322220020022): complete subsection reference.

- [arg_matchers](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0333011032020222-1031202212303322-2302201223131012-1121020102022323-2301011330121222-1302011333231300-3312333012333012-1133113110333323): complete subsection reference.

- [asn_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-3030233320003122-1011021332000221-3203021022013102-2130330003322122-1203302321002323-3211210201111313-0012132222031031-2310010203120022): complete subsection reference.

- [asn_matcher](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-3122310212320233-0031302121130202-1012120233130012-0322303310331332-3110033221113022-2311012020202221-2123000203032312-1012312101202212): complete subsection reference.

- [body_matcher](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1213311021203302-2332101131201022-1122221301221222-2300200300120230-2022132311130021-0100323233103233-1000133330002011-1123331220122111): complete subsection reference.

- [client_selector](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-3102120111001203-1112131330132320-2121010233221012-2320103020332010-2113223021322212-2311332300120013-1311000030211222-2132213323012211): complete subsection reference.

- [cookie_matchers](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-3331311132121000-2302123312033003-2201211023002032-1200012322011112-2032220231112221-0300003102013101-0133332332130220-3120231102333000): complete subsection reference.

- [disable_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-3333031012000322-3022131233200332-1222200332232033-2210120021231332-1023232322313002-0133223103121211-0003110221123323-0313331331022230): complete subsection reference.

- [domain_matcher](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1213313313333103-1013220331032213-2322322203300120-0232232023222301-1110333033030011-2120302233312122-3102011002331023-2102012213103220): complete subsection reference.

- [enable_captcha_challenge](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-3011132030011201-2211311232213300-0312302301232121-3330311310211332-1123332021213002-1232230220130032-1231303030102230-1132313121201323): complete subsection reference.

- [enable_javascript_challenge](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2310032320333030-1003211003002112-2220030321201321-2312212330103131-3102112121111313-0103301022112213-2332123333231022-0233120222133023): complete subsection reference.

<a id="canonical-3322130120002003-0302020013310232-2201122102313021-3022030323322023-0013132200113301-1010212321333223-1101220220021010-2010131211013323"></a>

<a id="canonical-3102301011323313-1103230221121323-2122022101331030-0220000123212023-0030203012020322-3220221333002321-2202220110001031-1130230210112302"></a>

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

- [headers](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2133220031201030-3320302321033311-1323020211311211-0030320331200130-2221030100003120-1110102330021130-2013211231232321-3022012322202123): complete subsection reference.

- [http_method](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0203122003302010-0311301031311231-1113123010123300-1231200102112212-0111013121233330-0102113203321103-2323133133131301-0110321113200323): complete subsection reference.

- [ip_matcher](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0201030211031131-1111310113132001-1132122322201233-0010012333222220-2001233311231202-3210131212113212-2302321102233223-0202323121313221): complete subsection reference.

- [ip_prefix_list](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0212033022130233-1300022103102203-3221233211032110-0022200332332213-3121222133232013-0122231011020300-3311100313212031-0300232013311002): complete subsection reference.

- [path](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-1323200003020201-0013112330130332-0230323022322321-2231311020200133-1312320220310220-1302221121320212-2103002221230003-0302010321213122): complete subsection reference.

- [query_params](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-0322033212130133-1113133131322212-0103013001021121-2103031200101123-3011211013203013-2000223231222111-3221121300020221-2012013210300023): complete subsection reference.

- [tls_fingerprint_matcher](data-sources--cdn_loadbalancer--reference--group-014.md#canonical-2311200300102223-2200011111023033-1303211002130001-2221330223002003-2133120032110001-1030331310222102-1303231310132020-0311310011103102): complete subsection reference.

<a id="canonical-1211213323333331-3123231310232023-0102332112110100-2011112333113110-1210213012301101-0233123203311333-0212000332013310-0211301122011232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.any_asn` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- policy_based_challenge.rule_list.rules.spec.any_asn

<a id="canonical-2022321201311001-2022123020113302-3222131231102021-1311102002333213-1123131210201310-2002202123113222-2310103212232222-3200032111122113"></a>

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

<a id="canonical-3001032210121312-2232021320102331-2221130312230313-0322032201301033-0313120033310020-1120133211031201-2201030012031023-2021320321112111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.any_client` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- policy_based_challenge.rule_list.rules.spec.any_client

<a id="canonical-2113323021232020-3321112213211132-1033001002002323-0122101202032233-1333011303320102-3113112231101102-3331313321200021-2100302020232202"></a>

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

<a id="canonical-2132021231130112-2123321303233202-1003030103123331-3220221301231301-0320121202232322-3113302121103110-3122011123303333-3113322220020022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.any_ip` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- policy_based_challenge.rule_list.rules.spec.any_ip

<a id="canonical-0101310011002212-2011232231303331-2233122211023322-3303222221312010-2112211320102003-2021103320312033-0212212020003212-1230213221000303"></a>

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

<a id="canonical-0333011032020222-1031202212303322-2302201223131012-1121020102022323-2301011330121222-1302011333231300-3312333012333012-1133113110333323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.arg_matchers` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- policy_based_challenge.rule_list.rules.spec.arg_matchers

<a id="canonical-1322122122110222-3110303013132013-1303330030102120-3013132233113310-1323130213222200-0230320203320222-2333303000002312-2123130000333031"></a>

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

<a id="canonical-3221113321110313-1200203021321201-0000330031210231-1113122002301000-0001033332230312-0010121023023132-0131320333333012-1011110200000012"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.arg_matchers`

- [check_not_present](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0011210200110321-3133122231213301-1222102311231331-1211320002312231-0001120223300232-0013112021212230-1102312123120102-1333022202101021): complete subsection reference.

- [check_present](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1132333220020301-0313333333301101-2303331233003010-1203001113201310-0223011021030001-0030121210122100-3102120300121022-1011013211233201): complete subsection reference.

<a id="canonical-2202113131132011-0322010331121202-1310230132013023-2110003320323213-3323302033321223-3220021310313120-1021003122303001-2202210322000230"></a>

<a id="canonical-0203033112012222-3003333303210021-1233222220303111-3320111301012220-2132211011313112-3122213023300011-2230103220003300-1032323021023012"></a>

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

- [item](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2222320022131003-0333230101200030-1001003332011102-2033322112101302-0231021300332122-3031123223220122-0331130110231333-3031211122333100): complete subsection reference.

<a id="canonical-3231120300203132-2330230302012302-3111301312033200-2332313322001120-1112322232123232-1113133023133213-2303122032012302-1332320001200311"></a>

<a id="canonical-1220313122331220-0301130332300001-3002012322021332-2132032030033020-3201013303220112-0220023330033202-1231002102231200-1331310100221223"></a>

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

<a id="canonical-0011210200110321-3133122231213301-1222102311231331-1211320002312231-0001120223300232-0013112021212230-1102312123120102-1333022202101021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0333011032020222-1031202212303322-2302201223131012-1121020102022323-2301011330121222-1302011333231300-3312333012333012-1133113110333323)
- policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present

<a id="canonical-0122102322213210-0302210001013330-3232303020001213-3223131200122111-2310012020113021-0230230001121003-3030320033020330-0111113300333202"></a>

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

<a id="canonical-1132333220020301-0313333333301101-2303331233003010-1203001113201310-0223011021030001-0030121210122100-3102120300121022-1011013211233201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0333011032020222-1031202212303322-2302201223131012-1121020102022323-2301011330121222-1302011333231300-3312333012333012-1133113110333323)
- policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present

<a id="canonical-0321323202123303-1030332330031132-3332132333031201-0303220001303213-2031110230103213-0310213231212132-1110310322131211-3323121132221322"></a>

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

<a id="canonical-2222320022131003-0333230101200030-1001003332011102-2033322112101302-0231021300332122-3031123223220122-0331130110231333-3031211122333100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.arg_matchers.item` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0333011032020222-1031202212303322-2302201223131012-1121020102022323-2301011330121222-1302011333231300-3312333012333012-1133113110333323)
- policy_based_challenge.rule_list.rules.spec.arg_matchers.item

<a id="canonical-3203033132011112-2310122200022302-2123300020213331-2132110322002322-2201110012010122-1130231210110233-3111013312111123-0200132010011213"></a>

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

<a id="canonical-0133313011302231-1102332313002012-3322031120120032-3333330300330233-0113100302223221-1013110031023030-3201020120301332-3320223133030113"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.arg_matchers.item`

<a id="canonical-2010313320233220-0103123321312123-1321231203100301-3103021012000311-2231223101222212-2111101221103333-0012002120021021-3031232023312203"></a>

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

<a id="canonical-2323231233102210-2122031333121332-0121121031201302-3310230022032001-0001103111311120-1230331303320230-2231220102003020-1220112033230221"></a>

<a id="canonical-1112222311100131-1313333221013030-1311302003223210-2211103021032221-0111003313313331-2122200331110101-3102010230020322-3222312013132303"></a>

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

<a id="canonical-1021203023201130-0211002330022231-0133213200110100-3230020323121130-2001220310122310-2331312211301300-1132023322030311-3211210101321232"></a>

<a id="canonical-1021203030113032-0302120303113003-2033033000023221-0332001013320330-2332222022020323-2233130200220002-3331220012231020-3330020103101012"></a>

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

<a id="canonical-3030233320003122-1011021332000221-3203021022013102-2130330003322122-1203302321002323-3211210201111313-0012132222031031-2310010203120022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.asn_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- policy_based_challenge.rule_list.rules.spec.asn_list

<a id="canonical-0220233031230013-3202233312323011-3303301123322123-3133320311212313-3102312031011220-2001133123303121-2220322012323223-0211222111333223"></a>

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

<a id="canonical-1213311101212330-3100301002121303-1222301021222321-2113021103121032-0302201221003122-3220032213030033-1213102302023113-1332020211101200"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.asn_list`

<a id="canonical-2213320223023323-1213331230011311-2130313111130021-0212302121310021-1300221121113212-0311323210012123-1112200100302003-2233130222132313"></a>

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

<a id="canonical-3122310212320233-0031302121130202-1012120233130012-0322303310331332-3110033221113022-2311012020202221-2123000203032312-1012312101202212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.asn_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- policy_based_challenge.rule_list.rules.spec.asn_matcher

<a id="canonical-2310202112322232-0113300203010002-3003320332102002-0200132213230021-0120301201132133-3112232221300003-0031320013331020-3320230130023121"></a>

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

<a id="canonical-0302111122012022-3023221111223313-3231200231113221-0330332313333021-3233022000220112-0121231012022303-3123222313210100-2000120233232232"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.asn_matcher`

- [asn_sets](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122002213211130-3201030122031323-1113020310232001-3312130023013032-2202101133313130-3023013310012313-2123011332213113-1331021200200330): complete subsection reference.

<a id="canonical-2122002213211130-3201030122031323-1113020310232001-3312130023013032-2202101133313130-3023013310012313-2123011332213113-1331021200200330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [policy_based_challenge.rule_list.rules.spec.asn_matcher](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-3122310212320233-0031302121130202-1012120233130012-0322303310331332-3110033221113022-2311012020202221-2123000203032312-1012312101202212)
- policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets

<a id="canonical-3120312033001211-2231322111213320-2113221333133133-1133130022312111-3323222311310012-2223022112110232-0333331101310012-3101211132231330"></a>

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

<a id="canonical-2123002003123331-2121322121020003-2213001023303022-1311013211202210-1003003001011302-2010310210320020-0032332220200200-0323323233202213"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets`

<a id="canonical-0121330121013302-1021232302022202-0022323013201010-1233123212123012-2220212032032010-0203210220133100-3232012301322213-0323333132013231"></a>

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

<a id="canonical-1300022321333030-1012201001211033-0210131010123212-0030030103030023-0330002030200201-0301313303100110-2200023131311033-3103330102102331"></a>

<a id="canonical-2002030010001311-0300201022321323-1123100030131333-0011223323232102-1001200221110330-0123201222011330-2133210031300100-3232131032323331"></a>

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

<a id="canonical-1100223223202201-1110300020112312-1130111211332223-3002330123202132-3000303031302120-2002200020101010-0132122301213220-0232210133111113"></a>

<a id="canonical-3001003123311332-3321323221113232-1110003302221032-2202032011301103-2300312231230200-2133031333130133-0212302131301313-2232213310231330"></a>

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

<a id="canonical-0013312200122311-2333113003010212-3133100323312031-3020111232303223-0302011303213000-3332102110303111-1213112030321003-0120131031030012"></a>

<a id="canonical-3333212301101101-1132221111010331-3002331330201211-2210233210233223-0112203330232212-0322032003333321-0220222101113032-0013322331303321"></a>

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

<a id="canonical-2300232101202000-1100301323002100-3122003013112312-3123011131011030-3032000231221013-1220032001333333-1023222210020300-2231311222020022"></a>

<a id="canonical-0312023322132022-0030032110300302-3120113231223021-2303012320131330-1120120022220222-1020120333320223-3030111133212313-3223031011111020"></a>

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

<a id="canonical-1213311021203302-2332101131201022-1122221301221222-2300200300120230-2022132311130021-0100323233103233-1000133330002011-1123331220122111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.body_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- policy_based_challenge.rule_list.rules.spec.body_matcher

<a id="canonical-0021232132332231-3223332321120113-3201201200022220-0122000031122303-0133332031130012-1312202320200002-1101300212001312-1130112331202120"></a>

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

<a id="canonical-2121303001201111-2221023203220123-1333333023330023-0202110221232322-3030102001320110-0003230001313332-1331312231311122-2102112301003110"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.body_matcher`

<a id="canonical-2132032230201331-0011122332301110-1302233312203020-0101122032013022-3012123331201201-1223233223033300-1231102120220013-1301212200323012"></a>

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

<a id="canonical-2321111220332203-2013233313030233-3010131221302231-1213001213323112-2021132323121120-0022010122012230-2022120111223200-3200013112303302"></a>

<a id="canonical-3302033322331100-1230100033213220-2302231102013103-0202302222221000-1001233322323321-2210100300021031-1033032323222112-1232322302130203"></a>

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

<a id="canonical-3032020033121332-1001222222312021-3132002312313110-0301113103230102-2301110123032231-3212010020013303-2201113000313022-3033031211332110"></a>

<a id="canonical-2231212320011101-3221112032202110-2332010202331301-0002303111200312-3023310233122102-3300322330121030-0002032333033332-3022020010313213"></a>

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

<a id="canonical-3102120111001203-1112131330132320-2121010233221012-2320103020332010-2113223021322212-2311332300120013-1311000030211222-2132213323012211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.client_selector` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- policy_based_challenge.rule_list.rules.spec.client_selector

<a id="canonical-3100302003100121-2101031212313311-1311211200120000-1311230200122221-0012312132010120-1001030032031120-3112020230310203-0212302200232321"></a>

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

<a id="canonical-2221301222131313-3321300312011223-2122011031110131-2030220210201232-3132332033020223-3132313111200203-3203110000103313-1123322030300012"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.client_selector`

<a id="canonical-2131123020323300-3330102312212223-0333023321332310-1332111232113223-0101100110211322-3110023321302223-1201000131021112-3132323002212030"></a>

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

<a id="canonical-3331311132121000-2302123312033003-2201211023002032-1200012322011112-2032220231112221-0300003102013101-0133332332130220-3120231102333000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.cookie_matchers` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers

<a id="canonical-2031303023020310-2103200303200230-3002122313010321-2031022003310100-1301103113000112-2232311000303201-1122231010203230-1121210232211223"></a>

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

<a id="canonical-3301320120132212-3231002032123200-2231332230213221-3132032033201132-3202323201120230-0113111012211210-2120102232301100-1001222132231203"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.cookie_matchers`

- [check_not_present](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-3112220113030003-1023031111132130-2220033031110223-1032201132133310-1330231202211113-1022011132301320-2303211202203221-3111103110113310): complete subsection reference.

- [check_present](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2023012013122221-2232310323120110-3321321301313210-3212121020112212-3110120110022013-3223200222002302-0031001220313001-3100132033311312): complete subsection reference.

<a id="canonical-1012313202332301-2232231121213032-0212202321210203-3313232200222023-3123300022230003-0103203230331211-1313112010130122-2101120010233103"></a>

<a id="canonical-0121031201102213-2223311230212023-0133211231012201-2011001220013101-2301120113011313-1312332031301300-0032132320101322-3130200333012020"></a>

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

- [item](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2330131222212101-2030300121322113-2132131311222121-1323000101322130-2123230323131122-1010323021231031-2213021200203211-3312130120323313): complete subsection reference.

<a id="canonical-3130122012201320-2223012321213233-0132302020100322-1101010030332232-1330023102031013-0232110020201313-3313110331321103-1120212021231201"></a>

<a id="canonical-1002313331010211-0230031223222101-2021203211113122-0123301211321012-3112100203311123-1302130111323233-0100100300333012-0101302201003212"></a>

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

<a id="canonical-3112220113030003-1023031111132130-2220033031110223-1032201132133310-1330231202211113-1022011132301320-2303211202203221-3111103110113310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-3331311132121000-2302123312033003-2201211023002032-1200012322011112-2032220231112221-0300003102013101-0133332332130220-3120231102333000)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present

<a id="canonical-2021002211123001-2031320102012002-1013131100103120-3311312013031120-3100232202101030-1333011031130111-1131010211123202-2010232322132102"></a>

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

<a id="canonical-2023012013122221-2232310323120110-3321321301313210-3212121020112212-3110120110022013-3223200222002302-0031001220313001-3100132033311312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-3331311132121000-2302123312033003-2201211023002032-1200012322011112-2032220231112221-0300003102013101-0133332332130220-3120231102333000)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present

<a id="canonical-2022223301220301-2322200011203323-2301131320120223-2230210000110311-3311101322112232-3003132133102311-2222320223132211-0201022003323230"></a>

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

<a id="canonical-2330131222212101-2030300121322113-2132131311222121-1323000101322130-2123230323131122-1010323021231031-2213021200203211-3312130120323313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.cookie_matchers.item` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-3331311132121000-2302123312033003-2201211023002032-1200012322011112-2032220231112221-0300003102013101-0133332332130220-3120231102333000)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers.item

<a id="canonical-2012212000313113-2100132011023101-2322332211231330-1323113332302112-3022030312002312-0231131131203220-3210103011132223-0322321302200122"></a>

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

<a id="canonical-3230023303031132-3121102333331023-3300301322202302-1210311121033030-0221313232221100-2101331133023312-2013013300332131-1211331300211322"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.cookie_matchers.item`

<a id="canonical-2213020100230102-0023213102223031-1333021322233101-1323310000233203-2122231213220112-1001101131322130-2002000102231313-3021333133223001"></a>

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

<a id="canonical-1012113110222113-3323211030021020-3100302221330013-2100112133011102-1232312232321203-2203022113102201-0102021131222210-3301333110310200"></a>

<a id="canonical-0302223130333321-3221013002202101-0001010220131223-3021312321232231-2301013211303303-1112321000313313-1120332131333013-3021033123110103"></a>

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

<a id="canonical-0301003113210222-2312030102021031-2330121023232303-2022333110311132-2100003033011220-2223332132123122-0312000331213210-2222302300133323"></a>

<a id="canonical-0122100102300212-1332202000103102-0322033322231220-0010201111100122-1030202013203101-0201303021021010-3200033121013332-1323332200031112"></a>

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

<a id="canonical-3333031012000322-3022131233200332-1222200332232033-2210120021231332-1023232322313002-0133223103121211-0003110221123323-0313331331022230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.disable_challenge` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- policy_based_challenge.rule_list.rules.spec.disable_challenge

<a id="canonical-2102211211013310-3131003323203000-3120032221011100-0103022333112211-0122031030333201-2233301101120131-0121030233110221-0013030011221031"></a>

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

<a id="canonical-1213313313333103-1013220331032213-2322322203300120-0232232023222301-1110333033030011-2120302233312122-3102011002331023-2102012213103220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `policy_based_challenge.rule_list.rules.spec.domain_matcher` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-0131031313232111-3220331123332012-0312332120231130-2123203201020201-3310121233132331-3122031122130322-2333201200113231-1123033220323222)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-1131232031102311-1311132230333002-2103010130221322-3211022033121010-2010132101303220-1213122123033021-1002230232002302-2212020120301223)
- [policy_based_challenge](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-1003021032130300-2322021331122233-2322332013300023-2032000221100232-0311111113021131-1303100011013231-0211010200312023-3120133233232032)
- [policy_based_challenge.rule_list](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210302300122322-3311000103111301-0000210222112330-3200131110220012-0220300113301323-2132222023303312-0130120203310230-0310103123301032)
- [policy_based_challenge.rule_list.rules](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-2122332200010113-2012230320001223-0302332030031011-2331201000111013-3213033011323222-2332001103111111-1211131122212231-2000003002130302)
- [policy_based_challenge.rule_list.rules.spec](data-sources--cdn_loadbalancer--reference--group-013.md#canonical-0210130300232123-3302100223103010-2112120312100031-3110232333220300-0200013000120103-0301102120002322-3222002311001313-1330100300023230)
- policy_based_challenge.rule_list.rules.spec.domain_matcher

<a id="canonical-3202031332022010-0212311330100022-2010330303211311-3011300030103012-3223011313321221-0202223020100212-0231332003322112-0010012203311110"></a>

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

<a id="canonical-0030230023320120-3211203331332111-2221330030320201-1303112020110113-0032210221212122-2003112013021012-1212120123002303-1013022023322120"></a>

### Direct properties for `policy_based_challenge.rule_list.rules.spec.domain_matcher`

<a id="canonical-1003311002013300-0103110201233102-2101002323011021-2020311130132101-3000033030333110-3101211033223313-0331112003002022-3020330110013210"></a>

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

<a id="canonical-3031310332133320-0210231322020112-0233023220313220-0211311110320123-1323023222112302-2023131131303110-2220321231301123-2210032011222222"></a>

<a id="canonical-2323131110032230-2302333330031331-0300030123213223-2333012012013200-2102223122121210-3202223221102011-1111212203100333-0013322003001030"></a>

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
